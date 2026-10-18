package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint }
type GenS2 struct { F0 string; F1 uint8; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint32; F1 int8 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint = 1
    var v3 int8 = 100
    v4 := []uint32{v1, (v1 & v1), v1}
    v1 = (v1 >> (((v2 | 604)) & 63))
    if true {
        for fi5 := 0; fi5 < 7; fi5++ {
            for fi6 := 0; fi6 < 10; fi6++ {
                var v7 uint8 = 128
                if (((false && (true || (!false))) || (!true)) && false) {
                    var v8 uint = (^(v2))
                    fmt.Println(fi5, v2)
                    var v9 uint = (v2 + 657)
                    _ = v8
                    _ = v9
                }
                switch "hello" {
                case "":
                    var v10 uint64 = 625
                    _ = v10
                case "a":
                    v11 := GenS1{F0: 1, F1: 0, F2: (v2 * uint(v1))}
                    _ = v11
                }
                _ = fi6
                _ = v7
            }
            for fi12 := 0; fi12 < 16; fi12++ {
                switch (("foo bar" + "\n\t") + "") {
                case "a":
                    if ("hello" < "") {
                        var v13 int64 = -302
                        v4 = append(v4, v1)
                        if false {
                            var v14 string = ((("hello" + "a") + ("\n\t" + "foo bar")) + "")
                            v4 = append(v4, 65535)
                            if ((fi12 + fi12) > (fi5 & fi5)) {
                                switch (v14 + ("\n\t" + v14)) {
                                case "\n\t":
                                    v1 += 2
                                case "hello":
                                    v14 = (v14 + v14)
                                case "a":
                                    var v15 int64 = (v13 >> (((v2 / (((v2 | v2)) | 1))) & 63))
                                    _ = v15
                                default:
                                    var v16 uint32 = v1
                                    _ = v16
                                }
                            }
                            _ = v14
                        } else {
                            switch (("\n\t" + "foo bar") + "") {
                            case "\n\t":
                                v13 += 769
                            case "a":
                                var v17 int8 = (^(v3))
                                _ = v17
                            case "foo bar":
                                var v18 int32 = 0
                                _ = v18
                            }
                        }
                        _ = len(v4)
                        _ = v13
                    } else {
                        if (!(!false)) {
                            v2 = v2
                        }
                    }
                }
                _ = fi12
            }
            _ = fi5
        }
    }
    var v19 int32 = 0
    fmt.Println(v19)
    var v20 int8 = (v3 + (v3 % ((int8(v19)) | 1)))
    var v21 uint64 = 4294967295
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 uint32 = v1
        for fi24 := 0; fi24 < 5; fi24++ {
            switch "hello" {
            case "hello":
                if (!(!true)) {
                    v3 -= (v20 | v20)
                } else {
                    v20 += (-(v20))
                    if false {
                        v25 := GenS1{F0: v21, F1: v21, F2: (v2 ^ v2)}
                        var v26 int64 = -997
                        if (v19 < 63) {
                            if (649 >= v1) {
                                if false {
                                    if (!(1 <= 443)) {
                                        switch -1 {
                                        case 2147483647:
                                            if ((v3 == (v3 - (v20 - v3))) && (true || true)) {
                                                v20 -= (v3 << ((v2) & 63))
                                                v3 = (v20 << (((v2 % (((v2 / ((v2) | 1))) | 1))) & 63))
                                                var v27 bool = false
                                                _ = v27
                                            } else {
                                                if false {
                                                    switch "" {
                                                    case "foo bar":
                                                        var v28 uint = v2
                                                        _ = v28
                                                    case "a":
                                                        v29 := GenS2{F0: "", F1: 2, F2: "a"}
                                                        _ = v29
                                                    }
                                                }
                                                var v30 int8 = int8(v2)
                                                _ = v30
                                            }
                                        }
                                        v31 := GenS3{F0: (v23 + (v1 & v23))}
                                        var v32 int = 42
                                        _ = v31
                                        _ = v32
                                    } else {
                                        var v33 uint = (v2 ^ (^(v2)))
                                        v20 = (v20 - v3)
                                        switch (("" + (("\n\t" + "\n\t") + ("\n\t" + ""))) + "foo bar") {
                                        case "foo bar":
                                            if true {
                                                fmt.Println(v33, fi22, v33)
                                                var v34 int = fi24
                                                _ = v34
                                            } else {
                                                if (!true) {
                                                    v35 := GenS2{F0: ((((("hello" + "foo bar") + "") + "hello") + "hello") + "hello"), F1: 128, F2: ""}
                                                    _ = v35
                                                } else {
                                                    v19 -= (v19 / (((v19 ^ 42)) | 1))
                                                }
                                                fmt.Println(v33)
                                                v1 = (v23 & v1)
                                            }
                                        case "\n\t":
                                            v36 := GenS4{F0: int(v1)}
                                            _ = v36
                                        default:
                                            v3 = 100
                                            if true {
                                                v2 += (v2 + v33)
                                            }
                                            v20 = v20
                                        }
                                        switch (-(fi22)) {
                                        case 244:
                                            fmt.Println(v1, fi22, v3)
                                        case -550:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        _ = v33
                                    }
                                    v4 = append(v4, 380)
                                }
                                if false {
                                    v21 = v21
                                    v26 -= int64(v3)
                                    if (0 <= (v20 >> ((v2) & 63))) {
                                        if (!((v3 | 1) >= v3)) {
                                            v2 = v2
                                        }
                                        v21 = (v21 % ((v21) | 1))
                                    } else {
                                        if true {
                                            v3 -= (v3 | v3)
                                            if false {
                                                if len(v4) > 0 {
                                                    v4[0] = v1
                                                }
                                                _ = len(v4)
                                            } else {
                                                if false {
                                                    switch ((("" + "hello") + "hello") + (("\n\t" + "a") + "")) {
                                                    case "hello":
                                                        var v37 uint8 = 42
                                                        _ = v37
                                                    case "a":
                                                        _ = len(v4)
                                                    default:
                                                        fmt.Println(v21)
                                                    }
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v38 uint32 = v23
                                                    v23 -= v23
                                                    var v39 uint = (v2 | (^(v2)))
                                                    _ = v38
                                                    _ = v39
                                                }
                                            }
                                            fmt.Println(v23, v3)
                                        } else {
                                            v21 -= v21
                                        }
                                        v26 -= 9223372036854775807
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                }
                                var v40 int64 = (-(v26))
                                _ = v40
                            } else {
                                if (fi22 >= 0) {
                                    var v41 uint8 = 198
                                    var v42 string = "a"
                                    v4 = append(v4, (v1 ^ (v23 + v1)))
                                    _ = v41
                                    _ = v42
                                }
                                v21 = v21
                                if len(v4) > 0 {
                                    v4[0] = v1
                                }
                            }
                            fmt.Println(v2)
                            if (false || false) {
                                var v43 float32 = float32((1.5 * 1.5))
                                v26 -= -1
                                _ = len(v4)
                                var v44 int16 = int16(v20)
                                _ = v43
                                _ = v44
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = (v23 | v23)
                                }
                            }
                            switch "hello" {
                            case "\n\t":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case "a":
                                v4 = append(v4, (v1 & (v23 ^ v1)))
                            default:
                                fmt.Println(v19, v1, v20)
                            }
                            var v45 uint8 = 127
                            v46 := GenS5{F0: 1, F1: 428, F2: 589}
                            if len(v4) > 0 {
                                v4[0] = v1
                            }
                            _ = v45
                            _ = v46
                        }
                        _ = v25
                        _ = v26
                    }
                }
            case "a":
                switch "" {
                case "foo bar":
                    var v47 uint32 = uint32(v21)
                    _ = v47
                case "\n\t":
                    var v48 uint16 = 873
                    _ = v48
                }
            case "\n\t":
                var v49 uint64 = (v21 + (v21 << ((65535) & 63)))
                _ = v49
            }
            _ = fi24
        }
        v4 = append(v4, (v23 ^ (v1 | uint32(v2))))
        _ = fi22
        _ = v23
    }
    v50 := GenS6{F0: (v3 & (v3 >> ((v2) & 63)))}
    for fi51 := 0; fi51 < 5; fi51++ {
        v52 := GenS7{F0: ((420 < -32768) || (false || true))}
        _ = fi51
        _ = v52
    }
    fmt.Println(v2)
    if true {
        v53 := GenS8{F0: 1}
        if false {
            if len(v4) > 0 {
                v4[0] = (v1 | 0)
            }
            if (!(0 <= v2)) {
                v20 = 36
            }
            if true {
                switch "\n\t" {
                case "":
                    switch 743 {
                    case -2147483648:
                        v3 -= 19
                    case 1:
                        v54 := GenS8{F0: 9223372036854775807}
                        _ = v54
                    case 128:
                        v3 += v3
                    default:
                        var v55 uint64 = v21
                        _ = v55
                    }
                case "hello":
                    switch ("a" + "hello") {
                    case "\n\t":
                        switch -2147483648 {
                        case 1:
                            v56 := GenS9{F0: 65535, F1: (v3 * v3)}
                            _ = v56
                        case -832:
                            _ = len(v4)
                        default:
                            for fi57 := 0; fi57 < 5; fi57++ {
                                fmt.Println(v19, v19, v2)
                                _ = fi57
                            }
                        }
                    case "a":
                        var v58 int16 = 1
                        _ = v58
                    case "hello":
                        if false {
                            var v59 uint8 = 127
                            v60 := GenS6{F0: v3}
                            _ = v59
                            _ = v60
                        }
                    }
                case "a":
                    var v61 int = -597
                    _ = v61
                }
            }
        }
        switch "a" {
        case "\n\t":
            if false {
                v62 := GenS4{F0: 32767}
                _ = v62
            } else {
                v63 := GenS9{F0: 42, F1: 100}
                var v64 uint32 = uint32(v2)
                _ = v63
                _ = v64
            }
        default:
            v21 += (v21 / ((v21) | 1))
        }
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v50
}
