package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 int64; F2 uint16 }

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]int)
    var v3 int16 = 1
    var v4 float64 = (-2.25 / 1e10)
    var v5 uint = 4294967295
    var v6 uint8 = 69
    var v7 uint16 = 255
    v5 -= 592
    v8 := []int{9223372036854775807, 32767, 2147483647, 128}
    var v9 uint64 = 18446744073709551615
    if true {
        switch int(v5) {
        case 32767:
            switch "hello" {
            case "a":
                fmt.Println(v3, v4)
            }
        case 897:
            v5 += v5
        }
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 uint64 = uint64(v6)
        v1 = 0
        _ = fi10
        _ = v11
    }
    switch "foo bar" {
    case "a":
        if (!(!((v6 >> ((0) & 63)) >= (^(v6))))) {
            for fi12 := 0; fi12 < 8; fi12++ {
                fmt.Println(v1)
                var v13 bool = true
                _ = fi12
                _ = v13
            }
        } else {
            if true {
                for fi14 := 0; fi14 < 20; fi14++ {
                    if true {
                        switch fi14 {
                        case 0:
                            v9 = (v9 - v9)
                        case -2147483648:
                            fmt.Println(v1, v7, fi14)
                        default:
                            for fi15 := 0; fi15 < 11; fi15++ {
                                v16 := GenS1{F0: v6, F1: v1}
                                _ = fi15
                                _ = v16
                            }
                        }
                    } else {
                        v4 = (v4 + ((v4 - ((v4 / -2.25) * ((v4 - (v4 + (v4 - (float64(3.14159) + ((v4 * 0.1) - float64(((((-2.25 - 0.0) * 0.0) * (((1.5 - 1.5) / 1.5) - (3.14159 / 1e10))) / 1e10))))))) / v4))) / v4))
                    }
                    _ = fi14
                }
                fmt.Println(v9)
            }
        }
    case "hello":
        var v17 int8 = 127
        _ = v17
    case "":
        if (425 >= -1) {
            fmt.Println(v5)
            if false {
                fmt.Println(v4)
                v3 = (^(v3))
                v18 := GenS2{F0: (!true)}
                _ = v18
            } else {
                v5 = (^(v5))
                var v19 uint32 = 4294967295
                _ = v19
            }
        }
    }
    v20 := GenS3{F0: -1}
    delete(v2, 0)
    v9 += v9
    var v21 int = 128
    var v22 uint16 = (^(v1))
    if (!false) {
        fmt.Println(v4, v6)
        if true {
            var v23 int16 = (v3 << (((v5 - v5)) & 63))
            for fi24 := 0; fi24 < 14; fi24++ {
                var v25 float64 = v4
                _ = fi24
                _ = v25
            }
            _ = v23
        }
    } else {
        if (-1 < 116) {
            v26 := GenS2{F0: true}
            _ = v26
        }
        if false {
            v3 += (-(v3))
        } else {
            v27 := GenS4{F0: -68}
            v28 := GenS4{F0: 127}
            _ = v27
            _ = v28
        }
    }
    var v29 string = ""
    if (true || (-1 == 314)) {
        v29 = v29
        v30 := GenS2{F0: (!((!true) && ((1 > 9223372036854775807) || true)))}
        v3 = (v3 % ((v3) | 1))
        _ = v30
    } else {
        v31 := GenS5{F0: v1, F1: 0, F2: v1}
        for fi32 := 0; fi32 < 1; fi32++ {
            v33 := GenS1{F0: uint8(v7), F1: (v7 + (v7 - v22))}
            if (false || true) {
                v5 = (v5 % ((255) | 1))
            } else {
                if (!false) {
                    switch v29 {
                    case "foo bar":
                        v34 := GenS5{F0: v22, F1: 801, F2: 1}
                        _ = v34
                    case "a":
                        v29 = (((((("\n\t" + v29) + v29) + v29) + ("" + ((v29 + v29) + (("a" + "a") + "hello")))) + v29) + (("" + "") + "a"))
                    case "hello":
                        switch ((v29 + ("a" + "foo bar")) + "") {
                        case "hello":
                            for fi35 := 0; fi35 < 8; fi35++ {
                                var v36 int64 = -425
                                var v37 int8 = 71
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                        default:
                            var v38 int32 = 0
                            _ = v38
                        }
                    default:
                        _ = len(v8)
                    }
                }
            }
            _ = fi32
            _ = v33
        }
        var v39 int = (v21 << ((v5) & 63))
        _ = v31
        _ = v39
    }
    for fi40 := 0; fi40 < 11; fi40++ {
        v6 += (v6 & (v6 % (((^(v6))) | 1)))
        _ = fi40
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        v3 = 299
        var v42 uint = v5
        if (v6 <= (v6 * uint8(v22))) {
            v21 += 1
            v1 = 255
            if (true || false) {
                v9 -= (v9 % ((0) | 1))
                if (1e10 != (v4 * 1e10)) {
                    v42 -= v42
                    v43 := GenS2{F0: false}
                    _ = v43
                }
            }
        }
        _ = fi41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
