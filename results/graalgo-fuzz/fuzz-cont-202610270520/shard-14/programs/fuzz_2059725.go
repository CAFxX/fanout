package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS4 struct { F0 int; F1 uint; F2 int }
func genFn0(p0 int32, p1 string) uint32 {
    v6 := GenS1{F0: 127}
    var v7 string = "hello"
    if false {
        var v8 uint64 = 18446744073709551615
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 742
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    var v4 uint8 = 255
    var v5 int32 = 4
    switch "a" {
    case "hello":
        v1 += genFn0((v5 - v5), "foo bar")
    case "a":
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(v5, v1, v5)
            _ = fi9
        }
    default:
        fmt.Println(v1)
    }
    v10 := GenS2{F0: 1, F1: -490}
    v11 := []int{128, 2147483647, 127}
    if false {
        var v12 uint = 730
        _ = v12
    } else {
        if len(v11) > 0 {
            v11[0] = 1
        }
        if ((v1 * genFn0(0, "hello")) >= v1) {
            for fi13 := 0; fi13 < 16; fi13++ {
                v1 -= (v1 >> ((0) & 63))
                if false {
                    v14 := GenS3{F0: (^(v1)), F1: v5, F2: v4}
                    v1 = (v1 >> ((4294967295) & 63))
                    if true {
                        switch "" {
                        case "a":
                            v5 = v5
                        case "foo bar":
                            if (0.0 >= (0.0 * 1.5)) {
                                v15 := GenS3{F0: (^(v1)), F1: (v5 & (v5 + v5)), F2: 128}
                                var v16 uint64 = 4294967295
                                _ = v15
                                _ = v16
                            } else {
                                var v17 float64 = -2.25
                                _ = v17
                            }
                        case "hello":
                            v18 := GenS4{F0: (fi13 - fi13), F1: 1, F2: fi13}
                            _ = v18
                        }
                    }
                    _ = v14
                } else {
                    var v19 uint32 = (v1 / (((v1 ^ v1)) | 1))
                    _ = v19
                }
                _ = fi13
            }
            var v20 uint64 = 65
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        v11 = append(v11, (fi21 ^ fi21))
        _ = fi21
    }
    v22 := GenS3{F0: (v1 ^ v1), F1: 42, F2: 222}
    v5 = v5
    v1 -= v1
    var v23 float32 = 1.5
    fmt.Println(v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
