package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int; F2 int8 }
type GenS3 struct { F0 uint16; F1 string; F2 int8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 int; F1 string; F2 int }
type GenS7 struct { F0 uint16; F1 int8; F2 int16 }
type GenS8 struct { F0 uint8 }
func genFn0(p0 bool) int32 {
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: "hello"}
        var v7 uint = 4294967295
        var v8 bool = true
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    if p0 {
        p0 = (128 > 127)
        v9 := GenS2{F0: false, F1: 2147483647, F2: -61}
        _ = v9
    } else {
        switch -2147483648 {
        case 4294967295:
            var v10 uint64 = 1
            _ = v10
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            switch "\n\t" {
            case "a":
                v12 := []int{128, (^(fi11))}
                _ = v12
            case "foo bar":
                p0 = false
            case "hello":
                if p0 {
                    for fi13 := 0; fi13 < 5; fi13++ {
                        var v14 int16 = 320
                        v14 += (v14 + (v14 | (^(v14))))
                        fmt.Println(v14, fi13)
                        _ = fi13
                        _ = v14
                    }
                    var v15 int16 = -1
                    _ = v15
                }
            default:
                var v16 int8 = 100
                _ = v16
            }
            p0 = true
            p0 = true
            _ = fi11
        }
    }
    switch 9223372036854775807 {
    case 32768:
        v17 := []int{1, 32767, 32767}
        _ = v17
    }
    var v18 bool = (-1 != 100)
    var v19 int8 = 42
    _ = p0
    _ = v18
    _ = v19
    return 42
}
func genFn1(p0 uint, p1 int32, p2 int64) int64 {
    switch ("\n\t" + ("a" + "hello")) {
    case "foo bar":
        var v20 int32 = p1
        _ = v20
    case "":
        p1 += (p1 | p1)
    case "a":
        p1 = genFn0((("" + "\n\t") == ("foo bar" + "a")))
    default:
        if true {
            for fi21 := 0; fi21 < 20; fi21++ {
                v22 := GenS2{F0: (18446744073709551615 != 4294967295), F1: (fi21 << ((252) & 63)), F2: -42}
                _ = fi21
                _ = v22
            }
            var v23 float32 = -2.25
            _ = v23
        } else {
            fmt.Println(p1)
        }
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 int64 = 9223372036854775807
        for fi26 := 0; fi26 < 1; fi26++ {
            p0 += (p0 % ((771) | 1))
            p1 += p1
            v25 = p2
            if (!false) {
                v27 := GenS2{F0: ("hello" > ""), F1: 970, F2: -102}
                switch (fi24 - 128) {
                case 0:
                    var v28 int32 = (p1 * p1)
                    _ = v28
                default:
                    var v29 bool = (1.5 > float64(0.0))
                    _ = v29
                }
                _ = v27
            } else {
                var v30 int = fi26
                v31 := GenS2{F0: (!true), F1: 1, F2: 16}
                var v32 uint64 = 0
                _ = v30
                _ = v31
                _ = v32
            }
            _ = fi26
        }
        if (-1 != (fi24 << ((p0) & 63))) {
            if true {
                var v33 float32 = 0.1
                for fi34 := 0; fi34 < 15; fi34++ {
                    v35 := []int{fi24, 9223372036854775807, (fi34 ^ 32767), (fi24 >> ((p0) & 63))}
                    v33 += (1e10 / 1e10)
                    if false {
                        if false {
                            var v36 float64 = 1.5
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                            v37 := GenS3{F0: 1, F1: "hello", F2: -1}
                            _ = v36
                            _ = v37
                        }
                        var v38 int64 = v25
                        switch "hello" {
                        case "\n\t":
                            var v39 uint8 = uint8(p0)
                            _ = v39
                        case "hello":
                            if false {
                                fmt.Println(p2)
                                v35 = append(v35, fi24)
                                var v40 uint = 18446744073709551615
                                if (true || (65535 == 4294967295)) {
                                    v41 := GenS1{F0: (("a" + "\n\t") + "a")}
                                    _ = v41
                                }
                                var v42 int64 = p2
                                _ = v40
                                _ = v42
                            }
                        case "":
                            var v43 int8 = 100
                            _ = v43
                        }
                        _ = v38
                    } else {
                        var v44 string = "foo bar"
                        if true {
                            var v45 bool = (false && (!false))
                            _ = v45
                        } else {
                            v33 = float32(1.5)
                            var v46 float32 = v33
                            v47 := GenS4{F0: (p1 << (((p0 - (p0 ^ 18446744073709551615))) & 63)), F1: v44}
                            var v48 uint32 = 816
                            var v49 uint16 = 255
                            _ = v46
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                        _ = v44
                    }
                    _ = fi34
                    _ = v35
                }
                v50 := []int{fi24}
                _ = v33
                _ = v50
            }
        } else {
            var v51 uint16 = 269
            var v52 uint32 = 2
            _ = v51
            _ = v52
        }
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 bool = (1 < 740)
    var v2 int = 128
    var v3 uint64 = 18446744073709551615
    var v4 uint8 = 0
    var v53 int = 128
    switch 127 {
    case 444:
        fmt.Println(v4)
    }
    v54 := []int{v53, v53, v53}
    var v55 int64 = -938
    var v56 int64 = v55
    for fi57 := 0; fi57 < 10; fi57++ {
        v2 -= v2
        _ = fi57
    }
    for fi58 := 0; fi58 < 11; fi58++ {
        if len(v54) >= 2 {
            v54 = v54[1:len(v54)]
        }
        v1 = v1
        _ = fi58
    }
    var v59 int32 = -1
    if v1 {
        if len(v54) >= 2 {
            v54 = v54[1:len(v54)]
        }
        if v1 {
            if v1 {
                v54 = append(v54, -167)
                v60 := GenS3{F0: 0, F1: ("" + ("" + "a")), F2: -95}
                _ = v60
            }
            var v61 uint = 255
            if (true && v1) {
                var v62 uint8 = (v4 - v4)
                if v1 {
                    fmt.Println(v56, v62)
                    var v63 string = "hello"
                    _ = v63
                }
                _ = v62
            } else {
                v64 := GenS5{F0: (v56 ^ v55), F1: 0}
                fmt.Println(v56, v53)
                _ = len(v54)
                _ = v64
            }
            _ = v61
        }
        if false {
            v3 += (v3 + v3)
            if v1 {
                var v65 int64 = (v55 + (^(v56)))
                if v1 {
                    v56 += (v56 << ((65535) & 63))
                    if len(v54) > 0 {
                        v54[0] = (v53 << ((2) & 63))
                    }
                } else {
                    var v66 uint32 = 0
                    if (v53 <= (v53 / ((v53) | 1))) {
                        var v67 uint32 = (v66 & uint32(v3))
                        _ = v67
                    }
                    v2 += v53
                    _ = v66
                }
                _ = v65
            } else {
                var v68 float64 = -2.25
                _ = v68
            }
        } else {
            if false {
                v54 = append(v54, int(v3))
            } else {
                var v69 string = ("\n\t" + "foo bar")
                var v70 bool = true
                if (!(!(9223372036854775807 != genFn1(255, genFn0((349 <= -694)), v55)))) {
                    var v71 float64 = 1.5
                    var v72 int64 = v56
                    _ = v71
                    _ = v72
                } else {
                    v53 += v2
                }
                _ = v69
                _ = v70
            }
            if ((-(v2)) != (v2 | (v53 | 42))) {
                switch v2 {
                case -140:
                    var v73 int64 = (v56 ^ v55)
                    _ = v73
                case 9223372036854775807:
                    for fi74 := 0; fi74 < 17; fi74++ {
                        var v75 int8 = 1
                        _ = fi74
                        _ = v75
                    }
                case 128:
                    var v76 uint64 = uint64(v4)
                    _ = v76
                }
                var v77 int8 = -127
                _ = v77
            } else {
                v78 := GenS3{F0: 32767, F1: "a", F2: -128}
                _ = v78
            }
            v55 = -1
        }
    } else {
        fmt.Println(v1, v53)
        v79 := GenS6{F0: (v2 >> ((2) & 63)), F1: "a", F2: -2147483648}
        _ = v79
    }
    var v80 int16 = 42
    var v81 int = (v53 << ((18446744073709551615) & 63))
    v82 := GenS4{F0: (v59 / (((v59 << ((65535) & 63))) | 1)), F1: (("\n\t" + "foo bar") + (("\n\t" + ("\n\t" + "foo bar")) + "hello"))}
    switch (v53 % ((v53) | 1)) {
    case -2147483648:
        var v83 int64 = v55
        _ = v83
    case 127:
        if len(v54) >= 2 {
            v54 = v54[1:len(v54)]
        }
    case 316:
        v4 = 42
    }
    switch ("hello" + "") {
    case "\n\t":
        if (!false) {
            v84 := GenS7{F0: 593, F1: -115, F2: -593}
            _ = v84
        } else {
            for fi85 := 0; fi85 < 11; fi85++ {
                if (!(v1 && v1)) {
                    var v86 uint32 = 2
                    _ = v86
                } else {
                    if ((!v1) || ((v59 >> ((123) & 63)) > (v59 >> ((254) & 63)))) {
                        for fi87 := 0; fi87 < 4; fi87++ {
                            switch 32767 {
                            case 127:
                                v54 = append(v54, fi85)
                            case -821:
                                var v88 int8 = -94
                                _ = v88
                            case -2147483648:
                                fmt.Println(fi85, fi85, v3)
                            }
                            switch (fi87 ^ (^(v2))) {
                            case -105:
                                var v89 uint = 255
                                _ = v89
                            }
                            _ = fi87
                        }
                        v90 := GenS8{F0: (v4 - v4)}
                        _ = v90
                    }
                }
                _ = fi85
            }
        }
    case "hello":
        var v91 float32 = float32(3.14159)
        _ = v91
    case "a":
        _ = len(v54)
    }
    var v92 int16 = v80
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v59:", v59)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v59
    _ = v80
    _ = v81
    _ = v82
    _ = v92
}
