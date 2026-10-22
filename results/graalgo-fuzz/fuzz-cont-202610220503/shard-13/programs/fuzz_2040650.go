package main

import (
    "fmt"
)

func genFn0(p0 int) string {
    p0 -= p0
    var v7 int8 = 42
    if (((3.14159 + 0.0) * 1.5) >= 1e10) {
        v7 += 1
        v7 -= v7
    }
    _ = p0
    _ = v7
    return "hello"
}

func main() {
    var v1 int16 = 847
    var v2 string = "hello"
    v3 := []string{(((((v2 + (("foo bar" + "") + "")) + v2) + "foo bar") + v2) + v2), "a", v2}
    var v4 uint64 = 0
    var v5 int16 = 1000
    var v6 int32 = 1
    var v8 uint8 = uint8(v4)
    _ = len(v3)
    v4 += (^(v4))
    v1 = -32768
    v4 += v4
    v3 = append(v3, "hello")
    fmt.Println(v5, v2, v8)
    var v9 uint = 255
    var v10 int = -809
    fmt.Println(v9, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
