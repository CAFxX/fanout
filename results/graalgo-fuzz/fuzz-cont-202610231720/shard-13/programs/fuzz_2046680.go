package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string, p1 int) int16 {
    p1 = (p1 ^ 1)
    if true {
        v7 := []int{p1, 2147483647, 273}
        p1 -= p1
        if (!false) {
            for fi8 := 0; fi8 < 4; fi8++ {
                _ = len(v7)
                _ = fi8
            }
        }
        _ = v7
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 9; fi10++ {
            p0 = "hello"
            fmt.Println(fi9)
            var v11 uint = 974
            _ = fi10
            _ = v11
        }
        p0 = "a"
        _ = fi9
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(p1)
        v13 := []int{(fi12 & 1)}
        if true {
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 int8 = 115
                v16 := GenS1{F0: 0}
                p0 = (("a" + p0) + ("foo bar" + "foo bar"))
                if true {
                    p0 = p0
                    if true {
                        var v17 uint8 = 121
                        if (!true) {
                            var v18 uint32 = 1
                            var v19 uint = 4294967295
                            v20 := GenS1{F0: (v17 | v17)}
                            _ = v18
                            _ = v19
                            _ = v20
                        } else {
                            if false {
                                fmt.Println(fi12)
                                if len(v13) > 0 {
                                    v13[0] = fi12
                                }
                            } else {
                                var v21 int32 = 633
                                _ = len(v13)
                                v13 = append(v13, fi12)
                                _ = v21
                            }
                            fmt.Println(fi12, fi14, fi12)
                        }
                        _ = v17
                    } else {
                        v13 = append(v13, (p1 >> ((0) & 63)))
                        v13 = append(v13, fi14)
                        var v22 int32 = -2147483648
                        _ = v22
                    }
                } else {
                    v13 = append(v13, 4294967295)
                    if (float64(0.1) < (0.1 * 1e10)) {
                        var v23 int64 = -9223372036854775808
                        v24 := GenS1{F0: 2}
                        _ = v23
                        _ = v24
                    }
                    var v25 int32 = 1
                    _ = v25
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            fmt.Println(p0, fi12, fi12)
            for fi26 := 0; fi26 < 4; fi26++ {
                switch "" {
                case "a":
                    if true {
                        fmt.Println(fi26, p1, p0)
                        p0 = p0
                    } else {
                        p1 = (fi26 % ((-556) | 1))
                        p1 = p1
                        v13 = append(v13, fi12)
                    }
                }
                var v27 bool = (false || false)
                _ = fi26
                _ = v27
            }
        } else {
            v28 := GenS1{F0: 255}
            _ = v28
        }
        _ = fi12
        _ = v13
    }
    p0 = ((p0 + "foo bar") + ((("foo bar" + p0) + (("foo bar" + (p0 + "\n\t")) + "a")) + (p0 + p0)))
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = v1
    v3 := []uint32{27}
    v4 := []int{128, 573, 235, 2147483647}
    var v5 bool = (1e10 < 1e10)
    v6 := make(map[string]bool)
    v1 = v1
    v5 = (v5 && false)
    for fi29 := 0; fi29 < 7; fi29++ {
        fmt.Println(v5, fi29, v5)
        fmt.Println(v2, fi29, v2)
        var v30 string = ""
        _ = fi29
        _ = v30
    }
    fmt.Println(v5, v5, v5)
    v31 := GenS2{F0: (v1 * (v2 ^ 172))}
    _ = len(v4)
    v32 := GenS1{F0: 6}
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
