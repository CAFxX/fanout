package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
type GenS2 struct { F0 uint16; F1 int; F2 uint8 }
type GenS3 struct { F0 uint16; F1 string; F2 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    var v3 int16 = -1
    v4 := []int{4294967295}
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int8 = 65
            switch 979 {
            case -1:
                v8 := GenS1{F0: -1, F1: 880, F2: (false || (!true))}
                _ = v8
            case 9223372036854775807:
                v3 += 0
            default:
                fmt.Println(v7)
            }
            _ = fi6
            _ = v7
        }
        if (!(!true)) {
            var v9 uint64 = 816
            v9 = v9
            v9 = v9
            _ = v9
        } else {
            switch "hello" {
            case "foo bar":
                v3 = v3
            default:
                v10 := GenS1{F0: 1, F1: 790, F2: (!true)}
                _ = v10
            }
        }
        v3 += 0
        _ = fi5
    }
    v1[0] = 444
    var v11 int8 = 100
    var v12 float32 = 3.14159
    v11 -= 1
    for fi13 := 0; fi13 < 20; fi13++ {
        fmt.Println(v12, v11)
        if true {
            v14 := GenS2{F0: 644, F1: (-(fi13)), F2: 128}
            _ = v14
        } else {
            v15 := GenS2{F0: 902, F1: fi13, F2: 145}
            _ = v15
        }
        v16 := GenS3{F0: 32767, F1: ("\n\t" + "foo bar"), F2: v11}
        _ = fi13
        _ = v16
    }
    if true {
        v17 := GenS3{F0: 1, F1: "a", F2: (^(v11))}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 19; fi18++ {
            v12 = (3.14159 * (v12 / (v12 * 3.14159)))
            for fi19 := 0; fi19 < 3; fi19++ {
                if (722 > 2) {
                    var v20 uint64 = 42
                    var v21 uint = 255
                    _ = v20
                    _ = v21
                }
                v22 := GenS3{F0: 65535, F1: "hello", F2: v11}
                switch (fi19 ^ fi19) {
                case 138:
                    v23 := GenS2{F0: 65535, F1: (-(fi18)), F2: 2}
                    _ = v23
                case 95:
                    if (128 != 1) {
                        switch "foo bar" {
                        case "a":
                            v24 := GenS4{F0: 189}
                            _ = v24
                        case "foo bar":
                            if true {
                                v12 += v12
                                fmt.Println(v12)
                            } else {
                                if (!true) {
                                    if len(v4) > 0 {
                                        v4[0] = fi19
                                    }
                                    switch (fi18 ^ (fi19 | fi18)) {
                                    case 32768:
                                        v12 -= ((v12 / 1.5) / 1.0)
                                    default:
                                        v3 = (^(v3))
                                    }
                                    if (!false) {
                                        var v25 int = (fi18 * (fi19 * fi18))
                                        var v26 int16 = -1
                                        _ = v25
                                        _ = v26
                                    } else {
                                        v27 := GenS3{F0: 42, F1: "hello", F2: 127}
                                        _ = len(v2)
                                        v3 = 605
                                        v28 := GenS5{F0: 42, F1: 1}
                                        _ = v27
                                        _ = v28
                                    }
                                    var v29 uint = 62
                                    _ = v29
                                }
                                v12 -= -2.25
                                fmt.Println(v3)
                            }
                        case "\n\t":
                            fmt.Println(v3, v12, fi18)
                        }
                        fmt.Println(v3, v3, v11)
                        v3 = (v3 | (v3 - -323))
                    } else {
                        v12 -= (v12 * (0.1 / float32(1e10)))
                        var v30 float32 = 3.14159
                        if ((32767 < 2) && false) {
                            v31 := GenS6{F0: (fi19 << ((545) & 63))}
                            _ = v31
                        } else {
                            var v32 uint32 = 2147483647
                            _ = v32
                        }
                        _ = v30
                    }
                }
                _ = fi19
                _ = v22
            }
            v11 = (v11 * (v11 * (v11 % (((v11 & 127)) | 1))))
            _ = fi18
        }
        v11 = (v11 & (v11 % ((-1) | 1)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
