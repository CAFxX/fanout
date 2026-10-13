package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 uint8; F1 uint64; F2 string }
type GenS9 struct { F0 int64; F1 uint8; F2 int16 }
func genFn0(p0 uint64) uint8 {
    fmt.Println(p0)
    var v4 uint64 = p0
    switch "\n\t" {
    case "a":
        v4 = (v4 + v4)
    default:
        v5 := GenS1{F0: 42, F1: "a"}
        _ = v5
    }
    _ = p0
    _ = v4
    return 174
}

func main() {
    v1 := []uint32{}
    var v2 int32 = -2147483648
    var v3 uint8 = 248
    for fi6 := 0; fi6 < 19; fi6++ {
        v2 = (v2 >> ((4294967295) & 63))
        _ = fi6
    }
    var v7 int16 = 32767
    if true {
        var v8 string = "\n\t"
        switch 4294967295 {
        case 0:
            if false {
                v7 = -1
                v8 = v8
                var v9 int16 = (v7 % ((-1) | 1))
                _ = v9
            } else {
                for fi10 := 0; fi10 < 14; fi10++ {
                    for fi11 := 0; fi11 < 4; fi11++ {
                        var v12 uint16 = 424
                        _ = len(v1)
                        _ = fi11
                        _ = v12
                    }
                    var v13 uint = 847
                    if (((false && false) && true) || true) {
                        var v14 uint8 = (v3 + (v3 | v3))
                        v15 := GenS2{F0: (v3 >> ((v13) & 63)), F1: false}
                        _ = v14
                        _ = v15
                    } else {
                        if false {
                            v13 = v13
                            fmt.Println(fi10, v8, v3)
                            if false {
                                if (false && false) {
                                    var v16 int16 = 0
                                    _ = v16
                                }
                                switch fi10 {
                                case -2147483648:
                                    for fi17 := 0; fi17 < 16; fi17++ {
                                        v13 = 896
                                        fmt.Println(v13, v3, v7)
                                        var v18 uint32 = 4294967295
                                        _ = fi17
                                        _ = v18
                                    }
                                }
                                var v19 float32 = 1.5
                                _ = v19
                            }
                        } else {
                            var v20 uint64 = 4294967295
                            switch ("a" + (("" + v8) + v8)) {
                            case "":
                                if (1e10 == float64(0.1)) {
                                    v8 = ""
                                    fmt.Println(v3, fi10)
                                    var v21 bool = false
                                    _ = v21
                                } else {
                                    v3 = (v3 - (v3 % ((127) | 1)))
                                    fmt.Println(v8, fi10, v20)
                                    v1 = append(v1, 42)
                                }
                            }
                            _ = v20
                        }
                        v22 := GenS3{F0: ("foo bar" + v8)}
                        switch 101 {
                        case 127:
                            v23 := GenS4{F0: (606 >= 32767), F1: -1}
                            _ = v23
                        }
                        _ = v22
                    }
                    _ = fi10
                    _ = v13
                }
            }
        case 127:
            v1 = append(v1, 453)
        case 1:
            for fi24 := 0; fi24 < 4; fi24++ {
                switch (fi24 << ((4294967295) & 63)) {
                case 9223372036854775807:
                    v25 := GenS3{F0: v8}
                    _ = v25
                case 128:
                    fmt.Println(v7)
                case -64:
                    v26 := GenS2{F0: v3, F1: true}
                    _ = v26
                default:
                    v8 = v8
                }
                _ = fi24
            }
        default:
            v8 = "hello"
        }
        _ = v8
    } else {
        v27 := GenS2{F0: genFn0(157), F1: false}
        switch 32767 {
        case 32767:
            v3 = 207
        default:
            v7 = -355
        }
        var v28 int32 = (-(v2))
        _ = v27
        _ = v28
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        var v29 uint = 65535
        var v30 uint8 = uint8(v29)
        _ = v29
        _ = v30
    }
    v31 := GenS1{F0: 98, F1: "\n\t"}
    for fi32 := 0; fi32 < 8; fi32++ {
        switch -9223372036854775808 {
        case 9223372036854775807:
            v7 = v7
        case 727:
            var v33 int = fi32
            _ = v33
        case 326:
            for fi34 := 0; fi34 < 8; fi34++ {
                var v35 int = (fi32 >> ((255) & 63))
                v2 = (v2 ^ (v2 % (((-(v2))) | 1)))
                v36 := GenS5{F0: -352, F1: (!((v2 + v2) > v2))}
                _ = fi34
                _ = v35
                _ = v36
            }
        default:
            if ((("a" + "a") + "a") != (("\n\t" + "a") + ("foo bar" + ((("foo bar" + "a") + "hello") + "")))) {
                v37 := GenS6{F0: 2147483648}
                var v38 float64 = 1e10
                v7 = (v7 * v7)
                _ = v37
                _ = v38
            }
        }
        _ = fi32
    }
    for fi39 := 0; fi39 < 10; fi39++ {
        v40 := GenS4{F0: (255 <= 255), F1: 0}
        var v41 float32 = (((1.5 + ((1.5 * ((-2.25 * 3.14159) / 1.5)) - ((0.1 + 0.0) / 1.0))) + 0.0) * 1.5)
        _ = fi39
        _ = v40
        _ = v41
    }
    v42 := GenS7{F0: (true && ((-2.25 - 1e10) != 0.1)), F1: (v7 ^ (v7 + int16(v3)))}
    switch ("\n\t" + "\n\t") {
    case "hello":
        v7 = v7
    case "":
        for fi43 := 0; fi43 < 12; fi43++ {
            v7 = (v7 - v7)
            _ = fi43
        }
    default:
        for fi44 := 0; fi44 < 16; fi44++ {
            if false {
                v2 = 313
            }
            _ = fi44
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 += 42
    v45 := GenS8{F0: (v3 + (v3 | v3)), F1: 18446744073709551615, F2: "\n\t"}
    v46 := GenS9{F0: 0, F1: (v3 | 128), F2: (^(v7))}
    var v47 int = 127
    v7 -= -366
    if (!false) {
        v1 = append(v1, 2147483648)
    }
    v47 -= v47
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v7
}
