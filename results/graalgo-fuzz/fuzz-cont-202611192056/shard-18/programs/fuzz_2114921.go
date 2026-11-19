package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int16; F1 string; F2 int64 }
func genFn0(p0 uint64) bool {
    var v6 int8 = -1
    switch 32768 {
    case 1:
        v7 := []int{-510, 42, 167, -1}
        _ = v7
    case 42:
        p0 += 1
    }
    _ = p0
    _ = v6
    return (0 < 46)
}
func genFn1(p0 uint8) int8 {
    var v8 int16 = 32767
    if (v8 <= (v8 % ((v8) | 1))) {
        p0 -= p0
        v9 := []int{9223372036854775807, 9223372036854775807, 637}
        _ = v9
    }
    v10 := []int{32767, 128, -453}
    v8 += v8
    _ = p0
    _ = v10
    _ = v8
    return -102
}

func main() {
    var v1 bool = (!true)
    v2 := []string{}
    v3 := []int{-2147483648, 42, 4294967295}
    var v4 float32 = ((-2.25 * 1e10) / (1.5 - 3.14159))
    v5 := make(map[uint32]uint32)
    var v11 uint32 = 42
    v12 := GenS1{F0: -915, F1: 787, F2: uint(v11)}
    v13 := GenS1{F0: 2147483647, F1: 2, F2: 938}
    if (383 < 18446744073709551615) {
        v2 = append(v2, ("\n\t" + ("a" + "foo bar")))
        if v1 {
            var v14 float64 = 1.5
            for fi15 := 0; fi15 < 7; fi15++ {
                v1 = v1
                var v16 float64 = 3.14159
                _ = fi15
                _ = v16
            }
            fmt.Println(v14, v11, v4)
            _ = v14
        }
    }
    var v17 string = "a"
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 uint32 = v11
            var v21 float64 = 3.14159
            var v22 float64 = ((v21 * float64(0.0)) - (v21 * float64((v4 * v4))))
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        switch (v17 + v17) {
        case "\n\t":
            var v23 int32 = 1
            _ = v23
        case "":
            var v24 int32 = 0
            _ = v24
        default:
            var v25 int64 = -1
            _ = v25
        }
        v11 = v11
        _ = fi18
    }
    fmt.Println(v4, v11)
    switch 32768 {
    case 1:
        var v26 float64 = 1.5
        _ = v26
    case 128:
        var v27 int64 = 154
        _ = v27
    default:
        var v28 uint64 = 941
        _ = v28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 float64 = 0.1
        _ = fi29
        _ = v30
    }
    v4 = v4
    v31 := GenS1{F0: -806, F1: 32767, F2: 2}
    v32 := GenS2{F0: -358, F1: ((v17 + v17) + v17)}
    fmt.Println(v11, v1, v11)
    v17 = v17
    if (v4 == (0.1 * ((-2.25 / (v4 / v4)) / float32((0.0 / float64(((1e10 / 1e10) - (v4 * 0.0)))))))) {
        fmt.Println(v17)
        for fi33 := 0; fi33 < 12; fi33++ {
            if ((0 != 25) || v1) {
                if v1 {
                    v34 := GenS2{F0: 187, F1: v17}
                    _ = v34
                }
                for fi35 := 0; fi35 < 3; fi35++ {
                    var v36 float32 = -2.25
                    var v37 int8 = 5
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            } else {
                v38 := GenS2{F0: -32768, F1: "a"}
                _ = v38
            }
            v39 := GenS3{F0: -32768, F1: ("foo bar" + v17), F2: 42}
            _ = fi33
            _ = v39
        }
    } else {
        v40 := GenS3{F0: int16(v11), F1: v17, F2: 1}
        v11 -= (v11 >> ((889) & 63))
        v41 := GenS3{F0: 1, F1: v17, F2: 1}
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
