package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 uint32; F2 uint16 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 int64; F1 int16; F2 int32 }
type GenS9 struct { F0 int; F1 bool; F2 int64 }
func genFn0() int8 {
    switch "hello" {
    case "a":
        var v7 uint16 = 146
        _ = v7
    case "hello":
        var v8 int = 32767
        _ = v8
    }
    var v9 bool = false
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 uint32 = 165
        v12 := []int{(fi10 - fi10), (-(fi10)), fi10, fi10}
        var v13 int32 = 127
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v9 = v9
    _ = v9
    return 0
}
func genFn1(p0 uint16) int8 {
    v14 := GenS1{F0: ("hello" + "hello"), F1: 1}
    v15 := GenS2{F0: 18446744073709551615, F1: -146}
    p0 -= (p0 * (p0 << ((18446744073709551615) & 63)))
    if (!false) {
        if true {
            var v16 int64 = 76
            p0 = 1
            var v17 int64 = -9223372036854775808
            _ = v16
            _ = v17
        } else {
            switch (("foo bar" + (("\n\t" + "a") + ("a" + "\n\t"))) + "hello") {
            case "foo bar":
                if (!true) {
                    switch "\n\t" {
                    case "":
                        for fi18 := 0; fi18 < 20; fi18++ {
                            var v19 bool = false
                            if (!v19) {
                                if v19 {
                                    if v19 {
                                        v20 := []int{}
                                        fmt.Println(fi18)
                                        v19 = v19
                                        _ = v20
                                    } else {
                                        var v21 bool = ((928 > 139) && (1 != 255))
                                        for fi22 := 0; fi22 < 2; fi22++ {
                                            v23 := GenS2{F0: 18446744073709551615, F1: -871}
                                            var v24 int8 = -41
                                            v24 += int8(p0)
                                            _ = fi22
                                            _ = v23
                                            _ = v24
                                        }
                                        v25 := GenS3{F0: (fi18 << ((65535) & 63))}
                                        _ = v21
                                        _ = v25
                                    }
                                } else {
                                    fmt.Println(fi18, fi18)
                                    p0 += p0
                                }
                                if false {
                                    v26 := []int{fi18, fi18}
                                    p0 -= (^(p0))
                                    _ = v26
                                } else {
                                    fmt.Println(p0, fi18)
                                }
                                var v27 float64 = ((3.14159 - float64(float32(1e10))) - float64(3.14159))
                                _ = v27
                            } else {
                                var v28 uint = 255
                                v28 = 255
                                _ = v28
                            }
                            _ = fi18
                            _ = v19
                        }
                    case "\n\t":
                        var v29 uint8 = 153
                        _ = v29
                    }
                } else {
                    var v30 uint64 = 18446744073709551615
                    var v31 uint = 18446744073709551615
                    _ = v30
                    _ = v31
                }
            case "a":
                v32 := []int{32768}
                _ = v32
            }
        }
    }
    _ = p0
    _ = v14
    _ = v15
    return 111
}

