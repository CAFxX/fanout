package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int; F1 string; F2 uint8 }
type GenS3 struct { F0 bool; F1 int64; F2 int64 }
type GenS4 struct { F0 int64; F1 int; F2 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8; F1 uint8 }
func genFn0() uint8 {
    var v5 uint = 18446744073709551615
    for fi6 := 0; fi6 < 19; fi6++ {
        switch "foo bar" {
        case "":
            fmt.Println(fi6, fi6)
        default:
            if (!(true || false)) {
                v5 = (^(v5))
                v5 -= (v5 % ((18446744073709551615) | 1))
            }
        }
        v7 := GenS1{F0: "", F1: fi6}
        _ = fi6
        _ = v7
    }
    switch ("a" + "foo bar") {
    case "":
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint64 = uint64(v5)
            _ = fi8
            _ = v9
        }
    case "a":
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 float64 = float64((0.0 * 1.5))
            if true {
                var v12 int8 = 1
                _ = v12
            }
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 string = "\n\t"
                fmt.Println(v11, v5)
                _ = fi13
                _ = v14
            }
            _ = fi10
            _ = v11
        }
    case "hello":
        v5 += 126
    }
    switch 9223372036854775807 {
    case 2147483647:
        v5 -= 42
    case 42:
        if true {
            var v15 uint = (v5 * v5)
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 uint64 = 190
                v18 := []int{(-(fi16)), (fi16 & fi16), (fi16 >> ((255) & 63))}
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v15
        } else {
            v5 = (v5 & v5)
            v19 := []int{855, 9223372036854775807}
            _ = v19
        }
    default:
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 int64 = 1
            v22 := GenS2{F0: (fi20 | (fi20 & (-(fi20)))), F1: "", F2: 2}
            fmt.Println(v5, v21)
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    fmt.Println(v5)
    _ = v5
    return 0
}
func genFn1(p0 bool, p1 int32) uint64 {
    v23 := GenS3{F0: true, F1: 0, F2: 42}
    var v24 uint8 = 128
    v24 = 128
    var v25 int64 = 9223372036854775807
    switch "" {
    case "\n\t":
        if ((v24 & 255) < (^(v24))) {
            var v26 int64 = v25
            var v27 uint = 228
            v28 := GenS4{F0: v25, F1: 32768, F2: 599}
            _ = v26
            _ = v27
            _ = v28
        } else {
            p1 = (p1 * (p1 * p1))
            for fi29 := 0; fi29 < 13; fi29++ {
                var v30 uint = 65535
                fmt.Println(v24, v25, v24)
                _ = fi29
                _ = v30
            }
            for fi31 := 0; fi31 < 15; fi31++ {
                v25 = (v25 % ((858) | 1))
                _ = fi31
            }
        }
    case "hello":
        v32 := []int{}
        _ = v32
    default:
        fmt.Println(p0, p1, p1)
    }
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    _ = v25
    return 42
}

