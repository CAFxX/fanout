package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint; F1 string; F2 uint }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint }
func genFn0() int32 {
    if (!true) {
        var v6 int16 = -32768
        var v7 uint32 = 4294967295
        v8 := []int{-890, 1}
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 int8 = 42
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int16 = 1
        switch 2147483647 {
        case 589:
            if true {
                for fi12 := 0; fi12 < 20; fi12++ {
                    v13 := []int{}
                    switch (-(fi12)) {
                    case -485:
                        if false {
                            _ = len(v13)
                        } else {
                            v11 += 1
                            v14 := GenS1{F0: false, F1: (2 == 255)}
                            _ = v14
                        }
                    case 128:
                        v11 = v11
                    default:
                        var v15 int32 = 1
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                }
                if false {
                    v9 = v9
                } else {
                    switch ("\n\t" + ("foo bar" + "")) {
                    case "hello":
                        var v16 string = ("a" + "a")
                        _ = v16
                    case "foo bar":
                        if (false && (0 > 625)) {
                            v17 := GenS2{F0: 317, F1: "a", F2: 375}
                            var v18 uint16 = 823
                            v11 += (v11 + v11)
                            _ = v17
                            _ = v18
                        } else {
                            v11 = v11
                            switch (-(fi10)) {
                            case 42:
                                if false {
                                    v19 := GenS3{F0: true, F1: -1}
                                    v11 -= v11
                                    _ = v19
                                }
                            case 32768:
                                var v20 int32 = -828
                                _ = v20
                            case 127:
                                if (2 == 255) {
                                    var v21 int16 = (v11 << ((65535) & 63))
                                    v22 := []int{fi10, (fi10 - (fi10 ^ (-(fi10))))}
                                    if false {
                                        v21 += (v11 & (v11 * v21))
                                        for fi23 := 0; fi23 < 19; fi23++ {
                                            fmt.Println(fi10, v21, v9)
                                            v9 -= -1
                                            v9 -= 100
                                            v11 += (v21 - (-(v11)))
                                            _ = fi23
                                        }
                                    }
                                    _ = v21
                                    _ = v22
                                } else {
                                    if (!true) {
                                        if (0 >= 65535) {
                                            var v24 uint8 = 255
                                            v25 := []int{fi10, fi10, fi10, (^(fi10))}
                                            _ = v24
                                            _ = v25
                                        } else {
                                            switch 1 {
                                            case 962:
                                                v26 := []int{fi10, -1, fi10, (fi10 ^ fi10)}
                                                _ = v26
                                            }
                                        }
                                    }
                                    if false {
                                        v27 := GenS3{F0: false, F1: 1}
                                        v9 = (-(v9))
                                        _ = v27
                                    } else {
                                        var v28 int32 = 425
                                        for fi29 := 0; fi29 < 5; fi29++ {
                                            v28 -= (v28 % (((v28 << ((704) & 63))) | 1))
                                            _ = fi29
                                        }
                                        var v30 int8 = (v9 * (v9 ^ (v9 % (((v9 - v9)) | 1))))
                                        _ = v28
                                        _ = v30
                                    }
                                }
                            }
                        }
                    case "\n\t":
                        var v31 uint16 = 0
                        _ = v31
                    default:
                        var v32 int32 = 182
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 16; fi33++ {
                        fmt.Println(v11, v11, v11)
                        switch (-(fi33)) {
                        case -2147483648:
                            v34 := []int{(fi10 << ((20) & 63)), (fi10 ^ (fi33 + 4294967295)), fi33}
                            _ = v34
                        case 127:
                            var v35 float64 = 0.1
                            _ = v35
                        case 2147483647:
                            v36 := []int{32767}
                            _ = v36
                        }
                        var v37 uint64 = 2
                        var v38 int8 = -128
                        _ = fi33
                        _ = v37
                        _ = v38
                    }
                    v11 += (-(v11))
                }
            }
        default:
            var v39 uint16 = 65535
            _ = v39
        }
        if ((!true) && false) {
            var v40 uint = 42
            for fi41 := 0; fi41 < 8; fi41++ {
                var v42 float64 = 1e10
                v40 = v40
                v11 += v11
                var v43 uint64 = 0
                _ = fi41
                _ = v42
                _ = v43
            }
            v40 = (v40 * (v40 - v40))
            _ = v40
        } else {
            fmt.Println(v9, fi10)
        }
        _ = fi10
        _ = v11
    }
    v44 := []int{4294967295}
    v9 -= (v9 << ((42) & 63))
    _ = v44
    _ = v9
    return -103
}
func genFn1(p0 int64, p1 int8) uint8 {
    p0 = p0
    var v45 int32 = -1
    var v46 uint32 = 721
    switch "hello" {
    case "hello":
        switch -9223372036854775808 {
        case -979:
            switch "foo bar" {
            case "hello":
                fmt.Println(p0, v45, v46)
            default:
                v46 = v46
            }
        }
    }
    _ = p0
    _ = p1
    _ = v45
    _ = v46
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 9223372036854775807
    var v3 int16 = -32768
    var v4 uint32 = 2147483647
    var v5 uint16 = 2
    for fi47 := 0; fi47 < 5; fi47++ {
        var v48 int64 = (v2 ^ v2)
        _ = fi47
        _ = v48
    }
    v4 = (v4 ^ (v4 | (v4 & 2147483647)))
    if (1e10 > 1.5) {
        v1[("hello" + "a")] = (61 == 127)
    } else {
        v49 := GenS3{F0: (v3 == (v3 % ((v3) | 1))), F1: 42}
        _ = v49
    }
    v50 := GenS4{F0: 1}
    for fi51 := 0; fi51 < 14; fi51++ {
        if false {
            v4 = v4
            v2 = 9223372036854775807
            v4 -= (v4 ^ v4)
        }
        _ = fi51
    }
    switch "a" {
    case "hello":
        for fi52 := 0; fi52 < 12; fi52++ {
            if (("a" + ("foo bar" + "hello")) <= ("" + "\n\t")) {
                for fi53 := 0; fi53 < 20; fi53++ {
                    if (true && false) {
                        var v54 string = ""
                        _ = v54
                    } else {
                        if false {
                            var v55 int8 = int8(v2)
                            v55 -= (v55 << ((0) & 63))
                            _ = v55
                        } else {
                            var v56 int = 63
                            _ = v56
                        }
                        v5 = (v5 * v5)
                        v2 -= v2
                    }
                    _ = fi53
                }
                if ((0.1 / 1e10) == float64(3.14159)) {
                    for fi57 := 0; fi57 < 11; fi57++ {
                        v58 := GenS5{F0: fi52}
                        delete(v1, "hello")
                        _ = fi57
                        _ = v58
                    }
                } else {
                    var v59 string = "hello"
                    _ = v59
                }
            } else {
                v60 := GenS1{F0: true, F1: (555 <= 203)}
                v61 := []int{fi52, fi52, fi52, (fi52 & (fi52 * (fi52 + fi52)))}
                _ = v60
                _ = v61
            }
            for fi62 := 0; fi62 < 4; fi62++ {
                v63 := []int{(fi62 ^ fi52)}
                _ = fi62
                _ = v63
            }
            _ = fi52
        }
    case "a":
        if (1 != 2) {
            if true {
                var v64 float64 = 0.0
                v65 := GenS6{F0: 4294967295}
                v66 := GenS7{F0: 739}
                _ = v64
                _ = v65
                _ = v66
            } else {
                v5 = 32767
            }
        }
    default:
        v5 = (^(v5))
    }
    v3 -= v3
    for fi67 := 0; fi67 < 10; fi67++ {
        v3 = (v3 << ((255) & 63))
        _ = fi67
    }
    v2 += (v2 % (((v2 / (((v2 << ((4294967295) & 63))) | 1))) | 1))
    if (127 == 1) {
        var v68 string = (("a" + "foo bar") + "hello")
        v69 := v1[((v68 + "hello") + "a")]
        for fi70 := 0; fi70 < 8; fi70++ {
            v4 = (^(v4))
            delete(v1, ("" + v68))
            var v71 uint8 = 196
            _ = fi70
            _ = v71
        }
        _ = v68
        _ = v69
    }
    v72 := GenS4{F0: -12}
    for fi73 := 0; fi73 < 10; fi73++ {
        if ((!((("hello" + "") + "hello") < "\n\t")) && true) {
            v3 -= (v3 & (v3 + 1000))
            fmt.Println(v4, v2)
            v74 := v1["hello"]
            _ = v74
        } else {
            v75 := []int{32767, fi73}
            var v76 uint = 42
            _ = v75
            _ = v76
        }
        v77 := GenS2{F0: 2, F1: "", F2: 1}
        _ = fi73
        _ = v77
    }
    v5 = v5
    var v78 int = 562
    if (!true) {
        switch "\n\t" {
        case "hello":
            v79 := []int{32767, (v78 & v78), (-(v78))}
            _ = v79
        case "foo bar":
            v2 = v2
        default:
            v1[""] = true
        }
        v80 := GenS5{F0: int(v2)}
        v78 += v78
        _ = v80
    } else {
        if (v78 < -572) {
            v81 := GenS5{F0: -549}
            v82 := []int{v78, v78, 1, (v78 ^ (^(v78)))}
            _ = v81
            _ = v82
        }
        v2 = 9223372036854775807
        fmt.Println(v4, v4, v2)
    }
    var v83 bool = true
    var v84 uint = 42
    v85 := GenS4{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v78:", v78)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v72
    _ = v78
    _ = v83
    _ = v84
    _ = v85
}
