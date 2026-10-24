package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint32 }

func main() {
    var v1 int64 = 1
    var v2 float32 = ((float32((3.14159 - 0.0)) - ((1.5 * 0.1) + 0.0)) - (3.14159 + float32((-2.25 + 0.0))))
    var v3 bool = ("foo bar" == "\n\t")
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 bool = true
        _ = fi4
        _ = v5
    }
    v6 := []int{-2147483648}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 bool = true
        var v9 uint = 65535
        _ = fi7
        _ = v8
        _ = v9
    }
    v3 = (!v3)
    var v10 float32 = (1.5 * 0.0)
    var v11 uint16 = 0
    v2 = v2
    if v3 {
        var v12 int64 = -235
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        fmt.Println(v2, fi13)
        var v14 float32 = float32((((3.14159 + float64(v10)) * 3.14159) * 1.5))
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 string = (("a" + ("hello" + "\n\t")) + "a")
            v17 := GenS1{F0: 0, F1: fi13, F2: 42}
            var v18 int64 = (v1 >> ((255) & 63))
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi13
        _ = v14
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
}
