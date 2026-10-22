package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int; F2 int8 }
func genFn0(p0 uint8, p1 uint8) bool {
    var v4 int16 = 719
    fmt.Println(v4, v4, p1)
    v4 -= (v4 & -467)
    v4 = int16(p0)
    v5 := []int{2147483647, -2147483648}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return false
}

func main() {
    var v1 uint = 761
    var v2 int8 = -31
    var v3 string = "foo bar"
    var v6 uint = (v1 + v1)
    v6 = (v1 - (^(v6)))
    if true {
        v6 -= (v1 >> ((477) & 63))
        v3 = (("a" + v3) + "hello")
    } else {
        switch -368 {
        case -312:
            for fi7 := 0; fi7 < 17; fi7++ {
                v1 = v1
                _ = fi7
            }
        default:
            v8 := GenS1{F0: (0 != 777)}
            _ = v8
        }
    }
    fmt.Println(v3, v6)
    switch 42 {
    case -9223372036854775808:
        for fi9 := 0; fi9 < 9; fi9++ {
            v3 = "hello"
            fmt.Println(v6)
            _ = fi9
        }
    case -2147483648:
        switch v3 {
        case "hello":
            v10 := []int{}
            _ = v10
        }
    case 32768:
        var v11 bool = false
        _ = v11
    default:
        v6 -= (^(v6))
    }
    var v12 string = (((v3 + ((v3 + v3) + "hello")) + v3) + ("a" + ((v3 + v3) + "\n\t")))
    var v13 uint64 = 616
    v14 := []int{128, 940, 32768, -742}
    var v15 float64 = (1.5 / (3.14159 * float64(-2.25)))
    if true {
        v2 += v2
    }
    if false {
        if (-1 > 0) {
            v2 += 100
        } else {
            if true {
                v16 := GenS1{F0: (42 >= 0)}
                v1 = v6
                switch v12 {
                case "a":
                    var v17 uint8 = 255
                    _ = v17
                case "hello":
                    v14 = append(v14, 128)
                case "foo bar":
                    if (822 >= 604) {
                        fmt.Println(v13, v15)
                        for fi18 := 0; fi18 < 7; fi18++ {
                            var v19 bool = true
                            _ = fi18
                            _ = v19
                        }
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            if true {
                                if len(v14) > 0 {
                                    v14[0] = 32767
                                }
                                var v20 int8 = -1
                                _ = v20
                            }
                        case "a":
                            v15 -= v15
                        case "\n\t":
                            if ((-9223372036854775808 == -9223372036854775808) && (0.1 > (v15 - 3.14159))) {
                                switch v3 {
                                case "hello":
                                    v1 = (v6 & (v6 << ((2) & 63)))
                                default:
                                    for fi21 := 0; fi21 < 6; fi21++ {
                                        var v22 int8 = v2
                                        fmt.Println(v1, v12, v2)
                                        fmt.Println(v13, v6, v12)
                                        _ = fi21
                                        _ = v22
                                    }
                                }
                                v12 = ("" + "foo bar")
                                _ = len(v14)
                            }
                        }
                    }
                }
                _ = v16
            } else {
                v23 := GenS1{F0: false}
                var v24 uint = v1
                v15 = 0.1
                _ = v23
                _ = v24
            }
            v25 := GenS1{F0: ((!false) && false)}
            if len(v14) > 0 {
                v14[0] = 578
            }
            _ = v25
        }
        var v26 float64 = 1e10
        _ = v26
    } else {
        v2 -= v2
        for fi27 := 0; fi27 < 7; fi27++ {
            v28 := GenS2{F0: 182, F1: fi27, F2: 100}
            v1 = 534
            v29 := GenS2{F0: 2, F1: fi27, F2: v2}
            _ = fi27
            _ = v28
            _ = v29
        }
        v6 = v6
    }
    var v30 int = 4294967295
    var v31 int = 127
    var v32 uint32 = 42
    if true {
        v33 := GenS1{F0: false}
        _ = v33
    } else {
        v1 += 229
    }
    _ = len(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v6
}