func main() {
    var v1 uint16 = 255
    v2 := []uint32{1, 2147483647, 42, 4294967295}
    var v3 bool = (0 >= 4294967295)
    var v4 int8 = 36
    v5 := make(map[string]uint32)
    var v6 uint32 = 42
    var v33 float64 = -2.25
    v34 := GenS4{F0: 4294967295}
    if ("" <= "\n\t") {
        var v35 bool = false
        _ = v35
    }
    v6 += v6
    _ = len(v2)
    if true {
        var v36 int8 = (v4 & (^(v4)))
        var v37 int64 = 946
        var v38 int16 = -1
        _ = v36
        _ = v37
        _ = v38
    }
    switch "hello" {
    case "":
        for fi39 := 0; fi39 < 14; fi39++ {
            for fi40 := 0; fi40 < 10; fi40++ {
                var v41 float32 = (0.0 + (3.14159 - 0.1))
                v42 := GenS5{F0: fi39, F1: v6, F2: (v1 ^ 65535)}
                v33 += (1e10 + v33)
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = len(v5)
            v2 = append(v2, v6)
            _ = fi39
        }
    case "foo bar":
        if false {
            var v43 float64 = 0.1
            _ = v43
        }
    case "hello":
        v3 = v3
    default:
        if false {
            if (v4 != v4) {
                if (v3 && ((v6 % ((4294967295) | 1)) > 2147483648)) {
                    if (!false) {
                        _ = len(v2)
                        fmt.Println(v1, v3, v6)
                    }
                    switch 32768 {
                    case -256:
                        for fi44 := 0; fi44 < 1; fi44++ {
                            var v45 int16 = 556
                            _ = fi44
                            _ = v45
                        }
                    case 2147483647:
                        v46 := GenS4{F0: 705}
                        _ = v46
                    case -2147483648:
                        _ = len(v2)
                    }
                } else {
                    if v3 {
                        for fi47 := 0; fi47 < 1; fi47++ {
                            v48 := GenS1{F0: "a", F1: (v6 >> ((65535) & 63))}
                            fmt.Println(v3, v1, v33)
                            _ = fi47
                            _ = v48
                        }
                        if v3 {
                            v49 := GenS2{F0: 382, F1: 42}
                            _ = v49
                        }
                        v3 = true
                    } else {
                        v5[("foo bar" + "")] = (v6 << ((181) & 63))
                        for fi50 := 0; fi50 < 15; fi50++ {
                            v3 = ((255 == 184) || (fi50 != fi50))
                            v51 := GenS6{F0: 189, F1: 0}
                            _ = fi50
                            _ = v51
                        }
                    }
                    for fi52 := 0; fi52 < 19; fi52++ {
                        for fi53 := 0; fi53 < 6; fi53++ {
                            if true {
                                v54 := GenS7{F0: v1, F1: fi52}
                                v3 = (!(0 > 1))
                                var v55 int64 = 1
                                _ = v54
                                _ = v55
                            } else {
                                if (false && (fi53 <= (fi52 >> ((402) & 63)))) {
                                    v3 = v3
                                    var v56 int16 = 42
                                    if true {
                                        fmt.Println(v3, v33, v33)
                                        if len(v2) > 0 {
                                            v2[0] = v6
                                        }
                                        var v57 int32 = -765
                                        _ = v57
                                    }
                                    _ = v56
                                } else {
                                    v58 := GenS3{F0: (-(fi53))}
                                    _ = v58
                                }
                            }
                            var v59 bool = (-370 != int64(fi53))
                            v3 = (int16(fi52) >= 42)
                            _ = fi53
                            _ = v59
                        }
                        v5["hello"] = v6
                        v1 -= (^(v1))
                        _ = fi52
                    }
                    var v60 int16 = -1
                    _ = v60
                }
                v5[(("" + "a") + ("hello" + (("" + ("foo bar" + ("a" + ""))) + (("a" + "\n\t") + "a"))))] = (^(v6))
                switch "a" {
                case "hello":
                    if (v6 >= (v6 >> ((2) & 63))) {
                        var v61 float32 = ((float32(1e10) - 1e10) - float32((v33 - (float64(0.1) + v33))))
                        v4 = v4
                        var v62 int16 = 941
                        _ = v61
                        _ = v62
                    }
                }
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v63 int64 = 9223372036854775807
            var v64 uint16 = uint16(v6)
            _ = v63
            _ = v64
        }
    }
    if v3 {
        var v65 uint32 = v6
        switch ("hello" + ("a" + "\n\t")) {
        case "foo bar":
            for fi66 := 0; fi66 < 11; fi66++ {
                v67 := GenS8{F0: 1, F1: 646, F2: 42}
                _ = fi66
                _ = v67
            }
        case "":
            fmt.Println(v3, v4)
        }
        switch ("foo bar" + ("\n\t" + "hello")) {
        case "":
            var v68 int32 = -2147483648
            _ = v68
        case "foo bar":
            v6 += 2
        default:
            v69 := GenS9{F0: -129, F1: ((^(v4)) != (v4 | (-(v4)))), F2: -879}
            _ = v69
        }
        _ = v65
    }
    fmt.Println(v4, v4, v1)
    _ = len(v2)
    for fi70 := 0; fi70 < 12; fi70++ {
        v71 := GenS6{F0: 2147483647, F1: -918}
        v6 = v6
        if ((v3 || (2147483647 < 42)) || ((1 >= 65535) || v3)) {
            v72 := GenS4{F0: (fi70 >> ((255) & 63))}
            switch (fi70 - -1) {
            case 580:
                var v73 int8 = -1
                _ = v73
            default:
                for fi74 := 0; fi74 < 4; fi74++ {
                    if true {
                        v2 = append(v2, v6)
                    } else {
                        v5["a"] = 717
                    }
                    var v75 int32 = 95
                    _ = fi74
                    _ = v75
                }
            }
            v33 = v33
            _ = v72
        } else {
            var v76 uint64 = 1
            fmt.Println(v3, v3, v6)
            _ = v76
        }
        _ = fi70
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
