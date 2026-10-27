package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 bool; F1 int }
func genFn0(p0 uint16) string {
    if (!(false && (42 >= 0))) {
        var v5 uint32 = 2147483647
        v6 := []int{}
        _ = len(v6)
        _ = v5
        _ = v6
    }
    var v7 uint32 = 2147483647
    v7 += (^(v7))
    _ = p0
    _ = v7
    return ""
}

func main() {
    v1 := []string{"a", "hello", "\n\t", "a"}
    v2 := make(map[uint32]int)
    var v3 int16 = 759
    v4 := []uint32{2147483647, 0, 2, 4294967295}
    fmt.Println(v3)
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: 2147483647, F1: 0}
        var v10 int16 = 32767
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: -9223372036854775808, F1: -1}
    v3 -= v3
    var v12 uint32 = 2
    v13 := v2[(v12 - (v12 << ((489) & 63)))]
    for fi14 := 0; fi14 < 12; fi14++ {
        fmt.Println(v3, fi14)
        v12 -= (v12 << ((4294967295) & 63))
        _ = fi14
    }
    v15 := GenS1{F0: (v13 * 32767), F1: 42}
    if false {
        v16 := GenS2{F0: 2147483647, F1: 4294967295}
        for fi17 := 0; fi17 < 11; fi17++ {
            v3 -= int16(fi17)
            _ = fi17
        }
        _ = v16
    } else {
        fmt.Println(v3, v13)
        if (44 < 45) {
            if (65535 < 2) {
                fmt.Println(v13)
                if true {
                    v18 := GenS2{F0: (v12 ^ (v12 ^ 435)), F1: 1}
                    switch "\n\t" {
                    case "":
                        v19 := GenS3{F0: (1e10 > ((0.0 / (3.14159 * 1.5)) - (((0.0 / (((float32(3.14159) - (-2.25 / 1.5)) + 1e10) + ((float32(1e10) - (-2.25 / 1.0)) + 0.0))) / (-2.25 - 3.14159)) * 3.14159))), F1: v13}
                        _ = v19
                    case "foo bar":
                        v20 := GenS2{F0: (v12 * (v12 / (((v12 << ((18446744073709551615) & 63))) | 1))), F1: 0}
                        _ = v20
                    case "a":
                        if (32767 >= 65535) {
                            var v21 uint = 65535
                            _ = v21
                        } else {
                            var v22 uint32 = v12
                            var v23 uint8 = 2
                            for fi24 := 0; fi24 < 17; fi24++ {
                                var v25 int8 = int8(v23)
                                if (false || false) {
                                    v23 += (v23 >> ((506) & 63))
                                    var v26 int64 = -9223372036854775808
                                    _ = v26
                                } else {
                                    v22 = 1
                                }
                                var v27 int8 = (v25 << ((1) & 63))
                                _ = fi24
                                _ = v25
                                _ = v27
                            }
                            _ = v22
                            _ = v23
                        }
                    default:
                        fmt.Println(v3, v13, v3)
                    }
                    fmt.Println(v12, v12, v12)
                    _ = v18
                }
                v12 -= v12
            } else {
                v13 -= v13
                v3 += v3
                for fi28 := 0; fi28 < 17; fi28++ {
                    v2[v12] = 32767
                    _ = fi28
                }
            }
        }
        var v29 int16 = v3
        _ = v29
    }
    var v30 int = 128
    var v31 uint32 = v12
    switch (("\n\t" + (("hello" + (("" + "a") + "hello")) + "foo bar")) + ("hello" + ("\n\t" + "\n\t"))) {
    case "foo bar":
        if true {
            var v32 float32 = ((float32(((-2.25 * 0.0) - (float64((-2.25 / float32((-2.25 - 0.0)))) * 1.5))) + (float32(3.14159) + -2.25)) / (1.5 * 1e10))
            fmt.Println(v3, v31, v3)
            _ = v32
        }
    case "":
        for fi33 := 0; fi33 < 13; fi33++ {
            v3 += v3
            _ = fi33
        }
    }
    fmt.Println(v30, v3)
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
