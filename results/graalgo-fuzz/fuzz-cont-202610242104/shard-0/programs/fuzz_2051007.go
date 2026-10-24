package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int }
type GenS2 struct { F0 uint32; F1 int64; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint }

func main() {
    var v1 float32 = 1.5
    var v2 uint16 = 416
    var v3 string = (((("\n\t" + ("" + "")) + "a") + ("" + "foo bar")) + "a")
    v4 := []string{(v3 + v3), v3, v3}
    var v5 int = 32768
    var v6 uint8 = 1
    switch 93 {
    case 127:
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                fmt.Println(v1, v3)
                if ((!true) || false) {
                    if len(v4) > 0 {
                        v4[0] = (v3 + "")
                    }
                }
                _ = fi8
            }
            var v9 int = v5
            _ = fi7
            _ = v9
        }
    case 876:
        fmt.Println(v2, v2, v6)
    case 70:
        var v10 string = (v3 + ("foo bar" + v3))
        _ = v10
    default:
        v3 = "a"
    }
    var v11 uint32 = 334
    if true {
        if (2 <= 65535) {
            var v12 string = ("hello" + v3)
            var v13 int = v5
            _ = v12
            _ = v13
        } else {
            if (false && true) {
                for fi14 := 0; fi14 < 4; fi14++ {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        fmt.Println(fi14, v2, v1)
                        v11 += 4294967295
                        _ = fi15
                    }
                    v2 -= (v2 + (v2 % ((v2) | 1)))
                    if false {
                        v2 = v2
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        v2 = (^(v2))
                    }
                    _ = fi14
                }
                v4 = append(v4, "")
            }
            v2 = (v2 | v2)
        }
        var v16 string = (v3 + v3)
        var v17 int8 = 0
        _ = v16
        _ = v17
    } else {
        fmt.Println(v3, v6, v6)
        fmt.Println(v11, v6, v1)
    }
    fmt.Println(v5, v6)
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    v5 -= v5
    var v18 uint8 = v6
    var v19 string = (v3 + (v3 + v3))
    if false {
        v3 = ((v19 + ((("foo bar" + "foo bar") + "hello") + v3)) + ("\n\t" + v19))
        var v20 float32 = 1.5
        switch -9223372036854775808 {
        case 758:
            switch (v3 + v19) {
            case "hello":
                var v21 int = (v5 >> ((760) & 63))
                _ = v21
            }
        default:
            if (2 >= (v2 & v2)) {
                var v22 uint32 = (v11 + v11)
                _ = v22
            } else {
                fmt.Println(v11, v18, v1)
                if (!true) {
                    v5 = v5
                    var v23 bool = (!(!true))
                    v3 = ((((v19 + "a") + (((v3 + v19) + v3) + v3)) + (v3 + v19)) + ("a" + "hello"))
                    _ = v23
                }
            }
        }
        _ = v20
    }
    if ((1e10 + 3.14159) > 0.0) {
        v24 := GenS1{F0: -121, F1: (!(v3 < v19)), F2: (-(v5))}
        _ = v24
    } else {
        if false {
            if (v2 > (v2 & 956)) {
                for fi25 := 0; fi25 < 10; fi25++ {
                    for fi26 := 0; fi26 < 4; fi26++ {
                        switch "" {
                        case "a":
                            if (0.0 == 0.1) {
                                fmt.Println(v3, fi26)
                            }
                        default:
                            v19 = v3
                        }
                        v6 = 128
                        v27 := GenS1{F0: 0, F1: false, F2: 127}
                        var v28 uint32 = v11
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    v29 := GenS2{F0: (v11 / ((65535) | 1)), F1: -1, F2: 100}
                    _ = fi25
                    _ = v29
                }
                v4 = append(v4, (("a" + ((v19 + v3) + v19)) + ""))
            } else {
                for fi30 := 0; fi30 < 4; fi30++ {
                    if true {
                        switch v3 {
                        case "\n\t":
                            for fi31 := 0; fi31 < 13; fi31++ {
                                var v32 uint32 = (v11 - (v11 * 4294967295))
                                switch 127 {
                                case 127:
                                    var v33 float32 = v1
                                    _ = v33
                                case 32768:
                                    v18 = v18
                                case -1:
                                    fmt.Println(v3)
                                }
                                var v34 uint64 = 0
                                _ = fi31
                                _ = v32
                                _ = v34
                            }
                        default:
                            for fi35 := 0; fi35 < 14; fi35++ {
                                v19 = "foo bar"
                                _ = fi35
                            }
                        }
                    } else {
                        v4 = append(v4, "foo bar")
                        for fi36 := 0; fi36 < 17; fi36++ {
                            if true {
                                var v37 uint32 = v11
                                switch fi36 {
                                case 128:
                                    if false {
                                        var v38 int16 = 1
                                        v37 -= (v11 & v37)
                                        _ = v38
                                    } else {
                                        v39 := GenS3{F0: 18446744073709551615, F1: 574}
                                        v4 = append(v4, "\n\t")
                                        fmt.Println(v3, fi36, v11)
                                        _ = v39
                                    }
                                case 32767:
                                    v19 = ((v19 + (v3 + v3)) + (v3 + "a"))
                                case 4294967295:
                                    switch fi36 {
                                    case 7:
                                        v6 -= 42
                                    case 128:
                                        v40 := GenS3{F0: 494, F1: 0}
                                        var v41 uint16 = v2
                                        _ = len(v4)
                                        v42 := GenS3{F0: 156, F1: 255}
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                default:
                                    var v43 uint64 = uint64(v2)
                                    _ = v43
                                }
                                fmt.Println(v19, v2, v2)
                                _ = v37
                            }
                            _ = fi36
                        }
                        v5 -= (-(fi30))
                    }
                    v44 := GenS1{F0: int8(fi30), F1: true, F2: 1}
                    _ = fi30
                    _ = v44
                }
                for fi45 := 0; fi45 < 8; fi45++ {
                    v6 = (v6 * 0)
                    v5 -= (v5 % (((fi45 - 42)) | 1))
                    _ = fi45
                }
                if len(v4) > 0 {
                    v4[0] = v3
                }
            }
            v11 -= v11
        }
        if len(v4) > 0 {
            v4[0] = v3
        }
    }
    var v46 uint8 = 2
    if (true || (!false)) {
        v46 -= (v18 + uint8(v11))
        var v47 float64 = 1.5
        _ = v47
    }
    v11 += v11
    if (((v3 + ((v3 + v19) + "a")) == "a") || false) {
        if true {
            v1 = v1
            for fi48 := 0; fi48 < 2; fi48++ {
                v4 = append(v4, v3)
                _ = fi48
            }
        } else {
            v19 = "foo bar"
            v1 = 3.14159
        }
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        var v50 uint8 = (^(v46))
        _ = fi49
        _ = v50
    }
    v11 = v11
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
