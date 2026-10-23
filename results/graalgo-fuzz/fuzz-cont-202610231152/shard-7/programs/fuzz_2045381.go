package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 string = ""
    var v2 bool = (!true)
    v3 := make(map[int]string)
    v4 := []string{(v1 + v1), ""}
    var v5 uint8 = 181
    var v6 uint32 = 16
    var v7 uint8 = (v5 - v5)
    var v8 uint16 = 2
    v3[73] = v1
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 int = fi9
        v10 -= (v10 + (fi9 / ((341) | 1)))
        _ = fi9
        _ = v10
    }
    var v11 int8 = 42
    fmt.Println(v6, v11)
    v7 -= (v7 % ((137) | 1))
    v12 := GenS1{F0: 128}
    if true {
        v6 += (v6 & (^(v6)))
        v2 = true
    } else {
        if (true && (v6 >= (v6 & v6))) {
            delete(v3, -929)
        } else {
            v1 = ""
            var v13 int32 = 42
            _ = v13
        }
        v8 -= v8
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
