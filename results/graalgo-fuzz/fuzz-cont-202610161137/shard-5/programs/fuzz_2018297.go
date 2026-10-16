package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := []uint32{2147483647, 2147483648, 642, 2147483647}
    var v2 uint16 = 1
    v3 := make(map[string]int)
    v4 := GenS1{F0: 514, F1: 2}
    for fi5 := 0; fi5 < 16; fi5++ {
        switch "a" {
        case "":
            if false {
                for fi6 := 0; fi6 < 9; fi6++ {
                    fmt.Println(fi6, fi6)
                    if (42 < 0) {
                        var v7 float32 = ((1.5 + 1.5) - 3.14159)
                        v2 -= v2
                        v8 := GenS1{F0: 1, F1: 18446744073709551615}
                        v2 -= v2
                        v9 := GenS1{F0: -1, F1: 65535}
                        if true {
                            if false {
                                v7 = float32(1e10)
                                var v10 bool = (!(4294967295 > 42))
                                _ = v10
                            }
                        }
                        _ = v7
                        _ = v8
                        _ = v9
                    } else {
                        switch ("a" + (("\n\t" + ("\n\t" + ("foo bar" + ""))) + "")) {
                        case "":
                            v2 += (^(v2))
                        case "\n\t":
                            var v11 string = "a"
                            _ = v11
                        case "foo bar":
                            var v12 uint64 = 743
                            _ = v12
                        default:
                            var v13 bool = ((v2 < (v2 >> ((559) & 63))) || true)
                            _ = v13
                        }
                        if false {
                            v2 += (v2 << ((42) & 63))
                            if (!false) {
                                var v14 int64 = -525
                                _ = v14
                            }
                            fmt.Println(v2, fi5)
                        } else {
                            switch "a" {
                            case "hello":
                                v15 := GenS2{F0: 104}
                                _ = v15
                            default:
                                if false {
                                    v2 += v2
                                } else {
                                    if true {
                                        v16 := GenS1{F0: 1000, F1: 65535}
                                        var v17 float64 = ((0.1 * -2.25) * float64(float32(3.14159)))
                                        _ = v16
                                        _ = v17
                                    }
                                    _ = len(v1)
                                }
                            }
                            v18 := GenS3{F0: int16(fi5), F1: 543}
                            _ = v18
                        }
                        var v19 bool = false
                        _ = v19
                    }
                    v20 := GenS2{F0: 35}
                    _ = fi6
                    _ = v20
                }
            } else {
                var v21 int16 = -1
                var v22 int16 = v21
                _ = v21
                _ = v22
            }
        case "a":
            var v23 float64 = 0.0
            _ = v23
        case "\n\t":
            v24 := GenS4{F0: 42}
            _ = v24
        default:
            v1 = append(v1, 2147483648)
        }
        _ = fi5
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        var v26 string = "a"
        _ = fi25
        _ = v26
    }
    delete(v3, "")
    v2 = v2
    var v27 string = ("foo bar" + "a")
    var v28 int32 = 42
    switch -9223372036854775808 {
    case 1:
        v29 := GenS4{F0: -9223372036854775808}
        _ = v29
    default:
        var v30 int8 = 0
        _ = v30
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    switch 1 {
    case -2147483648:
        _ = len(v3)
    default:
        fmt.Println(v2, v27)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
