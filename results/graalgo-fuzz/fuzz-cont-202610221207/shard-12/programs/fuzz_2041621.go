package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 157
    v2 := []string{"a", "hello"}
    var v3 int64 = -416
    var v4 uint16 = 240
    var v5 uint16 = 65535
    _ = len(v2)
    switch 127 {
    case 1:
        v2 = append(v2, "")
    case 9223372036854775807:
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 int = 496
            _ = fi6
            _ = v7
        }
    case 4294967295:
        fmt.Println(v5, v1, v1)
    default:
        v5 = 65535
    }
    _ = len(v2)
    v2 = append(v2, "hello")
    fmt.Println(v5, v1, v3)
    if true {
        v5 -= v5
    } else {
        var v8 int8 = 100
        fmt.Println(v3, v1, v8)
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        v4 += 32767
        var v10 int64 = -1
        _ = fi9
        _ = v10
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
