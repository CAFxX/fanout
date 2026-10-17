package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 string; F1 int16; F2 uint }
func genFn0(p0 int64) uint32 {
    v7 := []int{-601, 9223372036854775807, 128}
    v8 := GenS1{F0: 434, F1: 128}
    if (false || true) {
        p0 -= (p0 + (p0 ^ (p0 << ((0) & 63))))
    }
    p0 = p0
    _ = p0
    _ = v7
    _ = v8
    return 583
}
func genFn1() uint8 {
    v9 := GenS2{F0: 42, F1: 621}
    var v10 int16 = 0
    for fi11 := 0; fi11 < 16; fi11++ {
        v10 -= 463
        _ = fi11
    }
    fmt.Println(v10)
    v10 += (^(v10))
    _ = v10
    _ = v9
    return 247
}

func main() {
    var v1 int32 = 1
    var v2 uint16 = 255
    var v3 int = -9223372036854775808
    var v4 float64 = 1e10
    v5 := []uint32{177, 42, 65535, 4294967295}
    var v6 float32 = float32(v4)
    fmt.Println(v6, v2, v4)
    switch ("hello" + "hello") {
    case "a":
        var v12 int32 = -1
        _ = v12
    default:
        v13 := GenS3{F0: "", F1: -441, F2: 42}
        _ = v13
    }
    if false {
        var v14 uint = 2
        _ = v14
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v15 float32 = float32(3.14159)
    v4 += ((v4 / 0.1) - v4)
    for fi16 := 0; fi16 < 15; fi16++ {
        fmt.Println(v6, v3)
        v6 = v15
        _ = fi16
    }
    switch (v3 * 914) {
    case -180:
        for fi17 := 0; fi17 < 11; fi17++ {
            fmt.Println(fi17)
            v4 = v4
            _ = fi17
        }
    case 965:
        _ = len(v5)
    case 128:
        switch ("foo bar" + "hello") {
        case "hello":
            var v18 int64 = 42
            _ = v18
        case "":
            var v19 uint16 = 561
            _ = v19
        default:
            var v20 uint64 = 930
            _ = v20
        }
    }
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v21 int = (v3 / ((4294967295) | 1))
    v6 = (((v6 * ((v15 * (v15 / float32(v4))) * v15)) / float32((((3.14159 / v4) - v4) / 1.5))) * v15)
    switch "\n\t" {
    case "hello":
        v3 = -52
    case "":
        v2 = 32767
    default:
        for fi22 := 0; fi22 < 14; fi22++ {
            switch v3 {
            case 4294967295:
                for fi23 := 0; fi23 < 2; fi23++ {
                    v6 -= (((((float32(float64(v6)) - 3.14159) * float32((v4 - (float64(v6) - 1.5)))) / 1e10) + ((0.0 / 3.14159) / float32(float64(v15)))) / 0.1)
                    var v24 int64 = -9223372036854775808
                    _ = fi23
                    _ = v24
                }
            case 36:
                if false {
                    var v25 uint16 = v2
                    _ = v25
                } else {
                    v3 = fi22
                    if ((0.1 - v4) >= v4) {
                        if len(v5) > 0 {
                            v5[0] = 97
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
            case 32768:
                switch "" {
                case "foo bar":
                    switch v3 {
                    case -1:
                        var v26 float64 = v4
                        _ = v26
                    }
                }
            }
            v21 -= (v3 << ((0) & 63))
            v27 := GenS2{F0: -579, F1: 1}
            _ = fi22
            _ = v27
        }
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
