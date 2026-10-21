package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0() uint8 {
    v6 := []int{128, 1, -568}
    v7 := GenS1{F0: 42, F1: 127}
    if true {
        if (!(0.0 == 1.5)) {
            var v8 uint64 = 2
            _ = v8
        } else {
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(fi9)
                switch "\n\t" {
                case "foo bar":
                    for fi10 := 0; fi10 < 19; fi10++ {
                        fmt.Println(fi9, fi10)
                        var v11 uint = 42
                        fmt.Println(v11, fi9, v11)
                        _ = fi10
                        _ = v11
                    }
                case "\n\t":
                    var v12 int = (fi9 * fi9)
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 16; fi13++ {
                        v14 := GenS1{F0: 2147483647, F1: (fi13 & (^(fi9)))}
                        _ = fi13
                        _ = v14
                    }
                }
                if true {
                    if len(v6) > 0 {
                        v6[0] = fi9
                    }
                }
                _ = fi9
            }
            var v15 uint8 = 71
            _ = v15
        }
        var v16 uint32 = 195
        _ = v16
    }
    var v17 uint8 = 204
    _ = v17
    _ = v6
    _ = v7
    return (v17 / (((^(v17))) | 1))
}

func main() {
    var v1 uint16 = 42
    var v2 int64 = 9223372036854775807
    var v3 int8 = -128
    var v4 uint16 = v1
    v5 := make(map[uint32]string)
    fmt.Println(v3)
    if true {
        switch 127 {
        case 32767:
            var v18 float32 = 0.1
            _ = v18
        default:
            switch -1 {
            case -9223372036854775808:
                for fi19 := 0; fi19 < 11; fi19++ {
                    switch (fi19 + fi19) {
                    case 327:
                        v3 = v3
                    case 32767:
                        v3 += (v3 & -128)
                    case 734:
                        if (!(v1 == 32767)) {
                            if false {
                                v2 += (v2 >> ((1) & 63))
                                var v20 int = fi19
                                var v21 string = ("a" + "")
                                _ = v20
                                _ = v21
                            } else {
                                v22 := GenS2{F0: 1, F1: (^(v4))}
                                delete(v5, 2147483648)
                                _ = v22
                            }
                        } else {
                            v1 = 42
                            for fi23 := 0; fi23 < 6; fi23++ {
                                v1 = (v1 ^ (^(v4)))
                                _ = fi23
                            }
                            var v24 uint32 = 2147483647
                            _ = v24
                        }
                    }
                    _ = fi19
                }
            case -130:
                v2 = 168
            case 32767:
                v25 := []int{1, 42}
                _ = v25
            default:
                var v26 int8 = -128
                _ = v26
            }
        }
        v3 -= v3
        if ((false || (v3 != v3)) && false) {
            v3 -= 1
            var v27 string = ("a" + (("\n\t" + "foo bar") + (("foo bar" + ("" + "a")) + ("\n\t" + "a"))))
            switch v27 {
            case "foo bar":
                var v28 uint8 = 2
                _ = v28
            default:
                fmt.Println(v1)
            }
            _ = v27
        } else {
            var v29 float32 = 0.0
            for fi30 := 0; fi30 < 13; fi30++ {
                v2 = v2
                _ = fi30
            }
            v3 = (^(v3))
            _ = v29
        }
    }
    switch 343 {
    case 42:
        switch 128 {
        case -158:
            if false {
                v31 := []int{9223372036854775807, 0, 0}
                _ = v31
            }
        }
    }
    var v32 int = 4294967295
    if ((((1e10 + (1e10 * 1.5)) / ((3.14159 + (1e10 + 0.0)) - 0.0)) + 0.0) == 3.14159) {
        var v33 float32 = 1.5
        v34 := []int{v32}
        if (0 < 2) {
            for fi35 := 0; fi35 < 5; fi35++ {
                v32 = v32
                _ = fi35
            }
            if (1e10 > 1.5) {
                var v36 int = -902
                _ = v36
            }
        } else {
            v3 -= 100
        }
        _ = v33
        _ = v34
    }
    v1 += v1
    var v37 int32 = 855
    switch 207 {
    case 9223372036854775807:
        fmt.Println(v3)
    default:
        v2 += v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v5
}
