package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]uint32)
    var v4 float64 = 0.0
    v5 := []string{(((("hello" + "foo bar") + "a") + "a") + ("a" + "")), (("hello" + "a") + ("a" + ("a" + "a"))), ((("hello" + (("hello" + ("a" + ("\n\t" + "hello"))) + ("a" + ("a" + "hello")))) + (("a" + "\n\t") + "")) + "hello")}
    v6 := []uint32{2}
    v4 -= (v4 + v4)
    v2 = 9223372036854775807
    var v7 float32 = ((3.14159 - 3.14159) * -2.25)
    v4 = v4
    v4 = (float64(0.1) / 1e10)
    var v8 uint = 1
    v4 += v4
    v8 += 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
