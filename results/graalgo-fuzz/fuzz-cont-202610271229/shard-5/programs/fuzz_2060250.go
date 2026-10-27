package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint16 = 107
    var v2 float32 = 1e10
    var v3 float64 = 1e10
    v4 := make(map[uint32]string)
    var v5 string = ("\n\t" + "a")
    v1 += v1
    if (1 != 0) {
        v6 := []int{2147483647}
        _ = v6
    }
    v7 := []int{32768, 128}
    v8 := GenS1{F0: 1}
    switch v5 {
    case "foo bar":
        if true {
            v9 := GenS1{F0: 4294967295}
            v1 = v1
            _ = v9
        } else {
            var v10 uint8 = 0
            _ = v10
        }
    default:
        for fi11 := 0; fi11 < 5; fi11++ {
            switch fi11 {
            case 4294967295:
                v1 = 691
            default:
                v5 = ((v5 + v5) + "a")
            }
            _ = fi11
        }
    }
    v12 := GenS1{F0: 42}
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int16 = 0
        var v15 int32 = 1
        v15 = (v15 >> ((255) & 63))
        _ = fi13
        _ = v14
        _ = v15
    }
    if ((^(v1)) < 42) {
        if false {
            v5 = "foo bar"
        }
        var v16 int = 2147483647
        _ = v16
    }
    var v17 int32 = 2147483647
    var v18 uint16 = (^(v1))
    var v19 int = 32768
    var v20 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
