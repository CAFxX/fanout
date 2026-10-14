package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint16; F2 bool }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint64; F1 uint8; F2 int32 }
func genFn0(p0 uint8, p1 int16, p2 int8) uint64 {
    if true {
        p0 -= 2
    } else {
        switch 42 {
        case 4294967295:
            p2 = (p2 + (p2 + p2))
        }
    }
    v5 := GenS1{F0: 1, F1: 2}
    p1 = p1
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint8 = p0
        _ = fi6
        _ = v7
    }
    fmt.Println(p2, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 980
}

func main() {
    var v1 uint16 = 547
    var v2 uint16 = v1
    var v3 uint32 = 647
    var v4 int32 = -682
    v8 := []int{-468, 127}
    fmt.Println(v2, v4)
    if len(v8) > 0 {
        v8[0] = -2147483648
    }
    v9 := GenS1{F0: 372, F1: uint64(v2)}
    fmt.Println(v4)
    v10 := GenS2{F0: (v4 << ((672) & 63)), F1: (v1 / ((v2) | 1)), F2: false}
    switch "hello" {
    case "\n\t":
        var v11 string = "hello"
        _ = v11
    case "":
        for fi12 := 0; fi12 < 7; fi12++ {
            v4 -= (v4 / ((-1) | 1))
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi12
        }
    default:
        v3 = v3
    }
    var v13 int32 = (-(v4))
    var v14 bool = false
    switch "\n\t" {
    case "\n\t":
        for fi15 := 0; fi15 < 1; fi15++ {
            for fi16 := 0; fi16 < 3; fi16++ {
                if len(v8) > 0 {
                    v8[0] = fi16
                }
                switch (-(fi15)) {
                case 32767:
                    var v17 uint64 = 0
                    _ = v17
                case -376:
                    if (255 <= 325) {
                        var v18 float64 = ((float64(3.14159) / ((-2.25 - 1e10) * 0.1)) * (float64((((0.1 - -2.25) / ((1.5 * 1.5) + 0.1)) / (-2.25 - (1e10 * (0.1 - (0.0 + 0.1)))))) - -2.25))
                        _ = v18
                    }
                case -1:
                    v2 = v2
                }
                if len(v8) > 0 {
                    v8[0] = 32768
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                var v19 uint32 = 4294967295
                v20 := GenS1{F0: int64(fi15), F1: 1}
                _ = fi16
                _ = v19
                _ = v20
            }
            _ = fi15
        }
    }
    var v21 int8 = 0
    switch 0 {
    case -227:
        v22 := GenS1{F0: 9223372036854775807, F1: 18446744073709551615}
        _ = v22
    case -1:
        v23 := GenS3{F0: ("a" + "a"), F1: -1}
        _ = v23
    case 32767:
        v1 = (v2 * (v1 % ((2) | 1)))
    default:
        if (!v14) {
            switch 1 {
            case 0:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            case -980:
                v21 -= (-(v21))
            }
            v1 = 649
            fmt.Println(v1, v4, v1)
        } else {
            var v24 float64 = 1.5
            _ = v24
        }
    }
    var v25 uint64 = 2
    var v26 bool = (true || (false || true))
    _ = len(v8)
    v1 = v1
    if false {
        for fi27 := 0; fi27 < 15; fi27++ {
            _ = len(v8)
            _ = fi27
        }
    } else {
        var v28 float32 = 0.0
        if v26 {
            v8 = append(v8, -1)
        }
        if len(v8) > 0 {
            v8[0] = 4294967295
        }
        _ = v28
    }
    v29 := GenS4{F0: (^(v25)), F1: 61, F2: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
