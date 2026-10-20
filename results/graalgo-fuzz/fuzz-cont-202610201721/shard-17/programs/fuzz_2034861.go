package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (0.1 - -2.25)
    var v3 int32 = -864
    v4 := make(map[string]uint32)
    var v5 uint64 = 0
    if true {
        v3 -= v3
        var v6 int64 = -1
        v2 -= float32(0.1)
        _ = v6
    }
    var v7 string = "\n\t"
    var v8 int8 = -1
    var v9 uint = 734
    v10 := []int{}
    var v11 int8 = (v8 % ((v8) | 1))
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 bool = true
        _ = fi12
        _ = v13
    }
    var v14 int8 = v11
    switch 42 {
    case -9223372036854775808:
        v15 := GenS1{F0: v9, F1: v5}
        _ = v15
    default:
        if (1 != 127) {
            v16 := GenS2{F0: ((v7 + "a") + (((((v7 + v7) + (v7 + (v7 + v7))) + v7) + v7) + "")), F1: 221, F2: ((!false) || (!false))}
            var v17 uint = 18446744073709551615
            for fi18 := 0; fi18 < 7; fi18++ {
                var v19 bool = true
                _ = fi18
                _ = v19
            }
            _ = v16
            _ = v17
        } else {
            var v20 uint64 = v5
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 float64 = -2.25
        switch fi21 {
        case -288:
            fmt.Println(v11, v3, v14)
        case 2147483647:
            var v23 uint64 = 42
            _ = v23
        default:
            for fi24 := 0; fi24 < 16; fi24++ {
                _ = len(v10)
                v25 := GenS2{F0: (("a" + v7) + ((v7 + (("a" + ("" + "")) + v7)) + "hello")), F1: 65, F2: (v7 > v7)}
                if (!false) {
                    switch v7 {
                    case "foo bar":
                        if true {
                            _ = len(v10)
                        }
                    case "a":
                        v3 -= v3
                    case "":
                        v26 := GenS1{F0: 18446744073709551615, F1: (v5 % (((v5 - 689)) | 1))}
                        _ = v26
                    }
                    fmt.Println(v3, v8)
                }
                _ = fi24
                _ = v25
            }
        }
        v27 := v4[v7]
        _ = fi21
        _ = v22
        _ = v27
    }
    var v28 int64 = -492
    var v29 float32 = (float32(float64(v2)) / (v2 - v2))
    var v30 bool = true
    if len(v10) > 0 {
        v10[0] = -1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
