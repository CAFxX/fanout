package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 487
    v3 := make(map[int]uint32)
    var v4 float32 = -2.25
    var v5 int8 = 127
    var v6 string = "a"
    v7 := GenS1{F0: true, F1: -661}
    var v8 uint64 = 0
    var v9 uint64 = v8
    switch -1 {
    case -9223372036854775808:
        switch "hello" {
        case "\n\t":
            v8 -= v8
        case "a":
            if (4294967295 <= 664) {
                fmt.Println(v8)
                for fi10 := 0; fi10 < 2; fi10++ {
                    if ((!(!(!true))) || (4294967295 != 965)) {
                        v6 = (v6 + v6)
                        v5 = v5
                    }
                    v11 := v1[125]
                    _ = fi10
                    _ = v11
                }
            }
        }
    }
    fmt.Println(v6, v5)
    v12 := GenS2{F0: 151}
    v13 := []int{820}
    if len(v13) > 0 {
        v13[0] = 1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
