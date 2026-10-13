package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint16 }
func genFn0(p0 string) int64 {
    var v7 uint64 = 5
    p0 = p0
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 int16, p1 uint8, p2 bool) bool {
    for fi8 := 0; fi8 < 8; fi8++ {
        if (3.14159 < 1e10) {
            v9 := GenS1{F0: 42, F1: 100, F2: 2}
            if true {
                v10 := GenS2{F0: false, F1: (p1 + (p1 & p1)), F2: 0}
                var v11 int8 = -128
                v12 := []int{(-(fi8)), (fi8 * fi8), (fi8 ^ (fi8 * fi8)), fi8}
                _ = v10
                _ = v11
                _ = v12
            }
            _ = v9
        }
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 uint = 1
            var v15 int32 = -1
            v16 := []int{32767, int(p1)}
            if (0 != 255) {
                if p2 {
                    v17 := GenS2{F0: (!true), F1: (p1 << ((v14) & 63)), F2: 729}
                    _ = v17
                } else {
                    var v18 int32 = v15
                    p2 = (uint64(p1) > 0)
                    fmt.Println(p1, v15)
                    fmt.Println(p0, fi13, v18)
                    var v19 int32 = (v15 >> (((v14 * 1)) & 63))
                    _ = v18
                    _ = v19
                }
                p0 = 0
                v20 := GenS1{F0: 65535, F1: 42, F2: p1}
                _ = v20
            }
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            v22 := GenS2{F0: (0 == 562), F1: p1, F2: 908}
            var v23 uint64 = 4294967295
            fmt.Println(p0)
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi8
    }
    v24 := []int{-244}
    v25 := GenS1{F0: 1, F1: -10, F2: 218}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    return false
}

func main() {
    v1 := []uint32{2147483648, 1}
    var v2 int64 = 9223372036854775807
    var v3 int16 = -1
    v4 := make(map[string]string)
    v5 := make(map[int]int)
    v6 := make(map[int]string)
    var v26 uint8 = 255
    switch "\n\t" {
    case "":
        switch -1 {
        case -774:
            var v27 int16 = (-(v3))
            _ = v27
        case 2147483647:
            for fi28 := 0; fi28 < 19; fi28++ {
                switch fi28 {
                case 0:
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(fi28, v2)
                    case "a":
                        fmt.Println(v2, v3)
                    case "":
                        var v29 float64 = ((3.14159 / 0.1) / 1.0)
                        _ = v29
                    }
                case 435:
                    var v30 int32 = 42
                    _ = v30
                case 128:
                    switch (fi28 >> ((0) & 63)) {
                    case -283:
                        v31 := GenS2{F0: true, F1: v26, F2: 65535}
                        _ = v31
                    }
                default:
                    v3 = (v3 & (v3 % ((v3) | 1)))
                }
                _ = fi28
            }
        }
    case "\n\t":
        v6[-499] = "a"
    case "hello":
        v2 += genFn0("hello")
    }
    if true {
        if false {
            v26 += (v26 ^ v26)
        }
        v26 = v26
        v26 = (v26 + v26)
    } else {
        var v32 uint16 = 137
        for fi33 := 0; fi33 < 20; fi33++ {
            for fi34 := 0; fi34 < 17; fi34++ {
                _ = len(v1)
                var v35 int64 = -9223372036854775808
                v36 := GenS2{F0: false, F1: v26, F2: 255}
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = fi33
        }
        var v37 uint32 = 798
        _ = v32
        _ = v37
    }
    _ = len(v5)
    fmt.Println(v2, v3, v3)
    v3 += v3
    v38 := GenS2{F0: (2 == 1), F1: (^(v26)), F2: 32767}
    fmt.Println(v2, v2, v26)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
