package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint8; F2 bool }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 int64; F1 uint }
type GenS8 struct { F0 string; F1 uint16 }
type GenS9 struct { F0 uint16 }
func genFn0(p0 int32) uint64 {
    v7 := GenS1{F0: false, F1: 623, F2: true}
    switch "hello" {
    case "foo bar":
        fmt.Println(p0)
    }
    p0 += -2147483648
    _ = p0
    _ = v7
    return 424
}
func genFn1(p0 uint8) uint {
    var v8 int8 = 100
    p0 = p0
    v9 := GenS2{F0: 13}
    var v10 uint16 = 532
    var v11 uint64 = 831
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := []uint32{1}
    var v2 float64 = 3.14159
    var v3 bool = true
    var v4 uint32 = 2
    var v5 uint8 = 141
    var v6 string = "\n\t"
    switch "\n\t" {
    case "foo bar":
        v2 = 1e10
    case "hello":
        var v12 string = v6
        _ = v12
    default:
        var v13 float64 = 1.5
        _ = v13
    }
    v4 = v4
    if ((-658 == 128) && ((4294967295 < 663) || true)) {
        if len(v1) > 0 {
            v1[0] = (v4 << ((2) & 63))
        }
        for fi14 := 0; fi14 < 19; fi14++ {
            fmt.Println(v2, fi14)
            _ = fi14
        }
    } else {
        var v15 string = v6
        _ = v15
    }
    if (!v3) {
        for fi16 := 0; fi16 < 16; fi16++ {
            if (-2.25 >= ((((((((3.14159 - 1.5) * -2.25) - 1.5) / 1.0) + 3.14159) + ((((-2.25 / 3.14159) / 3.14159) + (0.1 / (0.0 + 3.14159))) + 1.5)) - (((((-2.25 + ((0.1 + 1.5) + 1.5)) * float32((1.5 / v2))) / ((-2.25 * 0.0) - float32(3.14159))) / 1.5) / (0.0 + 0.1))) - float32((v2 * 0.1)))) {
                switch "a" {
                case "hello":
                    for fi17 := 0; fi17 < 2; fi17++ {
                        v1 = append(v1, v4)
                        v2 = (-2.25 - v2)
                        var v18 int16 = 550
                        _ = fi17
                        _ = v18
                    }
                case "\n\t":
                    if (v3 && true) {
                        v3 = ((v4 - 1) > (^(v4)))
                    } else {
                        switch (fi16 * (fi16 | fi16)) {
                        case 32767:
                            _ = len(v1)
                        }
                    }
                }
                for fi19 := 0; fi19 < 17; fi19++ {
                    v20 := GenS3{F0: 515, F1: 2}
                    var v21 string = (v6 + "hello")
                    v4 = v4
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                var v22 float64 = float64(1e10)
                _ = v22
            } else {
                v23 := GenS3{F0: 32767, F1: (v5 | (v5 & (v5 * v5)))}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v24 := GenS3{F0: 842, F1: (v5 & v5)}
                _ = v23
                _ = v24
            }
            _ = fi16
        }
        switch v6 {
        case "foo bar":
            var v25 uint16 = 918
            _ = v25
        case "hello":
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := GenS4{F0: 42, F1: v5, F2: ((v5 * 0) >= (^(v5)))}
                v1 = append(v1, v4)
                if v3 {
                    fmt.Println(v3)
                    for fi28 := 0; fi28 < 15; fi28++ {
                        switch "\n\t" {
                        case "":
                            var v29 int64 = 42
                            _ = v29
                        case "foo bar":
                            if (v5 > (v5 / ((v5) | 1))) {
                                v3 = ((1 <= 42) && v3)
                            } else {
                                _ = len(v1)
                            }
                        case "\n\t":
                            switch (fi28 * (-(fi28))) {
                            case -2147483648:
                                v30 := GenS5{F0: 42, F1: 1}
                                _ = v30
                            default:
                                switch 452 {
                                case 2147483647:
                                    v1 = append(v1, v4)
                                }
                            }
                        }
                        _ = fi28
                    }
                }
                _ = fi26
                _ = v27
            }
        case "a":
            v1 = append(v1, 559)
        }
        v2 = v2
    } else {
        if (v3 && ((v6 >= v6) || v3)) {
            var v31 int8 = -1
            _ = v31
        } else {
            fmt.Println(v4, v6)
        }
        v5 = v5
        var v32 uint64 = 18446744073709551615
        _ = v32
    }
    if false {
        v33 := GenS3{F0: 707, F1: v5}
        var v34 int8 = 42
        _ = v33
        _ = v34
    } else {
        v4 += 2147483647
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 bool = v3
            _ = fi35
            _ = v36
        }
        fmt.Println(v4)
    }
    v5 -= 42
    var v37 int64 = -642
    for fi38 := 0; fi38 < 15; fi38++ {
        for fi39 := 0; fi39 < 12; fi39++ {
            if v3 {
                fmt.Println(v37)
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if v3 {
                    switch int(v5) {
                    case -9223372036854775808:
                        v40 := GenS1{F0: ((v4 - (^(v4))) >= v4), F1: uint16(v4), F2: ((1 == 493) || (v5 == 2))}
                        _ = v40
                    case 9223372036854775807:
                        v41 := GenS6{F0: 582, F1: 0}
                        _ = v41
                    case 32767:
                        v6 = v6
                    }
                    if len(v1) > 0 {
                        v1[0] = v4
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v6 = v6
                    var v42 uint32 = uint32(v5)
                    v43 := GenS7{F0: (v37 % (((^(v37))) | 1)), F1: 42}
                    v44 := GenS4{F0: 1, F1: v5, F2: v3}
                    if (!((v37 / ((42) | 1)) <= 1)) {
                        v2 = v2
                    } else {
                        switch (v6 + ("foo bar" + "")) {
                        case "":
                            v3 = v3
                        }
                        switch (fi38 + (fi39 << ((genFn1((v5 | (v5 % (((v5 << ((4294967295) & 63))) | 1))))) & 63))) {
                        case -9223372036854775808:
                            v45 := GenS6{F0: uint16(v4), F1: 42}
                            _ = v45
                        }
                        v4 = v4
                        fmt.Println(v6, fi38)
                        if v3 {
                            _ = len(v1)
                        }
                    }
                    v46 := GenS2{F0: 966}
                    var v47 float64 = (0.0 - -2.25)
                    _ = v42
                    _ = v43
                    _ = v44
                    _ = v46
                    _ = v47
                } else {
                    var v48 uint8 = (v5 << ((756) & 63))
                    _ = v48
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v6 = v6
            _ = fi39
        }
        var v49 float32 = 3.14159
        _ = fi38
        _ = v49
    }
    v50 := GenS3{F0: 255, F1: v5}
    switch -492 {
    case 0:
        v6 = (v6 + (v6 + (v6 + v6)))
    default:
        switch ("\n\t" + v6) {
        case "":
            for fi51 := 0; fi51 < 18; fi51++ {
                v52 := GenS8{F0: v6, F1: 255}
                if (127 <= v5) {
                    v53 := GenS2{F0: 2147483647}
                    _ = v53
                }
                for fi54 := 0; fi54 < 7; fi54++ {
                    fmt.Println(v4)
                    v3 = (!v3)
                    fmt.Println(fi54, v4, v4)
                    _ = fi54
                }
                _ = fi51
                _ = v52
            }
        case "\n\t":
            for fi55 := 0; fi55 < 17; fi55++ {
                if (v3 && (float64(0.1) > (v2 / (v2 - 1.5)))) {
                    v56 := GenS7{F0: (v37 << ((42) & 63)), F1: 2}
                    _ = v56
                }
                _ = fi55
            }
        case "hello":
            var v57 int16 = 969
            _ = v57
        }
    }
    fmt.Println(v2, v6, v4)
    v1 = append(v1, 42)
    if len(v1) > 0 {
        v1[0] = (v4 | v4)
    }
    for fi58 := 0; fi58 < 3; fi58++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v6 = ("foo bar" + (("hello" + ("" + v6)) + v6))
        _ = fi58
    }
    v1 = append(v1, (^(v4)))
    v59 := GenS9{F0: 2}
    for fi60 := 0; fi60 < 15; fi60++ {
        v6 = v6
        _ = fi60
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v50
    _ = v59
    _ = v6
}
