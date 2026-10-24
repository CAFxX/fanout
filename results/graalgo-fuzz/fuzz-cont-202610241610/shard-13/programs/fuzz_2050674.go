package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int32 }
func genFn0() int32 {
    var v4 string = "foo bar"
    v4 = (v4 + v4)
    var v5 uint64 = 42
    v5 += v5
    var v6 int = 4294967295
    _ = v4
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = (1.5 * 0.0)
    var v3 bool = true
    _ = len(v1)
    switch "foo bar" {
    case "":
        v3 = (-9223372036854775808 != 9223372036854775807)
    case "foo bar":
        var v7 uint8 = 1
        _ = v7
    case "hello":
        v3 = (0 != 0)
    default:
        var v8 int16 = 109
        _ = v8
    }
    if v3 {
        var v9 int16 = 234
        fmt.Println(v9)
        _ = v9
    } else {
        if (42 == 1) {
            v2 += -2.25
            if v3 {
                v2 = (v2 * float32(3.14159))
                fmt.Println(v3, v3)
                var v10 int64 = -1
                _ = v10
            }
            var v11 float64 = 1e10
            _ = v11
        }
        v2 += (v2 / v2)
    }
    v12 := GenS1{F0: v3, F1: 2}
    var v13 int32 = 117
    fmt.Println(v2, v13)
    v14 := GenS2{F0: 0}
    var v15 int16 = 1000
    fmt.Println(v2)
    fmt.Println(v3, v2)
    v16 := []int{32768, 0}
    fmt.Println(v13, v2, v15)
    v15 += (v15 % ((v15) | 1))
    v13 -= (v13 + v13)
    for fi17 := 0; fi17 < 13; fi17++ {
        if true {
            if v3 {
                for fi18 := 0; fi18 < 13; fi18++ {
                    var v19 uint32 = 2147483648
                    _ = fi18
                    _ = v19
                }
                v15 = (v15 << ((65535) & 63))
                v20 := GenS2{F0: genFn0()}
                _ = v20
            } else {
                var v21 uint32 = 2
                fmt.Println(v3, v2, v2)
                v22 := GenS1{F0: v3, F1: 2}
                _ = v21
                _ = v22
            }
        } else {
            v15 = (^(v15))
            v1[4294967295] = (-2.25 >= v2)
            for fi23 := 0; fi23 < 14; fi23++ {
                _ = len(v1)
                if v3 {
                    var v24 float32 = 0.1
                    switch "hello" {
                    case "foo bar":
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    case "hello":
                        if (42 >= -117) {
                            var v25 int = (^(fi17))
                            _ = v25
                        } else {
                            if len(v16) > 0 {
                                v16[0] = (-(fi23))
                            }
                        }
                    default:
                        var v26 float64 = 1.5
                        var v27 int8 = 100
                        _ = v26
                        _ = v27
                    }
                    _ = v24
                } else {
                    fmt.Println(v2, v15, v2)
                    v28 := GenS2{F0: v13}
                    var v29 int64 = 0
                    _ = v28
                    _ = v29
                }
                _ = fi23
            }
        }
        var v30 float32 = v2
        var v31 int8 = 30
        _ = fi17
        _ = v30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
