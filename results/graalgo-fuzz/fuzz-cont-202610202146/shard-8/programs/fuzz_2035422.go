package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int32; F2 string }
type GenS3 struct { F0 int32; F1 uint; F2 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 string; F1 uint; F2 bool }
func genFn0(p0 string, p1 int, p2 uint32) uint64 {
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: ((0 != 395) || (!false))}
        fmt.Println(fi5, p1)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 float64 = 0.0
        _ = fi7
        _ = v8
    }
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 4294967295
}
func genFn1() uint8 {
    if (42 > 632) {
        var v10 uint64 = 984
        _ = v10
    } else {
        var v11 uint64 = genFn0("", 2147483647, 42)
        switch "" {
        case "\n\t":
            v11 = (v11 * (v11 ^ v11))
        case "foo bar":
            var v12 int64 = -9223372036854775808
            _ = v12
        case "":
            var v13 int16 = 32767
            _ = v13
        default:
            v11 = (v11 ^ 0)
        }
        fmt.Println(v11)
        _ = v11
    }
    v14 := []int{-2147483648}
    if false {
        var v15 int32 = 2147483647
        v16 := GenS2{F0: 1, F1: 325, F2: ((("hello" + "a") + (("hello" + "") + "foo bar")) + "foo bar")}
        if len(v14) > 0 {
            v14[0] = -981
        }
        _ = v15
        _ = v16
    } else {
        v17 := GenS2{F0: 9223372036854775807, F1: 42, F2: ("\n\t" + "\n\t")}
        _ = v17
    }
    var v18 uint8 = 255
    _ = v14
    _ = v18
    return v18
}

