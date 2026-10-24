package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint8; F2 int16 }
type GenS3 struct { F0 int8; F1 uint; F2 uint8 }
type GenS4 struct { F0 int; F1 uint16; F2 int64 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 string; F1 int32; F2 bool }
type GenS9 struct { F0 uint32; F1 int16 }
func genFn0() int64 {
    var v5 uint8 = 127
    var v6 int64 = 771
    _ = v5
    _ = v6
    return (v6 >> ((255) & 63))
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    v3 := []uint32{}
    v4 := make(map[uint32]uint32)
    var v7 uint16 = 162
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            v3 = append(v3, 65535)
            v7 -= (^(v7))
            v10 := GenS1{F0: (^(v7))}
            _ = fi9
            _ = v10
        }
        if len(v3) > 0 {
            v3[0] = 53
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS2{F0: 0, F1: 226, F2: 0}
            _ = fi11
            _ = v12
        }
        _ = fi8
    }
    v7 = (^(v7))
    v13 := GenS2{F0: -808, F1: 2, F2: -24}
    v14 := GenS1{F0: (v7 >> ((65535) & 63))}
    switch "" {
    case "foo bar":
        fmt.Println(v7)
    case "a":
        if (0.0 <= (0.0 + 0.1)) {
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 float64 = -2.25
                _ = fi15
                _ = v16
            }
        }
    case "hello":
        v7 = (v7 + v7)
    default:
        switch -703 {
        case 32768:
            var v17 int64 = 42
            _ = v17
        }
    }
    v18 := GenS3{F0: -1, F1: 2, F2: 42}
    delete(v1, -2147483648)
    v7 = 2
    v7 += (^(v7))
    v3 = append(v3, 65535)
    var v19 int = 637
    if (false && (!true)) {
        for fi20 := 0; fi20 < 4; fi20++ {
            v19 = 32768
            _ = fi20
        }
    } else {
        fmt.Println(v7, v7)
        for fi21 := 0; fi21 < 5; fi21++ {
            for fi22 := 0; fi22 < 12; fi22++ {
                if (0 > 1) {
                    v7 = v7
                    switch fi22 {
                    case 127:
                        v1[v19] = false
                    case -537:
                        var v23 int32 = -2147483648
                        _ = v23
                    case 642:
                        switch "foo bar" {
                        case "a":
                            fmt.Println(fi22, v19)
                        default:
                            if (!(648 >= 0)) {
                                switch (v19 << ((4294967295) & 63)) {
                                case -1:
                                    _ = len(v4)
                                case 9223372036854775807:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case 32768:
                                    v7 = v7
                                default:
                                    var v24 float32 = 0.1
                                    _ = v24
                                }
                                if (true && true) {
                                    _ = len(v4)
                                    switch v19 {
                                    case 0:
                                        v2[42] = 42
                                    }
                                    v7 -= 0
                                    switch (v19 - -520) {
                                    case 4294967295:
                                        v25 := GenS4{F0: (v19 % ((fi21) | 1)), F1: v7, F2: 9223372036854775807}
                                        _ = v25
                                    default:
                                        var v26 uint64 = 18446744073709551615
                                        _ = v26
                                    }
                                    fmt.Println(v7)
                                }
                            }
                        }
                    }
                    v27 := GenS5{F0: -958, F1: 42}
                    v19 = (v19 / ((-878) | 1))
                    _ = v27
                }
                fmt.Println(fi21, fi21, v19)
                _ = fi22
            }
            v28 := GenS6{F0: ("" + "")}
            v7 = 255
            _ = fi21
            _ = v28
        }
        var v29 bool = false
        _ = v29
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        v31 := GenS7{F0: (v7 >> ((42) & 63))}
        for fi32 := 0; fi32 < 1; fi32++ {
            if true {
                v33 := GenS8{F0: "hello", F1: 0, F2: true}
                _ = v33
            }
            _ = fi32
        }
        v34 := GenS9{F0: 784, F1: 978}
        _ = fi30
        _ = v31
        _ = v34
    }
    for fi35 := 0; fi35 < 15; fi35++ {
        _ = len(v3)
        for fi36 := 0; fi36 < 19; fi36++ {
            var v37 int64 = -166
            _ = fi36
            _ = v37
        }
        v7 -= (v7 + v7)
        _ = fi35
    }
    var v38 float64 = 1.5
    delete(v1, 0)
    if (750 > v7) {
        fmt.Println(v7)
        v39 := GenS1{F0: v7}
        _ = v39
    } else {
        var v40 string = ("hello" + "a")
        var v41 float32 = -2.25
        _ = v40
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v7
}
