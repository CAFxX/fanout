package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int16; F2 uint16 }
type GenS3 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 42
    for fi4 := 0; fi4 < 4; fi4++ {
        v3 = (v3 * v3)
        fmt.Println(fi4)
        _ = fi4
    }
    var v5 int32 = 1
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(v3)
        var v7 bool = ((((1 > 549) || (v3 != v3)) && ((0 != 0) && true)) || true)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 32767}
    v3 += 42
    for fi9 := 0; fi9 < 10; fi9++ {
        v5 = (v5 | 2147483647)
        _ = fi9
    }
    var v10 int16 = 32767
    var v11 uint64 = 551
    switch 32768 {
    case -939:
        v2[726] = (42 <= 953)
    case 9223372036854775807:
        v3 -= (v3 * (v3 | (v3 & 210)))
    case -57:
        for fi12 := 0; fi12 < 9; fi12++ {
            v10 += v10
            v13 := []int{0, fi12, fi12, fi12}
            var v14 int = fi12
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    switch -14 {
    case 0:
        v15 := GenS2{F0: 100, F1: -460, F2: 42}
        _ = v15
    case -534:
        if false {
            var v16 uint8 = v3
            var v17 int64 = int64(v10)
            _ = v16
            _ = v17
        }
    }
    v18 := GenS1{F0: 406}
    v10 += (v10 - v10)
    var v19 uint = 255
    for fi20 := 0; fi20 < 4; fi20++ {
        _ = len(v2)
        fmt.Println(fi20)
        if ((!(4294967295 != (v11 + (v11 / (((v11 / ((42) | 1))) | 1))))) && (v3 < (v3 - (v3 - (v3 % ((v3) | 1)))))) {
            switch -9223372036854775808 {
            case 127:
                fmt.Println(v19, v5)
            default:
                _ = len(v1)
            }
            v21 := GenS1{F0: 1}
            _ = v21
        } else {
            v3 += 36
            var v22 float32 = 1e10
            fmt.Println(v11)
            _ = v22
        }
        _ = fi20
    }
    switch "a" {
    case "a":
        var v23 bool = (255 > 32767)
        _ = v23
    case "foo bar":
        _ = len(v2)
    case "":
        switch "a" {
        case "\n\t":
            v24 := GenS3{F0: (false || false), F1: 42}
            _ = v24
        case "a":
            if (uint(v3) < v19) {
                for fi25 := 0; fi25 < 4; fi25++ {
                    var v26 uint8 = v3
                    _ = fi25
                    _ = v26
                }
                v27 := GenS1{F0: 282}
                for fi28 := 0; fi28 < 14; fi28++ {
                    v29 := GenS3{F0: (true || (124 > -2)), F1: fi28}
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
        case "foo bar":
            v30 := []int{0}
            _ = v30
        default:
            var v31 int = 1
            _ = v31
        }
    default:
        v11 -= v11
    }
    var v32 uint8 = (v3 - (v3 ^ uint8(v19)))
    v33 := v1[486]
    v5 -= (v5 & (^(v5)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v5
    _ = v8
}
