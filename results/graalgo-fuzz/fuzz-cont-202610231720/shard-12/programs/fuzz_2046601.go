package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint; F2 uint64 }
func genFn0() uint {
    var v6 uint8 = 1
    var v7 int32 = -1
    _ = v6
    _ = v7
    return 637
}
func genFn1(p0 int64) uint64 {
    fmt.Println(p0)
    v8 := []int{0, 127, 2147483647, 4294967295}
    _ = len(v8)
    p0 -= (p0 - (-(p0)))
    p0 += p0
    _ = p0
    _ = v8
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 1
    var v3 int8 = 0
    var v4 string = ("foo bar" + "hello")
    v5 := []string{("foo bar" + (((v4 + v4) + v4) + v4)), "", "\n\t", (v4 + "a")}
    v9 := v1[(((v4 + ("foo bar" + v4)) + v4) + v4)]
    v10 := GenS1{F0: -32768, F1: 36, F2: genFn1(-9223372036854775808)}
    v11 := GenS1{F0: 42, F1: 42, F2: 42}
    for fi12 := 0; fi12 < 12; fi12++ {
        if false {
            v2 += 112
            fmt.Println(v2, v2)
        } else {
            var v13 int16 = 0
            var v14 int8 = 21
            var v15 string = (("\n\t" + "hello") + v4)
            _ = v13
            _ = v14
            _ = v15
        }
        if (48 < -747) {
            if (-1 != 1) {
                var v16 uint16 = 1
                if true {
                    var v17 int8 = v3
                    for fi18 := 0; fi18 < 1; fi18++ {
                        var v19 uint64 = 867
                        _ = fi18
                        _ = v19
                    }
                    _ = v17
                }
                for fi20 := 0; fi20 < 11; fi20++ {
                    fmt.Println(v3, fi12)
                    _ = fi20
                }
                _ = v16
            }
        }
        var v21 float64 = 3.14159
        _ = fi12
        _ = v21
    }
    var v22 float32 = 3.14159
    v9 += v9
    for fi23 := 0; fi23 < 10; fi23++ {
        if false {
            v22 += (v22 + -2.25)
        }
        _ = fi23
    }
    if ((v9 ^ (v9 & (v9 - (v9 & (v9 + (v9 >> (((^(v2))) & 63))))))) < 2) {
        var v24 int64 = 1
        var v25 int64 = (^(v24))
        _ = v24
        _ = v25
    } else {
        if (!false) {
            if (42 >= int32(v2)) {
                switch -1 {
                case 0:
                    v3 = v3
                case 734:
                    var v26 uint32 = v9
                    _ = v26
                case 42:
                    v27 := GenS2{F0: (v3 + (v3 >> ((genFn0()) & 63))), F1: (^(v2)), F2: 644}
                    _ = v27
                default:
                    var v28 int = -102
                    _ = v28
                }
            } else {
                for fi29 := 0; fi29 < 14; fi29++ {
                    v4 = "a"
                    if (-323 != 1000) {
                        if len(v5) > 0 {
                            v5[0] = ((v4 + (v4 + v4)) + v4)
                        }
                    }
                    for fi30 := 0; fi30 < 17; fi30++ {
                        v9 += (v9 % ((v9) | 1))
                        v9 -= v9
                        _ = fi30
                    }
                    _ = fi29
                }
                v1[("foo bar" + (v4 + v4))] = 326
                var v31 int64 = 9223372036854775807
                _ = v31
            }
            if (true && (0 <= 32767)) {
                var v32 int = -724
                fmt.Println(v3)
                _ = v32
            }
        } else {
            v9 -= (v9 & 0)
        }
        var v33 int64 = 9223372036854775807
        _ = v33
    }
    v9 += (v9 >> (((v2 + genFn0())) & 63))
    v3 = v3
    v9 += (^(v9))
    var v34 int8 = 61
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v9
}
