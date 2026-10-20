package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int16; F2 int64 }
type GenS3 struct { F0 int32; F1 int; F2 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    var v3 uint32 = 65535
    fmt.Println(v3)
    if false {
        v3 -= v3
        v4 := GenS1{F0: 205}
        v5 := []int{-473, 42, -2147483648}
        _ = v4
        _ = v5
    }
    fmt.Println(v3)
    var v6 uint8 = 74
    switch -2147483648 {
    case 32768:
        v7 := v2[4294967295]
        _ = v7
    case 158:
        switch "foo bar" {
        case "\n\t":
            if (float32(1e10) > -2.25) {
                var v8 uint8 = 15
                fmt.Println(v8, v3, v8)
                _ = v8
            }
        default:
            v3 = v3
        }
    }
    v3 = v3
    if (1 >= 1) {
        v9 := GenS1{F0: 4294967295}
        var v10 uint8 = v6
        _ = v10
        _ = v9
    } else {
        fmt.Println(v3, v3)
        var v11 string = "foo bar"
        _ = v11
    }
    var v12 bool = (true || true)
    switch ("hello" + (("\n\t" + "") + "foo bar")) {
    case "a":
        var v13 int32 = -1
        _ = v13
    case "foo bar":
        var v14 uint = 18446744073709551615
        _ = v14
    }
    v12 = ((!(-2.25 <= 1e10)) && true)
    fmt.Println(v3)
    var v15 int64 = 9223372036854775807
    switch "" {
    case "hello":
        v16 := GenS2{F0: v15, F1: 0, F2: v15}
        _ = v16
    case "a":
        if (true && (119 < 42)) {
            switch "" {
            case "":
                if true {
                    for fi17 := 0; fi17 < 14; fi17++ {
                        var v18 float64 = 3.14159
                        _ = fi17
                        _ = v18
                    }
                } else {
                    if v12 {
                        v19 := GenS3{F0: 0, F1: 127, F2: 91}
                        if v12 {
                            if true {
                                var v20 string = ""
                                v12 = (v12 && v12)
                                _ = v20
                            } else {
                                v21 := GenS1{F0: 812}
                                v22 := GenS3{F0: -350, F1: 4294967295, F2: -128}
                                _ = v21
                                _ = v22
                            }
                            switch "foo bar" {
                            case "foo bar":
                                v12 = v12
                            case "\n\t":
                                fmt.Println(v6)
                            case "hello":
                                if false {
                                    v6 = v6
                                    v12 = (4294967295 > 0)
                                }
                            }
                            v23 := []int{42, -2147483648, 128}
                            _ = v23
                        }
                        _ = v19
                    }
                    v6 += (v6 | v6)
                    var v24 int16 = 613
                    _ = v24
                }
            case "\n\t":
                var v25 string = "a"
                _ = v25
            case "hello":
                var v26 int8 = -104
                _ = v26
            }
        }
    }
    switch 4294967295 {
    case 912:
        v3 -= (v3 + 4294967295)
    case -9223372036854775808:
        for fi27 := 0; fi27 < 11; fi27++ {
            _ = len(v2)
            v15 = (v15 * (v15 | v15))
            var v28 int8 = -1
            _ = fi27
            _ = v28
        }
    default:
        var v29 int8 = -128
        _ = v29
    }
    if true {
        v15 += v15
        v30 := GenS2{F0: 9223372036854775807, F1: -324, F2: (v15 - (v15 * (-(v15))))}
        if ("\n\t" == ("hello" + "")) {
            var v31 float32 = 1e10
            _ = v31
        } else {
            v32 := GenS2{F0: (v15 | v15), F1: 1, F2: v15}
            _ = v32
        }
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v6
}
