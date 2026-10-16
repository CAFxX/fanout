package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    var v4 uint16 = 65535
    v3[-658] = 2147483648
    var v5 int32 = 2147483647
    fmt.Println(v5, v5)
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 float64 = 3.14159
        var v8 int = 4294967295
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v4, v5)
    if true {
        v5 -= v5
        var v9 string = (((("foo bar" + "foo bar") + "") + "foo bar") + "a")
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v5 = v5
        v1 = append(v1, 547)
        fmt.Println(v4, fi10)
        _ = fi10
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
