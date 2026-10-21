package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 int; F2 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint }
func genFn0(p0 int16, p1 uint8) int32 {
    if false {
        for fi4 := 0; fi4 < 1; fi4++ {
            var v5 bool = false
            for fi6 := 0; fi6 < 6; fi6++ {
                fmt.Println(p0, fi4)
                v5 = false
                _ = fi6
            }
            if v5 {
                var v7 uint = 338
                _ = v7
            }
            _ = fi4
            _ = v5
        }
    }
    var v8 int32 = 42
    _ = p0
    _ = p1
    _ = v8
    return (v8 >> ((255) & 63))
}
func genFn1(p0 uint, p1 string) uint8 {
    var v9 float64 = (1e10 + 1.5)
    if ((false && false) && false) {
        for fi10 := 0; fi10 < 5; fi10++ {
            v9 -= v9
            var v11 int8 = -48
            switch (p1 + "") {
            case "":
                switch fi10 {
                case 32767:
                    var v12 uint8 = 168
                    _ = v12
                }
            case "\n\t":
                fmt.Println(p0, p0)
            case "hello":
                switch -2147483648 {
                case -2147483648:
                    var v13 float64 = 0.0
                    _ = v13
                case 128:
                    fmt.Println(v11, p0)
                default:
                    p1 = ""
                }
            default:
                v14 := []int{fi10}
                _ = v14
            }
            _ = fi10
            _ = v11
        }
        v9 = 0.0
        fmt.Println(p1)
    } else {
        p1 = ("foo bar" + ((p1 + "foo bar") + p1))
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 uint8 = 255
        var v17 bool = (!false)
        var v18 float32 = (0.0 / 0.1)
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 0, F1: ((p1 + p1) + p1)}
    _ = p0
    _ = p1
    _ = v19
    _ = v9
    return 90
}

