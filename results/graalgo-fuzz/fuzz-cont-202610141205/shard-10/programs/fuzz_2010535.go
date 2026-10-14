package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint64; F1 int8; F2 int8 }
type GenS5 struct { F0 string; F1 string }
func genFn0(p0 uint64, p1 int) int64 {
    var v7 string = ("a" + ((("foo bar" + (("\n\t" + "") + ((("" + "a") + "") + "hello"))) + "\n\t") + ("" + "\n\t")))
    var v8 int16 = -32768
    var v9 int16 = (v8 / ((0) | 1))
    var v10 uint32 = 65535
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -9223372036854775808
}
func genFn1() int8 {
    if (!(((!(!true)) || false) || (-725 <= 32767))) {
        if true {
            switch 32767 {
            case -642:
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 int64 = -861
                    if (float32(0.0) >= 3.14159) {
                        switch "hello" {
                        case "\n\t":
                            v12 -= (^(v12))
                        case "foo bar":
                            fmt.Println(v12, v12)
                        }
                        if false {
                            fmt.Println(fi11, v12)
                        } else {
                            fmt.Println(v12, fi11)
                            for fi13 := 0; fi13 < 19; fi13++ {
                                fmt.Println(fi11, v12)
                                var v14 float32 = float32(((1e10 * (0.0 / 1.0)) * ((-2.25 - 0.1) / float64(((1e10 / 1.5) - 1.5)))))
                                switch fi13 {
                                case 127:
                                    v15 := []int{}
                                    _ = v15
                                case -1:
                                    v16 := GenS1{F0: 102, F1: 0, F2: true}
                                    _ = v16
                                case 0:
                                    if (735 >= -1) {
                                        if (((("hello" + "") + ("a" + ("\n\t" + ""))) + "a") < "") {
                                            var v17 string = ""
                                            var v18 uint16 = 401
                                            _ = v17
                                            _ = v18
                                        } else {
                                            v19 := GenS2{F0: 1}
                                            v20 := []int{}
                                            v21 := GenS1{F0: 127, F1: 18446744073709551615, F2: false}
                                            _ = v19
                                            _ = v20
                                            _ = v21
                                        }
                                        v22 := GenS3{F0: (v12 - 0), F1: 127}
                                        if true {
                                            var v23 uint16 = 535
                                            if (false || (("a" + "foo bar") != "hello")) {
                                                var v24 float32 = float32(float64(-2.25))
                                                if true {
                                                    var v25 int16 = 42
                                                    if (true || (true || false)) {
                                                        v26 := GenS2{F0: 1}
                                                        v14 = v24
                                                        _ = v26
                                                    }
                                                    switch fi11 {
                                                    case 578:
                                                        v27 := []int{42}
                                                        _ = v27
                                                    default:
                                                        v12 -= genFn0(1, 2147483647)
                                                    }
                                                    _ = v25
                                                }
                                                fmt.Println(fi13, fi11, v14)
                                                v24 = v24
                                                var v28 bool = ((v12 != -1) || true)
                                                v29 := []int{(fi13 & (fi11 << ((963) & 63))), -287, 1, (fi11 << ((18446744073709551615) & 63))}
                                                _ = v24
                                                _ = v28
                                                _ = v29
                                            } else {
                                                switch fi11 {
                                                case 128:
                                                    v30 := GenS1{F0: 0, F1: 584, F2: false}
                                                    _ = v30
                                                }
                                            }
                                            var v31 float64 = float64(v14)
                                            var v32 int32 = -2147483648
                                            _ = v23
                                            _ = v31
                                            _ = v32
                                        }
                                        _ = v22
                                    } else {
                                        switch fi11 {
                                        case -681:
                                            v33 := GenS2{F0: -20}
                                            _ = v33
                                        case 31:
                                            v14 = (v14 - float32(0.0))
                                        }
                                    }
                                }
                                fmt.Println(v14, v14, fi11)
                                v34 := []int{(^(fi11)), fi13, fi13, (fi11 - 4294967295)}
                                _ = fi13
                                _ = v14
                                _ = v34
                            }
                        }
                        v35 := GenS1{F0: 1, F1: 0, F2: true}
                        _ = v35
                    }
                    var v36 uint64 = 2
                    _ = fi11
                    _ = v12
                    _ = v36
                }
            }
            var v37 int8 = 11
            _ = v37
        } else {
            v38 := []int{4294967295, 1}
            var v39 float64 = (1e10 + -2.25)
            for fi40 := 0; fi40 < 3; fi40++ {
                if true {
                    fmt.Println(v39)
                    v41 := GenS4{F0: 1, F1: 42, F2: -72}
                    _ = v41
                }
                _ = fi40
            }
            _ = v38
            _ = v39
        }
    }
    var v42 uint8 = 127
    var v43 uint64 = 0
    v44 := GenS5{F0: "hello", F1: "foo bar"}
    v45 := []int{1, 133, 2147483647, 127}
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    return 0
}

