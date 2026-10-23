package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 bool }
type GenS2 struct { F0 uint32; F1 string; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 uint32; F2 int64 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []string{"", "\n\t"}
    v3 := []int{}
    var v4 uint64 = v1
    fmt.Println(v1, v4)
    v1 = v1
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    v1 = (v4 - 0)
    var v5 uint16 = 255
    switch 128 {
    case -9223372036854775808:
        var v6 uint16 = uint16(v4)
        _ = v6
    default:
        v3 = append(v3, 9223372036854775807)
    }
    var v7 uint = 0
    if ("foo bar" != "a") {
        var v8 uint32 = 236
        v4 -= (v1 / (((v1 / ((v1) | 1))) | 1))
        _ = v8
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v9 := GenS1{F0: 0, F1: 2147483647, F2: ((-296 >= 525) && false)}
    var v10 int64 = 42
    switch "" {
    case "\n\t":
        var v11 int64 = int64(v7)
        _ = v11
    case "":
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 float64 = -2.25
            v5 -= 0
            _ = fi12
            _ = v13
        }
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        v1 = v4
    }
    v7 += (v7 / (((v7 * v7)) | 1))
    var v14 bool = (false || (false && true))
    v15 := GenS2{F0: 4294967295, F1: ("foo bar" + "hello"), F2: (v1 << ((uint(v5)) & 63))}
    if ((-1 >= -141) || (true || (32768 <= -9223372036854775808))) {
        var v16 uint16 = 539
        var v17 int = 127
        _ = v16
        _ = v17
    } else {
        if (2147483647 < 128) {
            if (!v14) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if ((v5 / ((v5) | 1)) != 255) {
                    var v18 uint32 = 664
                    v19 := GenS3{F0: v10}
                    v14 = v14
                    _ = v18
                    _ = v19
                } else {
                    v3 = append(v3, 485)
                    var v20 int16 = 334
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    _ = v20
                }
                if v14 {
                    if (true && v14) {
                        var v21 uint16 = 342
                        _ = v21
                    }
                    if false {
                        v22 := GenS4{F0: v4}
                        v5 = 898
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 10; fi23++ {
                            var v24 int16 = -32768
                            switch (fi23 >> (((^(v7))) & 63)) {
                            case 2147483647:
                                if true {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            case -2147483648:
                                v14 = v14
                            default:
                                v25 := GenS5{F0: (42 >= -1), F1: 0, F2: (v10 & (-(v10)))}
                                _ = v25
                            }
                            _ = fi23
                            _ = v24
                        }
                        v14 = (v10 <= (v10 | -853))
                        v10 = v10
                    }
                    for fi26 := 0; fi26 < 12; fi26++ {
                        v27 := GenS5{F0: (v5 != v5), F1: 42, F2: (v10 / ((-1) | 1))}
                        _ = len(v2)
                        var v28 int8 = -9
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                }
            }
            var v29 uint8 = 2
            _ = v29
        }
        if v14 {
            var v30 uint64 = (v4 >> (((v7 / (((v7 - v7)) | 1))) & 63))
            for fi31 := 0; fi31 < 6; fi31++ {
                v32 := GenS2{F0: 4294967295, F1: ("foo bar" + "a"), F2: (v1 + v4)}
                _ = fi31
                _ = v32
            }
            if v14 {
                fmt.Println(v1, v5)
            } else {
                fmt.Println(v1, v7)
            }
            _ = v30
        } else {
            v10 = v10
            for fi33 := 0; fi33 < 8; fi33++ {
                fmt.Println(v5)
                _ = fi33
            }
        }
        var v34 uint64 = v4
        _ = v34
    }
    v5 += (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
