package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int32 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint32 }
func genFn0(p0 int, p1 uint16, p2 int8) string {
    v7 := []int{}
    var v8 int64 = 42
    if false {
        switch "" {
        case "foo bar":
            fmt.Println(p1, v8, p1)
        case "a":
            if true {
                p0 -= 990
                fmt.Println(v8, p1)
            }
        case "hello":
            fmt.Println(p2)
        default:
            if false {
                switch "a" {
                case "\n\t":
                    v8 = v8
                default:
                    switch (p0 - p0) {
                    case -40:
                        for fi9 := 0; fi9 < 6; fi9++ {
                            var v10 int8 = p2
                            _ = fi9
                            _ = v10
                        }
                    case -1:
                        _ = len(v7)
                    case 2147483647:
                        v8 += 0
                    }
                }
                p2 -= p2
            } else {
                p1 += 32767
                if (!(true || true)) {
                    for fi11 := 0; fi11 < 15; fi11++ {
                        v8 -= (v8 << ((18446744073709551615) & 63))
                        if false {
                            for fi12 := 0; fi12 < 15; fi12++ {
                                var v13 int8 = -128
                                if (!true) {
                                    fmt.Println(v8)
                                    if len(v7) > 0 {
                                        v7[0] = fi12
                                    }
                                } else {
                                    if len(v7) > 0 {
                                        v7[0] = fi12
                                    }
                                }
                                _ = fi12
                                _ = v13
                            }
                            v14 := GenS1{F0: 19, F1: 0, F2: 0}
                            p2 = p2
                            _ = v14
                        } else {
                            v8 = (v8 / ((v8) | 1))
                            var v15 float64 = ((3.14159 / -2.25) - float64(float32(((3.14159 / (0.1 * 3.14159)) / 1e10))))
                            v16 := GenS1{F0: 2, F1: fi11, F2: -2147483648}
                            _ = v15
                            _ = v16
                        }
                        v17 := GenS1{F0: 852, F1: (p0 - (fi11 & (fi11 & (p0 << ((4294967295) & 63))))), F2: 0}
                        _ = fi11
                        _ = v17
                    }
                    p2 = p2
                    _ = len(v7)
                }
            }
        }
        if false {
            if false {
                var v18 bool = (false && false)
                v18 = ((-1 == -1) && (-1 >= 32767))
                _ = v18
            }
        } else {
            var v19 int64 = v8
            _ = v19
        }
    }
    var v20 uint = 2
    v21 := GenS1{F0: (v20 >> ((221) & 63)), F1: (p0 >> ((292) & 63)), F2: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    _ = v7
    _ = v8
    return "a"
}
func genFn1(p0 uint, p1 string, p2 int64) uint8 {
    v22 := GenS1{F0: (p0 % ((18446744073709551615) | 1)), F1: 127, F2: -2147483648}
    var v23 int16 = 32767
    var v24 string = ""
    switch p1 {
    case "hello":
        switch p1 {
        case "foo bar":
            if ((4294967295 <= 154) || (v24 != ((v24 + ("foo bar" + p1)) + p1))) {
                fmt.Println(p2)
                v25 := []int{127, 0, 0, 1}
                _ = v25
            } else {
                v26 := GenS2{F0: 2, F1: (^(v23)), F2: 186}
                v27 := GenS2{F0: 0, F1: v23, F2: 629}
                _ = v26
                _ = v27
            }
        }
    }
    v24 = v24
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    _ = v24
    return 128
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{-1, 327, 1, 128}
    var v3 bool = ((587 <= 645) && true)
    var v4 float64 = 1.5
    var v5 int16 = -331
    v6 := make(map[int]bool)
    v2 = append(v2, 0)
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 float32 = float32(1.5)
        _ = fi28
        _ = v29
    }
    v4 -= -2.25
    v3 = v3
    fmt.Println(v3, v5, v3)
    for fi30 := 0; fi30 < 18; fi30++ {
        v4 = v4
        v4 += ((3.14159 / v4) * v4)
        var v31 uint8 = 1
        _ = fi30
        _ = v31
    }
    v4 = v4
    var v32 float64 = v4
    var v33 string = "hello"
    var v34 float64 = (float64(1.5) / 1e10)
    if v3 {
        for fi35 := 0; fi35 < 19; fi35++ {
            fmt.Println(v34)
            v34 -= 1.5
            _ = fi35
        }
        if v3 {
            v33 = v33
            v32 -= 3.14159
            for fi36 := 0; fi36 < 1; fi36++ {
                v3 = true
                v37 := GenS2{F0: genFn1(4294967295, ((v33 + "a") + v33), -1), F1: (-(v5)), F2: 65535}
                _ = fi36
                _ = v37
            }
        } else {
            var v38 int8 = -125
            v32 += float64((3.14159 / (0.1 * float32(v34))))
            v34 = (0.0 - v32)
            _ = v38
        }
    }
    v32 += v4
    v39 := v1[v33]
    v40 := v6[127]
    var v41 int16 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
