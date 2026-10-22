package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 string }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint32; F1 int16; F2 int32 }
type GenS6 struct { F0 string; F1 uint; F2 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 int32; F2 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "hello"
    var v3 string = v2
    v2 = v3
    v2 = v2
    fmt.Println(v2, v3)
    v2 = v2
    switch "\n\t" {
    case "hello":
        if (!true) {
            v3 = v2
            if (-95 >= -128) {
                fmt.Println(v2, v3)
                for fi4 := 0; fi4 < 16; fi4++ {
                    var v5 uint64 = 2
                    _ = fi4
                    _ = v5
                }
            }
            var v6 bool = (!(42 >= -1))
            _ = v6
        } else {
            v3 = ("a" + v3)
            for fi7 := 0; fi7 < 1; fi7++ {
                v3 = v2
                if (!(-758 >= -1)) {
                    var v8 uint16 = 255
                    if (387 > 18446744073709551615) {
                        v2 = v2
                    } else {
                        for fi9 := 0; fi9 < 8; fi9++ {
                            var v10 int = -516
                            v11 := GenS1{F0: -11, F1: "\n\t", F2: "\n\t"}
                            v8 += (v8 & 2)
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                        if (!true) {
                            v8 += 1
                        }
                    }
                    _ = len(v1)
                    _ = v8
                }
                for fi12 := 0; fi12 < 9; fi12++ {
                    if false {
                        v2 = v3
                        v2 = (("foo bar" + "\n\t") + ("\n\t" + "\n\t"))
                    }
                    v13 := GenS2{F0: (2 > 255), F1: v3}
                    fmt.Println(fi12, v2, fi7)
                    _ = fi12
                    _ = v13
                }
                _ = fi7
            }
            v14 := GenS1{F0: 1, F1: v2, F2: (((("hello" + ("\n\t" + ("a" + v3))) + ("\n\t" + (v2 + (v2 + (v3 + ""))))) + ("foo bar" + v2)) + v3)}
            _ = v14
        }
    case "\n\t":
        delete(v1, 797)
    default:
        v15 := GenS3{F0: 42}
        _ = v15
    }
    fmt.Println(v3, v2)
    switch ("foo bar" + v2) {
    case "foo bar":
        var v16 bool = true
        _ = v16
    case "hello":
        switch 9223372036854775807 {
        case 1:
            switch v2 {
            case "foo bar":
                v17 := GenS3{F0: 1000}
                _ = v17
            default:
                for fi18 := 0; fi18 < 13; fi18++ {
                    v19 := GenS1{F0: -128, F1: v2, F2: ""}
                    _ = fi18
                    _ = v19
                }
            }
        default:
            v20 := []int{9223372036854775807, -1, 127}
            _ = v20
        }
    case "a":
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 uint32 = 2147483648
            _ = fi21
            _ = v22
        }
    default:
        v23 := GenS4{F0: 896, F1: 1, F2: 32767}
        _ = v23
    }
    fmt.Println(v3)
    for fi24 := 0; fi24 < 13; fi24++ {
        v2 = "\n\t"
        var v25 int32 = 1
        if false {
            if (true || (-1 >= 157)) {
                if (402 > 65535) {
                    v3 = (v3 + v3)
                } else {
                    if false {
                        v26 := GenS5{F0: 2147483648, F1: 32767, F2: v25}
                        fmt.Println(v3)
                        _ = v26
                    } else {
                        var v27 int32 = (v25 - (v25 / ((42) | 1)))
                        for fi28 := 0; fi28 < 18; fi28++ {
                            if false {
                                if false {
                                    fmt.Println(fi28)
                                }
                            } else {
                                v27 = v27
                                v29 := GenS6{F0: "foo bar", F1: 65535, F2: 1}
                                v30 := []int{}
                                if len(v30) > 0 {
                                    v30[0] = fi28
                                }
                                switch -1 {
                                case 127:
                                    if len(v30) >= 2 {
                                        v30 = v30[1:len(v30)]
                                    }
                                case 32767:
                                    v31 := GenS5{F0: 0, F1: 0, F2: -610}
                                    fmt.Println(v3, fi24)
                                    _ = v31
                                case -2147483648:
                                    var v32 uint16 = 643
                                    _ = v32
                                }
                                var v33 uint32 = 0
                                _ = v29
                                _ = v30
                                _ = v33
                            }
                            var v34 float32 = 1e10
                            _ = fi28
                            _ = v34
                        }
                        v35 := GenS6{F0: "", F1: 0, F2: -1}
                        _ = v27
                        _ = v35
                    }
                }
                v36 := GenS3{F0: 0}
                _ = v36
            } else {
                v37 := GenS7{F0: 1}
                _ = v37
            }
        }
        _ = fi24
        _ = v25
    }
    if (!false) {
        var v38 string = v3
        switch -429 {
        case -686:
            v38 = "foo bar"
        case -21:
            v39 := GenS7{F0: 244}
            _ = v39
        default:
            _ = len(v1)
        }
        _ = v38
    } else {
        v40 := GenS8{F0: "", F1: 42, F2: (42 == -1)}
        if false {
            v41 := GenS4{F0: 2147483647, F1: 128, F2: 1000}
            switch 2147483647 {
            case 32768:
                v3 = v3
            default:
                var v42 int32 = 528
                _ = v42
            }
            var v43 uint32 = 164
            _ = v41
            _ = v43
        }
        _ = v40
    }
    var v44 float64 = (float64(((0.0 + ((1.5 - ((0.1 + (1.5 / 3.14159)) / (-2.25 + 0.1))) * ((3.14159 - float32(((((1e10 * float64(0.0)) - 3.14159) * 1.5) - (1.5 * float64(0.1))))) * 3.14159))) * 0.1)) + 1.5)
    delete(v1, 32768)
    v45 := GenS3{F0: 32767}
    var v46 int64 = -453
    v47 := []int{766, 32767, 175}
    v2 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
