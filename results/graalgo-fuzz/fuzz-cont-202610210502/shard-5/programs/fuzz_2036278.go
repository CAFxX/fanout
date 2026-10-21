package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int; F1 string; F2 int64 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 int64; F1 int8; F2 uint64 }
type GenS7 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint32) uint {
    var v4 uint32 = p0
    if false {
        v4 -= v4
        v5 := []int{-9223372036854775808, 128, -9223372036854775808}
        _ = v5
    }
    p0 = (v4 | (p0 / ((42) | 1)))
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 float32 = 1.5
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v4
    return 18446744073709551615
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]string)
    var v3 uint8 = 1
    v3 = (v3 >> ((2) & 63))
    if false {
        if false {
            v3 = (^(v3))
            switch "a" {
            case "foo bar":
                var v8 uint32 = 2147483647
                _ = v8
            default:
                v9 := v2[-275]
                _ = v9
            }
        } else {
            v3 = v3
            if ((!(1 < 407)) && (v3 < 42)) {
                var v10 int16 = -1
                _ = v10
            } else {
                if (1 <= -26) {
                    switch "hello" {
                    case "":
                        for fi11 := 0; fi11 < 5; fi11++ {
                            fmt.Println(fi11, v3)
                            v3 += (v3 / (((v3 - (v3 | 128))) | 1))
                            var v12 int8 = 127
                            _ = fi11
                            _ = v12
                        }
                    case "a":
                        fmt.Println(v3)
                    case "\n\t":
                        for fi13 := 0; fi13 < 19; fi13++ {
                            v14 := GenS1{F0: 0}
                            _ = fi13
                            _ = v14
                        }
                    }
                    if true {
                        var v15 uint = 0
                        for fi16 := 0; fi16 < 15; fi16++ {
                            fmt.Println(v3, v3)
                            fmt.Println(v15, fi16)
                            if (42 == 42) {
                                if (true || (("" >= ("hello" + "a")) && ((-32768 > -1) && (341 < 4294967295)))) {
                                    if (-425 == -1) {
                                        var v17 int16 = 1000
                                        fmt.Println(v17, v3)
                                        _ = v17
                                    }
                                } else {
                                    if (!false) {
                                        v15 = (v15 | 0)
                                    } else {
                                        v3 -= (v3 | v3)
                                    }
                                    v15 = v15
                                }
                                for fi18 := 0; fi18 < 2; fi18++ {
                                    fmt.Println(v15, fi18, fi16)
                                    _ = len(v2)
                                    delete(v2, 4294967295)
                                    _ = fi18
                                }
                                switch ("a" + "") {
                                case "hello":
                                    for fi19 := 0; fi19 < 17; fi19++ {
                                        if false {
                                            if (-946 == 536) {
                                                delete(v2, (fi16 - 32767))
                                                v3 += 169
                                                var v20 int64 = 697
                                                _ = v20
                                            } else {
                                                switch (fi19 ^ (fi16 & 9223372036854775807)) {
                                                case 4294967295:
                                                    if false {
                                                        var v21 float32 = float32(1.5)
                                                        var v22 int16 = 32767
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                case 102:
                                                    var v23 float32 = 1.5
                                                    _ = v23
                                                }
                                            }
                                            var v24 string = (("\n\t" + ("\n\t" + "")) + "")
                                            var v25 uint = genFn0(65535)
                                            _ = v24
                                            _ = v25
                                        } else {
                                            v15 = v15
                                        }
                                        delete(v2, (^(fi16)))
                                        var v26 int16 = -395
                                        var v27 uint8 = (v3 >> ((0) & 63))
                                        var v28 bool = ((v15 % ((v15) | 1)) == genFn0(1))
                                        _ = fi19
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                case "\n\t":
                                    var v29 bool = (1 < 0)
                                    _ = v29
                                case "a":
                                    fmt.Println(v3, v3)
                                }
                            } else {
                                v30 := GenS1{F0: int16(v3)}
                                for fi31 := 0; fi31 < 4; fi31++ {
                                    v3 -= 127
                                    _ = fi31
                                }
                                _ = v30
                            }
                            _ = fi16
                        }
                        var v32 float32 = (-2.25 + (float32(0.1) * float32(0.1)))
                        _ = v15
                        _ = v32
                    }
                } else {
                    v3 = v3
                }
                for fi33 := 0; fi33 < 11; fi33++ {
                    for fi34 := 0; fi34 < 4; fi34++ {
                        v3 -= (v3 & (v3 - 4))
                        _ = fi34
                    }
                    v3 += v3
                    _ = fi33
                }
            }
        }
        var v35 float32 = 0.1
        _ = v35
    } else {
        for fi36 := 0; fi36 < 8; fi36++ {
            var v37 string = ""
            v38 := v2[(fi36 ^ (fi36 + fi36))]
            for fi39 := 0; fi39 < 5; fi39++ {
                v40 := GenS2{F0: (^(fi36)), F1: 0}
                if (true && false) {
                    fmt.Println(v37)
                    var v41 uint = 693
                    _ = v41
                }
                fmt.Println(v3, fi39, v38)
                _ = fi39
                _ = v40
            }
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    switch -9223372036854775808 {
    case -1:
        fmt.Println(v3)
    case 2147483647:
        if true {
            var v42 bool = false
            v43 := v2[9223372036854775807]
            _ = v42
            _ = v43
        }
    default:
        for fi44 := 0; fi44 < 13; fi44++ {
            fmt.Println(fi44)
            v45 := GenS3{F0: 1, F1: 2147483648}
            _ = fi44
            _ = v45
        }
    }
    v3 += v3
    if (-385 == 42) {
        if false {
            var v46 uint = 2
            var v47 int = -346
            var v48 uint8 = 166
            _ = v46
            _ = v47
            _ = v48
        } else {
            var v49 int16 = 0
            _ = v49
        }
        for fi50 := 0; fi50 < 10; fi50++ {
            if (-1 == 711) {
                if (!false) {
                    for fi51 := 0; fi51 < 5; fi51++ {
                        fmt.Println(v3, fi51)
                        var v52 int = -413
                        _ = fi51
                        _ = v52
                    }
                    if (1 < -9223372036854775808) {
                        v3 += (v3 - (^(v3)))
                        v53 := GenS4{F0: 154, F1: "a", F2: 42}
                        _ = v53
                    } else {
                        var v54 uint64 = 42
                        _ = v54
                    }
                    var v55 uint = 18446744073709551615
                    _ = v55
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v56 uint16 = 1
                    if len(v1) > 0 {
                        v1[0] = 759
                    }
                    _ = v56
                }
                var v57 float64 = 1e10
                if (18446744073709551615 > 2) {
                    var v58 int16 = 32767
                    var v59 float64 = (v57 + float64(0.1))
                    _ = v58
                    _ = v59
                } else {
                    if false {
                        var v60 int16 = 1000
                        if true {
                            v60 -= 7
                            switch fi50 {
                            case -463:
                                fmt.Println(v60)
                            case 127:
                                switch "" {
                                case "":
                                    switch "a" {
                                    case "\n\t":
                                        for fi61 := 0; fi61 < 18; fi61++ {
                                            if true {
                                                if false {
                                                    v62 := v2[(fi61 + 32767)]
                                                    fmt.Println(v3, v62, v57)
                                                    v63 := GenS1{F0: v60}
                                                    _ = v62
                                                    _ = v63
                                                }
                                                v57 -= (1e10 + v57)
                                                var v64 bool = true
                                                switch "hello" {
                                                case "hello":
                                                    fmt.Println(v64)
                                                case "foo bar":
                                                    v64 = (889 < 587)
                                                case "":
                                                    switch (fi50 ^ (fi50 ^ fi61)) {
                                                    case 0:
                                                        v65 := GenS3{F0: -9223372036854775808, F1: 2147483648}
                                                        _ = v65
                                                    case 42:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        fmt.Println(fi61, v60, v60)
                                                    default:
                                                        switch "hello" {
                                                        case "a":
                                                            var v66 uint64 = 0
                                                            _ = v66
                                                        case "\n\t":
                                                            v67 := GenS5{F0: fi61, F1: 877}
                                                            _ = v67
                                                        case "hello":
                                                            v68 := GenS5{F0: 32768, F1: 2}
                                                            _ = v68
                                                        }
                                                    }
                                                }
                                                _ = v64
                                            }
                                            _ = fi61
                                        }
                                    default:
                                        var v69 int = 0
                                        _ = v69
                                    }
                                case "a":
                                    if true {
                                        v2[0] = "foo bar"
                                        v3 = v3
                                        delete(v2, fi50)
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = 2
                                        }
                                        fmt.Println(v60)
                                    }
                                default:
                                    var v70 int = -875
                                    _ = v70
                                }
                            default:
                                var v71 int = (fi50 * fi50)
                                _ = v71
                            }
                        } else {
                            v72 := GenS3{F0: -9223372036854775808, F1: 1}
                            var v73 uint8 = v3
                            for fi74 := 0; fi74 < 17; fi74++ {
                                switch "hello" {
                                case "":
                                    fmt.Println(v3, v73)
                                case "a":
                                    v75 := GenS6{F0: -9223372036854775808, F1: 0, F2: 18446744073709551615}
                                    _ = v75
                                }
                                v3 = v73
                                var v76 int16 = (v60 + (v60 * v60))
                                _ = fi74
                                _ = v76
                            }
                            _ = v72
                            _ = v73
                        }
                        fmt.Println(v60, fi50, v60)
                        _ = v60
                    } else {
                        fmt.Println(v57, fi50)
                    }
                }
                _ = v57
            }
            _ = fi50
        }
    }
    fmt.Println(v3)
    if true {
        v77 := GenS7{F0: true, F1: 32767}
        _ = v77
    } else {
        var v78 uint = 345
        _ = v78
    }
    if (false && (((1.5 - 0.1) / -2.25) <= 1e10)) {
        var v79 int32 = -2147483648
        _ = v79
    }
    var v80 uint16 = uint16(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v80
}
