package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -9223372036854775808
    var v3 int8 = 1
    switch "a" {
    case "a":
        var v4 bool = (127 > 0)
        _ = v4
    default:
        for fi5 := 0; fi5 < 2; fi5++ {
            v2 -= v2
            fmt.Println(v2, v3)
            _ = fi5
        }
    }
    var v6 string = ("\n\t" + "foo bar")
    v7 := []int{}
    v2 += v2
    var v8 uint8 = 1
    var v9 int32 = 1
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            delete(v1, (v6 + "\n\t"))
            v8 = (^(v8))
            _ = fi11
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi10
    }
    fmt.Println(v9)
    v9 += (v9 + (^(v9)))
    fmt.Println(v3, v2)
    v12 := GenS1{F0: 2147483647}
    if len(v7) > 0 {
        v7[0] = 4294967295
    }
    fmt.Println(v8, v9)
    v13 := GenS1{F0: 101}
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
