package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{166, 4294967295, 2147483648}
    var v2 uint32 = 2
    v3 := make(map[string]int)
    var v4 int64 = 163
    var v5 string = "hello"
    v2 += 0
    delete(v3, v5)
    if (false && false) {
        var v6 bool = (("\n\t" == ((("foo bar" + "hello") + (v5 + ((v5 + "\n\t") + (v5 + v5)))) + "\n\t")) || (231 < 0))
        _ = v6
    }
    switch 4294967295 {
    case 0:
        fmt.Println(v2, v4, v5)
    default:
        v5 = v5
    }
    var v7 uint = 42
    if false {
        var v8 int16 = 42
        delete(v3, v5)
        _ = v8
    }
    var v9 int8 = 100
    v2 = 2
    var v10 int32 = 1
    v11 := v3["foo bar"]
    v7 += v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
