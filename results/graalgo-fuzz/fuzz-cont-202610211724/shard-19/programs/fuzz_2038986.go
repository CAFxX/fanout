package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 int }
func genFn0(p0 int64, p1 uint8) int {
    v4 := GenS1{F0: 127, F1: true, F2: (p1 % (((^(p1))) | 1))}
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            fmt.Println(fi6, p1, p1)
            var v7 uint8 = p1
            v8 := []int{fi6, fi5, (^(fi6)), fi5}
            _ = fi6
            _ = v7
            _ = v8
        }
        p1 += (p1 / ((127) | 1))
        if (2 == 4294967295) {
            v9 := []int{}
            fmt.Println(p0)
            _ = v9
        }
        _ = fi5
    }
    var v10 uint16 = 633
    _ = p0
    _ = p1
    _ = v10
    _ = v4
    return 32767
}

func main() {
    var v1 uint8 = 127
    var v2 uint16 = 1
    var v3 int = 1
    if true {
        v11 := []int{(v3 << ((42) & 63)), int(v1)}
        v12 := GenS1{F0: (v1 >> ((42) & 63)), F1: true, F2: v1}
        _ = v11
        _ = v12
    }
    var v13 int64 = -1
    if (("\n\t" + "foo bar") == ("" + "a")) {
        v14 := GenS2{F0: 1}
        _ = v14
    } else {
        if true {
            v15 := []int{-2147483648, (v3 + genFn0(0, (v1 | v1)))}
            _ = v15
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 int64 = -9223372036854775808
            fmt.Println(v2, v3)
            var v18 uint16 = (v2 << ((65535) & 63))
            _ = fi16
            _ = v17
            _ = v18
        }
        if (0 < 42) {
            for fi19 := 0; fi19 < 14; fi19++ {
                fmt.Println(fi19, v2, v3)
                if ((("a" + "\n\t") == (("" + (((("\n\t" + "hello") + ("hello" + "foo bar")) + (("foo bar" + "a") + ("foo bar" + "\n\t"))) + "a")) + "")) && false) {
                    v20 := []int{-2147483648, 127, -1, genFn0(int64(v2), (v1 * (v1 - v1)))}
                    fmt.Println(v2, v2, v1)
                    if false {
                        fmt.Println(v13, v3)
                        var v21 float64 = (float64(1.5) - ((float64(0.1) + (1e10 * 1e10)) / (float64(-2.25) - float64(3.14159))))
                        v22 := GenS3{F0: (v2 * v2), F1: 1}
                        _ = v21
                        _ = v22
                    } else {
                        _ = len(v20)
                        var v23 int8 = 1
                        var v24 uint = 752
                        _ = v23
                        _ = v24
                    }
                    _ = v20
                } else {
                    v25 := GenS4{F0: -1}
                    if false {
                        for fi26 := 0; fi26 < 2; fi26++ {
                            v3 = fi19
                            fmt.Println(v13, v2, v13)
                            v27 := []int{}
                            v27 = append(v27, (fi19 >> ((1) & 63)))
                            if true {
                                if false {
                                    v28 := GenS5{F0: 2}
                                    if false {
                                        var v29 uint8 = v1
                                        _ = v29
                                    }
                                    _ = v28
                                } else {
                                    var v30 uint = 4294967295
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                    _ = v30
                                }
                                switch 393 {
                                case 32768:
                                    v31 := GenS6{F0: 43, F1: 9223372036854775807}
                                    _ = v31
                                case 127:
                                    _ = len(v27)
                                case 2147483647:
                                    if len(v27) > 0 {
                                        v27[0] = (fi26 | fi26)
                                    }
                                }
                                var v32 string = "foo bar"
                                _ = v32
                            } else {
                                if len(v27) >= 2 {
                                    v27 = v27[1:len(v27)]
                                }
                                var v33 int64 = 42
                                var v34 int64 = -9223372036854775808
                                _ = v33
                                _ = v34
                            }
                            _ = fi26
                            _ = v27
                        }
                        v3 -= fi19
                        if true {
                            v3 = v3
                        } else {
                            for fi35 := 0; fi35 < 3; fi35++ {
                                v13 += v13
                                fmt.Println(v3, v3)
                                v36 := []int{32767, fi35, 128}
                                _ = fi35
                                _ = v36
                            }
                            var v37 int = genFn0((v13 & v13), (^(v1)))
                            _ = v37
                        }
                    } else {
                        var v38 uint16 = 2
                        _ = v38
                    }
                    _ = v25
                }
                _ = fi19
            }
        } else {
            v39 := []int{v3, genFn0(1, 40), (v3 - (-(v3))), (v3 * (-(v3)))}
            v40 := GenS1{F0: (^(v1)), F1: false, F2: 0}
            fmt.Println(v2, v3, v1)
            _ = v39
            _ = v40
        }
    }
    var v41 int16 = 925
    var v42 uint16 = v2
    var v43 int = 9223372036854775807
    v44 := GenS4{F0: 2147483647}
    v1 = v1
    v45 := []int{}
    if false {
        fmt.Println(v41, v13, v42)
        var v46 int16 = v41
        if len(v45) > 0 {
            v45[0] = (v3 - (v3 / ((v43) | 1)))
        }
        _ = v46
    }
    fmt.Println(v3)
    if ((v1 + (v1 | (^(v1)))) > (^(v1))) {
        if (-1 > v13) {
            v47 := GenS2{F0: 870}
            _ = v47
        }
        var v48 uint64 = 380
        _ = v48
    }
    var v49 int32 = -2147483648
    var v50 int64 = (v13 & (v13 / ((v13) | 1)))
    v41 = v41
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v49
    _ = v50
}
