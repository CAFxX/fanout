package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int16 }

func main() {
    var v1 float64 = (1e10 + 1e10)
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v4 int64 = -321
    v5 := []string{}
    v6 := []string{"a", "", "a", "\n\t"}
    delete(v2, "\n\t")
    v4 = (v4 | (v4 ^ v4))
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            if false {
                var v9 uint32 = 665
                v10 := GenS1{F0: 255, F1: 42}
                var v11 bool = ("" > "\n\t")
                _ = v10
                _ = v11
                _ = v9
            } else {
                if (true || ((true && false) && true)) {
                    v1 = v1
                    var v12 float32 = 3.14159
                    v1 -= v1
                    _ = v12
                } else {
                    var v13 uint8 = 128
                    fmt.Println(fi7, v1)
                    _ = v13
                }
                fmt.Println(fi8)
            }
            _ = fi8
        }
        _ = fi7
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS1{F0: 361, F1: 65535}
        v16 := GenS1{F0: 4294967295, F1: 2}
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 2, F1: 255}
    var v18 float64 = 3.14159
    v19 := GenS2{F0: 42}
    v20 := GenS1{F0: 4294967295, F1: 0}
    var v21 int = 0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        v4 = v4
        if (!false) {
            var v23 int = v21
            fmt.Println(v4, v18)
            _ = v23
        } else {
            if true {
                if false {
                    for fi24 := 0; fi24 < 13; fi24++ {
                        v4 = v4
                        _ = fi24
                    }
                    var v25 uint = 0
                    if false {
                        if (1.5 > 1e10) {
                            var v26 int64 = v4
                            var v27 int8 = -121
                            switch (v21 / (((fi22 * -541)) | 1)) {
                            case 2147483647:
                                v21 = (v21 / (((v21 % (((-(v21))) | 1))) | 1))
                            case -853:
                                for fi28 := 0; fi28 < 17; fi28++ {
                                    var v29 uint16 = 42
                                    var v30 uint = (v25 - v25)
                                    v5 = append(v5, "foo bar")
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                            case -9223372036854775808:
                                fmt.Println(fi22, fi22, v25)
                            }
                            _ = v26
                            _ = v27
                        }
                        fmt.Println(v4, fi22, v21)
                    }
                    _ = v25
                } else {
                    var v31 bool = (1000 < 0)
                    _ = v31
                }
            } else {
                v1 += 0.0
                v32 := GenS3{F0: 4294967295, F1: 32767}
                fmt.Println(v21)
                _ = v32
            }
        }
        var v33 bool = false
        _ = fi22
        _ = v33
    }
    if (false || (0.1 != v18)) {
        v34 := GenS2{F0: -2147483648}
        fmt.Println(v18)
        v4 -= -9223372036854775808
        _ = v34
    } else {
        v18 = 0.0
        if (true || (42 < 0)) {
            v18 = v1
            switch 0 {
            case 9223372036854775807:
                switch -584 {
                case 4294967295:
                    v4 -= v4
                case -1:
                    fmt.Println(v1)
                case 877:
                    v18 += -2.25
                default:
                    v4 -= (^(v4))
                }
            case -792:
                _ = len(v5)
            default:
                var v35 uint8 = 0
                _ = v35
            }
            var v36 int32 = 840
            _ = v36
        } else {
            if (42 == -2147483648) {
                var v37 float32 = (1e10 - (1.5 + float32(0.1)))
                if ((v4 * v4) < v4) {
                    v37 += 1e10
                }
                for fi38 := 0; fi38 < 11; fi38++ {
                    switch v21 {
                    case 128:
                        v21 -= (^(v21))
                    case 1:
                        v18 = float64(1e10)
                    case 4294967295:
                        v39 := GenS2{F0: -1}
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 1; fi40++ {
                        fmt.Println(v21)
                        _ = fi40
                    }
                    var v41 string = "a"
                    _ = fi38
                    _ = v41
                }
                _ = v37
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
