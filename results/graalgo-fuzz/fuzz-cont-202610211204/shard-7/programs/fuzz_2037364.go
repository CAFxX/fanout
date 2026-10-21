package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 bool }

func main() {
    var v1 uint16 = 0
    var v2 uint = 445
    var v3 float32 = float32(((3.14159 * (0.1 * 0.0)) - 3.14159))
    v4 := []int{-9223372036854775808}
    v5 := []int{4294967295, 1}
    v6 := []string{"hello"}
    switch 705 {
    case 700:
        v3 = (v3 * v3)
    }
    v3 -= 3.14159
    var v7 uint = 348
    _ = len(v4)
    var v8 int = -602
    v9 := GenS1{F0: (-(v8)), F1: (^(v1))}
    for fi10 := 0; fi10 < 20; fi10++ {
        switch ("" + ("" + "a")) {
        case "foo bar":
            switch "\n\t" {
            case "\n\t":
                v11 := GenS2{F0: 254, F1: false, F2: 44}
                _ = v11
            case "":
                var v12 float64 = (0.0 - (-2.25 - ((0.1 - 1e10) + float64(v3))))
                _ = v12
            case "foo bar":
                v1 = v1
            }
        case "\n\t":
            switch (v8 * (v8 + (fi10 << ((v2) & 63)))) {
            case -2147483648:
                var v13 bool = (0 < 0)
                _ = v13
            case 1:
                v14 := GenS1{F0: 9223372036854775807, F1: 0}
                _ = v14
            }
        case "a":
            v3 += 1e10
        default:
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 uint16 = 42
                v17 := GenS3{F0: (v2 * (^(v7)))}
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        _ = fi10
    }
    var v18 float32 = -2.25
    for fi19 := 0; fi19 < 8; fi19++ {
        if (0 < 128) {
            var v20 uint16 = v1
            if (255 == 18) {
                v4 = append(v4, 9223372036854775807)
            }
            fmt.Println(v8, v2)
            _ = v20
        }
        var v21 int16 = 896
        var v22 uint16 = (v1 % ((uint16(v7)) | 1))
        _ = fi19
        _ = v21
        _ = v22
    }
    var v23 string = "foo bar"
    switch (v8 ^ (v8 >> ((4294967295) & 63))) {
    case 4294967295:
        v24 := GenS4{F0: (v7 ^ v7)}
        _ = v24
    default:
        for fi25 := 0; fi25 < 1; fi25++ {
            var v26 string = v23
            v1 = (v1 >> (((v7 ^ v7)) & 63))
            _ = fi25
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        v2 = (v2 - 2)
        v23 = (("\n\t" + v23) + "a")
        for fi28 := 0; fi28 < 1; fi28++ {
            if len(v5) > 0 {
                v5[0] = (fi27 >> ((v2) & 63))
            }
            v29 := GenS4{F0: v7}
            var v30 float32 = 1e10
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi27
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        var v32 int16 = -882
        for fi33 := 0; fi33 < 2; fi33++ {
            var v34 int = (fi31 << ((v2) & 63))
            fmt.Println(v2, v32, v7)
            if ((v2 * 436) <= v7) {
                fmt.Println(v2, v23, v23)
                if true {
                    v35 := GenS5{F0: (v34 <= (fi33 << ((v2) & 63))), F1: -51}
                    var v36 float64 = (1.5 * 3.14159)
                    _ = v35
                    _ = v36
                } else {
                    v7 -= v2
                }
            }
            _ = fi33
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    if (true || (0 <= 0)) {
        v37 := GenS3{F0: (v7 ^ v2)}
        var v38 uint8 = 2
        v4 = append(v4, (v8 + -35))
        _ = v37
        _ = v38
    } else {
        v39 := GenS2{F0: 255, F1: false, F2: 128}
        switch (v23 + (v23 + "a")) {
        case "\n\t":
            if ((v8 % (((v8 - 698)) | 1)) >= v8) {
                v1 = (v1 * (v1 - (v1 << ((v7) & 63))))
            }
        case "hello":
            v2 = (v7 >> ((v7) & 63))
        case "a":
            var v40 int = v8
            _ = v40
        default:
            for fi41 := 0; fi41 < 16; fi41++ {
                v18 = v3
                if false {
                    var v42 int = v8
                    if (2 >= 1) {
                        var v43 int32 = 1
                        switch (("" + v23) + v23) {
                        case "hello":
                            if (0 <= 65535) {
                                v2 += (^(v2))
                            } else {
                                v6 = append(v6, v23)
                                switch v23 {
                                case "hello":
                                    for fi44 := 0; fi44 < 3; fi44++ {
                                        v45 := GenS5{F0: (4294967295 < 2), F1: 127}
                                        v18 = (float32(-2.25) * v3)
                                        v46 := GenS6{F0: (!true)}
                                        _ = fi44
                                        _ = v45
                                        _ = v46
                                    }
                                case "":
                                    if (1 == 42) {
                                        v47 := GenS5{F0: ((1 > fi41) || false), F1: 100}
                                        _ = v47
                                    }
                                default:
                                    v18 = (float32(-2.25) - ((float32(3.14159) + float32(1e10)) - (float32(((1.5 * 0.1) - (float64(((v18 / 3.14159) / -2.25)) / -2.25))) / 1.0)))
                                }
                            }
                        }
                        _ = v43
                    } else {
                        fmt.Println(v23)
                        var v48 uint = v2
                        if (false && true) {
                            var v49 int = (-(v8))
                            _ = v49
                        }
                        _ = v48
                    }
                    _ = len(v6)
                    _ = v42
                }
                _ = fi41
            }
        }
        var v50 int64 = -107
        _ = v39
        _ = v50
    }
    var v51 float64 = ((3.14159 * 0.1) - 1.5)
    fmt.Println(v18)
    v4 = append(v4, -207)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
