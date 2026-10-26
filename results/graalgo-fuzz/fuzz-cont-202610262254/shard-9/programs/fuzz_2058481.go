package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint8 }
func genFn0(p0 uint8) int8 {
    for fi6 := 0; fi6 < 19; fi6++ {
        if true {
            v7 := GenS1{F0: -1, F1: 1000}
            if true {
                fmt.Println(p0)
            } else {
                v8 := []int{}
                var v9 int = (fi6 + fi6)
                p0 += (p0 * (p0 + p0))
                _ = v8
                _ = v9
            }
            _ = v7
        }
        p0 -= p0
        _ = fi6
    }
    var v10 int16 = -1
    v10 = v10
    _ = p0
    _ = v10
    return 79
}

func main() {
    var v1 uint = 65535
    v2 := make(map[string]int)
    var v3 float32 = 1e10
    v4 := []int{-675, -712}
    v5 := []string{"a"}
    var v11 uint64 = 0
    fmt.Println(v3, v11, v1)
    fmt.Println(v11)
    if (!true) {
        _ = len(v5)
        v1 = v1
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            if true {
                switch "\n\t" {
                case "":
                    switch fi12 {
                    case 42:
                        fmt.Println(fi12)
                        var v14 bool = (-861 > -281)
                        switch 9223372036854775807 {
                        case -1:
                            v3 -= (v3 / float32(-2.25))
                        default:
                            v15 := GenS1{F0: 123, F1: 1000}
                            _ = v15
                        }
                        _ = v14
                    case 161:
                        var v16 bool = (-2147483648 < 0)
                        _ = v16
                    }
                case "\n\t":
                    v3 = v3
                }
            } else {
                v17 := GenS1{F0: genFn0(42), F1: -664}
                _ = v17
            }
            _ = fi13
        }
        _ = fi12
    }
    fmt.Println(v11, v1)
    v18 := GenS2{F0: false, F1: 42, F2: 29}
    var v19 uint8 = 101
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
