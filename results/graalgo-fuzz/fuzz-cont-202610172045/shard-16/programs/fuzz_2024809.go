package main

import (
    "fmt"
)


func main() {
    var v1 int64 = -1
    v2 := []int{-2147483648, 32767, 641, -300}
    var v3 int = 32768
    v4 := make(map[uint32]bool)
    v3 = 0
    v3 = v3
    var v5 int32 = 0
    v5 = 304
    switch (v3 | 0) {
    case -1:
        var v6 int16 = 1000
        _ = v6
    case 9223372036854775807:
        var v7 uint16 = 0
        _ = v7
    }
    v1 -= v1
    var v8 uint16 = 32767
    v5 = (-(v5))
    switch (("a" + (("\n\t" + "foo bar") + "\n\t")) + "foo bar") {
    case "":
        var v9 float64 = float64(0.0)
        _ = v9
    }
    v5 = v5
    v8 = v8
    var v10 uint8 = 255
    var v11 uint32 = 2147483647
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 uint = 82
        if false {
            fmt.Println(v10, v13)
            v4[v11] = true
        }
        switch fi12 {
        case 4294967295:
            v5 = (v5 << ((1) & 63))
        case 0:
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 uint16 = v8
                v11 += (v11 - 4294967295)
                _ = fi14
                _ = v15
            }
        default:
            fmt.Println(v5)
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v10, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
