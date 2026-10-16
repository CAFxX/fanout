package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := []int{1, -627, -542}
    var v2 uint8 = 232
    var v3 int64 = 0
    v3 = v3
    for fi4 := 0; fi4 < 17; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := GenS1{F0: (v2 >> ((18446744073709551615) & 63)), F1: v2}
            var v7 float32 = -2.25
            v7 = 3.14159
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    _ = len(v1)
    fmt.Println(v3, v3)
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    var v8 string = "\n\t"
    v2 -= (v2 ^ v2)
    v9 := GenS2{F0: -128}
    var v10 bool = true
    switch 9223372036854775807 {
    case -329:
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS2{F0: 9}
            var v13 uint16 = 0
            _ = fi11
            _ = v12
            _ = v13
        }
    case 2147483647:
        switch "\n\t" {
        case "":
            var v14 uint = 1
            _ = v14
        case "\n\t":
            v3 -= 354
        default:
            v8 = v8
        }
    }
    v10 = (!false)
    if true {
        fmt.Println(v2, v8)
    } else {
        _ = len(v1)
    }
    v8 = (v8 + "\n\t")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
