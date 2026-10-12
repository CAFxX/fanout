package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 string }
func genFn0(p0 int32, p1 uint16) int64 {
    var v4 float32 = -2.25
    p0 += (p0 % ((-1) | 1))
    var v5 uint64 = 901
    if (128 <= 128) {
        v6 := GenS1{F0: p0}
        switch 42 {
        case -9223372036854775808:
            switch int(p0) {
            case 4294967295:
                if (-1 == 217) {
                    var v7 uint = 4294967295
                    var v8 uint64 = 16
                    _ = v7
                    _ = v8
                }
            }
        case 32767:
            v9 := GenS2{F0: p0, F1: "foo bar"}
            _ = v9
        case 292:
            var v10 float64 = 3.14159
            _ = v10
        }
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -9223372036854775808
}

func main() {
    v1 := []string{"\n\t"}
    var v2 float32 = -2.25
    var v3 int16 = 42
    fmt.Println(v2, v2)
    var v11 int32 = 868
    v3 += v3
    v1 = append(v1, (("" + "hello") + "\n\t"))
    v12 := GenS2{F0: (v11 / (((v11 | int32(v3))) | 1)), F1: ""}
    for fi13 := 0; fi13 < 11; fi13++ {
        fmt.Println(v3, v11, v3)
        var v14 uint8 = 42
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        fmt.Println(fi15)
        if false {
            for fi16 := 0; fi16 < 8; fi16++ {
                v17 := GenS1{F0: 631}
                if true {
                    fmt.Println(v11)
                    v2 = (float32(3.14159) * ((v2 * 0.1) / ((float32(0.1) + float32(1e10)) - ((0.1 + v2) + v2))))
                }
                _ = fi16
                _ = v17
            }
            var v18 float32 = (((0.0 + float32(1e10)) * 1e10) * (float32(3.14159) - 1e10))
            _ = v18
        }
        var v19 uint32 = 65535
        _ = fi15
        _ = v19
    }
    var v20 int16 = v3
    if true {
        var v21 int = 42
        if (0 < 255) {
            v21 = (-(v21))
        } else {
            for fi22 := 0; fi22 < 11; fi22++ {
                v23 := GenS1{F0: (v11 ^ v11)}
                _ = fi22
                _ = v23
            }
        }
        var v24 bool = false
        _ = v21
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
}
