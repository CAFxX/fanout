package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int16; F1 int16; F2 int64 }
type GenS3 struct { F0 int64; F1 int8; F2 uint }
func genFn0(p0 uint, p1 string) int32 {
    for fi6 := 0; fi6 < 19; fi6++ {
        switch ((p1 + "") + p1) {
        case "\n\t":
            var v7 uint64 = 803
            _ = v7
        case "":
            fmt.Println(fi6, p0, p0)
        default:
            var v8 uint8 = 2
            _ = v8
        }
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS1{F0: 1, F1: 825}
            _ = fi9
            _ = v10
        }
        v11 := GenS2{F0: 1, F1: 0, F2: -445}
        _ = fi6
        _ = v11
    }
    p0 = 0
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := []int{(fi12 % ((fi12) | 1)), (fi12 % ((fi12) | 1)), (fi12 / ((772) | 1))}
        _ = fi12
        _ = v13
    }
    p0 = p0
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    var v1 bool = (-2.25 < 0.1)
    var v2 uint64 = 18446744073709551615
    var v3 uint16 = 130
    v4 := make(map[string]string)
    var v5 float32 = 1.5
    var v14 uint8 = 128
    if (-52 < -128) {
        v15 := []int{32768, 128, 32768}
        _ = v15
    } else {
        fmt.Println(v14, v5, v5)
        var v16 int16 = 1
        _ = v16
    }
    v17 := GenS1{F0: 2147483647, F1: -9223372036854775808}
    for fi18 := 0; fi18 < 11; fi18++ {
        if ((fi18 + (fi18 % ((173) | 1))) >= (fi18 % ((fi18) | 1))) {
            switch ("\n\t" + ("" + "")) {
            case "a":
                for fi19 := 0; fi19 < 10; fi19++ {
                    v2 = v2
                    v3 -= 791
                    _ = fi19
                }
            default:
                fmt.Println(v3, v14, v2)
            }
        }
        var v20 int16 = 42
        _ = fi18
        _ = v20
    }
    v21 := []int{-9223372036854775808, -2147483648, 42}
    v14 += 90
    if (!v1) {
        v22 := GenS3{F0: 42, F1: 127, F2: 1}
        _ = v22
    }
    var v23 uint = 0
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    var v24 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
