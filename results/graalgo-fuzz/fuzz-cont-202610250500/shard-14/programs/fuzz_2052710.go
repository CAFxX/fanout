package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint }
type GenS3 struct { F0 string; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int8; F1 bool; F2 bool }
func genFn0(p0 int8, p1 int8, p2 uint) uint8 {
    for fi6 := 0; fi6 < 10; fi6++ {
        p0 += (p1 | (p0 | (p1 & p1)))
        v7 := []int{(fi6 ^ (fi6 + fi6))}
        _ = fi6
        _ = v7
    }
    var v8 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}
func genFn1(p0 string, p1 bool) string {
    p1 = (!(!((p1 || (("hello" + p0) != ("foo bar" + (("a" + "hello") + "\n\t")))) && (42 == 65535))))
    if false {
        fmt.Println(p0)
        v9 := []int{708, -456, 4294967295, 2147483647}
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS1{F0: 948, F1: 128}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    return (p0 + p0)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 int16 = -1
    var v4 float64 = 0.1
    v5 := make(map[uint32]uint32)
    for fi12 := 0; fi12 < 4; fi12++ {
        switch ((((("a" + "foo bar") + ("\n\t" + "\n\t")) + "foo bar") + (((("hello" + "\n\t") + "hello") + ("\n\t" + "\n\t")) + "a")) + "hello") {
        case "\n\t":
            fmt.Println(fi12, v3, v3)
        default:
            var v13 int32 = -2147483648
            _ = v13
        }
        if true {
            v14 := GenS2{F0: 1, F1: 2, F2: 580}
            var v15 uint = 4294967295
            var v16 string = ("\n\t" + ("foo bar" + ("\n\t" + "a")))
            _ = v14
            _ = v15
            _ = v16
        } else {
            v4 -= ((v4 - (float64(-2.25) * 1.5)) / float64((1e10 / 1.5)))
        }
        _ = len(v5)
        _ = fi12
    }
    var v17 float64 = (v4 + (v4 + 0.0))
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 uint = 0
        switch ("foo bar" + "a") {
        case "":
            v20 := []int{}
            _ = v20
        default:
            if false {
                var v21 int8 = 1
                _ = v21
            }
        }
        _ = fi18
        _ = v19
    }
    if (!(false || false)) {
        for fi22 := 0; fi22 < 15; fi22++ {
            v17 -= v4
            fmt.Println(fi22, v4)
            _ = fi22
        }
    }
    v23 := GenS1{F0: -1, F1: 137}
    v24 := []int{32767, 454}
    switch "\n\t" {
    case "hello":
        v4 = (0.1 / ((float64((1.5 - 3.14159)) * (-2.25 / ((((1.5 - 1e10) / v17) - 0.0) - v4))) * v17))
    }
    var v25 int8 = 106
    fmt.Println(v17)
    v26 := GenS3{F0: "\n\t", F1: 1, F2: 2}
    if true {
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 int8 = v25
            var v29 float64 = v4
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    switch "hello" {
    case "":
        v4 = (v17 / 0.1)
    default:
        v17 = (v4 / v4)
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 uint32 = 1
            switch "foo bar" {
            case "hello":
                _ = len(v5)
            case "\n\t":
                if true {
                    var v33 int = (fi30 + (fi31 - (fi31 ^ -1)))
                    if ((!(((540 != 1) && false) || false)) && ((v32 + v32) != 621)) {
                        v3 = (v3 / ((42) | 1))
                        switch "\n\t" {
                        case "hello":
                            var v34 float64 = ((float64(float32((0.0 / v4))) * (v17 * -2.25)) - 0.1)
                            _ = v34
                        case "foo bar":
                            if len(v24) > 0 {
                                v24[0] = v33
                            }
                        }
                        v35 := GenS4{F0: 103, F1: false, F2: true}
                        _ = v35
                    } else {
                        v25 += v25
                        v32 = 65535
                        v32 += 2
                    }
                    _ = v33
                }
            case "a":
                fmt.Println(v17, v32)
            default:
                _ = len(v24)
            }
            switch 42 {
            case -334:
                var v36 int64 = 42
                _ = v36
            case 128:
                var v37 int64 = 151
                _ = v37
            case 704:
                v3 += v3
            default:
                var v38 int16 = v3
                _ = v38
            }
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
