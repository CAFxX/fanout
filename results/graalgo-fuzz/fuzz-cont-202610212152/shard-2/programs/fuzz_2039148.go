package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    v4 := []int{32768, 451, 2147483647, -606}
    v5 := []int{}
    switch ("" + ("\n\t" + "a")) {
    case "foo bar":
        var v6 float32 = 1.5
        _ = v6
    case "a":
        var v7 bool = (42 != 42)
        _ = v7
    default:
        if (-2.25 > 1.5) {
            if false {
                var v8 int8 = -103
                v9 := GenS1{F0: v8, F1: (("" + "") + "hello")}
                v10 := GenS1{F0: v8, F1: "hello"}
                _ = v10
                _ = v8
                _ = v9
            } else {
                var v11 string = ""
                fmt.Println(v11)
                _ = v11
            }
        } else {
            var v12 uint32 = 318
            _ = v12
        }
    }
    if (427 >= -354) {
        var v13 int = 0
        if false {
            v14 := GenS2{F0: 137}
            v13 += (-(v13))
            v15 := v1[2147483647]
            _ = v14
            _ = v15
        } else {
            if (4294967295 >= 247) {
                var v16 uint8 = 206
                v17 := v3[2147483648]
                _ = v16
                _ = v17
            } else {
                var v18 bool = (4294967295 > 42)
                switch v13 {
                case 128:
                    v13 += 904
                case 1:
                    v13 = (v13 % ((9223372036854775807) | 1))
                case 4294967295:
                    switch (v13 * v13) {
                    case -852:
                        v2 = append(v2, 2147483647)
                    case 1:
                        v18 = false
                    case -9223372036854775808:
                        var v19 uint32 = 617
                        _ = v19
                    }
                }
                _ = v18
            }
        }
        if (-28 > 79) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if false {
                fmt.Println(v13)
            } else {
                v20 := v3[2147483647]
                _ = v20
            }
        }
        _ = v13
    }
    var v21 int16 = -815
    v21 = v21
    if false {
        var v22 int16 = v21
        for fi23 := 0; fi23 < 19; fi23++ {
            v5 = append(v5, (fi23 % ((fi23) | 1)))
            v21 = -32768
            _ = fi23
        }
        _ = v22
    }
    var v24 uint64 = 4294967295
    switch ("" + "a") {
    case "foo bar":
        v2 = append(v2, 2147483647)
    default:
        if false {
            var v25 int = 623
            for fi26 := 0; fi26 < 1; fi26++ {
                v27 := GenS1{F0: 127, F1: "a"}
                _ = fi26
                _ = v27
            }
            _ = v25
        }
    }
    v24 -= (v24 ^ v24)
    v24 += 1
    fmt.Println(v24, v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
