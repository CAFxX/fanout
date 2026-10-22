package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS3 struct { F0 int32; F1 string }
func genFn0() int8 {
    var v6 uint8 = 42
    if false {
        var v7 uint = 4294967295
        _ = v7
    }
    fmt.Println(v6)
    v8 := []int{}
    if false {
        fmt.Println(v6)
    }
    _ = v6
    _ = v8
    return -82
}

func main() {
    var v1 float64 = (0.1 - 1e10)
    var v2 int32 = 0
    var v3 int32 = v2
    var v4 bool = (!(v3 == 0))
    v5 := []int{128, -822, 0}
    var v9 float64 = v1
    if ((("hello" + "a") + (("a" + (("a" + "foo bar") + ("a" + ("" + "")))) + ("foo bar" + (("\n\t" + "\n\t") + "foo bar")))) <= "") {
        _ = len(v5)
    } else {
        var v10 uint64 = 936
        v2 = (v2 + -2147483648)
        _ = v10
    }
    var v11 uint64 = 0
    var v12 uint64 = 412
    if len(v5) > 0 {
        v5[0] = 32768
    }
    v13 := GenS1{F0: 342}
    _ = len(v5)
    v3 -= (^(v2))
    v12 = (v11 + v11)
    var v14 int = 279
    v12 = (v11 - 18446744073709551615)
    switch "" {
    case "hello":
        if ((float32((0.0 * v9)) * (float32((v1 - v1)) * (1.5 + 0.0))) != -2.25) {
            v2 = v3
            v5 = append(v5, -1)
        } else {
            v3 -= v2
            if (!(42 >= genFn0())) {
                switch v14 {
                case 32767:
                    var v15 int16 = -663
                    _ = v15
                case 32768:
                    v5 = append(v5, (v14 << ((42) & 63)))
                }
            } else {
                var v16 uint = 1
                fmt.Println(v16)
                fmt.Println(v14)
                _ = v16
            }
        }
    case "\n\t":
        var v17 uint = 2
        _ = v17
    case "foo bar":
        if true {
            switch v14 {
            case -624:
                if (0.1 > 1e10) {
                    switch -2147483648 {
                    case 814:
                        for fi18 := 0; fi18 < 13; fi18++ {
                            v1 = float64(1e10)
                            var v19 int32 = v3
                            v5 = append(v5, (fi18 * 325))
                            _ = fi18
                            _ = v19
                        }
                    case 1:
                        v20 := GenS2{F0: 65535, F1: int16(v2), F2: 1}
                        _ = v20
                    case -592:
                        v14 = 32767
                    }
                    var v21 uint64 = v11
                    v2 = (v2 * (v2 / ((0) | 1)))
                    _ = v21
                }
            default:
                for fi22 := 0; fi22 < 14; fi22++ {
                    for fi23 := 0; fi23 < 3; fi23++ {
                        v24 := GenS2{F0: 65535, F1: 552, F2: 211}
                        v25 := GenS2{F0: 588, F1: 32767, F2: 423}
                        _ = len(v5)
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    if (127 < genFn0()) {
                        v2 = 118
                        var v26 uint16 = 0
                        _ = v26
                    }
                    v1 -= (0.0 / v1)
                    _ = fi22
                }
            }
        }
    }
    v3 = (^(v2))
    if v4 {
        if v4 {
            _ = len(v5)
            if v4 {
                if (2 < 118) {
                    _ = len(v5)
                    var v27 uint32 = 241
                    _ = v27
                }
            }
        } else {
            v5 = append(v5, 385)
        }
        var v28 uint8 = 128
        _ = v28
    }
    v14 += v14
    v4 = true
    switch v14 {
    case 478:
        v29 := GenS1{F0: 42}
        _ = v29
    case -859:
        switch 32768 {
        case 853:
            v30 := GenS3{F0: (v3 % (((v3 | (v3 & (v2 % ((v3) | 1))))) | 1)), F1: ""}
            _ = v30
        case -9223372036854775808:
            var v31 uint64 = (^(v11))
            _ = v31
        default:
            v32 := GenS3{F0: v2, F1: "hello"}
            _ = v32
        }
    default:
        var v33 int8 = 100
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
