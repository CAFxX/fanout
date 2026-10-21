package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    var v1 float32 = (((-2.25 + (1e10 * 1.5)) - 1.5) / 1.5)
    v2 := make(map[int]string)
    var v3 int = 129
    var v4 int32 = 0
    v5 := make(map[uint32]int)
    var v6 uint = 338
    if ((("" + "hello") + "\n\t") < "\n\t") {
        v4 += (v4 % (((v4 * (v4 ^ v4))) | 1))
        v2[(-(v3))] = "hello"
        v7 := []int{(v3 % (((v3 / ((9223372036854775807) | 1))) | 1))}
        _ = v7
    } else {
        var v8 int16 = 1000
        _ = v8
    }
    _ = len(v2)
    var v9 string = "\n\t"
    v10 := GenS1{F0: true, F1: (v3 + 0)}
    v1 -= 0.1
    v11 := GenS1{F0: false, F1: v3}
    v9 = "\n\t"
    var v12 uint32 = 42
    if ((false && false) || (v12 > (v12 - (v12 % (((v12 | 2147483647)) | 1))))) {
        if ((v4 << ((v6) & 63)) >= v4) {
            v13 := GenS1{F0: true, F1: (v3 & 2147483647)}
            v14 := []int{v3, v3}
            _ = v13
            _ = v14
        } else {
            if (2 <= 4294967295) {
                v15 := []int{-2147483648, 1}
                var v16 int = v3
                var v17 bool = (v1 != (v1 - v1))
                _ = v15
                _ = v16
                _ = v17
            } else {
                var v18 float64 = 0.0
                _ = v18
            }
            switch v3 {
            case 1:
                v19 := GenS1{F0: true, F1: (-(v3))}
                _ = v19
            case 42:
                fmt.Println(v12, v1, v12)
            case 32767:
                v6 += 137
            }
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 int64 = 9223372036854775807
                v22 := []int{v3, fi20, 32767, 42}
                var v23 uint64 = 770
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        }
        v24 := GenS1{F0: (v6 <= (v6 - (^(v6)))), F1: v3}
        var v25 uint = (^(v6))
        _ = v24
        _ = v25
    } else {
        v9 = "a"
    }
    v26 := v5[(v12 + v12)]
    v4 = v4
    switch "" {
    case "hello":
        v27 := GenS1{F0: false, F1: v26}
        _ = v27
    case "a":
        if false {
            for fi28 := 0; fi28 < 13; fi28++ {
                var v29 int = (v26 + (v3 % (((fi28 % ((42) | 1))) | 1)))
                v6 = (v6 + 0)
                var v30 int32 = v4
                _ = fi28
                _ = v29
                _ = v30
            }
        }
    default:
        v5[v12] = v3
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        delete(v2, v26)
        _ = fi31
    }
    if ((v12 / (((v12 / (((^(v12))) | 1))) | 1)) >= v12) {
        var v32 float32 = v1
        v33 := GenS1{F0: false, F1: (v3 & v26)}
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
