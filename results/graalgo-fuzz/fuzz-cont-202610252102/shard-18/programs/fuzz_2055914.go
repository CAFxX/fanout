package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint; F1 uint16; F2 uint }
func genFn0(p0 int16, p1 uint64, p2 int32) int8 {
    for fi6 := 0; fi6 < 14; fi6++ {
        switch "" {
        case "\n\t":
            var v7 float64 = ((0.0 + 1e10) * 3.14159)
            _ = v7
        }
        _ = fi6
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(p0, fi8)
        _ = fi8
    }
    p0 += 42
    _ = p0
    _ = p1
    _ = p2
    return 100
}
func genFn1(p0 bool, p1 bool) uint16 {
    v9 := GenS1{F0: 297, F1: 2}
    if (!p1) {
        var v10 int64 = 42
        _ = v10
    } else {
        fmt.Println(p0, p1)
        if true {
            v11 := GenS1{F0: 159, F1: 42}
            _ = v11
        }
    }
    v12 := GenS2{F0: 32767, F1: 1, F2: 29}
    v13 := GenS3{F0: 255}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v9
    return 32767
}

func main() {
    var v1 uint8 = 143
    var v2 int32 = -2147483648
    var v3 uint32 = 114
    var v4 int8 = 115
    var v5 int16 = -577
    if (!(!(982 < 0))) {
        fmt.Println(v4, v1)
        var v14 uint16 = uint16(v3)
        v15 := GenS4{F0: (v3 / ((4294967295) | 1)), F1: (v4 + v4), F2: v1}
        _ = v14
        _ = v15
    }
    var v16 uint16 = genFn1(false, (false || true))
    for fi17 := 0; fi17 < 9; fi17++ {
        fmt.Println(v3, v1)
        if true {
            switch "" {
            case "hello":
                v18 := []int{fi17, (fi17 >> ((0) & 63))}
                _ = v18
            case "\n\t":
                var v19 int = fi17
                _ = v19
            case "a":
                v20 := GenS1{F0: (v16 ^ (v16 % (((v16 % ((v16) | 1))) | 1))), F1: 1}
                _ = v20
            default:
                var v21 uint16 = (v16 >> ((308) & 63))
                _ = v21
            }
        }
        _ = fi17
    }
    v5 = -1
    if (1 <= (v5 ^ -1)) {
        v1 = (v1 << ((65535) & 63))
        v22 := []int{42}
        switch 127 {
        case -897:
            v4 = (v4 / ((v4) | 1))
        case -1:
            v3 += uint32(v1)
        }
        _ = v22
    }
    var v23 uint32 = (v3 - v3)
    for fi24 := 0; fi24 < 16; fi24++ {
        v25 := GenS2{F0: v5, F1: uint8(v16), F2: v4}
        var v26 uint16 = v16
        if true {
            v27 := GenS5{F0: 42, F1: genFn1((v3 <= (v23 - v23)), (3.14159 <= 0.0)), F2: 4294967295}
            fmt.Println(v26, v5, fi24)
            _ = v27
        } else {
            var v28 uint32 = v23
            fmt.Println(v16, v28, v1)
            _ = v28
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        v30 := GenS2{F0: -550, F1: (v1 % ((v1) | 1)), F2: (v4 & v4)}
        _ = fi29
        _ = v30
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
