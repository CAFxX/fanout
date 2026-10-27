package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint8) uint {
    var v5 int = 9223372036854775807
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := []int{int(p0)}
        if ("\n\t" != "\n\t") {
            var v8 uint64 = 305
            _ = v8
        }
        p0 += (p0 / (((p0 << ((236) & 63))) | 1))
        _ = fi6
        _ = v7
    }
    switch -625 {
    case 1:
        p0 -= (^(p0))
    case 9223372036854775807:
        v9 := []int{(v5 + (v5 << ((673) & 63))), (v5 ^ -1), (v5 % ((v5) | 1))}
        _ = v9
    case 127:
        v10 := GenS1{F0: (false || true)}
        _ = v10
    }
    _ = p0
    _ = v5
    return 8
}
func genFn1(p0 int64) int64 {
    fmt.Println(p0)
    var v11 string = "hello"
    _ = p0
    _ = v11
    return (p0 ^ 1)
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -128
    var v3 uint = 214
    v4 := make(map[int]uint32)
    _ = len(v4)
    v3 = (v3 ^ (v3 * (v3 - v3)))
    fmt.Println(v2, v2)
    var v12 int32 = -2147483648
    v2 -= (v2 / ((27) | 1))
    v1[9223372036854775807] = "a"
    if false {
        v13 := []int{-1, 9223372036854775807, 9223372036854775807}
        _ = v13
    } else {
        if (413 == 0) {
            if (!true) {
                delete(v4, 4294967295)
            }
            fmt.Println(v2)
        } else {
            v14 := []int{}
            _ = v14
        }
        var v15 bool = (1 < 754)
        _ = v15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 float64 = 0.1
        _ = fi16
        _ = v17
    }
    v12 -= (v12 ^ (v12 % ((v12) | 1)))
    fmt.Println(v3)
    if false {
        v2 -= (v2 & 1)
        for fi18 := 0; fi18 < 2; fi18++ {
            fmt.Println(v2, fi18, v12)
            _ = fi18
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
