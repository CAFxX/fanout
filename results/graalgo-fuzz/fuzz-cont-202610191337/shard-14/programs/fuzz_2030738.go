package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int8; F1 int8; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = ((3.14159 - 0.0) + 0.0)
    var v3 int32 = 0
    var v4 float64 = (1.5 - 1e10)
    var v5 uint64 = 2
    v2 = (((((v2 / 0.1) + 1.5) / v2) - (v2 + 0.1)) / 1e10)
    switch -592 {
    case 9223372036854775807:
        var v6 string = "hello"
        _ = v6
    }
    v2 = v2
    if ((-32768 == 1000) && false) {
        v7 := GenS1{F0: -82, F1: 1}
        v3 -= (v3 | (v3 | v3))
        v8 := GenS1{F0: 42, F1: -276}
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        switch "foo bar" {
        case "":
            if true {
                var v10 int16 = 1000
                v5 -= (v5 / ((v5) | 1))
                fmt.Println(v10)
                _ = v10
            }
        default:
            v11 := GenS2{F0: "", F1: 42}
            _ = v11
        }
        _ = fi9
    }
    v4 -= 0.0
    for fi12 := 0; fi12 < 19; fi12++ {
        if (v4 != 1.5) {
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS1{F0: 100, F1: 9223372036854775807}
                _ = fi13
                _ = v14
            }
            var v15 float64 = -2.25
            var v16 int16 = 1
            _ = v15
            _ = v16
        } else {
            switch "foo bar" {
            case "hello":
                if true {
                    for fi17 := 0; fi17 < 6; fi17++ {
                        v18 := v1[2]
                        var v19 uint = 2
                        fmt.Println(v5)
                        v20 := GenS1{F0: -5, F1: -1}
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    if false {
                        switch fi12 {
                        case -211:
                            for fi21 := 0; fi21 < 7; fi21++ {
                                v4 -= v4
                                _ = fi21
                            }
                        case -439:
                            v4 = v4
                        default:
                            v4 = -2.25
                        }
                    }
                } else {
                    var v22 string = ("" + "hello")
                    _ = v22
                }
            case "":
                v5 -= 18446744073709551615
            case "\n\t":
                if false {
                    if ((v5 & 0) < (v5 * 2)) {
                        fmt.Println(v2, v2)
                        var v23 int32 = 952
                        for fi24 := 0; fi24 < 8; fi24++ {
                            if (2 == 710) {
                                v4 += float64(((v2 / float32(v4)) - float32(float64((v2 / v2)))))
                                fmt.Println(v3)
                                v25 := GenS3{F0: -128, F1: -9, F2: (fi12 - fi12)}
                                _ = v25
                            }
                            fmt.Println(v23)
                            switch "" {
                            case "\n\t":
                                switch "hello" {
                                case "a":
                                    v26 := v1[65535]
                                    _ = v26
                                case "":
                                    v27 := GenS2{F0: "hello", F1: 65535}
                                    _ = v27
                                case "hello":
                                    switch fi12 {
                                    case 0:
                                        if false {
                                            v1[2147483647] = 1
                                        }
                                    case 32768:
                                        var v28 bool = false
                                        _ = v28
                                    case 271:
                                        var v29 int32 = (v23 << ((574) & 63))
                                        _ = v29
                                    }
                                default:
                                    v30 := GenS2{F0: "foo bar", F1: 65535}
                                    _ = v30
                                }
                            case "a":
                                delete(v1, 42)
                            case "hello":
                                var v31 int32 = v3
                                _ = v31
                            }
                            _ = fi24
                        }
                        _ = v23
                    } else {
                        v2 -= 1.5
                    }
                    switch "a" {
                    case "hello":
                        v5 -= 0
                    case "":
                        v32 := []int{fi12}
                        _ = v32
                    default:
                        fmt.Println(v5, v4, v2)
                    }
                    fmt.Println(v2, v5, v3)
                } else {
                    v2 += (v2 - v2)
                    var v33 float64 = v4
                    _ = v33
                }
            default:
                delete(v1, 0)
            }
        }
        _ = fi12
    }
    v3 += (v3 + (v3 ^ (v3 % (((^(v3))) | 1))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
