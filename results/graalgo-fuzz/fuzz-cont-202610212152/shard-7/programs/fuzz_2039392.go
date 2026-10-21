package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int }
type GenS2 struct { F0 int64 }
func genFn0(p0 bool, p1 int) int {
    p1 += p1
    var v6 bool = false
    var v7 uint32 = 338
    v7 += 0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    v4 := []int{}
    v5 := make(map[int]bool)
    v3[32768] = 793
    if true {
        var v8 uint32 = 2147483647
        switch 4294967295 {
        case 0:
            v1 = (^(v1))
        case 32768:
            var v9 uint16 = 360
            _ = v9
        case 10:
            var v10 int64 = 159
            _ = v10
        default:
            v11 := GenS1{F0: 18446744073709551615, F1: false, F2: -9223372036854775808}
            _ = v11
        }
        _ = v8
    }
    fmt.Println(v1)
    v1 -= 1000
    var v12 int8 = 111
    fmt.Println(v12, v12)
    var v13 int = genFn0((377 > -2147483648), 9223372036854775807)
    if (78 != 42) {
        v13 += (v13 >> ((296) & 63))
        var v14 uint8 = 229
        _ = v14
    } else {
        v15 := GenS1{F0: 170, F1: (!((2 > 875) && (!(!true)))), F2: v13}
        var v16 uint32 = 467
        v16 += (v16 >> ((255) & 63))
        _ = v15
        _ = v16
    }
    var v17 uint = 0
    v1 -= (v1 >> ((v17) & 63))
    v12 += int8(v17)
    v18 := GenS2{F0: 9223372036854775807}
    for fi19 := 0; fi19 < 12; fi19++ {
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 uint8 = uint8(v17)
            _ = len(v5)
            var v22 uint32 = 42
            _ = fi20
            _ = v21
            _ = v22
        }
        fmt.Println(fi19, v17, v17)
        v23 := GenS2{F0: 0}
        _ = fi19
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
