package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    v3 := []string{"hello", "hello", ("hello" + "a")}
    var v4 int8 = 0
    if ("a" > "hello") {
        fmt.Println(v4)
    }
    v4 = -47
    var v5 uint32 = 2147483647
    v5 -= v5
    v4 += (^(v4))
    if (-425 >= 32767) {
        for fi6 := 0; fi6 < 11; fi6++ {
            if (true && true) {
                v4 -= 1
                v5 += 65535
            } else {
                var v7 int16 = 0
                if true {
                    v8 := GenS1{F0: 589}
                    v9 := GenS1{F0: 2}
                    switch fi6 {
                    case -862:
                        v7 += (v7 ^ v7)
                    case -248:
                        v10 := GenS1{F0: 0}
                        _ = v10
                    default:
                        fmt.Println(v7, v4)
                    }
                    _ = v8
                    _ = v9
                } else {
                    fmt.Println(v7, v7, v5)
                    v5 = (v5 & 65535)
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                }
                _ = v7
            }
            _ = fi6
        }
        var v11 float64 = 0.1
        switch "" {
        case "":
            v12 := GenS1{F0: 1}
            _ = v12
        case "foo bar":
            for fi13 := 0; fi13 < 6; fi13++ {
                v14 := GenS1{F0: 331}
                v15 := GenS1{F0: 1}
                v16 := GenS1{F0: 42}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        case "a":
            var v17 uint32 = v5
            _ = v17
        default:
            fmt.Println(v4, v5, v11)
        }
        _ = v11
    }
    v18 := GenS2{F0: 899, F1: (v4 / ((v4) | 1))}
    var v19 int16 = 42
    v3 = append(v3, "hello")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
