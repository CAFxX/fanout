package main

import (
    "fmt"
)


func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    var v4 int32 = 415
    var v5 uint8 = 31
    fmt.Println(v5)
    for fi6 := 0; fi6 < 16; fi6++ {
        v5 += 1
        _ = fi6
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(v4, v4)
    case "":
        switch 42 {
        case -1:
            for fi7 := 0; fi7 < 11; fi7++ {
                fmt.Println(v4, v4)
                _ = fi7
            }
        case -583:
            fmt.Println(v1)
        default:
            for fi8 := 0; fi8 < 13; fi8++ {
                for fi9 := 0; fi9 < 3; fi9++ {
                    v5 = (v5 | v5)
                    switch "hello" {
                    case "\n\t":
                        v4 -= 42
                    case "":
                        var v10 uint32 = 1
                        _ = v10
                    case "a":
                        _ = len(v2)
                    }
                    var v11 int64 = 1
                    _ = fi9
                    _ = v11
                }
                _ = fi8
            }
        }
    case "a":
        var v12 string = ("foo bar" + ("foo bar" + "a"))
        _ = v12
    }
    var v13 float32 = (0.0 / 0.1)
    v14 := []int{2147483647, -35, -414}
    var v15 int32 = v4
    switch "a" {
    case "a":
        _ = len(v14)
    case "\n\t":
        switch "" {
        case "":
            v14 = append(v14, -9223372036854775808)
        }
    }
    var v16 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
