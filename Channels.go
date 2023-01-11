import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)
// https://www.youtube.com/watch?v=VkGQFFl66X4&list=PLKOhSssLmvQQPFv-QqbRZ10rn3UdBLvVC&index=3

func DoWork() int{
  time.sleep(time.Second)
  return rand.Int(100)
}

func main(){
  dataChan := make(chan int) //declare channel
  
  go func(){ // declare new go routine
    wg := sync.WaitGroup{}
    for i:=0; i < 1000; i++{
      wg.Add(1) // wait for one more go routine
      result := DoWork()
      dataChan <- result //insert integer into the channel
    }
    wg.Wait() // block until all go routines are finished executing
    close(dataChan)
  }()
  
  for n:= range dataChan {
    fmt.Printf("n=%d", n) 
  }
  
}
