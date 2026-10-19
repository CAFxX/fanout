package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS3 struct { F0 int16 }
func genFn0() uint16 {
    var v5 uint32 = 1
    if ("hello" != "foo bar") {
        v6 := []int{2147483647, 32767, 445}
        fmt.Println(v5)
        _ = v6
    } else {
        v5 += (v5 | 1)
        for fi7 := 0; fi7 < 11; fi7++ {
            switch (fi7 ^ fi7) {
            case 8:
                switch "a" {
                case "":
                    var v8 int64 = 236
                    _ = v8
                default:
                    for fi9 := 0; fi9 < 3; fi9++ {
                        fmt.Println(fi9)
                        v10 := GenS1{F0: (2 <= 39), F1: fi9}
                        fmt.Println(fi9, fi9)
                        _ = fi9
                        _ = v10
                    }
                }
            default:
                v11 := []int{32767, fi7, 128, fi7}
                _ = v11
            }
            if (617 != 0) {
                var v12 int = -9223372036854775808
                _ = v12
            } else {
                var v13 uint64 = 42
                _ = v13
            }
            var v14 uint = 0
            _ = fi7
            _ = v14
        }
        switch "foo bar" {
        case "":
            v15 := GenS2{F0: 97, F1: (^(v5)), F2: 0}
            _ = v15
        }
    }
    _ = v5
    return 368
}
func genFn1(p0 uint8, p1 uint8, p2 int) int16 {
    if true {
        v16 := []int{p2, (p2 % ((p2) | 1))}
        var v17 float32 = 1e10
        var v18 bool = false
        _ = v16
        _ = v17
        _ = v18
    }
    fmt.Println(p1, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = -9223372036854775808
    var v3 int32 = -1
    var v4 uint64 = 2
    v19 := GenS2{F0: 183, F1: 763, F2: 100}
    v4 = v4
    var v20 int64 = (-(v2))
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := GenS3{F0: -1}
        _ = len(v1)
        if true {
            v23 := []int{(fi21 / ((fi21) | 1)), (fi21 / ((int(v3)) | 1)), (fi21 | fi21), fi21}
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 int32 = 1
        v4 -= 4294967295
        v20 = v2
        _ = fi24
        _ = v25
    }
    v20 = v2
    v4 = v4
    var v26 bool = false
    if (92 > 0) {
        var v27 bool = (false && true)
        var v28 float64 = 0.0
        for fi29 := 0; fi29 < 10; fi29++ {
            fmt.Println(v26)
            _ = fi29
        }
        _ = v27
        _ = v28
    } else {
        v30 := GenS3{F0: 42}
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
}
