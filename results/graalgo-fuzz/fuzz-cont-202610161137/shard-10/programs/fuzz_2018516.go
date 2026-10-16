package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int64 }
type GenS2 struct { F0 string }

func main() {
    var v1 float64 = 3.14159
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]int)
    v5 := []string{(("\n\t" + "hello") + "")}
    var v6 float64 = 1.5
    v2 = v2
    switch 962 {
    case 32768:
        fmt.Println(v1, v6)
    }
    var v7 uint16 = 344
    for fi8 := 0; fi8 < 1; fi8++ {
        if true {
            switch (fi8 / (((fi8 / ((1) | 1))) | 1)) {
            case 9223372036854775807:
                var v9 uint = uint(v2)
                _ = v9
            default:
                for fi10 := 0; fi10 < 4; fi10++ {
                    switch ("" + "hello") {
                    case "\n\t":
                        v1 = v6
                    default:
                        var v11 bool = true
                        _ = v11
                    }
                    var v12 float64 = v6
                    var v13 float32 = ((3.14159 - 1.5) / 1.0)
                    if false {
                        v12 += v1
                        fmt.Println(v7, v7)
                    }
                    if true {
                        v14 := v3[2]
                        _ = v14
                    } else {
                        fmt.Println(fi10)
                    }
                    _ = fi10
                    _ = v12
                    _ = v13
                }
            }
            v6 = (0.1 - v6)
        } else {
            var v15 uint8 = 4
            _ = v15
        }
        _ = fi8
    }
    if len(v5) > 0 {
        v5[0] = ("a" + "\n\t")
    }
    v6 = (v6 + v6)
    var v16 float32 = 1.5
    v17 := v4[(("a" + "a") + "a")]
    v17 -= 144
    switch v17 {
    case 128:
        v18 := GenS1{F0: v17, F1: "foo bar", F2: -9223372036854775808}
        _ = v18
    case -115:
        var v19 bool = true
        _ = v19
    }
    v20 := GenS2{F0: "foo bar"}
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    switch v17 {
    case -534:
        v17 -= (v17 >> ((1) & 63))
    default:
        var v21 int = int(v7)
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
