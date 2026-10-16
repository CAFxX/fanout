package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 string; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 int16; F1 int8; F2 int8 }
type GenS7 struct { F0 int64; F1 uint64 }
func genFn0() string {
    var v6 int16 = -486
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(v6, fi7)
        v8 := []int{fi7, (fi7 & fi7), 682}
        _ = fi7
        _ = v8
    }
    var v9 int32 = 2147483647
    v10 := GenS1{F0: 205}
    v11 := GenS2{F0: -1, F1: "a", F2: 220}
    _ = v10
    _ = v11
    _ = v6
    _ = v9
    return "a"
}

func main() {
    var v1 int32 = -93
    v2 := make(map[string]string)
    v3 := []int{-2147483648, 140, 128, -9223372036854775808}
    var v4 int64 = 61
    v5 := make(map[int]bool)
    switch (("foo bar" + "hello") + ("hello" + "foo bar")) {
    case "foo bar":
        var v12 int32 = v1
        _ = v12
    }
    var v13 uint64 = 2
    switch "a" {
    case "\n\t":
        switch -9223372036854775808 {
        case -2147483648:
            if (1 <= 72) {
                var v14 int32 = v1
                var v15 int16 = -952
                if (false && (float64(3.14159) > 1e10)) {
                    var v16 int32 = (v14 | v14)
                    v17 := GenS3{F0: int8(v1)}
                    switch "foo bar" {
                    case "":
                        v18 := GenS4{F0: false}
                        _ = v18
                    case "hello":
                        v4 += v4
                    case "a":
                        var v19 float64 = 3.14159
                        _ = v19
                    }
                    _ = v16
                    _ = v17
                } else {
                    v13 = (v13 + (v13 - (v13 * (v13 * v13))))
                }
                _ = v14
                _ = v15
            }
        case -205:
            switch 1 {
            case 42:
                fmt.Println(v13, v1, v4)
            case -659:
                var v20 float32 = (-2.25 - -2.25)
                _ = v20
            }
        case -235:
            v3 = append(v3, 128)
        default:
            switch -9223372036854775808 {
            case 127:
                v21 := GenS5{F0: 64, F1: 1}
                _ = v21
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    case "":
        fmt.Println(v13, v4)
    case "foo bar":
        fmt.Println(v13, v1, v13)
    }
    v4 -= 72
    switch "a" {
    case "":
        v4 = v4
    }
    if (((100 > -1) || (((9223372036854775807 >= 1) || true) && false)) || (128 != 2)) {
        fmt.Println(v1, v1, v1)
    }
    if true {
        var v22 int16 = 1000
        for fi23 := 0; fi23 < 5; fi23++ {
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 uint64 = (v13 * 537)
                v26 := GenS4{F0: true}
                if (int32(v22) >= (v1 + 42)) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v4, v22)
                }
                _ = fi24
                _ = v25
                _ = v26
            }
            v27 := GenS1{F0: 572}
            _ = fi23
            _ = v27
        }
        var v28 int32 = (v1 + 423)
        _ = v22
        _ = v28
    } else {
        var v29 int32 = v1
        _ = v29
    }
    v4 -= v4
    v1 = v1
    v4 += (v4 % (((v4 * (v4 & (-(v4))))) | 1))
    var v30 float32 = 3.14159
    switch -1 {
    case 2147483647:
        var v31 float32 = (v30 - float32(1e10))
        _ = v31
    default:
        fmt.Println(v30, v1, v13)
    }
    v32 := GenS6{F0: 1000, F1: -128, F2: 0}
    if len(v3) > 0 {
        v3[0] = 0
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        _ = len(v3)
        _ = fi33
    }
    v34 := GenS7{F0: v4, F1: 42}
    if (!true) {
        fmt.Println(v1)
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 uint16 = 638
            fmt.Println(v36)
            var v37 int32 = v1
            _ = fi35
            _ = v36
            _ = v37
        }
    } else {
        var v38 int32 = v1
        _ = v38
    }
    var v39 int8 = 33
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v34
    _ = v39
    _ = v4
    _ = v5
}
