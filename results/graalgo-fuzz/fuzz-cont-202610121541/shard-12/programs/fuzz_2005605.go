package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int32; F2 int64 }
type GenS5 struct { F0 uint; F1 int64 }
func genFn0(p0 int16) uint64 {
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = (p0 % (((p0 / (((p0 >> ((65535) & 63))) | 1))) | 1))
    v5 := GenS1{F0: 0, F1: (-2147483648 <= 1)}
    var v6 string = "\n\t"
    _ = p0
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint16, p1 int8) uint32 {
    var v7 uint16 = (p0 >> ((2) & 63))
    v8 := []int{-2147483648}
    if true {
        if true {
            var v9 uint64 = 4294967295
            if ((p0 * v7) > 729) {
                var v10 bool = false
                var v11 int32 = 0
                _ = v10
                _ = v11
            } else {
                if len(v8) > 0 {
                    v8[0] = -9223372036854775808
                }
                var v12 bool = (!true)
                _ = v12
            }
            _ = v9
        } else {
            var v13 int64 = 42
            var v14 int8 = (p1 * 42)
            _ = v13
            _ = v14
        }
    }
    p1 -= p1
    v15 := GenS2{F0: 4294967295, F1: 2}
    _ = p0
    _ = p1
    _ = v15
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    var v1 int32 = -1
    var v2 uint16 = 42
    v3 := []uint32{351, 1, 2147483647, 42}
    var v4 float64 = float64(float32((3.14159 + 1.5)))
    if (0 > v2) {
        fmt.Println(v4)
    } else {
        switch 9223372036854775807 {
        case 964:
            v1 = 42
        case 42:
            var v16 bool = (!(-659 > 276))
            _ = v16
        default:
            v3 = append(v3, 42)
        }
    }
    var v17 uint = 42
    var v18 int32 = -426
    var v19 int32 = v18
    for fi20 := 0; fi20 < 17; fi20++ {
        if len(v3) > 0 {
            v3[0] = 931
        }
        _ = fi20
    }
    switch 119 {
    case 0:
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int32 = (v18 * v18)
            _ = fi21
            _ = v22
        }
    case 2147483647:
        var v23 float64 = (((v4 + v4) / 1.5) / 1e10)
        _ = v23
    case -1:
        fmt.Println(v1, v2, v4)
    default:
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 string = ("\n\t" + ("hello" + ("foo bar" + "\n\t")))
            _ = fi24
            _ = v25
        }
    }
    var v26 int = -2147483648
    if (v2 <= v2) {
        var v27 float64 = 1e10
        if (0.1 != 1.5) {
            _ = len(v3)
        } else {
            var v28 uint = 255
            switch "\n\t" {
            case "foo bar":
                _ = len(v3)
            default:
                v29 := GenS2{F0: 2, F1: 255}
                _ = v29
            }
            _ = v28
        }
        _ = v27
    } else {
        switch ("\n\t" + "foo bar") {
        case "":
            _ = len(v3)
        default:
            v1 = v18
        }
        if false {
            v17 += v17
            if false {
                v30 := GenS3{F0: (v17 | (v17 % (((v17 / ((2) | 1))) | 1)))}
                _ = v30
            }
            var v31 float64 = v4
            _ = v31
        } else {
            v1 += v18
            var v32 uint8 = 255
            _ = v32
        }
    }
    v33 := GenS4{F0: false, F1: -2147483648, F2: 0}
    _ = len(v3)
    for fi34 := 0; fi34 < 3; fi34++ {
        for fi35 := 0; fi35 < 1; fi35++ {
            var v36 uint8 = 127
            if len(v3) > 0 {
                v3[0] = 658
            }
            if (v17 >= (v17 * (v17 - (v17 * 4294967295)))) {
                v2 -= (v2 * v2)
            }
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    var v37 int32 = (v1 - v1)
    if true {
        v38 := GenS5{F0: v17, F1: int64(v2)}
        _ = v38
    }
    if (1 <= 143) {
        var v39 int16 = 42
        if len(v3) > 0 {
            v3[0] = 2
        }
        if true {
            v39 = (v39 % ((v39) | 1))
        } else {
            v37 -= v18
            for fi40 := 0; fi40 < 6; fi40++ {
                v41 := GenS2{F0: 2, F1: 255}
                fmt.Println(v19, v26, v4)
                _ = fi40
                _ = v41
            }
        }
        _ = v39
    } else {
        var v42 float32 = 0.1
        _ = v42
    }
    var v43 int = v26
    v1 += 42
    var v44 uint32 = 42
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v45 := GenS3{F0: (v17 ^ 42)}
        _ = v45
    case "hello":
        v46 := GenS2{F0: 0, F1: 42}
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v43
    _ = v44
}
