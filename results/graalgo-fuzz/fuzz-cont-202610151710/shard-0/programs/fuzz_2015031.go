package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    var v4 bool = (!(116 >= 0))
    v4 = (0 != 1)
    var v5 float32 = 1e10
    for fi6 := 0; fi6 < 20; fi6++ {
        switch ("foo bar" + "") {
        case "foo bar":
            var v7 int = (fi6 * -560)
            _ = v7
        }
        var v8 int = fi6
        _ = fi6
        _ = v8
    }
    fmt.Println(v4)
    v5 = v5
    fmt.Println(v4)
    fmt.Println(v4)
    if (128 < 42) {
        if (v4 || true) {
            fmt.Println(v4, v4)
        }
        var v9 uint8 = 2
        var v10 int8 = 37
        _ = v10
        _ = v9
    }
    fmt.Println(v4, v4)
    var v11 string = "\n\t"
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(v11, v11)
        var v13 int16 = -375
        _ = fi12
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
