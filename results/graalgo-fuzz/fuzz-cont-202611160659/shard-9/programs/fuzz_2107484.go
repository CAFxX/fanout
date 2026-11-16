package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int64 }

func main() {
    var v1 int = 4294967295
    var v2 uint8 = 192
    var v3 int8 = -128
    v1 += v1
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := []int{(v1 & (v1 + (-(v1)))), v1, (fi4 * 32767), v1}
        v1 += (v1 >> ((0) & 63))
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 2147483648, F1: 18446744073709551615, F2: -529}
    v7 := GenS1{F0: 65535, F1: 398, F2: int64(v1)}
    for fi8 := 0; fi8 < 9; fi8++ {
        v1 = (-(v1))
        _ = fi8
    }
    if ((313 >= 32767) || true) {
        var v9 int16 = 0
        var v10 int16 = v9
        _ = v10
        _ = v9
    }
    v3 = (v3 & v3)
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 int64 = 1
        if true {
            v12 = (-(v12))
        } else {
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 int32 = -516
                _ = fi13
                _ = v14
            }
        }
        _ = fi11
        _ = v12
    }
    v2 += (v2 & v2)
    v1 = 127
    v15 := GenS1{F0: 888, F1: 1, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
