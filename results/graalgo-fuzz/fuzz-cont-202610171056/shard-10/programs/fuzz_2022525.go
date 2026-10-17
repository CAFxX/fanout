package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint = 512
    v2 := make(map[int]uint32)
    var v3 bool = false
    v4 := make(map[uint32]bool)
    v5 := []uint32{0, 4294967295}
    v6 := []string{"", "a"}
    var v7 uint8 = 1
    v8 := GenS1{F0: 32767, F1: -64}
    fmt.Println(v1, v3, v7)
    var v9 int32 = -1
    v7 -= (v7 % ((255) | 1))
    v9 -= 1
    for fi10 := 0; fi10 < 20; fi10++ {
        switch "a" {
        case "a":
            for fi11 := 0; fi11 < 12; fi11++ {
                v12 := GenS1{F0: 111, F1: 98}
                if v3 {
                    v3 = (!((101 < 1) && (53 < 1)))
                } else {
                    v7 -= (v7 - 11)
                    v13 := GenS2{F0: 2}
                    _ = v13
                }
                v14 := GenS3{F0: (^(v1))}
                _ = fi11
                _ = v12
                _ = v14
            }
        case "\n\t":
            var v15 int32 = v9
            _ = v15
        }
        _ = fi10
    }
    switch "" {
    case "":
        v16 := GenS4{F0: 32767, F1: v1}
        _ = v16
    default:
        delete(v4, 482)
    }
    if true {
        v17 := GenS5{F0: v9}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        switch -710 {
        case -159:
            fmt.Println(v1, v9)
        case 2147483647:
            if len(v6) > 0 {
                v6[0] = ("a" + "hello")
            }
        default:
            _ = len(v6)
        }
        _ = v17
    }
    v18 := v4[65535]
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
