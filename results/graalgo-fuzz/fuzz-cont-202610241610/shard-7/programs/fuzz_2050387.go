package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 int8; F2 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 int16; F1 int; F2 string }
func genFn0(p0 uint8, p1 uint8) int64 {
    var v7 uint = 4294967295
    switch "hello" {
    case "":
        for fi8 := 0; fi8 < 3; fi8++ {
            if (v7 != (v7 - v7)) {
                v9 := []int{(fi8 >> ((v7) & 63)), (fi8 ^ -103), (fi8 | fi8), fi8}
                _ = v9
            }
            switch fi8 {
            case 0:
                var v10 uint16 = 1
                _ = v10
            default:
                p1 = (^(p0))
            }
            _ = fi8
        }
    }
    if true {
        v11 := GenS1{F0: int(v7)}
        _ = v11
    } else {
        v12 := []int{1, 32767, -2147483648, -9223372036854775808}
        if true {
            fmt.Println(v7, v7, p1)
            if (float32(1.5) <= 1.5) {
                p0 = 75
                for fi13 := 0; fi13 < 16; fi13++ {
                    switch (fi13 << (((v7 >> (((v7 / (((v7 % ((v7) | 1))) | 1))) & 63))) & 63)) {
                    case 32768:
                        for fi14 := 0; fi14 < 6; fi14++ {
                            if false {
                                var v15 bool = true
                                v16 := GenS1{F0: fi13}
                                _ = v15
                                _ = v16
                            } else {
                                v17 := GenS1{F0: (fi14 % ((fi14) | 1))}
                                fmt.Println(v7)
                                _ = v17
                            }
                            var v18 int = -1
                            _ = fi14
                            _ = v18
                        }
                    default:
                        fmt.Println(p0)
                    }
                    var v19 bool = false
                    var v20 int32 = -2147483648
                    _ = fi13
                    _ = v19
                    _ = v20
                }
                v7 = (^(v7))
            }
        } else {
            if (!true) {
                p0 -= (p1 & p0)
                if (!false) {
                    switch "" {
                    case "a":
                        p1 = 42
                    case "\n\t":
                        p1 += (p1 & (p0 ^ p0))
                    case "":
                        switch 0 {
                        case 42:
                            var v21 int64 = 9223372036854775807
                            _ = v21
                        case 127:
                            v7 = (^(v7))
                        default:
                            var v22 uint16 = 560
                            _ = v22
                        }
                    }
                    var v23 uint = (v7 - 0)
                    var v24 string = ("" + (("hello" + "a") + "hello"))
                    _ = v23
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 18; fi25++ {
                        for fi26 := 0; fi26 < 15; fi26++ {
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            v27 := GenS1{F0: fi25}
                            v28 := GenS2{F0: 65535}
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        _ = fi25
                    }
                }
            } else {
                var v29 uint16 = 42
                _ = v29
            }
        }
        _ = v12
    }
    p1 = 2
    _ = p0
    _ = p1
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint8) uint64 {
    p0 = p0
    if (255 > 65535) {
        v30 := GenS2{F0: 255}
        v31 := []int{945, 461, 4294967295}
        _ = v30
        _ = v31
    } else {
        var v32 uint16 = 65535
        _ = v32
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        v34 := GenS3{F0: 42, F1: (p0 % ((p0) | 1))}
        var v35 int8 = 55
        _ = fi33
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        var v37 uint16 = 903
        _ = fi36
        _ = v37
    }
    switch -9223372036854775808 {
    case 2147483647:
        if (!(false || false)) {
            if false {
                v38 := []int{}
                if len(v38) > 0 {
                    v38[0] = 32767
                }
                _ = v38
            } else {
                p0 = (p0 | (^(p0)))
                for fi39 := 0; fi39 < 5; fi39++ {
                    var v40 int16 = 32767
                    if true {
                        var v41 float32 = 1e10
                        for fi42 := 0; fi42 < 10; fi42++ {
                            v43 := []int{}
                            _ = fi42
                            _ = v43
                        }
                        _ = v41
                    } else {
                        var v44 uint32 = 409
                        v45 := GenS4{F0: (fi39 == fi39)}
                        _ = v44
                        _ = v45
                    }
                    if (0 >= 2) {
                        var v46 int64 = genFn0(p0, (p0 >> ((18446744073709551615) & 63)))
                        _ = v46
                    } else {
                        var v47 int8 = -128
                        _ = v47
                    }
                    _ = fi39
                    _ = v40
                }
            }
        } else {
            var v48 float32 = float32(0.0)
            _ = v48
        }
    default:
        p0 = (^(p0))
    }
    _ = p0
    return 1
}

func main() {
    var v1 uint64 = 1
    var v2 float32 = (float32(3.14159) + float32(3.14159))
    v3 := make(map[int]uint32)
    var v4 float32 = 0.0
    v5 := []string{"hello", ""}
    var v6 uint8 = 42
    fmt.Println(v1, v1)
    var v49 int16 = -607
    for fi50 := 0; fi50 < 5; fi50++ {
        for fi51 := 0; fi51 < 9; fi51++ {
            var v52 int8 = -90
            switch (^(fi50)) {
            case -727:
                var v53 uint = 52
                _ = v53
            case 2147483647:
                var v54 uint64 = 734
                _ = v54
            }
            if false {
                v1 -= 317
            }
            v55 := GenS5{F0: 18446744073709551615, F1: v52, F2: (v49 ^ -489)}
            fmt.Println(v6, fi51, v52)
            _ = fi51
            _ = v52
            _ = v55
        }
        var v56 string = ""
        _ = fi50
        _ = v56
    }
    for fi57 := 0; fi57 < 10; fi57++ {
        v2 -= (v4 * ((1.5 + (v4 + ((v4 + float32(3.14159)) - float32((float64(((v2 * ((v4 / 1e10) - 0.0)) * (v4 - -2.25))) * 1e10))))) - 0.0))
        _ = fi57
    }
    var v58 uint = 2
    var v59 int64 = genFn0(127, 2)
    if (false && (58 > -128)) {
        var v60 int32 = -2147483648
        var v61 int64 = (v59 & v59)
        _ = v60
        _ = v61
    } else {
        if ((!true) || true) {
            v49 = (v49 & 0)
        } else {
            var v62 int = -508
            for fi63 := 0; fi63 < 16; fi63++ {
                for fi64 := 0; fi64 < 4; fi64++ {
                    fmt.Println(v2, v62, fi63)
                    if (!(v6 != v6)) {
                        v2 -= v4
                        switch (fi63 + (^(v62))) {
                        case 127:
                            switch ("foo bar" + "a") {
                            case "hello":
                                if false {
                                    if ((false || (-130 == (fi63 ^ (^(fi64))))) || (339 >= 2)) {
                                        delete(v3, (fi64 << ((v58) & 63)))
                                    }
                                    var v65 int16 = 32767
                                    fmt.Println(v1, v49, v58)
                                    _ = v65
                                }
                            case "":
                                delete(v3, fi64)
                            }
                        case 955:
                            v66 := GenS6{F0: false}
                            _ = v66
                        }
                    }
                    v67 := GenS7{F0: true, F1: 2147483647}
                    if true {
                        switch "a" {
                        case "":
                            v68 := GenS5{F0: v58, F1: 127, F2: (-(v49))}
                            _ = v68
                        }
                        v69 := GenS8{F0: (v49 & (v49 | int16(fi64))), F1: 0, F2: "hello"}
                        if (true && true) {
                            v58 = (v58 - v58)
                            if (true || (v6 == v6)) {
                                var v70 uint = v58
                                _ = v70
                            } else {
                                fmt.Println(v6, v62)
                                switch 1 {
                                case -9223372036854775808:
                                    v2 = (((0.1 - 1e10) - v2) / (3.14159 / 1.5))
                                default:
                                    var v71 float64 = -2.25
                                    _ = v71
                                }
                            }
                        } else {
                            var v72 uint = (v58 << ((v58) & 63))
                            _ = v72
                        }
                        var v73 int32 = 991
                        if false {
                            var v74 bool = (!true)
                            var v75 int64 = genFn0(v6, (v6 | v6))
                            v3[-540] = 2147483648
                            var v76 int8 = -128
                            _ = v74
                            _ = v75
                            _ = v76
                        } else {
                            fmt.Println(v4)
                        }
                        _ = v69
                        _ = v73
                    }
                    _ = fi64
                    _ = v67
                }
                var v77 uint16 = uint16(v1)
                v78 := GenS2{F0: v77}
                _ = fi63
                _ = v77
                _ = v78
            }
            var v79 int8 = 100
            _ = v62
            _ = v79
        }
    }
    fmt.Println(v2, v1, v59)
    var v80 uint16 = 744
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v80
}
