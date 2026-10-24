package main

import (
    "fmt"
)

func genFn0(p0 uint32) int {
    v7 := []int{834, -2147483648}
    p0 += (p0 << ((65535) & 63))
    switch "" {
    case "foo bar":
        p0 -= (p0 >> ((1) & 63))
    }
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 32767
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v4 uint8 = (^(v1))
    v5 := make(map[string]string)
    v6 := []string{"foo bar"}
    var v8 int8 = 100
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int32 = -1
        _ = fi9
        _ = v10
    }
    fmt.Println(v4, v8, v1)
    if len(v6) > 0 {
        v6[0] = "\n\t"
    }
    switch ("hello" + "hello") {
    case "hello":
        switch "a" {
        case "hello":
            fmt.Println(v4, v4, v1)
        default:
            v4 -= (v4 + v4)
        }
    default:
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 int = (fi11 - int(v1))
            var v13 uint16 = 161
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    var v14 int = genFn0(2147483647)
    var v15 float64 = 3.14159
    v1 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
