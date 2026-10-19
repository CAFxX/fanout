package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int16; F2 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 int }
type GenS4 struct { F0 uint8; F1 int32 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    var v3 int8 = 1
    v4 := []int{}
    switch -2147483648 {
    case -842:
        if (!true) {
            var v5 float32 = 0.0
            fmt.Println(v5)
            switch "a" {
            case "a":
                var v6 float32 = 0.1
                _ = v6
            case "foo bar":
                v5 = v5
            }
            _ = v5
        }
    default:
        v3 -= 86
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v7 := GenS1{F0: 0}
    v3 -= -62
    if false {
        var v8 int32 = 42
        v8 = -1
        _ = v8
    } else {
        v9 := GenS2{F0: 0, F1: 0, F2: 65535}
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int = fi10
            if true {
                v12 := v2[("" + "a")]
                switch "hello" {
                case "\n\t":
                    v13 := GenS3{F0: 848, F1: 0, F2: (^(v12))}
                    _ = v13
                case "a":
                    fmt.Println(v11)
                default:
                    switch (("foo bar" + "\n\t") + (("a" + "") + ("a" + ("hello" + "hello")))) {
                    case "foo bar":
                        for fi14 := 0; fi14 < 15; fi14++ {
                            switch (v11 & (v12 + v12)) {
                            case -1:
                                switch 32767 {
                                case 870:
                                    v3 = v3
                                case 949:
                                    if true {
                                        v15 := GenS4{F0: 40, F1: -1}
                                        v12 += fi10
                                        v16 := v2[("hello" + "foo bar")]
                                        _ = v15
                                        _ = v16
                                    } else {
                                        v17 := GenS3{F0: 18446744073709551615, F1: 270, F2: 786}
                                        _ = v17
                                    }
                                case 42:
                                    var v18 int32 = 556
                                    _ = v18
                                }
                            case 42:
                                v12 = (fi14 - fi14)
                            case 32767:
                                v1[(v11 >> ((4294967295) & 63))] = -1
                            default:
                                fmt.Println(v3, fi14, v3)
                            }
                            _ = fi14
                        }
                    case "hello":
                        v12 = (v12 - -886)
                    case "":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        var v19 float64 = (1e10 + (3.14159 / 1.0))
                        _ = v19
                    }
                }
                fmt.Println(v3)
                _ = v12
            }
            var v20 int8 = 127
            _ = fi10
            _ = v11
            _ = v20
        }
        _ = v9
    }
    if true {
        var v21 uint = 255
        _ = v21
    }
    var v22 int = 9223372036854775807
    v22 -= (v22 * (v22 * v22))
    v3 += (v3 * 0)
    var v23 int16 = 316
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi24 := 0; fi24 < 6; fi24++ {
            var v25 int32 = 1
            var v26 int16 = 0
            _ = fi24
            _ = v25
            _ = v26
        }
    } else {
        var v27 int8 = v3
        if true {
            var v28 int8 = v3
            for fi29 := 0; fi29 < 14; fi29++ {
                v28 = 72
                _ = fi29
            }
            _ = v28
        } else {
            var v30 int = -379
            v23 += 42
            var v31 bool = true
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 18; fi32++ {
            v33 := GenS4{F0: 128, F1: -895}
            _ = fi32
            _ = v33
        }
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v7
}
