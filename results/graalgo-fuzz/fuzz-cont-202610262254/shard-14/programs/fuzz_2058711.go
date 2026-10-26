package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 int64; F1 uint16; F2 uint }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 int16; F2 int32 }
type GenS10 struct { F0 int64; F1 bool; F2 int8 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint8; F1 uint }
func genFn0(p0 int, p1 int32) int64 {
    p0 = p0
    var v5 string = "hello"
    var v6 int64 = 1
    v7 := GenS1{F0: 230, F1: 2}
    if true {
        fmt.Println(v5, v5, p1)
        var v8 float64 = (((1e10 + 0.1) / 1e10) - float64(1e10))
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (-(v6))
}
func genFn1(p0 uint, p1 uint16) int16 {
    var v9 bool = true
    var v10 uint64 = 0
    fmt.Println(p0, p0, v10)
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 1
}

func main() {
    v1 := []string{(((("hello" + "hello") + "hello") + "hello") + "hello")}
    var v2 uint = 86
    v3 := []string{""}
    var v4 int = 0
    var v11 int = v4
    var v12 string = ("hello" + "hello")
    v2 -= (v2 & (v2 - 18446744073709551615))
    var v13 int16 = -32768
    var v14 uint8 = 1
    v13 -= genFn1(v2, uint16(v2))
    switch (v11 >> ((v2) & 63)) {
    case 32767:
        v13 += v13
    default:
        v4 -= v4
    }
    var v15 int64 = -9223372036854775808
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 16; fi17++ {
            if true {
                if false {
                    v18 := GenS2{F0: 32767, F1: 0}
                    _ = v18
                }
            }
            v15 = (v15 | (v15 & (-(v15))))
            _ = fi17
        }
        v19 := GenS3{F0: (^(v14)), F1: 2, F2: 2147483647}
        _ = fi16
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        if (!false) {
            v21 := GenS4{F0: 42, F1: v13}
            _ = v21
        }
        v22 := GenS5{F0: -79, F1: (((0.1 > float64(1e10)) || (262 != 18446744073709551615)) || (-14 != int8(v15)))}
        _ = fi20
        _ = v22
    }
    fmt.Println(v11)
    v4 -= int(v13)
    if ((true || true) && ((-(v13)) > (v13 + v13))) {
        switch (v4 & v11) {
        case -637:
            var v23 uint8 = (v14 << ((v2) & 63))
            _ = v23
        default:
            switch ("a" + "\n\t") {
            case "foo bar":
                for fi24 := 0; fi24 < 11; fi24++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v15 -= (v15 << ((v2) & 63))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi24
                }
            case "\n\t":
                v15 = 0
            default:
                var v25 bool = ((v14 & (v14 - v14)) >= v14)
                _ = v25
            }
        }
        switch v4 {
        case 42:
            if false {
                if false {
                    fmt.Println(v12)
                }
                v13 = genFn1((^(v2)), 42)
            } else {
                if false {
                    for fi26 := 0; fi26 < 15; fi26++ {
                        if true {
                            var v27 bool = (563 > 2)
                            _ = len(v1)
                            if (v27 && v27) {
                                var v28 int8 = 1
                                _ = v28
                            }
                            _ = v27
                        }
                        switch 536 {
                        case 0:
                            v12 = v12
                        case 923:
                            if ((3.14159 / float64(0.1)) != float64(0.0)) {
                                v1 = append(v1, v12)
                            }
                        case -2147483648:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = fi26
                    }
                }
                for fi29 := 0; fi29 < 17; fi29++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi29
                }
            }
        case 32768:
            var v30 uint16 = 32767
            _ = v30
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if false {
        switch "" {
        case "foo bar":
            var v31 string = v12
            _ = v31
        case "":
            v32 := GenS6{F0: 1, F1: v13}
            _ = v32
        default:
            for fi33 := 0; fi33 < 9; fi33++ {
                switch (("a" + "foo bar") + "\n\t") {
                case "hello":
                    v34 := GenS4{F0: 2, F1: genFn1((v2 / (((v2 * v2)) | 1)), 32767)}
                    _ = v34
                case "a":
                    fmt.Println(v11, v11, v12)
                default:
                    for fi35 := 0; fi35 < 9; fi35++ {
                        v36 := GenS7{F0: (^(v15)), F1: 255, F2: v2}
                        v15 = (v15 ^ v15)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi35
                        _ = v36
                    }
                }
                var v37 int64 = (v15 << (((^(v2))) & 63))
                _ = fi33
                _ = v37
            }
        }
        for fi38 := 0; fi38 < 10; fi38++ {
            v11 = 32767
            if false {
                _ = len(v3)
                if true {
                    switch (v12 + (((v12 + v12) + ((v12 + v12) + v12)) + v12)) {
                    case "a":
                        v13 = genFn1((v2 % ((v2) | 1)), 700)
                    default:
                        var v39 int = v4
                        _ = v39
                    }
                    v40 := GenS8{F0: v2}
                    _ = v40
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi38
        }
        if true {
            v41 := GenS9{F0: (v15 >> (((v2 << (((v2 * (^(v2)))) & 63))) & 63)), F1: (^(v13)), F2: 42}
            v4 += (v11 << ((v2) & 63))
            _ = v41
        }
    } else {
        if len(v1) > 0 {
            v1[0] = "a"
        }
    }
    v42 := GenS10{F0: (^(v15)), F1: true, F2: -1}
    v43 := GenS11{F0: v12}
    v12 = ""
    for fi44 := 0; fi44 < 5; fi44++ {
        for fi45 := 0; fi45 < 7; fi45++ {
            var v46 int64 = (v15 & -567)
            switch (v4 << ((23) & 63)) {
            case 32767:
                v47 := GenS12{F0: v14, F1: 0}
                _ = v47
            }
            switch (^(fi44)) {
            case -2147483648:
                v48 := GenS6{F0: -32, F1: (^(v13))}
                _ = v48
            case 4294967295:
                fmt.Println(v46, v2, fi44)
            default:
                v49 := GenS10{F0: int64(v14), F1: true, F2: -128}
                _ = v49
            }
            _ = fi45
            _ = v46
        }
        _ = fi44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
}
