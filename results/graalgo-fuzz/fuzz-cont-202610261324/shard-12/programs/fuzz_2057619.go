package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int16) string {
    fmt.Println(p0)
    p0 -= (p0 | 0)
    v7 := []int{-458, -9223372036854775808}
    v8 := GenS1{F0: -41, F1: ("a" + "\n\t"), F2: 2147483648}
    _ = p0
    _ = v7
    _ = v8
    return ("" + "hello")
}

func main() {
    v1 := []string{"a", "", "\n\t"}
    var v2 int8 = 42
    var v3 int16 = -820
    v4 := make(map[int]uint32)
    v5 := []uint32{916, 4294967295, 65535}
    v6 := make(map[uint32]uint32)
    switch (("a" + ("\n\t" + "\n\t")) + "foo bar") {
    case "foo bar":
        var v9 float32 = (3.14159 + (0.1 / -2.25))
        _ = v9
    case "":
        if (!(false && true)) {
            for fi10 := 0; fi10 < 19; fi10++ {
                v2 = (^(v2))
                _ = fi10
            }
        }
    default:
        var v11 float64 = float64(1.5)
        _ = v11
    }
    if true {
        v12 := GenS2{F0: v3}
        fmt.Println(v2)
        fmt.Println(v2)
        _ = v12
    } else {
        v2 = v2
        v2 += (v2 & (v2 & (v2 >> ((18446744073709551615) & 63))))
        var v13 int = -2147483648
        _ = v13
    }
    _ = len(v4)
    v14 := GenS2{F0: v3}
    var v15 string = ("a" + "foo bar")
    var v16 uint8 = 42
    v6[4294967295] = 42
    v15 = "a"
    v16 = v16
    var v17 int32 = -486
    if len(v5) > 0 {
        v5[0] = 732
    }
    if true {
        if (415 != 4294967295) {
            for fi18 := 0; fi18 < 15; fi18++ {
                if (fi18 == (fi18 - fi18)) {
                    v3 -= -823
                    v17 -= -1
                    v3 = v3
                }
                if false {
                    if false {
                        var v19 string = v15
                        _ = v19
                    } else {
                        var v20 int16 = v3
                        _ = v20
                    }
                    var v21 int16 = 360
                    _ = v21
                }
                var v22 int32 = v17
                _ = fi18
                _ = v22
            }
            if true {
                v23 := v4[9223372036854775807]
                var v24 uint = 1
                var v25 uint = 65535
                _ = v23
                _ = v24
                _ = v25
            }
        } else {
            v2 = v2
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
