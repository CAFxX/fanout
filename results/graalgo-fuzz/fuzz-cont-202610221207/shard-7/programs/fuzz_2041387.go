package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int32; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 int16) int64 {
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := []int{(fi7 << ((65535) & 63))}
        p0 = (p0 | p0)
        _ = len(v8)
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    return 42
}
func genFn1(p0 int32, p1 int, p2 string) int64 {
    v9 := []int{p1}
    p2 = p2
    p1 += (p1 / (((-(p1))) | 1))
    v10 := GenS1{F0: -17}
    switch (p2 + p2) {
    case "\n\t":
        v11 := GenS1{F0: p0}
        _ = v11
    case "a":
        var v12 uint32 = 949
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return int64(p1)
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = ""
    v3 := make(map[int]int)
    var v4 uint64 = 831
    v5 := []int{-621}
    var v6 int8 = 16
    if false {
        if ((!(false || true)) && true) {
            v13 := GenS1{F0: -2147483648}
            v6 = v6
            _ = v13
        }
        for fi14 := 0; fi14 < 9; fi14++ {
            if ((!false) && ((v4 >> ((621) & 63)) == v4)) {
                v5 = append(v5, fi14)
            }
            _ = fi14
        }
    } else {
        v15 := v1[v2]
        if (!false) {
            var v16 int8 = (v6 % ((100) | 1))
            _ = v16
        } else {
            fmt.Println(v6)
        }
        _ = v15
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 uint = 1
        if false {
            v19 := GenS2{F0: 248, F1: (true || true)}
            if true {
                for fi20 := 0; fi20 < 4; fi20++ {
                    if true {
                        v21 := GenS3{F0: v6, F1: ("hello" + v2)}
                        var v22 uint64 = 598
                        var v23 float32 = ((0.1 / (3.14159 + float32(1.5))) * (((float32(-2.25) + float32(((((1e10 * (-2.25 / (-2.25 / (3.14159 + ((float64(0.1) / 3.14159) * float64(1.5)))))) / 0.1) / 1.5) - (1e10 + (0.0 - 1.5))))) - float32((float64((((float32(3.14159) / float32(-2.25)) + 3.14159) * 3.14159)) + float64(0.0)))) / 1.0))
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        if (428 >= v4) {
                            v24 := GenS2{F0: (^(fi20)), F1: (v2 < v2)}
                            v1["a"] = (v4 < (^(v4)))
                            if true {
                                var v25 float32 = (1e10 / 3.14159)
                                v6 += v6
                                var v26 float32 = (-2.25 / 1.0)
                                _ = v25
                                _ = v26
                            }
                            _ = v24
                        }
                        v27 := GenS4{F0: 818, F1: 766, F2: (^(v4))}
                        v2 = v2
                        fmt.Println(fi20)
                        _ = v27
                    }
                    var v28 int = (fi17 & fi17)
                    _ = fi20
                    _ = v28
                }
                if len(v5) > 0 {
                    v5[0] = (^(fi17))
                }
                v18 = 573
            } else {
                fmt.Println(v2, v2)
                fmt.Println(v2, fi17, v4)
            }
            v6 += 100
            _ = v19
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v18 = (v18 >> (((^(v18))) & 63))
        }
        if false {
            v18 -= v18
            switch v2 {
            case "":
                fmt.Println(v18, v4, v4)
            }
        }
        _ = fi17
        _ = v18
    }
    v29 := GenS2{F0: 42, F1: (0 != 255)}
    var v30 float32 = 0.0
    var v31 uint8 = 127
    var v32 int = -9223372036854775808
    v33 := GenS5{F0: 42}
    if true {
        var v34 uint64 = 18446744073709551615
        _ = v34
    } else {
        v35 := GenS5{F0: 2147483647}
        switch v2 {
        case "hello":
            v32 = (v32 & (v32 * v32))
        case "a":
            v32 = v32
        default:
            for fi36 := 0; fi36 < 12; fi36++ {
                var v37 int32 = -2147483648
                _ = len(v3)
                v38 := GenS6{F0: 4294967295}
                _ = fi36
                _ = v37
                _ = v38
            }
        }
        _ = v35
    }
    if false {
        fmt.Println(v32, v2, v30)
        if len(v5) > 0 {
            v5[0] = v32
        }
    } else {
        v32 = v32
        var v39 uint64 = 0
        for fi40 := 0; fi40 < 16; fi40++ {
            switch v2 {
            case "foo bar":
                for fi41 := 0; fi41 < 14; fi41++ {
                    var v42 string = v2
                    switch (v32 % ((fi40) | 1)) {
                    case -667:
                        v31 -= (v31 * v31)
                    case -363:
                        var v43 uint32 = 822
                        _ = v43
                    default:
                        v32 -= (fi40 % ((fi40) | 1))
                        if false {
                            _ = len(v1)
                        } else {
                            v42 = ""
                        }
                    }
                    if true {
                        v31 += 1
                    } else {
                        var v44 int64 = -9223372036854775808
                        v39 = v39
                        _ = v44
                    }
                    _ = fi41
                    _ = v42
                }
            default:
                v45 := v1[(("\n\t" + v2) + v2)]
                _ = v45
            }
            _ = fi40
        }
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
