package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0(p0 int16) string {
    switch "" {
    case "hello":
        v6 := []int{-2147483648, 1}
        _ = v6
    case "":
        var v7 float64 = (0.1 + (0.1 * (float64(0.1) + 1e10)))
        _ = v7
    case "a":
        switch 42 {
        case -1:
            v8 := GenS1{F0: "foo bar", F1: 6}
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 18; fi9++ {
            p0 = (p0 & p0)
            _ = fi9
        }
    }
    switch "hello" {
    case "hello":
        p0 = (p0 << ((255) & 63))
    }
    _ = p0
    return ""
}

func main() {
    var v1 float64 = (0.1 - 0.1)
    var v2 string = "a"
    var v3 uint64 = 1
    v4 := make(map[uint32]uint32)
    v5 := []int{674, 127, 127}
    v10 := GenS2{F0: -1}
    v1 = v1
    v3 = 18446744073709551615
    v2 = v2
    v11 := v4[42]
    v11 = (^(v11))
    v1 = v1
    if false {
        if (float32(float64((3.14159 - float32(0.0)))) == -2.25) {
            var v12 string = (("\n\t" + ("" + v2)) + (("hello" + v2) + v2))
            if len(v5) > 0 {
                v5[0] = 0
            }
            _ = v12
        }
    } else {
        if true {
            if (!true) {
                v3 = (v3 & v3)
                for fi13 := 0; fi13 < 8; fi13++ {
                    v2 = ""
                    var v14 float32 = float32(-2.25)
                    for fi15 := 0; fi15 < 16; fi15++ {
                        v16 := GenS2{F0: 459}
                        v2 = v2
                        if true {
                            var v17 int = fi13
                            var v18 float64 = v1
                            _ = v17
                            _ = v18
                        }
                        if (!(!(false && true))) {
                            fmt.Println(v1, fi15)
                            fmt.Println(v14)
                            fmt.Println(fi15, fi15)
                            v11 = v11
                            var v19 uint32 = 2
                            _ = v19
                        }
                        _ = fi15
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                }
                fmt.Println(v1, v1)
            } else {
                fmt.Println(v11, v1)
                if false {
                    v20 := GenS3{F0: (v3 * v3), F1: 194}
                    _ = v20
                } else {
                    var v21 int16 = 1
                    v11 = 2147483648
                    _ = v21
                }
                v3 = 68
            }
        } else {
            if true {
                fmt.Println(v2, v3)
            } else {
                v1 = (0.0 + v1)
                for fi22 := 0; fi22 < 15; fi22++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi22
                }
            }
            if (("\n\t" + "hello") <= "\n\t") {
                var v23 int64 = 1
                _ = len(v4)
                _ = v23
            } else {
                if false {
                    if (!(!(!(!(0 >= -1))))) {
                        var v24 uint64 = 2
                        switch v2 {
                        case "\n\t":
                            fmt.Println(v3)
                        case "":
                            if len(v5) > 0 {
                                v5[0] = 127
                            }
                        case "hello":
                            if true {
                                fmt.Println(v3)
                                v25 := v4[v11]
                                _ = v25
                            }
                        default:
                            var v26 uint32 = 992
                            _ = v26
                        }
                        _ = v24
                    }
                    v3 = v3
                }
            }
        }
    }
    var v27 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
