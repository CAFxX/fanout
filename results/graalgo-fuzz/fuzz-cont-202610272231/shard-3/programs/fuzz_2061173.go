package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint64 }
func genFn0(p0 uint16, p1 string, p2 uint) uint {
    var v7 uint32 = 1
    v8 := GenS1{F0: 2}
    v9 := []int{-9223372036854775808, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p2
}

func main() {
    var v1 float64 = 0.1
    var v2 string = ""
    var v3 uint8 = 0
    var v4 uint32 = 2147483648
    v5 := make(map[int]int)
    var v6 float32 = ((0.0 + 3.14159) / float32(float64((1.5 - 1e10))))
    v10 := GenS2{F0: 178}
    v4 += (v4 >> ((744) & 63))
    var v11 uint32 = 4294967295
    if true {
        if (false && (592 < 752)) {
            fmt.Println(v1, v6, v11)
            fmt.Println(v4, v2)
        }
        v12 := GenS2{F0: 0}
        v6 = -2.25
        _ = v12
    } else {
        if ((42 >= 103) && true) {
            v13 := GenS3{F0: -32768, F1: 228, F2: 0}
            switch 0 {
            case 32768:
                v4 = v4
            }
            var v14 int16 = -32768
            _ = v13
            _ = v14
        }
    }
    v15 := GenS2{F0: 1}
    _ = len(v5)
    var v16 uint16 = 73
    v17 := GenS2{F0: (v16 % (((v16 - (v16 >> ((0) & 63)))) | 1))}
    v18 := []int{-9223372036854775808, 127, 2147483647, 32768}
    switch 4294967295 {
    case 32767:
        v19 := GenS4{F0: (v3 % ((128) | 1)), F1: 2147483647, F2: uint64(v4)}
        _ = v19
    case 274:
        var v20 int64 = 447
        _ = v20
    }
    var v21 bool = true
    var v22 uint = 818
    var v23 int64 = 42
    var v24 float64 = v1
    v21 = true
    for fi25 := 0; fi25 < 9; fi25++ {
        if len(v18) > 0 {
            v18[0] = (fi25 * 414)
        }
        for fi26 := 0; fi26 < 3; fi26++ {
            v1 += (v1 / (v1 * v1))
            v6 -= (v6 / (v6 * (1.5 - (v6 / (v6 * v6)))))
            if false {
                switch (fi26 ^ fi25) {
                case 1:
                    var v27 uint16 = v16
                    _ = v27
                case 9223372036854775807:
                    v24 -= (1.5 + v1)
                default:
                    v22 = uint(v3)
                }
                v18 = append(v18, 42)
                v18 = append(v18, 127)
            } else {
                if (true || (((v23 & v23) != (-(v23))) && v21)) {
                    v16 += (v16 & v16)
                    v22 -= v22
                }
            }
            _ = fi26
        }
        _ = fi25
    }
    v22 = (v22 & genFn0((v16 | (v16 & (v16 ^ (v16 + 32767)))), ("foo bar" + v2), v22))
    for fi28 := 0; fi28 < 5; fi28++ {
        fmt.Println(fi28)
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
