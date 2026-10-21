package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 string = ""
    v2 := make(map[string]string)
    v3 := []uint32{295, 0, 65535}
    v4 := []uint32{42, 1}
    var v5 uint16 = 0
    if false {
        var v6 uint16 = (v5 | v5)
        _ = v6
    }
    fmt.Println(v1, v1)
    v7 := GenS1{F0: 1, F1: 828, F2: 2}
    v8 := GenS2{F0: false, F1: 42, F2: (v5 * (^(v5)))}
    switch (v1 + ((v1 + v1) + v1)) {
    case "\n\t":
        switch "hello" {
        case "a":
            switch 127 {
            case 545:
                var v9 uint16 = v5
                _ = v9
            case -351:
                for fi10 := 0; fi10 < 10; fi10++ {
                    switch (fi10 | 32768) {
                    case 198:
                        for fi11 := 0; fi11 < 6; fi11++ {
                            v5 -= v5
                            v2[v1] = v1
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                            _ = fi11
                        }
                    }
                    _ = fi10
                }
            case 128:
                v5 += (v5 / (((v5 - 65535)) | 1))
            default:
                for fi12 := 0; fi12 < 3; fi12++ {
                    v5 -= (v5 & 0)
                    _ = fi12
                }
            }
        case "":
            var v13 uint32 = 27
            _ = v13
        default:
            v14 := GenS3{F0: 1}
            _ = v14
        }
    case "":
        v15 := v2[v1]
        _ = v15
    default:
        v1 = "a"
    }
    _ = len(v4)
    switch v1 {
    case "\n\t":
        v5 += 65535
    case "foo bar":
        var v16 int8 = 42
        _ = v16
    default:
        v17 := GenS2{F0: true, F1: 29, F2: 255}
        _ = v17
    }
    var v18 float64 = ((0.1 * 3.14159) * -2.25)
    v5 -= v5
    var v19 int16 = 129
    var v20 int = -982
    v5 -= v5
    _ = len(v4)
    if false {
        v21 := GenS3{F0: 4294967295}
        _ = v21
    }
    v20 -= v20
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
