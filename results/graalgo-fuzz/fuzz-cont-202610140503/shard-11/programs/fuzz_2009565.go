package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint32 }
func genFn0(p0 int) bool {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 18; fi5++ {
        p0 = p0
        var v6 int64 = 42
        _ = fi5
        _ = v6
    }
    p0 = -600
    var v7 float64 = 1.5
    v8 := []int{9223372036854775807, p0, (p0 * (p0 >> ((4294967295) & 63)))}
    _ = p0
    _ = v7
    _ = v8
    return (!false)
}
func genFn1() int8 {
    var v9 uint8 = 1
    v9 = (v9 >> ((971) & 63))
    v10 := GenS1{F0: 2, F1: 42, F2: 650}
    if true {
        v9 = (v9 & v9)
        if (2 > 2) {
            v11 := []int{32767}
            v12 := GenS1{F0: (v9 | v9), F1: -777, F2: 42}
            v9 = v9
            _ = v11
            _ = v12
        }
    }
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 string = "hello"
    v2 := make(map[int]string)
    var v3 uint = 1
    v4 := make(map[uint32]int)
    if false {
        v3 += v3
        v13 := GenS1{F0: 128, F1: 42, F2: 2147483648}
        v3 = 1
        _ = v13
    }
    v1 = v1
    if false {
        fmt.Println(v3, v3)
    }
    v3 = v3
    v14 := GenS1{F0: 191, F1: 1, F2: 454}
    v1 = "a"
    if (true || (0 != 0)) {
        switch -474 {
        case 32767:
            fmt.Println(v1, v1)
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := []int{(fi15 + (fi15 + (fi15 % ((fi15) | 1)))), -9223372036854775808, 277}
            v17 := GenS1{F0: 42, F1: 0, F2: 607}
            v1 = v1
            _ = fi15
            _ = v16
            _ = v17
        }
        v18 := []int{}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 16; fi19++ {
            v1 = ("foo bar" + v1)
            fmt.Println(v1, v1, v3)
            _ = fi19
        }
    }
    fmt.Println(v3, v1)
    fmt.Println(v1, v3)
    if (!true) {
        v20 := []int{-1, 32767, 1}
        var v21 uint64 = 4294967295
        _ = v20
        _ = v21
    }
    var v22 int16 = -539
    fmt.Println(v1, v22, v22)
    v23 := GenS1{F0: 134, F1: 334, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
