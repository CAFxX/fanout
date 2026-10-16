package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }
type GenS4 struct { F0 int; F1 int; F2 int8 }

func main() {
    var v1 int16 = 751
    v2 := make(map[int]int)
    v3 := []int{}
    v4 := make(map[string]bool)
    var v5 bool = (true || (!false))
    v5 = (!((!false) && (int64(v1) >= 42)))
    fmt.Println(v1, v1)
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(v1, v5)
            if (2147483647 <= 2147483648) {
                switch "" {
                case "foo bar":
                    var v8 int64 = 1
                    _ = v8
                case "hello":
                    v5 = v5
                }
                if ((718 >= 34) || v5) {
                    v9 := GenS1{F0: (-(fi6)), F1: 678, F2: "\n\t"}
                    _ = v9
                } else {
                    delete(v2, (fi6 ^ fi7))
                    switch ("hello" + "hello") {
                    case "hello":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        if (80 != 573) {
                            var v10 uint16 = 0
                            _ = v10
                        } else {
                            v11 := GenS1{F0: (^(fi7)), F1: 18446744073709551615, F2: "foo bar"}
                            v5 = v5
                            _ = v11
                        }
                    }
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi7
        }
        _ = fi6
    }
    if true {
        if false {
            var v12 uint16 = 0
            _ = v12
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        if v5 {
            var v13 float64 = (0.0 - 0.0)
            _ = v13
        } else {
            v1 += 32767
            var v14 float32 = (0.1 - 0.0)
            if true {
                v14 -= ((0.1 * 0.0) * float32(-2.25))
                v14 += 1e10
            }
            _ = v14
        }
        if ((1.5 / (float32(0.0) - ((((3.14159 / 0.1) * ((1e10 - 1e10) * float32(float64((((1e10 / 1.0) - 0.0) * 0.1))))) - float32(float64(1.5))) - 0.0))) == -2.25) {
            v4["hello"] = v5
            var v15 bool = true
            v5 = true
            _ = v15
        } else {
            v16 := GenS1{F0: 4294967295, F1: 0, F2: ""}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v17 := v2[672]
            _ = v16
            _ = v17
        }
    }
    v5 = v5
    var v18 int64 = 291
    for fi19 := 0; fi19 < 19; fi19++ {
        v5 = ((75 != 424) || (!true))
        _ = fi19
    }
    v20 := GenS2{F0: 524}
    switch 128 {
    case -5:
        var v21 bool = (("hello" + "hello") > "hello")
        _ = v21
    case -1:
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v22 := GenS3{F0: 127, F1: int(v18), F2: -1}
            var v23 float32 = float32((3.14159 / (float64(0.0) + (-2.25 + float64(0.0)))))
            _ = v22
            _ = v23
        } else {
            var v24 string = ""
            v3 = append(v3, 32768)
            if ((42 == 0) && ((0 <= 100) && v5)) {
                if v5 {
                    for fi25 := 0; fi25 < 3; fi25++ {
                        var v26 bool = ((fi25 + 948) >= (fi25 / ((1) | 1)))
                        v3 = append(v3, (fi25 ^ 4294967295))
                        _ = fi25
                        _ = v26
                    }
                    var v27 uint32 = 4294967295
                    for fi28 := 0; fi28 < 6; fi28++ {
                        if true {
                            for fi29 := 0; fi29 < 15; fi29++ {
                                fmt.Println(v18, v27, v1)
                                _ = fi29
                            }
                            for fi30 := 0; fi30 < 16; fi30++ {
                                fmt.Println(fi30)
                                if ((1.5 - 0.1) <= (0.1 / 1.5)) {
                                    var v31 string = v24
                                    v3 = append(v3, fi30)
                                    fmt.Println(v5, fi30, v1)
                                    if (v31 > "") {
                                        _ = len(v3)
                                        var v32 int32 = -1
                                        v2[fi30] = fi30
                                        _ = v32
                                    }
                                    _ = v31
                                } else {
                                    var v33 string = (v24 + ("hello" + "hello"))
                                    if (-2.25 != (float32(0.0) + -2.25)) {
                                        switch fi30 {
                                        case 32767:
                                            fmt.Println(v24)
                                        case 9223372036854775807:
                                            fmt.Println(v1)
                                        }
                                        if v5 {
                                            var v34 string = v24
                                            v27 += 2
                                            v5 = (815 == (v1 + v1))
                                            _ = v34
                                        } else {
                                            _ = len(v3)
                                            if (-128 <= 1) {
                                                fmt.Println(v33)
                                                fmt.Println(v1)
                                                v33 = v33
                                            } else {
                                                fmt.Println(v18, fi30)
                                                v2[fi28] = (fi30 * fi28)
                                                var v35 uint64 = 0
                                                _ = v35
                                            }
                                        }
                                        var v36 uint32 = (v27 - v27)
                                        _ = v36
                                    } else {
                                        v37 := GenS2{F0: 18446744073709551615}
                                        _ = v37
                                    }
                                    _ = len(v3)
                                    delete(v4, v33)
                                    var v38 bool = (v33 > v33)
                                    fmt.Println(v1, v33)
                                    v27 = v27
                                    _ = v33
                                    _ = v38
                                }
                                v18 -= (v18 | -1)
                                _ = fi30
                            }
                            v24 = v24
                        } else {
                            var v39 float32 = 0.1
                            v40 := GenS4{F0: 32768, F1: (fi28 << ((607) & 63)), F2: -1}
                            v41 := GenS1{F0: (fi28 % ((fi28) | 1)), F1: 615, F2: v24}
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        _ = fi28
                    }
                    _ = v27
                }
            }
            _ = v24
        }
    default:
        fmt.Println(v5, v18, v5)
    }
    v2[127] = 0
    v42 := GenS4{F0: 128, F1: 1, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
