package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0(p0 uint32, p1 uint32) int {
    p0 += (p1 - 42)
    var v4 float32 = -2.25
    switch -116 {
    case 128:
        switch -9223372036854775808 {
        case 579:
            for fi5 := 0; fi5 < 1; fi5++ {
                v4 += float32(0.0)
                var v6 uint16 = 2
                _ = fi5
                _ = v6
            }
        case -9223372036854775808:
            v4 = v4
        case -2147483648:
            for fi7 := 0; fi7 < 3; fi7++ {
                fmt.Println(v4, p0, p0)
                _ = fi7
            }
        default:
            p0 += (^(p0))
        }
    }
    var v8 int = -1
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return v8
}
func genFn1(p0 uint8, p1 int32) uint64 {
    if false {
        v9 := GenS1{F0: ((637 == 18446744073709551615) && false)}
        _ = v9
    } else {
        var v10 int16 = -775
        p1 = p1
        _ = v10
    }
    p1 = (p1 + (p1 >> ((42) & 63)))
    if false {
        if (18446744073709551615 > 18446744073709551615) {
            v11 := GenS2{F0: false, F1: ((3.14159 * -2.25) >= (3.14159 / float32(0.1)))}
            var v12 uint64 = 42
            _ = v11
            _ = v12
        }
        fmt.Println(p0, p0)
        var v13 int = 2147483647
        _ = v13
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 string = ""
    v2 := []uint32{}
    v3 := make(map[string]bool)
    var v14 uint = 18446744073709551615
    v1 = v1
    v15 := GenS3{F0: -128, F1: -391}
    var v16 bool = (4294967295 != 1)
    var v17 float32 = ((float32((0.1 + 1e10)) + (-2.25 / float32(-2.25))) * (-2.25 + float32((0.0 + 0.0))))
    var v18 float64 = 1e10
    if (888 != -1) {
        v16 = (!v16)
        v3[(("foo bar" + "\n\t") + "a")] = v16
        for fi19 := 0; fi19 < 1; fi19++ {
            v16 = (0 != 418)
            _ = fi19
        }
    }
    v17 -= (v17 - (v17 * 1.5))
    var v20 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
}
