package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint }
type GenS2 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int32, p1 uint64, p2 bool) uint64 {
    v5 := GenS1{F0: -1, F1: "hello", F2: 370}
    p0 = 42
    v6 := GenS1{F0: -128, F1: "hello", F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 2147483647
    v3 := make(map[int]uint32)
    var v4 uint8 = 0
    var v7 int32 = 2147483647
    var v8 bool = true
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 int16 = -1
        var v11 float64 = ((0.0 / 1.0) / float64(3.14159))
        var v12 float32 = (1e10 - 0.0)
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := v3[-676]
    var v14 int64 = 1
    if true {
        if (v8 || (v8 || false)) {
            v15 := GenS2{F0: 981, F1: -1, F2: 2147483647}
            _ = v15
        } else {
            fmt.Println(v4, v13, v14)
        }
        _ = len(v3)
    }
    var v16 float64 = (((1e10 * 1.5) * 3.14159) * float64(3.14159))
    v14 -= v14
    if (!(128 >= v4)) {
        switch ("\n\t" + (("hello" + "") + ("" + ("" + ("foo bar" + ("" + "\n\t")))))) {
        case "\n\t":
            if v8 {
                var v17 float32 = ((1.5 - 0.0) + -2.25)
                _ = v17
            } else {
                for fi18 := 0; fi18 < 8; fi18++ {
                    for fi19 := 0; fi19 < 20; fi19++ {
                        v7 -= (-(v7))
                        var v20 uint = 255
                        if (v8 || true) {
                            switch "a" {
                            case "hello":
                                switch fi19 {
                                case -306:
                                    fmt.Println(v8)
                                case 127:
                                    var v21 float64 = ((((v16 + v16) + 1.5) / float64(-2.25)) + float64(0.1))
                                    _ = v21
                                case 32768:
                                    var v22 int = fi19
                                    _ = v22
                                default:
                                    if false {
                                        delete(v1, uint32(v4))
                                    } else {
                                        if (int8(v20) == int8(v4)) {
                                            fmt.Println(fi18, v4)
                                            v14 = (-(v14))
                                            var v23 uint32 = (^(v13))
                                            _ = v23
                                        }
                                        switch "hello" {
                                        case "\n\t":
                                            if false {
                                                var v24 string = "a"
                                                _ = v24
                                            } else {
                                                v25 := GenS1{F0: 0, F1: (("" + "") + "hello"), F2: v20}
                                                var v26 uint8 = v4
                                                _ = v25
                                                _ = v26
                                            }
                                        case "a":
                                            v20 -= 18446744073709551615
                                            v27 := GenS2{F0: 42, F1: 65, F2: v13}
                                            _ = v27
                                        }
                                        var v28 uint8 = uint8(v2)
                                        _ = v28
                                    }
                                }
                            case "\n\t":
                                v29 := []int{}
                                _ = v29
                            case "foo bar":
                                v30 := []int{(^(fi19)), fi18, (fi18 % (((fi18 % ((fi19) | 1))) | 1))}
                                _ = v30
                            default:
                                fmt.Println(v20, v20)
                            }
                            var v31 float32 = 0.1
                            fmt.Println(v31, v13, v7)
                            _ = v31
                        } else {
                            v14 = (v14 | (v14 * 1))
                            fmt.Println(v14)
                            if v8 {
                                switch 446 {
                                case 32768:
                                    switch ("" + "foo bar") {
                                    case "\n\t":
                                        fmt.Println(v20, v20)
                                    case "foo bar":
                                        v32 := GenS3{F0: 2}
                                        _ = v32
                                    default:
                                        v7 -= 2147483647
                                    }
                                default:
                                    v16 += float64(-2.25)
                                }
                                v33 := []int{32767, fi18}
                                v8 = (4294967295 == uint64(v20))
                                var v34 float32 = 0.1
                                v20 += (v20 | (v20 - v20))
                                _ = v33
                                _ = v34
                            } else {
                                var v35 int8 = 0
                                v36 := GenS4{F0: 42}
                                _ = v35
                                _ = v36
                            }
                        }
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
            }
        case "a":
            fmt.Println(v2, v16, v14)
        }
    }
    var v37 int32 = int32(v2)
    v16 += (v16 + (0.0 - 3.14159))
    delete(v1, (v13 / ((v13) | 1)))
    for fi38 := 0; fi38 < 12; fi38++ {
        v39 := []int{(fi38 & (fi38 & -509)), 676, -351, fi38}
        _ = fi38
        _ = v39
    }
    v8 = v8
    v1[v13] = v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v7
    _ = v8
}
