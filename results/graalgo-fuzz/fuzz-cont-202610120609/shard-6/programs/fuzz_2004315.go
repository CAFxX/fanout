package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 int8; F1 int8; F2 int16 }
func genFn0(p0 uint64, p1 uint64) int8 {
    p0 += (p1 / (((p0 & (^(p0)))) | 1))
    if (1 >= -693) {
        v5 := GenS1{F0: true}
        _ = v5
    } else {
        var v6 uint32 = 2
        v7 := []int{9223372036854775807, -342, 1}
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    return 100
}
func genFn1(p0 int16, p1 int64, p2 bool) int32 {
    p2 = (728 == 42)
    if (0 != 2147483647) {
        switch ("hello" + (("a" + ("" + "a")) + "a")) {
        case "a":
            var v8 int16 = p0
            _ = v8
        case "\n\t":
            p1 -= (p1 >> ((866) & 63))
        case "hello":
            v9 := GenS1{F0: ((!false) || (0 > 8))}
            _ = v9
        default:
            fmt.Println(p2, p2)
        }
        switch -367 {
        case 4294967295:
            p2 = p2
        default:
            var v10 uint8 = 93
            _ = v10
        }
    }
    var v11 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (v11 & (v11 | 324))
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 42
    var v3 bool = false
    var v4 float32 = float32(float64((1e10 / 1.5)))
    fmt.Println(v3)
    v4 -= -2.25
    fmt.Println(v2, v2)
    v12 := GenS1{F0: (1 == 423)}
    v2 -= v2
    var v13 int64 = 826
    for fi14 := 0; fi14 < 20; fi14++ {
        switch ("" + "hello") {
        case "hello":
            var v15 uint = 117
            _ = v15
        }
        _ = fi14
    }
    v16 := v1[9223372036854775807]
    var v17 string = (v16 + (("hello" + v16) + ("" + (v16 + (v16 + v16)))))
    switch ("hello" + v17) {
    case "hello":
        switch v17 {
        case "a":
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 int64 = 42
                v16 = (v17 + v17)
                v20 := []int{}
                _ = fi18
                _ = v19
                _ = v20
            }
        case "":
            if true {
                var v21 int = 243
                var v22 uint64 = (v2 & v2)
                v23 := GenS1{F0: v3}
                _ = v21
                _ = v22
                _ = v23
            } else {
                var v24 float64 = -2.25
                for fi25 := 0; fi25 < 20; fi25++ {
                    if v3 {
                        v26 := v1[fi25]
                        for fi27 := 0; fi27 < 9; fi27++ {
                            if true {
                                v28 := GenS2{F0: int16(fi25)}
                                if true {
                                    var v29 uint32 = 4294967295
                                    fmt.Println(v3, v4, v2)
                                    var v30 int8 = -42
                                    _ = v29
                                    _ = v30
                                }
                                _ = v28
                            } else {
                                var v31 int8 = 42
                                _ = v31
                            }
                            v32 := GenS3{F0: 742}
                            v33 := GenS4{F0: 0}
                            _ = fi27
                            _ = v32
                            _ = v33
                        }
                        _ = v26
                    }
                    _ = fi25
                }
                v34 := GenS5{F0: genFn0(4294967295, (v2 << ((255) & 63))), F1: 2147483647}
                _ = v24
                _ = v34
            }
        case "\n\t":
            v35 := GenS2{F0: -110}
            _ = v35
        default:
            v2 = 42
        }
    case "foo bar":
        var v36 float64 = -2.25
        _ = v36
    case "a":
        v37 := GenS6{F0: 127, F1: genFn0(v2, (v2 ^ (v2 ^ 4294967295))), F2: 42}
        _ = v37
    default:
        if (!v3) {
            v38 := GenS5{F0: 4, F1: genFn1(42, v13, v3)}
            v39 := GenS4{F0: genFn1(1000, v13, (v3 && v3))}
            _ = v38
            _ = v39
        } else {
            for fi40 := 0; fi40 < 13; fi40++ {
                v41 := v1[254]
                for fi42 := 0; fi42 < 17; fi42++ {
                    var v43 int32 = 42
                    v44 := GenS4{F0: v43}
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                fmt.Println(v13, v4)
                _ = fi40
                _ = v41
            }
        }
    }
    v45 := v1[-1]
    var v46 uint16 = 136
    fmt.Println(v16, v45)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
}
