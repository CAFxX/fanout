package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int32; F2 uint }
func genFn0(p0 bool, p1 uint8, p2 bool) int16 {
    v5 := []int{128, 2147483647, 128, -961}
    p0 = p0
    var v6 int16 = -842
    if p0 {
        var v7 uint64 = 42
        switch -1 {
        case 2147483647:
            if true {
                v6 += 416
                v8 := GenS1{F0: v6, F1: -532, F2: uint32(p1)}
                p0 = p0
                _ = v8
            } else {
                fmt.Println(p2, p0, p2)
            }
        default:
            switch -464 {
            case -9223372036854775808:
                fmt.Println(v7, p2, v6)
            case 787:
                if len(v5) > 0 {
                    v5[0] = -9223372036854775808
                }
            case 1:
                if ((!p0) || p2) {
                    if (((^(p1)) <= p1) && p2) {
                        for fi9 := 0; fi9 < 10; fi9++ {
                            var v10 uint8 = p1
                            v7 = (v7 + 622)
                            var v11 uint8 = p1
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                        var v12 int16 = v6
                        _ = v12
                    } else {
                        var v13 float32 = (float32(3.14159) + 3.14159)
                        _ = v13
                    }
                    var v14 int8 = 100
                    _ = v14
                }
            default:
                var v15 int64 = 0
                _ = v15
            }
        }
        _ = v7
    } else {
        if false {
            fmt.Println(v6, p1, v6)
        } else {
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 int64 = 9223372036854775807
                var v18 string = "foo bar"
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        v6 -= v6
        fmt.Println(p1, p2, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return v6
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"", "foo bar", "a", ("hello" + (("a" + "\n\t") + "\n\t"))}
    var v3 int64 = 0
    v4 := make(map[uint32]bool)
    v19 := GenS1{F0: 32767, F1: 0, F2: 2147483647}
    if true {
        v20 := GenS2{F0: 61}
        var v21 uint8 = 0
        _ = len(v2)
        _ = v20
        _ = v21
    }
    var v22 string = ""
    v2 = append(v2, "")
    var v23 int = -2147483648
    var v24 uint32 = 42
    var v25 int32 = 998
    if true {
        var v26 bool = true
        _ = v26
    } else {
        v24 -= v24
        v27 := GenS3{F0: (1 != (v3 + v3)), F1: -1, F2: 1}
        _ = v27
    }
    v22 = v22
    v28 := GenS2{F0: 1}
    v23 = (^(v23))
    var v29 string = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
