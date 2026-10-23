package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 bool }
type GenS2 struct { F0 uint8; F1 uint16; F2 int64 }

func main() {
    var v1 string = "hello"
    var v2 int8 = 127
    var v3 uint = 4294967295
    var v4 uint32 = 2147483648
    var v5 uint8 = 1
    v6 := []uint32{}
    var v7 int8 = (v2 << (((^(v3))) & 63))
    v8 := GenS1{F0: (^(v5)), F1: -9223372036854775808, F2: (!false)}
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 int64 = 595
        if (!false) {
            var v11 int32 = -452
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    if len(v6) > 0 {
        v6[0] = (v4 << ((v3) & 63))
    }
    fmt.Println(v4, v3, v2)
    var v12 int32 = 2147483647
    v3 += 4294967295
    v13 := GenS1{F0: (v5 << ((991) & 63)), F1: 1, F2: (-2.25 < 3.14159)}
    if (713 < 0) {
        v14 := GenS1{F0: (v5 ^ (v5 - v5)), F1: -826, F2: (!(286 < 42))}
        var v15 uint64 = 698
        _ = v14
        _ = v15
    } else {
        var v16 uint32 = 2
        v3 = (v3 % (((v3 >> ((61) & 63))) | 1))
        for fi17 := 0; fi17 < 16; fi17++ {
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 float64 = -2.25
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
        _ = v16
    }
    fmt.Println(v7, v12)
    var v20 int16 = -1
    if (false || true) {
        v20 -= (v20 - (v20 * -1))
    }
    v7 = v7
    v21 := GenS2{F0: v5, F1: 42, F2: 0}
    var v22 int = 0
    fmt.Println(v3, v5, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
