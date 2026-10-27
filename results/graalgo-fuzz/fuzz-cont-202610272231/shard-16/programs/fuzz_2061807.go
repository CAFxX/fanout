package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint8, p1 int) uint32 {
    var v4 float64 = ((float64(1e10) + float64((-2.25 - ((0.0 - float32(1.5)) - 0.1)))) / 3.14159)
    v5 := []int{}
    var v6 int32 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 float64 = 1e10
    var v2 float64 = -2.25
    var v3 string = "foo bar"
    if (42 <= 2147483647) {
        v7 := GenS1{F0: (("foo bar" + "hello") + v3), F1: 0}
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := []int{fi8, fi8, (fi8 % ((fi8) | 1))}
        v9 = append(v9, (fi8 | (fi8 >> ((65535) & 63))))
        var v10 float32 = ((float32(v1) * float32((((3.14159 / (v2 + -2.25)) / v2) + ((((((-2.25 - v2) * v1) * v1) / (v1 + 0.1)) / -2.25) - (v2 * v1))))) * (0.1 * 1.5))
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 float32 = 1e10
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 bool = (738 <= 1)
        var v14 int32 = 38
        v15 := []int{fi12, (fi12 | fi12)}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v16 := []int{1}
    if true {
        var v17 int32 = -2147483648
        _ = v17
    } else {
        v2 += 1.5
        var v18 float64 = 0.1
        _ = v18
    }
    v19 := GenS1{F0: v3, F1: 42}
    if false {
        fmt.Println(v11, v3, v2)
        v16 = append(v16, -1)
    } else {
        var v20 float64 = 0.1
        _ = v20
    }
    if false {
        if false {
            v3 = ""
        } else {
            if false {
                v3 = "\n\t"
                if (((18446744073709551615 != 1) && true) && ("a" < v3)) {
                    fmt.Println(v2)
                    var v21 uint64 = 187
                    v11 = ((float32(0.1) - (v11 - ((1e10 * (-2.25 + float32(float64(((1e10 - 0.1) / (3.14159 - 0.0)))))) / (3.14159 / 0.1)))) / 0.1)
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 19; fi22++ {
                        var v23 int = 127
                        _ = fi22
                        _ = v23
                    }
                }
            } else {
                v3 = v3
                switch v3 {
                case "a":
                    v2 -= 0.0
                default:
                    switch v3 {
                    case "":
                        v1 += 0.0
                    default:
                        if (v2 > -2.25) {
                            v3 = v3
                            fmt.Println(v1)
                        }
                    }
                }
            }
            if false {
                if (v2 > v1) {
                    v2 = float64(v11)
                } else {
                    v1 += v1
                    if true {
                        fmt.Println(v3, v11)
                        v1 -= float64(v11)
                        v24 := GenS1{F0: "a", F1: 1}
                        _ = v24
                    } else {
                        fmt.Println(v3, v3, v11)
                    }
                    v25 := GenS2{F0: 500, F1: 4294967295}
                    _ = v25
                }
            } else {
                for fi26 := 0; fi26 < 2; fi26++ {
                    switch "hello" {
                    case "foo bar":
                        var v27 int32 = 1
                        _ = v27
                    case "a":
                        var v28 int64 = -1
                        _ = v28
                    case "\n\t":
                        v29 := GenS1{F0: v3, F1: 0}
                        _ = v29
                    }
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    _ = fi26
                }
            }
            v30 := GenS2{F0: -1, F1: 972}
            _ = v30
        }
        fmt.Println(v3, v3)
        var v31 int8 = 42
        _ = v31
    }
    if false {
        var v32 float32 = 1.5
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
