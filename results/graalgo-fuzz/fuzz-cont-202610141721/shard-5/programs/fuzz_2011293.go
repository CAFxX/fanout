package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint }
func genFn0(p0 uint, p1 string) uint {
    if (float64(0.0) >= 1e10) {
        v7 := []int{}
        switch ((((("" + "foo bar") + (("" + "hello") + "a")) + p1) + p1) + (p1 + "foo bar")) {
        case "":
            var v8 float32 = 0.1
            _ = v8
        default:
            var v9 int8 = 1
            _ = v9
        }
        _ = v7
    } else {
        var v10 string = "hello"
        v11 := []int{32767, 494}
        p0 += (^(p0))
        _ = v10
        _ = v11
    }
    var v12 uint16 = uint16(p0)
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := []int{-9223372036854775808}
        if false {
            p1 = "hello"
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            v15 := GenS1{F0: -310, F1: 1}
            _ = v15
        } else {
            var v16 int64 = -9223372036854775808
            _ = v16
        }
        var v17 uint16 = v12
        _ = fi13
        _ = v14
        _ = v17
    }
    fmt.Println(v12)
    fmt.Println(v12, p1, p1)
    _ = p0
    _ = p1
    _ = v12
    return p0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    v5 := []uint32{1, 91, 41, 2147483647}
    var v6 float64 = float64(0.1)
    v6 = 3.14159
    switch 42 {
    case 869:
        v6 -= float64((1.5 / 3.14159))
    case -9223372036854775808:
        var v18 string = ("a" + "foo bar")
        _ = v18
    default:
        v6 += 3.14159
    }
    v6 = ((1.5 + v6) * v6)
    v6 = v6
    v6 += v6
    var v19 int32 = -2147483648
    if false {
        for fi20 := 0; fi20 < 18; fi20++ {
            for fi21 := 0; fi21 < 9; fi21++ {
                if false {
                    var v22 uint8 = 128
                    if true {
                        v23 := GenS1{F0: 876, F1: 1}
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        fmt.Println(v6)
                        fmt.Println(fi20, v6)
                        _ = len(v5)
                        var v24 float32 = 1.5
                        _ = v23
                        _ = v24
                    } else {
                        v22 -= v22
                        var v25 float32 = 3.14159
                        _ = v25
                    }
                    _ = v22
                }
                var v26 int8 = 127
                _ = fi21
                _ = v26
            }
            v27 := GenS2{F0: 112, F1: 2, F2: 18446744073709551615}
            v28 := GenS2{F0: 703, F1: 255, F2: 42}
            _ = fi20
            _ = v27
            _ = v28
        }
    }
    delete(v2, ((("foo bar" + "foo bar") + "\n\t") + "hello"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
