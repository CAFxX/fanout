package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int) uint {
    v5 := GenS1{F0: 1, F1: "hello", F2: 859}
    fmt.Println(p0)
    var v6 string = "\n\t"
    switch v6 {
    case "":
        v7 := []int{}
        _ = v7
    default:
        v6 = ("\n\t" + "hello")
    }
    v8 := GenS1{F0: 671, F1: v6, F2: 561}
    _ = p0
    _ = v5
    _ = v6
    _ = v8
    return 4294967295
}

func main() {
    v1 := []int{}
    var v2 uint8 = 0
    var v3 int16 = -32768
    v4 := make(map[int]uint32)
    switch "foo bar" {
    case "":
        for fi9 := 0; fi9 < 8; fi9++ {
            delete(v4, fi9)
            _ = fi9
        }
    case "\n\t":
        if ((v3 % ((v3) | 1)) > (-(v3))) {
            v3 += v3
            for fi10 := 0; fi10 < 4; fi10++ {
                if true {
                    fmt.Println(v2)
                    for fi11 := 0; fi11 < 7; fi11++ {
                        v12 := GenS1{F0: -571, F1: "hello", F2: 913}
                        var v13 uint64 = 184
                        switch ("hello" + (("\n\t" + ("" + "\n\t")) + ("foo bar" + "hello"))) {
                        case "a":
                            var v14 uint = genFn0(fi11)
                            _ = v14
                        case "foo bar":
                            v2 = 2
                        }
                        var v15 bool = (0.1 > 0.1)
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 2; fi16++ {
                        v17 := GenS2{F0: -1}
                        _ = fi16
                        _ = v17
                    }
                }
                _ = fi10
            }
        } else {
            var v18 uint8 = v2
            _ = v18
        }
    }
    v19 := GenS1{F0: -224, F1: ("" + ""), F2: 42}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v20 := GenS3{F0: ((((float64(1.5) - float64(-2.25)) - 1.5) - 3.14159) != (0.0 / (-2.25 - ((float64((float32(float64((((((0.1 * float32(((((((0.0 - ((((-2.25 + float64((1e10 / 1.5))) - (1e10 / 1.0)) / (0.0 - 1.5)) * -2.25)) * ((1e10 + (1e10 / (float64(1e10) - 1.5))) + 0.1)) - (3.14159 / 1e10)) * 0.1) + 0.1) / 1.0))) + (-2.25 * float32(float64(((float32(0.0) / 1.5) - 1e10))))) + ((-2.25 - (-2.25 * 0.0)) + 0.0)) + ((((-2.25 - (1e10 * -2.25)) / float32(((float64(((-2.25 + (-2.25 / 3.14159)) * 1e10)) - 1.5) + 1.5))) + 1e10) / 1.5)) - 1e10))) + -2.25)) / 1.5) * 0.0)))), F1: 807}
    fmt.Println(v3)
    var v21 float32 = 0.1
    v21 += v21
    for fi22 := 0; fi22 < 17; fi22++ {
        switch fi22 {
        case 42:
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v2 -= (v2 + v2)
            }
        case 32768:
            var v23 float32 = 0.1
            _ = v23
        default:
            fmt.Println(fi22, fi22)
        }
        switch "" {
        case "\n\t":
            var v24 int32 = -1
            _ = v24
        case "foo bar":
            for fi25 := 0; fi25 < 7; fi25++ {
                v26 := GenS4{F0: 42, F1: 936}
                _ = fi25
                _ = v26
            }
        case "a":
            switch 128 {
            case -9223372036854775808:
                if ((int(v3) != (-(fi22))) && ((v3 - -1) >= v3)) {
                    if true {
                        v3 -= 32767
                    }
                    v2 += (^(v2))
                }
            case 127:
                if false {
                    v2 += v2
                } else {
                    var v27 bool = ((true && false) || ((251 <= 2) && ((510 == 0) || (!(18446744073709551615 != 853)))))
                    _ = v27
                }
            case 2147483647:
                v3 += -1
            }
        default:
            var v28 uint16 = 42
            _ = v28
        }
        _ = fi22
    }
    var v29 float32 = float32(3.14159)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v4
}
