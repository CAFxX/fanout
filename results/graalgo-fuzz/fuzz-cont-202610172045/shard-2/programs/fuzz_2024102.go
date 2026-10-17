package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int8; F1 int32 }
func genFn0(p0 int16, p1 int8, p2 uint32) int8 {
    if false {
        v6 := GenS1{F0: 76, F1: 2}
        switch 127 {
        case 128:
            for fi7 := 0; fi7 < 6; fi7++ {
                switch (fi7 >> ((0) & 63)) {
                case -876:
                    v8 := GenS1{F0: 4294967295, F1: 252}
                    _ = v8
                case 32767:
                    var v9 float32 = (-2.25 - ((3.14159 - float32(0.0)) + ((0.1 - (0.1 / -2.25)) - (-2.25 / 1.0))))
                    _ = v9
                case 1:
                    if true {
                        var v10 uint16 = 229
                        _ = v10
                    }
                default:
                    if false {
                        var v11 float64 = ((float64((-2.25 / (1.5 / float32(1.5)))) * 1.5) - 0.0)
                        var v12 uint32 = 42
                        p2 -= (p2 * (^(v12)))
                        _ = v11
                        _ = v12
                    }
                }
                var v13 uint64 = 741
                switch "" {
                case "a":
                    var v14 float64 = 0.0
                    _ = v14
                case "foo bar":
                    if true {
                        p2 += (p2 + (p2 >> ((255) & 63)))
                        for fi15 := 0; fi15 < 13; fi15++ {
                            v16 := []int{(^(fi7)), (fi7 & fi7)}
                            v17 := GenS1{F0: 696, F1: 0}
                            p2 = p2
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    }
                case "hello":
                    p2 += p2
                }
                _ = fi7
                _ = v13
            }
        }
        v18 := GenS1{F0: 887, F1: 0}
        _ = v18
        _ = v6
    } else {
        for fi19 := 0; fi19 < 11; fi19++ {
            fmt.Println(p2, p0)
            _ = fi19
        }
    }
    v20 := []int{60, 128, 316}
    var v21 int = 4294967295
    switch "" {
    case "foo bar":
        v20 = append(v20, (v21 >> ((255) & 63)))
    case "":
        p2 = p2
    case "\n\t":
        var v22 uint = 2
        _ = v22
    default:
        var v23 float32 = 3.14159
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return p1
}
func genFn1(p0 int8) uint16 {
    p0 += -60
    p0 += (p0 | 1)
    p0 += genFn0(-32768, p0, 0)
    var v24 float64 = (-2.25 / -2.25)
    _ = p0
    _ = v24
    return 255
}

func main() {
    v1 := []uint32{1, 154, 2147483647}
    v2 := make(map[int]bool)
    v3 := []string{"\n\t", "", "hello", ("foo bar" + "hello")}
    v4 := make(map[uint32]string)
    var v5 uint8 = 1
    v5 -= v5
    for fi25 := 0; fi25 < 7; fi25++ {
        v5 = 2
        _ = fi25
    }
    var v26 float32 = 1.5
    v27 := GenS1{F0: 1, F1: 42}
    fmt.Println(v26)
    v5 += (v5 >> ((0) & 63))
    if false {
        v28 := GenS1{F0: 0, F1: 644}
        var v29 string = "a"
        for fi30 := 0; fi30 < 16; fi30++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v1)
            _ = fi30
        }
        _ = v28
        _ = v29
    }
    var v31 bool = ((("\n\t" + "a") + ("" + "")) >= "")
    _ = len(v3)
    v32 := GenS2{F0: 42, F1: -2147483648}
    fmt.Println(v5, v31)
    fmt.Println(v31)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
