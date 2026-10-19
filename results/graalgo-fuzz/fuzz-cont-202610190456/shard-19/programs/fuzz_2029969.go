package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 int32; F2 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 int32; F1 int16 }
type GenS9 struct { F0 int32; F1 int; F2 int8 }
func genFn0(p0 uint) uint16 {
    p0 -= (^(p0))
    if true {
        v7 := []int{4294967295, -268, 697}
        v8 := GenS1{F0: ("a" == ((("a" + "hello") + "a") + "hello")), F1: -1}
        _ = v7
        _ = v8
    } else {
        v9 := GenS2{F0: 2}
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        switch "hello" {
        case "hello":
            var v11 float64 = -2.25
            _ = v11
        case "\n\t":
            if true {
                fmt.Println(p0, p0)
            }
        case "foo bar":
            var v12 int32 = 0
            _ = v12
        }
        for fi13 := 0; fi13 < 19; fi13++ {
            p0 += p0
            switch 128 {
            case -300:
                var v14 bool = false
                _ = v14
            case 1:
                fmt.Println(fi10, fi13, fi13)
            case 127:
                if true {
                    p0 -= (p0 >> ((p0) & 63))
                    var v15 int = fi13
                    var v16 int = 32768
                    _ = v15
                    _ = v16
                }
            default:
                v17 := []int{fi13, 9223372036854775807}
                _ = v17
            }
            v18 := []int{-2147483648, (fi13 - fi13)}
            _ = fi13
            _ = v18
        }
        for fi19 := 0; fi19 < 11; fi19++ {
            p0 += (p0 | p0)
            v20 := []int{(fi19 % ((42) | 1))}
            _ = fi19
            _ = v20
        }
        _ = fi10
    }
    switch -493 {
    case -2147483648:
        var v21 int32 = 697
        _ = v21
    case 128:
        var v22 int8 = -77
        _ = v22
    case -9223372036854775808:
        p0 += (p0 << ((p0) & 63))
    default:
        v23 := GenS2{F0: 4294967295}
        _ = v23
    }
    _ = p0
    return 1
}
func genFn1(p0 int16) uint16 {
    var v24 uint64 = 4294967295
    v24 += v24
    _ = p0
    _ = v24
    return genFn0(4294967295)
}

