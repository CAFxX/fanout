package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 bool; F2 string }
type GenS5 struct { F0 int32; F1 uint64; F2 int16 }
func genFn0(p0 uint32) int32 {
    v4 := []int{0}
    var v5 float64 = float64(0.0)
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[uint32]uint32)
    var v3 int = 9223372036854775807
    v6 := GenS1{F0: 2147483648, F1: "foo bar", F2: 401}
    fmt.Println(v3)
    var v7 uint16 = 970
    switch v3 {
    case 32768:
        v7 = (v7 + (v7 << ((447) & 63)))
    default:
        fmt.Println(v3, v7)
    }
    if (false && false) {
        _ = len(v1)
        v8 := GenS1{F0: 1, F1: ("" + ("\n\t" + ("a" + ("foo bar" + ("a" + "\n\t"))))), F2: 255}
        for fi9 := 0; fi9 < 19; fi9++ {
            if (0 >= 137) {
                v1 = append(v1, "")
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 int64 = 42
                    switch 127 {
                    case 128:
                        v12 := GenS2{F0: 127}
                        _ = v12
                    case 127:
                        if false {
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(fi9, v3)
                            case "hello":
                                fmt.Println(v11, v3, v7)
                            default:
                                if true {
                                    v13 := GenS1{F0: 42, F1: "a", F2: uint(v7)}
                                    if false {
                                        switch (fi10 + 9223372036854775807) {
                                        case -629:
                                            fmt.Println(v7)
                                        case 9223372036854775807:
                                            v14 := GenS3{F0: (fi9 << ((uint(v7)) & 63))}
                                            _ = v14
                                        case 128:
                                            var v15 int8 = 0
                                            if false {
                                                var v16 string = ("a" + "a")
                                                v11 -= 0
                                                _ = v16
                                            }
                                            delete(v2, 2147483647)
                                            v7 += 42
                                            _ = v15
                                        }
                                        fmt.Println(fi9)
                                        var v17 int32 = -2147483648
                                        switch ("foo bar" + ((("" + ("foo bar" + ((("foo bar" + "hello") + "foo bar") + ""))) + ("\n\t" + "a")) + "hello")) {
                                        case "hello":
                                            v18 := GenS4{F0: -55, F1: (!(-415 != 32767)), F2: ""}
                                            _ = v18
                                        case "":
                                            v19 := GenS4{F0: -1, F1: false, F2: (("" + ("hello" + "hello")) + "hello")}
                                            _ = v19
                                        default:
                                            delete(v2, 9)
                                        }
                                        _ = v17
                                    }
                                    v20 := v2[42]
                                    _ = v13
                                    _ = v20
                                } else {
                                    switch fi10 {
                                    case 0:
                                        switch "a" {
                                        case "\n\t":
                                            v21 := GenS5{F0: 1, F1: 1, F2: 42}
                                            _ = v21
                                        case "":
                                            v1 = append(v1, "\n\t")
                                        }
                                    case 709:
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                    }
                                    var v22 uint64 = 230
                                    _ = v22
                                }
                            }
                            v23 := GenS1{F0: 2147483647, F1: "a", F2: 233}
                            v3 -= (-(fi9))
                            _ = v23
                        } else {
                            var v24 int32 = 752
                            v1 = append(v1, "a")
                            var v25 bool = ((float64(1e10) * -2.25) == (((1e10 + float64(0.1)) * float64(float32((float64(0.1) * 1.5)))) / 3.14159))
                            var v26 int16 = 641
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(fi9, v11)
                            case "a":
                                v24 += genFn0(362)
                                v7 = v7
                                v2[uint32(v7)] = 2
                            }
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    case -23:
                        var v27 uint8 = 104
                        _ = v27
                    }
                    v7 += 335
                    _ = fi10
                    _ = v11
                }
            }
            for fi28 := 0; fi28 < 13; fi28++ {
                switch v3 {
                case -9223372036854775808:
                    var v29 bool = (42 >= 42)
                    _ = v29
                case 9223372036854775807:
                    if len(v1) > 0 {
                        v1[0] = (("" + "a") + (((("foo bar" + "\n\t") + ((("\n\t" + ("\n\t" + "hello")) + "a") + "\n\t")) + "\n\t") + ("" + ("\n\t" + "a"))))
                    }
                case 0:
                    if (!true) {
                        var v30 uint8 = 27
                        _ = v30
                    }
                default:
                    var v31 int32 = 286
                    _ = v31
                }
                var v32 int64 = 496
                var v33 bool = ((0.1 < float32(((1.5 - (0.1 / 3.14159)) * -2.25))) && false)
                _ = fi28
                _ = v32
                _ = v33
            }
            _ = fi9
        }
        _ = v8
    }
    var v34 float32 = (((0.1 * 0.1) / (3.14159 / 1e10)) * (0.0 + (float32((1.5 - float64(0.0))) / float32(1.5))))
    v35 := GenS1{F0: 2147483647, F1: "a", F2: 0}
    v34 = 3.14159
    var v36 float64 = (float64((0.0 / 0.1)) - ((3.14159 * (3.14159 - 1e10)) / 1e10))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v6
    _ = v7
}
