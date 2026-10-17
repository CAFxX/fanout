package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 string }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int8) string {
    if (983 <= 18446744073709551615) {
        p0 += (p0 | (p0 & p0))
        p0 += (p0 / ((p0) | 1))
        v4 := []int{128, 127, 0, -206}
        _ = v4
    }
    v5 := GenS1{F0: (false && true), F1: 18446744073709551615, F2: "hello"}
    v6 := []int{869, 1, 450}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 float32 = 1e10
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[string]string)
    v3 := []int{-9223372036854775808, 32767, -1, 9223372036854775807}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v8 int64 = -145
    if false {
        var v9 bool = false
        v2["hello"] = ""
        v9 = v9
        _ = v9
    }
    v8 -= -903
    v2["foo bar"] = "hello"
    var v10 float64 = 0.0
    var v11 int8 = -128
    fmt.Println(v10, v10, v11)
    delete(v2, "foo bar")
    v8 = (^(v1))
    v1 += 42
    var v12 int = 1
    fmt.Println(v12)
    v1 += v8
    if (0.0 <= 0.1) {
        var v13 float64 = float64(3.14159)
        _ = v13
    } else {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int8 = -1
            v11 += v11
            if false {
                switch fi14 {
                case 1:
                    v8 = (v8 - (v8 % ((v8) | 1)))
                case 128:
                    if (100 == v11) {
                        v16 := GenS2{F0: 42}
                        _ = v16
                    } else {
                        switch fi14 {
                        case 713:
                            v3 = append(v3, (fi14 >> ((49) & 63)))
                        case 831:
                            fmt.Println(v10)
                        }
                        var v17 uint32 = 2147483647
                        _ = len(v3)
                        _ = v17
                    }
                case 681:
                    v15 = v15
                }
                v8 = v8
            }
            _ = fi14
            _ = v15
        }
    }
    _ = len(v2)
    v18 := v2[("a" + ("" + "hello"))]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v8
}
