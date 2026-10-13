package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 bool }
type GenS2 struct { F0 int8; F1 int32; F2 uint32 }
func genFn0(p0 bool, p1 string, p2 uint32) uint16 {
    v7 := []int{127, 42}
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 454
}
func genFn1(p0 string, p1 uint16, p2 uint) int {
    v8 := []int{32767, 24, -9223372036854775808, -374}
    switch p0 {
    case "\n\t":
        var v9 int64 = 1
        _ = v9
    case "hello":
        if false {
            if (false && ((int64(p2) <= 42) || (-54 <= 2147483647))) {
                p2 -= p2
            }
            p2 = (p2 >> (((^(p2))) & 63))
            if (!(!(!(1.5 < float64((0.1 + (-2.25 - -2.25))))))) {
                for fi10 := 0; fi10 < 4; fi10++ {
                    p2 = p2
                    if true {
                        v11 := GenS1{F0: -40, F1: (fi10 * (fi10 * fi10)), F2: true}
                        fmt.Println(p1, p0, p2)
                        _ = v11
                    } else {
                        if (p2 == (p2 * p2)) {
                            p0 = (p0 + ("" + "foo bar"))
                            fmt.Println(p1, p1)
                            switch fi10 {
                            case -604:
                                fmt.Println(fi10, p1, p1)
                            case 0:
                                var v12 int64 = -9223372036854775808
                                _ = v12
                            }
                        } else {
                            p0 = "a"
                        }
                        fmt.Println(fi10, fi10)
                        v13 := GenS1{F0: -56, F1: fi10, F2: true}
                        _ = v13
                    }
                    switch "\n\t" {
                    case "":
                        var v14 float32 = 1e10
                        _ = v14
                    case "foo bar":
                        var v15 string = ""
                        _ = v15
                    case "a":
                        var v16 bool = false
                        _ = v16
                    }
                    _ = fi10
                }
                fmt.Println(p1)
                fmt.Println(p1, p2, p0)
            } else {
                v17 := GenS1{F0: 42, F1: 4294967295, F2: true}
                _ = v17
            }
        }
    case "":
        for fi18 := 0; fi18 < 8; fi18++ {
            p1 -= (^(p1))
            var v19 uint16 = genFn0(true, "\n\t", 2)
            for fi20 := 0; fi20 < 14; fi20++ {
                v21 := GenS2{F0: -128, F1: 1, F2: 1}
                v8 = append(v8, (fi18 ^ (fi20 % ((-687) | 1))))
                p2 = p2
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
    default:
        if true {
            p2 = p2
            fmt.Println(p0, p0, p1)
        } else {
            p1 += (^(p1))
            p0 = "foo bar"
            var v22 float64 = (3.14159 * 1.5)
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 9223372036854775807
    v3 := []string{"a"}
    var v4 uint16 = 1
    v5 := []int{-9223372036854775808, 1, 0}
    var v6 uint16 = v4
    v6 += 65535
    var v23 uint = 0
    v24 := GenS2{F0: 42, F1: 0, F2: 515}
    var v25 float64 = float64(0.1)
    if false {
        v25 -= (v25 * v25)
    } else {
        v26 := GenS1{F0: -128, F1: 2147483647, F2: true}
        for fi27 := 0; fi27 < 7; fi27++ {
            if (false || true) {
                v23 += v23
            } else {
                delete(v1, "hello")
                _ = len(v5)
                var v28 uint32 = 1
                _ = v28
            }
            switch "\n\t" {
            case "a":
                var v29 int16 = -201
                _ = v29
            }
            _ = fi27
        }
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 string = "\n\t"
            if false {
                var v32 int16 = -1
                fmt.Println(v31, v31, v2)
                _ = v32
            }
            _ = fi30
            _ = v31
        }
        _ = v26
    }
    var v33 bool = (!((v2 >> ((v23) & 63)) == (v2 | (v2 | (v2 / ((1) | 1))))))
    v34 := GenS1{F0: 1, F1: 909, F2: (42 == -2147483648)}
    for fi35 := 0; fi35 < 5; fi35++ {
        v6 = (v4 + 1)
        var v36 int32 = 1
        _ = fi35
        _ = v36
    }
    var v37 uint32 = 501
    var v38 float64 = (v25 + 1e10)
    for fi39 := 0; fi39 < 9; fi39++ {
        var v40 string = "hello"
        _ = fi39
        _ = v40
    }
    v4 = (v6 | genFn0((!(!(!false))), (("a" + "a") + ("a" + "")), 42))
    var v41 uint32 = (v37 % (((v37 + (v37 | (v37 & 2147483647)))) | 1))
    v33 = (v33 && ((!false) || v33))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
