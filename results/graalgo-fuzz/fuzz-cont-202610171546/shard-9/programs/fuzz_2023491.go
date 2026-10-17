package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 bool }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 int16 }
func genFn0() int32 {
    var v6 bool = false
    v7 := []int{-444, 127}
    if (!false) {
        var v8 int16 = 713
        if v6 {
            v6 = ((!true) || ((3.14159 == (0.0 / (3.14159 / ((((((0.1 + ((1.5 + (((float64(1.5) - 0.1) / 1.0) + float64((3.14159 + float32(3.14159))))) * (1.5 + 1e10))) + (0.0 - (1e10 * -2.25))) / 1.5) / (-2.25 - (((0.0 + 0.0) / (-2.25 + 1.5)) + 3.14159))) / float64(((-2.25 / float32(float64(0.1))) + ((float32(1e10) - (float32((-2.25 / 1.5)) * (0.1 / 3.14159))) / 0.1)))) + (1e10 / 1e10))))) || v6))
        }
        v7 = append(v7, -1)
        _ = v8
    } else {
        if v6 {
            v6 = (-1 != 295)
            var v9 uint16 = 1
            _ = v9
        } else {
            v10 := GenS1{F0: 2147483647, F1: "foo bar", F2: v6}
            var v11 uint64 = 1
            _ = v10
            _ = v11
        }
    }
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint8) bool {
    p0 = p0
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := []int{(fi12 & (fi12 * (fi12 | fi12))), fi12, fi12}
        if (p0 > p0) {
            _ = len(v13)
        }
        _ = fi12
        _ = v13
    }
    _ = p0
    return false
}

