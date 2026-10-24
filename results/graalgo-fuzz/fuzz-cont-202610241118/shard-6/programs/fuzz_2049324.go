package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 65535
    var v2 bool = true
    v3 := []int{}
    var v4 float64 = 3.14159
    var v5 int = 32768
    v6 := []string{"foo bar", ("foo bar" + "hello"), "hello"}
    v3 = append(v3, v5)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v6 = append(v6, (("" + "\n\t") + "hello"))
    fmt.Println(v5, v5, v4)
    v4 = (1e10 * (0.0 + (v4 / (v4 / v4))))
    var v7 uint32 = (^(v1))
    if false {
        v7 = (v1 << ((542) & 63))
        var v8 uint64 = 1
        v7 = (v7 ^ (v1 << ((18446744073709551615) & 63)))
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v5 = (v5 - (v5 / ((fi9) | 1)))
        _ = fi9
    }
    switch "" {
    case "foo bar":
        fmt.Println(v1, v7, v2)
    default:
        var v10 int64 = -946
        _ = v10
    }
    if v2 {
        if v2 {
            v7 = (v1 - v7)
            if len(v3) > 0 {
                v3[0] = 128
            }
        }
    }
    if len(v6) > 0 {
        v6[0] = "a"
    }
    var v11 uint8 = 128
    var v12 uint = uint(v1)
    var v13 string = ("hello" + ((("\n\t" + "foo bar") + "hello") + "a"))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
