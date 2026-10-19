package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 int }
type GenS3 struct { F0 int }
func genFn0(p0 int8, p1 int16) uint {
    var v6 int16 = p1
    v6 -= v6
    for fi7 := 0; fi7 < 5; fi7++ {
        p1 = (v6 - 1)
        v8 := GenS1{F0: 4294967295, F1: 11, F2: 828}
        if (true || (true || (!false))) {
            fmt.Println(p0, p0)
        } else {
            v9 := []int{}
            switch "\n\t" {
            case "\n\t":
                if true {
                    var v10 int32 = -2147483648
                    _ = v10
                } else {
                    fmt.Println(v6, p0)
                    var v11 int32 = -1
                    _ = len(v9)
                    _ = v11
                }
            }
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 4294967295
}
func genFn1() int {
    var v12 int8 = 3
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS1{F0: 1, F1: -2147483648, F2: 0}
        v15 := []int{fi13, 607}
        v16 := GenS2{F0: ((!(true || true)) || (!false)), F1: 932, F2: fi13}
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: (!true), F1: 1, F2: -9223372036854775808}
    if false {
        for fi18 := 0; fi18 < 19; fi18++ {
            fmt.Println(v12, v12)
            v12 = v12
            v12 -= v12
            _ = fi18
        }
    }
    _ = v12
    _ = v17
    return 1
}

func main() {
    var v1 float64 = (1e10 * 1.5)
    v2 := make(map[uint32]int)
    var v3 uint = 4
    v4 := []int{32768, -9223372036854775808, 128}
    v5 := make(map[string]uint32)
    v1 = ((v1 - v1) / ((v1 + -2.25) + (v1 + 0.1)))
    if (2 >= 941) {
        var v19 uint64 = 274
        _ = v19
    } else {
        if true {
            for fi20 := 0; fi20 < 3; fi20++ {
                v2[0] = (fi20 % (((fi20 | genFn1())) | 1))
                _ = fi20
            }
        }
        switch ("foo bar" + "hello") {
        case "hello":
            v1 = ((v1 / -2.25) * v1)
        case "a":
            for fi21 := 0; fi21 < 13; fi21++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch (fi21 ^ (fi21 * fi21)) {
                case 9223372036854775807:
                    fmt.Println(v3, fi21, v3)
                }
                var v22 int64 = 9223372036854775807
                _ = fi21
                _ = v22
            }
        case "\n\t":
            switch genFn1() {
            case 1:
                v3 = (v3 * 1)
            case -9223372036854775808:
                for fi23 := 0; fi23 < 10; fi23++ {
                    _ = len(v4)
                    _ = fi23
                }
            case 9223372036854775807:
                v1 += (0.0 * (v1 + float64(0.0)))
            default:
                if false {
                    for fi24 := 0; fi24 < 14; fi24++ {
                        v25 := GenS3{F0: (fi24 + (fi24 | 989))}
                        for fi26 := 0; fi26 < 8; fi26++ {
                            var v27 float64 = 0.0
                            v27 += ((-2.25 - 0.1) - ((1.5 - ((v27 / 0.1) + ((v1 + float64(-2.25)) - v1))) + v1))
                            delete(v2, 2147483648)
                            _ = fi26
                            _ = v27
                        }
                        _ = fi24
                        _ = v25
                    }
                    for fi28 := 0; fi28 < 12; fi28++ {
                        delete(v2, 441)
                        _ = fi28
                    }
                } else {
                    v1 = ((float64(1e10) - (1e10 / 1.5)) + (3.14159 * v1))
                }
            }
        default:
            v3 += (v3 * v3)
        }
    }
    v29 := GenS2{F0: (("hello" + ("a" + ("\n\t" + ("" + "\n\t")))) >= "a"), F1: 255, F2: -578}
    var v30 int64 = 0
    v3 -= (v3 / ((v3) | 1))
    v3 = v3
    v1 += (v1 + v1)
    v31 := GenS3{F0: 2147483647}
    for fi32 := 0; fi32 < 6; fi32++ {
        v3 = (v3 << (((v3 >> ((v3) & 63))) & 63))
        var v33 uint = (v3 & (v3 * v3))
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        _ = len(v2)
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
