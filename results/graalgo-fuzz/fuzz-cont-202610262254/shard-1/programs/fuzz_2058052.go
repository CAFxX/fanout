package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }

func main() {
    var v1 string = (("" + ("" + (("\n\t" + "a") + "a"))) + ((("a" + "a") + ("foo bar" + ("\n\t" + ("a" + (("hello" + ((("" + "foo bar") + (("foo bar" + "") + "foo bar")) + "hello")) + ("" + ("hello" + ("a" + "\n\t")))))))) + "a"))
    var v2 int8 = -66
    var v3 float64 = 1.5
    v4 := make(map[uint32]string)
    var v5 int = -9223372036854775808
    var v6 uint32 = 2
    v5 = -2147483648
    var v7 uint8 = 2
    v4[v6] = v1
    var v8 int16 = 0
    var v9 float32 = 0.1
    if (0 == 65535) {
        var v10 bool = true
        fmt.Println(v2, v2)
        var v11 int64 = 767
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 18446744073709551615, F1: 300}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
