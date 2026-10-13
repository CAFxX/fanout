package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 bool; F1 bool; F2 uint16 }
func genFn0() int8 {
    v5 := []int{}
    var v6 string = ("" + "")
    v7 := GenS1{F0: v6, F1: v6}
    v6 = "a"
    _ = v5
    _ = v6
    _ = v7
    return 119
}

func main() {
    var v1 int8 = -128
    v2 := []string{"\n\t", ""}
    v3 := make(map[string]int)
    var v4 uint32 = 83
    v1 = (v1 - (v1 & (^(v1))))
    for fi8 := 0; fi8 < 6; fi8++ {
        switch (("" + "") + "foo bar") {
        case "a":
            v4 = v4
        default:
            switch fi8 {
            case -532:
                if ((!((!false) || false)) && true) {
                    v1 = v1
                } else {
                    if false {
                        switch fi8 {
                        case -635:
                            fmt.Println(fi8)
                        case 128:
                            v9 := GenS1{F0: "hello", F1: (("hello" + "foo bar") + "")}
                            _ = v9
                        case -1:
                            if true {
                                v4 = (v4 / ((2) | 1))
                            } else {
                                if (!true) {
                                    if (true && ((42 != 220) && (0.0 > 1.5))) {
                                        v4 = v4
                                    }
                                } else {
                                    if ((-2.25 - 1.5) <= float64(-2.25)) {
                                        var v10 uint32 = v4
                                        for fi11 := 0; fi11 < 14; fi11++ {
                                            v12 := GenS1{F0: ("foo bar" + "foo bar"), F1: ""}
                                            _ = fi11
                                            _ = v12
                                        }
                                        if (4294967295 > 861) {
                                            for fi13 := 0; fi13 < 11; fi13++ {
                                                if (!(42 == 2)) {
                                                    _ = len(v3)
                                                }
                                                switch fi13 {
                                                case -1:
                                                    if true {
                                                        v10 += 4294967295
                                                    }
                                                }
                                                v14 := GenS2{F0: 826, F1: false}
                                                _ = fi13
                                                _ = v14
                                            }
                                            var v15 uint16 = 0
                                            _ = v15
                                        }
                                        _ = v10
                                    }
                                }
                            }
                        default:
                            v16 := GenS2{F0: 0, F1: true}
                            _ = v16
                        }
                    } else {
                        fmt.Println(v1, v1)
                    }
                    var v17 int = fi8
                    _ = v17
                }
            case 452:
                var v18 float64 = 3.14159
                _ = v18
            case 2147483647:
                for fi19 := 0; fi19 < 16; fi19++ {
                    v1 -= genFn0()
                    _ = fi19
                }
            }
        }
        v4 += (v4 % ((v4) | 1))
        if len(v2) > 0 {
            v2[0] = ("a" + "hello")
        }
        _ = fi8
    }
    var v20 uint16 = 394
    v4 -= 644
    for fi21 := 0; fi21 < 18; fi21++ {
        fmt.Println(fi21, fi21, fi21)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v22 string = "a"
        _ = fi21
        _ = v22
    }
    v1 += 1
    fmt.Println(v4, v4)
    for fi23 := 0; fi23 < 5; fi23++ {
        v24 := GenS2{F0: (^(v20)), F1: ("hello" <= "")}
        _ = fi23
        _ = v24
    }
    var v25 int = -2147483648
    var v26 uint64 = 42
    switch (v25 / (((v25 ^ 0)) | 1)) {
    case 684:
        if true {
            switch (v25 * int(v4)) {
            case 128:
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            case -1:
                var v27 float32 = -2.25
                _ = v27
            case -9223372036854775808:
                v25 = int(v4)
            }
            v28 := GenS2{F0: v20, F1: (false && (!true))}
            for fi29 := 0; fi29 < 10; fi29++ {
                for fi30 := 0; fi30 < 5; fi30++ {
                    _ = len(v2)
                    v25 = (fi30 | fi29)
                    v3["hello"] = 910
                    _ = fi30
                }
                var v31 int64 = 9223372036854775807
                switch v25 {
                case -9223372036854775808:
                    var v32 int = fi29
                    _ = v32
                case 504:
                    switch "\n\t" {
                    case "hello":
                        v20 = 0
                    case "a":
                        if false {
                            v25 += v25
                        } else {
                            v2 = append(v2, "\n\t")
                            var v33 float32 = 0.0
                            var v34 int32 = 42
                            _ = v33
                            _ = v34
                        }
                    case "foo bar":
                        v20 += v20
                    default:
                        v35 := GenS3{F0: false, F1: (-2147483648 > -263), F2: v20}
                        _ = v35
                    }
                }
                _ = fi29
                _ = v31
            }
            _ = v28
        }
    case -744:
        v4 = (v4 & v4)
    case -2147483648:
        v36 := GenS2{F0: (v20 << ((4294967295) & 63)), F1: true}
        _ = v36
    }
    var v37 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v4
}
