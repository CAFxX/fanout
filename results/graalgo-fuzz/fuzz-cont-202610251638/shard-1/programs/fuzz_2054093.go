package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 string; F1 bool; F2 int32 }
type GenS6 struct { F0 string; F1 bool }

func main() {
    v1 := []string{"", "hello"}
    var v2 int8 = -45
    v3 := make(map[uint32]int)
    if ((2147483647 == 65535) && (true || (2147483647 < -1))) {
        fmt.Println(v2)
    }
    if false {
        var v4 bool = false
        v5 := GenS1{F0: 1, F1: 9223372036854775807}
        _ = v4
        _ = v5
    }
    v6 := GenS2{F0: 4294967295, F1: 42}
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int8 = (v2 / (((v2 % ((v2) | 1))) | 1))
        v8 = (v2 / ((v8) | 1))
        v9 := GenS2{F0: 9223372036854775807, F1: 756}
        _ = fi7
        _ = v8
        _ = v9
    }
    v2 = v2
    var v10 uint32 = 987
    v10 -= v10
    for fi11 := 0; fi11 < 9; fi11++ {
        switch fi11 {
        case 29:
            v12 := GenS2{F0: (-(fi11)), F1: 42}
            _ = v12
        case -2147483648:
            v13 := GenS1{F0: 160, F1: fi11}
            _ = v13
        case 128:
            var v14 uint64 = 42
            _ = v14
        }
        for fi15 := 0; fi15 < 1; fi15++ {
            fmt.Println(v10)
            v1 = append(v1, ("a" + ("a" + "\n\t")))
            v10 += v10
            _ = len(v1)
            switch (fi15 / ((fi15) | 1)) {
            case 4294967295:
                var v16 float32 = (-2.25 / ((-2.25 + (0.1 * (1e10 - (1.5 - (-2.25 / float32(1e10)))))) * -2.25))
                _ = v16
            case 1:
                v2 = (v2 * (v2 | v2))
                v17 := GenS1{F0: 127, F1: fi11}
                _ = v17
            case 82:
                var v18 int32 = 2147483647
                _ = v18
            default:
                if (-2147483648 == -2147483648) {
                    if true {
                        v19 := GenS1{F0: 2, F1: 9223372036854775807}
                        v10 = v10
                        v2 = (v2 << ((437) & 63))
                        _ = v19
                    } else {
                        fmt.Println(fi15)
                    }
                    fmt.Println(fi15)
                    var v20 bool = ((v10 - v10) <= (v10 % ((v10) | 1)))
                    _ = v20
                }
            }
            var v21 int8 = (v2 ^ v2)
            var v22 int16 = 0
            _ = fi15
            _ = v21
            _ = v22
        }
        if (!(!false)) {
            var v23 uint64 = 2
            _ = v23
        } else {
            if ((true && true) && (4294967295 != 95)) {
                fmt.Println(fi11, fi11)
            }
            if (0 < 18446744073709551615) {
                v10 = 952
                switch (fi11 >> ((686) & 63)) {
                case 882:
                    var v24 float32 = 0.0
                    _ = v24
                case -1:
                    var v25 float64 = 0.1
                    _ = v25
                case -558:
                    switch (fi11 + -859) {
                    case -9223372036854775808:
                        var v26 int8 = (v2 & (v2 << ((4294967295) & 63)))
                        _ = v26
                    case 833:
                        delete(v3, (^(v10)))
                    }
                }
                for fi27 := 0; fi27 < 1; fi27++ {
                    fmt.Println(v2, v10, v2)
                    fmt.Println(v10)
                    _ = fi27
                }
            }
        }
        _ = fi11
    }
    var v28 uint16 = 1
    if (true && (2 > uint64(v10))) {
        if false {
            if (false || (!(0 == 1))) {
                for fi29 := 0; fi29 < 17; fi29++ {
                    v2 -= v2
                    if false {
                        switch fi29 {
                        case 968:
                            var v30 uint64 = 18446744073709551615
                            _ = v30
                        case 9223372036854775807:
                            var v31 bool = true
                            _ = v31
                        }
                    }
                    _ = fi29
                }
                for fi32 := 0; fi32 < 1; fi32++ {
                    var v33 int8 = 100
                    v10 = 988
                    for fi34 := 0; fi34 < 7; fi34++ {
                        v35 := GenS3{F0: 42, F1: v28, F2: fi32}
                        _ = fi34
                        _ = v35
                    }
                    _ = fi32
                    _ = v33
                }
            } else {
                v36 := GenS4{F0: ("" + ("a" + "foo bar")), F1: 97}
                var v37 int64 = 1
                _ = v36
                _ = v37
            }
            _ = len(v3)
        }
        for fi38 := 0; fi38 < 12; fi38++ {
            var v39 int64 = 42
            for fi40 := 0; fi40 < 4; fi40++ {
                fmt.Println(fi40, v39, v39)
                switch (fi40 << ((2) & 63)) {
                case 552:
                    v1 = append(v1, ((("\n\t" + "a") + "\n\t") + "foo bar"))
                case 42:
                    switch ("" + "a") {
                    case "a":
                        if false {
                            if (!true) {
                                v1 = append(v1, (("" + "hello") + "a"))
                            } else {
                                var v41 int64 = (v39 - 15)
                                v42 := v3[v10]
                                _ = v41
                                _ = v42
                            }
                            var v43 float32 = 0.1
                            switch (("foo bar" + "") + ((("" + "\n\t") + ("foo bar" + (("foo bar" + "\n\t") + ("\n\t" + (((("hello" + (("\n\t" + "a") + "a")) + "a") + "foo bar") + "\n\t"))))) + "")) {
                            case "foo bar":
                                var v44 int64 = v39
                                _ = v44
                            }
                            _ = v43
                        } else {
                            v39 -= -1
                        }
                    case "hello":
                        if true {
                            fmt.Println(v2)
                            var v45 bool = false
                            fmt.Println(v45)
                            fmt.Println(v2, v39)
                            _ = v45
                        }
                    default:
                        var v46 uint32 = (v10 - (^(v10)))
                        _ = v46
                    }
                case -982:
                    v10 = v10
                }
                switch "hello" {
                case "\n\t":
                    var v47 int16 = 981
                    _ = v47
                case "hello":
                    var v48 int8 = (v2 ^ v2)
                    v3[4294967295] = fi38
                    _ = v48
                case "foo bar":
                    var v49 uint32 = 65535
                    _ = v49
                }
                _ = fi40
            }
            v2 = v2
            _ = fi38
            _ = v39
        }
    } else {
        v1 = append(v1, "")
        switch (("" + "foo bar") + ((("hello" + ("a" + ("" + ""))) + "") + "\n\t")) {
        case "":
            var v50 int = -1
            _ = v50
        case "foo bar":
            var v51 uint16 = v28
            _ = v51
        case "\n\t":
            v28 = v28
        default:
            v2 -= v2
        }
    }
    v52 := GenS3{F0: uint64(v28), F1: (v28 - v28), F2: 2147483647}
    switch int(v2) {
    case 128:
        var v53 uint16 = 65535
        _ = v53
    case 1:
        var v54 int16 = 0
        _ = v54
    default:
        v55 := GenS5{F0: "\n\t", F1: (42 == 0), F2: 42}
        _ = v55
    }
    for fi56 := 0; fi56 < 11; fi56++ {
        switch fi56 {
        case 4294967295:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -9223372036854775808:
            var v57 uint = 1
            _ = v57
        case 1:
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        default:
            for fi58 := 0; fi58 < 5; fi58++ {
                var v59 uint16 = 319
                _ = fi58
                _ = v59
            }
        }
        for fi60 := 0; fi60 < 18; fi60++ {
            v61 := GenS6{F0: "a", F1: false}
            _ = fi60
            _ = v61
        }
        _ = fi56
    }
    var v62 int64 = 373
    v62 = v62
    var v63 string = "hello"
    _ = len(v1)
    for fi64 := 0; fi64 < 18; fi64++ {
        v1 = append(v1, (((v63 + "a") + ("foo bar" + (v63 + (("" + v63) + "a")))) + ("foo bar" + (v63 + ("\n\t" + ((v63 + v63) + v63))))))
        _ = fi64
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v52
    _ = v6
    _ = v62
    _ = v63
}
