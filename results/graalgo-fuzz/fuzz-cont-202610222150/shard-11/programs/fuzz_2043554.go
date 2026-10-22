package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint64 }
func genFn0(p0 string) string {
    v6 := []int{9223372036854775807, -2147483648, -264}
    v6 = append(v6, -1000)
    var v7 uint64 = 1
    if (!true) {
        var v8 uint8 = 2
        _ = v8
    } else {
        var v9 int16 = 32767
        switch p0 {
        case "a":
            fmt.Println(v9, p0)
        default:
            fmt.Println(p0, v9)
        }
        fmt.Println(v7, p0)
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    var v1 uint32 = 1
    v2 := []uint32{v1, 65535, (v1 - v1), v1}
    v3 := []string{(("" + (((("hello" + ("foo bar" + "hello")) + "foo bar") + ("hello" + "hello")) + "a")) + ""), ((("" + "") + "a") + "foo bar"), ("hello" + "")}
    v4 := make(map[int]int)
    var v5 int16 = 42
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(fi10, v5, v5)
        var v11 int8 = 127
        v11 = (v11 ^ 84)
        _ = fi10
        _ = v11
    }
    var v12 float32 = 0.0
    var v13 uint8 = 255
    v1 += (v1 - (^(v1)))
    v12 = ((v12 / v12) + v12)
    if true {
        v1 -= 579
        for fi14 := 0; fi14 < 2; fi14++ {
            v12 -= (v12 / v12)
            var v15 string = "hello"
            _ = fi14
            _ = v15
        }
        v16 := GenS1{F0: 1, F1: 1}
        _ = v16
    }
    var v17 int64 = 0
    var v18 bool = ((1 != 18446744073709551615) && true)
    var v19 uint32 = (v1 >> ((255) & 63))
    fmt.Println(v12, v17, v17)
    v12 += -2.25
    v17 = (v17 & -827)
    v1 -= (v1 * 2147483647)
    if true {
        v20 := GenS2{F0: 0, F1: v13, F2: 1}
        _ = v20
    } else {
        if v18 {
            var v21 int = 127
            for fi22 := 0; fi22 < 9; fi22++ {
                v1 += 2147483647
                _ = fi22
            }
            _ = v21
        }
        if v18 {
            if v18 {
                v5 -= v5
                var v23 bool = v18
                _ = v23
            } else {
                switch "foo bar" {
                case "hello":
                    if true {
                        switch -1 {
                        case 32767:
                            for fi24 := 0; fi24 < 5; fi24++ {
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    v19 -= v19
                                    v26 := GenS1{F0: v17, F1: -128}
                                    var v27 uint8 = 138
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                                if (0 < v19) {
                                    if (!(v17 > v17)) {
                                        fmt.Println(v1)
                                        v17 -= 9223372036854775807
                                        if len(v2) > 0 {
                                            v2[0] = v19
                                        }
                                    }
                                    for fi28 := 0; fi28 < 18; fi28++ {
                                        var v29 float32 = v12
                                        v19 = (v1 << ((867) & 63))
                                        var v30 float32 = 1e10
                                        _ = fi28
                                        _ = v29
                                        _ = v30
                                    }
                                    for fi31 := 0; fi31 < 8; fi31++ {
                                        var v32 string = "foo bar"
                                        _ = fi31
                                        _ = v32
                                    }
                                } else {
                                    if (42 < 42) {
                                        v18 = (100 != 100)
                                    }
                                }
                                fmt.Println(v19, v1)
                                _ = fi24
                            }
                        default:
                            v12 += -2.25
                        }
                    }
                default:
                    var v33 int16 = -1
                    _ = v33
                }
                _ = len(v4)
                _ = len(v3)
            }
        } else {
            if (false && false) {
                var v34 float32 = -2.25
                _ = v34
            } else {
                v35 := v4[32767]
                _ = v35
            }
            for fi36 := 0; fi36 < 20; fi36++ {
                v18 = true
                v1 -= v1
                _ = fi36
            }
            if true {
                var v37 float32 = 0.1
                var v38 uint = 1
                _ = v37
                _ = v38
            } else {
                v18 = true
                v13 -= (^(v13))
                v39 := GenS1{F0: (^(v17)), F1: 42}
                _ = v39
            }
        }
    }
    v13 = v13
    v40 := GenS1{F0: (v17 % ((-9223372036854775808) | 1)), F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
