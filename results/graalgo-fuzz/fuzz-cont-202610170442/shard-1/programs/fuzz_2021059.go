package main

import (
    "fmt"
)


func main() {
    var v1 int64 = -1
    var v2 bool = (!false)
    v3 := []string{}
    var v4 int8 = 1
    var v5 uint = 2
    fmt.Println(v1, v1)
    var v6 int16 = -32768
    var v7 uint16 = 65535
    if true {
        if ((((("hello" + "") + (((("\n\t" + "a") + "foo bar") + ("foo bar" + "")) + ("a" + "a"))) + "a") + ("a" + "a")) < "hello") {
            v1 = v1
            v5 -= (v5 - (v5 & (v5 ^ v5)))
        }
        var v8 uint32 = 2
        _ = v8
    }
    var v9 float64 = float64(1.5)
    var v10 uint64 = 1
    v9 -= float64(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
