package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 string; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint64; F1 bool }
func genFn0(p0 uint32) int16 {
    if (!(true || true)) {
        for fi6 := 0; fi6 < 1; fi6++ {
            fmt.Println(p0, p0)
            v7 := []int{128, 4294967295, (fi6 >> ((65535) & 63)), fi6}
            for fi8 := 0; fi8 < 4; fi8++ {
                p0 = p0
                if false {
                    p0 -= 42
                    v9 := GenS1{F0: 4294967295}
                    _ = v9
                }
                v10 := GenS2{F0: p0, F1: "\n\t", F2: 1}
                var v11 int32 = -157
                _ = fi8
                _ = v10
                _ = v11
            }
            _ = fi6
            _ = v7
        }
    } else {
        v12 := []int{-528}
        _ = v12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 int32 = 0
        _ = fi13
        _ = v14
    }
    v15 := []int{}
    v16 := GenS2{F0: (p0 + 4294967295), F1: "\n\t", F2: 0}
    _ = p0
    _ = v15
    _ = v16
    return 32767
}

func main() {
    var v1 string = "\n\t"
    var v2 float64 = (float64(0.1) + 1.5)
    v3 := make(map[int]int)
    var v4 int16 = 1
    v5 := make(map[uint32]string)
    fmt.Println(v1, v4, v2)
    if false {
        v2 = ((v2 / v2) / 3.14159)
        v1 = "hello"
    }
    if true {
        v17 := GenS3{F0: 18446744073709551615, F1: 503}
        v18 := v3[-470]
        _ = v17
        _ = v18
    } else {
        var v19 int8 = -1
        fmt.Println(v4, v1, v2)
        _ = v19
    }
    v20 := GenS4{F0: (v1 + "a"), F1: 0}
    for fi21 := 0; fi21 < 15; fi21++ {
        if false {
            fmt.Println(v1, v4, v2)
            var v22 uint = 783
            _ = v22
        }
        for fi23 := 0; fi23 < 13; fi23++ {
            switch int(v4) {
            case 1:
                v3[fi23] = (fi21 ^ fi21)
            case 42:
                v4 -= genFn0(2)
            default:
                var v24 string = "\n\t"
                _ = v24
            }
            if true {
                v25 := v3[fi21]
                switch (("" + v1) + v1) {
                case "\n\t":
                    var v26 float64 = 1e10
                    _ = v26
                case "":
                    var v27 int64 = -1
                    _ = v27
                default:
                    v25 += v25
                }
                switch (^(fi21)) {
                case -9223372036854775808:
                    var v28 float32 = 3.14159
                    _ = v28
                }
                _ = v25
            } else {
                if false {
                    fmt.Println(v4)
                }
            }
            _ = fi23
        }
        _ = fi21
    }
    var v29 uint16 = 32767
    v29 = (v29 ^ v29)
    v30 := GenS5{F0: 2, F1: false}
    if true {
        var v31 int16 = (v4 & genFn0(278))
        _ = v31
    }
    fmt.Println(v2, v2, v2)
    for fi32 := 0; fi32 < 5; fi32++ {
        fmt.Println(v29, fi32, v2)
        _ = fi32
    }
    var v33 uint8 = 0
    switch "\n\t" {
    case "\n\t":
        var v34 int64 = 0
        _ = v34
    case "a":
        var v35 float32 = 1e10
        _ = v35
    case "":
        v36 := []int{-1}
        _ = v36
    }
    v4 -= (-(v4))
    var v37 int = 574
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v37
    _ = v4
    _ = v5
}
