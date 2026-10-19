package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint32 }

func main() {
    v1 := []string{""}
    var v2 string = ""
    v3 := make(map[uint32]int)
    var v4 int32 = 2147483647
    var v5 int64 = -24
    var v6 float32 = 1e10
    for fi7 := 0; fi7 < 9; fi7++ {
        v4 = (-(v4))
        _ = fi7
    }
    v8 := GenS1{F0: 42, F1: 1}
    v9 := GenS1{F0: 795, F1: 195}
    switch "hello" {
    case "\n\t":
        if (-32768 > 730) {
            if false {
                v2 = ("foo bar" + v2)
                if false {
                    v4 = (^(v4))
                    var v10 uint = 27
                    _ = v10
                } else {
                    v11 := GenS1{F0: 1000, F1: 631}
                    var v12 int16 = 1000
                    _ = v11
                    _ = v12
                }
                var v13 int16 = -32768
                _ = v13
            } else {
                if false {
                    if true {
                        var v14 uint16 = 416
                        v15 := GenS1{F0: -1, F1: 1}
                        _ = v14
                        _ = v15
                    }
                    var v16 float32 = float32(3.14159)
                    if true {
                        if false {
                            var v17 bool = (v2 != v2)
                            _ = v17
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(v2, v16)
                            case "":
                                v4 -= (v4 / ((v4) | 1))
                            case "\n\t":
                                v18 := GenS1{F0: 42, F1: 0}
                                _ = v18
                            default:
                                v5 -= v5
                            }
                            v2 = v2
                            var v19 bool = true
                            _ = v19
                        }
                    }
                    _ = v16
                } else {
                    v20 := GenS1{F0: -1, F1: 4294967295}
                    var v21 uint = 255
                    v4 = 42
                    _ = v20
                    _ = v21
                }
                switch v2 {
                case "":
                    for fi22 := 0; fi22 < 14; fi22++ {
                        for fi23 := 0; fi23 < 14; fi23++ {
                            switch (fi22 & fi23) {
                            case -409:
                                var v24 string = ""
                                _ = v24
                            default:
                                v25 := GenS1{F0: -32768, F1: 0}
                                _ = v25
                            }
                            var v26 uint8 = 223
                            v6 = 0.1
                            _ = fi23
                            _ = v26
                        }
                        switch fi22 {
                        case -1:
                            var v27 float64 = (1.5 / 1e10)
                            _ = v27
                        case 32767:
                            switch fi22 {
                            case 504:
                                switch (("hello" + ("foo bar" + v2)) + v2) {
                                case "hello":
                                    _ = len(v3)
                                case "\n\t":
                                    var v28 uint32 = 2
                                    _ = v28
                                case "foo bar":
                                    var v29 uint16 = 827
                                    _ = v29
                                default:
                                    v4 += v4
                                }
                            case 2147483647:
                                var v30 uint16 = 273
                                _ = v30
                            default:
                                var v31 int16 = 66
                                _ = v31
                            }
                        default:
                            v5 -= (v5 >> ((199) & 63))
                        }
                        _ = fi22
                    }
                }
                switch v2 {
                case "foo bar":
                    var v32 uint16 = 1
                    _ = v32
                default:
                    for fi33 := 0; fi33 < 5; fi33++ {
                        for fi34 := 0; fi34 < 16; fi34++ {
                            fmt.Println(v5)
                            v6 = v6
                            _ = fi34
                        }
                        v35 := GenS2{F0: 42, F1: -128}
                        v36 := v3[0]
                        _ = fi33
                        _ = v35
                        _ = v36
                    }
                }
            }
        }
    case "foo bar":
        if (!false) {
            delete(v3, 4294967295)
        }
    }
    v6 = (v6 * (1e10 / v6))
    if false {
        v37 := GenS3{F0: 4294967295, F1: 100, F2: 366}
        _ = v37
    }
    if (18446744073709551615 > 42) {
        var v38 uint = 4294967295
        _ = v38
    }
    var v39 float32 = 3.14159
    var v40 int64 = 747
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
