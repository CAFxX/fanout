package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint; F1 int; F2 int32 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 uint16 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{9223372036854775807}
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    var v5 int8 = -1
    var v6 uint16 = 374
    v7 := v3["foo bar"]
    v5 = (v5 + v5)
    if (88 < 0) {
        v8 := v3[("hello" + "hello")]
        v8 = v7
        var v9 uint8 = 35
        _ = v8
        _ = v9
    } else {
        switch -1 {
        case -9223372036854775808:
            v6 = (v6 & (^(v6)))
        default:
            v7 += (v7 | (v7 + v7))
        }
        var v10 int16 = -1
        _ = v10
    }
    var v11 int16 = -482
    fmt.Println(v6, v6, v6)
    var v12 uint8 = 79
    if (0 >= 42) {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
    }
    if true {
        for fi13 := 0; fi13 < 17; fi13++ {
            _ = len(v3)
            switch "hello" {
            case "\n\t":
                var v14 int = 253
                _ = v14
            case "":
                var v15 float64 = float64(0.1)
                _ = v15
            default:
                if true {
                    if (false && false) {
                        fmt.Println(v11, v11)
                        v5 = 39
                        v6 += 42
                    }
                    var v16 uint64 = 1
                    fmt.Println(v7, v12, v16)
                    _ = v16
                }
            }
            _ = fi13
        }
        for fi17 := 0; fi17 < 1; fi17++ {
            v1[""] = ((!(v12 < (v12 / ((2) | 1)))) && true)
            var v18 bool = true
            for fi19 := 0; fi19 < 7; fi19++ {
                v6 = (v6 * (v6 >> ((4294967295) & 63)))
                v20 := GenS1{F0: (v12 + (v12 / ((243) | 1)))}
                switch (fi19 >> ((2) & 63)) {
                case 2147483647:
                    v2 = append(v2, 9223372036854775807)
                case 4294967295:
                    v11 = -1
                default:
                    v21 := GenS2{F0: 42, F1: 278}
                    _ = v21
                }
                v22 := GenS2{F0: 852, F1: (v7 >> ((2) & 63))}
                v6 = 255
                _ = fi19
                _ = v20
                _ = v22
            }
            _ = fi17
            _ = v18
        }
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := GenS3{F0: 1, F1: fi23, F2: 42}
            var v25 string = "foo bar"
            v11 += v11
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    var v26 int32 = -641
    var v27 uint32 = (v7 >> ((18) & 63))
    v28 := v3["hello"]
    v11 = 1
    switch ("\n\t" + ("" + "a")) {
    case "foo bar":
        v29 := GenS2{F0: 1, F1: (v7 + 4294967295)}
        _ = v29
    case "\n\t":
        switch int(v7) {
        case 128:
            v27 = v28
        }
    case "a":
        if (42 == 474) {
            for fi30 := 0; fi30 < 1; fi30++ {
                if false {
                    if (true || false) {
                        fmt.Println(v6, v28, v28)
                        for fi31 := 0; fi31 < 5; fi31++ {
                            if true {
                                v32 := GenS2{F0: uint64(v28), F1: (v28 - (v27 + v28))}
                                v12 += (^(v12))
                                _ = v32
                            } else {
                                v33 := GenS2{F0: 4294967295, F1: v28}
                                if true {
                                    if true {
                                        v34 := GenS1{F0: (v12 * (v12 & (v12 + uint8(v6))))}
                                        var v35 uint64 = 18446744073709551615
                                        _ = v34
                                        _ = v35
                                    }
                                    v36 := v1["hello"]
                                    _ = v36
                                } else {
                                    var v37 float32 = ((1.5 / 1e10) * float32((float64((0.1 + (3.14159 + 1.5))) + ((0.0 / 0.1) / 0.1))))
                                    switch "hello" {
                                    case "a":
                                        _ = len(v4)
                                    case "foo bar":
                                        if true {
                                            if (424 == (^(v26))) {
                                                if (!((!((^(v5)) > 127)) && (false && (false && (true || false))))) {
                                                    fmt.Println(v11)
                                                    var v38 string = "a"
                                                    _ = v38
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    var v39 uint = 2
                                                    _ = v39
                                                default:
                                                    var v40 int16 = -32768
                                                    _ = v40
                                                }
                                            }
                                            if (!true) {
                                                var v41 uint16 = (v6 ^ v6)
                                                if ((0 == 0) && (((((v37 + v37) + (v37 / ((-2.25 - ((3.14159 - v37) - ((1e10 - -2.25) * (1.5 - 1e10)))) - v37))) * v37) * ((((-2.25 + float32(1e10)) * float32(0.1)) * (1.5 / v37)) - 1.5)) == v37)) {
                                                    v11 = 1
                                                    fmt.Println(v7, v26)
                                                    v42 := v4["foo bar"]
                                                    v11 += (v11 - -58)
                                                    var v43 bool = true
                                                    _ = v42
                                                    _ = v43
                                                } else {
                                                    var v44 uint32 = v7
                                                    v3["\n\t"] = 2147483647
                                                    _ = v44
                                                }
                                                _ = v41
                                            } else {
                                                if true {
                                                    v45 := GenS3{F0: 4294967295, F1: (fi30 / ((fi31) | 1)), F2: (v26 & v26)}
                                                    v46 := GenS4{F0: (^(v27)), F1: -9223372036854775808}
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    v47 := GenS5{F0: "\n\t"}
                                                    v48 := GenS4{F0: (v28 * (v27 % ((v27) | 1))), F1: -459}
                                                    var v49 uint = 0
                                                    if true {
                                                        var v50 int8 = (-(v5))
                                                        var v51 uint64 = 1
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                }
                                                if true {
                                                    v52 := GenS3{F0: uint(v7), F1: fi30, F2: (-(v26))}
                                                    _ = v52
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v11 = 1000
                                                    if (true || (false || true)) {
                                                        var v53 string = "foo bar"
                                                        v5 += (v5 >> ((0) & 63))
                                                        v54 := GenS5{F0: ("a" + ("hello" + v53))}
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        var v55 float32 = v37
                                                        v11 += 315
                                                        var v56 float64 = 0.0
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                }
                                            }
                                            v11 += 1
                                        }
                                    default:
                                        var v57 uint8 = v12
                                        _ = v57
                                    }
                                    v58 := GenS4{F0: v27, F1: 0}
                                    _ = v37
                                    _ = v58
                                }
                                if (!((false || (v12 < v12)) || (!false))) {
                                    var v59 string = "foo bar"
                                    _ = v59
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    switch fi31 {
                                    case 1:
                                        v60 := GenS6{F0: 704, F1: (true && true)}
                                        _ = v60
                                    case 4294967295:
                                        switch ("foo bar" + "\n\t") {
                                        case "":
                                            v26 += (v26 / ((0) | 1))
                                        }
                                        v11 = -1
                                        if true {
                                            if true {
                                                var v61 float32 = 0.0
                                                if true {
                                                    _ = len(v2)
                                                    v27 += v7
                                                    if false {
                                                        v12 -= 42
                                                        fmt.Println(v7)
                                                        var v62 int64 = 1
                                                        _ = v62
                                                    }
                                                }
                                                _ = v61
                                            } else {
                                                v5 = (v5 << ((42) & 63))
                                                switch ("foo bar" + ("foo bar" + "\n\t")) {
                                                case "foo bar":
                                                    var v63 int = (^(fi30))
                                                    _ = v63
                                                case "":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case "\n\t":
                                                    var v64 string = (("a" + "foo bar") + (((("foo bar" + "a") + "\n\t") + ("a" + ((("" + "\n\t") + "hello") + ""))) + "hello"))
                                                    _ = v64
                                                default:
                                                    v65 := GenS5{F0: ""}
                                                    _ = v65
                                                }
                                            }
                                            v6 = 1
                                            var v66 uint64 = 1
                                            _ = v66
                                        } else {
                                            if false {
                                                if (!false) {
                                                    v5 -= -47
                                                    switch fi31 {
                                                    case -777:
                                                        v5 = (v5 - int8(v11))
                                                    case 42:
                                                        switch "" {
                                                        case "\n\t":
                                                            fmt.Println(v12)
                                                        case "":
                                                            if false {
                                                                v12 = v12
                                                            }
                                                        case "hello":
                                                            var v67 int32 = -1
                                                            _ = v67
                                                        }
                                                        switch "" {
                                                        case "a":
                                                            _ = len(v3)
                                                        case "hello":
                                                            if (1 >= 302) {
                                                                var v68 uint64 = 42
                                                                fmt.Println(v68, v7, fi30)
                                                                _ = v68
                                                            }
                                                        default:
                                                            v28 = (v7 << ((255) & 63))
                                                        }
                                                    default:
                                                        v6 += v6
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v7, fi31, v12)
                                                _ = len(v3)
                                            }
                                        }
                                    case 901:
                                        var v69 uint16 = 338
                                        _ = v69
                                    }
                                    if false {
                                        v12 = (^(v12))
                                        if (!(uint16(v12) > 1)) {
                                            var v70 uint8 = (^(v12))
                                            _ = v70
                                        } else {
                                            fmt.Println(v28, v7, v12)
                                            v7 = (^(v7))
                                        }
                                    } else {
                                        switch "foo bar" {
                                        case "":
                                            var v71 uint32 = v28
                                            _ = v71
                                        default:
                                            var v72 float64 = 0.1
                                            _ = v72
                                        }
                                        fmt.Println(v12, v7, v11)
                                        fmt.Println(v26)
                                        v11 = -1
                                    }
                                }
                                _ = v33
                            }
                            var v73 float64 = ((1.5 * float64(1.5)) * -2.25)
                            var v74 string = "\n\t"
                            _ = fi31
                            _ = v73
                            _ = v74
                        }
                        switch fi30 {
                        case -324:
                            for fi75 := 0; fi75 < 7; fi75++ {
                                var v76 int32 = (v26 | -1)
                                v77 := v1["a"]
                                v11 = (v11 - v11)
                                _ = fi75
                                _ = v76
                                _ = v77
                            }
                        }
                    } else {
                        v5 += v5
                        for fi78 := 0; fi78 < 17; fi78++ {
                            if true {
                                if (!(v11 != 32767)) {
                                    v5 = (v5 + (v5 / ((v5) | 1)))
                                    v6 += v6
                                    v6 += (^(v6))
                                }
                                var v79 uint32 = 4294967295
                                _ = v79
                            }
                            v80 := GenS7{F0: (v6 & v6)}
                            if true {
                                fmt.Println(v28)
                                var v81 uint8 = v12
                                if (true || false) {
                                    v6 = (v6 & (v6 << ((255) & 63)))
                                    var v82 int32 = (v26 + v26)
                                    switch "hello" {
                                    case "\n\t":
                                        v2 = append(v2, 4294967295)
                                    case "hello":
                                        fmt.Println(v28, v82)
                                    default:
                                        var v83 float64 = float64(float32((1e10 + 0.1)))
                                        _ = v83
                                    }
                                    _ = v82
                                }
                                _ = v81
                            }
                            _ = fi78
                            _ = v80
                        }
                        var v84 int64 = 42
                        _ = v84
                    }
                } else {
                    delete(v3, (("a" + ((("foo bar" + "hello") + "hello") + (("hello" + "hello") + ("hello" + "foo bar")))) + "hello"))
                    _ = len(v2)
                    if (false && (("" + "a") != "foo bar")) {
                        v4["\n\t"] = 65535
                        v28 = v7
                        v1["\n\t"] = true
                    }
                }
                switch (fi30 | fi30) {
                case -9223372036854775808:
                    var v85 bool = true
                    _ = v85
                default:
                    var v86 int64 = 9223372036854775807
                    _ = v86
                }
                _ = fi30
            }
            for fi87 := 0; fi87 < 2; fi87++ {
                var v88 int = fi87
                _ = fi87
                _ = v88
            }
        }
    default:
        fmt.Println(v5)
    }
    v5 += v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
