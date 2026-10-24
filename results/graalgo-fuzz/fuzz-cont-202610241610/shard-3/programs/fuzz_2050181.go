package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int16; F2 bool }
func genFn0(p0 bool, p1 int8, p2 int8) int8 {
    p0 = p0
    v4 := GenS1{F0: 1000, F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p2
}

func main() {
    var v1 bool = (-1 > 935)
    v2 := []int{9223372036854775807, -303, 127, -15}
    v3 := make(map[int]int)
    var v5 uint32 = 532
    v1 = v1
    v1 = (!false)
    v6 := GenS1{F0: 928, F1: 698}
    fmt.Println(v5, v5)
    v1 = true
    var v7 uint64 = 42
    fmt.Println(v1, v7, v7)
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: 42, F1: 960}
        if v1 {
            v10 := GenS1{F0: 4, F1: (-(fi8))}
            if v1 {
                v1 = v1
                v7 = v7
            } else {
                fmt.Println(v5, v7, fi8)
            }
            var v11 string = ("\n\t" + "foo bar")
            _ = v10
            _ = v11
        }
        v1 = (0 <= 151)
        _ = fi8
        _ = v9
    }
    fmt.Println(v5, v7)
    v12 := GenS1{F0: -32768, F1: -9223372036854775808}
    for fi13 := 0; fi13 < 11; fi13++ {
        if (((0.0 >= ((0.1 - (((0.1 * (1e10 - 3.14159)) * 0.1) + ((0.0 / 1.0) / -2.25))) - ((3.14159 / 3.14159) / 1.5))) && (v7 > (v7 >> ((0) & 63)))) || (42 != uint32(v7))) {
            if (-102 >= 0) {
                v5 -= (v5 * v5)
            }
        } else {
            v14 := GenS2{F0: 942}
            if v1 {
                var v15 string = "a"
                _ = v15
            } else {
                var v16 int64 = -9223372036854775808
                _ = v16
            }
            var v17 uint8 = 42
            _ = v14
            _ = v17
        }
        if v1 {
            v18 := GenS3{F0: 250}
            v5 = (v5 & 65535)
            _ = v18
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            v5 -= v5
            delete(v3, 0)
            _ = fi19
        }
        _ = fi13
    }
    var v20 uint32 = (v5 ^ 607)
    var v21 uint = 195
    v22 := GenS4{F0: (!false), F1: 154, F2: (((v7 * v7) != 476) && v1)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
