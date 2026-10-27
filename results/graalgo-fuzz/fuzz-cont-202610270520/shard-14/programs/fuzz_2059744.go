package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }
type GenS2 struct { F0 uint; F1 int8; F2 int64 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 int8 }
func genFn0(p0 string) int32 {
    var v4 uint8 = 95
    v5 := []int{32768}
    p0 = (p0 + (p0 + "\n\t"))
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint64 = 585
        var v8 uint16 = uint16(v7)
        p0 = p0
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint16 = 1
    _ = p0
    _ = v4
    _ = v5
    _ = v9
    return 42
}
func genFn1(p0 uint32) uint64 {
    p0 += 123
    var v10 float32 = -2.25
    if false {
        for fi11 := 0; fi11 < 20; fi11++ {
            p0 -= p0
            var v12 int = (fi11 | 32767)
            _ = fi11
            _ = v12
        }
        var v13 int = -9223372036854775808
        _ = v13
    } else {
        var v14 uint32 = 2
        p0 = (v14 + p0)
        fmt.Println(v14)
        _ = v14
    }
    _ = p0
    _ = v10
    return 1
}

func main() {
    v1 := []int{2147483647, -2147483648, 521}
    v2 := make(map[uint32]uint32)
    var v3 float64 = 0.1
    v3 = v3
    switch "hello" {
    case "foo bar":
        if false {
            var v15 int16 = -895
            v3 = (((v3 + v3) / v3) + 0.1)
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := v2[0]
                var v18 uint32 = 2147483647
                v17 = 474
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v15
        }
    case "hello":
        v19 := GenS1{F0: 32768, F1: 392, F2: ("hello" + "a")}
        _ = v19
    case "\n\t":
        v20 := GenS1{F0: 903, F1: 65535, F2: "hello"}
        _ = v20
    default:
        v3 = v3
    }
    v21 := GenS2{F0: 42, F1: 42, F2: -365}
    v22 := GenS1{F0: 127, F1: 4294967295, F2: ("\n\t" + ("a" + "\n\t"))}
    switch "foo bar" {
    case "":
        v3 += (v3 / v3)
    case "hello":
        switch 916 {
        case -2147483648:
            if (1 != 255) {
                v2[4294967295] = 4294967295
                v3 = 0.1
            } else {
                v3 = float64((((-2.25 - float32((-2.25 / v3))) + (1.5 / 0.1)) + -2.25))
                v3 += 1.5
                v23 := GenS1{F0: -1, F1: 42, F2: ("" + "hello")}
                _ = v23
            }
        case 128:
            if true {
                v2[1] = 65535
            }
        }
    case "foo bar":
        switch -338 {
        case 128:
            v2[2] = 0
        case 501:
            fmt.Println(v3)
        default:
            fmt.Println(v3)
        }
    default:
        switch "a" {
        case "":
            v3 -= float64(1.5)
        case "hello":
            var v24 int32 = -1
            _ = v24
        default:
            if true {
                v3 -= 1e10
            }
        }
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        if len(v1) > 0 {
            v1[0] = (fi25 + 4294967295)
        }
        _ = fi25
    }
    var v26 uint16 = 446
    for fi27 := 0; fi27 < 15; fi27++ {
        if ("" == ("foo bar" + "")) {
            v26 = v26
            if (false && ((v26 - v26) == 0)) {
                var v28 uint8 = 42
                _ = v28
            }
            for fi29 := 0; fi29 < 9; fi29++ {
                v1 = append(v1, (fi27 << ((18446744073709551615) & 63)))
                v30 := GenS3{F0: 1, F1: 398, F2: 18446744073709551615}
                _ = fi29
                _ = v30
            }
        }
        for fi31 := 0; fi31 < 6; fi31++ {
            v3 = (float64(1e10) - v3)
            var v32 uint = 862
            var v33 bool = true
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi27
    }
    v3 = 3.14159
    var v34 int16 = -639
    if (v26 >= 1) {
        if true {
            for fi35 := 0; fi35 < 13; fi35++ {
                v3 -= v3
                _ = fi35
            }
            _ = len(v2)
            if (!(42 != 42)) {
                var v36 uint32 = 42
                for fi37 := 0; fi37 < 14; fi37++ {
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    _ = fi37
                }
                _ = v36
            }
        }
        var v38 string = "hello"
        v39 := GenS1{F0: 128, F1: 551, F2: (v38 + (("foo bar" + "foo bar") + "\n\t"))}
        _ = v38
        _ = v39
    }
    var v40 float64 = v3
    switch "" {
    case "":
        for fi41 := 0; fi41 < 2; fi41++ {
            switch -9223372036854775808 {
            case 9223372036854775807:
                fmt.Println(v26, fi41)
            case 32767:
                var v42 int8 = 42
                _ = v42
            }
            v34 = 819
            var v43 int8 = -1
            _ = fi41
            _ = v43
        }
    case "hello":
        v44 := GenS4{F0: 0}
        _ = v44
    default:
        v45 := GenS5{F0: false, F1: ((("\n\t" + "\n\t") + ("a" + ("hello" + ("hello" + "hello")))) + "a")}
        _ = v45
    }
    if false {
        v40 -= 0.0
        var v46 int64 = -616
        _ = v46
    } else {
        if ((true && (42 < genFn0("\n\t"))) && (1.5 != (((0.1 * float64(0.1)) - v3) + float64((1.5 - 1e10))))) {
            v40 = 3.14159
            v34 = 910
            if true {
                var v47 string = "hello"
                _ = v47
            }
        } else {
            var v48 bool = (!((!(v34 == v34)) || (!(float64(1e10) == (v3 / (((v40 / ((-2.25 + v3) - v40)) - v3) * v3))))))
            v3 -= (3.14159 - v3)
            _ = v48
        }
        v1 = append(v1, -810)
    }
    switch "foo bar" {
    case "\n\t":
        var v49 uint16 = (v26 / ((v26) | 1))
        _ = v49
    case "a":
        var v50 int16 = (v34 - v34)
        _ = v50
    }
    var v51 float32 = 0.1
    v52 := GenS6{F0: 0}
    v34 -= 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v40:", v40)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v34
    _ = v40
    _ = v51
    _ = v52
}
