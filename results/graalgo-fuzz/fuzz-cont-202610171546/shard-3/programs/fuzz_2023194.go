package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
func genFn0(p0 bool) uint8 {
    var v7 int64 = 0
    var v8 int64 = (v7 << ((4294967295) & 63))
    fmt.Println(p0, p0, p0)
    v9 := GenS1{F0: 0}
    switch "hello" {
    case "\n\t":
        v8 -= (v7 * (-(v7)))
    case "a":
        v8 = 0
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 bool) int16 {
    fmt.Println(p0)
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS2{F0: 2147483647}
        fmt.Println(fi10, p0)
        _ = fi10
        _ = v11
    }
    _ = p0
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    v4 := make(map[int]bool)
    var v5 uint16 = 236
    var v6 uint32 = 42
    switch 631 {
    case 0:
        fmt.Println(v5, v6)
    default:
        v5 += (^(v5))
    }
    var v12 int32 = -2147483648
    if false {
        switch ("" + "") {
        case "foo bar":
            delete(v3, -2147483648)
        default:
            var v13 float64 = -2.25
            _ = v13
        }
    } else {
        v2[-241] = 1
        var v14 uint64 = 2
        v3[42] = uint32(v14)
        _ = v14
    }
    v15 := GenS1{F0: 36}
    v5 = 735
    if true {
        v12 = v12
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v4[2147483647] = false
        v5 -= (v5 % (((v5 >> ((676) & 63))) | 1))
        v6 = (^(v6))
        _ = fi16
    }
    v6 = (v6 + 734)
    v12 = int32(v5)
    if false {
        if (!false) {
            fmt.Println(v12, v6, v6)
            v12 += 2147483647
        } else {
            if ((v12 & 354) < 1) {
                v6 -= v6
                var v17 uint8 = genFn0(false)
                var v18 uint64 = 18446744073709551615
                _ = v17
                _ = v18
            } else {
                v6 = (^(v6))
                v6 -= (^(v6))
                var v19 uint32 = (v6 + v6)
                _ = v19
            }
            v5 = v5
        }
    }
    v20 := GenS1{F0: (^(v6))}
    for fi21 := 0; fi21 < 18; fi21++ {
        v22 := GenS1{F0: v6}
        _ = fi21
        _ = v22
    }
    v23 := []int{32767, 4294967295, 128}
    v24 := GenS2{F0: 32768}
    v6 = v6
    var v25 uint8 = 127
    v5 = uint16(v6)
    v12 -= 807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
