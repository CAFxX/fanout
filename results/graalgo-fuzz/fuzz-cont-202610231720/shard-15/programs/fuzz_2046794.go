package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int8 }
type GenS2 struct { F0 int8; F1 int32 }
func genFn0(p0 int, p1 bool, p2 int8) uint32 {
    if (0.0 >= (1e10 + float32(((0.0 * 0.1) / 1.5)))) {
        v6 := []int{p0}
        p1 = (42 > 134)
        for fi7 := 0; fi7 < 14; fi7++ {
            _ = len(v6)
            var v8 uint = 255
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        if (-758 > -32768) {
            p2 = 100
            v9 := []int{128, (p0 ^ p0)}
            if (!((p0 + int(p2)) == (p0 + p0))) {
                var v10 uint = 173
                var v11 int = p0
                _ = v10
                _ = v11
            } else {
                if false {
                    p0 = (p0 ^ (-(p0)))
                    if false {
                        fmt.Println(p2)
                        if p1 {
                            if p1 {
                                var v12 string = ""
                                if (!false) {
                                    v9 = append(v9, p0)
                                    fmt.Println(p0)
                                    var v13 uint = 2
                                    _ = v13
                                } else {
                                    if p1 {
                                        switch ((("hello" + v12) + (("a" + "") + (("\n\t" + ((v12 + ("foo bar" + (("hello" + "") + (v12 + "")))) + v12)) + ("a" + (v12 + (v12 + v12)))))) + v12) {
                                        case "hello":
                                            var v14 int = 203
                                            _ = v14
                                        default:
                                            _ = len(v9)
                                        }
                                        v12 = v12
                                    }
                                }
                                _ = v12
                            }
                            var v15 uint16 = 909
                            _ = v15
                        }
                    } else {
                        p2 = p2
                    }
                } else {
                    var v16 float32 = 0.0
                    _ = v16
                }
                p1 = p1
                fmt.Println(p1)
            }
            _ = v9
        }
    }
    if (!p1) {
        fmt.Println(p1, p2)
        if ("hello" < (("\n\t" + ("foo bar" + ("" + ("hello" + "")))) + "a")) {
            fmt.Println(p2, p0)
            fmt.Println(p1, p1, p0)
        }
        v17 := GenS1{F0: 42, F1: 65535, F2: -128}
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return 799
}
func genFn1() int8 {
    var v18 int64 = 42
    fmt.Println(v18)
    _ = v18
    return 42
}

func main() {
    v1 := []uint32{732}
    var v2 int32 = -1
    var v3 int64 = 0
    v4 := []int{9223372036854775807}
    var v5 int16 = 1
    fmt.Println(v2, v3)
    v19 := GenS1{F0: (v2 + (v2 + v2)), F1: genFn0(2147483647, true, 100), F2: 127}
    for fi20 := 0; fi20 < 19; fi20++ {
        v3 -= 0
        v21 := GenS1{F0: v2, F1: genFn0(128, true, -70), F2: -98}
        _ = fi20
        _ = v21
    }
    var v22 int32 = v2
    for fi23 := 0; fi23 < 7; fi23++ {
        if len(v4) > 0 {
            v4[0] = fi23
        }
        if (!false) {
            v2 += (-(v2))
            if true {
                v22 = (v2 / (((v22 + v2)) | 1))
                for fi24 := 0; fi24 < 19; fi24++ {
                    switch 9223372036854775807 {
                    case 42:
                        v4 = append(v4, fi24)
                    case -1:
                        if true {
                            fmt.Println(fi24, fi24, v22)
                            fmt.Println(v5, v3, v2)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    default:
                        v25 := GenS1{F0: (v22 | 1), F1: 0, F2: 0}
                        _ = v25
                    }
                    _ = fi24
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            var v26 int16 = 42
            _ = v26
        }
        var v27 uint16 = 0
        _ = fi23
        _ = v27
    }
    v5 -= 313
    var v28 bool = (!false)
    for fi29 := 0; fi29 < 18; fi29++ {
        if (float32((1.5 + float64(0.1))) != ((3.14159 / 0.1) + 0.1)) {
            var v30 bool = false
            var v31 float32 = (-2.25 / 1.0)
            for fi32 := 0; fi32 < 5; fi32++ {
                var v33 string = "foo bar"
                if (!v28) {
                    var v34 int64 = v3
                    v35 := GenS2{F0: 1, F1: 705}
                    _ = v34
                    _ = v35
                }
                var v36 uint32 = genFn0((fi32 ^ (fi32 ^ fi32)), true, 42)
                _ = fi32
                _ = v33
                _ = v36
            }
            _ = v30
            _ = v31
        }
        v3 -= v3
        v28 = true
        _ = fi29
    }
    var v37 uint64 = 4294967295
    var v38 int64 = (-(v3))
    v22 = v2
    if (0 < 484) {
        for fi39 := 0; fi39 < 14; fi39++ {
            var v40 int8 = 100
            v41 := GenS2{F0: v40, F1: v22}
            v5 -= v5
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    _ = len(v1)
    fmt.Println(v28, v22, v37)
    v42 := GenS2{F0: 0, F1: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v5
}
