package main

import (
    "fmt"
)

func genFn0(p0 int, p1 uint32) int8 {
    p0 = 42
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 uint64 = 18446744073709551615
        fmt.Println(p1, p1, v6)
        fmt.Println(p0)
        _ = fi5
        _ = v6
    }
    var v7 int = p0
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    var v1 uint = 42
    var v2 string = ("foo bar" + ("\n\t" + ""))
    v3 := make(map[int]int)
    var v4 uint64 = 1
    var v8 int8 = 42
    v2 = ("hello" + (v2 + "hello"))
    fmt.Println(v4)
    var v9 string = v2
    for fi10 := 0; fi10 < 5; fi10++ {
        fmt.Println(fi10, v8, v9)
        switch v2 {
        case "":
            _ = len(v3)
        case "foo bar":
            fmt.Println(fi10, v4, v9)
        case "\n\t":
            var v11 float32 = 0.0
            _ = v11
        }
        var v12 int64 = -17
        _ = fi10
        _ = v12
    }
    var v13 int64 = -1
    if (-2.25 != float64(((1.5 - float32(-2.25)) / ((1.5 + 1.5) / 1e10)))) {
        var v14 uint64 = (v4 | v4)
        v15 := []int{1, 2147483647}
        _ = v14
        _ = v15
    }
    v1 += v1
    v13 = 232
    for fi16 := 0; fi16 < 8; fi16++ {
        switch v9 {
        case "hello":
            var v17 float64 = 3.14159
            _ = v17
        default:
            fmt.Println(v1)
        }
        _ = fi16
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        if (255 < 0) {
            v2 = v9
        }
        if (v13 >= 1) {
            var v19 uint = (^(v1))
            _ = v19
        }
        _ = fi18
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
