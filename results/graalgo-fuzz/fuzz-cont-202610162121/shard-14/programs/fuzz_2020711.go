package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
func genFn0() uint64 {
    v7 := []int{128, 4294967295, -2147483648, 904}
    var v8 uint8 = 42
    v8 += (v8 + v8)
    if ((v8 ^ v8) < (^(v8))) {
        var v9 string = ""
        for fi10 := 0; fi10 < 10; fi10++ {
            if len(v7) > 0 {
                v7[0] = fi10
            }
            for fi11 := 0; fi11 < 9; fi11++ {
                v8 = (v8 >> ((0) & 63))
                _ = fi11
            }
            _ = fi10
        }
        _ = v9
    } else {
        v8 += 128
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 bool = (14 <= -34)
            if len(v7) > 0 {
                v7[0] = (fi12 + (fi12 + (fi12 / ((-2147483648) | 1))))
            }
            _ = fi12
            _ = v13
        }
    }
    var v14 uint = 255
    _ = v14
    _ = v7
    _ = v8
    return 664
}

func main() {
    var v1 string = "\n\t"
    var v2 uint8 = 134
    var v3 int64 = 9223372036854775807
    v4 := []string{}
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 770
    var v15 uint16 = 2
    v16 := GenS1{F0: (v15 | (^(v15))), F1: uint8(v15)}
    fmt.Println(v1, v6, v6)
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 float64 = ((0.0 / 3.14159) + (-2.25 + -2.25))
            v15 = uint16(v6)
            if (false || (v19 == v19)) {
                var v20 uint16 = 0
                var v21 int16 = 1000
                if false {
                    var v22 uint64 = 0
                    v20 = v20
                    _ = v22
                }
                _ = v20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 uint64 = genFn0()
            v25 := GenS1{F0: v15, F1: 1}
            v26 := GenS1{F0: v15, F1: 182}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi17
    }
    v27 := GenS1{F0: v15, F1: 2}
    var v28 uint64 = 2
    fmt.Println(v28, v1)
    v29 := GenS1{F0: (v15 * v15), F1: 2}
    var v30 int64 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
