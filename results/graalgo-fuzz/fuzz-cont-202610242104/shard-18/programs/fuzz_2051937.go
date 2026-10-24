package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int64; F1 int; F2 int16 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint16 }

func main() {
    var v1 int = 816
    v2 := make(map[string]uint32)
    var v3 uint = 2
    var v4 string = "hello"
    var v5 uint = 1
    v3 -= v5
    v4 = v4
    v6 := []int{v1, int(v5)}
    switch ("" + ((v4 + (v4 + v4)) + v4)) {
    case "foo bar":
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint64 = 766
            _ = fi7
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 17; fi9++ {
            _ = len(v2)
            switch int(v3) {
            case 32768:
                fmt.Println(v3)
            case -1:
                v1 = fi9
            default:
                v5 = 1
            }
            _ = fi9
        }
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        switch v4 {
        case "":
            if false {
                v4 = (("a" + "foo bar") + v4)
                v11 := GenS1{F0: (-(v1))}
                var v12 uint64 = 94
                _ = v11
                _ = v12
            }
        case "\n\t":
            var v13 int16 = 1000
            _ = v13
        default:
            var v14 float32 = (float32(1.5) / (3.14159 - 0.0))
            _ = v14
        }
        if false {
            for fi15 := 0; fi15 < 7; fi15++ {
                v3 = 4294967295
                v1 += fi10
                if (v5 >= v5) {
                    fmt.Println(v1, v5, v1)
                    v2[""] = 594
                } else {
                    if true {
                        var v16 int32 = -1
                        var v17 int32 = 529
                        v18 := GenS2{F0: -1, F1: 42}
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        var v19 string = "foo bar"
                        v5 += (^(v5))
                        if false {
                            v4 = (v19 + v4)
                            _ = len(v2)
                            var v20 float64 = (((((3.14159 - 0.1) / float64(((1e10 * 0.1) - -2.25))) / 1e10) - (0.0 / 1e10)) + 0.0)
                            _ = v20
                        } else {
                            v5 += v5
                        }
                        v4 = ((((v4 + ("foo bar" + "")) + (v19 + v19)) + v4) + v19)
                        _ = v19
                    }
                }
                _ = fi15
            }
            switch v1 {
            case -9223372036854775808:
                v3 = v5
            default:
                v21 := GenS1{F0: -191}
                _ = v21
            }
        } else {
            var v22 uint64 = 0
            v23 := GenS3{F0: -448, F1: (^(fi10)), F2: 1}
            _ = v22
            _ = v23
        }
        v5 = v5
        _ = fi10
    }
    switch (v4 + v4) {
    case "":
        var v24 int64 = 9223372036854775807
        _ = v24
    case "foo bar":
        v3 = (v3 & 18446744073709551615)
    case "hello":
        for fi25 := 0; fi25 < 7; fi25++ {
            var v26 uint8 = 165
            var v27 float64 = (((0.0 * 1.5) * (float64(float32(0.1)) * (1e10 + ((1.5 / 3.14159) - 3.14159)))) * (1e10 + 1e10))
            v27 = v27
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    var v28 int = v1
    v29 := GenS4{F0: true, F1: 0, F2: 313}
    if (0.1 >= float32(1e10)) {
        v2[(v4 + (v4 + v4))] = 1
        for fi30 := 0; fi30 < 19; fi30++ {
            var v31 uint16 = 0
            v4 = (v4 + (v4 + v4))
            var v32 uint8 = 0
            _ = fi30
            _ = v31
            _ = v32
        }
        if (!(!(481 != 65535))) {
            v5 = v3
        }
    } else {
        fmt.Println(v28)
        v1 += -9223372036854775808
    }
    _ = len(v6)
    for fi33 := 0; fi33 < 2; fi33++ {
        var v34 int64 = -688
        v35 := GenS3{F0: (v34 & (v34 & (-(v34)))), F1: -1, F2: 0}
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
