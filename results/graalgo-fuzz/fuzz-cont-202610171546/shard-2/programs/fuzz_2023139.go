package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint = 42
    v2 := []string{"", "hello", "", "a"}
    var v3 uint64 = 425
    var v4 int64 = -9223372036854775808
    var v5 int8 = 1
    var v6 uint64 = 18446744073709551615
    v5 = (^(v5))
    var v7 float32 = 0.0
    if true {
        v3 -= v6
        v8 := GenS1{F0: "hello"}
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        v1 = (v1 | v1)
        _ = fi9
    }
    v10 := GenS1{F0: "\n\t"}
    v2 = append(v2, "hello")
    if (!false) {
        v5 += (v5 >> ((v1) & 63))
        for fi11 := 0; fi11 < 17; fi11++ {
            v3 = v6
            _ = fi11
        }
    } else {
        if false {
            var v12 int8 = -1
            _ = v12
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
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
