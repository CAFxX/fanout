package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint }
type GenS2 struct { F0 int32; F1 int8; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 string }
func genFn0(p0 uint32) int64 {
    p0 += 2
    v7 := []int{}
    if true {
        fmt.Println(p0)
        var v8 int8 = -114
        fmt.Println(p0, p0)
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint16 = 32767
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v7
    return -1
}
func genFn1(p0 int16, p1 int64, p2 uint32) bool {
    if false {
        switch 128 {
        case 739:
            p0 -= (p0 ^ 0)
        case -2147483648:
            fmt.Println(p1, p0)
        case 9223372036854775807:
            if (-70 >= -97) {
                var v11 int64 = p1
                _ = v11
            }
        default:
            var v12 uint64 = 0
            _ = v12
        }
        fmt.Println(p2, p0)
    } else {
        if false {
            p2 -= 0
            p0 = 1000
        }
    }
    p0 = -32768
    v13 := GenS1{F0: (!(((true && ((-(p1)) > (p1 | p1))) || (-2147483648 >= int32(p0))) && (468 != 255))), F1: 1, F2: 700}
    if false {
        fmt.Println(p2, p2, p0)
        var v14 float64 = ((3.14159 * 0.1) / 3.14159)
        if (255 >= 403) {
            v15 := []int{32767, 42, 127}
            _ = v15
        } else {
            if (248 < p0) {
                v16 := []int{32768}
                v17 := GenS1{F0: true, F1: 705, F2: 18446744073709551615}
                _ = v16
                _ = v17
            }
            v18 := []int{32768, 127, 32767}
            _ = v18
        }
        _ = v14
    } else {
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := GenS2{F0: 2147483647, F1: 127, F2: 282}
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 5; fi21++ {
            v22 := GenS2{F0: -1, F1: -128, F2: fi21}
            for fi23 := 0; fi23 < 16; fi23++ {
                p1 = p1
                if ((p1 ^ (p1 / ((p1) | 1))) != (p1 << ((65535) & 63))) {
                    fmt.Println(fi23)
                    p0 = (^(p0))
                } else {
                    var v24 int32 = -1
                    var v25 bool = true
                    _ = v24
                    _ = v25
                }
                switch ("hello" + "") {
                case "":
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(p0, fi23, p0)
                    default:
                        var v26 uint8 = 128
                        var v27 uint64 = 26
                        _ = v26
                        _ = v27
                    }
                case "\n\t":
                    fmt.Println(p0)
                case "foo bar":
                    v28 := GenS3{F0: uint8(p2)}
                    _ = v28
                default:
                    v29 := []int{fi23, 128, (fi23 & (^(fi21))), (fi23 + fi23)}
                    _ = v29
                }
                v30 := GenS3{F0: 0}
                p1 = -9223372036854775808
                _ = fi23
                _ = v30
            }
            var v31 float32 = float32((1.5 - 0.1))
            _ = fi21
            _ = v22
            _ = v31
        }
        p0 = -1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return true
}

func main() {
    v1 := []uint32{42, 4294967295}
    var v2 int16 = 0
    v3 := make(map[string]bool)
    var v4 int = 32767
    var v5 int64 = -841
    var v6 float32 = 0.0
    v2 -= 32767
    var v32 float32 = ((v6 * ((float32((1e10 / 3.14159)) - v6) + v6)) * (v6 + v6))
    v33 := GenS1{F0: false, F1: 464, F2: 255}
    if false {
        switch v4 {
        case -521:
            v4 += -9223372036854775808
        default:
            for fi34 := 0; fi34 < 15; fi34++ {
                var v35 uint32 = 4294967295
                switch 633 {
                case 42:
                    var v36 uint8 = 255
                    _ = v36
                case 1:
                    var v37 string = "a"
                    _ = v37
                default:
                    var v38 int32 = -623
                    _ = v38
                }
                _ = fi34
                _ = v35
            }
        }
        if len(v1) > 0 {
            v1[0] = 42
        }
        switch 195 {
        case 275:
            v5 -= genFn0(0)
        default:
            var v39 int16 = 1000
            _ = v39
        }
    } else {
        v40 := v3[""]
        _ = v40
    }
    var v41 int64 = v5
    if true {
        v3["a"] = true
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        fmt.Println(fi42, fi42)
        var v43 int16 = v2
        _ = fi42
        _ = v43
    }
    var v44 uint = 182
    if (uint32(v44) == 219) {
        v45 := GenS4{F0: 127, F1: (("" + "hello") + (("\n\t" + "") + "a"))}
        v6 = ((float32((1.5 - 1.5)) + v32) * (-2.25 - 1e10))
        var v46 int8 = -43
        _ = v45
        _ = v46
    }
    var v47 int = (v4 * 128)
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v44
    _ = v47
    _ = v5
    _ = v6
}
