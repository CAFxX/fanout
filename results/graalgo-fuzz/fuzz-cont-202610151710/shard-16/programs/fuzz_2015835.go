package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint8; F1 int; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 198
    v3 := []string{("hello" + "foo bar"), ("hello" + "a"), "hello", ""}
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    v6 := make(map[uint32]int)
    v2 = (v2 ^ 1)
    v2 += v2
    if (("foo bar" <= "\n\t") || false) {
        if false {
            v3 = append(v3, "")
            v2 += v2
        } else {
            _ = len(v6)
            var v7 uint8 = (^(v2))
            var v8 uint64 = 1
            _ = v7
            _ = v8
        }
    }
    if false {
        for fi9 := 0; fi9 < 19; fi9++ {
            v2 -= (v2 / (((^(v2))) | 1))
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := GenS1{F0: -1, F1: -897}
                var v12 bool = false
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    }
    var v13 uint32 = 0
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS1{F0: -9223372036854775808, F1: 2147483647}
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 42, F1: 1}
    for fi17 := 0; fi17 < 9; fi17++ {
        if false {
            for fi18 := 0; fi18 < 7; fi18++ {
                v13 = v13
                _ = fi18
            }
        } else {
            switch "" {
            case "foo bar":
                var v19 uint64 = 42
                _ = v19
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v2 += v2
            }
            fmt.Println(fi17)
            v13 = (v13 / (((v13 | v13)) | 1))
        }
        if true {
            v20 := v1[fi17]
            v21 := GenS2{F0: 0, F1: 912, F2: v2}
            _ = v20
            _ = v21
        } else {
            var v22 float64 = ((1.5 - 3.14159) / 3.14159)
            var v23 bool = (!(1 != 0))
            switch -92 {
            case 9223372036854775807:
                if (false && (!true)) {
                    v13 += (v13 % ((42) | 1))
                } else {
                    v24 := GenS1{F0: -1, F1: 2147483647}
                    v25 := GenS2{F0: 1, F1: 9223372036854775807, F2: 177}
                    _ = v24
                    _ = v25
                }
            case 4294967295:
                v3 = append(v3, "a")
            default:
                if (1e10 != ((v22 + (v22 + v22)) * v22)) {
                    if (598 <= 2) {
                        if (0.1 != float32((1.5 * ((v22 + (((0.0 / 1.0) / float64(-2.25)) - v22)) + v22)))) {
                            switch "\n\t" {
                            case "hello":
                                var v26 int = (fi17 + (^(fi17)))
                                _ = v26
                            }
                            var v27 bool = (!false)
                            _ = v27
                        }
                        v22 += (1e10 + (v22 - ((-2.25 / (v22 / (v22 - ((v22 - -2.25) - v22)))) + (3.14159 - (v22 - v22)))))
                        var v28 int8 = 1
                        _ = v28
                    } else {
                        if true {
                            var v29 uint16 = 32767
                            _ = v29
                        }
                        var v30 bool = (v23 && (1.5 <= float32(v22)))
                        for fi31 := 0; fi31 < 4; fi31++ {
                            var v32 int = fi17
                            _ = fi31
                            _ = v32
                        }
                        _ = v30
                    }
                    for fi33 := 0; fi33 < 2; fi33++ {
                        if v23 {
                            fmt.Println(fi33)
                            var v34 uint64 = 2
                            _ = v34
                        } else {
                            v3 = append(v3, "\n\t")
                        }
                        if (!v23) {
                            _ = len(v3)
                            v35 := GenS3{F0: -2147483648, F1: 925}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v35
                        }
                        _ = fi33
                    }
                }
            }
            _ = v22
            _ = v23
        }
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
