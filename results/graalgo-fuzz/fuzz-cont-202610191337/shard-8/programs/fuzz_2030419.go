package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 bool; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool }

func main() {
    var v1 uint32 = 2
    v2 := []int{32768}
    v3 := []string{"a", "hello", ("\n\t" + "\n\t")}
    var v4 int32 = 42
    v5 := []int{-250}
    var v6 int = 81
    var v7 uint8 = 2
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: 100, F1: -1, F2: ("" + "foo bar")}
        switch (("" + "a") + "") {
        case "hello":
            if false {
                var v10 uint8 = (v7 + (v7 ^ 255))
                v7 += (v7 ^ 176)
                _ = v10
            }
        }
        _ = fi8
        _ = v9
    }
    v4 = v4
    for fi11 := 0; fi11 < 8; fi11++ {
        v4 -= v4
        _ = fi11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        v2 = append(v2, (fi12 >> ((65535) & 63)))
        _ = fi12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        v7 = v7
        var v14 int64 = 0
        var v15 uint64 = 632
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 string = "hello"
    if true {
        v1 -= v1
        v1 = (v1 << ((65535) & 63))
        if false {
            for fi17 := 0; fi17 < 7; fi17++ {
                if (1 <= 156) {
                    v6 -= (v6 - -773)
                }
                var v18 uint64 = 18446744073709551615
                _ = fi17
                _ = v18
            }
        } else {
            for fi19 := 0; fi19 < 5; fi19++ {
                if (!false) {
                    for fi20 := 0; fi20 < 9; fi20++ {
                        if true {
                            _ = len(v2)
                            v16 = v16
                            if (0.1 < float64(0.0)) {
                                if false {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    fmt.Println(v4, v1, fi19)
                                    switch v16 {
                                    case "\n\t":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case "hello":
                                        v1 -= (^(v1))
                                        v7 -= (v7 % (((v7 * v7)) | 1))
                                        fmt.Println(v7)
                                    }
                                    v21 := GenS1{F0: 0, F1: 1000, F2: v16}
                                    var v22 int64 = 1
                                    _ = v21
                                    _ = v22
                                } else {
                                    var v23 uint8 = (^(v7))
                                    _ = v23
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                        v24 := GenS2{F0: -1, F1: uint32(v7)}
                        var v25 uint32 = v1
                        _ = fi20
                        _ = v24
                        _ = v25
                    }
                    v6 += v6
                    var v26 string = ((((("\n\t" + v16) + (v16 + "foo bar")) + ("a" + v16)) + v16) + (v16 + "foo bar"))
                    _ = v26
                } else {
                    var v27 int8 = 1
                    v28 := GenS2{F0: (v27 | v27), F1: (v1 | v1)}
                    _ = v27
                    _ = v28
                }
                switch v16 {
                case "hello":
                    v6 = v6
                case "foo bar":
                    v1 += 4294967295
                case "a":
                    if len(v2) > 0 {
                        v2[0] = (v6 >> ((uint(v1)) & 63))
                    }
                default:
                    for fi29 := 0; fi29 < 20; fi29++ {
                        fmt.Println(fi29, v16)
                        v30 := GenS3{F0: (false || false), F1: -128, F2: (v1 | 4294967295)}
                        _ = fi29
                        _ = v30
                    }
                }
                _ = fi19
            }
        }
    } else {
        var v31 int = v6
        for fi32 := 0; fi32 < 10; fi32++ {
            if (0.0 != 1e10) {
                var v33 int = fi32
                switch 1 {
                case 2147483647:
                    v7 = v7
                case 284:
                    if (1 > 100) {
                        for fi34 := 0; fi34 < 8; fi34++ {
                            switch "\n\t" {
                            case "":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case "\n\t":
                                if false {
                                    v5 = append(v5, (fi32 / ((fi32) | 1)))
                                    var v35 float32 = float32((0.1 / 0.1))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v35
                                }
                            case "a":
                                if true {
                                    fmt.Println(fi32, v6, fi32)
                                    v36 := GenS3{F0: false, F1: 42, F2: 160}
                                    _ = v36
                                } else {
                                    _ = len(v3)
                                    _ = len(v5)
                                }
                            default:
                                switch v16 {
                                case "foo bar":
                                    v37 := GenS3{F0: false, F1: 14, F2: (^(v1))}
                                    _ = v37
                                case "a":
                                    var v38 bool = false
                                    _ = v38
                                case "\n\t":
                                    v39 := GenS4{F0: 65535}
                                    _ = v39
                                default:
                                    fmt.Println(fi32, v33)
                                }
                            }
                            _ = fi34
                        }
                        for fi40 := 0; fi40 < 12; fi40++ {
                            v41 := GenS5{F0: int64(v1)}
                            _ = fi40
                            _ = v41
                        }
                        fmt.Println(v6, v31)
                    } else {
                        v42 := GenS1{F0: -96, F1: 32767, F2: v16}
                        v31 -= (v33 + v33)
                        _ = v42
                    }
                }
                _ = v33
            } else {
                v43 := GenS6{F0: (0.0 < -2.25)}
                var v44 int32 = v4
                _ = v43
                _ = v44
            }
            v1 += 2
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi32
        }
        _ = len(v2)
        _ = v31
    }
    v45 := GenS2{F0: 42, F1: v1}
    fmt.Println(v6, v16, v1)
    v16 = "foo bar"
    fmt.Println(v4)
    _ = len(v3)
    v7 = 1
    var v46 uint8 = 127
    var v47 uint32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
