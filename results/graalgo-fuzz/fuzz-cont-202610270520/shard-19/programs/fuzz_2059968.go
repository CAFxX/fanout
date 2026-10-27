package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 bool; F1 int8; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 uint32; F2 int16 }
func genFn0(p0 string, p1 string) int64 {
    if true {
        v7 := []int{693, 9223372036854775807, 32768}
        _ = v7
    }
    switch p1 {
    case "":
        v8 := GenS1{F0: 80, F1: 42}
        _ = v8
    default:
        if (-32768 > 0) {
            fmt.Println(p0, p0)
            v9 := GenS1{F0: 65535, F1: 84}
            _ = v9
        } else {
            var v10 uint16 = 42
            _ = v10
        }
    }
    _ = p0
    _ = p1
    return -826
}

func main() {
    var v1 bool = true
    v2 := []uint32{2147483647, 797, 4294967295, 785}
    var v3 uint32 = 65535
    var v4 uint64 = 42
    var v5 uint16 = 1
    v6 := []string{"", "a", "foo bar"}
    fmt.Println(v4, v3, v4)
    var v11 string = "foo bar"
    var v12 int16 = 43
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int16 = v12
        fmt.Println(v1)
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        if (float64((0.1 - -2.25)) < float64(-2.25)) {
            switch (("foo bar" + "\n\t") + (v11 + ((v11 + v11) + v11))) {
            case "foo bar":
                var v16 int8 = 1
                _ = v16
            case "\n\t":
                switch (^(fi15)) {
                case 32767:
                    if v1 {
                        var v17 int8 = 127
                        _ = v17
                    }
                case 127:
                    v18 := GenS2{F0: ("a" + "\n\t"), F1: 650}
                    _ = v18
                case 751:
                    if len(v6) > 0 {
                        v6[0] = (("" + v11) + v11)
                    }
                }
            default:
                v12 = (v12 % ((-1) | 1))
            }
        }
        _ = fi15
    }
    var v19 string = "foo bar"
    switch -145 {
    case 484:
        v5 = 57
    default:
        fmt.Println(v19, v11)
    }
    if ((3.14159 > 1e10) && ((!v1) && (9223372036854775807 > genFn0(("\n\t" + ""), v11)))) {
        switch "" {
        case "":
            for fi20 := 0; fi20 < 2; fi20++ {
                v5 = v5
                _ = fi20
            }
        case "a":
            for fi21 := 0; fi21 < 11; fi21++ {
                v1 = (v3 < (v3 & v3))
                switch (^(fi21)) {
                case 4294967295:
                    v4 = 248
                case -1:
                    if (2 <= 4294967295) {
                        if (!v1) {
                            var v22 float64 = 3.14159
                            if false {
                                switch fi21 {
                                case -941:
                                    switch ("a" + "a") {
                                    case "hello":
                                        if true {
                                            var v23 string = "hello"
                                            _ = v23
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = 2
                                            }
                                            switch (fi21 >> ((4294967295) & 63)) {
                                            case -1:
                                                v24 := GenS1{F0: v3, F1: genFn0(v11, v11)}
                                                _ = v24
                                            case 32767:
                                                var v25 uint32 = (v3 / ((v3) | 1))
                                                _ = v25
                                            }
                                        }
                                    case "\n\t":
                                        v3 += v3
                                    }
                                case 145:
                                    v5 = 65535
                                default:
                                    v5 -= 768
                                }
                                var v26 uint64 = (v4 | 443)
                                v27 := GenS2{F0: v19, F1: 2147483647}
                                _ = v26
                                _ = v27
                            }
                            v28 := GenS2{F0: v11, F1: 164}
                            _ = v22
                            _ = v28
                        } else {
                            var v29 uint64 = v4
                            var v30 int64 = -1
                            _ = v29
                            _ = v30
                        }
                    }
                default:
                    var v31 int = fi21
                    _ = v31
                }
                if (v1 || v1) {
                    v4 += v4
                    for fi32 := 0; fi32 < 8; fi32++ {
                        switch (fi32 & (fi32 * -104)) {
                        case 986:
                            v6 = append(v6, (v11 + (v11 + (v19 + ((("" + "\n\t") + (v11 + v19)) + "")))))
                        case 343:
                            var v33 float32 = 1.5
                            _ = v33
                        case 0:
                            v3 = (v3 & (v3 & uint32(v5)))
                        default:
                            switch (("a" + (v11 + v19)) + v11) {
                            case "foo bar":
                                if (-1 <= -51) {
                                    fmt.Println(v5, v12, v12)
                                    var v34 uint32 = v3
                                    if ((v34 == v34) || true) {
                                        v2 = append(v2, (v34 & (v3 * (^(v34)))))
                                        var v35 int = fi21
                                        var v36 uint64 = 997
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = v34
                                }
                            case "\n\t":
                                v37 := GenS3{F0: ("hello" > v11), F1: 37, F2: (fi32 << ((491) & 63))}
                                _ = v37
                            default:
                                var v38 int = (fi32 & fi21)
                                fmt.Println(v5, fi21, v4)
                                _ = v38
                            }
                        }
                        var v39 uint32 = 2147483648
                        if len(v6) > 0 {
                            v6[0] = v11
                        }
                        _ = fi32
                        _ = v39
                    }
                } else {
                    v19 = "\n\t"
                }
                _ = fi21
            }
        default:
            if v1 {
                for fi40 := 0; fi40 < 15; fi40++ {
                    v11 = "a"
                    for fi41 := 0; fi41 < 18; fi41++ {
                        var v42 uint8 = 254
                        v43 := GenS4{F0: (v3 + (v3 ^ (^(v3))))}
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 11; fi44++ {
                        var v45 float64 = -2.25
                        _ = fi44
                        _ = v45
                    }
                    _ = fi40
                }
            }
        }
    }
    for fi46 := 0; fi46 < 15; fi46++ {
        var v47 float32 = 1.5
        var v48 float32 = 1e10
        if len(v2) > 0 {
            v2[0] = (^(v3))
        }
        _ = fi46
        _ = v47
        _ = v48
    }
    v11 = ("hello" + ((("hello" + (((v19 + v19) + "foo bar") + v11)) + "foo bar") + v19))
    if (2147483647 > 42) {
        v12 = (v12 * v12)
    } else {
        var v49 float64 = 3.14159
        switch -435 {
        case 2147483647:
            v50 := GenS5{F0: v19, F1: (^(v3)), F2: (v12 >> ((416) & 63))}
            _ = v50
        case 32768:
            var v51 int16 = v12
            _ = v51
        case 32767:
            var v52 uint16 = v5
            _ = v52
        default:
            _ = len(v2)
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v49
    }
    _ = len(v6)
    v3 = (v3 << ((18446744073709551615) & 63))
    for fi53 := 0; fi53 < 15; fi53++ {
        v12 = (v12 >> ((484) & 63))
        fmt.Println(v19)
        _ = fi53
    }
    v4 -= (v4 * 691)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
