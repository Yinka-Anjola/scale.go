package main 

import "fmt"

func main() {
        var value, result float64
        var fromUnit, toUnit string

        fmt.Println("Enter the value: ")
        fmt.Scanln(&value)

        fmt.Println("Enter a desired input unit (kilogram, pound or gram): ")
        fmt.Scanln(&fromUnit)

        fmt.Println("Enter a desired output unit (kilogram, pound or gram): ")
        fmt.Scanln(&toUnit)

        switch fromUnit{
        case "kilogram":
                switch toUnit{
                case "pound":
                        result = value * 2.205
                case "gram":
                        result = value * 1000
                default: 
                        fmt.Println("invalid input")
                        return
                }
        case "pound":
                switch toUnit{
                case "kilogram":
                        result = value * 0.453
                case "gram":
                        result = value * 453.592 
                default:
                        fmt.Println("invalid input")
                        return
                }
        case "gram":
                switch toUnit{
                case "kilogram":
                        result = value / 1000
                case "pound":
                        result = value / 453.592
                default:
                        fmt.Println("invalid input")
                        return
                }
        }
        fmt.Printf("%.2f %s is equal to %.2f %s\n", value, fromUnit, result, toUnit)
}