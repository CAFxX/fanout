package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 593
    var v3 float32 = 1e10
    var v4 int8 = 18
    v5 := make(map[string]int)
    var v6 float64 = (3.14159 + 1e10)
    v6 += (float64((v3 + v3)) - v6)
    v7 := []int{127, -753, int(v2)}
    if (0 <= 65535) {
        if true {
            for fi8 := 0; fi8 < 7; fi8++ {
                for fi9 := 0; fi9 < 13; fi9++ {
                    v4 += v4
                    fmt.Println(v3, v6)
                    v2 = v2
                    _ = fi9
                }
                for fi10 := 0; fi10 < 4; fi10++ {
                    switch "\n\t" {
                    case "a":
                        _ = len(v7)
                    case "\n\t":
                        if false {
                            var v11 int8 = 10
                            v12 := GenS1{F0: v2, F1: 1, F2: (v11 & 34)}
                            v1[""] = 720
                            _ = v11
                            _ = v12
                        } else {
                            if false {
                                v13 := GenS2{F0: v4}
                                _ = v13
                            } else {
                                var v14 uint8 = 1
                                var v15 uint8 = (v14 ^ v14)
                                fmt.Println(fi8, v2)
                                _ = v14
                                _ = v15
                            }
                            delete(v1, "a")
                            fmt.Println(v2)
                            v16 := GenS1{F0: (v2 - v2), F1: 127, F2: v4}
                            _ = v16
                        }
                    case "hello":
                        switch "a" {
                        case "\n\t":
                            fmt.Println(fi8, fi8, v2)
                        default:
                            v4 -= (v4 + (v4 * v4))
                        }
                    default:
                        var v17 float32 = 1e10
                        var v18 int16 = -32768
                        v19 := GenS3{F0: 127}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    switch "a" {
                    case "\n\t":
                        v4 -= int8(fi10)
                    case "foo bar":
                        v20 := v5["\n\t"]
                        _ = v20
                    }
                    v3 = float32(3.14159)
                    _ = fi10
                }
                for fi21 := 0; fi21 < 19; fi21++ {
                    v3 -= v3
                    _ = len(v1)
                    v7 = append(v7, 32768)
                    _ = fi21
                }
                _ = fi8
            }
        } else {
            for fi22 := 0; fi22 < 2; fi22++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                v6 = v6
                _ = fi22
            }
            v23 := GenS1{F0: v2, F1: 128, F2: (v4 * (v4 % (((v4 * (v4 * (v4 ^ v4)))) | 1)))}
            _ = v23
        }
        fmt.Println(v4, v6)
        if false {
            fmt.Println(v6, v4, v2)
        } else {
            v24 := v1["hello"]
            switch 268 {
            case -980:
                if false {
                    if true {
                        var v25 uint8 = 127
                        if false {
                            var v26 uint = 439
                            if false {
                                v26 += v26
                                v5[("hello" + "hello")] = 4294967295
                                fmt.Println(v26)
                            } else {
                                v27 := v5["hello"]
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 5; fi28++ {
                                delete(v5, "hello")
                                fmt.Println(v4, v25)
                                _ = fi28
                            }
                            _ = v26
                        }
                        _ = v25
                    }
                    var v29 int64 = -1
                    _ = v29
                } else {
                    switch 9223372036854775807 {
                    case 9223372036854775807:
                        v30 := GenS1{F0: v2, F1: 255, F2: v4}
                        _ = v30
                    }
                }
            case 32768:
                v31 := GenS1{F0: v2, F1: 2, F2: v4}
                _ = v31
            case 2147483647:
                v24 = (v24 & (v24 - v24))
            default:
                var v32 int = 32768
                _ = v32
            }
            _ = v24
        }
    }
    switch "foo bar" {
    case "":
        v2 += v2
    case "foo bar":
        v4 += v4
    case "hello":
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    v4 -= (v4 ^ v4)
    if false {
        var v33 int64 = -366
        _ = v33
    } else {
        var v34 float32 = 1.5
        _ = v34
    }
    v3 += (0.1 + v3)
    switch int(v2) {
    case 585:
        v6 += float64(v3)
    case -788:
        v35 := GenS2{F0: (v4 ^ v4)}
        _ = v35
    default:
        var v36 uint32 = 0
        _ = v36
    }
    v37 := GenS2{F0: (v4 * (v4 & v4))}
    switch "hello" {
    case "":
        v38 := GenS1{F0: v2, F1: 169, F2: (v4 | v4)}
        _ = v38
    case "hello":
        var v39 bool = true
        _ = v39
    case "foo bar":
        v40 := GenS1{F0: v2, F1: 217, F2: (v4 / (((^(v4))) | 1))}
        _ = v40
    default:
        var v41 int32 = 2147483647
        _ = v41
    }
    v42 := GenS1{F0: 474, F1: 1, F2: (v4 ^ v4)}
    v6 += float64(v3)
    v43 := GenS4{F0: -1}
    v7 = append(v7, 32767)
    v4 = v4
    v4 = (v4 ^ v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
