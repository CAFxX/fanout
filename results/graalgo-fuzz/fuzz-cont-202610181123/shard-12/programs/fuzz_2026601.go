package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }

func main() {
    var v1 uint32 = 99
    v2 := make(map[string]string)
    v3 := []int{-990}
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]string)
    fmt.Println(v1)
    if len(v3) > 0 {
        v3[0] = 763
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        v1 = (^(v1))
        _ = fi6
    }
    v1 = v1
    v1 = (v1 ^ 0)
    var v7 int64 = -772
    for fi8 := 0; fi8 < 4; fi8++ {
        if ((true && false) && false) {
            v9 := v4[645]
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 uint32 = 2147483647
                v12 := GenS1{F0: -24, F1: "foo bar"}
                var v13 bool = false
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            _ = v9
        }
        switch "\n\t" {
        case "foo bar":
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 uint64 = 0
                _ = fi14
                _ = v15
            }
        case "\n\t":
            for fi16 := 0; fi16 < 11; fi16++ {
                fmt.Println(fi16, fi8, v1)
                var v17 uint16 = 423
                if false {
                    var v18 uint32 = 2
                    fmt.Println(v18, v18)
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
        }
        _ = fi8
    }
    v7 = (-(v7))
    v4[65535] = 2
    switch 157 {
    case 32767:
        v1 -= (v1 >> ((18446744073709551615) & 63))
    }
    var v19 int64 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
