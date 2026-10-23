package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 uint }
func genFn0() int {
    v5 := []int{-127, 127, 1, 458}
    var v6 uint64 = 2
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint32 = 2147483648
        fmt.Println(v6, fi7, v8)
        if ((3.14159 <= 0.1) || (!(false && ("foo bar" != "\n\t")))) {
            switch 32767 {
            case -1:
                v9 := GenS1{F0: 39, F1: 2}
                _ = v9
            }
            if true {
                v6 = v6
            }
        } else {
            v6 = (v6 % ((v6) | 1))
            switch ("" + (("\n\t" + ("" + "")) + ("" + "a"))) {
            case "hello":
                v6 = v6
            case "foo bar":
                if (true && (42 == 190)) {
                    switch "" {
                    case "":
                        fmt.Println(v6, v6, fi7)
                    case "\n\t":
                        v10 := GenS1{F0: 127, F1: 176}
                        _ = v10
                    case "hello":
                        switch ("" + "\n\t") {
                        case "":
                            v11 := GenS1{F0: 42, F1: 42}
                            _ = v11
                        case "foo bar":
                            if false {
                                v12 := GenS1{F0: 255, F1: 1}
                                fmt.Println(v8)
                                for fi13 := 0; fi13 < 4; fi13++ {
                                    if (4294967295 < 963) {
                                        var v14 uint8 = 0
                                        if (false && false) {
                                            fmt.Println(v8, fi13, v6)
                                            fmt.Println(v14, v14, fi13)
                                        }
                                        _ = v14
                                    } else {
                                        var v15 int8 = 65
                                        switch fi13 {
                                        case -552:
                                            var v16 uint8 = 47
                                            _ = v16
                                        }
                                        switch fi13 {
                                        case 42:
                                            var v17 int32 = -2147483648
                                            _ = v17
                                        case 9223372036854775807:
                                            v6 = (v6 % ((167) | 1))
                                        case -361:
                                            switch "a" {
                                            case "\n\t":
                                                var v18 uint32 = (v8 * 2147483647)
                                                _ = v18
                                            case "":
                                                v6 -= v6
                                                var v19 bool = (false || false)
                                                _ = v19
                                            default:
                                                switch (^(fi13)) {
                                                case 1:
                                                    v20 := GenS2{F0: 18446744073709551615, F1: v6}
                                                    _ = v20
                                                }
                                            }
                                        default:
                                            var v21 uint = 4294967295
                                            _ = v21
                                        }
                                        v8 -= v8
                                        _ = v15
                                    }
                                    _ = fi13
                                }
                                _ = v12
                            } else {
                                v8 = v8
                                var v22 int = fi7
                                _ = v22
                            }
                        case "hello":
                            _ = len(v5)
                        default:
                            if false {
                                v6 -= (^(v6))
                                switch 0 {
                                case 32768:
                                    switch (fi7 + fi7) {
                                    case -135:
                                        v23 := GenS2{F0: 0, F1: 0}
                                        _ = v23
                                    case -675:
                                        fmt.Println(fi7, v8, v6)
                                    }
                                case -809:
                                    switch (fi7 + fi7) {
                                    case 0:
                                        var v24 string = "foo bar"
                                        _ = v24
                                    case 1:
                                        switch "foo bar" {
                                        case "\n\t":
                                            v6 = (v6 / ((v6) | 1))
                                        case "a":
                                            var v25 uint16 = 101
                                            _ = v25
                                        case "hello":
                                            var v26 int32 = -446
                                            _ = v26
                                        }
                                    }
                                case -9223372036854775808:
                                    v8 += (v8 | v8)
                                }
                                if false {
                                    if (0 <= 0) {
                                        var v27 uint64 = 748
                                        v28 := GenS1{F0: 136, F1: 255}
                                        _ = v27
                                        _ = v28
                                    } else {
                                        if (int16(v6) < 859) {
                                            v29 := GenS3{F0: 42}
                                            var v30 float32 = float32((1.5 * -2.25))
                                            var v31 int32 = -2147483648
                                            _ = v29
                                            _ = v30
                                            _ = v31
                                        } else {
                                            _ = len(v5)
                                            v5 = append(v5, -20)
                                            if false {
                                                v32 := GenS1{F0: 102, F1: 255}
                                                v6 = (v6 * v6)
                                                _ = v32
                                            } else {
                                                if len(v5) > 0 {
                                                    v5[0] = (^(fi7))
                                                }
                                            }
                                        }
                                        v8 -= v8
                                    }
                                    v8 -= (v8 / ((v8) | 1))
                                }
                            }
                        }
                    }
                } else {
                    var v33 uint16 = 490
                    var v34 int = fi7
                    _ = v33
                    _ = v34
                }
            }
        }
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 bool, p1 string, p2 int32) uint16 {
    for fi35 := 0; fi35 < 1; fi35++ {
        for fi36 := 0; fi36 < 4; fi36++ {
            v37 := []int{411, genFn0(), genFn0()}
            p2 -= 0
            var v38 int32 = (p2 >> ((4294967295) & 63))
            _ = fi36
            _ = v37
            _ = v38
        }
        var v39 uint16 = 2
        v40 := GenS4{F0: (!(p2 < (p2 - p2))), F1: 65535, F2: 12}
        _ = fi35
        _ = v39
        _ = v40
    }
    var v41 float64 = (1.5 / 1.0)
    p1 = "a"
    v42 := GenS4{F0: ((float64((float32(((v41 + v41) * ((float64(0.1) / 1.0) - v41))) * 1e10)) - (float64(3.14159) / float64((float32((float64((1.5 / 1e10)) + (v41 * 1e10))) + (float32(-2.25) / (0.0 - (1.5 - 1e10))))))) < ((v41 / v41) + v41)), F1: 42, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v41
    _ = v42
    return 42
}

func main() {
    v1 := []uint32{42, 2147483648}
    var v2 int32 = 2147483647
    var v3 float32 = 3.14159
    var v4 bool = (!((v2 << ((8) & 63)) == (^(v2))))
    for fi43 := 0; fi43 < 11; fi43++ {
        v44 := GenS5{F0: -32768, F1: 141}
        _ = fi43
        _ = v44
    }
    fmt.Println(v2, v4)
    for fi45 := 0; fi45 < 5; fi45++ {
        v4 = v4
        _ = fi45
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    for fi46 := 0; fi46 < 4; fi46++ {
        var v47 int16 = 1000
        for fi48 := 0; fi48 < 11; fi48++ {
            v47 = (v47 >> ((18446744073709551615) & 63))
            v49 := GenS6{F0: 255}
            _ = fi48
            _ = v49
        }
        fmt.Println(v3)
        _ = fi46
        _ = v47
    }
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    v3 = 1e10
    var v50 int8 = 100
    v3 = ((v3 * v3) + v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
}
