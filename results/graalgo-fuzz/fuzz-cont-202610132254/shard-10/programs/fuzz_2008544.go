package main

import (
    "fmt"
)

func genFn0() uint16 {
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    var v8 float32 = -2.25
    var v9 int32 = 1
    _ = v8
    _ = v9
    return 378
}

func main() {
    v1 := []string{"foo bar", "\n\t", ("hello" + ""), ("foo bar" + "")}
    var v2 int8 = -76
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    var v5 bool = ((-(v2)) > 100)
    for fi10 := 0; fi10 < 13; fi10++ {
        _ = len(v1)
        _ = fi10
    }
    fmt.Println(v5, v5)
    var v11 uint16 = 42
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 float32 = (0.1 * 0.1)
        _ = len(v3)
        switch "\n\t" {
        case "foo bar":
            var v14 float32 = (v13 * (v13 * v13))
            _ = v14
        case "\n\t":
            fmt.Println(v13)
        case "hello":
            v15 := v3[("foo bar" + "\n\t")]
            _ = v15
        default:
            v11 -= (v11 % (((v11 | v11)) | 1))
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v5)
    v2 = 29
    v2 = 0
    var v16 int16 = -53
    var v17 bool = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
