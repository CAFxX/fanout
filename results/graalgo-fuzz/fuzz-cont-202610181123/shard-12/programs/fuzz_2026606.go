package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
func genFn0(p0 string) uint {
    var v6 float32 = 3.14159
    v7 := GenS1{F0: 127}
    _ = p0
    _ = v6
    _ = v7
    return 431
}
func genFn1() int16 {
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint64 = 328
            fmt.Println(v10, fi9, fi8)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v11 uint64 = 2
    _ = v11
    return 0
}

func main() {
    var v1 float64 = 0.0
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]string)
    var v4 int32 = 0
    v5 := make(map[string]uint32)
    if false {
        v2 += -9223372036854775808
        v12 := GenS2{F0: 2147483648, F1: 781, F2: "a"}
        v2 = v2
        _ = v12
    }
    v13 := GenS2{F0: 0, F1: 0, F2: (("a" + "\n\t") + "foo bar")}
    if ((("hello" + "hello") > ((("a" + "foo bar") + "foo bar") + ("a" + "\n\t"))) || ((0.1 * ((v1 / v1) * (v1 * (v1 - v1)))) == (v1 + v1))) {
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS3{F0: 128, F1: "a"}
            v4 += (v4 >> ((1) & 63))
            v16 := GenS1{F0: -68}
            _ = fi14
            _ = v15
            _ = v16
        }
        v1 = v1
    } else {
        v4 -= 0
    }
    v4 = 295
    fmt.Println(v4)
    v17 := []int{2147483647}
    v4 -= v4
    v1 = v1
    var v18 uint64 = 18446744073709551615
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS4{F0: (v2 & 9223372036854775807)}
        _ = fi19
        _ = v20
    }
    switch 2147483647 {
    case 782:
        v21 := GenS5{F0: 65535}
        _ = v21
    case 127:
        var v22 int32 = v4
        _ = v22
    case 2147483647:
        for fi23 := 0; fi23 < 5; fi23++ {
            for fi24 := 0; fi24 < 14; fi24++ {
                if false {
                    v4 += (v4 >> ((336) & 63))
                    v17 = append(v17, -564)
                    fmt.Println(fi24)
                } else {
                    if false {
                        v5["a"] = 623
                        _ = len(v17)
                        v17 = append(v17, (fi24 | -9223372036854775808))
                    } else {
                        v2 = -37
                    }
                    v1 += 3.14159
                }
                var v25 int8 = 0
                _ = fi24
                _ = v25
            }
            fmt.Println(fi23)
            var v26 string = ("hello" + "a")
            _ = fi23
            _ = v26
        }
    default:
        var v27 int8 = -128
        _ = v27
    }
    var v28 int32 = 269
    _ = len(v17)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
