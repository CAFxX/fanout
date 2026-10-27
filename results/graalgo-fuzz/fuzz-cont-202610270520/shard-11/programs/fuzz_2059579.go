package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 string; F2 bool }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 string }
func genFn0(p0 int, p1 int32) string {
    var v5 uint16 = 1
    v6 := GenS1{F0: (p0 / ((p0) | 1)), F1: 18446744073709551615}
    v5 = (v5 & 578)
    p0 += p0
    var v7 string = ("hello" + "\n\t")
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return v7
}
func genFn1() int {
    var v8 int32 = 0
    switch 127 {
    case -1:
        for fi9 := 0; fi9 < 15; fi9++ {
            for fi10 := 0; fi10 < 3; fi10++ {
                v8 += (v8 % (((v8 / ((v8) | 1))) | 1))
                if false {
                    if (!false) {
                        switch fi9 {
                        case 32768:
                            v8 = -2147483648
                        case 395:
                            fmt.Println(v8)
                        default:
                            v8 -= v8
                        }
                        v11 := GenS2{F0: 2147483648}
                        _ = v11
                    }
                    var v12 uint64 = 42
                    _ = v12
                }
                var v13 int64 = 42
                _ = fi10
                _ = v13
            }
            v8 -= (v8 ^ 2147483647)
            fmt.Println(fi9)
            _ = fi9
        }
    case 10:
        var v14 float64 = 1.5
        _ = v14
    default:
        var v15 int32 = 2147483647
        _ = v15
    }
    _ = v8
    return -9223372036854775808
}

func main() {
    v1 := []int{4294967295, 4294967295, 32768, -509}
    var v2 int16 = -1
    v3 := make(map[string]int)
    v4 := make(map[uint32]bool)
    switch -2147483648 {
    case -9223372036854775808:
        fmt.Println(v2)
    case -42:
        v16 := GenS1{F0: 4294967295, F1: 4294967295}
        _ = v16
    }
    var v17 int8 = -1
    var v18 int = -101
    v17 -= (v17 & (v17 + v17))
    switch "\n\t" {
    case "hello":
        var v19 float64 = (1.5 + 1e10)
        _ = v19
    case "\n\t":
        if (!((((float64(0.1) + float64(float32((1.5 - (0.0 - (0.1 - 3.14159)))))) * -2.25) * ((1e10 + ((((0.0 * 3.14159) * 3.14159) + 0.0) * ((((-2.25 * 3.14159) / ((1.5 - 3.14159) + 0.0)) - 0.0) * 0.0))) + -2.25)) == (float64(0.1) * (0.0 * 1e10)))) {
            if false {
                var v20 float32 = float32(-2.25)
                v1 = append(v1, v18)
                _ = v20
            } else {
                var v21 int = v18
                v18 -= 128
                _ = v21
            }
            fmt.Println(v18, v18, v17)
        }
    default:
        v22 := GenS3{F0: -9223372036854775808, F1: ("foo bar" + ("\n\t" + "hello")), F2: false}
        _ = v22
    }
    v17 = v17
    fmt.Println(v2, v17)
    var v23 uint16 = 32767
    var v24 uint8 = 127
    if (!(!true)) {
        var v25 float32 = (-2.25 + 1.5)
        v18 = 955
        _ = v25
    } else {
        v2 = v2
        var v26 int32 = 0
        _ = v26
    }
    v2 = -1
    for fi27 := 0; fi27 < 17; fi27++ {
        if true {
            if true {
                _ = len(v1)
                v28 := GenS4{F0: 42, F1: ""}
                for fi29 := 0; fi29 < 13; fi29++ {
                    v30 := GenS5{F0: ("\n\t" + "\n\t")}
                    v17 = 127
                    _ = fi29
                    _ = v30
                }
                _ = v28
            }
            if true {
                delete(v3, "hello")
                v2 += v2
            } else {
                v18 -= (fi27 * (fi27 / (((v18 * (fi27 / (((fi27 & v18)) | 1)))) | 1)))
            }
        } else {
            switch "foo bar" {
            case "\n\t":
                v31 := GenS4{F0: -387, F1: "foo bar"}
                _ = v31
            case "foo bar":
                v32 := GenS4{F0: 2147483647, F1: ""}
                _ = v32
            case "a":
                fmt.Println(v2)
            default:
                var v33 int8 = 42
                _ = v33
            }
            v23 += v23
        }
        _ = fi27
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        v18 = (-(v18))
        var v35 int8 = (v17 + -95)
        _ = fi34
        _ = v35
    }
    fmt.Println(v17, v23)
    var v36 uint = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v36
    _ = v4
}
