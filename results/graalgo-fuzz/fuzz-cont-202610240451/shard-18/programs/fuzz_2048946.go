package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 string }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 int64; F2 uint }
func genFn0() uint16 {
    var v7 int8 = 42
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint32 = 1
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS1{F0: 0, F1: -128, F2: -2147483648}
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    if false {
        switch 127 {
        case 9223372036854775807:
            var v12 uint64 = 801
            _ = v12
        case 127:
            v7 -= (-(v7))
        case -391:
            if (!false) {
                v13 := GenS1{F0: 65535, F1: 42, F2: 474}
                var v14 int32 = -2147483648
                var v15 uint64 = 1
                _ = v13
                _ = v14
                _ = v15
            }
        default:
            v16 := []int{128, 0, 42, 32767}
            _ = v16
        }
        v17 := []int{434, -111, 32768, -1}
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := []int{(fi18 + (-(fi18))), fi18}
        var v20 int = (fi18 & 127)
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = v7
    return 1
}

func main() {
    var v1 int32 = 0
    v2 := make(map[int]int)
    var v3 uint64 = 4294967295
    var v4 int16 = 1
    var v5 uint16 = 32767
    var v6 int64 = 42
    fmt.Println(v1, v4)
    switch ("a" + "") {
    case "foo bar":
        if (90 <= 127) {
            v5 = 2
            _ = len(v2)
            switch ("foo bar" + (("foo bar" + (("a" + ("foo bar" + "a")) + ("hello" + ("a" + "a")))) + "hello")) {
            case "foo bar":
                fmt.Println(v3)
            }
        }
    case "hello":
        v5 = 42
    default:
        var v21 int64 = v6
        _ = v21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 string = "a"
        _ = fi22
        _ = v23
    }
    fmt.Println(v1)
    switch "a" {
    case "":
        var v24 uint16 = 1
        _ = v24
    case "foo bar":
        v25 := GenS2{F0: false, F1: 2, F2: "hello"}
        _ = v25
    }
    v6 = v6
    if ((!false) && true) {
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 uint16 = v5
            if (true && (!(42 >= v6))) {
                v3 += uint64(v5)
                var v28 bool = (!true)
                _ = v28
            }
            switch 1 {
            case -302:
                v29 := GenS3{F0: v6, F1: v6}
                _ = v29
            case 0:
                for fi30 := 0; fi30 < 16; fi30++ {
                    v31 := []int{int(v27)}
                    _ = fi30
                    _ = v31
                }
            case -483:
                for fi32 := 0; fi32 < 10; fi32++ {
                    v33 := []int{}
                    var v34 float32 = float32(-2.25)
                    v4 = v4
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
            _ = fi26
            _ = v27
        }
        var v35 uint16 = v5
        for fi36 := 0; fi36 < 17; fi36++ {
            var v37 bool = true
            _ = fi36
            _ = v37
        }
        _ = v35
    } else {
        fmt.Println(v6, v5)
        var v38 string = "hello"
        fmt.Println(v5, v4, v5)
        _ = v38
    }
    var v39 uint64 = (v3 % (((v3 >> ((1) & 63))) | 1))
    for fi40 := 0; fi40 < 8; fi40++ {
        for fi41 := 0; fi41 < 19; fi41++ {
            v2[int(v4)] = (fi41 >> ((42) & 63))
            _ = fi41
        }
        var v42 float64 = -2.25
        var v43 uint64 = (v39 | 4294967295)
        _ = fi40
        _ = v42
        _ = v43
    }
    var v44 float32 = 1.5
    var v45 int8 = 1
    var v46 int16 = (^(v4))
    var v47 int8 = int8(v5)
    if (941 <= 42) {
        v2[-9223372036854775808] = -682
    } else {
        fmt.Println(v4, v39)
        switch "foo bar" {
        case "\n\t":
            for fi48 := 0; fi48 < 8; fi48++ {
                fmt.Println(v45)
                _ = fi48
            }
        case "foo bar":
            v45 += (v45 ^ (v47 * 21))
        case "a":
            if (false || (755 < 2147483648)) {
                if false {
                    var v49 bool = false
                    var v50 int64 = 204
                    if (v49 && true) {
                        for fi51 := 0; fi51 < 13; fi51++ {
                            for fi52 := 0; fi52 < 4; fi52++ {
                                v46 = (^(v4))
                                var v53 int8 = -128
                                _ = fi52
                                _ = v53
                            }
                            v46 = (v4 % ((32767) | 1))
                            _ = fi51
                        }
                    }
                    _ = v49
                    _ = v50
                } else {
                    v54 := GenS3{F0: int64(v39), F1: v6}
                    _ = v54
                }
            }
        default:
            if true {
                v1 -= v1
            } else {
                var v55 string = "a"
                if (false && (!(!false))) {
                    v56 := GenS2{F0: false, F1: (v5 >> ((2) & 63)), F2: (("foo bar" + (v55 + "a")) + v55)}
                    switch -9223372036854775808 {
                    case 845:
                        var v57 float64 = (1e10 * (0.1 + 1.5))
                        _ = v57
                    case -875:
                        v6 -= (v6 & (v6 & v6))
                    case -73:
                        v58 := []int{9223372036854775807, 1}
                        _ = v58
                    default:
                        var v59 bool = (true || false)
                        _ = v59
                    }
                    _ = v56
                }
                if true {
                    switch 42 {
                    case 30:
                        switch v55 {
                        case "":
                            v44 = 3.14159
                        }
                    default:
                        v60 := GenS4{F0: (^(v1))}
                        _ = v60
                    }
                    var v61 int = -2147483648
                    _ = v61
                } else {
                    if (v55 == v55) {
                        var v62 float32 = (((v44 - (3.14159 * 0.0)) * 3.14159) + 3.14159)
                        fmt.Println(v47)
                        _ = v62
                    } else {
                        var v63 float32 = v44
                        v64 := GenS3{F0: v6, F1: -1}
                        fmt.Println(v5, v55, v3)
                        _ = v63
                        _ = v64
                    }
                }
                _ = v55
            }
        }
        v65 := GenS5{F0: (-(v1)), F1: (v6 * (-(v6))), F2: 348}
        _ = v65
    }
    var v66 uint8 = 255
    for fi67 := 0; fi67 < 2; fi67++ {
        for fi68 := 0; fi68 < 14; fi68++ {
            v69 := []int{}
            var v70 uint = 298
            v5 -= v5
            _ = fi68
            _ = v69
            _ = v70
        }
        switch ("hello" + "\n\t") {
        case "a":
            var v71 uint8 = uint8(v3)
            _ = v71
        default:
            for fi72 := 0; fi72 < 16; fi72++ {
                var v73 uint32 = 1
                var v74 int64 = 1
                _ = fi72
                _ = v73
                _ = v74
            }
        }
        fmt.Println(v39)
        _ = fi67
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v66
}
