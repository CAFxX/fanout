package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint64 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int; F1 int64; F2 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8 }
func genFn0(p0 int64, p1 int16) string {
    fmt.Println(p1)
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := GenS1{F0: -1, F1: (p0 / (((^(p0))) | 1)), F2: 42}
        _ = fi5
        _ = v6
    }
    v7 := []int{127, 287}
    if len(v7) > 0 {
        v7[0] = 893
    }
    _ = p0
    _ = p1
    _ = v7
    return "\n\t"
}
func genFn1() uint64 {
    switch "foo bar" {
    case "foo bar":
        v8 := []int{32768, 9223372036854775807, 4294967295, -128}
        _ = v8
    default:
        v9 := GenS1{F0: -128, F1: 139, F2: 4294967295}
        _ = v9
    }
    var v10 int64 = -9223372036854775808
    _ = v10
    return 847
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 4294967295
    v3 := make(map[int]uint32)
    v4 := []string{"foo bar", "", "", ("hello" + "hello")}
    v11 := GenS2{F0: 9223372036854775807, F1: 128}
    switch 0 {
    case 140:
        for fi12 := 0; fi12 < 15; fi12++ {
            v2 -= (v2 ^ v2)
            _ = fi12
        }
    case 934:
        switch 923 {
        case 706:
            var v13 uint16 = 285
            _ = v13
        case 19:
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := v3[fi14]
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi14
                _ = v15
            }
        case -851:
            v2 = 2
        default:
            var v16 uint8 = 1
            _ = v16
        }
    }
    v2 += 1
    fmt.Println(v2)
    for fi17 := 0; fi17 < 18; fi17++ {
        if (!(false && ((756 <= 2) && true))) {
            var v18 string = "foo bar"
            switch "a" {
            case "a":
                v19 := GenS3{F0: fi17, F1: 42, F2: (-(fi17))}
                _ = v19
            }
            var v20 uint16 = 255
            _ = v18
            _ = v20
        } else {
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 string = ("foo bar" + "hello")
                switch "a" {
                case "hello":
                    if true {
                        if (!true) {
                            v23 := GenS1{F0: -1, F1: 276, F2: v2}
                            fmt.Println(v22, fi17)
                            var v24 uint64 = v2
                            fmt.Println(v22, fi21)
                            _ = v23
                            _ = v24
                        }
                    } else {
                        v22 = v22
                        switch "\n\t" {
                        case "\n\t":
                            if true {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v25 := GenS4{F0: (-95 < 53)}
                                if (2147483648 > 109) {
                                    v2 += (v2 + (v2 << ((0) & 63)))
                                    var v26 int = 9223372036854775807
                                    v27 := GenS1{F0: 0, F1: -964, F2: (v2 * 792)}
                                    _ = v26
                                    _ = v27
                                } else {
                                    v22 = v22
                                    v4 = append(v4, (v22 + v22))
                                    v28 := GenS5{F0: 1}
                                    _ = v28
                                }
                                _ = v25
                            } else {
                                fmt.Println(v2, fi21, v2)
                                _ = len(v1)
                                v4 = append(v4, (("" + "a") + v22))
                            }
                        case "foo bar":
                            switch ("\n\t" + "") {
                            case "a":
                                if true {
                                    var v29 uint = 1
                                    if (true && (false && false)) {
                                        var v30 uint = v29
                                        var v31 int = fi17
                                        switch (fi17 % ((fi21) | 1)) {
                                        case 127:
                                            var v32 int32 = 1
                                            _ = v32
                                        case 4294967295:
                                            v2 -= v2
                                        default:
                                            v29 = 739
                                        }
                                        _ = v30
                                        _ = v31
                                    }
                                    var v33 uint32 = 4294967295
                                    _ = v29
                                    _ = v33
                                } else {
                                    v2 = v2
                                    v4 = append(v4, v22)
                                    v2 = (v2 + v2)
                                    v2 = 42
                                }
                            case "":
                                switch (fi21 << ((581) & 63)) {
                                case -196:
                                    var v34 int8 = 127
                                    _ = v34
                                }
                            default:
                                v2 = v2
                            }
                        case "a":
                            v35 := GenS2{F0: 42, F1: fi17}
                            _ = v35
                        default:
                            var v36 string = "foo bar"
                            _ = v36
                        }
                    }
                }
                _ = fi21
                _ = v22
            }
            v37 := GenS4{F0: (2 <= 0)}
            v2 -= genFn1()
            _ = v37
        }
        var v38 int32 = 525
        switch fi17 {
        case 404:
            fmt.Println(fi17, fi17)
        }
        _ = fi17
        _ = v38
    }
    for fi39 := 0; fi39 < 12; fi39++ {
        for fi40 := 0; fi40 < 6; fi40++ {
            v4 = append(v4, "hello")
            v41 := GenS2{F0: 1, F1: -9223372036854775808}
            if ((v2 & 2) <= v2) {
                var v42 uint8 = 127
                _ = v42
            } else {
                v4 = append(v4, "foo bar")
                var v43 string = (("hello" + "hello") + "")
                _ = v43
            }
            _ = fi40
            _ = v41
        }
        v44 := GenS6{F0: 100}
        _ = fi39
        _ = v44
    }
    if (v2 <= genFn1()) {
        fmt.Println(v2)
    } else {
        var v45 float32 = 1.5
        _ = v45
    }
    v2 += (^(v2))
    var v46 float32 = 0.1
    var v47 uint64 = 460
    v48 := GenS3{F0: -9223372036854775808, F1: 29, F2: 32768}
    var v49 int8 = 100
    fmt.Println(v2)
    if (-974 == -498) {
        var v50 uint16 = 724
        _ = v50
    } else {
        for fi51 := 0; fi51 < 5; fi51++ {
            var v52 uint32 = 2147483648
            switch "a" {
            case "a":
                switch fi51 {
                case 1:
                    v53 := GenS2{F0: 0, F1: (fi51 ^ int(v2))}
                    _ = v53
                case -371:
                    for fi54 := 0; fi54 < 5; fi54++ {
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        _ = fi54
                    }
                }
            case "hello":
                for fi55 := 0; fi55 < 14; fi55++ {
                    var v56 uint16 = 32767
                    var v57 int8 = (v49 * 0)
                    var v58 bool = false
                    v56 = 162
                    _ = fi55
                    _ = v56
                    _ = v57
                    _ = v58
                }
            case "\n\t":
                switch (-(fi51)) {
                case 972:
                    v59 := GenS2{F0: -9223372036854775808, F1: (-(fi51))}
                    _ = v59
                }
            }
            _ = fi51
            _ = v52
        }
        var v60 float64 = float64(1.5)
        _ = v60
    }
    v61 := GenS7{F0: v49}
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v61
}
