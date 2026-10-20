package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
func genFn0(p0 bool, p1 uint32) string {
    p0 = (32767 >= 1)
    var v6 float32 = float32(3.14159)
    p0 = false
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(v6, fi7)
        v6 = (((float32(1.5) / (v6 * -2.25)) / 1e10) - v6)
        _ = fi7
    }
    v6 = v6
    _ = p0
    _ = p1
    _ = v6
    return "hello"
}

func main() {
    v1 := []int{-1, 128, -9223372036854775808, 42}
    v2 := make(map[uint32]string)
    var v3 uint32 = 242
    v4 := make(map[uint32]bool)
    var v5 int16 = 189
    for fi8 := 0; fi8 < 7; fi8++ {
        switch fi8 {
        case 32767:
            if true {
                v9 := GenS1{F0: (fi8 % ((32767) | 1))}
                var v10 uint8 = 195
                _ = v10
                _ = v9
            }
        case 127:
            v3 = v3
        case 4294967295:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            fmt.Println(fi8, v3, fi8)
        }
        var v11 string = "a"
        v11 = v11
        _ = fi8
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        if ((0 != 4294967295) && (!(255 < 4294967295))) {
            v5 = v5
            fmt.Println(fi12, fi12, fi12)
            v13 := GenS2{F0: 208, F1: 18446744073709551615, F2: 897}
            _ = v13
        } else {
            for fi14 := 0; fi14 < 5; fi14++ {
                _ = len(v1)
                _ = fi14
            }
        }
        _ = fi12
    }
    v4[v3] = (18 < -9223372036854775808)
    if true {
        v15 := GenS1{F0: -356}
        switch "" {
        case "hello":
            if true {
                var v16 int32 = 42
                if true {
                    v17 := GenS3{F0: ("foo bar" + "foo bar")}
                    v18 := GenS1{F0: 464}
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v3 -= v3
                        _ = fi19
                    }
                    _ = v17
                    _ = v18
                } else {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        fmt.Println(v3, v5, v5)
                        _ = fi20
                    }
                    v21 := GenS4{F0: 1}
                    _ = v21
                }
                _ = v16
            }
        case "":
            switch -272 {
            case 903:
                var v22 int16 = (v5 << ((4294967295) & 63))
                _ = v22
            case -9223372036854775808:
                v5 = v5
            default:
                var v23 int64 = 42
                _ = v23
            }
        default:
            v5 += (^(v5))
        }
        _ = len(v4)
        _ = v15
    } else {
        for fi24 := 0; fi24 < 19; fi24++ {
            v5 = (v5 % ((v5) | 1))
            v25 := GenS2{F0: 127, F1: 1, F2: 18446744073709551615}
            _ = fi24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 uint = 2
        v3 -= 262
        for fi28 := 0; fi28 < 16; fi28++ {
            if ((true && (42 <= 14)) || false) {
                v29 := GenS3{F0: ""}
                _ = v29
            } else {
                var v30 uint64 = 1
                if len(v1) > 0 {
                    v1[0] = fi28
                }
                var v31 int64 = 9223372036854775807
                _ = v30
                _ = v31
            }
            var v32 int8 = 0
            fmt.Println(fi26)
            _ = fi28
            _ = v32
        }
        _ = fi26
        _ = v27
    }
    v5 += 32767
    var v33 bool = (0.1 >= ((1.5 * ((0.0 - ((((0.0 * 1.5) - 1.5) / 3.14159) * (0.0 * float64(0.0)))) - (((1e10 * ((float64(3.14159) - float64(-2.25)) * (((-2.25 * ((1e10 + 0.1) / float64(((((float32(3.14159) + 3.14159) / ((0.1 - -2.25) * 3.14159)) - (float32(((3.14159 * 0.1) - (3.14159 * -2.25))) * (0.1 + float32(0.1)))) + 3.14159)))) + 1.5) - (float64(1.5) / -2.25)))) + 0.1) / 1e10))) / -2.25))
    v34 := v2[v3]
    v5 = v5
    v35 := GenS1{F0: 32768}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
