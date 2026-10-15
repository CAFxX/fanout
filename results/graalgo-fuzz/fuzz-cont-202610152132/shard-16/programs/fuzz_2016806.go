package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint16 }
func genFn0() uint32 {
    v7 := []int{0, 42, 128, 127}
    var v8 uint8 = 242
    switch "" {
    case "a":
        var v9 uint = 0
        _ = v9
    case "foo bar":
        if false {
            v10 := GenS1{F0: 745, F1: 42, F2: 42}
            for fi11 := 0; fi11 < 7; fi11++ {
                switch "hello" {
                case "foo bar":
                    if true {
                        var v12 int16 = -929
                        var v13 bool = (-2.25 != -2.25)
                        for fi14 := 0; fi14 < 17; fi14++ {
                            var v15 bool = v13
                            _ = fi14
                            _ = v15
                        }
                        _ = v12
                        _ = v13
                    } else {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        fmt.Println(fi11, fi11)
                    }
                default:
                    if true {
                        v8 -= (v8 ^ 157)
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        for fi16 := 0; fi16 < 3; fi16++ {
                            v17 := GenS1{F0: 4294967295, F1: -178, F2: 886}
                            _ = fi16
                            _ = v17
                        }
                    }
                }
                _ = fi11
            }
            _ = v10
        }
    }
    if ((0.0 - 1.5) < 3.14159) {
        var v18 int64 = -9223372036854775808
        var v19 string = "a"
        _ = v18
        _ = v19
    }
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 1
    var v3 uint64 = 824
    var v4 int16 = -314
    v5 := make(map[string]bool)
    var v6 string = ""
    v2 += v4
    fmt.Println(v4, v2, v4)
    switch (v6 + v6) {
    case "hello":
        v4 += -1
    default:
        var v20 uint32 = 33
        _ = v20
    }
    if (!true) {
        v6 = v6
        v3 = 661
        fmt.Println(v6, v3)
    }
    v21 := []int{9223372036854775807}
    v4 = v4
    var v22 uint = 0
    var v23 float32 = -2.25
    fmt.Println(v3, v2)
    v2 = (v4 + 0)
    var v24 int32 = -194
    for fi25 := 0; fi25 < 12; fi25++ {
        fmt.Println(v23, v22, v3)
        v26 := GenS1{F0: v3, F1: 42, F2: 255}
        _ = fi25
        _ = v26
    }
    switch "" {
    case "foo bar":
        if true {
            switch "hello" {
            case "":
                var v27 int16 = (v4 & (-(v4)))
                _ = v27
            default:
                v4 += (v2 | v4)
            }
            v28 := GenS1{F0: 4294967295, F1: -9223372036854775808, F2: 255}
            _ = v28
        }
    case "":
        fmt.Println(v24, v6)
    case "a":
        for fi29 := 0; fi29 < 19; fi29++ {
            var v30 int8 = 0
            var v31 float64 = (1e10 / (-2.25 * 0.1))
            v22 += v22
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        switch v6 {
        case "":
            var v33 uint = v22
            _ = v33
        case "\n\t":
            v2 = (v4 - 0)
        case "hello":
            var v34 uint64 = (v3 / ((409) | 1))
            _ = v34
        }
        fmt.Println(fi32, v2)
        _ = fi32
    }
    v1[65535] = -2147483648
    switch 2147483647 {
    case -9223372036854775808:
        v35 := GenS1{F0: v3, F1: -1, F2: 0}
        _ = v35
    case -433:
        fmt.Println(v3)
    default:
        v24 -= 2147483647
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        var v37 int64 = -969
        _ = fi36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
