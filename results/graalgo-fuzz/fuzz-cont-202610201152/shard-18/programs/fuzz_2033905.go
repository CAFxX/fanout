package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }

func main() {
    var v1 int = 32767
    v2 := make(map[uint32]int)
    v3 := []int{(-(v1)), (v1 % ((42) | 1)), (v1 + (v1 >> ((550) & 63))), (-(v1))}
    var v4 uint64 = 0
    var v5 uint64 = v4
    var v6 int64 = -9223372036854775808
    v3 = append(v3, 9223372036854775807)
    var v7 bool = (929 > 2147483648)
    if v7 {
        v8 := v2[51]
        _ = v8
    }
    var v9 uint16 = 32767
    fmt.Println(v4, v5, v7)
    switch (("foo bar" + (((("foo bar" + "hello") + "") + "hello") + "hello")) + "\n\t") {
    case "":
        fmt.Println(v9, v4, v1)
    case "foo bar":
        var v10 uint16 = v9
        _ = v10
    case "hello":
        fmt.Println(v6, v7, v4)
    }
    v7 = false
    v9 = v9
    v7 = true
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: 32767, F1: "", F2: (v9 * v9)}
        _ = fi11
        _ = v12
    }
    switch 128 {
    case -9223372036854775808:
        var v13 string = ("hello" + ("hello" + ("" + ("\n\t" + ("a" + "hello")))))
        _ = v13
    case 4294967295:
        var v14 float64 = 0.0
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
