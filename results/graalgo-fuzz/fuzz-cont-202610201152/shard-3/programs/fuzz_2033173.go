package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]uint32)
    var v3 int32 = -1
    var v4 int16 = 1000
    delete(v2, 0)
    v5 := []int{32767, 9223372036854775807, -1}
    var v6 int8 = -22
    switch 127 {
    case -71:
        var v7 uint64 = 72
        _ = v7
    case 128:
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 14; fi9++ {
                v10 := v2[(-(fi8))]
                delete(v2, int(v6))
                var v11 int = fi8
                _ = fi9
                _ = v10
                _ = v11
            }
            if false {
                if false {
                    fmt.Println(fi8, fi8)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    var v12 uint32 = 752
                    var v13 int64 = -1
                    _ = v12
                    _ = v13
                }
                switch "hello" {
                case "a":
                    if false {
                        if len(v5) > 0 {
                            v5[0] = -9223372036854775808
                        }
                        if (false || (42 <= -666)) {
                            var v14 uint16 = 255
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 9; fi15++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi15
                        }
                    } else {
                        var v16 uint8 = 0
                        v17 := GenS1{F0: (v1 >> ((0) & 63))}
                        _ = v16
                        _ = v17
                    }
                case "hello":
                    fmt.Println(fi8)
                case "":
                    v6 -= (v6 << ((1) & 63))
                default:
                    for fi18 := 0; fi18 < 14; fi18++ {
                        v1 = (v1 << ((964) & 63))
                        v4 = (v4 | v4)
                        _ = fi18
                    }
                }
                if true {
                    var v19 uint32 = 2147483648
                    _ = v19
                } else {
                    if true {
                        v1 = (v1 / ((v1) | 1))
                    } else {
                        v4 += v4
                        var v20 int16 = -1
                        var v21 float32 = 0.1
                        _ = v20
                        _ = v21
                    }
                }
            }
            _ = fi8
        }
    case 32768:
        var v22 uint64 = 18446744073709551615
        _ = v22
    default:
        for fi23 := 0; fi23 < 13; fi23++ {
            if false {
                var v24 uint = 649
                v2[(fi23 >> ((0) & 63))] = 2147483647
                if true {
                    var v25 int = (fi23 / (((fi23 + 2147483647)) | 1))
                    if ((v25 | int(v24)) <= fi23) {
                        var v26 float64 = 0.1
                        _ = v26
                    } else {
                        fmt.Println(fi23)
                        v25 += fi23
                    }
                    _ = v25
                } else {
                    var v27 float32 = 3.14159
                    switch ("foo bar" + "foo bar") {
                    case "hello":
                        v1 -= v1
                    case "\n\t":
                        var v28 int16 = v4
                        _ = v28
                    }
                    _ = v27
                }
                _ = v24
            }
            _ = fi23
        }
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        if (42 == 18446744073709551615) {
            var v30 int = (fi29 % (((fi29 - int(v4))) | 1))
            _ = v30
        } else {
            fmt.Println(v4)
            v4 -= (v4 ^ (v4 << ((596) & 63)))
        }
        _ = fi29
    }
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    v4 += (v4 & (v4 / (((v4 * v4)) | 1)))
    v31 := GenS1{F0: (v1 << ((65535) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
