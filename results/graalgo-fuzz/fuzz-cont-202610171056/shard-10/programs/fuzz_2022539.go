package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 uint8; F1 bool }
func genFn0(p0 uint) string {
    p0 -= (p0 / (((p0 | p0)) | 1))
    if false {
        var v4 int32 = 843
        var v5 float32 = -2.25
        _ = v4
        _ = v5
    } else {
        fmt.Println(p0)
        if true {
            p0 = 1
            for fi6 := 0; fi6 < 10; fi6++ {
                if true {
                    fmt.Println(fi6, p0)
                    switch "hello" {
                    case "foo bar":
                        var v7 int = (fi6 * fi6)
                        _ = v7
                    case "hello":
                        v8 := []int{-905}
                        _ = v8
                    }
                } else {
                    fmt.Println(fi6)
                    var v9 bool = (!false)
                    _ = v9
                }
                if (!true) {
                    p0 += (p0 ^ (^(p0)))
                    for fi10 := 0; fi10 < 17; fi10++ {
                        var v11 uint16 = 2
                        fmt.Println(fi10, fi6)
                        _ = fi10
                        _ = v11
                    }
                    switch "" {
                    case "a":
                        v12 := []int{(fi6 | fi6), 42, (fi6 << (((p0 % ((835) | 1))) & 63)), -343}
                        _ = v12
                    }
                } else {
                    p0 = (p0 % ((p0) | 1))
                }
                v13 := GenS1{F0: (true && (42 != 127))}
                _ = fi6
                _ = v13
            }
            var v14 string = ("a" + "hello")
            _ = v14
        }
    }
    var v15 int64 = -123
    _ = p0
    _ = v15
    return ""
}

func main() {
    var v1 int8 = 0
    var v2 uint64 = 2
    v3 := make(map[int]int)
    delete(v3, 42)
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 uint = 4294967295
        var v18 uint32 = 2147483647
        var v19 float64 = 1e10
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    if true {
        var v20 float64 = (0.1 - 3.14159)
        v2 = 4294967295
        _ = v20
    } else {
        var v21 float32 = float32(-2.25)
        _ = v21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        v23 := GenS2{F0: 2147483648}
        var v24 uint8 = 255
        v25 := v3[fi22]
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    v26 := []int{32767, -2147483648}
    var v27 float64 = float64(float32(3.14159))
    if false {
        v1 = (v1 << ((723) & 63))
    } else {
        v28 := GenS3{F0: v2}
        _ = v28
    }
    switch -914 {
    case 42:
        var v29 uint16 = 42
        _ = v29
    case 5:
        v30 := GenS1{F0: (true || true)}
        _ = v30
    default:
        var v31 int16 = 287
        _ = v31
    }
    var v32 int = 32767
    _ = len(v3)
    if (!false) {
        for fi33 := 0; fi33 < 6; fi33++ {
            var v34 int32 = 42
            v34 += v34
            _ = fi33
            _ = v34
        }
        for fi35 := 0; fi35 < 19; fi35++ {
            for fi36 := 0; fi36 < 2; fi36++ {
                v37 := v3[(v32 >> ((642) & 63))]
                fmt.Println(v37, v32, v1)
                _ = fi36
                _ = v37
            }
            v27 += v27
            v38 := GenS2{F0: 0}
            _ = fi35
            _ = v38
        }
        if false {
            var v39 uint8 = 0
            v32 -= v32
            for fi40 := 0; fi40 < 14; fi40++ {
                _ = len(v26)
                _ = fi40
            }
            _ = v39
        } else {
            v27 = 1.5
        }
    } else {
        var v41 bool = true
        for fi42 := 0; fi42 < 16; fi42++ {
            v26 = append(v26, (fi42 - fi42))
            _ = fi42
        }
        v43 := GenS4{F0: 156, F1: 32767}
        _ = v41
        _ = v43
    }
    v26 = append(v26, v32)
    if ((-820 > 629) && (!true)) {
        _ = len(v3)
        var v44 int32 = 780
        var v45 string = "foo bar"
        _ = v44
        _ = v45
    }
    fmt.Println(v27, v32, v32)
    v27 = (((-2.25 + v27) * (1.5 + float64((-2.25 / ((((0.0 * 1.5) + (0.0 + 1.5)) * -2.25) + 0.0))))) - (v27 / v27))
    v46 := GenS5{F0: 0, F1: (!((1 >= 42) || true))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v46
}
