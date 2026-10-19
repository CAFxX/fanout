package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 bool }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int }

func main() {
    v1 := []int{1, 0, 937}
    v2 := []uint32{}
    v3 := make(map[int]string)
    var v4 bool = false
    _ = len(v3)
    if (1 >= 1) {
        v4 = (1 > -723)
        if (4294967295 >= 26) {
            v5 := GenS1{F0: 771, F1: 2147483648, F2: (false || (884 == 108))}
            if false {
                switch 32768 {
                case 32768:
                    _ = len(v2)
                case 32767:
                    for fi6 := 0; fi6 < 11; fi6++ {
                        var v7 float64 = 0.1
                        _ = fi6
                        _ = v7
                    }
                case -1:
                    v4 = (!((2147483647 > 0) || false))
                default:
                    delete(v3, 127)
                }
                fmt.Println(v4)
            } else {
                v4 = ((("hello" + (("a" + "hello") + ("a" + "a"))) + "a") == "a")
                for fi8 := 0; fi8 < 14; fi8++ {
                    if false {
                        switch ("a" + ("foo bar" + "a")) {
                        case "foo bar":
                            v4 = v4
                        case "":
                            for fi9 := 0; fi9 < 12; fi9++ {
                                if true {
                                    v4 = ((42 >= -769) && (127 >= 42))
                                    if len(v2) > 0 {
                                        v2[0] = 2147483647
                                    }
                                } else {
                                    if false {
                                        if v4 {
                                            var v10 int64 = 1
                                            _ = v10
                                        } else {
                                            v11 := GenS2{F0: 2, F1: 749}
                                            _ = v11
                                        }
                                        var v12 uint64 = 302
                                        _ = v12
                                    }
                                    var v13 int32 = -779
                                    _ = v13
                                }
                                switch fi9 {
                                case 9223372036854775807:
                                    switch 4 {
                                    case 32768:
                                        if (0 >= 957) {
                                            v14 := GenS3{F0: 4294967295}
                                            switch fi8 {
                                            case 32767:
                                                v15 := GenS4{F0: 42, F1: (fi8 / (((fi9 - (fi9 & -2147483648))) | 1))}
                                                _ = v15
                                            case -90:
                                                v4 = (true && (!v4))
                                            case 3:
                                                v16 := GenS4{F0: 255, F1: fi9}
                                                _ = v16
                                            }
                                            var v17 float32 = -2.25
                                            _ = v14
                                            _ = v17
                                        }
                                    }
                                case 2147483647:
                                    var v18 uint64 = 42
                                    _ = v18
                                case -231:
                                    if v4 {
                                        if (4294967295 < 14) {
                                            var v19 uint = 18446744073709551615
                                            _ = v19
                                        }
                                        switch -1 {
                                        case -933:
                                            v4 = ((0 != 91) && (784 <= 554))
                                        case -9223372036854775808:
                                            var v20 string = "a"
                                            _ = v20
                                        }
                                    }
                                default:
                                    switch fi8 {
                                    case -635:
                                        var v21 float32 = 0.0
                                        _ = v21
                                    case 792:
                                        v22 := GenS4{F0: 32767, F1: (fi9 | 2147483647)}
                                        switch 42 {
                                        case -568:
                                            v4 = v4
                                        }
                                        _ = v22
                                    case -9223372036854775808:
                                        var v23 string = ("" + "")
                                        _ = v23
                                    default:
                                        v24 := GenS1{F0: 544, F1: 224, F2: (((0 < 2147483648) || v4) || true)}
                                        _ = v24
                                    }
                                }
                                var v25 int8 = 48
                                _ = fi9
                                _ = v25
                            }
                        case "hello":
                            for fi26 := 0; fi26 < 18; fi26++ {
                                var v27 string = ""
                                v28 := GenS3{F0: 0}
                                v1 = append(v1, fi26)
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        default:
                            fmt.Println(fi8)
                        }
                        v4 = v4
                    } else {
                        v29 := GenS4{F0: 0, F1: fi8}
                        v4 = v4
                        _ = v29
                    }
                    _ = fi8
                }
                var v30 int16 = 200
                _ = v30
            }
            var v31 float64 = 0.1
            _ = v31
            _ = v5
        } else {
            var v32 string = "hello"
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        if (false || (!(!(-128 < 1)))) {
            var v34 string = "a"
            if len(v1) > 0 {
                v1[0] = (^(fi33))
            }
            var v35 float64 = 3.14159
            _ = v34
            _ = v35
        }
        _ = fi33
    }
    v4 = v4
    var v36 int64 = -66
    var v37 string = ""
    var v38 bool = v4
    var v39 uint64 = 769
    v4 = v38
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
