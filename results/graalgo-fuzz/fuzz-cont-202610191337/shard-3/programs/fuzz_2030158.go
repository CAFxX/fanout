package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint }
func genFn0() int16 {
    if false {
        var v4 uint = 448
        v4 = 42
        var v5 string = ("" + "hello")
        _ = v4
        _ = v5
    } else {
        v6 := []int{}
        _ = v6
    }
    var v7 int16 = -542
    v7 = (-(v7))
    if true {
        v8 := GenS1{F0: 42, F1: 35, F2: 65535}
        _ = v8
    }
    var v9 int64 = 0
    _ = v7
    _ = v9
    return (v7 ^ v7)
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{128, 736, 9223372036854775807, 9223372036854775807}
    var v3 int32 = 0
    if true {
        switch "" {
        case "":
            v3 = (v3 | 154)
        case "a":
            var v10 uint16 = 255
            _ = v10
        default:
            var v11 uint8 = 2
            _ = v11
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (1 >= 42) {
            switch ("a" + "\n\t") {
            case "foo bar":
                var v12 float32 = 1.5
                _ = v12
            }
        } else {
            v13 := GenS1{F0: 255, F1: 2, F2: 42}
            for fi14 := 0; fi14 < 19; fi14++ {
                v3 = -2147483648
                v3 = 2147483647
                var v15 uint64 = 0
                _ = fi14
                _ = v15
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v13
        }
    }
    var v16 uint16 = 42
    var v17 float32 = (0.1 / float32(1e10))
    fmt.Println(v17, v16, v17)
    v3 = int32(v16)
    var v18 uint64 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v19 := GenS1{F0: 600, F1: 0, F2: 4294967295}
    switch 32767 {
    case -679:
        v16 += 1
    case -179:
        for fi20 := 0; fi20 < 2; fi20++ {
            if false {
                v21 := GenS1{F0: 1, F1: 128, F2: 255}
                v18 = (v18 / (((v18 | (v18 << ((18446744073709551615) & 63)))) | 1))
                v17 = -2.25
                _ = v21
            } else {
                v18 += (v18 + 313)
                var v22 uint8 = 17
                _ = v22
            }
            _ = fi20
        }
    case -193:
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 bool = (false && true)
            _ = fi23
            _ = v24
        }
    }
    var v25 uint32 = 0
    v25 = v25
    var v26 int64 = 9223372036854775807
    for fi27 := 0; fi27 < 20; fi27++ {
        switch ("\n\t" + "hello") {
        case "hello":
            v17 -= (float32(0.1) / float32(1e10))
        default:
            fmt.Println(v18)
        }
        var v28 float32 = v17
        _ = len(v2)
        _ = fi27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
