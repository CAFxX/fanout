package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int64 = -796
    var v2 uint64 = 42
    var v3 int = 92
    var v4 int8 = 42
    v5 := make(map[string]bool)
    var v6 int64 = v1
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 7; fi8++ {
        v1 += v6
        var v9 int16 = -32768
        _ = fi8
        _ = v9
    }
    v2 = (v2 >> ((0) & 63))
    fmt.Println(v1, v6)
    v10 := GenS1{F0: 42}
    v3 -= 42
    var v11 uint64 = v2
    switch (v3 >> ((718) & 63)) {
    case 0:
        v2 = (v11 << ((113) & 63))
    case 9223372036854775807:
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 int16 = 38
            _ = fi12
            _ = v13
        }
    case 4294967295:
        v14 := v5["hello"]
        _ = v14
    default:
        if false {
            var v15 uint16 = 0
            var v16 float32 = (1.5 - 3.14159)
            _ = v15
            _ = v16
        } else {
            if true {
                v17 := []int{v3, (-(v3))}
                _ = v17
            } else {
                var v18 uint16 = 65535
                _ = v18
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
