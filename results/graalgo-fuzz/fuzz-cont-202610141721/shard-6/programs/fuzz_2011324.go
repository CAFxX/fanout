package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 string; F2 uint32 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 int; F1 int; F2 int16 }
func genFn0(p0 uint64, p1 int, p2 int) string {
    p1 = (p1 / (((^(p1))) | 1))
    var v6 float32 = 1.5
    var v7 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "hello"
}
func genFn1() int64 {
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int64 = 1
        v10 := GenS1{F0: ("foo bar" + ("a" + ("hello" + "\n\t")))}
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 uint8 = 255
            if true {
                v9 = (v9 | v9)
                switch "" {
                case "\n\t":
                    var v13 string = "foo bar"
                    _ = v13
                case "a":
                    var v14 uint32 = 0
                    _ = v14
                default:
                    v15 := []int{1, 0, (fi8 | 32767)}
                    _ = v15
                }
            } else {
                switch fi8 {
                case -9223372036854775808:
                    v16 := GenS1{F0: "\n\t"}
                    _ = v16
                case 42:
                    v17 := []int{(fi8 * fi11)}
                    _ = v17
                case 128:
                    v12 += v12
                }
                var v18 float32 = ((((-2.25 * (0.1 + 3.14159)) + 3.14159) / (float32(3.14159) - (1.5 + 1.5))) - 1.5)
                _ = v18
            }
            _ = fi11
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v19 := []int{9223372036854775807, 552, 846, 1}
    _ = v19
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 int8 = -128
    v4 := []uint32{717, 2147483648, 42}
    var v5 int16 = -32768
    if len(v4) > 0 {
        v4[0] = 863
    }
    switch "foo bar" {
    case "a":
        _ = len(v4)
    case "foo bar":
        _ = len(v1)
    case "":
        v20 := v2[0]
        _ = v20
    default:
        var v21 uint32 = 2147483647
        _ = v21
    }
    v3 = v3
    v22 := GenS1{F0: "a"}
    for fi23 := 0; fi23 < 10; fi23++ {
        v24 := v1[fi23]
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        if (!(0 < 255)) {
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 int16 = 0
                if true {
                    var v28 int32 = -1
                    switch (fi26 + (fi25 - fi26)) {
                    case -2147483648:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        if false {
                            if false {
                                v29 := GenS1{F0: ""}
                                if ((float32(1.5) > (1.5 - 3.14159)) || false) {
                                    _ = len(v2)
                                    if (!(2 < 65535)) {
                                        var v30 bool = true
                                        _ = v30
                                    }
                                } else {
                                    fmt.Println(v5)
                                }
                                switch "a" {
                                case "\n\t":
                                    var v31 int = fi26
                                    _ = v31
                                case "hello":
                                    v32 := GenS1{F0: ("" + "\n\t")}
                                    _ = v32
                                case "foo bar":
                                    switch ((((("\n\t" + ("a" + (("a" + "a") + "a"))) + "") + "foo bar") + (("foo bar" + "\n\t") + "")) + "\n\t") {
                                    case "hello":
                                        var v33 int8 = v3
                                        _ = v33
                                    default:
                                        v27 -= (v5 ^ v5)
                                    }
                                }
                                _ = v29
                            }
                            v5 = -1
                        } else {
                            if (false || (true || true)) {
                                if false {
                                    var v34 uint = 18446744073709551615
                                    _ = v34
                                }
                            }
                            v5 = 634
                        }
                    }
                    _ = v28
                } else {
                    v35 := GenS2{F0: v3}
                    v36 := GenS2{F0: v3}
                    v5 = v27
                    var v37 float64 = -2.25
                    _ = v35
                    _ = v36
                    _ = v37
                }
                _ = fi26
                _ = v27
            }
        } else {
            v38 := GenS1{F0: "a"}
            if true {
                switch fi25 {
                case -2147483648:
                    delete(v2, (^(fi25)))
                case -634:
                    var v39 uint16 = 255
                    _ = v39
                }
                var v40 bool = true
                _ = v40
            } else {
                switch "a" {
                case "\n\t":
                    for fi41 := 0; fi41 < 17; fi41++ {
                        v42 := GenS3{F0: 628, F1: "\n\t", F2: 2147483648}
                        _ = fi41
                        _ = v42
                    }
                case "hello":
                    fmt.Println(v5)
                default:
                    v3 += (v3 | -54)
                }
                if false {
                    v4 = append(v4, 212)
                    fmt.Println(v5, fi25, v3)
                    var v43 float32 = 0.1
                    _ = v43
                }
                for fi44 := 0; fi44 < 8; fi44++ {
                    var v45 int64 = 1
                    var v46 uint16 = 1
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            }
            _ = v38
        }
        switch (fi25 | fi25) {
        case 32767:
            var v47 uint16 = 1
            _ = v47
        case -898:
            v5 -= v5
        case -1:
            var v48 uint64 = 1
            _ = v48
        default:
            v5 = 256
        }
        switch ("a" + (("hello" + ((((("" + "") + "hello") + ("foo bar" + ("a" + "a"))) + "\n\t") + "hello")) + ("foo bar" + "a"))) {
        case "a":
            v3 = (^(v3))
        case "hello":
            _ = len(v1)
        }
        _ = fi25
    }
    for fi49 := 0; fi49 < 16; fi49++ {
        switch ("" + ("a" + ("foo bar" + "hello"))) {
        case "hello":
            if len(v4) > 0 {
                v4[0] = 42
            }
        case "foo bar":
            var v50 int16 = v5
            _ = v50
        default:
            var v51 float64 = 1e10
            _ = v51
        }
        var v52 int64 = genFn1()
        var v53 uint32 = 191
        _ = fi49
        _ = v52
        _ = v53
    }
    if len(v4) > 0 {
        v4[0] = 1
    }
    if false {
        if false {
            v3 = (v3 | (-(v3)))
            if true {
                if ((v3 % ((v3) | 1)) > v3) {
                    var v54 uint32 = 42
                    if true {
                        var v55 float32 = (0.1 + 1e10)
                        if false {
                            var v56 uint8 = 2
                            _ = v56
                        }
                        _ = v55
                    } else {
                        v57 := GenS3{F0: 1, F1: "foo bar", F2: v54}
                        var v58 int = -50
                        _ = v57
                        _ = v58
                    }
                    v59 := GenS1{F0: "hello"}
                    _ = v54
                    _ = v59
                } else {
                    v5 = 1000
                    var v60 int16 = v5
                    var v61 uint64 = 0
                    _ = v60
                    _ = v61
                }
                v62 := GenS4{F0: -93, F1: v3}
                _ = v62
            }
        }
        v63 := v1[127]
        v64 := GenS5{F0: 32767, F1: -465, F2: (-(v5))}
        _ = v63
        _ = v64
    }
    if true {
        switch -1 {
        case 1:
            var v65 string = ("\n\t" + "a")
            _ = v65
        case 2147483647:
            var v66 uint = 255
            _ = v66
        case 127:
            for fi67 := 0; fi67 < 16; fi67++ {
                var v68 uint16 = 527
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi67
                _ = v68
            }
        }
    }
    var v69 uint32 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v69
}
