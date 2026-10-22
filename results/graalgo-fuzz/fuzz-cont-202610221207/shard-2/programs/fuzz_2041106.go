package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int64; F1 int; F2 int16 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int16) int {
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: p0}
        v8 := []int{}
        var v9 float64 = (0.0 - -2.25)
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    p0 = 42
    _ = p0
    return 567
}
func genFn1(p0 int8, p1 int32, p2 uint16) int64 {
    var v10 uint64 = 18446744073709551615
    var v11 int8 = p0
    fmt.Println(v10, v11, v11)
    if (float32(((1.5 * 0.0) / 1.0)) != 0.1) {
        switch ("a" + ("" + "")) {
        case "":
            var v12 int64 = 42
            _ = v12
        }
    } else {
        v13 := GenS2{F0: (!(v11 == (p0 >> ((2) & 63)))), F1: 42, F2: 0}
        fmt.Println(p2)
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return -9223372036854775808
}

func main() {
    var v1 float32 = ((-2.25 + ((-2.25 / 1e10) * 3.14159)) + 0.0)
    var v2 int = -9223372036854775808
    v3 := make(map[int]int)
    v4 := []string{(("\n\t" + ("a" + "\n\t")) + (("\n\t" + ("a" + (("foo bar" + ("\n\t" + "")) + ""))) + "a")), "hello"}
    var v5 string = ("foo bar" + "a")
    v4 = append(v4, "\n\t")
    if (v5 == "foo bar") {
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 uint8 = 0
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    _ = len(v3)
    fmt.Println(v5, v2)
    if (!true) {
        _ = len(v3)
    } else {
        fmt.Println(v2)
    }
    v2 -= (-(v2))
    var v17 int8 = 0
    _ = len(v3)
    var v18 uint = 631
    if ((v1 + v1) <= v1) {
        var v19 float32 = float32(-2.25)
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 int = (fi20 / ((-585) | 1))
            for fi22 := 0; fi22 < 17; fi22++ {
                fmt.Println(v5, v1)
                switch genFn0(42) {
                case -9223372036854775808:
                    v21 -= 964
                case -483:
                    var v23 bool = false
                    v21 = 42
                    switch ("a" + "foo bar") {
                    case "\n\t":
                        if len(v4) > 0 {
                            v4[0] = ("hello" + v5)
                        }
                    case "hello":
                        if v23 {
                            var v24 uint = v18
                            _ = v24
                        } else {
                            var v25 uint16 = 239
                            fmt.Println(v19, v1, v18)
                            _ = v25
                        }
                    case "a":
                        fmt.Println(v17, fi20)
                    }
                    _ = v23
                default:
                    v19 -= (((v19 * (0.0 - (v1 / -2.25))) + v1) / (v19 + 1.5))
                    _ = len(v3)
                }
                _ = fi22
            }
            var v26 float32 = -2.25
            _ = fi20
            _ = v21
            _ = v26
        }
        _ = v19
    } else {
        v27 := GenS3{F0: 42, F1: 0}
        if false {
            for fi28 := 0; fi28 < 7; fi28++ {
                fmt.Println(v5, v2, v1)
                switch "foo bar" {
                case "":
                    for fi29 := 0; fi29 < 5; fi29++ {
                        if false {
                            v30 := GenS3{F0: 32767, F1: 4294967295}
                            v4 = append(v4, v5)
                            v1 = (float32(1.5) * v1)
                            _ = v30
                        } else {
                            if ((0.1 + ((1.5 - (1e10 + 3.14159)) - 0.1)) > 3.14159) {
                                v31 := GenS2{F0: true, F1: 128, F2: 127}
                                var v32 int8 = v17
                                if (true && true) {
                                    if true {
                                        v5 = v5
                                    } else {
                                        fmt.Println(v5, v18)
                                        var v33 bool = (true && false)
                                        if v33 {
                                            switch (fi29 - fi29) {
                                            case 373:
                                                switch (fi29 << ((v18) & 63)) {
                                                case 1:
                                                    v32 -= (v32 + (v17 ^ v17))
                                                case 290:
                                                    var v34 uint8 = 2
                                                    _ = v34
                                                default:
                                                    if len(v4) > 0 {
                                                        v4[0] = "hello"
                                                    }
                                                }
                                            }
                                            v5 = v5
                                            v35 := GenS4{F0: 9223372036854775807, F1: (fi28 ^ 988), F2: -32768}
                                            if (-859 <= 2147483647) {
                                                fmt.Println(v17)
                                            }
                                            _ = v35
                                        } else {
                                            v3[fi29] = int(v32)
                                            delete(v3, genFn0(0))
                                            v1 = v1
                                        }
                                        _ = v33
                                    }
                                    var v36 uint64 = 12
                                    var v37 bool = false
                                    _ = v36
                                    _ = v37
                                }
                                var v38 int32 = 42
                                _ = v31
                                _ = v32
                                _ = v38
                            } else {
                                if (128 > 42) {
                                    if ((551 != 46) && (true && false)) {
                                        v2 = v2
                                        var v39 int64 = 1
                                        _ = v39
                                    } else {
                                        v40 := GenS5{F0: -9223372036854775808}
                                        delete(v3, fi29)
                                        if false {
                                            var v41 int16 = 32767
                                            switch -1 {
                                            case 431:
                                                v42 := GenS4{F0: 154, F1: genFn0(-1), F2: v41}
                                                _ = v42
                                            case -9223372036854775808:
                                                fmt.Println(v2, v5)
                                            case -2147483648:
                                                var v43 uint8 = 114
                                                _ = v43
                                            default:
                                                var v44 uint8 = 2
                                                _ = v44
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = "a"
                                            }
                                            _ = v41
                                        } else {
                                            var v45 uint16 = 1
                                            if false {
                                                switch v5 {
                                                case "hello":
                                                    v45 += (v45 / ((v45) | 1))
                                                case "\n\t":
                                                    if false {
                                                        v18 = 466
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    } else {
                                                        v4 = append(v4, (v5 + "a"))
                                                        v2 = fi29
                                                        var v46 uint32 = 2147483647
                                                        _ = v46
                                                    }
                                                case "a":
                                                    v18 -= 255
                                                default:
                                                    v17 += v17
                                                    if false {
                                                        fmt.Println(fi28, v45)
                                                    } else {
                                                        v45 = 554
                                                        if false {
                                                            v2 = (fi29 / ((-30) | 1))
                                                        }
                                                    }
                                                }
                                            } else {
                                                v47 := GenS5{F0: int64(v2)}
                                                v1 = 0.0
                                                var v48 uint16 = (v45 + v45)
                                                _ = v47
                                                _ = v48
                                            }
                                            v4 = append(v4, "foo bar")
                                            var v49 uint16 = 42
                                            _ = v45
                                            _ = v49
                                        }
                                        var v50 uint32 = 490
                                        var v51 string = "hello"
                                        _ = v40
                                        _ = v50
                                        _ = v51
                                    }
                                    delete(v3, (-(fi28)))
                                    if true {
                                        v52 := v3[genFn0(32767)]
                                        fmt.Println(fi29)
                                        var v53 int32 = -1
                                        _ = v52
                                        _ = v53
                                    }
                                }
                                var v54 int64 = 0
                                if true {
                                    _ = len(v3)
                                    var v55 uint64 = 0
                                    var v56 uint8 = 157
                                    if true {
                                        var v57 bool = true
                                        v58 := GenS5{F0: (v54 - genFn1(v17, -598, 1))}
                                        v4 = append(v4, v5)
                                        _ = v57
                                        _ = v58
                                    } else {
                                        v1 = float32(1.5)
                                        switch (fi28 >> ((525) & 63)) {
                                        case 1:
                                            if (!((^(v2)) == (^(fi28)))) {
                                                v18 = v18
                                                v4 = append(v4, (v5 + "hello"))
                                            }
                                        }
                                    }
                                    _ = v55
                                    _ = v56
                                } else {
                                    switch ((("hello" + v5) + "a") + v5) {
                                    case "\n\t":
                                        v54 -= 1
                                    }
                                    v18 += v18
                                }
                                _ = v54
                            }
                        }
                        var v59 int16 = 42
                        fmt.Println(v1, v1, v5)
                        _ = fi29
                        _ = v59
                    }
                }
                _ = fi28
            }
        }
        v2 += (v2 - v2)
        _ = v27
    }
    fmt.Println(v18, v17, v1)
    switch (v2 % ((genFn0(32767)) | 1)) {
    case 42:
        switch "foo bar" {
        case "":
            for fi60 := 0; fi60 < 1; fi60++ {
                var v61 int16 = 42
                _ = fi60
                _ = v61
            }
        }
    case 32767:
        var v62 int16 = -1
        _ = v62
    case 0:
        fmt.Println(v17)
    default:
        for fi63 := 0; fi63 < 4; fi63++ {
            if true {
                fmt.Println(fi63, v5)
            } else {
                fmt.Println(v18)
                v1 = v1
                if (-2.25 < 0.1) {
                    var v64 int = genFn0(1)
                    _ = v64
                } else {
                    var v65 float32 = float32(1.5)
                    var v66 int32 = 1
                    var v67 float64 = 0.1
                    _ = v65
                    _ = v66
                    _ = v67
                }
            }
            if false {
                for fi68 := 0; fi68 < 16; fi68++ {
                    v69 := GenS1{F0: 1000}
                    v70 := GenS3{F0: 255, F1: 2}
                    _ = fi68
                    _ = v69
                    _ = v70
                }
            }
            var v71 float32 = ((0.1 * v1) - 0.0)
            _ = fi63
            _ = v71
        }
    }
    switch "hello" {
    case "a":
        v17 = (^(v17))
    case "foo bar":
        fmt.Println(v17, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
