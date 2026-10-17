package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int }
func genFn0() int8 {
    var v5 uint = 42
    if (!((((((!(1 <= -128)) || (float64(3.14159) < (0.0 + (0.0 / -2.25)))) && false) || false) || false) || ((0.0 <= ((-2.25 + 0.0) * (-2.25 / 0.1))) && (false && (-2.25 == 0.1))))) {
        v6 := []int{-1, 128, 128}
        _ = v6
    }
    _ = v5
    return -128
}

func main() {
    var v1 int = 32767
    var v2 int64 = -1
    v3 := []int{(^(v1)), 32767}
    var v4 int8 = -1
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    var v7 uint8 = 42
    var v8 int64 = int64(v1)
    _ = len(v3)
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint32 = 65535
        v1 += 0
        _ = fi9
        _ = v10
    }
    if ((genFn0() > (v4 & (v4 << ((18446744073709551615) & 63)))) || false) {
        if (false || (1.5 < float64(1e10))) {
            var v11 string = ""
            switch (((v11 + v11) + v11) + v11) {
            case "\n\t":
                v7 -= v7
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v11
        } else {
            if false {
                v8 = v8
            }
        }
    }
    v12 := GenS1{F0: 1}
    var v13 uint16 = 278
    for fi14 := 0; fi14 < 8; fi14++ {
        v7 += (v7 / (((^(v7))) | 1))
        var v15 uint8 = (v7 + (v7 / ((2) | 1)))
        v3 = append(v3, int(v13))
        _ = fi14
        _ = v15
    }
    var v16 int32 = -580
    var v17 float32 = 1.5
    switch "hello" {
    case "hello":
        v18 := GenS2{F0: (v1 + v1), F1: v1}
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
