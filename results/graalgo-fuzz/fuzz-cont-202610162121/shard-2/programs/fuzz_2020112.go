package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint64; F1 int16; F2 uint }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 uint8; F1 int; F2 int8 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 uint; F1 int8; F2 int64 }
type GenS8 struct { F0 int64; F1 string; F2 uint64 }
type GenS9 struct { F0 uint8; F1 uint16 }
type GenS10 struct { F0 int32 }

func main() {
    var v1 bool = true
    v2 := []int{-1}
    var v3 string = ("\n\t" + "a")
    var v4 int16 = 1
    switch -2147483648 {
    case 32767:
        if (42 != 2) {
            fmt.Println(v4)
        }
    default:
        var v5 string = (v3 + v3)
        _ = v5
    }
    var v6 int32 = 2147483647
    if len(v2) > 0 {
        v2[0] = -9223372036854775808
    }
    var v7 string = "a"
    fmt.Println(v3)
    var v8 int64 = -1
    if true {
        if false {
            if (!(4294967295 >= 4294967295)) {
                var v9 bool = true
                v6 = 0
                if false {
                    switch (("hello" + ("foo bar" + v3)) + (((v3 + "") + v3) + "a")) {
                    case "":
                        var v10 float64 = 0.1
                        _ = v10
                    case "a":
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if false {
                                v3 = "\n\t"
                                if len(v2) > 0 {
                                    v2[0] = -2147483648
                                }
                            }
                            v9 = v1
                        }
                    case "hello":
                        if true {
                            v11 := GenS1{F0: 128, F1: 0}
                            if true {
                                var v12 uint8 = 255
                                var v13 int8 = -127
                                _ = v12
                                _ = v13
                            }
                            _ = v11
                        } else {
                            if (48 != 0) {
                                var v14 string = ("a" + (v7 + v7))
                                for fi15 := 0; fi15 < 9; fi15++ {
                                    for fi16 := 0; fi16 < 13; fi16++ {
                                        switch (v14 + v3) {
                                        case "foo bar":
                                            v17 := GenS2{F0: 4294967295, F1: 1, F2: 493}
                                            _ = v17
                                        case "":
                                            v8 = (-(v8))
                                        default:
                                            v9 = true
                                            v2 = append(v2, 128)
                                        }
                                        v18 := GenS3{F0: false, F1: -1}
                                        _ = fi16
                                        _ = v18
                                    }
                                    _ = fi15
                                }
                                if v1 {
                                    fmt.Println(v6, v4, v9)
                                    v9 = (259 != 1)
                                } else {
                                    for fi19 := 0; fi19 < 1; fi19++ {
                                        var v20 int64 = (v8 >> ((255) & 63))
                                        _ = fi19
                                        _ = v20
                                    }
                                    for fi21 := 0; fi21 < 5; fi21++ {
                                        v22 := GenS3{F0: v9, F1: (-(fi21))}
                                        v14 = v7
                                        for fi23 := 0; fi23 < 15; fi23++ {
                                            _ = len(v2)
                                            v8 = (v8 + (v8 << ((65535) & 63)))
                                            v8 = (v8 | v8)
                                            if (float32(0.1) == 0.0) {
                                                if true {
                                                    v8 += v8
                                                }
                                            } else {
                                                v14 = v7
                                                v1 = v9
                                                v9 = v1
                                            }
                                            if (!false) {
                                                v24 := GenS4{F0: int16(fi21), F1: (v6 % ((int32(fi23)) | 1))}
                                                if v9 {
                                                    if (!(85 > -22)) {
                                                        v4 += (v4 & 1000)
                                                        if ((42 > 638) && v9) {
                                                            var v25 int = fi21
                                                            _ = v25
                                                        } else {
                                                            v14 = (v14 + v3)
                                                            if true {
                                                                var v26 int32 = (v6 * (v6 % ((v6) | 1)))
                                                                _ = v26
                                                            } else {
                                                                fmt.Println(fi21, v3, fi21)
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v27 := GenS5{F0: 128, F1: fi23, F2: 42}
                                                                var v28 float32 = (-2.25 / 1.0)
                                                                _ = v27
                                                                _ = v28
                                                            }
                                                        }
                                                        if false {
                                                            if (65535 == 18446744073709551615) {
                                                                var v29 uint32 = 2147483648
                                                                var v30 float64 = 0.1
                                                                var v31 uint = 2
                                                                _ = v29
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                            v32 := GenS6{F0: 4294967295, F1: 228}
                                                            _ = v32
                                                        } else {
                                                            var v33 int16 = v4
                                                            var v34 uint16 = 978
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi23 + 9223372036854775807)
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = -223
                                                    }
                                                    switch ("\n\t" + "hello") {
                                                    case "hello":
                                                        if len(v2) > 0 {
                                                            v2[0] = fi21
                                                        }
                                                    }
                                                }
                                                fmt.Println(fi23, v14)
                                                if (32767 <= (v4 / ((-142) | 1))) {
                                                    v2 = append(v2, fi21)
                                                } else {
                                                    var v35 string = (((v7 + "foo bar") + (v3 + v7)) + ("a" + ((v3 + v3) + v3)))
                                                    fmt.Println(v6, v6)
                                                    switch v14 {
                                                    case "foo bar":
                                                        v4 -= (v4 / (((v4 << ((83) & 63))) | 1))
                                                    case "a":
                                                        var v36 uint32 = 0
                                                        _ = v36
                                                    case "\n\t":
                                                        v7 = ("\n\t" + (v3 + v14))
                                                    }
                                                    _ = v35
                                                }
                                                _ = v24
                                            } else {
                                                v2 = append(v2, (fi21 & (-(fi23))))
                                            }
                                            v37 := GenS7{F0: 491, F1: int8(fi21), F2: 1}
                                            if len(v2) > 0 {
                                                v2[0] = (fi21 << ((726) & 63))
                                            }
                                            _ = fi23
                                            _ = v37
                                        }
                                        _ = fi21
                                        _ = v22
                                    }
                                }
                                _ = v14
                            }
                            if true {
                                if v9 {
                                    var v38 bool = ((v4 + v4) <= -32768)
                                    if (!true) {
                                        fmt.Println(v1, v9, v1)
                                    }
                                    v3 = ("" + ("hello" + (((v7 + "hello") + v3) + v3)))
                                    _ = v38
                                }
                                for fi39 := 0; fi39 < 4; fi39++ {
                                    switch "a" {
                                    case "\n\t":
                                        switch fi39 {
                                        case -735:
                                            if v1 {
                                                v6 = (v6 + (v6 >> ((18446744073709551615) & 63)))
                                                var v40 int32 = (^(v6))
                                                _ = v40
                                            }
                                        case -246:
                                            v1 = (105 >= 842)
                                        }
                                    default:
                                        v41 := GenS8{F0: v8, F1: v3, F2: 4294967295}
                                        _ = v41
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi39
                                }
                                v42 := GenS9{F0: 0, F1: 255}
                                _ = v42
                            } else {
                                if v9 {
                                    var v43 uint = 985
                                    _ = v43
                                }
                            }
                            var v44 int8 = -1
                            _ = v44
                        }
                    }
                } else {
                    var v45 uint8 = 42
                    _ = v45
                }
                _ = v9
            } else {
                v6 = v6
                if ((((1e10 <= (-2.25 + (((3.14159 * 0.1) - float32(float64(3.14159))) / (1.5 / (float32(0.1) - (-2.25 * ((((0.1 * 1e10) + -2.25) / 3.14159) + 3.14159))))))) && v1) || (65535 == 65535)) && (255 >= 2)) {
                    if (0 < 42) {
                        v7 = ("" + v3)
                        for fi46 := 0; fi46 < 16; fi46++ {
                            switch "hello" {
                            case "":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                var v47 float32 = 3.14159
                                _ = v47
                            }
                            _ = fi46
                        }
                    } else {
                        switch 32768 {
                        case 615:
                            v1 = (!((true || ((v8 >> ((2) & 63)) == v8)) && true))
                        case -895:
                            v7 = "\n\t"
                        default:
                            v48 := GenS7{F0: 255, F1: 116, F2: v8}
                            _ = v48
                        }
                        var v49 uint64 = 2
                        for fi50 := 0; fi50 < 17; fi50++ {
                            var v51 bool = ((false || false) || v1)
                            if (42 >= 127) {
                                var v52 string = ("" + (("foo bar" + (((("\n\t" + (v3 + (v3 + v3))) + (("\n\t" + "hello") + "\n\t")) + (("a" + v7) + (("a" + "\n\t") + v3))) + v7)) + v3))
                                v52 = (v52 + ((v7 + ("\n\t" + "a")) + v3))
                                var v53 uint64 = v49
                                _ = v52
                                _ = v53
                            } else {
                                v7 = "foo bar"
                                v2 = append(v2, (fi50 & (fi50 * fi50)))
                            }
                            _ = fi50
                            _ = v51
                        }
                        _ = v49
                    }
                    if true {
                        v2 = append(v2, 1)
                        var v54 float64 = (0.0 + 3.14159)
                        _ = v54
                    }
                } else {
                    v6 = 0
                }
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v55 := GenS7{F0: 0, F1: 102, F2: 1}
            _ = v55
        }
        v56 := GenS10{F0: -979}
        if len(v2) > 0 {
            v2[0] = -2147483648
        }
        _ = v56
    }
    for fi57 := 0; fi57 < 9; fi57++ {
        v58 := GenS4{F0: (v4 % ((v4) | 1)), F1: v6}
        _ = fi57
        _ = v58
    }
    var v59 float32 = (float32(1.5) * float32(-2.25))
    var v60 bool = (-942 <= (^(v8)))
    var v61 string = v7
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v7
    _ = v8
}
