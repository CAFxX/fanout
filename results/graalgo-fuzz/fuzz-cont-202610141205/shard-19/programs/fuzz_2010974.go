package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
func genFn0() int16 {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := []int{fi4, fi4, 128}
        if (true && false) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = fi4
        _ = v5
    }
    var v6 float64 = float64(((1e10 - 0.1) / 1.5))
    fmt.Println(v6)
    _ = v6
    return -756
}
func genFn1(p0 uint64, p1 uint, p2 uint) uint16 {
    v7 := []int{244, 0}
    v7 = append(v7, 42)
    _ = len(v7)
    switch 127 {
    case 2147483647:
        switch 1 {
        case 32768:
            var v8 string = ""
            _ = v8
        default:
            var v9 uint = 411
            _ = v9
        }
    case -529:
        p2 -= p2
    default:
        var v10 uint32 = 4294967295
        _ = v10
    }
    p2 -= (p2 * p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 442
}

func main() {
    v1 := []uint32{792}
    v2 := []uint32{782, 828, 2147483647, 450}
    v3 := make(map[int]uint32)
    var v11 float64 = 1e10
    v11 = v11
    if (-32768 != 1000) {
        v11 = v11
    } else {
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 int64 = 9223372036854775807
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            v15 := GenS1{F0: 2147483647, F1: 256}
            v16 := GenS1{F0: 274, F1: 42}
            _ = fi14
            _ = v15
            _ = v16
        }
        v17 := GenS1{F0: 32767, F1: genFn1(18446744073709551615, 0, 18446744073709551615)}
        _ = v17
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        v11 -= 3.14159
        var v19 int8 = 25
        _ = fi18
        _ = v19
    }
    fmt.Println(v11)
    var v20 int = -237
    v20 -= (v20 << ((18446744073709551615) & 63))
    var v21 float32 = 0.0
    switch "hello" {
    case "foo bar":
        fmt.Println(v11, v20)
    case "":
        delete(v3, (v20 - (v20 ^ 32768)))
    }
    var v22 int32 = -83
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
