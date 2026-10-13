package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 string; F2 int64 }
func genFn0(p0 uint8, p1 int32, p2 uint16) uint64 {
    v5 := GenS1{F0: p1}
    if (true && false) {
        v6 := GenS1{F0: p1}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 13; fi7++ {
            p2 = (p2 / (((p2 | p2)) | 1))
            _ = fi7
        }
        p1 += 2147483647
        p1 = p1
    }
    v8 := GenS2{F0: (0.1 > (1e10 - ((1e10 - -2.25) * 1.5))), F1: 0, F2: p1}
    v9 := []int{-1, 843, 94}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    _ = v9
    return 0
}
func genFn1() string {
    var v10 uint32 = 42
    fmt.Println(v10)
    switch "hello" {
    case "foo bar":
        fmt.Println(v10)
    case "":
        for fi11 := 0; fi11 < 5; fi11++ {
            switch (fi11 ^ (^(fi11))) {
            case 4294967295:
                v10 -= (v10 << ((738) & 63))
            }
            _ = fi11
        }
    case "a":
        var v12 int = 128
        _ = v12
    default:
        for fi13 := 0; fi13 < 14; fi13++ {
            switch fi13 {
            case 630:
                for fi14 := 0; fi14 < 17; fi14++ {
                    var v15 string = "hello"
                    v16 := []int{}
                    if (0.0 <= 1.5) {
                        var v17 int = 32768
                        switch v15 {
                        case "hello":
                            if ((v10 << ((2) & 63)) < (v10 % ((v10) | 1))) {
                                v18 := GenS2{F0: false, F1: 255, F2: 174}
                                _ = v18
                            } else {
                                v19 := GenS2{F0: (!false), F1: 42, F2: 2147483647}
                                _ = v19
                            }
                        case "a":
                            fmt.Println(v15, v10, fi13)
                        case "\n\t":
                            fmt.Println(v17, fi13, v17)
                        default:
                            if (32767 > -639) {
                                v20 := GenS3{F0: (42 >= 0)}
                                v10 -= v10
                                if len(v16) > 0 {
                                    v16[0] = (v17 >> ((2) & 63))
                                }
                                _ = v20
                            }
                        }
                        v10 = v10
                        var v21 float64 = 1e10
                        v10 += (^(v10))
                        _ = v17
                        _ = v21
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case 981:
                if (260 == 42) {
                    if false {
                        v10 += 2147483647
                        var v22 int64 = 9223372036854775807
                        v22 -= v22
                        _ = v22
                    }
                    var v23 int = fi13
                    _ = v23
                }
            case 4294967295:
                v24 := []int{fi13, 849, (fi13 + (^(fi13)))}
                _ = v24
            }
            var v25 int = fi13
            switch (-(fi13)) {
            case -52:
                if false {
                    switch ("a" + ("a" + "")) {
                    case "":
                        var v26 uint = 669
                        _ = v26
                    case "hello":
                        v27 := []int{}
                        _ = v27
                    case "\n\t":
                        v28 := GenS1{F0: 42}
                        _ = v28
                    default:
                        for fi29 := 0; fi29 < 12; fi29++ {
                            v10 -= (v10 & v10)
                            _ = fi29
                        }
                    }
                    fmt.Println(fi13)
                    var v30 float64 = 0.0
                    _ = v30
                }
            case -9223372036854775808:
                switch (-(v25)) {
                case 2147483647:
                    v31 := GenS4{F0: 1, F1: "\n\t", F2: 9223372036854775807}
                    _ = v31
                case 128:
                    v32 := []int{}
                    _ = v32
                default:
                    v10 += (v10 - v10)
                }
            default:
                v10 = (v10 >> ((18446744073709551615) & 63))
            }
            _ = fi13
            _ = v25
        }
    }
    _ = v10
    return "hello"
}

func main() {
    var v1 float32 = (3.14159 - (-2.25 / 3.14159))
    v2 := []string{"\n\t", (((("hello" + "") + "") + (("a" + ("foo bar" + ((("a" + "hello") + "") + "\n\t"))) + ("" + ""))) + ("hello" + "\n\t"))}
    var v3 float32 = 0.1
    v4 := []string{"hello", ((("foo bar" + ((("foo bar" + "foo bar") + ("foo bar" + "foo bar")) + (("" + "") + "foo bar"))) + "\n\t") + "foo bar")}
    v1 += v1
    if (false && false) {
        v1 += v1
        var v33 int32 = -239
        switch (("" + "a") + "\n\t") {
        case "foo bar":
            v34 := GenS4{F0: 4294967295, F1: (("" + "hello") + (("a" + "") + "foo bar")), F2: 123}
            _ = v34
        case "\n\t":
            for fi35 := 0; fi35 < 1; fi35++ {
                v1 += (0.0 - v3)
                if false {
                    v3 = float32(0.1)
                } else {
                    v33 = v33
                }
                _ = fi35
            }
        case "hello":
            v36 := GenS4{F0: 0, F1: "\n\t", F2: 42}
            _ = v36
        }
        _ = v33
    }
    fmt.Println(v1, v1)
    if len(v2) > 0 {
        v2[0] = ((((("hello" + "foo bar") + "\n\t") + "\n\t") + "foo bar") + "a")
    }
    v1 += float32(3.14159)
    _ = len(v4)
    if true {
        for fi37 := 0; fi37 < 9; fi37++ {
            var v38 uint8 = 128
            v3 -= (1.5 - v1)
            fmt.Println(v3, v1)
            _ = fi37
            _ = v38
        }
        v1 += (v3 * ((v1 + 0.0) - float32(((0.0 / float64((((1.5 * (v3 / (1e10 + (3.14159 + v3)))) - ((1e10 - v3) + (1e10 + v1))) * 0.1))) / 0.1))))
    }
    v1 -= v1
    switch ((("hello" + "") + ("a" + (("" + "foo bar") + "\n\t"))) + ((("" + "a") + "foo bar") + "hello")) {
    case "hello":
        _ = len(v2)
    case "":
        v1 = (float32(-2.25) * 1.5)
    case "\n\t":
        for fi39 := 0; fi39 < 7; fi39++ {
            fmt.Println(v3)
            _ = fi39
        }
    default:
        var v40 string = ""
        _ = v40
    }
    for fi41 := 0; fi41 < 17; fi41++ {
        fmt.Println(v1)
        v3 += v1
        if (false || (42 < genFn0(255, -676, 42))) {
            switch -1 {
            case 0:
                switch fi41 {
                case -1:
                    var v42 uint16 = 255
                    _ = v42
                case 42:
                    v43 := GenS4{F0: 2, F1: "", F2: 0}
                    _ = v43
                case 2147483647:
                    v3 += v3
                }
            }
        } else {
            if (true || (!true)) {
                fmt.Println(fi41)
            } else {
                var v44 int64 = 9223372036854775807
                var v45 uint16 = 1
                v3 -= (3.14159 - (v1 + float32(0.0)))
                _ = v44
                _ = v45
            }
            for fi46 := 0; fi46 < 19; fi46++ {
                var v47 uint8 = 84
                _ = fi46
                _ = v47
            }
        }
        _ = fi41
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1, v3)
    var v48 uint8 = 238
    if false {
        if (494 == -1) {
            v1 += (1.5 * 0.0)
            var v49 float64 = 1.5
            v48 = (v48 * v48)
            _ = v49
        } else {
            fmt.Println(v3, v3, v48)
            fmt.Println(v3, v48)
            var v50 int = 9223372036854775807
            _ = v50
        }
        if true {
            v1 += 0.1
            if (646 == 4294967295) {
                if (!true) {
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                } else {
                    var v51 uint64 = 2
                    _ = v51
                }
                fmt.Println(v1, v3)
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
            }
            var v52 uint64 = 4294967295
            _ = v52
        } else {
            if false {
                v48 += (v48 - (v48 - (v48 >> ((497) & 63))))
                v53 := GenS2{F0: false, F1: 5, F2: -2147483648}
                _ = v53
            }
            switch "foo bar" {
            case "foo bar":
                var v54 int64 = 42
                _ = v54
            case "a":
                var v55 int64 = -1
                _ = v55
            case "":
                v3 = -2.25
            }
        }
    } else {
        switch 127 {
        case 548:
            for fi56 := 0; fi56 < 16; fi56++ {
                var v57 uint8 = (^(v48))
                v57 = v57
                _ = fi56
                _ = v57
            }
        case -9223372036854775808:
            v48 -= v48
        case 9223372036854775807:
            fmt.Println(v1, v3)
        default:
            for fi58 := 0; fi58 < 8; fi58++ {
                fmt.Println(v1, v3)
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    var v59 uint32 = 65535
                    _ = v59
                default:
                    _ = len(v4)
                }
                _ = fi58
            }
        }
        var v60 uint64 = 90
        _ = v60
    }
    var v61 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v61
}
