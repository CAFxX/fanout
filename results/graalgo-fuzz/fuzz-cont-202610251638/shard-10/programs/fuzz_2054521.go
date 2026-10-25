package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int16 }
type GenS2 struct { F0 string; F1 uint; F2 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 uint }
type GenS4 struct { F0 int; F1 string; F2 uint8 }
type GenS5 struct { F0 bool; F1 uint16; F2 int64 }
type GenS6 struct { F0 string; F1 int; F2 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8; F1 uint8 }
type GenS9 struct { F0 int; F1 uint16; F2 uint64 }
func genFn0(p0 uint8) string {
    v7 := GenS1{F0: 42, F1: (false || true), F2: -555}
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            fmt.Println(p0, fi8)
            fmt.Println(p0)
            var v10 uint32 = 2
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS1{F0: 65535, F1: (0 == 0), F2: -277}
        v13 := []int{(fi11 * (fi11 % ((fi11) | 1)))}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 bool = ("a" <= ("foo bar" + "hello"))
    v15 := []int{32768}
    _ = p0
    _ = v14
    _ = v15
    _ = v7
    return "a"
}
func genFn1(p0 int8, p1 bool, p2 int) uint32 {
    var v16 uint8 = 204
    for fi17 := 0; fi17 < 10; fi17++ {
        if false {
            v18 := GenS2{F0: "\n\t", F1: 18446744073709551615, F2: 492}
            _ = v18
        }
        if p1 {
            v19 := []int{}
            _ = v19
        } else {
            fmt.Println(p2)
            switch "foo bar" {
            case "\n\t":
                if p1 {
                    v20 := GenS1{F0: 556, F1: ((340 > 245) && true), F2: -483}
                    _ = v20
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        v21 := []int{fi17, (p2 % (((fi17 + p2)) | 1)), (^(fi17))}
                        _ = v21
                    case "hello":
                        v22 := GenS2{F0: "foo bar", F1: 0, F2: 565}
                        _ = v22
                    default:
                        var v23 int32 = 2147483647
                        _ = v23
                    }
                    if false {
                        v24 := GenS2{F0: "", F1: 65535, F2: 768}
                        if false {
                            var v25 uint64 = 2
                            _ = v25
                        } else {
                            if false {
                                if (2147483647 < 65535) {
                                    var v26 string = "a"
                                    _ = v26
                                } else {
                                    fmt.Println(v16, v16)
                                    for fi27 := 0; fi27 < 9; fi27++ {
                                        fmt.Println(p0)
                                        var v28 bool = p1
                                        var v29 uint64 = 223
                                        _ = fi27
                                        _ = v28
                                        _ = v29
                                    }
                                }
                                switch ("" + ("" + "")) {
                                case "\n\t":
                                    v30 := []int{}
                                    _ = v30
                                case "":
                                    var v31 int16 = -32768
                                    _ = v31
                                }
                            }
                            var v32 float64 = 3.14159
                            var v33 string = ("\n\t" + ("foo bar" + ("\n\t" + "")))
                            _ = v32
                            _ = v33
                        }
                        _ = v24
                    } else {
                        if p1 {
                            p0 -= -11
                        } else {
                            var v34 uint = 1
                            _ = v34
                        }
                        fmt.Println(p1, v16, p2)
                        if (!(!(p1 || p1))) {
                            v35 := GenS2{F0: ("foo bar" + "a"), F1: 1, F2: 18446744073709551615}
                            for fi36 := 0; fi36 < 19; fi36++ {
                                v16 -= v16
                                if (!p1) {
                                    v37 := []int{(fi36 << ((99) & 63)), p2, 0}
                                    p1 = (1.5 >= (1e10 / 1.5))
                                    if (-200 > 42) {
                                        if false {
                                            var v38 uint16 = 65535
                                            if p1 {
                                                if (true || p1) {
                                                    var v39 float64 = 0.1
                                                    p2 = 1
                                                    _ = v39
                                                } else {
                                                    if len(v37) >= 2 {
                                                        v37 = v37[1:len(v37)]
                                                    }
                                                    fmt.Println(p0, p0)
                                                    var v40 int8 = (p0 + p0)
                                                    _ = v40
                                                }
                                                var v41 string = ""
                                                if true {
                                                    p0 -= 1
                                                    switch v41 {
                                                    case "a":
                                                        if (0 >= 1) {
                                                            p1 = false
                                                        }
                                                    case "hello":
                                                        var v42 float32 = (((1e10 - float32(0.0)) / (0.1 * -2.25)) + -2.25)
                                                        _ = v42
                                                    }
                                                    if ("hello" < "") {
                                                        if true {
                                                            var v43 int8 = p0
                                                            _ = v43
                                                        }
                                                        var v44 uint32 = 1
                                                        _ = v44
                                                    } else {
                                                        if len(v37) > 0 {
                                                            v37[0] = fi36
                                                        }
                                                        if len(v37) >= 2 {
                                                            v37 = v37[1:len(v37)]
                                                        }
                                                        var v45 uint32 = 509
                                                        _ = v45
                                                    }
                                                }
                                                switch (("foo bar" + "foo bar") + v41) {
                                                case "hello":
                                                    var v46 int8 = p0
                                                    _ = v46
                                                default:
                                                    switch v41 {
                                                    case "\n\t":
                                                        var v47 int = fi36
                                                        _ = v47
                                                    case "a":
                                                        fmt.Println(fi17, p0, v16)
                                                    case "foo bar":
                                                        if len(v37) >= 2 {
                                                            v37 = v37[1:len(v37)]
                                                        }
                                                    default:
                                                        p0 = p0
                                                    }
                                                }
                                                if (!p1) {
                                                    fmt.Println(p1)
                                                    v48 := GenS3{F0: 2, F1: 254, F2: 0}
                                                    _ = v48
                                                }
                                                _ = v41
                                            }
                                            fmt.Println(fi36, p1, p0)
                                            _ = v38
                                        }
                                        if len(v37) >= 2 {
                                            v37 = v37[1:len(v37)]
                                        }
                                        var v49 uint64 = 42
                                        _ = v49
                                    }
                                    _ = v37
                                }
                                _ = fi36
                            }
                            _ = v35
                        }
                    }
                    p2 += 9223372036854775807
                }
            default:
                if false {
                    switch "hello" {
                    case "hello":
                        p1 = p1
                    case "a":
                        for fi50 := 0; fi50 < 3; fi50++ {
                            var v51 uint8 = 0
                            var v52 int16 = -1
                            p2 = p2
                            v16 = (v16 % ((v16) | 1))
                            _ = fi50
                            _ = v51
                            _ = v52
                        }
                    case "foo bar":
                        v53 := []int{(fi17 << ((482) & 63)), fi17}
                        _ = v53
                    default:
                        p1 = (!p1)
                    }
                }
            }
            var v54 uint64 = 43
            _ = v54
        }
        _ = fi17
    }
    for fi55 := 0; fi55 < 7; fi55++ {
        v56 := []int{fi55, 9223372036854775807, p2, (fi55 << ((65535) & 63))}
        var v57 float32 = (1e10 * (1e10 + (3.14159 / float32(-2.25))))
        v57 = 0.0
        _ = fi55
        _ = v56
        _ = v57
    }
    var v58 int16 = -32768
    v59 := GenS4{F0: p2, F1: "hello", F2: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v58
    _ = v59
    return 1
}

