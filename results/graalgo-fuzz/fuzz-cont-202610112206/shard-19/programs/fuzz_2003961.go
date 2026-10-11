package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint16; F2 string }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS6 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint; F1 int; F2 int8 }
type GenS11 struct { F0 int8; F1 int; F2 int8 }
type GenS12 struct { F0 int16; F1 int64; F2 uint16 }
func genFn0() uint {
    var v4 uint32 = 125
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 uint32 = 65535
        v7 := GenS1{F0: 42}
        if false {
            v4 = (v6 ^ v4)
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: 1}
    var v9 float64 = 0.1
    fmt.Println(v9)
    _ = v4
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 uint, p1 uint64, p2 uint8) string {
    if true {
        var v10 uint16 = 255
        var v11 bool = (!("hello" <= ("hello" + (("foo bar" + "hello") + ""))))
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            for fi13 := 0; fi13 < 3; fi13++ {
                p1 = p1
                p1 -= 38
                fmt.Println(fi13, p2, fi12)
                _ = fi13
            }
            _ = fi12
        }
        var v14 int32 = -1
        _ = v14
    }
    var v15 uint16 = 728
    var v16 float64 = 3.14159
    var v17 int64 = -137
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := []int{(fi18 * 128), (fi18 >> ((genFn0()) & 63)), fi18}
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v17
    return ("\n\t" + "a")
}

