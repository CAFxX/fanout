package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 uint16 }
func genFn0(p0 int32) int {
    var v5 uint64 = 480
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: 796, F1: -1, F2: (v5 << ((826) & 63))}
        var v8 uint32 = 2
        v9 := []int{fi6}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 float64 = float64(-2.25)
    var v11 string = ((((("hello" + ("" + "hello")) + "hello") + "\n\t") + ("a" + ("a" + ""))) + "")
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := GenS2{F0: 2}
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 uint64 = 221
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v5
    return -362
}

func main() {
    var v1 int64 = 1
    var v2 uint64 = 4294967295
    var v3 bool = (!(0 != 1))
    v4 := make(map[uint32]bool)
    fmt.Println(v3, v1, v2)
    v16 := GenS3{F0: false}
    v4[uint32(v2)] = (!v3)
    _ = len(v4)
    var v17 int64 = (v1 ^ v1)
    if (!false) {
        _ = len(v4)
    } else {
        var v18 int8 = -110
        _ = v18
    }
    var v19 int = int(v17)
    var v20 uint8 = 128
    for fi21 := 0; fi21 < 11; fi21++ {
        switch (v19 | v19) {
        case -489:
            fmt.Println(v19, v1, v3)
        default:
            var v22 int = 867
            _ = v22
        }
        v1 = (v1 % ((1) | 1))
        for fi23 := 0; fi23 < 18; fi23++ {
            v2 -= 250
            if v3 {
                v24 := GenS4{F0: (!v3), F1: ""}
                _ = v24
            } else {
                v2 -= 1
                v25 := GenS5{F0: 169}
                v19 += 9223372036854775807
                v19 = -985
                if (901 > 0) {
                    v26 := []int{fi23, fi23, fi21}
                    _ = v26
                } else {
                    if v3 {
                        _ = len(v4)
                    }
                    fmt.Println(v1)
                    var v27 uint8 = 128
                    _ = v27
                }
                _ = v25
            }
            v19 = (v19 * (fi23 >> ((4294967295) & 63)))
            _ = fi23
        }
        _ = fi21
    }
    var v28 int8 = 103
    var v29 bool = v3
    if v29 {
        v30 := GenS6{F0: v2}
        var v31 int32 = -1
        _ = v30
        _ = v31
    }
    v32 := GenS7{F0: "hello", F1: 42}
    var v33 int = (^(v19))
    if (((v2 > v2) || v3) || v29) {
        for fi34 := 0; fi34 < 6; fi34++ {
            for fi35 := 0; fi35 < 14; fi35++ {
                _ = len(v4)
                fmt.Println(v28)
                v17 -= (v17 & (v1 | 9223372036854775807))
                _ = fi35
            }
            var v36 uint64 = (v2 % (((^(v2))) | 1))
            var v37 uint16 = 352
            _ = fi34
            _ = v36
            _ = v37
        }
        v33 = 977
        switch genFn0(-1) {
        case 1:
            var v38 int = (^(v19))
            _ = v38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
