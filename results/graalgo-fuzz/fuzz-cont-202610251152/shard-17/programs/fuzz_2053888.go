package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint8, p1 int) int64 {
    var v5 uint16 = 188
    var v6 string = ("" + "a")
    v5 = (^(v5))
    for fi7 := 0; fi7 < 18; fi7++ {
        p1 = -1
        for fi8 := 0; fi8 < 2; fi8++ {
            switch "foo bar" {
            case "hello":
                v9 := []int{(fi7 << ((891) & 63)), (fi8 + fi7), (fi7 >> ((854) & 63)), fi7}
                _ = v9
            case "foo bar":
                var v10 int16 = -32768
                _ = v10
            default:
                var v11 uint = 255
                _ = v11
            }
            p1 -= (^(fi7))
            v12 := []int{fi8}
            _ = fi8
            _ = v12
        }
        v5 -= 964
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 513
}
func genFn1(p0 uint8, p1 bool) uint64 {
    for fi13 := 0; fi13 < 16; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := GenS1{F0: 0, F1: 9223372036854775807}
            fmt.Println(p0, fi13)
            switch "hello" {
            case "":
                var v16 int32 = 1
                p0 -= p0
                _ = v16
            case "a":
                v17 := []int{fi13, fi13, (fi14 | (fi14 & -9223372036854775808))}
                _ = v17
            case "\n\t":
                var v18 int32 = -715
                _ = v18
            }
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    var v19 int64 = genFn0((p0 & (p0 << ((1) & 63))), 42)
    v20 := []int{602, -2147483648}
    v19 += v19
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 42
}

func main() {
    var v1 int8 = -21
    var v2 int32 = 42
    v3 := []int{826, 32768}
    v4 := make(map[string]bool)
    v1 += (v1 & v1)
    v21 := GenS1{F0: 1, F1: 0}
    v22 := GenS1{F0: 255, F1: -1}
    v3 = append(v3, 0)
    var v23 string = "a"
    v1 = v1
    var v24 int16 = -1
    for fi25 := 0; fi25 < 1; fi25++ {
        v23 = (v23 + ((((v23 + "\n\t") + v23) + "") + "\n\t"))
        v26 := GenS2{F0: 1, F1: 185}
        _ = fi25
        _ = v26
    }
    var v27 float32 = 1.5
    var v28 uint32 = 818
    if (true && false) {
        if true {
            fmt.Println(v24, v24, v23)
        } else {
            v29 := v4["a"]
            switch -487 {
            case 2147483647:
                fmt.Println(v28, v28, v28)
            case 61:
                var v30 int32 = v2
                _ = v30
            default:
                var v31 uint32 = v28
                _ = v31
            }
            for fi32 := 0; fi32 < 20; fi32++ {
                if len(v3) > 0 {
                    v3[0] = fi32
                }
                v27 += -2.25
                fmt.Println(fi32)
                _ = fi32
            }
            _ = v29
        }
        for fi33 := 0; fi33 < 13; fi33++ {
            v34 := GenS1{F0: 42, F1: 42}
            _ = fi33
            _ = v34
        }
        var v35 uint = 0
        _ = v35
    } else {
        if (!false) {
            v2 = v2
        } else {
            v23 = "hello"
            for fi36 := 0; fi36 < 3; fi36++ {
                for fi37 := 0; fi37 < 16; fi37++ {
                    fmt.Println(v23, fi37, v1)
                    var v38 string = ((v23 + (((v23 + v23) + v23) + v23)) + "")
                    _ = fi37
                    _ = v38
                }
                switch (v23 + v23) {
                case "":
                    switch (("" + (v23 + "foo bar")) + "a") {
                    case "hello":
                        v2 = (v2 - v2)
                    case "\n\t":
                        var v39 int16 = v24
                        _ = v39
                    default:
                        if (false || (72 <= 45)) {
                            var v40 int64 = genFn0(56, int(v2))
                            _ = v40
                        } else {
                            v27 = 1.5
                        }
                    }
                case "hello":
                    if false {
                        for fi41 := 0; fi41 < 4; fi41++ {
                            var v42 int8 = v1
                            if true {
                                v42 -= (v1 / (((^(v42))) | 1))
                            } else {
                                v1 -= v42
                                v43 := GenS1{F0: 18446744073709551615, F1: -369}
                                _ = v43
                            }
                            _ = fi41
                            _ = v42
                        }
                        if (221 == -9223372036854775808) {
                            for fi44 := 0; fi44 < 7; fi44++ {
                                switch fi36 {
                                case 4294967295:
                                    var v45 uint64 = 1
                                    _ = v45
                                case 1:
                                    v28 = (v28 * 814)
                                case 0:
                                    v46 := GenS3{F0: 337}
                                    _ = v46
                                default:
                                    switch fi36 {
                                    case -359:
                                        v23 = v23
                                    case 32767:
                                        v27 = (0.0 + 1e10)
                                    }
                                }
                                if (!false) {
                                    v47 := GenS4{F0: v23}
                                    _ = v47
                                } else {
                                    fmt.Println(fi44, v2, fi36)
                                    v24 = (-(v24))
                                }
                                v48 := GenS5{F0: -9223372036854775808, F1: (v28 + (v28 & v28))}
                                if false {
                                    var v49 uint64 = genFn1(255, (!true))
                                    var v50 bool = true
                                    v1 -= (-(v1))
                                    _ = v49
                                    _ = v50
                                }
                                var v51 int8 = v1
                                _ = fi44
                                _ = v48
                                _ = v51
                            }
                            if len(v3) > 0 {
                                v3[0] = (fi36 - (fi36 >> ((741) & 63)))
                            }
                        } else {
                            var v52 float32 = v27
                            _ = v52
                        }
                        if false {
                            var v53 string = (v23 + "hello")
                            _ = v53
                        } else {
                            delete(v4, v23)
                            v28 += v28
                        }
                    }
                case "foo bar":
                    v3 = append(v3, (-(fi36)))
                default:
                    var v54 int64 = genFn0(0, fi36)
                    _ = v54
                }
                if len(v3) > 0 {
                    v3[0] = fi36
                }
                _ = fi36
            }
            v28 = v28
        }
        switch -1 {
        case 1:
            v28 = (v28 >> ((4294967295) & 63))
        }
    }
    if (802 >= 2) {
        var v55 int16 = (^(v24))
        fmt.Println(v23)
        if true {
            for fi56 := 0; fi56 < 13; fi56++ {
                var v57 int16 = v24
                var v58 int8 = -115
                _ = fi56
                _ = v57
                _ = v58
            }
        } else {
            if (0.0 < 1e10) {
                fmt.Println(v23, v24)
            }
            v59 := GenS5{F0: 9223372036854775807, F1: (v28 ^ v28)}
            var v60 string = v23
            _ = v59
            _ = v60
        }
        _ = v55
    } else {
        v61 := GenS5{F0: genFn0(2, -150), F1: 2147483647}
        v24 += (v24 & v24)
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
