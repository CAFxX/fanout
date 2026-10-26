package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int32 }
func genFn0() uint {
    var v7 uint = 2
    if true {
        v7 += v7
        for fi8 := 0; fi8 < 4; fi8++ {
            v7 -= v7
            if true {
                var v9 string = "foo bar"
                _ = v9
            }
            v10 := []int{fi8, fi8, fi8}
            _ = fi8
            _ = v10
        }
        v7 -= (v7 + v7)
    }
    v7 += (v7 << ((v7) & 63))
    var v11 uint = (v7 | v7)
    _ = v11
    _ = v7
    return (v11 - v11)
}
func genFn1(p0 uint16, p1 uint64) uint64 {
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int16 = -1
        v14 := []int{fi12, fi12, fi12, (fi12 & (fi12 | (fi12 >> ((255) & 63))))}
        v15 := GenS1{F0: p1, F1: 118, F2: 366}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    p1 = (^(p1))
    _ = p0
    _ = p1
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 1
    var v3 int = 127
    var v4 uint = 0
    v5 := make(map[string]uint32)
    var v6 int16 = 1000
    var v16 float64 = 0.1
    v2 += (v2 >> ((uint(v2)) & 63))
    v17 := v5["hello"]
    var v18 uint8 = 255
    delete(v1, (v3 | 2147483647))
    if false {
        if ((-1 == -854) || false) {
            var v19 float64 = (3.14159 - v16)
            _ = v19
        } else {
            _ = len(v1)
            var v20 int64 = 42
            var v21 int8 = 23
            _ = v20
            _ = v21
        }
        fmt.Println(v3, v6)
        v2 = v2
    } else {
        v3 -= (v3 | v3)
        var v22 int64 = -9223372036854775808
        _ = len(v1)
        _ = v22
    }
    switch "" {
    case "foo bar":
        v18 -= v18
    }
    var v23 uint32 = 828
    delete(v5, "\n\t")
    for fi24 := 0; fi24 < 4; fi24++ {
        fmt.Println(v23, v17)
        _ = fi24
    }
    fmt.Println(v16)
    for fi25 := 0; fi25 < 18; fi25++ {
        fmt.Println(v18, v17, v6)
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
