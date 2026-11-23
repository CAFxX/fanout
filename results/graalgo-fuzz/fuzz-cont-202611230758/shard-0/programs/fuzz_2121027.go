package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int16) uint {
    var v4 bool = (-1 > 4294967295)
    for fi5 := 0; fi5 < 15; fi5++ {
        if v4 {
            v6 := GenS1{F0: 639, F1: -1}
            _ = v6
        } else {
            if ((((0.1 * -2.25) - 0.1) / 0.1) < (3.14159 * ((1.5 + 1.5) + 0.1))) {
                v7 := []int{32768, (fi5 % ((-923) | 1)), -777}
                _ = v7
            }
        }
        var v8 bool = ("a" < (("" + "\n\t") + "a"))
        _ = fi5
        _ = v8
    }
    v9 := GenS1{F0: p0, F1: 2147483647}
    _ = p0
    _ = v4
    _ = v9
    return 0
}

func main() {
    v1 := []string{}
    var v2 uint8 = 36
    var v3 int64 = -9223372036854775808
    v3 = (v3 | (v3 >> ((484) & 63)))
    switch 410 {
    case 128:
        var v10 string = "a"
        _ = v10
    case 32767:
        var v11 bool = true
        _ = v11
    }
    v2 += (v2 % ((v2) | 1))
    for fi12 := 0; fi12 < 5; fi12++ {
        switch "\n\t" {
        case "a":
            v3 -= (v3 - (-(v3)))
        default:
            v2 -= (v2 * (v2 << ((uint(v2)) & 63)))
        }
        var v13 float64 = 3.14159
        _ = fi12
        _ = v13
    }
    var v14 int64 = (v3 / ((9223372036854775807) | 1))
    var v15 uint8 = v2
    _ = len(v1)
    fmt.Println(v15)
    var v16 int8 = 1
    var v17 bool = (!false)
    var v18 int32 = -969
    v17 = (65535 > 0)
    v18 -= v18
    v14 -= v3
    if (false && v17) {
        var v19 int32 = v18
        v16 -= v16
        if (-1 <= -32768) {
            fmt.Println(v16, v17)
            for fi20 := 0; fi20 < 19; fi20++ {
                if (0 > 200) {
                    var v21 bool = v17
                    var v22 float64 = -2.25
                    if ((("hello" + "hello") + "\n\t") == (("hello" + "hello") + "a")) {
                        v23 := GenS1{F0: -79, F1: (v18 | v18)}
                        for fi24 := 0; fi24 < 13; fi24++ {
                            v19 += 42
                            switch (((("\n\t" + (("\n\t" + "\n\t") + "\n\t")) + "a") + "") + ("\n\t" + "")) {
                            case "hello":
                                v2 += (v2 & 2)
                            default:
                                v1 = append(v1, "foo bar")
                            }
                            _ = fi24
                        }
                        _ = v23
                    }
                    _ = v21
                    _ = v22
                } else {
                    if v17 {
                        v25 := GenS1{F0: 0, F1: (v18 & v18)}
                        for fi26 := 0; fi26 < 5; fi26++ {
                            v16 = -114
                            var v27 uint64 = 45
                            v28 := GenS2{F0: v3}
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        _ = v25
                    }
                    for fi29 := 0; fi29 < 8; fi29++ {
                        v17 = (!((v14 + 42) != 0))
                        if (v16 >= v16) {
                            var v30 int8 = 1
                            var v31 uint = 1
                            _ = v30
                            _ = v31
                        }
                        _ = fi29
                    }
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v16, v17)
                    case "":
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                    default:
                        var v32 uint = 436
                        _ = v32
                    }
                }
                _ = fi20
            }
            v15 += (v15 << ((2) & 63))
        } else {
            switch 32767 {
            case 1:
                if false {
                    v16 = v16
                    switch 0 {
                    case 2147483647:
                        for fi33 := 0; fi33 < 19; fi33++ {
                            for fi34 := 0; fi34 < 17; fi34++ {
                                v2 = v15
                                fmt.Println(fi33, v19, v17)
                                fmt.Println(v19, v19)
                                _ = fi34
                            }
                            _ = len(v1)
                            var v35 uint16 = 682
                            _ = fi33
                            _ = v35
                        }
                    case -934:
                        fmt.Println(v16, v14)
                    default:
                        fmt.Println(v3, v14, v2)
                    }
                    var v36 int32 = v18
                    _ = v36
                } else {
                    var v37 int = 42
                    switch v37 {
                    case 9223372036854775807:
                        fmt.Println(v19, v14)
                    }
                    fmt.Println(v17, v18)
                    _ = v37
                }
            default:
                v38 := GenS1{F0: -1, F1: v19}
                _ = v38
            }
        }
        _ = v19
    }
    if ((0.1 < -2.25) || v17) {
        for fi39 := 0; fi39 < 17; fi39++ {
            v1 = append(v1, "")
            _ = fi39
        }
        v2 += 88
    } else {
        v15 -= (^(v15))
        if len(v1) > 0 {
            v1[0] = "hello"
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
