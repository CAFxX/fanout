package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint; F2 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -1
    var v3 float32 = (3.14159 + 3.14159)
    var v4 int16 = -32768
    v5 := []string{((("\n\t" + "hello") + "foo bar") + (("foo bar" + "") + "\n\t")), ("hello" + "a"), "foo bar"}
    var v6 int32 = 42
    v7 := GenS1{F0: 42, F1: 4294967295, F2: v6}
    delete(v1, 0)
    for fi8 := 0; fi8 < 19; fi8++ {
        if true {
            if false {
                if true {
                    v2 = -9223372036854775808
                }
                for fi9 := 0; fi9 < 18; fi9++ {
                    switch fi8 {
                    case 32768:
                        var v10 uint32 = 2147483647
                        _ = v10
                    case -374:
                        switch ("a" + ("foo bar" + "foo bar")) {
                        case "hello":
                            if true {
                                var v11 bool = (!(false || true))
                                var v12 int64 = (v2 | v2)
                                var v13 string = "\n\t"
                                _ = v11
                                _ = v12
                                _ = v13
                            }
                        case "\n\t":
                            if len(v5) > 0 {
                                v5[0] = ("foo bar" + "")
                            }
                        case "foo bar":
                            v14 := v1[0]
                            _ = v14
                        }
                        var v15 string = "foo bar"
                        _ = v15
                    case 0:
                        if ((128 < 0) && ((3.14159 < 3.14159) && (156 > 787))) {
                            var v16 int16 = v4
                            v3 += ((v3 * (0.1 / v3)) * v3)
                            var v17 int = fi9
                            _ = v16
                            _ = v17
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if (42 == 200) {
                        var v18 uint64 = 4294967295
                        v1[538] = (true || (false || false))
                        v19 := GenS1{F0: 128, F1: 0, F2: v6}
                        v20 := GenS2{F0: 2, F1: -32768}
                        v4 = -32768
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    switch 2147483647 {
                    case 2147483647:
                        _ = len(v1)
                    case -2147483648:
                        v4 -= -588
                    case 336:
                        var v21 string = "hello"
                        _ = v21
                    }
                    v22 := GenS1{F0: 177, F1: 2147483648, F2: (v6 & (^(v6)))}
                    _ = fi9
                    _ = v22
                }
            }
            v23 := GenS3{F0: false, F1: (^(v6))}
            v2 -= v2
            _ = v23
        }
        v24 := GenS1{F0: 128, F1: 4294967295, F2: int32(v2)}
        _ = fi8
        _ = v24
    }
    var v25 float32 = v3
    if (42 == 1) {
        fmt.Println(v6, v25)
        v25 = v25
        for fi26 := 0; fi26 < 19; fi26++ {
            if false {
                for fi27 := 0; fi27 < 17; fi27++ {
                    var v28 bool = false
                    var v29 string = "foo bar"
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                switch "" {
                case "hello":
                    v25 += v3
                case "foo bar":
                    if (-1 < v2) {
                        fmt.Println(v3)
                        v30 := GenS4{F0: 1, F1: 18446744073709551615, F2: (v6 | int32(v4))}
                        v31 := v1[727]
                        _ = v30
                        _ = v31
                    }
                default:
                    var v32 uint8 = 255
                    _ = v32
                }
            } else {
                if false {
                    fmt.Println(v2)
                    v33 := GenS5{F0: 4294967295}
                    _ = v33
                }
                fmt.Println(fi26)
            }
            var v34 int = (fi26 ^ (fi26 - 42))
            _ = fi26
            _ = v34
        }
    }
    if (v4 >= (v4 >> ((4294967295) & 63))) {
        v35 := GenS6{F0: "hello"}
        for fi36 := 0; fi36 < 15; fi36++ {
            var v37 bool = ((^(fi36)) > (fi36 << ((1) & 63)))
            _ = fi36
            _ = v37
        }
        var v38 uint16 = 65535
        _ = v35
        _ = v38
    }
    var v39 int32 = (^(v6))
    var v40 int = 2147483647
    switch 62 {
    case -961:
        v40 += (v40 & (v40 - 1))
    }
    var v41 string = ("foo bar" + "foo bar")
    var v42 int = (-(v40))
    v5 = append(v5, ("foo bar" + "foo bar"))
    v43 := GenS2{F0: 4294967295, F1: (v4 * -242)}
    _ = len(v1)
    switch v42 {
    case -217:
        v39 = (^(v39))
    default:
        _ = len(v1)
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        var v45 uint = 1
        var v46 uint16 = 42
        v47 := GenS5{F0: 2147483647}
        _ = fi44
        _ = v45
        _ = v46
        _ = v47
    }
    for fi48 := 0; fi48 < 7; fi48++ {
        v40 = (fi48 / (((v42 - v40)) | 1))
        _ = fi48
    }
    if (-16 == 19) {
        for fi49 := 0; fi49 < 20; fi49++ {
            switch int(v4) {
            case 848:
                switch (fi49 >> ((255) & 63)) {
                case 42:
                    v50 := GenS3{F0: (0 <= 437), F1: (v6 & (v6 << ((0) & 63)))}
                    _ = v50
                default:
                    v51 := GenS7{F0: 9223372036854775807, F1: 697}
                    _ = v51
                }
            default:
                v52 := GenS4{F0: 32767, F1: 4294967295, F2: (v39 & (v39 & 1))}
                _ = v52
            }
            var v53 bool = (!(v41 <= "foo bar"))
            fmt.Println(v42)
            _ = fi49
            _ = v53
        }
        if ((1.5 / (v25 - float32(3.14159))) != v25) {
            var v54 int32 = (v6 * int32(v40))
            var v55 int8 = -34
            var v56 float32 = v3
            _ = v54
            _ = v55
            _ = v56
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
