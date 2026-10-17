package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int8; F1 int16; F2 int32 }
type GenS4 struct { F0 int64; F1 int64 }
func genFn0(p0 int16, p1 int32) uint16 {
    var v5 int8 = 127
    v6 := []int{0}
    if (65535 != 2) {
        v5 = v5
        if true {
            if ((0.0 / ((3.14159 / (0.1 * (float32(float64(1e10)) / float32((-2.25 - 1.5))))) / 1.0)) < 1.5) {
                for fi7 := 0; fi7 < 10; fi7++ {
                    v5 = v5
                    _ = fi7
                }
            }
        } else {
            if len(v6) > 0 {
                v6[0] = 32768
            }
            if (!(4294967295 >= 1)) {
                if false {
                    var v8 uint16 = 773
                    p1 = p1
                    _ = v8
                }
            } else {
                if (576 <= 4294967295) {
                    var v9 float32 = (1.5 / (0.0 + (float32(1.5) - 1e10)))
                    if true {
                        var v10 int32 = -2147483648
                        v9 += 1.5
                        _ = v10
                    } else {
                        if (4294967295 >= 373) {
                            v11 := GenS1{F0: 1, F1: 255, F2: p1}
                            if len(v6) > 0 {
                                v6[0] = 247
                            }
                            for fi12 := 0; fi12 < 6; fi12++ {
                                p0 = 866
                                _ = fi12
                            }
                            _ = v11
                        } else {
                            if (-1 > 42) {
                                switch 42 {
                                case 32768:
                                    if (true && (32768 == 128)) {
                                        for fi13 := 0; fi13 < 17; fi13++ {
                                            if (197 == 255) {
                                                if true {
                                                    var v14 int32 = 0
                                                    var v15 uint8 = 255
                                                    _ = v14
                                                    _ = v15
                                                } else {
                                                    fmt.Println(v5, v5)
                                                }
                                                if false {
                                                    if false {
                                                        fmt.Println(v5, v9, v9)
                                                    }
                                                    p0 -= p0
                                                }
                                                p1 += p1
                                            }
                                            v16 := GenS1{F0: 127, F1: 127, F2: p1}
                                            v5 += int8(p0)
                                            _ = fi13
                                            _ = v16
                                        }
                                        v9 = (v9 * v9)
                                    } else {
                                        fmt.Println(v5)
                                        if (65535 < 42) {
                                            var v17 uint32 = 42
                                            _ = v17
                                        } else {
                                            var v18 int64 = -9223372036854775808
                                            _ = v18
                                        }
                                        for fi19 := 0; fi19 < 4; fi19++ {
                                            if (true || (0 > 159)) {
                                                if (!false) {
                                                    for fi20 := 0; fi20 < 4; fi20++ {
                                                        var v21 string = ("" + "hello")
                                                        v22 := GenS2{F0: 255, F1: p1}
                                                        _ = fi20
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                } else {
                                                    var v23 uint32 = 2
                                                    _ = v23
                                                }
                                                var v24 int64 = 9223372036854775807
                                                var v25 bool = (true || true)
                                                _ = v24
                                                _ = v25
                                            }
                                            _ = fi19
                                        }
                                    }
                                default:
                                    var v26 float32 = 1e10
                                    _ = v26
                                }
                            }
                            var v27 int32 = (p1 * 2147483647)
                            _ = v27
                        }
                        if false {
                            if len(v6) > 0 {
                                v6[0] = 4294967295
                            }
                        } else {
                            for fi28 := 0; fi28 < 1; fi28++ {
                                var v29 int = fi28
                                switch (-(v29)) {
                                case 354:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                default:
                                    v30 := GenS2{F0: 128, F1: p1}
                                    _ = v30
                                }
                                _ = fi28
                                _ = v29
                            }
                        }
                    }
                    _ = v9
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int) uint64 {
    p0 += -1
    v31 := GenS3{F0: 42, F1: 1, F2: 0}
    fmt.Println(p0)
    _ = p0
    _ = v31
    return 42
}

func main() {
    v1 := []int{42}
    v2 := make(map[string]bool)
    var v3 string = ("hello" + ("a" + ("hello" + "hello")))
    var v4 bool = ((false && ((!false) && false)) || true)
    if (v4 || (0.0 == ((((((1e10 - (1.5 / 1e10)) / (-2.25 + 0.1)) * (0.1 + (1e10 + 3.14159))) / 1.5) - 0.1) / (float32(1e10) / (1e10 - (1.5 + 3.14159)))))) {
        var v32 int16 = 42
        _ = v32
    } else {
        v4 = ((v4 || (-1 <= 32768)) && ((42 == 18446744073709551615) && (255 >= 32767)))
    }
    switch (v3 + ((v3 + (v3 + ("" + v3))) + (v3 + "hello"))) {
    case "foo bar":
        v3 = v3
    case "hello":
        v4 = v4
    case "a":
        v3 = v3
    default:
        fmt.Println(v3)
    }
    var v33 float64 = -2.25
    var v34 uint = 0
    var v35 int = 32767
    v33 += v33
    if ((v34 / (((v34 / (((v34 & v34)) | 1))) | 1)) > v34) {
        var v36 uint8 = 54
        v37 := GenS4{F0: 9223372036854775807, F1: -498}
        var v38 int32 = 1
        _ = v36
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        if true {
            fmt.Println(fi39)
        } else {
            v2[v3] = v4
        }
        fmt.Println(v33, v4, v3)
        _ = fi39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
