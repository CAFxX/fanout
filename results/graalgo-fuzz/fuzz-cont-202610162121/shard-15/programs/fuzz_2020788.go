package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint }

func main() {
    v1 := []int{-139, 2147483647, 127, 32767}
    v2 := []int{520, 128, 252}
    var v3 uint8 = 2
    v4 := make(map[string]uint32)
    var v5 float64 = (float64(((0.0 * float32((0.1 + -2.25))) * 0.0)) - 3.14159)
    var v6 uint64 = 631
    var v7 uint16 = 0
    var v8 uint8 = v3
    fmt.Println(v5)
    v8 -= 42
    switch -377 {
    case -9223372036854775808:
        switch 9223372036854775807 {
        case -1:
            for fi9 := 0; fi9 < 2; fi9++ {
                v3 = 127
                _ = fi9
            }
        default:
            switch "hello" {
            case "":
                fmt.Println(v3, v3, v5)
            }
        }
    default:
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 18; fi11++ {
                switch fi10 {
                case 32768:
                    switch 9223372036854775807 {
                    case 767:
                        v8 += (^(v8))
                    case 32767:
                        v6 = 42
                    }
                case 127:
                    v8 = v8
                case 42:
                    v12 := GenS1{F0: 0}
                    _ = v12
                default:
                    if (!(!(1 != 32767))) {
                        var v13 int32 = -720
                        v3 = v8
                        switch fi11 {
                        case 9223372036854775807:
                            if (0 != 1) {
                                v7 = (v7 * v7)
                                var v14 float32 = 3.14159
                                var v15 int = fi10
                                _ = v14
                                _ = v15
                            } else {
                                var v16 int8 = 14
                                v17 := GenS2{F0: (v7 & v7)}
                                v13 -= -868
                                switch "" {
                                case "a":
                                    var v18 int64 = 113
                                    var v19 int16 = int16(v16)
                                    var v20 int16 = v19
                                    if true {
                                        v19 = (v20 % ((v19) | 1))
                                        var v21 uint64 = v6
                                        _ = v21
                                    }
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                }
                                v5 = 0.1
                                _ = v16
                                _ = v17
                            }
                        case 4294967295:
                            v3 -= (v8 ^ v8)
                        case 506:
                            var v22 int = fi10
                            _ = v22
                        }
                        _ = v13
                    }
                }
                _ = fi11
            }
            _ = fi10
        }
    }
    var v23 uint16 = (v7 + (^(v7)))
    v24 := GenS3{F0: -1, F1: v23, F2: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
