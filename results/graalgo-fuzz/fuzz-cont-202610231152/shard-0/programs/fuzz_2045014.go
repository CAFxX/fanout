package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 bool }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int16; F1 int64 }
func genFn0() string {
    v7 := []int{-1}
    var v8 uint = 974
    v8 += v8
    v8 = v8
    _ = v7
    _ = v8
    return "foo bar"
}
func genFn1(p0 bool) int {
    switch 1 {
    case 1:
        if (233 >= 0) {
            v9 := GenS1{F0: "a", F1: -1, F2: (4294967295 <= 891)}
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 int = -1
                v12 := GenS1{F0: ("\n\t" + "a"), F1: 820, F2: true}
                p0 = (196 >= 481)
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        } else {
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := GenS2{F0: -1, F1: 100}
                fmt.Println(fi13, p0)
                _ = fi13
                _ = v14
            }
        }
    default:
        var v15 uint8 = 111
        _ = v15
    }
    p0 = p0
    fmt.Println(p0)
    var v16 bool = (-838 != 628)
    _ = p0
    _ = v16
    return 9223372036854775807
}

func main() {
    var v1 float64 = (float64((0.0 - -2.25)) + 0.0)
    v2 := []uint32{2147483648, 65535, 1, 65535}
    var v3 int32 = 2147483647
    var v4 int8 = 127
    var v5 int16 = 1
    var v6 int8 = v4
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 = v6
    var v17 float64 = 0.1
    if false {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        if true {
            v2 = append(v2, 955)
        } else {
            var v18 int64 = 919
            _ = v18
        }
        v6 = (v6 * v4)
    }
    var v19 uint16 = 1
    switch 511 {
    case 32767:
        v1 -= v17
    case -161:
        if false {
            v19 -= (v19 + v19)
        } else {
            switch 6 {
            case 1:
                fmt.Println(v1, v3, v17)
            default:
                switch "\n\t" {
                case "\n\t":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "":
                    var v20 int8 = 100
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 6; fi21++ {
                        fmt.Println(v17)
                        _ = len(v2)
                        switch fi21 {
                        case 175:
                            if (128 >= 128) {
                                if false {
                                    v19 = v19
                                    switch fi21 {
                                    case 1:
                                        v4 = v6
                                    case -607:
                                        v2 = append(v2, 4294967295)
                                    default:
                                        var v22 uint16 = v19
                                        _ = v22
                                    }
                                    v4 += (v6 << ((42) & 63))
                                } else {
                                    v4 += 0
                                }
                            }
                        case -690:
                            for fi23 := 0; fi23 < 19; fi23++ {
                                v17 = v17
                                _ = len(v2)
                                var v24 string = "foo bar"
                                v25 := GenS1{F0: v24, F1: v3, F2: (v3 >= v3)}
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                        default:
                            if true {
                                _ = len(v2)
                                if (1.5 < (-2.25 * (0.1 / 1.0))) {
                                    if len(v2) > 0 {
                                        v2[0] = 40
                                    }
                                }
                            } else {
                                _ = len(v2)
                            }
                        }
                        _ = fi21
                    }
                }
            }
            v19 = 32767
        }
    default:
        v26 := GenS3{F0: v5, F1: 49}
        _ = v26
    }
    var v27 int16 = v5
    v5 = (v5 - v27)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
