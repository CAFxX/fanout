package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint32; F1 bool; F2 int64 }
type GenS5 struct { F0 int16; F1 string; F2 uint64 }
type GenS6 struct { F0 int; F1 uint32; F2 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64; F1 uint64; F2 int16 }
func genFn0(p0 bool, p1 uint8, p2 int64) int64 {
    var v7 float64 = 0.0
    var v8 uint64 = 306
    v9 := GenS1{F0: 1, F1: 937}
    for fi10 := 0; fi10 < 13; fi10++ {
        if ((p0 && (false && false)) || p0) {
            var v11 bool = (p1 > 255)
            _ = v11
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int32 = -127
            if ("hello" > "") {
                v13 = (v13 - (v13 / ((1) | 1)))
            }
            v13 = (^(v13))
            v14 := GenS2{F0: 32767, F1: -50}
            p2 = int64(v13)
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 527
}

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]int)
    v3 := []uint32{}
    v4 := make(map[string]int)
    var v5 string = ("\n\t" + "foo bar")
    var v6 int8 = -1
    v5 = v5
    fmt.Println(v1, v6, v1)
    for fi15 := 0; fi15 < 16; fi15++ {
        if false {
            _ = len(v2)
            v16 := GenS2{F0: 32767, F1: (v6 - (-(v6)))}
            for fi17 := 0; fi17 < 1; fi17++ {
                delete(v4, "foo bar")
                if len(v3) > 0 {
                    v3[0] = 0
                }
                _ = fi17
            }
            _ = v16
        } else {
            var v18 int8 = 127
            fmt.Println(v1)
            _ = v18
        }
        _ = fi15
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 uint8 = 2
        _ = fi19
        _ = v20
    }
    if (!(((3.14159 / 1.0) + 0.0) == 3.14159)) {
        for fi21 := 0; fi21 < 2; fi21++ {
            for fi22 := 0; fi22 < 2; fi22++ {
                if (814 >= 4294967295) {
                    v6 = (v6 * -15)
                    delete(v2, uint32(v1))
                } else {
                    var v23 uint8 = 1
                    v2[65535] = fi22
                    _ = v23
                }
                _ = fi22
            }
            _ = fi21
        }
        if false {
            for fi24 := 0; fi24 < 1; fi24++ {
                v4[v5] = (-(fi24))
                var v25 int32 = -1
                for fi26 := 0; fi26 < 6; fi26++ {
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                    if (true && ((!false) && (0 < 4294967295))) {
                        v5 = "foo bar"
                        var v27 uint = 42
                        if (false || true) {
                            var v28 float32 = 1e10
                            _ = v28
                        } else {
                            if ((-1 != (v25 + v25)) && true) {
                                v27 -= (v27 - (v1 | (v1 + v1)))
                                var v29 int64 = 1
                                v30 := GenS3{F0: 42, F1: (v29 | genFn0(true, 255, v29)), F2: 301}
                                v6 = v6
                                if ((0.1 - 1e10) > (((-2.25 + (-2.25 + -2.25)) / float32((float64((((0.1 + ((float32(3.14159) - float32((1.5 / 3.14159))) - (3.14159 - ((-2.25 / 1.5) / 1e10)))) - (float32(0.0) - 1.5)) + (1.5 / float32(1.5)))) - 1.5))) * float32(0.1))) {
                                    var v31 uint32 = 1
                                    v32 := GenS4{F0: v31, F1: false, F2: int64(v1)}
                                    v5 = v5
                                    _ = v31
                                    _ = v32
                                }
                                v5 = (v5 + "foo bar")
                                _ = v29
                                _ = v30
                            }
                            v5 = ""
                        }
                        v6 = v6
                        _ = v27
                    }
                    _ = fi26
                }
                _ = fi24
                _ = v25
            }
            v33 := GenS5{F0: 62, F1: "a", F2: 2}
            _ = v33
        } else {
            if true {
                delete(v4, (v5 + (v5 + ((v5 + (v5 + v5)) + (v5 + "a")))))
            }
        }
        for fi34 := 0; fi34 < 11; fi34++ {
            var v35 bool = (!((!false) || true))
            v1 = (v1 % ((65535) | 1))
            _ = fi34
            _ = v35
        }
    }
    if false {
        var v36 uint16 = 65535
        for fi37 := 0; fi37 < 2; fi37++ {
            v5 = ("" + v5)
            _ = fi37
        }
        for fi38 := 0; fi38 < 6; fi38++ {
            v39 := GenS2{F0: -32768, F1: (v6 & (-(v6)))}
            _ = fi38
            _ = v39
        }
        _ = v36
    } else {
        switch (v5 + "\n\t") {
        case "":
            v40 := GenS4{F0: 0, F1: false, F2: 42}
            _ = v40
        }
        fmt.Println(v5, v6)
        fmt.Println(v5, v6, v6)
    }
    fmt.Println(v6, v6)
    switch -2147483648 {
    case 128:
        fmt.Println(v5, v5, v5)
    default:
        for fi41 := 0; fi41 < 3; fi41++ {
            _ = len(v2)
            if (!true) {
                for fi42 := 0; fi42 < 3; fi42++ {
                    var v43 bool = true
                    var v44 int32 = 1
                    var v45 float64 = 0.0
                    _ = fi42
                    _ = v43
                    _ = v44
                    _ = v45
                }
                for fi46 := 0; fi46 < 8; fi46++ {
                    if true {
                        v1 += v1
                        if true {
                            if ((42 != 2) && true) {
                                v6 += -39
                                v3 = append(v3, 169)
                                var v47 int = fi46
                                var v48 uint64 = 42
                                _ = v47
                                _ = v48
                            }
                            v6 += (-(v6))
                            v1 = v1
                            v49 := GenS4{F0: 1, F1: (true || (3.14159 > -2.25)), F2: -661}
                            _ = v49
                        } else {
                            switch v5 {
                            case "foo bar":
                                var v50 uint16 = 255
                                _ = v50
                            default:
                                fmt.Println(v1, fi41)
                            }
                        }
                        v5 = (v5 + v5)
                    }
                    _ = fi46
                }
                if (true || true) {
                    for fi51 := 0; fi51 < 7; fi51++ {
                        v6 += int8(fi41)
                        v1 += (v1 / ((v1) | 1))
                        switch ((v5 + "hello") + v5) {
                        case "a":
                            var v52 int = fi51
                            _ = v52
                        case "":
                            v5 = ""
                            if true {
                                v53 := GenS1{F0: 32767, F1: 42}
                                switch "\n\t" {
                                case "\n\t":
                                    switch v5 {
                                    case "hello":
                                        v54 := GenS6{F0: (fi51 / ((-327) | 1)), F1: 65535, F2: 10}
                                        _ = v54
                                    case "":
                                        var v55 string = v5
                                        _ = v55
                                    }
                                default:
                                    fmt.Println(v6)
                                }
                                _ = v53
                            }
                            switch 9223372036854775807 {
                            case -2147483648:
                                v1 = (^(v1))
                            case -174:
                                var v56 float32 = float32(0.0)
                                _ = v56
                            }
                        case "foo bar":
                            v5 = ""
                        default:
                            v6 = -128
                        }
                        if false {
                            delete(v4, (v5 + "foo bar"))
                            delete(v4, v5)
                        } else {
                            if false {
                                switch (-(fi41)) {
                                case 0:
                                    if true {
                                        if (!true) {
                                            v57 := GenS6{F0: fi41, F1: 42, F2: 79}
                                            _ = v57
                                        }
                                    }
                                case 9223372036854775807:
                                    switch (^(fi41)) {
                                    case 9223372036854775807:
                                        delete(v4, v5)
                                    case -858:
                                        v3 = append(v3, 676)
                                    case 865:
                                        var v58 int32 = 0
                                        _ = v58
                                    }
                                case 611:
                                    switch (v5 + "\n\t") {
                                    case "\n\t":
                                        if true {
                                            fmt.Println(v1, fi41, fi51)
                                        } else {
                                            _ = len(v2)
                                            v59 := GenS3{F0: 2, F1: genFn0((65535 == 255), 0, -9223372036854775808), F2: 255}
                                            _ = v59
                                        }
                                    default:
                                        _ = len(v4)
                                    }
                                }
                                _ = len(v3)
                                v60 := GenS7{F0: 1}
                                _ = v60
                            } else {
                                v1 += (v1 & (v1 | v1))
                            }
                            v5 = ""
                            var v61 uint = (^(v1))
                            v3 = append(v3, 185)
                            switch (fi51 | (fi41 | fi41)) {
                            case 1:
                                v62 := v4[v5]
                                _ = v62
                            case -956:
                                v6 = 57
                            default:
                                var v63 int = (-(fi51))
                                _ = v63
                            }
                            _ = v61
                        }
                        _ = fi51
                    }
                    if false {
                        if true {
                            var v64 string = v5
                            fmt.Println(v5, v6)
                            _ = v64
                        } else {
                            fmt.Println(v5, fi41)
                        }
                        v65 := GenS4{F0: 2147483648, F1: (-32768 > -1), F2: -1}
                        for fi66 := 0; fi66 < 17; fi66++ {
                            var v67 int16 = 1
                            _ = fi66
                            _ = v67
                        }
                        _ = v65
                    } else {
                        v68 := GenS6{F0: (fi41 ^ (fi41 - 4294967295)), F1: 4294967295, F2: 39}
                        var v69 string = v5
                        if len(v3) > 0 {
                            v3[0] = 0
                        }
                        _ = v68
                        _ = v69
                    }
                }
            }
            _ = fi41
        }
    }
    v70 := GenS7{F0: 65535}
    var v71 int8 = (^(v6))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 = (((v5 + ("foo bar" + "a")) + v5) + ("" + v5))
    if (1000 != 1) {
        v72 := GenS8{F0: 2, F1: 4294967295, F2: -32768}
        v5 = (v5 + (v5 + v5))
        v73 := GenS8{F0: 515, F1: 927, F2: 1000}
        _ = v72
        _ = v73
    } else {
        v74 := GenS1{F0: 1000, F1: 522}
        v75 := v4["hello"]
        _ = v74
        _ = v75
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
}
