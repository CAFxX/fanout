package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int32, p1 string) uint {
    v4 := GenS1{F0: 18446744073709551615}
    v5 := []int{-9223372036854775808, 2147483647}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint64, p1 bool, p2 uint16) string {
    fmt.Println(p0)
    v6 := []int{-9223372036854775808, 575, 376, -924}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "\n\t"
}

func main() {
    var v1 bool = (2147483648 <= 54)
    var v2 string = "\n\t"
    var v3 uint8 = 85
    var v7 float32 = (3.14159 - ((float32(3.14159) / -2.25) + (-2.25 * 3.14159)))
    v7 += v7
    fmt.Println(v2, v1)
    fmt.Println(v2, v1, v1)
    v7 -= v7
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := []int{660, fi8}
        v7 += float32((float64(((v7 - v7) / float32(0.1))) * 1.5))
        var v10 bool = (1 != 348)
        _ = fi8
        _ = v10
        _ = v9
    }
    switch v2 {
    case "hello":
        fmt.Println(v2, v3)
    case "a":
        v1 = ((0 >= 2147483648) && ("hello" < "foo bar"))
    default:
        var v11 uint64 = 42
        _ = v11
    }
    var v12 float64 = (0.0 - 0.0)
    v13 := []int{}
    var v14 int8 = 100
    v3 -= (v3 >> ((65535) & 63))
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v3 = 0
    if (!false) {
        for fi15 := 0; fi15 < 6; fi15++ {
            if (v1 || (fi15 == (fi15 - (fi15 ^ fi15)))) {
                switch fi15 {
                case 1:
                    var v16 uint64 = 17
                    _ = v16
                case 127:
                    v1 = false
                case -48:
                    if len(v13) > 0 {
                        v13[0] = fi15
                    }
                default:
                    v12 = float64(0.0)
                }
                if v1 {
                    v13 = append(v13, 32767)
                }
            } else {
                var v17 float32 = (((0.1 / v7) + (v7 + (v7 - -2.25))) * float32((v12 * (((v12 / v12) * v12) + (v12 / ((v12 - -2.25) - v12))))))
                _ = v17
            }
            var v18 uint64 = 0
            _ = fi15
            _ = v18
        }
        fmt.Println(v1, v2)
    }
    var v19 uint16 = 243
    for fi20 := 0; fi20 < 10; fi20++ {
        if (262 == fi20) {
            fmt.Println(v3, v1, v1)
            v3 += v3
        } else {
            v19 -= 0
            v21 := GenS2{F0: (v3 + (v3 << ((18446744073709551615) & 63)))}
            v22 := GenS2{F0: 1}
            _ = v21
            _ = v22
        }
        var v23 int8 = v14
        v14 -= (v23 << ((839) & 63))
        _ = fi20
        _ = v23
    }
    v13 = append(v13, 127)
    for fi24 := 0; fi24 < 6; fi24++ {
        if v1 {
            if len(v13) > 0 {
                v13[0] = (-(fi24))
            }
            var v25 string = v2
            if v1 {
                switch ("\n\t" + ("\n\t" + (v25 + ("\n\t" + ("a" + ((v25 + ((v25 + "hello") + v25)) + (("a" + v2) + v25))))))) {
                case "\n\t":
                    v13 = append(v13, (fi24 >> ((1) & 63)))
                case "a":
                    for fi26 := 0; fi26 < 4; fi26++ {
                        v27 := GenS2{F0: 126}
                        var v28 int = 128
                        if (0.1 > float64((((v7 * (((0.0 + (v7 * v7)) - 3.14159) * ((0.0 - (1e10 + float32(0.1))) - float32((-2.25 / v12))))) - (v7 / 3.14159)) - (1e10 + v7)))) {
                            var v29 uint32 = 662
                            var v30 uint = 4294967295
                            if v1 {
                                v28 -= -598
                            }
                            _ = v29
                            _ = v30
                        }
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                case "foo bar":
                    var v31 int16 = 874
                    _ = v31
                default:
                    v19 -= v19
                }
                var v32 bool = ((fi24 % ((927) | 1)) != (fi24 << ((850) & 63)))
                if false {
                    if len(v13) > 0 {
                        v13[0] = fi24
                    }
                }
                _ = v32
            }
            _ = v25
        }
        v13 = append(v13, (fi24 - (-(fi24))))
        fmt.Println(v14)
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v7
}
