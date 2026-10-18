package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int16 }
func genFn0(p0 bool, p1 int64, p2 string) int {
    v5 := []int{4294967295, 28}
    fmt.Println(p2, p1)
    v6 := GenS1{F0: 1, F1: 785, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 894
}
func genFn1() uint16 {
    var v7 int = -9223372036854775808
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := []int{(v7 / (((v7 / ((v7) | 1))) | 1))}
        switch (fi8 >> ((540) & 63)) {
        case 42:
            var v10 uint8 = 42
            _ = v10
        default:
            fmt.Println(fi8, v7)
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint64 = 42
            _ = len(v9)
            var v13 int32 = 0
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    switch (^(v7)) {
    case 2147483647:
        v7 += v7
    case -9223372036854775808:
        fmt.Println(v7)
    }
    _ = v7
    return 294
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]uint32)
    var v3 uint = 4294967295
    v4 := []string{v1}
    fmt.Println(v1, v3)
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int = (fi14 << ((v3) & 63))
        v16 := GenS1{F0: 32767, F1: 42, F2: 913}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int16 = -360
    v2[250] = 4294967295
    switch v1 {
    case "\n\t":
        var v18 int64 = 0
        _ = v18
    default:
        var v19 string = v1
        _ = v19
    }
    var v20 uint32 = 889
    if (v3 > (v3 / ((v3) | 1))) {
        if false {
            for fi21 := 0; fi21 < 14; fi21++ {
                if true {
                    v22 := GenS1{F0: 65535, F1: 499, F2: (v17 % ((v17) | 1))}
                    v1 = ((((v1 + ("foo bar" + ("\n\t" + v1))) + (("" + v1) + "hello")) + v1) + "hello")
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        var v24 int8 = -128
                        if len(v4) > 0 {
                            v4[0] = v1
                        }
                        v25 := GenS1{F0: 1, F1: genFn1(), F2: v17}
                        fmt.Println(v24, v1, v20)
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
                _ = fi21
            }
        } else {
            if false {
                v17 += (v17 >> ((v3) & 63))
                if len(v4) > 0 {
                    v4[0] = v1
                }
            }
        }
        v1 = "foo bar"
    }
    var v26 float64 = float64(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
}
