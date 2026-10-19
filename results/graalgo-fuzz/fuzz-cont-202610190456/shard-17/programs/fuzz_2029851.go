package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint32; F1 bool }
func genFn0(p0 string, p1 uint16) int16 {
    v6 := []int{-893, -9223372036854775808, 32767}
    p0 = p0
    for fi7 := 0; fi7 < 13; fi7++ {
        p0 = (((((((p0 + (p0 + "hello")) + "") + p0) + p0) + "foo bar") + p0) + "\n\t")
        if (209 != 936) {
            _ = len(v6)
            fmt.Println(p0, fi7, p1)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            p1 -= p1
        }
        _ = fi7
    }
    v8 := GenS1{F0: -32768}
    v6 = append(v6, 127)
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 32767
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 1
    v3 := make(map[uint32]string)
    v4 := []uint32{2, 2, 2, 2147483648}
    var v5 float64 = 3.14159
    v9 := GenS2{F0: true, F1: ("" + "hello")}
    v2 -= (v2 * (v2 * v2))
    _ = len(v4)
    v10 := v3[2]
    var v11 int8 = 1
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS2{F0: true, F1: ((v10 + v10) + "\n\t")}
            var v15 uint16 = 2
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    v16 := GenS3{F0: 658, F1: true}
    var v17 uint64 = 161
    delete(v1, 566)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
