package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 bool; F2 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 329
    var v3 int16 = (v2 - 42)
    var v4 uint8 = 178
    var v5 int64 = 1
    for fi6 := 0; fi6 < 8; fi6++ {
        v4 -= (v4 | v4)
        _ = fi6
    }
    fmt.Println(v4)
    if true {
        if false {
            v2 = v3
            var v7 uint16 = 65535
            _ = v7
        }
        v2 += -767
        if false {
            v8 := []int{}
            var v9 int32 = -561
            _ = v8
            _ = v9
        } else {
            var v10 uint = 648
            v11 := []int{4294967295, -478, 32767}
            var v12 uint64 = 673
            _ = v10
            _ = v11
            _ = v12
        }
    }
    switch "a" {
    case "foo bar":
        var v13 bool = (false && (0 < 127))
        _ = v13
    case "\n\t":
        v5 += (^(v5))
    case "a":
        v14 := []int{-1, 42}
        _ = v14
    default:
        var v15 uint = 18446744073709551615
        _ = v15
    }
    if (!false) {
        if ((!(18446744073709551615 != 1)) || (1 == 902)) {
            var v16 int = 4294967295
            v2 = (v3 - 535)
            if (1e10 > 0.1) {
                v2 += (v3 - -32768)
                v17 := GenS1{F0: 0}
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 uint8 = v4
                    _ = fi18
                    _ = v19
                }
                _ = v17
            } else {
                var v20 uint8 = (v4 * (v4 | (v4 * 73)))
                v21 := GenS1{F0: -2147483648}
                v3 = int16(v16)
                _ = v20
                _ = v21
            }
            _ = v16
        }
    } else {
        v22 := GenS2{F0: v3, F1: (!((1.5 + 1e10) != (1.5 - ((-2.25 + 1.5) * 1.5)))), F2: 2147483647}
        switch -1 {
        case 0:
            v23 := []int{1, 1, 455, 318}
            _ = v23
        case 1:
            switch -536 {
            case 42:
                v4 = (v4 & v4)
            }
        case 9223372036854775807:
            var v24 float64 = (1.5 + 3.14159)
            _ = v24
        }
        v25 := []int{-38, 886}
        _ = v22
        _ = v25
    }
    v4 = v4
    v26 := GenS3{F0: "\n\t"}
    fmt.Println(v4, v3, v2)
    for fi27 := 0; fi27 < 2; fi27++ {
        v5 -= v5
        _ = fi27
    }
    fmt.Println(v3, v4, v2)
    v28 := GenS4{F0: false}
    v2 -= (v2 & v2)
    for fi29 := 0; fi29 < 5; fi29++ {
        fmt.Println(v5, v2, v4)
        v3 += (v3 / ((1) | 1))
        var v30 float64 = 0.0
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
