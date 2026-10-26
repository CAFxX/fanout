package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint8; F2 int16 }
type GenS3 struct { F0 int32; F1 bool }
func genFn0(p0 int8, p1 int8, p2 uint32) string {
    p0 += p1
    switch 32767 {
    case -2147483648:
        fmt.Println(p0)
    default:
        v6 := GenS1{F0: 951, F1: 0}
        _ = v6
    }
    var v7 int16 = 1000
    switch ("\n\t" + "") {
    case "foo bar":
        switch "a" {
        case "\n\t":
            if (true && false) {
                v8 := GenS2{F0: v7, F1: 42, F2: 42}
                if (0 >= 4294967295) {
                    var v9 int8 = p1
                    v9 -= (v9 ^ p0)
                    _ = v9
                }
                _ = v8
            } else {
                var v10 int8 = (-(p0))
                _ = v10
            }
        case "":
            var v11 float32 = -2.25
            _ = v11
        case "a":
            if (true && (1e10 <= ((1.5 / (((float64(((1.5 + 1.5) - 3.14159)) - (0.0 / 1.5)) - float64(3.14159)) - (0.1 * 0.0))) + 1e10))) {
                fmt.Println(p2, p1)
            } else {
                var v12 float32 = 1.5
                if (255 < 1) {
                    var v13 int32 = int32(v7)
                    _ = v13
                }
                switch "\n\t" {
                case "":
                    for fi14 := 0; fi14 < 11; fi14++ {
                        switch ("hello" + ("\n\t" + "a")) {
                        case "hello":
                            p0 -= p1
                        case "":
                            var v15 string = "\n\t"
                            _ = v15
                        case "\n\t":
                            p1 -= (p0 - p1)
                        default:
                            var v16 float32 = v12
                            _ = v16
                        }
                        _ = fi14
                    }
                case "a":
                    p1 += (p1 / (((^(p1))) | 1))
                case "hello":
                    switch -50 {
                    case 32767:
                        fmt.Println(v7)
                    case -754:
                        fmt.Println(p1, p0)
                    }
                default:
                    p0 = p1
                }
                _ = v12
            }
        }
    case "\n\t":
        if false {
            switch 42 {
            case 919:
                v7 -= int16(p2)
            }
        } else {
            v17 := []int{192, 850, int(p1)}
            if (!false) {
                switch 902 {
                case 2147483647:
                    p1 -= (^(p1))
                }
                if false {
                    var v18 int = int(p1)
                    v19 := GenS3{F0: 0, F1: false}
                    var v20 int64 = 0
                    _ = v18
                    _ = v19
                    _ = v20
                }
            } else {
                p1 = (p1 / ((p1) | 1))
                _ = len(v17)
            }
            if (false && true) {
                if len(v17) > 0 {
                    v17[0] = 42
                }
                if len(v17) > 0 {
                    v17[0] = 2147483647
                }
            } else {
                fmt.Println(p0, p0, v7)
                fmt.Println(p1)
            }
            _ = v17
        }
    }
    var v21 string = ("foo bar" + (("foo bar" + "hello") + "foo bar"))
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v7
    return v21
}
func genFn1(p0 uint, p1 int) int8 {
    var v22 uint = 63
    v23 := GenS2{F0: 0, F1: 127, F2: 32767}
    p1 += p1
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 99
}

func main() {
    var v1 int8 = -96
    v2 := make(map[uint32]int)
    var v3 uint16 = 32767
    v4 := []string{"foo bar", "", "foo bar"}
    v5 := []uint32{774}
    var v24 uint8 = 127
    v25 := GenS3{F0: -1, F1: ((v3 | (^(v3))) > 2)}
    if (v1 >= 0) {
        v1 = (v1 % ((67) | 1))
        var v26 int64 = 9223372036854775807
        for fi27 := 0; fi27 < 9; fi27++ {
            v26 = 518
            _ = fi27
        }
        _ = v26
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        var v29 uint32 = 65535
        _ = fi28
        _ = v29
    }
    v3 = (^(v3))
    delete(v2, 915)
    v30 := GenS2{F0: 0, F1: (v24 % (((^(v24))) | 1)), F2: 1}
    v24 -= (v24 | v24)
    v3 -= (v3 + (v3 ^ (v3 ^ v3)))
    var v31 bool = (true || true)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
