package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint; F1 int32; F2 int16 }
type GenS4 struct { F0 int; F1 int8; F2 int16 }

func main() {
    var v1 bool = (1 < 32768)
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    v4 := []int{}
    v1 = (true && (1 >= 255))
    var v5 float64 = 1.5
    var v6 uint32 = 4294967295
    v7 := GenS1{F0: 255}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v8 := v2[127]
    fmt.Println(v1)
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := v3[v8]
        v11 := GenS1{F0: 2}
        var v12 uint16 = 1
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    if v1 {
        v8 = (v8 + "")
        switch v8 {
        case "foo bar":
            var v13 int8 = 127
            _ = v13
        case "hello":
            v14 := v3[""]
            _ = v14
        default:
            v4 = append(v4, 2147483647)
        }
        var v15 int16 = 660
        _ = v15
    } else {
        v6 = (^(v6))
        var v16 float64 = (3.14159 + (v5 - float64(1.5)))
        _ = v16
    }
    v17 := GenS1{F0: 906}
    var v18 uint32 = (^(v6))
    v3[(v8 + (v8 + v8))] = -1
    switch 128 {
    case -1:
        switch 0 {
        case -29:
            v5 += (v5 / 1.0)
        case 318:
            _ = len(v2)
        case -941:
            if ("hello" >= v8) {
                if (!false) {
                    v19 := GenS1{F0: 1}
                    _ = v19
                }
                for fi20 := 0; fi20 < 15; fi20++ {
                    if (255 <= 0) {
                        for fi21 := 0; fi21 < 9; fi21++ {
                            if false {
                                switch fi20 {
                                case -865:
                                    switch v8 {
                                    case "a":
                                        _ = len(v4)
                                    case "hello":
                                        fmt.Println(v6, v5)
                                    case "":
                                        v8 = v8
                                    default:
                                        var v22 float32 = 1.5
                                        _ = v22
                                    }
                                }
                            } else {
                                if (460 < 1) {
                                    var v23 uint64 = 2
                                    _ = v23
                                } else {
                                    if (-174 <= -9223372036854775808) {
                                        var v24 int64 = 0
                                        v18 += (v6 ^ 571)
                                        v18 += v6
                                        _ = v24
                                    } else {
                                        v18 -= 65535
                                    }
                                }
                                v6 += (^(v6))
                                v25 := GenS2{F0: (v8 + v8), F1: 42}
                                var v26 int32 = 42
                                _ = v25
                                _ = v26
                            }
                            if v1 {
                                if ("" > ((v8 + "hello") + v8)) {
                                    v5 = ((1.5 + (3.14159 + ((0.1 + -2.25) * v5))) / 1.0)
                                    v27 := GenS2{F0: "\n\t", F1: int8(fi21)}
                                    if (!(2 <= 42)) {
                                        if v1 {
                                            _ = len(v2)
                                        } else {
                                            switch v8 {
                                            case "":
                                                var v28 int16 = 32767
                                                _ = v28
                                            case "\n\t":
                                                v29 := GenS1{F0: 0}
                                                _ = v29
                                            }
                                        }
                                        _ = len(v4)
                                        var v30 string = v8
                                        _ = v30
                                    }
                                    _ = v27
                                } else {
                                    if v1 {
                                        if (!false) {
                                            v6 = 431
                                            fmt.Println(fi20, v5)
                                            var v31 uint16 = 42
                                            v32 := GenS3{F0: uint(v31), F1: 0, F2: 1}
                                            _ = v31
                                            _ = v32
                                        } else {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v33 int = fi20
                                            v34 := GenS3{F0: 2, F1: 316, F2: -862}
                                            v35 := GenS4{F0: 9223372036854775807, F1: int8(v33), F2: -332}
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                    } else {
                                        _ = len(v4)
                                        v36 := v3[v8]
                                        _ = v36
                                    }
                                    if v1 {
                                        if (!((false || v1) || (65535 < 2))) {
                                            v1 = (-1 > fi21)
                                        }
                                    }
                                }
                            }
                            v18 += 251
                            _ = fi21
                        }
                        fmt.Println(v5, v6)
                    }
                    _ = fi20
                }
                v4 = append(v4, 435)
            }
        default:
            var v37 string = "hello"
            _ = v37
        }
    case 1:
        _ = len(v4)
    default:
        var v38 uint64 = 845
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
