package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint }
type GenS3 struct { F0 int64 }
func genFn0(p0 string, p1 uint32, p2 int) int64 {
    if true {
        for fi4 := 0; fi4 < 20; fi4++ {
            var v5 uint32 = (p1 >> ((0) & 63))
            _ = fi4
            _ = v5
        }
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: 127}
        if (p1 != (^(p1))) {
            v8 := GenS1{F0: 255}
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 bool = (false || (42 <= 1))
                if v10 {
                    v11 := GenS1{F0: 2}
                    p0 = "hello"
                    _ = v11
                } else {
                    v12 := []int{-2147483648, (fi9 + 124), (fi6 - (fi9 & p2))}
                    _ = v12
                }
                _ = fi9
                _ = v10
            }
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1() int16 {
    var v13 bool = (!(false || (2 <= 192)))
    v13 = v13
    v14 := []int{127, 1}
    _ = v13
    _ = v14
    return 709
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{-501}
    var v3 string = ""
    v3 = "\n\t"
    switch v3 {
    case "\n\t":
        v15 := GenS2{F0: 2147483647, F1: 2147483647, F2: 4294967295}
        _ = v15
    case "a":
        fmt.Println(v3)
    case "":
        v3 = v3
    default:
        if (0.1 == 1e10) {
            v3 = "a"
            delete(v1, v3)
            v3 = "a"
        } else {
            v16 := GenS2{F0: 154, F1: 4294967295, F2: 0}
            fmt.Println(v3)
            var v17 int8 = 127
            _ = v16
            _ = v17
        }
    }
    var v18 float32 = 0.0
    var v19 int64 = 9223372036854775807
    v2 = append(v2, -875)
    var v20 int = -9223372036854775808
    var v21 uint32 = 42
    _ = len(v2)
    v22 := GenS2{F0: -2147483648, F1: v21, F2: 18446744073709551615}
    if (665 > 782) {
        var v23 int32 = 231
        _ = v23
    }
    fmt.Println(v21)
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 float64 = (((0.0 / (0.1 + float64(v18))) - 1e10) * (((0.0 + (1e10 + -2.25)) / (1e10 - (-2.25 - 1e10))) / float64(float32(1e10))))
        _ = fi24
        _ = v25
    }
    v21 += v21
    var v26 int64 = (v19 << ((2) & 63))
    if true {
        v20 -= -2147483648
    } else {
        var v27 float32 = v18
        var v28 uint32 = 65535
        _ = len(v2)
        _ = v27
        _ = v28
    }
    v29 := GenS3{F0: v19}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v29
    _ = v3
}
