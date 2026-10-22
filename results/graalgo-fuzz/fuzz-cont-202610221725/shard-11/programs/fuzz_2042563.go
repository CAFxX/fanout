package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
func genFn0() int64 {
    v4 := GenS1{F0: -710, F1: 2147483647}
    switch ("foo bar" + "hello") {
    case "hello":
        v5 := []int{}
        _ = v5
    default:
        v6 := []int{}
        _ = v6
    }
    var v7 uint8 = 2
    if false {
        v7 -= v7
        v8 := []int{9223372036854775807}
        _ = v8
    } else {
        fmt.Println(v7)
        if true {
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 int16 = -1
                fmt.Println(v10)
                _ = fi9
                _ = v10
            }
        }
    }
    _ = v4
    _ = v7
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    var v3 int64 = 1
    var v11 uint8 = 0
    v3 += v3
    v3 += genFn0()
    fmt.Println(v3, v11)
    var v12 uint32 = 42
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(fi13, v3)
        _ = fi13
    }
    var v14 int64 = 0
    fmt.Println(v12)
    fmt.Println(v14, v14)
    switch 32767 {
    case -9223372036854775808:
        var v15 int = -1
        _ = v15
    case 734:
        var v16 float64 = 1.5
        _ = v16
    case 32767:
        _ = len(v1)
    default:
        v12 += v12
    }
    v17 := []int{-985}
    if true {
        for fi18 := 0; fi18 < 15; fi18++ {
            v3 += genFn0()
            _ = fi18
        }
        v19 := GenS1{F0: 42, F1: 42}
        if ((1e10 * ((0.1 * 1e10) + -2.25)) >= 1.5) {
            var v20 int32 = 42
            for fi21 := 0; fi21 < 13; fi21++ {
                var v22 uint8 = 97
                switch "hello" {
                case "\n\t":
                    switch (("\n\t" + "foo bar") + "\n\t") {
                    case "a":
                        var v23 int64 = (v14 * genFn0())
                        _ = v23
                    case "hello":
                        v3 = (v3 * v14)
                    }
                }
                fmt.Println(v11)
                _ = fi21
                _ = v22
            }
            v3 = genFn0()
            _ = v20
        }
        _ = v19
    }
    fmt.Println(v11)
    fmt.Println(v12, v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v17
    _ = v2
    _ = v3
}
