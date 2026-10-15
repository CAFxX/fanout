package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int32) int8 {
    for fi4 := 0; fi4 < 2; fi4++ {
        switch "\n\t" {
        case "":
            v5 := []int{fi4}
            _ = v5
        case "hello":
            p0 += p0
        }
        _ = fi4
    }
    p0 = (p0 + p0)
    _ = p0
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (0.1 + (1.5 - (float32(1e10) + 0.0)))
    var v3 int16 = 745
    v6 := []int{702, -688, 9223372036854775807}
    if false {
        if len(v6) > 0 {
            v6[0] = 2147483647
        }
        if false {
            var v7 uint = 65535
            for fi8 := 0; fi8 < 14; fi8++ {
                if true {
                    fmt.Println(v3)
                }
                for fi9 := 0; fi9 < 1; fi9++ {
                    var v10 int = (fi9 % ((fi9) | 1))
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
            _ = v7
        } else {
            for fi11 := 0; fi11 < 13; fi11++ {
                fmt.Println(v3, fi11)
                v12 := GenS1{F0: 128}
                fmt.Println(v3, v3, fi11)
                _ = fi11
                _ = v12
            }
            if false {
                v13 := GenS1{F0: 0}
                for fi14 := 0; fi14 < 11; fi14++ {
                    v2 += (v2 / (v2 + v2))
                    v3 = (v3 % (((^(v3))) | 1))
                    v3 -= (^(v3))
                    _ = fi14
                }
                if false {
                    if true {
                        var v15 float32 = (v2 + (v2 / (v2 - ((((((0.0 / float32(-2.25)) + -2.25) + ((v2 / (v2 * (float32(1e10) / v2))) + v2)) - v2) / v2) / v2))))
                        switch "" {
                        case "":
                            delete(v1, -204)
                        case "a":
                            for fi16 := 0; fi16 < 19; fi16++ {
                                v3 += v3
                                _ = fi16
                            }
                        case "foo bar":
                            if true {
                                switch "foo bar" {
                                case "hello":
                                    v17 := v1[128]
                                    _ = v17
                                }
                                for fi18 := 0; fi18 < 18; fi18++ {
                                    v3 = v3
                                    v15 += v2
                                    _ = fi18
                                }
                                var v19 uint32 = 42
                                _ = v19
                            }
                        default:
                            v3 = v3
                        }
                        _ = v15
                    }
                    v20 := v1[-2147483648]
                    if false {
                        if v20 {
                            v3 = (-(v3))
                        }
                        for fi21 := 0; fi21 < 14; fi21++ {
                            var v22 uint16 = 718
                            _ = fi21
                            _ = v22
                        }
                        v23 := GenS1{F0: 42}
                        _ = v23
                    }
                    _ = v20
                } else {
                    v24 := GenS2{F0: -27}
                    _ = v24
                }
                _ = v13
            }
        }
    }
    switch "hello" {
    case "a":
        v3 = (v3 | v3)
    case "hello":
        fmt.Println(v3, v3)
    case "foo bar":
        for fi25 := 0; fi25 < 20; fi25++ {
            if (-433 < 0) {
                v2 = (v2 + 1e10)
            } else {
                v3 = (-(v3))
                var v26 float64 = (1e10 + float64((((1e10 * v2) * (0.1 * 0.0)) * 0.0)))
                _ = v26
            }
            if true {
                _ = len(v6)
                _ = len(v1)
                if true {
                    v3 -= (v3 << ((1) & 63))
                    v27 := GenS2{F0: 42}
                    _ = v27
                }
            } else {
                v28 := GenS1{F0: 255}
                _ = v28
            }
            _ = fi25
        }
    default:
        v2 = (v2 / 1.0)
    }
    var v29 uint = 354
    if (((v29 % (((v29 ^ v29)) | 1)) >= (v29 - (v29 * v29))) || false) {
        fmt.Println(v2, v29, v29)
        fmt.Println(v29, v3, v29)
    }
    fmt.Println(v2, v3, v3)
    v29 += v29
    v2 = 1.5
    var v30 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v6
}
