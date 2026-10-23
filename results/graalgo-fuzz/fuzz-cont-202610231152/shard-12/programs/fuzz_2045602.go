package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 int32; F2 uint32 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 bool, p1 int8, p2 bool) uint {
    var v6 uint32 = 4294967295
    switch 4294967295 {
    case -2147483648:
        p0 = (127 <= -680)
    case 42:
        v7 := GenS1{F0: p1, F1: 9223372036854775807}
        _ = v7
    case -182:
        switch 520 {
        case -9223372036854775808:
            var v8 int32 = 0
            _ = v8
        case 128:
            if p0 {
                var v9 uint = 340
                _ = v9
            } else {
                fmt.Println(v6, v6, v6)
                for fi10 := 0; fi10 < 18; fi10++ {
                    switch "foo bar" {
                    case "":
                        p0 = false
                    case "a":
                        for fi11 := 0; fi11 < 9; fi11++ {
                            p0 = (127 > 42)
                            _ = fi11
                        }
                    default:
                        v12 := []int{fi10, 42}
                        _ = v12
                    }
                    _ = fi10
                }
                if (false && (42 > -466)) {
                    v6 -= (v6 % (((v6 | v6)) | 1))
                    fmt.Println(p1, p1)
                    v13 := []int{1}
                    _ = v13
                }
            }
        case 127:
            v14 := GenS2{F0: -2147483648, F1: p1}
            _ = v14
        }
    }
    v15 := []int{525, 106}
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 uint = 640
        v17 += (v17 % (((v17 % ((255) | 1))) | 1))
        _ = len(v15)
        _ = fi16
        _ = v17
    }
    p2 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v6
    return 65535
}
func genFn1() int8 {
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(fi18)
        v19 := GenS1{F0: -62, F1: fi18}
        v20 := []int{}
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 string = ("foo bar" + "hello")
    _ = v21
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 0.0
    v3 := []int{172, -392, 4294967295, 42}
    v4 := []int{1, 128, 450}
    var v5 int16 = 1
    if true {
        var v22 int8 = -128
        v22 -= (v22 * 29)
        v22 = v22
        _ = v22
    }
    v23 := GenS1{F0: -128, F1: 128}
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := GenS2{F0: (fi24 * (fi24 & fi24)), F1: 99}
        v26 := GenS3{F0: 255}
        v2 = (v2 - 1e10)
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int64 = -1
    v28 := GenS2{F0: -2147483648, F1: 1}
    var v29 int32 = int32(v27)
    if true {
        var v30 uint = 42
        _ = v30
    } else {
        switch "" {
        case "hello":
            var v31 int64 = (v27 & (^(v27)))
            _ = v31
        case "foo bar":
            fmt.Println(v27)
        case "\n\t":
            var v32 uint32 = 42
            _ = v32
        default:
            if true {
                fmt.Println(v29, v29)
            } else {
                var v33 int = -9223372036854775808
                switch "a" {
                case "":
                    v34 := GenS4{F0: (true && false), F1: v29, F2: 4294967295}
                    _ = v34
                case "\n\t":
                    v35 := v1["a"]
                    _ = v35
                case "foo bar":
                    if (!true) {
                        if (!true) {
                            var v36 float64 = float64(v2)
                            if false {
                                var v37 uint16 = 1
                                v38 := GenS5{F0: 2147483648}
                                v4 = append(v4, 9223372036854775807)
                                _ = v37
                                _ = v38
                            }
                            _ = v36
                        }
                        var v39 float32 = (((((v2 * v2) / v2) - v2) * 3.14159) - v2)
                        _ = v39
                    } else {
                        v40 := GenS2{F0: -229, F1: -128}
                        _ = v40
                    }
                default:
                    v5 = v5
                }
                v29 += -2147483648
                _ = v33
            }
        }
    }
    v41 := GenS4{F0: (42 != 2), F1: (^(v29)), F2: 2147483648}
    v42 := GenS5{F0: 430}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
