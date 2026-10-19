package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }

func main() {
    var v1 int64 = 0
    var v2 int32 = 514
    v3 := make(map[uint32]int)
    var v4 int = 4294967295
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            v1 = (v1 | (v1 & -1))
            var v7 float32 = 0.0
            v2 += v2
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    var v8 int = v4
    v9 := GenS1{F0: v4, F1: 255}
    if (!(!(true && (149 >= 273)))) {
        v10 := GenS1{F0: (^(v8)), F1: 42}
        _ = v10
    }
    var v11 bool = (false && true)
    delete(v3, 65535)
    v1 = (v1 - v1)
    var v12 string = "\n\t"
    v13 := v3[1]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
