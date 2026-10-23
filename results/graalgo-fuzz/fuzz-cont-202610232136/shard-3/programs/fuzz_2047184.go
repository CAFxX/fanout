package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
func genFn0(p0 string, p1 uint, p2 int16) uint64 {
    p0 = p0
    v7 := []int{-9223372036854775808, 522, 1}
    var v8 int64 = -494
    var v9 float64 = (3.14159 * (1e10 * (3.14159 + float64((float32(0.1) - 0.0)))))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 409
}

func main() {
    v1 := []uint32{0, 4294967295}
    var v2 string = "hello"
    var v3 int8 = -127
    var v4 int16 = 32767
    v5 := make(map[int]bool)
    var v6 float64 = (0.1 * ((0.0 * -2.25) / 1.0))
    var v10 uint = 0
    var v11 int = 127
    if ((v11 - (v11 << (((v10 | (v10 % ((v10) | 1)))) & 63))) <= v11) {
        if ((-(v4)) != -32768) {
            if false {
                fmt.Println(v11, v10)
            } else {
                v10 += v10
                var v12 int = 9223372036854775807
                _ = v12
            }
            if true {
                for fi13 := 0; fi13 < 7; fi13++ {
                    fmt.Println(fi13, v4, v11)
                    switch v11 {
                    case -730:
                        switch v2 {
                        case "foo bar":
                            v5[0] = true
                        case "":
                            v3 -= -128
                        case "a":
                            v10 += 42
                        default:
                            var v14 int64 = 9223372036854775807
                            _ = v14
                        }
                    case 2147483647:
                        fmt.Println(v2, v11, v10)
                    case -1:
                        fmt.Println(v4)
                    default:
                        var v15 int32 = 670
                        _ = v15
                    }
                    v1 = append(v1, 2147483647)
                    _ = fi13
                }
                if (32767 != -2147483648) {
                    v16 := GenS1{F0: (v11 & (v11 + (v11 ^ (v11 ^ (^(v11)))))), F1: 42, F2: 127}
                    if len(v1) > 0 {
                        v1[0] = 2147483647
                    }
                    _ = v16
                }
            }
            var v17 float64 = float64(0.1)
            _ = v17
        } else {
            v4 -= (v4 / (((v4 ^ (v4 ^ v4))) | 1))
            var v18 int32 = 2147483647
            v2 = v2
            _ = v18
        }
        switch v2 {
        case "":
            var v19 float64 = v6
            _ = v19
        case "foo bar":
            var v20 uint64 = 42
            _ = v20
        default:
            for fi21 := 0; fi21 < 8; fi21++ {
                if false {
                    v4 -= v4
                }
                _ = fi21
            }
        }
    } else {
        for fi22 := 0; fi22 < 13; fi22++ {
            v5[v11] = (true && (true || (2147483648 == 65535)))
            var v23 uint8 = 63
            var v24 uint16 = 32767
            _ = fi22
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 16; fi25++ {
            v3 -= v3
            _ = fi25
        }
        v26 := GenS1{F0: 128, F1: 770, F2: 2}
        _ = v26
    }
    if (!true) {
        if ((!true) || false) {
            v3 = v3
        }
        if true {
            v27 := GenS1{F0: v11, F1: 756, F2: 198}
            _ = v27
        }
        v2 = (v2 + v2)
    }
    var v28 float32 = 3.14159
    var v29 uint16 = 255
    if len(v1) > 0 {
        v1[0] = 832
    }
    var v30 float64 = 3.14159
    v30 -= v6
    v11 = (v11 + v11)
    if len(v1) > 0 {
        v1[0] = 2
    }
    var v31 uint = (v10 * (v10 >> ((4294967295) & 63)))
    switch v11 {
    case 0:
        var v32 string = ("hello" + (v2 + (v2 + v2)))
        _ = v32
    case 463:
        _ = len(v5)
    default:
        if (9223372036854775807 > -9223372036854775808) {
            fmt.Println(v4, v4, v6)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
