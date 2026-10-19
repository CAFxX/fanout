package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int16; F2 uint64 }

func main() {
    v1 := []uint32{2}
    var v2 uint64 = 6
    v3 := []string{"a", (("\n\t" + "a") + "foo bar")}
    var v4 float64 = -2.25
    var v5 int = 4294967295
    v6 := make(map[string]string)
    var v7 uint16 = 203
    fmt.Println(v7)
    var v8 int32 = 2147483647
    fmt.Println(v5, v4)
    var v9 uint = 18446744073709551615
    v5 += 2147483647
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: (^(v2))}
        _ = fi10
        _ = v11
    }
    switch "foo bar" {
    case "hello":
        v12 := GenS1{F0: 2}
        _ = v12
    default:
        delete(v6, ("foo bar" + "hello"))
    }
    fmt.Println(v7, v2)
    _ = len(v1)
    v7 += (v7 / ((v7) | 1))
    if ((v7 ^ v7) != (v7 | (v7 & v7))) {
        fmt.Println(v9, v7)
        fmt.Println(v2, v7, v7)
        for fi13 := 0; fi13 < 4; fi13++ {
            if true {
                v9 = v9
                if ((v7 | v7) == 672) {
                    var v14 uint = uint(v2)
                    for fi15 := 0; fi15 < 11; fi15++ {
                        if (42 > 2) {
                            if ("foo bar" < "a") {
                                v4 -= (v4 - (v4 * ((1.5 * (float64(0.1) / v4)) - float64(0.1))))
                                _ = len(v6)
                                v14 = (v14 * (v9 >> ((1) & 63)))
                            }
                            var v16 uint32 = 2147483647
                            v2 += (v2 | 4294967295)
                            _ = v16
                        }
                        _ = fi15
                    }
                    _ = v14
                } else {
                    v17 := GenS2{F0: int(v9)}
                    v7 += v7
                    _ = v17
                }
                v6[(("hello" + "\n\t") + "\n\t")] = (("\n\t" + "a") + "\n\t")
            }
            v18 := GenS2{F0: v5}
            fmt.Println(v8, v9)
            _ = fi13
            _ = v18
        }
    }
    v19 := GenS3{F0: v5, F1: 901, F2: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
