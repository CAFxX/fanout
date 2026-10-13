package main

import (
    "fmt"
)

func genFn0() uint16 {
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 int32 = -330
            var v8 bool = false
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 uint16 = 65535
        _ = fi5
        _ = v9
    }
    var v10 int16 = 32767
    _ = v10
    return 141
}
func genFn1(p0 uint32, p1 int8, p2 uint) uint32 {
    v11 := []int{680, 42}
    p1 -= int8(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (p0 - (^(p0)))
}

func main() {
    var v1 bool = true
    var v2 uint64 = 297
    var v3 int16 = -911
    var v4 uint = 255
    v12 := []int{-432, 1, 61, -1}
    var v13 uint64 = 4294967295
    switch 4294967295 {
    case 32768:
        fmt.Println(v13, v1, v1)
    default:
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 float64 = 1.5
            _ = fi14
            _ = v15
        }
    }
    v1 = ("a" != "\n\t")
    switch "hello" {
    case "hello":
        v1 = v1
    case "\n\t":
        var v16 int64 = 9223372036854775807
        _ = v16
    case "":
        fmt.Println(v1, v4)
    }
    var v17 int16 = (v3 >> ((v4) & 63))
    if len(v12) > 0 {
        v12[0] = -1
    }
    v2 -= 18446744073709551615
    var v18 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
