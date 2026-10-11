package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
func genFn0() bool {
    var v7 int8 = 15
    var v8 float64 = 0.1
    v7 -= (-(v7))
    v9 := GenS1{F0: -702, F1: 9223372036854775807}
    _ = v7
    _ = v8
    _ = v9
    return (-32768 < 42)
}
func genFn1() int64 {
    v10 := GenS1{F0: 0, F1: 0}
    var v11 uint64 = 4294967295
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 int16 = -1
    var v2 string = ("\n\t" + (("" + "hello") + "\n\t"))
    var v3 bool = (1 != 4294967295)
    var v4 int16 = (v1 * (v1 >> ((2) & 63)))
    var v5 int32 = -908
    var v6 int64 = -9223372036854775808
    v12 := GenS1{F0: v5, F1: v6}
    v13 := GenS1{F0: v5, F1: (v6 & -9223372036854775808)}
    if (1 == 2147483648) {
        v14 := GenS2{F0: 447, F1: -1, F2: 127}
        v5 += -963
        var v15 int8 = 1
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 3; fi16++ {
            fmt.Println(fi16)
            _ = fi16
        }
        v4 += (v4 / ((v1) | 1))
    }
    v5 = (v5 | (^(v5)))
    v1 -= v4
    var v17 float32 = 0.1
    v6 += (v6 % (((v6 ^ -9223372036854775808)) | 1))
    var v18 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
