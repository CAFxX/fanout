package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int32 }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0(p0 int, p1 uint32, p2 uint) bool {
    v6 := []int{}
    v7 := GenS1{F0: p0, F1: 743, F2: -1}
    fmt.Println(p1, p0)
    if true {
        for fi8 := 0; fi8 < 8; fi8++ {
            if true {
                for fi9 := 0; fi9 < 19; fi9++ {
                    switch ("" + "foo bar") {
                    case "\n\t":
                        var v10 uint8 = 42
                        _ = v10
                    case "a":
                        switch "" {
                        case "":
                            var v11 float64 = (1e10 + 3.14159)
                            _ = v11
                        case "a":
                            var v12 int8 = 0
                            _ = v12
                        case "\n\t":
                            v13 := GenS1{F0: (p0 + (-(p0))), F1: 1, F2: -1}
                            v14 := GenS2{F0: false, F1: 6}
                            _ = v13
                            _ = v14
                        }
                    case "foo bar":
                        p2 = p2
                    default:
                        if (("hello" != "foo bar") && false) {
                            p1 = (p1 * 488)
                            var v15 uint8 = 128
                            _ = v15
                        } else {
                            var v16 uint16 = 2
                            _ = v16
                        }
                    }
                    fmt.Println(p1)
                    if true {
                        v6 = append(v6, fi9)
                        p2 = (p2 << (((p2 >> ((465) & 63))) & 63))
                        p2 += 65535
                    } else {
                        p0 -= 4294967295
                        if (145 <= 1) {
                            p2 = 2
                        }
                        p2 = 448
                    }
                    _ = fi9
                }
            }
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return false
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]string)
    v3 := []uint32{0, 170, 65535}
    v4 := make(map[int]string)
    var v5 uint8 = 127
    var v17 int64 = -9223372036854775808
    v3 = append(v3, 65535)
    if len(v3) > 0 {
        v3[0] = 65535
    }
    var v18 int = 128
    v1 = (v1 + v1)
    var v19 float32 = float32((v1 - v1))
    delete(v4, 9223372036854775807)
    var v20 uint32 = 1
    if (!true) {
        for fi21 := 0; fi21 < 20; fi21++ {
            v20 = v20
            _ = fi21
        }
    }
    v22 := GenS1{F0: (v18 / (((v18 + v18)) | 1)), F1: -751, F2: 42}
    for fi23 := 0; fi23 < 11; fi23++ {
        _ = len(v3)
        var v24 int = (^(v18))
        _ = fi23
        _ = v24
    }
    switch (v18 << ((536) & 63)) {
    case 128:
        v20 = (v20 ^ v20)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
