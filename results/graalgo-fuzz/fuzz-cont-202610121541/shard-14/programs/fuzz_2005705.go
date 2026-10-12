package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0(p0 uint16, p1 uint) int {
    fmt.Println(p1)
    p1 = (p1 ^ p1)
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    return -240
}
func genFn1(p0 string) int16 {
    var v9 string = ("" + "")
    fmt.Println(p0)
    _ = p0
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = ((true && false) && true)
    var v3 uint32 = 557
    var v4 string = "foo bar"
    v5 := []string{v4, ("a" + (("\n\t" + "hello") + v4)), v4, ("\n\t" + v4)}
    var v6 int = 128
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 int32 = 1
        _ = fi10
        _ = v11
    }
    v3 = v3
    v12 := v1[(v6 | genFn0(0, 2))]
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v6)
    var v13 float64 = 1.5
    if v2 {
        delete(v1, v6)
    }
    switch -9223372036854775808 {
    case -555:
        if (!((0 == 2) || (v3 >= 296))) {
            for fi14 := 0; fi14 < 11; fi14++ {
                v2 = (31 > 168)
                _ = fi14
            }
            v3 -= 628
            for fi15 := 0; fi15 < 3; fi15++ {
                if (v2 || v2) {
                    delete(v1, int(v3))
                    if (v2 || false) {
                        switch "foo bar" {
                        case "":
                            for fi16 := 0; fi16 < 11; fi16++ {
                                fmt.Println(v6, v2, fi15)
                                if v2 {
                                    v3 = (v3 >> ((305) & 63))
                                } else {
                                    v12 = "a"
                                    if len(v5) > 0 {
                                        v5[0] = v12
                                    }
                                }
                                v2 = true
                                fmt.Println(fi15, v3, v2)
                                _ = fi16
                            }
                        case "hello":
                            var v17 bool = (v2 && true)
                            _ = v17
                        default:
                            for fi18 := 0; fi18 < 3; fi18++ {
                                var v19 uint16 = 666
                                _ = fi18
                                _ = v19
                            }
                        }
                    }
                }
                v6 -= (fi15 - fi15)
                if false {
                    for fi20 := 0; fi20 < 6; fi20++ {
                        fmt.Println(v6)
                        var v21 uint64 = 1
                        _ = fi20
                        _ = v21
                    }
                    v22 := GenS1{F0: 42, F1: (v12 + "\n\t")}
                    v1[fi15] = v4
                    _ = v22
                }
                _ = fi15
            }
        } else {
            var v23 int8 = 0
            _ = v23
        }
    case 1:
        var v24 int8 = -114
        _ = v24
    default:
        v25 := GenS1{F0: 65535, F1: v12}
        _ = v25
    }
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
