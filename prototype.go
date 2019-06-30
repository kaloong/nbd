package main

import (
    "time"
    "fmt"
    "os"
    "flag"
)

//Function to return last Business day
// func dummy() time.Weekday {
func getLastBusinessDay( t time.Time ) time.Weekday {
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

func main() {
    if len(os.Args) <= 1 {
        fmt.Printf("No command specified.\n")
        return
    }

    getCmd := flag.NewFlagSet("get", flag.ExitOnError)
    tzFlag := getCmd.String("t","","Timezone name")
    bdFlag := getCmd.String("b","","Business Date name")
    hpFlag := getCmd.String("h","","Get help")

    switch os.Args[1] {
    case "get":
        getCmd.Parse(os.Args[2:])
    default:
        fmt.Printf("%q is not valid command.\n", os.Args[1])
        os.Exit(2)
    }
    if getCmd.Parsed() { 
        if *hpFlag != "" {
            fmt.Println("Get help page")
            os.Exit(2)
        }
        if *tzFlag == "" {
            fmt.Println("Please provide timezone.")
            return
        }
        if *bdFlag == "" {
            fmt.Println("Please provide business date.")
            return
        }
        fmt.Printf("Timezone selected: %q\n",*tzFlag)
        fmt.Printf("Bus Date selected: %q\n",*bdFlag)

    }
    loc, err := time.LoadLocation (*tzFlag)
    if err != nil {
         panic(err)
    }
    formats := [] string{"20060102", "2006/01/02", "2006-01-02"}
    for i := range formats {
        if timeStamp, err := time.Parse(formats[i], *bdFlag); err == nil {
            //t := time.Now().Add(*bdFlag)
            fmt.Printf("- %v %v\n", timeStamp.In(loc), getLastBusinessDay(timeStamp.In(loc)) )
        }
    }
}
