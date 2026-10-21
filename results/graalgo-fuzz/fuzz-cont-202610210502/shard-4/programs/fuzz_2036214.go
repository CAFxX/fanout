package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 string = "hello"
    var v2 string = (v1 + v1)
    v3 := []uint32{}
    var v4 int16 = 32767
    var v5 int = 202
    var v6 int32 = -1
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v7 float64 = 1.5
    v5 -= v5
    for fi8 := 0; fi8 < 9; fi8++ {
        switch (v2 + (v2 + "\n\t")) {
        case "\n\t":
            fmt.Println(v6)
        case "foo bar":
            var v9 uint32 = 0
            _ = v9
        case "a":
            fmt.Println(v7, v5)
        }
        _ = fi8
    }
    var v10 uint64 = 361
    var v11 int16 = v4
    _ = len(v3)
    v4 = 445
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 bool = false
        var v14 uint64 = 0
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: v11}
    v16 := GenS1{F0: v11}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
