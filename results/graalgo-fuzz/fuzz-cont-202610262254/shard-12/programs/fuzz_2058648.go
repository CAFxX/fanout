package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
func genFn0() uint16 {
    switch 259 {
    case 128:
        var v6 uint32 = 168
        _ = v6
    case -927:
        v7 := []int{-2147483648, 42}
        _ = v7
    default:
        if ((-725 <= 9223372036854775807) || (true && (-128 == 100))) {
            var v8 string = "\n\t"
            switch v8 {
            case "a":
                var v9 float64 = (-2.25 - 0.1)
                _ = v9
            default:
                var v10 string = (v8 + "hello")
                _ = v10
            }
            _ = v8
        } else {
            v11 := []int{2147483647, 4294967295, 395}
            _ = v11
        }
    }
    switch -1 {
    case 0:
        var v12 int8 = 1
        _ = v12
    case 32768:
        var v13 float64 = (0.1 - ((0.1 / (-2.25 + -2.25)) + 0.0))
        _ = v13
    }
    switch "foo bar" {
    case "":
        var v14 uint32 = 0
        _ = v14
    default:
        var v15 uint64 = 42
        _ = v15
    }
    var v16 int64 = 297
    v16 = v16
    _ = v16
    return 32767
}
func genFn1() uint {
    switch ("hello" + "\n\t") {
    case "hello":
        v17 := []int{9223372036854775807, 4294967295, 790}
        _ = v17
    default:
        if true {
            if true {
                var v18 string = (("\n\t" + ("" + "\n\t")) + (("foo bar" + "foo bar") + ("foo bar" + "a")))
                _ = v18
            }
            var v19 uint16 = 640
            for fi20 := 0; fi20 < 18; fi20++ {
                if (!false) {
                    v21 := GenS1{F0: 2, F1: true}
                    _ = v21
                } else {
                    if true {
                        if true {
                            var v22 float64 = 0.1
                            _ = v22
                        } else {
                            v19 = v19
                            v23 := GenS1{F0: 0, F1: (!true)}
                            _ = v23
                        }
                    } else {
                        v19 = v19
                        fmt.Println(fi20, fi20)
                    }
                    switch fi20 {
                    case -9223372036854775808:
                        v19 = v19
                    case 127:
                        v19 = (v19 * genFn0())
                    case 952:
                        v24 := GenS1{F0: 562, F1: (100 > 100)}
                        _ = v24
                    }
                }
                var v25 int64 = -9223372036854775808
                v26 := GenS1{F0: 65535, F1: (v19 != v19)}
                _ = fi20
                _ = v25
                _ = v26
            }
            _ = v19
        }
    }
    var v27 uint32 = 939
    _ = v27
    return 545
}

func main() {
    var v1 uint32 = 383
    v2 := []int{2147483647, 2147483647, 0}
    v3 := []int{}
    v4 := make(map[int]int)
    var v5 uint8 = uint8(v1)
    fmt.Println(v1, v5)
    v1 += (v1 % (((v1 << ((1) & 63))) | 1))
    var v28 float32 = float32((float64(3.14159) - 3.14159))
    if len(v3) > 0 {
        v3[0] = 128
    }
    v5 = (v5 % ((v5) | 1))
    fmt.Println(v28, v5, v5)
    v1 -= v1
    var v29 int8 = 0
    var v30 int = -341
    var v31 int16 = int16(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
