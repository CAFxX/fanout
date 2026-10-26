package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    var v1 uint64 = 1
    var v2 float64 = -2.25
    var v3 uint = 737
    var v4 uint8 = uint8(v3)
    v3 = (v3 + v3)
    v5 := GenS1{F0: (v3 % ((1) | 1))}
    v6 := []int{int(v1), 42, -9223372036854775808}
    v2 = 1.5
    switch ("hello" + "foo bar") {
    case "":
        if false {
            var v7 uint16 = 553
            switch "a" {
            case "hello":
                v8 := GenS1{F0: v3}
                _ = v8
            case "\n\t":
                fmt.Println(v3)
            }
            _ = v7
        } else {
            fmt.Println(v2, v2)
        }
    }
    v1 += v1
    for fi9 := 0; fi9 < 2; fi9++ {
        switch fi9 {
        case -2147483648:
            v2 = -2.25
        default:
            if false {
                fmt.Println(v4, fi9, v4)
                v4 += 29
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        }
        v10 := GenS2{F0: ("\n\t" + ("foo bar" + "")), F1: 2147483647}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 uint64 = 2
        v13 := GenS2{F0: "hello", F1: 42}
        if false {
            v4 += v4
            fmt.Println(v3)
            var v14 bool = true
            _ = v14
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
