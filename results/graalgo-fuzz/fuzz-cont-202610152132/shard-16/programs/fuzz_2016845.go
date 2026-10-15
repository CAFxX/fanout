package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 bool; F1 int; F2 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 int64; F2 uint }
func genFn0(p0 string, p1 int16, p2 int8) string {
    v4 := GenS1{F0: p0}
    v5 := []int{42, -9223372036854775808, 42}
    var v6 int8 = -128
    v7 := GenS2{F0: 255, F1: 1}
    fmt.Println(v6, v6, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return ""
}
func genFn1(p0 int32, p1 bool) int32 {
    v8 := GenS3{F0: (p1 && ((0.0 * 0.1) <= (-2.25 / float64(((0.1 / (1.5 * (-2.25 * 1.5))) - 0.0))))), F1: 32768, F2: ("a" + "\n\t")}
    p1 = ("hello" < "hello")
    p0 -= -471
    _ = p0
    _ = p1
    _ = v8
    return (-(p0))
}

func main() {
    var v1 uint64 = 876
    var v2 int64 = 1
    v3 := make(map[string]bool)
    fmt.Println(v1, v1)
    for fi9 := 0; fi9 < 16; fi9++ {
        if true {
            if false {
                switch "\n\t" {
                case "hello":
                    v10 := []int{-387, (fi9 * 42), (fi9 << ((1) & 63))}
                    _ = v10
                case "":
                    fmt.Println(v2)
                case "foo bar":
                    v11 := GenS2{F0: 255, F1: 377}
                    _ = v11
                }
            } else {
                var v12 uint16 = 1
                if true {
                    fmt.Println(fi9, v12, fi9)
                } else {
                    v13 := GenS2{F0: (v12 | v12), F1: 2147483647}
                    _ = v13
                }
                if true {
                    fmt.Println(fi9, fi9, v12)
                    var v14 float64 = float64(3.14159)
                    _ = v14
                } else {
                    v12 = v12
                    v15 := GenS4{F0: 4294967295}
                    _ = v15
                }
                _ = v12
            }
            v16 := GenS2{F0: 2, F1: 4294967295}
            _ = v16
        } else {
            var v17 int16 = -45
            for fi18 := 0; fi18 < 9; fi18++ {
                v1 = (v1 ^ (v1 << ((4294967295) & 63)))
                _ = fi18
            }
            _ = v17
        }
        _ = fi9
    }
    v19 := GenS4{F0: v1}
    switch 656 {
    case 595:
        var v20 float32 = (0.1 * (0.1 - (1e10 * -2.25)))
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        if (false && (42 == 127)) {
            switch ("a" + ("a" + ("" + "\n\t"))) {
            case "\n\t":
                switch 2147483647 {
                case 2147483647:
                    v2 += (-(v2))
                case -2147483648:
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v2 -= -9223372036854775808
                        var v23 uint = 2
                        var v24 uint64 = v1
                        var v25 int64 = (-(v2))
                        var v26 int32 = 0
                        v23 -= v23
                        if false {
                            v27 := []int{-9223372036854775808, fi21}
                            v28 := GenS2{F0: 1, F1: 473}
                            v29 := GenS4{F0: (^(v1))}
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                }
            case "a":
                for fi30 := 0; fi30 < 4; fi30++ {
                    v1 += v1
                    switch fi21 {
                    case -852:
                        var v31 bool = (false || (float64(1e10) < (((0.1 - 1.5) + ((1.5 * (0.0 + -2.25)) + (-2.25 + 0.0))) / (((3.14159 * 1e10) + ((float64(0.0) - float64((1.5 + (float32(float64(3.14159)) + ((0.1 * 1.5) - float32(1e10)))))) + 3.14159)) / 1e10))))
                        _ = v31
                    case 657:
                        v1 += 18446744073709551615
                    case 487:
                        fmt.Println(fi21, fi21, fi30)
                    default:
                        var v32 int16 = 32767
                        _ = v32
                    }
                    _ = fi30
                }
            }
        } else {
            v2 = (-(v2))
            v1 = v1
        }
        for fi33 := 0; fi33 < 4; fi33++ {
            var v34 uint64 = (v1 | 42)
            v3[("" + "\n\t")] = false
            v34 -= (v34 >> ((2) & 63))
            _ = fi33
            _ = v34
        }
        _ = fi21
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        var v36 uint16 = 65535
        _ = fi35
        _ = v36
    }
    v37 := GenS2{F0: 335, F1: 211}
    v38 := GenS5{F0: "hello", F1: 593, F2: 255}
    var v39 float32 = (1.5 * 0.0)
    v1 -= v1
    fmt.Println(v2, v2, v1)
    var v40 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
