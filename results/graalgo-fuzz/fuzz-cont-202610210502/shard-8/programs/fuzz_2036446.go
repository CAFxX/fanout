package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 bool; F1 int }
func genFn0(p0 uint64, p1 bool) int16 {
    v7 := []int{2147483647, 4294967295, int(p0)}
    p0 = (p0 % (((p0 & p0)) | 1))
    v8 := GenS1{F0: 65535, F1: ((p0 * 607) == p0)}
    if p1 {
        v9 := GenS1{F0: 65535, F1: (true && p1)}
        fmt.Println(p0)
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 498
}

func main() {
    var v1 int8 = 42
    v2 := []string{}
    var v3 int = 9223372036854775807
    var v4 int16 = int16(v1)
    var v5 uint64 = 665
    v6 := make(map[uint32]uint32)
    var v10 bool = false
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint8 = 1
        _ = fi11
        _ = v12
    }
    if v10 {
        if true {
            v1 += v1
            fmt.Println(v10)
        } else {
            v13 := GenS1{F0: 0, F1: v10}
            v6[65535] = 65535
            var v14 float32 = float32((-2.25 - (((3.14159 - (0.0 / (((0.0 / 1e10) - 3.14159) + 1e10))) - -2.25) / 1.0)))
            _ = v13
            _ = v14
        }
        var v15 string = ""
        for fi16 := 0; fi16 < 7; fi16++ {
            switch (v3 & v3) {
            case 32768:
                var v17 int = (fi16 - -1)
                _ = v17
            case -796:
                v10 = v10
            default:
                switch v3 {
                case 0:
                    for fi18 := 0; fi18 < 11; fi18++ {
                        if ((((^(v5)) != 542) && (978 == 355)) && (false && ((-9223372036854775808 > 9223372036854775807) && v10))) {
                            if (-1 < 275) {
                                v5 += v5
                            }
                            v19 := GenS1{F0: 65535, F1: v10}
                            var v20 int32 = 42
                            delete(v6, 42)
                            switch "hello" {
                            case "hello":
                                var v21 int8 = (^(v1))
                                _ = v21
                            case "\n\t":
                                var v22 float64 = (1e10 * 0.1)
                                _ = v22
                            }
                            _ = v19
                            _ = v20
                        } else {
                            v10 = (65535 != 299)
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v23 bool = ((v15 + v15) == ("hello" + "a"))
                        _ = fi18
                        _ = v23
                    }
                case -1:
                    _ = len(v2)
                case 9223372036854775807:
                    var v24 int8 = -128
                    _ = v24
                }
            }
            _ = fi16
        }
        _ = v15
    }
    var v25 int8 = (v1 / (((v1 % ((v1) | 1))) | 1))
    if ("foo bar" > "") {
        if (-128 != v1) {
            v6[363] = 2147483647
            if ((-2147483648 > -51) && (1 == (v25 | v25))) {
                delete(v6, 0)
            }
        } else {
            fmt.Println(v5, v4, v10)
        }
        if true {
            var v26 uint = 65535
            var v27 float32 = ((3.14159 * (0.1 / 0.1)) * (((((0.1 - (0.0 / 3.14159)) + float32((0.1 + 1e10))) - 0.0) + float32(0.1)) + 1e10))
            v26 -= (^(v26))
            _ = v26
            _ = v27
        } else {
            switch 735 {
            case 32767:
                v25 = v25
            }
            v10 = true
        }
        for fi28 := 0; fi28 < 13; fi28++ {
            for fi29 := 0; fi29 < 7; fi29++ {
                v25 = -128
                fmt.Println(v5, v4)
                switch fi29 {
                case -460:
                    v5 = 4294967295
                case -2147483648:
                    v3 = (fi29 / (((fi28 ^ fi29)) | 1))
                case 32767:
                    v30 := GenS1{F0: 255, F1: v10}
                    _ = v30
                }
                _ = fi29
            }
            _ = fi28
        }
    } else {
        delete(v6, 2147483647)
    }
    v10 = (301 != 42)
    if v10 {
        var v31 uint = 924
        v32 := GenS2{F0: v10, F1: (v3 & v3)}
        v33 := GenS2{F0: (821 >= 2), F1: v3}
        _ = v31
        _ = v32
        _ = v33
    }
    v34 := GenS2{F0: false, F1: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
