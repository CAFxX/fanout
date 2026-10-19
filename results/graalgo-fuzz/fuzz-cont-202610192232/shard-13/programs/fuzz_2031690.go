package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint; F1 uint8; F2 int8 }
type GenS3 struct { F0 int32; F1 int; F2 uint64 }
func genFn0() uint32 {
    var v4 uint32 = 1
    v5 := GenS1{F0: 255, F1: -904}
    v6 := []int{32767, 0}
    v4 = v4
    _ = v4
    _ = v5
    _ = v6
    return 2147483648
}
func genFn1(p0 uint32) string {
    p0 += (^(p0))
    v7 := GenS2{F0: 18446744073709551615, F1: 128, F2: 42}
    _ = p0
    _ = v7
    return "a"
}

func main() {
    var v1 uint64 = 4294967295
    v2 := []uint32{1, 65535, 4294967295, 68}
    var v3 int8 = 1
    var v8 int8 = (-(v3))
    if len(v2) > 0 {
        v2[0] = 42
    }
    if (42 >= 9223372036854775807) {
        v9 := GenS3{F0: 2147483647, F1: int(v1), F2: (v1 << ((65535) & 63))}
        v8 -= (v3 << ((65535) & 63))
        var v10 int64 = 9223372036854775807
        _ = v10
        _ = v9
    }
    v8 = v3
    var v11 uint32 = 0
    v1 += 1
    var v12 uint64 = (v1 | 2)
    switch "a" {
    case "":
        if true {
            v13 := GenS2{F0: 2, F1: 0, F2: (^(v3))}
            v2 = append(v2, uint32(v12))
            _ = v13
        } else {
            for fi14 := 0; fi14 < 13; fi14++ {
                if false {
                    v15 := GenS3{F0: 42, F1: fi14, F2: (v12 % ((18446744073709551615) | 1))}
                    v16 := GenS1{F0: 227, F1: 0}
                    _ = v15
                    _ = v16
                }
                _ = fi14
            }
            var v17 bool = false
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(v12, v3, v12)
            var v19 uint32 = (v11 % (((v11 >> ((uint(v1)) & 63))) | 1))
            var v20 uint64 = v1
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    fmt.Println(v12, v1, v8)
    fmt.Println(v1)
    var v21 uint8 = 42
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v8
}
