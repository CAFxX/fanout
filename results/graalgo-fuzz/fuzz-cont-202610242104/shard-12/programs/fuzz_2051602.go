package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }

func main() {
    v1 := []uint32{8, 567}
    var v2 float64 = (0.1 * 1.5)
    v3 := []int{931}
    var v4 uint8 = 255
    var v5 uint = 2
    var v6 float32 = (1e10 * float32(-2.25))
    v4 -= v4
    var v7 uint32 = 2147483648
    fmt.Println(v6, v7, v4)
    v8 := GenS1{F0: 574, F1: ""}
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 uint8 = v4
        _ = fi9
        _ = v10
    }
    fmt.Println(v5, v2)
    var v11 uint = v5
    var v12 uint8 = uint8(v7)
    v13 := GenS1{F0: 0, F1: "hello"}
    v2 = 0.0
    v14 := GenS1{F0: 0, F1: ((("foo bar" + ("\n\t" + "\n\t")) + ("\n\t" + "\n\t")) + ("" + "a"))}
    fmt.Println(v2, v11, v7)
    v1 = append(v1, v7)
    if true {
        fmt.Println(v11, v11)
        var v15 int32 = 2147483647
        var v16 int32 = v15
        _ = v15
        _ = v16
    } else {
        v7 = (v7 & (^(v7)))
        v12 += 0
    }
    var v17 string = "foo bar"
    v17 = v17
    for fi18 := 0; fi18 < 13; fi18++ {
        v4 = v4
        var v19 int = (fi18 >> ((42) & 63))
        _ = fi18
        _ = v19
    }
    v17 = v17
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
