package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int }
type GenS2 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint16, p1 bool) uint32 {
    var v6 bool = ((p1 || (4294967295 < 255)) || p1)
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 bool = v6
        v9 := []int{4294967295, (fi7 << ((18446744073709551615) & 63)), fi7, 128}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1(p0 string, p1 int8, p2 uint32) uint32 {
    var v10 uint8 = 1
    v10 -= v10
    var v11 uint32 = (p2 * (p2 / ((p2) | 1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 151
}

func main() {
    var v1 int64 = 249
    var v2 uint8 = 127
    v3 := make(map[string]uint32)
    var v4 int16 = 0
    v5 := make(map[uint32]bool)
    var v12 int = 1
    v1 -= v1
    v1 += v1
    fmt.Println(v4, v4, v2)
    v13 := GenS1{F0: 2, F1: 2147483647, F2: v12}
    var v14 int8 = 106
    v15 := []int{(v12 & (v12 % ((2147483647) | 1))), 431, 32768}
    if true {
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint32 = 2147483647
            var v18 int64 = int64(v2)
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 uint8 = (v2 | v2)
        v14 = (v14 * (v14 >> ((18446744073709551615) & 63)))
        _ = v19
    }
    if false {
        _ = len(v15)
        if (!false) {
            fmt.Println(v4, v12)
            var v20 string = ("a" + ("hello" + ((("" + "\n\t") + ((("a" + ((("\n\t" + ("hello" + "foo bar")) + "foo bar") + ("\n\t" + "hello"))) + "foo bar") + "hello")) + ("" + "a"))))
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 float32 = ((-2.25 / 3.14159) * -2.25)
                _ = fi21
                _ = v22
            }
            _ = v20
        } else {
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 int = (v12 / (((-(v12))) | 1))
                switch "a" {
                case "foo bar":
                    v1 += (v1 % (((v1 | (v1 | (v1 * (v1 + (v1 | (v1 - (-(v1))))))))) | 1))
                case "":
                    v25 := v3[("hello" + ((("" + "") + "foo bar") + ""))]
                    _ = v25
                case "\n\t":
                    v12 = (fi23 & (fi23 ^ 2147483647))
                }
                v26 := GenS2{F0: v2, F1: -644}
                _ = fi23
                _ = v24
                _ = v26
            }
            v4 -= v4
        }
    } else {
        var v27 float64 = (1.5 / ((-2.25 + 0.0) / (1e10 * 0.1)))
        for fi28 := 0; fi28 < 16; fi28++ {
            fmt.Println(fi28, fi28)
            v27 = (float64(((0.0 / 1.0) * 3.14159)) * v27)
            _ = fi28
        }
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
