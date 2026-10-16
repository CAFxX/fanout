package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 int64; F2 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 uint; F2 int32 }

func main() {
    var v1 bool = true
    v2 := []string{}
    var v3 bool = (-125 == -2147483648)
    var v4 int = 32768
    switch v4 {
    case 0:
        for fi5 := 0; fi5 < 10; fi5++ {
            switch "a" {
            case "\n\t":
                var v6 float64 = ((1e10 - float64(0.1)) * 0.1)
                _ = v6
            case "hello":
                if v1 {
                    if ((v4 - v4) == 9223372036854775807) {
                        for fi7 := 0; fi7 < 1; fi7++ {
                            var v8 float64 = (1e10 - 3.14159)
                            var v9 float64 = 3.14159
                            _ = fi7
                            _ = v8
                            _ = v9
                        }
                        if (!(1.5 >= 3.14159)) {
                            if (257 <= 0) {
                                fmt.Println(v3)
                            }
                        } else {
                            switch "\n\t" {
                            case "a":
                                v10 := GenS1{F0: 2, F1: 42}
                                _ = v10
                            case "foo bar":
                                var v11 float64 = (((float64((1.5 / -2.25)) * 0.0) / (-2.25 * 1e10)) / 1.5)
                                _ = v11
                            }
                            for fi12 := 0; fi12 < 6; fi12++ {
                                v13 := GenS2{F0: 255, F1: 131}
                                if (-2147483648 == -2147483648) {
                                    v3 = (v3 || (42 != -100))
                                    v14 := GenS3{F0: 2}
                                    if (true && (v4 != fi5)) {
                                        var v15 int64 = 936
                                        v1 = v3
                                        if v3 {
                                            fmt.Println(v1)
                                        }
                                        _ = v15
                                    } else {
                                        var v16 uint64 = 42
                                        switch "" {
                                        case "\n\t":
                                            v17 := GenS1{F0: 2, F1: v16}
                                            _ = v17
                                        case "":
                                            fmt.Println(fi5)
                                        default:
                                            var v18 uint16 = 1
                                            _ = v18
                                        }
                                        _ = v16
                                    }
                                    _ = v14
                                } else {
                                    if (!true) {
                                        fmt.Println(v1, v4, v4)
                                    }
                                    if (!(127 <= 100)) {
                                        v3 = (0 < 18446744073709551615)
                                        switch -966 {
                                        case -324:
                                            var v19 int16 = -785
                                            _ = v19
                                        case 217:
                                            if (v1 && v3) {
                                                var v20 uint16 = 42
                                                _ = v20
                                            } else {
                                                var v21 uint = 873
                                                if (-591 == 0) {
                                                    var v22 int64 = 9223372036854775807
                                                    _ = v22
                                                }
                                                var v23 int16 = -1
                                                var v24 int32 = 2147483647
                                                _ = v21
                                                _ = v23
                                                _ = v24
                                            }
                                        }
                                        v4 = v4
                                        var v25 uint = 74
                                        v25 = 1
                                        v2 = append(v2, "")
                                        if false {
                                            v26 := GenS3{F0: 65535}
                                            if (32767 > 42) {
                                                if v3 {
                                                    var v27 int8 = 42
                                                    v4 = fi12
                                                    _ = v27
                                                }
                                            }
                                            _ = v26
                                        }
                                        _ = v25
                                    } else {
                                        v4 -= (fi5 * fi12)
                                        switch 476 {
                                        case -140:
                                            switch "\n\t" {
                                            case "hello":
                                                switch v4 {
                                                case 127:
                                                    v28 := GenS2{F0: 645, F1: 65535}
                                                    _ = v28
                                                case -887:
                                                    var v29 uint8 = 255
                                                    _ = v29
                                                default:
                                                    fmt.Println(v4)
                                                }
                                            default:
                                                var v30 int64 = -1
                                                _ = v30
                                            }
                                        default:
                                            var v31 uint = 1
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            switch -2147483648 {
                                            case 9223372036854775807:
                                                v4 += int(v31)
                                            case 32767:
                                                _ = len(v2)
                                            }
                                            _ = v31
                                        }
                                    }
                                    var v32 uint8 = 255
                                    v33 := GenS4{F0: 0, F1: 941, F2: 195}
                                    _ = v32
                                    _ = v33
                                }
                                _ = fi12
                                _ = v13
                            }
                            v34 := GenS3{F0: 2}
                            _ = v34
                        }
                    }
                } else {
                    var v35 uint = 147
                    v36 := GenS5{F0: 3, F1: -32768}
                    _ = v35
                    _ = v36
                }
            }
            var v37 int16 = int16(fi5)
            _ = fi5
            _ = v37
        }
    case 4294967295:
        var v38 int = 128
        _ = v38
    }
    v39 := GenS6{F0: 1}
    v40 := GenS7{F0: v1}
    if ((-2.25 <= (((1.5 - ((-2.25 * (1e10 + 0.1)) - float64((1e10 / float32(((float64((float32(float64(1e10)) + 0.0)) - -2.25) + 1.5)))))) - -2.25) * 0.1)) && true) {
        for fi41 := 0; fi41 < 10; fi41++ {
            if (!true) {
                v42 := GenS5{F0: 45, F1: 1}
                var v43 uint16 = 65535
                _ = v42
                _ = v43
            } else {
                var v44 int8 = -1
                switch v4 {
                case 0:
                    v45 := GenS7{F0: ((3.14159 >= 0.0) && true)}
                    _ = v45
                default:
                    for fi46 := 0; fi46 < 18; fi46++ {
                        var v47 int = 32768
                        v4 = v47
                        switch fi46 {
                        case 32768:
                            fmt.Println(v4, v3)
                        case 127:
                            switch "foo bar" {
                            case "\n\t":
                                var v48 uint64 = 219
                                _ = v48
                            case "a":
                                v49 := GenS8{F0: 42, F1: 386, F2: -128}
                                _ = v49
                            default:
                                v47 -= v47
                            }
                        default:
                            fmt.Println(v47, v3)
                        }
                        if (false || v1) {
                            fmt.Println(v47, fi46, v4)
                            var v50 uint16 = 65535
                            _ = v50
                        } else {
                            var v51 uint8 = 1
                            _ = v51
                        }
                        _ = fi46
                        _ = v47
                    }
                }
                _ = v44
            }
            _ = fi41
        }
        for fi52 := 0; fi52 < 1; fi52++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v53 bool = v3
            _ = fi52
            _ = v53
        }
        var v54 int = (v4 << ((42) & 63))
        _ = v54
    }
    v1 = ((!v1) && (v4 != v4))
    if (23 >= 0) {
        var v55 string = ("a" + "foo bar")
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v55
    }
    var v56 bool = false
    for fi57 := 0; fi57 < 15; fi57++ {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        _ = fi57
    }
    switch 4294967295 {
    case -2147483648:
        if v1 {
            v2 = append(v2, "hello")
            switch (v4 >> ((656) & 63)) {
            case 1:
                for fi58 := 0; fi58 < 3; fi58++ {
                    v4 = (fi58 / ((fi58) | 1))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi58
                }
            }
        } else {
            v59 := GenS8{F0: 1, F1: 9223372036854775807, F2: -61}
            for fi60 := 0; fi60 < 15; fi60++ {
                v2 = append(v2, "\n\t")
                var v61 uint = 4294967295
                _ = fi60
                _ = v61
            }
            var v62 int = 592
            _ = v59
            _ = v62
        }
    case 32767:
        _ = len(v2)
    case 2147483647:
        var v63 float32 = -2.25
        _ = v63
    }
    v3 = (v1 || v3)
    switch (("" + ("hello" + "hello")) + "hello") {
    case "\n\t":
        fmt.Println(v4, v1, v56)
    default:
        if v56 {
            v64 := GenS9{F0: 51}
            v56 = true
            var v65 int64 = -1
            _ = v64
            _ = v65
        } else {
            var v66 string = "a"
            var v67 int16 = 0
            _ = v66
            _ = v67
        }
    }
    v68 := GenS10{F0: false, F1: 4294967295, F2: -2147483648}
    v69 := GenS5{F0: 42, F1: 0}
    var v70 int64 = -1
    fmt.Println(v1, v4, v56)
    var v71 float64 = ((3.14159 / 0.1) / ((1e10 * (float64((float32(((3.14159 - 1.5) - (((((((3.14159 - 1e10) / ((0.1 + float64(1e10)) / 1.5)) * (0.1 * 1e10)) * float64(0.0)) - (((1.5 + 0.0) / 1.0) + 0.0)) * ((((0.1 / (0.1 - 0.0)) - (-2.25 * 1.5)) / float64(3.14159)) * (0.1 * 1e10))) * 1.5))) + (1.5 * 1.5))) + (1e10 - (3.14159 - 0.0)))) * (((1.5 + ((((0.1 + 1e10) / -2.25) / float64(1.5)) + (((0.0 - 0.1) - 3.14159) / (0.1 + 0.0)))) * ((0.0 * -2.25) + 3.14159)) + 0.1)))
    fmt.Println(v70)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v56
    _ = v68
    _ = v69
    _ = v70
    _ = v71
}
