package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 758
    var v3 int8 = 127
    var v4 uint32 = 419
    v2 += v2
    switch "\n\t" {
    case "hello":
        v2 = uint(v4)
    default:
        var v5 int32 = -80
        _ = v5
    }
    v3 -= (v3 | 0)
    if false {
        v2 = (v2 * v2)
    }
    v6 := GenS1{F0: ((!(-55 != 127)) || (true || false))}
    v4 -= v4
    _ = len(v1)
    fmt.Println(v2)
    if (0 < -78) {
        var v7 float32 = -2.25
        v8 := []int{925, -550, 127, 1}
        fmt.Println(v2)
        _ = v7
        _ = v8
    } else {
        switch 300 {
        case -9223372036854775808:
            if ((1e10 - ((3.14159 / (float32(1e10) - ((3.14159 - 3.14159) / float32((1e10 / (-2.25 * (float64(3.14159) * float64((1e10 + 1.5))))))))) * 0.1)) >= (((0.1 - (1e10 * ((float32(0.1) / 1.0) / 1.5))) - (1e10 / (1e10 + (1.5 - float32(0.0))))) - (float32(0.1) + -2.25))) {
                var v9 uint8 = 42
                v3 = (v3 / (((v3 ^ v3)) | 1))
                switch -186 {
                case 128:
                    fmt.Println(v4)
                case -266:
                    fmt.Println(v9, v2, v3)
                case -769:
                    for fi10 := 0; fi10 < 16; fi10++ {
                        var v11 int = 0
                        _ = fi10
                        _ = v11
                    }
                default:
                    v12 := []int{-2147483648, 231}
                    _ = v12
                }
                _ = v9
            }
        case 176:
            if false {
                switch "foo bar" {
                case "":
                    if false {
                        v4 = 821
                        var v13 int32 = 0
                        _ = v13
                    } else {
                        v3 += 1
                        if true {
                            if true {
                                v4 += v4
                            } else {
                                v14 := []int{0, 42, 154}
                                for fi15 := 0; fi15 < 20; fi15++ {
                                    v16 := GenS2{F0: "", F1: 0}
                                    v1[("hello" + "hello")] = fi15
                                    v4 = (v4 << ((v2) & 63))
                                    _ = fi15
                                    _ = v16
                                }
                                _ = v14
                            }
                        } else {
                            if true {
                                var v17 uint64 = 673
                                _ = v17
                            } else {
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    for fi19 := 0; fi19 < 4; fi19++ {
                                        v20 := GenS1{F0: (4294967295 <= (v2 ^ 4294967295))}
                                        if true {
                                            v21 := GenS2{F0: "hello", F1: 1}
                                            var v22 string = "hello"
                                            _ = v21
                                            _ = v22
                                        } else {
                                            var v23 uint8 = uint8(v2)
                                            v24 := []int{(fi18 + (fi18 & -714)), fi18}
                                            _ = v23
                                            _ = v24
                                        }
                                        _ = fi19
                                        _ = v20
                                    }
                                    _ = fi18
                                }
                                var v25 int16 = 1
                                v25 = v25
                                _ = v25
                            }
                        }
                    }
                case "foo bar":
                    for fi26 := 0; fi26 < 14; fi26++ {
                        if false {
                            var v27 int = 2147483647
                            _ = v27
                        } else {
                            for fi28 := 0; fi28 < 17; fi28++ {
                                fmt.Println(fi28, v2, v4)
                                _ = fi28
                            }
                            var v29 int8 = (v3 / ((v3) | 1))
                            _ = v29
                        }
                        var v30 uint32 = v4
                        var v31 uint64 = 131
                        _ = fi26
                        _ = v30
                        _ = v31
                    }
                case "a":
                    v4 -= v4
                }
                v32 := GenS3{F0: (v4 + v4)}
                v33 := GenS3{F0: v4}
                _ = v32
                _ = v33
            }
        case 2147483647:
            v34 := GenS4{F0: true}
            _ = v34
        }
        v1["foo bar"] = 858
        switch 848 {
        case 32768:
            for fi35 := 0; fi35 < 13; fi35++ {
                var v36 uint64 = 4294967295
                v2 = 18446744073709551615
                _ = fi35
                _ = v36
            }
        case 32767:
            var v37 int8 = -67
            _ = v37
        case 602:
            var v38 int = -660
            _ = v38
        default:
            v3 -= (-(v3))
        }
    }
    for fi39 := 0; fi39 < 6; fi39++ {
        v4 -= uint32(v2)
        if ((v3 * 123) == (v3 - v3)) {
            var v40 uint = (v2 - (v2 << ((v2) & 63)))
            _ = v40
        }
        _ = fi39
    }
    v41 := []int{}
    v42 := GenS5{F0: (v2 + uint(v4)), F1: 1}
    if len(v41) > 0 {
        v41[0] = 128
    }
    _ = len(v41)
    v43 := GenS2{F0: "", F1: 392}
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v6
}
