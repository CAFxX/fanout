package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint32; F1 int8; F2 uint16 }
func genFn0(p0 int32, p1 uint64) string {
    for fi6 := 0; fi6 < 19; fi6++ {
        if (-128 == 1) {
            fmt.Println(fi6, fi6)
            var v7 bool = true
            if true {
                fmt.Println(fi6, v7)
                var v8 int = (fi6 >> ((654) & 63))
                _ = v8
            } else {
                fmt.Println(p1)
                for fi9 := 0; fi9 < 4; fi9++ {
                    if (((-2.25 <= 1e10) && (false || true)) || true) {
                        p0 += (p0 * (p0 | (p0 & p0)))
                    } else {
                        v10 := GenS1{F0: "hello", F1: 211, F2: -300}
                        switch "" {
                        case "hello":
                            v11 := []int{(fi6 / ((fi6) | 1)), fi6}
                            _ = v11
                        default:
                            var v12 int32 = p0
                            _ = v12
                        }
                        var v13 int8 = 100
                        var v14 int16 = 1000
                        _ = v10
                        _ = v13
                        _ = v14
                    }
                    _ = fi9
                }
            }
            _ = v7
        }
        p0 += 2147483647
        _ = fi6
    }
    var v15 uint32 = 271
    v16 := GenS2{F0: v15}
    var v17 uint32 = v15
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := GenS3{F0: (p0 | p0), F1: 473}
        p0 += 2147483647
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    _ = v17
    return "\n\t"
}
func genFn1(p0 uint8) uint16 {
    var v20 string = "hello"
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS4{F0: 0, F1: -70, F2: 1}
        _ = fi21
        _ = v22
    }
    _ = p0
    _ = v20
    return 2
}

func main() {
    var v1 bool = (1.5 != 0.0)
    var v2 float64 = float64(((1.5 / 1.0) - float32(-2.25)))
    var v3 int16 = 678
    var v4 string = ""
    var v5 string = v4
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 uint8 = 128
        var v25 uint64 = 188
        if (1e10 == float64(0.1)) {
            for fi26 := 0; fi26 < 15; fi26++ {
                v27 := GenS2{F0: 2}
                _ = fi26
                _ = v27
            }
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println(v2, v1)
    var v28 int16 = (v3 | -1)
    for fi29 := 0; fi29 < 4; fi29++ {
        v2 = v2
        var v30 int16 = 742
        _ = fi29
        _ = v30
    }
    v5 = "foo bar"
    var v31 int = 1
    for fi32 := 0; fi32 < 19; fi32++ {
        var v33 int64 = 42
        var v34 uint32 = 0
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println(v28, v1)
    switch ("\n\t" + v5) {
    case "":
        if (v1 && (false && v1)) {
            v35 := []int{-288, 2147483647, -1, -928}
            _ = v35
        } else {
            v28 -= (v28 % ((v3) | 1))
        }
    case "a":
        if (float32(float64(1.5)) < float32(v2)) {
            var v36 float32 = 0.1
            v37 := GenS2{F0: 4294967295}
            v1 = ((((42 > 2) || (v1 && (false || false))) && (v1 || ((v36 < (3.14159 / v36)) && (9223372036854775807 <= int64(v31))))) && (((v28 % ((v3) | 1)) <= v3) || (((v4 + (((v4 + v4) + "foo bar") + (v5 + ("\n\t" + "\n\t")))) + v4) == "a")))
            _ = v36
            _ = v37
        }
    default:
        for fi38 := 0; fi38 < 16; fi38++ {
            var v39 uint16 = 0
            if true {
                var v40 int = (fi38 % (((v31 % ((fi38) | 1))) | 1))
                v1 = (0.1 >= 0.1)
                _ = v40
            }
            _ = fi38
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
