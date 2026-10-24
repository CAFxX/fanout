package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int8; F2 int32 }
type GenS3 struct { F0 int64; F1 int32; F2 int8 }
func genFn0(p0 int64) int32 {
    p0 += (p0 / ((p0) | 1))
    var v5 uint8 = 154
    if (false || false) {
        switch "\n\t" {
        case "":
            var v6 uint64 = 18446744073709551615
            _ = v6
        case "foo bar":
            if true {
                fmt.Println(p0)
            } else {
                p0 += (p0 | p0)
                v7 := []int{}
                _ = v7
            }
        }
    } else {
        switch 344 {
        case 2147483647:
            var v8 int8 = -125
            _ = v8
        case -587:
            p0 -= p0
        case 127:
            v9 := GenS1{F0: 42}
            _ = v9
        }
    }
    p0 -= 0
    if (!false) {
        v5 -= v5
    }
    _ = p0
    _ = v5
    return 2147483647
}

func main() {
    v1 := []string{"", ("" + "\n\t")}
    v2 := make(map[int]uint32)
    var v3 int64 = 514
    var v4 uint32 = 65535
    switch 1 {
    case 617:
        v3 = 9223372036854775807
    case -486:
        var v10 int32 = 42
        _ = v10
    case -9223372036854775808:
        v11 := GenS2{F0: 0, F1: -12, F2: 1}
        _ = v11
    default:
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := GenS1{F0: -1}
            if (127 <= 127) {
                if false {
                    if (!(false || (1 >= 142))) {
                        switch "\n\t" {
                        case "foo bar":
                            var v14 int = 127
                            _ = v14
                        default:
                            v1 = append(v1, "hello")
                        }
                    }
                    var v15 int32 = 0
                    var v16 uint8 = 56
                    _ = v15
                    _ = v16
                } else {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    for fi17 := 0; fi17 < 19; fi17++ {
                        v4 += (v4 / ((2147483648) | 1))
                        _ = fi17
                    }
                    for fi18 := 0; fi18 < 2; fi18++ {
                        if false {
                            v19 := GenS2{F0: 343, F1: 1, F2: 1}
                            switch (^(fi18)) {
                            case 2147483647:
                                v2[fi18] = (v4 % ((65535) | 1))
                            case 4294967295:
                                v4 = 42
                            }
                            fmt.Println(fi12)
                            var v20 int32 = -1
                            _ = v19
                            _ = v20
                        } else {
                            var v21 uint32 = (v4 % ((v4) | 1))
                            v21 += v4
                            _ = v21
                        }
                        v22 := v2[32767]
                        var v23 uint64 = 4294967295
                        delete(v2, (fi18 % (((fi18 ^ fi12)) | 1)))
                        _ = fi18
                        _ = v22
                        _ = v23
                    }
                }
                var v24 uint16 = 1
                if false {
                    v25 := GenS2{F0: v24, F1: -1, F2: -228}
                    _ = v25
                } else {
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    if (("a" == "") || (42 == 0)) {
                        v26 := GenS1{F0: 9223372036854775807}
                        _ = v26
                    }
                }
                _ = v24
            } else {
                v27 := GenS3{F0: v3, F1: 2147483647, F2: -128}
                _ = v27
            }
            var v28 int8 = 0
            _ = fi12
            _ = v13
            _ = v28
        }
    }
    v4 += v4
    v4 = v4
    v4 += (v4 * 2147483647)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v4, v4)
    case "":
        if len(v1) > 0 {
            v1[0] = ("" + (((("" + "") + "\n\t") + "a") + "foo bar"))
        }
    case "foo bar":
        v3 += v3
    default:
        v4 = v4
    }
    var v29 int16 = 92
    switch -112 {
    case 1:
        v30 := GenS2{F0: 32767, F1: 0, F2: 2147483647}
        _ = v30
    }
    v4 = v4
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 int64 = (^(v3))
        v33 := GenS1{F0: (v32 / ((v3) | 1))}
        if (2 < 1) {
            if false {
                v32 = v32
                switch ("\n\t" + ((("\n\t" + "hello") + "foo bar") + "")) {
                case "foo bar":
                    fmt.Println(v29, v4)
                case "":
                    switch 2147483647 {
                    case 596:
                        v2[fi31] = v4
                    case -1:
                        delete(v2, fi31)
                    default:
                        fmt.Println(v29, fi31)
                    }
                default:
                    switch (fi31 * fi31) {
                    case 32768:
                        var v34 int8 = 115
                        _ = v34
                    case 9223372036854775807:
                        v4 -= (v4 % (((v4 / ((112) | 1))) | 1))
                    default:
                        v29 += v29
                    }
                }
                var v35 uint16 = 255
                _ = v35
            } else {
                var v36 uint32 = (^(v4))
                _ = v36
            }
            v29 = (v29 / ((v29) | 1))
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    v4 = (v4 - (v4 & v4))
    v3 -= v3
    v37 := GenS3{F0: 42, F1: -421, F2: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v4
}
