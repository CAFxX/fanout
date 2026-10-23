package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{2, 2147483648}
    var v2 uint64 = 213
    v3 := []uint32{941}
    var v4 uint32 = 703
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint32 = v4
        _ = fi5
        _ = v6
    }
    var v7 uint8 = 1
    v4 = v4
    if true {
        fmt.Println(v4, v7)
        var v8 int16 = -32768
        for fi9 := 0; fi9 < 13; fi9++ {
            switch "" {
            case "\n\t":
                var v10 string = ("foo bar" + "foo bar")
                _ = v10
            case "hello":
                var v11 int64 = -9223372036854775808
                _ = v11
            default:
                var v12 uint16 = 65535
                _ = v12
            }
            _ = fi9
        }
        _ = v8
    } else {
        switch ("" + ("foo bar" + "\n\t")) {
        case "foo bar":
            fmt.Println(v7, v2)
        }
        fmt.Println(v2)
    }
    v2 += (v2 / ((v2) | 1))
    v4 = v4
    var v13 uint8 = (v7 - (v7 * (v7 * 246)))
    v1 = append(v1, (v4 * 2147483647))
    var v14 int = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
