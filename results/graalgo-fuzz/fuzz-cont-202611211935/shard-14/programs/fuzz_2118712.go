package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int32; F2 int16 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint; F1 int16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    var v4 uint32 = 2147483647
    var v5 int32 = -713
    var v6 float64 = 0.0
    var v7 uint8 = uint8(v4)
    var v8 int = 4294967295
    v7 = (v7 + 42)
    switch "" {
    case "a":
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS1{F0: v7}
            if (1000 <= int16(v4)) {
                for fi11 := 0; fi11 < 10; fi11++ {
                    v12 := v2[32767]
                    var v13 uint32 = (v4 % ((v12) | 1))
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                var v14 uint64 = 4294967295
                var v15 int16 = 317
                _ = v14
                _ = v15
            } else {
                v16 := GenS2{F0: v7, F1: v5, F2: 1000}
                v7 = v7
                v17 := GenS2{F0: v7, F1: v5, F2: 32767}
                _ = v16
                _ = v17
            }
            if false {
                v7 = 2
            }
            _ = fi9
            _ = v10
        }
    case "foo bar":
        var v18 float64 = (0.0 - 0.0)
        _ = v18
    case "hello":
        fmt.Println(v7, v8, v7)
    }
    if false {
        var v19 int32 = v5
        var v20 float64 = (v6 / (float64(((-2.25 + 3.14159) / 1.0)) * v6))
        v5 = (v5 << ((2) & 63))
        _ = v19
        _ = v20
    } else {
        v4 = (v4 + 2147483647)
        switch 127 {
        case 9223372036854775807:
            v4 -= v4
        case 42:
            if true {
                v2[(v8 >> ((255) & 63))] = v4
                v7 += v7
            }
        case 151:
            switch v8 {
            case 2147483647:
                var v21 float64 = (0.1 + -2.25)
                _ = v21
            case 42:
                v6 -= v6
            }
        default:
            switch (v8 * 42) {
            case 2147483647:
                v6 = float64((((1.5 * 0.1) - ((((1e10 - -2.25) * (((1.5 * (0.0 * ((3.14159 + (((1.5 * 0.1) - (0.0 - (3.14159 + 1e10))) + 1.5)) * 3.14159))) - -2.25) * 0.1)) * ((-2.25 * 1e10) / ((float32(float64(1.5)) / (float32((0.1 / (((float64((3.14159 / 1e10)) * 3.14159) + v6) / v6))) - 0.1)) / (-2.25 / float32(v6))))) - 0.0)) - 0.0))
            case 127:
                v22 := []int{}
                _ = v22
            case 32768:
                fmt.Println(v7, v4)
            }
        }
    }
    if (42 >= 18446744073709551615) {
        v23 := GenS3{F0: 102, F1: 0}
        var v24 uint64 = 42
        for fi25 := 0; fi25 < 9; fi25++ {
            for fi26 := 0; fi26 < 12; fi26++ {
                var v27 float64 = (((((float64(0.0) / (v6 - v6)) + v6) - ((v6 / (v6 / 3.14159)) * v6)) * float64(3.14159)) - -2.25)
                v5 -= int32(v7)
                if false {
                    if (-343 == int64(fi25)) {
                        v1[fi26] = (v7 < (^(v7)))
                        v28 := GenS2{F0: v7, F1: 0, F2: -32768}
                        _ = v28
                    }
                    v7 = 1
                    v29 := GenS4{F0: uint(v4), F1: 886}
                    _ = v29
                } else {
                    v27 = float64(1.5)
                    var v30 float32 = (3.14159 / 1e10)
                    _ = v30
                }
                _ = fi26
                _ = v27
            }
            var v31 int32 = 42
            _ = fi25
            _ = v31
        }
        _ = v23
        _ = v24
    } else {
        if false {
            switch 32767 {
            case -2147483648:
                fmt.Println(v5, v8, v5)
            case 518:
                v7 = v7
            }
            v7 -= (v7 ^ v7)
        }
        v32 := GenS3{F0: -32768, F1: v5}
        _ = v32
    }
    var v33 int = v8
    v8 += -9223372036854775808
    v2[(-(v33))] = (^(v4))
    for fi34 := 0; fi34 < 20; fi34++ {
        v8 += (v33 << ((0) & 63))
        var v35 float32 = 3.14159
        _ = len(v3)
        _ = fi34
        _ = v35
    }
    delete(v3, "a")
    v36 := GenS3{F0: 251, F1: v5}
    switch (v33 ^ v33) {
    case 32767:
        var v37 int = v33
        _ = v37
    }
    v38 := GenS1{F0: (v7 >> ((2) & 63))}
    v39 := GenS4{F0: 988, F1: 425}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
