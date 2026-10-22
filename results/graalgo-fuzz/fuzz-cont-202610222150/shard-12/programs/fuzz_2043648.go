package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 string }
func genFn0(p0 int) int8 {
    p0 = p0
    switch "" {
    case "hello":
        v5 := []int{}
        _ = v5
    }
    if ((!true) || (false || (2 < 2))) {
        switch "hello" {
        case "hello":
            for fi6 := 0; fi6 < 20; fi6++ {
                var v7 float32 = 1.5
                _ = fi6
                _ = v7
            }
        case "a":
            switch "hello" {
            case "foo bar":
                var v8 int64 = -575
                _ = v8
            case "hello":
                p0 += p0
            case "\n\t":
                var v9 float32 = (0.0 * -2.25)
                _ = v9
            default:
                if false {
                    v10 := []int{(^(p0))}
                    var v11 uint64 = 18446744073709551615
                    switch 32768 {
                    case 0:
                        p0 += -1
                    case 761:
                        if len(v10) > 0 {
                            v10[0] = p0
                        }
                    case -2147483648:
                        fmt.Println(p0, p0)
                    }
                    _ = v10
                    _ = v11
                }
            }
        case "":
            fmt.Println(p0)
        }
        p0 = p0
    }
    _ = p0
    return -1
}
func genFn1(p0 uint8, p1 uint8) uint32 {
    v12 := GenS1{F0: -1}
    fmt.Println(p1)
    var v13 int32 = 42
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 int32 = v13
            p1 += (p0 ^ (p1 & 1))
            v17 := []int{fi14, (fi15 + fi14), (^(fi14))}
            _ = fi15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 4; fi18++ {
            switch 32767 {
            case 1:
                v13 -= -192
                v13 = (^(v13))
            case 2147483647:
                switch "a" {
                case "":
                    var v19 uint8 = (p1 % ((p0) | 1))
                    _ = v19
                case "\n\t":
                    p1 -= (p1 >> ((42) & 63))
                case "foo bar":
                    var v20 float64 = 3.14159
                    _ = v20
                default:
                    v21 := []int{}
                    _ = v21
                }
            case -9223372036854775808:
                switch (^(fi14)) {
                case 850:
                    var v22 float32 = 3.14159
                    _ = v22
                case -669:
                    if true {
                        p0 = p1
                        p0 += p0
                    }
                case 32768:
                    v13 = (v13 | 0)
                }
                v23 := GenS2{F0: (!true), F1: "a"}
                var v24 uint8 = p1
                _ = v23
                _ = v24
            }
            _ = fi18
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 2147483648
}

func main() {
    v1 := []int{127, 32768, 32767}
    var v2 string = "hello"
    var v3 float64 = ((-2.25 - (3.14159 * (0.1 / 3.14159))) - 1e10)
    v4 := []uint32{656, 258, 1}
    var v25 float64 = float64((0.1 - 1.5))
    v3 = v25
    if len(v1) > 0 {
        v1[0] = 127
    }
    v26 := GenS1{F0: 0}
    v25 = v25
    if (0 >= 1) {
        var v27 uint = 2
        if len(v4) > 0 {
            v4[0] = 0
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v27
    } else {
        if (false || false) {
            for fi28 := 0; fi28 < 4; fi28++ {
                if false {
                    var v29 int32 = -1
                    _ = v29
                }
                switch (fi28 * fi28) {
                case -2147483648:
                    if (!(true || (true || ((-32768 < 42) || true)))) {
                        for fi30 := 0; fi30 < 12; fi30++ {
                            if len(v4) > 0 {
                                v4[0] = 33
                            }
                            _ = fi30
                        }
                    }
                case 770:
                    var v31 int16 = -988
                    _ = v31
                case -1:
                    var v32 bool = (!(-974 == 32767))
                    _ = v32
                default:
                    var v33 float32 = 1.5
                    _ = v33
                }
                v25 += 1e10
                _ = fi28
            }
            var v34 int16 = -32768
            _ = v34
        }
        fmt.Println(v3, v25)
        var v35 float32 = (0.1 / (((((0.0 / 1.0) / (0.1 - 1.5)) + 0.0) / float32((1.5 * v25))) - -2.25))
        _ = v35
    }
    v2 = ((v2 + "a") + v2)
    v25 = v25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
