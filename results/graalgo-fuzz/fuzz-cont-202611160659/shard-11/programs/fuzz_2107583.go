package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS4 struct { F0 int16; F1 string; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 uint16; F2 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint; F1 uint16 }
func genFn0(p0 bool) uint {
    p0 = (!(!false))
    fmt.Println(p0)
    var v5 int = -9223372036854775808
    v6 := GenS1{F0: (!true), F1: (v5 | 42)}
    var v7 uint = 757
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (v7 << ((1) & 63))
}
func genFn1(p0 int8, p1 uint) bool {
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := []int{fi8, fi8, (^(fi8))}
        v10 := GenS1{F0: (fi8 > (fi8 >> (((p1 << (((p1 & p1)) & 63))) & 63))), F1: fi8}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 float64 = 0.0
    if false {
        for fi12 := 0; fi12 < 1; fi12++ {
            switch (fi12 - (-(fi12))) {
            case 1:
                var v13 uint = (p1 ^ genFn0(false))
                _ = v13
            case 32768:
                v14 := []int{fi12, (fi12 * (fi12 << ((p1) & 63)))}
                _ = v14
            case 42:
                fmt.Println(v11)
            }
            _ = fi12
        }
        v15 := []int{4294967295, 127, -9223372036854775808}
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint = 255
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        if (!false) {
            if false {
                fmt.Println(p1, p0, p1)
            }
            v18 := []int{32767, -194, 9223372036854775807, 32768}
            _ = v18
        } else {
            p0 += -1
        }
        if ((699 <= p1) || (32767 >= 65535)) {
            for fi19 := 0; fi19 < 7; fi19++ {
                fmt.Println(p0, p1)
                p0 = p0
                _ = fi19
            }
            v20 := GenS1{F0: (!((685 > 255) || false)), F1: 32768}
            if false {
                v11 = v11
                switch 4294967295 {
                case -127:
                    switch ("" + "hello") {
                    case "a":
                        v21 := GenS2{F0: 1, F1: (-179 >= 2147483647)}
                        _ = v21
                    }
                default:
                    p0 -= (p0 | p0)
                }
                fmt.Println(p1, v11, v11)
            }
            _ = v20
        } else {
            for fi22 := 0; fi22 < 10; fi22++ {
                var v23 uint32 = 42
                if (p1 >= p1) {
                    if true {
                        fmt.Println(p0, p0)
                    }
                    var v24 uint8 = 255
                    _ = v24
                } else {
                    v25 := []int{fi22, fi22, (^(fi22)), 345}
                    if true {
                        if len(v25) > 0 {
                            v25[0] = fi22
                        }
                        var v26 uint32 = uint32(p1)
                        if ("hello" != (("a" + "a") + "foo bar")) {
                            var v27 uint64 = 18446744073709551615
                            v28 := GenS3{F0: 704, F1: uint64(p1), F2: 128}
                            _ = v27
                            _ = v28
                        }
                        _ = v26
                    }
                    _ = v25
                }
                _ = fi22
                _ = v23
            }
            if (false || (false || true)) {
                p1 += (p1 | (p1 + p1))
                var v29 bool = (true || (-9223372036854775808 == 10))
                _ = v29
            } else {
                v30 := []int{-91, 0, 9223372036854775807, 128}
                _ = v30
            }
        }
        if (51 != 687) {
            for fi31 := 0; fi31 < 2; fi31++ {
                v32 := GenS2{F0: fi31, F1: false}
                for fi33 := 0; fi33 < 7; fi33++ {
                    p0 -= (^(p0))
                    _ = fi33
                }
                _ = fi31
                _ = v32
            }
            v34 := GenS1{F0: true, F1: 655}
            v35 := GenS4{F0: 0, F1: ("" + "a"), F2: 0}
            _ = v34
            _ = v35
        } else {
            v11 -= v11
            switch -1 {
            case -2147483648:
                fmt.Println(p0)
            case 127:
                v36 := []int{-1, 127}
                _ = v36
            case -348:
                v37 := []int{945}
                _ = v37
            }
        }
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        var v39 uint16 = 1
        for fi40 := 0; fi40 < 6; fi40++ {
            v41 := []int{}
            var v42 uint8 = 127
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = fi38
        _ = v39
    }
    _ = p0
    _ = p1
    _ = v11
    return (1 > 1)
}

func main() {
    v1 := []string{("hello" + "")}
    var v2 uint = 2
    var v3 string = "foo bar"
    v4 := make(map[int]int)
    switch (v3 + ((("\n\t" + "") + "hello") + ("" + v3))) {
    case "a":
        v3 = v3
    case "":
        v43 := GenS5{F0: uint64(v2), F1: 819, F2: 18446744073709551615}
        _ = v43
    case "foo bar":
        for fi44 := 0; fi44 < 18; fi44++ {
            switch (fi44 ^ 2147483647) {
            case 426:
                var v45 int64 = -1
                _ = v45
            }
            _ = fi44
        }
    default:
        var v46 uint8 = 255
        _ = v46
    }
    v2 = v2
    switch ("\n\t" + v3) {
    case "\n\t":
        switch (v3 + v3) {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case "hello":
        for fi47 := 0; fi47 < 14; fi47++ {
            switch 32767 {
            case -1:
                v1 = append(v1, "\n\t")
            case 960:
                fmt.Println(v3, v3, fi47)
            case -9223372036854775808:
                for fi48 := 0; fi48 < 6; fi48++ {
                    var v49 int = (^(fi48))
                    var v50 int8 = 42
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            default:
                v3 = "a"
            }
            var v51 float32 = (3.14159 - -2.25)
            _ = fi47
            _ = v51
        }
    default:
        for fi52 := 0; fi52 < 20; fi52++ {
            switch ("" + v3) {
            case "hello":
                for fi53 := 0; fi53 < 14; fi53++ {
                    v3 = "hello"
                    v2 += (v2 << ((v2) & 63))
                    _ = fi53
                }
            }
            v54 := GenS6{F0: 0}
            v3 = ("a" + (v3 + ((v3 + v3) + ("hello" + v3))))
            _ = fi52
            _ = v54
        }
    }
    var v55 uint64 = 846
    if (true && false) {
        v56 := GenS7{F0: 2147483647, F1: 882, F2: v2}
        fmt.Println(v55, v3, v3)
        _ = v56
    } else {
        if false {
            v55 -= (v55 + 966)
            v3 = "foo bar"
            var v57 float32 = (0.1 / (((((-2.25 - 1e10) + 1.5) / 0.1) - 1.5) - 0.0))
            _ = v57
        } else {
            for fi58 := 0; fi58 < 5; fi58++ {
                if (0.1 == 0.1) {
                    fmt.Println(fi58, v55, fi58)
                    fmt.Println(v55, fi58, v55)
                }
                _ = fi58
            }
        }
    }
    delete(v4, -1)
    for fi59 := 0; fi59 < 6; fi59++ {
        v3 = v3
        _ = fi59
    }
    var v60 bool = true
    if (65535 <= (v2 & uint(v55))) {
        v1 = append(v1, ("hello" + ""))
        switch 2147483647 {
        case 32767:
            if true {
                switch (("" + ("foo bar" + (v3 + ((v3 + v3) + (v3 + v3))))) + v3) {
                case "a":
                    v55 -= (^(v55))
                default:
                    fmt.Println(v55, v3, v3)
                }
            }
        default:
            if (0 <= 127) {
                switch "a" {
                case "\n\t":
                    v4[-1] = 2147483647
                }
                v61 := GenS7{F0: -172, F1: 42, F2: (^(v2))}
                _ = v61
            } else {
                v62 := GenS8{F0: -1}
                _ = v62
            }
        }
        var v63 bool = false
        _ = v63
    }
    if ((v55 >> ((786) & 63)) != uint64(v2)) {
        for fi64 := 0; fi64 < 5; fi64++ {
            fmt.Println(v55)
            if len(v1) > 0 {
                v1[0] = ""
            }
            _ = fi64
        }
        v55 = (v55 & v55)
    } else {
        var v65 uint64 = v55
        _ = v65
    }
    fmt.Println(v3)
    var v66 float64 = 1e10
    fmt.Println(v3)
    var v67 uint64 = (^(v55))
    if ((1e10 >= -2.25) && false) {
        v68 := GenS8{F0: -128}
        _ = v68
    } else {
        switch 32768 {
        case 196:
            v69 := GenS3{F0: 1, F1: 18446744073709551615, F2: 128}
            _ = v69
        case -9223372036854775808:
            var v70 int32 = 1
            _ = v70
        case 132:
            switch -965 {
            case 9223372036854775807:
                fmt.Println(v67, v55, v2)
            case 2147483647:
                if true {
                    fmt.Println(v3)
                    v3 = ("foo bar" + v3)
                } else {
                    if (v2 != v2) {
                        if (0 == uint8(v55)) {
                            v66 = v66
                        } else {
                            v3 = v3
                            v3 = v3
                        }
                        switch v3 {
                        case "a":
                            v55 += 4294967295
                        }
                        v71 := GenS9{F0: genFn0((0 >= 0)), F1: 1}
                        _ = v71
                    }
                    v2 = (v2 & v2)
                    v72 := GenS1{F0: v60, F1: 128}
                    _ = v72
                }
            default:
                fmt.Println(v2, v3, v2)
            }
        }
        fmt.Println(v55)
        for fi73 := 0; fi73 < 15; fi73++ {
            fmt.Println(v60)
            var v74 int = (^(fi73))
            v2 += (v2 + v2)
            _ = fi73
            _ = v74
        }
    }
    v4[128] = 80
    for fi75 := 0; fi75 < 11; fi75++ {
        v4[(fi75 + -1)] = int(v2)
        v55 += 872
        if v60 {
            var v76 uint8 = 2
            _ = v76
        } else {
            var v77 int32 = -329
            _ = v77
        }
        _ = fi75
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v55:", v55)
    fmt.Println("v60:", v60)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v60
    _ = v66
    _ = v67
}
