package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int8 }
type GenS2 struct { F0 uint16; F1 string }

func main() {
    var v1 int16 = 0
    var v2 int = 128
    var v3 int = v2
    var v4 uint = 1
    switch (v3 + (v2 << (((^(v4))) & 63))) {
    case 390:
        v5 := []int{127}
        _ = v5
    case -549:
        v4 += v4
    case -9223372036854775808:
        v6 := GenS1{F0: 653, F1: -1, F2: 10}
        _ = v6
    default:
        v4 -= (^(v4))
    }
    if false {
        v1 = -447
        switch ("" + "") {
        case "a":
            v7 := GenS2{F0: 255, F1: ("\n\t" + "")}
            _ = v7
        }
    }
    v1 = (^(v1))
    v1 += -389
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS2{F0: 255, F1: "hello"}
        _ = fi8
        _ = v9
    }
    v10 := []int{9223372036854775807, v2, v3, 2147483647}
    fmt.Println(v3, v4)
    var v11 int = v2
    if false {
        switch 9223372036854775807 {
        case -2147483648:
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 bool = true
                _ = fi12
                _ = v13
            }
        }
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        fmt.Println(v4)
    } else {
        v11 = (v11 - 530)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
