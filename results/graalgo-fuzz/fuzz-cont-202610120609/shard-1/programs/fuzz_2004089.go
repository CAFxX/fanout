package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint8 }

func main() {
    var v1 string = "a"
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    var v4 int16 = 1
    v5 := []int{-877}
    v6 := GenS1{F0: v1, F1: -15, F2: 42}
    switch "hello" {
    case "\n\t":
        v4 = (v4 << ((18446744073709551615) & 63))
    case "hello":
        v7 := GenS2{F0: v4, F1: 1}
        _ = v7
    }
    fmt.Println(v1, v1)
    v1 = v1
    v1 = (v1 + "a")
    var v8 uint8 = 1
    var v9 uint16 = 2
    v9 = (v9 & (v9 * (v9 + v9)))
    var v10 float32 = 0.1
    _ = len(v5)
    fmt.Println(v10, v8)
    v1 = v1
    v10 = v10
    if (9223372036854775807 != 0) {
        v8 += (^(v8))
    } else {
        v1 = v1
        for fi11 := 0; fi11 < 13; fi11++ {
            if false {
                v1 = v1
                var v12 float32 = 0.0
                _ = v12
            }
            if false {
                v13 := GenS1{F0: v1, F1: -52, F2: 953}
                switch v1 {
                case "\n\t":
                    var v14 uint = 0
                    _ = v14
                case "hello":
                    var v15 bool = false
                    _ = v15
                case "foo bar":
                    v16 := GenS2{F0: v4, F1: v8}
                    _ = v16
                default:
                    var v17 uint16 = v9
                    _ = v17
                }
                if false {
                    if false {
                        var v18 int16 = v4
                        var v19 int32 = -1
                        _ = v18
                        _ = v19
                    } else {
                        v1 = v1
                    }
                } else {
                    fmt.Println(v8, v9, v8)
                    switch (fi11 + fi11) {
                    case 32767:
                        fmt.Println(v9, v8)
                    }
                    v1 = v1
                }
                _ = v13
            } else {
                if true {
                    v1 = v1
                    for fi20 := 0; fi20 < 12; fi20++ {
                        fmt.Println(fi20, fi20, fi11)
                        _ = fi20
                    }
                    var v21 int8 = 100
                    _ = v21
                }
            }
            _ = fi11
        }
        fmt.Println(v9, v4, v9)
    }
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
