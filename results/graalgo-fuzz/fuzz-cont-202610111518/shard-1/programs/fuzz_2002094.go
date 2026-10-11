package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int }
type GenS2 struct { F0 int32; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int; F2 int16 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 int64 = 0
    var v4 int32 = -1
    v5 := GenS1{F0: 2, F1: 0, F2: -489}
    v6 := GenS1{F0: 2147483647, F1: 946, F2: 128}
    fmt.Println(v3, v4)
    v7 := []int{}
    var v8 int64 = 399
    var v9 int16 = 32767
    for fi10 := 0; fi10 < 6; fi10++ {
        switch int(v3) {
        case 32767:
            fmt.Println(fi10)
        case 9223372036854775807:
            if (1e10 > ((((1.5 + 1.5) + 0.1) / 0.1) + float64(0.0))) {
                v11 := v1[(^(fi10))]
                switch fi10 {
                case 128:
                    if false {
                        v3 += 313
                        v12 := GenS2{F0: 0, F1: -900, F2: 192}
                        _ = v12
                    }
                default:
                    switch "" {
                    case "a":
                        _ = len(v7)
                    case "foo bar":
                        if false {
                            var v13 float32 = 3.14159
                            _ = v13
                        } else {
                            var v14 float32 = 1.5
                            switch 9223372036854775807 {
                            case -1:
                                v15 := GenS3{F0: 47, F1: (fi10 & fi10), F2: v9}
                                _ = v15
                            default:
                                v8 = -134
                            }
                            switch "" {
                            case "foo bar":
                                var v16 string = ""
                                _ = v16
                            }
                            _ = v14
                        }
                    default:
                        v17 := GenS4{F0: v9}
                        _ = v17
                    }
                }
                for fi18 := 0; fi18 < 11; fi18++ {
                    v19 := GenS1{F0: v11, F1: 201, F2: int(v9)}
                    _ = fi18
                    _ = v19
                }
                _ = v11
            } else {
                fmt.Println(v3, v9)
            }
        case 127:
            if len(v7) > 0 {
                v7[0] = (^(fi10))
            }
        }
        var v20 float32 = -2.25
        _ = fi10
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        delete(v2, 2147483647)
        _ = fi21
    }
    if ((3.14159 * 0.0) == -2.25) {
        _ = len(v2)
        v1[-2147483648] = 707
    }
    v9 += (^(v9))
    var v22 uint32 = 885
    fmt.Println(v8, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
