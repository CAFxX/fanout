package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint64; F2 bool }
func genFn0(p0 int32, p1 int32) bool {
    if (9223372036854775807 < 641) {
        v7 := []int{}
        _ = v7
    }
    var v8 int64 = -9223372036854775808
    var v9 uint32 = 42
    for fi10 := 0; fi10 < 12; fi10++ {
        p1 -= -1
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return true
}
func genFn1(p0 string, p1 uint16) uint {
    v11 := GenS1{F0: 844, F1: false, F2: 0}
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 string = p0
        switch (-(fi12)) {
        case 42:
            var v14 int64 = 1
            _ = v14
        default:
            v15 := []int{fi12, (fi12 & fi12)}
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v16 := GenS1{F0: 42, F1: (false || false), F2: 551}
    _ = p0
    _ = p1
    _ = v11
    _ = v16
    return 76
}

func main() {
    var v1 float64 = 1.5
    var v2 int16 = 0
    var v3 uint32 = 865
    var v4 float64 = v1
    var v5 uint8 = 222
    v6 := []uint32{(v3 * v3), v3, (v3 ^ (v3 * 1))}
    v6 = append(v6, (v3 >> ((1) & 63)))
    if (0.1 == (v4 - v4)) {
        switch (("a" + "") + "\n\t") {
        case "foo bar":
            var v17 uint16 = 480
            _ = v17
        default:
            v2 += (v2 * v2)
        }
    }
    v18 := GenS1{F0: 267, F1: false, F2: 2}
    _ = len(v6)
    fmt.Println(v4)
    for fi19 := 0; fi19 < 16; fi19++ {
        v4 -= 0.0
        _ = fi19
    }
    var v20 uint16 = 255
    if (65535 > 559) {
        if false {
            v21 := GenS1{F0: -1, F1: false, F2: v3}
            var v22 float64 = ((v1 / v4) / ((v1 - v4) - ((1e10 / (((1e10 - (v4 + 3.14159)) + v1) + v1)) - (0.1 * v4))))
            _ = v21
            _ = v22
        }
        v4 -= (0.1 * 0.0)
        fmt.Println(v2, v4, v5)
    }
    var v23 uint32 = v3
    switch -1 {
    case -722:
        v2 = v2
    case -1:
        var v24 string = ""
        _ = v24
    case 119:
        for fi25 := 0; fi25 < 1; fi25++ {
            fmt.Println(v2, fi25, v23)
            for fi26 := 0; fi26 < 11; fi26++ {
                var v27 int16 = v2
                var v28 float64 = v4
                if ((-47 < 285) && (!(127 >= (v5 & 117)))) {
                    switch "a" {
                    case "foo bar":
                        if (!(false && ((fi25 / (((fi25 + (-(fi26)))) | 1)) == (fi25 / ((fi26) | 1))))) {
                            v6 = append(v6, (v3 & 572))
                            var v29 float64 = v1
                            if len(v6) > 0 {
                                v6[0] = 0
                            }
                            _ = v29
                        } else {
                            v23 = v23
                            v2 += 517
                            var v30 uint8 = (v5 + v5)
                            _ = v30
                        }
                    case "":
                        var v31 uint16 = v20
                        _ = v31
                    case "hello":
                        v32 := GenS2{F0: -128, F1: 18446744073709551615, F2: (int64(v3) != int64(fi26))}
                        _ = v32
                    }
                } else {
                    v6 = append(v6, (v3 ^ (v3 / (((v23 / (((v3 << ((genFn1(("\n\t" + "foo bar"), v20)) & 63))) | 1))) | 1))))
                    v6 = append(v6, 42)
                }
                _ = fi26
                _ = v27
                _ = v28
            }
            v4 = v4
            _ = fi25
        }
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        for fi34 := 0; fi34 < 14; fi34++ {
            if len(v6) > 0 {
                v6[0] = uint32(v5)
            }
            v35 := GenS1{F0: 1, F1: false, F2: (v23 << ((239) & 63))}
            _ = fi34
            _ = v35
        }
        for fi36 := 0; fi36 < 7; fi36++ {
            v6 = append(v6, (v3 | v23))
            v5 -= (^(v5))
            var v37 string = ("" + ((("foo bar" + "a") + "hello") + "\n\t"))
            _ = fi36
            _ = v37
        }
        _ = fi33
    }
    if (42 == 2) {
        switch 32767 {
        case 538:
            var v38 int = 127
            _ = v38
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi39 := 0; fi39 < 2; fi39++ {
            fmt.Println(v20, fi39)
            _ = fi39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
