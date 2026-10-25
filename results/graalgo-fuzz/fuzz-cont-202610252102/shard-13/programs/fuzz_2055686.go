package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint8; F1 int16; F2 string }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int16 = 1
            var v9 float32 = 0.0
            _ = fi7
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 uint64 = 42
            var v12 int = fi6
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi6
    }
    var v13 uint64 = 0
    v13 = v13
    _ = v13
    return 188
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -1
    var v3 uint = 240
    var v4 uint32 = 2147483647
    v5 := []int{}
    v3 -= (v3 * (^(v3)))
    switch 128 {
    case -2147483648:
        v2 = v2
    case 128:
        var v14 uint = (v3 >> ((v3) & 63))
        _ = v14
    default:
        if (266 >= 42) {
            switch 9223372036854775807 {
            case -422:
                switch 42 {
                case 9223372036854775807:
                    for fi15 := 0; fi15 < 5; fi15++ {
                        v2 = 42
                        var v16 uint64 = 0
                        _ = fi15
                        _ = v16
                    }
                case 2147483647:
                    var v17 uint = 18446744073709551615
                    _ = v17
                default:
                    var v18 uint = uint(v4)
                    _ = v18
                }
            case -9223372036854775808:
                v4 -= (^(v4))
            case 127:
                var v19 int32 = -1
                _ = v19
            default:
                for fi20 := 0; fi20 < 20; fi20++ {
                    if false {
                        v21 := GenS1{F0: (v2 & v2), F1: (^(v4)), F2: -1}
                        v2 += -1
                        _ = v21
                    } else {
                        var v22 int64 = -1
                        var v23 uint32 = 42
                        _ = v22
                        _ = v23
                    }
                    if (32767 <= 1) {
                        fmt.Println(v2)
                        v2 = v2
                    }
                    var v24 float32 = (-2.25 + -2.25)
                    _ = fi20
                    _ = v24
                }
            }
            v2 += -128
        }
    }
    var v25 int32 = -339
    _ = len(v1)
    if false {
        v25 = v25
        switch "hello" {
        case "\n\t":
            if true {
                var v26 bool = true
                _ = v26
            } else {
                var v27 int8 = v2
                for fi28 := 0; fi28 < 3; fi28++ {
                    for fi29 := 0; fi29 < 11; fi29++ {
                        v2 = v2
                        v5 = append(v5, fi29)
                        if true {
                            var v30 string = ""
                            var v31 float32 = 0.0
                            _ = v30
                            _ = v31
                        }
                        _ = fi29
                    }
                    _ = fi28
                }
                _ = v27
            }
        }
        switch "foo bar" {
        case "\n\t":
            var v32 int = 788
            _ = v32
        case "foo bar":
            _ = len(v5)
        case "a":
            fmt.Println(v4, v3, v4)
        }
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        var v34 float64 = float64(float32(-2.25))
        _ = fi33
        _ = v34
    }
    _ = len(v1)
    switch ("foo bar" + "hello") {
    case "\n\t":
        v2 = v2
    case "":
        if (int16(v2) >= 0) {
            var v35 int64 = 898
            var v36 int8 = -1
            _ = v35
            _ = v36
        }
    case "a":
        for fi37 := 0; fi37 < 6; fi37++ {
            fmt.Println(v25, v4)
            var v38 bool = (int8(fi37) == int8(v3))
            _ = fi37
            _ = v38
        }
    }
    var v39 float32 = (((0.1 / 3.14159) + 3.14159) + (0.0 + 3.14159))
    for fi40 := 0; fi40 < 19; fi40++ {
        var v41 uint64 = 1
        if (uint64(v4) >= v41) {
            fmt.Println(v39, v25, v4)
            if true {
                v42 := GenS1{F0: (v2 + v2), F1: genFn0(), F2: 1}
                v43 := GenS1{F0: v2, F1: 0, F2: 1}
                _ = v42
                _ = v43
            }
            for fi44 := 0; fi44 < 7; fi44++ {
                v4 -= 0
                v4 += v4
                v45 := GenS2{F0: 127, F1: -1, F2: "\n\t"}
                _ = fi44
                _ = v45
            }
        }
        _ = fi40
        _ = v41
    }
    switch "" {
    case "a":
        var v46 float64 = 3.14159
        _ = v46
    case "foo bar":
        if (806 >= 9223372036854775807) {
            _ = len(v5)
            var v47 int32 = v25
            _ = v47
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
