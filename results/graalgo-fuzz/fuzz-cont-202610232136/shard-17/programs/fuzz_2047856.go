package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int8 }
type GenS2 struct { F0 bool; F1 string; F2 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 uint; F1 int; F2 uint32 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]string)
    var v3 float64 = 1e10
    var v4 bool = false
    var v5 uint = 255
    var v6 float32 = (1e10 + ((0.0 - 3.14159) / -2.25))
    v4 = v4
    v6 -= (1.5 / ((((v6 + v6) - v6) / (0.1 * float32((v3 - v3)))) * v6))
    var v7 int = -1
    switch "\n\t" {
    case "foo bar":
        var v8 uint32 = 0
        _ = v8
    case "":
        if (4294967295 >= 0) {
            v6 = (v6 / float32((v3 - v3)))
        } else {
            if v4 {
                if v4 {
                    switch v7 {
                    case -71:
                        for fi9 := 0; fi9 < 5; fi9++ {
                            v10 := GenS1{F0: 42, F1: 112, F2: 1}
                            v2[42] = "\n\t"
                            _ = fi9
                            _ = v10
                        }
                    case -9223372036854775808:
                        var v11 float64 = (v3 / 0.1)
                        _ = v11
                    }
                    switch ("foo bar" + "\n\t") {
                    case "":
                        v12 := GenS1{F0: -525, F1: 127, F2: 1}
                        _ = v12
                    case "a":
                        v13 := GenS2{F0: (!v4), F1: "", F2: 2147483647}
                        _ = v13
                    default:
                        v14 := GenS3{F0: 1000}
                        _ = v14
                    }
                }
                switch "hello" {
                case "":
                    fmt.Println(v7, v5, v3)
                case "a":
                    var v15 int8 = -128
                    _ = v15
                }
            } else {
                var v16 uint64 = 18446744073709551615
                _ = v16
            }
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 string = ""
                if v4 {
                    switch 128 {
                    case -1:
                        for fi19 := 0; fi19 < 1; fi19++ {
                            var v20 uint = 0
                            _ = fi19
                            _ = v20
                        }
                    case -119:
                        for fi21 := 0; fi21 < 9; fi21++ {
                            switch v7 {
                            case 17:
                                var v22 uint32 = 43
                                var v23 uint64 = 725
                                fmt.Println(fi21)
                                _ = v22
                                _ = v23
                            default:
                                var v24 uint64 = 0
                                _ = v24
                            }
                            v25 := GenS4{F0: 17, F1: (fi21 ^ v7)}
                            _ = fi21
                            _ = v25
                        }
                    case 754:
                        v26 := GenS1{F0: 0, F1: 0, F2: 1}
                        _ = v26
                    }
                    _ = len(v2)
                } else {
                    v7 += (-(fi17))
                    v27 := GenS5{F0: (v5 + v5), F1: (-(v7)), F2: 2147483648}
                    v3 += 1.5
                    _ = v27
                }
                _ = fi17
                _ = v18
            }
        }
    default:
        var v28 int8 = 2
        _ = v28
    }
    v3 = float64(0.0)
    var v29 int8 = 127
    v30 := GenS5{F0: v5, F1: v7, F2: 42}
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
