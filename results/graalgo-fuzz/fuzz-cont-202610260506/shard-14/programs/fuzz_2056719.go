package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS3 struct { F0 bool; F1 int16; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 uint16; F2 uint64 }
func genFn0(p0 uint32, p1 uint8, p2 string) int16 {
    for fi5 := 0; fi5 < 15; fi5++ {
        fmt.Println(fi5)
        p1 -= 128
        _ = fi5
    }
    var v6 int64 = 42
    for fi7 := 0; fi7 < 20; fi7++ {
        switch fi7 {
        case 32767:
            p0 += p0
        case 128:
            var v8 float32 = 1.5
            _ = v8
        default:
            p0 += (p0 / ((0) | 1))
        }
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 int8 = 1
            _ = fi9
            _ = v10
        }
        _ = fi7
    }
    v11 := []int{4294967295, -72, -651, 32768}
    if (-1 > 9223372036854775807) {
        if ((p2 + p2) < ((("hello" + p2) + ((("a" + "\n\t") + (("hello" + p2) + "foo bar")) + (("a" + "hello") + p2))) + "")) {
            var v12 uint8 = p1
            _ = v12
        } else {
            v6 -= 204
        }
        v13 := GenS1{F0: 255}
        _ = v13
    } else {
        v6 = (v6 * 501)
        v6 += (-(v6))
        var v14 float32 = ((0.0 + 0.0) - (1e10 / 0.1))
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    return 32767
}

func main() {
    var v1 int8 = 42
    var v2 int32 = 2147483647
    var v3 int8 = 42
    var v4 float64 = (3.14159 * float64(float32(float64(-2.25))))
    var v15 bool = ((4294967295 < 1) || false)
    switch -9223372036854775808 {
    case -2147483648:
        for fi16 := 0; fi16 < 4; fi16++ {
            fmt.Println(v2, v3, v4)
            switch "" {
            case "foo bar":
                v17 := GenS2{F0: 42, F1: 144, F2: 933}
                _ = v17
            default:
                if (false || false) {
                    v18 := []int{(fi16 % (((fi16 & 0)) | 1)), 4294967295, fi16}
                    _ = v18
                }
            }
            _ = fi16
        }
    default:
        v19 := []int{-572, 2147483647}
        _ = v19
    }
    v2 -= (v2 / ((-1) | 1))
    var v20 int16 = -1
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS3{F0: ((((("foo bar" + (("foo bar" + ("foo bar" + ("foo bar" + "foo bar"))) + "a")) + ("hello" + "hello")) + "foo bar") + "") >= (("foo bar" + ((("a" + "a") + ("foo bar" + "a")) + "foo bar")) + "\n\t")), F1: (v20 + 1000), F2: -805}
        switch (fi21 ^ (^(fi21))) {
        case 42:
            v23 := []int{}
            _ = v23
        case 492:
            v24 := GenS4{F0: 255}
            _ = v24
        case 0:
            var v25 float32 = 3.14159
            _ = v25
        }
        _ = fi21
        _ = v22
    }
    fmt.Println(v20)
    v26 := []int{-203, -2147483648}
    for fi27 := 0; fi27 < 7; fi27++ {
        v28 := GenS3{F0: true, F1: (-(v20)), F2: -913}
        v1 = 1
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        fmt.Println(v1, v4)
        var v30 uint = 622
        var v31 float64 = ((v4 / (float64(3.14159) + v4)) - float64(1e10))
        _ = fi29
        _ = v30
        _ = v31
    }
    _ = len(v26)
    switch "foo bar" {
    case "a":
        v32 := GenS1{F0: 4294967295}
        _ = v32
    case "":
        fmt.Println(v15, v1, v3)
    case "foo bar":
        if v15 {
            for fi33 := 0; fi33 < 6; fi33++ {
                var v34 int32 = (v2 ^ v2)
                _ = fi33
                _ = v34
            }
        } else {
            for fi35 := 0; fi35 < 8; fi35++ {
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                _ = fi35
            }
            v36 := GenS5{F0: v3, F1: 755, F2: 0}
            v1 = v1
            _ = v36
        }
    default:
        v37 := GenS4{F0: 55}
        _ = v37
    }
    var v38 float64 = -2.25
    v20 -= v20
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v38
    _ = v4
}
