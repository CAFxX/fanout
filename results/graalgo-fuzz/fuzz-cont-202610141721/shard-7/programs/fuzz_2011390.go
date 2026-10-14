package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int16; F2 uint16 }

func main() {
    var v1 int = -1
    v2 := make(map[int]bool)
    var v3 int = v1
    var v4 bool = true
    v5 := []int{}
    switch v3 {
    case 4294967295:
        var v6 float32 = (3.14159 + 3.14159)
        _ = v6
    default:
        v1 += (v1 / (((v1 ^ v3)) | 1))
    }
    switch v1 {
    case 128:
        v7 := GenS1{F0: 1, F1: 880}
        _ = v7
    case 960:
        if (4294967295 > 2147483647) {
            fmt.Println(v3, v1, v4)
        } else {
            fmt.Println(v4)
        }
    case -3:
        v8 := GenS2{F0: 4294967295}
        _ = v8
    }
    v1 = v3
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: 2147483647, F1: 1}
        _ = fi9
        _ = v10
    }
    v2[(v3 & (v1 | 1))] = true
    for fi11 := 0; fi11 < 8; fi11++ {
        switch "\n\t" {
        case "hello":
            v3 = fi11
        case "\n\t":
            for fi12 := 0; fi12 < 6; fi12++ {
                fmt.Println(v4, v3)
                if false {
                    if (v4 && (("" + (("foo bar" + "foo bar") + "\n\t")) >= "\n\t")) {
                        v4 = (115 >= 1)
                        v13 := GenS1{F0: 848, F1: 32767}
                        fmt.Println(v4, v3)
                        _ = v13
                    } else {
                        v3 += (v3 | (-(v1)))
                        v14 := GenS1{F0: 4294967295, F1: 0}
                        v3 = v3
                        _ = v14
                    }
                }
                v15 := GenS1{F0: 977, F1: 42}
                _ = fi12
                _ = v15
            }
        default:
            fmt.Println(v3, fi11, v1)
        }
        _ = len(v2)
        if len(v5) > 0 {
            v5[0] = (-(v3))
        }
        _ = fi11
    }
    fmt.Println(v3, v3, v1)
    if v4 {
        var v16 uint32 = 2147483647
        _ = v16
    } else {
        for fi17 := 0; fi17 < 17; fi17++ {
            v3 += (v1 + v1)
            _ = fi17
        }
    }
    var v18 int16 = 455
    v2[(^(v1))] = v4
    var v19 uint = 1
    var v20 float32 = 0.0
    if (65535 == v19) {
        if v4 {
            fmt.Println(v4, v19, v4)
            delete(v2, v3)
            v21 := GenS3{F0: (!v4), F1: 1, F2: 704}
            _ = v21
        }
        fmt.Println(v3, v4, v4)
        var v22 bool = (!v4)
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
