package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "time"
)

var amountPayable, dropoffPrice, amount int

func main() {
  var reader = bufio.NewReader(os.Stdin)
  fmt.Println("----------------------------------")
  fmt.Print("|")
  fmt.Print("WELCOME TO CABBY BOOKING PLATFORM")
  fmt.Println("|")
  fmt.Println("----------------------------------")
  fmt.Println("Pickup and dropoff locations: ikono, ikot ekpene, uyo, essien udim, ibesikpo, eket, onna")
  fmt.Println("----------------------------------")

  //Pickup and drop-off points as locations with key as name and value as price
  var locations = map[string]int{
    "ikono":       0,
    "ikot ekpene": 300,
    "essien udim": 500,
    "uyo":         800,
    "ibesikpo":    1000,
    "eket":        1300,
    "onna":        1600,
  }

  for {
    //Get customer's pickup point
    fmt.Print("What is your pickup point?-> ")
    pickup, _ := reader.ReadString('\n')
    // convert CRLF to LF
    pickup = strings.Replace(pickup, "\n", "", -1)

    //Checking if pickup location does exist
    pickupPrice, ok := locations[pickup]

    if ok {
      fmt.Printf("%s is available for pickup\n", pickup)

      for {
        //Get Customer's drop-off point
        fmt.Print("What is your dropoff point?-> ")
        dropoff, _ := reader.ReadString('\n')
        // convert CRLF to LF
        dropoff = strings.Replace(dropoff, "\n", "", -1)

        //Check if drop off location  is available
        dropoffPrice, ok = locations[dropoff]
        if ok {
          fmt.Printf("%s is available for dropoff\n", dropoff)
          break
        } else {
          continue
        }

      }

      //Calculate cost for the delivery
      if dropoffPrice > pickupPrice {
        amountPayable = dropoffPrice - pickupPrice
        fmt.Printf("Cost of delivery is N%d\n", amountPayable)
      } else {
        amountPayable = pickupPrice - dropoffPrice
        fmt.Printf("Cost of delivery is N%d\n", amountPayable)
      }

      //Get Customer's confirmation
      fmt.Print("Confirm delivery (yes/no)?-> ")
      confirm, _ := reader.ReadString('\n')
      // convert CRLF to LF
      confirm = strings.Replace(confirm, "\n", "", -1)
      if strings.Compare("yes", confirm) == 0 {
        fmt.Println("Processing delivery.................10%")
        time.Sleep(2 * time.Second)
        fmt.Println("Processing delivery.................20%")
        time.Sleep(2 * time.Second)
        fmt.Println("Processing delivery.................40%")
        time.Sleep(2 * time.Second)
        fmt.Println("Processing delivery.................60%")
        time.Sleep(2 * time.Second)
        fmt.Println("Processing delivery.................80%")
        time.Sleep(2 * time.Second)
        fmt.Println("Delivery Successfully completed.................100%")

        for {
          tryPayment()
          count := 0
          count++
          if count == 5 {
            fmt.Println("YOU WILL BE REPORTED TO THE POLICE IMMEDIATELY")
            time.Sleep(2 * time.Second)
            fmt.Println("Sending notification to the nearest police station....90%")
            time.Sleep(2 * time.Second)
            fmt.Println("Police already on their way to your location")
          } else if amount >= amountPayable {
            break
          }
        }

        if amount > amountPayable {
          balance := amount - amountPayable
          fmt.Printf("Thank your patronage, your balance is N%d\n", balance)
        }
        //Get Customer's amount payable
        fmt.Print("We will appreciate if you give us a tip?-> ")
        var amount uint16
        fmt.Scan(&amount)

        if int(amount) <= 0 {
          fmt.Println("<<<<<<<<<<<<<<<<<<<You are a stingy fool>>>>>>>>>>>>>>>")

        } else if amount < uint16(amountPayable) && amount > 0 {
          fmt.Println("<<<<<<<<<<<<<<<Thank you very much>>>>>>>>>>>>>>>")

        } else {
          fmt.Println("<<<<<<<<<<<<<<<<<gracias mucho>>>>>>>>>>>>>>>>")

        }

      } else {
        fmt.Println("Sorry! We can not proceed with your request. Try again")
      }

    } else {
      fmt.Println("<<<<<<<<<<<<<<<<<Bye, laters>>>>>>>>>>>>>>>")

    }

    // if strings.Compare("hi", text) == 0 {
    //   fmt.Println("hello, Yourself")
    // }

  }
}

func tryPayment() {
  //Get Customer's amount payable
  fmt.Print("Confirm amount for your delivery?-> ")
  fmt.Scan(&amount)
  if amount < amountPayable {
    fmt.Printf("Sorry! The amount is not correct, your cost of delivery is %d\n", amountPayable)

  }
}
