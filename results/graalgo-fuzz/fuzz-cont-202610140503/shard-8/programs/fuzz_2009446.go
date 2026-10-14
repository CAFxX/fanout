package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := []uint32{768, 813, 4294967295, 1}
    v2 := make(map[int]bool)
    var v3 float32 = (3.14159 * 1.5)
    v4 := make(map[int]int)
    v5 := []string{("\n\t" + ("" + ""))}
    switch 244 {
    case -94:
        v3 -= v3
    case 32768:
        var v6 int16 = -351
        _ = v6
    }
    if false {
        var v7 uint32 = 42
        if (127 < -101) {
            v8 := v4[-339]
            var v9 uint64 = 42
            for fi10 := 0; fi10 < 12; fi10++ {
                for fi11 := 0; fi11 < 20; fi11++ {
                    var v12 uint32 = (^(v7))
                    v13 := GenS1{F0: (4 != 9223372036854775807), F1: false, F2: -51}
                    v3 -= (float32((0.0 / (((0.0 - 0.1) - (0.0 - 1.5)) * float64((v3 / 0.1))))) * (0.1 * v3))
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                fmt.Println(v8, v8, fi10)
                _ = fi10
            }
            _ = v8
            _ = v9
        } else {
            v7 = 1
        }
        fmt.Println(v3)
        _ = v7
    }
    switch "" {
    case "\n\t":
        if (1000 == -781) {
            if (!false) {
                v14 := GenS1{F0: false, F1: (((v3 / (-2.25 * v3)) * float32((float64(v3) * 1.5))) == (v3 + v3)), F2: -1}
                _ = v14
            } else {
                v3 -= v3
                v15 := GenS1{F0: (4294967295 < 2147483647), F1: (!true), F2: -114}
                switch ("foo bar" + "hello") {
                case "hello":
                    if (206 > 18446744073709551615) {
                        if false {
                            for fi16 := 0; fi16 < 12; fi16++ {
                                fmt.Println(v3)
                                if (-520 != 1) {
                                    for fi17 := 0; fi17 < 3; fi17++ {
                                        v18 := v2[fi16]
                                        _ = fi17
                                        _ = v18
                                    }
                                    switch (fi16 >> ((1) & 63)) {
                                    case -9223372036854775808:
                                        switch "" {
                                        case "a":
                                            v3 = -2.25
                                        case "":
                                            switch "\n\t" {
                                            case "hello":
                                                var v19 int16 = 42
                                                _ = v19
                                            case "\n\t":
                                                v20 := GenS1{F0: ("hello" < "a"), F1: (!false), F2: -93}
                                                _ = v20
                                            case "foo bar":
                                                var v21 int64 = -52
                                                _ = v21
                                            default:
                                                var v22 uint = 65535
                                                _ = v22
                                            }
                                        }
                                    case 0:
                                        fmt.Println(v3, v3)
                                    case 540:
                                        v3 = 0.0
                                    }
                                    var v23 uint = 0
                                    _ = v23
                                }
                                v3 -= v3
                                _ = fi16
                            }
                        } else {
                            v3 = 0.1
                        }
                    }
                }
                _ = v15
            }
            var v24 uint64 = 53
            _ = v24
        } else {
            v25 := v4[2147483647]
            v26 := GenS2{F0: v25}
            var v27 int64 = -9223372036854775808
            _ = v25
            _ = v26
            _ = v27
        }
    case "a":
        v28 := GenS3{F0: 943}
        _ = v28
    default:
        v29 := v4[0]
        _ = v29
    }
    fmt.Println(v3)
    v3 = v3
    var v30 uint64 = 18446744073709551615
    v30 = 2
    var v31 float32 = (v3 + (float32((1e10 / 0.1)) - v3))
    v31 = float32(float64(((v3 + (v3 / (v31 + v3))) / (v31 / (v3 - v31)))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
