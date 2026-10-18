package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 int8 }

func main() {
    var v1 int16 = 32767
    var v2 int8 = 45
    var v3 uint8 = 2
    var v4 uint = 255
    v5 := make(map[string]string)
    v3 = (^(v3))
    _ = len(v5)
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{fi6}
        var v8 uint32 = 791
        for fi9 := 0; fi9 < 16; fi9++ {
            v8 = (v8 % ((v8) | 1))
            _ = fi9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v10 int16 = (v1 ^ v1)
    switch 42 {
    case 4294967295:
        v11 := v5[(("a" + "a") + "")]
        _ = v11
    case -473:
        v4 = (v4 * 255)
    }
    if (!false) {
        fmt.Println(v10)
    }
    v1 = 1000
    var v12 int = 44
    if true {
        var v13 uint8 = 128
        if (true || false) {
            if true {
                v14 := []int{(v12 / ((v12) | 1))}
                if false {
                    var v15 uint16 = 812
                    var v16 uint16 = v15
                    for fi17 := 0; fi17 < 15; fi17++ {
                        if (0.0 > 1e10) {
                            if ((v1 == v10) || ((((((-2.25 / 0.1) / 1.5) + (0.0 / (1.5 / float64(1.5)))) + 3.14159) != 3.14159) && (false && (fi17 != fi17)))) {
                                var v18 int16 = (v10 << ((v4) & 63))
                                v16 = 0
                                _ = v18
                            } else {
                                if ((42 < 903) || (!(!false))) {
                                    if (false || (v12 != (fi17 / ((v12) | 1)))) {
                                        for fi19 := 0; fi19 < 17; fi19++ {
                                            v16 = (v16 - 61)
                                            switch ("" + "hello") {
                                            case "a":
                                                var v20 uint = (v4 * 2)
                                                _ = v20
                                            case "foo bar":
                                                if false {
                                                    var v21 string = "hello"
                                                    v5[(v21 + "")] = ((("hello" + "a") + "\n\t") + v21)
                                                    v22 := GenS1{F0: 1}
                                                    _ = v21
                                                    _ = v22
                                                }
                                            case "\n\t":
                                                if true {
                                                    v4 = (v4 % ((v4) | 1))
                                                } else {
                                                    delete(v5, ((("a" + "foo bar") + "hello") + "a"))
                                                    v23 := GenS2{F0: 2147483647}
                                                    _ = v23
                                                }
                                            }
                                            v24 := GenS3{F0: 18446744073709551615, F1: v2}
                                            _ = fi19
                                            _ = v24
                                        }
                                        var v25 uint8 = (^(v13))
                                        if false {
                                            v26 := GenS2{F0: 65535}
                                            for fi27 := 0; fi27 < 15; fi27++ {
                                                if false {
                                                    v5["\n\t"] = ("hello" + "hello")
                                                    var v28 int64 = -1
                                                    if len(v14) > 0 {
                                                        v14[0] = 42
                                                    }
                                                    _ = v28
                                                }
                                                v3 = 1
                                                v10 -= v10
                                                _ = fi27
                                            }
                                            v29 := v5["a"]
                                            _ = v26
                                            _ = v29
                                        } else {
                                            v30 := GenS4{F0: v15}
                                            _ = v30
                                        }
                                        _ = v25
                                    }
                                } else {
                                    var v31 uint32 = 1
                                    for fi32 := 0; fi32 < 15; fi32++ {
                                        v10 = (^(v1))
                                        _ = fi32
                                    }
                                    _ = len(v5)
                                    _ = v31
                                }
                            }
                            var v33 float64 = (1.5 / (1e10 * 1e10))
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 3; fi34++ {
                            if (919 != 18446744073709551615) {
                                var v35 uint32 = 723
                                _ = v35
                            }
                            var v36 float64 = 3.14159
                            _ = fi34
                            _ = v36
                        }
                        _ = fi17
                    }
                    _ = v15
                    _ = v16
                } else {
                    if (((!true) && ((!true) || false)) || false) {
                        var v37 uint = (v4 % ((v4) | 1))
                        fmt.Println(v2, v12)
                        _ = v37
                    }
                    v1 -= (v10 / (((v10 << ((v4) & 63))) | 1))
                }
                _ = v14
            } else {
                var v38 int64 = 946
                _ = v38
            }
            var v39 bool = (42 < (v13 >> (((^(v4))) & 63)))
            v3 += v3
            _ = v39
        }
        _ = v13
    } else {
        if false {
            var v40 string = ""
            var v41 int64 = 42
            var v42 int32 = 42
            _ = v40
            _ = v41
            _ = v42
        } else {
            var v43 uint8 = v3
            fmt.Println(v4, v43)
            _ = v43
        }
    }
    v44 := GenS5{F0: (3.14159 >= (3.14159 / float32((3.14159 / float64(0.1))))), F1: (v2 - 0)}
    v45 := []int{(v12 + 42), int(v1), v12, (v12 + v12)}
    switch ("hello" + ("" + "foo bar")) {
    case "\n\t":
        var v46 uint32 = 1
        _ = v46
    case "foo bar":
        if len(v45) > 0 {
            v45[0] = 0
        }
    default:
        if (1.5 >= (0.0 / 1.0)) {
            var v47 uint64 = 42
            _ = len(v5)
            if (!(!true)) {
                if true {
                    var v48 uint32 = 2
                    var v49 float64 = float64(float32((float64(1.5) * 3.14159)))
                    _ = v48
                    _ = v49
                } else {
                    v5[""] = ("" + "foo bar")
                    var v50 bool = (((v4 - (^(v4))) >= v4) || false)
                    _ = v50
                }
                var v51 uint32 = 42
                _ = v51
            }
            _ = v47
        }
    }
    v2 = (-(v2))
    var v52 int = v12
    if false {
        switch (("a" + ("" + "a")) + ((("" + "\n\t") + "a") + "hello")) {
        case "":
            if ((259 == 100) || false) {
                if (false || true) {
                    v45 = append(v45, (v52 / ((v12) | 1)))
                } else {
                    fmt.Println(v2, v2)
                    for fi53 := 0; fi53 < 3; fi53++ {
                        for fi54 := 0; fi54 < 9; fi54++ {
                            v2 += (v2 / ((61) | 1))
                            _ = fi54
                        }
                        var v55 uint8 = (v3 >> ((42) & 63))
                        v56 := v5["a"]
                        _ = fi53
                        _ = v55
                        _ = v56
                    }
                    if (true || false) {
                        v57 := GenS4{F0: 255}
                        switch "foo bar" {
                        case "a":
                            switch v12 {
                            case 32767:
                                var v58 uint64 = 654
                                _ = v58
                            case 128:
                                v10 += (v1 << ((v4) & 63))
                            }
                        }
                        _ = v57
                    } else {
                        var v59 int32 = -2147483648
                        fmt.Println(v3)
                        _ = v59
                    }
                }
                for fi60 := 0; fi60 < 1; fi60++ {
                    for fi61 := 0; fi61 < 15; fi61++ {
                        switch "hello" {
                        case "\n\t":
                            v10 += v10
                        case "":
                            var v62 int8 = 96
                            _ = v62
                        case "a":
                            v5[""] = (("a" + ("hello" + "a")) + "a")
                        }
                        var v63 uint8 = 1
                        _ = fi61
                        _ = v63
                    }
                    _ = fi60
                }
                var v64 int8 = (-(v2))
                _ = v64
            } else {
                v65 := GenS3{F0: 18446744073709551615, F1: 0}
                _ = v65
            }
        case "\n\t":
            fmt.Println(v1, v3)
        }
        fmt.Println(v1, v4)
        fmt.Println(v4, v2, v4)
    } else {
        for fi66 := 0; fi66 < 12; fi66++ {
            if true {
                fmt.Println(v3)
                var v67 uint64 = 333
                v68 := GenS1{F0: 1}
                _ = v67
                _ = v68
            }
            _ = fi66
        }
        fmt.Println(v4, v4, v3)
    }
    fmt.Println(v4)
    v69 := GenS2{F0: 763}
    for fi70 := 0; fi70 < 11; fi70++ {
        v71 := GenS4{F0: 42}
        _ = fi70
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v52
    _ = v69
}
