package main

import ( "fmt"
          "sync"
          "time"
      )



type Chops struct{
mutx sync.Mutex
chopstickNumber int
}

type Philosopher struct{
leftChopstick *Chops
rightChopStick *Chops
philosopherNumber int
}


var c1 = make(chan int,1)
var c2 = make(chan int,1)
func getPermissionFromHost(p *Philosopher, wg2 *sync.WaitGroup){
    select {
            case   <-c1:
                        fmt.Println("received on c1 by philosopher", p.philosopherNumber)
                        wg2.Done()
            case  <-c2:
                         fmt.Println("received on c2 by philosopher", p.philosopherNumber)
                         wg2.Done()
    }
}


func (p Philosopher) philosophersEat(wg1 *sync.WaitGroup){
var wg2 sync.WaitGroup
wg2.Add(1)
  go getPermissionFromHost(&p, &wg2)
wg2.Wait()
//fmt.Println("------------------------------- start")
 for i:=0;i<3;i++{
    (p.rightChopStick).mutx.Lock()
    (p.leftChopstick).mutx.Lock()
    fmt.Println("starting to eat :" ,p.philosopherNumber)
    time.Sleep(50 * time.Millisecond) // used for verifying that only 2 goroutines are running concurrently
    fmt.Println("finishing eating:" ,p.philosopherNumber)
    (p.leftChopstick).mutx.Unlock()
    (p.rightChopStick).mutx.Unlock()
   }
//fmt.Println("-------------------------------end")
      select{
        case c1 <- 2:
                      fmt.Println("send to c1 by philosopher", p.philosopherNumber)
        case c2 <- 2:
                      fmt.Println("send to c2 by philosopher", p.philosopherNumber)
        }
wg1.Done()
}

func main(){
var wg1 sync.WaitGroup
chopsticks := make([]*Chops,0)
c1 <- 1
c2 <- 2
for i :=0 ; i<5 ; i++{
chopsticks=append(chopsticks,&Chops{chopstickNumber:i})
}
philosophers := make([]*Philosopher,0)
for i :=0 ; i<5 ; i++{
philosophers =append(philosophers,&Philosopher{leftChopstick:chopsticks[i],rightChopStick:chopsticks[(i+1)%5],philosopherNumber:i+1})
}
for i :=0 ; i<5 ; i++{
wg1.Add(1)
go philosophers[i].philosophersEat(&wg1)
}
wg1.Wait()
}
