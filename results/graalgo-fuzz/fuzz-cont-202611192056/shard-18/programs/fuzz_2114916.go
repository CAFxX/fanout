package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 bool; F1 int64; F2 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 int16; F1 int64; F2 int64 }
type GenS9 struct { F0 uint16; F1 bool }

func main() {
    var v1 uint = 576
    v2 := make(map[string]string)
    var v3 int64 = -404
    var v4 int32 = 206
    var v5 int8 = 1
    v6 := []int{32767}
    fmt.Println(v4)
    if true {
        if true {
            switch "foo bar" {
            case "":
                var v7 int16 = 73
                _ = v7
            case "\n\t":
                var v8 int32 = (v4 | v4)
                _ = v8
            default:
                var v9 uint64 = 2
                _ = v9
            }
            v10 := v2[""]
            _ = v10
        } else {
            if false {
                fmt.Println(v3, v4, v1)
                switch "foo bar" {
                case "foo bar":
                    for fi11 := 0; fi11 < 16; fi11++ {
                        v5 -= -128
                        v12 := GenS1{F0: 1000, F1: 42}
                        _ = fi11
                        _ = v12
                    }
                case "a":
                    var v13 uint16 = 60
                    _ = v13
                case "\n\t":
                    var v14 uint16 = 65535
                    _ = v14
                }
            } else {
                for fi15 := 0; fi15 < 14; fi15++ {
                    for fi16 := 0; fi16 < 13; fi16++ {
                        v17 := GenS2{F0: (float64(0.1) > 1.5), F1: (-(v3)), F2: -338}
                        var v18 string = (("foo bar" + ("" + "\n\t")) + "a")
                        v4 -= (v4 + (v4 | v4))
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    _ = fi15
                }
                v5 -= 42
            }
            v19 := GenS3{F0: (v1 | (v1 >> ((1) & 63)))}
            _ = v19
        }
        var v20 int = 9223372036854775807
        _ = v20
    } else {
        var v21 uint64 = 0
        v5 += 0
        delete(v2, ("\n\t" + ("" + "foo bar")))
        _ = v21
    }
    v22 := GenS3{F0: v1}
    for fi23 := 0; fi23 < 14; fi23++ {
        fmt.Println(fi23)
        _ = fi23
    }
    if (true || false) {
        for fi24 := 0; fi24 < 13; fi24++ {
            switch ((("\n\t" + ("foo bar" + "a")) + "a") + ("\n\t" + "foo bar")) {
            case "\n\t":
                v25 := GenS1{F0: 1000, F1: 20}
                _ = v25
            }
            switch "\n\t" {
            case "hello":
                fmt.Println(fi24)
            case "\n\t":
                if (uint16(v1) == 2) {
                    var v26 int32 = v4
                    _ = v26
                }
            case "a":
                var v27 int8 = v5
                _ = v27
            default:
                var v28 uint = 36
                _ = v28
            }
            delete(v2, "")
            _ = fi24
        }
        v4 = v4
    }
    var v29 uint = v1
    fmt.Println(v3, v1, v4)
    var v30 uint32 = 1
    v30 = (^(v30))
    v30 = (v30 / (((v30 ^ (v30 & (v30 * 2147483647)))) | 1))
    for fi31 := 0; fi31 < 4; fi31++ {
        for fi32 := 0; fi32 < 8; fi32++ {
            fmt.Println(v1, v3)
            if (0.0 < -2.25) {
                fmt.Println(v5, v3)
            }
            _ = fi32
        }
        var v33 uint = (v1 & (v29 & v1))
        switch (^(fi31)) {
        case 127:
            for fi34 := 0; fi34 < 7; fi34++ {
                v4 = (^(v4))
                v35 := GenS4{F0: 658}
                fmt.Println(v29)
                _ = fi34
                _ = v35
            }
        case -1:
            fmt.Println(v1)
        default:
            var v36 uint64 = 757
            _ = v36
        }
        _ = fi31
        _ = v33
    }
    v4 += 184
    if false {
        v30 = v30
        v4 -= v4
        v2[""] = ("\n\t" + "")
    } else {
        fmt.Println(v5)
    }
    switch -2147483648 {
    case -2147483648:
        if (!false) {
            if (v1 >= (v29 ^ (v1 / ((v1) | 1)))) {
                for fi37 := 0; fi37 < 18; fi37++ {
                    switch "" {
                    case "":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    case "\n\t":
                        var v38 int16 = 1
                        _ = v38
                    case "hello":
                        v39 := GenS1{F0: -32768, F1: 1}
                        _ = v39
                    default:
                        var v40 int16 = 42
                        _ = v40
                    }
                    _ = fi37
                }
                _ = len(v6)
            } else {
                v3 += -894
            }
        } else {
            switch 128 {
            case -14:
                var v41 int16 = 980
                _ = v41
            case 1:
                if (!(-2.25 > (3.14159 + 1.5))) {
                    switch 546 {
                    case 128:
                        v42 := GenS4{F0: 32767}
                        _ = v42
                    case 42:
                        _ = len(v2)
                    case -2147483648:
                        fmt.Println(v30, v1, v5)
                    }
                    v2[(("foo bar" + "hello") + "a")] = ("a" + "")
                } else {
                    var v43 uint8 = 208
                    v44 := GenS4{F0: uint16(v1)}
                    _ = v43
                    _ = v44
                }
            default:
                var v45 int64 = v3
                _ = v45
            }
            switch "" {
            case "hello":
                fmt.Println(v3, v4, v29)
            case "a":
                if (((float64((0.0 * 1.5)) / 0.1) > 1e10) && (!false)) {
                    v46 := GenS5{F0: -1}
                    _ = v46
                } else {
                    var v47 string = ""
                    var v48 bool = true
                    if v48 {
                        v5 -= v5
                        v3 -= (^(v3))
                    } else {
                        fmt.Println(v5, v47)
                        v49 := GenS3{F0: v1}
                        _ = v49
                    }
                    _ = v47
                    _ = v48
                }
            }
        }
    case 32767:
        v1 += 431
    case 4294967295:
        if (true || true) {
            for fi50 := 0; fi50 < 3; fi50++ {
                switch (("foo bar" + "a") + "foo bar") {
                case "":
                    fmt.Println(v4)
                case "foo bar":
                    for fi51 := 0; fi51 < 20; fi51++ {
                        v29 += v1
                        fmt.Println(fi50, v29)
                        _ = fi51
                    }
                case "hello":
                    var v52 float32 = 0.1
                    _ = v52
                }
                for fi53 := 0; fi53 < 18; fi53++ {
                    switch fi53 {
                    case 0:
                        v2[(("foo bar" + "\n\t") + "\n\t")] = "foo bar"
                    case 128:
                        v4 += 2147483647
                    }
                    v4 -= (v4 - 76)
                    _ = fi53
                }
                switch 9223372036854775807 {
                case 32767:
                    if len(v6) > 0 {
                        v6[0] = (fi50 / (((fi50 << (((v29 * 42)) & 63))) | 1))
                    }
                case 9223372036854775807:
                    for fi54 := 0; fi54 < 11; fi54++ {
                        v29 = (^(v1))
                        switch -9223372036854775808 {
                        case 2147483647:
                            if false {
                                switch "" {
                                case "\n\t":
                                    var v55 int8 = 15
                                    v56 := GenS4{F0: 0}
                                    _ = v55
                                    _ = v56
                                }
                            } else {
                                if (2147483647 == (v4 << ((245) & 63))) {
                                    var v57 uint32 = (^(v30))
                                    var v58 int64 = (v3 | int64(v29))
                                    v30 -= 2147483648
                                    _ = v57
                                    _ = v58
                                }
                                v1 -= uint(v30)
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                        case 127:
                            fmt.Println(fi50)
                        case 128:
                            v5 = v5
                        }
                        _ = len(v2)
                        v59 := GenS6{F0: (^(v29)), F1: v3}
                        if true {
                            var v60 uint = (^(v1))
                            v61 := GenS4{F0: 0}
                            _ = v60
                            _ = v61
                        }
                        _ = fi54
                        _ = v59
                    }
                case 128:
                    var v62 int64 = v3
                    _ = v62
                }
                _ = fi50
            }
        } else {
            for fi63 := 0; fi63 < 5; fi63++ {
                var v64 string = ("" + "")
                v30 += 4294967295
                var v65 float32 = (3.14159 / 1e10)
                _ = fi63
                _ = v64
                _ = v65
            }
            v66 := GenS7{F0: 12, F1: 255}
            v2[("hello" + "")] = "a"
            _ = v66
        }
    default:
        v5 -= v5
    }
    _ = len(v2)
    if ((1.5 - -2.25) != 1e10) {
        for fi67 := 0; fi67 < 2; fi67++ {
            v3 = (^(v3))
            _ = fi67
        }
        if (-690 == 1000) {
            var v68 uint8 = 0
            for fi69 := 0; fi69 < 5; fi69++ {
                v70 := GenS8{F0: 456, F1: v3, F2: v3}
                _ = fi69
                _ = v70
            }
            var v71 bool = (!true)
            _ = v68
            _ = v71
        } else {
            v72 := GenS9{F0: uint16(v30), F1: true}
            _ = v72
        }
        v4 -= v4
    } else {
        v73 := GenS2{F0: (true || (!(!true))), F1: 0, F2: 42}
        _ = len(v6)
        v74 := GenS7{F0: 255, F1: 0}
        _ = v73
        _ = v74
    }
    var v75 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v75
}
