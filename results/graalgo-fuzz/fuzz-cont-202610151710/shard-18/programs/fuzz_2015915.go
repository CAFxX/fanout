package main

import (
    "fmt"
)


func main() {
    v1 := []string{(("foo bar" + ("a" + ("foo bar" + (("a" + ("foo bar" + ((("\n\t" + "hello") + "a") + "foo bar"))) + "\n\t")))) + ("a" + "\n\t")), ((("a" + (("a" + ("\n\t" + ("" + "hello"))) + "")) + "foo bar") + "foo bar"), "foo bar", (("hello" + "") + "\n\t")}
    var v2 uint8 = 102
    var v3 float32 = 1.5
    v4 := make(map[string]bool)
    v5 := make(map[uint32]string)
    var v6 uint32 = 4294967295
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    delete(v4, "foo bar")
    switch "hello" {
    case "\n\t":
        var v7 uint8 = 163
        _ = v7
    }
    switch -2147483648 {
    case -554:
        var v8 bool = true
        _ = v8
    case 32767:
        var v9 uint64 = 2
        _ = v9
    }
    if (3.14159 > ((1.5 / 1.5) - 3.14159)) {
        var v10 bool = (269 > 970)
        _ = v10
    } else {
        fmt.Println(v6, v2)
        v2 -= v2
    }
    v2 -= (v2 - 0)
    var v11 float64 = float64(float32(((1e10 + (1.5 * 0.0)) - 1.5)))
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
