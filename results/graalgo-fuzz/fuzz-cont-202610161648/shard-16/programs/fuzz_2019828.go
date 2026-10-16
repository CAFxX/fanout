package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int }
func genFn0() int16 {
    var v4 uint32 = 471
    v4 = (v4 + (v4 | v4))
    v4 = (v4 >> ((2) & 63))
    for fi5 := 0; fi5 < 16; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            v4 = v4
            _ = fi6
        }
        _ = fi5
    }
    _ = v4
    return 42
}

func main() {
    v1 := []uint32{2147483647, 827, 2, 42}
    v2 := make(map[string]string)
    var v3 uint = 854
    v7 := GenS1{F0: 882, F1: -52}
    var v8 float64 = 0.0
    var v9 int8 = 127
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: -32768, F1: 0}
        delete(v2, "foo bar")
        _ = fi10
        _ = v11
    }
    switch "hello" {
    case "\n\t":
        delete(v2, "")
    }
    v9 -= (v9 + v9)
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS1{F0: 204, F1: 307}
        v8 = v8
        v3 -= 255
        _ = fi12
        _ = v13
    }
    _ = len(v1)
    v3 = (v3 / ((566) | 1))
    var v14 int16 = 722
    if true {
        v9 += (v9 | v9)
        _ = len(v1)
        for fi15 := 0; fi15 < 19; fi15++ {
            fmt.Println(v14)
            _ = fi15
        }
    }
    switch 542 {
    case 0:
        for fi16 := 0; fi16 < 8; fi16++ {
            v14 += (v14 << ((18446744073709551615) & 63))
            _ = fi16
        }
    }
    var v17 int16 = genFn0()
    v9 = v9
    switch 1 {
    case -844:
        v18 := GenS2{F0: 878}
        _ = v18
    case 645:
        v3 = 4294967295
    case -775:
        var v19 int32 = -997
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
