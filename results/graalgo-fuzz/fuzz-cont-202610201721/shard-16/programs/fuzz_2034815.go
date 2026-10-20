package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 string }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint; F1 int32; F2 uint }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 bool; F1 int16; F2 uint }
func genFn0(p0 int8) uint16 {
    if false {
        v6 := GenS1{F0: 1, F1: 2, F2: "\n\t"}
        var v7 int16 = -23
        _ = v6
        _ = v7
    } else {
        switch 2147483647 {
        case -2147483648:
            if false {
                var v8 int = 696
                fmt.Println(p0, v8)
                if (42 < 315) {
                    v9 := GenS1{F0: 4294967295, F1: 1, F2: ("" + "\n\t")}
                    _ = v9
                }
                _ = v8
            }
        case 2147483647:
            var v10 int16 = 0
            _ = v10
        case 42:
            var v11 uint16 = 571
            _ = v11
        default:
            var v12 uint64 = 2
            _ = v12
        }
        switch ("a" + "foo bar") {
        case "a":
            p0 = p0
        case "":
            switch "\n\t" {
            case "\n\t":
                if (65535 < 1) {
                    var v13 float32 = 3.14159
                    if (971 > 2) {
                        var v14 uint = 618
                        p0 = p0
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 6; fi15++ {
                        for fi16 := 0; fi16 < 15; fi16++ {
                            var v17 string = ""
                            if (!true) {
                                var v18 uint64 = 0
                                v19 := []int{fi15}
                                _ = v18
                                _ = v19
                            }
                            v20 := []int{(^(fi16)), 9223372036854775807, fi15, fi16}
                            v13 = float32(3.14159)
                            v13 -= v13
                            _ = fi16
                            _ = v17
                            _ = v20
                        }
                        _ = fi15
                    }
                    _ = v13
                } else {
                    p0 = (-(p0))
                    v21 := GenS2{F0: -22, F1: (p0 * (p0 | (p0 * (p0 & p0))))}
                    p0 = -128
                    _ = v21
                }
            default:
                var v22 int = 2147483647
                _ = v22
            }
        case "foo bar":
            var v23 bool = (-511 != 1)
            _ = v23
        default:
            if false {
                var v24 uint64 = 951
                _ = v24
            } else {
                for fi25 := 0; fi25 < 13; fi25++ {
                    switch "\n\t" {
                    case "hello":
                        v26 := []int{128, (-(fi25)), (fi25 / ((fi25) | 1)), (fi25 + fi25)}
                        _ = v26
                    }
                    fmt.Println(p0)
                    _ = fi25
                }
            }
        }
    }
    p0 += p0
    fmt.Println(p0)
    switch ("" + "") {
    case "\n\t":
        var v27 uint64 = 4294967295
        _ = v27
    case "hello":
        var v28 int = 42
        _ = v28
    case "foo bar":
        if (!false) {
            v29 := []int{362, 9223372036854775807}
            _ = v29
        }
    }
    _ = p0
    return 34
}
func genFn1() bool {
    v30 := GenS2{F0: -1, F1: 100}
    var v31 bool = ((3.14159 > (1e10 / 1.0)) || false)
    v32 := GenS3{F0: 4294967295, F1: 42, F2: 4294967295}
    _ = v30
    _ = v31
    _ = v32
    return true
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]string)
    var v3 int8 = 127
    var v4 uint16 = 42
    var v5 uint64 = 2
    for fi33 := 0; fi33 < 2; fi33++ {
        var v34 int8 = 127
        switch "" {
        case "":
            fmt.Println(fi33, v4, v3)
        case "a":
            var v35 bool = false
            _ = v35
        default:
            fmt.Println(fi33, fi33)
        }
        if false {
            v36 := GenS1{F0: 2147483647, F1: 42, F2: "\n\t"}
            for fi37 := 0; fi37 < 9; fi37++ {
                switch "" {
                case "\n\t":
                    if false {
                        v38 := GenS1{F0: 0, F1: 0, F2: ("hello" + "\n\t")}
                        v39 := []int{0, (-(fi33)), int(v4)}
                        fmt.Println(fi33, fi37, v34)
                        _ = v38
                        _ = v39
                    }
                case "a":
                    v40 := GenS4{F0: 0, F1: 2}
                    _ = v40
                default:
                    v34 -= 42
                }
                var v41 int8 = (^(v3))
                switch ("foo bar" + ((("\n\t" + "foo bar") + "\n\t") + ((("a" + "\n\t") + ("" + "\n\t")) + (("foo bar" + "hello") + "\n\t")))) {
                case "foo bar":
                    v42 := []int{(fi33 - (fi33 / ((-962) | 1))), (fi33 | -9223372036854775808), (fi33 % (((^(fi33))) | 1))}
                    _ = v42
                }
                switch ("hello" + "") {
                case "a":
                    v43 := GenS4{F0: 518, F1: 969}
                    _ = v43
                case "hello":
                    v34 += (v34 % ((1) | 1))
                default:
                    var v44 string = "foo bar"
                    _ = v44
                }
                _ = fi37
                _ = v41
            }
            v5 += (^(v5))
            _ = v36
        } else {
            for fi45 := 0; fi45 < 3; fi45++ {
                if (("hello" != "foo bar") && (4 > 726)) {
                    v34 = (-(v34))
                    var v46 string = ("foo bar" + "\n\t")
                    _ = v46
                }
                _ = fi45
            }
        }
        _ = fi33
        _ = v34
    }
    switch "hello" {
    case "hello":
        for fi47 := 0; fi47 < 20; fi47++ {
            var v48 string = ""
            v5 = 738
            var v49 int = (fi47 / ((fi47) | 1))
            _ = fi47
            _ = v48
            _ = v49
        }
    case "foo bar":
        var v50 float64 = (0.0 * (1e10 * -2.25))
        _ = v50
    case "a":
        switch ("" + "\n\t") {
        case "":
            var v51 bool = false
            _ = v51
        case "hello":
            delete(v2, "hello")
        case "\n\t":
            v52 := GenS3{F0: 136, F1: 2147483647, F2: 645}
            _ = v52
        default:
            v4 += genFn0(v3)
        }
    default:
        if ((v5 << ((828) & 63)) != 677) {
            if (!(((true || (0.0 <= 1.5)) || (32767 <= int16(v5))) && ((0.1 * 1.5) == float64(-2.25)))) {
                for fi53 := 0; fi53 < 18; fi53++ {
                    var v54 uint16 = (v4 - (^(v4)))
                    _ = fi53
                    _ = v54
                }
                if true {
                    switch ("foo bar" + "foo bar") {
                    case "":
                        if (-1 == 42) {
                            var v55 uint = 702
                            _ = v55
                        }
                    default:
                        var v56 int16 = 42
                        _ = v56
                    }
                }
                if (-9223372036854775808 > 9223372036854775807) {
                    v4 -= genFn0((v3 / ((v3) | 1)))
                }
            } else {
                var v57 uint = 4294967295
                v58 := []int{-314, 127, 0}
                _ = v57
                _ = v58
            }
            var v59 int32 = 42
            _ = v59
        }
    }
    var v60 int64 = -1
    fmt.Println(v4, v5)
    v5 = v5
    var v61 int64 = (v60 + 0)
    _ = len(v2)
    for fi62 := 0; fi62 < 10; fi62++ {
        v63 := GenS4{F0: 0, F1: 42}
        _ = fi62
        _ = v63
    }
    v64 := GenS5{F0: true, F1: -648, F2: 1}
    if (false && (-1 < -822)) {
        v4 += (v4 % (((v4 / ((0) | 1))) | 1))
        for fi65 := 0; fi65 < 19; fi65++ {
            v66 := GenS5{F0: true, F1: 0, F2: 575}
            _ = fi65
            _ = v66
        }
    } else {
        v5 = (^(v5))
        var v67 uint = 4294967295
        _ = v67
    }
    switch (("\n\t" + "hello") + ("foo bar" + (("foo bar" + "foo bar") + ("" + "\n\t")))) {
    case "":
        v68 := []int{}
        _ = v68
    default:
        if true {
            v5 -= 323
            var v69 int = -2147483648
            var v70 string = "\n\t"
            _ = v69
            _ = v70
        }
    }
    v5 = (v5 + (v5 << ((255) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v61
    _ = v64
}