func main() {
    var v1 int64 = -1
    var v2 uint16 = 566
    var v3 float32 = 3.14159
    v4 := make(map[uint32]int)
    var v33 uint64 = 683
    var v34 int64 = (v1 >> ((18446744073709551615) & 63))
    switch "\n\t" {
    case "\n\t":
        for fi35 := 0; fi35 < 17; fi35++ {
            v36 := []int{(^(fi35)), fi35}
            for fi37 := 0; fi37 < 10; fi37++ {
                fmt.Println(v34)
                switch (fi37 >> ((42) & 63)) {
                case 512:
                    var v38 int16 = 32767
                    _ = v38
                case 0:
                    v2 = (v2 + v2)
                    v39 := GenS5{F0: "hello"}
                    _ = v39
                case 128:
                    var v40 uint32 = 42
                    _ = v40
                default:
                    if false {
                        if (170 == 97) {
                            if (float32(1.5) > v3) {
                                v34 -= 0
                                delete(v4, uint32(v2))
                                if false {
                                    v4[42] = 2147483647
                                }
                                if false {
                                    v1 += 0
                                    var v41 int8 = -43
                                    _ = v41
                                }
                            }
                        }
                        var v42 float64 = 0.0
                        _ = v42
                    } else {
                        v33 -= v33
                        if ((-(fi37)) < int(v33)) {
                            fmt.Println(fi37, fi37, v3)
                        }
                    }
                }
                v43 := GenS6{F0: 456}
                _ = fi37
                _ = v43
            }
            _ = fi35
            _ = v36
        }
    default:
        var v44 int16 = 42
        _ = v44
    }
    var v45 uint32 = 2
    v3 = v3
    for fi46 := 0; fi46 < 8; fi46++ {
        v47 := v4[0]
        _ = fi46
        _ = v47
    }
    var v48 int = 512
    if false {
        if true {
            if (498 < 1) {
                v49 := []int{}
                v1 -= v1
                if false {
                    for fi50 := 0; fi50 < 17; fi50++ {
                        var v51 bool = (-2.25 > v3)
                        _ = fi50
                        _ = v51
                    }
                    if true {
                        v33 = v33
                        switch "" {
                        case "a":
                            v33 = (v33 << ((65535) & 63))
                        default:
                            switch "a" {
                            case "a":
                                v48 = v48
                            case "":
                                v34 += 42
                            }
                        }
                        if false {
                            v52 := v4[v45]
                            _ = v52
                        }
                    } else {
                        v53 := GenS3{F0: (!(!(0.0 == 0.1))), F1: (v34 * 1), F2: v1}
                        v45 = (v45 + (v45 ^ (v45 ^ (v45 >> ((556) & 63)))))
                        var v54 uint16 = v2
                        _ = v53
                        _ = v54
                    }
                } else {
                    for fi55 := 0; fi55 < 10; fi55++ {
                        delete(v4, 239)
                        var v56 uint8 = 255
                        _ = fi55
                        _ = v56
                    }
                    fmt.Println(v45)
                }
                _ = v49
            } else {
                var v57 int16 = -898
                if (true || ((0.1 - 1e10) <= -2.25)) {
                    v45 = (v45 ^ v45)
                }
                v48 -= v48
                _ = v57
            }
        } else {
            v58 := v4[(v45 - 2147483648)]
            if false {
                var v59 uint16 = 2
                v60 := GenS3{F0: (55 > 2), F1: (v1 + (^(v34))), F2: v1}
                _ = v59
                _ = v60
            } else {
                v61 := GenS4{F0: v1, F1: v48, F2: (v48 ^ v48)}
                if ((true && (v34 == v1)) || false) {
                    var v62 uint16 = 42
                    switch v58 {
                    case 127:
                        v2 += (v2 & 42)
                    default:
                        if ((v33 == uint64(v62)) && false) {
                            v63 := GenS7{F0: int32(v48)}
                            _ = v63
                        } else {
                            fmt.Println(v48, v45)
                        }
                    }
                    _ = v62
                } else {
                    _ = len(v4)
                }
                _ = v61
            }
            _ = v58
        }
    }
    for fi64 := 0; fi64 < 2; fi64++ {
        var v65 string = "hello"
        v3 = 1.5
        _ = fi64
        _ = v65
    }
    fmt.Println(v3, v2, v34)
    var v66 int8 = -93
    switch -100 {
    case 128:
        for fi67 := 0; fi67 < 10; fi67++ {
            v2 = 32767
            v34 = v34
            if false {
                delete(v4, v45)
                v68 := v4[v45]
                var v69 uint64 = v33
                _ = v68
                _ = v69
            }
            _ = fi67
        }
    default:
        switch "a" {
        case "\n\t":
            for fi70 := 0; fi70 < 12; fi70++ {
                if ((4294967295 < 894) && true) {
                    v33 = (^(v33))
                    var v71 bool = (3.14159 != 0.1)
                    delete(v4, v45)
                    _ = v71
                } else {
                    var v72 bool = (float32(0.1) >= -2.25)
                    v73 := GenS8{F0: 6}
                    _ = v72
                    _ = v73
                }
                v74 := GenS9{F0: -2, F1: 0}
                var v75 int64 = v34
                _ = fi70
                _ = v74
                _ = v75
            }
        }
    }
    switch "\n\t" {
    case "foo bar":
        switch (-(v48)) {
        case 164:
            v76 := []int{}
            _ = v76
        case 1:
            switch -9223372036854775808 {
            case -603:
                var v77 string = "hello"
                _ = v77
            case 32767:
                v34 = (v1 + 0)
            }
        }
    case "a":
        if true {
            v3 -= (v3 + 0.1)
            switch 853 {
            case 42:
                v1 -= v34
            case 32767:
                v4[v45] = v48
            case 0:
                if false {
                    var v78 uint16 = (v2 - (v2 << ((127) & 63)))
                    _ = v78
                }
            }
            var v79 int32 = 1
            _ = v79
        } else {
            var v80 uint = 18446744073709551615
            switch ("a" + (("foo bar" + "hello") + ("hello" + "\n\t"))) {
            case "foo bar":
                v3 = (((0.0 - ((((0.0 + 0.0) / (v3 - v3)) - v3) + (0.0 / -2.25))) - v3) / (v3 * ((v3 / v3) * v3)))
            }
            for fi81 := 0; fi81 < 13; fi81++ {
                if false {
                    switch (("a" + "hello") + "\n\t") {
                    case "foo bar":
                        v48 += fi81
                    case "":
                        v66 = int8(v80)
                    default:
                        v45 = v45
                    }
                    v45 = 4294967295
                } else {
                    var v82 float32 = v3
                    _ = v82
                }
                v83 := GenS3{F0: (623 < (v80 - (v80 - v80))), F1: v34, F2: (v1 + 45)}
                _ = fi81
                _ = v83
            }
            _ = v80
        }
    case "":
        switch (v48 >> ((42) & 63)) {
        case 2147483647:
            fmt.Println(v33, v45, v3)
        default:
            v2 = 2
        }
    }
    var v84 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v66:", v66)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v45
    _ = v48
    _ = v66
    _ = v84
}
