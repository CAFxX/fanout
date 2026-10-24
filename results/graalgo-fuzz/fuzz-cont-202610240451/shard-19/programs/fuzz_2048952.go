package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint; F1 int32; F2 int16 }
func genFn0(p0 uint64, p1 uint8, p2 bool) int {
    v7 := []int{32767}
    v8 := GenS1{F0: 65535, F1: p0, F2: 49}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1(p0 uint16) int64 {
    v9 := []int{-60, 9223372036854775807}
    var v10 uint64 = 18446744073709551615
    _ = p0
    _ = v10
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 string = "a"
    var v2 int = 4294967295
    v3 := make(map[string]bool)
    var v4 uint32 = 109
    var v5 uint32 = 492
    v6 := []int{(v2 | (-(v2))), int(v5), v2}
    v5 = v4
    _ = len(v6)
    var v11 uint8 = 255
    for fi12 := 0; fi12 < 12; fi12++ {
        if false {
            v1 = (("hello" + (v1 + (v1 + v1))) + "hello")
            for fi13 := 0; fi13 < 12; fi13++ {
                v1 = "foo bar"
                _ = fi13
            }
        }
        if len(v6) > 0 {
            v6[0] = v2
        }
        _ = fi12
    }
    v1 = "hello"
    v1 = ""
    if (true && ((v4 - (v4 - (v4 & uint32(v11)))) <= 0)) {
        if false {
            fmt.Println(v1, v1)
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            v11 -= (v11 * v11)
            _ = fi14
        }
        v4 = (^(v5))
    } else {
        v15 := GenS1{F0: 195, F1: 18446744073709551615, F2: 1}
        for fi16 := 0; fi16 < 20; fi16++ {
            fmt.Println(v1, v5, v2)
            switch (v2 << ((1) & 63)) {
            case 127:
                var v17 int = -870
                _ = v17
            case 0:
                var v18 uint32 = (v5 << ((140) & 63))
                _ = v18
            default:
                if (!false) {
                    v4 -= (v4 ^ uint32(v11))
                    if true {
                        for fi19 := 0; fi19 < 13; fi19++ {
                            if true {
                                v6 = append(v6, int(v11))
                                v5 = (v5 % (((^(v4))) | 1))
                            } else {
                                var v20 int16 = 0
                                var v21 float32 = 1e10
                                _ = v20
                                _ = v21
                            }
                            v22 := GenS2{F0: 2, F1: 1, F2: 0}
                            _ = fi19
                            _ = v22
                        }
                    }
                } else {
                    for fi23 := 0; fi23 < 1; fi23++ {
                        var v24 uint32 = v4
                        _ = fi23
                        _ = v24
                    }
                    v3[v1] = true
                    var v25 uint8 = (v11 & (v11 << ((170) & 63)))
                    _ = v25
                }
            }
            _ = fi16
        }
        _ = v15
    }
    var v26 string = v1
    if true {
        v1 = "a"
        for fi27 := 0; fi27 < 1; fi27++ {
            if (1e10 > (1e10 - ((float64((1e10 / 0.1)) + 0.1) * 0.0))) {
                if true {
                    v2 -= fi27
                } else {
                    v26 = ("" + "a")
                    v4 = (v4 & (v5 >> ((4294967295) & 63)))
                    fmt.Println(fi27)
                }
            } else {
                var v28 float64 = 0.0
                _ = v28
            }
            v3[""] = true
            _ = fi27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
