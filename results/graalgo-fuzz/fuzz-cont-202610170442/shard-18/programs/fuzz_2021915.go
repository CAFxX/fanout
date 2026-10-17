package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint; F1 uint64; F2 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 bool) int64 {
    var v4 int16 = 32767
    if (0 > 65535) {
        var v5 uint64 = 327
        _ = v5
    }
    var v6 uint = 65535
    _ = p0
    _ = v4
    _ = v6
    return 1
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]bool)
    var v3 int8 = 1
    switch "a" {
    case "\n\t":
        v1 += v1
    case "hello":
        for fi7 := 0; fi7 < 10; fi7++ {
            switch "" {
            case "\n\t":
                for fi8 := 0; fi8 < 13; fi8++ {
                    v9 := GenS1{F0: (v1 << ((18446744073709551615) & 63)), F1: int16(v3)}
                    _ = fi8
                    _ = v9
                }
            default:
                v1 = v1
            }
            var v10 uint8 = 110
            var v11 int32 = -2147483648
            _ = fi7
            _ = v10
            _ = v11
        }
    case "foo bar":
        v12 := []int{9223372036854775807, 42, -254, int(v1)}
        _ = v12
    }
    v13 := GenS2{F0: 18446744073709551615, F1: (^(v1)), F2: v3}
    if (!false) {
        v1 = 4294967295
        fmt.Println(v1, v3)
        v1 = (v1 << ((255) & 63))
    } else {
        v14 := []int{-867}
        _ = v14
    }
    v1 = v1
    v15 := GenS1{F0: 4294967295, F1: -476}
    v3 = v3
    v16 := GenS1{F0: (v1 | (v1 - v1)), F1: 1}
    v17 := GenS1{F0: v1, F1: -368}
    v18 := GenS3{F0: 32767}
    delete(v2, ("\n\t" + "a"))
    fmt.Println(v1, v1)
    v3 = 42
    v19 := GenS2{F0: 961, F1: 2, F2: (v3 ^ v3)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
