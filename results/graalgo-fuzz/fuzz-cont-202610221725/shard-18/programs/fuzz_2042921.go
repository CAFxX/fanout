package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -59
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    v5 := []int{-972}
    fmt.Println(v2)
    if len(v5) > 0 {
        v5[0] = 543
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi7 := 0; fi7 < 6; fi7++ {
            switch fi6 {
            case 1:
                switch (fi7 - 513) {
                case 990:
                    var v8 uint64 = 2
                    _ = v8
                case 9223372036854775807:
                    if true {
                        v9 := GenS1{F0: -9223372036854775808}
                        if true {
                            v2 += (v2 << ((42) & 63))
                            if (128 > 168) {
                                var v10 int8 = (-(v2))
                                fmt.Println(fi6, v2)
                                v2 = (v10 + 42)
                                _ = v10
                            } else {
                                if (2 < 255) {
                                    fmt.Println(fi6, v2)
                                }
                            }
                        }
                        var v11 float32 = (3.14159 - float32(0.1))
                        var v12 int32 = 62
                        _ = v11
                        _ = v12
                        _ = v9
                    }
                }
            case 128:
                var v13 uint8 = 42
                _ = v13
            }
            var v14 uint16 = 416
            v14 -= (v14 << ((4294967295) & 63))
            _ = fi7
            _ = v14
        }
        _ = fi6
    }
    if (true && false) {
        switch ("\n\t" + "a") {
        case "\n\t":
            fmt.Println(v2)
        case "a":
            delete(v1, -2147483648)
        }
    }
    if (v2 <= v2) {
        if (true && false) {
            fmt.Println(v2)
            for fi15 := 0; fi15 < 12; fi15++ {
                for fi16 := 0; fi16 < 13; fi16++ {
                    switch fi15 {
                    case 32767:
                        if (875 == 1) {
                            var v17 int8 = v2
                            fmt.Println(fi16)
                            if false {
                                switch -9223372036854775808 {
                                case 32768:
                                    v2 = v17
                                case 127:
                                    v18 := GenS2{F0: 0, F1: 42, F2: 0}
                                    _ = v18
                                }
                            }
                            fmt.Println(v2, v2, fi16)
                            fmt.Println(v17, v2)
                            _ = v17
                        } else {
                            if (!(!false)) {
                                v2 = v2
                            } else {
                                v2 += 0
                                var v19 int64 = -56
                                _ = v19
                            }
                            switch "hello" {
                            case "":
                                if ((171 < 32767) || ((3.14159 / 3.14159) == 0.0)) {
                                    var v20 uint16 = 0
                                    _ = v20
                                } else {
                                    v2 += (^(v2))
                                }
                            case "hello":
                                _ = len(v5)
                                v2 += (v2 << ((616) & 63))
                                switch "\n\t" {
                                case "hello":
                                    var v21 uint8 = 42
                                    _ = v21
                                default:
                                    delete(v3, "hello")
                                }
                            case "foo bar":
                                var v22 float32 = 1e10
                                _ = v22
                            default:
                                v23 := GenS2{F0: 1, F1: 255, F2: 932}
                                _ = v23
                            }
                        }
                    case 128:
                        var v24 int8 = v2
                        _ = v24
                    default:
                        fmt.Println(fi16)
                    }
                    _ = fi16
                }
                _ = fi15
            }
            var v25 uint32 = 65535
            _ = v25
        }
        if (false && (!true)) {
            for fi26 := 0; fi26 < 3; fi26++ {
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 uint16 = 2
                    v28 += (v28 | (v28 ^ v28))
                    _ = fi27
                    _ = v28
                }
                switch "hello" {
                case "\n\t":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "foo bar":
                    fmt.Println(v2, v2)
                default:
                    var v29 int64 = -1
                    _ = v29
                }
                v2 += v2
                _ = fi26
            }
        }
    } else {
        var v30 int = -2147483648
        _ = v30
    }
    v2 += v2
    v31 := v3[("a" + "a")]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
