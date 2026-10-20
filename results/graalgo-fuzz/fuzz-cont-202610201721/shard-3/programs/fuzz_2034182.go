package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int; F2 string }
type GenS4 struct { F0 int64; F1 uint16; F2 uint }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint16) bool {
    p0 = (^(p0))
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: 187, F1: -715, F2: 846}
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 int = 32767
            fmt.Println(fi9, fi7, fi9)
            var v11 int = v10
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    var v12 uint = 65535
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := GenS1{F0: 74, F1: 798, F2: 42}
            var v16 int64 = 9223372036854775807
            var v17 int64 = -9223372036854775808
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        if (-1 != 42) {
            v12 = v12
            v18 := []int{fi13, (fi13 << ((v12) & 63))}
            var v19 int64 = -1
            _ = v18
            _ = v19
        }
        _ = fi13
    }
    if true {
        if true {
            var v20 int8 = 59
            v21 := GenS2{F0: v12, F1: 255}
            _ = v20
            _ = v21
        } else {
            for fi22 := 0; fi22 < 18; fi22++ {
                if true {
                    switch "hello" {
                    case "":
                        fmt.Println(v12, v12)
                    case "foo bar":
                        if true {
                            v23 := []int{(fi22 / ((fi22) | 1)), (fi22 * (-(fi22)))}
                            fmt.Println(v12)
                            v12 += (v12 * (v12 << ((v12) & 63)))
                            _ = v23
                        }
                    case "\n\t":
                        p0 -= (p0 - 383)
                    }
                }
                for fi24 := 0; fi24 < 12; fi24++ {
                    p0 += (p0 >> (((v12 >> ((v12) & 63))) & 63))
                    var v25 uint64 = 0
                    if (65535 == 336) {
                        v12 = v12
                        fmt.Println(p0)
                        if false {
                            fmt.Println(p0, v25, fi22)
                            var v26 uint = (v12 % ((v12) | 1))
                            _ = v26
                        }
                    }
                    var v27 float32 = 3.14159
                    var v28 int32 = 0
                    _ = fi24
                    _ = v25
                    _ = v27
                    _ = v28
                }
                p0 = p0
                _ = fi22
            }
            if false {
                switch ("a" + "foo bar") {
                case "hello":
                    v29 := []int{}
                    _ = v29
                case "foo bar":
                    if (!(2 <= 42)) {
                        var v30 uint16 = (p0 * (p0 * p0))
                        p0 = p0
                        v31 := []int{-807}
                        _ = v30
                        _ = v31
                    }
                default:
                    for fi32 := 0; fi32 < 17; fi32++ {
                        var v33 string = "hello"
                        v34 := GenS2{F0: v12, F1: 42}
                        var v35 int8 = -111
                        _ = fi32
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                }
                for fi36 := 0; fi36 < 20; fi36++ {
                    var v37 uint64 = 18446744073709551615
                    for fi38 := 0; fi38 < 11; fi38++ {
                        p0 += p0
                        v39 := []int{}
                        _ = fi38
                        _ = v39
                    }
                    fmt.Println(v37)
                    _ = fi36
                    _ = v37
                }
                var v40 float32 = -2.25
                _ = v40
            } else {
                switch ("foo bar" + "a") {
                case "hello":
                    for fi41 := 0; fi41 < 3; fi41++ {
                        v42 := GenS3{F0: 208, F1: fi41, F2: ("\n\t" + "hello")}
                        _ = fi41
                        _ = v42
                    }
                default:
                    var v43 int16 = -364
                    _ = v43
                }
            }
            var v44 int8 = 0
            _ = v44
        }
        var v45 string = "hello"
        _ = v45
    } else {
        for fi46 := 0; fi46 < 5; fi46++ {
            v47 := []int{}
            v48 := GenS3{F0: 129, F1: fi46, F2: "a"}
            _ = fi46
            _ = v47
            _ = v48
        }
        var v49 uint16 = 2
        _ = v49
    }
    _ = p0
    _ = v12
    return true
}
func genFn1() uint32 {
    for fi50 := 0; fi50 < 6; fi50++ {
        fmt.Println(fi50)
        switch fi50 {
        case 457:
            var v51 string = "hello"
            _ = v51
        case 1:
            var v52 string = ("foo bar" + "a")
            _ = v52
        case 394:
            v53 := GenS3{F0: 127, F1: (fi50 / ((32767) | 1)), F2: (("foo bar" + "hello") + "\n\t")}
            _ = v53
        default:
            var v54 uint32 = 671
            _ = v54
        }
        var v55 float64 = (-2.25 - 0.0)
        _ = fi50
        _ = v55
    }
    switch -2147483648 {
    case 42:
        v56 := GenS2{F0: 4294967295, F1: 2}
        _ = v56
    default:
        var v57 int32 = 1
        _ = v57
    }
    for fi58 := 0; fi58 < 2; fi58++ {
        if true {
            v59 := GenS2{F0: 456, F1: 216}
            for fi60 := 0; fi60 < 2; fi60++ {
                var v61 uint16 = 255
                if (("a" + "foo bar") == ("" + (("a" + "a") + "a"))) {
                    v62 := GenS4{F0: 9223372036854775807, F1: (v61 - v61), F2: 18446744073709551615}
                    v61 = (v61 << ((883) & 63))
                    v63 := GenS2{F0: 18446744073709551615, F1: 127}
                    _ = v62
                    _ = v63
                } else {
                    switch fi60 {
                    case 1:
                        v64 := GenS1{F0: 2, F1: 42, F2: -1}
                        _ = v64
                    }
                }
                var v65 bool = (42 >= -818)
                _ = fi60
                _ = v61
                _ = v65
            }
            _ = v59
        } else {
            var v66 uint64 = 4294967295
            if false {
                v66 = (v66 - (v66 >> ((65535) & 63)))
            }
            var v67 float64 = (float64((1e10 - -2.25)) * (0.0 / (1e10 / 3.14159)))
            _ = v66
            _ = v67
        }
        var v68 float32 = ((0.0 + float32((-2.25 / 1.0))) + float32((-2.25 - float64(-2.25))))
        _ = fi58
        _ = v68
    }
    v69 := []int{127, -783}
    _ = v69
    return 1
}

