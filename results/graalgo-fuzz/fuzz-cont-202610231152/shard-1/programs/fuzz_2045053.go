package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 string }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 906
    var v3 int32 = -1
    v4 := make(map[int]string)
    var v5 uint = 1
    v6 := []uint32{2147483647, 917, 861, 2}
    fmt.Println(v3, v3)
    var v7 int32 = v3
    for fi8 := 0; fi8 < 20; fi8++ {
        if false {
            if (("foo bar" + "hello") == "hello") {
                v3 = v3
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v9 := GenS1{F0: 2147483648}
                _ = v9
            }
        }
        _ = fi8
    }
    v10 := GenS2{F0: 488, F1: ""}
    fmt.Println(v5, v3)
    var v11 bool = false
    v12 := GenS1{F0: 402}
    var v13 float64 = 1e10
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
