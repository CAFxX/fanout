package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }
type GenS3 struct { F0 string; F1 uint16; F2 bool }
type GenS4 struct { F0 uint32; F1 int }
func genFn0(p0 bool) int16 {
    v6 := []int{42}
    p0 = (!((1000 < -32768) || false))
    _ = p0
    _ = v6
    return 100
}

func main() {
    v1 := []int{}
    var v2 float32 = 0.1
    v3 := make(map[uint32]bool)
    var v4 uint = 2
    var v5 uint32 = 37
    var v7 float64 = (0.0 + float64(((v2 + ((v2 - v2) + v2)) / ((1e10 - (float32((-2.25 * 0.0)) - 0.0)) + ((v2 + (1e10 * v2)) - (v2 - 1.5))))))
    var v8 uint16 = 42
    v9 := GenS1{F0: -117, F1: (^(v8)), F2: ("\n\t" + (("foo bar" + "a") + "hello"))}
    fmt.Println(v4, v5, v4)
    fmt.Println(v7, v4)
    if true {
        if (233 != (v4 >> ((1) & 63))) {
            var v10 uint = v4
            _ = v10
        } else {
            if (2 >= 4294967295) {
                if len(v1) > 0 {
                    v1[0] = 1
                }
                v11 := GenS1{F0: 1, F1: (v8 / ((uint16(v5)) | 1)), F2: ""}
                _ = v11
            }
        }
    }
    var v12 string = "a"
    if (!false) {
        var v13 int = 32768
        _ = v13
    } else {
        var v14 int8 = 0
        v15 := GenS1{F0: (v14 - (v14 / (((v14 * v14)) | 1))), F1: v8, F2: v12}
        delete(v3, v5)
        _ = v14
        _ = v15
    }
    var v16 int32 = -1
    v8 = 42
    if (!false) {
        v17 := GenS2{F0: 42, F1: 32767, F2: (v8 >> (((v4 - (v4 * (v4 + v4)))) & 63))}
        for fi18 := 0; fi18 < 7; fi18++ {
            v12 = v12
            _ = fi18
        }
        var v19 uint8 = 127
        _ = v17
        _ = v19
    }
    if false {
        if (-372 >= 42) {
            v20 := v3[2]
            for fi21 := 0; fi21 < 7; fi21++ {
                fmt.Println(v8, v7, v16)
                _ = fi21
            }
            _ = v20
        }
        for fi22 := 0; fi22 < 8; fi22++ {
            if false {
                switch v12 {
                case "\n\t":
                    _ = len(v3)
                default:
                    v1 = append(v1, (fi22 * (^(fi22))))
                }
                fmt.Println(v16, v16)
                v23 := GenS3{F0: "", F1: (v8 + 1), F2: false}
                _ = v23
            }
            v24 := GenS4{F0: 116, F1: fi22}
            delete(v3, v5)
            _ = fi22
            _ = v24
        }
    }
    v25 := GenS4{F0: v5, F1: -262}
    for fi26 := 0; fi26 < 19; fi26++ {
        var v27 float32 = 1e10
        _ = fi26
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
