package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int16 }

func main() {
    v1 := []int{}
    v2 := []uint32{1, 65535, 4294967295}
    var v3 uint = 65535
    v4 := make(map[string]bool)
    var v5 uint = (^(v3))
    var v6 int32 = -1
    if false {
        if len(v1) > 0 {
            v1[0] = 42
        }
        v3 = v5
    }
    var v7 int16 = 977
    if ((true || true) || (135 > 1)) {
        v8 := GenS1{F0: 128, F1: 1, F2: 2}
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := GenS1{F0: 142, F1: 0, F2: 0}
            var v11 float64 = (-2.25 / (1.5 * 0.1))
            if false {
                for fi12 := 0; fi12 < 8; fi12++ {
                    v13 := GenS1{F0: 121, F1: 2, F2: 127}
                    if (false || true) {
                        switch "hello" {
                        case "hello":
                            if (-9223372036854775808 != 42) {
                                var v14 uint = v5
                                v7 -= (v7 >> (((v3 ^ (v14 % ((v5) | 1)))) & 63))
                                _ = v14
                            }
                        case "":
                            v7 += v7
                            v3 += 1
                            v6 = v6
                        default:
                            fmt.Println(v5, v3)
                        }
                    } else {
                        var v15 float64 = v11
                        fmt.Println(v5)
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                }
            } else {
                var v16 uint = (v5 % ((807) | 1))
                var v17 uint8 = 0
                var v18 string = ""
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    } else {
        if (18446744073709551615 != 0) {
            v1 = append(v1, -2147483648)
        }
    }
    var v19 string = "foo bar"
    var v20 uint8 = 0
    if true {
        if true {
            fmt.Println(v7)
        } else {
            v7 += v7
            var v21 int16 = v7
            _ = v21
        }
        v22 := GenS1{F0: (^(v20)), F1: (v20 * 127), F2: (v20 % (((v20 * v20)) | 1))}
        _ = v22
    } else {
        v20 += 128
    }
    v5 = (v5 & (v5 % ((18446744073709551615) | 1)))
    var v23 float32 = 1e10
    _ = len(v1)
    v2 = append(v2, 65535)
    if false {
        v19 = "hello"
        var v24 uint32 = 601
        if (-809 == 481) {
            var v25 int8 = 100
            var v26 float64 = (1e10 - 1e10)
            _ = v25
            _ = v26
        }
        _ = v24
    } else {
        v1 = append(v1, 128)
        if true {
            for fi27 := 0; fi27 < 17; fi27++ {
                v20 += uint8(v3)
                _ = fi27
            }
        } else {
            var v28 uint16 = 1
            switch ((((v19 + v19) + (v19 + (((((v19 + v19) + v19) + (((v19 + ("" + ("" + ""))) + v19) + v19)) + v19) + (((("hello" + "") + "hello") + ("hello" + ("" + "hello"))) + "\n\t")))) + v19) + v19) {
            case "foo bar":
                fmt.Println(v20, v19, v3)
            case "hello":
                var v29 int16 = v7
                _ = v29
            case "":
                var v30 bool = (!(false || true))
                _ = v30
            default:
                fmt.Println(v19, v3, v23)
            }
            _ = v28
        }
    }
    switch v19 {
    case "foo bar":
        var v31 float64 = 0.1
        _ = v31
    case "":
        var v32 uint8 = v20
        _ = v32
    case "hello":
        v6 -= (v6 ^ (v6 + 241))
    default:
        for fi33 := 0; fi33 < 3; fi33++ {
            v3 = (v5 >> ((v3) & 63))
            _ = fi33
        }
    }
    v20 -= v20
    v19 = (((((("" + v19) + v19) + (v19 + v19)) + "\n\t") + ((v19 + v19) + v19)) + "a")
    if (!false) {
        var v34 bool = true
        switch v19 {
        case "":
            for fi35 := 0; fi35 < 20; fi35++ {
                var v36 int8 = 42
                var v37 int32 = v6
                _ = fi35
                _ = v36
                _ = v37
            }
        default:
            if v34 {
                v20 = (^(v20))
                var v38 uint32 = 65535
                v39 := GenS2{F0: 347, F1: (v7 - (v7 / ((32767) | 1)))}
                _ = v38
                _ = v39
            }
        }
        _ = v34
    } else {
        var v40 int = -558
        switch "hello" {
        case "foo bar":
            v41 := v4[v19]
            _ = v41
        case "a":
            var v42 int8 = -1
            _ = v42
        case "\n\t":
            v43 := GenS1{F0: 0, F1: (^(v20)), F2: (^(v20))}
            _ = v43
        }
        v44 := GenS2{F0: 0, F1: v7}
        _ = v40
        _ = v44
    }
    var v45 float64 = ((1e10 * float64(v23)) / float64(v23))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
