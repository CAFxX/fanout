package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 bool) uint {
    p0 = (p0 || p0)
    var v5 int8 = -128
    var v6 int16 = 77
    _ = p0
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 bool = false
    switch ("a" + ("foo bar" + "a")) {
    case "":
        v7 := v3[2]
        _ = v7
    case "hello":
        var v8 uint64 = 18446744073709551615
        _ = v8
    default:
        if v4 {
            if (!false) {
                v4 = v4
            } else {
                v4 = (128 != 42)
                v4 = false
            }
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 uint8 = 255
                _ = fi9
                _ = v10
            }
            fmt.Println(v4)
        }
    }
    if v4 {
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 uint = 839
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 uint32 = 2147483647
                if v4 {
                    switch fi13 {
                    case 127:
                        fmt.Println(v14, v14)
                    case 297:
                        if (1 <= 1) {
                            var v15 uint32 = v14
                            var v16 string = ("hello" + "\n\t")
                            _ = v15
                            _ = v16
                        }
                    default:
                        v17 := GenS1{F0: "hello", F1: -1, F2: 2147483648}
                        _ = v17
                    }
                    if v4 {
                        var v18 uint16 = 2
                        var v19 float64 = (3.14159 - float64(1.5))
                        fmt.Println(v14, v12)
                        _ = v18
                        _ = v19
                    }
                    fmt.Println(fi13, fi13)
                }
                _ = fi13
                _ = v14
            }
            v20 := GenS1{F0: "", F1: 42, F2: 1}
            _ = fi11
            _ = v12
            _ = v20
        }
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 string = ""
            if (!v4) {
                v22 = "foo bar"
            } else {
                v22 = "a"
            }
            v23 := GenS2{F0: 77}
            _ = fi21
            _ = v22
            _ = v23
        }
    } else {
        v4 = v4
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 int64 = -1
            switch "" {
            case "hello":
                for fi26 := 0; fi26 < 14; fi26++ {
                    if (2 >= 2) {
                        v27 := v1["\n\t"]
                        if true {
                            v28 := GenS1{F0: "a", F1: 127, F2: 42}
                            if (9223372036854775807 != -9223372036854775808) {
                                v29 := GenS2{F0: 127}
                                fmt.Println(v25)
                                switch "\n\t" {
                                case "hello":
                                    var v30 int = 327
                                    _ = v30
                                case "\n\t":
                                    v31 := []int{(fi26 << ((4294967295) & 63)), (fi24 - (fi24 - 4294967295)), 1, fi26}
                                    _ = v31
                                case "foo bar":
                                    v25 = -430
                                default:
                                    v25 = 0
                                }
                                _ = v29
                            } else {
                                v32 := GenS1{F0: ("" + "\n\t"), F1: 49, F2: 2147483648}
                                v33 := GenS2{F0: -61}
                                _ = v32
                                _ = v33
                            }
                            if (float32(float64(3.14159)) == 0.0) {
                                if (1 != 1) {
                                    var v34 float64 = 1.5
                                    var v35 uint64 = 0
                                    var v36 int = (fi26 / (((-(fi26))) | 1))
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                } else {
                                    if (!true) {
                                        v37 := GenS1{F0: "", F1: 73, F2: 1}
                                        if v27 {
                                            switch fi24 {
                                            case 128:
                                                if (("" + ("hello" + "foo bar")) >= ("\n\t" + "a")) {
                                                    v38 := v3[566]
                                                    _ = v38
                                                } else {
                                                    v27 = (fi26 >= (fi24 << ((4294967295) & 63)))
                                                }
                                            case 9223372036854775807:
                                                switch "foo bar" {
                                                case "":
                                                    switch "hello" {
                                                    case "hello":
                                                        v27 = v27
                                                    case "\n\t":
                                                        var v39 float64 = ((0.0 + 1.5) / 3.14159)
                                                        _ = v39
                                                    default:
                                                        var v40 int64 = v25
                                                        _ = v40
                                                    }
                                                default:
                                                    v41 := []int{751, fi26, fi26, fi24}
                                                    var v42 string = "foo bar"
                                                    v43 := GenS2{F0: -1}
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                            case 127:
                                                v27 = (true && (-320 == (v25 << ((2) & 63))))
                                            default:
                                                var v44 bool = v4
                                                var v45 int16 = 1
                                                _ = v44
                                                _ = v45
                                            }
                                            if true {
                                                var v46 bool = (42 == -2147483648)
                                                _ = v46
                                            } else {
                                                v47 := v3[2147483648]
                                                var v48 uint = 912
                                                v49 := GenS2{F0: 1}
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                            }
                                        }
                                        v50 := GenS1{F0: "", F1: 0, F2: 618}
                                        _ = v37
                                        _ = v50
                                    }
                                    v51 := v2[""]
                                    switch "\n\t" {
                                    case "a":
                                        switch 32768 {
                                        case 262:
                                            v51 -= (-(v51))
                                        case 128:
                                            v52 := []int{}
                                            if false {
                                                v25 = -950
                                                switch ("" + ((("a" + "") + "\n\t") + "hello")) {
                                                case "foo bar":
                                                    if ((!v4) && true) {
                                                        if v27 {
                                                            fmt.Println(v27, fi24, fi24)
                                                        } else {
                                                            v1["\n\t"] = v27
                                                            v25 -= v25
                                                        }
                                                    } else {
                                                        _ = len(v52)
                                                    }
                                                case "a":
                                                    var v53 int = (^(v51))
                                                    if ((("hello" + "") + "\n\t") != "hello") {
                                                        v25 = 550
                                                    } else {
                                                        var v54 int32 = -2147483648
                                                        var v55 uint32 = 2147483648
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                    v4 = (4 != 2)
                                                    _ = v53
                                                default:
                                                    if ((1 > 241) || (-702 != -2147483648)) {
                                                        v56 := GenS1{F0: "\n\t", F1: -1, F2: 4294967295}
                                                        v25 += 42
                                                        _ = v56
                                                    }
                                                    if (2147483647 < -2147483648) {
                                                        v57 := GenS3{F0: 32767}
                                                        var v58 float64 = 0.1
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        if (v25 < (v25 * (v25 + 9223372036854775807))) {
                                                            var v59 uint = 0
                                                            var v60 float64 = float64(((float32(1e10) * (1e10 / -2.25)) / 3.14159))
                                                            v61 := GenS4{F0: (v25 * (v25 & (v25 ^ v25)))}
                                                            var v62 uint64 = 2
                                                            var v63 int64 = v25
                                                            _ = v59
                                                            _ = v60
                                                            _ = v61
                                                            _ = v62
                                                            _ = v63
                                                        } else {
                                                            v27 = (249 <= 2)
                                                            var v64 bool = (!(42 != 2147483647))
                                                            _ = v64
                                                        }
                                                        var v65 string = "hello"
                                                        var v66 string = "foo bar"
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    fmt.Println(v51, v4)
                                                    if false {
                                                        fmt.Println(v4)
                                                        var v67 uint32 = 0
                                                        _ = v67
                                                    } else {
                                                        delete(v3, 42)
                                                        v51 = fi26
                                                    }
                                                }
                                            } else {
                                                v68 := GenS5{F0: 255, F1: (v25 - (-(v25)))}
                                                _ = v68
                                            }
                                            _ = v52
                                        }
                                    case "":
                                        fmt.Println(v51, fi24)
                                    case "hello":
                                        var v69 uint8 = 108
                                        _ = v69
                                    }
                                    _ = v51
                                }
                                v70 := GenS1{F0: "foo bar", F1: 127, F2: 150}
                                _ = v70
                            }
                            _ = v28
                        }
                        var v71 uint64 = 826
                        var v72 uint32 = 42
                        delete(v1, "a")
                        v73 := GenS1{F0: ("\n\t" + (((("hello" + (("foo bar" + "\n\t") + "foo bar")) + "a") + ("foo bar" + "a")) + "hello")), F1: -1, F2: (v72 + 2147483648)}
                        _ = v27
                        _ = v71
                        _ = v72
                        _ = v73
                    } else {
                        v74 := []int{}
                        switch (fi24 & (fi26 + fi24)) {
                        case 32767:
                            v75 := GenS6{F0: 65535}
                            _ = v75
                        case 1:
                            var v76 uint = 744
                            var v77 int64 = (v25 / ((v25) | 1))
                            if true {
                                v76 -= v76
                                var v78 int64 = (v25 >> (((v76 - (v76 ^ v76))) & 63))
                                v4 = (!((v4 || (!(!((v25 << ((358) & 63)) <= (v25 + -9223372036854775808))))) && false))
                                fmt.Println(fi26, v76, v78)
                                if len(v74) >= 2 {
                                    v74 = v74[1:len(v74)]
                                }
                                v3[2] = "a"
                                var v79 int16 = 615
                                _ = v78
                                _ = v79
                            }
                            _ = v76
                            _ = v77
                        default:
                            switch fi26 {
                            case 0:
                                if (42 > 1) {
                                    v25 = v25
                                    var v80 int64 = (v25 % ((-1) | 1))
                                    _ = len(v74)
                                    _ = v80
                                }
                            case 170:
                                v25 = (v25 >> ((609) & 63))
                            case 909:
                                if (!(2147483647 > 1)) {
                                    var v81 int32 = -2147483648
                                    _ = v81
                                } else {
                                    if v4 {
                                        fmt.Println(fi24, fi26)
                                        _ = len(v3)
                                        fmt.Println(fi26, v4, fi26)
                                    } else {
                                        if (v25 != v25) {
                                            switch ("hello" + ("foo bar" + "")) {
                                            case "hello":
                                                fmt.Println(fi24, v4, fi24)
                                            }
                                        }
                                        var v82 bool = (0 >= -1)
                                        _ = v82
                                    }
                                    if ((((!v4) && true) || v4) && v4) {
                                        v83 := GenS3{F0: 32767}
                                        fmt.Println(v4)
                                        _ = v83
                                    } else {
                                        var v84 int = 4294967295
                                        var v85 uint32 = 2147483647
                                        v86 := GenS4{F0: 42}
                                        _ = v84
                                        _ = v85
                                        _ = v86
                                    }
                                }
                            default:
                                delete(v1, "foo bar")
                            }
                        }
                        fmt.Println(fi24)
                        _ = v74
                    }
                    v87 := []int{(fi24 - fi26), 229, (^(fi26))}
                    v4 = (1 <= fi24)
                    _ = fi26
                    _ = v87
                }
            }
            _ = fi24
            _ = v25
        }
        var v88 bool = (18446744073709551615 < 673)
        _ = v88
    }
    var v89 int = 0
    for fi90 := 0; fi90 < 7; fi90++ {
        v91 := []int{649, (v89 / ((v89) | 1))}
        fmt.Println(v89, fi90)
        _ = fi90
        _ = v91
    }
    fmt.Println(v89, v4)
    v4 = ((v89 << ((48) & 63)) == v89)
    var v92 uint = 832
    v93 := GenS1{F0: "", F1: 0, F2: 648}
    if false {
        v89 = v89
    } else {
        switch "" {
        case "foo bar":
            var v94 uint = v92
            _ = v94
        case "":
            v89 = v89
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v89:", v89)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v89
    _ = v92
    _ = v93
}
