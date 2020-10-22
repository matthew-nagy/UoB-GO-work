package main

import (
	"container/list"
	"flag"
	"fmt"
	"math/rand"
	"time"
	"sync"
	"strconv"
)

var debug *bool

var failed int

type Key struct{
	locked bool
	readLock *sync.Mutex
}

func canLockKey(x *Key)bool{
	toret := false
	x.readLock.Lock()
	if x.locked == false{
		toret = true
		x.locked = true
	}
	x.readLock.Unlock()
	return toret
}
func unlockKey(x *Key){
	x.readLock.Lock()
	x.locked = false
	x.readLock.Unlock()
}

func isOK(from, to int, mutexInfo []Key)bool{
	if canLockKey(&mutexInfo[from]){
		if canLockKey(&mutexInfo[to]){
			return true
		}else{
			unlockKey(&mutexInfo[from])
		}
	}
	failed = failed + 1
	return false
}

// An executor is a type of a worker goroutine that handles the incoming transactions.
func executor(bank *bank, executorId int, transactionQueue chan transaction, done chan<- bool, mutexInfo []Key) {
	for {
		searching := true
		from := ""
		to := ""
		fromI := 0
		toI := 0
		t := <-transactionQueue
		for searching{

			from = bank.getAccountName(t.from)
			to = bank.getAccountName(t.to)
			
			fromI = t.from
			toI = t.to
			if fromI == toI{
				if(canLockKey(&mutexInfo[fromI])){
					searching = false
				}else{
					transactionQueue <- t
					t = <-transactionQueue
				}
			}else{
				if isOK(fromI, toI, mutexInfo){
					searching = false
				}else{
					transactionQueue <- t
					t = <-transactionQueue
				}
			}
		}

		 bank.lockAccount(t.from, strconv.Itoa(executorId))
		fmt.Println("Executor\t", executorId, "locked account", from)
		 bank.lockAccount(t.to, strconv.Itoa(executorId))
		 fmt.Println("Executor\t", executorId, "locked account", to)

		fmt.Println("Executor\t", executorId, "attempting transaction from", from, "to", to)
		e := bank.addInProgress(t, executorId) // Removing this line will break visualisations.
		bank.execute(t, executorId)

		 bank.unlockAccount(t.from, strconv.Itoa(executorId))
		 fmt.Println("Executor\t", executorId, "unlocked account", from)
		 bank.unlockAccount(t.to, strconv.Itoa(executorId))
		 fmt.Println("Executor\t", executorId, "unlocked account", to)

		unlockKey(&mutexInfo[fromI])
		if fromI != toI{
			unlockKey(&mutexInfo[toI])
		}

		bank.removeCompleted(e, executorId) // Removing this line will break visualisations.
		done <- true
	}
}

func toChar(i int) rune {
	return rune('A' + i)
}

// main creates a bank and executors that will be handling the incoming transactions.
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	debug = flag.Bool("debug", false, "generate DOT graphs of the state of the bank")
	flag.Parse()

	bankSize := 6 // Must be even for correct visualisation.
	transactions := 1000

	accounts := make([]*account, bankSize)
	for i := range accounts {
		accounts[i] = &account{name: string(toChar(i)), balance: 1000}
	}

	bank := bank{
		accounts:               accounts,
		transactionsInProgress: list.New(),
		gen:                    newGenerator(),
	}

	startSum := bank.sum()

	transactionQueue := make(chan transaction, transactions)
	expectedMoneyTransferred := 0
	for i := 0; i < transactions; i++ {
		t := bank.getTransaction()
		expectedMoneyTransferred += t.amount
		transactionQueue <- t
	}

	done := make(chan bool)

	lockInfo := make([]Key, 6)
	for i := 0; i < bankSize; i++{
		lockInfo[i] = Key{locked: false, readLock: &sync.Mutex{}}
	}


	for i := 0; i < bankSize; i++ {
		go executor(&bank, i, transactionQueue, done, lockInfo)
	}
		//go executor(&bank, 0, transactionQueue, done, lockInfo)


	for total := 0; total < transactions; total++ {
		fmt.Println("Completed transactions\t", total)
		<-done
	}

	fmt.Println()
	fmt.Println("Caught failures: ", failed)
	fmt.Println("Expected transferred", expectedMoneyTransferred)
	fmt.Println("Actual transferred", bank.moneyTransferred)
	fmt.Println("Expected sum", startSum)
	fmt.Println("Actual sum", bank.sum())
	if bank.sum() != startSum {
		panic("sum of the account balances does not much the starting sum")
	} else if len(transactionQueue) > 0 {
		panic("not all transactions have been executed")
	} else if bank.moneyTransferred != expectedMoneyTransferred {
		panic("incorrect amount of money was transferred")
	} else {
		fmt.Println("The bank works!")
	}
}
