package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 string; F2 int }
type GenS4 struct { F0 int; F1 uint16; F2 int }
type GenS5 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint32, p1 int32, p2 string) uint8 {
    var v4 uint = 2
    v5 := []int{315, -812, 690}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2
}
func genFn1() int32 {
    var v6 int8 = 42
    v7 := []int{42, -297}
    var v8 string = "foo bar"
    var v9 float64 = (1.5 * -2.25)
    v10 := GenS1{F0: genFn0(65535, 0, v8)}
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    for fi11 := 0; fi11 < 7; fi11++ {
        switch fi11 {
        case -884:
            var v12 uint8 = 0
            _ = v12
        case -695:
            if true {
                if ("foo bar" < "a") {
                    switch 32767 {
                    case -1:
                        if (0.0 == 3.14159) {
                            v13 := []int{}
                            _ = v13
                        }
                    case 764:
                        if true {
                            var v14 float64 = ((-2.25 / 1.5) * 1.5)
                            switch "\n\t" {
                            case "\n\t":
                                var v15 int32 = -2147483648
                                _ = v15
                            case "foo bar":
                                var v16 uint32 = 930
                                _ = v16
                            default:
                                var v17 uint16 = 255
                                _ = v17
                            }
                            delete(v2, 836)
                            _ = v14
                        } else {
                            if (false && (true && false)) {
                                v18 := []int{(fi11 & (fi11 >> ((552) & 63)))}
                                _ = v18
                            } else {
                                var v19 uint8 = 82
                                v19 -= (v19 / ((genFn0(0, -1, (("" + "") + "\n\t"))) | 1))
                                for fi20 := 0; fi20 < 14; fi20++ {
                                    v21 := []int{-2147483648, (fi11 | fi20), -2}
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    var v22 int8 = 37
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                                _ = v19
                            }
                            switch "foo bar" {
                            case "":
                                fmt.Println(fi11)
                            case "\n\t":
                                var v23 float64 = (float64(0.1) / 3.14159)
                                _ = v23
                            default:
                                for fi24 := 0; fi24 < 6; fi24++ {
                                    fmt.Println(fi11)
                                    _ = len(v1)
                                    var v25 string = "a"
                                    fmt.Println(fi24)
                                    var v26 int = (fi11 % (((fi11 & 42)) | 1))
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                            }
                            if (false && true) {
                                var v27 int64 = 578
                                switch (fi11 >> ((0) & 63)) {
                                case -2147483648:
                                    v27 = (-(v27))
                                case -9223372036854775808:
                                    for fi28 := 0; fi28 < 5; fi28++ {
                                        v27 += (v27 | (-(v27)))
                                        delete(v1, 2)
                                        _ = fi28
                                    }
                                }
                                for fi29 := 0; fi29 < 18; fi29++ {
                                    var v30 int8 = 100
                                    v30 -= (v30 << ((0) & 63))
                                    v31 := []int{(fi29 - 9223372036854775807)}
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                                _ = v27
                            }
                        }
                    default:
                        var v32 int = -2147483648
                        _ = v32
                    }
                } else {
                    var v33 bool = false
                    for fi34 := 0; fi34 < 7; fi34++ {
                        v35 := []int{(fi34 ^ (fi11 - (fi34 * fi11))), fi11, fi11}
                        switch 129 {
                        case 128:
                            v36 := v3[2]
                            fmt.Println(fi11)
                            switch "a" {
                            case "a":
                                _ = len(v3)
                            case "\n\t":
                                v37 := GenS1{F0: 136}
                                _ = v37
                            }
                            _ = v36
                        case -636:
                            v38 := GenS2{F0: 482}
                            _ = v38
                        case -530:
                            if ((1e10 / 3.14159) >= (3.14159 + (1e10 / (3.14159 * ((1e10 / (0.0 + (float64(-2.25) * 0.1))) - 3.14159))))) {
                                v39 := GenS1{F0: 21}
                                var v40 uint16 = 32767
                                v41 := GenS1{F0: 2}
                                _ = v39
                                _ = v40
                                _ = v41
                            } else {
                                v42 := GenS1{F0: genFn0(2147483648, 1, ("hello" + (("" + ("\n\t" + "")) + "a")))}
                                if (4294967295 > 4294967295) {
                                    var v43 uint16 = 530
                                    var v44 int = (^(fi34))
                                    _ = v43
                                    _ = v44
                                }
                                _ = v42
                            }
                        default:
                            fmt.Println(v33, fi34)
                        }
                        v33 = v33
                        _ = fi34
                        _ = v35
                    }
                    _ = v33
                }
            }
        case 128:
            var v45 float32 = ((float32((1.5 - (1e10 / 1.0))) / 1.0) + (-2.25 - (0.1 - float32(-2.25))))
            _ = v45
        }
        for fi46 := 0; fi46 < 18; fi46++ {
            v47 := []int{(fi11 ^ (^(fi46))), fi11, (fi46 / ((-9223372036854775808) | 1)), fi11}
            _ = fi46
            _ = v47
        }
        _ = fi11
    }
    var v48 uint8 = 1
    v3[0] = ("\n\t" + ((((("\n\t" + "hello") + ("foo bar" + "foo bar")) + (((("hello" + "\n\t") + "foo bar") + (("hello" + "a") + "hello")) + ("" + "a"))) + (("hello" + "\n\t") + ("" + (("foo bar" + "a") + "hello")))) + ""))
    if (699 >= -789) {
        if (!((true || (!true)) || true)) {
            v48 += (v48 - (v48 & (v48 / (((v48 | (v48 * 255))) | 1))))
            fmt.Println(v48)
        } else {
            v49 := GenS1{F0: (^(v48))}
            v50 := GenS2{F0: 349}
            for fi51 := 0; fi51 < 15; fi51++ {
                var v52 uint = 4294967295
                _ = fi51
                _ = v52
            }
            _ = v49
            _ = v50
        }
    } else {
        if false {
            for fi53 := 0; fi53 < 1; fi53++ {
                var v54 float32 = 3.14159
                var v55 int = fi53
                v56 := GenS1{F0: v48}
                _ = fi53
                _ = v54
                _ = v55
                _ = v56
            }
            fmt.Println(v48)
        } else {
            var v57 uint8 = (v48 + v48)
            v57 -= genFn0(1, int32(v48), "hello")
            var v58 bool = (4294967295 > 0)
            _ = v57
            _ = v58
        }
        v48 = v48
    }
    var v59 uint16 = 0
    v3[0] = ""
    fmt.Println(v48, v59)
    var v60 string = ("\n\t" + ("hello" + "foo bar"))
    switch "foo bar" {
    case "\n\t":
        v61 := GenS3{F0: 2147483647, F1: v60, F2: 9223372036854775807}
        _ = v61
    case "foo bar":
        v60 = (v60 + v60)
    }
    v48 = v48
    if (1 >= -128) {
        fmt.Println(v59, v48)
    } else {
        v1[2147483648] = true
    }
    v48 = v48
    v62 := []int{551, -33, 32767, 9223372036854775807}
    if true {
        for fi63 := 0; fi63 < 2; fi63++ {
            if true {
                if false {
                    v48 = (v48 - (v48 | 1))
                }
            } else {
                for fi64 := 0; fi64 < 13; fi64++ {
                    v48 = (v48 ^ 128)
                    _ = fi64
                }
                var v65 uint64 = 4294967295
                var v66 bool = (3.14159 < 3.14159)
                _ = v65
                _ = v66
            }
            _ = len(v2)
            _ = fi63
        }
        delete(v1, 2147483647)
        fmt.Println(v59)
    } else {
        v60 = ((v60 + v60) + (v60 + "hello"))
        if ((!(true && false)) || false) {
            fmt.Println(v48, v48, v59)
            var v67 uint8 = v48
            fmt.Println(v67)
            _ = v67
        } else {
            switch 0 {
            case 441:
                var v68 bool = false
                _ = v68
            case 128:
                if len(v62) > 0 {
                    v62[0] = -9223372036854775808
                }
            case -536:
                v69 := GenS4{F0: 212, F1: (^(v59)), F2: -471}
                _ = v69
            }
        }
    }
    v70 := GenS3{F0: 4294967295, F1: "", F2: -9223372036854775808}
    v71 := GenS5{F0: -9223372036854775808, F1: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v48:", v48)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v59
    _ = v60
    _ = v62
    _ = v70
    _ = v71
}
