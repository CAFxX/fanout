package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS2 struct { F0 int16; F1 int64 }

func main() {
    var v1 bool = false
    v2 := make(map[int]string)
    var v3 uint16 = 255
    v4 := []uint32{2, 42}
    switch "\n\t" {
    case "foo bar":
        v5 := GenS1{F0: 1, F1: 0, F2: -2147483648}
        _ = v5
    case "a":
        var v6 uint8 = 47
        _ = v6
    }
    _ = len(v4)
    for fi7 := 0; fi7 < 19; fi7++ {
        _ = len(v4)
        _ = fi7
    }
    var v8 uint = 255
    if false {
        if (2147483647 > 2147483648) {
            var v9 uint32 = 963
            for fi10 := 0; fi10 < 5; fi10++ {
                v8 = (v8 ^ v8)
                _ = len(v2)
                v11 := GenS1{F0: 4294967295, F1: 42, F2: -1}
                _ = fi10
                _ = v11
            }
            _ = v9
        }
    } else {
        delete(v2, 127)
    }
    var v12 uint8 = 0
    if (57 == (v12 & (v12 % (((^(v12))) | 1)))) {
        v13 := GenS2{F0: 32767, F1: 0}
        _ = v13
    } else {
        var v14 string = ""
        fmt.Println(v3, v3)
        _ = v14
    }
    v1 = (v8 <= 925)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
