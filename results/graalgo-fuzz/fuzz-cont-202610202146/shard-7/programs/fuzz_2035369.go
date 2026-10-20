package main

import (
    "fmt"
)


func main() {
    var v1 int64 = 1
    var v2 int64 = -10
    var v3 uint8 = 65
    v4 := make(map[int]bool)
    var v5 int = 42
    var v6 int = (v5 >> ((42) & 63))
    v6 += 32768
    for fi7 := 0; fi7 < 4; fi7++ {
        if (2147483647 == 20) {
            var v8 bool = true
            var v9 uint8 = (v3 % ((187) | 1))
            _ = v8
            _ = v9
        } else {
            v1 += 128
            v2 -= 9223372036854775807
        }
        _ = fi7
    }
    var v10 int = (v6 + -583)
    fmt.Println(v1, v10)
    var v11 uint16 = 332
    fmt.Println(v1)
    var v12 int64 = (-(v2))
    v11 -= v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
