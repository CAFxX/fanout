package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int16 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 bool }
func genFn0(p0 uint64) uint8 {
    var v4 float32 = (1.5 / 0.1)
    fmt.Println(v4)
    var v5 float64 = (1e10 * 0.1)
    var v6 int32 = 0
    var v7 uint64 = 18446744073709551615
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 127
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 800
    var v3 int32 = 1
    v3 += (^(v3))
    if (false && false) {
        for fi8 := 0; fi8 < 11; fi8++ {
            switch "" {
            case "":
                switch 1 {
                case -10:
                    v3 = v3
                }
            default:
                if (false && false) {
                    fmt.Println(v3)
                    fmt.Println(fi8, fi8, v2)
                } else {
                    for fi9 := 0; fi9 < 1; fi9++ {
                        var v10 bool = true
                        if (false || v10) {
                            v11 := GenS1{F0: (fi9 / ((fi9) | 1)), F1: v2, F2: -1}
                            switch fi9 {
                            case 1:
                                if (float32((float64((float32(1e10) / 1.5)) / 1.0)) != 0.0) {
                                    v12 := GenS1{F0: 200, F1: v2, F2: -1}
                                    _ = len(v1)
                                    _ = v12
                                }
                            case 4294967295:
                                var v13 uint8 = 0
                                _ = v13
                            }
                            fmt.Println(v10, v2, v3)
                            v14 := GenS1{F0: fi9, F1: (v2 >> ((65535) & 63)), F2: -259}
                            _ = v11
                            _ = v14
                        } else {
                            var v15 uint16 = 32767
                            v15 = 0
                            v10 = v10
                            _ = v15
                        }
                        fmt.Println(fi8)
                        _ = fi9
                        _ = v10
                    }
                    v3 = -2147483648
                }
            }
            v16 := GenS1{F0: (fi8 ^ fi8), F1: (v2 >> ((533) & 63)), F2: int16(v2)}
            _ = fi8
            _ = v16
        }
        v17 := []int{42}
        var v18 int8 = 42
        _ = v17
        _ = v18
    }
    switch ("hello" + "a") {
    case "foo bar":
        v3 = -2147483648
    case "a":
        v2 = 388
    }
    fmt.Println(v3)
    v2 = (v2 | v2)
    var v19 uint = uint(v2)
    v19 += 4294967295
    delete(v1, -9223372036854775808)
    if true {
        v20 := GenS2{F0: (true || (128 > 1)), F1: 9223372036854775807}
        _ = v20
    } else {
        if true {
            v2 += 199
            switch -9223372036854775808 {
            case 9223372036854775807:
                fmt.Println(v19, v3, v19)
            case 127:
                if true {
                    v3 = (v3 | (-(v3)))
                    switch ("" + (("a" + "hello") + "foo bar")) {
                    case "":
                        var v21 uint32 = 42
                        _ = v21
                    case "a":
                        switch (("" + "a") + "hello") {
                        case "\n\t":
                            v22 := GenS3{F0: true}
                            _ = v22
                        case "hello":
                            v23 := GenS3{F0: (65535 == 898)}
                            _ = v23
                        default:
                            switch "\n\t" {
                            case "a":
                                v3 -= -1
                            }
                        }
                    default:
                        var v24 uint32 = (v2 + (v2 * 42))
                        _ = v24
                    }
                    var v25 string = ((("\n\t" + "\n\t") + (("foo bar" + "hello") + "a")) + ("hello" + (("" + "hello") + "hello")))
                    _ = v25
                }
            case 32767:
                var v26 int8 = 42
                _ = v26
            }
            v19 -= (v19 - 255)
        }
        var v27 bool = (!(!false))
        for fi28 := 0; fi28 < 11; fi28++ {
            switch "" {
            case "":
                var v29 int8 = 42
                _ = v29
            }
            for fi30 := 0; fi30 < 12; fi30++ {
                var v31 int16 = 737
                fmt.Println(v27, v27)
                _ = fi30
                _ = v31
            }
            fmt.Println(v19)
            _ = fi28
        }
        _ = v27
    }
    var v32 uint64 = 1
    if (int(v19) == -188) {
        switch 680 {
        case 32767:
            v32 += (v32 - (^(v32)))
        }
    }
    v32 -= v32
    if false {
        fmt.Println(v2, v19, v19)
    } else {
        v33 := v1[2147483647]
        v34 := GenS2{F0: false, F1: v33}
        _ = v33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v32
}
