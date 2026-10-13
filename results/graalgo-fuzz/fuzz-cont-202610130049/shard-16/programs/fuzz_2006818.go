package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int8; F1 int; F2 int32 }
type GenS4 struct { F0 uint64; F1 int8; F2 int8 }
type GenS5 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS7 struct { F0 uint32; F1 int16 }
func genFn0(p0 int64, p1 uint64, p2 uint8) uint {
    p0 -= p0
    if (0.0 <= 3.14159) {
        v4 := GenS1{F0: 117, F1: 127, F2: (p2 + p2)}
        v5 := GenS2{F0: int32(p2), F1: 1}
        _ = v4
        _ = v5
    } else {
        for fi6 := 0; fi6 < 3; fi6++ {
            p1 += 946
            _ = fi6
        }
        v7 := []int{2147483647, -9223372036854775808}
        fmt.Println(p2, p1)
        _ = v7
    }
    p1 = (p1 << ((8) & 63))
    _ = p0
    _ = p1
    _ = p2
    return uint(p1)
}

func main() {
    var v1 uint32 = 570
    v2 := []uint32{v1, (v1 | 42), v1, v1}
    var v3 uint32 = 4294967295
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        if ((0.0 + 1e10) != (0.1 * 1e10)) {
            v10 := GenS3{F0: -1, F1: -475, F2: 525}
            var v11 int32 = 42
            _ = v10
            _ = v11
        } else {
            fmt.Println(fi9, v3)
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            _ = len(v2)
            _ = fi12
        }
        v13 := GenS1{F0: v3, F1: 127, F2: 42}
        _ = fi9
        _ = v13
    }
    v14 := GenS4{F0: 42, F1: int8(v3), F2: 100}
    v1 += v3
    v15 := GenS5{F0: -1, F1: 2, F2: 426}
    var v16 int = 190
    v1 = v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 += 0
    var v17 int8 = 108
    if (!(v17 == -121)) {
        switch (v16 & (v16 % ((1) | 1))) {
        case -2147483648:
            v17 += v17
        }
        if (0 == 4294967295) {
            fmt.Println(v1, v16, v1)
        } else {
            if (false || (0.1 == float32(0.1))) {
                v17 += v17
            }
            v18 := GenS6{F0: 0, F1: 0, F2: 42}
            _ = v18
        }
        var v19 float32 = (3.14159 * 1.5)
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        v1 = v3
        if (244 >= 91) {
            var v21 float64 = ((0.1 / 1e10) * (3.14159 / -2.25))
            var v22 uint8 = 128
            v2 = append(v2, (^(v1)))
            _ = v21
            _ = v22
        } else {
            v17 -= (v17 + 42)
        }
        var v23 int32 = 1
        _ = fi20
        _ = v23
    }
    v24 := GenS7{F0: v1, F1: -32768}
    v17 = (v17 ^ v17)
    var v25 uint16 = 2
    fmt.Println(v25, v16, v1)
    for fi26 := 0; fi26 < 5; fi26++ {
        if (!true) {
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 int = (v16 | v16)
                _ = fi27
                _ = v28
            }
            switch ((("\n\t" + "hello") + ("foo bar" + "foo bar")) + (("hello" + "hello") + "hello")) {
            case "":
                v29 := GenS2{F0: -2147483648, F1: (v17 / (((v17 - -84)) | 1))}
                _ = v29
            case "hello":
                v16 = -97
            case "foo bar":
                v30 := GenS6{F0: 157, F1: 0, F2: v25}
                _ = v30
            }
        } else {
            v1 = v1
            var v31 uint16 = v25
            for fi32 := 0; fi32 < 3; fi32++ {
                switch (fi26 % ((128) | 1)) {
                case -533:
                    v3 -= v3
                    v1 = 4294967295
                case 1:
                    v33 := GenS3{F0: v17, F1: (fi32 % (((fi26 / ((42) | 1))) | 1)), F2: 565}
                    _ = v33
                case -1:
                    fmt.Println(v3, v16)
                }
                var v34 string = "foo bar"
                v17 -= v17
                _ = fi32
                _ = v34
            }
            _ = v31
        }
        if (!(1.5 < -2.25)) {
            for fi35 := 0; fi35 < 3; fi35++ {
                v36 := GenS4{F0: 18446744073709551615, F1: 0, F2: 42}
                _ = fi35
                _ = v36
            }
        }
        switch 4294967295 {
        case 1:
            v3 = v1
        case 32768:
            var v37 bool = true
            _ = v37
        default:
            var v38 int32 = -332
            _ = v38
        }
        _ = fi26
    }
    fmt.Println(v16, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
