package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16, p1 bool, p2 int32) int {
    var v6 int = 128
    v7 := GenS1{F0: 42, F1: (^(v6))}
    switch ("" + "foo bar") {
    case "a":
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS2{F0: int16(p2)}
            _ = fi8
            _ = v9
        }
    }
    p2 = p2
    for fi10 := 0; fi10 < 3; fi10++ {
        v6 += (fi10 | fi10)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 int8 = 0
    v2 := []uint32{0, 2147483648}
    var v3 int = 496
    v4 := []uint32{15, 2, 2147483648}
    var v5 int16 = 0
    if false {
        var v11 bool = false
        _ = v11
    }
    var v12 int32 = -2147483648
    v13 := GenS2{F0: -325}
    var v14 float32 = 3.14159
    fmt.Println(v5, v12)
    fmt.Println(v3)
    if false {
        switch (v3 - v3) {
        case 953:
            var v15 uint64 = 4294967295
            _ = v15
        }
        v12 = -336
        var v16 uint16 = 42
        _ = v16
    } else {
        if false {
            v14 = (v14 + 0.0)
        }
    }
    var v17 int64 = -1
    switch "" {
    case "\n\t":
        var v18 int64 = 42
        _ = v18
    }
    v3 = (v3 >> ((688) & 63))
    fmt.Println(v17, v3, v3)
    var v19 float64 = -2.25
    v17 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
