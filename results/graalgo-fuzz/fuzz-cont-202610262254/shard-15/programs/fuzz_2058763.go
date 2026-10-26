package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint64; F2 int16 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int; F1 int8; F2 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint }
func genFn0(p0 bool, p1 int) int {
    for fi4 := 0; fi4 < 9; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            switch "a" {
            case "hello":
                var v6 uint = 42
                _ = v6
            }
            switch 750 {
            case 0:
                v7 := GenS1{F0: 100}
                _ = v7
            case -102:
                v8 := GenS1{F0: -53}
                _ = v8
            default:
                v9 := []int{}
                p0 = false
                if (p0 && true) {
                    v10 := GenS1{F0: -54}
                    p1 -= fi4
                    _ = v10
                }
                _ = v9
            }
            p0 = p0
            _ = fi5
        }
        if false {
            v11 := []int{}
            fmt.Println(p1, fi4)
            _ = v11
        } else {
            if p0 {
                if (false || p0) {
                    var v12 uint64 = 18446744073709551615
                    switch "a" {
                    case "":
                        v13 := []int{(p1 - fi4)}
                        _ = v13
                    case "a":
                        if (-1 < -1) {
                            v14 := []int{(p1 * (fi4 % ((32767) | 1))), (fi4 << ((224) & 63)), p1}
                            _ = v14
                        } else {
                            if ((42 != -1) || (float64(float32(0.1)) >= float64((1.5 / 1e10)))) {
                                v15 := GenS1{F0: 1}
                                _ = v15
                            }
                            switch p1 {
                            case 128:
                                p0 = ((v12 ^ (^(v12))) < v12)
                            case 42:
                                for fi16 := 0; fi16 < 3; fi16++ {
                                    v12 += (^(v12))
                                    if p0 {
                                        fmt.Println(v12)
                                        v17 := GenS1{F0: 87}
                                        var v18 int8 = 0
                                        _ = v17
                                        _ = v18
                                    } else {
                                        var v19 uint = 637
                                        _ = v19
                                    }
                                    v20 := GenS1{F0: 31}
                                    _ = fi16
                                    _ = v20
                                }
                            case 4294967295:
                                for fi21 := 0; fi21 < 8; fi21++ {
                                    switch "hello" {
                                    case "\n\t":
                                        v22 := GenS2{F0: 127}
                                        _ = v22
                                    case "hello":
                                        switch "\n\t" {
                                        case "\n\t":
                                            fmt.Println(fi21)
                                        case "":
                                            if p0 {
                                                v23 := []int{p1, fi21, fi4, fi4}
                                                _ = v23
                                            }
                                        case "hello":
                                            var v24 uint16 = 0
                                            _ = v24
                                        default:
                                            var v25 uint32 = 947
                                            _ = v25
                                        }
                                    case "":
                                        v26 := []int{fi21, (fi21 % (((-(fi21))) | 1)), -9223372036854775808}
                                        _ = v26
                                    }
                                    _ = fi21
                                }
                            default:
                                switch (fi4 | 836) {
                                case 4294967295:
                                    v27 := GenS1{F0: -1}
                                    _ = v27
                                case 127:
                                    fmt.Println(p0, fi4, v12)
                                case 42:
                                    var v28 bool = (!true)
                                    _ = v28
                                default:
                                    if true {
                                        v29 := GenS3{F0: -128, F1: (^(v12)), F2: 32767}
                                        switch (p1 - -126) {
                                        case -9223372036854775808:
                                            p0 = p0
                                        }
                                        var v30 uint32 = 1
                                        _ = v29
                                        _ = v30
                                    }
                                }
                            }
                        }
                    }
                    _ = v12
                } else {
                    var v31 int64 = 1
                    _ = v31
                }
            }
        }
        v32 := []int{(fi4 ^ 2147483647), p1, (fi4 | (p1 & fi4))}
        _ = fi4
        _ = v32
    }
    p1 -= (-(p1))
    for fi33 := 0; fi33 < 8; fi33++ {
        v34 := GenS2{F0: 255}
        p1 = 9223372036854775807
        p0 = p0
        _ = fi33
        _ = v34
    }
    _ = p0
    _ = p1
    return (p1 / (((p1 / (((p1 ^ 4294967295)) | 1))) | 1))
}
func genFn1() uint64 {
    v35 := []int{}
    var v36 int32 = 0
    v36 += v36
    v36 += v36
    _ = v35
    _ = v36
    return 375
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    var v3 bool = false
    v37 := GenS4{F0: "hello", F1: 32767}
    v38 := GenS2{F0: 205}
    switch ((("foo bar" + "a") + "a") + (("foo bar" + "hello") + "\n\t")) {
    case "foo bar":
        fmt.Println(v3)
    default:
        _ = len(v1)
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        v40 := []int{(-(fi39))}
        switch "hello" {
        case "":
            var v41 uint8 = 100
            _ = v41
        case "a":
            fmt.Println(fi39)
        }
        if (!v3) {
            switch (fi39 ^ (fi39 | (fi39 / (((fi39 ^ fi39)) | 1)))) {
            case 4294967295:
                delete(v1, genFn0(((100 <= -128) && v3), (fi39 / ((32767) | 1))))
            }
            var v42 string = "hello"
            v3 = (false || (4294967295 != 533))
            _ = v42
        }
        _ = fi39
        _ = v40
    }
    if (((((-2.25 / 1.0) / ((1.5 * float64(0.1)) / ((1.5 + (1e10 + ((1e10 + 1.5) / (1.5 - float64(-2.25))))) + 0.1))) * 1e10) * (1.5 - -2.25)) > -2.25) {
        v2[127] = ("hello" + (("\n\t" + ("foo bar" + "foo bar")) + "hello"))
        v2[620] = ("foo bar" + "\n\t")
    } else {
        var v43 bool = (591 == -1)
        fmt.Println(v3)
        _ = v43
    }
    var v44 bool = v3
    if (810 != 255) {
        v44 = v3
    }
    if (!v3) {
        v45 := []int{}
        var v46 int64 = -9223372036854775808
        if false {
            if v3 {
                v47 := GenS5{F0: -2147483648, F1: 41, F2: ""}
                v48 := GenS3{F0: -83, F1: 42, F2: 537}
                v2[int(v46)] = (((((("\n\t" + (("foo bar" + "foo bar") + "\n\t")) + "a") + "\n\t") + "") + "a") + (("\n\t" + "a") + (("hello" + "") + ("hello" + "hello"))))
                _ = v47
                _ = v48
            }
            v46 -= -1
            var v49 uint64 = 0
            _ = v49
        }
        _ = v45
        _ = v46
    } else {
        var v50 int16 = 32767
        v44 = (!v44)
        _ = v50
    }
    if true {
        switch genFn0(v44, -79) {
        case 2147483647:
            var v51 bool = v44
            _ = v51
        case 127:
            for fi52 := 0; fi52 < 12; fi52++ {
                var v53 uint8 = 127
                v44 = ((42 < 2147483647) || v44)
                for fi54 := 0; fi54 < 8; fi54++ {
                    if v3 {
                        if ((true && (("" + ("a" + "\n\t")) >= "a")) || (!true)) {
                            if true {
                                v55 := v2[(-(fi54))]
                                _ = v55
                            } else {
                                fmt.Println(v53)
                                v3 = v3
                            }
                            v56 := []int{}
                            var v57 uint32 = 1
                            _ = v56
                            _ = v57
                        } else {
                            if v3 {
                                if v3 {
                                    v58 := GenS1{F0: -128}
                                    fmt.Println(v44)
                                    _ = v58
                                }
                                fmt.Println(v53, fi54, fi54)
                            }
                        }
                        var v59 int32 = 42
                        var v60 uint8 = v53
                        _ = v59
                        _ = v60
                    }
                    var v61 bool = (-57 != -65)
                    v3 = true
                    fmt.Println(v53, fi54)
                    var v62 uint32 = 153
                    _ = fi54
                    _ = v61
                    _ = v62
                }
                _ = fi52
                _ = v53
            }
        case -9223372036854775808:
            switch 583 {
            case 361:
                if false {
                    var v63 float32 = float32(((3.14159 * 0.1) - 1.5))
                    _ = v63
                }
            case 4294967295:
                switch 4294967295 {
                case 127:
                    v64 := GenS6{F0: genFn1()}
                    _ = v64
                default:
                    v65 := GenS3{F0: 112, F1: 42, F2: 42}
                    _ = v65
                }
            }
        }
        if (0.1 >= (-2.25 / float64((((float32(0.0) + 1.5) / float32((-2.25 * ((0.1 / 3.14159) - (float64(((1.5 - 0.0) * (-2.25 * (1e10 + 0.0)))) / (1e10 - 1.5)))))) * 1.5)))) {
            for fi66 := 0; fi66 < 1; fi66++ {
                var v67 uint = 1
                _ = fi66
                _ = v67
            }
            v3 = v44
            v3 = (4294967295 <= 727)
        } else {
            for fi68 := 0; fi68 < 20; fi68++ {
                var v69 string = ""
                _ = fi68
                _ = v69
            }
            for fi70 := 0; fi70 < 1; fi70++ {
                v44 = v44
                fmt.Println(fi70)
                v71 := []int{}
                _ = fi70
                _ = v71
            }
            v72 := []int{}
            _ = v72
        }
        v44 = true
    } else {
        var v73 uint64 = genFn1()
        switch 128 {
        case 994:
            for fi74 := 0; fi74 < 20; fi74++ {
                v75 := GenS7{F0: 1}
                for fi76 := 0; fi76 < 15; fi76++ {
                    v77 := GenS3{F0: 21, F1: genFn1(), F2: -795}
                    if true {
                        v3 = true
                        var v78 uint32 = 4294967295
                        var v79 uint16 = 151
                        _ = v78
                        _ = v79
                    }
                    _ = fi76
                    _ = v77
                }
                v80 := GenS2{F0: 255}
                _ = fi74
                _ = v75
                _ = v80
            }
        case 32767:
            switch genFn0(true, 950) {
            case 32768:
                for fi81 := 0; fi81 < 18; fi81++ {
                    for fi82 := 0; fi82 < 3; fi82++ {
                        var v83 uint32 = 65535
                        v2[32768] = ("" + "foo bar")
                        v73 = (v73 >> ((65535) & 63))
                        var v84 bool = (false || ((0 >= 183) && (227 >= 1)))
                        fmt.Println(fi81, v83, v44)
                        _ = fi82
                        _ = v83
                        _ = v84
                    }
                    fmt.Println(fi81, v3)
                    v85 := []int{fi81}
                    _ = fi81
                    _ = v85
                }
            default:
                v73 = (v73 & v73)
            }
        default:
            fmt.Println(v44)
        }
        v44 = (v44 && true)
        _ = v73
    }
    for fi86 := 0; fi86 < 14; fi86++ {
        fmt.Println(v3, v3)
        v44 = v44
        _ = fi86
    }
    v1[-9223372036854775808] = ("a" + ("foo bar" + ""))
    switch ("" + ((("a" + ((("" + "foo bar") + "") + ("hello" + ("hello" + "a")))) + "") + "\n\t")) {
    case "":
        var v87 uint8 = 127
        _ = v87
    case "a":
        for fi88 := 0; fi88 < 6; fi88++ {
            var v89 int32 = 42
            var v90 float64 = (-2.25 / (1e10 * 1e10))
            for fi91 := 0; fi91 < 16; fi91++ {
                v90 = v90
                _ = fi91
            }
            _ = fi88
            _ = v89
            _ = v90
        }
    case "hello":
        var v92 uint8 = 213
        _ = v92
    default:
        var v93 uint8 = 42
        _ = v93
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v44
}
