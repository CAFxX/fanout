package main

import (
    "fmt"
)


func main() {
    var v1 bool = false
    var v2 int16 = 107
    v3 := make(map[uint32]string)
    v4 := make(map[string]bool)
    v5 := []uint32{}
    v6 := make(map[int]bool)
    for fi7 := 0; fi7 < 6; fi7++ {
        _ = len(v5)
        _ = fi7
    }
    fmt.Println(v1)
    if v1 {
        for fi8 := 0; fi8 < 1; fi8++ {
            if (false && (42 == 20)) {
                v6[fi8] = v1
                var v9 uint64 = 0
                v4[(("foo bar" + "hello") + "a")] = (((fi8 * (fi8 - (fi8 & (^(fi8))))) > (fi8 + int(v9))) && ((!((((("foo bar" + "a") + "") + "hello") == "\n\t") || (-1 >= -1))) && v1))
                _ = v9
            }
            v1 = (4294967295 == 4294967295)
            _ = fi8
        }
    }
    v2 = (v2 << ((18446744073709551615) & 63))
    v1 = v1
    switch "foo bar" {
    case "foo bar":
        for fi10 := 0; fi10 < 2; fi10++ {
            v1 = true
            v2 = (v2 << ((4294967295) & 63))
            _ = fi10
        }
    default:
        v11 := v4["a"]
        _ = v11
    }
    v2 = (-(v2))
    v1 = (v1 || true)
    var v12 uint32 = 932
    var v13 float32 = (1e10 * 3.14159)
    fmt.Println(v2, v1)
    v1 = (v12 <= (^(v12)))
    fmt.Println(v12, v2)
    var v14 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
