package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 bool; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint }
func genFn0() string {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 string = ("foo bar" + "foo bar")
        fmt.Println(v8, v8)
        _ = fi7
        _ = v8
    }
    switch -230 {
    case -56:
        v9 := GenS1{F0: 639, F1: 100, F2: 736}
        _ = v9
    }
    if (678 <= 1) {
        v10 := []int{-983}
        switch -329 {
        case -1:
            var v11 int = 128
            _ = v11
        default:
            var v12 uint64 = 495
            _ = v12
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := GenS1{F0: (fi13 / ((9223372036854775807) | 1)), F1: 100, F2: 4294967295}
            if false {
                var v15 uint8 = 255
                _ = v15
            }
            _ = len(v10)
            _ = fi13
            _ = v14
        }
        _ = v10
    } else {
        v16 := GenS1{F0: -742, F1: 1, F2: 499}
        var v17 int = 816
        fmt.Println(v17)
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 42, F1: -1, F2: 938}
    _ = v18
    return (("a" + "\n\t") + ("" + ("" + "")))
}
func genFn1() int64 {
    v19 := []int{2147483647, 128, -1, -744}
    var v20 int16 = 1000
    if (65535 <= 255) {
        v20 = (v20 - v20)
        switch "" {
        case "":
            if (!(42 == -288)) {
                if ((1.5 + 3.14159) != -2.25) {
                    fmt.Println(v20)
                    for fi21 := 0; fi21 < 20; fi21++ {
                        switch fi21 {
                        case 32768:
                            if true {
                                var v22 float64 = (3.14159 + -2.25)
                                for fi23 := 0; fi23 < 1; fi23++ {
                                    var v24 uint8 = 127
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    v25 := GenS1{F0: (fi21 | fi21), F1: -1, F2: 42}
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                                v26 := GenS2{F0: 543, F1: "foo bar"}
                                _ = v22
                                _ = v26
                            }
                        case 9223372036854775807:
                            fmt.Println(v20, fi21)
                        case 4294967295:
                            var v27 uint32 = 2147483648
                            _ = v27
                        default:
                            if true {
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                            } else {
                                if true {
                                    var v28 string = "\n\t"
                                    _ = v28
                                } else {
                                    var v29 string = "foo bar"
                                    var v30 uint16 = 493
                                    _ = v29
                                    _ = v30
                                }
                                switch (fi21 << ((973) & 63)) {
                                case -1:
                                    for fi31 := 0; fi31 < 13; fi31++ {
                                        var v32 int = fi31
                                        _ = fi31
                                        _ = v32
                                    }
                                default:
                                    _ = len(v19)
                                }
                                var v33 bool = (255 > 1)
                                _ = v33
                            }
                        }
                        for fi34 := 0; fi34 < 16; fi34++ {
                            v35 := GenS3{F0: false, F1: 291, F2: -1}
                            var v36 string = (((("hello" + (("\n\t" + "") + ("a" + ("\n\t" + "foo bar")))) + "\n\t") + "\n\t") + "")
                            var v37 string = v36
                            switch ("\n\t" + "") {
                            case "a":
                                switch (-(fi21)) {
                                case 289:
                                    var v38 string = "hello"
                                    _ = v38
                                case -645:
                                    fmt.Println(v20)
                                case 2147483647:
                                    var v39 uint32 = 747
                                    _ = v39
                                default:
                                    if ((v37 + v37) != (((v37 + (v37 + "")) + v37) + (v37 + ""))) {
                                        _ = len(v19)
                                        var v40 string = ("hello" + "a")
                                        if true {
                                            v36 = (v40 + ("foo bar" + (v40 + ((((((((("foo bar" + (((v36 + (v40 + v40)) + "foo bar") + "a")) + ((("a" + v40) + (("hello" + ((v36 + (v40 + "\n\t")) + v40)) + (v37 + v40))) + ("\n\t" + ((("\n\t" + v36) + v36) + "a")))) + v37) + (v36 + (("foo bar" + v36) + "foo bar"))) + (v36 + v37)) + v37) + v37) + "a") + "\n\t"))))
                                        }
                                        _ = v40
                                    }
                                }
                            }
                            var v41 string = (v36 + (("a" + "foo bar") + ""))
                            _ = fi34
                            _ = v35
                            _ = v36
                            _ = v37
                            _ = v41
                        }
                        _ = fi21
                    }
                    for fi42 := 0; fi42 < 3; fi42++ {
                        v20 = -32768
                        if len(v19) > 0 {
                            v19[0] = (fi42 | fi42)
                        }
                        fmt.Println(fi42, v20)
                        _ = fi42
                    }
                }
            }
        case "a":
            fmt.Println(v20)
        }
    }
    v19 = append(v19, 32767)
    _ = v19
    _ = v20
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int8 = -1
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    var v5 float64 = 1e10
    v6 := make(map[int]bool)
    var v43 bool = true
    v43 = ("hello" > (("hello" + "foo bar") + "foo bar"))
    for fi44 := 0; fi44 < 19; fi44++ {
        switch ("foo bar" + ("" + "hello")) {
        case "foo bar":
            if true {
                v43 = (285 == 0)
                v2 = (^(v2))
            }
        default:
            v45 := GenS3{F0: (v43 && ((-2147483648 == 42) || v43)), F1: (v1 ^ (^(v1))), F2: 1}
            _ = v45
        }
        for fi46 := 0; fi46 < 11; fi46++ {
            var v47 uint64 = v1
            _ = fi46
            _ = v47
        }
        _ = fi44
    }
    v43 = false
    v1 -= 230
    for fi48 := 0; fi48 < 10; fi48++ {
        var v49 bool = v43
        v1 = 4294967295
        var v50 int32 = 724
        _ = fi48
        _ = v49
        _ = v50
    }
    delete(v4, 9223372036854775807)
    v51 := GenS4{F0: 192}
    v2 -= (v2 - 0)
    for fi52 := 0; fi52 < 20; fi52++ {
        v53 := GenS2{F0: 536, F1: "\n\t"}
        _ = fi52
        _ = v53
    }
    v54 := []int{-2147483648, -243, 127, 2147483647}
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v51
    _ = v54
    _ = v6
}
