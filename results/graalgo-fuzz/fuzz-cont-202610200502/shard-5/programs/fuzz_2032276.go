package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint8 }
func genFn0(p0 int8) int16 {
    if true {
        var v4 float64 = 1.5
        var v5 int = 4294967295
        _ = v4
        _ = v5
    }
    v6 := GenS1{F0: 65535, F1: 149, F2: 2}
    _ = p0
    _ = v6
    return -500
}
func genFn1(p0 string) uint32 {
    fmt.Println(p0)
    switch p0 {
    case "a":
        fmt.Println(p0)
    case "foo bar":
        var v7 int32 = -205
        _ = v7
    default:
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := []int{(fi8 % (((fi8 - 4294967295)) | 1))}
            var v10 bool = true
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    p0 = ("foo bar" + "\n\t")
    p0 = (p0 + p0)
    fmt.Println(p0)
    _ = p0
    return 648
}

func main() {
    v1 := []int{2147483647}
    v2 := make(map[string]bool)
    var v3 int16 = -416
    fmt.Println(v3)
    v3 += v3
    v3 -= (v3 - v3)
    fmt.Println(v3)
    v3 = genFn0(-12)
    var v11 int32 = -1
    var v12 uint = 42
    var v13 int32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
