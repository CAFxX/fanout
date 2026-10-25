package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int16 = 32767
    var v2 uint32 = 805
    var v3 uint = 65535
    v4 := make(map[int]bool)
    var v5 float64 = -2.25
    var v6 uint8 = 74
    var v7 uint16 = 615
    var v8 bool = (true || true)
    v9 := GenS1{F0: (^(v1))}
    if (true && ((!false) || (((v8 && ((v7 == 1) && (true && true))) || v8) && v8))) {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 bool = v8
            for fi12 := 0; fi12 < 8; fi12++ {
                _ = len(v4)
                var v13 uint64 = 419
                _ = fi12
                _ = v13
            }
            var v14 bool = false
            _ = fi10
            _ = v11
            _ = v14
        }
        if (472 != 42) {
            var v15 uint16 = (v7 >> ((4294967295) & 63))
            fmt.Println(v1, v6)
            _ = v15
        }
    } else {
        var v16 int16 = 0
        _ = v16
    }
    v2 = (^(v2))
    var v17 int64 = 42
    var v18 uint16 = (v7 % ((v7) | 1))
    v2 = (^(v2))
    v5 = 1.5
    var v19 uint8 = (v6 | uint8(v18))
    v19 -= 42
    var v20 uint = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
