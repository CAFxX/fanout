package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 bool; F1 uint; F2 int16 }
type GenS5 struct { F0 uint64; F1 uint8; F2 string }
type GenS6 struct { F0 int64; F1 string; F2 int32 }
func genFn0() int32 {
    if false {
        var v7 bool = (false && false)
        _ = v7
    } else {
        switch (("\n\t" + "") + "hello") {
        case "foo bar":
            var v8 int8 = 100
            _ = v8
        case "a":
            v9 := GenS1{F0: 42}
            _ = v9
        case "\n\t":
            var v10 uint8 = 255
            _ = v10
        }
        v11 := GenS1{F0: 42}
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        switch "a" {
        case "hello":
            var v13 uint16 = 1
            _ = v13
        default:
            v14 := []int{32768, (fi12 & fi12), (-(fi12)), fi12}
            _ = v14
        }
        v15 := GenS1{F0: -833}
        var v16 int64 = -479
        _ = fi12
        _ = v15
        _ = v16
    }
    var v17 uint64 = 42
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 int8 = 55
        var v20 bool = true
        var v21 int16 = 484
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    v17 = (v17 % ((572) | 1))
    _ = v17
    return -2147483648
}
func genFn1(p0 bool, p1 uint8) uint64 {
    if true {
        var v22 string = ("hello" + ("a" + "foo bar"))
        _ = v22
    }
    v23 := GenS1{F0: 207}
    _ = p0
    _ = p1
    _ = v23
    return 4294967295
}

