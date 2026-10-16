package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int16, p1 string) int16 {
    v5 := GenS1{F0: 9223372036854775807}
    var v6 uint32 = 42
    v7 := []int{-1, int(p0)}
    fmt.Println(v6, p0, v6)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (-(p0))
}

func main() {
    v1 := []int{9223372036854775807, -9223372036854775808, 148}
    var v2 int8 = -128
    v3 := []string{("foo bar" + (("a" + "a") + ((("a" + "") + "a") + "a"))), ""}
    var v4 int64 = 9223372036854775807
    for fi8 := 0; fi8 < 11; fi8++ {
        v4 += 9223372036854775807
        if (((-2.25 * -2.25) != (float64(1.5) / 1.5)) && true) {
            var v9 float32 = float32(((1.5 / -2.25) / 3.14159))
            fmt.Println(v4, v2, v4)
            _ = v9
        } else {
            var v10 float64 = (0.0 - -2.25)
            _ = v10
        }
        _ = fi8
    }
    if (v2 == 1) {
        v2 -= (v2 ^ -53)
        for fi11 := 0; fi11 < 7; fi11++ {
            if true {
                for fi12 := 0; fi12 < 11; fi12++ {
                    _ = len(v3)
                    switch "" {
                    case "\n\t":
                        v2 = (v2 >> ((18446744073709551615) & 63))
                    }
                    v13 := GenS1{F0: (-(v4))}
                    _ = fi12
                    _ = v13
                }
                var v14 int16 = 42
                _ = v14
            } else {
                _ = len(v3)
            }
            var v15 bool = true
            _ = fi11
            _ = v15
        }
        switch 1 {
        case 0:
            if (!(!(!true))) {
                switch 1 {
                case -9223372036854775808:
                    v16 := GenS1{F0: 1}
                    _ = v16
                }
                fmt.Println(v2, v2)
                for fi17 := 0; fi17 < 1; fi17++ {
                    fmt.Println(fi17, v4)
                    if (fi17 >= (fi17 & int(v4))) {
                        v18 := GenS1{F0: (v4 ^ v4)}
                        v19 := GenS1{F0: -1}
                        for fi20 := 0; fi20 < 8; fi20++ {
                            v2 = v2
                            fmt.Println(v4, fi20)
                            _ = fi20
                        }
                        _ = v18
                        _ = v19
                    }
                    _ = fi17
                }
            } else {
                for fi21 := 0; fi21 < 2; fi21++ {
                    v3 = append(v3, "foo bar")
                    if (!false) {
                        fmt.Println(v4, v4)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v4 = v4
                    } else {
                        var v22 float64 = 1e10
                        _ = v22
                    }
                    fmt.Println(v4, fi21, fi21)
                    _ = fi21
                }
            }
        case 127:
            var v23 uint = 0
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        fmt.Println(v4)
        var v25 int16 = 42
        v26 := GenS1{F0: (-(v4))}
        _ = fi24
        _ = v25
        _ = v26
    }
    switch -531 {
    case -841:
        var v27 int8 = v2
        _ = v27
    case 9223372036854775807:
        var v28 bool = (32767 != -32768)
        _ = v28
    }
    v2 = (v2 | (v2 * 0))
    var v29 uint8 = 130
    var v30 uint16 = uint16(v29)
    for fi31 := 0; fi31 < 12; fi31++ {
        v4 = v4
        if false {
            for fi32 := 0; fi32 < 15; fi32++ {
                v30 = 0
                v2 += -1
                _ = fi32
            }
            var v33 uint16 = v30
            _ = v33
        }
        if true {
            v3 = append(v3, "")
        } else {
            v34 := GenS1{F0: 42}
            switch "foo bar" {
            case "a":
                var v35 string = (("foo bar" + "a") + ("foo bar" + "\n\t"))
                _ = v35
            case "\n\t":
                v30 = 65535
            case "foo bar":
                var v36 int16 = genFn0(42, ("" + "\n\t"))
                _ = v36
            default:
                if false {
                    v4 = (v4 | (v4 + (v4 | int64(v2))))
                    v37 := GenS2{F0: 255, F1: v29}
                    _ = v37
                }
            }
            _ = v34
        }
        _ = fi31
    }
    v29 -= (v29 - v29)
    v30 += v30
    v30 += v30
    var v38 uint32 = 1
    fmt.Println(v30, v38, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v4
}
