package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int16; F2 bool }
type GenS5 struct { F0 uint64; F1 uint16; F2 bool }
func genFn0(p0 uint64, p1 int16) uint8 {
    p1 += (p1 ^ (^(p1)))
    fmt.Println(p1, p1)
    v6 := GenS1{F0: 0, F1: 4294967295, F2: -861}
    v7 := []int{42, 32767, int(p1), 42}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 90
}
func genFn1(p0 uint32, p1 int16, p2 uint) int16 {
    var v8 int64 = 9223372036854775807
    if (!(uint16(p2) > 32767)) {
        for fi9 := 0; fi9 < 1; fi9++ {
            if true {
                switch "" {
                case "a":
                    var v10 uint = p2
                    _ = v10
                case "hello":
                    var v11 uint64 = uint64(p0)
                    _ = v11
                case "foo bar":
                    for fi12 := 0; fi12 < 3; fi12++ {
                        if (!false) {
                            if true {
                                p1 -= -749
                            }
                            var v13 uint8 = 255
                            v14 := []int{2147483647, fi9, fi12}
                            _ = v13
                            _ = v14
                        }
                        var v15 uint16 = 42
                        _ = fi12
                        _ = v15
                    }
                }
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 uint64 = 42
                    v18 := GenS1{F0: 1, F1: p0, F2: (fi16 & (fi16 | fi9))}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                for fi19 := 0; fi19 < 8; fi19++ {
                    var v20 bool = false
                    _ = fi19
                    _ = v20
                }
            } else {
                if false {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        v8 = v8
                        _ = fi21
                    }
                } else {
                    fmt.Println(p2, v8)
                }
            }
            v22 := []int{fi9, (fi9 >> ((18446744073709551615) & 63))}
            _ = fi9
            _ = v22
        }
    } else {
        var v23 uint = 65535
        _ = v23
    }
    var v24 bool = false
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 float64 = 0.1
        switch ("\n\t" + "") {
        case "":
            for fi27 := 0; fi27 < 9; fi27++ {
                if v24 {
                    var v28 string = ""
                    _ = v28
                }
                v29 := []int{fi27, (-(fi27))}
                _ = fi27
                _ = v29
            }
        }
        switch ("hello" + "") {
        case "foo bar":
            if (!((!v24) && (!v24))) {
                var v30 uint = (p2 & 42)
                _ = v30
            }
        case "\n\t":
            if (-1 > 1) {
                v24 = v24
                if (true && ((!false) || v24)) {
                    if true {
                        var v31 int8 = 19
                        v32 := GenS1{F0: 255, F1: (p0 - (p0 << ((65535) & 63))), F2: 4294967295}
                        _ = v31
                        _ = v32
                    }
                    fmt.Println(p2, v24, p2)
                    var v33 int16 = (p1 << (((^(p2))) & 63))
                    _ = v33
                } else {
                    v26 += -2.25
                    var v34 uint64 = 1
                    _ = v34
                }
                fmt.Println(p1)
            } else {
                p2 = p2
            }
        default:
            var v35 bool = false
            _ = v35
        }
        _ = fi25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v8
    return -32768
}

