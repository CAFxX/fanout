package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int32; F2 int }
type GenS4 struct { F0 int; F1 uint16 }
func genFn0(p0 int8) uint {
    v6 := GenS1{F0: "", F1: 1000}
    var v7 int32 = 980
    _ = p0
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    var v3 int8 = -93
    var v4 int8 = (v3 / ((100) | 1))
    var v5 int64 = -1
    v5 -= 868
    v2[("foo bar" + "hello")] = -208
    var v8 int64 = v5
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            fmt.Println(v8, v4)
            _ = fi10
        }
        _ = fi9
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            if (-263 >= 0) {
                if (v1 && v1) {
                    switch 9223372036854775807 {
                    case -1:
                        if v1 {
                            fmt.Println(fi12, v8, v4)
                            if false {
                                switch (("a" + "foo bar") + "hello") {
                                case "foo bar":
                                    v8 -= v8
                                }
                            } else {
                                if true {
                                    v13 := []int{}
                                    _ = v13
                                } else {
                                    fmt.Println(v8, v5, v8)
                                    var v14 string = ("hello" + "")
                                    _ = v14
                                }
                                v5 -= v5
                            }
                            var v15 uint16 = 2
                            var v16 uint16 = 781
                            _ = v15
                            _ = v16
                        } else {
                            var v17 int = 9223372036854775807
                            fmt.Println(fi12, v1, v3)
                            v18 := GenS2{F0: (v5 & -586)}
                            _ = v17
                            _ = v18
                        }
                    case 32767:
                        var v19 bool = (3.14159 == 0.0)
                        _ = v19
                    case 289:
                        if v1 {
                            v20 := GenS2{F0: v5}
                            _ = v20
                        }
                    }
                    v4 -= (^(v3))
                    v21 := GenS2{F0: v8}
                    _ = v21
                } else {
                    v22 := GenS2{F0: v8}
                    v23 := GenS3{F0: (!((float32((((-2.25 + 1.5) - 3.14159) * float64(0.1))) * (-2.25 * 3.14159)) <= 3.14159)), F1: 0, F2: (fi12 ^ fi12)}
                    var v24 bool = true
                    _ = v22
                    _ = v23
                    _ = v24
                }
            }
            v4 -= (v4 ^ v4)
            if (!(!(0 == 42))) {
                v25 := GenS4{F0: fi11, F1: 255}
                _ = v25
            }
            _ = fi12
        }
        _ = fi11
    }
    switch 672 {
    case 832:
        fmt.Println(v5)
    case -556:
        delete(v2, "a")
    default:
        v26 := []int{128}
        _ = v26
    }
    var v27 uint16 = 32767
    v27 += v27
    v28 := []int{42, 856, 128}
    v4 = (-(v4))
    for fi29 := 0; fi29 < 11; fi29++ {
        if v1 {
            var v30 int64 = v5
            _ = v30
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
