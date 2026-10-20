package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }
type GenS2 struct { F0 int; F1 uint16; F2 int32 }
func genFn0(p0 int64, p1 uint16) uint {
    if false {
        var v7 uint8 = 16
        var v8 string = ((("" + "hello") + "\n\t") + "foo bar")
        _ = v7
        _ = v8
    } else {
        fmt.Println(p0, p0)
    }
    p0 = p0
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 int8 = -1
    v2 := make(map[string]string)
    var v3 int64 = 1
    v4 := make(map[int]bool)
    v5 := make(map[string]bool)
    var v6 int8 = (v1 - v1)
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 int = 1
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 32767
    v6 += v1
    switch -871 {
    case 2147483647:
        var v12 int16 = 1000
        _ = v12
    case -525:
        v11 += 2
    case 127:
        if (0 < 65535) {
            v13 := GenS1{F0: 6, F1: false, F2: 4294967295}
            fmt.Println(v11, v3, v1)
            v1 = v1
            _ = v13
        }
    default:
        v14 := []int{32768, 9223372036854775807, 900, -439}
        _ = v14
    }
    delete(v5, "")
    var v15 float64 = 0.0
    v16 := []int{287}
    var v17 int16 = -661
    v18 := GenS1{F0: 0, F1: (true && (("a" + "hello") >= "\n\t")), F2: 42}
    fmt.Println(v6, v15)
    v19 := GenS2{F0: 0, F1: (v11 ^ (v11 * v11)), F2: 1}
    switch "hello" {
    case "\n\t":
        var v20 bool = true
        _ = v20
    case "":
        fmt.Println(v17, v1)
    case "hello":
        var v21 string = ""
        _ = v21
    }
    var v22 int32 = 1
    var v23 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
