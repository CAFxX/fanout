package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int64; F2 int16 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int }

func main() {
    var v1 float64 = (1e10 * 1.5)
    v2 := []int{-169, 551}
    v3 := make(map[string]int)
    var v4 int8 = 100
    switch "hello" {
    case "foo bar":
        if false {
            v5 := GenS1{F0: 315}
            _ = v5
        }
    case "\n\t":
        v4 += (v4 * v4)
    case "hello":
        v2 = append(v2, 32768)
    }
    v4 += (v4 * 100)
    v6 := GenS1{F0: 42}
    v1 += v1
    v1 = 3.14159
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint16 = 631
        v4 = 92
        switch (("foo bar" + ("" + "")) + "") {
        case "\n\t":
            var v9 int16 = 32767
            _ = v9
        case "foo bar":
            for fi10 := 0; fi10 < 5; fi10++ {
                if false {
                    var v11 uint8 = 232
                    var v12 int16 = 1
                    v1 = 1e10
                    _ = v11
                    _ = v12
                } else {
                    fmt.Println(v8, v4, v4)
                    var v13 int8 = 127
                    v14 := GenS1{F0: 1000}
                    var v15 float64 = 1.5
                    _ = v13
                    _ = v14
                    _ = v15
                }
                if (!false) {
                    v16 := GenS2{F0: -9223372036854775808, F1: 82, F2: 42}
                    if true {
                        if true {
                            v17 := GenS2{F0: 1, F1: 98, F2: 1000}
                            v4 = (v4 - 4)
                            v4 -= (v4 ^ v4)
                            delete(v3, "a")
                            _ = v17
                        }
                        var v18 float64 = 1e10
                        v8 += 42
                        _ = v18
                    }
                    if (!false) {
                        var v19 float64 = ((0.1 + (((1e10 * (-2.25 / v1)) / 0.1) - v1)) / v1)
                        var v20 int32 = -2147483648
                        var v21 uint = 42
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    _ = v16
                } else {
                    if false {
                        v8 += (v8 >> ((42) & 63))
                    } else {
                        if true {
                            if true {
                                v22 := GenS3{F0: -1, F1: 794}
                                var v23 float64 = v1
                                _ = v22
                                _ = v23
                            }
                            var v24 float32 = float32(0.0)
                            _ = v24
                        }
                        _ = len(v2)
                    }
                    var v25 int = (fi7 / ((fi10) | 1))
                    _ = v25
                }
                var v26 uint8 = 225
                _ = fi10
                _ = v26
            }
        case "a":
            for fi27 := 0; fi27 < 16; fi27++ {
                v4 -= -109
                if true {
                    var v28 int16 = 32767
                    _ = v28
                } else {
                    if (v8 != v8) {
                        v1 += v1
                        if len(v2) > 0 {
                            v2[0] = (fi27 ^ (fi27 * fi7))
                        }
                    } else {
                        if true {
                            v29 := GenS4{F0: 2, F1: 0}
                            var v30 uint32 = 0
                            _ = v29
                            _ = v30
                        } else {
                            if (true && true) {
                                v8 -= (v8 & (v8 + (v8 | (v8 ^ v8))))
                                fmt.Println(v4, fi27, v4)
                                if false {
                                    var v31 string = "a"
                                    v4 = (v4 ^ (v4 * (v4 ^ v4)))
                                    if (128 != 42) {
                                        v1 = v1
                                    } else {
                                        if ((v31 + v31) != ((((v31 + v31) + "hello") + v31) + v31)) {
                                            fmt.Println(v4, v4)
                                            var v32 int32 = -606
                                            _ = v32
                                        } else {
                                            v4 = v4
                                            v31 = v31
                                            var v33 int64 = 9223372036854775807
                                            _ = v33
                                        }
                                        var v34 int16 = 32767
                                        _ = v34
                                    }
                                    _ = v31
                                } else {
                                    if (("hello" + "") == "hello") {
                                        var v35 uint8 = 77
                                        _ = v35
                                    }
                                    v1 += v1
                                    var v36 int8 = (-(v4))
                                    _ = v36
                                }
                            } else {
                                if (v4 > 100) {
                                    v37 := GenS5{F0: 0}
                                    v38 := GenS2{F0: 42, F1: -927, F2: 79}
                                    _ = v37
                                    _ = v38
                                }
                            }
                        }
                        var v39 float64 = v1
                        v40 := GenS1{F0: int16(v8)}
                        _ = v39
                        _ = v40
                    }
                    switch ("a" + "\n\t") {
                    case "":
                        v3["foo bar"] = (-(fi27))
                    case "a":
                        v8 -= (^(v8))
                    case "foo bar":
                        fmt.Println(v1)
                    default:
                        var v41 bool = ("a" == "")
                        _ = v41
                    }
                    fmt.Println(fi7, v1, fi27)
                }
                _ = fi27
            }
        }
        _ = fi7
        _ = v8
    }
    switch 32767 {
    case 32768:
        v4 = (v4 / (((v4 | 0)) | 1))
    }
    var v42 int16 = 0
    if true {
        _ = len(v3)
    } else {
        v43 := GenS6{F0: -261}
        _ = v43
    }
    for fi44 := 0; fi44 < 1; fi44++ {
        for fi45 := 0; fi45 < 9; fi45++ {
            if true {
                switch (fi44 & 9223372036854775807) {
                case 32767:
                    var v46 uint64 = 0
                    _ = v46
                case 9223372036854775807:
                    fmt.Println(fi44)
                case -1:
                    fmt.Println(v42, fi45, fi45)
                default:
                    var v47 float64 = float64((float32((v1 / (v1 / v1))) / 1e10))
                    _ = v47
                }
                var v48 uint8 = 255
                _ = v48
            }
            v4 = v4
            var v49 float32 = float32(v1)
            var v50 int16 = 32767
            var v51 uint64 = 18446744073709551615
            _ = fi45
            _ = v49
            _ = v50
            _ = v51
        }
        v52 := GenS4{F0: 920, F1: 0}
        _ = fi44
        _ = v52
    }
    v3["\n\t"] = -9223372036854775808
    fmt.Println(v42)
    for fi53 := 0; fi53 < 17; fi53++ {
        v1 = (v1 / (-2.25 + v1))
        for fi54 := 0; fi54 < 17; fi54++ {
            switch fi53 {
            case 42:
                var v55 int = fi54
                _ = v55
            default:
                v1 = -2.25
            }
            _ = fi54
        }
        v1 -= v1
        _ = fi53
    }
    fmt.Println(v4, v1, v4)
    if (-9223372036854775808 != -9223372036854775808) {
        v4 = (v4 * v4)
        v42 = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v6
}
