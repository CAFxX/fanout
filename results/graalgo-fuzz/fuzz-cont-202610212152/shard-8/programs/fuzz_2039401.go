package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint16; F1 bool; F2 string }
type GenS3 struct { F0 bool; F1 uint32; F2 int8 }
func genFn0(p0 string, p1 uint16, p2 uint64) uint16 {
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            v7 := GenS1{F0: -1, F1: 0}
            var v8 uint16 = (p1 % (((p1 % ((489) | 1))) | 1))
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: -1, F1: 1}
        _ = fi5
        _ = v9
    }
    v10 := []int{127, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}

func main() {
    var v1 uint32 = 810
    v2 := make(map[int]int)
    var v3 int = -2147483648
    var v4 float32 = 1.5
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: 42, F1: 1}
        for fi13 := 0; fi13 < 14; fi13++ {
            switch v3 {
            case 1:
                switch "\n\t" {
                case "foo bar":
                    var v14 string = ""
                    _ = v14
                default:
                    fmt.Println(fi11, v1, v1)
                }
            case 9223372036854775807:
                fmt.Println(v4, fi13)
            case 739:
                v3 -= -519
            default:
                v15 := []int{v3}
                _ = v15
            }
            var v16 int8 = 109
            var v17 float64 = (3.14159 + ((1.5 + 1e10) / -2.25))
            v18 := GenS2{F0: 310, F1: false, F2: "a"}
            _ = fi13
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi11
        _ = v12
    }
    v4 -= (v4 + ((v4 / v4) * (v4 / v4)))
    switch "a" {
    case "a":
        if (true && (2147483647 > -2147483648)) {
            for fi19 := 0; fi19 < 7; fi19++ {
                v2[fi19] = (-(fi19))
                _ = fi19
            }
        }
    case "\n\t":
        delete(v2, -140)
    case "":
        v1 += v1
    default:
        v3 += v3
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        fmt.Println(fi20)
        var v21 uint64 = 0
        _ = fi20
        _ = v21
    }
    v22 := []int{v3, (v3 + -796), (v3 ^ v3), (v3 & (-(v3)))}
    var v23 int = (v3 << ((1) & 63))
    if (!(2147483648 != v1)) {
        var v24 float32 = (0.0 - (v4 * (-2.25 - float32((float64((v4 - (1.5 * (v4 + 0.0)))) / 1.5)))))
        var v25 uint = 255
        if ("" > "foo bar") {
            v24 -= float32(((0.1 + 1e10) + 3.14159))
            for fi26 := 0; fi26 < 17; fi26++ {
                if ((-1 <= 42) && false) {
                    var v27 uint16 = 623
                    switch "" {
                    case "hello":
                        if (((-2.25 / float32((-2.25 + (1.5 / float64(((v4 - v4) - (v24 + v24))))))) / (v4 - float32(float64(float32(0.1))))) >= float32(0.1)) {
                            if len(v22) > 0 {
                                v22[0] = (v3 - (^(fi26)))
                            }
                            for fi28 := 0; fi28 < 6; fi28++ {
                                delete(v2, (v3 * fi26))
                                var v29 int16 = 1000
                                fmt.Println(v27)
                                _ = fi28
                                _ = v29
                            }
                            var v30 uint16 = 2
                            _ = v30
                        }
                    }
                    for fi31 := 0; fi31 < 9; fi31++ {
                        v23 -= 747
                        switch 32768 {
                        case -265:
                            v32 := GenS2{F0: 1, F1: ((v1 >> (((v25 >> ((v25) & 63))) & 63)) >= 2147483648), F2: ""}
                            _ = v32
                        case 2147483647:
                            switch (fi31 << ((v25) & 63)) {
                            case -310:
                                var v33 float64 = ((((float64(v4) / 1.5) + (3.14159 * float64(v4))) + 3.14159) - (0.1 * 1.5))
                                _ = v33
                            }
                        default:
                            _ = len(v2)
                        }
                        _ = fi31
                    }
                    _ = v27
                } else {
                    v2[(fi26 >> (((v25 << ((v25) & 63))) & 63))] = fi26
                    var v34 uint8 = uint8(v25)
                    _ = v34
                }
                v2[9223372036854775807] = v23
                v25 = (v25 % ((v25) | 1))
                _ = fi26
            }
            v35 := GenS2{F0: 255, F1: true, F2: ("a" + "hello")}
            _ = v35
        } else {
            v24 += 3.14159
            v36 := GenS3{F0: ((true && (78 <= -1)) && ((((v4 * v4) * v24) != (float32(0.1) + 3.14159)) && true)), F1: 2, F2: 42}
            _ = v36
        }
        _ = v24
        _ = v25
    } else {
        var v37 int = -2147483648
        v23 += (v3 + (v37 & v3))
        var v38 int32 = -2147483648
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 19; fi39++ {
        fmt.Println(fi39, v4)
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
