package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
func genFn0(p0 uint64, p1 uint, p2 uint) bool {
    switch 4294967295 {
    case 0:
        for fi6 := 0; fi6 < 17; fi6++ {
            switch 128 {
            case -164:
                v7 := []int{(fi6 >> ((p1) & 63)), 325}
                _ = v7
            default:
                fmt.Println(p2, p1)
            }
            fmt.Println(p2, p0, fi6)
            _ = fi6
        }
    case -1:
        v8 := GenS1{F0: 2147483648, F1: 391}
        _ = v8
    default:
        p1 = p1
    }
    v9 := GenS1{F0: 4294967295, F1: 838}
    var v10 int64 = 9223372036854775807
    fmt.Println(p2, p0, p2)
    var v11 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return ((127 > 127) || (-9 > 128))
}

func main() {
    v1 := []string{"a", ("foo bar" + ""), ("hello" + "")}
    var v2 uint = 1
    var v3 int = 118
    var v4 uint64 = 0
    v5 := make(map[uint32]string)
    v4 += (v4 - 2)
    var v12 int32 = 0
    v4 += 4294967295
    v3 += (v3 % (((v3 >> ((4294967295) & 63))) | 1))
    v2 += (v2 - v2)
    v2 += (v2 & (v2 - (v2 + v2)))
    fmt.Println(v3, v3, v4)
    v13 := GenS1{F0: 2, F1: 1}
    v2 += (v2 << (((^(v2))) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
