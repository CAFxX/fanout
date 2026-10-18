package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint }
func genFn0(p0 int8, p1 int, p2 int64) int {
    p0 = p0
    if (int(p2) >= p1) {
        p0 = 0
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 - -2147483648)
}
func genFn1(p0 int16, p1 uint16) bool {
    p0 += (^(p0))
    if (0.0 <= 0.1) {
        p0 = p0
        var v7 uint32 = 42
        _ = v7
    }
    v8 := GenS1{F0: -2147483648, F1: 0, F2: 2147483647}
    _ = p0
    _ = p1
    _ = v8
    return (0 != 127)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    var v3 int16 = 670
    var v4 float32 = 0.0
    v5 := []uint32{2147483647, 0, 301}
    var v6 bool = (!true)
    v4 -= (3.14159 - float32(float64(3.14159)))
    v2[("\n\t" + ("" + ("hello" + ("" + "foo bar"))))] = 0
    for fi9 := 0; fi9 < 12; fi9++ {
        v3 = v3
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 uint64 = 0
            var v12 int = (fi9 & genFn0(20, fi10, -1))
            v6 = false
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 int32 = -2147483648
        _ = fi9
        _ = v13
    }
    fmt.Println(v4)
    v14 := GenS1{F0: genFn0(127, 394, 9223372036854775807), F1: v3, F2: 4294967295}
    fmt.Println(v3, v3)
    for fi15 := 0; fi15 < 16; fi15++ {
        fmt.Println(fi15, v3, v6)
        v3 -= (-(v3))
        for fi16 := 0; fi16 < 11; fi16++ {
            v6 = (386 < 255)
            _ = fi16
        }
        _ = fi15
    }
    fmt.Println(v4, v6, v4)
    switch "\n\t" {
    case "a":
        fmt.Println(v3, v3)
    case "hello":
        v17 := GenS1{F0: 0, F1: (v3 % ((v3) | 1)), F2: 65535}
        _ = v17
    case "":
        var v18 float64 = 3.14159
        _ = v18
    default:
        for fi19 := 0; fi19 < 5; fi19++ {
            v6 = (686 != 788)
            for fi20 := 0; fi20 < 10; fi20++ {
                v6 = (3.14159 == ((3.14159 - -2.25) - ((-2.25 + float64((v4 - (v4 * v4)))) / 0.1)))
                var v21 int64 = 89
                _ = fi20
                _ = v21
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi19
        }
    }
    switch -9223372036854775808 {
    case 32768:
        if (1 <= genFn0(0, -2147483648, -976)) {
            var v22 int32 = -727
            switch 4294967295 {
            case 127:
                switch 715 {
                case -499:
                    if len(v5) > 0 {
                        v5[0] = 4294967295
                    }
                case -571:
                    v1[42] = ""
                case -9223372036854775808:
                    var v23 int16 = (v3 << ((0) & 63))
                    _ = v23
                default:
                    v3 = v3
                }
            default:
                for fi24 := 0; fi24 < 6; fi24++ {
                    if v6 {
                        fmt.Println(v22)
                    } else {
                        var v25 float64 = 3.14159
                        _ = v25
                    }
                    _ = fi24
                }
            }
            _ = v22
        }
    case -2147483648:
        var v26 int16 = v3
        _ = v26
    }
    var v27 uint = 2
    fmt.Println(v4, v3, v3)
    v6 = v6
    v28 := GenS1{F0: 1, F1: v3, F2: 1}
    for fi29 := 0; fi29 < 13; fi29++ {
        var v30 int64 = 1
        if (!(v6 && true)) {
            v30 -= (^(v30))
            fmt.Println(v27)
        } else {
            v4 = 0.1
            v4 = v4
            v3 = v3
        }
        if (905 != 42) {
            v4 += 3.14159
        }
        _ = fi29
        _ = v30
    }
    if ((0 == -1) || (v4 > v4)) {
        var v31 uint = v27
        _ = v31
    } else {
        _ = len(v1)
        for fi32 := 0; fi32 < 2; fi32++ {
            v33 := GenS2{F0: 127, F1: (v27 << ((255) & 63))}
            var v34 uint16 = 2
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    v35 := GenS1{F0: genFn0(100, 1, 42), F1: v3, F2: 65535}
    v27 += v27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
