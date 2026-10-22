package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int16 }
type GenS2 struct { F0 int64; F1 int; F2 int32 }
type GenS3 struct { F0 bool; F1 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 1
    v3 := make(map[string]string)
    switch "" {
    case "\n\t":
        if (!(true && ((v2 & v2) > (-(v2))))) {
            v4 := GenS1{F0: 127, F1: 0, F2: -32768}
            _ = v4
        } else {
            if (346 > -486) {
                if true {
                    var v5 int32 = (^(v2))
                    var v6 int = 643
                    v3["\n\t"] = "a"
                    _ = v5
                    _ = v6
                } else {
                    if (!true) {
                        fmt.Println(v2)
                        v7 := GenS2{F0: 0, F1: -2147483648, F2: 2147483647}
                        if (!(42 <= (v2 * (v2 - 699)))) {
                            fmt.Println(v2)
                            if true {
                                fmt.Println(v2)
                                var v8 int16 = int16(v2)
                                _ = v8
                            }
                            var v9 uint16 = 635
                            _ = v9
                        }
                        _ = v7
                    }
                    v10 := []int{0}
                    var v11 bool = (42 != 1)
                    _ = v10
                    _ = v11
                }
                var v12 uint8 = 128
                v12 = (v12 | 2)
                _ = v12
            } else {
                v2 -= (v2 - (v2 + v2))
                if false {
                    var v13 uint16 = 1
                    _ = v13
                }
            }
            v14 := GenS2{F0: 0, F1: 32767, F2: 837}
            _ = v14
        }
    default:
        var v15 uint = 4294967295
        _ = v15
    }
    var v16 int16 = 0
    v17 := []int{32768, 2147483647, 297, 1}
    var v18 bool = ((true && true) || (false && true))
    fmt.Println(v2, v16, v16)
    if ((!true) || (v18 && (!v18))) {
        fmt.Println(v2, v16)
        _ = len(v3)
    } else {
        var v19 int8 = -1
        _ = v19
    }
    v18 = (410 < 42)
    fmt.Println(v2, v16)
    var v20 uint8 = 42
    v21 := GenS3{F0: (!(255 <= 2)), F1: "foo bar"}
    v22 := GenS2{F0: 1, F1: 32767, F2: v2}
    delete(v1, 4294967295)
    v18 = (!((!(-139 != 0)) || false))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
