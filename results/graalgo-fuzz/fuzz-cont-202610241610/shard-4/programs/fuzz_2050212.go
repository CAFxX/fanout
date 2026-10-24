package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint16 = 865
    v2 := []uint32{898, 65535}
    var v3 uint8 = 255
    var v4 float32 = 1.5
    v5 := []uint32{42, 65535, 42, 2147483647}
    v6 := make(map[int]string)
    v4 = (v4 - v4)
    if (0 < 816) {
        var v7 float64 = (0.0 + float64((0.1 / float32((-2.25 + float64(0.1))))))
        if false {
            v2 = append(v2, 4294967295)
            v4 -= v4
        } else {
            v8 := GenS1{F0: -1}
            for fi9 := 0; fi9 < 17; fi9++ {
                for fi10 := 0; fi10 < 10; fi10++ {
                    v4 = float32((-2.25 * ((v7 / v7) + 3.14159)))
                    _ = fi10
                }
                v11 := GenS1{F0: 723}
                _ = fi9
                _ = v11
            }
            _ = v8
        }
        _ = v7
    } else {
        var v12 float64 = (1e10 / float64((v4 + (-2.25 / v4))))
        for fi13 := 0; fi13 < 17; fi13++ {
            fmt.Println(v1)
            for fi14 := 0; fi14 < 15; fi14++ {
                if ((^(v1)) >= v1) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v15 uint8 = 2
                    v12 += v12
                    _ = v15
                }
                v3 = (v3 & v3)
                _ = fi14
            }
            for fi16 := 0; fi16 < 15; fi16++ {
                switch fi16 {
                case -749:
                    var v17 bool = (v12 == 0.0)
                    _ = v17
                case -9223372036854775808:
                    var v18 float32 = 0.0
                    _ = v18
                case -437:
                    switch fi16 {
                    case 9223372036854775807:
                        fmt.Println(v1, fi13)
                    default:
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(fi13, v3, fi16)
                        case "foo bar":
                            v3 -= v3
                        }
                    }
                default:
                    fmt.Println(v1, v4, v4)
                }
                _ = fi16
            }
            _ = fi13
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            switch -1 {
            case 5:
                fmt.Println(v1, fi19)
            case 42:
                if len(v5) > 0 {
                    v5[0] = 4294967295
                }
            case 9223372036854775807:
                fmt.Println(v12)
            default:
                v3 += (v3 | (v3 | (^(v3))))
            }
            _ = fi19
        }
        _ = v12
    }
    v4 = ((1.5 * 1e10) / v4)
    v1 = v1
    if (-109 != 117) {
        v4 = (1e10 / v4)
        var v20 int8 = -94
        var v21 int64 = 0
        _ = v20
        _ = v21
    }
    v3 += v3
    v6[-9223372036854775808] = (("" + "a") + "foo bar")
    var v22 float32 = v4
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((false && (669 > 520)) && (true && true)) {
        var v23 uint8 = (v3 % (((v3 / (((v3 % (((v3 / ((v3) | 1))) | 1))) | 1))) | 1))
        _ = v23
    } else {
        if true {
            for fi24 := 0; fi24 < 10; fi24++ {
                v3 = 127
                var v25 uint16 = 963
                _ = fi24
                _ = v25
            }
        }
    }
    if true {
        v26 := GenS1{F0: 344}
        _ = v26
    }
    var v27 int8 = 102
    v28 := GenS1{F0: 1000}
    _ = len(v2)
    var v29 bool = true
    v1 = 491
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
