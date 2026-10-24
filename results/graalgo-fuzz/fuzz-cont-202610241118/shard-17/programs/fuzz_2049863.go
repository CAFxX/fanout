package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
func genFn0(p0 bool) int32 {
    switch 0 {
    case 128:
        var v4 int16 = 1000
        _ = v4
    }
    v5 := GenS1{F0: "foo bar", F1: 113}
    if true {
        var v6 int8 = -8
        v7 := []int{}
        v6 -= 42
        _ = v6
        _ = v7
    } else {
        var v8 int32 = 368
        if (1 != 2) {
            if true {
                p0 = (578 < -84)
                for fi9 := 0; fi9 < 6; fi9++ {
                    p0 = p0
                    v10 := GenS1{F0: ("\n\t" + (("\n\t" + "hello") + "\n\t")), F1: 100}
                    _ = fi9
                    _ = v10
                }
                var v11 string = "foo bar"
                _ = v11
            }
            if (65535 <= 2147483648) {
                v12 := GenS1{F0: "", F1: 0}
                _ = v12
            } else {
                v13 := GenS2{F0: 8}
                fmt.Println(v8)
                for fi14 := 0; fi14 < 6; fi14++ {
                    var v15 int16 = 32767
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
            v16 := GenS2{F0: -603}
            _ = v16
        }
        v17 := GenS1{F0: "foo bar", F1: 1}
        _ = v17
        _ = v8
    }
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 string, p1 uint8, p2 uint64) int32 {
    v18 := GenS1{F0: p0, F1: 127}
    v19 := []int{385}
    v19 = append(v19, -9223372036854775808)
    if (((1 > 1) && ((float32(float64(0.0)) - 0.1) > 1e10)) && (false || (42 <= 363))) {
        fmt.Println(p2, p2)
    } else {
        for fi20 := 0; fi20 < 7; fi20++ {
            v19 = append(v19, -144)
            if true {
                p1 += p1
            } else {
                switch fi20 {
                case 0:
                    p1 += (p1 - (p1 / ((215) | 1)))
                default:
                    fmt.Println(p0, fi20, p2)
                }
                switch 359 {
                case 127:
                    fmt.Println(p0)
                }
            }
            _ = fi20
        }
        var v21 int32 = 0
        _ = v21
    }
    if (true && (((p1 - (p1 | p1)) < (p1 / ((p1) | 1))) || (true && false))) {
        p2 = (^(p2))
    } else {
        if (!false) {
            switch 948 {
            case 9223372036854775807:
                if true {
                    if true {
                        p2 = p2
                        var v22 uint16 = 591
                        fmt.Println(p1, p1)
                        _ = v22
                    }
                    v23 := GenS2{F0: -1}
                    _ = v23
                }
            case 0:
                for fi24 := 0; fi24 < 3; fi24++ {
                    v19 = append(v19, fi24)
                    var v25 uint8 = p1
                    var v26 string = (p0 + "a")
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            default:
                fmt.Println(p2, p0, p1)
            }
            if (42 >= 42) {
                for fi27 := 0; fi27 < 12; fi27++ {
                    fmt.Println(p1, p0, p0)
                    if (!(p1 >= 5)) {
                        var v28 float64 = 0.0
                        _ = v28
                    }
                    _ = len(v19)
                    _ = fi27
                }
            } else {
                switch -212 {
                case 32767:
                    switch -495 {
                    case 2147483647:
                        p1 += p1
                    case 542:
                        p0 = (("foo bar" + ("a" + p0)) + p0)
                    }
                default:
                    fmt.Println(p1)
                }
                if len(v19) > 0 {
                    v19[0] = 0
                }
            }
            for fi29 := 0; fi29 < 13; fi29++ {
                p2 = 1
                var v30 float32 = (0.1 + 1.5)
                _ = fi29
                _ = v30
            }
        }
        p0 = (p0 + (p0 + "a"))
        v19 = append(v19, 127)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return 801
}

func main() {
    v1 := []int{128, -9223372036854775808, -10}
    v2 := make(map[int]int)
    var v3 uint8 = 85
    var v31 uint16 = 842
    fmt.Println(v31, v31)
    v3 = v3
    var v32 float32 = ((0.1 * float32(-2.25)) - 1e10)
    v33 := GenS3{F0: ("hello" + ("foo bar" + ""))}
    fmt.Println(v32, v3, v31)
    var v34 float64 = (((1.5 * (0.1 - 1e10)) * float64(v32)) + float64((0.0 / (v32 - float32(0.1)))))
    var v35 int16 = 42
    if false {
        v36 := GenS2{F0: -865}
        switch ((("a" + "\n\t") + "a") + "a") {
        case "a":
            _ = len(v1)
        case "hello":
            _ = len(v1)
        default:
            fmt.Println(v34, v31, v32)
        }
        _ = v36
    }
    if len(v1) > 0 {
        v1[0] = 42
    }
    v35 += (v35 >> ((196) & 63))
    v37 := GenS2{F0: -2147483648}
    v32 -= ((0.0 - ((((v32 + v32) * v32) / 1.0) * v32)) + 1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v37
}
