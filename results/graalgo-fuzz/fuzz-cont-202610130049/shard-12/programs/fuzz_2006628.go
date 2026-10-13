package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint8; F1 uint64 }
func genFn0(p0 bool, p1 uint64, p2 int16) int {
    fmt.Println(p0, p1, p1)
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 uint = 4294967295
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    v1 := []string{"a", (("\n\t" + "foo bar") + "foo bar")}
    var v2 string = "hello"
    v3 := make(map[uint32]uint32)
    var v4 float32 = 0.0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v7 := GenS1{F0: 0, F1: v2}
    var v8 int64 = 0
    v8 = v8
    switch v2 {
    case "foo bar":
        v2 = v2
    default:
        v1 = append(v1, "foo bar")
    }
    v9 := v3[2147483647]
    v10 := GenS1{F0: 42, F1: v2}
    var v11 string = ("a" + (v2 + v2))
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 uint16 = 0
        if (v4 >= 0.0) {
            for fi14 := 0; fi14 < 20; fi14++ {
                if ((0.0 + 1e10) <= 3.14159) {
                    switch v2 {
                    case "\n\t":
                        var v15 uint16 = v13
                        _ = v15
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v16 := v3[(v9 >> ((2) & 63))]
                        _ = v16
                    }
                    switch fi12 {
                    case 0:
                        var v17 float64 = 1.5
                        _ = v17
                    case 164:
                        switch (v2 + v11) {
                        case "a":
                            if true {
                                v18 := GenS2{F0: 42, F1: 319}
                                v19 := GenS2{F0: 1, F1: 2}
                                var v20 int = (fi14 & fi14)
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                        case "":
                            var v21 uint = 0
                            _ = v21
                        default:
                            v22 := GenS2{F0: uint8(v9), F1: 155}
                            _ = v22
                        }
                    case 42:
                        v9 += 2
                    }
                    if (1 != 0) {
                        if (!(-95 > 100)) {
                            if (fi12 == fi12) {
                                var v23 uint8 = 54
                                _ = v23
                            }
                        }
                        var v24 uint = 0
                        v8 += (v8 >> ((v24) & 63))
                        _ = v24
                    }
                } else {
                    _ = len(v3)
                }
                _ = fi14
            }
            fmt.Println(v2, fi12)
            delete(v3, v9)
        } else {
            var v25 uint = 534
            var v26 uint64 = 18446744073709551615
            _ = v25
            _ = v26
        }
        _ = fi12
        _ = v13
    }
    var v27 int8 = -38
    if (-1 == -605) {
        if (-2147483648 == -2147483648) {
            var v28 uint64 = 2
            v9 -= (v9 - 2147483647)
            _ = v28
        }
    }
    fmt.Println(v9, v11, v27)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
