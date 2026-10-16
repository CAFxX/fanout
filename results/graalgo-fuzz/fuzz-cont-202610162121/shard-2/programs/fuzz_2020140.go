package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int }
type GenS2 struct { F0 uint16; F1 uint64; F2 int }
func genFn0(p0 int64, p1 uint8) int8 {
    if (2 != 2147483647) {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(p0, p1)
        case "\n\t":
            var v6 uint32 = 2
            _ = v6
        default:
            switch "foo bar" {
            case "hello":
                fmt.Println(p1, p1)
            case "\n\t":
                v7 := GenS1{F0: 227, F1: 18446744073709551615, F2: -443}
                _ = v7
            }
        }
        var v8 uint32 = 42
        _ = v8
    }
    if (p1 >= p1) {
        fmt.Println(p1, p1)
        var v9 uint = 461
        fmt.Println(p0, p0)
        _ = v9
    } else {
        if (false || true) {
            var v10 int64 = (^(p0))
            var v11 uint32 = 25
            _ = v10
            _ = v11
        }
        v12 := []int{-554, -1}
        for fi13 := 0; fi13 < 5; fi13++ {
            for fi14 := 0; fi14 < 9; fi14++ {
                p0 = (p0 + p0)
                var v15 uint8 = (^(p1))
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
        _ = v12
    }
    v16 := GenS2{F0: 2, F1: 4294967295, F2: -329}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v16
    return 1
}

func main() {
    var v1 float32 = -2.25
    v2 := []int{}
    v3 := []string{(("" + "foo bar") + ("" + "hello")), ("a" + "foo bar"), "hello", "\n\t"}
    v4 := make(map[uint32]int)
    var v5 int16 = 42
    var v17 bool = false
    v17 = (-128 >= -1)
    var v18 uint16 = 2
    for fi19 := 0; fi19 < 12; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            if (!(-2.25 <= 3.14159)) {
                if ("\n\t" >= "foo bar") {
                    v1 -= ((v1 - 1e10) / v1)
                    fmt.Println(fi20, v18, v17)
                    switch (fi20 | (fi19 % ((243) | 1))) {
                    case 4294967295:
                        var v21 int64 = 399
                        _ = v21
                    }
                }
                if (65535 <= 988) {
                    v22 := GenS2{F0: (v18 + (v18 * v18)), F1: 0, F2: 0}
                    _ = len(v4)
                    var v23 int16 = 1000
                    _ = v22
                    _ = v23
                } else {
                    v1 += (v1 - v1)
                    var v24 bool = v17
                    var v25 string = "hello"
                    v5 += v5
                    _ = v24
                    _ = v25
                }
            }
            _ = fi20
        }
        _ = fi19
    }
    var v26 uint32 = 788
    var v27 int = 997
    v26 = (v26 % (((v26 / ((v26) | 1))) | 1))
    var v28 uint = 255
    if ((42 <= (^(v26))) && true) {
        switch v27 {
        case 9223372036854775807:
            var v29 float64 = (1e10 * (1e10 + (1.5 - 0.0)))
            _ = v29
        }
        if len(v3) > 0 {
            v3[0] = ""
        }
        for fi30 := 0; fi30 < 1; fi30++ {
            var v31 int16 = v5
            for fi32 := 0; fi32 < 7; fi32++ {
                var v33 uint8 = 128
                _ = fi32
                _ = v33
            }
            _ = fi30
            _ = v31
        }
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
