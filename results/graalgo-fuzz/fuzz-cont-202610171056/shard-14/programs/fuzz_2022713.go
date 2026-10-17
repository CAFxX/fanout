package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int32 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 int8; F1 int8; F2 bool }
type GenS8 struct { F0 uint64; F1 int8; F2 uint8 }

func main() {
    var v1 int16 = 42
    var v2 uint32 = 2
    var v3 int16 = v1
    v4 := make(map[string]bool)
    v5 := GenS1{F0: 0, F1: (v2 % (((v2 ^ (v2 * v2))) | 1)), F2: 1}
    var v6 uint32 = (v2 * 42)
    v6 += (v2 * 1)
    v7 := []int{}
    switch "a" {
    case "foo bar":
        for fi8 := 0; fi8 < 17; fi8++ {
            if false {
                v3 = (v3 >> ((2) & 63))
            }
            v9 := GenS1{F0: 485, F1: (v6 % (((v2 ^ v6)) | 1)), F2: 42}
            _ = fi8
            _ = v9
        }
    case "":
        v10 := GenS1{F0: 1, F1: (v6 % (((^(v2))) | 1)), F2: 42}
        _ = v10
    case "hello":
        v11 := GenS2{F0: 90, F1: false}
        _ = v11
    }
    _ = len(v7)
    _ = len(v4)
    var v12 uint16 = 255
    if true {
        var v13 int = 2147483647
        v14 := v4[""]
        v15 := GenS2{F0: -1, F1: ((v14 || v14) && (((v6 ^ (v2 | 4294967295)) > (v6 / ((v2) | 1))) && (!false)))}
        _ = v13
        _ = v14
        _ = v15
    } else {
        v16 := GenS1{F0: -1, F1: v2, F2: 124}
        switch 0 {
        case -124:
            switch ("" + "a") {
            case "a":
                v17 := v4["a"]
                _ = v17
            case "\n\t":
                v2 -= v2
            }
        }
        switch 830 {
        case -155:
            for fi18 := 0; fi18 < 19; fi18++ {
                v6 -= v2
                v6 = v2
                for fi19 := 0; fi19 < 8; fi19++ {
                    v1 = 0
                    switch "hello" {
                    case "a":
                        var v20 uint32 = 65535
                        _ = v20
                    }
                    v21 := GenS3{F0: fi18, F1: v2}
                    _ = fi19
                    _ = v21
                }
                _ = fi18
            }
        case 1:
            if ((v12 <= (v12 - 0)) || true) {
                v4["a"] = true
                fmt.Println(v1, v2, v2)
            } else {
                var v22 uint16 = v12
                var v23 int = 128
                switch 4294967295 {
                case -9223372036854775808:
                    v24 := GenS1{F0: -1, F1: (v2 >> ((4294967295) & 63)), F2: -1}
                    _ = v24
                }
                _ = v22
                _ = v23
            }
        case 2147483647:
            v3 = -345
        default:
            v3 += -1
        }
        _ = v16
    }
    switch -885 {
    case 32768:
        fmt.Println(v6, v3, v1)
    case 706:
        switch -1 {
        case 559:
            for fi25 := 0; fi25 < 20; fi25++ {
                v26 := GenS3{F0: fi25, F1: 344}
                v27 := GenS3{F0: (^(fi25)), F1: v2}
                _ = fi25
                _ = v26
                _ = v27
            }
        case 412:
            fmt.Println(v6)
        case 9223372036854775807:
            v28 := GenS4{F0: 32768, F1: 1}
            _ = v28
        default:
            switch ("\n\t" + ("" + (("foo bar" + "\n\t") + "\n\t"))) {
            case "foo bar":
                v3 += v1
            }
        }
    case 42:
        v4[""] = true
    }
    if true {
        var v29 int64 = -9223372036854775808
        if (v3 != v3) {
            v30 := GenS5{F0: -334}
            _ = v30
        } else {
            if len(v7) > 0 {
                v7[0] = 128
            }
            var v31 bool = (!false)
            _ = v31
        }
        _ = v29
    }
    v32 := v4["a"]
    if v32 {
        for fi33 := 0; fi33 < 5; fi33++ {
            v6 = v6
            var v34 uint16 = v12
            _ = fi33
            _ = v34
        }
    } else {
        var v35 int64 = 42
        fmt.Println(v32, v12)
        if v32 {
            var v36 int = 1
            _ = v36
        } else {
            v37 := GenS3{F0: 28, F1: (v6 / ((v2) | 1))}
            _ = v37
        }
        _ = v35
    }
    if (4294967295 >= 149) {
        var v38 int32 = 42
        _ = v38
    } else {
        for fi39 := 0; fi39 < 20; fi39++ {
            for fi40 := 0; fi40 < 6; fi40++ {
                var v41 int32 = 0
                switch fi40 {
                case 1:
                    v12 += (v12 & 490)
                default:
                    v6 -= 254
                }
                _ = fi40
                _ = v41
            }
            v1 = (v3 | 1)
            _ = fi39
        }
        var v42 int8 = 0
        _ = v42
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(v6, v12, v32)
    delete(v4, "")
    switch "a" {
    case "foo bar":
        if v32 {
            if (2 >= 0) {
                if false {
                    for fi43 := 0; fi43 < 6; fi43++ {
                        if v32 {
                            var v44 int16 = 1
                            var v45 int = (^(fi43))
                            _ = v44
                            _ = v45
                        } else {
                            if ((!(false || v32)) && v32) {
                                v46 := v4[("hello" + "hello")]
                                if (v32 && (("" + "hello") != ("" + ("a" + ("foo bar" + "hello"))))) {
                                    if false {
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        for fi47 := 0; fi47 < 17; fi47++ {
                                            v48 := GenS6{F0: true, F1: 453}
                                            _ = fi47
                                            _ = v48
                                        }
                                    }
                                } else {
                                    v6 = v2
                                    v49 := GenS7{F0: 127, F1: 0, F2: ((!v32) && v32)}
                                    _ = v49
                                }
                                _ = v46
                            }
                            v50 := v4["foo bar"]
                            _ = v50
                        }
                        if (v32 || v32) {
                            for fi51 := 0; fi51 < 18; fi51++ {
                                v52 := GenS8{F0: uint64(v12), F1: -91, F2: 128}
                                var v53 int16 = (v3 + v1)
                                if (v32 && ("\n\t" != ("" + (("" + "foo bar") + "")))) {
                                    v4[""] = v32
                                    fmt.Println(fi51)
                                } else {
                                    v54 := GenS7{F0: 1, F1: 127, F2: (!(!false))}
                                    v55 := GenS1{F0: -2147483648, F1: (^(v2)), F2: 174}
                                    _ = v54
                                    _ = v55
                                }
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                        }
                        _ = fi43
                    }
                }
                switch "a" {
                case "hello":
                    fmt.Println(v32, v2)
                }
            }
            v1 += (v1 + (v3 / ((1) | 1)))
        } else {
            var v56 uint32 = 65535
            _ = v56
        }
    case "hello":
        var v57 uint = 18446744073709551615
        _ = v57
    case "\n\t":
        for fi58 := 0; fi58 < 18; fi58++ {
            var v59 float64 = ((float64(0.1) - 1e10) / -2.25)
            var v60 float32 = -2.25
            _ = fi58
            _ = v59
            _ = v60
        }
    default:
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            var v61 uint16 = v12
            _ = v61
        case "":
            v62 := GenS3{F0: 0, F1: (v6 / (((^(v2))) | 1))}
            _ = v62
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
