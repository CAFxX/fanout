package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int16; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint16; F1 bool }
func genFn0(p0 uint32) int16 {
    v5 := []int{}
    switch "hello" {
    case "a":
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 18; fi7++ {
                fmt.Println(fi6)
                _ = fi7
            }
            v8 := GenS1{F0: 127, F1: "foo bar"}
            if false {
                if true {
                    fmt.Println(p0, fi6)
                }
                p0 = 2147483647
            }
            _ = fi6
            _ = v8
        }
    default:
        p0 = p0
    }
    var v9 uint64 = 18446744073709551615
    p0 = p0
    var v10 string = ("hello" + ("a" + "a"))
    _ = p0
    _ = v10
    _ = v5
    _ = v9
    return 1000
}
func genFn1(p0 string, p1 bool) uint16 {
    v11 := GenS1{F0: 69, F1: (("" + p0) + (("hello" + "hello") + p0))}
    p0 = "\n\t"
    _ = p0
    _ = p1
    _ = v11
    return 2
}

func main() {
    var v1 uint64 = 451
    var v2 string = "foo bar"
    var v3 int64 = 42
    v4 := make(map[uint32]int)
    delete(v4, 0)
    v2 = v2
    if (!(true || (v2 != (v2 + "hello")))) {
        var v12 uint64 = v1
        var v13 int16 = 1000
        if false {
            var v14 bool = ("a" == v2)
            _ = v14
        }
        _ = v12
        _ = v13
    } else {
        var v15 float64 = 0.1
        fmt.Println(v2, v1, v1)
        _ = v15
    }
    v16 := GenS2{F0: 1, F1: -1}
    if true {
        for fi17 := 0; fi17 < 15; fi17++ {
            if false {
                if false {
                    var v18 uint = 1
                    v19 := []int{(fi17 >> (((v18 << ((v18) & 63))) & 63))}
                    _ = v18
                    _ = v19
                } else {
                    var v20 int32 = 0
                    fmt.Println(fi17, v3)
                    if false {
                        for fi21 := 0; fi21 < 10; fi21++ {
                            v2 = "foo bar"
                            if false {
                                v22 := GenS3{F0: 0, F1: (v20 % ((v20) | 1)), F2: 226}
                                v3 += v3
                                var v23 string = v2
                                _ = v22
                                _ = v23
                            } else {
                                v3 = (v3 % ((v3) | 1))
                            }
                            switch (fi21 / (((fi21 >> ((0) & 63))) | 1)) {
                            case 0:
                                var v24 float32 = float32((3.14159 + -2.25))
                                _ = v24
                            case -876:
                                if (1.5 <= -2.25) {
                                    if (!(false && ((v3 / (((v3 + (v3 % ((v3) | 1)))) | 1)) <= (^(v3))))) {
                                        if true {
                                            v25 := GenS3{F0: 839, F1: 658, F2: 1}
                                            _ = v25
                                        }
                                        v3 += (v3 / ((v3) | 1))
                                        v2 = v2
                                    } else {
                                        switch ("\n\t" + v2) {
                                        case "hello":
                                            var v26 int64 = 42
                                            _ = v26
                                        }
                                        v20 = v20
                                    }
                                }
                            default:
                                v27 := []int{fi21, 1, fi17, (-(fi21))}
                                _ = v27
                            }
                            _ = fi21
                        }
                    }
                    _ = v20
                }
            }
            if false {
                var v28 float32 = 0.0
                _ = v28
            }
            for fi29 := 0; fi29 < 5; fi29++ {
                var v30 uint = 1
                var v31 int32 = 2147483647
                v31 -= (v31 + (v31 | (-(v31))))
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi17
        }
        var v32 uint16 = 65535
        _ = v32
    } else {
        if (964 == 2147483647) {
            for fi33 := 0; fi33 < 17; fi33++ {
                switch ("foo bar" + v2) {
                case "a":
                    v34 := GenS2{F0: -1, F1: 2147483647}
                    _ = v34
                case "foo bar":
                    v35 := GenS4{F0: 65535, F1: (30 < 128)}
                    _ = v35
                }
                if (!((fi33 << ((4294967295) & 63)) != fi33)) {
                    v36 := []int{fi33, 369, fi33}
                    v37 := GenS1{F0: 161, F1: v2}
                    if true {
                        fmt.Println(v2, v2, v3)
                    } else {
                        var v38 float32 = 0.0
                        _ = v38
                    }
                    _ = v36
                    _ = v37
                }
                v3 += (v3 % (((v3 >> ((913) & 63))) | 1))
                _ = fi33
            }
            if false {
                v2 = ""
            }
        } else {
            for fi39 := 0; fi39 < 12; fi39++ {
                var v40 float32 = float32((3.14159 + 0.0))
                v41 := []int{}
                _ = fi39
                _ = v40
                _ = v41
            }
            v42 := GenS4{F0: uint16(v1), F1: true}
            if false {
                if false {
                    var v43 int32 = 893
                    switch "a" {
                    case "":
                        if (v1 > 4294967295) {
                            v44 := GenS2{F0: v43, F1: 1}
                            if (!false) {
                                var v45 bool = ((v3 == -959) || (uint32(v1) >= 42))
                                if (!v45) {
                                    v46 := GenS2{F0: v43, F1: -1}
                                    switch "a" {
                                    case "":
                                        var v47 uint16 = 709
                                        _ = v47
                                    default:
                                        v3 = v3
                                    }
                                    for fi48 := 0; fi48 < 15; fi48++ {
                                        v49 := GenS2{F0: v43, F1: int32(v3)}
                                        _ = fi48
                                        _ = v49
                                    }
                                    _ = v46
                                }
                                var v50 uint32 = 65535
                                _ = v45
                                _ = v50
                            } else {
                                v51 := []int{128, -995, 4294967295, 32768}
                                var v52 uint = 0
                                var v53 uint16 = 1
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                            fmt.Println(v43, v43, v1)
                            _ = v44
                        } else {
                            fmt.Println(v1, v2, v1)
                        }
                    case "foo bar":
                        v54 := GenS3{F0: 1, F1: v43, F2: 93}
                        _ = v54
                    }
                    var v55 uint8 = 42
                    _ = v43
                    _ = v55
                } else {
                    v56 := GenS2{F0: -2147483648, F1: -890}
                    var v57 float32 = float32(float64(0.1))
                    _ = v56
                    _ = v57
                }
                var v58 uint = 498
                var v59 int32 = -2147483648
                _ = v58
                _ = v59
            }
            _ = v42
        }
    }
    v2 = ((v2 + ((("" + "\n\t") + "") + v2)) + "")
    v60 := GenS4{F0: 1, F1: ((v1 % (((^(v1))) | 1)) != (v1 >> ((818) & 63)))}
    if (!true) {
        var v61 int32 = 42
        _ = v61
    }
    delete(v4, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v60
}
