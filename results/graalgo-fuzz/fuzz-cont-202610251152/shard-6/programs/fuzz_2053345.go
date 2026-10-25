package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint64; F1 uint64 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint32 = 852
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 uint32 = (^(v7))
            fmt.Println(fi6)
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    if (127 != 42) {
        var v10 bool = false
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v10)
        case "hello":
            fmt.Println(v10)
        }
        v11 := []int{-898, -1}
        _ = v10
        _ = v11
    }
    var v12 float32 = 0.0
    fmt.Println(v12)
    _ = v12
    return 128
}

func main() {
    v1 := []int{128, 9223372036854775807, -2147483648}
    var v2 uint32 = 600
    v3 := make(map[string]int)
    var v4 int32 = 2147483647
    v5 := []uint32{v2}
    v2 += (v2 - v2)
    v13 := GenS1{F0: 672, F1: true}
    v2 += (^(v2))
    var v14 uint64 = 18446744073709551615
    var v15 uint64 = 4294967295
    v3[("hello" + "foo bar")] = 238
    if (42 < 32767) {
        for fi16 := 0; fi16 < 14; fi16++ {
            switch (("a" + "a") + "foo bar") {
            case "foo bar":
                v17 := GenS1{F0: 18446744073709551615, F1: (!false)}
                _ = v17
            case "":
                for fi18 := 0; fi18 < 8; fi18++ {
                    v2 = (v2 / ((v2) | 1))
                    var v19 uint64 = (v15 + (v15 << ((110) & 63)))
                    var v20 uint16 = uint16(v2)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            case "a":
                v15 += (v15 - v14)
            }
            if (true && false) {
                v2 -= (v2 & (v2 >> ((1) & 63)))
            } else {
                if false {
                    var v21 string = "\n\t"
                    _ = v21
                }
            }
            fmt.Println(v14)
            _ = fi16
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        v22 := GenS1{F0: 329, F1: true}
        for fi23 := 0; fi23 < 14; fi23++ {
            v3[""] = fi23
            fmt.Println(v4, v15)
            _ = fi23
        }
        _ = v22
    }
    v24 := GenS2{F0: (v15 + 4294967295), F1: (v4 & v4)}
    v25 := GenS1{F0: 42, F1: true}
    var v26 uint16 = 851
    _ = len(v3)
    v27 := GenS3{F0: (v26 / (((v26 & (v26 | v26))) | 1))}
    if false {
        fmt.Println(v14)
    }
    if (v26 == 575) {
        var v28 uint32 = v2
        _ = v28
    } else {
        v29 := GenS4{F0: -32768, F1: 18}
        v3["a"] = 32768
        _ = v29
    }
    v4 += (v4 * (v4 & 2147483647))
    if (1.5 <= (((3.14159 - 0.0) + (1.5 - -2.25)) + ((1e10 * 1e10) * (1.5 + ((1.5 + (-2.25 - -2.25)) + 0.0))))) {
        var v30 uint = 4294967295
        fmt.Println(v4, v26, v30)
        _ = v30
    }
    v31 := GenS5{F0: v15, F1: v15}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
