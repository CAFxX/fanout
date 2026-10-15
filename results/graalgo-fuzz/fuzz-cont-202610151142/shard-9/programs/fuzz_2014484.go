package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint16; F1 bool; F2 int16 }
type GenS4 struct { F0 int64; F1 int; F2 int64 }
type GenS5 struct { F0 uint16; F1 int16; F2 int16 }
type GenS6 struct { F0 uint64; F1 bool; F2 string }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int8; F1 string }
func genFn0(p0 int, p1 uint64) uint {
    v7 := GenS1{F0: "foo bar"}
    v8 := GenS2{F0: (p0 | p0), F1: (^(p1))}
    fmt.Println(p1, p1)
    if (18446744073709551615 > 1) {
        v9 := []int{p0, -233, p0}
        var v10 float64 = 0.1
        _ = len(v9)
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        if (1e10 <= -2.25) {
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 float64 = float64(((((0.1 * 1.5) + ((1.5 * 1e10) + 1.5)) + float32(1.5)) + 0.0))
                p1 += p1
                switch 127 {
                case -739:
                    v13 = (v13 * -2.25)
                case 0:
                    if (129 >= 65535) {
                        if false {
                            v14 := GenS3{F0: 387, F1: false, F2: -32768}
                            p0 += p0
                            v13 = 1e10
                            _ = v14
                        } else {
                            fmt.Println(p0, v13, p0)
                            v15 := GenS1{F0: "a"}
                            _ = v15
                        }
                        if (!true) {
                            var v16 string = ""
                            v17 := []int{(fi12 + (fi12 >> ((0) & 63))), p0}
                            _ = v16
                            _ = v17
                        } else {
                            fmt.Println(fi11, p0, p0)
                        }
                        switch "a" {
                        case "a":
                            p1 = (^(p1))
                        case "":
                            fmt.Println(p0, p0, fi11)
                        case "\n\t":
                            p1 = p1
                        }
                    } else {
                        var v18 int32 = 2147483647
                        _ = v18
                    }
                case 9223372036854775807:
                    if true {
                        p1 = 2
                        v19 := GenS4{F0: -847, F1: p0, F2: 36}
                        if (false || (!(false || false))) {
                            var v20 int16 = 702
                            if false {
                                var v21 int64 = -1
                                v22 := GenS2{F0: 1, F1: 972}
                                _ = v21
                                _ = v22
                            } else {
                                var v23 string = "a"
                                var v24 uint8 = 42
                                _ = v23
                                _ = v24
                            }
                            fmt.Println(p1, fi11)
                            if true {
                                v20 = (v20 | v20)
                            } else {
                                var v25 uint = 1
                                _ = v25
                            }
                            _ = v20
                        }
                        _ = v19
                    }
                    fmt.Println(p0)
                default:
                    var v26 int32 = 1
                    _ = v26
                }
                var v27 float32 = (1e10 / 1e10)
                v28 := GenS2{F0: 333, F1: 42}
                _ = fi12
                _ = v13
                _ = v27
                _ = v28
            }
        } else {
            var v29 uint64 = 156
            if (!false) {
                v30 := GenS1{F0: "\n\t"}
                p1 = p1
                var v31 float32 = 3.14159
                _ = v30
                _ = v31
            } else {
                p0 = -998
            }
            fmt.Println(v29, p1, fi11)
            _ = v29
        }
        p0 = (fi11 / ((32768) | 1))
        v32 := GenS2{F0: fi11, F1: (p1 & p1)}
        _ = fi11
        _ = v32
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 335
}
func genFn1() bool {
    var v33 int8 = 47
    for fi34 := 0; fi34 < 15; fi34++ {
        v35 := GenS5{F0: 65535, F1: 1, F2: -32768}
        v36 := GenS2{F0: (fi34 + (fi34 << ((679) & 63))), F1: 178}
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 float64 = float64(1.5)
    _ = v33
    _ = v37
    return false
}

func main() {
    v1 := []string{"a", ("\n\t" + (("foo bar" + "foo bar") + ("a" + "hello"))), ("foo bar" + "a"), ("\n\t" + (("foo bar" + "") + "\n\t"))}
    var v2 int = -437
    var v3 int8 = 1
    var v4 uint8 = 127
    var v5 bool = ((!false) || true)
    var v6 int64 = int64(v3)
    var v38 int = v2
    for fi39 := 0; fi39 < 1; fi39++ {
        v4 += 127
        if (!(-2.25 == (3.14159 * float32(0.1)))) {
            if len(v1) > 0 {
                v1[0] = ("\n\t" + "a")
            }
        } else {
            if (v5 && v5) {
                if true {
                    v38 = 32767
                    fmt.Println(v3, v4, v2)
                }
                fmt.Println(v2)
                var v40 float32 = 1.5
                _ = v40
            }
            if (true || false) {
                v3 += (v3 | 96)
                for fi41 := 0; fi41 < 11; fi41++ {
                    fmt.Println(v5)
                    var v42 string = "\n\t"
                    _ = fi41
                    _ = v42
                }
            }
            v43 := GenS5{F0: 42, F1: 32767, F2: int16(v6)}
            _ = v43
        }
        _ = fi39
    }
    v2 -= v2
    v6 += v6
    for fi44 := 0; fi44 < 5; fi44++ {
        var v45 int = fi44
        _ = fi44
        _ = v45
    }
    v6 -= v6
    for fi46 := 0; fi46 < 4; fi46++ {
        fmt.Println(v4, v6, v5)
        if (2147483647 > 685) {
            if true {
                var v47 uint8 = 2
                for fi48 := 0; fi48 < 17; fi48++ {
                    v4 = (v47 | (v47 & v4))
                    if (129 > uint(v47)) {
                        v38 += int(v6)
                        v49 := GenS2{F0: fi46, F1: 2}
                        v50 := GenS5{F0: 2, F1: 32767, F2: -504}
                        _ = v49
                        _ = v50
                    } else {
                        fmt.Println(v2)
                    }
                    var v51 float32 = (-2.25 + float32(1.5))
                    _ = fi48
                    _ = v51
                }
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                _ = v47
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v52 uint64 = 42
            _ = v52
        } else {
            v53 := GenS6{F0: 1, F1: true, F2: ""}
            v54 := GenS5{F0: 42, F1: -358, F2: 0}
            v38 = v38
            _ = v53
            _ = v54
        }
        _ = fi46
    }
    _ = len(v1)
    v55 := GenS3{F0: 487, F1: (true && (true && (v3 == v3))), F2: 42}
    v3 = (v3 | 42)
    switch "a" {
    case "foo bar":
        for fi56 := 0; fi56 < 3; fi56++ {
            for fi57 := 0; fi57 < 5; fi57++ {
                var v58 float64 = 0.1
                v59 := GenS7{F0: "foo bar"}
                _ = fi57
                _ = v58
                _ = v59
            }
            switch (fi56 & v2) {
            case -77:
                var v60 float32 = (1e10 + 1e10)
                _ = v60
            case 32768:
                var v61 int8 = (v3 * v3)
                _ = v61
            case 0:
                if v5 {
                    var v62 int16 = 32767
                    _ = v62
                } else {
                    v1 = append(v1, ("\n\t" + ("a" + "a")))
                    _ = len(v1)
                    _ = len(v1)
                }
            }
            _ = fi56
        }
    default:
        _ = len(v1)
    }
    var v63 string = "a"
    v64 := GenS3{F0: 877, F1: false, F2: 72}
    fmt.Println(v63, v38, v63)
    switch -2147483648 {
    case 0:
        v1 = append(v1, (v63 + (("foo bar" + (v63 + "hello")) + (("foo bar" + ((v63 + ("foo bar" + ("a" + "foo bar"))) + v63)) + v63))))
    default:
        switch "a" {
        case "a":
            fmt.Println(v38, v63)
        case "":
            var v65 float32 = -2.25
            _ = v65
        default:
            if v5 {
                switch (v38 + -766) {
                case -302:
                    v66 := GenS2{F0: 42, F1: 4294967295}
                    _ = v66
                default:
                    if (false && v5) {
                        v67 := GenS4{F0: v6, F1: v2, F2: v6}
                        _ = v67
                    }
                }
            } else {
                v68 := GenS5{F0: 0, F1: 0, F2: 0}
                _ = v68
            }
        }
    }
    if v5 {
        var v69 uint16 = 42
        if (!v5) {
            v6 = v6
            v4 -= v4
        }
        if true {
            v1 = append(v1, "a")
            v6 += v6
            if true {
                if ((v4 >> ((1) & 63)) <= (v4 * 0)) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                for fi70 := 0; fi70 < 16; fi70++ {
                    v71 := GenS8{F0: fi70}
                    _ = fi70
                    _ = v71
                }
            } else {
                for fi72 := 0; fi72 < 20; fi72++ {
                    var v73 uint32 = 725
                    _ = fi72
                    _ = v73
                }
            }
        } else {
            var v74 int16 = int16(v4)
            _ = v74
        }
        _ = v69
    } else {
        v75 := GenS9{F0: v3, F1: ""}
        v76 := GenS1{F0: "a"}
        v5 = (65535 != 313)
        _ = v75
        _ = v76
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v63
    _ = v64
}