func main() {
    v1 := []string{"a", "\n\t", ("hello" + (("a" + "\n\t") + "foo bar")), ("a" + "")}
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    var v4 string = ""
    v5 := []uint32{379, 2}
    var v6 uint = 758
    v25 := GenS3{F0: v6}
    var v26 bool = (((255 != 0) && (-757 != 1)) && (true || (true && (2 >= 1))))
    v4 = v4
    switch -1 {
    case 1:
        v5 = append(v5, 664)
    case 191:
        for fi27 := 0; fi27 < 18; fi27++ {
            v4 = v4
            v26 = (!(!true))
            _ = fi27
        }
    default:
        if (1 >= 42) {
            var v28 bool = ((-110 <= 42) && false)
            _ = v28
        } else {
            var v29 uint32 = 2147483647
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v30 := GenS3{F0: v6}
            _ = v29
            _ = v30
        }
    }
    switch "" {
    case "foo bar":
        for fi31 := 0; fi31 < 5; fi31++ {
            v32 := GenS3{F0: v6}
            if false {
                switch v4 {
                case "":
                    v33 := GenS1{F0: (!v26), F1: 0}
                    _ = v33
                case "\n\t":
                    fmt.Println(v4, v26, v26)
                default:
                    v26 = (812 > 1)
                }
                v4 = v4
            }
            fmt.Println(v6, fi31, fi31)
            _ = fi31
            _ = v32
        }
    case "\n\t":
        if (((("" + "\n\t") > ("hello" + "\n\t")) && (v26 || false)) || true) {
            v34 := GenS2{F0: 2147483647}
            _ = v34
        } else {
            var v35 uint64 = 592
            v36 := GenS4{F0: 612, F1: 42, F2: (v4 + (v4 + v4))}
            _ = v35
            _ = v36
        }
    default:
        v4 = ((v4 + v4) + ("\n\t" + v4))
    }
    var v37 bool = v26
    v38 := GenS2{F0: 924}
    if (!(42 < 2)) {
        var v39 int = 807
        if (uint32(v6) == 2147483648) {
            v4 = (v4 + "a")
            for fi40 := 0; fi40 < 1; fi40++ {
                if (1 <= 1) {
                    if (v37 && true) {
                        for fi41 := 0; fi41 < 11; fi41++ {
                            v42 := GenS3{F0: v6}
                            v37 = v37
                            _ = fi41
                            _ = v42
                        }
                        switch v4 {
                        case "":
                            if len(v1) > 0 {
                                v1[0] = v4
                            }
                        case "a":
                            if len(v1) > 0 {
                                v1[0] = v4
                            }
                        case "hello":
                            v6 = 45
                        default:
                            var v43 int64 = -1
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 7; fi44++ {
                            fmt.Println(v4, v26)
                            switch v39 {
                            case 32768:
                                v45 := GenS5{F0: 2147483647}
                                _ = v45
                            default:
                                switch ((v4 + v4) + ((v4 + "") + "foo bar")) {
                                case "hello":
                                    v4 = v4
                                }
                            }
                            _ = fi44
                        }
                    }
                    v46 := GenS5{F0: 0}
                    if len(v5) > 0 {
                        v5[0] = 365
                    }
                    _ = v46
                } else {
                    var v47 uint = (v6 % ((v6) | 1))
                    var v48 string = "\n\t"
                    for fi49 := 0; fi49 < 6; fi49++ {
                        var v50 int = fi40
                        _ = fi49
                        _ = v50
                    }
                    _ = v47
                    _ = v48
                }
                for fi51 := 0; fi51 < 2; fi51++ {
                    if (9223372036854775807 < -195) {
                        var v52 uint64 = 869
                        delete(v2, ("" + ("hello" + (v4 + ("hello" + v4)))))
                        _ = v52
                    } else {
                        v53 := GenS6{F0: 87, F1: v6, F2: 1}
                        var v54 int16 = -32768
                        switch v4 {
                        case "hello":
                            if false {
                                v55 := GenS7{F0: "hello", F1: 183}
                                _ = v55
                            }
                        case "":
                            var v56 int64 = -147
                            _ = v56
                        }
                        _ = v53
                        _ = v54
                    }
                    if ((2 != 2) || v26) {
                        if false {
                            v39 = 42
                            fmt.Println(v39)
                        }
                        if v26 {
                            if false {
                                switch ("" + ((v4 + (v4 + "foo bar")) + v4)) {
                                case "hello":
                                    fmt.Println(fi40, fi51, v4)
                                case "":
                                    if (!v26) {
                                        var v57 uint64 = 42
                                        _ = v57
                                    }
                                }
                                fmt.Println(fi40, fi51)
                                v58 := GenS8{F0: 1000, F1: 0}
                                _ = v58
                            } else {
                                v1 = append(v1, ("" + ("foo bar" + "\n\t")))
                                v39 = (v39 & v39)
                                switch v39 {
                                case 2147483647:
                                    var v59 uint32 = 0
                                    _ = v59
                                case -271:
                                    switch fi40 {
                                    case 127:
                                        v4 = v4
                                        if v37 {
                                            v60 := GenS1{F0: ((1 < fi51) || (false || v37)), F1: 361}
                                            v37 = true
                                            _ = v60
                                        } else {
                                            delete(v2, "")
                                            var v61 string = (("\n\t" + v4) + (("\n\t" + (("" + (v4 + v4)) + (v4 + ""))) + ((v4 + v4) + v4)))
                                            var v62 int8 = -94
                                            switch (v39 & 365) {
                                            case 4294967295:
                                                fmt.Println(v39)
                                            case 1:
                                                var v63 uint8 = 42
                                                _ = v63
                                            }
                                            _ = v61
                                            _ = v62
                                        }
                                        v4 = v4
                                        var v64 int8 = 42
                                        var v65 uint16 = 0
                                        _ = v64
                                        _ = v65
                                    case 32767:
                                        v37 = v26
                                    case -249:
                                        var v66 uint = (v6 / ((v6) | 1))
                                        _ = v66
                                    }
                                    fmt.Println(fi51, fi51, fi51)
                                    var v67 uint8 = 127
                                    _ = v67
                                }
                                _ = len(v1)
                                if len(v5) > 0 {
                                    v5[0] = 65535
                                }
                            }
                            if v37 {
                                v68 := GenS8{F0: -2147483648, F1: -561}
                                switch (v39 * (^(v39))) {
                                case 0:
                                    v69 := GenS9{F0: 121, F1: fi40, F2: int8(fi51)}
                                    _ = v69
                                case 9223372036854775807:
                                    v70 := GenS9{F0: -946, F1: (v39 << ((173) & 63)), F2: 100}
                                    _ = v70
                                case -2147483648:
                                    if v26 {
                                        var v71 uint32 = 0
                                        _ = v71
                                    }
                                default:
                                    var v72 float32 = (0.0 - 3.14159)
                                    _ = v72
                                }
                                var v73 int16 = 0
                                _ = v68
                                _ = v73
                            }
                            var v74 int64 = 0
                            _ = v74
                        } else {
                            fmt.Println(v37)
                            var v75 uint64 = 4294967295
                            _ = v75
                        }
                        if len(v1) > 0 {
                            v1[0] = v4
                        }
                    }
                    var v76 uint64 = 18446744073709551615
                    _ = fi51
                    _ = v76
                }
                _ = fi40
            }
            var v77 uint16 = 32767
            _ = v77
        }
        _ = v39
    } else {
        v26 = (42 < 0)
        for fi78 := 0; fi78 < 8; fi78++ {
            v6 -= v6
            var v79 bool = (!v26)
            _ = fi78
            _ = v79
        }
    }
    if true {
        if v26 {
            var v80 uint32 = 132
            v81 := GenS5{F0: 2}
            _ = v80
            _ = v81
        } else {
            var v82 float64 = 1.5
            _ = v82
        }
        var v83 int = 32768
        _ = v83
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    v4 = "hello"
    v5 = append(v5, 2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
