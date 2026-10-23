package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 bool }
type GenS2 struct { F0 uint64; F1 string; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint16, p1 int, p2 string) uint64 {
    p2 = ("a" + "")
    fmt.Println(p0, p0)
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 uint64 = 2
        _ = fi4
        _ = v5
    }
    p0 = 375
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 uint16, p1 int) int32 {
    var v6 uint = 18446744073709551615
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(p0, p0, p1)
        if (1 <= 42) {
            if (!(197 <= -1)) {
                switch 539 {
                case 345:
                    fmt.Println(fi7, p1)
                case -9223372036854775808:
                    v8 := GenS1{F0: -965, F1: v6, F2: true}
                    _ = v8
                }
            } else {
                p1 = fi7
            }
            fmt.Println(p0, v6)
        }
        _ = fi7
    }
    if true {
        var v9 float64 = 0.0
        v10 := []int{(^(p1))}
        _ = v10
        _ = v9
    }
    var v11 bool = true
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 0
    var v3 uint = 0
    var v12 uint8 = 168
    var v13 uint8 = 0
    var v14 uint8 = 216
    v15 := GenS2{F0: 18446744073709551615, F1: ("hello" + ("\n\t" + ("a" + "\n\t"))), F2: v2}
    v13 += (^(v12))
    var v16 float64 = (((1e10 - (1.5 * -2.25)) * ((float64(1.5) - (-2.25 / 1.0)) + float64(1e10))) - (float64(1e10) + 3.14159))
    v14 += (v14 << ((v3) & 63))
    v16 = (float64((float32(v16) + 0.1)) - v16)
    var v17 int16 = -1
    for fi18 := 0; fi18 < 7; fi18++ {
        v2 -= uint16(v3)
        for fi19 := 0; fi19 < 15; fi19++ {
            if ("foo bar" != "a") {
                fmt.Println(fi19)
                v2 += v2
            }
            if (true && (false && (1 < -69))) {
                v14 = 203
                v20 := GenS2{F0: 18446744073709551615, F1: ("" + "a"), F2: v2}
                _ = v20
            } else {
                var v21 uint16 = v2
                v22 := GenS1{F0: -1, F1: 255, F2: ((v3 - (v3 ^ 0)) <= (v3 << ((v3) & 63)))}
                fmt.Println(v17)
                _ = v21
                _ = v22
            }
            if true {
                v23 := GenS3{F0: 42}
                switch "a" {
                case "a":
                    v16 -= ((1e10 + v16) * 0.1)
                case "hello":
                    if ((v14 | 2) > (v14 * v13)) {
                        switch "\n\t" {
                        case "a":
                            var v24 int32 = 2147483647
                            _ = v24
                        }
                        v16 += v16
                    } else {
                        if false {
                            var v25 int8 = -1
                            v26 := GenS4{F0: (v17 << ((v3) & 63)), F1: v2}
                            _ = v25
                            _ = v26
                        } else {
                            var v27 int = fi18
                            _ = v27
                        }
                    }
                case "foo bar":
                    v12 = v13
                default:
                    var v28 int8 = -79
                    _ = v28
                }
                _ = v23
            }
            v29 := []int{-2147483648, (-(fi19)), 1}
            _ = fi19
            _ = v29
        }
        _ = fi18
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
