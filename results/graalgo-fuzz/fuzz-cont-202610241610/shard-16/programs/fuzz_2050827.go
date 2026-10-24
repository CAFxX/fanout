package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 2
    v3 := make(map[int]int)
    v4 := []uint32{287, 2147483648, 1}
    _ = len(v4)
    switch -637 {
    case -979:
        v5 := v3[0]
        _ = v5
    case 4294967295:
        var v6 bool = true
        _ = v6
    case 251:
        switch 127 {
        case 374:
            var v7 float32 = float32((((1e10 * (3.14159 - -2.25)) * 1e10) / -2.25))
            _ = v7
        default:
            v2 = v2
        }
    default:
        var v8 uint64 = 912
        _ = v8
    }
    v2 = 134
    switch ("a" + ("\n\t" + (("foo bar" + ("hello" + "")) + "hello"))) {
    case "a":
        if len(v4) > 0 {
            v4[0] = 1
        }
    case "foo bar":
        v2 += 127
    }
    v2 += v2
    var v9 int64 = 42
    v2 = (v2 + v2)
    v2 -= 152
    v2 -= (v2 / ((127) | 1))
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
