package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 float32 = (1.5 * 1e10)
    var v4 int16 = 1
    v5 := GenS1{F0: 627, F1: 4294967295, F2: 42}
    v6 := GenS1{F0: 0, F1: 267, F2: 946}
    v3 = v3
    fmt.Println(v3)
    switch 9223372036854775807 {
    case -219:
        v7 := v1[168]
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: 65535, F1: 18446744073709551615, F2: (v4 * 32767)}
        v3 -= v3
        if false {
            var v10 float32 = 3.14159
            v11 := []int{}
            v12 := GenS1{F0: 1, F1: 0, F2: 32767}
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 uint32 = 4294967295
        for fi15 := 0; fi15 < 19; fi15++ {
            switch "" {
            case "foo bar":
                v3 = v3
            case "hello":
                fmt.Println(fi13)
            default:
                if true {
                    var v16 uint32 = 39
                    _ = v16
                } else {
                    v17 := []int{fi15, fi13, (fi13 % (((fi15 % ((fi13) | 1))) | 1))}
                    _ = v17
                }
            }
            v3 += ((v3 * 3.14159) * (1.5 / 1.0))
            switch "hello" {
            case "a":
                var v18 int64 = 1
                _ = v18
            case "foo bar":
                var v19 uint32 = v14
                _ = v19
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    if true {
        switch 9223372036854775807 {
        case 32768:
            var v20 uint16 = 689
            _ = v20
        }
    }
    var v21 uint = 4294967295
    v21 -= 866
    if (206 <= 1) {
        var v22 uint16 = 32767
        fmt.Println(v3)
        _ = v22
    } else {
        var v23 int = 389
        _ = v23
    }
    if (2147483647 < -2147483648) {
        var v24 float32 = v3
        var v25 string = (("foo bar" + "hello") + "foo bar")
        _ = v24
        _ = v25
    } else {
        fmt.Println(v4, v4)
        v3 += v3
        v26 := []int{-1, 128}
        _ = v26
    }
    var v27 float32 = 0.1
    v3 = v27
    fmt.Println(v3, v4, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
