package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint8 = 6
    var v2 uint = 65535
    v3 := []int{1, 625}
    var v4 uint16 = 202
    var v5 uint16 = (v4 >> (((^(v2))) & 63))
    var v6 bool = false
    var v7 bool = (float64(0.0) < 1.5)
    for fi8 := 0; fi8 < 14; fi8++ {
        switch "a" {
        case "\n\t":
            var v9 int = (fi8 * (fi8 >> (((^(v2))) & 63)))
            _ = v9
        default:
            if false {
                fmt.Println(fi8)
                v10 := GenS1{F0: 965, F1: ("\n\t" + "\n\t")}
                _ = v10
            } else {
                v11 := GenS1{F0: 2147483648, F1: "hello"}
                fmt.Println(v2, v2, v7)
                _ = len(v3)
                _ = v11
            }
        }
        _ = len(v3)
        _ = fi8
    }
    v4 = (v5 / ((v4) | 1))
    if true {
        v4 = v5
        v1 -= (v1 % (((^(v1))) | 1))
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 int = 300
            fmt.Println(v7)
            _ = fi12
            _ = v13
        }
    }
    v14 := GenS2{F0: (v1 & v1)}
    switch "a" {
    case "hello":
        v4 += v4
    default:
        v6 = v7
    }
    var v15 int32 = -600
    _ = len(v3)
    _ = len(v3)
    _ = len(v3)
    v1 = v1
    v16 := GenS1{F0: 1, F1: (("" + ((("" + "hello") + "hello") + "\n\t")) + "hello")}
    v1 = 244
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
