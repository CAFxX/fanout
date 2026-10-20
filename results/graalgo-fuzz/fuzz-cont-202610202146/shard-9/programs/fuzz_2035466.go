package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int32; F1 int32; F2 int16 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 uint }
func genFn0() int8 {
    var v6 int16 = 0
    var v7 uint = 373
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int8 = 1
        var v10 int8 = 1
        _ = fi8
        _ = v10
        _ = v9
    }
    switch (("" + "") + (("hello" + "\n\t") + "foo bar")) {
    case "":
        v7 = v7
    default:
        var v11 uint32 = 971
        _ = v11
    }
    v12 := GenS1{F0: 0, F1: "\n\t"}
    _ = v12
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int = 42
    var v2 bool = (0 != 1)
    var v3 string = ""
    var v4 bool = (1 < 1)
    v5 := make(map[uint32]uint32)
    var v13 uint32 = 42
    fmt.Println(v13, v3, v3)
    v14 := GenS2{F0: -2147483648, F1: -93, F2: 42}
    var v15 int32 = 1
    fmt.Println(v13, v3)
    var v16 float32 = -2.25
    v17 := []int{}
    switch "foo bar" {
    case "hello":
        var v18 uint16 = uint16(v13)
        _ = v18
    case "":
        fmt.Println(v1)
    default:
        switch (v1 * -1) {
        case -934:
            var v19 bool = v4
            _ = v19
        case 32768:
            v13 -= v13
        case 315:
            fmt.Println(v1, v16, v3)
        default:
            fmt.Println(v13)
        }
    }
    var v20 int = 32768
    for fi21 := 0; fi21 < 2; fi21++ {
        if (320 >= 4294967295) {
            if (!v2) {
                v16 -= (-2.25 / 1.0)
                v22 := GenS3{F0: -128, F1: 65535}
                _ = v22
            } else {
                switch (v3 + (v3 + ("\n\t" + ("\n\t" + "foo bar")))) {
                case "":
                    for fi23 := 0; fi23 < 11; fi23++ {
                        v24 := GenS4{F0: (0 != -788)}
                        _ = fi23
                        _ = v24
                    }
                }
            }
            v16 += v16
        }
        v3 = (v3 + (((v3 + (v3 + "hello")) + v3) + (v3 + v3)))
        _ = fi21
    }
    v20 = 127
    fmt.Println(v20, v4, v3)
    v25 := GenS1{F0: 486, F1: v3}
    for fi26 := 0; fi26 < 2; fi26++ {
        v2 = false
        _ = fi26
    }
    v27 := GenS5{F0: 340, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
