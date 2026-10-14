package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint8; F2 uint }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0(p0 bool, p1 uint64) bool {
    fmt.Println(p1)
    var v4 int64 = -1
    if p0 {
        p1 = (p1 << ((0) & 63))
    } else {
        var v5 int = 0
        var v6 uint = 448
        var v7 uint16 = 2
        _ = v5
        _ = v6
        _ = v7
    }
    v4 -= (v4 % (((v4 * 1)) | 1))
    if ((!false) || true) {
        p1 = (p1 + 4294967295)
        var v8 int32 = -1
        _ = v8
    } else {
        p0 = (1.5 < (float32(-2.25) / 1e10))
        var v9 uint64 = 18446744073709551615
        if (((float32((0.1 * (1e10 * (-2.25 / 1.0)))) - 1.5) <= 0.1) || p0) {
            p1 = p1
            v10 := GenS1{F0: ("a" + ("foo bar" + "a"))}
            _ = v10
        } else {
            var v11 uint64 = 2
            var v12 uint64 = 1
            var v13 float64 = 0.1
            _ = v11
            _ = v12
            _ = v13
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    return (0 == 991)
}
func genFn1(p0 string, p1 string) int32 {
    p1 = ""
    var v14 string = (p0 + p0)
    var v15 string = ("" + ("a" + ((("hello" + (p1 + ("" + p1))) + "") + p1)))
    if true {
        fmt.Println(p0, p1)
        v15 = ""
        p1 = v14
    } else {
        p0 = "foo bar"
        v16 := []int{128, 32768, -573}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -259
    var v3 uint = 283
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            v19 := GenS2{F0: 42, F1: 127, F2: 65535}
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v20 uint8 = 21
    v21 := GenS3{F0: false, F1: 65535}
    for fi22 := 0; fi22 < 11; fi22++ {
        switch (^(fi22)) {
        case 602:
            v3 += 18446744073709551615
        case -358:
            var v23 int64 = 0
            _ = v23
        case 128:
            fmt.Println(v2, fi22)
        default:
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 int64 = -1
                if false {
                    var v26 int8 = -1
                    _ = len(v1)
                    _ = v26
                } else {
                    var v27 float64 = 1e10
                    v20 = (^(v20))
                    switch -1 {
                    case 32767:
                        v28 := []int{}
                        _ = v28
                    case 32768:
                        v29 := []int{fi22}
                        _ = v29
                    }
                    v30 := GenS2{F0: 127, F1: v20, F2: 0}
                    _ = v27
                    _ = v30
                }
                var v31 bool = true
                _ = fi24
                _ = v25
                _ = v31
            }
        }
        _ = fi22
    }
    v1[4294967295] = 2147483647
    var v32 uint32 = 4294967295
    var v33 int8 = -128
    var v34 uint = (v3 * v3)
    v2 = v2
    var v35 int64 = 9223372036854775807
    var v36 uint8 = v20
    for fi37 := 0; fi37 < 4; fi37++ {
        fmt.Println(v3)
        _ = fi37
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        fmt.Println(v35, v35)
        var v39 uint8 = (v36 % ((2) | 1))
        _ = fi38
        _ = v39
    }
    v40 := GenS2{F0: (v33 / ((v33) | 1)), F1: uint8(v32), F2: 49}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v40
}
