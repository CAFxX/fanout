package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int64 }
func genFn0() uint {
    var v4 uint8 = 127
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 int32 = -166
        fmt.Println(fi5, v6)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 0, F1: -35, F2: 0}
    _ = v4
    _ = v7
    return 168
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = -2.25
    v3 := []string{("" + "a"), "foo bar"}
    var v8 uint32 = 4294967295
    for fi9 := 0; fi9 < 16; fi9++ {
        _ = len(v1)
        _ = fi9
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v8 -= 4294967295
    var v10 uint64 = 42
    var v11 float32 = -2.25
    v10 = v10
    if false {
        var v12 uint64 = v10
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 float64 = (1.5 / 1e10)
        var v15 int16 = 1000
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 0, F1: 0, F2: 0}
    for fi17 := 0; fi17 < 1; fi17++ {
        v2 = ((((v2 / (-2.25 * v2)) / -2.25) + v2) * (float32(1.5) / ((v2 / (v2 - v2)) + v2)))
        var v18 uint16 = 224
        v19 := GenS1{F0: v18, F1: -9223372036854775808, F2: -308}
        _ = fi17
        _ = v18
        _ = v19
    }
    v8 = v8
    switch (("" + "foo bar") + ("" + ("a" + "foo bar"))) {
    case "hello":
        if false {
            var v20 uint16 = 298
            if (false || (false || (!true))) {
                v2 = v2
            } else {
                if false {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        v22 := GenS1{F0: 1, F1: 896, F2: -1}
                        var v23 uint16 = (v20 | (v20 + (^(v20))))
                        if (-1 >= 1) {
                            v20 -= v20
                        }
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                } else {
                    v24 := GenS1{F0: uint16(v8), F1: 42, F2: 0}
                    _ = v24
                }
                var v25 uint64 = 42
                v26 := GenS1{F0: (v20 >> ((411) & 63)), F1: 42, F2: 42}
                _ = v25
                _ = v26
            }
            v11 -= 1e10
            _ = v20
        } else {
            v2 -= 0.0
            var v27 uint = 4294967295
            _ = v27
        }
    case "\n\t":
        fmt.Println(v11, v10, v10)
    default:
        var v28 int = -9223372036854775808
        _ = v28
    }
    v10 = v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
