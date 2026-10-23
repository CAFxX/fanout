package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint64; F1 int; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 string; F2 uint8 }
func genFn0(p0 int8, p1 uint) uint16 {
    fmt.Println(p1)
    v6 := []int{0, 9223372036854775807}
    var v7 int32 = 1
    var v8 uint16 = 861
    if (169 != 341) {
        v9 := GenS1{F0: 42}
        if (-60 > 32767) {
            p1 += 42
            if (((true || (!true)) && (874 >= 1)) && ((float32(-2.25) >= ((0.0 - 0.0) - 1e10)) && true)) {
                v10 := GenS1{F0: 42}
                _ = v10
            } else {
                switch "\n\t" {
                case "hello":
                    var v11 uint8 = 0
                    _ = v11
                default:
                    var v12 int8 = p0
                    _ = v12
                }
            }
            v7 -= v7
        }
        var v13 int = -9223372036854775808
        _ = v13
        _ = v9
    } else {
        v14 := GenS1{F0: 128}
        p1 = 18446744073709551615
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint64, p1 int8, p2 string) int64 {
    if true {
        fmt.Println(p1, p1)
        var v15 int8 = -71
        if false {
            for fi16 := 0; fi16 < 16; fi16++ {
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 uint8 = 0
                    fmt.Println(v18, fi17, p1)
                    fmt.Println(fi16, fi17)
                    _ = fi17
                    _ = v18
                }
                v19 := []int{(-(fi16))}
                v20 := GenS2{F0: -9223372036854775808, F1: v15}
                _ = fi16
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 12; fi21++ {
                if true {
                    if (true && true) {
                        if false {
                            if (-1 < fi21) {
                                if (128 > 42) {
                                    for fi22 := 0; fi22 < 19; fi22++ {
                                        var v23 float64 = 0.1
                                        _ = fi22
                                        _ = v23
                                    }
                                    if true {
                                        v15 = (v15 + v15)
                                        fmt.Println(v15, p1)
                                    }
                                    v24 := GenS2{F0: 9223372036854775807, F1: (v15 % ((-128) | 1))}
                                    _ = v24
                                }
                            }
                            p1 = p1
                            var v25 bool = true
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 7; fi26++ {
                                var v27 uint64 = (p0 + (p0 << ((65535) & 63)))
                                _ = fi26
                                _ = v27
                            }
                        }
                        var v28 float64 = ((1.5 - (0.1 - ((((0.0 / (-2.25 * 1e10)) + 1e10) - 0.0) / -2.25))) / 3.14159)
                        _ = v28
                    }
                    v29 := []int{}
                    _ = v29
                }
                v30 := []int{(fi21 * (fi21 - fi21)), fi21, fi21}
                _ = fi21
                _ = v30
            }
        }
        _ = v15
    } else {
        switch (p2 + p2) {
        case "hello":
            p0 = (p0 >> ((0) & 63))
        case "foo bar":
            v31 := GenS1{F0: 127}
            _ = v31
        }
        fmt.Println(p2, p0)
    }
    p1 = (p1 - (p1 * -1))
    var v32 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    return -1
}

func main() {
    var v1 float32 = 3.14159
    var v2 float32 = (0.1 * 1.5)
    v3 := make(map[string]int)
    var v4 int8 = 0
    var v5 int64 = 9223372036854775807
    v1 = v2
    var v33 int = 824
    v5 = v5
    v34 := []int{}
    v1 = v2
    if false {
        v33 = v33
        if true {
            v35 := GenS2{F0: genFn1(462, v4, "\n\t"), F1: v4}
            var v36 float32 = (1e10 / v2)
            if (false || ((42 == 0) && false)) {
                _ = len(v3)
                var v37 int64 = (v5 - genFn1(42, v4, "hello"))
                _ = v37
            }
            _ = v35
            _ = v36
        }
    }
    fmt.Println(v1, v2)
    var v38 float64 = 0.1
    if len(v34) >= 2 {
        v34 = v34[1:len(v34)]
    }
    if len(v34) > 0 {
        v34[0] = 32768
    }
    v39 := GenS3{F0: 42, F1: (-(v33)), F2: ("a" + "\n\t")}
    v40 := GenS3{F0: 42, F1: 42, F2: "foo bar"}
    fmt.Println(v38)
    v41 := GenS4{F0: genFn0((v4 - 127), 1)}
    v42 := GenS5{F0: (v33 + v33), F1: ("foo bar" + "\n\t"), F2: 1}
    fmt.Println(v4, v1, v33)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
