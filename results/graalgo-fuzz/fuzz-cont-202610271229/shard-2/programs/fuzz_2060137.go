package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint }
type GenS2 struct { F0 string; F1 int32; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 int8; F2 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 uint }
func genFn0(p0 bool, p1 uint16) uint8 {
    switch 9223372036854775807 {
    case 0:
        p0 = (((^(p1)) <= (p1 - (p1 | p1))) && p0)
    case -631:
        var v4 int = 127
        _ = v4
    case -708:
        p1 -= p1
    default:
        fmt.Println(p0, p1)
    }
    p1 += (^(p1))
    v5 := GenS1{F0: 1, F1: ("hello" + "\n\t"), F2: 869}
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1(p0 uint64) bool {
    if true {
        fmt.Println(p0)
    }
    fmt.Println(p0)
    fmt.Println(p0)
    if true {
        if (1e10 != 0.1) {
            v6 := GenS1{F0: (p0 % ((p0) | 1)), F1: (("foo bar" + "a") + ""), F2: 65535}
            p0 += p0
            fmt.Println(p0)
            _ = v6
        } else {
            p0 = p0
        }
        p0 -= p0
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int8 = 0
            var v9 uint = 859
            p0 += (p0 % ((p0) | 1))
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        fmt.Println(p0)
        v10 := GenS2{F0: "a", F1: -834, F2: 127}
        _ = v10
    }
    _ = p0
    return true
}

func main() {
    var v1 uint = 879
    var v2 uint16 = 0
    var v3 float64 = 0.1
    v2 = v2
    v3 = (0.0 - ((1e10 / (-2.25 - ((3.14159 + v3) + (v3 - ((((1.5 * v3) * v3) + v3) / 0.1))))) + v3))
    var v11 int8 = -128
    v1 = 39
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            fmt.Println(v11)
            if false {
                v1 -= (v1 | v1)
                switch "foo bar" {
                case "hello":
                    v3 -= ((v3 * v3) - (-2.25 + v3))
                case "a":
                    v1 -= 203
                case "\n\t":
                    fmt.Println(v2, fi13, fi12)
                }
                var v14 uint = (^(v1))
                switch ("hello" + "") {
                case "":
                    v3 += 3.14159
                case "hello":
                    v15 := []int{71, fi12}
                    _ = v15
                case "a":
                    v2 = 558
                default:
                    v2 -= v2
                }
                _ = v14
            } else {
                var v16 uint64 = 0
                _ = v16
            }
            v1 = (^(v1))
            _ = fi13
        }
        v1 -= v1
        v3 += 1.5
        _ = fi12
    }
    v3 = (v3 + -2.25)
    v1 = (^(v1))
    if false {
        v11 -= v11
        if (false && false) {
            v3 = v3
        } else {
            for fi17 := 0; fi17 < 2; fi17++ {
                v3 = v3
                fmt.Println(v3)
                switch "" {
                case "a":
                    var v18 uint8 = 130
                    _ = v18
                case "hello":
                    v2 = (v2 >> ((2) & 63))
                case "foo bar":
                    v3 = 0.0
                default:
                    if true {
                        switch "" {
                        case "a":
                            fmt.Println(v1, v2)
                        }
                        var v19 int8 = v11
                        v20 := GenS3{F0: (65535 < 2147483648)}
                        _ = v19
                        _ = v20
                    } else {
                        fmt.Println(v2, v11, fi17)
                    }
                }
                _ = fi17
            }
            v21 := []int{}
            switch 204 {
            case -208:
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 uint64 = 2
                    if true {
                        switch "\n\t" {
                        case "\n\t":
                            v21 = append(v21, (fi22 - -9223372036854775808))
                        case "a":
                            v24 := GenS4{F0: 0, F1: -1, F2: -32768}
                            _ = v24
                        case "hello":
                            v11 -= (v11 + (v11 & v11))
                        default:
                            v1 = (v1 ^ 255)
                        }
                        v25 := GenS5{F0: -411}
                        v26 := GenS6{F0: 2147483647, F1: v1}
                        _ = v25
                        _ = v26
                    } else {
                        fmt.Println(v2, v23)
                    }
                    _ = fi22
                    _ = v23
                }
            }
            _ = v21
        }
    }
    var v27 int32 = 641
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
}
