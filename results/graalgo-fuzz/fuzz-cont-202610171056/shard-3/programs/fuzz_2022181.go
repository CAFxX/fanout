package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 string }

func main() {
    var v1 bool = true
    v2 := []int{-2147483648, 128}
    v3 := make(map[uint32]bool)
    var v4 uint16 = 1
    v5 := make(map[uint32]bool)
    var v6 int32 = 42
    v4 -= (v4 << ((175) & 63))
    switch "" {
    case "hello":
        for fi7 := 0; fi7 < 13; fi7++ {
            v4 = (^(v4))
            for fi8 := 0; fi8 < 16; fi8++ {
                switch (fi7 >> ((0) & 63)) {
                case 0:
                    v6 += v6
                case 42:
                    fmt.Println(v1, v4, v4)
                    v2 = append(v2, 4294967295)
                    var v9 float32 = (0.0 / 1.0)
                    _ = v9
                default:
                    var v10 uint32 = 0
                    switch fi8 {
                    case 128:
                        v10 = (^(v10))
                    default:
                        if true {
                            var v11 uint = 646
                            v12 := GenS1{F0: "", F1: int32(v10), F2: ("" + "a")}
                            _ = v11
                            _ = v12
                        }
                    }
                    switch (("" + "a") + "a") {
                    case "":
                        if len(v2) > 0 {
                            v2[0] = (fi8 >> ((752) & 63))
                        }
                    }
                    _ = v10
                }
                fmt.Println(v4, fi7, fi8)
                _ = fi8
            }
            _ = fi7
        }
    case "\n\t":
        if (0.1 != (0.1 / 1.0)) {
            if v1 {
                var v13 float32 = 0.1
                v1 = true
                for fi14 := 0; fi14 < 15; fi14++ {
                    if (-9223372036854775808 > 9223372036854775807) {
                        var v15 uint8 = 145
                        fmt.Println(v4, v15, fi14)
                        _ = v15
                    }
                    if false {
                        delete(v5, 65535)
                        for fi16 := 0; fi16 < 9; fi16++ {
                            var v17 int64 = 144
                            fmt.Println(v6, v17)
                            _ = fi16
                            _ = v17
                        }
                    } else {
                        if (v4 > 42) {
                            if (0 <= 42) {
                                for fi18 := 0; fi18 < 10; fi18++ {
                                    v19 := GenS1{F0: ("a" + ""), F1: v6, F2: "hello"}
                                    v4 = (v4 / ((v4) | 1))
                                    _ = fi18
                                    _ = v19
                                }
                            }
                            switch -9223372036854775808 {
                            case 2147483647:
                                if false {
                                    for fi20 := 0; fi20 < 7; fi20++ {
                                        var v21 int16 = -614
                                        v13 += (1.5 / float32(-2.25))
                                        _ = fi20
                                        _ = v21
                                    }
                                }
                            case 32:
                                switch fi14 {
                                case 127:
                                    switch fi14 {
                                    case 128:
                                        var v22 float32 = float32(1.5)
                                        _ = v22
                                    case 1:
                                        v5[205] = (42 < 2)
                                    default:
                                        v23 := GenS1{F0: "\n\t", F1: v6, F2: "hello"}
                                        _ = v23
                                    }
                                case -965:
                                    v13 = v13
                                case 9223372036854775807:
                                    var v24 int32 = (v6 ^ 42)
                                    _ = v24
                                }
                            default:
                                var v25 uint8 = 0
                                _ = v25
                            }
                        } else {
                            v6 += 42
                            switch ("\n\t" + "") {
                            case "hello":
                                var v26 float32 = (float32(float64(v13)) * v13)
                                _ = v26
                            case "\n\t":
                                fmt.Println(v13, v6, fi14)
                            }
                            v13 -= v13
                        }
                    }
                    fmt.Println(v6, v1)
                    _ = fi14
                }
                _ = v13
            }
            var v27 uint64 = 42
            v1 = (0.0 < 1.5)
            _ = v27
        }
    case "a":
        v5[817] = (!true)
    default:
        var v28 int32 = (v6 + (v6 % (((v6 | (v6 ^ v6))) | 1)))
        _ = v28
    }
    v1 = (644 != 1000)
    v1 = v1
    var v29 int64 = 0
    var v30 float64 = 1.5
    for fi31 := 0; fi31 < 15; fi31++ {
        v3[991] = false
        _ = fi31
    }
    _ = len(v5)
    v1 = (true && (65535 > 65535))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