func main() {
    var v1 string = ""
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    var v20 int64 = 1
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 uint = 889
        var v23 bool = (!(v1 == ""))
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS3{F0: 2147483648, F1: 65535}
    var v25 string = (v1 + "\n\t")
    delete(v2, v25)
    _ = len(v3)
    var v26 uint64 = 18446744073709551615
    v25 = (v25 + "hello")
    v20 = (v20 - (v20 * (-(v20))))
    for fi27 := 0; fi27 < 12; fi27++ {
        v20 = (^(v20))
        for fi28 := 0; fi28 < 5; fi28++ {
            v1 = v25
            v29 := GenS2{F0: int8(v20)}
            if false {
                fmt.Println(v25)
            } else {
                if (0 > 1) {
                    v30 := []int{(fi28 | fi27), (fi27 + fi27), fi28}
                    if false {
                        v26 += 42
                        v1 = (((("" + ((v25 + "") + "foo bar")) + ("\n\t" + (v25 + v1))) + v1) + v25)
                        var v31 uint16 = 324
                        _ = v31
                    }
                    if false {
                        if (-486 == 42) {
                            switch v25 {
                            case "foo bar":
                                switch fi28 {
                                case -9223372036854775808:
                                    if false {
                                        if false {
                                            _ = len(v3)
                                            if (-32768 > -1) {
                                                if (127 <= -74) {
                                                    if len(v30) >= 2 {
                                                        v30 = v30[1:len(v30)]
                                                    }
                                                    var v32 bool = (false || (!true))
                                                    var v33 uint = 2
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    var v34 uint64 = v26
                                                    v35 := GenS4{F0: genFn0(), F1: 40, F2: ((("hello" + v1) + v1) + v1)}
                                                    _ = v34
                                                    _ = v35
                                                }
                                                fmt.Println(v26, v1)
                                            }
                                        } else {
                                            if true {
                                                v1 = "hello"
                                                v36 := GenS1{F0: (^(v26))}
                                                var v37 uint64 = v26
                                                _ = v36
                                                _ = v37
                                            } else {
                                                fmt.Println(fi27)
                                            }
                                        }
                                    } else {
                                        v38 := GenS5{F0: 930, F1: 4294967295, F2: 1}
                                        v39 := GenS3{F0: 894, F1: 2}
                                        var v40 string = ((v25 + v1) + (v1 + (((("" + "hello") + v25) + "") + ((v1 + v1) + v1))))
                                        v20 = (v20 % (((^(v20))) | 1))
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                case 128:
                                    v41 := GenS6{F0: v26, F1: (v20 << ((975) & 63)), F2: 873}
                                    _ = v41
                                case -1:
                                    var v42 int8 = 0
                                    _ = v42
                                default:
                                    v43 := GenS7{F0: 618}
                                    _ = v43
                                }
                            default:
                                var v44 float32 = (float32(3.14159) + (1e10 / float32(0.1)))
                                _ = v44
                            }
                            _ = len(v2)
                            var v45 uint64 = (v26 >> ((21) & 63))
                            _ = v45
                        } else {
                            switch (-(fi27)) {
                            case 128:
                                v25 = v25
                            case 2147483647:
                                var v46 int = 435
                                _ = v46
                            case 42:
                                if false {
                                    v47 := GenS8{F0: ((21 > 143) || (!true))}
                                    switch v25 {
                                    case "a":
                                        v2["hello"] = v25
                                    case "hello":
                                        var v48 int = 697
                                        _ = v48
                                    case "foo bar":
                                        var v49 float64 = 1e10
                                        _ = v49
                                    }
                                    _ = v47
                                } else {
                                    if (false && true) {
                                        fmt.Println(fi27)
                                    }
                                }
                                var v50 uint16 = 42
                                v51 := GenS9{F0: false}
                                _ = v50
                                _ = v51
                            default:
                                v52 := GenS8{F0: true}
                                _ = v52
                            }
                            if len(v30) >= 2 {
                                v30 = v30[1:len(v30)]
                            }
                        }
                        var v53 float32 = 3.14159
                        v54 := GenS3{F0: 12, F1: 2147483647}
                        _ = v53
                        _ = v54
                    }
                    v20 += (v20 | (v20 | (v20 * v20)))
                    if (int64(fi28) > (v20 - (v20 + -1))) {
                        fmt.Println(v25)
                        v55 := GenS3{F0: 0, F1: 573}
                        var v56 uint8 = 1
                        _ = v55
                        _ = v56
                    }
                    if (((true && true) && (!true)) || true) {
                        v30 = append(v30, (fi28 + fi27))
                        if (((float64(0.1) / 1e10) + -2.25) < 0.0) {
                            v25 = "a"
                            if (false || (1 == -567)) {
                                v57 := GenS10{F0: 40, F1: fi27, F2: -128}
                                var v58 float32 = -2.25
                                _ = v57
                                _ = v58
                            } else {
                                var v59 string = (v25 + v1)
                                _ = v59
                            }
                        }
                    }
                    _ = v30
                } else {
                    if ((!false) || false) {
                        v25 = "hello"
                        if true {
                            var v60 uint16 = 32767
                            _ = v60
                        }
                        fmt.Println(v1, v1, v1)
                    } else {
                        if (4294967295 < 348) {
                            v61 := []int{(fi27 >> ((genFn0()) & 63)), fi28, 42, fi27}
                            if (true || (false || true)) {
                                v26 -= (v26 ^ v26)
                                if (1 < 32767) {
                                    switch v25 {
                                    case "\n\t":
                                        if true {
                                            if (true || (!true)) {
                                                v61 = append(v61, (fi27 | 42))
                                            }
                                            switch 9223372036854775807 {
                                            case 2147483647:
                                                var v62 int8 = 88
                                                _ = v62
                                            case 575:
                                                if len(v61) > 0 {
                                                    v61[0] = fi27
                                                }
                                            default:
                                                fmt.Println(fi28, fi27, v25)
                                            }
                                        } else {
                                            var v63 int16 = 137
                                            var v64 int16 = v63
                                            var v65 float32 = 0.0
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        }
                                    default:
                                        switch fi27 {
                                        case -767:
                                            var v66 bool = false
                                            _ = v66
                                        case 4294967295:
                                            v61 = append(v61, (fi28 % ((-2147483648) | 1)))
                                        default:
                                            if true {
                                                v2[("a" + "foo bar")] = v25
                                                var v67 int64 = (-(v20))
                                                switch "hello" {
                                                case "":
                                                    v67 = 9223372036854775807
                                                default:
                                                    v25 = ((("hello" + ("\n\t" + v1)) + (v25 + v1)) + ("a" + ("hello" + "foo bar")))
                                                }
                                                fmt.Println(v1)
                                                _ = v67
                                            }
                                            fmt.Println(fi28, v25, fi28)
                                            v68 := GenS5{F0: 370, F1: 870, F2: 619}
                                            _ = v68
                                        }
                                    }
                                } else {
                                    fmt.Println(v1, fi28)
                                    v26 += (v26 | v26)
                                }
                            } else {
                                fmt.Println(v1, fi27, v25)
                                if false {
                                    if len(v61) >= 2 {
                                        v61 = v61[1:len(v61)]
                                    }
                                    if (!false) {
                                        var v69 bool = (v20 >= (v20 * (^(v20))))
                                        var v70 uint32 = 2
                                        v71 := GenS1{F0: 2}
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                    } else {
                                        v26 -= 2
                                    }
                                    v25 = v25
                                }
                                fmt.Println(fi27, v20)
                            }
                            v25 = (v25 + v25)
                            _ = v61
                        }
                        if true {
                            v72 := GenS8{F0: false}
                            switch "" {
                            case "foo bar":
                                var v73 uint = 65535
                                _ = v73
                            }
                            _ = v72
                        } else {
                            v74 := []int{(fi27 - fi27), 410, fi27}
                            if true {
                                if true {
                                    var v75 uint = genFn0()
                                    switch fi27 {
                                    case 9223372036854775807:
                                        if true {
                                            if true {
                                                v1 = "a"
                                            } else {
                                                v3[v25] = 675
                                                fmt.Println(v25, fi28)
                                            }
                                            fmt.Println(fi28, v25)
                                            var v76 float64 = 3.14159
                                            _ = v76
                                        }
                                        if false {
                                            var v77 uint16 = 255
                                            _ = v77
                                        }
                                        var v78 int16 = int16(fi28)
                                        _ = v78
                                    default:
                                        v79 := GenS4{F0: 255, F1: 1, F2: (((v1 + (("foo bar" + "") + ((v25 + v25) + v1))) + v1) + (v1 + ("" + ((((((v25 + "a") + v1) + "foo bar") + "a") + "") + (v1 + (("a" + ("hello" + (v25 + ""))) + (((v1 + v25) + v1) + v1)))))))}
                                        _ = v79
                                    }
                                    _ = v75
                                } else {
                                    v20 += (v20 - (v20 / ((0) | 1)))
                                    v80 := GenS7{F0: (fi28 | fi28)}
                                    var v81 int8 = -65
                                    _ = v80
                                    _ = v81
                                }
                                fmt.Println(v20)
                            } else {
                                _ = len(v2)
                                fmt.Println(v25, v26)
                                v20 -= v20
                            }
                            _ = v74
                        }
                    }
                }
            }
            _ = fi28
            _ = v29
        }
        if true {
            if false {
                v26 -= 1
                v82 := GenS11{F0: 0, F1: fi27, F2: 0}
                v83 := GenS11{F0: -128, F1: 367, F2: 100}
                _ = v82
                _ = v83
            } else {
                fmt.Println(v20, fi27)
            }
            fmt.Println(v1, v26)
            fmt.Println(v20)
        }
        _ = fi27
    }
    if ((!false) || true) {
        fmt.Println(v20, v26, v25)
        if true {
            var v84 int = 32767
            v2[v25] = "\n\t"
            _ = v84
        }
        for fi85 := 0; fi85 < 11; fi85++ {
            v86 := GenS12{F0: 807, F1: v20, F2: 0}
            var v87 float32 = (3.14159 * (((((float32((3.14159 * (1.5 / 1e10))) - (-2.25 + (3.14159 * 3.14159))) + (0.1 * (0.0 + 1.5))) / ((1e10 / 1.0) / 1.0)) - (1e10 / 1e10)) - (0.0 / 3.14159)))
            _ = fi85
            _ = v86
            _ = v87
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
