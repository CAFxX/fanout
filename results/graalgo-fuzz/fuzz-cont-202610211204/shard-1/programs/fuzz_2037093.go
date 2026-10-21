package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint }
type GenS2 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint16 }

func main() {
    var v1 int64 = -1
    v2 := make(map[int]string)
    var v3 uint32 = 2
    var v4 uint = 601
    if true {
        v3 = v3
    }
    v5 := GenS1{F0: (true || ((42 < 304) && false)), F1: 127, F2: 161}
    var v6 float64 = -2.25
    var v7 int64 = v1
    switch 4294967295 {
    case 679:
        v8 := []int{}
        _ = v8
    default:
        v6 = 1.5
    }
    var v9 string = "foo bar"
    v1 = v7
    var v10 string = (((v9 + v9) + "hello") + "")
    v11 := GenS1{F0: ((false || false) || (0 <= 60)), F1: 42, F2: (v4 | 65535)}
    var v12 uint64 = 2
    v13 := []int{}
    fmt.Println(v3)
    var v14 uint16 = 1
    if (true || ((v14 + 16) >= 32767)) {
        v6 = 1e10
        v15 := GenS2{F0: 0, F1: (v3 - (v3 << ((uint(v3)) & 63))), F2: v14}
        v16 := GenS3{F0: v4, F1: (^(v14))}
        _ = v15
        _ = v16
    } else {
        if (-2.25 > 3.14159) {
            if true {
                v7 = (v7 >> (((v4 | 4294967295)) & 63))
                var v17 uint = v4
                for fi18 := 0; fi18 < 1; fi18++ {
                    if false {
                        if ((-2147483648 >= 127) && (v14 >= v14)) {
                            switch -1 {
                            case 0:
                                if true {
                                    var v19 uint32 = v3
                                    _ = v19
                                } else {
                                    v6 = v6
                                    v4 += (v17 << (((v4 | (v4 * v17))) & 63))
                                }
                            case 42:
                                var v20 uint64 = (v12 + v12)
                                _ = v20
                            case -395:
                                for fi21 := 0; fi21 < 13; fi21++ {
                                    if (255 < 130) {
                                        var v22 int64 = (v7 - (^(v1)))
                                        v6 += v6
                                        _ = v22
                                    }
                                    var v23 bool = (v6 == v6)
                                    _ = fi21
                                    _ = v23
                                }
                            }
                            v24 := GenS2{F0: 42, F1: 42, F2: v14}
                            var v25 bool = false
                            _ = v24
                            _ = v25
                        } else {
                            switch fi18 {
                            case 381:
                                v7 += -674
                            }
                        }
                    } else {
                        if len(v13) > 0 {
                            v13[0] = (fi18 + (fi18 >> (((v4 % ((522) | 1))) & 63)))
                        }
                    }
                    _ = fi18
                }
                _ = v17
            } else {
                v13 = append(v13, -2147483648)
            }
            v13 = append(v13, -483)
        } else {
            for fi26 := 0; fi26 < 16; fi26++ {
                var v27 int8 = 127
                switch ("" + "\n\t") {
                case "":
                    v14 = (v14 >> ((0) & 63))
                case "foo bar":
                    v7 += 9223372036854775807
                case "\n\t":
                    var v28 int = 597
                    _ = v28
                default:
                    v29 := GenS3{F0: 4294967295, F1: (v14 >> (((v4 | 255)) & 63))}
                    _ = v29
                }
                _ = fi26
                _ = v27
            }
            if false {
                if (1 > -1) {
                    if false {
                        var v30 uint = (v4 * (v4 / ((2) | 1)))
                        v7 += 0
                        _ = v30
                    } else {
                        switch ("" + v10) {
                        case "\n\t":
                            switch v10 {
                            case "hello":
                                for fi31 := 0; fi31 < 14; fi31++ {
                                    v9 = v9
                                    v2[(-(fi31))] = v9
                                    delete(v2, -9223372036854775808)
                                    _ = fi31
                                }
                            case "a":
                                var v32 int32 = 1
                                _ = v32
                            default:
                                v4 = v4
                            }
                        case "foo bar":
                            var v33 bool = (!(42 <= -36))
                            _ = v33
                        default:
                            v34 := GenS1{F0: false, F1: 0, F2: (v4 - 89)}
                            _ = v34
                        }
                    }
                }
                v35 := GenS3{F0: (v4 % ((42) | 1)), F1: (v14 | uint16(v4))}
                _ = v35
            }
            if true {
                var v36 float32 = 0.1
                _ = v36
            }
        }
        v37 := GenS2{F0: 32767, F1: 771, F2: (v14 ^ (v14 >> ((65535) & 63)))}
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
