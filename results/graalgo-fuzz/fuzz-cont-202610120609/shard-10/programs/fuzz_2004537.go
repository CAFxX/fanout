package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int8) int32 {
    v4 := []int{2147483647}
    switch -9223372036854775808 {
    case -2147483648:
        for fi5 := 0; fi5 < 20; fi5++ {
            p0 -= p0
            p0 = p0
            if false {
                v6 := GenS1{F0: 0, F1: 1}
                for fi7 := 0; fi7 < 18; fi7++ {
                    if (false || false) {
                        fmt.Println(fi5)
                    } else {
                        _ = len(v4)
                        switch (fi7 / ((4294967295) | 1)) {
                        case -9223372036854775808:
                            p0 = -1
                        }
                    }
                    var v8 float64 = 1.5
                    var v9 uint = 4294967295
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                p0 -= (p0 / ((p0) | 1))
                _ = v6
            } else {
                p0 = (p0 + int8(fi5))
            }
            _ = fi5
        }
    case 9223372036854775807:
        fmt.Println(p0)
    }
    _ = p0
    _ = v4
    return 42
}
func genFn1(p0 bool) uint {
    p0 = ((!false) && (p0 && true))
    p0 = (1 < 303)
    if p0 {
        v10 := GenS1{F0: 218, F1: 42}
        switch "foo bar" {
        case "\n\t":
            if p0 {
                v11 := []int{-2147483648, 2147483647}
                _ = v11
            }
        case "foo bar":
            p0 = ((255 == 0) && ((p0 && false) && true))
        }
        fmt.Println(p0)
        _ = v10
    } else {
        var v12 uint8 = 0
        var v13 uint64 = 4294967295
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := []int{-937, -310, 4294967295}
            v16 := GenS1{F0: v13, F1: 2}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v12
        _ = v13
    }
    _ = p0
    return 954
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 0
    v3 := make(map[string]bool)
    var v17 bool = false
    v18 := GenS1{F0: 68, F1: 18446744073709551615}
    var v19 bool = false
    var v20 string = "hello"
    var v21 int16 = 42
    for fi22 := 0; fi22 < 12; fi22++ {
        if ((!v19) && (100 > 0)) {
            v23 := []int{(fi22 + (^(fi22))), 127}
            _ = v23
        }
        fmt.Println(v21, v2, fi22)
        _ = fi22
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        v21 = 32767
        fmt.Println(v20, v19)
        _ = fi24
    }
    if v17 {
        var v25 bool = (32768 != -2147483648)
        _ = v25
    } else {
        v26 := []int{-2147483648}
        _ = v26
    }
    v27 := []int{0, 9223372036854775807, 127, 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
}
