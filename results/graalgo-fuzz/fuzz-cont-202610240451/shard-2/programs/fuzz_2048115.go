package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
func genFn0(p0 int, p1 int) uint64 {
    p1 = (^(p1))
    var v5 string = ("foo bar" + "hello")
    _ = p0
    _ = p1
    _ = v5
    return 462
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint8 = 33
    var v3 string = "a"
    var v4 string = "\n\t"
    var v6 uint8 = 1
    var v7 int32 = -2147483648
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int = fi8
        v3 = ((v4 + ((v3 + v3) + "\n\t")) + "hello")
        var v10 float64 = -2.25
        _ = fi8
        _ = v10
        _ = v9
    }
    v2 += (v2 >> ((2) & 63))
    fmt.Println(v7)
    v11 := GenS1{F0: ("" + v4), F1: ""}
    var v12 uint32 = (v1 + (v1 * v1))
    v7 += v7
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
