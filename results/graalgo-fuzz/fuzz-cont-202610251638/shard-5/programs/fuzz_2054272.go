package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint64, p1 bool) int32 {
    v6 := []int{}
    if false {
        var v7 int = 127
        _ = v7
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        p1 = (127 <= 42)
    }
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int8 = 1
    var v3 int32 = 740
    var v4 string = "foo bar"
    var v5 uint32 = 2147483648
    v3 = (^(v3))
    v1 -= v1
    var v8 uint8 = 42
    v3 = v3
    fmt.Println(v3, v8)
    var v9 int8 = v2
    fmt.Println(v1, v4, v4)
    var v10 uint = 4294967295
    v11 := []int{}
    for fi12 := 0; fi12 < 7; fi12++ {
        _ = len(v11)
        v1 = v1
        if (uint64(v5) > 504) {
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 int64 = (v1 * (-(v1)))
                var v15 uint = (v10 * (v10 / ((18446744073709551615) | 1)))
                var v16 uint16 = 65535
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            if false {
                fmt.Println(v5, v9)
                if false {
                    switch v4 {
                    case "foo bar":
                        fmt.Println(v8)
                    default:
                        v4 = ((((((v4 + "") + v4) + (v4 + ("" + v4))) + ((((v4 + v4) + "foo bar") + (((v4 + v4) + (v4 + v4)) + v4)) + v4)) + v4) + v4)
                    }
                    fmt.Println(v1)
                    var v17 uint8 = v8
                    _ = v17
                } else {
                    v2 -= v2
                    v8 += v8
                    for fi18 := 0; fi18 < 10; fi18++ {
                        var v19 int32 = (-(v3))
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        var v20 uint32 = 2147483648
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                }
                v21 := GenS1{F0: 42, F1: v5}
                _ = v21
            } else {
                if false {
                    for fi22 := 0; fi22 < 15; fi22++ {
                        var v23 bool = (1 != -995)
                        v24 := GenS2{F0: uint8(v10)}
                        var v25 float32 = ((0.1 / 3.14159) * (0.1 - 0.0))
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                }
                if (1 >= 18446744073709551615) {
                    var v26 uint16 = 0
                    v5 = 63
                    _ = v26
                } else {
                    v27 := GenS2{F0: (^(v8))}
                    if len(v11) > 0 {
                        v11[0] = 2147483647
                    }
                    _ = v27
                }
                fmt.Println(v10, v5)
            }
            fmt.Println(v5)
        }
        _ = fi12
    }
    if len(v11) > 0 {
        v11[0] = 32768
    }
    var v28 uint32 = v5
    _ = len(v11)
    var v29 int16 = -23
    if false {
        v9 -= v2
        if len(v11) > 0 {
            v11[0] = 1
        }
    } else {
        fmt.Println(v28, v29, v9)
        for fi30 := 0; fi30 < 17; fi30++ {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi30
        }
        if (2 != v10) {
            v10 -= v10
            for fi31 := 0; fi31 < 5; fi31++ {
                for fi32 := 0; fi32 < 7; fi32++ {
                    fmt.Println(v28, v2)
                    v9 -= v9
                    var v33 uint64 = 18446744073709551615
                    _ = fi32
                    _ = v33
                }
                fmt.Println(v10, v3)
                var v34 uint8 = (^(v8))
                _ = fi31
                _ = v34
            }
            v11 = append(v11, -633)
        }
    }
    v9 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
