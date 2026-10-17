package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 bool }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 string; F1 bool; F2 uint }
type GenS4 struct { F0 int; F1 int; F2 int32 }
func genFn0(p0 int64) int8 {
    p0 = p0
    v5 := GenS1{F0: 947, F1: 42, F2: false}
    if false {
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS2{F0: (-(p0)), F1: 0}
            _ = fi6
            _ = v7
        }
        v8 := GenS2{F0: (^(p0)), F1: 0}
        for fi9 := 0; fi9 < 10; fi9++ {
            for fi10 := 0; fi10 < 2; fi10++ {
                v11 := []int{(fi9 << ((2) & 63)), fi10, (fi9 | -721)}
                p0 = 0
                fmt.Println(fi9, fi10, p0)
                _ = fi10
                _ = v11
            }
            var v12 int32 = -2147483648
            _ = fi9
            _ = v12
        }
        _ = v8
    }
    _ = p0
    _ = v5
    return 127
}
func genFn1() int64 {
    if false {
        var v13 uint32 = 42
        _ = v13
    } else {
        v14 := []int{-1, -235}
        _ = v14
    }
    v15 := GenS1{F0: 0, F1: -76, F2: (!true)}
    for fi16 := 0; fi16 < 4; fi16++ {
        switch "a" {
        case "foo bar":
            if true {
                switch fi16 {
                case 2147483647:
                    for fi17 := 0; fi17 < 5; fi17++ {
                        if (42 < 0) {
                            var v18 uint32 = 42
                            fmt.Println(v18, fi17)
                            _ = v18
                        } else {
                            v19 := []int{fi16}
                            var v20 float64 = 3.14159
                            if (true || false) {
                                fmt.Println(fi17, fi17)
                            } else {
                                v20 += 1e10
                                if false {
                                    v20 -= float64((((((1e10 - (((((-2.25 * (0.1 * 1.5)) * 0.1) / ((((((((1.5 - 1.5) / 0.1) * 3.14159) + 0.0) / 1.0) + 3.14159) / 1.0) - (0.1 - ((float32(0.0) + 0.1) - float32(v20))))) / 3.14159) - float32(v20))) + (0.1 / 1e10)) + 0.1) - 0.1) + 0.1))
                                    v21 := GenS2{F0: 1, F1: 4294967295}
                                    fmt.Println(fi16)
                                    var v22 bool = (42 < 65535)
                                    _ = v21
                                    _ = v22
                                } else {
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                }
                                var v23 int32 = 1
                                _ = v23
                            }
                            _ = v19
                            _ = v20
                        }
                        _ = fi17
                    }
                case 133:
                    for fi24 := 0; fi24 < 9; fi24++ {
                        switch ("foo bar" + "hello") {
                        case "":
                            var v25 float64 = (1e10 + 0.1)
                            _ = v25
                        case "a":
                            var v26 int8 = genFn0(int64(fi16))
                            _ = v26
                        case "hello":
                            var v27 uint64 = 18446744073709551615
                            _ = v27
                        }
                        _ = fi24
                    }
                case 32768:
                    v28 := []int{-9223372036854775808, (^(fi16))}
                    _ = v28
                default:
                    var v29 string = "a"
                    _ = v29
                }
                var v30 int64 = 1
                if true {
                    if false {
                        if ((127 == 0) && (int(v30) != -214)) {
                            if (!((v30 * 0) >= (v30 / (((v30 / ((v30) | 1))) | 1)))) {
                                v31 := []int{fi16, int(v30), 32767, (fi16 / (((-(fi16))) | 1))}
                                _ = v31
                            } else {
                                switch (("hello" + ("\n\t" + (("" + "hello") + "hello"))) + "foo bar") {
                                case "foo bar":
                                    for fi32 := 0; fi32 < 12; fi32++ {
                                        var v33 string = (("" + "a") + "hello")
                                        var v34 uint = 455
                                        var v35 uint32 = 2147483648
                                        _ = fi32
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                default:
                                    if true {
                                        v36 := []int{(fi16 * (fi16 & fi16))}
                                        v37 := GenS1{F0: -1, F1: -1, F2: ((3.14159 - -2.25) != 1.5)}
                                        _ = v36
                                        _ = v37
                                    } else {
                                        fmt.Println(v30, fi16)
                                        v30 = (v30 + (v30 & (v30 % ((v30) | 1))))
                                    }
                                }
                                for fi38 := 0; fi38 < 10; fi38++ {
                                    v39 := []int{(fi16 + (fi16 & (fi38 + fi38))), (^(fi38)), fi38}
                                    var v40 uint16 = 32767
                                    if true {
                                        var v41 int16 = 1000
                                        var v42 int8 = 103
                                        var v43 float32 = 1e10
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                            }
                            fmt.Println(fi16)
                            if false {
                                var v44 uint64 = 430
                                _ = v44
                            } else {
                                v45 := GenS3{F0: ("" + "a"), F1: (((!false) && false) && ((!((!false) || true)) && (0 != 0))), F2: 190}
                                _ = v45
                            }
                        }
                    }
                    var v46 string = ("hello" + "")
                    var v47 int8 = genFn0((v30 + (v30 / ((142) | 1))))
                    _ = v46
                    _ = v47
                }
                _ = v30
            } else {
                v48 := GenS3{F0: "", F1: (2147483648 < 1), F2: 65535}
                _ = v48
            }
        }
        _ = fi16
    }
    _ = v15
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 2
    var v2 float64 = 1e10
    v3 := make(map[int]int)
    v4 := []uint32{0, 435, 572, 42}
    fmt.Println(v1)
    switch 1 {
    case 0:
        _ = len(v3)
    }
    for fi49 := 0; fi49 < 12; fi49++ {
        switch "" {
        case "a":
            switch 358 {
            case 32767:
                var v50 uint64 = 0
                _ = v50
            case -1:
                for fi51 := 0; fi51 < 12; fi51++ {
                    switch fi49 {
                    case -126:
                        v2 = (v2 - (0.0 * v2))
                    case 0:
                        v52 := GenS4{F0: -2147483648, F1: fi49, F2: 786}
                        _ = v52
                    }
                    switch "\n\t" {
                    case "":
                        var v53 float64 = ((0.1 / (float64(0.1) - v2)) / v2)
                        _ = v53
                    default:
                        fmt.Println(v1, v1, v1)
                    }
                    _ = fi51
                }
            case 2147483647:
                for fi54 := 0; fi54 < 3; fi54++ {
                    var v55 string = ((("" + (("\n\t" + (("a" + "a") + "foo bar")) + ("" + "\n\t"))) + ("hello" + ("hello" + ("a" + "")))) + ("hello" + ("foo bar" + ("a" + ("foo bar" + "")))))
                    var v56 float64 = v2
                    var v57 int32 = 869
                    _ = fi54
                    _ = v55
                    _ = v56
                    _ = v57
                }
            default:
                fmt.Println(v1)
            }
        }
        var v58 int32 = 1
        _ = fi49
        _ = v58
    }
    v59 := GenS3{F0: ("hello" + "\n\t"), F1: false, F2: 0}
    v1 += v1
    var v60 float64 = v2
    var v61 int32 = -2147483648
    v4 = append(v4, 2147483648)
    var v62 int = 127
    v61 = int32(v62)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v60
    _ = v61
    _ = v62
}
