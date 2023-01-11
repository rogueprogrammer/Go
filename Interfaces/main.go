package main

type BankAccount interface {
  GetBalance() int
  Deposit(amount int)
  Withdraw(amount int) error
}

func main(){
  wf := NewWellsFargo()
  wf.Deposit(100)
  wf.Deposit(200)
  if err := wf.Withdraw(5) ; err != nil {
    panic(err) 
  }
  
}
