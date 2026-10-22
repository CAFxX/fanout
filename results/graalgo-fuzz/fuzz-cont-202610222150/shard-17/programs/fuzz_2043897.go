package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 string }
type GenS2 struct { F0 int64; F1 uint; F2 uint64 }

func main() {
    v1 := []uint32{2, 4294967295, 42}
    var v2 bool = ((!true) && (32767 != 0))
    var v3 float64 = 3.14159
    v4 := make(map[int]bool)
    v3 += -2.25
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 float32 = 0.1
            v2 = false
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: -204, F1: 4294967295, F2: ""}
        for fi9 := 0; fi9 < 18; fi9++ {
            v2 = v2
            fmt.Println(fi9, fi9)
            if v2 {
                var v10 float32 = -2.25
                if v2 {
                    fmt.Println(fi9, v3)
                    switch (-(fi9)) {
                    case 1:
                        if true {
                            var v11 int16 = 1000
                            _ = v11
                        }
                    case -9223372036854775808:
                        var v12 uint = 376
                        _ = v12
                    case -124:
                        if true {
                            var v13 int32 = 999
                            v4[fi9] = (2147483647 != v13)
                            _ = v13
                        }
                    }
                    v3 = 3.14159
                }
                _ = v10
            } else {
                if false {
                    var v14 uint = 18446744073709551615
                    _ = v14
                }
                if (fi9 > fi5) {
                    fmt.Println(v2, fi9)
                    v15 := v4[fi5]
                    _ = v15
                } else {
                    fmt.Println(fi5, v2, v2)
                }
            }
            _ = fi9
        }
        _ = fi5
        _ = v8
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := GenS2{F0: 9223372036854775807, F1: 317, F2: 18446744073709551615}
        if true {
            var v18 int = fi16
            switch v18 {
            case -1:
                var v19 string = "foo bar"
                _ = v19
            default:
                if v2 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v2 = (1 <= -128)
                    var v20 float32 = -2.25
                    _ = v20
                } else {
                    v18 += (^(fi16))
                    fmt.Println(v3, v2, v3)
                }
            }
            _ = v18
        } else {
            v2 = v2
        }
        _ = fi16
        _ = v17
    }
    switch 127 {
    case -330:
        v4[42] = false
    case -2147483648:
        var v21 int8 = 85
        _ = v21
    case -680:
        var v22 int = -9223372036854775808
        _ = v22
    default:
        if ((float64((0.1 - (0.0 * 0.0))) * ((v3 - v3) / v3)) >= 0.1) {
            fmt.Println(v3, v3)
        }
    }
    fmt.Println(v3, v3)
    v2 = (100 != -128)
    v1 = append(v1, 610)
    if (!v2) {
        _ = len(v4)
        v3 -= 0.1
    } else {
        if (!((255 > 255) || ((!(((1.5 - (3.14159 * 1.5)) < 1e10) && (2147483647 == 326))) || true))) {
            v2 = true
        }
    }
    v2 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
