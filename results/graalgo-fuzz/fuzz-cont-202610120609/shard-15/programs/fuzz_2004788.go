package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[string]string)
    var v2 bool = (-1 < 876)
    var v3 int64 = -575
    var v4 uint16 = 1
    var v5 int64 = (-(v3))
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := v1["a"]
        v8 := []int{fi6, fi6, (fi6 / (((fi6 / (((fi6 + fi6)) | 1))) | 1))}
        _ = fi6
        _ = v7
        _ = v8
    }
    v3 -= v5
    var v9 float64 = -2.25
    var v10 uint16 = (^(v4))
    var v11 float32 = ((1e10 / 3.14159) + 1e10)
    var v12 uint = 615
    var v13 int64 = (v5 >> ((255) & 63))
    v14 := []int{-1, 320}
    v11 = float32((((v9 / v9) * -2.25) * ((1e10 + float64(v11)) * (v9 / -2.25))))
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS1{F0: fi15}
        switch ("a" + ("hello" + "a")) {
        case "a":
            if v2 {
                v4 -= v10
                var v17 uint32 = 4294967295
                _ = v17
            }
        case "hello":
            var v18 int64 = v13
            _ = v18
        case "\n\t":
            v13 = (v3 + v13)
        default:
            var v19 float32 = ((v11 + -2.25) / (float32(1.5) - v11))
            _ = v19
        }
        _ = len(v1)
        _ = fi15
        _ = v16
    }
    v3 = v5
    if (!v2) {
        var v20 int16 = -749
        for fi21 := 0; fi21 < 15; fi21++ {
            for fi22 := 0; fi22 < 15; fi22++ {
                v23 := v1[("\n\t" + ("a" + "\n\t"))]
                v24 := GenS1{F0: fi22}
                fmt.Println(v20, v10)
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
