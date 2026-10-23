package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint8 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 int64) uint8 {
    v7 := GenS1{F0: 4294967295, F1: 42, F2: 128}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 1
}
func genFn1() int8 {
    var v8 float64 = -2.25
    var v9 float64 = -2.25
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(fi10, v9, fi10)
        if (0 == genFn0(-1)) {
            v8 += float64(1.5)
        }
        _ = fi10
    }
    var v11 string = ("a" + "")
    v9 = ((1e10 * 0.0) - (v8 * v9))
    _ = v11
    _ = v8
    _ = v9
    return 60
}

func main() {
    var v1 uint32 = 0
    var v2 float64 = (3.14159 * 3.14159)
    var v3 float64 = 1e10
    var v4 string = ((("a" + "") + ("a" + "")) + (("hello" + "hello") + "a"))
    v5 := []uint32{2147483648, 4294967295, 42}
    var v6 int16 = 1
    v4 = (("a" + ("hello" + "a")) + v4)
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            if false {
                if true {
                    switch (v4 + ("\n\t" + v4)) {
                    case "hello":
                        var v14 int8 = -69
                        _ = v14
                    default:
                        fmt.Println(v4, v1)
                    }
                    if (-495 > -1) {
                        var v15 uint64 = uint64(v1)
                        var v16 uint32 = 517
                        v2 = (0.1 / v2)
                        _ = v15
                        _ = v16
                    }
                } else {
                    var v17 float32 = float32(float64(1.5))
                    var v18 int16 = -179
                    _ = v17
                    _ = v18
                }
                if true {
                    v2 = v2
                }
                _ = len(v5)
            } else {
                v3 = (float64(3.14159) - 1.5)
                v1 += 2147483648
                if (1 > 42) {
                    if (42 < 478) {
                        switch "a" {
                        case "foo bar":
                            _ = len(v5)
                        case "\n\t":
                            switch (v4 + (v4 + v4)) {
                            case "":
                                var v19 uint16 = 255
                                _ = v19
                            case "\n\t":
                                v20 := GenS1{F0: 0, F1: 254, F2: uint8(v1)}
                                _ = v20
                            }
                        case "a":
                            v2 = 1e10
                        }
                        _ = len(v5)
                        v3 -= v3
                    }
                } else {
                    v1 += (^(v1))
                }
            }
            if false {
                v3 = float64(((((1.5 * (-2.25 / 1e10)) + ((1e10 / (3.14159 + 3.14159)) / (((1e10 * 0.1) + 1e10) + 1e10))) - (0.0 * (float32(1e10) - 0.1))) / 1.0))
            } else {
                var v21 int64 = 42
                _ = v21
            }
            _ = fi13
        }
        _ = fi12
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        if (true && ((!(!((18446744073709551615 <= 192) || false))) && false)) {
            if true {
                switch "" {
                case "foo bar":
                    var v23 uint = 529
                    _ = v23
                case "\n\t":
                    v5 = append(v5, v1)
                }
                var v24 string = "foo bar"
                _ = v24
            } else {
                v25 := GenS1{F0: 2, F1: 812, F2: 42}
                _ = v25
            }
            var v26 float32 = -2.25
            for fi27 := 0; fi27 < 6; fi27++ {
                if ((0 <= 2) && true) {
                    v28 := GenS1{F0: 327, F1: 1, F2: 127}
                    if ("" >= (("hello" + v4) + (((("a" + v4) + ("\n\t" + v4)) + v4) + v4))) {
                        v29 := GenS1{F0: 0, F1: -1, F2: 2}
                        fmt.Println(v26, v1, fi27)
                        _ = v29
                    }
                    var v30 float32 = (0.1 - 0.1)
                    _ = v28
                    _ = v30
                } else {
                    var v31 float64 = v2
                    _ = v31
                }
                switch ("\n\t" + "a") {
                case "":
                    if false {
                        if len(v5) > 0 {
                            v5[0] = 1
                        }
                        v26 = v26
                        switch v4 {
                        case "\n\t":
                            switch ("" + "a") {
                            case "foo bar":
                                var v32 int = (fi27 | (fi22 * (fi22 - 127)))
                                v3 = (v3 / ((v3 * 1.5) * v3))
                                v33 := GenS1{F0: 65535, F1: 42, F2: 42}
                                _ = v32
                                _ = v33
                            case "":
                                switch "\n\t" {
                                case "hello":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            }
                        case "a":
                            fmt.Println(fi27)
                        case "":
                            var v34 uint = 18446744073709551615
                            _ = v34
                        }
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v35 int64 = 0
                        _ = v35
                    }
                case "a":
                    _ = len(v5)
                case "\n\t":
                    var v36 int32 = 2147483647
                    _ = v36
                default:
                    v6 += (^(v6))
                }
                v2 -= ((1.5 / (((v2 - v2) + (((v2 / (v3 / v2)) + v2) + ((float64(v26) * v3) - v2))) + (1.5 - float64(float32(0.0))))) * ((-2.25 - 1.5) - float64(0.0)))
                _ = fi27
            }
            _ = v26
        }
        fmt.Println(v1)
        v2 = v3
        _ = fi22
    }
    fmt.Println(v3, v4, v2)
    switch 4294967295 {
    case -14:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if false {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        var v38 uint32 = (v1 * 177)
        v2 = (v3 / v3)
        if (216 < -1) {
            _ = len(v5)
            if false {
                if (42 <= 436) {
                    var v39 uint = 4294967295
                    v5 = append(v5, v1)
                    _ = v39
                } else {
                    if len(v5) > 0 {
                        v5[0] = (v38 + 4294967295)
                    }
                    var v40 int8 = -32
                    fmt.Println(v6, v38, v4)
                    _ = v40
                }
                for fi41 := 0; fi41 < 1; fi41++ {
                    if false {
                        v42 := GenS2{F0: ((2 <= 255) && true), F1: ((v3 + (((1e10 + -2.25) + float64(float32(v3))) / v2)) < float64(((1e10 - float32((v3 * (v2 * float64(float32(v2)))))) + 3.14159)))}
                        if (42 != 830) {
                            v1 = (v1 ^ (v38 - (^(v38))))
                        }
                        v1 = 5
                        _ = v42
                    } else {
                        if true {
                            switch fi37 {
                            case 32768:
                                _ = len(v5)
                            case 42:
                                if len(v5) > 0 {
                                    v5[0] = (v1 % ((2) | 1))
                                }
                            case 127:
                                v43 := GenS1{F0: 586, F1: -444, F2: 0}
                                v1 = v38
                                _ = v43
                            default:
                                switch fi37 {
                                case -638:
                                    v6 -= v6
                                }
                            }
                        }
                    }
                    _ = fi41
                }
                v6 -= (v6 >> ((663) & 63))
            } else {
                if true {
                    var v44 uint32 = v1
                    var v45 uint32 = (v1 * v1)
                    _ = v44
                    _ = v45
                } else {
                    v3 = v2
                    fmt.Println(v6)
                    v2 = v2
                }
            }
        } else {
            v46 := GenS1{F0: 18446744073709551615, F1: 0, F2: 92}
            fmt.Println(v3, fi37, v38)
            _ = v46
        }
        _ = fi37
        _ = v38
    }
    var v47 int8 = -128
    var v48 float64 = ((3.14159 - (((v3 - v3) + 1e10) / v2)) + 1.5)
    for fi49 := 0; fi49 < 8; fi49++ {
        for fi50 := 0; fi50 < 5; fi50++ {
            v51 := GenS1{F0: 4294967295, F1: 42, F2: genFn0(-1)}
            fmt.Println(v3)
            if false {
                if false {
                    v52 := GenS3{F0: uint(v1), F1: (-(v6))}
                    _ = v52
                }
            } else {
                v47 = 100
            }
            _ = fi50
            _ = v51
        }
        switch (fi49 + (fi49 | (fi49 - (fi49 % (((^(fi49))) | 1))))) {
        case 127:
            v5 = append(v5, 116)
        default:
            if (v4 < v4) {
                fmt.Println(v3, v2, v48)
                var v53 string = v4
                _ = v53
            }
        }
        v54 := GenS2{F0: (v6 <= (v6 % (((v6 - 32767)) | 1))), F1: true}
        _ = fi49
        _ = v54
    }
    v5 = append(v5, (v1 * (^(v1))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
