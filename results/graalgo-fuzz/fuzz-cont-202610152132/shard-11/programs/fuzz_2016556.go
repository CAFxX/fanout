package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int16 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int) uint8 {
    for fi4 := 0; fi4 < 12; fi4++ {
        p0 += (p0 & fi4)
        _ = fi4
    }
    v5 := []int{(-(p0)), 2147483647, 127, p0}
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint16 = 255
        if (v7 < v7) {
            fmt.Println(fi6, fi6, fi6)
            v7 += v7
            v7 = (v7 << ((255) & 63))
        }
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: (!(true && (!false))), F1: 215, F2: 32767}
    switch "a" {
    case "":
        if true {
            var v9 uint32 = 841
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v9
        } else {
            if true {
                switch (^(p0)) {
                case 955:
                    var v10 int8 = 100
                    _ = v10
                case 0:
                    if len(v5) > 0 {
                        v5[0] = 0
                    }
                default:
                    for fi11 := 0; fi11 < 5; fi11++ {
                        _ = len(v5)
                        switch (^(fi11)) {
                        case 9223372036854775807:
                            if ((-32768 == 595) && ((1 < -119) && (!true))) {
                                p0 = -1
                                var v12 uint8 = 128
                                _ = v12
                            } else {
                                p0 += (p0 >> ((65535) & 63))
                            }
                        }
                        for fi13 := 0; fi13 < 8; fi13++ {
                            switch "hello" {
                            case "":
                                v14 := GenS1{F0: true, F1: 208, F2: -69}
                                _ = v14
                            case "foo bar":
                                fmt.Println(p0, p0)
                            case "a":
                                p0 = (-(fi13))
                            }
                            if (!false) {
                                var v15 float64 = (0.0 + 1e10)
                                var v16 uint32 = 1
                                p0 -= (fi11 & (fi11 + 4294967295))
                                var v17 int16 = 327
                                _ = v15
                                _ = v16
                                _ = v17
                            } else {
                                v18 := GenS2{F0: "a", F1: 0}
                                _ = v18
                            }
                            _ = fi13
                        }
                        _ = fi11
                    }
                }
                for fi19 := 0; fi19 < 9; fi19++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi19
                }
            } else {
                _ = len(v5)
                for fi20 := 0; fi20 < 6; fi20++ {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        _ = len(v5)
                        _ = fi21
                    }
                    for fi22 := 0; fi22 < 20; fi22++ {
                        fmt.Println(fi20, fi22)
                        p0 = 128
                        _ = fi22
                    }
                    _ = fi20
                }
            }
        }
    case "a":
        var v23 uint32 = 21
        _ = v23
    }
    _ = p0
    _ = v5
    _ = v8
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{2147483647, 4294967295}
    v3 := make(map[uint32]uint32)
    var v24 uint32 = 4294967295
    var v25 uint64 = 346
    v26 := GenS3{F0: 0, F1: (v25 + 306)}
    fmt.Println(v24, v24)
    v27 := GenS4{F0: (^(v24)), F1: 128}
    fmt.Println(v24, v24)
    _ = len(v2)
    v3[(v24 | (v24 & (v24 - 1)))] = (^(v24))
    if false {
        if (!(1.5 >= ((1e10 + float32(1.5)) / 3.14159))) {
            if ((100 == 0) && false) {
                v25 = v25
                if false {
                    var v28 int64 = -213
                    v29 := GenS5{F0: v24, F1: 2}
                    _ = v28
                    _ = v29
                } else {
                    var v30 int16 = 1
                    var v31 int32 = 0
                    _ = v30
                    _ = v31
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v2 = append(v2, 9223372036854775807)
            }
        }
        v25 -= 18446744073709551615
    }
    if true {
        var v32 int32 = -1
        if true {
            v33 := GenS6{F0: uint32(v25), F1: v25}
            switch "\n\t" {
            case "hello":
                var v34 int = -34
                _ = v34
            case "":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "foo bar":
                if ((v25 << ((1) & 63)) <= (v25 + 18446744073709551615)) {
                    var v35 int32 = (v32 | v32)
                    var v36 uint32 = (v24 & v24)
                    _ = v35
                    _ = v36
                }
            default:
                v2 = append(v2, -736)
            }
            v24 = 0
            _ = v33
        } else {
            var v37 int32 = v32
            _ = v37
        }
        _ = v32
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v24 = (^(v24))
    }
    var v38 int64 = 1
    var v39 string = "foo bar"
    if len(v2) > 0 {
        v2[0] = -2147483648
    }
    fmt.Println(v39, v25, v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v38
    _ = v39
}
