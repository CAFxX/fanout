package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int32 = 2147483647
    var v2 uint32 = 4294967295
    var v3 int32 = (v1 * v1)
    switch 128 {
    case 177:
        v2 = v2
    }
    for fi4 := 0; fi4 < 13; fi4++ {
        fmt.Println(fi4, fi4)
        _ = fi4
    }
    v3 += (v3 * v3)
    var v5 uint = 4294967295
    var v6 int32 = (v1 << ((22) & 63))
    v7 := []int{32767}
    fmt.Println(v5, v3)
    fmt.Println(v5, v1)
    v2 = v2
    var v8 float32 = -2.25
    var v9 uint32 = (v2 * v2)
    if false {
        switch -2147483648 {
        case 5:
            var v10 uint = 877
            _ = v10
        }
        v6 = v1
        for fi11 := 0; fi11 < 12; fi11++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v12 := GenS1{F0: 1}
            _ = fi11
            _ = v12
        }
    }
    v1 = 0
    fmt.Println(v1, v6)
    for fi13 := 0; fi13 < 9; fi13++ {
        v8 = v8
        switch "hello" {
        case "foo bar":
            var v14 int64 = -9223372036854775808
            _ = v14
        default:
            if len(v7) > 0 {
                v7[0] = fi13
            }
        }
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
