package main

import (
    "time"
    "fmt"
)
//Function to return last Business day
// func dummy() time.Weekday {
func dummy( t time.Time ) time.Weekday {
    fmt.Println("- Its a :", time.Time(t).Weekday())
    switch today := time.Time(t).Weekday() ; {
    //switch today := time.Now().Weekday() ; {
    // if it is Sunday
    case today == 0 :
        fmt.Println("- case 0",today+5)
        return today+5
    // if it is Monday
    case today == 1 :
        fmt.Println("- case 1",today+4)
        return today+4
    // if it is Tuesday
    case today == 2 :
        fmt.Println("- case 2",today-1)
        return today-1
    // if it is Wednesday
    case today == 3 :
        fmt.Println("- case 3",today-1)
        return today-1
    // if it is Thursday
    case today == 4 :
        fmt.Println("- case 4",today-1)
        return today-1
    // if it is Friday
    case today == 5 :
        fmt.Println("- case 5",today-1)
        return today-1
    // if it is Saturday
    case today == 6 :
        fmt.Println("- case 6",today-1)
        return today-1
    default:
        fmt.Println("- case d",today+1)
        return today-1
    }
}

func main()  {
//    t := time.Date(2019, 6, 18 , 0, 0, 0, 0, time.UTC)
  //  fmt.Printf("%v\n",t.Weekday() )
    //fmt.Printf("%v\n",dummy(t) )
   
    //time.In("UTC")
    timeStamp, _ := time.Parse("2006/01/02", "2019/06/30" )
    //timeStamp, _ := time.Parse(time.RFC3339, "2019-02-19T00:00:00+00:00" )
    if timeStamp.IsZero()  {
       fmt.Println("Is zero")
   } else {
       fmt.Println(timeStamp) 
   }
}
