import "fmt"
// https://www.youtube.com/watch?v=VkGQFFl66X4&list=PLKOhSssLmvQQPFv-QqbRZ10rn3UdBLvVC&index=3

func DoWork() int{
  time.sleep(time.Second)
  return rand.Int(100)
}

func main(){
  dataChan := make(chan int)
  
  go func(){
    for i:=0; i < 1000; i++{
      result := DoWork()
      dataChan <- result
    }
    close(dataChan)
  }()
  
  for n:= range dataChan {
    fmt.Printf("n=%d", n) 
  }
  
}
