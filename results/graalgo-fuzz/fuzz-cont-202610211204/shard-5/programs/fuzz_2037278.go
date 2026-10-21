package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int64 }
type GenS2 struct { F0 uint16; F1 int16; F2 int }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 bool; F1 uint16 }

func main() {
    var v1 int = 1
    var v2 int32 = 1
    var v3 uint8 = 128
    var v4 int8 = 127
    v5 := []int{(v1 | (v1 ^ 127))}
    fmt.Println(v2)
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(v1, v1, v3)
        _ = fi6
    }
    var v7 float64 = -2.25
    if (("\n\t" > ("foo bar" + ("a" + ""))) && true) {
        if ((-110 > 987) && (v7 > v7)) {
            var v8 uint32 = 192
            v7 -= (v7 * ((v7 + v7) - 1.5))
            _ = v8
        }
        switch v1 {
        case 0:
            var v9 uint16 = 1
            _ = v9
        case -2147483648:
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 int64 = 9223372036854775807
                _ = fi10
                _ = v11
            }
        default:
            var v12 int32 = (v2 & v2)
            _ = v12
        }
        v13 := GenS1{F0: v3, F1: 1, F2: 1}
        _ = v13
    } else {
        v14 := GenS2{F0: 255, F1: -1, F2: v1}
        v2 = v2
        _ = v14
    }
    v5 = append(v5, (^(v1)))
    fmt.Println(v2)
    if false {
        var v15 uint32 = 1
        v15 = v15
        _ = v15
    }
    v16 := GenS3{F0: true, F1: v3}
    v4 = (v4 / ((-1) | 1))
    switch (v1 & v1) {
    case -1:
        var v17 int64 = 42
        _ = v17
    case -2147483648:
        fmt.Println(v7, v3, v2)
    default:
        v5 = append(v5, (v1 + v1))
    }
    var v18 uint16 = 0
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 uint64 = 42
        v20 -= v20
        v1 += (-(fi19))
        _ = fi19
        _ = v20
    }
    v7 -= v7
    if (-494 > -1) {
        for fi21 := 0; fi21 < 3; fi21++ {
            v7 -= 0.1
            var v22 int16 = -349
            _ = fi21
            _ = v22
        }
        var v23 int = (v1 + v1)
        var v24 uint = 466
        _ = v23
        _ = v24
    }
    var v25 float64 = v7
    v26 := GenS4{F0: (2147483648 < 2147483648), F1: v18}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
