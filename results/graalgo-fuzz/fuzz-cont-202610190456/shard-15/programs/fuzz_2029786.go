package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint16; F1 uint32; F2 bool }
type GenS4 struct { F0 bool; F1 uint64; F2 string }

func main() {
    v1 := []int{4294967295}
    var v2 int16 = -32768
    var v3 int = 128
    for fi4 := 0; fi4 < 20; fi4++ {
        if false {
            v3 += v3
            v3 = fi4
            fmt.Println(v3)
        } else {
            v3 -= (v3 / ((fi4) | 1))
        }
        v2 = (v2 / (((^(v2))) | 1))
        var v5 float64 = 3.14159
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        v3 = fi6
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int8 = 42
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 int64 = -9223372036854775808
    var v10 uint8 = 225
    var v11 uint = 0
    if false {
        if false {
            switch "hello" {
            case "foo bar":
                switch "" {
                case "hello":
                    if (3 > 32767) {
                        var v12 uint16 = 0
                        var v13 uint64 = 762
                        var v14 int16 = (v2 << ((v11) & 63))
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                case "foo bar":
                    for fi15 := 0; fi15 < 9; fi15++ {
                        _ = len(v1)
                        v16 := GenS1{F0: -1, F1: v9, F2: 4294967295}
                        switch "hello" {
                        case "foo bar":
                            v11 = v11
                        case "":
                            v3 = fi15
                        case "hello":
                            var v17 string = ("a" + "foo bar")
                            _ = v17
                        }
                        _ = fi15
                        _ = v16
                    }
                case "\n\t":
                    v3 -= -942
                default:
                    var v18 int32 = 0
                    _ = v18
                }
            case "a":
                v19 := GenS2{F0: (v11 + v11), F1: false}
                _ = v19
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        } else {
            v1 = append(v1, 1)
            v20 := GenS2{F0: (^(v11)), F1: false}
            switch (v3 * v3) {
            case -2147483648:
                v2 -= (v2 | (v2 | (v2 % ((32767) | 1))))
            case 898:
                v1 = append(v1, v3)
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v20
        }
        var v21 uint = 1
        _ = v21
    } else {
        var v22 int64 = (v9 ^ 865)
        v23 := GenS1{F0: 42, F1: (v22 >> ((v11) & 63)), F2: (v3 << (((v11 * v11)) & 63))}
        if len(v1) > 0 {
            v1[0] = (^(v3))
        }
        _ = v22
        _ = v23
    }
    v1 = append(v1, 870)
    v1 = append(v1, (v3 ^ v3))
    if (v3 < (v3 & v3)) {
        fmt.Println(v2, v2)
        if (294 > 42) {
            var v24 uint32 = 4294967295
            for fi25 := 0; fi25 < 9; fi25++ {
                v24 = (v24 | v24)
                v26 := GenS3{F0: 65535, F1: uint32(v10), F2: (!(4294967295 != uint64(v24)))}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi25
                _ = v26
            }
            if false {
                v27 := GenS3{F0: 968, F1: (v24 / ((4294967295) | 1)), F2: ((v3 * v3) > v3)}
                _ = v27
            } else {
                var v28 int8 = 127
                _ = v28
            }
            _ = v24
        } else {
            v2 = -1
        }
        fmt.Println(v3)
    } else {
        if len(v1) > 0 {
            v1[0] = (v3 | v3)
        }
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 int16 = (v2 * v2)
            switch 9223372036854775807 {
            case 32767:
                v10 -= v10
            default:
                if (!(!(((("a" + "\n\t") + "") + "\n\t") >= "a"))) {
                    if (fi29 > (v3 % ((fi29) | 1))) {
                        v9 += (v9 << (((v11 * (v11 * (v11 & (v11 | (v11 ^ 1)))))) & 63))
                        var v31 uint64 = 854
                        _ = v31
                    } else {
                        v32 := GenS2{F0: 782, F1: (true && false)}
                        var v33 int8 = -1
                        v9 = v9
                        _ = v32
                        _ = v33
                    }
                }
            }
            v34 := GenS4{F0: (false || false), F1: 1, F2: ("hello" + "hello")}
            _ = fi29
            _ = v30
            _ = v34
        }
        fmt.Println(v9, v10, v3)
    }
    var v35 uint8 = (v10 % (((^(v10))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v9
}
