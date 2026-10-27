package main

import (
    "fmt"
)


func main() {
    var v1 int32 = 42
    var v2 float64 = 1.5
    var v3 int = 32767
    fmt.Println(v2)
    v4 := []int{(v3 - v3), (v3 << ((255) & 63))}
    if false {
        var v5 float64 = 0.1
        var v6 string = ("a" + "foo bar")
        _ = v5
        _ = v6
    } else {
        v2 += v2
    }
    v3 = v3
    if true {
        var v7 int = (v3 >> ((0) & 63))
        v2 -= v2
        _ = v7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int8 = 0
        _ = fi8
        _ = v9
    }
    var v10 float64 = 3.14159
    var v11 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
