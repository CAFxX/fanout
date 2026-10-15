package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 string; F1 int64; F2 uint }
func genFn0(p0 int64, p1 int32) int32 {
    p0 = -9223372036854775808
    switch -690 {
    case 9223372036854775807:
        v6 := GenS1{F0: 114, F1: 42}
        _ = v6
    case -9223372036854775808:
        v7 := []int{-838, -933, 127}
        _ = v7
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 uint = 303
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 float32 = 0.1
    v5 := []int{469, 42}
    switch 32768 {
    case -973:
        v5 = append(v5, 32768)
    case 253:
        v1 = v1
    case 2147483647:
        var v8 int8 = 127
        _ = v8
    default:
        switch ("a" + "foo bar") {
        case "":
            v1 = v1
        default:
            for fi9 := 0; fi9 < 12; fi9++ {
                if true {
                    v4 -= v4
                    if (("foo bar" + ("hello" + "a")) <= ("hello" + "\n\t")) {
                        if len(v5) > 0 {
                            v5[0] = (fi9 / (((fi9 * fi9)) | 1))
                        }
                    }
                    v4 -= v4
                }
                _ = fi9
            }
        }
    }
    fmt.Println(v4, v1)
    v1 = 255
    var v10 string = (((((("foo bar" + "hello") + "foo bar") + "a") + ("foo bar" + "hello")) + "hello") + "foo bar")
    if false {
        v10 = v10
        v4 -= ((float32(0.1) - 1.5) / 1.0)
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: 583, F1: 0}
        v13 := GenS2{F0: (v10 + ""), F1: 42, F2: (v1 & v1)}
        switch (fi11 / ((fi11) | 1)) {
        case 32767:
            var v14 uint32 = 2147483648
            _ = v14
        case 128:
            for fi15 := 0; fi15 < 10; fi15++ {
                v10 = (("hello" + v10) + ((v10 + "\n\t") + v10))
                _ = fi15
            }
        case 9223372036854775807:
            v4 += (((v4 / ((((v4 - 1.5) / 0.1) - ((float32(((0.0 + ((-2.25 / 1e10) / 1.5)) / 1.5)) * float32(1e10)) / v4)) * v4)) + 0.0) * v4)
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    var v16 uint16 = 931
    for fi17 := 0; fi17 < 5; fi17++ {
        v1 -= v1
        v5 = append(v5, fi17)
        v1 = v1
        _ = fi17
    }
    var v18 int16 = -838
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
