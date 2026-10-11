package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint32; F2 int8 }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0(p0 int16, p1 bool, p2 int32) bool {
    p0 += (-(p0))
    v7 := []int{4294967295, -228, -2147483648, -523}
    v7 = append(v7, 32768)
    if len(v7) > 0 {
        v7[0] = 496
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}

func main() {
    v1 := []uint32{101, 65535, 2147483647}
    var v2 float64 = float64(-2.25)
    v3 := []string{}
    var v4 bool = false
    v5 := make(map[string]uint32)
    var v6 int64 = -1
    var v8 int8 = 127
    v9 := GenS1{F0: 459, F1: 2147483648}
    v8 = (^(v8))
    v10 := v5[(("a" + ("foo bar" + "a")) + "a")]
    if (65535 != 1) {
        for fi11 := 0; fi11 < 17; fi11++ {
            if v4 {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi11
        }
        v10 = v10
        v12 := GenS2{F0: v4, F1: v10, F2: v8}
        _ = v12
    } else {
        switch "foo bar" {
        case "foo bar":
            if (!v4) {
                var v13 int = -573
                fmt.Println(v10, v8, v13)
                if (1 <= 1000) {
                    fmt.Println(v6, v2)
                    var v14 int8 = (v8 % ((v8) | 1))
                    switch (("a" + (("a" + "foo bar") + (("foo bar" + "\n\t") + (("" + "foo bar") + ("hello" + "foo bar"))))) + ("\n\t" + "foo bar")) {
                    case "hello":
                        var v15 uint64 = 2
                        _ = v15
                    case "foo bar":
                        if len(v1) > 0 {
                            v1[0] = v10
                        }
                    }
                    _ = v14
                } else {
                    v5[("" + "\n\t")] = 0
                }
                _ = v13
            }
        }
    }
    fmt.Println(v4, v2)
    _ = len(v5)
    v4 = (42 == -9223372036854775808)
    var v16 int16 = -1
    v17 := GenS3{F0: v8, F1: 2}
    var v18 uint64 = 1
    _ = len(v5)
    fmt.Println(v18, v4, v10)
    var v19 uint64 = (v18 - 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
