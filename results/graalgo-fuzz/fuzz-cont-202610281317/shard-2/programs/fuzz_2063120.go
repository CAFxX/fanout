package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int16 }

func main() {
    v1 := []int{32767, 0}
    var v2 int8 = 0
    var v3 uint = 18446744073709551615
    v4 := make(map[string]int)
    v5 := []string{"hello"}
    v2 = -1
    var v6 uint64 = 18446744073709551615
    var v7 int16 = 42
    v2 = (v2 * v2)
    for fi8 := 0; fi8 < 17; fi8++ {
        if true {
            switch "foo bar" {
            case "foo bar":
                v6 += v6
            case "a":
                if false {
                    v9 := v4["\n\t"]
                    _ = v9
                } else {
                    switch "" {
                    case "hello":
                        if false {
                            v2 += v2
                            _ = len(v4)
                            var v10 uint32 = 2
                            _ = v10
                        } else {
                            v11 := GenS1{F0: 2147483647, F1: 42}
                            _ = v11
                        }
                    case "":
                        v12 := GenS2{F0: true}
                        _ = v12
                    }
                    v1 = append(v1, fi8)
                }
            case "\n\t":
                var v13 uint = v3
                _ = v13
            default:
                var v14 int = fi8
                _ = v14
            }
        }
        _ = fi8
    }
    v3 -= (v3 / (((v3 - v3)) | 1))
    var v15 int64 = -9223372036854775808
    var v16 bool = (false || (!((1.5 - (0.1 - (0.0 + 0.0))) > 0.0)))
    _ = len(v4)
    v17 := GenS3{F0: (65535 == 42), F1: v7}
    fmt.Println(v15)
    var v18 uint64 = v6
    v4[""] = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
