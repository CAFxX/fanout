package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint64 }
func genFn0(p0 int64, p1 int8, p2 int64) int32 {
    v7 := GenS1{F0: 90, F1: true, F2: 18446744073709551615}
    var v8 string = "foo bar"
    p0 += p0
    if (873 != 143) {
        var v9 float64 = -2.25
        switch 127 {
        case 575:
            var v10 uint32 = 2
            _ = v10
        case 9223372036854775807:
            fmt.Println(v8)
        case 32768:
            for fi11 := 0; fi11 < 12; fi11++ {
                for fi12 := 0; fi12 < 7; fi12++ {
                    v9 = (v9 * (v9 / v9))
                    p0 += 0
                    _ = fi12
                }
                _ = fi11
            }
        }
        for fi13 := 0; fi13 < 8; fi13++ {
            v8 = (v8 + v8)
            p0 = p0
            _ = fi13
        }
        _ = v9
    } else {
        switch -819 {
        case 0:
            fmt.Println(p0)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 417
}
func genFn1(p0 int, p1 int16) uint32 {
    if true {
        v14 := []int{}
        v15 := GenS1{F0: 4294967295, F1: (true || (((0.0 / ((3.14159 * 0.1) + float32(-2.25))) <= (3.14159 - (1e10 + float32(3.14159)))) && (-128 >= 110))), F2: 42}
        var v16 float32 = 1e10
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := []int{(-(p0))}
    switch p0 {
    case 0:
        fmt.Println(p0, p0)
    default:
        p1 = 1
    }
    _ = p0
    _ = p1
    _ = v17
    return 166
}

func main() {
    var v1 uint64 = 76
    var v2 float64 = (((1.5 + 3.14159) - 0.1) / -2.25)
    v3 := make(map[int]int)
    var v4 int64 = 678
    var v5 string = ("\n\t" + "")
    v6 := make(map[uint32]int)
    v2 -= (float64((-2.25 - 1.5)) - (float64(0.0) + (0.0 * -2.25)))
    v18 := []int{2147483647, 9223372036854775807, 962, 1}
    switch v5 {
    case "a":
        for fi19 := 0; fi19 < 9; fi19++ {
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            if false {
                var v20 uint8 = 128
                _ = v20
            } else {
                v4 = -1
                v4 = (v4 >> ((350) & 63))
            }
            _ = fi19
        }
    case "foo bar":
        v4 -= v4
    }
    fmt.Println(v5, v5)
    v4 += 9223372036854775807
    fmt.Println(v5, v1)
    if len(v18) > 0 {
        v18[0] = -9223372036854775808
    }
    var v21 int16 = 0
    v4 = (v4 % ((1) | 1))
    v21 += (v21 >> ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
