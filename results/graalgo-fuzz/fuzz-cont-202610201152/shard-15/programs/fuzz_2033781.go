package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 bool; F2 int8 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 uint8; F2 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS9 struct { F0 int64; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 uint64 = 2
    v4 := []int{32768, 504, 1}
    switch "hello" {
    case "a":
        v5 := GenS1{F0: v3}
        _ = v5
    default:
        for fi6 := 0; fi6 < 2; fi6++ {
            v3 -= 1
            var v7 string = "foo bar"
            _ = fi6
            _ = v7
        }
    }
    v3 = v3
    if (-922 != 435) {
        v8 := GenS1{F0: v3}
        var v9 uint64 = 1
        _ = v8
        _ = v9
    } else {
        var v10 uint64 = v3
        _ = v10
    }
    var v11 uint8 = 42
    v11 = v11
    if (4294967295 != 42) {
        if true {
            var v12 float64 = (3.14159 * 1.5)
            _ = v12
        } else {
            switch -424 {
            case -419:
                switch ("\n\t" + "hello") {
                case "":
                    fmt.Println(v11, v11)
                case "foo bar":
                    _ = len(v4)
                case "a":
                    var v13 uint8 = 166
                    _ = v13
                }
            case -2147483648:
                if (!(v11 < (v11 - (v11 * (v11 / (((v11 % ((v11) | 1))) | 1)))))) {
                    switch ("hello" + "hello") {
                    case "hello":
                        v11 = v11
                    case "a":
                        switch 127 {
                        case 42:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    case "\n\t":
                        v3 = (v3 >> ((4294967295) & 63))
                    default:
                        var v14 uint8 = (^(v11))
                        _ = v14
                    }
                }
            case 32767:
                var v15 uint = 65535
                _ = v15
            }
            switch "hello" {
            case "hello":
                switch 0 {
                case 127:
                    switch ("" + ("hello" + "a")) {
                    case "hello":
                        v3 = 87
                    case "foo bar":
                        for fi16 := 0; fi16 < 18; fi16++ {
                            if len(v4) > 0 {
                                v4[0] = fi16
                            }
                            _ = fi16
                        }
                    case "\n\t":
                        v2[("\n\t" + ("\n\t" + "foo bar"))] = false
                    default:
                        if false {
                            fmt.Println(v11)
                            switch "foo bar" {
                            case "foo bar":
                                v17 := GenS1{F0: 771}
                                _ = v17
                            }
                        }
                    }
                case 346:
                    for fi18 := 0; fi18 < 8; fi18++ {
                        for fi19 := 0; fi19 < 3; fi19++ {
                            var v20 float32 = (0.1 - ((1.5 + -2.25) * 1.5))
                            v21 := GenS1{F0: v3}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        v22 := GenS2{F0: -863}
                        _ = fi18
                        _ = v22
                    }
                case 1:
                    if ((255 >= 191) && false) {
                        if (!true) {
                            fmt.Println(v3, v11)
                            var v23 int16 = 284
                            for fi24 := 0; fi24 < 17; fi24++ {
                                v23 += 953
                                if false {
                                    if true {
                                        for fi25 := 0; fi25 < 2; fi25++ {
                                            if (!true) {
                                                if false {
                                                    var v26 string = "\n\t"
                                                    v23 = (v23 << ((137) & 63))
                                                    _ = v26
                                                }
                                                switch fi24 {
                                                case 9223372036854775807:
                                                    fmt.Println(fi24)
                                                default:
                                                    var v27 uint16 = 138
                                                    _ = v27
                                                }
                                                v23 = int16(v3)
                                                var v28 string = ""
                                                switch fi25 {
                                                case -352:
                                                    var v29 int16 = (-(v23))
                                                    _ = v29
                                                case -2147483648:
                                                    v4 = append(v4, fi24)
                                                }
                                                _ = v28
                                            } else {
                                                v30 := GenS3{F0: (fi25 ^ fi25), F1: (v23 <= v23), F2: -79}
                                                if (1.5 >= 0.0) {
                                                    if (734 > 65535) {
                                                        v31 := GenS2{F0: (v23 >> ((uint(v11)) & 63))}
                                                        _ = v31
                                                    } else {
                                                        if true {
                                                            var v32 uint64 = v3
                                                            var v33 uint8 = v11
                                                            v34 := GenS1{F0: v32}
                                                            _ = v32
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                        switch (fi25 / (((fi24 ^ 127)) | 1)) {
                                                        case 128:
                                                            v23 += (v23 % ((32767) | 1))
                                                        case 32767:
                                                            if true {
                                                                var v35 uint16 = 255
                                                                _ = v35
                                                            }
                                                        }
                                                        fmt.Println(v11, v23)
                                                        if false {
                                                            fmt.Println(v23, fi25)
                                                            if true {
                                                                var v36 int32 = -2147483648
                                                                fmt.Println(fi25, fi24, fi25)
                                                                switch fi24 {
                                                                case 128:
                                                                    var v37 int16 = (v23 ^ (v23 | (v23 / (((v23 & (v23 >> ((822) & 63)))) | 1))))
                                                                    _ = v37
                                                                case 32768:
                                                                    var v38 uint16 = 2
                                                                    _ = v38
                                                                case 681:
                                                                    v39 := v1[64]
                                                                    _ = v39
                                                                default:
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                                _ = v36
                                                            } else {
                                                                switch "hello" {
                                                                case "":
                                                                    v40 := GenS4{F0: (v3 - (^(v3))), F1: true}
                                                                    _ = v40
                                                                case "a":
                                                                    var v41 bool = (("" + "hello") == "a")
                                                                    _ = v41
                                                                case "hello":
                                                                    v42 := GenS3{F0: (fi24 | fi25), F1: (true || false), F2: 68}
                                                                    _ = v42
                                                                default:
                                                                    v3 += (v3 & 725)
                                                                }
                                                                if len(v4) > 0 {
                                                                    v4[0] = (fi25 >> ((255) & 63))
                                                                }
                                                                fmt.Println(fi24)
                                                            }
                                                        }
                                                        var v43 uint = 42
                                                        _ = v43
                                                    }
                                                    v44 := GenS5{F0: 84}
                                                    fmt.Println(fi25)
                                                    _ = v44
                                                }
                                                v4 = append(v4, fi25)
                                                var v45 float64 = ((((3.14159 + ((0.1 - 0.1) * float64((1e10 * (1.5 - 0.1))))) * 1.5) * (0.1 + -2.25)) / 3.14159)
                                                fmt.Println(fi24, v11)
                                                v46 := GenS2{F0: v23}
                                                if len(v4) > 0 {
                                                    v4[0] = (-(fi25))
                                                }
                                                _ = v30
                                                _ = v45
                                                _ = v46
                                            }
                                            var v47 string = ""
                                            v48 := GenS6{F0: 255, F1: (^(v11)), F2: 65535}
                                            switch 127 {
                                            case 230:
                                                if false {
                                                    delete(v1, 0)
                                                    fmt.Println(fi25)
                                                } else {
                                                    v47 = "foo bar"
                                                    if true {
                                                        v49 := GenS7{F0: v3}
                                                        v50 := GenS5{F0: (v11 / ((v11) | 1))}
                                                        _ = len(v1)
                                                        v3 += 1
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                    if true {
                                                        delete(v1, 641)
                                                        var v51 int8 = -128
                                                        _ = v51
                                                    }
                                                }
                                            case -2147483648:
                                                switch v47 {
                                                case "foo bar":
                                                    fmt.Println(fi25, v23, fi25)
                                                case "a":
                                                    if (!false) {
                                                        _ = len(v4)
                                                    } else {
                                                        if false {
                                                            var v52 uint16 = 255
                                                            if false {
                                                                v23 = (v23 ^ v23)
                                                                v53 := GenS2{F0: (v23 | (^(v23)))}
                                                                _ = v53
                                                            } else {
                                                                var v54 bool = true
                                                                _ = v54
                                                            }
                                                            _ = v52
                                                        }
                                                        var v55 bool = true
                                                        _ = v55
                                                    }
                                                case "\n\t":
                                                    v23 = v23
                                                default:
                                                    delete(v2, v47)
                                                }
                                            case 128:
                                                var v56 int16 = (-(v23))
                                                _ = v56
                                            }
                                            v47 = (v47 + v47)
                                            _ = fi25
                                            _ = v47
                                            _ = v48
                                        }
                                        for fi57 := 0; fi57 < 12; fi57++ {
                                            v23 = 757
                                            v23 -= (v23 >> ((949) & 63))
                                            _ = fi57
                                        }
                                        var v58 int8 = 0
                                        _ = v58
                                    } else {
                                        v3 += v3
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        for fi59 := 0; fi59 < 5; fi59++ {
                                            fmt.Println(fi59, v3, fi59)
                                            _ = fi59
                                        }
                                    }
                                    delete(v2, "foo bar")
                                    v60 := GenS8{F0: v23, F1: 136, F2: 42}
                                    _ = v60
                                } else {
                                    v4 = append(v4, fi24)
                                    fmt.Println(v3, v23, v3)
                                }
                                v23 = (v23 ^ v23)
                                _ = fi24
                            }
                            _ = v23
                        } else {
                            for fi61 := 0; fi61 < 11; fi61++ {
                                if (("" + "hello") != ("foo bar" + "a")) {
                                    v11 += (v11 ^ v11)
                                    for fi62 := 0; fi62 < 18; fi62++ {
                                        v63 := v1[638]
                                        var v64 int8 = 1
                                        fmt.Println(v3, v3, v3)
                                        _ = fi62
                                        _ = v63
                                        _ = v64
                                    }
                                    v65 := GenS9{F0: -9223372036854775808, F1: 4294967295}
                                    _ = v65
                                }
                                _ = fi61
                            }
                            fmt.Println(v11, v3)
                        }
                    }
                }
            case "foo bar":
                switch "" {
                case "hello":
                    v66 := GenS2{F0: -32768}
                    _ = v66
                case "foo bar":
                    fmt.Println(v3)
                case "a":
                    var v67 uint32 = 114
                    _ = v67
                }
            }
        }
        _ = len(v1)
    } else {
        var v68 int16 = 0
        var v69 float64 = 0.1
        _ = v68
        _ = v69
    }
    v11 -= 128
    switch -2147483648 {
    case 0:
        v11 += (v11 | 139)
    case 128:
        switch (("a" + "foo bar") + ("" + "\n\t")) {
        case "\n\t":
            v70 := GenS7{F0: v3}
            _ = v70
        case "foo bar":
            v3 -= (v3 / (((v3 >> ((326) & 63))) | 1))
        }
    }
    var v71 uint = 65535
    v3 = (v3 / ((1) | 1))
    switch 42 {
    case 2147483647:
        var v72 uint64 = v3
        _ = v72
    case 32768:
        for fi73 := 0; fi73 < 1; fi73++ {
            fmt.Println(v11, v71)
            fmt.Println(v11)
            _ = fi73
        }
    case 42:
        v3 = v3
    }
    var v74 uint16 = uint16(v71)
    v71 = v71
    var v75 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v71:", v71)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v71
    _ = v74
    _ = v75
}
