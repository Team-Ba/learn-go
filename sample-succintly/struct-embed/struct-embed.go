package main

import "fmt"

type Discount struct {
  percent       float32
  promotionID   string
}

type ManagersSpecial struct {
  Discount
  extraOff      float32
}

func (d* Discount) double() {
  d.percent *= 2
}

func (ms ManagersSpecial) calculate(oPrice float32) float32 {
  return oPrice - (oPrice / 100 * ms.Discount.percent) - ms.extraOff
}

func main() {
  ms := ManagersSpecial {
    Discount : Discount {
                 percent       : 15.00,
                 promotionID   : "January",
               },
    extraOff : 10.00,
  }
  
  price := float32(100.00)
  
  fmt.Println("Manager's Special")
  fmt.Printf("Discount: %v", ms.Discount.percent)
  fmt.Println()
  fmt.Printf("Extra off: %.2f", ms.extraOff)
  fmt.Println()
  fmt.Printf("Price: %.2f", price)
  fmt.Println()
  fmt.Printf("Discounted price: %.2f", ms.calculate(price))
  
  fmt.Println()
  fmt.Println()
  ms.Discount.double()
  
  fmt.Println("Manager's Special")
  fmt.Printf("Discount: %v", ms.Discount.percent)
  fmt.Println()
  fmt.Printf("Extra off: %.2f", ms.extraOff)
  fmt.Println()
  fmt.Printf("Price: %.2f", price)
  fmt.Println()
  fmt.Printf("Discounted price: %.2f", ms.calculate(price))
}