func main() {
    var v1 int8 = 99
    var v2 uint32 = 4294967295
    var v3 float64 = (float64(0.0) + (((-2.25 / 1.5) / 1e10) - 3.14159))
    v4 := []int{}
    var v5 string = "a"
    var v6 uint = 112
    v2 += 4294967295
    var v24 bool = true
    var v25 int16 = -32768
    switch ("hello" + v5) {
    case "foo bar":
        for fi26 := 0; fi26 < 12; fi26++ {
            for fi27 := 0; fi27 < 2; fi27++ {
                v25 = v25
                if (!v24) {
                    var v28 float32 = (((float32((v3 * v3)) - 0.0) * float32((v3 * v3))) / float32(v3))
                    _ = v28
                }
                _ = fi27
            }
            fmt.Println(v3, v1, v24)
            _ = fi26
        }
    case "":
        switch "\n\t" {
        case "\n\t":
            switch -2147483648 {
            case 0:
                var v29 string = v5
                _ = v29
            case 42:
                fmt.Println(v25)
            default:
                for fi30 := 0; fi30 < 13; fi30++ {
                    fmt.Println(v2, v5, v25)
                    _ = fi30
                }
            }
        case "":
            switch -906 {
            case -1:
                for fi31 := 0; fi31 < 6; fi31++ {
                    switch fi31 {
                    case -1:
                        v32 := GenS2{F0: (^(v2))}
                        _ = v32
                    case 42:
                        var v33 bool = true
                        _ = v33
                    default:
                        fmt.Println(v25, fi31, v6)
                    }
                    _ = fi31
                }
            case 1:
                for fi34 := 0; fi34 < 19; fi34++ {
                    fmt.Println(v5, v3, v3)
                    fmt.Println(v24)
                    _ = fi34
                }
            case 128:
                v3 += v3
            }
        case "hello":
            for fi35 := 0; fi35 < 19; fi35++ {
                fmt.Println(fi35)
                _ = fi35
            }
        }
    }
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    v24 = v24
    for fi36 := 0; fi36 < 13; fi36++ {
        v2 += (^(v2))
        if v24 {
            var v37 uint32 = (^(v2))
            v25 -= (v25 >> ((926) & 63))
            v24 = ((false || (172 < 1)) || v24)
            _ = v37
        } else {
            if v24 {
                v2 -= (^(v2))
                switch (v5 + ("\n\t" + "a")) {
                case "\n\t":
                    var v38 uint32 = v2
                    _ = v38
                default:
                    v39 := GenS3{F0: v24, F1: 0}
                    _ = v39
                }
            } else {
                var v40 bool = ((v2 | v2) <= (v2 | (v2 - (v2 + (v2 & v2)))))
                var v41 string = (v5 + ("\n\t" + v5))
                switch fi36 {
                case 42:
                    if (476 <= 2) {
                        for fi42 := 0; fi42 < 18; fi42++ {
                            var v43 uint32 = v2
                            v2 += (v2 | v43)
                            _ = fi42
                            _ = v43
                        }
                        var v44 float32 = float32((float64(float32(v3)) / (0.0 + v3)))
                        _ = v44
                    }
                default:
                    if v24 {
                        v45 := GenS1{F0: v25}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v46 float64 = 3.14159
                        _ = v45
                        _ = v46
                    }
                }
                _ = v40
                _ = v41
            }
            v5 = v5
            for fi47 := 0; fi47 < 1; fi47++ {
                v5 = v5
                v6 = v6
                var v48 uint16 = 124
                _ = fi47
                _ = v48
            }
        }
        _ = fi36
    }
    for fi49 := 0; fi49 < 9; fi49++ {
        v50 := GenS4{F0: (4294967295 == 234), F1: v6, F2: (v25 - (v25 * v25))}
        _ = fi49
        _ = v50
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v51 := GenS3{F0: false, F1: 42}
    if (v24 || (930 != 2)) {
        var v52 int16 = (^(v25))
        v52 -= v52
        _ = v52
    } else {
        if (!false) {
            for fi53 := 0; fi53 < 10; fi53++ {
                if (42 > genFn1(true, 0)) {
                    switch -9223372036854775808 {
                    case 9223372036854775807:
                        _ = len(v4)
                    case 697:
                        for fi54 := 0; fi54 < 16; fi54++ {
                            var v55 uint8 = 228
                            fmt.Println(v25, v2)
                            _ = fi54
                            _ = v55
                        }
                    case 689:
                        fmt.Println(v2, v5, v24)
                    }
                    fmt.Println(v6, v24, v3)
                    fmt.Println(v1)
                }
                v1 -= v1
                v2 -= v2
                _ = fi53
            }
        }
    }
    var v56 uint8 = 255
    for fi57 := 0; fi57 < 10; fi57++ {
        v58 := GenS5{F0: 0, F1: (v56 >> ((1) & 63)), F2: "foo bar"}
        v4 = append(v4, 524)
        if true {
            fmt.Println(v2, fi57)
            switch v5 {
            case "\n\t":
                _ = len(v4)
            case "foo bar":
                v4 = append(v4, (fi57 << ((65535) & 63)))
            case "hello":
                if len(v4) > 0 {
                    v4[0] = -1
                }
            default:
                var v59 uint32 = (v2 + (v2 ^ 42))
                _ = v59
            }
        }
        _ = fi57
        _ = v58
    }
    var v60 bool = ((^(v2)) != (v2 - v2))
    fmt.Println(v60, v56, v24)
    switch 9223372036854775807 {
    case 42:
        var v61 uint = (v6 & (v6 >> ((77) & 63)))
        _ = v61
    case 1:
        v4 = append(v4, 32768)
    default:
        if (v60 && (0 > -2147483648)) {
            v62 := GenS2{F0: (v2 ^ 1)}
            v63 := GenS6{F0: 388, F1: v5, F2: int32(v6)}
            var v64 float32 = -2.25
            _ = v62
            _ = v63
            _ = v64
        } else {
            var v65 bool = false
            _ = len(v4)
            _ = v65
        }
    }
    if true {
        if len(v4) > 0 {
            v4[0] = -496
        }
        switch int(v25) {
        case 213:
            v56 += (^(v56))
        case 4294967295:
            if true {
                var v66 string = "\n\t"
                _ = v66
            }
        default:
            v4 = append(v4, -2147483648)
        }
        if (-760 == -32768) {
            var v67 int16 = -368
            v3 -= v3
            _ = v67
        }
    }
    v24 = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v56
    _ = v6
    _ = v60
}
