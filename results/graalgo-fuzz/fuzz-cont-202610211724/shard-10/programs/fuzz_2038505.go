package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }
type GenS2 struct { F0 uint32; F1 bool; F2 int16 }
type GenS3 struct { F0 bool; F1 uint16 }

func main() {
    var v1 uint32 = 507
    var v2 uint = 0
    var v3 int8 = 127
    var v4 float64 = float64(0.1)
    var v5 uint = 65535
    if true {
        fmt.Println(v2, v5)
        var v6 bool = true
        fmt.Println(v6, v4, v4)
        _ = v6
    } else {
        v2 = 4294967295
    }
    v1 = v1
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{int(v2), (fi7 * (fi7 - fi7)), fi7, (fi7 - (^(fi7)))}
        if len(v8) > 0 {
            v8[0] = fi7
        }
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 42, F2: ""}
    if (v3 < (v3 << (((v5 * v2)) & 63))) {
        v3 += v3
        var v10 uint8 = 176
        fmt.Println(v4, v5)
        _ = v10
    }
    v11 := GenS1{F0: (v2 % ((v5) | 1)), F1: -32768, F2: "hello"}
    switch "hello" {
    case "":
        switch -494 {
        case -2147483648:
            var v12 int8 = v3
            _ = v12
        default:
            v5 += 4294967295
        }
    }
    var v13 uint32 = v1
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 float64 = v4
        var v16 uint32 = (v13 - v1)
        _ = fi14
        _ = v15
        _ = v16
    }
    switch "" {
    case "foo bar":
        var v17 bool = true
        _ = v17
    case "":
        var v18 int32 = -1
        _ = v18
    }
    switch "\n\t" {
    case "foo bar":
        var v19 int64 = 0
        _ = v19
    case "a":
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 uint16 = 65535
            _ = fi20
            _ = v21
        }
    case "":
        fmt.Println(v13, v3, v1)
    default:
        for fi22 := 0; fi22 < 17; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                if (true && true) {
                    v24 := []int{}
                    _ = v24
                } else {
                    var v25 uint = (v2 | v2)
                    if false {
                        v26 := []int{fi23}
                        v27 := GenS1{F0: 332, F1: -1, F2: "\n\t"}
                        if false {
                            fmt.Println(v5)
                            v13 += (v1 % ((v13) | 1))
                            if false {
                                if false {
                                    if len(v26) >= 2 {
                                        v26 = v26[1:len(v26)]
                                    }
                                    var v28 int8 = (v3 / ((-128) | 1))
                                    _ = v28
                                } else {
                                    v2 += (v5 / ((65535) | 1))
                                }
                                var v29 string = "hello"
                                var v30 uint32 = uint32(v5)
                                if (false || (false || (!(0.1 >= -2.25)))) {
                                    var v31 int64 = 1
                                    v32 := GenS1{F0: v25, F1: 1, F2: "foo bar"}
                                    v31 += (^(v31))
                                    _ = v31
                                    _ = v32
                                } else {
                                    var v33 float64 = v4
                                    if (!(true || (304 <= 2147483647))) {
                                        fmt.Println(fi22, v29, v33)
                                        var v34 uint8 = uint8(v2)
                                        if (533 >= -1) {
                                            v35 := GenS2{F0: (v13 ^ (^(v1))), F1: (92 > v34), F2: 125}
                                            if ((386 <= uint(v13)) || (!(189 < 1))) {
                                                switch (^(fi22)) {
                                                case 2147483647:
                                                    if true {
                                                        var v36 float64 = (-2.25 - v4)
                                                        switch v29 {
                                                        case "\n\t":
                                                            _ = len(v26)
                                                        }
                                                        var v37 int16 = -32768
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                case -392:
                                                    var v38 int16 = 0
                                                    _ = v38
                                                }
                                                fmt.Println(v34, v1)
                                            }
                                            if false {
                                                var v39 uint64 = 18446744073709551615
                                                if false {
                                                    v26 = append(v26, fi23)
                                                    v3 = v3
                                                    if true {
                                                        var v40 int8 = v3
                                                        v30 -= 805
                                                        _ = v40
                                                    } else {
                                                        v33 = v33
                                                    }
                                                    var v41 uint8 = uint8(v13)
                                                    _ = v41
                                                } else {
                                                    _ = len(v26)
                                                }
                                                _ = v39
                                            }
                                            _ = v35
                                        }
                                        _ = v34
                                    } else {
                                        var v42 int64 = 849
                                        if true {
                                            if len(v26) >= 2 {
                                                v26 = v26[1:len(v26)]
                                            }
                                            if true {
                                                v26 = append(v26, (^(fi22)))
                                                fmt.Println(v29, v3)
                                            } else {
                                                var v43 bool = (2 == 18446744073709551615)
                                                var v44 uint64 = 963
                                                fmt.Println(v42, v4)
                                                v3 += (^(v3))
                                                _ = v43
                                                _ = v44
                                            }
                                            var v45 uint16 = 2
                                            _ = v45
                                        } else {
                                            if false {
                                                v26 = append(v26, (fi23 ^ 127))
                                                if true {
                                                    var v46 uint16 = 255
                                                    _ = v46
                                                } else {
                                                    var v47 uint = (v2 & (v25 & 4294967295))
                                                    switch fi23 {
                                                    case -9223372036854775808:
                                                        var v48 int8 = 0
                                                        _ = v48
                                                    }
                                                    if len(v26) > 0 {
                                                        v26[0] = fi22
                                                    }
                                                    _ = v47
                                                }
                                                switch fi23 {
                                                case 0:
                                                    v1 = v1
                                                }
                                            }
                                        }
                                        v42 -= -1
                                        if ((v13 + (v30 + v30)) > 2147483647) {
                                            v25 += (v2 + (v2 << ((549) & 63)))
                                            var v49 uint32 = v1
                                            _ = v49
                                        }
                                        v50 := GenS1{F0: (v25 & (v5 >> ((v5) & 63))), F1: 32767, F2: v29}
                                        _ = v42
                                        _ = v50
                                    }
                                    var v51 uint8 = 127
                                    v52 := GenS2{F0: (v30 & v1), F1: (0.1 >= float32(float64(1.5))), F2: -1}
                                    fmt.Println(v13)
                                    _ = v33
                                    _ = v51
                                    _ = v52
                                }
                                var v53 int = fi23
                                _ = v29
                                _ = v30
                                _ = v53
                            }
                        } else {
                            fmt.Println(v2, v4)
                        }
                        _ = v26
                        _ = v27
                    }
                    var v54 int = fi22
                    _ = v25
                    _ = v54
                }
                switch ("" + "\n\t") {
                case "\n\t":
                    fmt.Println(v2)
                case "foo bar":
                    fmt.Println(v3)
                case "hello":
                    var v55 int64 = -1
                    _ = v55
                default:
                    var v56 uint8 = 255
                    _ = v56
                }
                switch "hello" {
                case "foo bar":
                    var v57 float64 = (((v4 - v4) * float64(((0.0 / 3.14159) + -2.25))) + v4)
                    _ = v57
                default:
                    if (42 >= (v3 ^ v3)) {
                        v3 = 127
                    } else {
                        v58 := GenS2{F0: (v13 & 65535), F1: ("foo bar" == (("hello" + (((("foo bar" + "") + "") + (("a" + "\n\t") + (("foo bar" + "hello") + ""))) + "a")) + "foo bar")), F2: 0}
                        v59 := []int{(fi22 * 32768), fi22, (fi23 / (((fi22 & (fi22 & -677))) | 1)), fi22}
                        _ = v58
                        _ = v59
                    }
                }
                _ = fi23
            }
            for fi60 := 0; fi60 < 4; fi60++ {
                v61 := GenS3{F0: true, F1: 390}
                var v62 uint16 = 1
                var v63 uint64 = 560
                _ = fi60
                _ = v61
                _ = v62
                _ = v63
            }
            v64 := []int{}
            _ = fi22
            _ = v64
        }
    }
    var v65 float32 = (1e10 / 1.0)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v9
}
