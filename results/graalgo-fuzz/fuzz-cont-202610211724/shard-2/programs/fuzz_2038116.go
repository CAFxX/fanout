package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    var v3 uint32 = 2147483647
    v4 := []int{-2147483648, 523, -9223372036854775808, 2147483647}
    var v5 uint = 1
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            v5 = 383
            v2[973] = "hello"
            _ = fi8
        }
        _ = fi7
    }
    v9 := GenS1{F0: false, F1: 0}
    fmt.Println(v3, v5)
    v10 := GenS2{F0: 168}
    _ = len(v6)
    fmt.Println(v3, v3)
    var v11 uint16 = uint16(v3)
    v12 := GenS3{F0: 1, F1: 4294967295}
    for fi13 := 0; fi13 < 19; fi13++ {
        v11 += v11
        _ = fi13
    }
    var v14 bool = (!(((v3 + v3) > (v3 & v3)) && false))
    v3 = uint32(v11)
    if (-2.25 >= 1e10) {
        if (("foo bar" + "a") == ("foo bar" + (("hello" + "") + "a"))) {
            var v15 int32 = 1
            v16 := GenS3{F0: int32(v5), F1: (^(v3))}
            if (!(!v14)) {
                v5 -= v5
                v17 := GenS2{F0: 0}
                _ = v17
            } else {
                if (1 != v5) {
                    v3 += 1
                }
                switch 434 {
                case 2147483647:
                    var v18 uint16 = v11
                    _ = v18
                case 128:
                    var v19 uint32 = (v3 / (((v3 & 65535)) | 1))
                    _ = v19
                case -1:
                    for fi20 := 0; fi20 < 20; fi20++ {
                        for fi21 := 0; fi21 < 1; fi21++ {
                            var v22 uint64 = 18446744073709551615
                            var v23 uint8 = 128
                            var v24 uint8 = v23
                            _ = fi21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        fmt.Println(v3, v3)
                        _ = fi20
                    }
                default:
                    if true {
                        for fi25 := 0; fi25 < 10; fi25++ {
                            v26 := GenS4{F0: -444, F1: 172, F2: v11}
                            _ = fi25
                            _ = v26
                        }
                        fmt.Println(v15, v14, v14)
                        var v27 bool = (32768 == 128)
                        _ = v27
                    } else {
                        v11 = v11
                        v28 := GenS2{F0: 693}
                        _ = v28
                    }
                }
            }
            _ = v15
            _ = v16
        } else {
            v29 := GenS1{F0: (-114 < 0), F1: 42}
            _ = v29
        }
        if v14 {
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 float32 = 0.1
                _ = fi30
                _ = v31
            }
            switch "hello" {
            case "foo bar":
                if v14 {
                    _ = len(v4)
                    if len(v4) > 0 {
                        v4[0] = 42
                    }
                } else {
                    fmt.Println(v5, v11, v3)
                    var v32 float32 = (-2.25 + 0.1)
                    _ = v32
                }
            case "hello":
                fmt.Println(v5, v5, v14)
            }
        } else {
            switch "" {
            case "a":
                v11 += 986
            case "\n\t":
                switch -1 {
                case 2147483647:
                    if false {
                        switch (("" + "a") + ("\n\t" + "hello")) {
                        case "foo bar":
                            if true {
                                switch 128 {
                                case -76:
                                    fmt.Println(v3, v3, v14)
                                default:
                                    var v33 float32 = (0.0 * 0.1)
                                    _ = v33
                                }
                            }
                        case "a":
                            v11 = (v11 | (v11 / ((785) | 1)))
                        default:
                            if v14 {
                                fmt.Println(v5)
                                if ("hello" > "\n\t") {
                                    delete(v6, -949)
                                    fmt.Println(v11, v11, v5)
                                    switch 0 {
                                    case 9223372036854775807:
                                        if (42 <= -625) {
                                            v34 := GenS3{F0: 1, F1: (^(v3))}
                                            for fi35 := 0; fi35 < 13; fi35++ {
                                                if v14 {
                                                    v1[""] = (!(!(!v14)))
                                                } else {
                                                    v5 += v5
                                                    if v14 {
                                                        v36 := GenS3{F0: -1, F1: v3}
                                                        _ = v36
                                                    } else {
                                                        var v37 int = 4294967295
                                                        v14 = v14
                                                        _ = v37
                                                    }
                                                }
                                                _ = fi35
                                            }
                                            _ = v34
                                        }
                                    }
                                } else {
                                    for fi38 := 0; fi38 < 9; fi38++ {
                                        v4 = append(v4, fi38)
                                        v5 += (v5 + 2)
                                        v39 := GenS5{F0: 255}
                                        _ = fi38
                                        _ = v39
                                    }
                                }
                            }
                        }
                        _ = len(v4)
                    }
                default:
                    if v14 {
                        for fi40 := 0; fi40 < 16; fi40++ {
                            var v41 float32 = (0.0 / (0.1 + (-2.25 + (3.14159 * 3.14159))))
                            if (!((v41 / ((-2.25 + (0.0 + v41)) - v41)) != 0.0)) {
                                if v14 {
                                    fmt.Println(v11)
                                } else {
                                    v5 = v5
                                    v5 = v5
                                    v5 -= (^(v5))
                                }
                            }
                            var v42 int16 = -32768
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                        fmt.Println(v3, v5, v11)
                        v43 := GenS2{F0: 42}
                        _ = v43
                    }
                }
            }
        }
    }
    if len(v4) > 0 {
        v4[0] = 128
    }
    v11 = v11
    var v44 int = -9223372036854775808
    if v14 {
        v45 := v2[uint32(v5)]
        var v46 int32 = 42
        _ = v45
        _ = v46
    }
    var v47 uint32 = (v3 % ((v3) | 1))
    var v48 bool = v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v9
}
