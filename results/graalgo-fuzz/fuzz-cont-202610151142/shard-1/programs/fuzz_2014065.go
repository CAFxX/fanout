package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 string; F1 uint16; F2 uint64 }
func genFn0(p0 bool, p1 uint32) uint {
    v7 := GenS1{F0: -802, F1: 0}
    p1 -= 4294967295
    fmt.Println(p1, p1)
    v8 := GenS2{F0: 42}
    if false {
        if (("hello" >= "") || p0) {
            v9 := []int{-2147483648}
            var v10 string = "a"
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 20; fi11++ {
                for fi12 := 0; fi12 < 16; fi12++ {
                    v13 := []int{(fi11 | -168)}
                    switch "\n\t" {
                    case "\n\t":
                        var v14 uint16 = 65535
                        _ = v14
                    case "foo bar":
                        _ = len(v13)
                        switch ("foo bar" + "foo bar") {
                        case "hello":
                            var v15 float64 = 3.14159
                            fmt.Println(fi11, p0, fi12)
                            var v16 uint8 = 1
                            _ = v15
                            _ = v16
                        case "\n\t":
                            p1 = (p1 << ((255) & 63))
                        default:
                            p1 -= p1
                        }
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    default:
                        p1 -= p1
                    }
                    _ = fi12
                    _ = v13
                }
                p1 = 632
                _ = fi11
            }
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 int32 = -883
    var v2 float64 = ((((3.14159 - 3.14159) + (1.5 * -2.25)) + 1e10) / (((0.0 + 0.0) - ((0.1 - ((((0.1 * (0.0 * ((-2.25 + (-2.25 * (((-2.25 / 1.0) - ((3.14159 - (((-2.25 * 1.5) + 1e10) - ((-2.25 + -2.25) / (-2.25 / -2.25)))) + (((0.0 + 1.5) * ((0.0 + (float64((1e10 * (3.14159 / ((3.14159 / float32(1.5)) / (1e10 / -2.25))))) / 0.1)) + float64(3.14159))) * 0.0))) + 3.14159))) - 0.1))) + (((((float64((((-2.25 - 0.1) / 3.14159) + (1.5 / float32((1e10 + 1.5))))) + -2.25) * (float64(1.5) * 0.1)) / 0.1) - float64((float32(-2.25) - (float32((0.0 - 1.5)) + 1.5)))) - 1e10)) / ((-2.25 * ((3.14159 - float64((0.0 * 0.1))) / 0.1)) * -2.25)) + (1.5 * ((-2.25 / 1.5) - float64(3.14159))))) / 1.5)) / 3.14159))
    v3 := make(map[string]int)
    v4 := []uint32{}
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]bool)
    if false {
        if true {
            v17 := GenS3{F0: 65535, F1: 42, F2: 215}
            fmt.Println(v1, v2)
            _ = v17
        } else {
            var v18 int16 = 32767
            _ = v18
        }
        var v19 int16 = -32768
        _ = v19
    }
    fmt.Println(v1, v1)
    v2 = (v2 + (0.0 * 0.1))
    if true {
        v20 := GenS3{F0: 0, F1: 2, F2: 68}
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 float64 = v2
            _ = fi21
            _ = v22
        }
        _ = v20
    } else {
        if (42 == 480) {
            v2 += v2
            v23 := v5[9223372036854775807]
            fmt.Println(v2, v2, v2)
            _ = v23
        }
    }
    switch "hello" {
    case "a":
        _ = len(v5)
    case "foo bar":
        v2 += float64(((1.5 * 0.1) + 0.0))
    case "hello":
        v2 = v2
    }
    if true {
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 int16 = 321
            _ = fi24
            _ = v25
        }
        fmt.Println(v2, v2)
    }
    v26 := GenS1{F0: -1, F1: -535}
    var v27 int8 = 29
    var v28 uint = 1
    switch 32768 {
    case 590:
        if true {
            for fi29 := 0; fi29 < 10; fi29++ {
                v28 = 65535
                _ = fi29
            }
        }
    default:
        for fi30 := 0; fi30 < 3; fi30++ {
            v2 = -2.25
            var v31 int = fi30
            if len(v4) > 0 {
                v4[0] = 42
            }
            _ = fi30
            _ = v31
        }
    }
    var v32 float32 = (1.5 / 3.14159)
    switch ("hello" + "a") {
    case "hello":
        for fi33 := 0; fi33 < 1; fi33++ {
            for fi34 := 0; fi34 < 1; fi34++ {
                if true {
                    if (65535 != 609) {
                        v2 += ((v2 / v2) - v2)
                    }
                }
                fmt.Println(fi34, v28, v1)
                switch (("" + "foo bar") + ("a" + "foo bar")) {
                case "\n\t":
                    _ = len(v5)
                case "a":
                    v32 -= ((v32 / v32) / (v32 * (3.14159 + (1.5 + 1e10))))
                case "":
                    var v35 string = ("\n\t" + "hello")
                    _ = v35
                default:
                    var v36 int32 = v1
                    _ = v36
                }
                _ = fi34
            }
            for fi37 := 0; fi37 < 17; fi37++ {
                v28 = v28
                _ = fi37
            }
            _ = fi33
        }
    case "foo bar":
        v32 -= float32(v2)
    case "\n\t":
        fmt.Println(v1, v27)
    }
    if false {
        v38 := GenS4{F0: v28, F1: "foo bar"}
        _ = v38
    } else {
        for fi39 := 0; fi39 < 14; fi39++ {
            for fi40 := 0; fi40 < 13; fi40++ {
                var v41 bool = true
                _ = fi40
                _ = v41
            }
            if false {
                v28 -= 281
                if true {
                    fmt.Println(v1, v27)
                }
            } else {
                for fi42 := 0; fi42 < 9; fi42++ {
                    v43 := GenS5{F0: "a", F1: 0, F2: 2}
                    _ = fi42
                    _ = v43
                }
                v32 += 0.0
                v44 := GenS2{F0: 1000}
                _ = v44
            }
            if true {
                var v45 uint32 = 4294967295
                if true {
                    v28 = (v28 * (v28 + v28))
                }
                _ = v45
            } else {
                var v46 int16 = -1
                var v47 uint8 = 2
                _ = v46
                _ = v47
            }
            _ = fi39
        }
    }
    var v48 int32 = (v1 * v1)
    if (false || (!(v1 == v1))) {
        fmt.Println(v32)
        v1 -= (v1 * -1)
        if true {
            var v49 int = 454
            _ = v49
        } else {
            v27 = v27
        }
    } else {
        v50 := GenS2{F0: -1}
        fmt.Println(v28, v32, v48)
        _ = v50
    }
    v48 += v48
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
