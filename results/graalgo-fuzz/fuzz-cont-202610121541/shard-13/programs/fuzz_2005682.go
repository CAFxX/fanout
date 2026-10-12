package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 string; F1 uint8 }

func main() {
    var v1 int = 4294967295
    v2 := []uint32{46, 2147483647}
    var v3 int32 = 135
    v4 := make(map[uint32]int)
    v5 := make(map[string]bool)
    v1 += v1
    var v6 uint = 797
    if false {
        if true {
            v1 += (v1 + v1)
        }
        v7 := GenS1{F0: 42, F1: -128, F2: (799 <= 2)}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS1{F0: 42, F1: -128, F2: (true && (42 != -9223372036854775808))}
            v2 = append(v2, 0)
            _ = fi8
            _ = v9
        }
        v6 -= 18446744073709551615
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        switch v1 {
        case 324:
            v11 := GenS2{F0: "hello", F1: 1}
            _ = v11
        }
        var v12 int16 = 1000
        _ = fi10
        _ = v12
    }
    fmt.Println(v3, v1)
    var v13 int = v1
    for fi14 := 0; fi14 < 10; fi14++ {
        v3 = v3
        var v15 uint32 = 42
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        switch (v13 * (v13 >> ((v6) & 63))) {
        case 0:
            switch (fi16 % ((fi16) | 1)) {
            case -2147483648:
                var v17 float64 = -2.25
                _ = v17
            }
        default:
            v3 -= -655
        }
        _ = fi16
    }
    if (-9223372036854775808 == 0) {
        fmt.Println(v13, v1, v3)
    }
    var v18 uint = v6
    v19 := GenS2{F0: ("" + "foo bar"), F1: 63}
    v13 = v1
    var v20 uint64 = 386
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
