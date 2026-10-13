package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 string = ("a" + "hello")
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 int16 = 32767
    var v5 int8 = -22
    v6 := v3[4294967295]
    var v7 int32 = -1
    if false {
        var v8 uint16 = 65535
        _ = v8
    }
    v7 = v7
    v4 += v4
    for fi9 := 0; fi9 < 15; fi9++ {
        fmt.Println(v5, v1, v1)
        var v10 int32 = v7
        _ = fi9
        _ = v10
    }
    _ = len(v2)
    var v11 uint64 = 1
    if true {
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := GenS1{F0: 42, F1: -929}
            var v14 string = ("\n\t" + v1)
            if false {
                switch "foo bar" {
                case "hello":
                    _ = len(v2)
                case "a":
                    v5 = (v5 * (v5 + v5))
                case "foo bar":
                    v15 := GenS1{F0: 9223372036854775807, F1: (v4 / ((v4) | 1))}
                    _ = v15
                }
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        v16 := GenS1{F0: -9223372036854775808, F1: v4}
        _ = v16
    }
    v17 := GenS2{F0: -152}
    var v18 uint64 = (v11 / ((18446744073709551615) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
