package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 42
    var v2 string = ""
    var v3 uint32 = 505
    var v4 int16 = -638
    v5 := []int{-9223372036854775808, -506, -22}
    fmt.Println(v4, v1)
    v1 -= v1
    v3 -= 65535
    v1 = v1
    for fi6 := 0; fi6 < 9; fi6++ {
        _ = len(v5)
        _ = fi6
    }
    fmt.Println(v4, v4, v3)
    switch (v2 + (((("foo bar" + "foo bar") + "foo bar") + "") + v2)) {
    case "a":
        if len(v5) > 0 {
            v5[0] = 1
        }
    case "\n\t":
        var v7 uint16 = 42
        _ = v7
    default:
        fmt.Println(v1, v1, v2)
    }
    var v8 int = 128
    var v9 bool = (v8 >= v8)
    var v10 uint8 = 200
    switch 32768 {
    case 42:
        switch (v2 + "\n\t") {
        case "hello":
            var v11 float32 = ((-2.25 + (0.0 / (1.5 / 1e10))) / -2.25)
            _ = v11
        }
    case 127:
        var v12 uint8 = (^(v10))
        _ = v12
    default:
        var v13 uint8 = (v10 >> ((1) & 63))
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
