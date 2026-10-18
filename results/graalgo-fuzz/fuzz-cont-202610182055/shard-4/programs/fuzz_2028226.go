package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint, p1 int32) bool {
    if (true || true) {
        p0 = (p0 + 659)
        p1 = (^(p1))
        if true {
            p1 = p1
        } else {
            if (!false) {
                var v5 int64 = 626
                v6 := GenS1{F0: 255}
                _ = v5
                _ = v6
            }
            v7 := GenS2{F0: uint8(p0)}
            var v8 uint = p0
            _ = v7
            _ = v8
        }
    } else {
        var v9 float64 = (0.0 * (((float64(1e10) / float64(0.1)) * 0.0) + 3.14159))
        if (-1 < 1) {
            v9 = (0.1 / (v9 * (v9 - 0.0)))
        }
        _ = v9
    }
    var v10 uint16 = 255
    v11 := GenS3{F0: (-(p1))}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return (294 == 0)
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{}
    v3 := []int{}
    var v4 uint32 = 185
    var v12 float32 = (((0.1 * ((3.14159 + ((1.5 / (float32(1.5) - 0.0)) - 0.1)) + 0.1)) + 0.1) + float32(3.14159))
    var v13 uint64 = 42
    fmt.Println(v12)
    var v14 float64 = 1.5
    for fi15 := 0; fi15 < 2; fi15++ {
        v1[fi15] = ((3.14159 - (v12 / v12)) > (((3.14159 + v12) - v12) - float32(-2.25)))
        if true {
            var v16 float32 = float32(v14)
            v13 -= v13
            _ = v16
        } else {
            var v17 int16 = 1000
            var v18 uint16 = 255
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    if false {
        var v19 uint = 0
        _ = v19
    } else {
        _ = len(v2)
        for fi20 := 0; fi20 < 17; fi20++ {
            for fi21 := 0; fi21 < 4; fi21++ {
                if (736 < 1) {
                    var v22 int32 = 2147483647
                    v23 := GenS4{F0: -365}
                    _ = v22
                    _ = v23
                }
                if (42 <= 1) {
                    switch fi20 {
                    case -415:
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        _ = len(v1)
                    case -562:
                        switch (fi20 * fi21) {
                        case -2147483648:
                            v3 = append(v3, -2147483648)
                        case 32767:
                            v24 := GenS2{F0: 42}
                            _ = v24
                        default:
                            var v25 string = ("\n\t" + (("" + "\n\t") + ("hello" + ("a" + ((("" + "foo bar") + "") + ("\n\t" + "foo bar"))))))
                            v26 := GenS5{F0: ((v25 + v25) + v25), F1: (v25 + ((v25 + v25) + "hello"))}
                            var v27 float64 = v14
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    var v28 int = (fi21 & fi20)
                    if (true || true) {
                        var v29 int = fi21
                        _ = v29
                    } else {
                        var v30 float32 = 1e10
                        _ = v30
                    }
                    _ = v28
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    switch (fi20 & 0) {
                    case -832:
                        var v31 bool = true
                        _ = v31
                    case -9223372036854775808:
                        var v32 float64 = float64((v12 / v12))
                        _ = v32
                    }
                }
                if (!(false || (61 != 42))) {
                    v33 := v1[(fi21 ^ 1000)]
                    fmt.Println(v12, v12, v12)
                    _ = v33
                } else {
                    var v34 uint16 = 65535
                    v14 = ((0.1 + ((float64(float32(v14)) * v14) + v14)) - (1.5 / v14))
                    _ = v34
                }
                _ = fi21
            }
            _ = fi20
        }
        v35 := GenS6{F0: 0}
        _ = v35
    }
    v2 = append(v2, "\n\t")
    v12 -= (v12 - v12)
    v12 -= 3.14159
    if true {
        v12 = -2.25
        var v36 uint64 = v13
        v12 -= ((0.1 - 3.14159) + v12)
        _ = v36
    }
    v14 += v14
    var v37 int = 32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
