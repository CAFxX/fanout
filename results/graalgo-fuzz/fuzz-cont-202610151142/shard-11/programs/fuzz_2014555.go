package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint) uint {
    if false {
        v7 := GenS1{F0: -266, F1: 2147483647}
        _ = v7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int64 = 492
        _ = fi8
        _ = v9
    }
    _ = p0
    return (p0 + p0)
}

func main() {
    var v1 uint32 = 113
    var v2 uint8 = 128
    var v3 int16 = 1000
    var v4 float64 = 1.5
    v5 := make(map[string]uint32)
    var v6 int16 = v3
    v6 = int16(v2)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 float32 = 1e10
        v12 := GenS1{F0: 9223372036854775807, F1: (fi10 * 1)}
        var v13 int16 = (-(v6))
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        v2 = (v2 ^ uint8(v1))
        _ = fi14
    }
    v1 = (^(v1))
    if true {
        for fi15 := 0; fi15 < 4; fi15++ {
            v2 = (v2 >> ((2) & 63))
            _ = fi15
        }
        fmt.Println(v4, v1)
    } else {
        v6 += v6
        switch ("\n\t" + "") {
        case "a":
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS2{F0: (!true), F1: (^(v1))}
                var v18 bool = true
                v1 -= (^(v1))
                _ = fi16
                _ = v17
                _ = v18
            }
        case "\n\t":
            v4 -= v4
        default:
            fmt.Println(v3, v3)
        }
    }
    var v19 string = ("foo bar" + "a")
    var v20 int = 32767
    v20 = 9223372036854775807
    if false {
        v2 = v2
        fmt.Println(v2, v19, v6)
        v21 := GenS2{F0: ((^(v2)) <= (v2 | (v2 / ((v2) | 1)))), F1: v1}
        _ = v21
    } else {
        v5["\n\t"] = (v1 & v1)
        v20 -= (v20 * (v20 / ((v20) | 1)))
    }
    switch 128 {
    case 610:
        v19 = v19
    case -195:
        switch v20 {
        case 384:
            v2 += (v2 | (v2 - (v2 / ((v2) | 1))))
        }
    default:
        switch (v20 * 22) {
        case -650:
            v22 := GenS1{F0: 42, F1: v20}
            _ = v22
        case -9223372036854775808:
            if true {
                delete(v5, (v19 + "\n\t"))
                var v23 int32 = 377
                v24 := []int{(v20 % ((-9223372036854775808) | 1))}
                _ = v23
                _ = v24
            }
        }
    }
    var v25 bool = (true && (0 >= v1))
    var v26 int8 = 100
    fmt.Println(v1)
    v5[v19] = 908
    v27 := GenS1{F0: 0, F1: (v20 | (-(v20)))}
    v4 = v4
    fmt.Println(v4, v25, v19)
    var v28 int8 = v26
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
