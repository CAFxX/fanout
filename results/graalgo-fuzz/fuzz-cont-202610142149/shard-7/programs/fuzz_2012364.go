package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint16; F1 int8; F2 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 0
    var v3 uint32 = 0
    for fi4 := 0; fi4 < 3; fi4++ {
        switch ((("a" + "\n\t") + ("a" + ("foo bar" + ("\n\t" + "foo bar")))) + "foo bar") {
        case "hello":
            v5 := []int{fi4}
            _ = v5
        }
        _ = fi4
    }
    var v6 uint = 134
    v1[-1] = 238
    v2 = 2
    switch "a" {
    case "\n\t":
        fmt.Println(v2, v3, v6)
    default:
        fmt.Println(v6, v2, v6)
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 float64 = 1.5
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: (v3 ^ uint32(v2))}
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 uint16 = (v2 % (((v2 % (((v2 % ((v2) | 1))) | 1))) | 1))
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: 9223372036854775807, F1: (!false)}
    v13 := GenS1{F0: (^(v3))}
    _ = len(v1)
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 int32 = 530
            _ = len(v1)
            _ = fi15
            _ = v16
        }
        if ((!true) || false) {
            for fi17 := 0; fi17 < 9; fi17++ {
                if false {
                    if true {
                        v2 += (v2 << (((v6 % ((638) | 1))) & 63))
                        v6 = v6
                        v18 := []int{(fi17 + fi14), fi17}
                        if len(v18) > 0 {
                            v18[0] = -9223372036854775808
                        }
                        fmt.Println(fi17)
                        if false {
                            v1[fi14] = 32768
                            v2 = v2
                            var v19 uint8 = 222
                            _ = v19
                        }
                        _ = v18
                    }
                } else {
                    v20 := GenS1{F0: 350}
                    _ = v20
                }
                var v21 uint64 = 1
                if (true || ("" > (("\n\t" + "") + "hello"))) {
                    var v22 uint32 = 2
                    _ = v22
                }
                _ = fi17
                _ = v21
            }
        } else {
            var v23 uint64 = 0
            v24 := GenS3{F0: (^(fi14)), F1: (fi14 >> ((v6) & 63))}
            _ = v23
            _ = v24
        }
        if (!true) {
            v25 := GenS4{F0: v2, F1: -62, F2: 0}
            if true {
                fmt.Println(v6)
                v1[(^(fi14))] = 128
                var v26 int8 = 127
                _ = v26
            } else {
                fmt.Println(fi14, v6)
            }
            _ = v25
        } else {
            if false {
                if (!true) {
                    v27 := GenS1{F0: (v3 << (((^(v6))) & 63))}
                    v28 := GenS5{F0: fi14}
                    _ = v27
                    _ = v28
                } else {
                    switch (fi14 ^ 4294967295) {
                    case 128:
                        for fi29 := 0; fi29 < 9; fi29++ {
                            switch ((("a" + ("foo bar" + ("" + ""))) + "hello") + "a") {
                            case "foo bar":
                                fmt.Println(v6, fi14)
                            case "":
                                if true {
                                    v30 := []int{(fi29 << ((v6) & 63)), (^(fi29))}
                                    if false {
                                        v3 = (v3 * (^(v3)))
                                        if false {
                                            var v31 int16 = 42
                                            v3 = v3
                                            _ = v31
                                        } else {
                                            fmt.Println(fi14)
                                            _ = len(v1)
                                            v32 := GenS6{F0: -9223372036854775808, F1: (false || true)}
                                            _ = v32
                                        }
                                        v3 += (v3 | v3)
                                        var v33 bool = true
                                        var v34 uint16 = 186
                                        var v35 int8 = 100
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                    fmt.Println(v6, fi29)
                                    if (true || (340 > -802)) {
                                        if (!true) {
                                            var v36 int8 = 100
                                            var v37 int = fi29
                                            v6 -= 1
                                            _ = v36
                                            _ = v37
                                        }
                                        if true {
                                            fmt.Println(v2, fi29)
                                        }
                                    } else {
                                        fmt.Println(v2, v3, fi29)
                                        switch ((((("a" + "\n\t") + "a") + ("\n\t" + "hello")) + "hello") + ("hello" + ((("hello" + "a") + "\n\t") + ("" + "foo bar")))) {
                                        case "hello":
                                            if len(v30) > 0 {
                                                v30[0] = 128
                                            }
                                        case "foo bar":
                                            v38 := GenS7{F0: -1, F1: v3}
                                            _ = v38
                                        }
                                        var v39 int16 = 1000
                                        fmt.Println(v39)
                                        _ = v39
                                    }
                                    v3 = (v3 - v3)
                                    _ = v30
                                }
                            }
                            v40 := GenS1{F0: (v3 % ((v3) | 1))}
                            v41 := GenS8{F0: 0, F1: (v2 / ((1) | 1))}
                            switch "a" {
                            case "foo bar":
                                var v42 uint = (v6 + v6)
                                _ = v42
                            case "":
                                switch 42 {
                                case 32767:
                                    switch fi29 {
                                    case -892:
                                        switch (fi14 + (-(fi14))) {
                                        case 127:
                                            v3 = (v3 << (((v6 + (^(v6)))) & 63))
                                        case -9223372036854775808:
                                            if true {
                                                if (uint32(v6) != (v3 % (((v3 - 42)) | 1))) {
                                                    v6 += (^(v6))
                                                } else {
                                                    if (2147483647 >= 0) {
                                                        v3 = v3
                                                        v43 := GenS8{F0: 585, F1: (v2 / (((v2 | v2)) | 1))}
                                                        _ = v43
                                                    }
                                                }
                                            }
                                        case 196:
                                            if true {
                                                v44 := []int{fi29, fi14, (^(fi29))}
                                                _ = v44
                                            } else {
                                                fmt.Println(fi14, fi14)
                                                v1[fi29] = 32768
                                                fmt.Println(v6, v2)
                                                v6 = (v6 - v6)
                                            }
                                        default:
                                            var v45 uint64 = 836
                                            _ = v45
                                        }
                                    case -30:
                                        var v46 int16 = 42
                                        _ = v46
                                    case -615:
                                        var v47 uint = uint(v2)
                                        _ = v47
                                    }
                                case 691:
                                    var v48 uint8 = 127
                                    _ = v48
                                default:
                                    if (true || ("hello" != "hello")) {
                                        var v49 uint8 = 103
                                        fmt.Println(v49, v3)
                                        var v50 uint8 = v49
                                        var v51 int = -1
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    }
                                }
                            case "a":
                                v6 = 42
                            default:
                                if (!((!false) || (("foo bar" + "foo bar") >= ("foo bar" + (("foo bar" + "") + "foo bar"))))) {
                                    v52 := GenS9{F0: 457}
                                    _ = v52
                                }
                                v53 := GenS4{F0: (v2 - (v2 & (v2 >> (((v6 ^ v6)) & 63)))), F1: int8(v3), F2: 2147483647}
                                _ = v53
                            }
                            _ = fi29
                            _ = v40
                            _ = v41
                        }
                    case 42:
                        if (0.1 <= float32((((0.1 * ((((((float64(float32(1e10)) - 1.5) / float64(3.14159)) - 1.5) - 0.0) / 1.5) + 1.5)) - 0.1) - 1e10))) {
                            switch fi14 {
                            case 127:
                                v54 := GenS2{F0: fi14, F1: (float32(float64((float32(1.5) * -2.25))) > 0.1)}
                                _ = v54
                            case 128:
                                v55 := GenS9{F0: 2147483647}
                                _ = v55
                            case 569:
                                v2 = (^(v2))
                            }
                            v2 += 42
                        }
                    }
                    fmt.Println(v2)
                    v3 -= v3
                }
                v6 = v6
                v3 = v3
            }
            var v56 float32 = -2.25
            fmt.Println(v3, fi14)
            _ = v56
        }
        _ = fi14
    }
    if false {
        var v57 uint64 = 0
        _ = v57
    } else {
        var v58 uint64 = 4294967295
        _ = v58
    }
    for fi59 := 0; fi59 < 18; fi59++ {
        var v60 uint16 = (^(v2))
        fmt.Println(v2, v2)
        var v61 float32 = 0.1
        _ = fi59
        _ = v60
        _ = v61
    }
    v6 = uint(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
