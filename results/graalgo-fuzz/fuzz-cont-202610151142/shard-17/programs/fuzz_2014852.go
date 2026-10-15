package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 uint16; F1 int; F2 uint16 }
type GenS7 struct { F0 int32; F1 uint64; F2 uint16 }

func main() {
    var v1 int32 = -945
    v2 := []string{"foo bar", (("a" + "hello") + "foo bar"), "\n\t"}
    v3 := make(map[string]uint32)
    var v4 int8 = -31
    switch -1 {
    case 182:
        v5 := GenS1{F0: 0}
        _ = v5
    case -1:
        v1 = v1
    case 2147483647:
        for fi6 := 0; fi6 < 15; fi6++ {
            for fi7 := 0; fi7 < 14; fi7++ {
                switch (fi7 + fi6) {
                case -9223372036854775808:
                    v1 = -181
                case 32767:
                    v4 = 1
                }
                var v8 int64 = -863
                v9 := GenS1{F0: v1}
                _ = fi7
                _ = v8
                _ = v9
            }
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                for fi10 := 0; fi10 < 4; fi10++ {
                    delete(v3, ("a" + (("" + "foo bar") + ("" + ("a" + "hello")))))
                    _ = fi10
                }
            case "":
                fmt.Println(fi6, v4, v4)
            case "hello":
                for fi11 := 0; fi11 < 3; fi11++ {
                    var v12 float32 = 1.5
                    var v13 uint16 = 267
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            _ = fi6
        }
    }
    if (("hello" + (("hello" + "\n\t") + "\n\t")) >= "foo bar") {
        var v14 int8 = (v4 + v4)
        fmt.Println(v1, v1)
        _ = v14
    } else {
        var v15 uint64 = 2
        _ = v15
    }
    switch "\n\t" {
    case "\n\t":
        v4 = v4
    default:
        v16 := GenS1{F0: v1}
        _ = v16
    }
    var v17 string = ("a" + "")
    var v18 int = -850
    for fi19 := 0; fi19 < 1; fi19++ {
        if true {
            for fi20 := 0; fi20 < 7; fi20++ {
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                var v21 int16 = -1
                if true {
                    fmt.Println(fi20)
                    if false {
                        if (1 != 127) {
                            if true {
                                switch ("" + "a") {
                                case "foo bar":
                                    v4 = 42
                                case "":
                                    var v22 float32 = -2.25
                                    if ((1.5 / 1e10) <= float64((v22 * (1.5 * 1.5)))) {
                                        v23 := GenS1{F0: v1}
                                        _ = v23
                                    } else {
                                        v24 := GenS2{F0: 2147483648}
                                        v25 := GenS3{F0: 421, F1: (v21 << ((769) & 63))}
                                        _ = v24
                                        _ = v25
                                    }
                                    var v26 uint64 = 18446744073709551615
                                    _ = v22
                                    _ = v26
                                case "\n\t":
                                    v18 = v18
                                    v17 = v17
                                default:
                                    fmt.Println(v18, fi19)
                                }
                                if len(v2) > 0 {
                                    v2[0] = "\n\t"
                                }
                                if ((42 >= -969) && (v4 < (-(v4)))) {
                                    if (!false) {
                                        if ((0.0 <= (float32(1.5) + 1.5)) || false) {
                                            v17 = v17
                                        }
                                        v27 := GenS1{F0: (v1 / ((v1) | 1))}
                                        _ = v27
                                    }
                                    v2 = append(v2, v17)
                                    v21 -= v21
                                }
                            }
                        } else {
                            _ = len(v2)
                            switch fi19 {
                            case -68:
                                if (v17 > ("hello" + "")) {
                                    var v28 float32 = -2.25
                                    v4 += v4
                                    fmt.Println(v21)
                                    _ = v28
                                } else {
                                    var v29 uint16 = 292
                                    if false {
                                        var v30 int64 = -9223372036854775808
                                        _ = v30
                                    } else {
                                        delete(v3, v17)
                                        var v31 int = v18
                                        _ = v31
                                    }
                                    _ = v29
                                }
                            case 127:
                                fmt.Println(v18, v4)
                            }
                            v2 = append(v2, "hello")
                        }
                    } else {
                        var v32 uint = 65535
                        v4 -= (v4 / ((v4) | 1))
                        _ = v32
                    }
                    v33 := GenS1{F0: (v1 << ((1) & 63))}
                    _ = v33
                }
                _ = fi20
                _ = v21
            }
            v34 := GenS4{F0: "hello"}
            v4 = v4
            _ = v34
        }
        switch fi19 {
        case 890:
            if false {
                var v35 int64 = 165
                _ = v35
            } else {
                fmt.Println(v4)
            }
        }
        v17 = ((v17 + (("hello" + "foo bar") + v17)) + ((v17 + v17) + "a"))
        _ = fi19
    }
    switch "foo bar" {
    case "hello":
        v36 := GenS3{F0: 18446744073709551615, F1: -1}
        _ = v36
    case "a":
        v1 += (v1 << ((255) & 63))
    default:
        for fi37 := 0; fi37 < 12; fi37++ {
            for fi38 := 0; fi38 < 19; fi38++ {
                switch fi37 {
                case 128:
                    var v39 uint16 = 0
                    _ = v39
                case -2147483648:
                    fmt.Println(v4, v17, v18)
                case -231:
                    v40 := GenS2{F0: 2147483647}
                    _ = v40
                }
                _ = fi38
            }
            _ = fi37
        }
    }
    if true {
        switch v17 {
        case "\n\t":
            for fi41 := 0; fi41 < 5; fi41++ {
                if false {
                    v2 = append(v2, v17)
                } else {
                    if (0.1 <= -2.25) {
                        switch fi41 {
                        case 127:
                            var v42 int64 = -1
                            _ = v42
                        case 32768:
                            var v43 int = 128
                            _ = v43
                        }
                        v44 := GenS5{F0: 0, F1: -11}
                        v1 += 0
                        _ = v44
                    } else {
                        var v45 int = (v18 << ((4294967295) & 63))
                        _ = v45
                    }
                    switch v17 {
                    case "a":
                        if (false && false) {
                            fmt.Println(v18, fi41, v4)
                            v4 = (v4 + int8(v18))
                            for fi46 := 0; fi46 < 6; fi46++ {
                                switch 748 {
                                case 0:
                                    v2 = append(v2, "a")
                                case 32768:
                                    v47 := GenS4{F0: "foo bar"}
                                    _ = v47
                                }
                                if false {
                                    var v48 uint64 = 708
                                    _ = v48
                                }
                                v1 += int32(v18)
                                _ = fi46
                            }
                        } else {
                            var v49 uint8 = 127
                            var v50 uint = 4294967295
                            v17 = (v17 + (v17 + v17))
                            _ = v49
                            _ = v50
                        }
                    case "foo bar":
                        switch (-(v18)) {
                        case 2147483647:
                            v1 -= v1
                        }
                    case "hello":
                        v51 := v3[(v17 + ("\n\t" + v17))]
                        _ = v51
                    }
                    v3["hello"] = 471
                }
                switch -66 {
                case -9223372036854775808:
                    switch (v18 + fi41) {
                    case -1:
                        if len(v2) > 0 {
                            v2[0] = v17
                        }
                    default:
                        fmt.Println(v1, v1, v18)
                    }
                case 2147483647:
                    var v52 int32 = (v1 % ((v1) | 1))
                    _ = v52
                case 4294967295:
                    if (!false) {
                        v53 := GenS6{F0: 255, F1: 127, F2: 32767}
                        v4 = (-(v4))
                        v54 := GenS7{F0: v1, F1: 48, F2: 924}
                        _ = v53
                        _ = v54
                    }
                }
                _ = fi41
            }
        case "hello":
            var v55 int16 = 42
            _ = v55
        default:
            _ = len(v2)
        }
        var v56 uint16 = 255
        _ = v56
    } else {
        var v57 int = v18
        v4 += (v4 << ((255) & 63))
        _ = v57
    }
    v58 := v3[("foo bar" + "hello")]
    for fi59 := 0; fi59 < 13; fi59++ {
        if false {
            v18 += (fi59 - v18)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch "\n\t" {
            case "foo bar":
                delete(v3, v17)
            case "\n\t":
                _ = len(v2)
            case "":
                var v60 uint32 = v58
                _ = v60
            }
        } else {
            var v61 uint64 = 793
            _ = v61
        }
        v17 = v17
        var v62 int16 = -320
        _ = fi59
        _ = v62
    }
    fmt.Println(v4)
    var v63 uint64 = 304
    if true {
        switch (v18 % ((v18) | 1)) {
        case 127:
            v2 = append(v2, v17)
        }
        var v64 float32 = 1.5
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v58:", v58)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v58
    _ = v63
}
