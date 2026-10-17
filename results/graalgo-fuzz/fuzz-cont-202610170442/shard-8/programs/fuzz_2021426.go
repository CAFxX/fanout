package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint16 }
type GenS2 struct { F0 int; F1 int16 }
func genFn0() int32 {
    v6 := GenS1{F0: 127, F1: -9223372036854775808, F2: 186}
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 string = "hello"
        for fi9 := 0; fi9 < 10; fi9++ {
            v8 = v8
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    var v10 int16 = -312
    v10 -= 1000
    v10 = v10
    _ = v10
    _ = v6
    return 0
}
func genFn1() int64 {
    var v11 int8 = 18
    v11 += -128
    _ = v11
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 1
    var v3 int = 0
    v4 := make(map[int]int)
    var v5 int = v3
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            switch "foo bar" {
            case "foo bar":
                var v14 int64 = 42
                _ = v14
            case "hello":
                var v15 uint16 = 2
                _ = v15
            case "\n\t":
                var v16 uint16 = 407
                _ = v16
            default:
                v3 -= (fi13 * fi13)
            }
            _ = fi13
        }
        v2 = -1
        _ = fi12
    }
    v5 += (v3 ^ int(v2))
    if false {
        var v17 string = "hello"
        _ = v17
    }
    switch (("foo bar" + "foo bar") + (("\n\t" + "hello") + "a")) {
    case "foo bar":
        v3 += 128
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 uint = 332
        v20 := []int{(v5 >> (((v19 << (((^(v19))) & 63))) & 63))}
        v2 -= (v2 ^ v2)
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: v5, F1: -1}
    v2 = -1
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 int64 = genFn1()
        v24 := GenS2{F0: (v3 << ((852) & 63)), F1: (v2 << ((42) & 63))}
        _ = fi22
        _ = v23
        _ = v24
    }
    delete(v1, 2147483647)
    v25 := GenS1{F0: -1, F1: int(v2), F2: 649}
    v26 := GenS1{F0: 42, F1: (v3 - v3), F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
