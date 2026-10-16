package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }
type GenS2 struct { F0 int; F1 int16; F2 uint8 }
func genFn0(p0 string, p1 uint, p2 int8) bool {
    if true {
        v6 := []int{}
        if (true || (p1 >= 2)) {
            for fi7 := 0; fi7 < 6; fi7++ {
                var v8 uint32 = 2147483648
                _ = fi7
                _ = v8
            }
            var v9 int32 = 518
            _ = v9
        } else {
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
        }
        _ = v6
    } else {
        switch 0 {
        case 32767:
            p0 = (p0 + "a")
        case 412:
            p1 = (p1 % ((1) | 1))
        case 643:
            var v10 int = 32768
            _ = v10
        }
        fmt.Println(p2, p0, p2)
    }
    switch (p0 + p0) {
    case "hello":
        v11 := []int{-704, 32768, 0, -778}
        _ = v11
    }
    var v12 float64 = ((1e10 + 0.1) * (0.0 / 0.1))
    var v13 int = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return ((false || false) || (483 > 88))
}

func main() {
    var v1 float32 = -2.25
    var v2 uint16 = 1
    v3 := []string{}
    var v4 uint = 690
    var v5 uint8 = 2
    if (!((!(!(((1 != -9223372036854775808) && (false && (!true))) || false))) || true)) {
        var v14 int8 = 51
        fmt.Println(v1, v14)
        _ = v14
    }
    _ = len(v3)
    var v15 uint = v4
    fmt.Println(v2)
    switch 633 {
    case 421:
        fmt.Println(v2, v15)
    case -2147483648:
        v4 = uint(v2)
    case 4294967295:
        v3 = append(v3, "a")
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := GenS1{F0: 564, F1: -133, F2: v5}
            if false {
                if (834 < uint32(v5)) {
                    var v19 bool = (false && (fi16 <= int(v2)))
                    v1 -= v1
                    fmt.Println(v5)
                    v1 = 3.14159
                    v20 := GenS1{F0: -826, F1: -1, F2: 0}
                    _ = v19
                    _ = v20
                } else {
                    if (v1 < float32(-2.25)) {
                        var v21 int16 = 1000
                        fmt.Println(v1)
                        var v22 int8 = 0
                        _ = v21
                        _ = v22
                    } else {
                        var v23 float32 = float32(1.5)
                        var v24 bool = true
                        _ = v23
                        _ = v24
                    }
                    if true {
                        v15 += (^(v15))
                        fmt.Println(v15, v5, fi16)
                    }
                }
                var v25 uint8 = v5
                _ = v25
            }
            fmt.Println(v5, fi17, fi17)
            _ = fi17
            _ = v18
        }
        v1 -= float32(-2.25)
        _ = fi16
    }
    switch 32767 {
    case 744:
        var v26 uint8 = 2
        _ = v26
    }
    var v27 uint8 = (v5 + v5)
    for fi28 := 0; fi28 < 10; fi28++ {
        switch "" {
        case "foo bar":
            if true {
                var v29 string = "hello"
                v15 = v4
                if ((false && (1 == v5)) || true) {
                    _ = len(v3)
                    var v30 float64 = 0.0
                    if true {
                        v4 = (v4 * (v15 + 2))
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                    } else {
                        v31 := GenS2{F0: fi28, F1: 0, F2: 0}
                        switch v29 {
                        case "":
                            var v32 float64 = v30
                            _ = v32
                        case "foo bar":
                            for fi33 := 0; fi33 < 8; fi33++ {
                                var v34 string = "\n\t"
                                _ = fi33
                                _ = v34
                            }
                        default:
                            var v35 uint = (v15 - 42)
                            _ = v35
                        }
                        _ = v31
                    }
                    _ = v30
                }
                _ = v29
            }
        }
        v15 = (v4 & (v15 + (v4 % ((v15) | 1))))
        switch "foo bar" {
        case "\n\t":
            v2 = v2
        case "":
            v36 := GenS1{F0: 627, F1: 42, F2: (v5 + v5)}
            _ = v36
        default:
            fmt.Println(v2, v27)
        }
        _ = fi28
    }
    v15 -= (v15 | v4)
    v37 := GenS1{F0: 0, F1: 0, F2: (v5 * (v27 + (v5 + v27)))}
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
