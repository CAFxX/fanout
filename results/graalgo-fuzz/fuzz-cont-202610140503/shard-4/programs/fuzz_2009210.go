package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS3 struct { F0 string; F1 string; F2 int32 }
type GenS4 struct { F0 int; F1 bool; F2 int32 }

func main() {
    var v1 uint64 = 4294967295
    var v2 int16 = 32767
    var v3 uint32 = 1
    var v4 float64 = float64((0.1 + 1e10))
    var v5 uint32 = 65535
    for fi6 := 0; fi6 < 10; fi6++ {
        v1 += v1
        _ = fi6
    }
    fmt.Println(v2, v4)
    v5 = v5
    var v7 float32 = (1e10 + float32(-2.25))
    v8 := []int{-1, 892, -788}
    v7 -= v7
    v4 = ((v4 / (v4 + 0.0)) / float64(v7))
    if len(v8) > 0 {
        v8[0] = 0
    }
    if true {
        if (0 >= -29) {
            v9 := GenS1{F0: 2}
            if false {
                var v10 uint = 1
                var v11 uint8 = 181
                switch "" {
                case "":
                    var v12 int32 = 2147483647
                    _ = v12
                case "foo bar":
                    for fi13 := 0; fi13 < 8; fi13++ {
                        v1 -= v1
                        var v14 int8 = 0
                        if (!(!(v1 == (v1 >> ((v10) & 63))))) {
                            var v15 uint = 65535
                            var v16 uint64 = 679
                            _ = v15
                            _ = v16
                        }
                        _ = fi13
                        _ = v14
                    }
                case "hello":
                    if false {
                        var v17 float32 = 1e10
                        for fi18 := 0; fi18 < 4; fi18++ {
                            v10 = (v10 >> ((v10) & 63))
                            _ = fi18
                        }
                        _ = v17
                    } else {
                        v3 = v3
                        v8 = append(v8, 2147483647)
                        var v19 string = "hello"
                        _ = v19
                    }
                }
                _ = v10
                _ = v11
            } else {
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 uint64 = v1
                    switch fi20 {
                    case 4294967295:
                        if len(v8) > 0 {
                            v8[0] = (^(fi20))
                        }
                    case 9223372036854775807:
                        if true {
                            v21 = v21
                            v4 -= 0.1
                            if false {
                                var v22 uint64 = v1
                                _ = v22
                            }
                        } else {
                            switch "\n\t" {
                            case "":
                                for fi23 := 0; fi23 < 3; fi23++ {
                                    if (((^(v3)) < (v5 | (v3 / ((4294967295) | 1)))) && false) {
                                        if (float64(1e10) < -2.25) {
                                            v8 = append(v8, (fi23 | -1))
                                            v24 := GenS2{F0: 65535, F1: 1, F2: 4294967295}
                                            _ = v24
                                        }
                                        if ((^(v1)) < (v1 % ((v21) | 1))) {
                                            var v25 bool = true
                                            if v25 {
                                                v21 = (v21 / (((v21 + 560)) | 1))
                                                switch "" {
                                                case "foo bar":
                                                    var v26 float32 = 1.5
                                                    _ = v26
                                                default:
                                                    _ = len(v8)
                                                    var v27 uint64 = v1
                                                    v4 = v4
                                                    v28 := GenS1{F0: 2}
                                                    if (!(false && ((float64(3.14159) <= 1.5) && (!v25)))) {
                                                        v7 = v7
                                                        v29 := GenS2{F0: 1, F1: -128, F2: v21}
                                                        v25 = (((fi23 | -17) >= (fi23 | fi20)) && (42 <= 0))
                                                        _ = v29
                                                    }
                                                    var v30 float64 = 0.1
                                                    if true {
                                                        v27 = 2
                                                        if (1 > int64(v2)) {
                                                            if (-1 == 1) {
                                                                var v31 uint = 70
                                                                _ = v31
                                                            }
                                                            var v32 uint8 = 135
                                                            _ = v32
                                                        }
                                                        if v25 {
                                                            var v33 int16 = -32768
                                                            _ = v33
                                                        } else {
                                                            fmt.Println(v21)
                                                            var v34 int64 = -9223372036854775808
                                                            v35 := GenS1{F0: 931}
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                        _ = len(v8)
                                                    } else {
                                                        switch ((((("hello" + "") + "") + "\n\t") + ("" + "\n\t")) + "") {
                                                        case "foo bar":
                                                            var v36 uint64 = 0
                                                            _ = v36
                                                        case "\n\t":
                                                            v37 := GenS3{F0: "a", F1: "a", F2: -893}
                                                            _ = v37
                                                        default:
                                                            if (42 < -39) {
                                                                _ = len(v8)
                                                                var v38 int8 = 127
                                                                _ = v38
                                                            } else {
                                                                v39 := GenS4{F0: fi20, F1: (v4 < v4), F2: 2147483647}
                                                                v4 -= (v4 / v4)
                                                                v2 -= v2
                                                                _ = v39
                                                            }
                                                        }
                                                        v40 := GenS1{F0: 0}
                                                        _ = v40
                                                    }
                                                    var v41 float32 = (1.5 / v7)
                                                    switch (fi20 >> ((914) & 63)) {
                                                    case 769:
                                                        v2 += v2
                                                    case 32767:
                                                        v1 -= v21
                                                    default:
                                                        v30 = v4
                                                    }
                                                    var v42 int64 = 1
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                    _ = v27
                                                    _ = v28
                                                    _ = v30
                                                    _ = v41
                                                    _ = v42
                                                }
                                                v3 = v3
                                            }
                                            v8 = append(v8, fi20)
                                            _ = v25
                                        }
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                    }
                                    _ = fi23
                                }
                            case "a":
                                var v43 uint16 = 65535
                                _ = v43
                            case "hello":
                                v44 := GenS4{F0: (fi20 >> ((1) & 63)), F1: ((!true) || true), F2: -759}
                                _ = v44
                            default:
                                var v45 bool = true
                                _ = v45
                            }
                            if len(v8) > 0 {
                                v8[0] = (fi20 | fi20)
                            }
                        }
                    case 127:
                        for fi46 := 0; fi46 < 10; fi46++ {
                            v5 -= (v3 * v3)
                            _ = fi46
                        }
                    default:
                        var v47 uint8 = 224
                        _ = v47
                    }
                    if (v7 < v7) {
                        v21 = v1
                    }
                    _ = fi20
                    _ = v21
                }
            }
            v3 = v5
            _ = v9
        }
    } else {
        var v48 string = (((("\n\t" + "") + "\n\t") + ((((("a" + "foo bar") + "\n\t") + "hello") + "foo bar") + "a")) + "foo bar")
        switch 42 {
        case -2147483648:
            v49 := GenS3{F0: v48, F1: "\n\t", F2: -1}
            _ = v49
        case 127:
            var v50 uint32 = v3
            _ = v50
        default:
            for fi51 := 0; fi51 < 5; fi51++ {
                var v52 uint16 = 65535
                _ = fi51
                _ = v52
            }
        }
        _ = v48
    }
    v3 = (v3 - (v3 / ((2) | 1)))
    var v53 float32 = v7
    for fi54 := 0; fi54 < 2; fi54++ {
        v3 = (v5 | 65535)
        v4 = 0.1
        _ = fi54
    }
    fmt.Println(v7, v5)
    var v55 int64 = -1
    var v56 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v55
    _ = v56
    _ = v7
    _ = v8
}
