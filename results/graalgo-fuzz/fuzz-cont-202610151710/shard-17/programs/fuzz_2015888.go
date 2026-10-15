package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int32 }
type GenS2 struct { F0 int16; F1 int32; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int64; F1 int16; F2 int }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 int }

func main() {
    var v1 uint = 2
    var v2 uint32 = 852
    var v3 string = ""
    var v4 float32 = 3.14159
    v5 := []string{v3, v3, "foo bar"}
    for fi6 := 0; fi6 < 16; fi6++ {
        fmt.Println(v1, v3)
        _ = fi6
    }
    if (4294967295 != 1) {
        fmt.Println(v1, v1, v4)
    }
    switch "hello" {
    case "hello":
        for fi7 := 0; fi7 < 17; fi7++ {
            fmt.Println(v2, v4)
            _ = fi7
        }
    case "\n\t":
        if (444 == 4294967295) {
            if len(v5) > 0 {
                v5[0] = (v3 + v3)
            }
            if false {
                if (!true) {
                    v4 -= float32(float64(v4))
                    for fi8 := 0; fi8 < 11; fi8++ {
                        if (!(false || false)) {
                            var v9 uint16 = uint16(v1)
                            v3 = v3
                            switch fi8 {
                            case 321:
                                fmt.Println(fi8)
                            case 0:
                                v10 := GenS1{F0: (fi8 & 32768), F1: (v9 | (v9 / ((v9) | 1))), F2: -249}
                                _ = v10
                            default:
                                var v11 uint = v1
                                _ = v11
                            }
                            _ = v9
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            for fi12 := 0; fi12 < 18; fi12++ {
                                switch v3 {
                                case "\n\t":
                                    var v13 string = v3
                                    _ = v13
                                case "hello":
                                    if true {
                                        if ((!true) && false) {
                                            _ = len(v5)
                                        } else {
                                            _ = len(v5)
                                            var v14 float64 = ((((0.0 + 0.1) / (0.0 + 1e10)) * 3.14159) * 0.1)
                                            var v15 int = fi12
                                            _ = v14
                                            _ = v15
                                        }
                                    }
                                case "":
                                    if ("\n\t" > v3) {
                                        fmt.Println(v4, fi8)
                                    } else {
                                        var v16 uint16 = 0
                                        if false {
                                            var v17 int8 = int8(v16)
                                            if true {
                                                if false {
                                                    v2 += v2
                                                } else {
                                                    switch (fi12 * fi12) {
                                                    case 32767:
                                                        if (int64(v16) < -1) {
                                                            v18 := GenS1{F0: (-(fi12)), F1: 32767, F2: 646}
                                                            _ = v18
                                                        }
                                                    case -569:
                                                        var v19 int16 = -32768
                                                        _ = v19
                                                    case 9223372036854775807:
                                                        if true {
                                                            if false {
                                                                v17 -= v17
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                switch v3 {
                                                                case "\n\t":
                                                                    v20 := GenS2{F0: 42, F1: 2147483647, F2: v3}
                                                                    _ = len(v5)
                                                                    if (!(642 <= -539)) {
                                                                        v17 = v17
                                                                    } else {
                                                                        fmt.Println(v3, v4)
                                                                        v16 -= (^(v16))
                                                                    }
                                                                    var v21 uint8 = 2
                                                                    _ = v20
                                                                    _ = v21
                                                                case "hello":
                                                                    switch (fi12 & 527) {
                                                                    case 599:
                                                                        var v22 uint64 = 0
                                                                        _ = v22
                                                                    }
                                                                }
                                                                if true {
                                                                    fmt.Println(v3)
                                                                    if false {
                                                                        var v23 float32 = -2.25
                                                                        var v24 float32 = float32((((0.1 + -2.25) * (0.1 - 0.0)) * -2.25))
                                                                        v16 += v16
                                                                        _ = v23
                                                                        _ = v24
                                                                    } else {
                                                                        _ = len(v5)
                                                                        if false {
                                                                            var v25 int32 = -746
                                                                            switch (fi12 | (fi12 | (fi8 / (((fi8 + (-(fi12)))) | 1)))) {
                                                                            case 712:
                                                                                fmt.Println(v3, fi12, fi12)
                                                                            default:
                                                                                fmt.Println(v1)
                                                                            }
                                                                            _ = v25
                                                                        }
                                                                        v26 := GenS3{F0: 1}
                                                                        v4 += ((((v4 + float32((3.14159 - 1e10))) - float32(1e10)) - (((v4 + (float32(float64(v4)) + v4)) / (v4 / 3.14159)) + 1e10)) * v4)
                                                                        _ = v26
                                                                    }
                                                                }
                                                                v27 := GenS2{F0: 36, F1: 1, F2: ((((v3 + (v3 + v3)) + v3) + (v3 + v3)) + (v3 + ("a" + v3)))}
                                                                _ = v27
                                                            }
                                                            v1 = (v1 + v1)
                                                            v17 -= (v17 * (v17 | 1))
                                                        }
                                                    default:
                                                        if ((v17 * v17) != (v17 | (-(v17)))) {
                                                            if true {
                                                                switch (v3 + "\n\t") {
                                                                case "a":
                                                                    v28 := GenS1{F0: int(v17), F1: v16, F2: -1}
                                                                    _ = v28
                                                                case "hello":
                                                                    if (1000 >= -32768) {
                                                                        _ = len(v5)
                                                                    }
                                                                }
                                                                var v29 uint = v1
                                                                _ = v29
                                                            } else {
                                                                if (4294967295 >= v1) {
                                                                    var v30 float64 = (-2.25 - 1e10)
                                                                    v31 := GenS4{F0: 1, F1: v2, F2: 42}
                                                                    var v32 float32 = 1.5
                                                                    _ = v30
                                                                    _ = v31
                                                                    _ = v32
                                                                } else {
                                                                    var v33 bool = false
                                                                    v3 = (v3 + (("a" + (("\n\t" + v3) + "hello")) + "a"))
                                                                    v34 := GenS5{F0: 1, F1: -32768, F2: fi12}
                                                                    _ = v33
                                                                    _ = v34
                                                                }
                                                                v1 = 4294967295
                                                            }
                                                        } else {
                                                            var v35 int16 = 1
                                                            _ = v35
                                                        }
                                                    }
                                                    v36 := GenS5{F0: -712, F1: 42, F2: (fi8 + fi8)}
                                                    _ = v36
                                                }
                                            } else {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                fmt.Println(v4)
                                                fmt.Println(v16, v17)
                                            }
                                            var v37 uint16 = v16
                                            _ = v17
                                            _ = v37
                                        } else {
                                            v38 := GenS6{F0: 4294967295, F1: 65535}
                                            _ = v38
                                        }
                                        var v39 uint = (v1 / ((65535) | 1))
                                        _ = v16
                                        _ = v39
                                    }
                                }
                                var v40 uint8 = 42
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi12
                                _ = v40
                            }
                            for fi41 := 0; fi41 < 18; fi41++ {
                                if len(v5) > 0 {
                                    v5[0] = v3
                                }
                                var v42 uint32 = uint32(v1)
                                v3 = v3
                                var v43 float64 = 0.0
                                _ = len(v5)
                                _ = fi41
                                _ = v42
                                _ = v43
                            }
                        }
                        v1 -= v1
                        _ = fi8
                    }
                }
                var v44 bool = true
                fmt.Println(v3)
                _ = v44
            } else {
                v45 := GenS7{F0: -9223372036854775808}
                _ = v45
            }
        }
    case "":
        _ = len(v5)
    default:
        var v46 int16 = int16(v2)
        _ = v46
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v47 float32 = 1.5
    if (-1 <= -260) {
        var v48 int16 = -1
        _ = v48
    }
    v1 = (v1 / ((v1) | 1))
    v1 -= 1
    var v49 uint64 = 4294967295
    var v50 uint32 = uint32(v1)
    var v51 float64 = ((-2.25 * (0.0 * ((-2.25 * 1e10) - 3.14159))) - 3.14159)
    v2 += (v50 ^ 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
