package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 string; F1 int8; F2 string }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 int64; F1 int16 }
type GenS10 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int32, p1 uint8, p2 uint32) int {
    var v4 int32 = (p0 << ((2) & 63))
    if true {
        v5 := []int{32768, 128, -9223372036854775808, 127}
        for fi6 := 0; fi6 < 8; fi6++ {
            if false {
                var v7 int32 = (p0 + 42)
                _ = v7
            } else {
                switch fi6 {
                case -938:
                    v4 = v4
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                v8 := GenS1{F0: 148, F1: v4}
                _ = len(v5)
                _ = v8
            }
            if false {
                fmt.Println(fi6)
                for fi9 := 0; fi9 < 2; fi9++ {
                    fmt.Println(v4)
                    _ = fi9
                }
            } else {
                if (false || true) {
                    if (!(1.5 >= 0.1)) {
                        if (255 != 255) {
                            var v10 uint8 = (p1 - p1)
                            for fi11 := 0; fi11 < 15; fi11++ {
                                var v12 uint8 = (v10 * p1)
                                v4 = (v4 * v4)
                                _ = fi11
                                _ = v12
                            }
                            _ = v10
                        } else {
                            p1 += (p1 ^ p1)
                            for fi13 := 0; fi13 < 6; fi13++ {
                                _ = len(v5)
                                fmt.Println(fi6, v4, v4)
                                _ = fi13
                            }
                        }
                        var v14 uint32 = (p2 * p2)
                        var v15 uint = 65535
                        _ = v14
                        _ = v15
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        fmt.Println(p1)
                    }
                    v16 := GenS1{F0: 127, F1: (v4 & (p0 + (p0 | 42)))}
                    _ = v16
                }
            }
            _ = fi6
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 int16 = 716
            var v19 uint8 = 42
            if len(v5) > 0 {
                v5[0] = (fi17 ^ fi17)
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v5
    } else {
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 uint64 = 2
            if false {
                if false {
                    p2 = (p2 >> ((1) & 63))
                } else {
                    var v22 int16 = -1
                    fmt.Println(p1)
                    fmt.Println(p0, p0)
                    _ = v22
                }
                p2 -= p2
            }
            v23 := GenS1{F0: (p1 ^ uint8(v21)), F1: p0}
            _ = fi20
            _ = v21
            _ = v23
        }
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 32768
}
func genFn1(p0 uint) uint16 {
    if true {
        v24 := []int{}
        _ = len(v24)
        _ = v24
    }
    fmt.Println(p0)
    if false {
        v25 := GenS2{F0: "\n\t"}
        _ = v25
    } else {
        p0 = (p0 & (p0 >> ((p0) & 63)))
    }
    p0 = (p0 / ((p0) | 1))
    var v26 int = -9223372036854775808
    _ = p0
    _ = v26
    return 65535
}

func main() {
    var v1 int64 = -1
    var v2 float64 = float64(float32((1e10 + (0.1 + -2.25))))
    v3 := make(map[int]bool)
    v27 := GenS2{F0: "foo bar"}
    fmt.Println(v1, v2)
    if ("a" < ("hello" + "hello")) {
        for fi28 := 0; fi28 < 13; fi28++ {
            v29 := v3[805]
            v3[fi28] = false
            for fi30 := 0; fi30 < 2; fi30++ {
                var v31 string = ("hello" + "foo bar")
                _ = fi30
                _ = v31
            }
            _ = fi28
            _ = v29
        }
        v2 = 0.1
        var v32 string = "a"
        _ = v32
    } else {
        switch "" {
        case "a":
            if false {
                for fi33 := 0; fi33 < 8; fi33++ {
                    v1 = (v1 & v1)
                    v34 := GenS3{F0: "foo bar", F1: 887, F2: 0}
                    _ = fi33
                    _ = v34
                }
            }
        case "hello":
            v35 := GenS4{F0: 4294967295}
            _ = v35
        case "foo bar":
            v1 = 9223372036854775807
        }
        for fi36 := 0; fi36 < 19; fi36++ {
            if (v1 != v1) {
                v2 = 3.14159
                fmt.Println(v2, v2)
                fmt.Println(fi36, v2, fi36)
            } else {
                switch (fi36 >> ((18446744073709551615) & 63)) {
                case -234:
                    for fi37 := 0; fi37 < 4; fi37++ {
                        switch ("\n\t" + (("hello" + "foo bar") + "a")) {
                        case "hello":
                            fmt.Println(fi37)
                        default:
                            v38 := []int{fi36, fi37, fi36, genFn0(867, 42, 65535)}
                            _ = v38
                        }
                        var v39 int64 = 9223372036854775807
                        v40 := GenS5{F0: 702, F1: 0}
                        _ = fi37
                        _ = v39
                        _ = v40
                    }
                }
            }
            switch (("\n\t" + (("a" + ("\n\t" + ("\n\t" + (("\n\t" + "") + "hello")))) + (("a" + ("" + "a")) + ("hello" + "")))) + "foo bar") {
            case "a":
                v41 := GenS1{F0: 128, F1: -2147483648}
                _ = v41
            case "":
                var v42 uint8 = 0
                _ = v42
            default:
                v43 := GenS6{F0: "", F1: genFn0(0, 47, 65535)}
                _ = v43
            }
            _ = fi36
        }
        v2 = v2
    }
    v1 = v1
    if true {
        var v44 uint16 = 255
        v45 := v3[-482]
        switch (("a" + "a") + "hello") {
        case "a":
            if v45 {
                v46 := []int{genFn0(-369, 151, 28), 127}
                if v45 {
                    var v47 uint32 = 65535
                    fmt.Println(v1)
                    _ = v47
                }
                v1 = (v1 * v1)
                _ = v46
            }
        default:
            v48 := GenS7{F0: ("" + "foo bar"), F1: 127, F2: ("a" + ("" + ("a" + (("hello" + ("\n\t" + "")) + "a"))))}
            _ = v48
        }
        _ = v44
        _ = v45
    } else {
        v2 += (v2 + float64(0.1))
        if true {
            v49 := GenS2{F0: ""}
            var v50 int16 = 0
            v1 += v1
            _ = v49
            _ = v50
        }
        if (4294967295 != 652) {
            for fi51 := 0; fi51 < 5; fi51++ {
                if true {
                    var v52 float64 = (0.1 + 0.0)
                    _ = v52
                } else {
                    fmt.Println(v1, v1, fi51)
                }
                v53 := GenS7{F0: "a", F1: 127, F2: (((("foo bar" + "") + "a") + "\n\t") + ("a" + "hello"))}
                _ = fi51
                _ = v53
            }
            var v54 uint32 = 340
            _ = v54
        } else {
            for fi55 := 0; fi55 < 9; fi55++ {
                var v56 float64 = (v2 * v2)
                v1 = v1
                _ = fi55
                _ = v56
            }
        }
    }
    v1 = v1
    var v57 int16 = 0
    var v58 int8 = 0
    for fi59 := 0; fi59 < 4; fi59++ {
        if (v2 > (1e10 + v2)) {
            var v60 int32 = 678
            var v61 uint64 = 894
            var v62 uint32 = 0
            _ = v60
            _ = v61
            _ = v62
        } else {
            for fi63 := 0; fi63 < 20; fi63++ {
                v64 := []int{fi59}
                switch ("\n\t" + "a") {
                case "hello":
                    if true {
                        if false {
                            switch (("\n\t" + "foo bar") + "foo bar") {
                            case "hello":
                                v65 := GenS8{F0: (fi59 % ((fi59) | 1)), F1: 42}
                                _ = v65
                            case "\n\t":
                                v66 := GenS9{F0: (v1 / ((-657) | 1)), F1: 778}
                                _ = v66
                            default:
                                v1 = v1
                            }
                        } else {
                            var v67 int8 = v58
                            var v68 float32 = (0.0 * 1.5)
                            _ = v67
                            _ = v68
                        }
                        fmt.Println(v58, v58)
                    }
                default:
                    fmt.Println(v1, v1, v57)
                }
                _ = fi63
                _ = v64
            }
            var v69 int16 = v57
            v58 = v58
            _ = v69
        }
        var v70 int8 = v58
        _ = fi59
        _ = v70
    }
    fmt.Println(v57)
    if ((-(v1)) > (v1 % ((v1) | 1))) {
        for fi71 := 0; fi71 < 10; fi71++ {
            v2 += (-2.25 * (v2 + v2))
            v57 = (v57 >> ((1) & 63))
            v72 := GenS4{F0: 1}
            _ = fi71
            _ = v72
        }
        v3[99] = true
        var v73 string = ("\n\t" + "\n\t")
        _ = v73
    } else {
        v74 := GenS7{F0: ("foo bar" + "a"), F1: int8(v57), F2: ("" + "foo bar")}
        var v75 uint64 = 42
        _ = v74
        _ = v75
    }
    switch -9223372036854775808 {
    case -256:
        var v76 uint = 954
        _ = v76
    case 4294967295:
        var v77 bool = false
        _ = v77
    }
    for fi78 := 0; fi78 < 7; fi78++ {
        if (("foo bar" < (("\n\t" + "") + ((("\n\t" + ("" + ("\n\t" + (("\n\t" + "a") + (("foo bar" + "hello") + ("foo bar" + ("hello" + ((("\n\t" + "") + ("hello" + "a")) + "\n\t")))))))) + "\n\t") + "a"))) && (v58 != v58)) {
            for fi79 := 0; fi79 < 4; fi79++ {
                v80 := v3[(-(fi78))]
                _ = fi79
                _ = v80
            }
        }
        switch (fi78 | (fi78 + (^(fi78)))) {
        case 1000:
            switch "hello" {
            case "a":
                var v81 uint = 4294967295
                _ = v81
            case "\n\t":
                v82 := GenS10{F0: 255, F1: 1}
                _ = v82
            case "":
                var v83 int32 = 42
                _ = v83
            }
        case 9:
            v84 := GenS7{F0: ("hello" + "\n\t"), F1: v58, F2: "foo bar"}
            _ = v84
        case 559:
            switch "\n\t" {
            case "\n\t":
                v85 := GenS3{F0: "hello", F1: 0, F2: 288}
                _ = v85
            case "":
                var v86 float32 = -2.25
                _ = v86
            default:
                v1 = (v1 + v1)
            }
        }
        for fi87 := 0; fi87 < 14; fi87++ {
            v58 += v58
            switch "a" {
            case "a":
                v2 = float64((3.14159 - 0.1))
                switch genFn0(-1, 2, 2) {
                case -414:
                    v57 += v57
                case 569:
                    v88 := v3[fi78]
                    _ = v88
                }
                v3[genFn0(-1, 255, 0)] = false
            default:
                var v89 uint64 = 680
                _ = v89
            }
            _ = fi87
        }
        _ = fi78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v57
    _ = v58
}
