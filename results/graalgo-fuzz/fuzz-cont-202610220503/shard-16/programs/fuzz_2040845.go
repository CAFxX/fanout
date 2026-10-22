package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int32; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint64, p1 uint8, p2 int16) uint8 {
    var v5 uint64 = p0
    if true {
        var v6 uint64 = 42
        _ = v6
    } else {
        p1 = 42
    }
    var v7 float64 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return 128
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    v4 := []string{"foo bar", "\n\t"}
    var v8 int64 = -2
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            switch fi10 {
            case 42:
                v8 -= (v8 / ((v8) | 1))
            default:
                v11 := GenS1{F0: (("a" + "a") + ""), F1: (v8 * 0)}
                _ = v11
            }
            _ = fi10
        }
        v12 := GenS1{F0: "\n\t", F1: (v8 + (v8 ^ (-(v8))))}
        if false {
            fmt.Println(v8, fi9)
            var v13 int64 = v8
            v1[781] = "hello"
            _ = v13
        } else {
            fmt.Println(fi9, v8)
            var v14 float32 = 1.5
            _ = v14
        }
        _ = fi9
        _ = v12
    }
    v8 -= (v8 / (((v8 << ((2) & 63))) | 1))
    v15 := v2[859]
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 uint8 = 2
        fmt.Println(v8, fi16, v17)
        _ = fi16
        _ = v17
    }
    if len(v4) > 0 {
        v4[0] = ("\n\t" + v15)
    }
    v15 = (v15 + (("hello" + "\n\t") + v15))
    v8 = v8
    if true {
        if (52 > 42) {
            v18 := GenS1{F0: v15, F1: (v8 >> ((2) & 63))}
            _ = v18
        } else {
            for fi19 := 0; fi19 < 14; fi19++ {
                fmt.Println(fi19, fi19)
                v20 := GenS2{F0: 1, F1: (false && true), F2: 565}
                _ = fi19
                _ = v20
            }
            if (!(236 == 0)) {
                var v21 int64 = (v8 & v8)
                fmt.Println(v21, v15)
                for fi22 := 0; fi22 < 14; fi22++ {
                    var v23 int64 = 0
                    v24 := GenS3{F0: 42, F1: 1}
                    if true {
                        var v25 uint64 = 4294967295
                        v15 = ""
                        _ = v25
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = v21
            }
        }
        var v26 float32 = (3.14159 / 1.0)
        if true {
            _ = len(v4)
            v26 -= ((3.14159 / -2.25) + ((v26 - (1e10 * (3.14159 / v26))) / (((3.14159 * (0.0 + ((3.14159 / 1.0) - (v26 + v26)))) - 1e10) + v26)))
            for fi27 := 0; fi27 < 10; fi27++ {
                var v28 uint8 = genFn0(1, 46, 614)
                var v29 int = 975
                _ = fi27
                _ = v28
                _ = v29
            }
        } else {
            fmt.Println(v15, v8)
            var v30 uint16 = 255
            _ = v30
        }
        _ = v26
    }
    fmt.Println(v15)
    var v31 uint16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v8
}
