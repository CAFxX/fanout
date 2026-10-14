package main

import (
    "fmt"
)


func main() {
    var v1 bool = false
    var v2 int32 = 0
    v3 := make(map[string]int)
    var v4 int16 = -1
    var v5 uint8 = 255
    for fi6 := 0; fi6 < 3; fi6++ {
        _ = len(v3)
        _ = fi6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 string = "hello"
        _ = fi7
        _ = v8
    }
    v2 += 1
    var v9 string = "a"
    v5 = (v5 - v5)
    fmt.Println(v1, v1, v9)
    v1 = (!(v1 && v1))
    v10 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
