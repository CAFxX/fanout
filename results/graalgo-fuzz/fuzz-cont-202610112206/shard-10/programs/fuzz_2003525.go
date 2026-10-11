package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int32; F2 int }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint; F1 string; F2 uint32 }
type GenS6 struct { F0 uint32; F1 int; F2 int16 }
func genFn0() string {
    var v4 uint16 = 386
    v5 := GenS1{F0: -9223372036854775808}
    v4 -= v4
    switch ("foo bar" + (("" + "foo bar") + "a")) {
    case "\n\t":
        var v6 int16 = 1
        _ = v6
    case "foo bar":
        var v7 int32 = -1
        _ = v7
    case "":
        v8 := []int{-2147483648, 641}
        _ = v8
    default:
        v4 = v4
    }
    var v9 int8 = 103
    _ = v4
    _ = v5
    _ = v9
    return "a"
}
func genFn1() int32 {
    if true {
        var v10 float64 = 1.5
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 uint = 4294967295
        var v13 int8 = -29
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int64 = -548
    _ = v14
    return -1
}

func main() {
    var v1 float32 = -2.25
    var v2 uint32 = 803
    var v3 int64 = 9223372036854775807
    fmt.Println(v1, v2, v1)
    var v15 uint = 255
    switch ("" + ("\n\t" + "\n\t")) {
    case "a":
        v3 -= 559
    case "":
        var v16 int16 = 1000
        _ = v16
    default:
        v2 -= 2147483647
    }
    switch "" {
    case "":
        switch "a" {
        case "\n\t":
            for fi17 := 0; fi17 < 9; fi17++ {
                for fi18 := 0; fi18 < 19; fi18++ {
                    v19 := []int{}
                    var v20 float32 = 0.0
                    v2 += (v2 + v2)
                    if (true && true) {
                        var v21 float32 = 1e10
                        _ = v21
                    } else {
                        v15 = (v15 & 2)
                        v22 := GenS1{F0: v3}
                        v2 -= uint32(v15)
                        _ = v22
                    }
                    v23 := GenS1{F0: (v3 >> ((4294967295) & 63))}
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v23
                }
                v3 = (v3 / (((v3 + (v3 | 639))) | 1))
                _ = fi17
            }
        default:
            v24 := []int{}
            _ = v24
        }
    case "hello":
        switch ((("" + "") + "\n\t") + "") {
        case "\n\t":
            v25 := GenS1{F0: 961}
            _ = v25
        case "a":
            var v26 uint8 = 2
            _ = v26
        }
    default:
        v27 := GenS2{F0: "foo bar"}
        _ = v27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 uint = (v15 / ((0) | 1))
        var v30 string = "foo bar"
        if false {
            var v31 uint16 = 65535
            _ = v31
        } else {
            for fi32 := 0; fi32 < 11; fi32++ {
                v33 := []int{(fi32 >> ((941) & 63)), (-(fi32)), (^(fi28)), -2147483648}
                var v34 int32 = 1
                _ = fi32
                _ = v33
                _ = v34
            }
            for fi35 := 0; fi35 < 4; fi35++ {
                v36 := []int{(fi28 & (^(fi35)))}
                _ = len(v36)
                _ = len(v36)
                _ = fi35
                _ = v36
            }
            var v37 int = fi28
            _ = v37
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    if (-2147483648 <= 127) {
        if (!false) {
            switch "a" {
            case "a":
                var v38 bool = false
                _ = v38
            case "foo bar":
                var v39 string = ("a" + "")
                _ = v39
            default:
                var v40 float32 = 3.14159
                _ = v40
            }
            if false {
                var v41 string = "foo bar"
                v42 := []int{128, 462}
                var v43 int8 = -128
                _ = v41
                _ = v42
                _ = v43
            } else {
                v44 := []int{379, -140, 4294967295, -321}
                v45 := GenS3{F0: (v3 ^ v3), F1: -1, F2: -358}
                _ = v44
                _ = v45
            }
        } else {
            v2 = v2
            v46 := []int{4294967295, 186, -1, 42}
            _ = v46
        }
        for fi47 := 0; fi47 < 6; fi47++ {
            fmt.Println(fi47, v3)
            v15 = v15
            v15 = v15
            _ = fi47
        }
    } else {
        v48 := GenS4{F0: "", F1: (v2 * v2)}
        v49 := GenS5{F0: v15, F1: "hello", F2: v2}
        for fi50 := 0; fi50 < 1; fi50++ {
            for fi51 := 0; fi51 < 7; fi51++ {
                switch fi50 {
                case 360:
                    fmt.Println(v15)
                }
                v15 += (^(v15))
                var v52 float32 = v1
                _ = fi51
                _ = v52
            }
            _ = fi50
        }
        _ = v48
        _ = v49
    }
    fmt.Println(v2, v15)
    var v53 uint16 = 32767
    v54 := GenS1{F0: (v3 % (((v3 >> ((4294967295) & 63))) | 1))}
    if ((v53 * v53) > v53) {
        var v55 string = ""
        v56 := []int{}
        _ = v55
        _ = v56
    }
    if (814 < 390) {
        if (false && true) {
            if true {
                if (!(((1.5 / 1.5) + 3.14159) != (1e10 + float64(1e10)))) {
                    v57 := []int{128, 32767, 9223372036854775807, 32768}
                    var v58 string = ((("a" + "") + ("" + "a")) + ("hello" + "\n\t"))
                    _ = v57
                    _ = v58
                } else {
                    if (3.14159 != ((-2.25 / 1.0) / (((1.5 - 1e10) + (3.14159 / (3.14159 - ((float64(-2.25) / 3.14159) / 0.1)))) + (0.0 / 1.5)))) {
                        v2 = v2
                        v15 = 18446744073709551615
                    } else {
                        switch 4294967295 {
                        case -1:
                            if (("a" + "hello") <= "") {
                                v3 -= -42
                                var v59 int8 = 127
                                _ = v59
                            }
                        case 152:
                            v60 := GenS3{F0: int64(v2), F1: 1, F2: -1}
                            _ = v60
                        case 32768:
                            switch "a" {
                            case "foo bar":
                                var v61 uint64 = 1
                                _ = v61
                            case "a":
                                var v62 uint = (v15 & (v15 * v15))
                                _ = v62
                            default:
                                var v63 int32 = genFn1()
                                _ = v63
                            }
                        }
                        var v64 uint64 = 922
                        switch ("hello" + ("a" + "a")) {
                        case "\n\t":
                            v65 := GenS3{F0: v3, F1: -1, F2: 32767}
                            _ = v65
                        }
                        _ = v64
                    }
                    v66 := []int{}
                    _ = v66
                }
                v1 -= 3.14159
                if false {
                    var v67 uint32 = 544
                    var v68 uint8 = 83
                    v69 := []int{-9223372036854775808, 162}
                    _ = v67
                    _ = v68
                    _ = v69
                } else {
                    for fi70 := 0; fi70 < 5; fi70++ {
                        v3 += v3
                        fmt.Println(v53, v2)
                        _ = fi70
                    }
                }
            } else {
                v2 -= (v2 >> (((v15 ^ v15)) & 63))
                var v71 uint16 = (v53 - 65535)
                _ = v71
            }
            if true {
                var v72 uint8 = 34
                _ = v72
            }
        }
        switch "\n\t" {
        case "":
            var v73 int64 = v3
            _ = v73
        case "\n\t":
            var v74 bool = false
            _ = v74
        default:
            fmt.Println(v53, v3)
        }
        var v75 int8 = 42
        _ = v75
    } else {
        v76 := GenS6{F0: (^(v2)), F1: 287, F2: 42}
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v53
    _ = v54
}
