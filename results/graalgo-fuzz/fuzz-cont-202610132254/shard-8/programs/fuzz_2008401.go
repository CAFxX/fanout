package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint16 }
func genFn0() uint8 {
    var v7 float64 = (-2.25 - 3.14159)
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(v7, v7)
        _ = fi8
    }
    fmt.Println(v7)
    var v9 int = 32767
    v10 := []int{}
    _ = v10
    _ = v7
    _ = v9
    return 128
}
func genFn1() uint {
    if (-809 != -2147483648) {
        switch 32767 {
        case 32768:
            v11 := []int{}
            _ = v11
        }
        v12 := []int{32768, -1, 32768, 128}
        _ = v12
    }
    var v13 uint16 = 1
    fmt.Println(v13)
    fmt.Println(v13)
    _ = v13
    return 0
}

func main() {
    var v1 uint16 = 0
    var v2 int = 0
    var v3 uint64 = 0
    var v4 uint64 = 993
    v5 := []int{}
    var v6 uint8 = 2
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v14 := GenS1{F0: ("" + ""), F1: -1}
    v15 := GenS2{F0: 734, F1: (v1 + v1)}
    v2 -= (v2 / ((128) | 1))
    for fi16 := 0; fi16 < 17; fi16++ {
        _ = len(v5)
        _ = fi16
    }
    var v17 uint32 = uint32(v6)
    fmt.Println(v3, v2)
    var v18 bool = (false && ((v3 + (^(v3))) > (^(v4))))
    if (42 <= 2147483647) {
        if false {
            _ = len(v5)
            var v19 int = 32767
            _ = v19
        }
    } else {
        v5 = append(v5, v2)
    }
    var v20 uint = 4294967295
    fmt.Println(v18)
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 bool = v18
        fmt.Println(v3, v17)
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
