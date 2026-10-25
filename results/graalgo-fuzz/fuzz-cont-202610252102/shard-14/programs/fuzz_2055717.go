package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint32 }
func genFn0() int {
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := []int{}
        var v9 uint = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint8 = 42
    v10 += (v10 ^ v10)
    var v11 uint8 = v10
    _ = v10
    _ = v11
    return -1
}
func genFn1() int32 {
    if true {
        var v12 uint16 = 2
        v13 := []int{42, 1}
        if len(v13) > 0 {
            v13[0] = -412
        }
        _ = v12
        _ = v13
    } else {
        v14 := []int{-1, 1, 32768, -1}
        var v15 uint8 = 52
        _ = v14
        _ = v15
    }
    var v16 int = 4294967295
    var v17 uint64 = 42
    v18 := []int{}
    _ = v16
    _ = v17
    _ = v18
    return 0
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    var v3 uint16 = 331
    var v4 bool = (!false)
    var v5 uint = 3
    var v6 int32 = 800
    if (v4 || v4) {
        var v19 int16 = -1
        _ = v19
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 uint64 = 0
            v22 := GenS1{F0: 42}
            _ = fi20
            _ = v21
            _ = v22
        }
        v6 += 0
    }
    v23 := []int{-338, -92}
    if len(v23) > 0 {
        v23[0] = 151
    }
    v5 = v5
    var v24 int32 = (-(v6))
    v25 := GenS1{F0: 1}
    switch 9223372036854775807 {
    case 9223372036854775807:
        if (!(599 < -9223372036854775808)) {
            var v26 bool = (v6 > v24)
            v27 := GenS2{F0: genFn1(), F1: 2147483647}
            _ = v26
            _ = v27
        }
    }
    fmt.Println(v4, v4)
    v28 := GenS3{F0: 42}
    v29 := GenS1{F0: -1}
    if v4 {
        v23 = append(v23, -500)
        v24 = v24
    } else {
        delete(v1, int(v3))
        v3 += (v3 % (((v3 ^ (v3 * (v3 & 0)))) | 1))
    }
    fmt.Println(v4, v4, v6)
    switch "foo bar" {
    case "hello":
        v30 := GenS2{F0: -347, F1: 32767}
        _ = v30
    case "a":
        var v31 int32 = 2147483647
        _ = v31
    case "foo bar":
        v3 -= (^(v3))
    }
    fmt.Println(v3)
    fmt.Println(v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
