package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64, p1 int) int16 {
    v6 := GenS1{F0: int32(p1), F1: 1}
    var v7 int = p1
    var v8 int8 = 30
    var v9 int32 = -1
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1000
}
func genFn1(p0 int64, p1 uint, p2 int32) uint8 {
    v10 := []int{-1, 341}
    switch ((((("foo bar" + "foo bar") + "hello") + "foo bar") + "a") + "hello") {
    case "":
        if true {
            fmt.Println(p2, p1, p2)
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 int = (fi11 << ((2) & 63))
                _ = fi11
                _ = v12
            }
        } else {
            var v13 int64 = -470
            v14 := GenS1{F0: (p2 / (((p2 | -1)) | 1)), F1: 240}
            fmt.Println(p0, p0)
            _ = v13
            _ = v14
        }
    default:
        if len(v10) > 0 {
            v10[0] = 204
        }
    }
    v15 := GenS2{F0: (p2 & (p2 ^ (p2 >> (((^(p1))) & 63)))), F1: ("" + "hello")}
    v10 = append(v10, 327)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v15
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 0
    v3 := make(map[int]string)
    var v4 float32 = 3.14159
    var v5 string = (("a" + "hello") + "\n\t")
    for fi16 := 0; fi16 < 11; fi16++ {
        v4 = (v4 - v4)
        var v17 int64 = v2
        _ = fi16
        _ = v17
    }
    fmt.Println(v5)
    v5 = ""
    v4 -= v4
    if (v2 <= v2) {
        var v18 int8 = 42
        v1[791] = ""
        _ = v18
    }
    if (2 >= genFn1(1, 0, 815)) {
        var v19 int16 = 77
        v5 = ""
        _ = v19
    } else {
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := GenS2{F0: -973, F1: "hello"}
            var v22 int64 = (^(v2))
            if (false && false) {
                var v23 uint8 = 1
                v24 := []int{fi20, 4294967295}
                v25 := GenS1{F0: 1, F1: 0}
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        v4 += (0.0 - (3.14159 + float32(((-2.25 - 0.1) + (((((0.0 + 1e10) / -2.25) / 0.1) + 1.5) + 0.0)))))
    }
    switch 584 {
    case 0:
        v26 := []int{}
        _ = v26
    case 4294967295:
        v27 := GenS3{F0: genFn0(383, -602)}
        _ = v27
    default:
        v2 = v2
    }
    if (659 != 32767) {
        switch ((((v5 + (v5 + v5)) + (((v5 + v5) + "foo bar") + v5)) + v5) + v5) {
        case "hello":
            for fi28 := 0; fi28 < 19; fi28++ {
                fmt.Println(v2)
                var v29 float64 = ((0.1 - 0.1) - 0.0)
                _ = fi28
                _ = v29
            }
        case "\n\t":
            v3[-9223372036854775808] = ("a" + (v5 + v5))
        }
    } else {
        var v30 uint = 255
        v31 := []int{}
        _ = v30
        _ = v31
    }
    var v32 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
