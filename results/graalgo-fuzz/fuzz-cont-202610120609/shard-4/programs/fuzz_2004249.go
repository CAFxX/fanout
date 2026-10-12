package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 string; F2 int8 }
type GenS3 struct { F0 uint }
func genFn0() uint32 {
    var v5 float32 = 3.14159
    switch "foo bar" {
    case "":
        v5 = float32(0.0)
    case "hello":
        var v6 int = -9223372036854775808
        _ = v6
    }
    if (18446744073709551615 > 436) {
        v7 := []int{127, 0, 32767}
        v8 := GenS1{F0: 127}
        fmt.Println(v5)
        _ = v7
        _ = v8
    } else {
        var v9 float64 = (((1.5 * -2.25) * float64(float32((1e10 + (0.1 + 0.0))))) / 1.5)
        _ = v9
    }
    _ = v5
    return 322
}

func main() {
    var v1 int8 = 42
    var v2 int32 = -198
    v3 := []int{-816}
    var v4 string = ("a" + "foo bar")
    v10 := GenS2{F0: 1, F1: ("foo bar" + v4), F2: (^(v1))}
    fmt.Println(v1, v2, v4)
    var v11 float32 = 0.0
    v4 = v4
    var v12 bool = (111 != 42)
    switch -67 {
    case 2147483647:
        v3 = append(v3, 42)
    case 432:
        var v13 string = ("" + ((v4 + "\n\t") + (v4 + v4)))
        _ = v13
    default:
        if ((128 > 0) && (v12 || v12)) {
            var v14 float32 = 1.5
            var v15 float64 = float64(((v14 - v11) - 1e10))
            _ = v14
            _ = v15
        }
    }
    v2 += v2
    if true {
        v3 = append(v3, 4294967295)
        var v16 int32 = v2
        v17 := GenS3{F0: 65535}
        _ = v16
        _ = v17
    } else {
        if (((v1 / (((v1 >> ((0) & 63))) | 1)) > (v1 + v1)) && (!(0 < -683))) {
            var v18 int64 = -9223372036854775808
            if v12 {
                v3 = append(v3, -106)
            } else {
                for fi19 := 0; fi19 < 14; fi19++ {
                    switch ((v4 + ((("a" + (v4 + v4)) + "") + "")) + "a") {
                    case "a":
                        for fi20 := 0; fi20 < 12; fi20++ {
                            switch (fi20 >> ((1) & 63)) {
                            case -392:
                                var v21 uint32 = 567
                                var v22 int8 = 124
                                fmt.Println(v4, v1)
                                _ = v21
                                _ = v22
                            }
                            _ = fi20
                        }
                    case "":
                        v2 -= v2
                    case "\n\t":
                        var v23 float32 = 3.14159
                        _ = v23
                    }
                    switch v4 {
                    case "":
                        v2 = v2
                    case "hello":
                        v24 := GenS1{F0: -31}
                        _ = v24
                    case "\n\t":
                        v3 = append(v3, (fi19 ^ 128))
                    default:
                        var v25 int64 = (v18 * v18)
                        _ = v25
                    }
                    _ = fi19
                }
                if v12 {
                    fmt.Println(v4, v4)
                    if (!(((float64((3.14159 - (0.0 * float32(float64((0.0 * 3.14159)))))) / float64((3.14159 * v11))) <= 1e10) || true)) {
                        var v26 float64 = 0.0
                        v2 -= v2
                        if (!v12) {
                            var v27 string = ((v4 + "") + ("hello" + (v4 + "\n\t")))
                            var v28 string = v27
                            _ = v27
                            _ = v28
                        } else {
                            _ = len(v3)
                        }
                        _ = v26
                    }
                }
                v11 = v11
            }
            _ = v18
        }
        v12 = (-1 < 1000)
        v11 += v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
