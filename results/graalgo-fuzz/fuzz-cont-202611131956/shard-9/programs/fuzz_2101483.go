package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint32 }
func genFn0() int16 {
    var v5 int32 = -1
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := []int{(fi6 << ((1) & 63)), (fi6 >> ((255) & 63)), fi6, (fi6 | fi6)}
        if (127 >= 202) {
            switch fi6 {
            case 32767:
                v8 := GenS1{F0: 0, F1: -862, F2: -1}
                _ = v8
            default:
                switch -233 {
                case -565:
                    for fi9 := 0; fi9 < 7; fi9++ {
                        var v10 string = "foo bar"
                        var v11 uint16 = 1
                        if true {
                            switch v10 {
                            case "":
                                if false {
                                    var v12 int8 = -1
                                    v11 += v11
                                    if false {
                                        v13 := GenS2{F0: -556}
                                        if true {
                                            fmt.Println(v12, fi9, v10)
                                        } else {
                                            var v14 int16 = 1000
                                            switch fi6 {
                                            case 9223372036854775807:
                                                v15 := GenS1{F0: 632, F1: -1, F2: -2147483648}
                                                _ = v15
                                            case 32767:
                                                switch (fi6 | 127) {
                                                case 276:
                                                    v11 = v11
                                                default:
                                                    var v16 int16 = v14
                                                    _ = v16
                                                }
                                            case -474:
                                                v11 -= (^(v11))
                                            }
                                            _ = v14
                                        }
                                        v17 := GenS1{F0: 550, F1: 9223372036854775807, F2: int32(v11)}
                                        _ = v13
                                        _ = v17
                                    }
                                    _ = v12
                                }
                            case "\n\t":
                                switch (fi9 & (fi9 * (fi9 ^ fi6))) {
                                case -110:
                                    var v18 int64 = 1
                                    _ = v18
                                default:
                                    fmt.Println(v10)
                                }
                            case "hello":
                                var v19 uint64 = 895
                                _ = v19
                            default:
                                v11 -= v11
                            }
                            if (-32768 < -345) {
                                _ = len(v7)
                            } else {
                                var v20 bool = (-1 == 0)
                                v10 = "foo bar"
                                _ = v20
                            }
                        }
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                case 2147483647:
                    var v21 int32 = -684
                    _ = v21
                case 0:
                    var v22 uint = 18446744073709551615
                    _ = v22
                }
            }
        } else {
            switch 32768 {
            case 42:
                var v23 bool = false
                _ = v23
            }
            switch (("a" + "") + (("hello" + ("a" + "foo bar")) + "foo bar")) {
            case "\n\t":
                v5 = v5
            default:
                var v24 int16 = 42
                _ = v24
            }
        }
        if false {
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 uint16 = 65535
                v7 = append(v7, 4294967295)
                _ = fi25
                _ = v26
            }
        }
        _ = fi6
        _ = v7
    }
    v27 := GenS3{F0: -32768}
    v5 += v5
    v28 := []int{-836, 9223372036854775807, 4294967295, 9223372036854775807}
    _ = v27
    _ = v28
    _ = v5
    return -1
}

func main() {
    var v1 string = "\n\t"
    var v2 uint32 = 2
    v3 := []int{498, 2147483647, 0}
    v4 := make(map[int]string)
    if false {
        if false {
            var v29 uint32 = (v2 | (^(v2)))
            var v30 int32 = 2147483647
            for fi31 := 0; fi31 < 13; fi31++ {
                for fi32 := 0; fi32 < 16; fi32++ {
                    v29 += v2
                    var v33 uint = 65535
                    _ = len(v4)
                    _ = len(v3)
                    fmt.Println(fi31)
                    v30 = (^(v30))
                    v34 := GenS1{F0: 4294967295, F1: 0, F2: (v30 / (((-(v30))) | 1))}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                switch fi31 {
                case 245:
                    switch "hello" {
                    case "a":
                        fmt.Println(v1, fi31)
                    case "hello":
                        var v35 uint8 = 2
                        _ = v35
                    default:
                        v2 = v29
                    }
                case 638:
                    v1 = "\n\t"
                }
                _ = fi31
            }
            _ = v29
            _ = v30
        }
        v36 := GenS1{F0: 4294967295, F1: 0, F2: 42}
        _ = v36
    } else {
        if false {
            switch "" {
            case "foo bar":
                var v37 float64 = (float64(0.1) - (0.1 + 3.14159))
                _ = v37
            default:
                v38 := GenS3{F0: -32768}
                _ = v38
            }
        }
        _ = len(v3)
    }
    var v39 uint8 = 1
    for fi40 := 0; fi40 < 11; fi40++ {
        for fi41 := 0; fi41 < 2; fi41++ {
            v42 := GenS4{F0: 42, F1: 4294967295, F2: 0}
            v2 = v2
            v1 = "hello"
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    v43 := GenS3{F0: genFn0()}
    v1 = "hello"
    switch 42 {
    case 128:
        for fi44 := 0; fi44 < 13; fi44++ {
            var v45 float64 = (3.14159 - float64(1.5))
            switch v1 {
            case "foo bar":
                if len(v3) > 0 {
                    v3[0] = (fi44 % ((fi44) | 1))
                }
            case "\n\t":
                var v46 uint64 = 2
                _ = v46
            default:
                var v47 uint = 1
                _ = v47
            }
            _ = fi44
            _ = v45
        }
    default:
        if (0.0 > -2.25) {
            var v48 bool = ((2 >= 2) && true)
            var v49 uint64 = 2
            _ = v48
            _ = v49
        } else {
            v39 += v39
            var v50 uint = 65535
            _ = v50
        }
    }
    v51 := GenS5{F0: v2}
    var v52 uint8 = v39
    var v53 float32 = (0.0 + 1e10)
    for fi54 := 0; fi54 < 7; fi54++ {
        var v55 uint32 = (^(v2))
        v4[(fi54 % (((fi54 & -44)) | 1))] = "a"
        _ = fi54
        _ = v55
    }
    switch ("\n\t" + "") {
    case "\n\t":
        for fi56 := 0; fi56 < 10; fi56++ {
            var v57 int = -1
            fmt.Println(fi56)
            _ = fi56
            _ = v57
        }
    case "a":
        if false {
            for fi58 := 0; fi58 < 4; fi58++ {
                for fi59 := 0; fi59 < 5; fi59++ {
                    switch 270 {
                    case 199:
                        fmt.Println(fi58, fi58)
                    }
                    v52 = v52
                    _ = fi59
                }
                delete(v4, fi58)
                _ = fi58
            }
            var v60 uint = 1
            _ = v60
        } else {
            delete(v4, 42)
        }
    case "hello":
        v39 = v39
    default:
        var v61 uint16 = 65535
        _ = v61
    }
    v62 := GenS3{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v43
    _ = v51
    _ = v52
    _ = v53
    _ = v62
}
