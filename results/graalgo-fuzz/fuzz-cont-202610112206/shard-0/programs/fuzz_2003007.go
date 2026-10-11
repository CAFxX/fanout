package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int16 }
func genFn0() uint8 {
    var v5 uint64 = 310
    switch 42 {
    case 127:
        if true {
            var v6 bool = (32 == 42)
            _ = v6
        } else {
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 int64 = 105
                fmt.Println(v5)
                _ = fi7
                _ = v8
            }
        }
    default:
        var v9 uint32 = 2147483647
        _ = v9
    }
    _ = v5
    return 88
}

func main() {
    var v1 uint8 = 127
    var v2 int8 = 0
    v3 := []int{4294967295, -9223372036854775808, 785, 1}
    var v4 uint32 = 9
    if (!((-1 > 0) && (313 <= 42))) {
        if ((v4 * 1) >= (v4 % ((v4) | 1))) {
            _ = len(v3)
            for fi10 := 0; fi10 < 8; fi10++ {
                if (v4 <= (v4 % ((v4) | 1))) {
                    fmt.Println(fi10, v4)
                    fmt.Println(v1, v2, v1)
                }
                for fi11 := 0; fi11 < 7; fi11++ {
                    v2 = v2
                    _ = fi11
                }
                _ = fi10
            }
        } else {
            var v12 int32 = -1
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 uint8 = v1
                fmt.Println(v4)
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := GenS1{F0: (^(v1)), F1: (v4 + (^(v4))), F2: 911}
            v17 := GenS1{F0: v1, F1: v4, F2: 32767}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    v18 := GenS1{F0: v1, F1: v4, F2: 0}
    v1 += (^(v1))
    v19 := GenS1{F0: 0, F1: v4, F2: 42}
    v1 = 0
    for fi20 := 0; fi20 < 6; fi20++ {
        v2 -= v2
        _ = fi20
    }
    fmt.Println(v1)
    v4 = 331
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
