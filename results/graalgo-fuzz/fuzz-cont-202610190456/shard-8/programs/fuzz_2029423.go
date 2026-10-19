package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int8; F1 bool; F2 int64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{252, 9223372036854775807}
    var v3 float32 = (0.1 / 1.5)
    var v4 uint8 = 255
    v4 += v4
    if (9223372036854775807 >= 4294967295) {
        var v5 uint8 = v4
        for fi6 := 0; fi6 < 2; fi6++ {
            fmt.Println(v3, fi6)
            if len(v2) > 0 {
                v2[0] = fi6
            }
            v3 += 0.0
            _ = fi6
        }
        v5 += (v5 + v5)
        _ = v5
    }
    v7 := GenS1{F0: 381}
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: 65535}
        _ = fi8
        _ = v9
    }
    v3 = 1.5
    switch "hello" {
    case "hello":
        v3 = 3.14159
    case "a":
        v10 := GenS2{F0: -2147483648, F1: 787}
        _ = v10
    default:
        if false {
            for fi11 := 0; fi11 < 13; fi11++ {
                v3 = v3
                _ = fi11
            }
            if true {
                var v12 float32 = (float32((1.5 / -2.25)) + v3)
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 uint8 = (v4 & (v4 % ((v4) | 1)))
                    v2 = append(v2, 1)
                    v12 = (float32(0.0) + (v12 - v12))
                    _ = fi13
                    _ = v14
                }
                fmt.Println(v4, v4, v4)
                _ = v12
            }
            v3 += float32((-2.25 / 1.0))
        } else {
            if (!true) {
                v15 := GenS1{F0: 42}
                if (!false) {
                    var v16 int64 = -11
                    v17 := GenS3{F0: 0, F1: 2, F2: 1}
                    _ = v16
                    _ = v17
                } else {
                    if (((1.5 / -2.25) <= 3.14159) && (42 == 4294967295)) {
                        for fi18 := 0; fi18 < 19; fi18++ {
                            if true {
                                v4 = v4
                            } else {
                                v4 = 42
                                if ((fi18 == fi18) || true) {
                                    var v19 uint = 2
                                    var v20 int8 = 100
                                    _ = v19
                                    _ = v20
                                }
                                fmt.Println(v3, v3)
                            }
                            _ = fi18
                        }
                    }
                    v3 += v3
                    if false {
                        var v21 float32 = v3
                        switch ("foo bar" + "") {
                        case "\n\t":
                            v22 := GenS4{F0: 127, F1: false, F2: 42}
                            _ = v22
                        case "a":
                            var v23 uint32 = 4294967295
                            _ = v23
                        }
                        switch ("" + "hello") {
                        case "":
                            switch ("a" + "foo bar") {
                            case "hello":
                                switch 42 {
                                case 1:
                                    var v24 uint32 = 65535
                                    _ = v24
                                case -990:
                                    fmt.Println(v4)
                                }
                            case "foo bar":
                                for fi25 := 0; fi25 < 13; fi25++ {
                                    for fi26 := 0; fi26 < 12; fi26++ {
                                        v21 = 3.14159
                                        v3 = -2.25
                                        _ = fi26
                                    }
                                    _ = fi25
                                }
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        case "\n\t":
                            v27 := v1[2147483648]
                            _ = v27
                        case "foo bar":
                            var v28 int8 = 127
                            _ = v28
                        }
                        _ = v21
                    }
                }
                for fi29 := 0; fi29 < 7; fi29++ {
                    v3 -= ((v3 - v3) - ((v3 * (((0.1 - v3) - 3.14159) / -2.25)) / v3))
                    v4 -= 2
                    v3 += 1.5
                    _ = fi29
                }
                _ = v15
            }
        }
    }
    _ = len(v1)
    v30 := v1[143]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v7
}
