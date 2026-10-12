package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 string; F1 int8; F2 int8 }
type GenS8 struct { F0 int8; F1 uint64; F2 int16 }
func genFn0(p0 int8, p1 int32, p2 int64) uint32 {
    v6 := GenS1{F0: true}
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483648
}

func main() {
    var v1 float64 = (-2.25 - float64(float32(-2.25)))
    var v2 int64 = -173
    v3 := []string{(("a" + "hello") + ""), "", "foo bar"}
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    fmt.Println(v1, v1)
    var v7 float64 = float64((0.0 - 0.0))
    if false {
        var v8 float32 = 3.14159
        if true {
            v5["a"] = genFn0(-128, -957, (v2 * v2))
            v2 = (v2 * v2)
        }
        var v9 bool = ((!(false || true)) && (v7 >= v7))
        _ = v8
        _ = v9
    }
    var v10 float64 = float64(0.0)
    var v11 int16 = 627
    switch 4294967295 {
    case 4294967295:
        v7 -= 3.14159
    case 9223372036854775807:
        v10 -= ((v10 / (v1 / (-2.25 * v1))) / (1e10 + -2.25))
    case 42:
        if ((!true) && (4294967295 < 803)) {
            v7 = 1e10
            var v12 float64 = ((v7 + v10) - v10)
            var v13 uint16 = 2
            _ = v12
            _ = v13
        }
    }
    switch ("" + "foo bar") {
    case "foo bar":
        var v14 int8 = 100
        _ = v14
    case "a":
        fmt.Println(v11)
    }
    v11 = (v11 - v11)
    var v15 int8 = 100
    v16 := GenS1{F0: false}
    if false {
        v11 -= (v11 % (((-(v11))) | 1))
        for fi17 := 0; fi17 < 6; fi17++ {
            if (!true) {
                for fi18 := 0; fi18 < 9; fi18++ {
                    switch fi18 {
                    case -1:
                        if false {
                            v7 += v7
                            v19 := GenS1{F0: (!false)}
                            var v20 bool = true
                            _ = v19
                            _ = v20
                        } else {
                            switch "foo bar" {
                            case "a":
                                _ = len(v5)
                            case "hello":
                                switch "foo bar" {
                                case "":
                                    var v21 uint64 = 1
                                    _ = v21
                                case "foo bar":
                                    fmt.Println(fi17)
                                case "\n\t":
                                    v10 += ((v7 / v7) - 1e10)
                                }
                            case "\n\t":
                                if false {
                                    var v22 uint = 65535
                                    fmt.Println(v11)
                                    if (!(false || (255 > 127))) {
                                        fmt.Println(v2, v11)
                                        v23 := GenS1{F0: true}
                                        _ = v23
                                    }
                                    switch fi17 {
                                    case -9223372036854775808:
                                        switch "hello" {
                                        case "a":
                                            v2 -= (v2 | (v2 ^ (v2 & (v2 | v2))))
                                        }
                                        v2 -= 1
                                        v24 := GenS1{F0: (!(0.0 >= 1e10))}
                                        _ = v24
                                    case 32767:
                                        v25 := GenS2{F0: (fi18 * (fi18 / ((32768) | 1)))}
                                        _ = v25
                                    case -1:
                                        var v26 int64 = (v2 + (^(v2)))
                                        _ = v26
                                    default:
                                        v15 = (v15 - v15)
                                    }
                                    switch 2147483647 {
                                    case -264:
                                        var v27 uint = 2
                                        _ = v27
                                    case 4294967295:
                                        if true {
                                            var v28 uint = (v22 & (v22 | (v22 * v22)))
                                            v2 = v2
                                            if ((^(v28)) >= (v22 + 0)) {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v29 int64 = -223
                                                    _ = v29
                                                case "hello":
                                                    v30 := GenS3{F0: (v22 >> ((255) & 63))}
                                                    _ = v30
                                                case "foo bar":
                                                    v31 := GenS2{F0: fi17}
                                                    _ = v31
                                                default:
                                                    switch fi18 {
                                                    case -2147483648:
                                                        v32 := GenS1{F0: (v11 != v11)}
                                                        _ = v32
                                                    case 9223372036854775807:
                                                        if true {
                                                            if false {
                                                                if false {
                                                                    v33 := GenS4{F0: -2147483648, F1: 1}
                                                                    _ = v33
                                                                }
                                                                fmt.Println(v10, v7)
                                                            } else {
                                                                v28 = v28
                                                            }
                                                            fmt.Println(v15, v1)
                                                            if true {
                                                                delete(v4, 2147483647)
                                                                switch int(v11) {
                                                                case 1:
                                                                    v34 := GenS5{F0: (v15 % ((v15) | 1))}
                                                                    _ = v34
                                                                case 32768:
                                                                    if (!false) {
                                                                        var v35 uint16 = 32767
                                                                        var v36 int32 = 42
                                                                        v37 := GenS3{F0: v22}
                                                                        _ = v35
                                                                        _ = v36
                                                                        _ = v37
                                                                    } else {
                                                                        v38 := GenS4{F0: 658, F1: 319}
                                                                        var v39 int8 = int8(fi17)
                                                                        v39 += 1
                                                                        _ = v38
                                                                        _ = v39
                                                                    }
                                                                case -9223372036854775808:
                                                                    var v40 int16 = (v11 ^ -32768)
                                                                    _ = v40
                                                                default:
                                                                    v41 := GenS2{F0: fi18}
                                                                    _ = v41
                                                                }
                                                            } else {
                                                                v42 := GenS6{F0: (^(v11)), F1: "hello"}
                                                                v3 = append(v3, "")
                                                                var v43 float32 = 1.5
                                                                _ = v42
                                                                _ = v43
                                                            }
                                                        } else {
                                                            var v44 int16 = (v11 >> ((42) & 63))
                                                            fmt.Println(v11)
                                                            switch (fi17 + fi18) {
                                                            case 42:
                                                                if false {
                                                                    fmt.Println(v7, v2, v44)
                                                                } else {
                                                                    if (1 >= 271) {
                                                                        v3 = append(v3, "\n\t")
                                                                    }
                                                                    v11 += v44
                                                                }
                                                            }
                                                            v44 = -323
                                                            _ = v44
                                                        }
                                                    }
                                                }
                                            } else {
                                                var v45 int32 = 411
                                                var v46 bool = false
                                                v2 -= v2
                                                _ = v45
                                                _ = v46
                                            }
                                            v15 = v15
                                            v2 -= (-(v2))
                                            _ = v28
                                        }
                                    case 2147483647:
                                        v2 = (^(v2))
                                    }
                                    _ = v22
                                }
                            }
                            var v47 string = ("" + (("" + ("hello" + "hello")) + ("foo bar" + ("a" + "\n\t"))))
                            _ = v47
                        }
                    case 127:
                        delete(v5, "hello")
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = fi18
                }
            } else {
                var v48 int8 = (v15 + (v15 & 0))
                _ = v48
            }
            _ = len(v3)
            _ = fi17
        }
    } else {
        fmt.Println(v1, v2)
        v7 = v7
    }
    var v49 uint = 4294967295
    if true {
        fmt.Println(v49)
    } else {
        var v50 string = ("foo bar" + ("\n\t" + ""))
        _ = v50
    }
    if (!(!(!true))) {
        var v51 string = ("\n\t" + "hello")
        var v52 int8 = (v15 / (((-(v15))) | 1))
        for fi53 := 0; fi53 < 18; fi53++ {
            for fi54 := 0; fi54 < 8; fi54++ {
                v52 += v52
                v52 -= 39
                _ = fi54
            }
            _ = fi53
        }
        _ = v51
        _ = v52
    }
    v55 := GenS7{F0: "\n\t", F1: v15, F2: (v15 % ((1) | 1))}
    for fi56 := 0; fi56 < 6; fi56++ {
        delete(v5, "a")
        delete(v4, 0)
        _ = fi56
    }
    switch -2147483648 {
    case 2147483647:
        v49 = (v49 << (((^(v49))) & 63))
    case 293:
        fmt.Println(v15, v7)
    case 1:
        for fi57 := 0; fi57 < 4; fi57++ {
            v58 := GenS8{F0: v15, F1: 2, F2: (v11 % (((^(v11))) | 1))}
            for fi59 := 0; fi59 < 6; fi59++ {
                v2 -= (v2 % ((9223372036854775807) | 1))
                switch 780 {
                case 2147483647:
                    var v60 uint32 = 2
                    _ = v60
                }
                _ = fi59
            }
            if (!(fi57 < fi57)) {
                var v61 uint64 = 18446744073709551615
                v61 = (v61 + v61)
                fmt.Println(v15, v10, v2)
                _ = v61
            }
            _ = fi57
            _ = v58
        }
    default:
        _ = len(v3)
    }
    for fi62 := 0; fi62 < 4; fi62++ {
        v63 := GenS6{F0: (v11 - v11), F1: "a"}
        _ = len(v4)
        for fi64 := 0; fi64 < 17; fi64++ {
            v49 = 1
            _ = fi64
        }
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v55
    _ = v7
}