func main() {
    var v1 string = "a"
    var v2 uint16 = 32767
    v3 := make(map[string]int)
    for fi20 := 0; fi20 < 15; fi20++ {
        v1 = ""
        fmt.Println(fi20, v2, v1)
        v2 -= v2
        _ = fi20
    }
    var v21 int32 = 42
    v1 = "hello"
    if (!((v21 ^ v21) < 0)) {
        v21 += -1
    }
    if ((false || (620 > 0)) && false) {
        v22 := v3[(v1 + v1)]
        var v23 uint32 = 2147483647
        _ = v22
        _ = v23
    }
    if (!((((float32(0.0) + (((3.14159 + 1e10) * -2.25) + 0.0)) > 0.1) || (1 >= 284)) || true)) {
        if true {
            if true {
                if false {
                    v24 := GenS1{F0: 0, F1: v1}
                    if (genFn0(364, 0) != genFn0(42, 211)) {
                        for fi25 := 0; fi25 < 5; fi25++ {
                            var v26 int64 = -992
                            if (0 == 2) {
                                fmt.Println(v26, fi25)
                                var v27 uint32 = 1
                                _ = v27
                            } else {
                                fmt.Println(fi25, v2, v2)
                                var v28 int16 = 42
                                _ = v28
                            }
                            _ = fi25
                            _ = v26
                        }
                        if true {
                            v29 := GenS1{F0: -20, F1: v1}
                            v21 += genFn0(-32768, 255)
                            _ = v29
                        }
                        switch -268 {
                        case 0:
                            for fi30 := 0; fi30 < 17; fi30++ {
                                var v31 string = ((v1 + (("foo bar" + ((v1 + v1) + ("hello" + (v1 + v1)))) + v1)) + (v1 + (("foo bar" + v1) + v1)))
                                v32 := []int{-885}
                                if (false || ((int64(v2) >= 42) && false)) {
                                    var v33 uint8 = 42
                                    if ((!false) && true) {
                                        v3[v31] = (fi30 & (fi30 | -1))
                                    }
                                    switch (((v31 + v1) + v31) + v1) {
                                    case "foo bar":
                                        v1 = ""
                                    case "hello":
                                        v33 = 105
                                    case "\n\t":
                                        var v34 uint64 = 18446744073709551615
                                        _ = v34
                                    }
                                    _ = v33
                                }
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                    } else {
                        v21 = (v21 * (v21 | v21))
                        var v35 int16 = 32767
                        _ = v35
                    }
                    var v36 float64 = (float64(0.0) * -2.25)
                    _ = v24
                    _ = v36
                }
                if false {
                    fmt.Println(v1, v2, v21)
                    if true {
                        v37 := v3[v1]
                        if true {
                            v38 := []int{v37, (v37 & v37), -798}
                            _ = v38
                        } else {
                            v21 = v21
                            v39 := v3["foo bar"]
                            switch v1 {
                            case "\n\t":
                                for fi40 := 0; fi40 < 3; fi40++ {
                                    var v41 int8 = 0
                                    for fi42 := 0; fi42 < 12; fi42++ {
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v43 float32 = 1e10
                                            _ = v43
                                        case "hello":
                                            v44 := GenS2{F0: 2}
                                            _ = v44
                                        case "":
                                            var v45 string = ("\n\t" + "foo bar")
                                            v3[("a" + "\n\t")] = fi40
                                            switch (v39 - (-(fi42))) {
                                            case 32768:
                                                if ((v45 + "\n\t") <= "hello") {
                                                    var v46 int16 = 42
                                                    v2 = 32767
                                                    _ = v46
                                                } else {
                                                    v3["a"] = 1
                                                    v47 := GenS3{F0: fi40, F1: 393, F2: v1}
                                                    _ = v47
                                                }
                                            case 20:
                                                v48 := []int{(fi42 | fi42), (v37 >> ((42) & 63)), v37, (^(v39))}
                                                _ = v48
                                            }
                                            _ = v45
                                        }
                                        _ = fi42
                                    }
                                    switch v37 {
                                    case 128:
                                        var v49 string = v1
                                        _ = v49
                                    case 27:
                                        v50 := GenS4{F0: v41}
                                        _ = v50
                                    case 4294967295:
                                        v51 := []int{(-(fi40)), fi40, 858}
                                        _ = v51
                                    default:
                                        var v52 float64 = (1.5 + (3.14159 + 0.0))
                                        _ = v52
                                    }
                                    _ = fi40
                                    _ = v41
                                }
                            case "a":
                                var v53 uint32 = 2147483648
                                _ = v53
                            case "foo bar":
                                v54 := GenS1{F0: -9223372036854775808, F1: ("hello" + "\n\t")}
                                _ = v54
                            }
                            _ = v39
                        }
                        _ = v37
                    }
                    var v55 int16 = 1000
                    _ = v55
                }
                switch 665 {
                case 32768:
                    fmt.Println(v1, v2, v1)
                case 32767:
                    var v56 uint32 = 687
                    _ = v56
                }
            }
        }
        for fi57 := 0; fi57 < 20; fi57++ {
            v58 := []int{}
            if (false || (787 != 42)) {
                var v59 int32 = -1
                switch v1 {
                case "hello":
                    v59 = -1
                }
                _ = v59
            } else {
                fmt.Println(v2)
                fmt.Println(v1)
            }
            if true {
                v21 -= v21
                if true {
                    switch fi57 {
                    case 1:
                        var v60 float64 = float64(float32((-2.25 - 0.1)))
                        _ = v60
                    default:
                        v1 = v1
                    }
                    for fi61 := 0; fi61 < 9; fi61++ {
                        var v62 float32 = 1e10
                        _ = fi61
                        _ = v62
                    }
                }
                for fi63 := 0; fi63 < 18; fi63++ {
                    fmt.Println(v2, v1, v1)
                    if (!(0.1 >= (-2.25 - 3.14159))) {
                        fmt.Println(v2, fi63)
                        v21 -= int32(fi57)
                        if (false || ((!(!false)) || false)) {
                            if (v21 > -1) {
                                if false {
                                    fmt.Println(v1, fi63, v21)
                                    var v64 int32 = (v21 << ((4294967295) & 63))
                                    v65 := v3[v1]
                                    _ = v64
                                    _ = v65
                                }
                                v66 := GenS1{F0: -9223372036854775808, F1: (v1 + (v1 + "hello"))}
                                v1 = ""
                                _ = v66
                            }
                            v67 := GenS5{F0: 18446744073709551615}
                            _ = v67
                        }
                    }
                    _ = fi63
                }
            }
            _ = fi57
            _ = v58
        }
    }
    v1 = ((v1 + (v1 + v1)) + "a")
    if ((false || true) && true) {
        v2 += v2
        if true {
            v3[v1] = -116
            v1 = v1
            for fi68 := 0; fi68 < 12; fi68++ {
                for fi69 := 0; fi69 < 13; fi69++ {
                    var v70 int16 = 1
                    v71 := []int{}
                    _ = fi69
                    _ = v70
                    _ = v71
                }
                var v72 float64 = 1.5
                v3[""] = -595
                _ = fi68
                _ = v72
            }
        } else {
            v2 = (v2 % (((v2 ^ 488)) | 1))
            fmt.Println(v1, v21)
        }
        fmt.Println(v21)
    } else {
        var v73 uint = 2
        if (!true) {
            delete(v3, v1)
            var v74 int32 = 42
            _ = v74
        } else {
            v1 = "foo bar"
        }
        _ = v73
    }
    var v75 int8 = 100
    switch ("\n\t" + v1) {
    case "hello":
        var v76 uint = 101
        _ = v76
    default:
        fmt.Println(v21, v75)
    }
    for fi77 := 0; fi77 < 4; fi77++ {
        v2 -= 791
        fmt.Println(v21, v2)
        if (2 >= 4294967295) {
            switch v1 {
            case "\n\t":
                switch ((v1 + ((v1 + ("" + (v1 + "a"))) + ("" + "a"))) + v1) {
                case "hello":
                    v78 := []int{(fi77 + (fi77 * fi77)), 128}
                    _ = v78
                case "\n\t":
                    var v79 float32 = 1.5
                    _ = v79
                default:
                    v80 := []int{fi77, (fi77 ^ fi77), fi77, (fi77 >> ((1) & 63))}
                    _ = v80
                }
            case "foo bar":
                fmt.Println(v75, v75, v75)
            default:
                _ = len(v3)
            }
            v81 := []int{32767}
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v21, v75, v2)
            }
            _ = v81
        }
        _ = fi77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v75
}
