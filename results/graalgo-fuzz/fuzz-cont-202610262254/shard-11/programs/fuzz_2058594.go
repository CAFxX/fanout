package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 string }
type GenS2 struct { F0 int32; F1 bool; F2 int32 }

func main() {
    var v1 uint8 = 42
    var v2 uint16 = 1
    var v3 uint8 = 127
    v4 := make(map[uint32]bool)
    var v5 int16 = -475
    _ = len(v4)
    v1 += v1
    v1 = (v3 + v1)
    switch -2147483648 {
    case 4294967295:
        switch "a" {
        case "foo bar":
            fmt.Println(v3, v1, v1)
        case "hello":
            v6 := []int{127, 4294967295}
            _ = v6
        }
    case -709:
        for fi7 := 0; fi7 < 11; fi7++ {
            if true {
                fmt.Println(v2, fi7)
                var v8 int16 = 42
                var v9 int32 = -2147483648
                _ = v8
                _ = v9
            } else {
                _ = len(v4)
                fmt.Println(v3)
                if false {
                    v2 = (v2 | (v2 << ((4294967295) & 63)))
                } else {
                    v10 := GenS1{F0: (fi7 + 166), F1: 18446744073709551615, F2: "\n\t"}
                    var v11 uint32 = 802
                    _ = v10
                    _ = v11
                }
            }
            _ = fi7
        }
    }
    v12 := []int{}
    switch "a" {
    case "":
        delete(v4, 2147483647)
    }
    var v13 int = 0
    if len(v12) > 0 {
        v12[0] = v13
    }
    v14 := GenS2{F0: 403, F1: (2147483648 < 63), F2: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
