package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int8; F1 int; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }

func main() {
    v1 := []int{}
    var v2 uint32 = 2147483648
    var v3 int16 = -32768
    v4 := make(map[string]bool)
    v5 := make(map[string]bool)
    v2 -= (v2 + v2)
    _ = len(v1)
    fmt.Println(v2)
    v6 := GenS1{F0: 943, F1: (v3 & v3)}
    var v7 int = 9223372036854775807
    switch 0 {
    case 128:
        v8 := GenS1{F0: 18446744073709551615, F1: (-(v3))}
        _ = v8
    case -2147483648:
        fmt.Println(v3, v3)
    default:
        for fi9 := 0; fi9 < 15; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                if (int32(fi9) == -2147483648) {
                    fmt.Println(v7, v2, v7)
                }
                var v11 string = ""
                fmt.Println(v11, fi9)
                fmt.Println(v11)
                _ = fi10
                _ = v11
            }
            v3 -= (v3 >> ((1) & 63))
            v2 = (v2 >> ((42) & 63))
            _ = fi9
        }
    }
    var v12 float32 = -2.25
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint8 = 162
        for fi15 := 0; fi15 < 2; fi15++ {
            switch "foo bar" {
            case "foo bar":
                if (v14 != 160) {
                    v16 := GenS2{F0: 0, F1: v7, F2: v2}
                    _ = v16
                } else {
                    switch "\n\t" {
                    case "a":
                        if (2 > 96) {
                            var v17 int8 = -128
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if (!(v3 <= (v3 << ((906) & 63)))) {
                                v1 = append(v1, (v7 >> ((255) & 63)))
                            }
                            _ = v17
                        }
                    case "foo bar":
                        var v18 string = ("foo bar" + "foo bar")
                        _ = v18
                    case "":
                        fmt.Println(v12, v2)
                    default:
                        v19 := GenS2{F0: 76, F1: v7, F2: v2}
                        _ = v19
                    }
                    v12 += (v12 - -2.25)
                    v20 := GenS1{F0: 67, F1: v3}
                    if true {
                        v21 := GenS3{F0: 518}
                        if (!((v3 + (v3 + v3)) < v3)) {
                            var v22 float64 = ((0.1 / 0.1) * 0.1)
                            v7 = (-(fi15))
                            v23 := GenS4{F0: (v2 >> ((uint(v2)) & 63))}
                            fmt.Println(v22, v12, v22)
                            _ = v22
                            _ = v23
                        }
                        _ = v21
                    }
                    _ = v20
                }
            case "a":
                var v24 bool = (!(((((1.5 / -2.25) + ((1.5 * 0.1) * (0.1 / (-2.25 + (1.5 - 1e10))))) / (1e10 + 0.1)) / ((0.1 + -2.25) * 1e10)) >= 0.1))
                _ = v24
            }
            v1 = append(v1, fi13)
            v1 = append(v1, v7)
            _ = fi15
        }
        var v25 float64 = 1.5
        _ = fi13
        _ = v14
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 bool = ("foo bar" > "a")
            _ = fi27
            _ = v28
        }
        v3 -= -32768
        _ = fi26
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        var v30 uint16 = 42
        _ = fi29
        _ = v30
    }
    delete(v5, (("hello" + "foo bar") + "a"))
    var v31 bool = true
    var v32 string = "foo bar"
    v3 -= v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
