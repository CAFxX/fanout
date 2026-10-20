package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 825
    v2 := []string{}
    var v3 uint = 700
    var v4 float32 = float32(((0.0 - -2.25) * 0.0))
    v4 = v4
    if true {
        v1 -= v1
        var v5 int32 = -1
        var v6 int = 42
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int32 = 42
        _ = fi7
        _ = v8
    }
    switch 32768 {
    case 1:
        var v9 int16 = 1
        _ = v9
    case 518:
        v4 -= 0.1
    }
    var v10 uint32 = 946
    fmt.Println(v1, v4)
    var v11 int32 = -532
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
