package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint8; F1 bool; F2 int64 }
func genFn0(p0 uint64) int32 {
    p0 = 1
    v7 := GenS1{F0: 508, F1: (^(p0))}
    var v8 uint16 = 1
    v9 := []int{}
    fmt.Println(v8)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = false
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    var v5 uint64 = 590
    v6 := make(map[int]int)
    if (-746 >= -498) {
        if v2 {
            var v10 uint32 = 4294967295
            v11 := []int{1}
            for fi12 := 0; fi12 < 16; fi12++ {
                v5 -= (^(v5))
                if (!false) {
                    var v13 int8 = 42
                    _ = v13
                } else {
                    switch "" {
                    case "\n\t":
                        var v14 float64 = 3.14159
                        _ = v14
                    case "hello":
                        var v15 uint32 = (v10 % ((v10) | 1))
                        _ = v15
                    }
                }
                v5 -= (^(v5))
                _ = fi12
            }
            _ = v10
            _ = v11
        } else {
            var v16 int64 = -1
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 uint32 = 0
            fmt.Println(fi17, v5, v18)
            if (false || v2) {
                v19 := GenS2{F0: 1, F1: false, F2: -331}
                v18 = 227
                _ = v19
            } else {
                if v2 {
                    var v20 string = (("hello" + ("foo bar" + "foo bar")) + ("hello" + "a"))
                    var v21 int32 = genFn0((v5 << ((730) & 63)))
                    _ = v20
                    _ = v21
                }
                v5 = v5
            }
            _ = fi17
            _ = v18
        }
        v22 := []int{32768, 619}
        _ = v22
    }
    var v23 int64 = 66
    v5 = 18446744073709551615
    v24 := v6[128]
    v25 := []int{-566, v24, (v24 + v24), (v24 * (v24 / (((^(v24))) | 1)))}
    v24 += v24
    var v26 uint = 0
    if (!(65535 < 4)) {
        var v27 int64 = 969
        var v28 float32 = 0.0
        _ = v27
        _ = v28
    }
    fmt.Println(v5, v23, v23)
    for fi29 := 0; fi29 < 17; fi29++ {
        fmt.Println(v23)
        var v30 int16 = 42
        if (0.0 >= 0.0) {
            switch (fi29 - v24) {
            case 687:
                _ = len(v25)
            }
        } else {
            switch -2147483648 {
            case 42:
                switch 128 {
                case 9223372036854775807:
                    var v31 float64 = 1.5
                    _ = v31
                }
            default:
                var v32 float32 = float32(1.5)
                _ = v32
            }
            for fi33 := 0; fi33 < 12; fi33++ {
                fmt.Println(v5, v26)
                _ = fi33
            }
            fmt.Println(v30, v24)
        }
        _ = fi29
        _ = v30
    }
    v24 -= 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
