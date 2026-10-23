package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []int{-52, 32768, -676}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 2
    v4 := make(map[int]string)
    var v5 uint16 = 229
    v6 := GenS1{F0: 42}
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint32 = 65535
        var v9 int64 = -1
        _ = fi7
        _ = v8
        _ = v9
    }
    v1 = append(v1, 918)
    fmt.Println(v3, v3)
    for fi10 := 0; fi10 < 2; fi10++ {
        v5 = v5
        _ = fi10
    }
    v5 = v5
    fmt.Println(v5, v3)
    var v11 uint8 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4[-568] = "\n\t"
    var v12 int = 874
    if ((((!false) || ((-2.25 * float32(((0.0 - (3.14159 - 1.5)) * (0.1 + (-2.25 / 1.0))))) != -2.25)) && false) || (!(false || false))) {
        switch "hello" {
        case "hello":
            v3 -= 1
        case "\n\t":
            v12 -= v12
        case "":
            for fi13 := 0; fi13 < 17; fi13++ {
                v14 := GenS1{F0: -74}
                for fi15 := 0; fi15 < 1; fi15++ {
                    if (!true) {
                        delete(v4, 0)
                        fmt.Println(v12)
                        var v16 int64 = -1
                        switch fi13 {
                        case 496:
                            v12 += -255
                        case 32768:
                            v5 = 667
                        case 773:
                            v5 += uint16(v3)
                        }
                        fmt.Println(v16, v5, v3)
                        v3 = 42
                        _ = v16
                    }
                    if (!false) {
                        var v17 string = "foo bar"
                        var v18 uint8 = v11
                        var v19 uint64 = (v3 / ((1) | 1))
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        v5 = v5
                        fmt.Println(v12)
                    }
                    _ = fi15
                }
                _ = fi13
                _ = v14
            }
        }
        var v20 float64 = 1.5
        _ = v20
    }
    if (563 > 502) {
        var v21 int8 = -1
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 uint = 18446744073709551615
            _ = fi22
            _ = v23
        }
        var v24 int8 = 99
        _ = v21
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
