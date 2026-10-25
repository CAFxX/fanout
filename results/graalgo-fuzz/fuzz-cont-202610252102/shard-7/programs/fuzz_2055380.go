package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 int16; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint32; F1 int64 }

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]bool)
    var v3 uint8 = 0
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]uint32)
    v6 := []string{}
    var v7 int16 = 704
    var v8 string = "\n\t"
    v9 := GenS1{F0: 2, F1: false, F2: (1 >= -907)}
    v7 -= 42
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int32 = 0
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        switch "hello" {
        case "":
            v13 := GenS2{F0: (fi12 % ((-1) | 1))}
            _ = v13
        case "\n\t":
            for fi14 := 0; fi14 < 7; fi14++ {
                switch (-(fi12)) {
                case 1:
                    v3 = (v3 | 96)
                default:
                    v15 := GenS3{F0: (v7 << ((997) & 63))}
                    _ = v15
                }
                _ = len(v4)
                _ = fi14
            }
        case "a":
            for fi16 := 0; fi16 < 16; fi16++ {
                v17 := GenS4{F0: (v7 | 42)}
                var v18 int = (fi16 & 411)
                v7 -= (v7 ^ (v7 - 1000))
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        if (false && false) {
            switch v8 {
            case "foo bar":
                v8 = v8
            default:
                var v19 float64 = 1.5
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 1; fi20++ {
            v3 = (v3 ^ (v3 - 127))
            _ = fi20
        }
        _ = fi12
    }
    if false {
        v21 := GenS3{F0: v7}
        var v22 uint8 = (v3 ^ v3)
        _ = v21
        _ = v22
    }
    v23 := GenS5{F0: -879, F1: v8}
    if (!(((false && false) || true) && false)) {
        v6 = append(v6, ((v8 + (v8 + v8)) + v8))
        var v24 uint = 1
        _ = v24
    }
    if true {
        switch 128 {
        case 4294967295:
            var v25 int = 4
            _ = v25
        case 285:
            delete(v5, 2147483647)
        }
        switch 848 {
        case 967:
            if (1 > -9223372036854775808) {
                if (!(-9223372036854775808 < 9223372036854775807)) {
                    fmt.Println(v3)
                    v26 := GenS2{F0: -465}
                    v1 += 105
                    _ = v26
                }
                var v27 float64 = (1e10 * float64(float32(3.14159)))
                _ = v27
            } else {
                v28 := GenS6{F0: (v7 & (v7 + 1000)), F1: 255, F2: v3}
                fmt.Println(v3, v8, v7)
                var v29 uint8 = (^(v3))
                _ = v28
                _ = v29
            }
        case -545:
            for fi30 := 0; fi30 < 7; fi30++ {
                for fi31 := 0; fi31 < 3; fi31++ {
                    var v32 string = "\n\t"
                    v33 := GenS7{F0: 2, F1: -9223372036854775808}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = fi30
            }
        default:
            for fi34 := 0; fi34 < 8; fi34++ {
                var v35 int32 = -242
                _ = fi34
                _ = v35
            }
        }
    } else {
        v36 := GenS2{F0: 127}
        v7 = v7
        v8 = "hello"
        _ = v36
    }
    var v37 uint = 18446744073709551615
    var v38 int8 = (v1 + v1)
    fmt.Println(v1, v3, v7)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
