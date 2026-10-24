package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16 }
func genFn0() int {
    v5 := GenS1{F0: 18446744073709551615, F1: ("a" + "hello")}
    var v6 int8 = 1
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 42, F1: "a"}
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return 127
}
func genFn1() int {
    v9 := GenS2{F0: 0}
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int8 = 100
        v12 := []int{(fi10 % ((128) | 1)), fi10}
        if (4294967295 <= 2147483648) {
            var v13 int32 = -863
            if (((0.1 / ((0.1 - 1e10) * 3.14159)) != 1e10) || (-680 == -113)) {
                var v14 int16 = 1000
                for fi15 := 0; fi15 < 3; fi15++ {
                    v16 := GenS2{F0: 495}
                    var v17 uint8 = 255
                    if false {
                        v18 := GenS1{F0: 284, F1: "foo bar"}
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v13, v11)
                            var v19 int32 = -724
                            _ = v19
                        case "a":
                            if (!true) {
                                fmt.Println(v17, fi15, fi10)
                                fmt.Println(v17)
                            } else {
                                switch ("hello" + "hello") {
                                case "foo bar":
                                    switch fi15 {
                                    case 32768:
                                        if true {
                                            fmt.Println(fi15, v14, v14)
                                            v20 := GenS1{F0: 725, F1: "\n\t"}
                                            v12 = append(v12, 9223372036854775807)
                                            var v21 uint64 = 4294967295
                                            switch ("\n\t" + "") {
                                            case "hello":
                                                var v22 uint = 65535
                                                _ = v22
                                            case "\n\t":
                                                var v23 int = int(v17)
                                                _ = v23
                                            case "foo bar":
                                                if false {
                                                    if ((v14 / (((v14 | (v14 / ((42) | 1)))) | 1)) > (v14 + (v14 + (v14 >> ((641) & 63))))) {
                                                        v24 := GenS2{F0: 2}
                                                        _ = v24
                                                    } else {
                                                        switch fi15 {
                                                        case 128:
                                                            var v25 bool = (0.0 == 3.14159)
                                                            _ = v25
                                                        default:
                                                            v21 = (v21 - (v21 & (v21 % ((v21) | 1))))
                                                        }
                                                        switch ("" + "hello") {
                                                        case "hello":
                                                            if true {
                                                                v26 := GenS3{F0: (fi10 & genFn0()), F1: uint32(v21)}
                                                                v12 = append(v12, (fi10 - fi15))
                                                                if true {
                                                                    fmt.Println(v11, v14, fi10)
                                                                    v27 := GenS3{F0: (fi15 | (fi10 ^ genFn0())), F1: 65535}
                                                                    _ = v27
                                                                } else {
                                                                    fmt.Println(fi15, v13)
                                                                    if true {
                                                                        fmt.Println(v11, v21, v17)
                                                                    } else {
                                                                        v11 += 127
                                                                        var v28 int32 = v13
                                                                        _ = v28
                                                                    }
                                                                    fmt.Println(fi15, v17, v14)
                                                                    if false {
                                                                        v17 = v17
                                                                        fmt.Println(v11)
                                                                        if true {
                                                                            var v29 uint32 = 4294967295
                                                                            _ = v29
                                                                        } else {
                                                                            v21 -= 42
                                                                            var v30 int = -152
                                                                            _ = v30
                                                                        }
                                                                    }
                                                                }
                                                                v31 := GenS3{F0: (fi10 | (fi10 >> ((239) & 63))), F1: 65535}
                                                                _ = v26
                                                                _ = v31
                                                            } else {
                                                                v32 := GenS3{F0: (fi15 + (fi15 + fi15)), F1: 6}
                                                                fmt.Println(v11, fi10, v17)
                                                                var v33 float64 = ((-2.25 * (1e10 - 1e10)) + 3.14159)
                                                                v34 := GenS4{F0: v17}
                                                                _ = v32
                                                                _ = v33
                                                                _ = v34
                                                            }
                                                        }
                                                        switch int(v13) {
                                                        case 1:
                                                            v13 -= (v13 % ((v13) | 1))
                                                        case 815:
                                                            v35 := GenS3{F0: int(v14), F1: 753}
                                                            _ = v35
                                                        case 249:
                                                            v14 = (^(v14))
                                                        default:
                                                            fmt.Println(v11)
                                                        }
                                                        v36 := GenS1{F0: v21, F1: (("hello" + "a") + ("" + "\n\t"))}
                                                        _ = v36
                                                    }
                                                    v14 -= v14
                                                    switch "hello" {
                                                    case "foo bar":
                                                        fmt.Println(v17, v11, v21)
                                                        if len(v12) > 0 {
                                                            v12[0] = fi15
                                                        }
                                                    case "a":
                                                        v13 += v13
                                                    case "\n\t":
                                                        if len(v12) > 0 {
                                                            v12[0] = fi10
                                                        }
                                                    }
                                                    var v37 uint16 = 42
                                                    var v38 string = "foo bar"
                                                    _ = v37
                                                    _ = v38
                                                }
                                            default:
                                                if (-2.25 > (-2.25 / 1.0)) {
                                                    if (true || false) {
                                                        v11 += (v11 * 126)
                                                    } else {
                                                        v39 := GenS3{F0: fi15, F1: 205}
                                                        var v40 float64 = float64(0.1)
                                                        _ = v39
                                                        _ = v40
                                                    }
                                                }
                                            }
                                            _ = v20
                                            _ = v21
                                        }
                                    default:
                                        v41 := GenS5{F0: (v14 % ((-32768) | 1))}
                                        _ = v41
                                    }
                                default:
                                    var v42 bool = (false && (885 != 299))
                                    _ = v42
                                }
                                var v43 uint32 = 4294967295
                                v43 += (v43 >> ((65535) & 63))
                                v44 := GenS2{F0: 65535}
                                v45 := GenS6{F0: 32767}
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                        default:
                            v11 = v11
                        }
                        fmt.Println(v14, v13, v13)
                        _ = v18
                    }
                    v17 += 255
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v11 = (v11 % (((v11 >> ((2) & 63))) | 1))
                _ = v14
            }
            var v46 int8 = (v11 & (v11 & 0))
            _ = v13
            _ = v46
        } else {
            v11 -= v11
            var v47 int8 = v11
            var v48 int16 = 761
            _ = v47
            _ = v48
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        var v49 int8 = 1
        _ = v49
    } else {
        v50 := []int{9223372036854775807}
        _ = v50
    }
    var v51 string = (((("a" + "foo bar") + ("\n\t" + "foo bar")) + "a") + (((("foo bar" + "a") + "\n\t") + ("a" + "\n\t")) + "hello"))
    if (18446744073709551615 < 0) {
        if (!(4294967295 >= 18446744073709551615)) {
            for fi52 := 0; fi52 < 3; fi52++ {
                var v53 float64 = 0.1
                _ = fi52
                _ = v53
            }
        } else {
            var v54 float64 = (1e10 * 3.14159)
            var v55 uint16 = 2
            _ = v54
            _ = v55
        }
        v56 := []int{42}
        fmt.Println(v51)
        _ = v56
    }
    _ = v51
    _ = v9
    return -787
}

