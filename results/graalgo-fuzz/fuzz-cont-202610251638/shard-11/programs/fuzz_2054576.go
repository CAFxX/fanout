package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int64; F1 uint8; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 uint64; F2 string }
type GenS6 struct { F0 uint32; F1 bool; F2 int16 }
type GenS7 struct { F0 int64; F1 int }
func genFn0(p0 uint8, p1 uint, p2 uint8) bool {
    for fi6 := 0; fi6 < 3; fi6++ {
        p0 = p2
        _ = fi6
    }
    v7 := []int{9223372036854775807, -766, -2147483648}
    v7 = append(v7, -823)
    var v8 uint32 = uint32(p1)
    var v9 int64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return ((^(p2)) != (p0 | 128))
}
func genFn1(p0 int64) int64 {
    v10 := []int{-817, 690}
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS1{F0: 1, F1: 544, F2: 1}
        for fi13 := 0; fi13 < 9; fi13++ {
            p0 = (p0 ^ p0)
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    if len(v10) > 0 {
        v10[0] = -539
    }
    fmt.Println(p0)
    _ = p0
    _ = v10
    return (p0 / ((467) | 1))
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    v4 := make(map[string]string)
    v5 := []string{"hello", "a", "a"}
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    v14 := GenS2{F0: -9223372036854775808, F1: "\n\t"}
    var v15 int64 = 77
    if false {
        switch 162 {
        case 495:
            var v16 uint16 = 1
            _ = v16
        default:
            for fi17 := 0; fi17 < 4; fi17++ {
                v5 = append(v5, "a")
                _ = fi17
            }
        }
        var v18 float32 = 1e10
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 int32 = 42
            var v21 uint64 = 676
            var v22 string = "a"
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = v18
    } else {
        fmt.Println(v15)
        v1 += v1
        _ = len(v3)
    }
    v1 = (v1 + ((float32((0.1 * (3.14159 / 1.0))) / 3.14159) / 3.14159))
    if true {
        var v23 int8 = -128
        _ = v23
    }
    if (false || (!true)) {
        if false {
            for fi24 := 0; fi24 < 20; fi24++ {
                var v25 int16 = -906
                v26 := GenS3{F0: (v15 | 1), F1: 255, F2: "a"}
                if true {
                    var v27 uint16 = 2
                    v28 := GenS4{F0: true}
                    switch (fi24 - int(v15)) {
                    case -709:
                        var v29 uint = 18446744073709551615
                        _ = v29
                    case -534:
                        v30 := GenS2{F0: (v15 >> ((509) & 63)), F1: "\n\t"}
                        _ = v30
                    default:
                        v27 -= (v27 << ((613) & 63))
                    }
                    _ = v27
                    _ = v28
                }
                _ = fi24
                _ = v25
                _ = v26
            }
            if true {
                var v31 uint16 = 293
                v32 := GenS2{F0: v15, F1: ("" + ("" + "hello"))}
                _ = v31
                _ = v32
            }
            var v33 uint32 = 0
            _ = v33
        }
    }
    if true {
        if (!false) {
            v34 := GenS4{F0: ((!false) || true)}
            _ = v34
        }
    } else {
        fmt.Println(v1, v15)
        var v35 int = 2147483647
        fmt.Println(v1, v1, v1)
        _ = v35
    }
    var v36 float64 = 0.0
    v37 := GenS2{F0: (v15 << ((2) & 63)), F1: "\n\t"}
    v38 := GenS5{F0: 255, F1: 42, F2: ("a" + "\n\t")}
    switch "foo bar" {
    case "":
        v36 -= float64((v1 / (((v1 * (3.14159 * v1)) + 1e10) * -2.25)))
    }
    if ((-128 > 114) || true) {
        var v39 uint = 711
        _ = v39
    } else {
        var v40 uint8 = 1
        if (2 == 133) {
            v41 := GenS6{F0: 0, F1: true, F2: -1}
            v42 := GenS7{F0: (^(v15)), F1: 32767}
            v1 += v1
            _ = v41
            _ = v42
        }
        var v43 int8 = -1
        _ = v40
        _ = v43
    }
    v36 = 1e10
    var v44 bool = (120 != 32767)
    v15 = v15
    v5 = append(v5, "a")
    switch ("" + ("\n\t" + ((("hello" + ("a" + "hello")) + "\n\t") + ("" + ("hello" + "hello"))))) {
    case "foo bar":
        v44 = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v44
    _ = v5
}
