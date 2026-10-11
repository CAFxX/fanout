package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int8 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 string; F1 int64; F2 uint64 }
func genFn0(p0 int8) uint {
    p0 -= (^(p0))
    p0 = (p0 << ((42) & 63))
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 float64 = 3.14159
        v8 := []int{}
        if (0 < 42) {
            if (749 == (fi6 ^ (fi6 | (fi6 * fi6)))) {
                p0 += (p0 * (p0 + -128))
            }
            v7 = (v7 / 1.5)
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 float64 = float64(-2.25)
                switch "" {
                case "hello":
                    v8 = append(v8, fi6)
                case "\n\t":
                    _ = len(v8)
                default:
                    v7 -= v10
                }
                _ = fi9
                _ = v10
            }
        } else {
            var v11 bool = (1 < 255)
            var v12 bool = (0 <= 42)
            v13 := GenS1{F0: 1, F1: (4294967295 <= 1), F2: (p0 << ((458) & 63))}
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    return 893
}

func main() {
    var v1 string = ((("a" + "foo bar") + "") + "")
    var v2 uint = 4294967295
    var v3 uint32 = 828
    v4 := []uint32{v3, (v3 * (^(v3))), 2147483647, v3}
    var v5 int = 1
    for fi14 := 0; fi14 < 13; fi14++ {
        if (9223372036854775807 <= -1) {
            if false {
                var v15 float32 = -2.25
                fmt.Println(v2, v5)
                _ = v15
            } else {
                v3 = v3
            }
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            if (!false) {
                v1 = (v1 + (v1 + (v1 + (v1 + v1))))
                fmt.Println(v2, fi14, fi16)
            } else {
                switch int(v3) {
                case 2147483647:
                    fmt.Println(v1, fi16, v5)
                case -9223372036854775808:
                    if len(v4) > 0 {
                        v4[0] = 437
                    }
                }
                var v17 uint8 = 127
                var v18 uint8 = (v17 ^ (^(v17)))
                _ = v17
                _ = v18
            }
            v1 = ("a" + v1)
            var v19 float64 = -2.25
            _ = fi16
            _ = v19
        }
        _ = fi14
    }
    if len(v4) > 0 {
        v4[0] = (v3 / (((v3 / (((^(v3))) | 1))) | 1))
    }
    v2 = (v2 ^ (v2 | v2))
    if (128 != 0) {
        switch (v5 * (^(v5))) {
        case 878:
            v4 = append(v4, (v3 % ((v3) | 1)))
        case 189:
            var v20 int16 = -233
            _ = v20
        default:
            v21 := GenS2{F0: -1, F1: -32768}
            _ = v21
        }
    }
    v22 := GenS1{F0: uint8(v2), F1: (1000 == -100), F2: -128}
    for fi23 := 0; fi23 < 3; fi23++ {
        if (-1 != -845) {
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
        for fi24 := 0; fi24 < 14; fi24++ {
            v25 := GenS2{F0: -679, F1: -32768}
            var v26 int16 = 966
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi23
    }
    var v27 float32 = 3.14159
    if false {
        if true {
            var v28 string = "\n\t"
            v29 := GenS3{F0: v1, F1: -838, F2: 2}
            _ = v28
            _ = v29
        }
        var v30 float64 = 1e10
        switch (v1 + v1) {
        case "hello":
            var v31 string = v1
            _ = v31
        }
        _ = v30
    } else {
        var v32 int32 = 0
        for fi33 := 0; fi33 < 13; fi33++ {
            var v34 string = v1
            _ = fi33
            _ = v34
        }
        _ = v32
    }
    v4 = append(v4, 42)
    var v35 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
