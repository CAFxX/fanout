package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 string; F1 uint16 }
func genFn0(p0 int64, p1 uint8, p2 int) uint64 {
    v4 := []int{}
    var v5 bool = (42 > 18446744073709551615)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 int) int32 {
    var v6 float32 = -2.25
    p0 = (p0 | (p0 >> ((255) & 63)))
    _ = p0
    _ = v6
    return -1
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = (v1 & v1)
    v3 := make(map[string]bool)
    v1 = 2
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: -1, F1: 42}
        switch -2147483648 {
        case 32768:
            if (3.14159 <= 0.0) {
                switch "foo bar" {
                case "\n\t":
                    if true {
                        fmt.Println(v1)
                        var v9 int = fi7
                        _ = v9
                    } else {
                        v10 := GenS1{F0: -812, F1: 254}
                        _ = v10
                    }
                case "a":
                    for fi11 := 0; fi11 < 8; fi11++ {
                        v2 += (v2 - (v2 % ((219) | 1)))
                        _ = fi11
                    }
                }
            }
        case 2147483647:
            fmt.Println(v2)
        case -355:
            var v12 uint = 369
            _ = v12
        default:
            fmt.Println(fi7, v2, v1)
        }
        _ = fi7
        _ = v8
    }
    var v13 int = 32767
    var v14 bool = (32767 != 2)
    var v15 uint = 65535
    if ((v15 != 16) && ((0.0 - (3.14159 + 3.14159)) != (3.14159 + (-2.25 * -2.25)))) {
        for fi16 := 0; fi16 < 12; fi16++ {
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 uint64 = (^(v2))
                v15 -= v15
                v19 := []int{(fi17 << ((v15) & 63))}
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi16
        }
        v20 := GenS2{F0: 42, F1: 682}
        _ = v20
    } else {
        var v21 int16 = -1
        fmt.Println(v15)
        _ = v21
    }
    switch v13 {
    case 1:
        if (3.14159 == ((((1e10 / ((1e10 + 1.5) * 3.14159)) - -2.25) + 0.1) / (((1e10 + (-2.25 - (3.14159 * 1e10))) / -2.25) * float64(3.14159)))) {
            delete(v3, "")
        }
    case 93:
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 float32 = (1e10 - ((((1e10 / (0.1 * 0.1)) + (((0.0 - 1e10) * 0.0) - -2.25)) * (((0.0 / (1.5 / 0.1)) / float32(float64(1e10))) - -2.25)) + float32(3.14159)))
            _ = fi22
            _ = v23
        }
    default:
        if false {
            v2 += v2
        } else {
            fmt.Println(v15, v14, v15)
            var v24 string = ""
            var v25 uint32 = 42
            _ = v24
            _ = v25
        }
    }
    v26 := GenS3{F0: (("" + "foo bar") + "\n\t"), F1: 732}
    for fi27 := 0; fi27 < 18; fi27++ {
        fmt.Println(v2, v15)
        var v28 uint32 = 934
        _ = fi27
        _ = v28
    }
    var v29 uint16 = 1
    fmt.Println(v14, v29)
    switch -527 {
    case 0:
        var v30 float64 = (((-2.25 / 3.14159) * 0.1) - (float64(0.0) * 1e10))
        _ = v30
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        for fi32 := 0; fi32 < 2; fi32++ {
            switch 127 {
            case 42:
                v33 := []int{(v13 ^ 972), (fi31 + 0)}
                _ = v33
            case 4294967295:
                if v14 {
                    var v34 uint = uint(v1)
                    v14 = (!false)
                    var v35 uint8 = 2
                    _ = v34
                    _ = v35
                }
            default:
                v36 := []int{fi31, 32768}
                _ = v36
            }
            _ = fi32
        }
        if v14 {
            fmt.Println(fi31, fi31)
            v29 += (v29 << ((uint(v29)) & 63))
        } else {
            var v37 int16 = 1
            v37 += (v37 | (v37 % (((-(v37))) | 1)))
            v38 := GenS2{F0: v37, F1: 65535}
            _ = v37
            _ = v38
        }
        _ = fi31
    }
    var v39 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v39
}