func main() {
    var v1 uint8 = 128
    var v2 int32 = 29
    var v3 int16 = 1
    var v4 uint64 = 1
    v5 := []string{"a"}
    v6 := make(map[string]string)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v70 := GenS1{F0: v4, F1: (v3 - (v3 | (^(v3)))), F2: 21}
    if false {
        v3 = (v3 * (v3 + (v3 & 110)))
        var v71 int = 32767
        var v72 float32 = 1.5
        _ = v71
        _ = v72
    }
    for fi73 := 0; fi73 < 11; fi73++ {
        if true {
            v74 := v6[("hello" + ("foo bar" + ""))]
            v75 := v6[("" + ((("foo bar" + (v74 + v74)) + v74) + (((v74 + v74) + v74) + v74)))]
            _ = v74
            _ = v75
        } else {
            var v76 float64 = ((3.14159 + 3.14159) - 3.14159)
            var v77 uint8 = (v1 - 168)
            _ = v76
            _ = v77
        }
        v78 := GenS3{F0: v1, F1: fi73, F2: "foo bar"}
        v79 := GenS5{F0: v2}
        _ = fi73
        _ = v78
        _ = v79
    }
    switch -2147483648 {
    case -453:
        if true {
            fmt.Println(v4, v1)
            var v80 uint = 1
            v1 = (v1 - (v1 - (^(v1))))
            _ = v80
        }
    case -1:
        v81 := GenS1{F0: v4, F1: v3, F2: -1}
        _ = v81
    case -2147483648:
        fmt.Println(v4, v4, v2)
    default:
        var v82 int32 = v2
        _ = v82
    }
    v4 -= v4
    var v83 uint = 18446744073709551615
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi84 := 0; fi84 < 15; fi84++ {
        for fi85 := 0; fi85 < 17; fi85++ {
            v4 += 458
            v3 += (v3 << ((v83) & 63))
            switch (^(fi85)) {
            case 356:
                v86 := v6[("\n\t" + "")]
                _ = v86
            case 32767:
                if true {
                    switch "" {
                    case "":
                        var v87 uint8 = uint8(v4)
                        _ = v87
                    case "foo bar":
                        _ = len(v6)
                    }
                    var v88 uint32 = 125
                    var v89 float64 = 0.0
                    v5 = append(v5, "hello")
                    v90 := GenS4{F0: 291, F1: 42, F2: 2}
                    if false {
                        v91 := GenS4{F0: 100, F1: 1, F2: (^(v83))}
                        _ = v91
                    }
                    _ = v88
                    _ = v89
                    _ = v90
                }
            }
            _ = fi85
        }
        v1 += v1
        _ = fi84
    }
    _ = len(v6)
    v3 += 840
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v83
}
