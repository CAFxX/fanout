package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]int)
    v2 := []string{("\n\t" + "foo bar")}
    v3 := make(map[int]uint32)
    var v4 uint = 221
    var v5 int = -565
    v5 = -1
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int = fi6
        if len(v2) > 0 {
            v2[0] = ""
        }
        var v8 string = "\n\t"
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int64 = 0
    _ = len(v2)
    v9 += v9
    var v10 int64 = v9
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
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
