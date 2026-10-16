package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint32; F2 int64 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool; F1 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    var v3 uint32 = 365
    var v4 uint16 = 821
    v5 := GenS1{F0: 0}
    v4 -= 2
    v4 = (v4 | v4)
    v6 := GenS2{F0: (("hello" + (("foo bar" + ("\n\t" + "a")) + "")) + ""), F1: 953, F2: -9223372036854775808}
    switch "foo bar" {
    case "":
        v3 += v3
    case "foo bar":
        _ = len(v1)
    case "a":
        v3 = 42
    }
    if true {
        v3 -= (v3 * (v3 / ((v3) | 1)))
        if (-1 > 42) {
            fmt.Println(v4, v4)
        } else {
            var v7 int16 = 0
            switch 0 {
            case -2147483648:
                if false {
                    var v8 uint = 0
                    v3 = 2147483647
                    _ = v8
                } else {
                    if ((1 != (-(v7))) && (((("hello" + (((("foo bar" + "") + "") + "a") + "foo bar")) + (((("foo bar" + "\n\t") + "") + "a") + "a")) + "") >= (("hello" + "hello") + ("" + "foo bar")))) {
                        v7 -= v7
                        _ = len(v2)
                    }
                }
            case -596:
                v9 := GenS2{F0: "\n\t", F1: 995, F2: 57}
                _ = v9
            default:
                v10 := GenS1{F0: v3}
                _ = v10
            }
            _ = v7
        }
    } else {
        v4 -= (v4 | v4)
        var v11 float64 = 0.0
        fmt.Println(v11)
        _ = v11
    }
    v1[204] = 65535
    var v12 float64 = (float64((1e10 / 1.5)) + (float64(float32((-2.25 / (((0.1 - (float64(1e10) / 1.0)) - (-2.25 - (3.14159 / 0.1))) - (-2.25 + 0.0))))) - (1.5 * -2.25)))
    fmt.Println(v12, v4)
    fmt.Println(v12)
    if false {
        switch 9223372036854775807 {
        case 618:
            v4 -= v4
        case 42:
            for fi13 := 0; fi13 < 12; fi13++ {
                for fi14 := 0; fi14 < 8; fi14++ {
                    fmt.Println(v4, fi13, v12)
                    v12 -= ((v12 + ((v12 * ((v12 - float64(1.5)) / -2.25)) * v12)) - ((v12 - v12) / (0.1 + (-2.25 * (-2.25 + 0.0)))))
                    _ = fi14
                }
                for fi15 := 0; fi15 < 20; fi15++ {
                    v12 = 3.14159
                    fmt.Println(v12)
                    v16 := v1[v3]
                    _ = fi15
                    _ = v16
                }
                _ = fi13
            }
        default:
            switch 127 {
            case 2147483647:
                v17 := GenS3{F0: 0, F1: v4}
                _ = v17
            case 32768:
                delete(v2, "")
            case 42:
                if false {
                    v18 := GenS4{F0: 255}
                    if false {
                        switch ("foo bar" + "") {
                        case "":
                            for fi19 := 0; fi19 < 14; fi19++ {
                                fmt.Println(v4, v3)
                                var v20 int16 = -285
                                _ = fi19
                                _ = v20
                            }
                        case "a":
                            switch 32767 {
                            case -255:
                                v4 = 42
                            case 598:
                                v21 := GenS5{F0: ((-42 != int8(v3)) && (!(0.0 <= float32(-2.25))))}
                                _ = v21
                            case 929:
                                v12 -= 0.1
                            default:
                                var v22 uint64 = 18446744073709551615
                                _ = v22
                            }
                        case "foo bar":
                            v4 = (v4 << ((126) & 63))
                        default:
                            if false {
                                if false {
                                    v23 := GenS2{F0: "\n\t", F1: (^(v3)), F2: 9223372036854775807}
                                    v24 := GenS6{F0: 749}
                                    v12 -= v12
                                    _ = v23
                                    _ = v24
                                }
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    v12 = 1.5
                                    delete(v1, 2147483648)
                                    if (65535 >= 1) {
                                        var v26 uint8 = 127
                                        _ = v26
                                    } else {
                                        v27 := GenS5{F0: (32767 > 32767)}
                                        v4 -= v4
                                        v3 -= (^(v3))
                                        _ = v27
                                    }
                                    _ = fi25
                                }
                            }
                        }
                        for fi28 := 0; fi28 < 19; fi28++ {
                            var v29 int8 = -48
                            _ = fi28
                            _ = v29
                        }
                        v3 -= (v3 >> ((42) & 63))
                    }
                    var v30 int = -278
                    _ = v18
                    _ = v30
                }
            default:
                v31 := GenS7{F0: (697 < 1), F1: true}
                _ = v31
            }
        }
    } else {
        v12 -= (1.5 * (((v12 - v12) + (v12 - v12)) - v12))
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 float64 = (float64(0.0) / (1.5 - v12))
        v3 += (v3 / (((v3 - (v3 >> ((0) & 63)))) | 1))
        _ = fi32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