func main() {
    v1 := []string{"foo bar", ("\n\t" + ((("hello" + ("\n\t" + ("" + "foo bar"))) + "foo bar") + "a")), "\n\t", (("\n\t" + "foo bar") + "\n\t")}
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v5 int = 1
    switch (-(v5)) {
    case -2147483648:
        var v36 int = (v5 | (v5 << ((1) & 63)))
        _ = v36
    }
    switch (v5 & v5) {
    case -521:
        if true {
            v37 := GenS1{F0: 75, F1: 213, F2: (v5 / (((v5 & 151)) | 1))}
            for fi38 := 0; fi38 < 4; fi38++ {
                switch 128 {
                case -752:
                    fmt.Println(fi38, v5)
                case -2147483648:
                    v39 := GenS1{F0: 128, F1: 2, F2: (fi38 - 802)}
                    _ = v39
                case 32767:
                    var v40 uint16 = 11
                    _ = v40
                }
                v41 := GenS2{F0: 127}
                for fi42 := 0; fi42 < 3; fi42++ {
                    fmt.Println(v5, v5)
                    v5 -= fi42
                    v5 -= (fi42 / ((127) | 1))
                    switch "" {
                    case "a":
                        if (2147483648 > 0) {
                            v5 -= -512
                            _ = len(v4)
                        } else {
                            v5 = fi38
                        }
                    case "foo bar":
                        var v43 uint16 = 1
                        _ = v43
                    default:
                        switch "hello" {
                        case "a":
                            v5 += 743
                        case "foo bar":
                            v44 := GenS2{F0: 127}
                            _ = v44
                        case "":
                            if (true || (false || false)) {
                                v45 := GenS1{F0: 0, F1: 0, F2: (fi38 << ((18446744073709551615) & 63))}
                                if false {
                                    v5 = (fi38 ^ (^(v5)))
                                    v46 := GenS1{F0: 0, F1: 2147483647, F2: (fi38 - v5)}
                                    fmt.Println(v5)
                                    v47 := GenS1{F0: 127, F1: 0, F2: (fi38 << ((2) & 63))}
                                    var v48 int = (v5 ^ fi38)
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                }
                                _ = v45
                            } else {
                                var v49 int16 = 0
                                if ((fi42 ^ (fi42 | 32768)) >= -677) {
                                    v50 := GenS3{F0: 42}
                                    if (true || (!(false || (true || (0.0 > -2.25))))) {
                                        if false {
                                            delete(v4, "a")
                                            if (1.5 <= (float32((1.5 - 0.0)) + 0.0)) {
                                                v5 = -9223372036854775808
                                                var v51 uint32 = 65535
                                                if (!(962 != 65535)) {
                                                    fmt.Println(fi38, v49)
                                                    v52 := v2[(^(v51))]
                                                    _ = v52
                                                }
                                                v51 = (v51 & 65535)
                                                _ = v51
                                            } else {
                                                v5 -= fi38
                                                if false {
                                                    v5 = (fi42 >> ((276) & 63))
                                                    switch "a" {
                                                    case "hello":
                                                        var v53 int = (fi38 ^ (fi42 + (fi42 & (fi42 + fi42))))
                                                        _ = v53
                                                    case "\n\t":
                                                        if true {
                                                            fmt.Println(v49, fi38)
                                                        } else {
                                                            var v54 bool = (!("\n\t" <= "hello"))
                                                            _ = len(v4)
                                                            _ = v54
                                                        }
                                                    case "":
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            var v55 bool = ("hello" > "foo bar")
                                                            _ = v55
                                                        case "":
                                                            var v56 uint16 = 32767
                                                            _ = v56
                                                        }
                                                        v57 := GenS1{F0: 1, F1: 0, F2: (fi38 + (fi38 - fi38))}
                                                        _ = v57
                                                    }
                                                    fmt.Println(fi42)
                                                }
                                            }
                                        }
                                    }
                                    var v58 uint32 = 1
                                    var v59 uint64 = 820
                                    _ = v50
                                    _ = v58
                                    _ = v59
                                } else {
                                    var v60 int64 = 400
                                    var v61 uint8 = 42
                                    _ = v60
                                    _ = v61
                                }
                                _ = v49
                            }
                        }
                    }
                    var v62 string = ""
                    v62 = (v62 + "hello")
                    _ = fi42
                    _ = v62
                }
                _ = fi38
                _ = v41
            }
            v5 = (v5 & v5)
            _ = v37
        }
    }
    v63 := GenS1{F0: genFn0(588, -861), F1: 0, F2: -729}
    for fi64 := 0; fi64 < 14; fi64++ {
        var v65 int8 = -27
        v66 := GenS1{F0: 127, F1: 939, F2: int(v65)}
        _ = fi64
        _ = v65
        _ = v66
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v5)
    }
    v67 := v4["foo bar"]
    var v68 float64 = (((-2.25 + (0.1 - 0.0)) + (0.1 + 1.5)) / ((float64(1.5) * float64((1.5 * 0.0))) + (float64(0.1) * -2.25)))
    v5 = v5
    fmt.Println(v68, v67)
    v5 = 56
    v3[-115] = false
    for fi69 := 0; fi69 < 15; fi69++ {
        fmt.Println(v67)
        _ = fi69
    }
    if (true && ((v67 < ("foo bar" + (v67 + ((v67 + v67) + v67)))) && false)) {
        v5 += v5
        for fi70 := 0; fi70 < 4; fi70++ {
            if (-9223372036854775808 == 285) {
                fmt.Println(v67, v67)
                for fi71 := 0; fi71 < 17; fi71++ {
                    if (genFn0(1, 1000) != 42) {
                        if ((!false) && true) {
                            switch -9223372036854775808 {
                            case 0:
                                var v72 int64 = 42
                                _ = v72
                            default:
                                fmt.Println(v68, fi71, v68)
                            }
                        } else {
                            v67 = v67
                            v73 := GenS4{F0: 198, F1: 0, F2: (419 >= (fi70 % ((fi71) | 1)))}
                            _ = v73
                        }
                        v68 = (v68 + v68)
                    }
                    _ = fi71
                }
            } else {
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v67)
                }
                var v74 bool = (true || true)
                switch v67 {
                case "":
                    delete(v4, v67)
                case "foo bar":
                    v75 := GenS3{F0: -1}
                    _ = v75
                default:
                    v76 := GenS5{F0: 2, F1: 1, F2: (65535 < 2147483647)}
                    _ = v76
                }
                _ = v74
            }
            v77 := GenS1{F0: 145, F1: 331, F2: v5}
            fmt.Println(fi70, v5)
            _ = fi70
            _ = v77
        }
        fmt.Println(v5, v67)
    } else {
        var v78 uint64 = 42
        _ = len(v2)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v78
    }
    v68 += 1.5
    var v79 bool = true
    for fi80 := 0; fi80 < 16; fi80++ {
        v79 = ((3.14159 * 1e10) == 3.14159)
        for fi81 := 0; fi81 < 9; fi81++ {
            fmt.Println(v67, v5, v67)
            var v82 uint8 = 42
            _ = fi81
            _ = v82
        }
        _ = fi80
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v67
    _ = v68
    _ = v79
}
