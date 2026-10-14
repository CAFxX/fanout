package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 string }
func genFn0(p0 bool, p1 uint32) int8 {
    if (-9223372036854775808 <= -1) {
        p0 = false
    }
    v7 := []int{128, -9223372036854775808, -304}
    var v8 int64 = 680
    var v9 uint32 = (^(p1))
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int64 = (^(v8))
            v8 = 0
            switch "a" {
            case "foo bar":
                if false {
                    switch (((("foo bar" + "hello") + ("foo bar" + "a")) + "hello") + "foo bar") {
                    case "hello":
                        if p0 {
                            v13 := GenS1{F0: (4294967295 == 704), F1: 1}
                            v12 = v8
                            _ = v13
                        } else {
                            var v14 uint16 = 905
                            v14 += 255
                            _ = v14
                        }
                    case "":
                        if ((!(p0 && true)) || (uint8(p1) > 0)) {
                            var v15 uint8 = 127
                            v12 = (v12 + v12)
                            _ = v15
                        } else {
                            var v16 bool = p0
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            var v17 int16 = 1000
                            _ = v16
                            _ = v17
                        }
                    }
                } else {
                    if (0 > 42) {
                        fmt.Println(fi11, v12, p1)
                    } else {
                        var v18 uint = 42
                        if ((v9 + (^(p1))) == 65535) {
                            if len(v7) > 0 {
                                v7[0] = int(v12)
                            }
                            v7 = append(v7, 128)
                        } else {
                            p1 -= (v9 % ((v9) | 1))
                        }
                        var v19 bool = true
                        _ = v18
                        _ = v19
                    }
                    var v20 int = fi10
                    _ = v20
                }
            case "hello":
                v9 += v9
            default:
                v21 := GenS1{F0: p0, F1: 4294967295}
                _ = v21
            }
            _ = fi11
            _ = v12
        }
        v7 = append(v7, -31)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 int, p1 int8, p2 uint64) uint32 {
    switch (("" + ("a" + "hello")) + "hello") {
    case "\n\t":
        p0 -= 127
    case "hello":
        v22 := []int{p0, p0}
        _ = v22
    }
    var v23 int = p0
    v24 := GenS1{F0: (2147483647 == -2147483648), F1: 4294967295}
    p0 -= (-(v23))
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    return 65535
}

func main() {
    var v1 int8 = 0
    var v2 int64 = 1
    var v3 uint16 = 2
    var v4 float64 = (0.1 / (1.5 + 0.1))
    v5 := []int{0}
    v6 := []int{128, -1}
    fmt.Println(v3)
    fmt.Println(v4)
    if len(v6) > 0 {
        v6[0] = 62
    }
    var v25 uint32 = uint32(v3)
    v6 = append(v6, -268)
    v26 := GenS2{F0: "\n\t"}
    v3 = (v3 ^ v3)
    if (942 <= 323) {
        if true {
            _ = len(v6)
            for fi27 := 0; fi27 < 17; fi27++ {
                for fi28 := 0; fi28 < 1; fi28++ {
                    var v29 uint8 = 128
                    var v30 uint64 = 850
                    if false {
                        if (true || (v30 > v30)) {
                            var v31 int8 = (v1 ^ genFn0(true, 528))
                            var v32 int32 = 2147483647
                            if true {
                                switch fi27 {
                                case 9223372036854775807:
                                    v31 = v31
                                case 232:
                                    var v33 int16 = -1
                                    _ = v33
                                }
                                var v34 float64 = 0.0
                                if (1e10 == 0.0) {
                                    if false {
                                        if true {
                                            v29 += v29
                                            fmt.Println(v4)
                                            var v35 float64 = (1.5 / (((float64(float32(v34)) + (v34 - (v34 - 1e10))) * (3.14159 / v4)) * v34))
                                            v36 := GenS2{F0: "\n\t"}
                                            _ = v35
                                            _ = v36
                                        }
                                        v25 = uint32(v30)
                                        v37 := GenS1{F0: false, F1: 42}
                                        _ = v37
                                    }
                                }
                                switch (-(fi27)) {
                                case 128:
                                    fmt.Println(v32, v3)
                                case -2147483648:
                                    v3 = v3
                                default:
                                    var v38 int64 = v2
                                    _ = v38
                                }
                                _ = v34
                            }
                            v29 += (v29 & 151)
                            switch "\n\t" {
                            case "\n\t":
                                v39 := GenS1{F0: true, F1: 4294967295}
                                _ = v39
                            case "a":
                                fmt.Println(v30, fi28)
                                v3 = (v3 & uint16(v30))
                                var v40 float64 = v4
                                _ = v40
                            default:
                                v3 -= v3
                            }
                            _ = v31
                            _ = v32
                        } else {
                            v25 += genFn1(fi28, (^(v1)), (v30 * (v30 & uint64(v3))))
                        }
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = fi27
            }
        } else {
            var v41 uint16 = v3
            _ = v41
        }
        _ = len(v6)
    } else {
        var v42 float64 = 3.14159
        if len(v6) > 0 {
            v6[0] = 128
        }
        _ = v42
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
