package main

import (
    "fmt"
)


func main() {
    v1 := []string{}
    v2 := make(map[string]string)
    var v3 uint32 = 4294967295
    var v4 string = ("a" + (("a" + "foo bar") + "\n\t"))
    v5 := make(map[uint32]string)
    var v6 uint8 = 255
    if (true || (!true)) {
        v4 = ((v4 + "a") + (((v4 + v4) + "foo bar") + "a"))
        fmt.Println(v3, v4)
        fmt.Println(v4, v6, v3)
    }
    var v7 bool = (true && true)
    v3 = 862
    var v8 uint16 = 706
    switch 4294967295 {
    case -539:
        var v9 uint8 = 0
        _ = v9
    }
    var v10 uint16 = 975
    var v11 uint8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
