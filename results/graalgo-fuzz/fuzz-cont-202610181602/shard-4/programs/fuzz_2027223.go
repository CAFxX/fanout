package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 bool; F1 int16; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int32; F1 int8; F2 int16 }
type GenS5 struct { F0 string; F1 uint8; F2 uint }
type GenS6 struct { F0 bool; F1 uint16 }
func genFn0() uint16 {
    v7 := GenS1{F0: 0, F1: 2}
    v8 := GenS2{F0: (-1 == 0), F1: 538, F2: 0}
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int64 = 9223372036854775807
        v11 := GenS3{F0: -1, F1: 1}
        _ = fi9
        _ = v10
        _ = v11
    }
    if false {
        v12 := GenS4{F0: -1, F1: 127, F2: 42}
        var v13 float64 = (-2.25 / 1.5)
        _ = v12
        _ = v13
    }
    _ = v7
    _ = v8
    return 992
}
func genFn1(p0 bool, p1 int) bool {
    switch "" {
    case "foo bar":
        v14 := []int{}
        _ = v14
    case "a":
        var v15 int8 = -1
        _ = v15
    }
    if (255 < 127) {
        v16 := GenS1{F0: 1, F1: 18446744073709551615}
        var v17 int = p1
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 float64 = float64(((0.1 * float32(1e10)) / 1e10))
    v2 := make(map[string]uint32)
    var v3 int32 = 286
    var v4 uint = 13
    var v5 int8 = -1
    var v6 int8 = (v5 % (((v5 % (((v5 | v5)) | 1))) | 1))
    switch 179 {
    case -9223372036854775808:
        if false {
            v1 = (0.0 + 0.0)
            v3 = -1
            v18 := GenS4{F0: (^(v3)), F1: (v5 % ((127) | 1)), F2: 0}
            _ = v18
        }
    }
    fmt.Println(v5, v1)
    if false {
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := v2["\n\t"]
            _ = fi19
            _ = v20
        }
        v3 -= v3
    }
    v21 := []int{-2147483648, -196, 0}
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 int = 42
        if true {
            v1 = (v1 + v1)
            v24 := GenS5{F0: "a", F1: 110, F2: 2}
            if (481 == 867) {
                v25 := GenS4{F0: (v3 >> ((0) & 63)), F1: -38, F2: 42}
                var v26 float64 = (((0.1 / 1.0) + 1e10) * (-2.25 - (v1 + (1e10 / (-2.25 + v1)))))
                _ = v25
                _ = v26
            }
            _ = v24
        } else {
            if false {
                var v27 string = "a"
                _ = v27
            }
            v28 := GenS3{F0: 1, F1: 212}
            switch "\n\t" {
            case "\n\t":
                for fi29 := 0; fi29 < 9; fi29++ {
                    if (false || false) {
                        fmt.Println(v5)
                    }
                    v30 := v2[("\n\t" + "hello")]
                    _ = fi29
                    _ = v30
                }
            case "a":
                if true {
                    if false {
                        var v31 uint8 = 0
                        _ = v31
                    }
                } else {
                    if (!false) {
                        v32 := v2[""]
                        _ = v32
                    }
                    if true {
                        delete(v2, "hello")
                        if (v4 > v4) {
                            v33 := GenS2{F0: (0.0 <= (v1 + v1)), F1: 32767, F2: 32767}
                            _ = v33
                        } else {
                            v34 := GenS6{F0: (!false), F1: 42}
                            var v35 int64 = 42
                            _ = v34
                            _ = v35
                        }
                    } else {
                        v6 = (v5 | -1)
                        for fi36 := 0; fi36 < 5; fi36++ {
                            _ = len(v21)
                            _ = fi36
                        }
                        fmt.Println(v6)
                    }
                    var v37 bool = true
                    _ = v37
                }
            }
            _ = v28
        }
        _ = fi22
        _ = v23
    }
    fmt.Println(v6)
    switch -9223372036854775808 {
    case 46:
        v3 = int32(v5)
    case 2147483647:
        v3 += (v3 & (-(v3)))
    case 414:
        for fi38 := 0; fi38 < 10; fi38++ {
            fmt.Println(v3)
            var v39 bool = (false && ((0.0 <= 1e10) || false))
            v5 = v5
            _ = fi38
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        var v41 int64 = 537
        var v42 float32 = ((-2.25 + float32((0.1 / (v1 + v1)))) + float32((((float64(0.0) * ((-2.25 / (v1 / (v1 * float64((-2.25 - float32((v1 + v1))))))) - (-2.25 + v1))) - float64(float32(v1))) / float64(1e10))))
        _ = fi40
        _ = v41
        _ = v42
    }
    delete(v2, (("" + "\n\t") + ""))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
