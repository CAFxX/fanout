package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int16; F1 int; F2 int8 }
type GenS5 struct { F0 int8; F1 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 42
    var v3 bool = true
    v4 := []string{"a", "a", "a"}
    var v5 int8 = 100
    var v6 int16 = 42
    _ = len(v1)
    v5 += v5
    if v3 {
        if (!(false || v3)) {
            if v3 {
                for fi7 := 0; fi7 < 10; fi7++ {
                    v4 = append(v4, "hello")
                    if ((v2 & (v2 + (v2 % ((v2) | 1)))) >= v2) {
                        v8 := GenS1{F0: 128}
                        fmt.Println(v2)
                        _ = v8
                    }
                    _ = fi7
                }
                fmt.Println(v3)
            } else {
                if (!v3) {
                    if (((3.14159 - (0.0 / -2.25)) / (0.0 + -2.25)) > (((0.0 + 1e10) * 1.5) * float64(1e10))) {
                        _ = len(v4)
                        for fi9 := 0; fi9 < 4; fi9++ {
                            var v10 int16 = v6
                            switch (fi9 ^ fi9) {
                            case 128:
                                v3 = false
                            case 127:
                                v1[2] = (fi9 | 9223372036854775807)
                            case -193:
                                for fi11 := 0; fi11 < 12; fi11++ {
                                    if (178 < 62) {
                                        var v12 float32 = 1.5
                                        _ = v12
                                    }
                                    _ = fi11
                                }
                            default:
                                if v3 {
                                    _ = len(v4)
                                    var v13 int16 = (v10 | v10)
                                    fmt.Println(v2, v2)
                                    _ = v13
                                } else {
                                    v14 := GenS2{F0: 1}
                                    _ = v14
                                }
                            }
                            _ = len(v4)
                            _ = fi9
                            _ = v10
                        }
                    } else {
                        var v15 int8 = v5
                        _ = v15
                    }
                    if v3 {
                        v16 := GenS3{F0: 0, F1: 4294967295}
                        v17 := GenS2{F0: 42}
                        _ = v16
                        _ = v17
                    }
                } else {
                    for fi18 := 0; fi18 < 9; fi18++ {
                        var v19 uint = 553
                        _ = fi18
                        _ = v19
                    }
                }
                var v20 uint16 = 1
                var v21 uint64 = 791
                _ = v20
                _ = v21
            }
        }
    }
    v22 := GenS1{F0: 243}
    v3 = ("foo bar" == "hello")
    v3 = (!true)
    v23 := GenS4{F0: (v6 << (((v2 | 255)) & 63)), F1: 32768, F2: v5}
    var v24 string = "\n\t"
    fmt.Println(v6)
    v25 := GenS5{F0: v5, F1: ("hello" + "a")}
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
