package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int16; F2 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 999
    v3 := []int{128, 0}
    v4 := []string{("" + ""), "a", "hello"}
    v5 := []string{"a", "", ""}
    var v6 uint64 = 42
    _ = len(v1)
    switch "foo bar" {
    case "\n\t":
        _ = len(v1)
    case "hello":
        if false {
            var v7 uint16 = 42
            fmt.Println(v6, v7, v2)
            var v8 uint16 = 646
            _ = v7
            _ = v8
        }
    default:
        _ = len(v4)
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(v2, fi9, v6)
        v2 = fi9
        for fi10 := 0; fi10 < 13; fi10++ {
            if true {
                var v11 float64 = -2.25
                _ = v11
            }
            var v12 float64 = 1.5
            _ = fi10
            _ = v12
        }
        _ = fi9
    }
    switch (v2 / ((v2) | 1)) {
    case 32767:
        switch -2147483648 {
        case 867:
            var v13 float64 = (-2.25 / float64(((float32((0.0 - 0.0)) + 1.5) / ((((0.1 / (((3.14159 + (3.14159 * (1e10 + (0.1 / 3.14159)))) + -2.25) / -2.25)) + 3.14159) / float32(1.5)) / 1.5))))
            _ = v13
        }
    case -861:
        if true {
            if ((-40 <= 0) || true) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v14 int = v2
                _ = v14
            }
            if false {
                v15 := v1[2147483648]
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 int = (fi16 / ((507) | 1))
                    v15 = (v15 && ((233 < 68) && v15))
                    _ = fi16
                    _ = v17
                }
                v15 = (!v15)
                _ = v15
            }
            var v18 int = v2
            _ = v18
        } else {
            var v19 uint64 = 4294967295
            delete(v1, 1)
            for fi20 := 0; fi20 < 18; fi20++ {
                v21 := GenS1{F0: 2147483647}
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    }
    v22 := GenS2{F0: 845, F1: 42, F2: -1}
    v23 := GenS2{F0: 1, F1: 1000, F2: -461}
    fmt.Println(v6)
    delete(v1, 867)
    v6 = 155
    var v24 float32 = (0.1 / ((1.5 / 0.1) - 1.5))
    v6 = (v6 ^ 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
