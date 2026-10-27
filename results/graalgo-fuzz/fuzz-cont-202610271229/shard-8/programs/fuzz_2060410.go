package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint32 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int32 = -2147483648
    var v3 uint32 = 42
    var v4 string = "foo bar"
    v5 := []string{v4, v4, ((((v4 + v4) + v4) + (v4 + v4)) + (v4 + "foo bar")), ("\n\t" + v4)}
    fmt.Println(v3)
    var v6 float32 = (0.1 + (((((0.1 - (0.1 - 1e10)) * 3.14159) / (1e10 * 1e10)) / (1.5 / 3.14159)) * 1e10))
    v7 := GenS1{F0: v2, F1: 1, F2: v3}
    var v8 float32 = (float32(-2.25) - ((1e10 + (1e10 * ((3.14159 / v6) + 3.14159))) - (1e10 / v6)))
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint16 = 0
        if ((255 == 255) || (v10 == v10)) {
            var v11 uint = 65535
            var v12 int = fi9
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 20; fi13++ {
            if ((fi13 ^ -2147483648) < fi13) {
                fmt.Println(fi9)
                var v14 float32 = 1e10
                var v15 int = (fi13 >> ((500) & 63))
                _ = len(v5)
                _ = v14
                _ = v15
            } else {
                var v16 uint8 = 1
                var v17 string = (("" + v4) + ((v4 + (v4 + ((v4 + ("\n\t" + v4)) + v4))) + v4))
                _ = v16
                _ = v17
            }
            v2 += 2147483647
            _ = fi13
        }
        _ = fi9
        _ = v10
    }
    var v18 int8 = 1
    if true {
        v18 = (v18 * (v18 / (((^(v18))) | 1)))
    }
    var v19 float64 = float64(v8)
    if true {
        v2 = (v2 & v2)
        var v20 float32 = (v6 / 1.0)
        v20 += 1e10
        _ = v20
    } else {
        fmt.Println(v3)
        if (false && (32767 <= 255)) {
            for fi21 := 0; fi21 < 10; fi21++ {
                for fi22 := 0; fi22 < 19; fi22++ {
                    if false {
                        if (-681 == (-(v2))) {
                            var v23 uint8 = 255
                            var v24 uint32 = 2147483647
                            v2 -= (v2 + 42)
                            _ = v23
                            _ = v24
                        }
                        var v25 int16 = -286
                        fmt.Println(v2)
                        _ = v25
                    }
                    fmt.Println(v8)
                    _ = fi22
                }
                _ = fi21
            }
        }
    }
    var v26 bool = (!(!(32767 < 161)))
    v3 -= (^(v3))
    if (!v26) {
        v4 = "hello"
    } else {
        var v27 int16 = 553
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
