package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int; F2 uint8 }

func main() {
    var v1 int16 = -32768
    var v2 int32 = 42
    v3 := []int{-2147483648, -929, 32767, -863}
    v4 := GenS1{F0: "foo bar", F1: 2, F2: 2147483648}
    var v5 int32 = -1
    _ = len(v3)
    v6 := GenS1{F0: ("\n\t" + "hello"), F1: 4294967295, F2: 434}
    var v7 int8 = 42
    switch 9223372036854775807 {
    case 4294967295:
        v8 := GenS2{F0: 1, F1: -1, F2: 127}
        _ = v8
    case -970:
        _ = len(v3)
    default:
        switch -1 {
        case -980:
            fmt.Println(v5, v1, v5)
        case 32768:
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS2{F0: 42, F1: (fi9 * fi9), F2: 42}
                _ = fi9
                _ = v10
            }
        }
    }
    var v11 float32 = 1e10
    if (-1 <= 42) {
        for fi12 := 0; fi12 < 13; fi12++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := GenS2{F0: 2, F1: (fi12 & (fi13 | (-(fi13)))), F2: 42}
                _ = fi13
                _ = v14
            }
            var v15 int32 = (v2 >> ((767) & 63))
            _ = fi12
            _ = v15
        }
    }
    v5 = int32(v7)
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 int32 = v2
        v18 := GenS1{F0: "a", F1: 1, F2: 91}
        _ = fi16
        _ = v17
        _ = v18
    }
    v5 += (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
