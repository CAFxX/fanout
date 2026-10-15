package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint; F1 int32; F2 uint32 }
func genFn0() int {
    var v5 uint32 = 0
    if false {
        var v6 int = 128
        var v7 uint64 = 2
        if false {
            v8 := []int{v6, (v6 << ((uint(v7)) & 63)), v6}
            var v9 uint32 = (v5 * v5)
            fmt.Println(v6, v6)
            _ = v8
            _ = v9
        } else {
            fmt.Println(v6, v6)
            v5 -= (v5 ^ (^(v5)))
        }
        _ = v6
        _ = v7
    } else {
        var v10 int16 = -32768
        var v11 uint64 = 18446744073709551615
        _ = v10
        _ = v11
    }
    var v12 bool = true
    v13 := GenS1{F0: (831 < 42), F1: 127}
    if (!((v12 && (false || (1 >= 42))) && v12)) {
        v5 = v5
        v14 := GenS2{F0: 2, F1: -1, F2: 2}
        _ = v14
    }
    _ = v12
    _ = v13
    _ = v5
    return 257
}

func main() {
    v1 := []string{("\n\t" + "hello")}
    var v2 int = -9223372036854775808
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    var v15 uint16 = 1
    v2 -= genFn0()
    var v16 bool = true
    _ = len(v4)
    v17 := GenS1{F0: v16, F1: 42}
    v16 = (1 < 647)
    v15 = 65535
    v3["foo bar"] = true
    for fi18 := 0; fi18 < 17; fi18++ {
        if v16 {
            v15 += 32767
            v2 = v2
        }
        if v16 {
            var v19 string = ("\n\t" + "\n\t")
            for fi20 := 0; fi20 < 14; fi20++ {
                switch -890 {
                case 740:
                    v21 := GenS1{F0: false, F1: 127}
                    _ = v21
                }
                _ = len(v4)
                _ = fi20
            }
            v15 = v15
            _ = v19
        } else {
            v15 = v15
            for fi22 := 0; fi22 < 2; fi22++ {
                fmt.Println(v15, fi18, v2)
                _ = fi22
            }
            v16 = (!false)
        }
        for fi23 := 0; fi23 < 2; fi23++ {
            var v24 uint64 = 4294967295
            _ = fi23
            _ = v24
        }
        _ = fi18
    }
    switch genFn0() {
    case 32767:
        switch (-(v2)) {
        case -2147483648:
            var v25 uint = 4294967295
            _ = v25
        case 2147483647:
            fmt.Println(v2, v2, v16)
        }
    case 535:
        var v26 int = (v2 | -9223372036854775808)
        _ = v26
    case -1:
        if false {
            var v27 float64 = 1.5
            fmt.Println(v2, v15)
            _ = v27
        }
    default:
        v16 = v16
    }
    v15 += v15
    switch "hello" {
    case "\n\t":
        if (286 <= 2) {
            v16 = (42 < 0)
            _ = len(v4)
            fmt.Println(v2)
        }
    }
    v28 := GenS1{F0: true, F1: 63}
    if (!true) {
        var v29 float64 = (float64(0.1) - 3.14159)
        _ = v29
    } else {
        v2 -= 909
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
