package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint; F2 int16 }
func genFn0(p0 int, p1 int16) int8 {
    var v6 bool = (42 == 2)
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(p0, v6)
            _ = fi8
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            if ((("foo bar" + ("a" + ("foo bar" + "\n\t"))) + "") == "a") {
                switch "hello" {
                case "foo bar":
                    v6 = false
                case "a":
                    switch "hello" {
                    case "\n\t":
                        v10 := GenS1{F0: 65535}
                        _ = v10
                    case "hello":
                        var v11 uint = 1
                        _ = v11
                    default:
                        fmt.Println(v6, p0)
                    }
                case "hello":
                    if (-2.25 != (1e10 + 3.14159)) {
                        var v12 string = "a"
                        if v6 {
                            p0 = -550
                            if (!(v6 && false)) {
                                var v13 int = fi7
                                fmt.Println(v12, p1)
                                fmt.Println(v6, v13)
                                _ = v13
                            } else {
                                v12 = v12
                            }
                        }
                        _ = v12
                    }
                default:
                    v14 := GenS1{F0: 18446744073709551615}
                    _ = v14
                }
                p0 = (-(fi7))
                v15 := GenS1{F0: 734}
                v16 := []int{}
                var v17 uint32 = 42
                _ = v15
                _ = v16
                _ = v17
            } else {
                if (1 < 0) {
                    p1 -= p1
                    var v18 bool = ((p1 == p1) || false)
                    _ = v18
                }
                var v19 int = fi9
                if (373 == -2147483648) {
                    v6 = v6
                    v20 := []int{(p0 - fi7)}
                    _ = v20
                }
                p1 = (-(p1))
                v19 = (v19 >> ((1) & 63))
                _ = v19
            }
            v21 := GenS1{F0: 1}
            _ = fi9
            _ = v21
        }
        v6 = v6
        _ = fi7
    }
    v22 := []int{p0, -1, p0, p0}
    var v23 int16 = p1
    var v24 uint16 = 73
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    _ = v24
    _ = v6
    return 63
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -2147483648
    var v3 string = "foo bar"
    v4 := make(map[string]uint32)
    v5 := []int{-9223372036854775808}
    var v25 bool = false
    v2 = (v2 % ((v2) | 1))
    switch -2147483648 {
    case 9223372036854775807:
        var v26 string = v3
        _ = v26
    case 216:
        var v27 uint8 = 255
        _ = v27
    }
    v2 = v2
    var v28 int8 = 79
    v29 := GenS2{F0: 201}
    if v25 {
        for fi30 := 0; fi30 < 13; fi30++ {
            if false {
                fmt.Println(v3, fi30)
            }
            v28 = v28
            v3 = (v3 + (v3 + v3))
            _ = fi30
        }
        var v31 string = ("foo bar" + v3)
        var v32 float32 = 1e10
        _ = v31
        _ = v32
    }
    v28 += -1
    if v25 {
        if v25 {
            v2 -= (^(v2))
        } else {
            v2 = v2
        }
    } else {
        fmt.Println(v3)
    }
    if true {
        v33 := GenS2{F0: 22}
        switch 4294967295 {
        case 0:
            fmt.Println(v2, v3, v25)
        case 391:
            var v34 float64 = (1e10 * (-2.25 + -2.25))
            _ = v34
        }
        _ = len(v5)
        _ = v33
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        var v36 uint32 = 4294967295
        switch v3 {
        case "\n\t":
            for fi37 := 0; fi37 < 5; fi37++ {
                v36 = v36
                _ = fi37
            }
        case "a":
            fmt.Println(v36)
        default:
            if v25 {
                if v25 {
                    v38 := v4[(v3 + "a")]
                    var v39 int = -229
                    _ = v38
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 14; fi40++ {
                        var v41 float64 = -2.25
                        _ = fi40
                        _ = v41
                    }
                }
            }
        }
        for fi42 := 0; fi42 < 12; fi42++ {
            v43 := GenS3{F0: (v2 << ((0) & 63)), F1: 4294967295, F2: 1}
            v28 += v28
            _ = fi42
            _ = v43
        }
        _ = fi35
        _ = v36
    }
    var v44 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
