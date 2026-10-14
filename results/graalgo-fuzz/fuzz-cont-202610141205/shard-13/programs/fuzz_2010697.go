package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64; F1 uint32 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]uint32)
    var v3 uint64 = 42
    if true {
        if false {
            v4 := GenS1{F0: ("a" + "")}
            fmt.Println(v1, v3)
            switch ("foo bar" + "") {
            case "hello":
                v5 := GenS1{F0: ((("" + "foo bar") + "foo bar") + "\n\t")}
                _ = v5
            default:
                var v6 int8 = -1
                _ = v6
            }
            _ = v4
        } else {
            var v7 uint = 160
            switch "" {
            case "a":
                var v8 int16 = 42
                _ = v8
            default:
                var v9 int32 = 2147483647
                _ = v9
            }
            _ = v7
        }
        v10 := GenS2{F0: 0, F1: 2147483647, F2: (v3 ^ 1)}
        v3 = (v3 - 4294967295)
        _ = v10
    } else {
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS3{F0: 352}
            var v13 int16 = 42
            _ = fi11
            _ = v12
            _ = v13
        }
        switch 32768 {
        case 432:
            v1 = v1
        case 128:
            for fi14 := 0; fi14 < 12; fi14++ {
                switch ("" + ("" + "hello")) {
                case "hello":
                    var v15 uint32 = 2147483648
                    _ = v15
                case "a":
                    for fi16 := 0; fi16 < 9; fi16++ {
                        var v17 uint = 42
                        var v18 bool = (308 <= -1)
                        var v19 float32 = 0.1
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                case "":
                    if true {
                        v20 := GenS2{F0: 2147483647, F1: 0, F2: v3}
                        v21 := []int{fi14, -579, -1}
                        v3 += (v3 - (v3 - (^(v3))))
                        _ = v20
                        _ = v21
                    } else {
                        var v22 uint8 = v1
                        var v23 int16 = int16(v1)
                        if true {
                            for fi24 := 0; fi24 < 10; fi24++ {
                                delete(v2, (fi24 % ((-1) | 1)))
                                v25 := []int{(fi14 + fi14), (^(fi24)), fi24}
                                var v26 uint = 18446744073709551615
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            fmt.Println(v22, v22)
                            switch 4294967295 {
                            case 983:
                                for fi27 := 0; fi27 < 8; fi27++ {
                                    var v28 uint16 = 2
                                    v28 = 32767
                                    switch (fi14 % ((fi27) | 1)) {
                                    case 42:
                                        v29 := GenS3{F0: -2147483648}
                                        _ = v29
                                    case -744:
                                        switch (^(fi27)) {
                                        case 2147483647:
                                            if false {
                                                var v30 int16 = -572
                                                _ = v30
                                            }
                                        case 1:
                                            v31 := GenS1{F0: "\n\t"}
                                            _ = v31
                                        }
                                    case 4294967295:
                                        var v32 uint16 = (v28 - (v28 * 32767))
                                        _ = v32
                                    }
                                    _ = fi27
                                    _ = v28
                                }
                            default:
                                v33 := []int{-962}
                                _ = v33
                            }
                        }
                        _ = v22
                        _ = v23
                    }
                }
                fmt.Println(fi14, fi14)
                v34 := v2[fi14]
                _ = fi14
                _ = v34
            }
        }
    }
    var v35 int16 = -735
    switch -2147483648 {
    case -9223372036854775808:
        for fi36 := 0; fi36 < 12; fi36++ {
            var v37 int = fi36
            if false {
                v38 := GenS4{F0: 518, F1: 1}
                var v39 uint = 507
                _ = v38
                _ = v39
            }
            v35 += v35
            _ = fi36
            _ = v37
        }
    case -486:
        fmt.Println(v35, v35, v35)
    case 0:
        if true {
            var v40 uint8 = uint8(v3)
            var v41 int32 = -209
            v40 += v1
            _ = v40
            _ = v41
        } else {
            v1 = v1
        }
    default:
        v42 := GenS5{F0: v3, F1: 0}
        _ = v42
    }
    var v43 string = "a"
    var v44 float64 = 3.14159
    v35 = v35
    if (42 <= -70) {
        if true {
            var v45 string = "a"
            _ = v45
        }
        v46 := []int{0, 128}
        _ = v46
    } else {
        for fi47 := 0; fi47 < 20; fi47++ {
            v48 := GenS1{F0: v43}
            fmt.Println(v44, v3, v44)
            _ = fi47
            _ = v48
        }
    }
    for fi49 := 0; fi49 < 8; fi49++ {
        v50 := v2[(fi49 << ((2) & 63))]
        if (!((v3 ^ 0) < (v3 >> ((357) & 63)))) {
            for fi51 := 0; fi51 < 5; fi51++ {
                if (float64(0.1) != v44) {
                    switch 1 {
                    case 9223372036854775807:
                        var v52 uint32 = uint32(v1)
                        _ = v52
                    }
                } else {
                    fmt.Println(v35)
                }
                var v53 bool = (!false)
                _ = fi51
                _ = v53
            }
            v54 := GenS6{F0: 42, F1: ((fi49 != 128) && (fi49 != fi49))}
            if (65535 > v50) {
                if false {
                    var v55 string = (v43 + v43)
                    var v56 float64 = float64(float32((0.1 * 1e10)))
                    _ = v55
                    _ = v56
                } else {
                    for fi57 := 0; fi57 < 18; fi57++ {
                        var v58 uint = 255
                        switch int(v1) {
                        case -1:
                            v50 += (v50 << ((643) & 63))
                        }
                        v59 := GenS1{F0: "\n\t"}
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                    v1 = v1
                    if true {
                        for fi60 := 0; fi60 < 8; fi60++ {
                            if (!((v3 * 841) <= (v3 >> ((uint(v50)) & 63)))) {
                                switch "hello" {
                                case "":
                                    if false {
                                        var v61 int16 = (v35 / (((-(v35))) | 1))
                                        fmt.Println(fi60)
                                        _ = v61
                                    } else {
                                        switch v43 {
                                        case "a":
                                            v1 = (v1 << ((255) & 63))
                                        }
                                    }
                                case "\n\t":
                                    v62 := GenS6{F0: -9223372036854775808, F1: false}
                                    _ = v62
                                case "foo bar":
                                    var v63 int16 = 32767
                                    _ = v63
                                }
                                var v64 int8 = 100
                                v65 := GenS7{F0: v1}
                                var v66 int32 = 275
                                _ = v64
                                _ = v65
                                _ = v66
                            } else {
                                if true {
                                    var v67 uint8 = v1
                                    if (439 == v35) {
                                        v68 := []int{(-(fi60)), 688, -9223372036854775808, -9223372036854775808}
                                        if false {
                                            var v69 int16 = -32768
                                            switch 271 {
                                            case -944:
                                                if (945 <= 65535) {
                                                    var v70 int32 = -2147483648
                                                    var v71 int16 = -206
                                                    _ = v70
                                                    _ = v71
                                                } else {
                                                    switch (fi60 ^ -1) {
                                                    case 0:
                                                        var v72 uint16 = 947
                                                        _ = v72
                                                    case 9223372036854775807:
                                                        if true {
                                                            fmt.Println(v69)
                                                        }
                                                    case 4294967295:
                                                        v43 = v43
                                                    default:
                                                        fmt.Println(v35, v43)
                                                    }
                                                }
                                            case -824:
                                                v43 = ((v43 + (("a" + ((v43 + ((((v43 + "\n\t") + "") + v43) + (("hello" + v43) + ((v43 + ("foo bar" + "a")) + v43)))) + v43)) + ((v43 + "hello") + v43))) + v43)
                                            case -538:
                                                if (65535 == v50) {
                                                    var v73 uint64 = v3
                                                    var v74 string = "hello"
                                                    v73 -= uint64(v67)
                                                    _ = v73
                                                    _ = v74
                                                }
                                            default:
                                                var v75 uint16 = 2
                                                _ = v75
                                            }
                                            _ = v69
                                        } else {
                                            _ = len(v2)
                                            v2[(fi60 % ((fi49) | 1))] = v50
                                            fmt.Println(v35, v50, v67)
                                        }
                                        _ = v68
                                    }
                                    _ = v67
                                }
                                v76 := []int{}
                                switch "hello" {
                                case "a":
                                    if true {
                                        var v77 float32 = float32((v44 + v44))
                                        if len(v76) >= 2 {
                                            v76 = v76[1:len(v76)]
                                        }
                                        v78 := GenS1{F0: v43}
                                        _ = v77
                                        _ = v78
                                    }
                                }
                                var v79 uint16 = 32767
                                v3 -= (v3 + (^(v3)))
                                _ = v76
                                _ = v79
                            }
                            if (true || false) {
                                v35 -= -32768
                                v80 := GenS1{F0: v43}
                                fmt.Println(v43)
                                _ = v80
                            }
                            switch (fi60 | (fi60 & -334)) {
                            case 397:
                                v81 := GenS7{F0: (v1 + v1)}
                                _ = v81
                            case 32767:
                                var v82 int32 = 0
                                _ = v82
                            case -809:
                                var v83 uint64 = v3
                                _ = v83
                            default:
                                v1 += v1
                            }
                            _ = fi60
                        }
                        v50 = (^(v50))
                    } else {
                        var v84 uint64 = (v3 << ((379) & 63))
                        v85 := GenS8{F0: 0, F1: 2147483648}
                        if true {
                            v50 += 726
                            var v86 int64 = -1
                            _ = v86
                        } else {
                            switch (fi49 & fi49) {
                            case 32767:
                                v50 -= (^(v50))
                            }
                            var v87 bool = true
                            v88 := []int{(fi49 - 32767), -9223372036854775808, fi49, -131}
                            _ = v87
                            _ = v88
                        }
                        _ = v84
                        _ = v85
                    }
                }
            } else {
                v89 := GenS2{F0: (v50 << ((774) & 63)), F1: 0, F2: v3}
                _ = v89
            }
            _ = v54
        } else {
            v50 += 4294967295
            fmt.Println(v35, v35)
            fmt.Println(v3)
        }
        v50 -= uint32(v1)
        _ = fi49
        _ = v50
    }
    v90 := []int{1, 680, -890}
    _ = len(v90)
    v91 := GenS4{F0: 682, F1: -2147483648}
    fmt.Println(v43, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v90:", len(v90), v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v43
    _ = v44
    _ = v90
    _ = v91
}
