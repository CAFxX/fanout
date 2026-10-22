package main

import (
    "fmt"
)


func main() {
    var v1 int64 = -673
    var v2 int16 = -1
    v3 := []uint32{}
    v4 := make(map[int]int)
    v2 -= (v2 / (((v2 + 42)) | 1))
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(v2, v1)
        _ = fi5
    }
    switch ("a" + ("" + ("hello" + ""))) {
    case "a":
        var v6 float32 = 1.5
        _ = v6
    }
    var v7 uint16 = 998
    v2 = -32768
    var v8 uint8 = 251
    fmt.Println(v2, v1)
    v1 += 321
    var v9 float64 = 1e10
    fmt.Println(v9)
    if (127 <= 682) {
        var v10 uint8 = v8
        _ = v10
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 uint16 = 255
        for fi13 := 0; fi13 < 14; fi13++ {
            fmt.Println(v7, fi13)
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
