package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int64 }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0(p0 uint16, p1 int64, p2 uint16) int32 {
    p1 = 12
    v5 := []int{870, 32767, 32767}
    p1 -= -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2147483647
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[string]int)
    var v3 int64 = 425
    v4 := make(map[uint32]string)
    v6 := GenS1{F0: 439, F1: 1, F2: v3}
    v7 := []int{}
    switch "hello" {
    case "":
        _ = len(v7)
    case "hello":
        for fi8 := 0; fi8 < 11; fi8++ {
            switch (fi8 * 920) {
            case -9223372036854775808:
                var v9 int8 = -1
                _ = v9
            case -729:
                var v10 float64 = ((-2.25 / (1.5 / float64(0.1))) / -2.25)
                _ = v10
            default:
                if true {
                    if (2147483647 == fi8) {
                        _ = len(v2)
                    }
                    fmt.Println(v1, v1, v1)
                }
            }
            if false {
                var v11 float32 = 3.14159
                v12 := GenS2{F0: fi8, F1: 42}
                if false {
                    var v13 string = "foo bar"
                    var v14 uint8 = 246
                    _ = v13
                    _ = v14
                }
                _ = v11
                _ = v12
            }
            _ = fi8
        }
    default:
        v15 := GenS1{F0: 801, F1: 2, F2: v3}
        _ = v15
    }
    var v16 uint = 211
    var v17 int = 127
    var v18 int32 = 1
    var v19 int = v17
    v16 = (v16 + 255)
    var v20 uint16 = 42
    switch v17 {
    case 32768:
        if false {
            if true {
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 float64 = 1e10
                    v18 += (v18 >> ((313) & 63))
                    v17 = -1
                    _ = fi21
                    _ = v22
                }
                var v23 uint = 267
                _ = v23
            }
            v24 := GenS1{F0: v19, F1: 42, F2: 329}
            _ = v24
        } else {
            fmt.Println(v19, v3)
        }
    case 128:
        var v25 bool = (!false)
        _ = v25
    }
    if len(v7) > 0 {
        v7[0] = -9223372036854775808
    }
    var v26 bool = ((v18 | v18) == (v18 * genFn0((v20 ^ v20), 42, (v20 / (((v20 - v20)) | 1)))))
    v1 += ((1.5 * (3.14159 + v1)) - (v1 + (-2.25 / v1)))
    var v27 string = "a"
    for fi28 := 0; fi28 < 11; fi28++ {
        v18 += (v18 + (v18 / ((v18) | 1)))
        _ = fi28
    }
    switch "a" {
    case "":
        fmt.Println(v26, v17, v19)
    case "foo bar":
        for fi29 := 0; fi29 < 20; fi29++ {
            v17 = (v19 >> ((v16) & 63))
            for fi30 := 0; fi30 < 12; fi30++ {
                v16 -= (^(v16))
                _ = fi30
            }
            for fi31 := 0; fi31 < 17; fi31++ {
                _ = len(v7)
                switch (fi29 % ((fi31) | 1)) {
                case 42:
                    var v32 uint32 = 1
                    _ = v32
                }
                _ = fi31
            }
            _ = fi29
        }
    default:
        var v33 uint64 = 217
        _ = v33
    }
    fmt.Println(v18, v20, v26)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
