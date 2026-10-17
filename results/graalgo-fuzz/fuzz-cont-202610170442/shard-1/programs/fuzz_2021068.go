package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 string; F2 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int16 = 1000
    v2 := []string{"", "foo bar", "hello"}
    var v3 int16 = 42
    v4 := make(map[string]uint32)
    delete(v4, "a")
    var v5 uint = 1
    switch "hello" {
    case "hello":
        switch ("a" + ("\n\t" + "\n\t")) {
        case "a":
            switch 127 {
            case -746:
                v3 = (v3 - (v1 & -32768))
            }
        case "":
            fmt.Println(v5, v1, v3)
        case "foo bar":
            if (1 > -114) {
                for fi6 := 0; fi6 < 5; fi6++ {
                    if (801 > v5) {
                        v7 := GenS1{F0: 128, F1: 9223372036854775807}
                        _ = v7
                    }
                    if (fi6 < (fi6 & fi6)) {
                        if false {
                            v5 -= (v5 + v5)
                            for fi8 := 0; fi8 < 13; fi8++ {
                                v9 := GenS2{F0: "a"}
                                _ = fi8
                                _ = v9
                            }
                        }
                    } else {
                        if (42 == 42) {
                            for fi10 := 0; fi10 < 5; fi10++ {
                                var v11 uint = v5
                                _ = fi10
                                _ = v11
                            }
                            v5 -= v5
                        } else {
                            var v12 string = "hello"
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 12; fi13++ {
                            switch ("foo bar" + "hello") {
                            case "a":
                                fmt.Println(v3, v3)
                            case "foo bar":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            v1 += v3
                            _ = fi13
                        }
                    }
                    _ = fi6
                }
                v1 = -1
                for fi14 := 0; fi14 < 3; fi14++ {
                    v5 = (v5 << ((v5) & 63))
                    for fi15 := 0; fi15 < 2; fi15++ {
                        var v16 string = ("foo bar" + "foo bar")
                        v5 -= v5
                        var v17 float32 = 0.1
                        v5 = (v5 / (((v5 ^ v5)) | 1))
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
            } else {
                v5 -= v5
                delete(v4, ("hello" + "\n\t"))
            }
        default:
            v1 -= (v1 >> (((v5 % ((1) | 1))) & 63))
        }
    case "":
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    default:
        for fi18 := 0; fi18 < 11; fi18++ {
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 uint8 = 2
                if (42 < 620) {
                    v21 := GenS1{F0: v20, F1: 0}
                    _ = len(v2)
                    _ = v21
                }
                switch ("foo bar" + ((("a" + "\n\t") + "hello") + "a")) {
                case "hello":
                    switch ("" + "a") {
                    case "":
                        v3 -= v1
                    case "foo bar":
                        _ = len(v4)
                    }
                case "\n\t":
                    if (-30 == 42) {
                        if (1 < 174) {
                            v20 = 1
                            if ((v5 | (v5 & (v5 & (v5 / ((1) | 1))))) == v5) {
                                v22 := GenS1{F0: v20, F1: (^(fi19))}
                                switch "\n\t" {
                                case "a":
                                    v5 -= (^(v5))
                                case "\n\t":
                                    v5 = (v5 / ((v5) | 1))
                                case "foo bar":
                                    var v23 int16 = 1
                                    _ = v23
                                }
                                if (uint64(v20) != 42) {
                                    v24 := v4["hello"]
                                    var v25 int8 = 0
                                    v1 += (v1 % ((-897) | 1))
                                    _ = v24
                                    _ = v25
                                }
                                _ = v22
                            } else {
                                var v26 int64 = 161
                                _ = len(v2)
                                v27 := v4[(((("a" + ("a" + ("foo bar" + "foo bar"))) + (("\n\t" + ("foo bar" + "")) + "foo bar")) + (("foo bar" + "foo bar") + ("" + "hello"))) + "a")]
                                v28 := GenS1{F0: v20, F1: (fi19 - (fi18 << ((v5) & 63)))}
                                v29 := GenS1{F0: 15, F1: fi19}
                                var v30 bool = (!(!(100 == 127)))
                                var v31 string = "a"
                                _ = v26
                                _ = v27
                                _ = v28
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        } else {
                            if (v20 < v20) {
                                var v32 int64 = 0
                                _ = v32
                            }
                            v4["a"] = 65535
                            v33 := GenS3{F0: 1, F1: ("foo bar" + ("foo bar" + ("foo bar" + ("a" + ("\n\t" + "\n\t"))))), F2: false}
                            _ = v33
                        }
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v34 int16 = int16(fi19)
                        var v35 int = 2147483647
                        _ = v34
                        _ = v35
                    }
                case "a":
                    if (true && true) {
                        v36 := GenS4{F0: -1}
                        _ = v36
                    } else {
                        v20 -= (v20 << ((v5) & 63))
                        var v37 int64 = -9223372036854775808
                        v37 -= v37
                        _ = v37
                    }
                }
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    }
    switch "a" {
    case "a":
        if true {
            v4["a"] = 788
        }
    case "foo bar":
        var v38 uint8 = 255
        _ = v38
    case "":
        if (42 <= 532) {
            switch 9223372036854775807 {
            case 592:
                _ = len(v4)
            }
            v39 := GenS3{F0: 42, F1: "", F2: (true || false)}
            var v40 uint64 = 549
            _ = v39
            _ = v40
        }
    }
    var v41 bool = false
    v5 += v5
    v42 := GenS4{F0: 9223372036854775807}
    for fi43 := 0; fi43 < 18; fi43++ {
        v1 += v1
        _ = fi43
    }
    if (uint32(v5) > 1) {
        fmt.Println(v5, v41, v3)
    } else {
        if false {
            v44 := GenS5{F0: 32767, F1: 0}
            fmt.Println(v1)
            _ = v44
        } else {
            v45 := v4["a"]
            fmt.Println(v5, v3, v3)
            _ = v45
        }
        var v46 uint16 = 0
        if (-2.25 != (0.0 * -2.25)) {
            var v47 uint8 = 92
            _ = v47
        } else {
            v1 = (v3 | v1)
            v48 := GenS6{F0: 9223372036854775807, F1: (v46 & (v46 * 1))}
            var v49 uint32 = 2147483647
            _ = v48
            _ = v49
        }
        _ = v46
    }
    switch 315 {
    case 1:
        v4[""] = 375
    case 32767:
        v1 += -517
    default:
        if (42 > 0) {
            var v50 int16 = v3
            v41 = (2147483647 <= 0)
            _ = v50
        }
    }
    if v41 {
        if false {
            v3 -= (v3 + v1)
            var v51 int8 = -1
            _ = v51
        }
    } else {
        switch ("" + ("hello" + "a")) {
        case "\n\t":
            var v52 int32 = 1
            _ = v52
        }
    }
    switch "hello" {
    case "":
        _ = len(v2)
    }
    switch 42 {
    case -907:
        switch 604 {
        case -29:
            for fi53 := 0; fi53 < 11; fi53++ {
                v54 := v4["a"]
                for fi55 := 0; fi55 < 9; fi55++ {
                    if (9223372036854775807 == -9223372036854775808) {
                        switch "foo bar" {
                        case "\n\t":
                            delete(v4, "hello")
                        }
                        _ = len(v4)
                        var v56 uint8 = 187
                        _ = v56
                    }
                    v2 = append(v2, "foo bar")
                    _ = fi55
                }
                _ = fi53
                _ = v54
            }
        }
    case 42:
        v2 = append(v2, "hello")
    }
    v57 := v4["a"]
    fmt.Println(v5, v5, v41)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v57
}