func main() {
    v1 := []int{-1, 32767, 2147483647, 549}
    v2 := make(map[int]string)
    var v3 uint = 65535
    v4 := make(map[string]bool)
    var v5 string = "foo bar"
    var v14 int32 = 1
    for fi15 := 0; fi15 < 8; fi15++ {
        fmt.Println(fi15)
        _ = fi15
    }
    var v16 float64 = 3.14159
    v14 = (v14 | -56)
    v5 = v5
    v3 -= v3
    v17 := GenS1{F0: genFn0(), F1: v5, F2: true}
    if true {
        v3 = v3
    }
    var v18 uint = (^(v3))
    var v19 int64 = 42
    v14 = (v14 * (v14 - (^(v14))))
    for fi20 := 0; fi20 < 2; fi20++ {
        if false {
            var v21 float64 = ((1e10 + v16) / 0.1)
            _ = v21
        } else {
            var v22 uint16 = 836
            if ((((3.14159 * -2.25) - (3.14159 / 1.5)) + ((0.1 * 0.0) + -2.25)) >= 3.14159) {
                fmt.Println(v16)
                v16 = float64(0.1)
            } else {
                if true {
                    if (v22 < (^(v22))) {
                        v23 := GenS2{F0: (!(1 >= 255)), F1: 735}
                        v19 = v19
                        for fi24 := 0; fi24 < 5; fi24++ {
                            var v25 uint16 = v22
                            _ = fi24
                            _ = v25
                        }
                        _ = v23
                    } else {
                        for fi26 := 0; fi26 < 13; fi26++ {
                            fmt.Println(v22)
                            delete(v2, int(v22))
                            if (!false) {
                                fmt.Println(v5)
                            } else {
                                if true {
                                    var v27 bool = (((v14 ^ (v14 & v14)) == v14) || true)
                                    v4[("foo bar" + (v5 + v5))] = v27
                                    v27 = (v27 || (647 >= 65535))
                                    _ = v27
                                } else {
                                    var v28 int16 = 42
                                    v1 = append(v1, fi20)
                                    if false {
                                        if (!((43 == 2) && (!true))) {
                                            var v29 uint8 = 135
                                            v30 := GenS3{F0: v19, F1: 1}
                                            fmt.Println(fi20)
                                            _ = v29
                                            _ = v30
                                        }
                                        var v31 int8 = -128
                                        fmt.Println(v14, v3)
                                        var v32 int8 = v31
                                        _ = v31
                                        _ = v32
                                    }
                                    _ = v28
                                }
                                fmt.Println(v22)
                                v33 := GenS1{F0: (v14 >> (((v18 >> ((v3) & 63))) & 63)), F1: (v5 + "foo bar"), F2: true}
                                if false {
                                    _ = len(v4)
                                    fmt.Println(v18, fi20)
                                }
                                if (42 > 37) {
                                    if true {
                                        var v34 uint32 = 379
                                        var v35 uint32 = (^(v34))
                                        var v36 float32 = 3.14159
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    } else {
                                        v16 = v16
                                        v19 -= v19
                                    }
                                    var v37 uint = v18
                                    switch (v5 + ((v5 + ("a" + v5)) + ((v5 + "\n\t") + v5))) {
                                    case "hello":
                                        v19 += v19
                                    }
                                    v38 := v2[fi26]
                                    v37 = uint(v22)
                                    _ = v37
                                    _ = v38
                                }
                                if false {
                                    v39 := v4[(("a" + v5) + ((v5 + v5) + ((("foo bar" + (v5 + "")) + v5) + (v5 + "foo bar"))))]
                                    _ = v39
                                }
                                var v40 bool = (v3 != (^(v3)))
                                v41 := GenS2{F0: v40, F1: 1}
                                _ = v33
                                _ = v40
                                _ = v41
                            }
                            _ = fi26
                        }
                        if false {
                            switch ("a" + (("hello" + "\n\t") + ("\n\t" + "a"))) {
                            case "a":
                                v19 = v19
                            case "\n\t":
                                v42 := GenS4{F0: 835}
                                _ = v42
                            default:
                                v3 += v18
                            }
                            var v43 float64 = (v16 * ((v16 + (((float64(0.0) + v16) / float64(float32(float64((1.5 + 1e10))))) + float64((0.1 - 0.1)))) / v16))
                            var v44 int64 = v19
                            _ = v43
                            _ = v44
                        }
                    }
                    if false {
                        for fi45 := 0; fi45 < 18; fi45++ {
                            v46 := v4[("\n\t" + v5)]
                            var v47 float32 = (0.0 / 0.1)
                            var v48 bool = v46
                            _ = fi45
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                        var v49 bool = false
                        _ = v49
                    } else {
                        v1 = append(v1, (fi20 % ((fi20) | 1)))
                        for fi50 := 0; fi50 < 4; fi50++ {
                            var v51 int64 = -1
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi50
                            _ = v51
                        }
                    }
                    v52 := v4[""]
                    _ = v52
                } else {
                    for fi53 := 0; fi53 < 5; fi53++ {
                        switch (fi53 >> ((v18) & 63)) {
                        case 979:
                            fmt.Println(fi53, v5, fi53)
                        }
                        if false {
                            var v54 uint32 = 0
                            var v55 int8 = 0
                            switch "hello" {
                            case "":
                                v5 = (v5 + v5)
                            case "\n\t":
                                var v56 uint32 = v54
                                _ = v56
                            }
                            _ = v54
                            _ = v55
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi53
                    }
                    var v57 uint = v3
                    v58 := GenS5{F0: v19, F1: -32768}
                    _ = v57
                    _ = v58
                }
                switch 818 {
                case 4294967295:
                    if false {
                        fmt.Println(v14, v22)
                    } else {
                        if false {
                            var v59 uint16 = (v22 << ((v3) & 63))
                            _ = v59
                        } else {
                            for fi60 := 0; fi60 < 19; fi60++ {
                                var v61 uint64 = 587
                                switch fi60 {
                                case -570:
                                    var v62 float32 = float32(v16)
                                    _ = v62
                                case -2147483648:
                                    v16 -= v16
                                case 582:
                                    v19 += (v19 * v19)
                                default:
                                    v19 += v19
                                }
                                _ = fi60
                                _ = v61
                            }
                            v3 = (v3 & v3)
                        }
                    }
                case -597:
                    v22 += (v22 | v22)
                case -503:
                    if ((v22 ^ v22) <= 65535) {
                        v16 -= 1e10
                    }
                }
                v63 := GenS3{F0: (v19 << ((v3) & 63)), F1: 2}
                _ = v63
            }
            _ = v22
        }
        _ = len(v2)
        _ = fi20
    }
    _ = len(v4)
    var v64 int8 = 0
    v19 += -1
    v18 += v3
    v65 := GenS4{F0: (v3 ^ v3)}
    if (!(true && false)) {
        var v66 float64 = (v16 * v16)
        _ = v66
    } else {
        var v67 int16 = 0
        _ = v67
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v65
}
