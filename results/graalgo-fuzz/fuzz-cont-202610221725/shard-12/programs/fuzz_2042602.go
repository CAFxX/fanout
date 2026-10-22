package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int64 }
type GenS2 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint64, p1 uint8) uint {
    p0 += (^(p0))
    var v4 uint = 941
    v5 := GenS1{F0: 1, F1: 42, F2: 1}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (v4 / ((197) | 1))
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 975
    var v3 float64 = ((1e10 - -2.25) * 1e10)
    v6 := GenS1{F0: 9223372036854775807, F1: 305, F2: 494}
    v7 := GenS1{F0: -9223372036854775808, F1: 1, F2: 0}
    var v8 bool = false
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 uint = 0
            v8 = (v8 && ((fi10 << (((^(v11))) & 63)) != fi9))
            _ = fi10
            _ = v11
        }
        var v12 int16 = (v2 << ((65535) & 63))
        _ = fi9
        _ = v12
    }
    v8 = (v8 || (false || v8))
    var v13 uint64 = 66
    for fi14 := 0; fi14 < 17; fi14++ {
        v8 = ((245 < 32767) || ((-(v2)) == (v2 / ((v2) | 1))))
        v3 = v3
        _ = fi14
    }
    switch -1 {
    case 888:
        var v15 uint16 = 189
        _ = v15
    case -2147483648:
        switch "hello" {
        case "a":
            v16 := GenS2{F0: 1, F1: -1}
            _ = v16
        case "\n\t":
            v17 := GenS1{F0: -391, F1: -575, F2: -240}
            _ = v17
        default:
            v18 := GenS1{F0: -243, F1: -972, F2: 0}
            _ = v18
        }
    default:
        v13 += v13
    }
    var v19 int8 = 100
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 uint8 = 128
        v13 = v13
        switch (fi20 >> ((65535) & 63)) {
        case 42:
            delete(v1, fi20)
        case 32767:
            var v22 uint32 = 134
            _ = v22
        }
        _ = fi20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
