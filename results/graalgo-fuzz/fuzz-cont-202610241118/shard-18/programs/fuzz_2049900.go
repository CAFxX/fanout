package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 string }
type GenS3 struct { F0 int16; F1 uint; F2 uint }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 int16; F2 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string; F1 int64; F2 uint16 }
type GenS9 struct { F0 int8; F1 uint32 }
type GenS10 struct { F0 uint32 }
func genFn0(p0 int8, p1 int16, p2 bool) int {
    if p2 {
        var v7 int16 = p1
        if (-432 != 683) {
            v8 := []int{32767, 777}
            fmt.Println(v7, p1)
            _ = v8
        }
        fmt.Println(p0)
        _ = v7
    } else {
        v9 := []int{0, 127}
        var v10 int8 = 42
        for fi11 := 0; fi11 < 1; fi11++ {
            if (45 > 0) {
                v12 := GenS1{F0: p1}
                _ = v12
            }
            _ = fi11
        }
        _ = v10
        _ = v9
    }
    var v13 uint16 = 914
    fmt.Println(p1, p0, p2)
    v14 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return -9223372036854775808
}
func genFn1(p0 int64, p1 uint32, p2 int64) int {
    p2 += p0
    switch "foo bar" {
    case "a":
        v15 := GenS1{F0: 42}
        _ = v15
    case "hello":
        p1 = 2147483647
    case "\n\t":
        for fi16 := 0; fi16 < 11; fi16++ {
            if true {
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 int = 254
                    var v19 string = "hello"
                    if ((v18 >> ((586) & 63)) <= (v18 << ((uint(p1)) & 63))) {
                        var v20 int16 = -60
                        var v21 int16 = v20
                        _ = v20
                        _ = v21
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v22 := GenS1{F0: -385}
                if (1e10 == -2.25) {
                    switch ("hello" + ("hello" + ("" + "\n\t"))) {
                    case "hello":
                        if ("" > (("" + ((("a" + "hello") + (("a" + "hello") + "")) + ("hello" + "\n\t"))) + "hello")) {
                            v23 := []int{fi16, 127, genFn0(int8(p0), 1, true)}
                            for fi24 := 0; fi24 < 1; fi24++ {
                                var v25 uint8 = uint8(p1)
                                v26 := GenS2{F0: false, F1: (1 >= 18446744073709551615), F2: "\n\t"}
                                p1 += p1
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            v27 := GenS3{F0: 0, F1: 4294967295, F2: 791}
                            _ = v23
                            _ = v27
                        } else {
                            switch (("a" + "\n\t") + (("foo bar" + ("\n\t" + ("foo bar" + "a"))) + (((("hello" + ("\n\t" + (("" + (("" + ((("\n\t" + "a") + "foo bar") + (("foo bar" + "foo bar") + ("foo bar" + ("a" + ("a" + "\n\t")))))) + "\n\t")) + (("" + "hello") + ("hello" + "\n\t"))))) + "foo bar") + "a") + ""))) {
                            case "a":
                                for fi28 := 0; fi28 < 20; fi28++ {
                                    p0 -= int64(fi16)
                                    var v29 string = ("" + "foo bar")
                                    var v30 uint32 = p1
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                            case "hello":
                                if true {
                                    v31 := []int{}
                                    fmt.Println(fi16, p0, fi16)
                                    for fi32 := 0; fi32 < 9; fi32++ {
                                        fmt.Println(fi16, p1)
                                        _ = fi32
                                    }
                                    _ = v31
                                }
                            case "foo bar":
                                v33 := GenS4{F0: 98, F1: (("hello" + ("hello" + (("hello" + "hello") + ("foo bar" + "a")))) + ("hello" + "\n\t"))}
                                _ = v33
                            }
                        }
                    case "\n\t":
                        var v34 uint = 0
                        _ = v34
                    }
                    if false {
                        p0 = p0
                        if (4294967295 != 4294967295) {
                            var v35 uint64 = 4294967295
                            fmt.Println(p0, p2, p1)
                            _ = v35
                        } else {
                            if false {
                                var v36 bool = (((!((true || (1.5 > 3.14159)) || (-2147483648 <= -2147483648))) || ((float32(0.0) + float32(-2.25)) <= 0.1)) && (0 < 115))
                                _ = v36
                            } else {
                                v37 := []int{(fi16 ^ 127), 32767, 2}
                                _ = v37
                            }
                            var v38 float64 = (1.5 / 3.14159)
                            _ = v38
                        }
                    }
                } else {
                    p1 -= (p1 << ((0) & 63))
                    if (!(3.14159 != 1.5)) {
                        fmt.Println(p0, p0, p2)
                        var v39 string = ((("" + "foo bar") + "a") + "hello")
                        var v40 int32 = 42
                        _ = v39
                        _ = v40
                    } else {
                        p2 += -9223372036854775808
                        var v41 uint8 = 149
                        _ = v41
                    }
                }
                _ = v22
            }
            v42 := GenS5{F0: 42}
            _ = fi16
            _ = v42
        }
    default:
        if false {
            if true {
                for fi43 := 0; fi43 < 3; fi43++ {
                    var v44 int64 = (p0 & p0)
                    _ = fi43
                    _ = v44
                }
                v45 := GenS3{F0: -340, F1: 284, F2: 2}
                v46 := GenS3{F0: -914, F1: 241, F2: 2}
                _ = v45
                _ = v46
            } else {
                switch ("" + "foo bar") {
                case "\n\t":
                    v47 := GenS5{F0: -2147483648}
                    _ = v47
                case "":
                    var v48 uint8 = 2
                    _ = v48
                case "a":
                    var v49 uint64 = uint64(p1)
                    _ = v49
                }
                p2 = (^(p2))
                switch 127 {
                case 32768:
                    p2 += (^(p0))
                case 0:
                    v50 := GenS2{F0: false, F1: true, F2: (("foo bar" + "hello") + "a")}
                    _ = v50
                }
            }
            v51 := GenS6{F0: 255, F1: -1, F2: 175}
            _ = v51
        } else {
            if false {
                v52 := GenS5{F0: 52}
                v53 := []int{}
                _ = v52
                _ = v53
            }
            fmt.Println(p1, p1)
        }
    }
    var v54 int = 0
    v55 := []int{(^(v54)), (v54 - v54), 2147483647, v54}
    _ = p0
    _ = p1
    _ = p2
    _ = v54
    _ = v55
    return genFn0(-1, 32767, false)
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    var v3 uint64 = 4294967295
    var v4 int64 = 42
    var v5 uint8 = 72
    var v6 uint64 = 18446744073709551615
    var v56 bool = true
    v57 := GenS7{F0: 42}
    for fi58 := 0; fi58 < 1; fi58++ {
        v59 := GenS8{F0: "\n\t", F1: (v4 % (((v4 << ((4294967295) & 63))) | 1)), F2: 319}
        switch "hello" {
        case "foo bar":
            if ("\n\t" > "\n\t") {
                switch fi58 {
                case -1:
                    var v60 bool = v56
                    _ = v60
                case -74:
                    fmt.Println(v3, v3)
                default:
                    if ("hello" != (("" + "foo bar") + "foo bar")) {
                        v61 := GenS7{F0: -1}
                        fmt.Println(fi58, v56)
                        v62 := []int{}
                        _ = v61
                        _ = v62
                    }
                }
            }
        case "":
            v63 := GenS5{F0: 1}
            _ = v63
        case "hello":
            if ((true && v56) && (v4 <= (v4 - v4))) {
                fmt.Println(v56, v4, v3)
                v64 := GenS9{F0: 1, F1: 1}
                var v65 int16 = -224
                _ = v64
                _ = v65
            }
        default:
            var v66 float64 = 1.5
            _ = v66
        }
        _ = fi58
        _ = v59
    }
    v56 = v56
    v67 := GenS10{F0: 2147483647}
    var v68 uint8 = v5
    fmt.Println(v3, v3)
    v68 -= v68
    v69 := []int{2147483647, 9223372036854775807}
    v4 += 9223372036854775807
    v70 := GenS6{F0: 65535, F1: 1000, F2: v5}
    v2[""] = v56
    var v71 int8 = -103
    _ = len(v69)
    _ = len(v2)
    if (((!v56) || v56) || ((v6 != (^(v6))) && ((!false) && true))) {
        var v72 int16 = -32768
        _ = v72
    }
    if (-32768 != -32768) {
        if (42 == 65535) {
            v4 = v4
        }
    } else {
        var v73 string = ""
        v71 += v71
        switch 127 {
        case -2147483648:
            switch 9223372036854775807 {
            case 130:
                if v56 {
                    fmt.Println(v6, v73)
                }
            case 0:
                fmt.Println(v6, v4, v4)
            default:
                _ = len(v2)
            }
        case 32768:
            switch int(v71) {
            case -9223372036854775808:
                v68 = (v68 | v68)
            }
        case 2147483647:
            v69 = append(v69, 2147483647)
        default:
            fmt.Println(v3)
        }
        _ = v73
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v69:", len(v69), v69)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v6
    _ = v67
    _ = v68
    _ = v69
    _ = v70
    _ = v71
}
