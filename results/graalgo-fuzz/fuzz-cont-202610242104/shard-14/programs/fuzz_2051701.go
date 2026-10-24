package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 uint8; F1 int16; F2 bool }
func genFn0(p0 int32, p1 uint64) int32 {
    fmt.Println(p0, p0)
    if (-9223372036854775808 != -607) {
        if (!true) {
            var v4 uint16 = 1
            _ = v4
        } else {
            v5 := GenS1{F0: false}
            if false {
                switch "hello" {
                case "":
                    fmt.Println(p1, p0)
                }
                for fi6 := 0; fi6 < 18; fi6++ {
                    fmt.Println(p1, p1, fi6)
                    _ = fi6
                }
                p0 = (p0 + 902)
            } else {
                for fi7 := 0; fi7 < 18; fi7++ {
                    v8 := []int{}
                    var v9 int32 = p0
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                v10 := GenS1{F0: false}
                _ = v10
            }
            _ = v5
        }
        if (!((0.0 / (((1e10 + 3.14159) / (3.14159 * 1.5)) * (((-2.25 + float64(0.1)) - (((0.0 - ((-2.25 * 1.5) * (1.5 * 0.1))) + 3.14159) * 1e10)) / 3.14159))) == 0.0)) {
            var v11 uint8 = 42
            if (!(true || (false || (!false)))) {
                if (-32768 == 51) {
                    v11 = (v11 * 128)
                    for fi12 := 0; fi12 < 12; fi12++ {
                        switch "hello" {
                        case "foo bar":
                            v13 := GenS1{F0: false}
                            _ = v13
                        case "a":
                            v14 := []int{-9223372036854775808, (-(fi12))}
                            _ = v14
                        }
                        v15 := GenS1{F0: true}
                        v16 := GenS1{F0: true}
                        _ = fi12
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(p0)
                }
                if (!(-1 != (p0 << ((65535) & 63)))) {
                    for fi17 := 0; fi17 < 5; fi17++ {
                        v11 -= v11
                        _ = fi17
                    }
                    v18 := GenS2{F0: (v11 | (v11 | v11)), F1: -1, F2: 255}
                    var v19 string = ((("" + "foo bar") + ("a" + "")) + "hello")
                    _ = v18
                    _ = v19
                }
                var v20 float64 = (((1.5 - 0.0) * float64(0.1)) - 1e10)
                _ = v20
            } else {
                var v21 int32 = p0
                _ = v21
            }
            _ = v11
        }
    } else {
        var v22 int = -202
        _ = v22
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{0, -1, 32768, 127}
    var v3 string = (("a" + "a") + ("foo bar" + ""))
    for fi23 := 0; fi23 < 13; fi23++ {
        fmt.Println(v3, v3)
        _ = fi23
    }
    _ = len(v1)
    v3 = (v3 + "")
    var v24 uint8 = 124
    if true {
        v3 = "\n\t"
        var v25 int64 = 0
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        v27 := v1[2147483647]
        v28 := GenS3{F0: 32767, F1: v24}
        _ = fi26
        _ = v27
        _ = v28
    }
    switch v3 {
    case "a":
        fmt.Println(v3, v24)
    case "":
        if false {
            v3 = v3
            switch v3 {
            case "\n\t":
                if len(v2) > 0 {
                    v2[0] = -9223372036854775808
                }
            case "hello":
                _ = len(v2)
            case "foo bar":
                v24 -= v24
            default:
                delete(v1, 672)
            }
            delete(v1, 166)
        } else {
            var v29 uint = 42
            var v30 uint64 = 42
            _ = v29
            _ = v30
        }
    case "hello":
        var v31 float64 = 1.5
        _ = v31
    }
    v24 += 255
    var v32 string = v3
    var v33 bool = false
    for fi34 := 0; fi34 < 9; fi34++ {
        fmt.Println(fi34, v24, v24)
        for fi35 := 0; fi35 < 8; fi35++ {
            fmt.Println(v3, fi34, v24)
            _ = len(v2)
            _ = fi35
        }
        var v36 uint8 = 0
        _ = fi34
        _ = v36
    }
    var v37 float32 = (1e10 - 0.0)
    v38 := GenS4{F0: 32767, F1: v33}
    v39 := GenS5{F0: (v24 <= (v24 >> ((65535) & 63))), F1: -128}
    v40 := GenS6{F0: v24, F1: 32767, F2: (0 <= 65535)}
    var v41 float32 = ((0.0 / ((v37 + 0.0) * v37)) + v37)
    for fi42 := 0; fi42 < 11; fi42++ {
        v2 = append(v2, fi42)
        v3 = ((v3 + v32) + "a")
        v41 -= v41
        _ = fi42
    }
    var v43 uint16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    _ = v43
}
