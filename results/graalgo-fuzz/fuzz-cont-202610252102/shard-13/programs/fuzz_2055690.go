package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint64 }
type GenS2 struct { F0 int8; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool) int64 {
    fmt.Println(p0)
    p0 = ((32767 <= 0) && (42 < 42))
    v4 := []int{}
    fmt.Println(p0)
    v5 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: 4294967295}
    _ = p0
    _ = v4
    _ = v5
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    v3 := []string{}
    var v6 uint8 = 47
    if (597 < 189) {
        if len(v3) > 0 {
            v3[0] = "a"
        }
        v3 = append(v3, "a")
        var v7 int32 = 2147483647
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 int64 = 469
        var v10 int64 = genFn0(false)
        var v11 float64 = 0.1
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 int = (fi12 + -2147483648)
            var v15 uint = 4294967295
            if false {
                switch "a" {
                case "a":
                    fmt.Println(fi12, v6, v6)
                    fmt.Println(v15, v15, fi12)
                    if false {
                        var v16 int8 = int8(fi12)
                        if ((0.1 - (3.14159 * (float64(3.14159) * ((1.5 * 0.1) - float64((0.0 / 0.1)))))) == ((1e10 * ((0.0 / ((1.5 / (-2.25 + float64(0.1))) - 1.5)) - (((1.5 * (1.5 - (1.5 * 0.1))) / (1e10 - float64(-2.25))) - 1e10))) / 0.1)) {
                            var v17 uint64 = 592
                            _ = v17
                        } else {
                            if true {
                                if true {
                                    v3 = append(v3, ("" + "hello"))
                                } else {
                                    v18 := GenS2{F0: (v16 ^ -1), F1: 32767, F2: 1}
                                    _ = v18
                                }
                                v19 := v1[((((("foo bar" + (("\n\t" + "\n\t") + "")) + ("" + "hello")) + "foo bar") + (((("\n\t" + "\n\t") + (("a" + "foo bar") + "a")) + "hello") + "a")) + "foo bar")]
                                var v20 int = (-(fi12))
                                _ = v19
                                _ = v20
                            } else {
                                v6 = v6
                                v6 -= 235
                            }
                        }
                        var v21 int = 986
                        v22 := GenS1{F0: 80, F1: v21, F2: 200}
                        var v23 bool = (genFn0((-221 == 2147483647)) != 52)
                        _ = v16
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                }
            } else {
                switch v14 {
                case 42:
                    if ("hello" == ((("foo bar" + "\n\t") + ((("hello" + "\n\t") + ("\n\t" + "")) + ("" + ("hello" + ("foo bar" + "foo bar"))))) + "hello")) {
                        fmt.Println(fi13, v6, v14)
                        if false {
                            _ = len(v3)
                            var v24 int16 = 0
                            _ = v24
                        }
                        if (1 != 4294967295) {
                            var v25 uint8 = 128
                            delete(v2, ("hello" + "a"))
                            v26 := GenS1{F0: 4294967295, F1: 1, F2: 310}
                            _ = v25
                            _ = v26
                        } else {
                            var v27 float64 = 1.5
                            _ = v27
                        }
                    }
                case 352:
                    v28 := GenS3{F0: (v6 % ((v6) | 1))}
                    _ = v28
                case 2147483647:
                    var v29 int64 = 637
                    _ = v29
                default:
                    var v30 uint32 = 2147483648
                    _ = v30
                }
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        var v31 int = fi12
        _ = fi12
        _ = v31
    }
    var v32 bool = (9223372036854775807 < -9223372036854775808)
    var v33 int32 = 2147483647
    var v34 float32 = (float32((float64(-2.25) + 0.1)) / 3.14159)
    for fi35 := 0; fi35 < 1; fi35++ {
        v32 = (629 > (^(v33)))
        _ = fi35
    }
    v34 += v34
    var v36 string = "\n\t"
    var v37 uint8 = 18
    if (!v32) {
        var v38 string = v36
        var v39 uint16 = 605
        v40 := GenS3{F0: (v37 - v6)}
        _ = v38
        _ = v39
        _ = v40
    } else {
        v34 += -2.25
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    delete(v2, v36)
    var v41 float32 = (-2.25 * v34)
    for fi42 := 0; fi42 < 20; fi42++ {
        var v43 int = fi42
        _ = fi42
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v41:", v41)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v41
    _ = v6
}
