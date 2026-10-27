package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int }
type GenS2 struct { F0 int32; F1 bool }
func genFn0(p0 uint8) uint32 {
    if true {
        var v6 uint8 = (p0 * 0)
        _ = v6
    }
    var v7 uint = 4
    _ = p0
    _ = v7
    return uint32(p0)
}
func genFn1() uint16 {
    if true {
        v8 := GenS1{F0: 2, F1: false, F2: 1}
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(fi9)
            fmt.Println(fi9)
            for fi10 := 0; fi10 < 5; fi10++ {
                fmt.Println(fi9)
                fmt.Println(fi10, fi10)
                v11 := []int{9223372036854775807, fi10, fi9}
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        var v12 int64 = -9223372036854775808
        _ = v12
        _ = v8
    }
    var v13 int16 = -32768
    v14 := []int{-518, 9223372036854775807, 0, 1}
    var v15 int64 = 378
    v13 = v13
    _ = v13
    _ = v14
    _ = v15
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 int = 818
    v4 := make(map[string]bool)
    v5 := []string{}
    var v16 float32 = 3.14159
    var v17 uint8 = 28
    for fi18 := 0; fi18 < 15; fi18++ {
        if (v16 <= v16) {
            var v19 uint16 = genFn1()
            _ = v19
        }
        var v20 float32 = v16
        _ = fi18
        _ = v20
    }
    var v21 float64 = (1.5 + float64(v16))
    switch "\n\t" {
    case "foo bar":
        var v22 uint32 = 803
        _ = v22
    case "\n\t":
        var v23 int = v3
        _ = v23
    case "":
        v3 += v3
    default:
        v21 = 0.1
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        if ((-2147483648 < 1) && false) {
            if true {
                var v25 float64 = 0.1
                if true {
                    for fi26 := 0; fi26 < 8; fi26++ {
                        delete(v4, (("foo bar" + "") + "\n\t"))
                        _ = fi26
                    }
                    var v27 int64 = -9223372036854775808
                    fmt.Println(v27, v17)
                    _ = v27
                } else {
                    fmt.Println(v16, v3, v21)
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                }
                _ = v25
            }
        } else {
            v28 := GenS2{F0: 2147483647, F1: (2147483648 >= uint32(v17))}
            var v29 uint64 = 1
            _ = v28
            _ = v29
        }
        _ = fi24
    }
    _ = len(v2)
    _ = len(v2)
    var v30 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
