package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0(p0 uint64, p1 uint) uint32 {
    if true {
        var v5 uint32 = 2
        _ = v5
    }
    v6 := []int{-152, -2147483648, 4294967295}
    if len(v6) > 0 {
        v6[0] = 42
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v6
    return 959
}
func genFn1(p0 string) int64 {
    var v7 int32 = 2147483647
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]string)
    var v3 bool = false
    var v4 float32 = float32(-2.25)
    v8 := []int{9223372036854775807, 718, 128, 42}
    if false {
        switch "foo bar" {
        case "a":
            fmt.Println(v3)
        default:
            var v9 int = 128
            _ = v9
        }
        v10 := GenS1{F0: 675}
        _ = v10
    } else {
        var v11 int = -9223372036854775808
        if len(v8) > 0 {
            v8[0] = v11
        }
        var v12 float64 = -2.25
        _ = v11
        _ = v12
    }
    fmt.Println(v3, v3)
    for fi13 := 0; fi13 < 5; fi13++ {
        v14 := GenS1{F0: 4294967295}
        v8 = append(v8, fi13)
        _ = fi13
        _ = v14
    }
    var v15 bool = (18446744073709551615 > 18446744073709551615)
    var v16 int64 = -9223372036854775808
    var v17 uint16 = 1
    if (0 <= 0) {
        var v18 int32 = 2147483647
        _ = v18
    } else {
        if v3 {
            switch 4294967295 {
            case 690:
                for fi19 := 0; fi19 < 18; fi19++ {
                    var v20 int16 = 42
                    fmt.Println(v17)
                    switch int(v20) {
                    case 128:
                        var v21 uint8 = 0
                        _ = v21
                    case 127:
                        switch "" {
                        case "foo bar":
                            v20 = int16(v17)
                        default:
                            v22 := GenS2{F0: "\n\t", F1: 42}
                            _ = v22
                        }
                    case 766:
                        var v23 uint8 = 2
                        _ = v23
                    default:
                        v15 = v15
                    }
                    _ = fi19
                    _ = v20
                }
            case 128:
                v24 := GenS2{F0: "\n\t", F1: 1}
                _ = v24
            }
            var v25 bool = (255 <= 0)
            for fi26 := 0; fi26 < 18; fi26++ {
                v17 += (v17 + (^(v17)))
                var v27 uint = 200
                _ = fi26
                _ = v27
            }
            _ = v25
        } else {
            var v28 int16 = -1
            _ = v28
        }
        for fi29 := 0; fi29 < 3; fi29++ {
            v15 = (false || (403 != 4294967295))
            fmt.Println(v15, v16)
            _ = fi29
        }
        if false {
            _ = len(v8)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
