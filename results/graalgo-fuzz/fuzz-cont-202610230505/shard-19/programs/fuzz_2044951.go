package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int64; F2 int64 }
type GenS4 struct { F0 bool; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 uint }
func genFn0(p0 string) uint16 {
    fmt.Println(p0)
    v4 := []int{}
    _ = p0
    _ = v4
    return 32767
}

func main() {
    var v1 string = "foo bar"
    var v2 int64 = 1
    var v3 float32 = (1.5 * (0.1 / -2.25))
    v2 += (v2 % ((v2) | 1))
    if true {
        if ((0 > -1) && false) {
            v3 += 0.0
        } else {
            switch -1 {
            case -430:
                v1 = v1
            case 127:
                var v5 uint8 = 127
                _ = v5
            case 976:
                var v6 uint = 255
                _ = v6
            }
            for fi7 := 0; fi7 < 13; fi7++ {
                v2 = (v2 ^ v2)
                switch ((("foo bar" + "a") + "\n\t") + "hello") {
                case "\n\t":
                    fmt.Println(v1, v2, v1)
                }
                _ = fi7
            }
            switch v1 {
            case "foo bar":
                v1 = v1
            case "\n\t":
                switch ((("hello" + ("\n\t" + "")) + v1) + "foo bar") {
                case "a":
                    v8 := GenS1{F0: 32767}
                    _ = v8
                default:
                    v9 := GenS2{F0: 42}
                    _ = v9
                }
            default:
                v3 += v3
            }
        }
    } else {
        switch "" {
        case "":
            fmt.Println(v2, v1, v2)
        case "hello":
            v10 := []int{}
            _ = v10
        case "\n\t":
            v11 := []int{4294967295, 1}
            _ = v11
        }
        var v12 int16 = -172
        v2 = (v2 | (v2 * 0))
        _ = v12
    }
    var v13 int16 = -1
    if (!false) {
        if (193 > 49) {
            var v14 float32 = (float32(0.1) * v3)
            v15 := GenS3{F0: 132, F1: (v2 / (((^(v2))) | 1)), F2: v2}
            var v16 uint64 = 4294967295
            _ = v14
            _ = v15
            _ = v16
        }
        v2 -= v2
    } else {
        v3 += v3
        var v17 uint32 = 4294967295
        fmt.Println(v1, v1, v17)
        _ = v17
    }
    fmt.Println(v2, v1)
    if false {
        if (v2 != -9223372036854775808) {
            v3 = (((((v3 * ((float32(float64((float32(-2.25) * v3))) / 3.14159) - (1.5 + v3))) - 0.0) - (v3 + 0.1)) - v3) * v3)
        } else {
            v3 = v3
        }
        v3 += v3
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 string = ((v1 + ("foo bar" + ("hello" + v1))) + v1)
            var v20 int16 = v13
            _ = fi18
            _ = v19
            _ = v20
        }
    } else {
        switch v1 {
        case "\n\t":
            var v21 float32 = -2.25
            _ = v21
        default:
            fmt.Println(v3)
        }
    }
    switch 128 {
    case 127:
        switch 128 {
        case 4294967295:
            v2 += v2
        case -959:
            fmt.Println(v1)
        default:
            if true {
                var v22 uint16 = 255
                _ = v22
            } else {
                v23 := GenS4{F0: (4294967295 == 65535), F1: 0, F2: 68}
                fmt.Println(v1, v13)
                if false {
                    v24 := []int{-1, 933, -2147483648}
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 5; fi25++ {
                        v26 := GenS5{F0: genFn0("foo bar"), F1: 2147483648, F2: 100}
                        v27 := []int{(-(fi25)), 9223372036854775807, (fi25 ^ fi25)}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
                _ = v23
            }
        }
    case -1:
        switch 42 {
        case 32767:
            v2 = v2
        default:
            var v28 int = 2147483647
            _ = v28
        }
    case 2147483647:
        v29 := GenS1{F0: 2}
        _ = v29
    default:
        switch v1 {
        case "\n\t":
            for fi30 := 0; fi30 < 16; fi30++ {
                v31 := GenS6{F0: 32767, F1: 42, F2: genFn0(v1)}
                if false {
                    v13 += v13
                } else {
                    var v32 uint32 = 1
                    v13 = v13
                    _ = v32
                }
                switch 455 {
                case 9223372036854775807:
                    v3 += 3.14159
                case 32768:
                    var v33 bool = true
                    _ = v33
                case -1:
                    if false {
                        fmt.Println(v13, v13, v13)
                    } else {
                        switch (fi30 * fi30) {
                        case 42:
                            v34 := []int{42, -883, -1}
                            _ = v34
                        case 128:
                            v35 := GenS7{F0: 693, F1: 128}
                            _ = v35
                        default:
                            v36 := []int{}
                            _ = v36
                        }
                    }
                default:
                    v3 = ((v3 / v3) / (v3 * (0.0 - 0.0)))
                }
                _ = fi30
                _ = v31
            }
        case "":
            var v37 bool = false
            _ = v37
        }
    }
    v38 := GenS8{F0: 42}
    var v39 uint32 = 2147483648
    if false {
        for fi40 := 0; fi40 < 18; fi40++ {
            var v41 uint16 = 65535
            v42 := GenS5{F0: genFn0(""), F1: (v39 & (v39 | (^(v39)))), F2: 0}
            _ = fi40
            _ = v41
            _ = v42
        }
    } else {
        if (v1 >= v1) {
            v43 := []int{-5, 127}
            fmt.Println(v3, v13)
            var v44 uint8 = 238
            _ = v43
            _ = v44
        } else {
            v1 = (v1 + (v1 + v1))
        }
        for fi45 := 0; fi45 < 12; fi45++ {
            v46 := GenS6{F0: 997, F1: (v39 << ((386) & 63)), F2: 159}
            v13 = v13
            v47 := []int{(^(fi45)), int(v39)}
            _ = fi45
            _ = v46
            _ = v47
        }
    }
    var v48 uint = 905
    v48 += 858
    v49 := GenS3{F0: 1, F1: -137, F2: v2}
    var v50 bool = true
    var v51 bool = (0.1 < 0.0)
    if (42 <= 4294967295) {
        v48 = (^(v48))
        if v50 {
            if (("foo bar" + (v1 + "a")) != v1) {
                v3 -= v3
            }
            if v51 {
                v2 = v2
                for fi52 := 0; fi52 < 7; fi52++ {
                    if true {
                        switch int(v39) {
                        case 4294967295:
                            switch (fi52 | fi52) {
                            case -54:
                                var v53 uint8 = 0
                                _ = v53
                            case 926:
                                var v54 int = fi52
                                _ = v54
                            }
                        default:
                            fmt.Println(v39)
                        }
                        v51 = v51
                        var v55 bool = v50
                        _ = v55
                    } else {
                        var v56 string = (v1 + "hello")
                        _ = v56
                    }
                    var v57 uint64 = 18446744073709551615
                    v50 = (1 <= (v39 + uint32(v57)))
                    _ = fi52
                    _ = v57
                }
                for fi58 := 0; fi58 < 20; fi58++ {
                    var v59 int = (fi58 + (fi58 * (fi58 | (fi58 % ((fi58) | 1)))))
                    v60 := GenS8{F0: (v48 ^ (v48 ^ v48))}
                    _ = fi58
                    _ = v59
                    _ = v60
                }
            }
            for fi61 := 0; fi61 < 4; fi61++ {
                v62 := []int{fi61, (fi61 * -2147483648), (fi61 ^ fi61)}
                _ = fi61
                _ = v62
            }
        } else {
            fmt.Println(v3, v50)
            if true {
                switch "hello" {
                case "a":
                    if true {
                        v51 = (v1 != (v1 + v1))
                    }
                case "":
                    for fi63 := 0; fi63 < 6; fi63++ {
                        var v64 float32 = v3
                        v51 = v50
                        _ = fi63
                        _ = v64
                    }
                }
            } else {
                v48 = (v48 % ((v48) | 1))
                if (!(!(-128 == -28))) {
                    v3 = float32(1.5)
                    v65 := []int{}
                    for fi66 := 0; fi66 < 3; fi66++ {
                        _ = len(v65)
                        _ = fi66
                    }
                    _ = v65
                } else {
                    v2 -= (v2 * v2)
                    for fi67 := 0; fi67 < 10; fi67++ {
                        v1 = v1
                        _ = fi67
                    }
                }
            }
        }
        v68 := []int{-412, 42}
        _ = v68
    }
    for fi69 := 0; fi69 < 7; fi69++ {
        var v70 float64 = float64((((v3 * v3) / v3) + 1e10))
        _ = fi69
        _ = v70
    }
    v51 = v50
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v48
    _ = v49
    _ = v50
    _ = v51
}
