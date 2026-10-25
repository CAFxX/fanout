package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int64; F1 bool; F2 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 bool }

func main() {
    var v1 int = 32767
    var v2 uint = 328
    var v3 int64 = -812
    var v4 uint = (^(v2))
    var v5 int32 = 2147483647
    var v6 float64 = ((float64(0.0) + -2.25) * 1.5)
    v7 := GenS1{F0: 0}
    if (false && (true || true)) {
        v8 := GenS2{F0: 42}
        _ = v8
    } else {
        v9 := GenS2{F0: 177}
        var v10 float32 = (3.14159 + 0.0)
        _ = v10
        _ = v9
    }
    switch 42 {
    case 128:
        v11 := GenS1{F0: 0}
        _ = v11
    default:
        v12 := []int{}
        _ = v12
    }
    if true {
        if (!false) {
            var v13 float64 = 3.14159
            for fi14 := 0; fi14 < 19; fi14++ {
                v15 := GenS3{F0: -128, F1: ("" + "a")}
                _ = fi14
                _ = v15
            }
            if false {
                v13 += (v13 + 1.5)
            } else {
                var v16 uint8 = 128
                v17 := []int{(^(v1)), (v1 / ((v1) | 1)), (-(v1))}
                v18 := GenS4{F0: v3, F1: ("\n\t" > "\n\t"), F2: 1}
                _ = v16
                _ = v17
                _ = v18
            }
            _ = v13
        } else {
            var v19 int64 = 24
            if (1 != 255) {
                switch (v1 >> (((v4 >> (((^(v4))) & 63))) & 63)) {
                case -161:
                    fmt.Println(v2)
                case 830:
                    switch "hello" {
                    case "foo bar":
                        v5 = v5
                    case "hello":
                        v20 := GenS5{F0: uint8(v2)}
                        _ = v20
                    case "":
                        if (false || ((v1 < 128) && true)) {
                            if ((true && true) || ((^(v3)) > 42)) {
                                v1 = (v1 - v1)
                                v21 := GenS3{F0: int8(v2), F1: ("foo bar" + ("" + ("hello" + "a")))}
                                v22 := GenS3{F0: -70, F1: "foo bar"}
                                _ = v21
                                _ = v22
                            }
                            var v23 int32 = v5
                            v24 := []int{v1, (v1 + v1)}
                            _ = v23
                            _ = v24
                        } else {
                            if false {
                                switch ("a" + "\n\t") {
                                case "foo bar":
                                    v4 -= v4
                                case "":
                                    v25 := []int{}
                                    _ = v25
                                default:
                                    var v26 float64 = v6
                                    _ = v26
                                }
                            }
                            for fi27 := 0; fi27 < 2; fi27++ {
                                var v28 int32 = (^(v5))
                                v19 -= 1
                                var v29 uint64 = 116
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            v30 := GenS1{F0: 1}
                            _ = v30
                        }
                    default:
                        for fi31 := 0; fi31 < 20; fi31++ {
                            v32 := []int{fi31}
                            fmt.Println(v19, v5)
                            _ = fi31
                            _ = v32
                        }
                    }
                case 741:
                    if (42 == (v5 ^ v5)) {
                        if ((42 >= v2) && (((583 < 194) && ((^(v5)) != v5)) && (true && ((v1 * v1) >= 32768)))) {
                            fmt.Println(v1, v4)
                        } else {
                            if false {
                                v33 := GenS3{F0: 42, F1: (("hello" + "\n\t") + "foo bar")}
                                fmt.Println(v2, v3, v6)
                                _ = v33
                            } else {
                                switch "a" {
                                case "hello":
                                    v5 = v5
                                case "":
                                    var v34 string = "\n\t"
                                    _ = v34
                                default:
                                    if (((1e10 < (0.0 / -2.25)) || (120 <= 1000)) && ("foo bar" == "hello")) {
                                        v35 := []int{}
                                        switch "" {
                                        case "hello":
                                            var v36 string = ("\n\t" + "foo bar")
                                            _ = v36
                                        case "\n\t":
                                            for fi37 := 0; fi37 < 17; fi37++ {
                                                if len(v35) > 0 {
                                                    v35[0] = fi37
                                                }
                                                if (128 == 2) {
                                                    v38 := GenS2{F0: 127}
                                                    _ = v38
                                                }
                                                var v39 int = (^(v1))
                                                _ = fi37
                                                _ = v39
                                            }
                                        case "foo bar":
                                            v1 += 42
                                        default:
                                            fmt.Println(v2, v4, v1)
                                        }
                                        v40 := GenS6{F0: 1}
                                        _ = v35
                                        _ = v40
                                    } else {
                                        fmt.Println(v6, v3)
                                        if true {
                                            switch 32768 {
                                            case 0:
                                                var v41 uint64 = 970
                                                _ = v41
                                            }
                                        }
                                    }
                                }
                                var v42 bool = ((((false && (!(false && true))) && true) && true) && false)
                                v43 := []int{}
                                _ = v42
                                _ = v43
                            }
                            if ((v1 | v1) == v1) {
                                if false {
                                    v44 := []int{v1, v1}
                                    var v45 int8 = -1
                                    for fi46 := 0; fi46 < 15; fi46++ {
                                        switch "a" {
                                        case "":
                                            var v47 string = "hello"
                                            _ = v47
                                        }
                                        _ = fi46
                                    }
                                    _ = v44
                                    _ = v45
                                }
                                var v48 uint16 = 42
                                v49 := []int{}
                                _ = v48
                                _ = v49
                            } else {
                                v50 := GenS7{F0: false}
                                v6 -= float64(float32((0.1 + v6)))
                                _ = v50
                            }
                        }
                        var v51 uint16 = 42
                        _ = v51
                    } else {
                        v19 = -493
                        v52 := GenS8{F0: 1}
                        if true {
                            v53 := []int{(v1 | -1), (v1 ^ (-(v1))), v1}
                            _ = v53
                        } else {
                            var v54 uint16 = 255
                            _ = v54
                        }
                        _ = v52
                    }
                }
                var v55 bool = ("a" >= "")
                _ = v55
            }
            _ = v19
        }
        switch v1 {
        case 32767:
            v6 -= float64(-2.25)
        }
        v1 -= (v1 ^ (v1 * -726))
    }
    v56 := []int{(^(v1))}
    var v57 int16 = 1000
    var v58 float64 = v6
    var v59 float32 = ((1e10 - (((0.1 - float32(0.0)) / 1.0) + 0.1)) / (-2.25 + (1e10 / 0.1)))
    switch -989 {
    case 2147483647:
        v5 -= (v5 >> ((v2) & 63))
    }
    if true {
        switch v1 {
        case 32767:
            v60 := GenS9{F0: true}
            _ = v60
        case 127:
            v57 = v57
        }
        for fi61 := 0; fi61 < 18; fi61++ {
            var v62 float64 = float64(v59)
            var v63 int32 = -527
            for fi64 := 0; fi64 < 9; fi64++ {
                fmt.Println(v5, v4, v63)
                _ = fi64
            }
            _ = fi61
            _ = v62
            _ = v63
        }
        var v65 float64 = v6
        _ = v65
    } else {
        _ = len(v56)
        v59 += v59
        _ = len(v56)
    }
    var v66 int16 = v57
    var v67 int32 = v5
    var v68 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v66
    _ = v67
    _ = v68
    _ = v7
}
