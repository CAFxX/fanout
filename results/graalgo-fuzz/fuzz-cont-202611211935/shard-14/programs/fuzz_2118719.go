package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint32 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float32 = -2.25
    var v3 uint32 = 810
    var v4 int8 = -128
    var v5 uint16 = 96
    for fi6 := 0; fi6 < 9; fi6++ {
        if ((fi6 >= -2147483648) || ((false && (!false)) && true)) {
            fmt.Println(fi6, v4, v4)
        }
        var v7 float64 = 3.14159
        v8 := []int{fi6, (fi6 ^ fi6), fi6, (fi6 & fi6)}
        _ = fi6
        _ = v7
        _ = v8
    }
    if true {
        v1 = v1
        v9 := GenS1{F0: 1, F1: 18446744073709551615}
        fmt.Println(v4, v2, v3)
        _ = v9
    } else {
        v1 += int64(v5)
        if (2 >= 2) {
            fmt.Println(v4, v1, v1)
            for fi10 := 0; fi10 < 14; fi10++ {
                v11 := []int{}
                v5 -= v5
                if false {
                    _ = len(v11)
                    var v12 string = "hello"
                    v13 := GenS1{F0: (v5 + (v5 + v5)), F1: 904}
                    _ = v12
                    _ = v13
                } else {
                    switch fi10 {
                    case 0:
                        var v14 int16 = 32767
                        _ = v14
                    case 719:
                        var v15 uint32 = v3
                        _ = v15
                    default:
                        fmt.Println(v5, fi10, v3)
                    }
                }
                _ = fi10
                _ = v11
            }
        } else {
            v1 -= v1
            var v16 uint32 = (v3 >> ((4294967295) & 63))
            _ = v16
        }
        v17 := GenS2{F0: 42, F1: (v3 * (v3 + (^(v3))))}
        _ = v17
    }
    if true {
        switch "" {
        case "":
            var v18 uint32 = (^(v3))
            _ = v18
        case "a":
            fmt.Println(v4)
        }
        var v19 uint64 = 1
        _ = v19
    }
    fmt.Println(v4, v1, v5)
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 uint8 = 128
        for fi22 := 0; fi22 < 2; fi22++ {
            fmt.Println(v3, v4)
            var v23 int64 = (v1 >> ((4294967295) & 63))
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    v24 := GenS1{F0: (v5 / ((v5) | 1)), F1: 792}
    var v25 int = 1
    if false {
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 int64 = -1
            var v28 float32 = v2
            _ = fi26
            _ = v27
            _ = v28
        }
        if (!false) {
            v3 += 42
            var v29 string = "\n\t"
            _ = v29
        }
        v30 := []int{v25}
        _ = v30
    } else {
        var v31 uint64 = 42
        v32 := GenS1{F0: v5, F1: v31}
        var v33 int16 = -1
        _ = v31
        _ = v32
        _ = v33
    }
    v3 += v3
    if ((255 < (v5 ^ (v5 - 65535))) || false) {
        v3 -= (v3 - v3)
        var v34 uint32 = v3
        _ = v34
    }
    var v35 int8 = (v4 + (v4 + v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
