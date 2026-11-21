package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint64 }
func genFn0() int8 {
    var v6 int64 = -1
    var v7 uint8 = 2
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int16 = 1000
        var v10 uint16 = 255
        var v11 int16 = v9
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := []int{}
        v6 = (-(v6))
        if true {
            if true {
                fmt.Println(v7)
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 bool = (!((1e10 <= 0.0) && (299 >= 18446744073709551615)))
                    if v15 {
                        switch (fi14 + (fi14 ^ (fi14 * (fi14 - (^(fi14)))))) {
                        case -1:
                            v15 = v15
                        }
                    } else {
                        switch fi12 {
                        case 127:
                            v6 = (^(v6))
                        case 32768:
                            fmt.Println(fi12, fi12)
                        case 665:
                            v16 := GenS1{F0: 0}
                            _ = v16
                        }
                        var v17 int32 = 364
                        _ = v17
                    }
                    var v18 int32 = 893
                    _ = fi14
                    _ = v15
                    _ = v18
                }
                v7 = 127
            } else {
                for fi19 := 0; fi19 < 8; fi19++ {
                    switch "\n\t" {
                    case "":
                        var v20 float64 = 1e10
                        _ = len(v13)
                        switch (fi19 * (fi12 >> ((179) & 63))) {
                        case -2147483648:
                            var v21 string = "a"
                            _ = v21
                        case 42:
                            var v22 int32 = -198
                            _ = v22
                        case 2147483647:
                            if len(v13) > 0 {
                                v13[0] = 0
                            }
                        }
                        _ = v20
                    case "foo bar":
                        var v23 int8 = 0
                        _ = v23
                    default:
                        var v24 float32 = -2.25
                        _ = v24
                    }
                    _ = fi19
                }
                v6 += v6
            }
        }
        _ = fi12
        _ = v13
    }
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 bool = (!(727 <= 255))
    v2 := []string{"a"}
    var v3 uint32 = 2147483648
    var v4 float32 = (1e10 * 1e10)
    var v5 int32 = 1
    fmt.Println(v1, v5, v5)
    v25 := GenS2{F0: 393}
    for fi26 := 0; fi26 < 4; fi26++ {
        v3 = (v3 << ((597) & 63))
        var v27 float32 = float32((1.5 * 0.1))
        v28 := GenS1{F0: -32768}
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 float64 = 1.5
    v30 := GenS2{F0: 42}
    v31 := GenS3{F0: v1, F1: 1}
    for fi32 := 0; fi32 < 13; fi32++ {
        v3 += (v3 % ((v3) | 1))
        switch fi32 {
        case 0:
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
        case 667:
            v5 -= (v5 & (v5 + -202))
        case -2147483648:
            var v33 uint8 = 127
            _ = v33
        }
        _ = fi32
    }
    v2 = append(v2, (("hello" + "") + "a"))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
