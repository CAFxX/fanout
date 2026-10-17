package main

import (
    "fmt"
)


func main() {
    var v1 int8 = -1
    var v2 uint64 = 107
    var v3 uint64 = v2
    v4 := []string{((("" + (("a" + "hello") + "\n\t")) + "hello") + "a"), "", "foo bar"}
    var v5 int16 = 32767
    var v6 string = "hello"
    _ = len(v4)
    v5 = v5
    for fi7 := 0; fi7 < 11; fi7++ {
        _ = len(v4)
        _ = len(v4)
        v1 = v1
        _ = fi7
    }
    var v8 int = 1
    v4 = append(v4, v6)
    for fi9 := 0; fi9 < 7; fi9++ {
        v2 = v3
        var v10 uint16 = 40
        _ = fi9
        _ = v10
    }
    switch "\n\t" {
    case "":
        if len(v4) > 0 {
            v4[0] = ("foo bar" + v6)
        }
    case "foo bar":
        var v11 uint16 = 580
        _ = v11
    case "a":
        v2 -= v2
    default:
        var v12 int16 = v5
        _ = v12
    }
    var v13 int16 = v5
    var v14 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
