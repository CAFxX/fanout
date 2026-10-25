package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 int8 }
type GenS8 struct { F0 int32; F1 int8 }

func main() {
    v1 := []int{-904, 1}
    var v2 string = "a"
    var v3 float32 = (0.1 / -2.25)
    var v4 uint8 = 216
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]string)
    if (!(-1 == -2147483648)) {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int32 = -1
            _ = fi7
            _ = v8
        }
        switch v2 {
        case "foo bar":
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 uint16 = 1
                var v11 uint16 = (^(v10))
                v12 := v5[0]
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        }
        v6[895] = (v2 + "a")
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 bool = (!false)
        if false {
            var v15 int64 = 1
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    v2 = "hello"
    var v16 int64 = -1
    _ = len(v1)
    if (9223372036854775807 <= v16) {
        if (!((120 != 2147483648) || (980 < 641))) {
            _ = len(v1)
        } else {
            v4 = 0
            if true {
                switch v2 {
                case "foo bar":
                    v17 := GenS1{F0: 250}
                    _ = v17
                case "":
                    _ = len(v6)
                case "\n\t":
                    v3 -= -2.25
                default:
                    switch 1 {
                    case 32767:
                        fmt.Println(v3, v16)
                    }
                }
                v4 += v4
            }
        }
        var v18 int8 = -65
        v2 = ""
        _ = v18
    } else {
        fmt.Println(v3, v16, v4)
        for fi19 := 0; fi19 < 15; fi19++ {
            for fi20 := 0; fi20 < 15; fi20++ {
                v21 := GenS1{F0: 4294967295}
                var v22 float32 = ((((v3 / (v3 / v3)) + (v3 / (float32(0.0) - float32(3.14159)))) - v3) - -2.25)
                var v23 float32 = v22
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            delete(v6, 440)
            _ = fi19
        }
    }
    v3 = 0.0
    switch int(v16) {
    case 32768:
        _ = len(v5)
    case -2147483648:
        var v24 float32 = 0.1
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 uint8 = v4
        if false {
            fmt.Println(v3, fi25, v2)
            v4 = v4
        } else {
            for fi27 := 0; fi27 < 14; fi27++ {
                if false {
                    switch v2 {
                    case "a":
                        if ("\n\t" < "foo bar") {
                            v28 := GenS2{F0: v26}
                            var v29 uint64 = 0
                            v30 := GenS1{F0: 1}
                            _ = v28
                            _ = v29
                            _ = v30
                        } else {
                            if true {
                                switch (fi25 << ((318) & 63)) {
                                case 4294967295:
                                    v16 += (v16 << ((0) & 63))
                                case 128:
                                    v16 += -1
                                case 42:
                                    v31 := GenS2{F0: (^(v4))}
                                    _ = v31
                                }
                            }
                            v2 = v2
                            v16 = (v16 * (v16 * (v16 - v16)))
                            switch fi27 {
                            case 128:
                                if (v26 >= 128) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v6[42] = (v2 + ((v2 + v2) + ("" + (v2 + "foo bar"))))
                                } else {
                                    var v32 int = (fi27 % ((42) | 1))
                                    v32 -= 32767
                                    _ = v32
                                }
                                fmt.Println(v3)
                            }
                            var v33 int16 = 32767
                            _ = v33
                        }
                    default:
                        v16 = (v16 & (v16 % ((v16) | 1)))
                    }
                    v34 := GenS3{F0: 0, F1: 127}
                    v35 := GenS4{F0: 42}
                    _ = v34
                    _ = v35
                }
                var v36 uint8 = v26
                fmt.Println(v26, v3)
                _ = fi27
                _ = v36
            }
            for fi37 := 0; fi37 < 2; fi37++ {
                v1 = append(v1, 980)
                if (227 <= 0) {
                    v4 = v26
                    var v38 uint16 = 42
                    v26 = v26
                    v16 = (-(v16))
                    if (127 > 58) {
                        fmt.Println(v3, fi37)
                    }
                    var v39 string = v2
                    _ = v38
                    _ = v39
                }
                _ = fi37
            }
        }
        var v40 uint16 = 877
        _ = fi25
        _ = v26
        _ = v40
    }
    for fi41 := 0; fi41 < 16; fi41++ {
        for fi42 := 0; fi42 < 14; fi42++ {
            v43 := v6[2147483648]
            var v44 uint64 = 77
            var v45 float64 = 1.5
            _ = fi42
            _ = v43
            _ = v44
            _ = v45
        }
        var v46 int = fi41
        var v47 int32 = 2147483647
        _ = fi41
        _ = v46
        _ = v47
    }
    var v48 uint8 = 11
    v49 := GenS4{F0: 65535}
    v2 = ((v2 + v2) + "hello")
    delete(v6, 962)
    v50 := GenS4{F0: 2}
    for fi51 := 0; fi51 < 7; fi51++ {
        for fi52 := 0; fi52 < 8; fi52++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v53 int32 = -2147483648
            if true {
                var v54 bool = false
                switch ("foo bar" + (("foo bar" + v2) + v2)) {
                case "foo bar":
                    v48 = (^(v4))
                    _ = len(v5)
                    var v55 bool = v54
                    v56 := GenS3{F0: 2147483648, F1: -1}
                    _ = v55
                    _ = v56
                case "":
                    v4 -= v48
                default:
                    var v57 int = 2147483647
                    _ = v57
                }
                switch (fi52 * (fi51 << ((0) & 63))) {
                case 128:
                    switch 1 {
                    case 440:
                        var v58 int8 = 0
                        _ = v58
                    }
                default:
                    if v54 {
                        var v59 uint32 = 194
                        var v60 uint32 = (v59 & v59)
                        v48 = v4
                        _ = v59
                        _ = v60
                    }
                }
                _ = v54
            } else {
                switch (fi51 % ((fi52) | 1)) {
                case -1:
                    _ = len(v5)
                case 2147483647:
                    var v61 int64 = 1
                    _ = v61
                case 4294967295:
                    var v62 int64 = v16
                    _ = v62
                default:
                    switch v2 {
                    case "\n\t":
                        if true {
                            v48 += 0
                            if true {
                                var v63 uint64 = 18446744073709551615
                                if (1 > 1000) {
                                    if false {
                                        fmt.Println(v16, v53, v3)
                                    } else {
                                        switch (fi52 - fi51) {
                                        case 32767:
                                            v64 := GenS1{F0: 2}
                                            _ = v64
                                        default:
                                            v3 += v3
                                        }
                                    }
                                    switch "hello" {
                                    case "hello":
                                        v2 = v2
                                        v65 := GenS5{F0: 216}
                                        switch v2 {
                                        case "a":
                                            var v66 float64 = ((1e10 - -2.25) * 3.14159)
                                            _ = v66
                                        case "hello":
                                            var v67 int64 = -9223372036854775808
                                            _ = v67
                                        case "foo bar":
                                            switch (fi51 ^ (fi51 / (((fi52 << ((830) & 63))) | 1))) {
                                            case 161:
                                                v68 := GenS1{F0: 935}
                                                _ = v68
                                            }
                                        default:
                                            var v69 int32 = (v53 + (v53 >> ((369) & 63)))
                                            _ = v69
                                        }
                                        _ = v65
                                    default:
                                        fmt.Println(v3, v4)
                                    }
                                    v2 = "\n\t"
                                }
                                switch "a" {
                                case "foo bar":
                                    v70 := GenS2{F0: v48}
                                    v71 := GenS4{F0: 65535}
                                    switch 4294967295 {
                                    case 4294967295:
                                        v72 := GenS2{F0: v4}
                                        _ = v72
                                    case 379:
                                        v48 = v48
                                    default:
                                        v4 = (v48 / ((77) | 1))
                                    }
                                    _ = v70
                                    _ = v71
                                case "a":
                                    v53 = v53
                                default:
                                    v16 = v16
                                }
                                _ = v63
                            } else {
                                v16 += v16
                                v3 = 0.1
                                if true {
                                    fmt.Println(v2)
                                    v73 := GenS6{F0: 0}
                                    _ = v73
                                } else {
                                    v2 = v2
                                    var v74 float32 = (float32(1e10) / (float32(-2.25) + 0.1))
                                    v75 := GenS7{F0: false, F1: -1}
                                    _ = v74
                                    _ = v75
                                }
                            }
                            v4 = 2
                        }
                    default:
                        v4 -= v48
                    }
                }
            }
            _ = fi52
            _ = v53
        }
        v48 = (v48 >> ((0) & 63))
        fmt.Println(v16, v3, v16)
        _ = fi51
    }
    if (2147483647 == -2147483648) {
        v76 := GenS8{F0: 42, F1: -128}
        if (v48 <= v4) {
            fmt.Println(v2)
            var v77 int = 32768
            _ = v77
        } else {
            v78 := GenS4{F0: 251}
            fmt.Println(v48, v2, v16)
            _ = v78
        }
        _ = v76
    } else {
        v79 := GenS4{F0: 1}
        _ = v79
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
