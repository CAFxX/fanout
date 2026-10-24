package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 string; F2 bool }
func genFn0(p0 int8, p1 int16) int64 {
    var v4 uint64 = 4294967295
    v5 := GenS1{F0: 2}
    var v6 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 bool = (255 != 1)
    var v2 int8 = 42
    var v3 int16 = 32767
    if (v1 && (!v1)) {
        if (1e10 > -2.25) {
            var v7 int32 = 2147483647
            v3 -= 32767
            _ = v7
        } else {
            v8 := []int{128, -2147483648, 1}
            v2 += -89
            _ = v8
        }
    } else {
        v3 += (v3 % ((v3) | 1))
    }
    v9 := GenS2{F0: 0}
    v10 := GenS3{F0: 18446744073709551615, F1: ("" + ("" + "\n\t")), F2: v1}
    v3 = v3
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int = (fi11 % (((^(fi11))) | 1))
        _ = fi11
        _ = v12
    }
    var v13 uint16 = 65535
    v14 := []int{}
    switch 32767 {
    case -2147483648:
        v2 = 1
    }
    v14 = append(v14, 2147483647)
    v2 += (v2 * v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
