package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int16 }
func genFn0(p0 uint8, p1 int32) int {
    p1 -= p1
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: (-(fi4)), F1: 1, F2: -32768}
        _ = fi4
        _ = v5
    }
    var v6 int8 = -1
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint8 = p0
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return -9223372036854775808
}
func genFn1(p0 uint64) uint {
    p0 += (p0 - 0)
    p0 = (p0 % (((^(p0))) | 1))
    _ = p0
    return 0
}

func main() {
    var v1 int16 = 738
    var v2 uint32 = 65535
    v3 := make(map[int]bool)
    var v9 int = genFn0(119, -2147483648)
    if false {
        var v10 uint8 = 157
        _ = v10
    }
    v9 += v9
    if false {
        v1 += 182
    }
    var v11 int8 = -128
    v9 = v9
    v1 -= v1
    var v12 int32 = int32(v2)
    if true {
        if true {
            v13 := []int{0, 2147483647}
            var v14 uint16 = 960
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 4; fi15++ {
            v16 := v3[403]
            _ = fi15
            _ = v16
        }
    }
    if (true || true) {
        v3[0] = (v12 > (v12 * -1))
    }
    var v17 uint32 = (v2 & 2147483647)
    v9 = (v9 / (((^(v9))) | 1))
    switch "" {
    case "foo bar":
        if true {
            delete(v3, v9)
            if (v11 == v11) {
                fmt.Println(v2, v12)
            } else {
                v17 = (v2 << ((4294967295) & 63))
                switch ("a" + (("hello" + "") + ("a" + "a"))) {
                case "a":
                    for fi18 := 0; fi18 < 4; fi18++ {
                        fmt.Println(v12, v9)
                        if false {
                            fmt.Println(v9)
                            v19 := []int{}
                            _ = v19
                        }
                        _ = fi18
                    }
                case "foo bar":
                    var v20 uint8 = 152
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 12; fi21++ {
                        var v22 uint32 = 0
                        _ = fi21
                        _ = v22
                    }
                }
            }
            for fi23 := 0; fi23 < 13; fi23++ {
                var v24 int = (fi23 - (fi23 ^ 2147483647))
                _ = len(v3)
                _ = len(v3)
                _ = fi23
                _ = v24
            }
        } else {
            var v25 bool = false
            _ = v25
        }
    }
    v9 = (v9 ^ 672)
    v26 := []int{v9}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v9
}
