package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
func genFn0(p0 int64, p1 int, p2 string) int {
    for fi4 := 0; fi4 < 1; fi4++ {
        p0 += (-(p0))
        _ = fi4
    }
    for fi5 := 0; fi5 < 18; fi5++ {
        switch 2147483647 {
        case 0:
            v6 := GenS1{F0: (fi5 | fi5), F1: 204}
            _ = v6
        }
        var v7 int8 = 127
        _ = fi5
        _ = v7
    }
    p1 = (p1 ^ (p1 - 623))
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{4294967295}
    var v3 int32 = -2147483648
    v3 = v3
    v1[48] = 232
    var v8 string = "a"
    v8 = "foo bar"
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            v1[65535] = 2
            _ = fi10
        }
        var v11 string = ""
        _ = fi9
        _ = v11
    }
    var v12 float32 = (1.5 + 0.1)
    v3 = 42
    v13 := GenS1{F0: -1, F1: 90}
    var v14 uint = 255
    if (!false) {
        fmt.Println(v14)
    } else {
        if true {
            var v15 int32 = 223
            _ = v15
        } else {
            var v16 int64 = 1
            _ = v16
        }
    }
    fmt.Println(v12, v12, v8)
    v14 += (v14 - v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
}
