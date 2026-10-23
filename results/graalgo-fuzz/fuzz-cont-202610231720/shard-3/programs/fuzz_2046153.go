package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
func genFn0(p0 string) uint {
    if false {
        switch -649 {
        case 218:
            if false {
                if (!false) {
                    v4 := GenS1{F0: 42, F1: 2}
                    switch p0 {
                    case "foo bar":
                        var v5 int64 = 1
                        _ = v5
                    }
                    v6 := GenS1{F0: 42, F1: 255}
                    _ = v4
                    _ = v6
                } else {
                    v7 := []int{9223372036854775807, 32767}
                    _ = v7
                }
            } else {
                for fi8 := 0; fi8 < 6; fi8++ {
                    fmt.Println(p0, fi8)
                    _ = fi8
                }
            }
        case -2147483648:
            for fi9 := 0; fi9 < 14; fi9++ {
                v10 := GenS1{F0: 2147483647, F1: 1}
                var v11 int16 = 32767
                _ = fi9
                _ = v10
                _ = v11
            }
        case 4294967295:
            v12 := []int{}
            _ = v12
        }
        var v13 float64 = 1.5
        _ = v13
    } else {
        if (0 < 181) {
            var v14 uint8 = 150
            var v15 float64 = (3.14159 + 1.5)
            switch p0 {
            case "":
                v16 := []int{32767, 1, 42}
                _ = v16
            default:
                fmt.Println(v14)
            }
            _ = v14
            _ = v15
        }
        v17 := []int{127, -2147483648, 0}
        _ = v17
    }
    fmt.Println(p0)
    switch "" {
    case "a":
        var v18 int16 = -1
        _ = v18
    default:
        fmt.Println(p0)
    }
    _ = p0
    return 42
}
func genFn1(p0 uint8, p1 int) int64 {
    var v19 float64 = -2.25
    p1 = (-(p1))
    _ = p0
    _ = p1
    _ = v19
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "\n\t"
    v3 := make(map[uint32]bool)
    var v20 uint = 432
    var v21 int = 42
    delete(v1, 0)
    v22 := []int{v21, v21, (v21 * (v21 >> ((v20) & 63))), v21}
    v23 := GenS1{F0: uint32(v20), F1: 0}
    v1[42] = "\n\t"
    for fi24 := 0; fi24 < 4; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            switch (v2 + (("\n\t" + "hello") + "foo bar")) {
            case "":
                v20 -= v20
            default:
                fmt.Println(fi25)
            }
            v26 := GenS1{F0: 847, F1: 2}
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    fmt.Println(v20, v2)
    v27 := GenS1{F0: 2147483648, F1: uint8(v20)}
    v28 := GenS1{F0: 4294967295, F1: 128}
    var v29 uint8 = uint8(v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
