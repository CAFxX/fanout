package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0(p0 int64, p1 bool) uint {
    switch "hello" {
    case "\n\t":
        v5 := []int{}
        _ = v5
    }
    var v6 int = -9223372036854775808
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 bool = (-175 >= 1000)
        for fi9 := 0; fi9 < 5; fi9++ {
            switch (("\n\t" + "a") + "") {
            case "a":
                v10 := GenS1{F0: 2}
                _ = v10
            case "hello":
                v11 := GenS1{F0: 89}
                _ = v11
            }
            var v12 uint32 = 65535
            _ = fi9
            _ = v12
        }
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 uint8 = 30
            v15 := []int{-2147483648, fi13, (v6 | 2147483647)}
            v16 := GenS1{F0: (v14 & v14)}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 42
}
func genFn1(p0 uint16) int16 {
    v17 := []int{-9223372036854775808, -35, 32768}
    var v18 bool = false
    if v18 {
        v18 = v18
        var v19 uint = 42
        _ = len(v17)
        _ = v19
    }
    var v20 int8 = -1
    _ = p0
    _ = v17
    _ = v18
    _ = v20
    return int16(v20)
}

func main() {
    var v1 float64 = float64(((-2.25 * 1.5) / (3.14159 + 0.0)))
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    v4 := []uint32{0, 4294967295}
    var v21 uint16 = 65535
    if true {
        v3[("\n\t" + "a")] = (v21 >= v21)
        switch 32767 {
        case 4294967295:
            v21 -= (^(v21))
        case -528:
            var v22 uint8 = 2
            _ = v22
        case -2147483648:
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 int8 = 1
                v25 := GenS1{F0: 0}
                _ = fi23
                _ = v24
                _ = v25
            }
        }
    }
    var v26 uint8 = 143
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v26, v1, v21)
    default:
        _ = len(v4)
    }
    var v27 int32 = -1
    var v28 int8 = 59
    fmt.Println(v28, v1)
    if true {
        fmt.Println(v26, v28)
        fmt.Println(v26, v26)
        fmt.Println(v1)
    } else {
        var v29 bool = true
        var v30 int32 = -740
        _ = len(v4)
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        var v32 int8 = v28
        _ = fi31
        _ = v32
    }
    switch "hello" {
    case "":
        for fi33 := 0; fi33 < 13; fi33++ {
            switch (fi33 - fi33) {
            case 42:
                v26 += v26
            case 2147483647:
                v34 := GenS1{F0: (v26 & v26)}
                _ = v34
            case -1:
                for fi35 := 0; fi35 < 18; fi35++ {
                    if (("a" != "a") || (v1 > 1.5)) {
                        var v36 uint32 = 812
                        fmt.Println(v26, v27)
                        _ = v36
                    }
                    var v37 uint16 = v21
                    _ = fi35
                    _ = v37
                }
            default:
                if (42 == -89) {
                    v1 += (v1 + 0.1)
                }
            }
            _ = fi33
        }
    case "foo bar":
        v2[106] = 4294967295
    }
    v38 := GenS2{F0: (v27 * v27), F1: 1}
    fmt.Println(v26, v21, v26)
    for fi39 := 0; fi39 < 9; fi39++ {
        v28 = v28
        var v40 uint32 = 2
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v4
}
