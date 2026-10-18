package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int8; F1 int16; F2 string }
func genFn0(p0 uint32, p1 int64) uint {
    v4 := GenS1{F0: 255, F1: 238, F2: true}
    switch -166 {
    case 2147483647:
        v5 := []int{600}
        _ = v5
    case 4294967295:
        if false {
            v6 := GenS1{F0: 1, F1: 65535, F2: false}
            v7 := []int{4294967295, 42, 4294967295}
            v7 = append(v7, 32767)
            _ = v6
            _ = v7
        }
    default:
        p0 += 42
    }
    _ = p0
    _ = p1
    _ = v4
    return 663
}
func genFn1(p0 uint16) string {
    var v8 uint32 = 1
    for fi9 := 0; fi9 < 12; fi9++ {
        if (255 <= 42) {
            v10 := GenS1{F0: (p0 | (p0 & 32767)), F1: (p0 * p0), F2: true}
            if (460 > -986) {
                var v11 string = ("foo bar" + "")
                var v12 bool = (0 != 255)
                _ = v11
                _ = v12
            }
            fmt.Println(v8, p0)
            _ = v10
        }
        v13 := GenS1{F0: (p0 - (^(p0))), F1: 32767, F2: true}
        if true {
            var v14 int8 = -128
            fmt.Println(v8, v14)
            v15 := GenS2{F0: -32768}
            _ = v14
            _ = v15
        }
        _ = fi9
        _ = v13
    }
    if (int64(p0) != 42) {
        fmt.Println(v8, v8)
        v16 := GenS3{F0: genFn0(2, -186), F1: genFn0(2147483647, 855)}
        var v17 int = 0
        _ = v16
        _ = v17
    } else {
        v18 := GenS2{F0: 32767}
        _ = v18
    }
    fmt.Println(v8, v8)
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := GenS3{F0: 4294967295, F1: 2}
        v8 += v8
        if (false || true) {
            var v21 int64 = 571
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = v8
    return "a"
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 0.1
    var v3 float32 = ((1.5 / (v2 * 1e10)) * (v2 - v2))
    delete(v1, -754)
    var v22 int16 = 0
    var v23 uint64 = 1
    v22 = (v22 - 1)
    v24 := GenS2{F0: v22}
    v2 += v2
    v2 -= (v3 - (1e10 - float32(((0.0 / (((1e10 / 3.14159) + -2.25) + (((1e10 * ((((1.5 - 0.1) * 0.1) / 1.5) - -2.25)) * (1e10 + ((0.1 / 1.0) - (0.1 * 3.14159)))) + 0.1))) / float64(0.1)))))
    delete(v1, 50)
    for fi25 := 0; fi25 < 2; fi25++ {
        if true {
            switch fi25 {
            case -9223372036854775808:
                v26 := GenS4{F0: 0, F1: (v22 / ((v22) | 1)), F2: "hello"}
                _ = v26
            default:
                if ((65535 >= 18446744073709551615) || (!false)) {
                    if false {
                        v27 := []int{fi25, (fi25 - (fi25 | fi25)), fi25, 248}
                        var v28 int64 = 1
                        _ = v27
                        _ = v28
                    } else {
                        v3 += float32(0.1)
                        v29 := GenS2{F0: v22}
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 11; fi30++ {
                        v2 = 0.1
                        _ = fi30
                    }
                    v22 += (-(v22))
                } else {
                    var v31 int = 2147483647
                    _ = v31
                }
            }
            if (-128 < 76) {
                if (!(4294967295 <= (v23 * 272))) {
                    fmt.Println(fi25)
                } else {
                    fmt.Println(v22, v2, v2)
                }
                v3 += v2
                var v32 float64 = 0.1
                _ = v32
            }
            v3 = 3.14159
        }
        if (!true) {
            _ = len(v1)
            v33 := []int{fi25}
            _ = v33
        }
        _ = fi25
    }
    switch ("" + "foo bar") {
    case "foo bar":
        for fi34 := 0; fi34 < 17; fi34++ {
            fmt.Println(v3, v22, fi34)
            _ = len(v1)
            _ = fi34
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
