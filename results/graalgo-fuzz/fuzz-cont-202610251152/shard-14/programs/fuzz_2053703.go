package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 string }
type GenS2 struct { F0 int; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int; F2 uint32 }

func main() {
    var v1 int32 = 42
    var v2 int = 32767
    var v3 uint = 2
    v4 := make(map[string]int)
    v5 := make(map[uint32]string)
    var v6 uint16 = 0
    fmt.Println(v2)
    var v7 int = (v2 >> ((65535) & 63))
    var v8 int32 = (v1 | -115)
    for fi9 := 0; fi9 < 3; fi9++ {
        v6 += v6
        _ = fi9
    }
    switch 128 {
    case 42:
        v3 = 0
    default:
        switch v2 {
        case 798:
            var v10 float64 = 1.5
            _ = v10
        case 42:
            for fi11 := 0; fi11 < 7; fi11++ {
                delete(v5, 42)
                var v12 int = fi11
                _ = fi11
                _ = v12
            }
        }
    }
    var v13 int8 = -128
    v14 := []int{32767}
    if (!false) {
        var v15 uint8 = uint8(v6)
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 bool = false
            switch (fi16 ^ fi17) {
            case -2147483648:
                v14 = append(v14, (v2 % ((0) | 1)))
            case -9223372036854775808:
                if true {
                    switch "foo bar" {
                    case "foo bar":
                        v7 += fi16
                    case "":
                        fmt.Println(v18, v18)
                    default:
                        var v19 uint16 = uint16(v3)
                        _ = v19
                    }
                    v20 := GenS1{F0: 731, F1: -9223372036854775808, F2: ""}
                    _ = v20
                } else {
                    if (319 >= 0) {
                        v21 := GenS2{F0: v7, F1: 4294967295, F2: 100}
                        var v22 string = "a"
                        _ = v21
                        _ = v22
                    }
                    if (0.0 == (1e10 / (1e10 * float32((0.0 - 1e10))))) {
                        fmt.Println(v7)
                    } else {
                        fmt.Println(fi17)
                        var v23 uint16 = v6
                        var v24 int32 = v1
                        _ = v23
                        _ = v24
                    }
                }
            case 441:
                if v18 {
                    v25 := GenS3{F0: (v3 % (((v3 - v3)) | 1))}
                    fmt.Println(v13)
                    v13 += (v13 | v13)
                    _ = v25
                }
            default:
                var v26 string = "a"
                _ = v26
            }
            fmt.Println(fi17, v13)
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    fmt.Println(v1, v2, v8)
    v27 := GenS4{F0: v8, F1: 32767, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
