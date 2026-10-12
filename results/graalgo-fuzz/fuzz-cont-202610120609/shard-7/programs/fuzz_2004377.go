package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int32 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int64; F1 bool; F2 uint64 }
type GenS4 struct { F0 int; F1 int8; F2 int8 }

func main() {
    var v1 uint32 = 1
    v2 := []int{-337, 683, 32768, -1}
    var v3 uint = 0
    v4 := []int{42, -2147483648}
    v5 := make(map[string]int)
    var v6 int64 = 0
    v3 -= (v3 * (v3 | (v3 | v3)))
    if true {
        for fi7 := 0; fi7 < 20; fi7++ {
            if (("\n\t" + "foo bar") >= ("a" + "foo bar")) {
                fmt.Println(v1, fi7, fi7)
            } else {
                for fi8 := 0; fi8 < 4; fi8++ {
                    if (true && ((0 >= 0) && (-32768 >= int16(v3)))) {
                        var v9 uint8 = 210
                        var v10 float32 = ((0.0 / 1.0) / -2.25)
                        _ = v10
                        _ = v9
                    } else {
                        var v11 int8 = 109
                        if true {
                            v6 = (^(v6))
                            var v12 int64 = v6
                            v12 -= v12
                            _ = v12
                        } else {
                            v1 += (v1 << ((0) & 63))
                            _ = len(v5)
                        }
                        if false {
                            v6 -= v6
                            var v13 uint8 = uint8(v1)
                            switch "hello" {
                            case "a":
                                if (!((^(fi7)) != int(v3))) {
                                    v14 := GenS1{F0: v1, F1: true, F2: 0}
                                    v6 += (v6 & -388)
                                    if (v11 < (v11 + (-(v11)))) {
                                        var v15 int8 = (v11 % (((v11 % ((v11) | 1))) | 1))
                                        _ = v15
                                    } else {
                                        v11 -= v11
                                        var v16 uint32 = 42
                                        v17 := GenS2{F0: true, F1: (fi7 + (-(fi8)))}
                                        _ = v16
                                        _ = v17
                                    }
                                    _ = v14
                                }
                            case "foo bar":
                                v5[(("hello" + "a") + "foo bar")] = int(v1)
                            default:
                                var v18 int32 = -165
                                _ = v18
                            }
                            _ = v13
                        } else {
                            v19 := GenS3{F0: 824, F1: ((true || false) || (!false)), F2: 4294967295}
                            var v20 int32 = -782
                            v5["\n\t"] = 42
                            switch "foo bar" {
                            case "":
                                var v21 float64 = 0.1
                                _ = v21
                            case "foo bar":
                                v1 = v1
                            case "hello":
                                if true {
                                    var v22 uint16 = 0
                                    v6 = (v6 >> ((47) & 63))
                                    v3 += (^(v3))
                                    _ = v22
                                } else {
                                    var v23 bool = (91 == 1)
                                    fmt.Println(v6)
                                    _ = v23
                                }
                            default:
                                v20 = 0
                            }
                            _ = v19
                            _ = v20
                        }
                        _ = v11
                    }
                    _ = fi8
                }
                fmt.Println(fi7)
                fmt.Println(fi7, fi7)
            }
            var v24 float32 = 0.0
            _ = fi7
            _ = v24
        }
    } else {
        var v25 bool = false
        switch -429 {
        case 1:
            if (true && (-136 == 325)) {
                v5["hello"] = -9223372036854775808
            }
        case 0:
            var v26 uint32 = (v1 >> (((^(v3))) & 63))
            _ = v26
        case -9223372036854775808:
            var v27 bool = (1.5 != float32(-2.25))
            _ = v27
        }
        for fi28 := 0; fi28 < 10; fi28++ {
            var v29 int8 = -92
            _ = fi28
            _ = v29
        }
        _ = v25
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        for fi31 := 0; fi31 < 12; fi31++ {
            v6 += v6
            var v32 float32 = 1.5
            _ = fi31
            _ = v32
        }
        _ = len(v4)
        _ = fi30
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = (v3 - v3)
    var v33 int16 = -32768
    for fi34 := 0; fi34 < 20; fi34++ {
        if true {
            v35 := GenS4{F0: 0, F1: 1, F2: -128}
            v33 = (-(v33))
            v36 := GenS1{F0: (^(v1)), F1: (!false), F2: 720}
            _ = v35
            _ = v36
        }
        var v37 int16 = (v33 << (((v3 ^ (^(v3)))) & 63))
        var v38 int = (fi34 << ((65535) & 63))
        _ = fi34
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        var v40 int = 42
        _ = fi39
        _ = v40
    }
    fmt.Println(v33, v33)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
