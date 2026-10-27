package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int; F2 uint16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    var v3 int16 = 324
    var v4 bool = (1.5 == (1e10 / -2.25))
    for fi5 := 0; fi5 < 7; fi5++ {
        v4 = v4
        var v6 string = "a"
        _ = fi5
        _ = v6
    }
    switch ("hello" + "foo bar") {
    case "\n\t":
        v4 = v4
    case "a":
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                v2[705] = ("foo bar" + (("hello" + ("foo bar" + "\n\t")) + "a"))
                _ = fi8
            }
            _ = fi7
        }
    }
    v3 += 1000
    v3 = (v3 | v3)
    var v9 bool = false
    if v4 {
        for fi10 := 0; fi10 < 15; fi10++ {
            _ = len(v1)
            _ = fi10
        }
    }
    var v11 uint = 833
    fmt.Println(v11, v3, v9)
    v3 = (-(v3))
    fmt.Println(v11, v3, v3)
    var v12 uint = v11
    v3 -= (v3 + (v3 * (v3 * 1000)))
    v13 := GenS1{F0: 226}
    v2[65535] = ("" + "")
    fmt.Println(v4, v3)
    if (v12 >= v12) {
        v14 := GenS1{F0: 42}
        var v15 uint16 = 887
        v3 = 1
        _ = v14
        _ = v15
    } else {
        v12 += 18446744073709551615
        fmt.Println(v11, v12, v3)
    }
    if (-2147483648 < 0) {
        switch "a" {
        case "\n\t":
            v16 := GenS1{F0: 0}
            _ = v16
        case "":
            v17 := GenS1{F0: 4294967295}
            _ = v17
        case "foo bar":
            if v4 {
                for fi18 := 0; fi18 < 5; fi18++ {
                    for fi19 := 0; fi19 < 5; fi19++ {
                        fmt.Println(v4, fi18)
                        var v20 int32 = 2147483647
                        _ = fi19
                        _ = v20
                    }
                    v21 := GenS2{F0: (v3 & 32767), F1: 2147483647, F2: 411}
                    var v22 uint = (v12 & (^(v12)))
                    _ = fi18
                    _ = v21
                    _ = v22
                }
                v11 = v12
                var v23 bool = true
                _ = v23
            }
        default:
            v24 := []int{-458}
            _ = v24
        }
    }
    v25 := []int{32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v9
}
