package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 255
    var v3 uint8 = 42
    v4 := []string{"\n\t", ("hello" + ("foo bar" + (("\n\t" + (("foo bar" + "a") + "\n\t")) + ("foo bar" + (("foo bar" + "\n\t") + "hello"))))), ("a" + ("hello" + "foo bar"))}
    var v5 uint = 18446744073709551615
    var v6 float32 = ((1e10 / 1e10) + -2.25)
    var v7 float64 = 0.1
    v5 += uint(v3)
    var v8 int8 = 57
    var v9 uint8 = v3
    fmt.Println(v6, v7)
    switch -949 {
    case 2147483647:
        var v10 uint8 = v9
        _ = v10
    case 42:
        v1[233] = false
    }
    var v11 int64 = 9223372036854775807
    v11 -= v11
    v2 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
