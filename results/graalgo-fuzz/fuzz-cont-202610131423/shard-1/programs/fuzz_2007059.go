package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int; F1 uint }
func genFn0(p0 uint32, p1 int) bool {
    var v6 uint8 = 1
    v7 := GenS1{F0: (!true), F1: 4294967295, F2: -1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (((0.0 + 0.0) * (float32(3.14159) + 0.0)) >= (3.14159 * float32(-2.25)))
}

func main() {
    v1 := []string{"hello"}
    v2 := []string{"foo bar"}
    v3 := make(map[string]int)
    var v4 uint16 = 255
    v5 := []string{""}
    for fi8 := 0; fi8 < 15; fi8++ {
        _ = len(v2)
        v4 = v4
        fmt.Println(fi8, fi8)
        _ = fi8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            _ = len(v1)
            v2 = append(v2, "hello")
            var v11 int = fi9
            _ = fi10
            _ = v11
        }
        if (0 >= 2) {
            var v12 int8 = 92
            if true {
                fmt.Println(v12)
            } else {
                var v13 int = (fi9 - fi9)
                v14 := GenS1{F0: ((0.1 != (3.14159 / 1e10)) || false), F1: 796, F2: 128}
                _ = v13
                _ = v14
            }
            fmt.Println(v4, v4)
            _ = v12
        } else {
            v4 = (v4 - v4)
            v15 := GenS1{F0: ((837 >= 0) || ((((float64(0.1) / 3.14159) / 1.0) - (3.14159 + 0.1)) == 1.5)), F1: 0, F2: 1}
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            _ = v15
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 float64 = 3.14159
            if true {
                var v18 float32 = ((1e10 * float32(v17)) - 1e10)
                _ = v18
            } else {
                v19 := GenS2{F0: 805}
                if (!false) {
                    v20 := GenS2{F0: 65535}
                    switch fi16 {
                    case -977:
                        v21 := GenS2{F0: 0}
                        _ = v21
                    case 700:
                        v22 := GenS3{F0: ("foo bar" + ("a" + "foo bar")), F1: 0}
                        _ = v22
                    }
                    var v23 float64 = 1e10
                    v23 = v17
                    _ = v20
                    _ = v23
                }
                fmt.Println(v4)
                v24 := GenS1{F0: (2 != 350), F1: 65535, F2: -2147483648}
                if ((("a" > "") && false) || (127 > uint8(v4))) {
                    var v25 float32 = 1.5
                    _ = v25
                } else {
                    var v26 uint = 2
                    _ = v26
                }
                _ = v19
                _ = v24
            }
            _ = len(v3)
            _ = fi16
            _ = v17
        }
        _ = fi9
    }
    v27 := GenS3{F0: "\n\t", F1: 1}
    if false {
        if true {
            var v28 float64 = 0.1
            _ = v28
        } else {
            var v29 uint64 = uint64(v4)
            _ = v29
        }
        for fi30 := 0; fi30 < 14; fi30++ {
            v31 := GenS4{F0: false, F1: 42}
            _ = fi30
            _ = v31
        }
        switch 128 {
        case 32767:
            v4 = 255
        default:
            v4 = v4
        }
    }
    var v32 string = "\n\t"
    for fi33 := 0; fi33 < 8; fi33++ {
        if false {
            v34 := v3[v32]
            fmt.Println(v4)
            switch fi33 {
            case 9223372036854775807:
                v34 = fi33
            case -2147483648:
                v35 := GenS5{F0: (-(fi33)), F1: 18446744073709551615}
                _ = v35
            case 18:
                delete(v3, ("foo bar" + ("" + ("" + "a"))))
            default:
                var v36 uint = 436
                _ = v36
            }
            _ = v34
        }
        if true {
            v37 := GenS1{F0: (1.5 < (-2.25 * 0.0)), F1: 255, F2: -1}
            _ = v37
        }
        v2 = append(v2, "")
        _ = fi33
    }
    var v38 int32 = 769
    v4 = 0
    v38 = v38
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v5
}