func main() {
    var v1 uint8 = 127
    var v2 int8 = 1
    v3 := make(map[uint32]int)
    var v4 int = -2147483648
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := []int{(fi19 >> ((346) & 63)), (fi19 - 32768), (fi19 ^ (fi19 & 2147483647)), v4}
        for fi21 := 0; fi21 < 11; fi21++ {
            v2 = v2
            _ = fi21
        }
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        _ = fi19
        _ = v20
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 uint16 = 32767
        _ = fi22
        _ = v23
    }
    delete(v3, 2)
    for fi24 := 0; fi24 < 11; fi24++ {
        _ = len(v3)
        var v25 int32 = -713
        if ("\n\t" > (("a" + (((("foo bar" + ("\n\t" + (("" + "foo bar") + "\n\t"))) + ("foo bar" + ("a" + "a"))) + ("hello" + "hello")) + "")) + ("" + "hello"))) {
            var v26 uint32 = 838
            _ = v26
        }
        _ = fi24
        _ = v25
    }
    v27 := GenS2{F0: (v4 >> ((821) & 63)), F1: 677, F2: "hello"}
    for fi28 := 0; fi28 < 16; fi28++ {
        fmt.Println(v4)
        for fi29 := 0; fi29 < 20; fi29++ {
            if false {
                v30 := GenS3{F0: 227, F1: 72, F2: 255}
                _ = v30
            } else {
                if (!true) {
                    var v31 uint = 180
                    _ = v31
                }
                v32 := GenS2{F0: 32767, F1: 1, F2: ("hello" + "\n\t")}
                if ((-2147483648 > 0) || (true || true)) {
                    switch (("foo bar" + ("hello" + "a")) + "\n\t") {
                    case "hello":
                        var v33 int64 = 9223372036854775807
                        _ = v33
                    case "foo bar":
                        v34 := GenS1{F0: (-372 != 2147483647)}
                        _ = v34
                    }
                }
                _ = v32
            }
            _ = fi29
        }
        var v35 uint32 = 65535
        _ = fi28
        _ = v35
    }
    v2 = 127
    fmt.Println(v4)
    v36 := GenS3{F0: -384, F1: 473, F2: 65535}
    if ((v2 ^ (v2 | (v2 / (((^(v2))) | 1)))) < 42) {
        switch (("foo bar" + ("a" + "foo bar")) + "") {
        case "":
            if false {
                for fi37 := 0; fi37 < 16; fi37++ {
                    v2 = 1
                    for fi38 := 0; fi38 < 4; fi38++ {
                        v4 = (^(fi38))
                        v1 = genFn1()
                        _ = fi38
                    }
                    _ = fi37
                }
                v2 = v2
            }
        case "foo bar":
            for fi39 := 0; fi39 < 16; fi39++ {
                if ((!false) || ((-2.25 / ((1.5 - 0.1) * 0.1)) <= (3.14159 - (3.14159 / -2.25)))) {
                    fmt.Println(fi39)
                    var v40 float32 = float32(float64(float32(0.1)))
                    _ = v40
                }
                switch (fi39 % (((^(fi39))) | 1)) {
                case 9223372036854775807:
                    v41 := []int{(fi39 / ((127) | 1)), 42, v4, v4}
                    _ = v41
                case 42:
                    if false {
                        if (v2 == 55) {
                            fmt.Println(v4)
                        } else {
                            var v42 float64 = 0.1
                            var v43 float32 = ((((float32(v42) / 0.1) + 3.14159) - float32(((0.1 * (1e10 + v42)) + v42))) - 3.14159)
                            _ = v42
                            _ = v43
                        }
                        delete(v3, 4294967295)
                        for fi44 := 0; fi44 < 20; fi44++ {
                            if false {
                                if (!false) {
                                    fmt.Println(v2, fi44)
                                }
                                var v45 float64 = 3.14159
                                v46 := GenS4{F0: 425}
                                if ("" < "hello") {
                                    v45 = (v45 - v45)
                                    if true {
                                        var v47 uint8 = (^(v1))
                                        _ = v47
                                    }
                                    switch (-(fi39)) {
                                    case 578:
                                        var v48 uint8 = genFn1()
                                        _ = v48
                                    }
                                }
                                _ = v45
                                _ = v46
                            } else {
                                v4 = v4
                            }
                            fmt.Println(fi44, v1, v2)
                            _ = fi44
                        }
                    }
                case 2147483647:
                    if false {
                        v2 = v2
                        for fi49 := 0; fi49 < 3; fi49++ {
                            fmt.Println(v4)
                            v50 := []int{}
                            v3[42] = 588
                            _ = fi49
                            _ = v50
                        }
                        var v51 uint64 = 1
                        _ = v51
                    } else {
                        v2 += (v2 << ((255) & 63))
                    }
                }
                _ = fi39
            }
        case "\n\t":
            if false {
                var v52 float32 = 1.5
                v53 := GenS2{F0: -998, F1: -1, F2: (((("" + (("" + "foo bar") + ((("hello" + "a") + ("hello" + "")) + ""))) + (((("hello" + "\n\t") + ("hello" + ("foo bar" + "a"))) + "a") + "")) + "a") + ("" + "a"))}
                for fi54 := 0; fi54 < 18; fi54++ {
                    var v55 int8 = (v2 | (v2 ^ (v2 + (-(v2)))))
                    for fi56 := 0; fi56 < 9; fi56++ {
                        switch "" {
                        case "foo bar":
                            if (!true) {
                                v52 = v52
                            } else {
                                var v57 bool = ((v1 & (v1 & (v1 - 253))) <= v1)
                                _ = v57
                            }
                        case "":
                            v52 = v52
                        case "hello":
                            v58 := GenS5{F0: (v55 % ((-128) | 1)), F1: (-1 == 0)}
                            _ = v58
                        default:
                            if (695 <= 2147483647) {
                                switch "\n\t" {
                                case "":
                                    if (710 >= 1) {
                                        v2 += v55
                                        v4 -= 127
                                    } else {
                                        fmt.Println(v55, v2)
                                        if false {
                                            if true {
                                                v59 := GenS3{F0: 2147483647, F1: 42, F2: 491}
                                                if false {
                                                    fmt.Println(v55)
                                                    v60 := []int{}
                                                    _ = v60
                                                }
                                                _ = v59
                                            } else {
                                                if (32767 >= 820) {
                                                    v4 -= fi56
                                                }
                                            }
                                            var v61 string = ""
                                            var v62 uint32 = 0
                                            _ = v61
                                            _ = v62
                                        }
                                        v63 := []int{fi56, v4, v4}
                                        _ = v63
                                    }
                                case "foo bar":
                                    var v64 float32 = 1.5
                                    _ = v64
                                default:
                                    v65 := []int{(fi54 + fi56), v4, (fi54 >> ((42) & 63)), (fi56 / ((fi56) | 1))}
                                    _ = v65
                                }
                                fmt.Println(v4, v52)
                                var v66 uint32 = 42
                                _ = v66
                            } else {
                                if (61 > -2147483648) {
                                    v55 = int8(v1)
                                }
                            }
                        }
                        if false {
                            v55 = 127
                            fmt.Println(v55, fi56, v4)
                            v67 := GenS6{F0: 42, F1: 1}
                            _ = v67
                        }
                        _ = fi56
                    }
                    _ = fi54
                    _ = v55
                }
                _ = v52
                _ = v53
            } else {
                v68 := GenS7{F0: "", F1: 4294967295, F2: true}
                var v69 int = (v4 * v4)
                _ = v68
                _ = v69
            }
        }
        var v70 uint8 = v1
        var v71 uint8 = 255
        _ = v70
        _ = v71
    } else {
        for fi72 := 0; fi72 < 3; fi72++ {
            var v73 float64 = -2.25
            if true {
                var v74 float64 = v73
                v73 -= (v74 / ((v74 - (1.5 * (1e10 - v74))) * v73))
                _ = v74
            }
            if (("hello" + ("foo bar" + ((("" + "a") + "\n\t") + "\n\t"))) < "hello") {
                v75 := v3[2147483647]
                _ = v75
            } else {
                var v76 uint8 = v1
                _ = v76
            }
            _ = fi72
            _ = v73
        }
        v2 += v2
        var v77 uint16 = 255
        _ = v77
    }
    fmt.Println(v1, v1, v2)
    var v78 uint = 255
    var v79 float32 = 1.5
    v80 := GenS5{F0: 42, F1: true}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v78
    _ = v79
    _ = v80
}
