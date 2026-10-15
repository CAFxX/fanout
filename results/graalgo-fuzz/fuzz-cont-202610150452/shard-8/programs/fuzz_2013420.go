package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8; F1 int64; F2 int8 }
func genFn0(p0 uint8) uint {
    fmt.Println(p0)
    if (-9223372036854775808 > 303) {
        var v7 int = -92
        var v8 uint8 = 151
        _ = v7
        _ = v8
    } else {
        var v9 bool = true
        _ = v9
    }
    switch 128 {
    case -1:
        v10 := GenS1{F0: -706, F1: 0, F2: 0}
        _ = v10
    case 128:
        fmt.Println(p0)
    case 261:
        p0 -= (p0 & p0)
    }
    switch 4294967295 {
    case -159:
        var v11 int8 = int8(p0)
        _ = v11
    case 128:
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 int = (^(fi12))
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 int64 = 9223372036854775807
                _ = fi14
                _ = v15
            }
            if (-635 != 0) {
                v13 += fi12
                var v16 float32 = -2.25
                _ = v16
            }
            _ = fi12
            _ = v13
        }
    case 120:
        if false {
            for fi17 := 0; fi17 < 11; fi17++ {
                for fi18 := 0; fi18 < 11; fi18++ {
                    v19 := GenS1{F0: -773, F1: 2147483648, F2: 42}
                    v20 := []int{}
                    fmt.Println(p0, p0, p0)
                    fmt.Println(fi18, fi18)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                switch "foo bar" {
                case "a":
                    fmt.Println(fi17, p0)
                case "\n\t":
                    fmt.Println(p0, p0)
                case "foo bar":
                    fmt.Println(p0)
                }
                _ = fi17
            }
            if true {
                for fi21 := 0; fi21 < 18; fi21++ {
                    v22 := []int{127, (fi21 << ((752) & 63))}
                    switch "\n\t" {
                    case "a":
                        for fi23 := 0; fi23 < 14; fi23++ {
                            var v24 uint = 42
                            v24 = (v24 | 2)
                            _ = fi23
                            _ = v24
                        }
                    default:
                        var v25 int8 = 22
                        _ = v25
                    }
                    if len(v22) > 0 {
                        v22[0] = fi21
                    }
                    _ = fi21
                    _ = v22
                }
            }
        } else {
            if ((("a" + ("hello" + ("foo bar" + "\n\t"))) + "foo bar") > "") {
                var v26 int64 = 0
                fmt.Println(v26, p0)
                _ = v26
            }
        }
    }
    p0 = (p0 | 128)
    _ = p0
    return 1
}
func genFn1(p0 int8, p1 uint8) uint64 {
    v27 := GenS1{F0: -302, F1: 225, F2: 1}
    v28 := []int{-1, 1, -1}
    fmt.Println(p0, p1)
    if true {
        _ = len(v28)
        fmt.Println(p0, p0)
        var v29 int8 = p0
        _ = v29
    } else {
        v30 := GenS1{F0: 2147483647, F1: 2, F2: -2147483648}
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    return 580
}

func main() {
    v1 := []int{-938, 850, -2147483648}
    v2 := make(map[string]bool)
    var v3 int32 = 478
    var v4 bool = false
    v5 := make(map[int]int)
    var v6 float32 = float32(3.14159)
    delete(v5, 1)
    v31 := GenS1{F0: v3, F1: 573, F2: 1}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 = v4
    v32 := GenS2{F0: 255, F1: 1, F2: -123}
    v33 := GenS1{F0: v3, F1: 812, F2: (v3 ^ (v3 | -970))}
    if v4 {
        v6 = (((v6 + 3.14159) * v6) + float32(float64(((0.1 - (((v6 - v6) - v6) + v6)) / -2.25))))
        var v34 uint64 = 4294967295
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v34
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        switch "" {
        case "\n\t":
            if v4 {
                v6 += (v6 / 1.5)
            }
        case "a":
            _ = len(v1)
        case "":
            var v36 uint8 = 42
            _ = v36
        }
        _ = fi35
    }
    v4 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
