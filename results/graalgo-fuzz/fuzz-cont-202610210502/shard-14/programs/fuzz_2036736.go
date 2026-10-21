package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 uint16 = 1
    var v2 uint32 = 618
    var v3 uint = 535
    var v4 int8 = 65
    var v5 int8 = 100
    v6 := make(map[uint32]uint32)
    v7 := []int{-2147483648, -9223372036854775808}
    for fi8 := 0; fi8 < 7; fi8++ {
        if false {
            for fi9 := 0; fi9 < 7; fi9++ {
                switch "" {
                case "\n\t":
                    v10 := GenS1{F0: -9223372036854775808}
                    _ = v10
                case "hello":
                    switch ("" + "\n\t") {
                    case "":
                        v11 := GenS1{F0: (^(fi9))}
                        _ = v11
                    case "a":
                        v1 = 0
                    }
                }
                var v12 uint8 = 255
                _ = fi9
                _ = v12
            }
        } else {
            v13 := GenS2{F0: (v5 % ((v5) | 1))}
            _ = v13
        }
        if (2 == 32767) {
            if len(v7) > 0 {
                v7[0] = fi8
            }
            var v14 uint16 = (^(v1))
            _ = v14
        } else {
            v3 -= (v3 & (v3 >> (((v3 >> (((v3 * 255)) & 63))) & 63)))
            if false {
                var v15 float32 = -2.25
                _ = v15
            } else {
                switch ("" + ("" + "")) {
                case "hello":
                    var v16 string = ""
                    _ = v16
                case "a":
                    switch "foo bar" {
                    case "foo bar":
                        delete(v6, (v2 & v2))
                    case "a":
                        v17 := GenS2{F0: (v4 * v5)}
                        _ = v17
                    default:
                        for fi18 := 0; fi18 < 10; fi18++ {
                            var v19 uint8 = 255
                            if len(v7) > 0 {
                                v7[0] = 32768
                            }
                            v4 = 42
                            _ = fi18
                            _ = v19
                        }
                    }
                case "foo bar":
                    var v20 int32 = 821
                    _ = v20
                }
                var v21 int = fi8
                _ = v21
            }
        }
        _ = fi8
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        fmt.Println(v3)
        _ = fi22
    }
    v23 := GenS2{F0: v5}
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 uint64 = 18446744073709551615
        _ = fi24
        _ = v25
    }
    v2 = v2
    v1 -= (v1 + (v1 ^ 65535))
    switch ("foo bar" + "hello") {
    case "":
        var v26 int64 = 0
        _ = v26
    case "\n\t":
        v27 := v6[(^(v2))]
        _ = v27
    default:
        if false {
            if false {
                switch 9223372036854775807 {
                case 4294967295:
                    fmt.Println(v5)
                case 127:
                    fmt.Println(v1, v3, v1)
                case 32768:
                    fmt.Println(v1, v2, v1)
                }
                fmt.Println(v5, v4)
            } else {
                var v28 uint32 = 42
                _ = v28
            }
        }
    }
    v29 := GenS3{F0: 2, F1: v2}
    switch ("hello" + "\n\t") {
    case "foo bar":
        v3 = (v3 - v3)
    }
    fmt.Println(v2)
    switch 0 {
    case -920:
        v4 -= (v5 - v5)
    default:
        var v30 float64 = -2.25
        _ = v30
    }
    var v31 uint8 = 98
    v3 += v3
    var v32 float32 = ((1e10 + 0.1) - 3.14159)
    fmt.Println(v32, v32, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
