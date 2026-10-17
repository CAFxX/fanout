package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 string; F2 uint64 }
type GenS6 struct { F0 bool }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: 4294967295, F1: 100, F2: fi4}
        var v6 int64 = -1
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int16 = 1
            v8 = (v8 & -32768)
            _ = fi7
            _ = v8
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    v9 := GenS1{F0: 1, F1: 62, F2: -9223372036854775808}
    v10 := GenS1{F0: 467, F1: 14, F2: 32768}
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 float32 = float32(1.5)
    var v3 uint16 = 472
    switch "\n\t" {
    case "foo bar":
        var v11 int = -759
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 6; fi12++ {
            switch ("\n\t" + "") {
            case "hello":
                var v13 uint64 = 1
                _ = v13
            default:
                v14 := GenS1{F0: 4294967295, F1: 1, F2: (^(fi12))}
                _ = v14
            }
            _ = fi12
        }
    default:
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS2{F0: 2}
            var v17 int8 = 127
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    var v18 uint8 = 0
    var v19 float64 = 1e10
    if true {
        v2 += v2
        if (255 == 18446744073709551615) {
            var v20 int32 = 730
            _ = v20
        }
        if len(v1) > 0 {
            v1[0] = 42
        }
    }
    v21 := GenS2{F0: 18446744073709551615}
    v22 := GenS3{F0: -42, F1: (v18 / ((v18) | 1)), F2: uint16(v18)}
    v18 = (v18 + v18)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 -= (v3 >> ((998) & 63))
    v23 := GenS2{F0: 42}
    v24 := GenS4{F0: "\n\t"}
    var v25 uint = 1
    _ = len(v1)
    if ((false && true) && (!((true || (v19 <= v19)) && (127 < 1)))) {
        var v26 int8 = 1
        if (false || false) {
            v19 -= (float64((v2 + v2)) * (v19 + v19))
        } else {
            var v27 uint32 = uint32(v25)
            v28 := GenS5{F0: 42, F1: ("foo bar" + "a"), F2: 18446744073709551615}
            _ = v27
            _ = v28
        }
        v25 -= (v25 / ((v25) | 1))
        _ = v26
    } else {
        fmt.Println(v18, v3, v2)
        v29 := GenS5{F0: 2, F1: "a", F2: 1}
        _ = v29
    }
    v30 := GenS6{F0: (100 <= 112)}
    v3 = v3
    fmt.Println(v3, v3)
    v25 += v25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
}
