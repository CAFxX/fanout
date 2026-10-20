package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 string; F1 string }
func genFn0() string {
    var v4 uint8 = 2
    v5 := []int{4294967295, 127, -553}
    if true {
        switch 42 {
        case 127:
            v6 := GenS1{F0: 42}
            _ = v6
        case 1:
            fmt.Println(v4)
        case -214:
            v5 = append(v5, 42)
        }
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(v4, fi7)
        v4 = (v4 | v4)
        if (65535 <= 32767) {
            for fi8 := 0; fi8 < 5; fi8++ {
                v4 -= 102
                if false {
                    var v9 float32 = (1.5 / 1e10)
                    _ = v9
                }
                fmt.Println(fi7, fi7, fi7)
                v4 = v4
                if len(v5) > 0 {
                    v5[0] = fi7
                }
                _ = fi8
            }
            var v10 uint = 0
            _ = v10
        }
        _ = fi7
    }
    _ = v4
    _ = v5
    return ("hello" + "a")
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ("hello" + "")
    v3 := []int{}
    _ = len(v3)
    if (-259 >= 4294967295) {
        if (!false) {
            v2 = v2
            if true {
                v11 := GenS1{F0: 2}
                var v12 bool = (!(127 != 50))
                _ = v11
                _ = v12
            } else {
                var v13 uint64 = 0
                _ = v13
            }
        } else {
            v2 = "hello"
        }
        var v14 float32 = ((float32(3.14159) + ((3.14159 / 3.14159) / (float32(float64(3.14159)) * (3.14159 * ((3.14159 + ((-2.25 / 1.0) - 1.5)) * -2.25))))) / 1.0)
        _ = v14
    }
    _ = len(v1)
    v2 = "foo bar"
    _ = len(v3)
    _ = len(v1)
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int8 = -75
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: 4294967295, F1: 0}
    v18 := v1[-2147483648]
    var v19 int64 = 485
    v20 := GenS2{F0: 593, F1: uint8(v18)}
    if (false && (746 <= 740)) {
        if (!false) {
            var v21 string = ""
            _ = v21
        }
        var v22 int16 = 42
        _ = v22
    }
    if false {
        if true {
            v23 := GenS1{F0: 566}
            fmt.Println(v2, v18, v2)
            if false {
                var v24 uint8 = 48
                var v25 uint8 = v24
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v24
                _ = v25
            }
            _ = v23
        } else {
            fmt.Println(v2, v2)
            var v26 string = ((((v2 + ("" + v2)) + (v2 + v2)) + "foo bar") + ("\n\t" + ""))
            v2 = v2
            _ = v26
        }
        v27 := GenS2{F0: 9223372036854775807, F1: 1}
        var v28 int16 = 32767
        _ = v27
        _ = v28
    } else {
        for fi29 := 0; fi29 < 6; fi29++ {
            if true {
                if (v2 < v2) {
                    var v30 uint = 4294967295
                    v2 = "\n\t"
                    _ = v30
                }
                for fi31 := 0; fi31 < 2; fi31++ {
                    if (!false) {
                        v2 = v2
                        var v32 int = (fi29 ^ fi29)
                        _ = v32
                    } else {
                        if true {
                            var v33 bool = (1 <= 742)
                            var v34 float32 = 0.1
                            _ = v33
                            _ = v34
                        }
                        var v35 uint16 = 1
                        if (fi31 < (fi31 % ((1) | 1))) {
                            var v36 int8 = 29
                            _ = v36
                        } else {
                            var v37 float64 = 1.5
                            if (v19 < 0) {
                                switch fi31 {
                                case -2147483648:
                                    fmt.Println(v18, v37)
                                default:
                                    if false {
                                        _ = len(v1)
                                        var v38 int = fi29
                                        v37 += float64(0.0)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v38
                                    } else {
                                        v37 = 0.0
                                        v39 := GenS2{F0: fi29, F1: 127}
                                        v18 = v18
                                        v40 := GenS3{F0: v2, F1: ((v2 + "a") + (("" + "") + v2))}
                                        _ = v39
                                        _ = v40
                                    }
                                }
                            } else {
                                var v41 bool = true
                                _ = v41
                            }
                            _ = v37
                        }
                        _ = v35
                    }
                    _ = fi31
                }
            }
            var v42 int8 = -1
            var v43 int8 = v42
            _ = fi29
            _ = v42
            _ = v43
        }
    }
    v18 = (v18 + 1)
    fmt.Println(v19)
    if ((823 >= uint(v18)) && (!(true && (568 > -9223372036854775808)))) {
        var v44 int8 = 0
        var v45 int = -828
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 11; fi46++ {
        var v47 int32 = -1
        v48 := GenS3{F0: (v2 + "a"), F1: v2}
        _ = fi46
        _ = v47
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
