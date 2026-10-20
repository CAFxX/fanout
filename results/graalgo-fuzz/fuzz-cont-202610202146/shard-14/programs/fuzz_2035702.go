package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint32 }
func genFn0(p0 bool, p1 uint16, p2 int32) string {
    p0 = (-1 > 9223372036854775807)
    if p0 {
        fmt.Println(p0, p0, p1)
        fmt.Println(p2, p1)
        if p0 {
            var v6 uint64 = 2
            _ = v6
        }
    } else {
        if true {
            v7 := GenS1{F0: 2147483648}
            _ = v7
        } else {
            switch -2147483648 {
            case 42:
                var v8 float32 = -2.25
                _ = v8
            case -87:
                var v9 uint32 = 2
                _ = v9
            default:
                v10 := []int{}
                _ = v10
            }
            v11 := []int{}
            _ = v11
        }
        v12 := []int{-37, -2147483648, 128}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return (("foo bar" + "") + "\n\t")
}
func genFn1() int16 {
    var v13 float64 = 0.1
    fmt.Println(v13)
    for fi14 := 0; fi14 < 7; fi14++ {
        switch fi14 {
        case 42:
            var v15 int = 1
            _ = v15
        }
        _ = fi14
    }
    v16 := GenS2{F0: 103, F1: 2, F2: 65535}
    var v17 int8 = -128
    _ = v13
    _ = v16
    _ = v17
    return -32768
}

func main() {
    v1 := []string{""}
    v2 := make(map[uint32]int)
    var v3 float64 = -2.25
    v4 := make(map[int]bool)
    v5 := make(map[string]uint32)
    fmt.Println(v3)
    v18 := GenS2{F0: -128, F1: 127, F2: 2}
    for fi19 := 0; fi19 < 11; fi19++ {
        if true {
            var v20 int = (fi19 ^ fi19)
            _ = v20
        } else {
            var v21 uint64 = 605
            var v22 int = fi19
            var v23 int8 = -1
            _ = v21
            _ = v22
            _ = v23
        }
        if (!((48 > 42) && true)) {
            v3 = v3
            v3 = 1e10
        } else {
            v3 += v3
            switch (fi19 >> ((2) & 63)) {
            case 0:
                v3 -= v3
            default:
                for fi24 := 0; fi24 < 13; fi24++ {
                    fmt.Println(fi24, v3, fi19)
                    _ = fi24
                }
            }
        }
        var v25 float64 = (3.14159 + (((v3 / 3.14159) * 0.0) / v3))
        _ = fi19
        _ = v25
    }
    v3 += v3
    var v26 float64 = (float64(((0.0 + 0.0) / ((0.1 * 1e10) / 0.1))) + v3)
    v3 = (((float64(0.1) * 1e10) * v3) / 1e10)
    fmt.Println(v3, v26)
    var v27 float32 = 0.0
    v27 -= v27
    v28 := GenS1{F0: 2147483648}
    v3 = (float64(3.14159) / v26)
    if true {
        for fi29 := 0; fi29 < 11; fi29++ {
            if (820 > 831) {
                fmt.Println(v26)
            }
            fmt.Println(v3, v26)
            _ = fi29
        }
    } else {
        fmt.Println(v27)
        switch "" {
        case "\n\t":
            for fi30 := 0; fi30 < 14; fi30++ {
                fmt.Println(v27, fi30)
                _ = fi30
            }
        case "foo bar":
            for fi31 := 0; fi31 < 20; fi31++ {
                v27 = v27
                _ = fi31
            }
        case "a":
            fmt.Println(v27)
        }
        var v32 uint8 = 2
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
