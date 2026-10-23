package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int16 }

func main() {
    var v1 uint8 = 0
    var v2 float64 = 0.1
    v3 := []int{127, -773}
    v4 := []string{"a"}
    v2 = float64(-2.25)
    v2 = ((((((1e10 / v2) / v2) - (v2 * (float64((-2.25 + 1.5)) - -2.25))) + v2) / (v2 + v2)) - v2)
    for fi5 := 0; fi5 < 2; fi5++ {
        fmt.Println(v1)
        _ = fi5
    }
    v2 = float64(-2.25)
    for fi6 := 0; fi6 < 12; fi6++ {
        if true {
            var v7 uint32 = 4294967295
            _ = v7
        } else {
            fmt.Println(v1, v1, v1)
            for fi8 := 0; fi8 < 12; fi8++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi8
            }
            for fi9 := 0; fi9 < 3; fi9++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v10 uint32 = 2
                var v11 bool = (4294967295 != fi9)
                switch "a" {
                case "foo bar":
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                case "hello":
                    v10 += uint32(v1)
                case "a":
                    var v12 float64 = v2
                    _ = len(v4)
                    switch fi6 {
                    case 53:
                        var v13 int64 = -1
                        _ = v13
                    case 42:
                        v14 := GenS1{F0: "\n\t", F1: 0}
                        _ = v14
                    default:
                        fmt.Println(v2, v1)
                    }
                    _ = v12
                }
                v2 = 3.14159
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        v1 += v1
        _ = fi6
    }
    v15 := GenS2{F0: 591}
    fmt.Println(v1, v2)
    var v16 uint16 = 2
    v17 := GenS2{F0: 207}
    v18 := GenS3{F0: -775, F1: 211}
    v16 = (v16 / ((v16) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
