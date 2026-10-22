package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[string]string)
    v2 := []int{127}
    v3 := make(map[string]string)
    var v4 int32 = 42
    var v5 string = "foo bar"
    var v6 int = 4294967295
    v7 := GenS1{F0: 42}
    v4 += (^(v4))
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int64 = 834
        v6 = 9223372036854775807
        _ = fi8
        _ = v9
    }
    if (!(78 >= 4294967295)) {
        if len(v2) > 0 {
            v2[0] = v6
        }
        var v10 uint16 = 320
        delete(v3, v5)
        _ = v10
    }
    var v11 string = ""
    v4 = 42
    if len(v2) > 0 {
        v2[0] = v6
    }
    var v12 int32 = (v4 + v4)
    v4 -= v12
    v13 := GenS1{F0: 0}
    if len(v2) > 0 {
        v2[0] = -320
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(v11, v12)
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 int16 = 1
            v5 = ("foo bar" + (v5 + v11))
            var v17 uint64 = 18446744073709551615
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 int = fi14
        _ = fi14
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
