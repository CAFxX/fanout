package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
func genFn0() int8 {
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 float64 = 0.0
        if true {
            var v7 string = ("a" + "hello")
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        if false {
            var v9 int8 = 0
            v9 -= (v9 / (((^(v9))) | 1))
            var v10 uint64 = 0
            _ = v10
            _ = v9
        }
        switch 1 {
        case 733:
            var v11 int8 = 0
            _ = v11
        }
        fmt.Println(fi8)
        _ = fi8
    }
    var v12 uint32 = 42
    v12 -= (v12 % ((v12) | 1))
    _ = v12
    return -69
}
func genFn1() int16 {
    v13 := []int{}
    var v14 int16 = -627
    v14 = (v14 ^ (v14 ^ -453))
    _ = v13
    _ = v14
    return (v14 + (v14 - (v14 % (((v14 << ((18446744073709551615) & 63))) | 1))))
}

func main() {
    var v1 int = 32767
    v2 := make(map[string]int)
    var v3 int8 = -128
    v4 := []int{v1, (v1 << ((723) & 63)), -865}
    var v15 uint64 = 1
    v3 -= (v3 + v3)
    v3 = 42
    if (-109 == 2147483647) {
        v16 := GenS1{F0: 128, F1: 65535}
        v3 = 42
        _ = v16
    }
    if len(v4) > 0 {
        v4[0] = v1
    }
    v3 = (v3 + (v3 / ((v3) | 1)))
    v17 := v2["foo bar"]
    v3 = (-(v3))
    v17 = v1
    v18 := GenS1{F0: 2, F1: 1}
    v17 -= (v17 << ((2) & 63))
    if false {
        v17 = v17
    }
    v19 := GenS1{F0: 1, F1: 860}
    if len(v4) > 0 {
        v4[0] = (v1 * (v17 << ((0) & 63)))
    }
    if true {
        v3 = genFn0()
    } else {
        var v20 int = (v17 ^ 32768)
        for fi21 := 0; fi21 < 13; fi21++ {
            for fi22 := 0; fi22 < 17; fi22++ {
                if (false || false) {
                    var v23 int32 = 1
                    v23 = v23
                    switch (fi22 + (v1 - 139)) {
                    case -2147483648:
                        if (32767 <= 1000) {
                            var v24 int64 = 787
                            _ = v24
                        }
                    case 277:
                        v20 = v1
                    }
                    v25 := GenS1{F0: 2, F1: uint32(v15)}
                    _ = v23
                    _ = v25
                } else {
                    if len(v4) > 0 {
                        v4[0] = (fi22 >> ((18446744073709551615) & 63))
                    }
                }
                _ = fi22
            }
            if ((((-1 > 1) && true) || (false && false)) || false) {
                v4 = append(v4, v17)
            }
            _ = fi21
        }
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
