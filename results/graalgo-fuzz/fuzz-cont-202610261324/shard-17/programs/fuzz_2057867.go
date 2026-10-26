package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int; F2 int }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint; F1 int32; F2 int8 }

func main() {
    var v1 uint = 1
    var v2 int = 2147483647
    var v3 uint16 = uint16(v1)
    var v4 int16 = -200
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: 18446744073709551615, F1: 1}
    fmt.Println(v3)
    v7 := []int{}
    fmt.Println(v2, v4)
    v8 := GenS2{F0: int(v3)}
    v9 := GenS3{F0: (v1 - 18446744073709551615), F1: v2, F2: 0}
    v1 -= uint(v3)
    _ = len(v7)
    var v10 float32 = (0.0 - (0.0 + (0.1 * (float32(3.14159) - -2.25))))
    if false {
        for fi11 := 0; fi11 < 2; fi11++ {
            fmt.Println(v2, v1)
            var v12 uint = v1
            _ = fi11
            _ = v12
        }
        if ((v1 & v1) == v1) {
            var v13 string = ""
            var v14 int64 = 42
            _ = v13
            _ = v14
        }
        var v15 uint64 = 18446744073709551615
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 uint16 = (v3 >> (((^(v1))) & 63))
            v1 -= 2
            _ = fi17
            _ = v18
        }
        fmt.Println(v10)
        if (713 >= 329) {
            var v19 uint = v1
            var v20 int = fi16
            for fi21 := 0; fi21 < 19; fi21++ {
                if (137 <= 45) {
                    switch "" {
                    case "a":
                        v22 := GenS4{F0: v1, F1: 1}
                        _ = v22
                    case "hello":
                        if ("" != "") {
                            var v23 string = "a"
                            fmt.Println(fi16, fi21, fi21)
                            _ = v23
                        }
                    case "\n\t":
                        v10 = float32((0.0 + (1.5 - 1.5)))
                    default:
                        var v24 string = "\n\t"
                        _ = v24
                    }
                } else {
                    _ = len(v5)
                }
                v25 := GenS5{F0: (v19 ^ (v19 >> (((v19 / ((v19) | 1))) & 63))), F1: -1, F2: 42}
                if true {
                    v19 = (v19 + (v1 << (((v1 ^ v1)) & 63)))
                }
                _ = fi21
                _ = v25
            }
            _ = v19
            _ = v20
        } else {
            v1 -= 18446744073709551615
        }
        _ = fi16
    }
    var v26 int32 = -1
    switch "a" {
    case "":
        var v27 bool = (false && false)
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
