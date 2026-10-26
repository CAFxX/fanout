package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]string)
    v3 := []int{128, 1, 32767}
    fmt.Println(v1)
    if (false && (v1 != v1)) {
        for fi4 := 0; fi4 < 10; fi4++ {
            var v5 int = (-(fi4))
            _ = fi4
            _ = v5
        }
    }
    v6 := GenS1{F0: 431}
    if len(v3) > 0 {
        v3[0] = 42
    }
    v1 = (v1 >> ((18446744073709551615) & 63))
    var v7 uint16 = (^(v1))
    fmt.Println(v1)
    var v8 int8 = 127
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint32 = 2147483647
        v1 = v7
        var v11 uint32 = (v10 >> ((902) & 63))
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int16 = -32768
    fmt.Println(v1, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
