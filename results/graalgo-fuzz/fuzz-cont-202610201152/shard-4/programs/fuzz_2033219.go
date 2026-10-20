package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int32 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS5 struct { F0 uint; F1 uint64; F2 int }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 string; F1 int8; F2 uint16 }
func genFn0() int32 {
    v6 := []int{}
    switch ("\n\t" + (("\n\t" + ("\n\t" + ("\n\t" + ("\n\t" + ("hello" + "hello"))))) + "a")) {
    case "\n\t":
        var v7 string = "a"
        _ = v7
    case "hello":
        var v8 uint32 = 870
        _ = v8
    default:
        var v9 int16 = 765
        _ = v9
    }
    var v10 uint8 = 222
    fmt.Println(v10)
    var v11 int8 = 84
    _ = v10
    _ = v11
    _ = v6
    return -1
}
func genFn1(p0 int) uint64 {
    fmt.Println(p0)
    fmt.Println(p0)
    if true {
        var v12 int32 = 42
        v13 := []int{p0}
        var v14 uint = 42
        _ = v12
        _ = v13
        _ = v14
    } else {
        fmt.Println(p0)
        fmt.Println(p0)
    }
    _ = p0
    return 1
}

func main() {
    var v1 int = -9223372036854775808
    v2 := []int{(v1 + v1), (v1 & v1), v1}
    v3 := []string{"", ("" + "\n\t"), "a", ""}
    v4 := make(map[uint32]bool)
    v5 := make(map[string]bool)
    var v15 bool = (v1 < (^(v1)))
    switch v1 {
    case 32767:
        v16 := v4[851]
        _ = v16
    case 47:
        delete(v4, 0)
    case 32768:
        v15 = true
    }
    fmt.Println(v15, v15)
    v1 += v1
    for fi17 := 0; fi17 < 13; fi17++ {
        if (0 == 123) {
            var v18 uint16 = 65535
            _ = v18
        } else {
            for fi19 := 0; fi19 < 5; fi19++ {
                fmt.Println(v1, v1)
                _ = fi19
            }
            v4[65535] = (v15 && (133 == -278))
            v20 := GenS1{F0: v15, F1: genFn0(), F2: -1}
            _ = v20
        }
        for fi21 := 0; fi21 < 8; fi21++ {
            v1 -= fi17
            if v15 {
                switch (fi17 + (fi21 * 9223372036854775807)) {
                case 0:
                    fmt.Println(fi21)
                default:
                    var v22 int32 = 797
                    _ = v22
                }
                if true {
                    if (false || (2 >= 386)) {
                        var v23 float32 = 1.5
                        _ = v23
                    } else {
                        switch 104 {
                        case -598:
                            if ((1e10 / 1e10) < -2.25) {
                                v24 := GenS1{F0: true, F1: 1, F2: -59}
                                _ = v24
                            } else {
                                if v15 {
                                    v1 -= (fi17 * fi21)
                                    var v25 uint = 478
                                    if v15 {
                                        v26 := v5["\n\t"]
                                        _ = v26
                                    }
                                    fmt.Println(fi17, fi21, v15)
                                    _ = v25
                                } else {
                                    v27 := GenS2{F0: 127, F1: 2147483647}
                                    _ = v27
                                }
                                if ((2 == 42) && v15) {
                                    var v28 int16 = -32768
                                    var v29 uint8 = 255
                                    _ = v28
                                    _ = v29
                                }
                                if (int64(fi21) != -1) {
                                    v1 -= fi17
                                }
                            }
                        case -9223372036854775808:
                            switch (("hello" + ("" + "\n\t")) + "") {
                            case "foo bar":
                                v4[447] = (fi17 >= (fi17 >> ((1) & 63)))
                            }
                        }
                        fmt.Println(v1, v1)
                        var v30 float64 = 1.5
                        var v31 uint16 = 122
                        var v32 uint = 571
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    var v33 uint64 = 288
                    _ = v33
                } else {
                    switch (v1 / ((29) | 1)) {
                    case 42:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -64:
                        if len(v2) > 0 {
                            v2[0] = (fi17 * -682)
                        }
                    case 202:
                        v34 := GenS3{F0: 42, F1: (v1 - v1)}
                        if len(v2) > 0 {
                            v2[0] = 32767
                        }
                        var v35 int64 = 785
                        _ = v34
                        _ = v35
                    }
                    if (64 > 0) {
                        v15 = (947 > genFn0())
                        if false {
                            if (!true) {
                                v15 = v15
                                v1 -= (fi21 << ((799) & 63))
                            } else {
                                v36 := GenS4{F0: 1, F1: 65535, F2: 517}
                                v5["foo bar"] = v15
                                v1 = 32768
                                _ = v36
                            }
                        } else {
                            var v37 uint16 = 255
                            _ = v37
                        }
                        v38 := GenS2{F0: 255, F1: 42}
                        _ = v38
                    }
                }
                switch "foo bar" {
                case "foo bar":
                    v15 = v15
                case "":
                    fmt.Println(v1, fi17)
                default:
                    v1 = 128
                }
            }
            var v39 int32 = 863
            _ = fi21
            _ = v39
        }
        var v40 int16 = -32768
        _ = fi17
        _ = v40
    }
    v15 = true
    for fi41 := 0; fi41 < 4; fi41++ {
        var v42 uint8 = 2
        _ = fi41
        _ = v42
    }
    switch 780 {
    case 32767:
        var v43 uint64 = genFn1(v1)
        _ = v43
    case 0:
        fmt.Println(v1, v1)
    case 1:
        var v44 float64 = (float64(1e10) + 1.5)
        _ = v44
    default:
        if v15 {
            for fi45 := 0; fi45 < 13; fi45++ {
                for fi46 := 0; fi46 < 9; fi46++ {
                    delete(v5, "hello")
                    v15 = true
                    _ = fi46
                }
                v47 := GenS5{F0: 18446744073709551615, F1: 2, F2: v1}
                v48 := v4[2147483648]
                _ = fi45
                _ = v47
                _ = v48
            }
        }
    }
    if true {
        var v49 int8 = -51
        delete(v4, 4294967295)
        _ = v49
    } else {
        for fi50 := 0; fi50 < 11; fi50++ {
            var v51 float32 = 0.0
            v3 = append(v3, (("\n\t" + "a") + "a"))
            _ = fi50
            _ = v51
        }
        var v52 string = ""
        v15 = (v52 <= "")
        _ = v52
    }
    v53 := GenS6{F0: 2147483648, F1: "\n\t"}
    v1 += -478
    if (!(0 <= 348)) {
        delete(v5, ("foo bar" + ((((((("" + ("a" + "foo bar")) + "\n\t") + "\n\t") + "a") + "a") + (("hello" + "a") + "\n\t")) + ((("foo bar" + "a") + "\n\t") + "\n\t"))))
    } else {
        if v15 {
            v54 := GenS4{F0: 42, F1: 255, F2: 42}
            if (!v15) {
                delete(v5, ("a" + ("\n\t" + "foo bar")))
                if false {
                    var v55 uint = 255
                    _ = v55
                } else {
                    v1 = 646
                    var v56 uint8 = 128
                    switch v1 {
                    case 128:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = v56
                }
                var v57 float32 = 0.1
                _ = v57
            }
            _ = v54
        }
        var v58 int8 = 1
        var v59 int = (v1 + (v1 + 4294967295))
        _ = v58
        _ = v59
    }
    for fi60 := 0; fi60 < 15; fi60++ {
        for fi61 := 0; fi61 < 7; fi61++ {
            var v62 uint32 = 2147483648
            if true {
                if false {
                    _ = len(v4)
                    v15 = true
                } else {
                    v63 := GenS6{F0: v62, F1: "\n\t"}
                    v62 -= 65535
                    _ = v63
                }
            } else {
                v2 = append(v2, (^(fi60)))
            }
            _ = fi61
            _ = v62
        }
        if (1e10 > 0.1) {
            for fi64 := 0; fi64 < 6; fi64++ {
                var v65 uint8 = 74
                _ = fi64
                _ = v65
            }
            fmt.Println(fi60, v15, v1)
        } else {
            v66 := GenS2{F0: 127, F1: 316}
            _ = v66
        }
        _ = fi60
    }
    for fi67 := 0; fi67 < 13; fi67++ {
        fmt.Println(v1)
        _ = fi67
    }
    switch "foo bar" {
    case "":
        var v68 string = ""
        _ = v68
    case "\n\t":
        _ = len(v4)
    case "foo bar":
        v69 := GenS7{F0: ("\n\t" + (("hello" + ("a" + ("a" + "foo bar"))) + ((("" + "\n\t") + (("foo bar" + "foo bar") + "hello")) + "foo bar"))), F1: 127, F2: 901}
        _ = v69
    default:
        v70 := GenS4{F0: 127, F1: 42, F2: 18446744073709551615}
        _ = v70
    }
    if v15 {
        _ = len(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
}
