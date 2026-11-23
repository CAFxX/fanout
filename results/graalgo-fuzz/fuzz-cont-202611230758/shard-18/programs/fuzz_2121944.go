package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 int }
func genFn0(p0 uint, p1 int16) uint16 {
    v5 := GenS1{F0: 18446744073709551615}
    var v6 uint32 = 1
    v7 := []int{-962, 4294967295, -9223372036854775808, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 1e10
    var v3 int = 32768
    v4 := []int{}
    var v8 uint64 = 1
    var v9 uint32 = 2147483648
    v10 := GenS2{F0: (v8 & (v8 >> ((255) & 63))), F1: (^(v9)), F2: 763}
    v11 := GenS1{F0: (v8 - (v8 << ((uint(v8)) & 63)))}
    v12 := GenS3{F0: (v9 % ((v9) | 1))}
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint = 4294967295
        _ = fi13
        _ = v14
    }
    v2 = 1.5
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 bool = ((v3 != 583) || false)
        var v17 uint32 = 694
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = len(v1)
    v8 = (v8 - v8)
    var v18 int32 = 2147483647
    if false {
        for fi19 := 0; fi19 < 8; fi19++ {
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 float64 = 3.14159
                if false {
                    fmt.Println(v21, v2)
                    v18 = -2147483648
                    v22 := GenS4{F0: 1, F1: (v3 ^ (fi20 | 2147483647))}
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        v9 -= (v9 % ((v9) | 1))
        fmt.Println(v8)
    }
    v9 = 4294967295
    switch int(v8) {
    case 128:
        var v23 int64 = 9223372036854775807
        _ = v23
    case -430:
        v8 = (v8 | v8)
    case -59:
        if (!(!(42 != 939))) {
            v24 := GenS4{F0: 100, F1: (^(v3))}
            var v25 int8 = 6
            v3 += (v3 >> ((0) & 63))
            _ = v24
            _ = v25
        }
    }
    var v26 uint64 = v8
    v3 = (v3 - (v3 & (-(v3))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
