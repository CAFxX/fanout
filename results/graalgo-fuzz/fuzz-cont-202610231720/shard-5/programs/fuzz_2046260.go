package main

import (
    "fmt"
)


func main() {
    var v1 bool = true
    v2 := []string{"hello", ("a" + (("foo bar" + ("foo bar" + ("a" + ("" + ("foo bar" + ("foo bar" + "\n\t")))))) + "foo bar"))}
    v3 := make(map[string]bool)
    var v4 float64 = 0.1
    v5 := make(map[string]int)
    switch ("hello" + "\n\t") {
    case "a":
        for fi6 := 0; fi6 < 11; fi6++ {
            v4 = (1.5 * 0.0)
            var v7 uint16 = 553
            v1 = (((1 <= 1000) || v1) || (128 > 1))
            _ = fi6
            _ = v7
        }
    case "foo bar":
        var v8 float32 = (0.0 - 1.5)
        _ = v8
    case "":
        var v9 uint8 = 112
        _ = v9
    default:
        fmt.Println(v1)
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v4 += (1.5 + v4)
        v4 = v4
        fmt.Println(v1, fi10)
        _ = fi10
    }
    v1 = false
    v1 = (557 != -222)
    delete(v5, (("a" + "") + "hello"))
    var v11 bool = v1
    fmt.Println(v11, v11, v4)
    v4 = (float64((1e10 * 0.0)) - v4)
    if true {
        v1 = false
        var v12 int64 = -9223372036854775808
        v12 += v12
        _ = v12
    }
    v4 = float64(-2.25)
    var v13 uint8 = 128
    var v14 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
