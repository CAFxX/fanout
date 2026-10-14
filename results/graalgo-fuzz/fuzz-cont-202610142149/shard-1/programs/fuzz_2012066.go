package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint8, p1 int8, p2 uint32) uint8 {
    v6 := GenS1{F0: 32767}
    switch "hello" {
    case "\n\t":
        fmt.Println(p1)
    case "hello":
        for fi7 := 0; fi7 < 13; fi7++ {
            fmt.Println(fi7, p0, p2)
            var v8 uint16 = 0
            _ = fi7
            _ = v8
        }
    case "a":
        var v9 uint64 = 859
        _ = v9
    }
    switch ("" + ("a" + "\n\t")) {
    case "hello":
        v10 := []int{811, 9223372036854775807}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 74
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[int]bool)
    var v3 float64 = 3.14159
    var v4 uint8 = 128
    var v5 string = "foo bar"
    fmt.Println(v4)
    var v11 int64 = 9223372036854775807
    v12 := []int{2147483647, 63, -1, -529}
    v4 = (v4 % ((v4) | 1))
    var v13 int32 = 144
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            fmt.Println(v13, fi15, fi14)
            _ = fi15
        }
        delete(v2, (fi14 / ((fi14) | 1)))
        _ = fi14
    }
    fmt.Println(v5)
    switch -681 {
    case -817:
        var v16 uint16 = 0
        _ = v16
    case 32768:
        var v17 float64 = (3.14159 * v3)
        _ = v17
    case 4294967295:
        v1 -= 65535
    default:
        v18 := GenS1{F0: -757}
        _ = v18
    }
    fmt.Println(v11, v11, v3)
    v11 = (v11 / (((^(v11))) | 1))
    fmt.Println(v11, v11)
    fmt.Println(v4)
    var v19 int64 = v11
    var v20 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
