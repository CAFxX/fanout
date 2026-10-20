package main

import (
    "fmt"
)


func main() {
    var v1 int = 860
    v2 := make(map[uint32]string)
    v3 := make(map[string]bool)
    var v4 int32 = 42
    fmt.Println(v1)
    var v5 int16 = 32767
    v2[32] = "\n\t"
    v4 = v4
    v4 -= v4
    var v6 uint = 861
    var v7 uint8 = 255
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int16 = v5
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        switch v1 {
        case 1:
            var v11 int16 = (-(v5))
            _ = v11
        case -865:
            v5 -= (^(v5))
        case 32767:
            fmt.Println(fi10, v6, v6)
        }
        _ = fi10
    }
    v5 -= (v5 | (v5 - v5))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
