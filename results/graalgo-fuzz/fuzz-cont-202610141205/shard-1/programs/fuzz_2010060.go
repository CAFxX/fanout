package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0(p0 bool) int8 {
    var v4 float32 = 1.5
    v5 := []int{}
    _ = p0
    _ = v4
    _ = v5
    return 72
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint16 = 2
    var v3 float64 = 0.1
    v2 = v2
    v3 = 1.5
    v6 := []int{32768, 128, 127, 42}
    var v7 uint = 0
    switch "a" {
    case "":
        fmt.Println(v2)
    case "\n\t":
        v7 = 42
    case "foo bar":
        switch "" {
        case "hello":
            var v8 int = 128
            _ = v8
        case "\n\t":
            for fi9 := 0; fi9 < 9; fi9++ {
                if (343 == 123) {
                    if len(v6) > 0 {
                        v6[0] = fi9
                    }
                    if true {
                        v7 = (^(v7))
                        switch 1 {
                        case 128:
                            var v10 int32 = 697
                            _ = v10
                        case 1:
                            v2 -= 32767
                        case -338:
                            v7 -= v7
                        }
                    }
                } else {
                    if false {
                        var v11 uint = 4294967295
                        if ((((191 > 32767) && true) || ((18446744073709551615 == 0) && (true || (((-1 == int16(v11)) || true) || true)))) && true) {
                            var v12 float32 = float32((v3 + 3.14159))
                            v2 += v2
                            if (18446744073709551615 < 2) {
                                for fi13 := 0; fi13 < 5; fi13++ {
                                    var v14 uint32 = 492
                                    _ = fi13
                                    _ = v14
                                }
                                v15 := GenS1{F0: 58, F1: 101, F2: 255}
                                _ = v15
                            } else {
                                if len(v6) > 0 {
                                    v6[0] = (fi9 + 2147483647)
                                }
                                v12 = v12
                                v16 := GenS1{F0: 128, F1: (v7 * v11), F2: 42}
                                _ = v16
                            }
                            _ = v12
                        }
                        _ = v11
                    }
                    v7 = (v7 >> ((255) & 63))
                    v17 := GenS1{F0: 127, F1: (v7 & 2), F2: 0}
                    _ = v17
                }
                if len(v6) > 0 {
                    v6[0] = (fi9 - 101)
                }
                _ = fi9
            }
        default:
            fmt.Println(v1, v2, v1)
        }
    }
    var v18 uint16 = v2
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS1{F0: 1, F1: 65535, F2: 163}
        _ = fi19
        _ = v20
    }
    var v21 uint64 = 0
    v22 := GenS2{F0: v1}
    var v23 bool = (v21 < v21)
    var v24 int8 = -65
    var v25 uint16 = v18
    var v26 uint8 = 2
    if (true || (!(v3 == (v3 + v3)))) {
        _ = len(v6)
        var v27 uint = v7
        _ = v27
    }
    fmt.Println(v21, v26, v23)
    v28 := GenS3{F0: (("hello" + "") + "foo bar"), F1: 911}
    v18 = 633
    if v23 {
        v2 += uint16(v21)
        var v29 uint16 = (v25 & (v25 ^ uint16(v7)))
        v30 := GenS3{F0: "", F1: 4294967295}
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v6
    _ = v7
}
