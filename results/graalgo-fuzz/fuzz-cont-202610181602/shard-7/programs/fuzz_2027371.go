package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 string }

func main() {
    var v1 int16 = 473
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 bool = (127 <= 100)
    v5 := []uint32{0, 42, 991}
    var v6 string = (("" + "") + "hello")
    var v7 uint8 = 42
    v6 = v6
    v8 := GenS1{F0: 128, F1: (!v4)}
    v1 = 32767
    if (!false) {
        var v9 uint8 = v7
        var v10 int = -9223372036854775808
        var v11 uint64 = 2
        _ = v10
        _ = v11
        _ = v9
    } else {
        var v12 float32 = (3.14159 - 3.14159)
        fmt.Println(v1, v4, v12)
        var v13 int32 = 0
        _ = v12
        _ = v13
    }
    switch 127 {
    case 509:
        v6 = ""
    case 1:
        v1 = 1000
    case -9223372036854775808:
        v7 = v7
    }
    v14 := GenS2{F0: (v6 + "")}
    var v15 uint8 = (^(v7))
    v6 = v6
    v15 = v7
    fmt.Println(v7, v1, v7)
    switch v6 {
    case "":
        var v16 int64 = 516
        _ = v16
    case "a":
        var v17 int32 = -2147483648
        _ = v17
    }
    if v4 {
        for fi18 := 0; fi18 < 19; fi18++ {
            switch (v6 + v6) {
            case "":
                v19 := v3[(fi18 - 535)]
                _ = v19
            default:
                fmt.Println(fi18)
            }
            var v20 uint16 = 42
            _ = fi18
            _ = v20
        }
        v6 = "\n\t"
        var v21 uint32 = 4294967295
        _ = v21
    }
    var v22 int16 = (-(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
