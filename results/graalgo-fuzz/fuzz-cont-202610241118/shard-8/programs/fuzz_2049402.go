package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
func genFn0(p0 bool, p1 uint64) int16 {
    v6 := []int{}
    p0 = p0
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            fmt.Println(fi7)
            var v9 int32 = -1
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if false {
        switch -923 {
        case 148:
            var v10 uint64 = 18446744073709551615
            _ = v10
        case 127:
            v11 := GenS1{F0: -9223372036854775808, F1: 397, F2: false}
            _ = v11
        default:
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS1{F0: -900, F1: 65535, F2: (255 <= 42)}
                _ = fi12
                _ = v13
            }
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return -569
}
func genFn1() int8 {
    v14 := []int{128, 2147483647, 1, 127}
    var v15 float64 = 3.14159
    _ = v14
    _ = v15
    return -128
}

func main() {
    var v1 int32 = 0
    v2 := make(map[int]int)
    var v3 uint = 255
    var v4 uint16 = 390
    v5 := make(map[uint32]int)
    v16 := GenS1{F0: -9223372036854775808, F1: 65535, F2: (!true)}
    var v17 uint64 = 656
    v17 = (v17 + (v17 & 18446744073709551615))
    v4 -= (v4 & v4)
    fmt.Println(v17, v17)
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 int16 = 32767
        _ = fi18
        _ = v19
    }
    var v20 int8 = -1
    v4 += (v4 ^ v4)
    var v21 string = ("\n\t" + "\n\t")
    v22 := []int{4294967295, 9223372036854775807, 4294967295}
    fmt.Println(v21)
    v2[-2147483648] = -573
    fmt.Println(v20, v20)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
