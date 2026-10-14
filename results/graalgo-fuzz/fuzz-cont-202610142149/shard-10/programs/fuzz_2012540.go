package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint16; F1 string; F2 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 uint }

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 42
    var v3 int = 32767
    var v4 int64 = 9223372036854775807
    var v5 uint64 = (v2 >> ((358) & 63))
    var v6 string = ""
    v7 := []int{v3, (v3 - (^(v3))), (v3 << ((18446744073709551615) & 63)), -525}
    var v8 int16 = -930
    if (127 <= 127) {
        switch (v6 + v6) {
        case "":
            if len(v7) > 0 {
                v7[0] = v3
            }
        }
        if (1 < (v5 << ((4294967295) & 63))) {
            for fi9 := 0; fi9 < 13; fi9++ {
                v6 = "foo bar"
                v8 += (v8 ^ int16(v4))
                switch "a" {
                case "":
                    fmt.Println(v6)
                case "hello":
                    v10 := GenS1{F0: 2147483647, F1: ("a" + ((v6 + "foo bar") + "\n\t"))}
                    _ = v10
                case "a":
                    v6 = (v6 + (v6 + ((v6 + v6) + v6)))
                }
                _ = fi9
            }
            var v11 uint32 = 4294967295
            for fi12 := 0; fi12 < 14; fi12++ {
                v13 := GenS1{F0: (^(v11)), F1: (v6 + v6)}
                var v14 bool = true
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        } else {
            if (int(v5) == -179) {
                if len(v7) > 0 {
                    v7[0] = -501
                }
            } else {
                v15 := GenS1{F0: 2, F1: v6}
                if false {
                    switch v3 {
                    case 785:
                        switch v3 {
                        case 456:
                            if (!("foo bar" <= "a")) {
                                v16 := GenS1{F0: 2147483647, F1: v6}
                                fmt.Println(v4, v6, v2)
                                if (0 <= 65535) {
                                    for fi17 := 0; fi17 < 11; fi17++ {
                                        v7 = append(v7, (^(fi17)))
                                        _ = fi17
                                    }
                                } else {
                                    delete(v1, (v3 << ((537) & 63)))
                                    v18 := v1[v3]
                                    switch v3 {
                                    case 2147483647:
                                        var v19 bool = (!(float32(0.0) >= 1.5))
                                        _ = v19
                                    case 9223372036854775807:
                                        if true {
                                            if true {
                                                v20 := GenS2{F0: 32767, F1: v6, F2: false}
                                                var v21 int16 = (v8 & (v8 ^ v8))
                                                var v22 int8 = -128
                                                _ = v20
                                                _ = v21
                                                _ = v22
                                            }
                                        } else {
                                            var v23 int8 = int8(v8)
                                            v24 := GenS3{F0: ("" + v6)}
                                            _ = v23
                                            _ = v24
                                        }
                                    }
                                    _ = v18
                                }
                                _ = v16
                            }
                        case 128:
                            if (1.5 < 1.5) {
                                var v25 int16 = (v8 << ((1) & 63))
                                if ((float32(1e10) > (3.14159 / (0.1 * 1.5))) || false) {
                                    v3 -= (v3 / (((v3 - (v3 >> ((65535) & 63)))) | 1))
                                } else {
                                    if true {
                                        v4 = v4
                                        switch -9223372036854775808 {
                                        case 42:
                                            var v26 bool = true
                                            _ = v26
                                        }
                                        if true {
                                            v27 := v1[-358]
                                            _ = v27
                                        } else {
                                            var v28 uint = 683
                                            _ = v28
                                        }
                                    } else {
                                        for fi29 := 0; fi29 < 13; fi29++ {
                                            v8 = v8
                                            v30 := GenS1{F0: 237, F1: v6}
                                            _ = fi29
                                            _ = v30
                                        }
                                        _ = len(v1)
                                        if (65535 > 32767) {
                                            v5 = (v2 ^ 482)
                                            for fi31 := 0; fi31 < 14; fi31++ {
                                                for fi32 := 0; fi32 < 5; fi32++ {
                                                    v5 -= v5
                                                    switch (v6 + v6) {
                                                    case "a":
                                                        var v33 int64 = v4
                                                        _ = v33
                                                    }
                                                    v6 = (("" + v6) + (((v6 + v6) + "hello") + (v6 + (((v6 + v6) + v6) + ((v6 + ((v6 + ("foo bar" + ((v6 + "\n\t") + (v6 + v6)))) + v6)) + "foo bar")))))
                                                    _ = fi32
                                                }
                                                _ = fi31
                                            }
                                            v3 += (v3 - v3)
                                        }
                                    }
                                    switch v3 {
                                    case 128:
                                        v4 = v4
                                    }
                                    switch ((v6 + "foo bar") + (((v6 + (v6 + "hello")) + "foo bar") + v6)) {
                                    case "foo bar":
                                        v25 += (v25 + v25)
                                    }
                                }
                                _ = v25
                            }
                        case 32768:
                            switch v3 {
                            case -637:
                                v34 := GenS2{F0: 1, F1: ("foo bar" + v6), F2: false}
                                _ = v34
                            case 32767:
                                _ = len(v1)
                            default:
                                for fi35 := 0; fi35 < 1; fi35++ {
                                    for fi36 := 0; fi36 < 10; fi36++ {
                                        v37 := GenS1{F0: 2147483648, F1: (v6 + (v6 + (v6 + (v6 + v6))))}
                                        var v38 uint = 255
                                        _ = fi36
                                        _ = v37
                                        _ = v38
                                    }
                                    v39 := GenS4{F0: fi35, F1: 42, F2: 250}
                                    _ = fi35
                                    _ = v39
                                }
                            }
                        }
                    default:
                        for fi40 := 0; fi40 < 15; fi40++ {
                            for fi41 := 0; fi41 < 13; fi41++ {
                                fmt.Println(v6, v2, v5)
                                _ = fi41
                            }
                            fmt.Println(v4, v4)
                            _ = fi40
                        }
                    }
                    v4 = (v4 / ((v4) | 1))
                }
                v5 = v5
                _ = v15
            }
        }
    } else {
        v8 += v8
        var v42 uint = uint(v5)
        _ = v42
    }
    var v43 uint64 = (v2 % ((v5) | 1))
    for fi44 := 0; fi44 < 6; fi44++ {
        if (608 > 18446744073709551615) {
            v3 = v3
            var v45 int32 = 1
            for fi46 := 0; fi46 < 18; fi46++ {
                if (!(128 < (fi44 << ((2) & 63)))) {
                    v47 := GenS2{F0: 255, F1: (v6 + ("" + ("a" + ((v6 + "a") + "foo bar")))), F2: ((v8 * (v8 << ((2) & 63))) >= -28)}
                    v2 += 401
                    v5 += v43
                    _ = v47
                }
                v2 -= (v43 >> ((556) & 63))
                v43 -= (v43 % (((v2 * 953)) | 1))
                _ = fi46
            }
            _ = v45
        } else {
            if len(v7) > 0 {
                v7[0] = (^(fi44))
            }
            if false {
                var v48 uint32 = 598
                var v49 uint16 = 42
                v3 = 1
                _ = v48
                _ = v49
            }
        }
        _ = fi44
    }
    for fi50 := 0; fi50 < 14; fi50++ {
        v6 = ((v6 + "\n\t") + "\n\t")
        delete(v1, v3)
        var v51 int64 = (v4 * (v4 + (v4 ^ 9223372036854775807)))
        _ = fi50
        _ = v51
    }
    var v52 uint = 4294967295
    switch v6 {
    case "a":
        if true {
            var v53 string = (v6 + (v6 + ((v6 + v6) + "hello")))
            v54 := GenS1{F0: 1, F1: "hello"}
            var v55 uint32 = 2147483647
            _ = v53
            _ = v54
            _ = v55
        }
    default:
        v56 := GenS5{F0: 0, F1: 2}
        _ = v56
    }
    fmt.Println(v4)
    var v57 int64 = (v4 * (^(v4)))
    for fi58 := 0; fi58 < 12; fi58++ {
        fmt.Println(v3)
        var v59 bool = true
        for fi60 := 0; fi60 < 16; fi60++ {
            switch (fi58 | (fi58 | (fi58 - 42))) {
            case -680:
                var v61 int32 = -1
                _ = v61
            case 323:
                _ = len(v1)
            case -372:
                if (!v59) {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    if len(v7) > 0 {
                        v7[0] = v3
                    }
                    v7 = append(v7, (v3 ^ 9223372036854775807))
                }
            }
            v62 := GenS6{F0: 18446744073709551615}
            _ = fi60
            _ = v62
        }
        _ = fi58
        _ = v59
    }
    v3 += (v3 | 32767)
    fmt.Println(v8, v6)
    var v63 uint = (v52 + (v52 + v52))
    v64 := GenS1{F0: 65535, F1: "hello"}
    v4 += (v57 | v57)
    v65 := GenS4{F0: -489, F1: 65535, F2: uint8(v2)}
    if false {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        if len(v7) > 0 {
            v7[0] = 32767
        }
        var v66 uint32 = 42
        _ = v66
    } else {
        v67 := v1[(v3 % ((v3) | 1))]
        var v68 float64 = 0.0
        _ = v67
        _ = v68
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v52
    _ = v57
    _ = v6
    _ = v63
    _ = v64
    _ = v65
    _ = v7
    _ = v8
}
