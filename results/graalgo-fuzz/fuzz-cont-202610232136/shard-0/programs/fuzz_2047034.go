package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint32; F1 int8; F2 uint16 }
func genFn0(p0 int32, p1 uint32) bool {
    p1 = 65535
    var v6 float32 = -2.25
    _ = p0
    _ = p1
    _ = v6
    return true
}
func genFn1(p0 int16, p1 uint8, p2 int8) uint16 {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 string = "a"
        v8 = v8
        fmt.Println(fi7, v8, v8)
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: (-(p2)), F1: (p0 > (p0 % ((p0) | 1)))}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 32767
}

func main() {
    var v1 int64 = 0
    var v2 string = "\n\t"
    var v3 uint = 42
    var v4 float64 = ((((1.5 * float64(0.0)) / float64(((0.1 - (1.5 - 0.1)) + ((((0.0 / 1.5) - 0.0) / (3.14159 - 1e10)) * (((0.1 - (((3.14159 + (0.1 * -2.25)) * ((((((1e10 - ((0.0 + ((0.1 - 0.0) + 1e10)) * -2.25)) - 1e10) * ((-2.25 + float32(((((1e10 * (1.5 - 0.0)) / 3.14159) + 1.5) * (-2.25 * 0.0)))) / 1.0)) / 1e10) * 3.14159) + 1.5)) * 0.0)) * -2.25) + (-2.25 + (3.14159 / 1.0))))))) - 1.5) - 3.14159)
    var v5 int64 = (v1 * v1)
    var v10 uint32 = 1
    switch -2147483648 {
    case -111:
        v4 = v4
    case 32767:
        if false {
            v11 := GenS1{F0: 1, F1: true}
            _ = v11
        }
    case -9223372036854775808:
        for fi12 := 0; fi12 < 16; fi12++ {
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 float32 = 0.0
                var v15 float64 = (-2.25 * v4)
                _ = fi13
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 17; fi16++ {
                if (fi16 != (fi12 >> ((v3) & 63))) {
                    switch v2 {
                    case "hello":
                        switch v2 {
                        case "":
                            fmt.Println(fi12, v1)
                        case "hello":
                            if true {
                                fmt.Println(v10, v4)
                                if true {
                                    v1 = 42
                                }
                                var v17 int16 = 312
                                _ = v17
                            } else {
                                if (fi16 != fi16) {
                                    var v18 int = 2147483647
                                    v1 += (^(v1))
                                    _ = v18
                                }
                            }
                        default:
                            v19 := []int{}
                            _ = v19
                        }
                    case "":
                        var v20 string = v2
                        _ = v20
                    }
                    var v21 uint8 = 127
                    _ = v21
                } else {
                    v2 = ((v2 + v2) + v2)
                    var v22 uint8 = uint8(v3)
                    _ = v22
                }
                _ = fi16
            }
            _ = fi12
        }
    }
    fmt.Println(v1)
    if true {
        fmt.Println(v2, v2)
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        fmt.Println(v3)
        fmt.Println(v4, fi23)
        v2 = ("" + v2)
        _ = fi23
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        if (v2 < v2) {
            v25 := GenS1{F0: -1, F1: (!true)}
            _ = v25
        } else {
            var v26 uint8 = 128
            for fi27 := 0; fi27 < 2; fi27++ {
                v3 = v3
                _ = fi27
            }
            _ = v26
        }
        _ = fi24
    }
    v28 := []int{42, -512, 9223372036854775807}
    v2 = v2
    fmt.Println(v4, v2)
    if false {
        var v29 uint64 = 1
        _ = v29
    }
    var v30 uint64 = 2
    v3 = (v3 * v3)
    fmt.Println(v10, v5, v1)
    var v31 string = ((v2 + "") + v2)
    for fi32 := 0; fi32 < 18; fi32++ {
        v1 = v1
        for fi33 := 0; fi33 < 2; fi33++ {
            if len(v28) >= 2 {
                v28 = v28[1:len(v28)]
            }
            _ = fi33
        }
        _ = fi32
    }
    v34 := GenS2{F0: (v10 << (((^(v3))) & 63)), F1: 100, F2: 0}
    v31 = (v2 + "foo bar")
    v35 := GenS1{F0: 1, F1: ((^(v3)) != v3)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
