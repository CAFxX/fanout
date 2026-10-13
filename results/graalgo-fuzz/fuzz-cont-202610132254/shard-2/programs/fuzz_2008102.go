package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint16, p1 int) uint16 {
    v7 := []int{2147483647, (-(p1))}
    v8 := GenS1{F0: 0, F1: "hello"}
    var v9 float32 = (1e10 + 3.14159)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (p0 - (p0 - (p0 >> ((0) & 63))))
}
func genFn1(p0 int, p1 int32, p2 uint16) bool {
    switch "hello" {
    case "foo bar":
        v10 := GenS2{F0: "hello", F1: 127}
        _ = v10
    default:
        v11 := []int{(p0 >> ((891) & 63)), (-(p0))}
        _ = v11
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]bool)
    v3 := []uint32{}
    v4 := make(map[uint32]bool)
    var v5 int32 = -1
    var v6 float32 = (3.14159 + 1.5)
    if ((true && true) && true) {
        _ = len(v4)
        _ = len(v4)
    } else {
        _ = len(v4)
    }
    if (17 < 68) {
        fmt.Println(v1)
    } else {
        v12 := GenS3{F0: v5, F1: 192}
        _ = v12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(v1)
        if (896 >= 0) {
            var v14 uint = 1
            v5 = (v5 % ((v5) | 1))
            _ = v14
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 uint = 1
            if ((1 < 127) || true) {
                v4[65535] = false
            } else {
                fmt.Println(v6, fi15, fi13)
                delete(v2, -950)
                v2[(fi13 + fi15)] = (false || true)
            }
            _ = fi15
            _ = v16
        }
        _ = fi13
    }
    fmt.Println(v6, v1, v6)
    v3 = append(v3, 809)
    var v17 uint32 = 2
    var v18 int16 = -1
    if false {
        _ = len(v3)
    }
    fmt.Println(v5, v1, v17)
    var v19 float32 = 3.14159
    if (!false) {
        var v20 bool = true
        var v21 float32 = v6
        _ = v20
        _ = v21
    } else {
        switch "\n\t" {
        case "a":
            if (4294967295 == 1) {
                switch 2147483647 {
                case -390:
                    var v22 int = 32768
                    _ = v22
                case -726:
                    var v23 uint16 = 255
                    _ = v23
                }
                v1 = 255
                for fi24 := 0; fi24 < 4; fi24++ {
                    switch ("hello" + "") {
                    case "a":
                        v6 = v6
                    case "\n\t":
                        v25 := GenS4{F0: (v17 % (((v17 / ((uint32(v1)) | 1))) | 1))}
                        _ = v25
                    case "foo bar":
                        switch fi24 {
                        case 4294967295:
                            var v26 uint = 2
                            _ = v26
                        case 1:
                            var v27 uint64 = 18446744073709551615
                            _ = v27
                        case 128:
                            if false {
                                v28 := v2[fi24]
                                v1 -= (v1 >> ((255) & 63))
                                v17 = (^(v17))
                                _ = v28
                            }
                        default:
                            for fi29 := 0; fi29 < 12; fi29++ {
                                v30 := v4[v17]
                                if v30 {
                                    switch fi24 {
                                    case 832:
                                        v18 = v18
                                    default:
                                        if v30 {
                                            var v31 bool = v30
                                            _ = v31
                                        }
                                    }
                                }
                                _ = fi29
                                _ = v30
                            }
                        }
                    default:
                        var v32 string = "foo bar"
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 5; fi33++ {
                        if false {
                            if false {
                                fmt.Println(v6, v5)
                            }
                        }
                        _ = fi33
                    }
                    v5 -= -157
                    _ = fi24
                }
            }
        }
        v34 := GenS5{F0: true}
        _ = v34
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 string = "foo bar"
        _ = fi35
        _ = v36
    }
    _ = len(v4)
    v17 += (v17 >> ((18446744073709551615) & 63))
    v37 := GenS1{F0: v5, F1: "hello"}
    if (v5 != 0) {
        switch "foo bar" {
        case "a":
            var v38 uint32 = v17
            _ = v38
        case "\n\t":
            v39 := GenS3{F0: 322, F1: v1}
            _ = v39
        }
        v18 = v18
        switch 32767 {
        case 128:
            var v40 bool = (0.0 >= 0.1)
            _ = v40
        case 0:
            var v41 int8 = 1
            _ = v41
        case 4294967295:
            var v42 uint8 = (v1 % (((v1 & (^(v1)))) | 1))
            _ = v42
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    v18 = (-(v18))
    v6 = ((v6 - v6) + float32((-2.25 / 1e10)))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
