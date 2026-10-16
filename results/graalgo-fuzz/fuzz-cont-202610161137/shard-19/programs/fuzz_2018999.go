package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0(p0 int, p1 uint32) uint32 {
    v6 := []int{p0, (p0 ^ (p0 << ((0) & 63)))}
    var v7 uint = 2
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int = (^(fi8))
        _ = fi8
        _ = v9
    }
    fmt.Println(p0)
    switch "" {
    case "\n\t":
        p0 += p0
    case "a":
        _ = len(v6)
    case "":
        v10 := GenS1{F0: -18, F1: 1}
        _ = v10
    default:
        for fi11 := 0; fi11 < 5; fi11++ {
            p1 = 316
            if (v7 == 347) {
                if (false && (42 == 808)) {
                    var v12 bool = false
                    var v13 int = (^(p0))
                    switch "foo bar" {
                    case "\n\t":
                        for fi14 := 0; fi14 < 15; fi14++ {
                            p1 = (p1 >> ((837) & 63))
                            v7 -= (^(v7))
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = fi14
                        }
                    case "foo bar":
                        v15 := GenS1{F0: -1, F1: 2}
                        _ = v15
                    case "a":
                        var v16 uint8 = 137
                        _ = v16
                    }
                    _ = v12
                    _ = v13
                } else {
                    if ((((!false) || true) || ((fi11 | (^(p0))) >= 32768)) || (p0 >= 524)) {
                        var v17 float32 = 0.1
                        var v18 uint = v7
                        _ = v17
                        _ = v18
                    } else {
                        fmt.Println(v7)
                        p1 = (p1 * p1)
                        var v19 float64 = 1.5
                        _ = v19
                    }
                    v7 -= 255
                }
            }
            if false {
                var v20 bool = (!false)
                if (!((-1 == -29) || (98 == 1))) {
                    v7 -= 0
                    var v21 bool = false
                    _ = v21
                }
                _ = v20
            } else {
                v22 := GenS1{F0: 1, F1: 1}
                _ = v22
            }
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 | p1)
}
func genFn1(p0 uint16) uint8 {
    if true {
        fmt.Println(p0)
    }
    fmt.Println(p0)
    var v23 int16 = 1
    var v24 float32 = (-2.25 * 0.0)
    _ = p0
    _ = v23
    _ = v24
    return 32
}

func main() {
    var v1 int32 = -467
    v2 := []string{("" + ""), "", ""}
    var v3 int16 = 696
    v4 := make(map[string]bool)
    var v5 int8 = 84
    var v25 uint64 = 0
    v1 = 42
    v5 = (v5 | v5)
    v3 -= -842
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 bool = (("\n\t" + "\n\t") >= ((("foo bar" + ("\n\t" + "foo bar")) + "\n\t") + ("foo bar" + "\n\t")))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v28 int64 = 9223372036854775807
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 int32 = 1
    var v30 uint = 255
    v30 = (^(v30))
    for fi31 := 0; fi31 < 13; fi31++ {
        fmt.Println(v25, v25, v25)
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
