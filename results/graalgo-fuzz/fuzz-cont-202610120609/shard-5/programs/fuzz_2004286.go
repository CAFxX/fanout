package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 bool }
type GenS2 struct { F0 int64; F1 string; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 bool; F2 uint }
type GenS8 struct { F0 int }
func genFn0(p0 int64) string {
    var v7 uint32 = 0
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            fmt.Println(fi8, fi8)
            v10 := []int{}
            if len(v10) > 0 {
                v10[0] = int(p0)
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v11 string = ""
    v7 -= (v7 * v7)
    _ = p0
    _ = v11
    _ = v7
    return v11
}
func genFn1(p0 int16) int8 {
    v12 := GenS1{F0: 173, F1: 42, F2: true}
    for fi13 := 0; fi13 < 11; fi13++ {
        p0 += p0
        _ = fi13
    }
    _ = p0
    _ = v12
    return -6
}

func main() {
    v1 := []string{}
    v2 := []uint32{0}
    var v3 float64 = ((0.1 / ((float64((1.5 + 1e10)) + float64(-2.25)) * 0.1)) - 1.5)
    var v4 string = ((("hello" + ("a" + "a")) + "a") + "a")
    var v5 float32 = 3.14159
    v6 := make(map[string]string)
    switch 127 {
    case 32768:
        v14 := GenS1{F0: 255, F1: -121, F2: (453 > -372)}
        _ = v14
    default:
        v4 = ("foo bar" + "\n\t")
    }
    v15 := v6[v4]
    v3 = 1.5
    if len(v2) > 0 {
        v2[0] = 1
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            fmt.Println(fi17, v4)
            v18 := GenS2{F0: 0, F1: (v4 + ((v4 + (v4 + "a")) + "foo bar")), F2: 75}
            fmt.Println(v5)
            var v19 string = "a"
            v2 = append(v2, 339)
            var v20 float64 = -2.25
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            fmt.Println(fi16, v5)
            v22 := v6[v4]
            _ = fi21
            _ = v22
        }
        _ = fi16
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v15 = "a"
    if true {
        var v23 bool = true
        if (128 == -95) {
            for fi24 := 0; fi24 < 19; fi24++ {
                switch (fi24 / ((fi24) | 1)) {
                case -9223372036854775808:
                    for fi25 := 0; fi25 < 17; fi25++ {
                        var v26 uint32 = 199
                        if (v23 || (2 >= uint(v26))) {
                            v27 := GenS1{F0: 128, F1: 1, F2: (128 > 61)}
                            if v23 {
                                if v23 {
                                    if ((v26 ^ v26) != v26) {
                                        v3 = ((v3 / (0.0 * v3)) / ((((float64(v5) + v3) - v3) + 0.1) / 1.5))
                                    } else {
                                        _ = len(v6)
                                        v28 := GenS3{F0: 415, F1: (v26 / (((v26 >> ((2) & 63))) | 1))}
                                        var v29 int64 = 982
                                        if v23 {
                                            v30 := GenS4{F0: v4, F1: 0}
                                            var v31 int = (fi25 ^ 718)
                                            _ = v30
                                            _ = v31
                                        } else {
                                            var v32 uint32 = (^(v26))
                                            delete(v6, v4)
                                            _ = v32
                                        }
                                        _ = v28
                                        _ = v29
                                    }
                                    var v33 uint8 = 0
                                    _ = v33
                                }
                                delete(v6, (v15 + ((v15 + v15) + "hello")))
                            }
                            switch ("a" + v4) {
                            case "foo bar":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "hello":
                                v34 := GenS5{F0: 0}
                                _ = v34
                            case "\n\t":
                                if true {
                                    v26 -= (^(v26))
                                    var v35 uint8 = 127
                                    v5 += v5
                                    var v36 uint8 = v35
                                    _ = v35
                                    _ = v36
                                }
                            }
                            _ = v27
                        }
                        _ = fi25
                        _ = v26
                    }
                case 758:
                    fmt.Println(v4)
                case 554:
                    var v37 int32 = 42
                    _ = v37
                }
                _ = fi24
            }
            delete(v6, "a")
            for fi38 := 0; fi38 < 1; fi38++ {
                var v39 string = v15
                for fi40 := 0; fi40 < 19; fi40++ {
                    if ((v23 || v23) && false) {
                        v41 := v6[((v39 + v39) + v15)]
                        v23 = v23
                        v41 = ("" + v39)
                        _ = v41
                    } else {
                        _ = len(v6)
                        var v42 string = v4
                        if v23 {
                            v43 := GenS1{F0: 255, F1: 102, F2: (((42 == 5) || v23) || (!v23))}
                            if v23 {
                                v44 := GenS6{F0: 42}
                                switch (fi40 / ((fi38) | 1)) {
                                case 32767:
                                    var v45 uint32 = 2147483648
                                    _ = v45
                                case 128:
                                    fmt.Println(v3, v39)
                                case 255:
                                    if ((!v23) || (27 <= 42)) {
                                        v46 := GenS7{F0: "\n\t", F1: v23, F2: 463}
                                        v47 := GenS8{F0: (fi40 + -860)}
                                        _ = len(v6)
                                        _ = v46
                                        _ = v47
                                    }
                                default:
                                    v4 = (v42 + v42)
                                }
                                _ = v44
                            }
                            _ = v43
                        }
                        _ = v42
                    }
                    if len(v1) > 0 {
                        v1[0] = (v15 + (v4 + (v15 + v39)))
                    }
                    var v48 bool = true
                    _ = fi40
                    _ = v48
                }
                _ = fi38
                _ = v39
            }
        }
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
