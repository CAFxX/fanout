package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint; F1 bool; F2 int8 }
func genFn0(p0 uint32, p1 uint16, p2 uint32) uint {
    p1 = p1
    fmt.Println(p0, p0)
    v7 := GenS1{F0: (-1 == 127), F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}
func genFn1(p0 int8) int8 {
    if false {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 string = ("" + "\n\t")
            var v10 int8 = p0
            _ = fi8
            _ = v10
            _ = v9
        }
        p0 = (p0 % (((p0 / ((-1) | 1))) | 1))
        if false {
            switch 4294967295 {
            case 0:
                v11 := []int{-302, 0, 452}
                _ = v11
            case 32767:
                v12 := []int{32768, -2147483648, 42}
                _ = v12
            default:
                fmt.Println(p0)
            }
            var v13 uint16 = 1
            _ = v13
        }
    }
    switch ("hello" + "foo bar") {
    case "hello":
        var v14 uint32 = 2147483648
        _ = v14
    case "a":
        p0 = p0
    case "\n\t":
        p0 += (-(p0))
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 string = "\n\t"
        p0 = 127
        if false {
            var v17 uint8 = 0
            switch (fi15 >> ((0) & 63)) {
            case -2147483648:
                v18 := GenS1{F0: true, F1: genFn0(2, 2, 326)}
                _ = v18
            default:
                var v19 bool = (false && false)
                _ = v19
            }
            var v20 int8 = p0
            _ = v17
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    _ = p0
    return 42
}

func main() {
    v1 := []int{83, 970}
    var v2 uint16 = 942
    v3 := []string{"hello", "", "a"}
    var v4 string = (("a" + "foo bar") + "foo bar")
    var v5 bool = (false || false)
    var v6 int64 = 0
    var v21 int64 = v6
    var v22 string = v4
    if (384 > -312) {
        fmt.Println(v4)
    }
    if (!true) {
        fmt.Println(v4, v5)
    }
    switch (v4 + v22) {
    case "":
        v4 = (v4 + "a")
    case "a":
        v23 := GenS2{F0: 655, F1: (2 > 2), F2: -128}
        _ = v23
    default:
        if len(v1) > 0 {
            v1[0] = 127
        }
    }
    v4 = (v22 + v22)
    v2 = (v2 / ((313) | 1))
    v22 = v4
    if v5 {
        for fi24 := 0; fi24 < 6; fi24++ {
            _ = len(v3)
            _ = fi24
        }
        v6 += (^(v6))
    }
    if (!true) {
        var v25 int8 = 100
        fmt.Println(v25, v21)
        v26 := GenS1{F0: (true && false), F1: 467}
        _ = v25
        _ = v26
    }
    if (0 == int64(v2)) {
        for fi27 := 0; fi27 < 12; fi27++ {
            v4 = v4
            _ = fi27
        }
        fmt.Println(v2)
        var v28 int8 = 70
        _ = v28
    }
    var v29 uint = 157
    _ = len(v1)
    v4 = v4
    switch "" {
    case "":
        switch -9223372036854775808 {
        case 672:
            v22 = v22
        case -1:
            _ = len(v3)
        case 9223372036854775807:
            v30 := GenS1{F0: ((v29 - v29) > 4294967295), F1: v29}
            _ = v30
        }
    }
    v22 = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