func main() {
    var v1 int = 32767
    var v2 uint64 = 2
    var v3 int64 = -1
    v4 := make(map[string]bool)
    var v5 float32 = 0.0
    v6 := []int{(^(v1)), v1, 42, v1}
    switch "a" {
    case "foo bar":
        if (("a" == "foo bar") || false) {
            v1 = v1
            v3 -= (v3 >> ((774) & 63))
        }
    case "hello":
        var v46 uint8 = 255
        _ = v46
    default:
        if (65535 != 1) {
            if false {
                var v47 bool = (464 > 32767)
                _ = v47
            } else {
                var v48 int64 = 42
                var v49 int = v1
                fmt.Println(v48, v48, v49)
                _ = v48
                _ = v49
            }
            var v50 int8 = 127
            _ = v50
        }
    }
    var v51 uint32 = 65535
    v4["a"] = true
    v5 = float32(0.0)
    v52 := v4[("foo bar" + ("foo bar" + ("\n\t" + "foo bar")))]
    if (!v52) {
        if v52 {
            fmt.Println(v51, v51)
            if true {
                v1 = v1
                var v53 bool = (v51 != 666)
                _ = v53
            } else {
                switch (("foo bar" + "\n\t") + "\n\t") {
                case "":
                    fmt.Println(v52)
                case "hello":
                    if len(v6) > 0 {
                        v6[0] = -9223372036854775808
                    }
                case "foo bar":
                    fmt.Println(v3)
                default:
                    switch v1 {
                    case 42:
                        fmt.Println(v52, v2)
                    case -805:
                        v2 -= uint64(v51)
                    default:
                        for fi54 := 0; fi54 < 4; fi54++ {
                            v55 := GenS5{F0: "", F1: "\n\t"}
                            _ = fi54
                            _ = v55
                        }
                    }
                }
                switch "a" {
                case "hello":
                    _ = len(v4)
                default:
                    switch 2147483647 {
                    case -947:
                        var v56 bool = (v52 || (0 < (v2 & 0)))
                        _ = v56
                    case 9223372036854775807:
                        if ((!v52) && v52) {
                            for fi57 := 0; fi57 < 15; fi57++ {
                                var v58 uint64 = 42
                                v59 := GenS5{F0: "", F1: (("a" + "foo bar") + "a")}
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                            if v52 {
                                var v60 uint64 = (v2 % (((v2 >> ((732) & 63))) | 1))
                                if v52 {
                                    var v61 int32 = -756
                                    _ = v61
                                }
                                v62 := GenS1{F0: 255, F1: v60, F2: (!v52)}
                                _ = v60
                                _ = v62
                            } else {
                                var v63 float64 = (0.0 * ((((float64(float32((3.14159 + -2.25))) + -2.25) - (3.14159 * 1e10)) * float64(v5)) + (1.5 - ((3.14159 + (0.0 - 0.1)) / -2.25))))
                                if len(v6) > 0 {
                                    v6[0] = (-(v1))
                                }
                                _ = v63
                            }
                        } else {
                            switch (v1 % (((v1 % ((v1) | 1))) | 1)) {
                            case -2147483648:
                                v51 = v51
                            }
                            for fi64 := 0; fi64 < 6; fi64++ {
                                switch "foo bar" {
                                case "hello":
                                    if v52 {
                                        fmt.Println(v3, v5, v1)
                                        fmt.Println(fi64)
                                    } else {
                                        v65 := GenS3{F0: -962, F1: -128}
                                        _ = v65
                                    }
                                }
                                v5 -= -2.25
                                _ = fi64
                            }
                            for fi66 := 0; fi66 < 15; fi66++ {
                                var v67 bool = true
                                _ = fi66
                                _ = v67
                            }
                        }
                    default:
                        var v68 int64 = (-(v3))
                        _ = v68
                    }
                }
            }
        } else {
            var v69 uint = 1
            switch "hello" {
            case "a":
                v6 = append(v6, (v1 ^ 979))
            case "hello":
                v2 = 961
            case "":
                var v70 int8 = 10
                _ = v70
            }
            var v71 int8 = 127
            _ = v69
            _ = v71
        }
    }
    var v72 int64 = 67
    if (v2 >= v2) {
        var v73 int = 616
        _ = v73
    }
    var v74 int64 = 911
    v75 := GenS3{F0: 545, F1: 42}
    var v76 bool = (false && (!v52))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v72:", v72)
    fmt.Println("v74:", v74)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
    _ = v72
    _ = v74
    _ = v75
    _ = v76
}
