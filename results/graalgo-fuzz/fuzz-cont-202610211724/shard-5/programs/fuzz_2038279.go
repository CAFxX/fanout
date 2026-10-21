package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int32; F2 int8 }
func genFn0(p0 uint64, p1 int32) bool {
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: 32767, F1: 7}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint16 = 32767
        var v10 int8 = -1
        var v11 float64 = ((1e10 * 1e10) * ((-2.25 + 0.1) * 1e10))
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    p0 = (p0 ^ (^(p0)))
    _ = p0
    _ = p1
    return (802 >= -773)
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{-991, 0}
    v3 := []int{}
    var v4 uint8 = 1
    var v5 float32 = 0.0
    if true {
        v5 += (3.14159 / v5)
        fmt.Println(v5, v5)
    } else {
        switch "a" {
        case "foo bar":
            var v12 int8 = 100
            _ = v12
        }
    }
    v13 := GenS1{F0: 88, F1: 412}
    v5 += v5
    v4 = v4
    if len(v3) > 0 {
        v3[0] = 128
    }
    v2 = append(v2, 42)
    switch "foo bar" {
    case "foo bar":
        if false {
            if (!false) {
                var v14 int8 = 26
                _ = len(v3)
                var v15 uint16 = 664
                _ = v14
                _ = v15
            } else {
                fmt.Println(v4, v5)
            }
        } else {
            switch "foo bar" {
            case "a":
                for fi16 := 0; fi16 < 16; fi16++ {
                    v17 := GenS1{F0: 1, F1: 2147483647}
                    if false {
                        fmt.Println(fi16)
                    } else {
                        _ = len(v1)
                        v4 = v4
                    }
                    v4 = 0
                    _ = fi16
                    _ = v17
                }
            }
            if (v4 == 127) {
                v18 := GenS1{F0: 65535, F1: 719}
                var v19 int8 = 1
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v18
                _ = v19
            } else {
                v20 := GenS1{F0: 32767, F1: 0}
                _ = v20
            }
        }
    case "\n\t":
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 uint32 = 42
            _ = fi21
            _ = v22
        }
    }
    fmt.Println(v4)
    for fi23 := 0; fi23 < 12; fi23++ {
        var v24 int32 = -277
        _ = fi23
        _ = v24
    }
    switch ("foo bar" + "") {
    case "":
        v25 := GenS2{F0: 281, F1: -2147483648, F2: 3}
        _ = v25
    }
    var v26 int32 = -803
    v4 = (v4 % ((87) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
