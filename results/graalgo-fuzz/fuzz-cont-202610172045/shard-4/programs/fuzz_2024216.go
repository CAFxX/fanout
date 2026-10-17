package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS4 struct { F0 string; F1 int16; F2 uint8 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 uint8; F1 uint8; F2 uint }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int8; F1 int16 }
func genFn0() int8 {
    var v4 bool = (42 > 42)
    v5 := []int{}
    v4 = (-32768 < 32767)
    for fi6 := 0; fi6 < 2; fi6++ {
        _ = len(v5)
        _ = fi6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint32 = 2147483647
        v9 := GenS1{F0: 42, F1: -1, F2: 13}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v4
    _ = v5
    return 127
}
func genFn1() uint16 {
    var v10 uint32 = 2147483648
    v10 = (v10 ^ v10)
    _ = v10
    return 42
}

func main() {
    v1 := []string{"", ""}
    var v2 float32 = (3.14159 * 0.1)
    var v3 int16 = 42
    switch ("foo bar" + ("\n\t" + "hello")) {
    case "foo bar":
        v3 += (v3 - 1000)
    case "a":
        v2 = (v2 * 3.14159)
    default:
        var v11 int = -2147483648
        _ = v11
    }
    v1 = append(v1, "")
    if true {
        for fi12 := 0; fi12 < 15; fi12++ {
            for fi13 := 0; fi13 < 2; fi13++ {
                v1 = append(v1, ("hello" + "a"))
                switch "\n\t" {
                case "":
                    if false {
                        if true {
                            v2 = 0.1
                        } else {
                            v3 += v3
                            v14 := GenS2{F0: 903, F1: 42, F2: 255}
                            v3 += (v3 << ((400) & 63))
                            _ = v14
                        }
                        switch "foo bar" {
                        case "\n\t":
                            switch (fi12 % ((fi12) | 1)) {
                            case -2147483648:
                                var v15 int32 = 776
                                _ = v15
                            case -350:
                                v3 = 0
                            case -4:
                                switch ("a" + "\n\t") {
                                case "\n\t":
                                    var v16 uint = 18446744073709551615
                                    _ = v16
                                case "a":
                                    var v17 uint64 = 0
                                    _ = v17
                                default:
                                    switch (("hello" + "") + "") {
                                    case "a":
                                        if len(v1) > 0 {
                                            v1[0] = "foo bar"
                                        }
                                    case "\n\t":
                                        switch 0 {
                                        case 127:
                                            var v18 string = ("a" + (("hello" + "a") + "\n\t"))
                                            _ = v18
                                        }
                                        v19 := GenS1{F0: 1000, F1: -397, F2: -32768}
                                        if ((fi13 << ((2) & 63)) == 42) {
                                            var v20 int8 = 42
                                            switch 1 {
                                            case -949:
                                                v20 = 127
                                            }
                                            switch "foo bar" {
                                            case "foo bar":
                                                v21 := GenS1{F0: v3, F1: 614, F2: (v3 ^ v3)}
                                                _ = v21
                                            case "hello":
                                                var v22 bool = true
                                                _ = v22
                                            }
                                            _ = v20
                                        }
                                        _ = v19
                                    case "":
                                        if true {
                                            _ = len(v1)
                                            var v23 uint16 = 255
                                            if false {
                                                var v24 float32 = ((v2 * v2) / (0.1 + v2))
                                                _ = v24
                                            } else {
                                                v3 = (v3 + (v3 / (((v3 - v3)) | 1)))
                                                var v25 int = (fi13 % (((-(fi12))) | 1))
                                                var v26 bool = (!true)
                                                fmt.Println(v3, v25, fi13)
                                                var v27 uint8 = 128
                                                v27 = (^(v27))
                                                _ = v25
                                                _ = v26
                                                _ = v27
                                            }
                                            _ = v23
                                        } else {
                                            var v28 uint8 = 158
                                            _ = v28
                                        }
                                    }
                                }
                            }
                        default:
                            fmt.Println(fi12, fi12, fi12)
                        }
                    }
                case "hello":
                    switch "a" {
                    case "hello":
                        var v29 string = "\n\t"
                        _ = v29
                    case "":
                        fmt.Println(v2)
                    }
                case "foo bar":
                    var v30 uint = 2
                    _ = v30
                }
                if false {
                    _ = len(v1)
                    var v31 uint = 18446744073709551615
                    _ = v31
                } else {
                    if false {
                        v32 := GenS2{F0: 18446744073709551615, F1: 53, F2: 65535}
                        v33 := GenS3{F0: 544, F1: 1, F2: 2}
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        _ = v32
                        _ = v33
                    } else {
                        switch (^(fi13)) {
                        case 1:
                            _ = len(v1)
                        case -94:
                            var v34 int16 = (v3 / (((v3 - v3)) | 1))
                            _ = v34
                        default:
                            switch fi12 {
                            case -9223372036854775808:
                                if true {
                                    if true {
                                        var v35 uint64 = 2
                                        var v36 float32 = float32(((1e10 / (-2.25 - (1.5 + (-2.25 / (3.14159 / 1.5))))) / 1.5))
                                        v37 := GenS4{F0: "\n\t", F1: 0, F2: 128}
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                    v3 -= v3
                                    _ = len(v1)
                                } else {
                                    v2 = (((-2.25 - (float32(1e10) * ((float32(float64(-2.25)) - 1e10) * float32(0.1)))) * v2) - (3.14159 - v2))
                                    v38 := GenS5{F0: 32767, F1: false}
                                    v39 := GenS6{F0: 0, F1: 1, F2: 934}
                                    switch (fi13 << ((2) & 63)) {
                                    case 32767:
                                        v40 := GenS7{F0: (-(fi13))}
                                        _ = v40
                                    case -2147483648:
                                        var v41 uint64 = 972
                                        _ = v41
                                    case -880:
                                        if false {
                                            v3 = 1000
                                            var v42 int8 = -45
                                            var v43 bool = ((genFn0() == (v42 << ((42) & 63))) && ((fi13 & (fi12 | fi13)) <= fi13))
                                            if len(v1) > 0 {
                                                v1[0] = ""
                                            }
                                            v44 := GenS8{F0: 0}
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        }
                                    }
                                    _ = v38
                                    _ = v39
                                }
                                var v45 int32 = 1
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v45
                            case 0:
                                v46 := GenS9{F0: 42}
                                _ = v46
                            case 2147483647:
                                var v47 int8 = 100
                                _ = v47
                            default:
                                fmt.Println(fi12)
                            }
                        }
                        var v48 int16 = (v3 ^ v3)
                        fmt.Println(fi13, v48)
                        _ = v48
                    }
                    v2 = float32(-2.25)
                    var v49 string = "foo bar"
                    var v50 int8 = -11
                    if true {
                        if false {
                            var v51 float64 = (1.5 - (1.5 / 1e10))
                            fmt.Println(v3, v2)
                            v1 = append(v1, v49)
                            _ = v51
                        }
                        if false {
                            _ = len(v1)
                            var v52 float64 = ((float64(((v2 * (((v2 / (((v2 / 0.1) + (((v2 - ((v2 + v2) - v2)) * (1e10 - v2)) - v2)) - v2)) + v2) * float32((0.0 - (1e10 * -2.25))))) + (-2.25 / v2))) * float64((v2 + 1e10))) / 1e10)
                            fmt.Println(v2, v50)
                            if (true && (int32(v3) != -2147483648)) {
                                var v53 float64 = float64(0.1)
                                var v54 int8 = genFn0()
                                _ = v53
                                _ = v54
                            }
                            if (false || true) {
                                switch v49 {
                                case "hello":
                                    if (!false) {
                                        switch v49 {
                                        case "hello":
                                            var v55 bool = (!(((v2 + (0.1 + v2)) / (3.14159 / 3.14159)) < v2))
                                            _ = v55
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = v49
                                            }
                                        }
                                    }
                                    if false {
                                        var v56 int32 = -1
                                        if len(v1) > 0 {
                                            v1[0] = ("a" + "")
                                        }
                                        _ = v56
                                    }
                                default:
                                    fmt.Println(fi13, fi12, v52)
                                }
                                if false {
                                    v50 = v50
                                }
                            }
                            _ = v52
                        } else {
                            v57 := GenS10{F0: (v50 >> ((65535) & 63)), F1: (-(v3))}
                            fmt.Println(fi12, v49)
                            v50 += (v50 * v50)
                            fmt.Println(fi12, fi13)
                            var v58 int64 = 0
                            _ = v57
                            _ = v58
                        }
                    } else {
                        v2 = (0.0 - v2)
                        _ = len(v1)
                    }
                    v50 -= (v50 >> ((0) & 63))
                    _ = v49
                    _ = v50
                }
                v3 = 0
                _ = fi13
            }
            var v59 uint16 = 65535
            _ = fi12
            _ = v59
        }
    } else {
        var v60 float32 = 0.1
        fmt.Println(v60, v3, v60)
        if true {
            for fi61 := 0; fi61 < 18; fi61++ {
                fmt.Println(v3, fi61)
                var v62 int = fi61
                _ = fi61
                _ = v62
            }
            _ = len(v1)
        }
        _ = v60
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2)
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    v63 := GenS1{F0: v3, F1: -1, F2: (^(v3))}
    for fi64 := 0; fi64 < 5; fi64++ {
        var v65 float32 = (1.5 - (1e10 - 0.1))
        switch fi64 {
        case 943:
            v3 = v3
        }
        v66 := GenS9{F0: 169}
        _ = fi64
        _ = v65
        _ = v66
    }
    if ((0.0 <= 0.1) && ((((v2 * v2) * v2) <= 1e10) && (((0 != 42) && ((!true) && (!true))) || (!(!true))))) {
        if ((!(!false)) || (v2 > ((0.1 / v2) + float32(1.5)))) {
            v2 = 1.5
        }
        var v67 string = "a"
        _ = v67
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v63
}
