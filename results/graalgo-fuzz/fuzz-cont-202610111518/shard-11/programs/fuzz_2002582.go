package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 string; F1 uint; F2 int }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int16; F1 uint32 }
func genFn0(p0 int64, p1 bool) uint16 {
    if (!p1) {
        if p1 {
            if (false || (p1 && p1)) {
                p1 = p1
                var v6 float64 = 0.1
                _ = v6
            }
            fmt.Println(p1, p1)
            for fi7 := 0; fi7 < 15; fi7++ {
                v8 := GenS1{F0: 127, F1: 42}
                fmt.Println(p1, p0)
                _ = fi7
                _ = v8
            }
        }
    } else {
        var v9 uint = 1
        if p1 {
            p1 = (0 != -1)
            p0 -= (p0 ^ p0)
        }
        _ = v9
    }
    p1 = p1
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := []int{fi10, 32767, (fi10 % ((fi10) | 1)), (fi10 * fi10)}
        if ((("\n\t" + "hello") + ("hello" + "\n\t")) > "") {
            var v12 float32 = 1.5
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := GenS2{F0: "", F1: 340, F2: (fi10 - fi13)}
                v11 = append(v11, 1)
                var v15 int = 0
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        } else {
            fmt.Println(p0)
            for fi16 := 0; fi16 < 1; fi16++ {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                p1 = (p1 && false)
                p1 = (p1 && p1)
                _ = fi16
            }
        }
        var v17 uint = 792
        _ = fi10
        _ = v11
        _ = v17
    }
    var v18 float64 = 3.14159
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v18
    return 32767
}
func genFn1(p0 uint8) uint64 {
    p0 -= (^(p0))
    if false {
        v19 := []int{32768, 127, 32767, 1}
        _ = v19
    }
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    var v1 float32 = (float32(1e10) * 0.1)
    v2 := make(map[string]int)
    var v3 int16 = 42
    var v4 float64 = 0.1
    v5 := make(map[string]int)
    v20 := GenS2{F0: ("" + (("hello" + (("a" + ("hello" + "foo bar")) + "\n\t")) + ("" + ""))), F1: 1, F2: 0}
    if false {
        v21 := []int{350, 32768, 0, 42}
        v22 := GenS1{F0: 33, F1: 77}
        _ = v21
        _ = v22
    }
    var v23 uint64 = 782
    v23 = (v23 * 4294967295)
    v24 := []int{140}
    v4 = v4
    if false {
        var v25 uint16 = 255
        var v26 bool = true
        if v26 {
            fmt.Println(v3)
        }
        _ = v25
        _ = v26
    } else {
        fmt.Println(v4, v1, v3)
    }
    if true {
        if false {
            v27 := GenS1{F0: -128, F1: 1}
            v4 = 0.1
            _ = v27
        }
        switch -9223372036854775808 {
        case 323:
            var v28 uint16 = genFn0(-1, false)
            _ = v28
        case -606:
            var v29 int8 = -128
            _ = v29
        }
        v30 := GenS1{F0: -127, F1: 127}
        _ = v30
    } else {
        if true {
            switch "foo bar" {
            case "\n\t":
                v31 := GenS3{F0: true, F1: "a"}
                _ = v31
            case "a":
                v4 = 1.5
            }
            v24 = append(v24, 127)
        }
        switch int(v23) {
        case 42:
            switch -2147483648 {
            case 953:
                var v32 uint64 = v23
                _ = v32
            case -2147483648:
                if false {
                    _ = len(v24)
                }
            case 42:
                var v33 uint8 = 115
                _ = v33
            }
        default:
            switch "" {
            case "hello":
                for fi34 := 0; fi34 < 13; fi34++ {
                    fmt.Println(v1, v1)
                    v35 := GenS2{F0: "", F1: 2, F2: fi34}
                    _ = fi34
                    _ = v35
                }
            case "foo bar":
                var v36 float32 = 0.0
                _ = v36
            default:
                for fi37 := 0; fi37 < 10; fi37++ {
                    v3 -= 42
                    v38 := v2["\n\t"]
                    switch "hello" {
                    case "\n\t":
                        v39 := GenS4{F0: v38, F1: -98}
                        _ = v39
                    case "hello":
                        var v40 int16 = v3
                        _ = v40
                    case "":
                        var v41 float32 = (3.14159 * float32(3.14159))
                        _ = v41
                    }
                    _ = fi37
                    _ = v38
                }
            }
        }
    }
    fmt.Println(v4, v23, v23)
    for fi42 := 0; fi42 < 5; fi42++ {
        if (1 > -2147483648) {
            v3 += -1
        }
        if true {
            if (true && false) {
                var v43 int32 = 2147483647
                _ = len(v24)
                _ = v43
            } else {
                for fi44 := 0; fi44 < 19; fi44++ {
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                    _ = fi44
                }
                v45 := GenS5{F0: (^(v3)), F1: 874}
                var v46 string = ""
                _ = v45
                _ = v46
            }
            switch -9223372036854775808 {
            case -9223372036854775808:
                var v47 int = -1
                _ = v47
            }
        } else {
            v4 -= v4
        }
        _ = fi42
    }
    switch 1 {
    case -150:
        if false {
            if false {
                var v48 int8 = 1
                _ = v48
            }
        }
    case 42:
        v23 = (v23 + 391)
    case 32767:
        if false {
            if (uint16(v23) != 771) {
                var v49 uint16 = 838
                v23 = (v23 & (^(v23)))
                if true {
                    for fi50 := 0; fi50 < 9; fi50++ {
                        v51 := GenS4{F0: fi50, F1: -1}
                        _ = fi50
                        _ = v51
                    }
                    switch "hello" {
                    case "":
                        fmt.Println(v23, v23)
                    default:
                        var v52 int32 = -1
                        _ = v52
                    }
                }
                _ = v49
            }
        }
    default:
        v1 = 0.1
    }
    for fi53 := 0; fi53 < 11; fi53++ {
        for fi54 := 0; fi54 < 14; fi54++ {
            v1 += ((v1 * v1) - (1.5 - ((-2.25 * (((float32((v4 * 1.5)) + v1) * v1) * -2.25)) / v1)))
            var v55 uint = 18446744073709551615
            _ = fi54
            _ = v55
        }
        v56 := GenS3{F0: true, F1: "foo bar"}
        for fi57 := 0; fi57 < 10; fi57++ {
            var v58 int64 = -9223372036854775808
            v58 += (v58 * v58)
            v59 := GenS1{F0: 0, F1: 128}
            _ = fi57
            _ = v58
            _ = v59
        }
        _ = fi53
        _ = v56
    }
    switch 372 {
    case -9223372036854775808:
        var v60 float64 = 1.5
        _ = v60
    }
    var v61 int64 = 1
    for fi62 := 0; fi62 < 20; fi62++ {
        var v63 bool = ((!((((v3 / ((v3) | 1)) < 428) && false) && (uint8(v23) < 2))) && false)
        for fi64 := 0; fi64 < 14; fi64++ {
            v1 = v1
            fmt.Println(v63)
            var v65 int16 = (v3 - v3)
            _ = fi64
            _ = v65
        }
        _ = fi62
        _ = v63
    }
    var v66 int16 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v61:", v61)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v66
}
