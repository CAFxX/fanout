package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{}
    var v2 int = -2147483648
    v3 := []string{(("hello" + "foo bar") + "")}
    var v4 float32 = -2.25
    v2 = (v2 ^ v2)
    var v5 uint16 = 2
    if ((!(141 > 2147483648)) || (float64(((float32(1e10) * 0.0) / (float32(1.5) / v4))) == 1.5)) {
        fmt.Println(v4, v4, v4)
    }
    switch (("hello" + "foo bar") + "") {
    case "":
        fmt.Println(v5, v4)
    }
    var v6 int32 = -1
    v2 = v2
    var v7 uint = 25
    if (false || false) {
        fmt.Println(v4, v4, v5)
        switch v2 {
        case 9223372036854775807:
            var v8 uint8 = uint8(v7)
            _ = v8
        case 224:
            var v9 float32 = (0.0 - ((3.14159 + v4) - (float32(((0.0 + float64(v4)) / 0.1)) * (1e10 / v4))))
            _ = v9
        case -157:
            var v10 uint8 = 2
            _ = v10
        }
        v6 += v6
    } else {
        v3 = append(v3, ("a" + "a"))
        v2 = 32768
        switch "foo bar" {
        case "\n\t":
            v7 += v7
        }
    }
    fmt.Println(v4, v2, v7)
    var v11 int8 = -90
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
