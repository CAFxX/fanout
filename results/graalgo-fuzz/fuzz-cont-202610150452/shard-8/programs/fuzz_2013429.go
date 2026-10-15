package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 float64 = (1.5 - 1e10)
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    v4 := GenS1{F0: 1000, F1: 0}
    var v5 uint64 = 89
    if len(v2) > 0 {
        v2[0] = 42
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 uint = 255
        v5 += 1
        for fi8 := 0; fi8 < 19; fi8++ {
            switch "hello" {
            case "a":
                v7 += (^(v7))
            }
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    switch "" {
    case "\n\t":
        var v9 uint32 = 42
        _ = v9
    case "":
        var v10 uint64 = v5
        _ = v10
    case "hello":
        switch -172 {
        case -925:
            if false {
                v5 = (v5 ^ v5)
                v1 = v1
            }
        default:
            v5 -= (^(v5))
        }
    default:
        if len(v2) > 0 {
            v2[0] = uint32(v5)
        }
    }
    var v11 float64 = float64(1.5)
    if false {
        v11 += (float64((3.14159 + float32((float64(-2.25) - v1)))) / 1.0)
        if (false && (32767 < -821)) {
            var v12 int = 42
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 int32 = 2147483647
            if (true && ("foo bar" >= ("foo bar" + ("hello" + "a")))) {
                fmt.Println(v1, fi13)
                switch (-(fi13)) {
                case 1:
                    var v15 int32 = v14
                    _ = v15
                case 404:
                    var v16 int8 = 100
                    _ = v16
                default:
                    v5 += 42
                }
                var v17 int32 = (v14 >> ((819) & 63))
                _ = v17
            } else {
                if (false && (627 < 459)) {
                    if false {
                        fmt.Println(v5, fi13, fi13)
                        var v18 uint16 = 65535
                        for fi19 := 0; fi19 < 10; fi19++ {
                            var v20 uint8 = 1
                            _ = fi19
                            _ = v20
                        }
                        _ = v18
                    } else {
                        v21 := GenS2{F0: (v14 | (v14 % ((-1) | 1)))}
                        var v22 int8 = 100
                        for fi23 := 0; fi23 < 17; fi23++ {
                            v24 := GenS2{F0: 175}
                            if (fi23 > fi23) {
                                if (!(fi23 < (fi23 + fi23))) {
                                    v14 -= (v14 + 2147483647)
                                    var v25 uint16 = 84
                                    _ = v25
                                }
                                var v26 int16 = 0
                                _ = v26
                            } else {
                                v1 = (v11 / v11)
                            }
                            var v27 int32 = -2147483648
                            _ = fi23
                            _ = v24
                            _ = v27
                        }
                        _ = v21
                        _ = v22
                    }
                    var v28 int16 = 222
                    _ = v28
                }
            }
            var v29 string = ("a" + "")
            _ = fi13
            _ = v14
            _ = v29
        }
    }
    v30 := GenS3{F0: 220}
    for fi31 := 0; fi31 < 18; fi31++ {
        _ = len(v3)
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
