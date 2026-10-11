package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint64 }
func genFn0(p0 uint32, p1 string, p2 uint64) uint8 {
    fmt.Println(p1, p0)
    var v7 int16 = 1
    for fi8 := 0; fi8 < 11; fi8++ {
        p2 += 18446744073709551615
        _ = fi8
    }
    v7 = (v7 & 1000)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 uint64, p1 int64, p2 int32) int32 {
    v9 := []int{-391, -105, 2, 284}
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p2
}

func main() {
    var v1 int = -9223372036854775808
    var v2 uint16 = 65535
    var v3 int = v1
    v4 := []int{(^(v3))}
    var v5 string = "a"
    var v6 uint = 65535
    v6 += v6
    var v10 uint8 = 1
    v4 = append(v4, v3)
    switch v5 {
    case "\n\t":
        for fi11 := 0; fi11 < 20; fi11++ {
            fmt.Println(fi11)
            _ = fi11
        }
    default:
        if false {
            var v12 int16 = -1
            _ = v12
        }
    }
    v13 := GenS1{F0: 1, F1: uint64(v10)}
    if (!true) {
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := GenS2{F0: 42, F1: 52}
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v5 = v5
                v5 = v5
            }
            v10 += genFn0(65535, "foo bar", 18446744073709551615)
            var v17 int64 = int64(v10)
            _ = fi16
            _ = v17
        }
    } else {
        var v18 int8 = int8(v6)
        _ = v18
    }
    var v19 uint64 = 18446744073709551615
    _ = len(v4)
    v1 = (v1 << (((v6 + uint(v2))) & 63))
    v4 = append(v4, (v1 << ((v6) & 63)))
    v20 := GenS3{F0: v6, F1: (v10 << (((v6 / (((v6 % ((42) | 1))) | 1))) & 63)), F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
