package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint8; F2 bool }

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 42
    var v3 int8 = 0
    var v4 uint64 = 4294967295
    v5 := make(map[int]int)
    if false {
        var v6 float64 = -2.25
        v7 := GenS1{F0: -566, F1: "hello", F2: v2}
        _ = v6
        _ = v7
    }
    if false {
        v8 := GenS1{F0: 2147483647, F1: "a", F2: (v2 + v2)}
        v2 -= 42
        _ = v8
    }
    if false {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := []int{}
            v11 := GenS2{F0: (0.1 >= 3.14159), F1: uint16(v4), F2: v2}
            fmt.Println(fi9, v2, v4)
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    if (127 < 38) {
        v12 := []int{-607, -862, 4294967295}
        v3 -= (v3 & (v3 | v3))
        _ = v12
    }
    if (false && true) {
        v13 := v5[32767]
        var v14 uint8 = 85
        var v15 uint32 = 767
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 int = 2147483647
    fmt.Println(v2)
    v17 := []int{(-(v16))}
    if true {
        fmt.Println(v4)
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        var v18 int16 = int16(v4)
        _ = v18
    }
    _ = len(v1)
    v19 := GenS2{F0: false, F1: 847, F2: (v2 ^ (v2 % ((uint16(v4)) | 1)))}
    for fi20 := 0; fi20 < 10; fi20++ {
        _ = len(v5)
        var v21 int64 = 0
        _ = fi20
        _ = v21
    }
    if true {
        switch "foo bar" {
        case "a":
            delete(v1, v16)
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            v4 = 2
            v16 = 0
            _ = fi22
        }
        switch (-(v16)) {
        case 2147483647:
            if (382 >= (^(v2))) {
                var v23 int8 = v3
                for fi24 := 0; fi24 < 12; fi24++ {
                    for fi25 := 0; fi25 < 19; fi25++ {
                        v3 = (v3 | (v23 ^ (v3 - v3)))
                        _ = fi25
                    }
                    var v26 string = "hello"
                    fmt.Println(v23, v26)
                    _ = fi24
                    _ = v26
                }
                _ = v23
            }
        default:
            if (false || ((!true) || (32767 <= -1))) {
                v27 := GenS2{F0: (3.14159 > 0.0), F1: 1, F2: v2}
                _ = v27
            } else {
                var v28 bool = false
                _ = v28
            }
        }
    } else {
        fmt.Println(v4, v4, v16)
        if true {
            if true {
                var v29 uint8 = 70
                _ = v29
            } else {
                v16 += v16
                v4 = v4
            }
            v30 := GenS3{F0: -1, F1: 10, F2: ((v2 << ((18446744073709551615) & 63)) == (v2 / (((^(v2))) | 1)))}
            var v31 int16 = -32768
            _ = v30
            _ = v31
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
