package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 string }
func genFn0() int32 {
    v5 := GenS1{F0: 255, F1: 4294967295, F2: "hello"}
    var v6 int64 = -9223372036854775808
    v7 := []int{}
    v8 := GenS1{F0: 42, F1: -373, F2: ("a" + "")}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    var v1 uint8 = 2
    var v2 int16 = 0
    var v3 uint32 = 1
    var v4 uint16 = 748
    v9 := []int{127}
    var v10 int = -9223372036854775808
    var v11 uint = 42
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    if ((v10 ^ v10) != (^(v10))) {
        v3 = v3
        if (!false) {
            v2 = (v2 + v2)
        }
    }
    v11 += v11
    fmt.Println(v1)
    switch 32767 {
    case -541:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    default:
        for fi12 := 0; fi12 < 16; fi12++ {
            v10 = v10
            _ = fi12
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
