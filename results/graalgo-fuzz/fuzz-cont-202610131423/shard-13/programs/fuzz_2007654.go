package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0(p0 uint16) bool {
    p0 = (p0 / ((p0) | 1))
    p0 = p0
    v4 := []int{4294967295, 351, 2147483647}
    _ = p0
    _ = v4
    return ("" <= "")
}

func main() {
    var v1 uint32 = 65535
    var v2 int64 = 9223372036854775807
    var v3 bool = (255 != 41)
    v5 := GenS1{F0: -1, F1: -1, F2: 271}
    if (0 <= 42) {
        v6 := GenS1{F0: -1, F1: 0, F2: 531}
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                v9 := GenS1{F0: (fi8 + fi8), F1: (v2 | (v2 - (v2 - v2))), F2: 18446744073709551615}
                v3 = v3
                _ = fi8
                _ = v9
            }
            v1 = (v1 & (v1 << ((1) & 63)))
            _ = fi7
        }
        if (((!(!(712 <= 0))) || v3) || v3) {
            if v3 {
                fmt.Println(v1)
                v10 := []int{127}
                v1 += (v1 << ((0) & 63))
                _ = v10
            } else {
                v11 := []int{432, 127}
                _ = v11
            }
            var v12 float32 = float32(-2.25)
            _ = v12
        }
        _ = v6
    }
    v13 := GenS1{F0: 128, F1: v2, F2: 2}
    var v14 int16 = -745
    v15 := GenS2{F0: 1, F1: 255}
    v16 := []int{127, 32767, 0, 749}
    v1 = (v1 >> ((0) & 63))
    var v17 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v5
}
