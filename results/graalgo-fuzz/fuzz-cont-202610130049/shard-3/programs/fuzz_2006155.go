package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int64 }
type GenS2 struct { F0 string }
func genFn0(p0 string, p1 int64, p2 int8) int32 {
    if false {
        v7 := GenS1{F0: -993, F1: (127 >= 837), F2: -200}
        if (((p0 + p0) + "hello") <= p0) {
            var v8 uint32 = 693
            fmt.Println(v8, p0, v8)
            _ = v8
        } else {
            var v9 string = p0
            _ = v9
        }
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint = 42
            _ = fi10
            _ = v11
        }
        _ = v7
    }
    p0 = (p0 + p0)
    switch p0 {
    case "hello":
        fmt.Println(p1)
    case "":
        var v12 int = 127
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 float32 = 1e10
        p1 -= p1
        v14 -= ((0.1 - (v14 - v14)) / (((-2.25 * (v14 * -2.25)) + float32(0.1)) + v14))
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]bool)
    var v3 float64 = ((1.5 + 1e10) - 0.0)
    v4 := []string{}
    var v5 int32 = -1
    v6 := make(map[string]int)
    _ = len(v4)
    var v15 uint = 1
    var v16 float32 = -2.25
    if (!false) {
        _ = len(v4)
        if (118 >= 42) {
            v1 -= v1
            v1 = (v1 | v1)
        } else {
            for fi17 := 0; fi17 < 12; fi17++ {
                switch (fi17 | (fi17 & fi17)) {
                case 32767:
                    fmt.Println(fi17, v15, v1)
                case -626:
                    v15 = v15
                case 1:
                    var v18 int = (fi17 + (^(fi17)))
                    _ = v18
                default:
                    v3 += float64(((v16 + v16) + 0.1))
                }
                switch "\n\t" {
                case "foo bar":
                    v19 := GenS1{F0: -9223372036854775808, F1: false, F2: 752}
                    _ = v19
                }
                _ = fi17
            }
            var v20 bool = true
            _ = v20
        }
        v16 = 0.1
    } else {
        if false {
            _ = len(v6)
            var v21 uint8 = 1
            _ = v21
        } else {
            v3 += (v3 - 1e10)
            for fi22 := 0; fi22 < 4; fi22++ {
                if ((((-2.25 * (v16 - v16)) / (v16 + v16)) - v16) == float32(v3)) {
                    if false {
                        var v23 int = 128
                        _ = v23
                    }
                } else {
                    fmt.Println(v5, v15, v15)
                    if (!(true || ((!(!true)) || true))) {
                        v3 -= v3
                    } else {
                        v1 += (v1 & 730)
                        v16 += v16
                        _ = len(v6)
                    }
                    var v24 uint = 474
                    _ = v24
                }
                var v25 int16 = -1
                v3 -= (((v3 / v3) * float64(v16)) - (v3 * (v3 + (v3 * float64((float32(float64(v16)) + (3.14159 + (3.14159 * 0.1))))))))
                _ = fi22
                _ = v25
            }
            if ((2 == 975) || true) {
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 uint64 = 970
                    _ = fi26
                    _ = v27
                }
            }
        }
    }
    if (-128 >= 127) {
        fmt.Println(v15, v16)
        var v28 uint16 = (v1 % ((v1) | 1))
        _ = v28
    } else {
        fmt.Println(v15, v3, v5)
    }
    v29 := v6["a"]
    for fi30 := 0; fi30 < 3; fi30++ {
        v31 := GenS2{F0: "a"}
        _ = fi30
        _ = v31
    }
    var v32 uint32 = uint32(v1)
    v29 = int(v32)
    var v33 int = (v29 % (((v29 ^ v29)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
