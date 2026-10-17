package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int }

func main() {
    var v1 uint64 = 1
    v2 := []string{"", ("hello" + "hello"), ("foo bar" + "hello"), "a"}
    var v3 float32 = float32((1e10 + 0.0))
    var v4 float64 = 1e10
    var v5 uint8 = 97
    var v6 int8 = -1
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: (fi7 % (((fi7 ^ (fi7 | 0))) | 1)), F1: (v1 - 0), F2: 255}
        if (0 >= 42) {
            v6 -= v6
            var v9 int = -170
            _ = v9
        }
        v1 += (v1 * v1)
        _ = fi7
        _ = v8
    }
    switch "foo bar" {
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    }
    v1 = v1
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS1{F0: -372, F1: v1, F2: 255}
        v12 := GenS1{F0: fi10, F1: v1, F2: 626}
        v1 = 18446744073709551615
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int32 = 42
    if (1.5 != (v3 * v3)) {
        var v14 uint = 1
        if true {
            v15 := GenS1{F0: 128, F1: 18446744073709551615, F2: (^(v14))}
            v5 = (v5 | v5)
            _ = v15
        } else {
            if (-1 == -9223372036854775808) {
                v14 = v14
                for fi16 := 0; fi16 < 13; fi16++ {
                    v17 := GenS2{F0: v5, F1: -745}
                    _ = fi16
                    _ = v17
                }
                fmt.Println(v14)
            }
        }
        v18 := GenS2{F0: v5, F1: int16(v14)}
        _ = v14
        _ = v18
    }
    v2 = append(v2, ("" + "a"))
    if (!(("\n\t" != ("" + "")) && false)) {
        if ((-(v13)) >= (v13 / ((-1) | 1))) {
            v6 = v6
        }
        fmt.Println(v4, v5)
        var v19 int16 = -625
        _ = v19
    } else {
        v20 := GenS3{F0: int(v1)}
        if false {
            switch ("hello" + "") {
            case "a":
                var v21 int64 = 42
                _ = v21
            case "":
                fmt.Println(v1, v5)
            case "foo bar":
                var v22 bool = true
                _ = v22
            default:
                switch -638 {
                case -506:
                    v13 = (v13 - (v13 & v13))
                case -16:
                    var v23 uint16 = 78
                    _ = v23
                }
            }
            for fi24 := 0; fi24 < 8; fi24++ {
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                if true {
                    v1 += v1
                    if true {
                        var v25 string = "\n\t"
                        _ = v25
                    }
                    v13 = -190
                } else {
                    var v26 uint16 = 65535
                    v27 := GenS1{F0: fi24, F1: 416, F2: 1}
                    _ = v26
                    _ = v27
                }
                var v28 uint64 = v1
                _ = fi24
                _ = v28
            }
            var v29 float64 = 0.0
            _ = v29
        }
        for fi30 := 0; fi30 < 2; fi30++ {
            if true {
                _ = len(v2)
                if true {
                    v3 += float32(v4)
                    for fi31 := 0; fi31 < 5; fi31++ {
                        var v32 uint64 = 4294967295
                        _ = fi31
                        _ = v32
                    }
                }
            }
            if (!true) {
                var v33 uint32 = 1
                _ = v33
            }
            for fi34 := 0; fi34 < 2; fi34++ {
                v6 = v6
                switch "\n\t" {
                case "":
                    var v35 int8 = v6
                    _ = v35
                case "foo bar":
                    var v36 uint8 = v5
                    _ = v36
                case "a":
                    v4 -= v4
                }
                var v37 uint8 = 2
                _ = fi34
                _ = v37
            }
            _ = fi30
        }
        _ = v20
    }
    if true {
        v38 := GenS2{F0: (v5 | (v5 % ((42) | 1))), F1: -457}
        var v39 bool = (65535 > 936)
        _ = v38
        _ = v39
    } else {
        _ = len(v2)
    }
    var v40 uint64 = v1
    var v41 uint8 = uint8(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