func main() {
    var v1 string = ("a" + "hello")
    var v2 float32 = 0.1
    var v3 float32 = v2
    var v4 float32 = 0.1
    v5 := make(map[string]uint32)
    var v6 bool = true
    v60 := []int{-1, 128, 1}
    if len(v60) >= 2 {
        v60 = v60[1:len(v60)]
    }
    switch (v1 + "a") {
    case "":
        v61 := GenS5{F0: (0 == 0), F1: 42, F2: -414}
        _ = v61
    }
    var v62 bool = v6
    if (-2.25 == 1.5) {
        v4 -= v3
        v62 = v62
    }
    v63 := GenS6{F0: v1, F1: 32767, F2: 42}
    var v64 uint = 1
    v65 := GenS1{F0: 255, F1: v62, F2: 1}
    for fi66 := 0; fi66 < 3; fi66++ {
        _ = len(v5)
        v62 = (!(21 >= 65535))
        fmt.Println(v3, v2, v1)
        _ = fi66
    }
    if true {
        v6 = (1 > 9223372036854775807)
        fmt.Println(v6, v4)
        if len(v60) > 0 {
            v60[0] = 32768
        }
    }
    v2 += -2.25
    switch "hello" {
    case "a":
        var v67 int16 = 0
        _ = v67
    case "hello":
        for fi68 := 0; fi68 < 9; fi68++ {
            v69 := GenS2{F0: "\n\t", F1: 1, F2: 422}
            _ = fi68
            _ = v69
        }
    case "":
        for fi70 := 0; fi70 < 18; fi70++ {
            if v62 {
                v71 := GenS7{F0: 0}
                _ = v71
            }
            v5[v1] = 42
            v72 := GenS2{F0: (v1 + v1), F1: 722, F2: v64}
            _ = fi70
            _ = v72
        }
    default:
        switch ("hello" + v1) {
        case "a":
            var v73 bool = ((!(0 != 116)) || v6)
            _ = v73
        case "foo bar":
            fmt.Println(v64, v6)
        case "":
            v74 := GenS8{F0: 12, F1: 127}
            _ = v74
        }
    }
    switch -334 {
    case -114:
        if v6 {
            v3 -= 0.0
        }
    case 128:
        var v75 bool = false
        _ = v75
    case 603:
        v76 := GenS9{F0: 9223372036854775807, F1: 499, F2: 0}
        _ = v76
    default:
        var v77 uint8 = 23
        _ = v77
    }
    if len(v60) > 0 {
        v60[0] = 955
    }
    v6 = v6
    if true {
        v62 = false
    }
    for fi78 := 0; fi78 < 13; fi78++ {
        v6 = v62
        fmt.Println(v2, v1, v4)
        v4 += (v3 + float32(1e10))
        _ = fi78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v60:", len(v60), v60)
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
