package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 int32; F2 int }
func genFn0(p0 int, p1 uint64, p2 uint) uint {
    var v6 uint8 = 42
    var v7 uint16 = 601
    v8 := GenS1{F0: 32767, F1: int64(p1)}
    if (-2147483648 != 127) {
        var v9 string = "foo bar"
        fmt.Println(v9, v7)
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 int8, p1 bool, p2 string) bool {
    if ((p0 / ((p0) | 1)) < (p0 << ((1) & 63))) {
        for fi10 := 0; fi10 < 19; fi10++ {
            p2 = ("" + ((p2 + p2) + (p2 + "hello")))
            if false {
                if p1 {
                    switch p2 {
                    case "hello":
                        fmt.Println(fi10, p2, p2)
                    case "foo bar":
                        v11 := GenS2{F0: 1, F1: 811}
                        _ = v11
                    case "":
                        var v12 uint64 = 4294967295
                        _ = v12
                    default:
                        p0 = (p0 * (p0 + (p0 * (p0 & (p0 + (p0 - p0))))))
                    }
                }
                switch ("\n\t" + p2) {
                case "hello":
                    if false {
                        v13 := []int{(-(fi10))}
                        for fi14 := 0; fi14 < 4; fi14++ {
                            p1 = true
                            if (!(-9223372036854775808 != 42)) {
                                p1 = (true || (!true))
                            }
                            var v15 int = 2147483647
                            _ = fi14
                            _ = v15
                        }
                        _ = v13
                    } else {
                        v16 := GenS1{F0: -584, F1: 1}
                        _ = v16
                    }
                }
            }
            p1 = p1
            _ = fi10
        }
    }
    v17 := GenS3{F0: 4294967295}
    switch p2 {
    case "":
        if (!p1) {
            if (!(1 >= 156)) {
                v18 := GenS4{F0: 9223372036854775807}
                p0 = p0
                _ = v18
            } else {
                if (!p1) {
                    fmt.Println(p0)
                    var v19 uint64 = 582
                    _ = v19
                }
            }
            fmt.Println(p0, p2, p0)
        }
    default:
        v20 := GenS4{F0: 455}
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return p1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -799
    var v3 float32 = ((((-2.25 * 0.0) - float32((0.0 + -2.25))) / 1.5) * (1.5 / 1.0))
    var v4 int16 = 1000
    var v5 int16 = -51
    fmt.Println(v4)
    v3 -= v3
    fmt.Println(v5)
    switch int(v2) {
    case 127:
        var v21 string = "\n\t"
        _ = v21
    }
    v22 := GenS5{F0: 96, F1: (v2 + -2147483648), F2: -849}
    fmt.Println(v5)
    v23 := v1[("a" + "")]
    v3 = (float32(1e10) * v3)
    switch 42 {
    case 9223372036854775807:
        var v24 int = 9223372036854775807
        _ = v24
    case 32767:
        v4 = v5
    default:
        var v25 uint64 = 1
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
