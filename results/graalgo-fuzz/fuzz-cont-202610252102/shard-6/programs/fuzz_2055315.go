package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int64; F2 bool }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    var v3 int = 889
    var v4 uint = 42
    var v5 int64 = -1
    v6 := make(map[int]bool)
    v7 := GenS1{F0: 42}
    v8 := []int{int(v5), v3}
    if (true && (1 <= 0)) {
        v9 := GenS1{F0: 42}
        _ = v9
    } else {
        v3 -= v3
    }
    v10 := GenS2{F0: (false || true)}
    v11 := GenS2{F0: true}
    v12 := GenS2{F0: (65535 != 0)}
    for fi13 := 0; fi13 < 10; fi13++ {
        switch fi13 {
        case -925:
            for fi14 := 0; fi14 < 18; fi14++ {
                v3 = 32768
                _ = fi14
            }
        case -2147483648:
            var v15 float64 = 0.1
            _ = v15
        case 1:
            v16 := GenS1{F0: 1}
            _ = v16
        default:
            for fi17 := 0; fi17 < 20; fi17++ {
                fmt.Println(v4)
                delete(v1, 2147483647)
                _ = fi17
            }
        }
        _ = fi13
    }
    var v18 uint16 = 65535
    var v19 uint8 = 212
    var v20 float32 = float32(0.1)
    switch v3 {
    case 0:
        if true {
            v2[uint32(v19)] = false
        }
    case 4294967295:
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    case 9223372036854775807:
        delete(v2, 0)
    }
    if true {
        switch (v3 << ((4294967295) & 63)) {
        case -437:
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 float64 = 0.0
                if (float64(float32(3.14159)) <= v22) {
                    switch "foo bar" {
                    case "foo bar":
                        _ = len(v2)
                    }
                } else {
                    switch fi21 {
                    case 4294967295:
                        v22 -= (v22 / v22)
                    case 128:
                        v23 := v6[fi21]
                        _ = v23
                    default:
                        v18 -= (v18 << ((4294967295) & 63))
                    }
                    var v24 float64 = v22
                    _ = v24
                }
                switch fi21 {
                case 2147483647:
                    if ((4294967295 != 2147483647) || (-1 < 1)) {
                        switch 711 {
                        case 2147483647:
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        case -2147483648:
                            v19 -= 0
                        case 42:
                            v25 := GenS3{F0: 773}
                            _ = v25
                        default:
                            for fi26 := 0; fi26 < 8; fi26++ {
                                if true {
                                    var v27 uint = v4
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    _ = v27
                                }
                                v28 := GenS2{F0: (708 <= fi26)}
                                _ = fi26
                                _ = v28
                            }
                        }
                        var v29 uint8 = (v19 >> ((v4) & 63))
                        v29 = (v29 % (((v19 ^ (^(v19)))) | 1))
                        _ = v29
                    }
                default:
                    v5 = (v5 * v5)
                }
                _ = fi21
                _ = v22
            }
        }
    }
    switch v3 {
    case 1:
        v30 := GenS4{F0: -2147483648, F1: 9223372036854775807, F2: false}
        _ = v30
    default:
        switch "a" {
        case "\n\t":
            delete(v6, -155)
        case "foo bar":
            v19 = 2
        default:
            for fi31 := 0; fi31 < 20; fi31++ {
                var v32 float32 = (v20 - float32(((1e10 / float64((((((0.1 + v20) * 3.14159) + 3.14159) + v20) - v20))) - 3.14159)))
                var v33 bool = true
                var v34 int32 = 1
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
        }
    }
    v35 := GenS2{F0: true}
    for fi36 := 0; fi36 < 14; fi36++ {
        var v37 float64 = 3.14159
        if true {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            var v38 uint64 = 42
            for fi39 := 0; fi39 < 8; fi39++ {
                v5 += (v5 & v5)
                switch "foo bar" {
                case "\n\t":
                    var v40 int64 = v5
                    _ = v40
                case "a":
                    var v41 int8 = -128
                    _ = v41
                case "hello":
                    var v42 int8 = -120
                    _ = v42
                default:
                    if true {
                        v38 -= uint64(v4)
                        if len(v8) > 0 {
                            v8[0] = (fi36 << (((v4 * (v4 >> (((^(v4))) & 63)))) & 63))
                        }
                    } else {
                        v43 := GenS5{F0: (v19 / ((v19) | 1))}
                        var v44 bool = ("" == ("a" + "a"))
                        _ = v43
                        _ = v44
                    }
                }
                _ = fi39
            }
            _ = v38
        }
        _ = fi36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
