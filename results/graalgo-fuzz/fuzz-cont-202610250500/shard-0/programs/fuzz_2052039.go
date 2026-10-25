package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }

func main() {
    v1 := []string{"", ("foo bar" + "\n\t"), "a"}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    var v4 uint = 65535
    v5 := make(map[int]uint32)
    v6 := GenS1{F0: 42}
    v4 -= (v4 | v4)
    v4 += v4
    var v7 uint8 = 165
    var v8 uint16 = 255
    v2[14] = 4294967295
    switch 4294967295 {
    case 42:
        v1 = append(v1, "a")
    case -438:
        if false {
            var v9 float64 = 1.5
            fmt.Println(v9)
            _ = v9
        } else {
            var v10 int8 = -11
            if false {
                for fi11 := 0; fi11 < 10; fi11++ {
                    var v12 string = "hello"
                    _ = fi11
                    _ = v12
                }
            } else {
                v4 = v4
                if (!true) {
                    var v13 uint8 = v7
                    _ = v13
                }
            }
            switch "hello" {
            case "foo bar":
                for fi14 := 0; fi14 < 2; fi14++ {
                    v15 := GenS1{F0: v7}
                    _ = fi14
                    _ = v15
                }
            case "\n\t":
                fmt.Println(v8)
            case "a":
                var v16 uint64 = 42
                _ = v16
            default:
                if (true || (false && true)) {
                    for fi17 := 0; fi17 < 15; fi17++ {
                        var v18 int16 = -1
                        if (!((fi17 >> ((v4) & 63)) >= fi17)) {
                            v19 := GenS1{F0: uint8(v4)}
                            v4 += v4
                            _ = v19
                        } else {
                            v7 -= v7
                            if true {
                                if (true && (v7 != v7)) {
                                    v10 = v10
                                } else {
                                    v20 := v3[("hello" + "\n\t")]
                                    for fi21 := 0; fi21 < 14; fi21++ {
                                        v2[2147483647] = 4294967295
                                        var v22 string = "foo bar"
                                        v4 -= uint(v7)
                                        _ = fi21
                                        _ = v22
                                    }
                                    v10 -= -1
                                    _ = v20
                                }
                            } else {
                                if (!(1 != 2)) {
                                    var v23 int32 = -1
                                    v8 -= 255
                                    switch "" {
                                    case "\n\t":
                                        if true {
                                            v7 += 127
                                            var v24 int64 = 0
                                            _ = v24
                                        }
                                    case "hello":
                                        v8 += 770
                                    case "foo bar":
                                        v25 := GenS1{F0: v7}
                                        _ = v25
                                    }
                                    _ = v23
                                }
                                fmt.Println(v10, v18, fi17)
                                var v26 int16 = int16(v8)
                                _ = v26
                            }
                            v10 = int8(v4)
                        }
                        v1 = append(v1, "hello")
                        _ = fi17
                        _ = v18
                    }
                    v27 := GenS2{F0: "\n\t", F1: (v7 >> ((83) & 63)), F2: (2147483647 > -714)}
                    _ = v27
                } else {
                    var v28 float64 = 3.14159
                    var v29 int16 = -396
                    _ = len(v5)
                    _ = v28
                    _ = v29
                }
            }
            _ = v10
        }
    case 715:
        var v30 int8 = 1
        _ = v30
    default:
        var v31 float64 = (0.0 + 0.0)
        _ = v31
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        v33 := GenS2{F0: "", F1: v7, F2: (false && true)}
        v7 += (v7 | 0)
        v4 -= (v4 | v4)
        _ = fi32
        _ = v33
    }
    if (-9223372036854775808 > 59) {
        v7 -= 237
        fmt.Println(v4, v7, v4)
    }
    var v34 uint = (v4 << ((v4) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
