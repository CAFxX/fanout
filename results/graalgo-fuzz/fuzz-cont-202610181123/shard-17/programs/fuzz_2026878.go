package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS3 struct { F0 int16; F1 int32; F2 uint8 }
func genFn0(p0 uint16) string {
    v6 := []int{4294967295, -553, -613, 0}
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint64 = 18446744073709551615
        v9 := GenS1{F0: 20, F1: false}
        _ = fi7
        _ = v8
        _ = v9
    }
    if len(v6) > 0 {
        v6[0] = -379
    }
    _ = p0
    _ = v6
    return "hello"
}
func genFn1(p0 uint, p1 uint64) uint8 {
    var v10 uint8 = 0
    var v11 string = ((("a" + "") + ("foo bar" + "a")) + ("" + "\n\t"))
    var v12 uint64 = p1
    v13 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return v10
}

func main() {
    var v1 uint32 = 667
    v2 := []string{"foo bar"}
    var v3 int = 4294967295
    v4 := []int{v3, -2147483648, (v3 | (v3 + 128))}
    var v5 uint64 = 2
    var v14 float64 = 1e10
    if len(v4) > 0 {
        v4[0] = v3
    }
    fmt.Println(v3, v1)
    v5 += (v5 * 1)
    fmt.Println(v1, v5)
    v1 -= v1
    v1 = v1
    var v15 int64 = 708
    v14 += float64((-2.25 + ((-2.25 + (0.1 + 1e10)) - 1.5)))
    switch "\n\t" {
    case "foo bar":
        v16 := GenS2{F0: 65535, F1: 255, F2: 1}
        _ = v16
    case "":
        v5 = (^(v5))
    case "a":
        var v17 int8 = -32
        _ = v17
    }
    if true {
        if (false && false) {
            v3 += v3
        } else {
            for fi18 := 0; fi18 < 6; fi18++ {
                switch 128 {
                case -9223372036854775808:
                    switch (v3 >> ((945) & 63)) {
                    case 9223372036854775807:
                        v3 = (v3 * (v3 * (^(fi18))))
                    case 2147483647:
                        if len(v4) > 0 {
                            v4[0] = 42
                        }
                    case 127:
                        fmt.Println(v3, v3)
                    }
                }
                switch 0 {
                case 0:
                    v3 = 4294967295
                case -197:
                    var v19 int16 = 0
                    _ = v19
                default:
                    var v20 int = (fi18 % (((fi18 << ((831) & 63))) | 1))
                    _ = v20
                }
                var v21 int8 = -122
                _ = fi18
                _ = v21
            }
            fmt.Println(v5)
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v22 := GenS3{F0: -795, F1: 194, F2: 1}
        _ = v22
    }
    fmt.Println(v3, v1, v3)
    var v23 int = -685
    var v24 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
