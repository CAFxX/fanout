package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 bool }

func main() {
    v1 := []string{"", ("hello" + ("" + ("\n\t" + ("\n\t" + "")))), (("hello" + "foo bar") + "\n\t")}
    v2 := make(map[uint32]uint32)
    var v3 int32 = 1
    var v4 uint16 = 0
    switch ("hello" + "a") {
    case "":
        var v5 bool = ((((42 > 119) && (v3 == v3)) || ("hello" == "a")) && ((((42 >= 0) && (35 == 947)) && false) || (((((float64((-2.25 * 0.0)) + (float64(-2.25) + 0.1)) - 1.5) <= -2.25) && false) && false)))
        _ = v5
    case "foo bar":
        var v6 uint8 = 114
        _ = v6
    default:
        var v7 float32 = 0.1
        _ = v7
    }
    v4 += 448
    var v8 uint8 = 127
    var v9 uint = 65535
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    if len(v1) > 0 {
        v1[0] = "a"
    }
    if true {
        v10 := v2[2147483647]
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := GenS1{F0: -1, F1: uint8(v9), F2: true}
            switch "\n\t" {
            case "":
                fmt.Println(fi11, v8)
            default:
                switch "\n\t" {
                case "":
                    v10 = (v10 / (((v10 * v10)) | 1))
                case "\n\t":
                    if (1000 == 0) {
                        var v13 float64 = (((1e10 - 1.5) / 1.0) / (3.14159 + (0.0 * 3.14159)))
                        _ = v13
                    } else {
                        switch fi11 {
                        case 2147483647:
                            switch fi11 {
                            case -1:
                                v3 = 2147483647
                            case 32768:
                                v14 := GenS1{F0: -128, F1: uint8(v4), F2: false}
                                _ = v14
                            }
                        case 4294967295:
                            var v15 int32 = -1
                            _ = v15
                        case -236:
                            switch fi11 {
                            case 128:
                                fmt.Println(fi11)
                            case -266:
                                for fi16 := 0; fi16 < 18; fi16++ {
                                    v17 := GenS1{F0: 127, F1: 255, F2: false}
                                    _ = fi16
                                    _ = v17
                                }
                            }
                        }
                    }
                }
            }
            var v18 int64 = 1
            _ = fi11
            _ = v12
            _ = v18
        }
        _ = v10
    } else {
        v4 += (v4 << ((v9) & 63))
        v8 += (v8 & (^(v8)))
        fmt.Println(v9, v9)
    }
    v4 = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
