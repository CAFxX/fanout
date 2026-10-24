package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int32; F1 int16; F2 bool }

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 2
    var v3 int64 = -9223372036854775808
    v4 := []uint32{65535}
    var v5 uint16 = 32767
    if ((false || ((9223372036854775807 != 9223372036854775807) || (false || (1 < 183)))) && (true || (1 < -277))) {
        if true {
            fmt.Println(v2, v5, v2)
            var v6 string = "a"
            _ = v6
        }
    }
    var v7 uint64 = 919
    var v8 bool = (!true)
    for fi9 := 0; fi9 < 3; fi9++ {
        if false {
            for fi10 := 0; fi10 < 2; fi10++ {
                v3 += (v3 - -1)
                _ = fi10
            }
            v1[fi9] = fi9
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 uint32 = 65535
                var v13 uint = 876
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        _ = fi9
    }
    switch 0 {
    case 32768:
        var v14 uint32 = 4294967295
        _ = v14
    }
    if len(v4) > 0 {
        v4[0] = 865
    }
    var v15 uint = 18446744073709551615
    if true {
        if (!false) {
            var v16 string = ("" + "hello")
            _ = v16
        } else {
            v3 = -9223372036854775808
        }
        v3 = v3
    } else {
        switch 128 {
        case 32768:
            var v17 float64 = 3.14159
            _ = v17
        case 2147483647:
            v18 := GenS1{F0: -115}
            _ = v18
        }
    }
    if ((((0.1 / 3.14159) - 3.14159) <= (float64(-2.25) - (0.0 + 1e10))) || v8) {
        v15 -= (v15 * 2)
    } else {
        v19 := GenS2{F0: 2, F1: 820, F2: v2}
        v20 := GenS2{F0: uint32(v2), F1: v5, F2: (v2 ^ (v2 % ((v7) | 1)))}
        _ = len(v4)
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        if (!(v8 && ((1.5 * (-2.25 - 1e10)) != ((1e10 - 0.0) + (float64((1e10 / ((0.1 + (1e10 + ((0.0 * 0.0) + ((1e10 + (((3.14159 * (-2.25 + 0.0)) / -2.25) - 0.1)) * 1.5)))) + ((float32(3.14159) + 0.1) - (0.0 / (float32((3.14159 * 1.5)) + 0.0)))))) + -2.25))))) {
            v3 = 42
        }
        v22 := GenS3{F0: (!(((1e10 / ((-2.25 + 1.5) + float32(0.1))) / 1.0) == 3.14159))}
        _ = fi21
        _ = v22
    }
    v23 := GenS4{F0: -9, F1: (-140 <= v3)}
    var v24 uint32 = 4294967295
    v2 = (v7 + (v2 & v2))
    var v25 int16 = 32767
    if (v25 < int16(v15)) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (v3 != (v3 >> ((v15) & 63))) {
            var v26 uint16 = (v5 << ((445) & 63))
            _ = v26
        } else {
            var v27 uint16 = 42
            if false {
                switch -1 {
                case 127:
                    v28 := GenS5{F0: -388, F1: (^(v25)), F2: ((v3 ^ 42) > v3)}
                    _ = v28
                default:
                    var v29 uint = (v15 + 0)
                    _ = v29
                }
            } else {
                v24 = 0
                fmt.Println(v25, v3, v27)
            }
            if ((v15 % ((951) | 1)) > (^(v15))) {
                v25 = (v25 >> ((v15) & 63))
                v30 := GenS1{F0: -765}
                var v31 float32 = float32(0.0)
                _ = v30
                _ = v31
            }
            _ = v27
        }
    }
    v4 = append(v4, v24)
    for fi32 := 0; fi32 < 12; fi32++ {
        var v33 uint8 = 2
        if (42 <= v33) {
            var v34 int32 = -2147483648
            fmt.Println(v8, v7)
            v25 -= (v25 >> (((v15 / ((4294967295) | 1))) & 63))
            _ = v34
        } else {
            if v8 {
                v15 = v15
                v15 -= (v15 >> (((v15 << (((v15 << (((^(v15))) & 63))) & 63))) & 63))
            }
        }
        var v35 uint64 = v7
        _ = fi32
        _ = v33
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
