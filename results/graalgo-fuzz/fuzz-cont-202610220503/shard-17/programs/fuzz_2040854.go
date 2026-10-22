package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint; F1 int8; F2 uint8 }
func genFn0(p0 uint, p1 int16, p2 string) int {
    var v7 uint16 = 42
    v8 := GenS1{F0: false}
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 uint32 = 636
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -35
}
func genFn1(p0 uint16) uint8 {
    v11 := GenS1{F0: (42 >= 325)}
    p0 += (^(p0))
    if true {
        fmt.Println(p0)
        for fi12 := 0; fi12 < 12; fi12++ {
            p0 -= (p0 - p0)
            v13 := []int{genFn0(42, 0, "")}
            _ = fi12
            _ = v13
        }
    } else {
        p0 += (p0 << ((0) & 63))
        if ((true || false) && (!false)) {
            for fi14 := 0; fi14 < 6; fi14++ {
                switch -587 {
                case 4294967295:
                    p0 = (p0 & (^(p0)))
                case 845:
                    for fi15 := 0; fi15 < 9; fi15++ {
                        p0 = 0
                        v16 := GenS1{F0: true}
                        _ = fi15
                        _ = v16
                    }
                default:
                    p0 -= (p0 % (((p0 % ((65535) | 1))) | 1))
                }
                _ = fi14
            }
            fmt.Println(p0)
        } else {
            v17 := GenS1{F0: false}
            _ = v17
        }
        v18 := []int{2147483647, -1, 9223372036854775807, -38}
        _ = v18
    }
    _ = p0
    _ = v11
    return 1
}

func main() {
    var v1 uint = 255
    var v2 uint64 = 63
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    v5 := []uint32{65535, 4294967295, 65535}
    var v6 int8 = 42
    v2 += (v2 & (v2 % ((v2) | 1)))
    v1 = v1
    if true {
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            for fi19 := 0; fi19 < 7; fi19++ {
                for fi20 := 0; fi20 < 13; fi20++ {
                    switch ("a" + "\n\t") {
                    case "":
                        var v21 string = "\n\t"
                        _ = v21
                    case "\n\t":
                        var v22 uint64 = 0
                        _ = v22
                    case "a":
                        v23 := GenS1{F0: true}
                        _ = v23
                    default:
                        if false {
                            v6 = -102
                            v2 = v2
                            if true {
                                fmt.Println(fi19, fi20, v2)
                                v6 = 0
                                fmt.Println(fi20, v6)
                            }
                        } else {
                            v6 = 29
                            v24 := GenS2{F0: 0, F1: 157, F2: 42}
                            _ = v24
                        }
                    }
                    v25 := GenS3{F0: 4294967295, F1: 100, F2: 255}
                    _ = fi20
                    _ = v25
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi19
            }
        }
        var v26 int8 = 0
        _ = v26
    }
    var v27 int16 = -792
    v2 = 470
    var v28 bool = (442 <= 394)
    var v29 float64 = ((1.5 + ((-2.25 + (1.5 - 1e10)) + float64(0.0))) + 1e10)
    fmt.Println(v28, v1)
    var v30 float32 = ((0.0 - 1e10) * (-2.25 * (((1.5 + (1e10 * (0.1 * 1.5))) / (1e10 / (((1.5 - (3.14159 / 1.0)) - (1.5 / ((0.0 - 3.14159) / 1.0))) * (float32((v29 + (1e10 - 3.14159))) / float32(v29))))) * (1e10 - 0.0))))
    if v28 {
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 uint = 2
            _ = fi31
            _ = v32
        }
    } else {
        if true {
            v6 = v6
        }
        v6 -= (v6 & v6)
    }
    v30 = v30
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