func main() {
    var v1 int16 = 1
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 int = 0
    switch ("\n\t" + "\n\t") {
    case "":
        v57 := GenS3{F0: (-(v4)), F1: 0}
        _ = v57
    case "foo bar":
        fmt.Println(v1)
    }
    v4 = (v4 % (((-(v4))) | 1))
    v4 += (^(v4))
    if false {
        fmt.Println(v4)
        var v58 uint32 = 837
        _ = v58
    }
    v4 = v4
    var v59 int16 = (v1 + (v1 * 1))
    for fi60 := 0; fi60 < 19; fi60++ {
        var v61 float64 = 0.0
        v1 += v1
        var v62 int16 = 1
        _ = fi60
        _ = v61
        _ = v62
    }
    _ = len(v3)
    for fi63 := 0; fi63 < 7; fi63++ {
        v4 = fi63
        var v64 string = "\n\t"
        v1 = (v1 << ((688) & 63))
        _ = fi63
        _ = v64
    }
    _ = len(v3)
    var v65 bool = ((v1 + (v1 ^ (^(v59)))) < (v1 << ((272) & 63)))
    for fi66 := 0; fi66 < 4; fi66++ {
        v67 := []int{v4}
        _ = fi66
        _ = v67
    }
    var v68 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("v65:", v65)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v65
    _ = v68
}
