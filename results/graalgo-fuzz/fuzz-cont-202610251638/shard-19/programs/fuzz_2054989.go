package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int32 = 42
    v2 := []uint32{864, 2147483648, 0}
    var v3 int = 73
    v4 := []string{("" + "foo bar"), "hello", "hello", "hello"}
    v3 -= (^(v3))
    v2 = append(v2, 1)
    var v5 int16 = 0
    if false {
        _ = len(v2)
        v5 = v5
    } else {
        switch (v3 + 863) {
        case -1:
            fmt.Println(v1, v1)
        case 32768:
            if ((("foo bar" + "\n\t") + "hello") <= ("" + "\n\t")) {
                v3 -= (v3 << ((65535) & 63))
                switch ("a" + "") {
                case "\n\t":
                    fmt.Println(v5, v3, v3)
                case "a":
                    for fi6 := 0; fi6 < 10; fi6++ {
                        switch fi6 {
                        case -233:
                            v1 = (v1 / (((v1 + v1)) | 1))
                        case 127:
                            v5 -= (v5 * (v5 & (v5 / ((v5) | 1))))
                        case 228:
                            v5 -= v5
                        default:
                            var v7 uint64 = 1
                            _ = v7
                        }
                        if (true || true) {
                            if true {
                                if false {
                                    for fi8 := 0; fi8 < 18; fi8++ {
                                        var v9 int = 9223372036854775807
                                        var v10 int64 = -869
                                        _ = fi8
                                        _ = v10
                                        _ = v9
                                    }
                                } else {
                                    if true {
                                        v11 := GenS1{F0: 42, F1: 0, F2: 42}
                                        _ = v11
                                    } else {
                                        for fi12 := 0; fi12 < 2; fi12++ {
                                            v13 := GenS1{F0: 4294967295, F1: 127, F2: 0}
                                            v14 := GenS2{F0: 255}
                                            _ = fi12
                                            _ = v13
                                            _ = v14
                                        }
                                    }
                                }
                            }
                        }
                        for fi15 := 0; fi15 < 10; fi15++ {
                            v2 = append(v2, 2147483648)
                            v16 := GenS2{F0: 1}
                            v1 -= 2147483647
                            _ = fi15
                            _ = v16
                        }
                        _ = fi6
                    }
                default:
                    switch v3 {
                    case 159:
                        switch "foo bar" {
                        case "a":
                            v3 += -9223372036854775808
                        case "foo bar":
                            switch "hello" {
                            case "":
                                if true {
                                    var v17 uint16 = 708
                                    _ = v17
                                } else {
                                    var v18 int = -887
                                    var v19 int16 = v5
                                    var v20 uint = 42
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                }
                            default:
                                var v21 int16 = 0
                                _ = v21
                            }
                        default:
                            _ = len(v2)
                        }
                    }
                }
            }
        }
        var v22 float32 = 0.1
        fmt.Println(v22, v22)
        _ = v22
    }
    var v23 int32 = v1
    switch "hello" {
    case "foo bar":
        v3 += (^(v3))
    case "\n\t":
        if false {
            if (!(!false)) {
                if false {
                    fmt.Println(v3, v5)
                    for fi24 := 0; fi24 < 10; fi24++ {
                        v5 += -1
                        v5 = (^(v5))
                        v25 := GenS2{F0: 127}
                        _ = fi24
                        _ = v25
                    }
                }
                var v26 int64 = 1
                _ = v26
            }
        } else {
            v1 = 0
            var v27 int32 = v1
            _ = v27
        }
    case "":
        var v28 int = v3
        _ = v28
    }
    v29 := GenS2{F0: 1}
    var v30 int8 = 0
    v1 -= v23
    if false {
        switch (^(v3)) {
        case -920:
            switch 1 {
            case 9223372036854775807:
                if ((!(0 > (v5 >> ((65535) & 63)))) && true) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            case -460:
                if len(v2) > 0 {
                    v2[0] = 42
                }
            default:
                v31 := GenS2{F0: 184}
                _ = v31
            }
        case -2147483648:
            var v32 string = ("" + "foo bar")
            _ = v32
        }
    } else {
        fmt.Println(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
