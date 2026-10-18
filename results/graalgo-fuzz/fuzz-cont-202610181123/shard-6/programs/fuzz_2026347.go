package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int8 }
type GenS2 struct { F0 int32; F1 int; F2 uint16 }
type GenS3 struct { F0 int8; F1 bool; F2 int }
func genFn0(p0 int) string {
    v5 := GenS1{F0: 255, F1: -2147483648, F2: 1}
    var v6 uint16 = 255
    v7 := GenS1{F0: 2, F1: -1, F2: 100}
    v8 := []int{127, (p0 + p0), (p0 ^ 0)}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return "a"
}

func main() {
    var v1 uint = 735
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    v4 := []uint32{42, 2147483648, 154, 406}
    if false {
        if ((v1 << ((1) & 63)) < v1) {
            v1 = (v1 % ((v1) | 1))
            if true {
                fmt.Println(v1)
                var v9 int64 = 9223372036854775807
                delete(v3, int(v1))
                _ = v9
            } else {
                v10 := GenS1{F0: 51, F1: 947, F2: 1}
                v11 := GenS1{F0: 229, F1: 42, F2: -128}
                _ = v10
                _ = v11
            }
            fmt.Println(v1)
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 uint64 = 2
            switch (fi12 / ((fi12) | 1)) {
            case 128:
                v13 += (v13 >> ((122) & 63))
            case 1:
                var v14 int8 = 127
                _ = v14
            default:
                v15 := GenS2{F0: -941, F1: 890, F2: 54}
                _ = v15
            }
            if (!(32768 > fi12)) {
                v1 = v1
                v2[fi12] = 0
                v16 := GenS1{F0: 128, F1: 0, F2: -1}
                _ = v16
            }
            _ = fi12
            _ = v13
        }
    }
    var v17 uint32 = 2147483647
    v2[2147483647] = 4294967295
    v18 := GenS3{F0: -58, F1: (true && true), F2: 9223372036854775807}
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS3{F0: 127, F1: false, F2: (fi19 * fi19)}
        if false {
            v21 := GenS3{F0: -52, F1: ((v17 - (v17 >> ((4294967295) & 63))) >= (v17 ^ (v17 / ((0) | 1)))), F2: 128}
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    var v22 float32 = (3.14159 / 1.0)
    v1 += 700
    v4 = append(v4, (^(v17)))
    v22 = v22
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
