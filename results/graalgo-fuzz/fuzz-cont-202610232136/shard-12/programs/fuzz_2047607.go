package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int8; F2 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 int16; F1 string; F2 uint }
func genFn0(p0 bool, p1 int16) uint8 {
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 uint8 = 14
        var v8 uint8 = v7
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{-9223372036854775808, 2147483647, 2147483647}
    switch 885 {
    case -842:
        var v10 float64 = (1e10 * 0.0)
        _ = v10
    case -914:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    default:
        p1 = (-(p1))
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(fi11)
        p0 = p0
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v9
    return 142
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    var v3 int32 = 673
    v4 := make(map[uint32]bool)
    var v5 uint = 476
    v12 := GenS1{F0: 0, F1: 1}
    v13 := []int{2147483647, 42, 32767, 127}
    v3 = v3
    v3 = v3
    v14 := GenS2{F0: 479, F1: 2, F2: 127}
    var v15 uint8 = genFn0((819 != 4294967295), 42)
    v16 := v1[(("foo bar" + "hello") + "a")]
    v17 := GenS1{F0: 2147483648, F1: 2147483647}
    v18 := GenS2{F0: 139, F1: 0, F2: v15}
    switch (v16 + ((((v16 + "hello") + v16) + ("hello" + "foo bar")) + v16)) {
    case "hello":
        v13 = append(v13, 32768)
    case "foo bar":
        if (true && (-9223372036854775808 >= 42)) {
            var v19 int32 = (v3 / ((v3) | 1))
            switch v16 {
            case "foo bar":
                v16 = "\n\t"
            }
            _ = len(v13)
            _ = v19
        } else {
            for fi20 := 0; fi20 < 6; fi20++ {
                switch v16 {
                case "foo bar":
                    var v21 uint16 = 2
                    _ = v21
                }
                if false {
                    if len(v13) > 0 {
                        v13[0] = (fi20 * (fi20 + (fi20 | 9223372036854775807)))
                    }
                    switch 0 {
                    case 127:
                        var v22 string = ((("a" + "\n\t") + v16) + "a")
                        _ = v22
                    }
                    var v23 uint8 = 2
                    _ = v23
                }
                if (true && true) {
                    for fi24 := 0; fi24 < 12; fi24++ {
                        var v25 float32 = 0.0
                        _ = fi24
                        _ = v25
                    }
                    v5 += v5
                    v16 = v16
                } else {
                    v16 = "foo bar"
                    v5 = (v5 + (v5 % (((v5 << ((v5) & 63))) | 1)))
                }
                _ = fi20
            }
        }
    default:
        switch v16 {
        case "a":
            if false {
                v26 := GenS3{F0: v15, F1: -128, F2: -2147483648}
                _ = v26
            } else {
                v16 = v16
                v15 += (v15 - genFn0(false, -32768))
                switch 714 {
                case 128:
                    v27 := GenS4{F0: (v16 + ("hello" + v16))}
                    _ = v27
                case 841:
                    v15 = 0
                }
            }
        case "":
            v28 := GenS5{F0: -1, F1: int8(v5)}
            _ = v28
        case "hello":
            v29 := GenS6{F0: 108, F1: "foo bar", F2: 905}
            _ = v29
        default:
            var v30 int = 745
            _ = v30
        }
    }
    var v31 float32 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
