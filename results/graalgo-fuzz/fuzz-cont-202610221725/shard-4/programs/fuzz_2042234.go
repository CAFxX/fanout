package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int32; F1 uint8; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint8; F2 bool }
type GenS5 struct { F0 uint32; F1 int64; F2 uint32 }

func main() {
    var v1 int = 1
    var v2 int = v1
    v3 := make(map[uint32]string)
    switch v2 {
    case -9223372036854775808:
        var v4 uint64 = 190
        _ = v4
    case 42:
        v5 := []int{(v1 & (v2 << ((65535) & 63))), (v2 | (v2 * 32768))}
        _ = v5
    }
    v6 := GenS1{F0: 4294967295, F1: "a"}
    v2 = -1
    v1 += (v2 << ((262) & 63))
    var v7 int32 = -2147483648
    if true {
        fmt.Println(v7, v7, v1)
    }
    var v8 uint64 = 0
    v1 = (v2 * v1)
    fmt.Println(v8, v1, v2)
    if (!(!(786 > 1))) {
        if true {
            var v9 bool = true
            _ = v9
        }
    } else {
        fmt.Println(v8, v2, v8)
        v10 := GenS2{F0: (v7 ^ 37), F1: 2, F2: 42}
        _ = v10
    }
    if (1 >= -271) {
        var v11 float64 = 0.0
        var v12 string = "foo bar"
        _ = v11
        _ = v12
    }
    var v13 int64 = 0
    _ = len(v3)
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS1{F0: 829, F1: "foo bar"}
        v16 := GenS3{F0: "a"}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int64 = v13
    var v18 int = v2
    var v19 int8 = 0
    if (true && (0.1 <= 1.5)) {
        v19 -= v19
        switch "foo bar" {
        case "":
            for fi20 := 0; fi20 < 12; fi20++ {
                v21 := []int{(fi20 % (((v2 >> ((1) & 63))) | 1)), v18}
                if false {
                    v1 -= 127
                    v8 = (v8 / (((v8 | 2)) | 1))
                    for fi22 := 0; fi22 < 7; fi22++ {
                        v23 := GenS4{F0: -597, F1: 127, F2: true}
                        var v24 uint32 = 1
                        switch (v1 - -1) {
                        case 4294967295:
                            v25 := GenS5{F0: v24, F1: 88, F2: (^(v24))}
                            _ = v25
                        case -1:
                            v19 = -1
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                } else {
                    v26 := GenS4{F0: 0, F1: 127, F2: (!((v1 / ((v1) | 1)) < v18))}
                    fmt.Println(v1, v8)
                    v27 := GenS3{F0: ("" + "")}
                    _ = v26
                    _ = v27
                }
                v13 -= v13
                _ = fi20
                _ = v21
            }
        case "\n\t":
            var v28 int16 = 32767
            _ = v28
        default:
            v13 = (v17 / ((-1) | 1))
        }
    } else {
        var v29 int16 = 42
        if true {
            fmt.Println(v19)
        } else {
            v30 := GenS1{F0: 2147483647, F1: "a"}
            var v31 int = -9223372036854775808
            _ = v30
            _ = v31
        }
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
