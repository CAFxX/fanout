package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint = 2
    v2 := make(map[int]bool)
    var v3 int16 = -1
    var v4 uint = 18446744073709551615
    var v5 uint64 = 2
    var v6 uint16 = 255
    for fi7 := 0; fi7 < 19; fi7++ {
        if (true && (true && (((0.1 * 0.0) + -2.25) != 0.1))) {
            v8 := GenS1{F0: "foo bar", F1: 362, F2: (v1 & 23)}
            _ = v8
        } else {
            v3 = (v3 * (v3 << (((v1 >> ((255) & 63))) & 63)))
            fmt.Println(v4)
            switch ("hello" + "hello") {
            case "":
                for fi9 := 0; fi9 < 11; fi9++ {
                    if (!true) {
                        v6 = v6
                        v4 = 1
                    }
                    switch -1 {
                    case 4294967295:
                        v4 = v4
                    case 32768:
                        switch "hello" {
                        case "foo bar":
                            v6 = (v6 / ((uint16(v1)) | 1))
                        case "hello":
                            var v10 int64 = 0
                            _ = v10
                        default:
                            var v11 uint16 = (v6 + (v6 & v6))
                            _ = v11
                        }
                    default:
                        var v12 uint8 = 165
                        _ = v12
                    }
                    _ = fi9
                }
            case "hello":
                var v13 uint32 = 695
                _ = v13
            case "a":
                for fi14 := 0; fi14 < 1; fi14++ {
                    if true {
                        v15 := GenS2{F0: v6, F1: 0}
                        if (("foo bar" == "foo bar") || false) {
                            v1 = (v4 ^ v4)
                            fmt.Println(v6, fi7, fi14)
                        } else {
                            v16 := GenS2{F0: v6, F1: 896}
                            var v17 int32 = -1
                            _ = v16
                            _ = v17
                        }
                        _ = v15
                    }
                    _ = fi14
                }
            }
        }
        for fi18 := 0; fi18 < 9; fi18++ {
            v4 += v4
            _ = fi18
        }
        _ = fi7
    }
    v19 := GenS1{F0: "\n\t", F1: 42, F2: (v1 / ((v4) | 1))}
    _ = len(v2)
    if (v6 == v6) {
        if true {
            switch "hello" {
            case "foo bar":
                var v20 uint = v1
                _ = v20
            case "a":
                fmt.Println(v4, v5, v4)
            default:
                fmt.Println(v5)
            }
            for fi21 := 0; fi21 < 13; fi21++ {
                v22 := GenS3{F0: uint64(v6)}
                v4 += 65535
                var v23 float32 = 3.14159
                _ = fi21
                _ = v22
                _ = v23
            }
        } else {
            v24 := []int{-1, -1, 9223372036854775807, 127}
            if len(v24) > 0 {
                v24[0] = 944
            }
            switch 4294967295 {
            case 2147483647:
                v25 := GenS2{F0: (v6 - (^(v6))), F1: 0}
                _ = v25
            default:
                for fi26 := 0; fi26 < 13; fi26++ {
                    fmt.Println(v6, v5, v3)
                    var v27 float64 = 3.14159
                    var v28 float64 = v27
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            _ = v24
        }
        for fi29 := 0; fi29 < 11; fi29++ {
            fmt.Println(v5)
            _ = fi29
        }
    }
    var v30 bool = false
    var v31 bool = v30
    v32 := GenS2{F0: v6, F1: 42}
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
