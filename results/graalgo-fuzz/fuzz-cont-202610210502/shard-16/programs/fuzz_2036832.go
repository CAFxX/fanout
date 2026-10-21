package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 370
    v2 := []string{((("foo bar" + "foo bar") + "foo bar") + "\n\t"), "", ("\n\t" + "\n\t")}
    var v3 string = "hello"
    v4 := make(map[int]string)
    v5 := make(map[uint32]string)
    var v6 int32 = 2147483647
    var v7 uint8 = 0
    var v8 uint64 = 0
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 bool = false
        _ = fi9
        _ = v10
    }
    v1 -= 2147483648
    var v11 uint16 = 1
    var v12 uint8 = (^(v7))
    v12 = v7
    fmt.Println(v11, v12, v8)
    var v13 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
