package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
func genFn0(p0 string, p1 bool, p2 int8) bool {
    p0 = "foo bar"
    p1 = p1
    for fi4 := 0; fi4 < 11; fi4++ {
        p0 = p0
        for fi5 := 0; fi5 < 3; fi5++ {
            p2 = 1
            var v6 int64 = 0
            p2 -= p2
            _ = fi5
            _ = v6
        }
        var v7 uint8 = 128
        _ = fi4
        _ = v7
    }
    p0 = ""
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return ((((p0 + ((p0 + ((((("foo bar" + p0) + "\n\t") + p0) + "a") + p0)) + (((p0 + "hello") + (("a" + p0) + (p0 + "\n\t"))) + "a"))) + p0) + (p0 + p0)) != "hello")
}

func main() {
    var v1 bool = false
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    v1 = v1
    var v8 int32 = -2147483648
    fmt.Println(v1, v8)
    var v9 bool = false
    v10 := []int{-662, 316, -2147483648, 4294967295}
    var v11 int8 = -128
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(v9)
        _ = fi12
    }
    switch 1 {
    case -9223372036854775808:
        if v9 {
            for fi13 := 0; fi13 < 12; fi13++ {
                fmt.Println(fi13, v8, v9)
                _ = fi13
            }
            fmt.Println(v1)
            if len(v10) > 0 {
                v10[0] = -715
            }
        } else {
            var v14 bool = v1
            v15 := GenS1{F0: 1, F1: -473}
            fmt.Println(v14, v8)
            _ = v14
            _ = v15
        }
    case 2147483647:
        var v16 int = 128
        _ = v16
    case -838:
        for fi17 := 0; fi17 < 11; fi17++ {
            if v9 {
                v18 := v2[("foo bar" + ("foo bar" + "\n\t"))]
                _ = v18
            }
            fmt.Println(v11)
            v9 = (1 <= (v8 & (v8 * v8)))
            _ = fi17
        }
    default:
        v1 = ((0.0 + 3.14159) <= (0.1 * (1e10 + (((1.5 + 3.14159) / 1.5) * float64(0.1)))))
    }
    switch "a" {
    case "foo bar":
        if (v1 || (1000 > 0)) {
            var v19 float32 = 3.14159
            var v20 float64 = -2.25
            if v1 {
                for fi21 := 0; fi21 < 2; fi21++ {
                    v8 += -2147483648
                    switch (fi21 ^ (fi21 << ((42) & 63))) {
                    case -1:
                        _ = len(v3)
                    }
                    fmt.Println(v1, v8)
                    _ = fi21
                }
                if v1 {
                    for fi22 := 0; fi22 < 9; fi22++ {
                        _ = len(v10)
                        v23 := GenS1{F0: 26, F1: fi22}
                        _ = fi22
                        _ = v23
                    }
                    _ = len(v10)
                    v1 = v1
                }
                v10 = append(v10, 2147483647)
            }
            _ = v19
            _ = v20
        } else {
            if len(v10) > 0 {
                v10[0] = 127
            }
            switch 0 {
            case 1:
                fmt.Println(v8, v8, v8)
            }
        }
    default:
        if (!false) {
            delete(v2, "\n\t")
        }
    }
    v8 = (v8 / ((v8) | 1))
    v8 = v8
    switch 32767 {
    case 929:
        v9 = (1 != v11)
    case 790:
        for fi24 := 0; fi24 < 13; fi24++ {
            if v1 {
                if (!("" >= "\n\t")) {
                    if (v9 && v1) {
                        v25 := GenS1{F0: 767, F1: fi24}
                        v1 = (float64((0.1 * 1e10)) <= 3.14159)
                        for fi26 := 0; fi26 < 17; fi26++ {
                            v11 = 127
                            if (-105 > 42) {
                                v8 = v8
                                v27 := v2["\n\t"]
                                if (!((!true) && false)) {
                                    var v28 int16 = 42
                                    _ = v28
                                } else {
                                    var v29 uint8 = 42
                                    var v30 uint = 1
                                    _ = v29
                                    _ = v30
                                }
                                _ = v27
                            } else {
                                v9 = (311 <= -679)
                                v31 := v3[127]
                                _ = v31
                            }
                            _ = fi26
                        }
                        _ = v25
                    }
                    fmt.Println(v1)
                    v11 = v11
                }
            }
            _ = fi24
        }
    default:
        var v32 uint16 = 50
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
