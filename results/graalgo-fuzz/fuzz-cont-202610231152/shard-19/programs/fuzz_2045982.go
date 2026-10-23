package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint16 = 42
    v2 := []string{"a", ("\n\t" + (("" + "hello") + "hello")), "a"}
    var v3 int16 = -32768
    var v4 uint32 = 36
    if (227 >= 128) {
        switch 1 {
        case -2147483648:
            var v5 int32 = 1
            _ = v5
        default:
            switch ("\n\t" + "hello") {
            case "foo bar":
                for fi6 := 0; fi6 < 1; fi6++ {
                    v1 += v1
                    _ = fi6
                }
            case "a":
                switch "foo bar" {
                case "\n\t":
                    if false {
                        var v7 bool = (!false)
                        switch "a" {
                        case "\n\t":
                            var v8 float64 = 1.5
                            _ = v8
                        default:
                            v2 = append(v2, "\n\t")
                        }
                        _ = len(v2)
                        _ = v7
                    } else {
                        _ = len(v2)
                        v9 := GenS1{F0: false}
                        _ = v9
                    }
                case "a":
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                case "":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            case "":
                fmt.Println(v1, v4)
            }
        }
    }
    var v10 uint8 = 2
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            v1 += v1
            fmt.Println(fi12)
            fmt.Println(fi12)
            _ = fi12
        }
        fmt.Println(fi11, v1)
        v4 = 2147483648
        _ = fi11
    }
    if true {
        var v13 int32 = 0
        _ = v13
    }
    var v14 string = "\n\t"
    var v15 int32 = -2147483648
    var v16 uint8 = v10
    v14 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
