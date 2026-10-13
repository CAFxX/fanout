package main

import (
    "fmt"
)


func main() {
    var v1 string = "foo bar"
    var v2 bool = (!(false || true))
    var v3 uint16 = 384
    var v4 uint64 = 18446744073709551615
    var v5 float64 = 1.5
    fmt.Println(v3, v1, v2)
    var v6 string = (v1 + v1)
    v2 = (v2 && v2)
    v7 := []int{537, 823}
    if true {
        v5 = v5
    } else {
        fmt.Println(v6, v5, v6)
    }
    switch 32768 {
    case 128:
        fmt.Println(v5)
    }
    var v8 bool = false
    var v9 int64 = -648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
