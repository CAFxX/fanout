package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int; F1 uint16; F2 uint64 }
func genFn0(p0 int16, p1 int64, p2 uint32) int8 {
    if (255 <= 423) {
        p1 += (p1 * (p1 & p1))
        p0 += (p0 % ((p0) | 1))
    }
    var v6 int64 = 1
    var v7 int = -467
    v8 := []int{v7, v7, (v7 | (^(v7)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]string)
    var v4 uint32 = v2
    v5 := make(map[string]string)
    fmt.Println(v4)
    var v9 uint16 = 274
    var v10 int8 = 0
    v11 := GenS1{F0: true}
    fmt.Println(v9)
    for fi12 := 0; fi12 < 11; fi12++ {
        if true {
            if false {
                v13 := GenS2{F0: 4294967295, F1: 1}
                v1[fi12] = "foo bar"
                _ = v13
            } else {
                v2 += v4
            }
            if true {
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 bool = true
                    _ = fi14
                    _ = v15
                }
            } else {
                var v16 int16 = -32768
                _ = v16
            }
        }
        v3[v2] = "hello"
        v17 := []int{fi12, (fi12 + -1)}
        _ = fi12
        _ = v17
    }
    if (536 == v9) {
        v9 = v9
        for fi18 := 0; fi18 < 7; fi18++ {
            v19 := GenS3{F0: (fi18 % ((fi18) | 1)), F1: (v9 >> ((18446744073709551615) & 63)), F2: 18446744073709551615}
            if true {
                v20 := GenS2{F0: 32, F1: 9223372036854775807}
                _ = v20
            }
            var v21 uint16 = (v9 & v9)
            _ = fi18
            _ = v19
            _ = v21
        }
        v22 := []int{127, -333, 614}
        _ = v22
    }
    v10 = v10
    v4 = (v2 >> ((333) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
