package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int = 4294967295
    var v2 uint = 0
    var v3 int = (v1 * (^(v1)))
    var v4 uint = (v2 - 65535)
    var v5 float64 = (float64((float32((1.5 - 3.14159)) - float32(0.0))) - 3.14159)
    for fi6 := 0; fi6 < 8; fi6++ {
        switch fi6 {
        case 42:
            var v7 float64 = float64(1.5)
            _ = v7
        case 693:
            v8 := GenS1{F0: 18446744073709551615}
            _ = v8
        case 4294967295:
            if true {
                if false {
                    for fi9 := 0; fi9 < 17; fi9++ {
                        v4 += v4
                        v5 -= v5
                        _ = fi9
                    }
                }
            } else {
                var v10 bool = ((0 < 42) || (1000 >= 1000))
                _ = v10
            }
        default:
            var v11 uint8 = 86
            _ = v11
        }
        _ = fi6
    }
    v5 -= (v5 + (0.0 - 3.14159))
    v12 := []int{(v1 >> (((v4 | 4294967295)) & 63)), v1, v3}
    var v13 int32 = -2147483648
    v2 -= (v4 << (((v4 << (((v4 + (^(v4)))) & 63))) & 63))
    switch "\n\t" {
    case "foo bar":
        v13 += v13
    }
    fmt.Println(v13, v1)
    switch "a" {
    case "":
        v1 = v3
    case "foo bar":
        if false {
            if ((!true) && (0 > 2)) {
                if len(v12) > 0 {
                    v12[0] = (-(v3))
                }
            } else {
                if true {
                    fmt.Println(v2)
                } else {
                    _ = len(v12)
                    var v14 uint64 = 42
                    _ = v14
                }
            }
            fmt.Println(v1)
            if false {
                v12 = append(v12, v1)
            } else {
                v3 -= (v1 & (v3 & (v3 * (v3 ^ v1))))
                if (v4 >= 65535) {
                    var v15 uint32 = 1
                    _ = v15
                } else {
                    var v16 uint16 = 162
                    v2 += (v4 % ((4294967295) | 1))
                    var v17 int = int(v4)
                    _ = v16
                    _ = v17
                }
                v5 -= float64((0.0 + float32(0.1)))
            }
        }
    case "\n\t":
        v18 := GenS1{F0: v4}
        _ = v18
    }
    v3 += (v1 % ((32767) | 1))
    v4 = 65535
    for fi19 := 0; fi19 < 16; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            if len(v12) > 0 {
                v12[0] = 42
            }
            v4 = v2
            _ = fi20
        }
        _ = fi19
    }
    v2 -= (v2 >> ((v4) & 63))
    for fi21 := 0; fi21 < 20; fi21++ {
        for fi22 := 0; fi22 < 4; fi22++ {
            _ = len(v12)
            _ = fi22
        }
        var v23 int8 = -128
        _ = fi21
        _ = v23
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        if len(v12) > 0 {
            v12[0] = fi24
        }
        _ = fi24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 uint64 = 188
        for fi27 := 0; fi27 < 8; fi27++ {
            fmt.Println(v2, v26)
            _ = fi27
        }
        switch "" {
        case "\n\t":
            v1 = (^(v3))
        case "":
            fmt.Println(v5)
        default:
            v1 -= 778
        }
        _ = fi25
        _ = v26
    }
    fmt.Println(v4, v5, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
