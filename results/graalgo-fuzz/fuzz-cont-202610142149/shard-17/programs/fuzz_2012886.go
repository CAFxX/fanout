package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 940
    v3 := make(map[uint32]string)
    delete(v3, 0)
    for fi4 := 0; fi4 < 12; fi4++ {
        v2 += (v2 - 42)
        if (2 >= 255) {
            v5 := GenS1{F0: (!false), F1: 180, F2: 128}
            _ = v5
        } else {
            v3[4294967295] = (("hello" + ("\n\t" + "hello")) + "foo bar")
            var v6 int32 = v2
            fmt.Println(v6)
            _ = v6
        }
        _ = fi4
    }
    fmt.Println(v2)
    if (9223372036854775807 <= 9223372036854775807) {
        if true {
            if false {
                v2 = (^(v2))
                v7 := []int{-764, 643}
                for fi8 := 0; fi8 < 15; fi8++ {
                    v2 = -485
                    _ = fi8
                }
                _ = v7
            }
        }
    } else {
        v9 := GenS2{F0: 1, F1: 590}
        _ = v9
    }
    fmt.Println(v2)
    v10 := []int{2147483647, 4294967295, -2147483648, 128}
    var v11 int8 = -34
    v11 += v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
