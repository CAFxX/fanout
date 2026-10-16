package main

import (
    "fmt"
)


func main() {
    v1 := []string{("a" + "a"), "hello", "a"}
    var v2 uint16 = 961
    var v3 int64 = 9223372036854775807
    var v4 int8 = -93
    var v5 uint16 = (v2 >> ((255) & 63))
    var v6 uint16 = v2
    var v7 float32 = 1e10
    if len(v1) > 0 {
        v1[0] = ((("a" + ("" + ("" + (("hello" + "foo bar") + ("hello" + ("hello" + "")))))) + "\n\t") + ("\n\t" + "a"))
    }
    if (1 >= 514) {
        v7 = (v7 / float32(1e10))
    }
    v1 = append(v1, "hello")
    if (42 >= 872) {
        for fi8 := 0; fi8 < 3; fi8++ {
            v7 -= v7
            _ = fi8
        }
        for fi9 := 0; fi9 < 1; fi9++ {
            v7 += 0.0
            switch "foo bar" {
            case "a":
                v4 -= v4
            }
            _ = fi9
        }
        var v10 bool = true
        _ = v10
    } else {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint32 = 521
            v7 -= float32((1e10 - (0.0 * -2.25)))
            v6 = 42
            _ = fi11
            _ = v12
        }
    }
    v3 = v3
    for fi13 := 0; fi13 < 10; fi13++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi13
    }
    var v14 int32 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
