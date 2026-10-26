package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int }
func genFn0() int16 {
    var v6 bool = (!(4294967295 != 42))
    var v7 float32 = (0.0 - 1e10)
    var v8 uint8 = 244
    var v9 float32 = (float32((1e10 - -2.25)) / v7)
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 uint64, p1 uint16) bool {
    v10 := GenS1{F0: 255, F1: 174, F2: 128}
    if (false && ((687 >= -2147483648) || ("hello" != "hello"))) {
        p1 -= p1
        var v11 int32 = 42
        if (1000 != 136) {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 int16 = 42
                switch ("hello" + "foo bar") {
                case "foo bar":
                    if (!(!(0 != -9223372036854775808))) {
                        v13 += (v13 + (v13 >> ((4294967295) & 63)))
                    } else {
                        var v14 uint32 = 0
                        for fi15 := 0; fi15 < 1; fi15++ {
                            v16 := GenS1{F0: 18446744073709551615, F1: (p0 & 2), F2: (fi15 - -812)}
                            p1 = 1
                            fmt.Println(v11, v11, fi15)
                            _ = fi15
                            _ = v16
                        }
                        _ = v14
                    }
                default:
                    for fi17 := 0; fi17 < 8; fi17++ {
                        if false {
                            if true {
                                fmt.Println(p0, p0, p1)
                            }
                            v13 = (v13 & (v13 >> ((2) & 63)))
                            v18 := []int{(fi12 & fi17), 2147483647}
                            v19 := GenS1{F0: 0, F1: (^(p0)), F2: (^(fi12))}
                            _ = v18
                            _ = v19
                        } else {
                            fmt.Println(p1)
                            v11 += v11
                            var v20 uint64 = p0
                            p0 = (p0 + p0)
                            _ = v20
                        }
                        _ = fi17
                    }
                }
                for fi21 := 0; fi21 < 8; fi21++ {
                    if false {
                        v22 := []int{(fi12 << ((284) & 63)), fi12}
                        _ = v22
                    } else {
                        switch fi12 {
                        case 510:
                            v23 := []int{}
                            _ = v23
                        case 127:
                            var v24 bool = false
                            if (1.5 > 0.1) {
                                v25 := []int{(fi12 - fi12), (fi12 ^ (-(fi21))), 910}
                                p1 += (p1 & 0)
                                _ = v25
                            }
                            _ = v24
                        case -638:
                            fmt.Println(fi21)
                        }
                        p1 = p1
                        switch "hello" {
                        case "foo bar":
                            v26 := []int{fi21, fi21, fi21}
                            _ = v26
                        case "":
                            v27 := GenS1{F0: 611, F1: p0, F2: fi12}
                            _ = v27
                        }
                    }
                    v28 := GenS1{F0: 4294967295, F1: p0, F2: (^(fi12))}
                    v29 := []int{fi21, fi21}
                    if (1000 >= (-(v13))) {
                        v30 := GenS1{F0: 4294967295, F1: p0, F2: (-(fi21))}
                        _ = v30
                    }
                    _ = fi21
                    _ = v28
                    _ = v29
                }
                _ = fi12
                _ = v13
            }
        }
        _ = v11
    }
    fmt.Println(p1, p1)
    v31 := []int{}
    var v32 int8 = -1
    _ = p0
    _ = p1
    _ = v10
    _ = v31
    _ = v32
    return (4294967295 >= 927)
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]string)
    var v3 uint16 = 2
    v4 := make(map[string]string)
    v5 := []int{127, 1, 51, 0}
    v5 = append(v5, 929)
    var v33 int64 = 9223372036854775807
    v33 -= (v33 * v33)
    v34 := GenS1{F0: 0, F1: uint64(v3), F2: -145}
    v1 += -1
    fmt.Println(v3, v33, v1)
    if false {
        var v35 int16 = 42
        var v36 uint32 = 540
        _ = v35
        _ = v36
    } else {
        v3 = v3
    }
    if (2147483647 < 2147483647) {
        v3 = v3
    }
    var v37 int64 = 1
    fmt.Println(v1)
    fmt.Println(v3, v37)
    var v38 int8 = 0
    v33 += v33
    v37 = v33
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
