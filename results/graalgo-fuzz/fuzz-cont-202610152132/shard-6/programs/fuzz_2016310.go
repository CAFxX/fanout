package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint }
type GenS2 struct { F0 int32; F1 int8 }

func main() {
    v1 := []int{0, 237, 4294967295, -350}
    var v2 uint8 = 42
    v3 := make(map[int]bool)
    var v4 string = "hello"
    v5 := GenS1{F0: 2, F1: -473, F2: 1}
    v4 = v4
    var v6 uint = 65535
    switch 9223372036854775807 {
    case -1:
        v2 += 1
    case 250:
        fmt.Println(v6, v6)
    case 385:
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 int = 303
            switch v4 {
            case "foo bar":
                for fi9 := 0; fi9 < 14; fi9++ {
                    var v10 int64 = -1
                    fmt.Println(v10, v4)
                    fmt.Println(v10, v10, v10)
                    var v11 int16 = -32768
                    delete(v3, -136)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            _ = fi7
            _ = v8
        }
    default:
        var v12 string = v4
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        fmt.Println(v6, v6, v6)
        _ = fi13
    }
    fmt.Println(v4, v2, v4)
    v6 += (^(v6))
    v14 := GenS2{F0: 0, F1: 42}
    v2 = uint8(v6)
    var v15 float64 = -2.25
    var v16 uint64 = 2
    v16 = v16
    v17 := v3[4294967295]
    v18 := GenS1{F0: 2, F1: -9223372036854775808, F2: (v6 + 332)}
    for fi19 := 0; fi19 < 13; fi19++ {
        fmt.Println(v15)
        v17 = v17
        _ = fi19
    }
    var v20 int16 = 559
    var v21 string = (v4 + "\n\t")
    v4 = (v21 + v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
