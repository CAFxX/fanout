package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint }
type GenS2 struct { F0 uint; F1 uint64 }

func main() {
    var v1 float64 = (1e10 + -2.25)
    var v2 string = (("\n\t" + "") + "foo bar")
    v3 := []string{v2, "a", "a"}
    var v4 uint = 42
    v5 := []string{"\n\t", (("a" + v2) + v2)}
    v6 := GenS1{F0: 255, F1: 0, F2: 354}
    var v7 float64 = v1
    if len(v3) > 0 {
        v3[0] = v2
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        if (!(uint32(v4) > 0)) {
            var v9 float64 = ((v7 + ((0.0 + float64(3.14159)) * float64((((((0.1 - 0.1) / 0.1) * 0.0) / 1.5) / 1.5)))) - v7)
            var v10 uint32 = 0
            var v11 string = v2
            _ = v10
            _ = v11
            _ = v9
        } else {
            var v12 int = fi8
            _ = v12
        }
        fmt.Println(v2, v7)
        if (!(32767 < 1)) {
            fmt.Println(v1, v7, fi8)
            switch (("" + (v2 + v2)) + "hello") {
            case "hello":
                v2 = v2
            case "":
                switch "foo bar" {
                case "hello":
                    v1 = v1
                case "":
                    if true {
                        for fi13 := 0; fi13 < 14; fi13++ {
                            var v14 int = 4294967295
                            _ = fi13
                            _ = v14
                        }
                    }
                default:
                    var v15 float32 = -2.25
                    _ = v15
                }
            case "a":
                v5 = append(v5, v2)
            }
        }
        _ = fi8
    }
    v16 := GenS2{F0: v4, F1: 0}
    v17 := GenS1{F0: 19, F1: 1, F2: (v4 | 255)}
    fmt.Println(v4, v2)
    var v18 uint32 = 51
    v19 := GenS1{F0: 42, F1: v18, F2: v4}
    var v20 uint32 = 82
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
