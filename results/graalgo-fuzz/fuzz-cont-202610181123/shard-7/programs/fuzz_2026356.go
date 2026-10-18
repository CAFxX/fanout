package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int8 }

func main() {
    var v1 int8 = 0
    var v2 int8 = -53
    v3 := []uint32{1, 0, 2147483648}
    var v4 int8 = v1
    v5 := make(map[uint32]uint32)
    v6 := GenS1{F0: 289, F1: 292}
    var v7 int32 = -2147483648
    v8 := v5[963]
    v8 += v8
    if (42 < -379) {
        var v9 int32 = v7
        _ = v9
    } else {
        v4 = (^(v1))
        fmt.Println(v4, v7, v8)
    }
    fmt.Println(v2, v8)
    var v10 float32 = 3.14159
    for fi11 := 0; fi11 < 9; fi11++ {
        v5[(v8 >> ((248) & 63))] = 2147483648
        for fi12 := 0; fi12 < 16; fi12++ {
            fmt.Println(v7)
            _ = len(v5)
            _ = fi12
        }
        for fi13 := 0; fi13 < 5; fi13++ {
            fmt.Println(v10, v8, v10)
            var v14 uint16 = 1
            _ = fi13
            _ = v14
        }
        _ = fi11
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 uint64 = 1
        if ((-(v4)) > 127) {
            v17 := GenS1{F0: 18446744073709551615, F1: (fi15 & fi15)}
            var v18 uint64 = 384
            _ = v17
            _ = v18
        } else {
            var v19 uint8 = 127
            _ = v19
        }
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 float64 = 1.5
            _ = fi20
            _ = v21
        }
        _ = fi15
        _ = v16
    }
    if false {
        var v22 bool = true
        if (true && (float64((((v10 / v10) + ((v10 * (1.5 / 1e10)) + (((3.14159 - (-2.25 / (float32(1e10) + float32(-2.25)))) - float32(3.14159)) - (v10 * (float32((0.0 - 1.5)) * v10))))) * float32(((1.5 / -2.25) - 3.14159)))) > 1.5)) {
            v23 := GenS2{F0: 50, F1: v22, F2: 122}
            _ = v23
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            v3 = append(v3, (v8 | v8))
            if v22 {
                v1 = 75
                if v22 {
                    fmt.Println(fi24, v4)
                }
            }
            _ = fi24
        }
        _ = v22
    }
    if (-52 < v7) {
        v25 := GenS1{F0: 4294967295, F1: -879}
        if true {
            for fi26 := 0; fi26 < 10; fi26++ {
                switch "" {
                case "\n\t":
                    for fi27 := 0; fi27 < 17; fi27++ {
                        var v28 int16 = 390
                        var v29 int8 = 0
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                default:
                    var v30 float32 = v10
                    _ = v30
                }
                fmt.Println(v8)
                _ = fi26
            }
            if true {
                v10 += (v10 * 1.5)
            } else {
                v8 -= 65535
            }
        } else {
            if (187 >= 32768) {
                for fi31 := 0; fi31 < 7; fi31++ {
                    for fi32 := 0; fi32 < 20; fi32++ {
                        v33 := GenS3{F0: -828, F1: -696}
                        v10 += v10
                        fmt.Println(v7, v8)
                        _ = fi32
                        _ = v33
                    }
                    fmt.Println(v2, v4)
                    _ = fi31
                }
            }
            var v34 int32 = (v7 & v7)
            switch "a" {
            case "\n\t":
                v34 = 2147483647
            case "foo bar":
                for fi35 := 0; fi35 < 13; fi35++ {
                    switch "hello" {
                    case "\n\t":
                        _ = len(v5)
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    var v36 uint8 = uint8(v8)
                    var v37 int64 = 0
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            case "hello":
                v5[v8] = v8
            }
            _ = v34
        }
        for fi38 := 0; fi38 < 4; fi38++ {
            v10 -= 0.0
            v39 := GenS3{F0: 1000, F1: 282}
            if (!false) {
                var v40 int = (fi38 / ((fi38) | 1))
                _ = v40
            }
            _ = fi38
            _ = v39
        }
        _ = v25
    }
    if true {
        v41 := GenS3{F0: 1, F1: (v7 ^ 821)}
        _ = v41
    }
    if (!false) {
        v4 = (v2 << ((869) & 63))
        v42 := GenS4{F0: 42, F1: -1}
        _ = v42
    }
    if ((!((v7 + v7) < -2147483648)) && false) {
        var v43 int64 = -9223372036854775808
        if (!(false && true)) {
            v8 = (v8 - v8)
        }
        _ = v43
    } else {
        var v44 uint32 = v8
        v45 := GenS4{F0: 401, F1: 32768}
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 3; fi46++ {
        v7 = (-(v7))
        _ = fi46
    }
    fmt.Println(v1, v1, v2)
    switch ("" + (("foo bar" + "") + (("a" + "a") + ""))) {
    case "a":
        v47 := v5[v8]
        _ = v47
    default:
        v48 := GenS5{F0: -59}
        _ = v48
    }
    if (true || true) {
        var v49 uint = 2
        _ = v49
    } else {
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
