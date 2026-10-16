package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]uint32)
    var v2 string = ("" + ("foo bar" + "foo bar"))
    v3 := []int{32767}
    fmt.Println(v2)
    for fi4 := 0; fi4 < 2; fi4++ {
        fmt.Println(v2, fi4)
        _ = fi4
    }
    var v5 uint64 = 557
    _ = len(v1)
    _ = len(v1)
    v6 := v1[v2]
    v5 = v5
    fmt.Println(v5, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
