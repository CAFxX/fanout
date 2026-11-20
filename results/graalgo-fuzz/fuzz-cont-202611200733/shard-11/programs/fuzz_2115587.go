package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 int16 = 0
    v2 := []uint32{989, 836}
    var v3 int64 = 1
    v3 += (v3 ^ 0)
    var v4 uint16 = 42
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 818, F1: 985}
        var v7 float32 = 3.14159
        _ = fi5
        _ = v6
        _ = v7
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v8 int32 = -613
    v9 := GenS1{F0: 323, F1: v8}
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int8 = 39
        if (!false) {
            _ = len(v2)
            if (!false) {
                var v12 uint16 = (v4 + 1)
                v13 := GenS2{F0: 870}
                var v14 string = "\n\t"
                _ = v12
                _ = v13
                _ = v14
            } else {
                fmt.Println(v8)
            }
            var v15 string = "a"
            _ = v15
        } else {
            fmt.Println(v3, fi10)
            v4 = v4
        }
        _ = fi10
        _ = v11
    }
    v8 += (v8 & int32(v3))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
