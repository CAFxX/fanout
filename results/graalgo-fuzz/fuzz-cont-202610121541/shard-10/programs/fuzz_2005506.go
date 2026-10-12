package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint64 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint64; F1 int32; F2 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 int64 }
func genFn0(p0 string, p1 int64, p2 uint32) bool {
    p1 += p1
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return (2 >= 518)
}
func genFn1(p0 int16, p1 int32) int32 {
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 uint8 = 199
        if true {
            for fi7 := 0; fi7 < 3; fi7++ {
                p0 -= (-(p0))
                _ = fi7
            }
            switch fi5 {
            case 42:
                v8 := GenS1{F0: p0, F1: p0, F2: 0}
                _ = v8
            }
            var v9 int32 = (p1 / ((p1) | 1))
            _ = v9
        } else {
            var v10 float32 = (float32(float64(1e10)) / (3.14159 - (0.0 * (0.1 - float32(0.0)))))
            v11 := GenS2{F0: 599, F1: "foo bar"}
            _ = v10
            _ = v11
        }
        v6 = 1
        _ = fi5
        _ = v6
    }
    var v12 uint32 = 159
    p1 -= (p1 - p1)
    var v13 uint8 = 33
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return (-(p1))
}

func main() {
    var v1 string = ("" + (("a" + "\n\t") + ""))
    var v2 uint16 = 1
    v3 := []uint32{}
    v4 := make(map[int]string)
    v14 := v4[128]
    if (false || true) {
        _ = len(v3)
    } else {
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 bool = true
            v14 = (("foo bar" + "\n\t") + v14)
            _ = fi15
            _ = v16
        }
        fmt.Println(v1, v2)
    }
    v1 = "\n\t"
    v17 := GenS3{F0: 4294967295, F1: 42, F2: (-288 == -1)}
    for fi18 := 0; fi18 < 7; fi18++ {
        v1 = ("foo bar" + v1)
        if ((0 >= 758) || false) {
            v2 = 434
            v14 = v14
        } else {
            var v19 int64 = 186
            var v20 uint32 = 969
            _ = v19
            _ = v20
        }
        switch int(v2) {
        case 590:
            var v21 uint16 = v2
            _ = v21
        case 42:
            v22 := GenS4{F0: 787}
            _ = v22
        }
        _ = fi18
    }
    if false {
        for fi23 := 0; fi23 < 14; fi23++ {
            switch fi23 {
            case 32767:
                fmt.Println(fi23, v1, v14)
            case -2147483648:
                var v24 uint64 = 4294967295
                _ = v24
            case 2147483647:
                v25 := GenS1{F0: 32767, F1: 1, F2: 160}
                _ = v25
            default:
                for fi26 := 0; fi26 < 9; fi26++ {
                    v27 := GenS5{F0: 3}
                    var v28 int8 = 100
                    if true {
                        v29 := GenS6{F0: 416, F1: 9223372036854775807}
                        _ = v29
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            for fi30 := 0; fi30 < 6; fi30++ {
                var v31 int = fi30
                _ = fi30
                _ = v31
            }
            _ = fi23
        }
    }
    v14 = "a"
    if true {
        for fi32 := 0; fi32 < 18; fi32++ {
            _ = len(v3)
            _ = fi32
        }
        fmt.Println(v14)
    } else {
        for fi33 := 0; fi33 < 9; fi33++ {
            _ = len(v3)
            var v34 uint64 = 4294967295
            v2 -= (v2 / ((v2) | 1))
            _ = fi33
            _ = v34
        }
        v4[1] = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
