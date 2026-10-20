package main

import (
    "fmt"
)


func main() {
    var v1 float64 = -2.25
    var v2 int = -9223372036854775808
    v3 := make(map[string]string)
    v4 := []string{("\n\t" + "foo bar"), "\n\t"}
    var v5 uint64 = 1
    var v6 float32 = 1.5
    _ = len(v3)
    if (151 > 525) {
        v6 = ((v6 + v6) / 1.0)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v7 uint32 = 2
        _ = v7
    }
    switch (^(v2)) {
    case 9223372036854775807:
        var v8 float32 = (v6 - -2.25)
        _ = v8
    case -651:
        fmt.Println(v5, v2, v5)
    case 128:
        v5 = 1
    }
    v6 += float32((1.5 / float64(3.14159)))
    v5 = (v5 % ((v5) | 1))
    var v9 uint8 = 215
    var v10 int64 = 42
    var v11 int8 = -1
    v2 += (v2 * (^(v2)))
    var v12 uint16 = 889
    var v13 float64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
