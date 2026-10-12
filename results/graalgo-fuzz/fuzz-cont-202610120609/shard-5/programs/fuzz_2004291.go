package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
func genFn0(p0 int8, p1 int32) uint16 {
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: (("" + "\n\t") + "\n\t"), F1: 0}
        _ = fi6
        _ = v7
    }
    var v8 uint64 = 18446744073709551615
    var v9 float64 = (0.1 - 3.14159)
    switch "\n\t" {
    case "a":
        v10 := []int{-842, 32768, -9223372036854775808, 32767}
        _ = v10
    case "\n\t":
        var v11 uint64 = (v8 + (v8 % ((982) | 1)))
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -919
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    var v5 uint16 = 931
    if false {
        v2 -= v2
        var v12 int16 = (v2 & (v2 * (v2 & v2)))
        var v13 uint = 18446744073709551615
        _ = v12
        _ = v13
    } else {
        switch "" {
        case "foo bar":
            v2 -= 563
        case "a":
            v14 := GenS1{F0: "", F1: (v2 - (v2 >> ((425) & 63)))}
            _ = v14
        case "\n\t":
            v2 = -32768
        }
        var v15 uint32 = 969
        _ = v15
    }
    v5 = v5
    switch "a" {
    case "foo bar":
        if (128 != -9223372036854775808) {
            v16 := []int{}
            var v17 int16 = v2
            var v18 string = ((("" + "") + "foo bar") + "a")
            _ = v16
            _ = v17
            _ = v18
        }
    }
    if (0.0 != -2.25) {
        v2 = (v2 & -339)
    } else {
        fmt.Println(v5, v2)
        var v19 string = "hello"
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := []int{fi20, fi20}
        v22 := GenS1{F0: ("foo bar" + "hello"), F1: 1}
        v5 -= genFn0(-91, -2147483648)
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            _ = len(v4)
        }
        _ = fi23
    }
    v24 := []int{-170, 1}
    switch 42 {
    case 127:
        if false {
            if (v5 > v5) {
                if (1 != 42) {
                    fmt.Println(v5, v2)
                    v5 -= (v5 | v5)
                    var v25 bool = (!true)
                    _ = v25
                }
                v26 := GenS1{F0: "\n\t", F1: (v2 >> ((2) & 63))}
                v27 := GenS1{F0: "hello", F1: v2}
                _ = v26
                _ = v27
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
