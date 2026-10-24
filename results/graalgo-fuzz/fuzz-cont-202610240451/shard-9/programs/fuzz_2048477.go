package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
func genFn0() int {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 float32 = ((1.5 / 1e10) - 1.5)
        _ = fi5
        _ = v6
    }
    var v7 uint = 216
    _ = v7
    return 128
}
func genFn1(p0 uint16, p1 int) uint16 {
    p0 = 895
    if ((true || (-895 < 42)) || true) {
        var v8 int32 = 42
        p0 -= 742
        _ = v8
    } else {
        v9 := []int{(p1 >> ((4294967295) & 63)), (p1 * (p1 ^ -238)), (p1 - (p1 % ((4294967295) | 1))), p1}
        p1 += genFn0()
        switch p1 {
        case 9223372036854775807:
            if (false && (false && false)) {
                switch p1 {
                case 127:
                    if true {
                        var v10 int = 1
                        _ = v10
                    } else {
                        for fi11 := 0; fi11 < 17; fi11++ {
                            fmt.Println(fi11, p0, fi11)
                            _ = fi11
                        }
                        _ = len(v9)
                        if (false || (p1 >= 32768)) {
                            _ = len(v9)
                            fmt.Println(p1, p0)
                        } else {
                            fmt.Println(p1, p1)
                        }
                    }
                case -960:
                    var v12 uint8 = 127
                    _ = v12
                case -770:
                    for fi13 := 0; fi13 < 9; fi13++ {
                        _ = len(v9)
                        _ = fi13
                    }
                }
                var v14 float64 = 0.0
                for fi15 := 0; fi15 < 16; fi15++ {
                    switch "\n\t" {
                    case "hello":
                        switch "hello" {
                        case "\n\t":
                            _ = len(v9)
                        }
                    }
                    switch "foo bar" {
                    case "":
                        p1 += -525
                    case "hello":
                        var v16 string = "foo bar"
                        _ = v16
                    case "a":
                        fmt.Println(v14, v14)
                    default:
                        v17 := GenS1{F0: -637}
                        _ = v17
                    }
                    v9 = append(v9, (p1 >> ((1) & 63)))
                    _ = fi15
                }
                _ = v14
            } else {
                p1 += (p1 | p1)
                switch "\n\t" {
                case "":
                    fmt.Println(p0, p0)
                default:
                    fmt.Println(p0, p1)
                }
                var v18 int64 = -611
                _ = v18
            }
        default:
            if true {
                if len(v9) > 0 {
                    v9[0] = (p1 * -500)
                }
                v19 := GenS1{F0: 1000}
                var v20 uint16 = (p0 / ((2) | 1))
                _ = v19
                _ = v20
            } else {
                _ = len(v9)
                v21 := GenS1{F0: -1}
                _ = v21
            }
        }
        _ = v9
    }
    var v22 bool = (617 >= 4294967295)
    p1 += p1
    p1 = p1
    _ = p0
    _ = p1
    _ = v22
    return (p0 ^ p0)
}

func main() {
    var v1 int32 = 359
    var v2 int = -490
    var v3 int16 = 507
    var v4 uint64 = 1
    v1 = -1
    switch -1 {
    case 32767:
        v4 += v4
    case 521:
        for fi23 := 0; fi23 < 8; fi23++ {
            if false {
                if ((128 == 1) && (v4 > 224)) {
                    if (0 > 255) {
                        var v24 int64 = 9223372036854775807
                        _ = v24
                    }
                    var v25 int8 = 84
                    v2 = (^(v2))
                    _ = v25
                }
            } else {
                fmt.Println(v1, v3)
            }
            v3 -= v3
            _ = fi23
        }
    }
    var v26 bool = false
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 string = "\n\t"
        _ = fi27
        _ = v28
    }
    fmt.Println(v3)
    if (v3 <= (v3 * v3)) {
        var v29 int = -887
        v30 := GenS2{F0: 1}
        v1 += v1
        _ = v29
        _ = v30
    }
    v31 := []int{genFn0()}
    v3 = v3
    fmt.Println(v3, v1)
    var v32 int64 = 0
    var v33 uint64 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
