package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint8; F2 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{}
    var v3 float64 = float64((0.1 + 1.5))
    if true {
        v3 = (float64((1.5 + 0.1)) / v3)
        if (false || false) {
            var v4 uint8 = 47
            switch -2147483648 {
            case -528:
                var v5 int8 = -128
                _ = v5
            }
            _ = v4
        }
        var v6 bool = false
        _ = v6
    } else {
        v3 = v3
        switch "\n\t" {
        case "":
            for fi7 := 0; fi7 < 10; fi7++ {
                if (503 > -1) {
                    fmt.Println(fi7, v3)
                    fmt.Println(v3)
                } else {
                    for fi8 := 0; fi8 < 19; fi8++ {
                        if false {
                            _ = len(v1)
                        } else {
                            switch (fi8 * (fi8 << ((752) & 63))) {
                            case 9223372036854775807:
                                v3 -= (float64(((3.14159 / ((((0.0 * 0.1) / float32(v3)) / 0.1) - 0.0)) - 1.5)) - v3)
                            default:
                                v3 = ((v3 - -2.25) / (v3 * v3))
                            }
                        }
                        if (127 == -1) {
                            if (255 > 16) {
                                fmt.Println(fi8, fi7)
                                if (!false) {
                                    var v9 uint16 = 32767
                                    var v10 uint32 = 2147483648
                                    _ = v10
                                    _ = v9
                                } else {
                                    fmt.Println(fi8, v3)
                                    if ("\n\t" >= "foo bar") {
                                        v11 := GenS1{F0: -128, F1: 42, F2: ("" + ((("\n\t" + "hello") + "a") + "a"))}
                                        _ = v11
                                    }
                                }
                            }
                            var v12 uint32 = 4294967295
                            _ = v12
                        } else {
                            v3 += v3
                            switch fi7 {
                            case 32767:
                                v3 = v3
                            case 32768:
                                var v13 int16 = 189
                                _ = v13
                            default:
                                v14 := v1[2]
                                _ = v14
                            }
                        }
                        _ = fi8
                    }
                    v3 += (((-2.25 + (v3 - (v3 - 1e10))) - float64(1.5)) / (v3 / v3))
                    switch (("" + "foo bar") + "\n\t") {
                    case "\n\t":
                        for fi15 := 0; fi15 < 2; fi15++ {
                            var v16 int16 = 1000
                            _ = fi15
                            _ = v16
                        }
                    case "hello":
                        v3 = v3
                    case "foo bar":
                        if true {
                            var v17 float64 = v3
                            var v18 uint32 = 601
                            _ = v17
                            _ = v18
                        }
                    }
                }
                _ = len(v2)
                _ = fi7
            }
        case "\n\t":
            for fi19 := 0; fi19 < 16; fi19++ {
                v3 = 1.5
                _ = fi19
            }
        case "a":
            var v20 int16 = 32767
            _ = v20
        default:
            v21 := GenS1{F0: 0, F1: 4294967295, F2: "foo bar"}
            _ = v21
        }
    }
    v3 = (v3 * 0.0)
    switch "hello" {
    case "hello":
        if false {
            switch 0 {
            case -1:
                v22 := GenS1{F0: 19, F1: 42, F2: ""}
                _ = v22
            case -2147483648:
                var v23 float32 = float32((v3 / 1.5))
                _ = v23
            case 127:
                switch -1 {
                case -556:
                    if true {
                        var v24 uint32 = 524
                        _ = len(v1)
                        _ = v24
                    }
                }
            }
            var v25 bool = true
            _ = v25
        } else {
            for fi26 := 0; fi26 < 19; fi26++ {
                v27 := GenS2{F0: -199}
                _ = fi26
                _ = v27
            }
        }
    case "foo bar":
        var v28 uint = 65535
        _ = v28
    }
    v3 = float64(((1e10 * 0.0) + -2.25))
    var v29 float32 = 0.0
    v3 = ((-2.25 / 1e10) - 0.0)
    _ = len(v2)
    v30 := v1[65535]
    var v31 bool = (!false)
    var v32 bool = (!(!(((v3 + (v3 - (-2.25 + 1.5))) >= ((((0.1 * (((v3 + v3) + float64(v29)) * v3)) - 1.5) - ((v3 * ((v3 * v3) - v3)) / v3)) * float64(float32((v3 + (1e10 * 1e10)))))) && v31)))
    for fi33 := 0; fi33 < 1; fi33++ {
        v34 := GenS3{F0: -9223372036854775808, F1: 129, F2: 42}
        v35 := GenS2{F0: -32768}
        _ = fi33
        _ = v34
        _ = v35
    }
    v29 = float32(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
