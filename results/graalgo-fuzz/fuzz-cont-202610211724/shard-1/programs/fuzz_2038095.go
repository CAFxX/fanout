package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int32 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint8; F1 string; F2 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 bool; F2 int8 }
type GenS7 struct { F0 uint16; F1 bool; F2 bool }

func main() {
    var v1 int = 0
    var v2 float32 = -2.25
    var v3 float64 = ((-2.25 + float64(3.14159)) * (0.0 / (float64(((((((3.14159 + float32(((0.0 - -2.25) - ((0.1 + -2.25) - 1.5)))) + v2) / (float32(1.5) / v2)) + (0.1 + v2)) / (((((1e10 * v2) / 1.5) * v2) - v2) - (float32((1e10 / (float64((((v2 + v2) * v2) + v2)) / -2.25))) + v2))) / 3.14159)) + -2.25)))
    var v4 int16 = 1
    v5 := []string{"", ("\n\t" + ("\n\t" + "hello")), "foo bar"}
    var v6 int = v1
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v7 float64 = (1e10 * (v3 / v3))
        fmt.Println(v3)
        _ = v7
    }
    var v8 string = "a"
    v4 = (v4 + 32767)
    fmt.Println(v1)
    if (755 > 2147483648) {
        var v9 string = "hello"
        if false {
            v5 = append(v5, v9)
            v8 = v9
            v10 := GenS1{F0: v4, F1: (v1 * (v1 & 2147483647)), F2: 2147483647}
            _ = v10
        }
        _ = v9
    } else {
        var v11 uint = 18446744073709551615
        fmt.Println(v6)
        _ = v11
    }
    if true {
        v12 := GenS1{F0: v4, F1: (v1 ^ v1), F2: 2147483647}
        var v13 bool = (!true)
        _ = v12
        _ = v13
    }
    v5 = append(v5, ("\n\t" + v8))
    for fi14 := 0; fi14 < 11; fi14++ {
        switch v8 {
        case "a":
            _ = len(v5)
        case "foo bar":
            switch v8 {
            case "\n\t":
                var v15 float64 = 0.0
                _ = v15
            case "a":
                if true {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        v1 -= v6
                        fmt.Println(v1)
                        _ = fi16
                    }
                    v2 = (v2 + float32(float64((((v2 + 0.1) * v2) - float32(0.0)))))
                }
            default:
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
            }
        }
        var v17 uint16 = 32767
        _ = fi14
        _ = v17
    }
    switch v6 {
    case 32767:
        var v18 uint = 65535
        _ = v18
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v1 -= v1
    if (!(-128 < 43)) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        if ((42 == 62) && true) {
            var v19 int32 = 2147483647
            var v20 int = (v1 - v6)
            var v21 int16 = -32768
            _ = v19
            _ = v20
            _ = v21
        } else {
            v1 = -1
            if true {
                v22 := GenS1{F0: (v4 - v4), F1: (v6 & 42), F2: -2147483648}
                var v23 int16 = v4
                for fi24 := 0; fi24 < 10; fi24++ {
                    fmt.Println(v23)
                    if (v2 < (1.5 + 0.0)) {
                        v25 := GenS2{F0: (v23 * (v4 ^ v4)), F1: 358}
                        var v26 int32 = 1
                        if false {
                            _ = len(v5)
                            switch ((v8 + (v8 + "")) + "foo bar") {
                            case "\n\t":
                                v27 := GenS1{F0: (v23 >> ((18446744073709551615) & 63)), F1: v6, F2: v26}
                                _ = v27
                            }
                        }
                        _ = v25
                        _ = v26
                    }
                    for fi28 := 0; fi28 < 4; fi28++ {
                        switch v8 {
                        case "a":
                            var v29 float32 = ((0.1 * float32(((v3 + ((3.14159 + (float64((1e10 - (v2 + (float32(0.1) / 0.1)))) - (v3 - (v3 + v3)))) * v3)) / (0.1 / -2.25)))) + float32(v3))
                            _ = v29
                        default:
                            var v30 float32 = v2
                            _ = v30
                        }
                        v1 = (v1 >> ((42) & 63))
                        _ = fi28
                    }
                    _ = fi24
                }
                _ = v22
                _ = v23
            } else {
                var v31 bool = true
                for fi32 := 0; fi32 < 20; fi32++ {
                    switch 42 {
                    case 32768:
                        var v33 uint64 = 972
                        _ = v33
                    }
                    _ = fi32
                }
                var v34 int32 = -109
                _ = v31
                _ = v34
            }
        }
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        var v36 uint32 = 886
        _ = fi35
        _ = v36
    }
    fmt.Println(v8, v6, v2)
    for fi37 := 0; fi37 < 3; fi37++ {
        fmt.Println(v6, v8)
        v8 = v8
        if false {
            v6 = (fi37 % ((v6) | 1))
            if false {
                if false {
                    var v38 int8 = 62
                    v39 := GenS3{F0: 128, F1: 232}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v38
                    _ = v39
                } else {
                    var v40 string = (v8 + (v8 + "a"))
                    v41 := GenS4{F0: 128, F1: v8, F2: v4}
                    var v42 bool = ((255 < 196) || (!(!(!(true && false)))))
                    _ = v40
                    _ = v41
                    _ = v42
                }
                var v43 int = v1
                _ = v43
            }
            for fi44 := 0; fi44 < 17; fi44++ {
                switch fi37 {
                case -1:
                    fmt.Println(v1, v6, fi44)
                case -192:
                    if (2 != 738) {
                        fmt.Println(v1, v6)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        if false {
                            var v45 uint64 = 0
                            v6 = fi37
                            _ = len(v5)
                            _ = v45
                        }
                        v8 = v8
                    }
                    v5 = append(v5, (("\n\t" + v8) + ""))
                    v6 = 42
                    v3 += (1e10 / v3)
                    v46 := GenS5{F0: 2147483647}
                    var v47 int32 = 42
                    _ = v46
                    _ = v47
                case 32767:
                    switch int(v4) {
                    case 42:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case 1:
                        if len(v5) > 0 {
                            v5[0] = (v8 + "hello")
                        }
                    }
                }
                _ = fi44
            }
        } else {
            v48 := GenS2{F0: 0, F1: 65535}
            for fi49 := 0; fi49 < 7; fi49++ {
                v3 = ((v3 * float64((v2 * v2))) / (-2.25 * v3))
                v8 = (v8 + v8)
                if (true && ((!false) || (true || (true && true)))) {
                    v50 := GenS6{F0: true, F1: (0 < 18446744073709551615), F2: 1}
                    var v51 int = fi49
                    _ = v50
                    _ = v51
                }
                v6 -= (^(fi49))
                _ = fi49
            }
            if (-128 == -1) {
                v4 -= v4
                fmt.Println(v4)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                if (230 <= 1) {
                    if true {
                        fmt.Println(v8, v2, v3)
                        if (true || false) {
                            v2 = ((-2.25 / (v2 * v2)) - v2)
                            v3 = v3
                        }
                        v52 := GenS7{F0: 7, F1: (false || ((((1 < -1) || false) || (-243 > 42)) || true)), F2: (-206 > 1)}
                        _ = v52
                    }
                    var v53 int16 = v4
                    _ = v53
                }
                var v54 uint16 = 146
                _ = v54
            }
            _ = v48
        }
        _ = fi37
    }
    if len(v5) > 0 {
        v5[0] = v8
    }
    var v55 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v8
}
