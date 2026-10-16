package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"foo bar", "a"}
    var v3 uint32 = 2
    v4 := make(map[string]string)
    var v5 int32 = -377
    switch ("foo bar" + "a") {
    case "a":
        v4[("\n\t" + ("a" + "\n\t"))] = "\n\t"
    case "":
        v5 = 433
    default:
        var v6 int16 = -1
        _ = v6
    }
    v4["a"] = "a"
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(fi7, v3)
        _ = fi7
    }
    fmt.Println(v5, v3)
    v2 = append(v2, (("\n\t" + "a") + ((("\n\t" + ("" + "")) + "foo bar") + "foo bar")))
    fmt.Println(v5, v5)
    if true {
        fmt.Println(v5, v3)
    }
    var v8 int64 = -1
    _ = len(v2)
    v8 += (v8 / ((v8) | 1))
    for fi9 := 0; fi9 < 10; fi9++ {
        if (42 <= int8(fi9)) {
            fmt.Println(v3, v5)
        }
        switch "\n\t" {
        case "a":
            v5 = v5
        }
        var v10 int = fi9
        _ = fi9
        _ = v10
    }
    var v11 float64 = (0.1 * 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
