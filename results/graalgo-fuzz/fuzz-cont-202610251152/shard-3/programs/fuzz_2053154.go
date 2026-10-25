package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int }
type GenS2 struct { F0 uint }

func main() {
    var v1 int8 = 0
    v2 := []int{177}
    var v3 int16 = int16(v1)
    var v4 uint64 = 0
    v5 := []string{"foo bar", "hello", ("a" + "a")}
    var v6 uint32 = 36
    var v7 uint64 = v4
    var v8 bool = (true && ((v7 & v4) != v4))
    v3 -= (v3 >> ((255) & 63))
    if false {
        var v9 int64 = 9223372036854775807
        for fi10 := 0; fi10 < 17; fi10++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v11 := GenS1{F0: v6, F1: v3, F2: (-(fi10))}
            var v12 uint8 = 2
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    if v8 {
        var v13 bool = v8
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = len(v5)
        _ = v13
    } else {
        v14 := GenS2{F0: 0}
        fmt.Println(v1, v8, v8)
        fmt.Println(v3)
        _ = v14
    }
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    var v15 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
