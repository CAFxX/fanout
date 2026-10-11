package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint16; F2 int64 }
func genFn0(p0 int32, p1 uint) string {
    if (-1 < 125) {
        var v4 float32 = 1e10
        p1 = (p1 / (((p1 & p1)) | 1))
        _ = v4
    } else {
        var v5 float64 = (-2.25 * 0.1)
        _ = v5
    }
    switch 32768 {
    case -81:
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 uint64 = 455
            _ = fi6
            _ = v7
        }
    case 128:
        p0 = (p0 - -2147483648)
    default:
        var v8 float32 = 0.0
        _ = v8
    }
    p1 += p1
    var v9 uint32 = 2147483647
    switch 42 {
    case -1:
        if (1 >= 42) {
            if false {
                v10 := []int{233, 922}
                _ = v10
            }
            var v11 uint16 = 606
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 bool = (-2.25 == 1e10)
                var v14 int8 = -119
                fmt.Println(fi12)
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        } else {
            v15 := GenS1{F0: ("\n\t" + ("\n\t" + ("\n\t" + "foo bar")))}
            v16 := []int{0, 32768}
            _ = v15
            _ = v16
        }
    case 1:
        fmt.Println(p1, p0)
    case -619:
        if ("hello" < ("foo bar" + ("" + "a"))) {
            for fi17 := 0; fi17 < 7; fi17++ {
                v9 -= (v9 - v9)
                _ = fi17
            }
            p1 -= p1
        }
    default:
        switch "hello" {
        case "":
            if true {
                var v18 int64 = 913
                _ = v18
            } else {
                p1 = p1
                var v19 int8 = -38
                p1 += p1
                _ = v19
            }
        default:
            v20 := []int{962}
            _ = v20
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return "hello"
}

func main() {
    v1 := []int{1, 42}
    var v2 int8 = -57
    v3 := make(map[int]uint32)
    v2 += -2
    v21 := GenS1{F0: ""}
    var v22 string = ("foo bar" + "foo bar")
    v22 = v22
    for fi23 := 0; fi23 < 6; fi23++ {
        fmt.Println(v22, fi23)
        v3[fi23] = 2
        _ = fi23
    }
    if (42 <= 1) {
        v24 := GenS2{F0: false, F1: 2147483648}
        _ = len(v1)
        _ = v24
    }
    v25 := GenS3{F0: -1, F1: 32767, F2: 1}
    switch v22 {
    case "hello":
        for fi26 := 0; fi26 < 1; fi26++ {
            v27 := GenS3{F0: -2147483648, F1: 2, F2: -1}
            _ = fi26
            _ = v27
        }
    case "a":
        switch ("a" + (v22 + v22)) {
        case "a":
            var v28 float64 = 1.5
            _ = v28
        case "":
            v22 = (v22 + v22)
        }
    }
    if len(v1) > 0 {
        v1[0] = -658
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
}
