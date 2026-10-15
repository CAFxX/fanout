package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0(p0 int, p1 int16) int16 {
    if false {
        switch "\n\t" {
        case "foo bar":
            p1 += (p1 & -1)
        }
    } else {
        var v5 uint32 = 4294967295
        v6 := GenS1{F0: -184}
        v5 += 280
        _ = v5
        _ = v6
    }
    var v7 float32 = 0.0
    _ = p0
    _ = p1
    _ = v7
    return p1
}

func main() {
    var v1 string = "\n\t"
    v2 := []string{v1, v1}
    v3 := make(map[string]uint32)
    var v4 int16 = 288
    v8 := GenS2{F0: 42, F1: -1}
    v1 = v1
    v9 := GenS1{F0: 767}
    for fi10 := 0; fi10 < 1; fi10++ {
        v4 = (^(v4))
        v1 = v1
        _ = fi10
    }
    delete(v3, "\n\t")
    v11 := GenS3{F0: 18446744073709551615, F1: 42}
    _ = len(v3)
    switch 32768 {
    case -337:
        v4 -= v4
    }
    var v12 uint = 604
    var v13 uint16 = 2
    var v14 uint32 = uint32(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
