package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
type GenS2 struct { F0 int8; F1 int; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int64; F2 bool }
type GenS4 struct { F0 int64; F1 string; F2 int16 }
type GenS5 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int }
func genFn0(p0 string, p1 uint16) uint64 {
    p0 = p0
    for fi7 := 0; fi7 < 1; fi7++ {
        p1 = p1
        var v8 bool = true
        p1 = p1
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        if true {
            switch "hello" {
            case "hello":
                v10 := GenS1{F0: (65535 == 65535), F1: -319, F2: -256}
                _ = v10
            case "foo bar":
                fmt.Println(p0, p1, p1)
            case "a":
                for fi11 := 0; fi11 < 1; fi11++ {
                    p1 += (p1 | 2)
                    var v12 int16 = 0
                    if true {
                        var v13 uint64 = 42
                        p0 = p0
                        var v14 uint32 = 42
                        _ = v13
                        _ = v14
                    }
                    _ = fi11
                    _ = v12
                }
            }
            var v15 int = fi9
            for fi16 := 0; fi16 < 5; fi16++ {
                v17 := []int{fi9}
                if len(v17) > 0 {
                    v17[0] = 2147483647
                }
                _ = fi16
                _ = v17
            }
            _ = v15
        } else {
            fmt.Println(p0, fi9, fi9)
            v18 := []int{-2147483648, 127, (fi9 / ((32768) | 1))}
            _ = v18
        }
        v19 := GenS1{F0: false, F1: -656, F2: -32768}
        _ = fi9
        _ = v19
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 uint16) bool {
    v20 := []int{760, -9223372036854775808, 816}
    if true {
        switch 32768 {
        case 704:
            var v21 uint16 = p0
            _ = v21
        }
        p0 = 140
        v20 = append(v20, 42)
    } else {
        if true {
            var v22 uint = 4294967295
            v22 += 65535
            switch ("\n\t" + ("foo bar" + ("a" + "a"))) {
            case "\n\t":
                v23 := GenS1{F0: true, F1: 42, F2: 0}
                _ = v23
            case "":
                p0 = (p0 - (p0 << ((713) & 63)))
            default:
                p0 += p0
            }
            _ = v22
        }
        p0 -= (p0 / ((p0) | 1))
    }
    fmt.Println(p0)
    p0 = (p0 << ((255) & 63))
    _ = len(v20)
    _ = p0
    _ = v20
    return (-1 > -32768)
}

func main() {
    var v1 int = 0
    v2 := []string{("\n\t" + "a"), "a"}
    v3 := []int{32768, -35}
    var v4 int = v1
    v5 := []int{(-(v1)), -9223372036854775808, (v1 ^ (v4 & v4))}
    var v6 uint8 = 2
    fmt.Println(v4, v6)
    v4 = (v1 >> ((4294967295) & 63))
    var v24 uint16 = 255
    for fi25 := 0; fi25 < 1; fi25++ {
        switch fi25 {
        case -1:
            fmt.Println(v4, v24)
        }
        switch "" {
        case "foo bar":
            v1 = v4
        case "\n\t":
            switch "a" {
            case "foo bar":
                if (float32(0.1) != ((0.1 * ((1e10 * (-2.25 + -2.25)) / 3.14159)) + 0.1)) {
                    var v26 int8 = 42
                    var v27 uint8 = (v6 / ((v6) | 1))
                    _ = v26
                    _ = v27
                }
            case "":
                for fi28 := 0; fi28 < 8; fi28++ {
                    fmt.Println(v4)
                    switch (v4 - (v1 | 32767)) {
                    case 74:
                        v1 = -2147483648
                    case 564:
                        v1 -= (v1 & fi25)
                    case -773:
                        var v29 float32 = 1e10
                        _ = v29
                    }
                    _ = fi28
                }
            case "a":
                for fi30 := 0; fi30 < 12; fi30++ {
                    var v31 uint16 = 255
                    var v32 int = (-(v1))
                    if false {
                        v32 = -176
                    } else {
                        v33 := GenS1{F0: (4294967295 > 65535), F1: 0, F2: 0}
                        var v34 int32 = -408
                        _ = v33
                        _ = v34
                    }
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            default:
                if (!true) {
                    v1 = fi25
                    var v35 bool = (false && (817 >= genFn0("hello", uint16(v6))))
                    _ = v35
                }
            }
        case "":
            if true {
                v1 -= (v4 / (((v1 / ((v4) | 1))) | 1))
                switch "foo bar" {
                case "\n\t":
                    v6 = uint8(v24)
                case "hello":
                    v36 := GenS1{F0: true, F1: 471, F2: 42}
                    _ = v36
                case "":
                    v37 := GenS1{F0: (!true), F1: 2147483647, F2: 42}
                    _ = v37
                default:
                    var v38 uint32 = 42
                    _ = v38
                }
            }
        }
        _ = fi25
    }
    switch "" {
    case "foo bar":
        var v39 uint = 833
        _ = v39
    case "\n\t":
        var v40 string = "hello"
        _ = v40
    case "a":
        v41 := GenS2{F0: 100, F1: v4, F2: (v6 % (((^(v6))) | 1))}
        _ = v41
    default:
        for fi42 := 0; fi42 < 17; fi42++ {
            var v43 int64 = -422
            _ = fi42
            _ = v43
        }
    }
    switch "hello" {
    case "foo bar":
        v44 := GenS3{F0: 2, F1: 1, F2: true}
        _ = v44
    case "a":
        v4 -= (v1 + v4)
    case "\n\t":
        fmt.Println(v6)
    }
    var v45 uint = 4294967295
    if true {
        v1 += 9223372036854775807
    }
    var v46 uint32 = 0
    v47 := GenS1{F0: false, F1: 1, F2: 1000}
    fmt.Println(v1)
    switch v4 {
    case 127:
        v48 := GenS2{F0: 42, F1: v1, F2: v6}
        _ = v48
    default:
        var v49 int8 = -105
        _ = v49
    }
    if true {
        v50 := GenS4{F0: 42, F1: "", F2: int16(v45)}
        fmt.Println(v6, v46, v4)
        for fi51 := 0; fi51 < 14; fi51++ {
            if false {
                switch "a" {
                case "foo bar":
                    v52 := GenS1{F0: false, F1: -2147483648, F2: 1}
                    _ = v52
                case "\n\t":
                    switch "a" {
                    case "\n\t":
                        v45 += v45
                    }
                case "hello":
                    v24 = v24
                }
            }
            _ = fi51
        }
        _ = v50
    } else {
        if ((42 == 0) && (false && true)) {
            switch v4 {
            case 1:
                var v53 int32 = -1
                _ = v53
            case 127:
                if len(v2) > 0 {
                    v2[0] = (("foo bar" + "") + "foo bar")
                }
            }
        }
        var v54 int32 = int32(v6)
        _ = v54
    }
    if (!((!true) && (0 <= v6))) {
        switch 127 {
        case -2147483648:
            var v55 float32 = 1.5
            _ = v55
        }
        switch ("a" + "") {
        case "":
            var v56 bool = (643 < 2147483647)
            _ = v56
        case "hello":
            var v57 uint8 = (v6 % ((v6) | 1))
            _ = v57
        case "foo bar":
            fmt.Println(v4)
        }
    } else {
        var v58 int16 = 42
        var v59 float64 = -2.25
        _ = v58
        _ = v59
    }
    for fi60 := 0; fi60 < 1; fi60++ {
        switch "\n\t" {
        case "\n\t":
            var v61 uint16 = 65535
            _ = v61
        case "foo bar":
            if false {
                if true {
                    var v62 float64 = 0.0
                    v63 := GenS5{F0: 119, F1: 2, F2: v46}
                    v45 = 0
                    _ = v62
                    _ = v63
                }
                for fi64 := 0; fi64 < 7; fi64++ {
                    switch v4 {
                    case -121:
                        var v65 int64 = -234
                        _ = v65
                    case 32767:
                        _ = len(v5)
                    case 42:
                        v3 = append(v3, v4)
                    default:
                        v5 = append(v5, 32768)
                    }
                    _ = fi64
                }
            }
        case "a":
            if ((1.5 < 1e10) && true) {
                if (!false) {
                    if (!true) {
                        v4 = (fi60 / ((v1) | 1))
                        for fi66 := 0; fi66 < 13; fi66++ {
                            fmt.Println(v1, v24, fi66)
                            v6 = v6
                            v4 = (fi60 << ((4294967295) & 63))
                            v4 = (v4 % (((v1 - v4)) | 1))
                            v24 = uint16(v45)
                            _ = fi66
                        }
                    } else {
                        v45 = (v45 | (^(v45)))
                        if len(v3) > 0 {
                            v3[0] = -2147483648
                        }
                        v2 = append(v2, ("hello" + "a"))
                    }
                    for fi67 := 0; fi67 < 20; fi67++ {
                        if true {
                            v4 = (fi60 + int(v45))
                            v4 -= v1
                        } else {
                            var v68 uint32 = 4294967295
                            _ = v68
                        }
                        v2 = append(v2, "")
                        v69 := GenS4{F0: 42, F1: ((((("" + "") + "a") + "a") + "\n\t") + ("" + "foo bar")), F2: 252}
                        fmt.Println(fi67, v45)
                        _ = fi67
                        _ = v69
                    }
                }
                var v70 int = 362
                _ = v70
            }
        }
        _ = fi60
    }
    var v71 float64 = 1.5
    if false {
        var v72 float64 = v71
        var v73 bool = (v24 <= (v24 << ((v45) & 63)))
        switch v1 {
        case 4294967295:
            if ((v24 ^ 2) < v24) {
                var v74 int64 = 1
                var v75 bool = (false && (false && ((v74 << (((v45 ^ (v45 % ((v45) | 1)))) & 63)) != 0)))
                if (v24 > uint16(v46)) {
                    v46 = v46
                } else {
                    fmt.Println(v6)
                    fmt.Println(v75, v75, v4)
                }
                _ = v74
                _ = v75
            }
        case 429:
            if (0 <= 42) {
                switch ("hello" + ("a" + "\n\t")) {
                case "foo bar":
                    var v76 int32 = 42
                    _ = v76
                }
                var v77 float64 = -2.25
                _ = v77
            } else {
                var v78 int = 9223372036854775807
                _ = v78
            }
        default:
            v1 = 128
        }
        _ = v72
        _ = v73
    } else {
        v24 = uint16(v46)
        var v79 uint16 = v24
        var v80 string = ""
        _ = v79
        _ = v80
    }
    v81 := GenS6{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v71
    _ = v81
}
