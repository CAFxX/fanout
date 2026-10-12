package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 bool; F2 int }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: 9223372036854775807, F1: 1}
        var v8 int64 = 882
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(fi9)
        var v10 uint32 = 444
        _ = fi9
        _ = v10
    }
    var v11 int16 = -43
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS2{F0: 35, F1: true, F2: -2147483648}
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 float32 = (1e10 + 0.0)
            _ = fi14
            _ = v15
        }
        if false {
            v16 := []int{1, fi12}
            for fi17 := 0; fi17 < 1; fi17++ {
                if (106 != 100) {
                    v18 := GenS2{F0: 1, F1: true, F2: 32768}
                    switch ("foo bar" + "foo bar") {
                    case "\n\t":
                        var v19 uint = 233
                        _ = v19
                    case "":
                        fmt.Println(fi17, fi12, fi12)
                    case "a":
                        fmt.Println(fi17, fi17)
                    }
                    fmt.Println(v11)
                    _ = v18
                }
                var v20 int16 = v11
                fmt.Println(v11, fi12)
                fmt.Println(v11)
                _ = fi17
                _ = v20
            }
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    switch 4294967295 {
    case 469:
        fmt.Println(v11)
    default:
        var v21 int32 = -560
        _ = v21
    }
    _ = v11
    return 333
}
func genFn1(p0 uint64, p1 int16, p2 uint) uint64 {
    p2 -= (p2 - p2)
    p1 = p1
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := GenS3{F0: 2, F1: 1}
        _ = fi22
        _ = v23
    }
    p1 = (p1 % (((p1 >> (((p2 & (p2 * (p2 - (^(p2)))))) & 63))) | 1))
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    var v1 bool = true
    var v2 string = ("hello" + ("hello" + "foo bar"))
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    v5 := []uint32{}
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    v24 := GenS3{F0: 255, F1: -269}
    v2 = "hello"
    v2 = "a"
    v1 = v1
    v2 = v2
    if len(v5) > 0 {
        v5[0] = 42
    }
    var v25 int16 = 1000
    _ = len(v5)
    if (!((v1 || true) || v1)) {
        var v26 int32 = 615
        fmt.Println(v1, v2)
        for fi27 := 0; fi27 < 12; fi27++ {
            if ((2 == 2) || (v1 || v1)) {
                v26 = v26
            }
            if (float64(float32(((float64(float32(-2.25)) - float64(float32(float64(0.1)))) * (1.5 / 1e10)))) > 0.0) {
                if (-1 >= (v25 | (^(v25)))) {
                    switch v2 {
                    case "foo bar":
                        var v28 float32 = ((1.5 * 1.5) / 1.0)
                        _ = v28
                    default:
                        fmt.Println(v26, fi27)
                    }
                    v29 := GenS1{F0: 1, F1: 128}
                    _ = v29
                }
            }
            var v30 int = (fi27 ^ (fi27 * 42))
            _ = fi27
            _ = v30
        }
        _ = v26
    }
    fmt.Println(v1, v1)
    for fi31 := 0; fi31 < 3; fi31++ {
        var v32 uint8 = 1
        _ = fi31
        _ = v32
    }
    v25 -= v25
    v33 := GenS2{F0: genFn1(4294967295, (v25 % (((v25 * v25)) | 1)), 269), F1: true, F2: 1}
    var v34 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
