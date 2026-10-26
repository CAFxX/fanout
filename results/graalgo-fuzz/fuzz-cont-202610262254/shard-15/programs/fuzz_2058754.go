package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int; F1 int; F2 int64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS8 struct { F0 int32; F1 uint8 }

func main() {
    var v1 uint32 = 2147483647
    var v2 int16 = -32768
    v3 := make(map[uint32]int)
    var v4 int16 = (-(v2))
    var v5 int = 32767
    v6 := GenS1{F0: -2147483648, F1: 698, F2: true}
    switch (v5 ^ -9223372036854775808) {
    case 0:
        var v7 uint32 = 42
        _ = v7
    case -2147483648:
        if false {
            if true {
                var v8 float64 = ((-2.25 - 1e10) / 1.0)
                _ = v8
            } else {
                switch "a" {
                case "hello":
                    v1 = 975
                case "\n\t":
                    var v9 float64 = (1.5 * (1e10 * (0.0 / (3.14159 + ((0.1 * (((1e10 + 1e10) / 1.5) + 3.14159)) / 3.14159)))))
                    _ = v9
                case "foo bar":
                    if (false || (-125 > 42)) {
                        v4 = int16(v1)
                        fmt.Println(v1, v5, v4)
                        for fi10 := 0; fi10 < 6; fi10++ {
                            delete(v3, (v1 % ((v1) | 1)))
                            if (0.1 >= 0.0) {
                                v11 := GenS2{F0: (v1 % (((v1 + 1)) | 1))}
                                if ((!(true || true)) && (0 >= -128)) {
                                    v1 += (v1 * (^(v1)))
                                    if false {
                                        if true {
                                            switch "a" {
                                            case "":
                                                v12 := []int{(v5 >> ((255) & 63)), (^(fi10)), v5, (^(v5))}
                                                _ = v12
                                            }
                                            switch -1 {
                                            case 4294967295:
                                                v13 := GenS1{F0: 0, F1: 2, F2: ((("foo bar" + "hello") + "hello") >= "")}
                                                _ = v13
                                            default:
                                                delete(v3, v1)
                                            }
                                            v14 := GenS3{F0: "foo bar", F1: 557}
                                            _ = v14
                                        }
                                    }
                                    v15 := GenS1{F0: -1, F1: 255, F2: false}
                                    _ = v15
                                } else {
                                    v16 := GenS4{F0: (false || true), F1: -156}
                                    var v17 int64 = -854
                                    _ = v16
                                    _ = v17
                                }
                                var v18 float32 = float32((0.1 * 0.0))
                                _ = v11
                                _ = v18
                            }
                            _ = fi10
                        }
                    }
                }
            }
        } else {
            for fi19 := 0; fi19 < 9; fi19++ {
                for fi20 := 0; fi20 < 3; fi20++ {
                    switch ("foo bar" + "") {
                    case "\n\t":
                        var v21 string = "a"
                        _ = v21
                    case "hello":
                        switch "" {
                        case "foo bar":
                            var v22 uint16 = 255
                            _ = v22
                        default:
                            fmt.Println(v4)
                            v23 := v3[v1]
                            v23 = fi19
                            _ = v23
                        }
                    case "":
                        if true {
                            v24 := GenS5{F0: 2147483647, F1: v5, F2: 9223372036854775807}
                            if false {
                                var v25 float64 = 0.0
                                _ = v25
                            } else {
                                _ = len(v3)
                                switch 4294967295 {
                                case 946:
                                    switch fi19 {
                                    case -2147483648:
                                        switch "\n\t" {
                                        case "":
                                            _ = len(v3)
                                        case "hello":
                                            fmt.Println(fi20, fi20, fi19)
                                        case "foo bar":
                                            v26 := v3[(v1 & v1)]
                                            _ = v26
                                        }
                                    case 42:
                                        v2 = v2
                                    case 862:
                                        v2 -= v2
                                    default:
                                        var v27 uint16 = 2
                                        _ = v27
                                    }
                                }
                            }
                            v28 := GenS3{F0: "", F1: 42}
                            _ = v24
                            _ = v28
                        } else {
                            v29 := v3[v1]
                            v4 = v4
                            _ = v29
                        }
                        if (false || false) {
                            v5 = fi19
                            v30 := GenS6{F0: "\n\t"}
                            v1 = (v1 << ((0) & 63))
                            _ = v30
                        } else {
                            var v31 uint16 = 2
                            var v32 uint32 = v1
                            _ = v31
                            _ = v32
                        }
                        var v33 float64 = (-2.25 - 0.1)
                        _ = v33
                    default:
                        v2 += v2
                    }
                    if false {
                        v2 = 32767
                        if true {
                            v4 += -958
                            fmt.Println(v5, fi20, v2)
                        } else {
                            if (fi20 > fi20) {
                                fmt.Println(v2, v1)
                                v5 -= (v5 - fi20)
                                var v34 uint8 = 127
                                _ = v34
                            }
                            v35 := GenS3{F0: ("\n\t" + ("foo bar" + (("hello" + "hello") + "\n\t"))), F1: 2}
                            var v36 uint32 = (v1 & v1)
                            _ = v35
                            _ = v36
                        }
                        v37 := GenS7{F0: v4, F1: 83, F2: 170}
                        fmt.Println(fi20)
                        if false {
                            var v38 int8 = 1
                            v39 := []int{fi19, 32768}
                            var v40 bool = (false || true)
                            _ = v38
                            _ = v39
                            _ = v40
                        } else {
                            var v41 int64 = 0
                            _ = v41
                        }
                        _ = v37
                    }
                    _ = fi20
                }
                _ = fi19
            }
            if true {
                fmt.Println(v4)
                switch "foo bar" {
                case "":
                    v2 += (v4 + -593)
                case "hello":
                    for fi42 := 0; fi42 < 9; fi42++ {
                        var v43 uint16 = 1
                        var v44 string = "\n\t"
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                case "\n\t":
                    v5 += (v5 % ((0) | 1))
                }
            } else {
                v2 = (v4 ^ v2)
            }
        }
    case -816:
        var v45 int = (^(v5))
        _ = v45
    default:
        if (!(!true)) {
            switch v5 {
            case 119:
                v46 := GenS2{F0: (^(v1))}
                _ = v46
            case 753:
                if (!true) {
                    v5 = (v5 / ((128) | 1))
                } else {
                    v47 := v3[v1]
                    var v48 bool = ((1 == 638) || ((false || (false || ((true && false) || false))) && true))
                    _ = v47
                    _ = v48
                }
            }
            v2 += -32768
        }
    }
    fmt.Println(v1, v1, v2)
    v49 := GenS7{F0: (v2 + (-(v4))), F1: 128, F2: 128}
    if (false || false) {
        v50 := GenS3{F0: ((("hello" + ("" + ((("\n\t" + "hello") + ("" + "a")) + "\n\t"))) + "\n\t") + "foo bar"), F1: 306}
        v1 = v1
        fmt.Println(v5, v2)
        _ = v50
    }
    var v51 int64 = 9223372036854775807
    v51 -= (v51 & (v51 | v51))
    for fi52 := 0; fi52 < 17; fi52++ {
        for fi53 := 0; fi53 < 19; fi53++ {
            var v54 float32 = 0.0
            if (879 > 1) {
                v54 = (v54 + 0.1)
                var v55 uint32 = v1
                v4 = 0
                _ = v55
            } else {
                var v56 float64 = float64(v54)
                if (!(670 >= 4294967295)) {
                    v2 += (v4 << ((42) & 63))
                    var v57 int16 = v2
                    _ = v57
                }
                if true {
                    v2 -= -252
                    if (!false) {
                        var v58 string = ((("\n\t" + (((("a" + "a") + "") + ("a" + ("\n\t" + ("a" + "")))) + "")) + "a") + ((("foo bar" + "foo bar") + "a") + "hello"))
                        fmt.Println(v51, v54, v1)
                        v59 := []int{(fi52 + -925)}
                        _ = v58
                        _ = v59
                    }
                }
                _ = v56
            }
            if true {
                var v60 uint8 = 255
                delete(v3, (^(v1)))
                var v61 uint8 = v60
                _ = v60
                _ = v61
            } else {
                v62 := []int{v5}
                v3[(v1 >> ((1) & 63))] = (fi52 << ((18446744073709551615) & 63))
                var v63 int = fi52
                _ = v62
                _ = v63
            }
            _ = fi53
            _ = v54
        }
        for fi64 := 0; fi64 < 2; fi64++ {
            if (!(65535 <= 594)) {
                v4 = v2
            }
            if false {
                v65 := []int{}
                v5 = fi52
                v66 := GenS8{F0: 2147483647, F1: 0}
                _ = v65
                _ = v66
            }
            _ = fi64
        }
        fmt.Println(v5, fi52, v2)
        _ = fi52
    }
    fmt.Println(v1, v4, v5)
    fmt.Println(v5, v2, v1)
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        v2 -= (v4 >> ((18446744073709551615) & 63))
    case "foo bar":
        v5 = v5
    case "hello":
        var v67 bool = false
        _ = v67
    default:
        v68 := GenS1{F0: 0, F1: 448, F2: (28 <= 255)}
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v51
    _ = v6
}
