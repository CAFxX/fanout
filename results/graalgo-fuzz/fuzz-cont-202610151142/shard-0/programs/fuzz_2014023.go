package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int }
func genFn0(p0 int64) string {
    p0 += (p0 & (p0 | (p0 % ((-1) | 1))))
    v4 := []int{9223372036854775807}
    fmt.Println(p0)
    var v5 int64 = -138
    _ = p0
    _ = v4
    _ = v5
    return "foo bar"
}
func genFn1(p0 uint8) uint64 {
    var v6 float32 = (3.14159 * 0.1)
    v7 := []int{9223372036854775807, 256, 993, -435}
    var v8 float64 = (((1e10 * 1e10) + 0.0) * 0.0)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int32 = 0
    var v2 uint8 = 214
    v3 := make(map[uint32]int)
    v2 = 0
    switch "hello" {
    case "hello":
        switch -98 {
        case 32768:
            if (-18 > 0) {
                if (4294967295 > genFn1(128)) {
                    var v9 uint32 = 2147483647
                    var v10 int64 = 287
                    _ = v10
                    _ = v9
                }
                fmt.Println(v1, v1)
                v11 := GenS1{F0: 251, F1: 557, F2: 4294967295}
                _ = v11
            } else {
                var v12 int32 = v1
                v13 := GenS2{F0: -353, F1: -363}
                v14 := v3[2147483648]
                _ = v12
                _ = v13
                _ = v14
            }
        default:
            delete(v3, 65)
        }
    }
    switch 0 {
    case -1:
        var v15 int64 = 9223372036854775807
        _ = v15
    case 1:
        switch -1 {
        case -9223372036854775808:
            v16 := GenS1{F0: -478, F1: 182, F2: 255}
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 string = "foo bar"
        var v19 uint16 = 86
        fmt.Println(v2, v19)
        _ = fi17
        _ = v18
        _ = v19
    }
    if false {
        v2 = v2
        for fi20 := 0; fi20 < 1; fi20++ {
            if (2147483647 >= 2147483648) {
                var v21 int = (fi20 | int(v1))
                _ = v21
            } else {
                var v22 uint = 824
                switch fi20 {
                case 397:
                    switch -2147483648 {
                    case 2147483647:
                        fmt.Println(v22, fi20, v2)
                    case 967:
                        var v23 string = "foo bar"
                        _ = v23
                    case -2147483648:
                        fmt.Println(fi20, fi20)
                    }
                default:
                    var v24 bool = true
                    _ = v24
                }
                _ = v22
            }
            _ = fi20
        }
        v3[164] = 9223372036854775807
    }
    v2 = v2
    for fi25 := 0; fi25 < 6; fi25++ {
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 uint16 = 255
            v1 = (v1 + v1)
            var v28 bool = true
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    if (42 < -32768) {
        var v29 uint64 = 2
        _ = v29
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        v31 := GenS3{F0: -9223372036854775808}
        for fi32 := 0; fi32 < 4; fi32++ {
            v33 := GenS4{F0: 404, F1: uint64(v2), F2: fi30}
            _ = fi32
            _ = v33
        }
        _ = fi30
        _ = v31
    }
    v1 = v1
    v1 = -296
    var v34 int64 = int64(v2)
    _ = len(v3)
    for fi35 := 0; fi35 < 18; fi35++ {
        v2 = (v2 << ((1) & 63))
        var v36 int8 = 42
        if false {
            v2 = (v2 << ((176) & 63))
            fmt.Println(v36)
            v37 := GenS3{F0: v34}
            _ = v37
        }
        _ = fi35
        _ = v36
    }
    v38 := GenS2{F0: v1, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
}
