package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int }
func genFn0(p0 bool, p1 uint16) int64 {
    var v7 string = "\n\t"
    v8 := GenS1{F0: 42, F1: 100, F2: 32768}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 string, p1 string) string {
    var v9 int = 0
    v10 := []int{v9}
    switch p0 {
    case "hello":
        var v11 int16 = -32768
        _ = v11
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return ""
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[int]uint32)
    v5 := []string{"hello"}
    v6 := make(map[string]string)
    v1 -= 0.1
    fmt.Println(v1)
    fmt.Println(v1)
    if false {
        switch 9223372036854775807 {
        case 32768:
            v1 -= (v1 - (v1 * v1))
        case -674:
            var v12 int = 757
            _ = v12
        case 0:
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 int16 = 1
                _ = fi13
                _ = v14
            }
        }
        fmt.Println(v1)
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint32 = 42
        var v17 bool = (!(0 == 0))
        for fi18 := 0; fi18 < 5; fi18++ {
            switch (fi18 >> ((42) & 63)) {
            case 0:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case -5:
                v17 = v17
            case 23:
                v19 := GenS1{F0: -1, F1: -54, F2: 127}
                _ = v19
            default:
                if (!(!v17)) {
                    v17 = v17
                }
            }
            switch "" {
            case "foo bar":
                var v20 float64 = v1
                _ = v20
            case "\n\t":
                fmt.Println(fi18, v16, v1)
            default:
                var v21 int = (fi18 % (((fi15 >> ((4294967295) & 63))) | 1))
                _ = v21
            }
            fmt.Println(v17)
            _ = fi18
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(v1)
    var v22 bool = true
    var v23 uint16 = 900
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
