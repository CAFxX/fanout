package main

import (
    "fmt"
)


func main() {
    var v1 uint = 897
    var v2 float64 = ((1e10 / 1.0) - 0.1)
    var v3 string = "a"
    v4 := make(map[string]bool)
    delete(v4, "a")
    var v5 int32 = 174
    v6 := []int{-90}
    var v7 int = 128
    v4[(v3 + v3)] = false
    v2 = v2
    var v8 int16 = 32767
    var v9 int = v7
    v4[("\n\t" + "hello")] = (351 <= 65535)
    _ = len(v4)
    var v10 bool = (-9223372036854775808 > -949)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
