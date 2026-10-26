package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]int)
    var v3 float32 = (1e10 / ((3.14159 + (1.5 * ((0.1 + float32(((0.1 * ((1e10 * -2.25) / v1)) - (float64((0.1 * 1.5)) + v1)))) / float32(1e10)))) * 1e10))
    v1 += (v1 - (v1 - ((0.1 - 1.5) / 1.5)))
    v4 := v2[32768]
    var v5 int = v4
    var v6 uint8 = 0
    if true {
        var v7 float64 = (v1 / (v1 / 0.1))
        fmt.Println(v1, v3, v6)
        if (-128 >= 1) {
            var v8 uint16 = 42
            for fi9 := 0; fi9 < 2; fi9++ {
                switch "" {
                case "":
                    var v10 uint = 2
                    _ = v10
                case "foo bar":
                    var v11 uint64 = 0
                    _ = v11
                default:
                    switch ("foo bar" + "") {
                    case "\n\t":
                        v12 := []int{int(v8), v5}
                        _ = v12
                    case "foo bar":
                        fmt.Println(v3, v1, v4)
                    case "a":
                        fmt.Println(v4, v7)
                    default:
                        if (!true) {
                            delete(v2, (fi9 ^ 9223372036854775807))
                            var v13 string = "hello"
                            _ = v13
                        }
                    }
                }
                var v14 uint8 = (v6 | v6)
                if false {
                    v4 += v5
                    var v15 uint16 = 943
                    _ = v15
                } else {
                    v16 := GenS1{F0: 1}
                    _ = v16
                }
                _ = fi9
                _ = v14
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                fmt.Println(v1, v6)
                _ = fi17
            }
            _ = v8
        } else {
            v18 := GenS1{F0: 42}
            v19 := GenS1{F0: -1}
            _ = v18
            _ = v19
        }
        _ = v7
    }
    if (1 == 1000) {
        if (v1 <= -2.25) {
            if true {
                switch "" {
                case "hello":
                    switch "foo bar" {
                    case "foo bar":
                        var v20 int8 = -101
                        _ = v20
                    }
                case "":
                    var v21 int16 = -187
                    _ = v21
                }
                if true {
                    switch "" {
                    case "":
                        for fi22 := 0; fi22 < 6; fi22++ {
                            v1 = v1
                            v4 = fi22
                            _ = fi22
                        }
                    case "a":
                        var v23 uint32 = 2147483647
                        _ = v23
                    }
                }
            } else {
                v24 := []int{v5, (v5 & 643), 127, (v5 + v4)}
                v5 = (v5 << ((65535) & 63))
                var v25 uint32 = 1
                _ = v24
                _ = v25
            }
            var v26 int = v4
            _ = v26
        }
        var v27 int16 = 937
        v4 -= v4
        _ = v27
    } else {
        if false {
            var v28 int16 = int16(v6)
            v29 := GenS1{F0: (v28 | (v28 % ((572) | 1)))}
            v28 += v28
            _ = v28
            _ = v29
        }
    }
    v30 := []int{(v4 << ((18446744073709551615) & 63))}
    var v31 int8 = 127
    var v32 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
