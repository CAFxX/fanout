package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 bool; F1 int16; F2 int16 }
func genFn0(p0 string, p1 int32, p2 int32) uint32 {
    v4 := []int{0, 688}
    v5 := GenS1{F0: 42, F1: 42}
    if true {
        if len(v4) > 0 {
            v4[0] = 9223372036854775807
        }
    } else {
        v6 := GenS2{F0: false, F1: -1, F2: -130}
        if ((0 <= 0) && true) {
            var v7 uint8 = 59
            _ = len(v4)
            _ = v7
        }
        if false {
            fmt.Println(p1, p0)
        }
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 596
}

func main() {
    var v1 int = -974
    var v2 uint16 = 167
    var v3 uint8 = 1
    var v8 int = (v1 % ((-263) | 1))
    switch -1 {
    case -728:
        fmt.Println(v3, v8, v3)
    default:
        fmt.Println(v3, v1, v1)
    }
    var v9 uint = 121
    v2 -= (v2 % (((v2 >> (((v9 << ((v9) & 63))) & 63))) | 1))
    for fi10 := 0; fi10 < 1; fi10++ {
        v2 = 675
        v11 := GenS2{F0: true, F1: -32768, F2: 0}
        v12 := GenS2{F0: (uint32(v3) < 162), F1: 0, F2: -302}
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        v13 := []int{v8}
        for fi14 := 0; fi14 < 3; fi14++ {
            fmt.Println(v1)
            _ = fi14
        }
        _ = v13
    } else {
        switch "foo bar" {
        case "":
            if (!(2 >= 476)) {
                v3 += v3
            }
        case "\n\t":
            fmt.Println(v8)
        case "foo bar":
            v15 := GenS1{F0: v2, F1: 450}
            _ = v15
        }
    }
    v16 := GenS1{F0: (v2 & v2), F1: (v9 >> ((18446744073709551615) & 63))}
    var v17 uint = v9
    var v18 int16 = 1
    fmt.Println(v2)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
