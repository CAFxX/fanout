package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 42
    v2 := make(map[int]bool)
    v3 := []string{"foo bar", "a", "hello", "a"}
    var v4 string = ("hello" + "a")
    fmt.Println(v1, v4)
    for fi5 := 0; fi5 < 16; fi5++ {
        fmt.Println(v4, fi5, fi5)
        _ = fi5
    }
    v1 = 1
    if true {
        var v6 uint32 = 477
        _ = v6
    }
    fmt.Println(v1, v1)
    fmt.Println(v1, v4)
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(v4)
        _ = len(v2)
        _ = fi7
    }
    if len(v3) > 0 {
        v3[0] = ""
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
