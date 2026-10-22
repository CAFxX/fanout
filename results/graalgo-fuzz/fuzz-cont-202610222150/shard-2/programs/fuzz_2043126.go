package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint64 = 0
    var v2 int64 = -9223372036854775808
    var v3 int8 = 110
    v4 := make(map[uint32]string)
    v5 := []string{"foo bar", ""}
    v1 += (^(v1))
    var v6 int32 = -1
    v7 := v4[317]
    v3 -= v3
    v1 -= 2
    fmt.Println(v2)
    if true {
        v4[1] = (v7 + v7)
    }
    v7 = "hello"
    if (false && (int16(v6) >= 42)) {
        var v8 int8 = (^(v3))
        v9 := GenS1{F0: 1000}
        _ = v8
        _ = v9
    }
    var v10 string = ((v7 + (("\n\t" + v7) + v7)) + v7)
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 bool = (!true)
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
