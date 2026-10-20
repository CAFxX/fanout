package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 uint32; F1 int64; F2 int64 }
type GenS8 struct { F0 bool; F1 bool }
type GenS9 struct { F0 uint; F1 uint32; F2 int8 }
func genFn0(p0 string, p1 uint8, p2 bool) uint16 {
    var v6 int = 907
    if p2 {
        for fi7 := 0; fi7 < 1; fi7++ {
            if ((^(p1)) <= 42) {
                var v8 uint32 = 2
                var v9 bool = ("foo bar" < "a")
                _ = v8
                _ = v9
            }
            p1 = (^(p1))
            _ = fi7
        }
        p0 = p0
        fmt.Println(v6, p1)
    } else {
        if p2 {
            switch "a" {
            case "a":
                p1 -= p1
            case "":
                if (-2 < -84) {
                    v10 := []int{v6, v6}
                    for fi11 := 0; fi11 < 2; fi11++ {
                        for fi12 := 0; fi12 < 3; fi12++ {
                            p2 = (!p2)
                            _ = fi12
                        }
                        switch 32767 {
                        case 1:
                            fmt.Println(p1)
                        case 32767:
                            var v13 uint16 = 42
                            _ = v13
                        default:
                            var v14 float64 = 0.1
                            _ = v14
                        }
                        _ = fi11
                    }
                    for fi15 := 0; fi15 < 14; fi15++ {
                        p0 = "a"
                        if (!(1 <= 878)) {
                            for fi16 := 0; fi16 < 14; fi16++ {
                                var v17 uint64 = 2
                                v18 := GenS1{F0: (-1 <= 941)}
                                if ((p2 || true) || (!(!(!(!(!p2)))))) {
                                    v6 += fi16
                                    v19 := GenS2{F0: 0}
                                    v20 := GenS2{F0: int32(fi15)}
                                    _ = v19
                                    _ = v20
                                } else {
                                    if ((128 != p1) || ((^(p1)) < 2)) {
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                        var v21 int16 = 1
                                        switch 32768 {
                                        case 127:
                                            v6 -= fi16
                                        case -628:
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                        }
                                        v22 := GenS2{F0: 42}
                                        fmt.Println(v6, v21)
                                        _ = len(v10)
                                        _ = v21
                                        _ = v22
                                    } else {
                                        var v23 uint64 = v17
                                        _ = v23
                                    }
                                    v24 := GenS3{F0: -426, F1: 607}
                                    if (false || p2) {
                                        v25 := GenS1{F0: false}
                                        if len(v10) > 0 {
                                            v10[0] = (fi16 >> ((18446744073709551615) & 63))
                                        }
                                        _ = v25
                                    } else {
                                        v26 := GenS4{F0: p1, F1: 612}
                                        _ = v26
                                    }
                                    var v27 string = "foo bar"
                                    v28 := GenS3{F0: -1, F1: 65535}
                                    _ = v24
                                    _ = v27
                                    _ = v28
                                }
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        } else {
                            p0 = "hello"
                            var v29 uint = 42
                            _ = v29
                        }
                        _ = fi15
                    }
                    _ = v10
                } else {
                    var v30 uint8 = 42
                    for fi31 := 0; fi31 < 3; fi31++ {
                        v32 := []int{(fi31 * (v6 + (v6 << ((65535) & 63)))), (v6 & fi31), 742}
                        _ = fi31
                        _ = v32
                    }
                    var v33 uint64 = 4294967295
                    _ = v30
                    _ = v33
                }
            case "foo bar":
                v6 += v6
            }
            for fi34 := 0; fi34 < 1; fi34++ {
                fmt.Println(fi34)
                var v35 float32 = ((-2.25 / 1.5) * (float32(float64((3.14159 - 0.1))) - 0.1))
                _ = fi34
                _ = v35
            }
        }
        p0 = ("a" + p0)
    }
    if p2 {
        p2 = (p2 || ((^(v6)) == 693))
        if (0.1 >= 1e10) {
            var v36 uint64 = 0
            switch (p0 + p0) {
            case "\n\t":
                var v37 float32 = 1e10
                _ = v37
            case "a":
                for fi38 := 0; fi38 < 15; fi38++ {
                    var v39 int8 = 100
                    _ = fi38
                    _ = v39
                }
            case "foo bar":
                var v40 uint16 = 42
                _ = v40
            default:
                for fi41 := 0; fi41 < 1; fi41++ {
                    p0 = ""
                    p0 = "foo bar"
                    _ = fi41
                }
            }
            v42 := []int{32767, (v6 ^ v6), -374}
            _ = v36
            _ = v42
        }
    } else {
        for fi43 := 0; fi43 < 10; fi43++ {
            var v44 uint = 2
            _ = fi43
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 9; fi45++ {
        v46 := GenS4{F0: (p1 ^ 0), F1: -2147483648}
        v47 := []int{}
        v47 = append(v47, fi45)
        _ = fi45
        _ = v46
        _ = v47
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 538
}
func genFn1() int64 {
    for fi48 := 0; fi48 < 9; fi48++ {
        var v49 int32 = 1
        v49 -= 0
        _ = fi48
        _ = v49
    }
    var v50 bool = false
    _ = v50
    return -1
}

func main() {
    var v1 float32 = (1e10 * 1e10)
    var v2 string = ("hello" + ("foo bar" + "foo bar"))
    var v3 uint8 = 148
    v4 := []string{}
    var v5 uint8 = 110
    v4 = append(v4, (v2 + (v2 + ("foo bar" + (v2 + ("\n\t" + v2))))))
    _ = len(v4)
    v1 -= -2.25
    for fi51 := 0; fi51 < 1; fi51++ {
        switch (v2 + v2) {
        case "a":
            for fi52 := 0; fi52 < 13; fi52++ {
                v53 := GenS4{F0: 128, F1: 2147483647}
                var v54 uint16 = 436
                _ = fi52
                _ = v53
                _ = v54
            }
        case "foo bar":
            if true {
                var v55 int8 = 42
                _ = v55
            }
        case "\n\t":
            if (!true) {
                fmt.Println(fi51, fi51)
            } else {
                if true {
                    var v56 int64 = 42
                    v57 := GenS5{F0: genFn0(((v2 + (v2 + v2)) + ("" + ((("foo bar" + v2) + "foo bar") + (v2 + ((v2 + v2) + (v2 + (v2 + v2))))))), 128, (!false))}
                    _ = v56
                    _ = v57
                } else {
                    fmt.Println(v1)
                    v58 := GenS5{F0: 65535}
                    _ = v58
                }
            }
        }
        _ = fi51
    }
    var v59 int8 = 0
    for fi60 := 0; fi60 < 10; fi60++ {
        v59 = (v59 - v59)
        var v61 float32 = float32(((3.14159 * 0.0) * 0.0))
        if (false && (v59 < (^(v59)))) {
            for fi62 := 0; fi62 < 18; fi62++ {
                var v63 string = (v2 + v2)
                var v64 int32 = 412
                _ = fi62
                _ = v63
                _ = v64
            }
            switch int(v5) {
            case -2147483648:
                fmt.Println(v61, v3, v1)
            case 127:
                var v65 int32 = 0
                _ = v65
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        } else {
            v61 += v61
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi60
        _ = v61
    }
    var v66 string = "hello"
    switch 4294967295 {
    case 128:
        for fi67 := 0; fi67 < 16; fi67++ {
            var v68 int = (fi67 * 0)
            _ = fi67
            _ = v68
        }
    default:
        switch -6 {
        case -690:
            var v69 int64 = -1
            _ = v69
        }
    }
    switch "hello" {
    case "foo bar":
        var v70 float64 = (0.0 - 0.0)
        _ = v70
    }
    if (float64(1.5) < (1e10 + 1.5)) {
        v2 = v66
        var v71 uint32 = 2
        _ = v71
    }
    for fi72 := 0; fi72 < 15; fi72++ {
        var v73 uint32 = 749
        _ = fi72
        _ = v73
    }
    switch -214 {
    case -184:
        if ((^(v3)) > v5) {
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch -209 {
                case -301:
                    var v74 bool = false
                    _ = v74
                case 280:
                    if true {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v1, v2)
                    }
                case 1:
                    v75 := GenS6{F0: 1, F1: (v2 + "foo bar")}
                    _ = v75
                }
                switch ("hello" + (v2 + v66)) {
                case "foo bar":
                    for fi76 := 0; fi76 < 11; fi76++ {
                        if true {
                            _ = len(v4)
                            switch (v2 + "a") {
                            case "\n\t":
                                var v77 string = "hello"
                                _ = v77
                            }
                            switch (fi76 - 403) {
                            case -9223372036854775808:
                                var v78 int8 = (v59 - 0)
                                _ = v78
                            case 32768:
                                fmt.Println(v66, v1)
                            default:
                                v66 = v66
                            }
                        }
                        _ = fi76
                    }
                case "\n\t":
                    v1 -= float32(3.14159)
                case "hello":
                    var v79 string = "a"
                    _ = v79
                }
            }
            if (208 < 705) {
                if (true && (127 == 9223372036854775807)) {
                    var v80 uint8 = 0
                    _ = v80
                }
                for fi81 := 0; fi81 < 16; fi81++ {
                    if (255 <= 0) {
                        v4 = append(v4, (v2 + ""))
                        v3 = (v3 & v5)
                        v5 -= v3
                    }
                    var v82 float32 = (v1 + v1)
                    _ = fi81
                    _ = v82
                }
            } else {
                switch 315 {
                case 784:
                    var v83 float32 = (-2.25 - v1)
                    _ = v83
                default:
                    fmt.Println(v3, v3)
                }
                var v84 uint16 = 702
                _ = v84
            }
            switch "a" {
            case "a":
                var v85 uint = 255
                _ = v85
            case "foo bar":
                if ("" <= "hello") {
                    for fi86 := 0; fi86 < 2; fi86++ {
                        v66 = "\n\t"
                        _ = fi86
                    }
                    if false {
                        v2 = v66
                    }
                }
            case "":
                v87 := GenS7{F0: 465, F1: -9223372036854775808, F2: genFn1()}
                _ = v87
            }
        }
    }
    v88 := GenS6{F0: v3, F1: (v66 + (v66 + "a"))}
    var v89 string = ((((v2 + v66) + (v66 + "foo bar")) + v66) + "\n\t")
    if false {
        if ((v59 << ((1) & 63)) == v59) {
            var v90 uint64 = 1
            v91 := GenS8{F0: (2147483647 >= -222), F1: true}
            if false {
                v92 := GenS2{F0: 2147483647}
                var v93 bool = (((1000 > 1000) || false) || (false || (v59 == v59)))
                _ = v92
                _ = v93
            } else {
                switch 2147483647 {
                case -303:
                    v94 := GenS5{F0: 0}
                    _ = v94
                default:
                    var v95 int = int(v5)
                    _ = v95
                }
                var v96 bool = false
                v97 := GenS9{F0: 416, F1: 254, F2: v59}
                _ = v96
                _ = v97
            }
            _ = v90
            _ = v91
        }
    } else {
        for fi98 := 0; fi98 < 12; fi98++ {
            v1 = 1e10
            v59 = -128
            _ = fi98
        }
        var v99 int32 = 42
        _ = v99
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v66:", v66)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v66
    _ = v88
    _ = v89
}
