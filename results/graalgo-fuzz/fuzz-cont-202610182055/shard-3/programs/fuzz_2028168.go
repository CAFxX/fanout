package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint; F1 uint16; F2 uint64 }
func genFn0(p0 int8, p1 int32, p2 int16) int32 {
    fmt.Println(p1, p0)
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := []int{}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1(p0 uint32) int8 {
    if ((("foo bar" + (("hello" + "foo bar") + "")) + "a") >= "a") {
        p0 = p0
        for fi7 := 0; fi7 < 9; fi7++ {
            p0 = (p0 + (p0 >> ((255) & 63)))
            p0 -= (p0 & (p0 >> ((4294967295) & 63)))
            v8 := []int{fi7}
            _ = fi7
            _ = v8
        }
    } else {
        fmt.Println(p0)
        if (394 != 454) {
            for fi9 := 0; fi9 < 6; fi9++ {
                if true {
                    p0 -= 2147483647
                } else {
                    var v10 float32 = float32((1e10 + 1.5))
                    _ = v10
                }
                for fi11 := 0; fi11 < 2; fi11++ {
                    if false {
                        v12 := []int{(fi9 - (fi9 >> ((18446744073709551615) & 63))), -2147483648}
                        if (false && (!false)) {
                            p0 -= p0
                            var v13 float32 = float32(3.14159)
                            var v14 int32 = genFn0(42, 0, -293)
                            _ = v13
                            _ = v14
                        } else {
                            var v15 int64 = 0
                            _ = v15
                        }
                        _ = v12
                    } else {
                        var v16 string = "hello"
                        var v17 uint32 = p0
                        _ = v16
                        _ = v17
                    }
                    _ = fi11
                }
                _ = fi9
            }
            v18 := []int{9223372036854775807}
            var v19 int = -848
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 7; fi20++ {
                switch fi20 {
                case -945:
                    var v21 bool = (!((42 > 485) || false))
                    _ = v21
                }
                switch 32767 {
                case 216:
                    fmt.Println(fi20)
                case 32767:
                    fmt.Println(fi20, fi20)
                }
                _ = fi20
            }
            fmt.Println(p0)
        }
        var v22 int8 = 90
        _ = v22
    }
    p0 += p0
    if false {
        v23 := GenS1{F0: 2, F1: 199}
        _ = v23
    }
    if true {
        switch "a" {
        case "\n\t":
            var v24 int32 = 42
            _ = v24
        case "hello":
            switch 148 {
            case 127:
                switch (("foo bar" + "") + ("hello" + "")) {
                case "foo bar":
                    var v25 float32 = ((((float32(3.14159) - 1e10) * float32(0.0)) * 1.5) * 0.0)
                    _ = v25
                case "a":
                    p0 = 42
                default:
                    fmt.Println(p0)
                }
            case 32768:
                p0 = p0
            case 0:
                if false {
                    var v26 uint = 429
                    if false {
                        if (false || (!true)) {
                            switch -914 {
                            case 42:
                                if true {
                                    v27 := []int{9223372036854775807}
                                    _ = v27
                                }
                            }
                        }
                    }
                    v28 := []int{709}
                    _ = v26
                    _ = v28
                }
            }
        default:
            v29 := GenS1{F0: 169, F1: 134}
            _ = v29
        }
    } else {
        var v30 uint16 = 579
        p0 += p0
        _ = v30
    }
    _ = p0
    return -65
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 42
    var v3 uint = 255
    var v4 float32 = 0.0
    v31 := GenS2{F0: 167}
    v32 := []int{}
    v4 = v4
    v33 := GenS1{F0: 2, F1: 128}
    var v34 int32 = 1
    var v35 int8 = -128
    for fi36 := 0; fi36 < 20; fi36++ {
        switch (fi36 % ((fi36) | 1)) {
        case 45:
            var v37 float64 = 0.0
            _ = v37
        case 32767:
            v38 := GenS3{F0: (fi36 | fi36), F1: "foo bar"}
            _ = v38
        case 2147483647:
            fmt.Println(v1, v34)
        default:
            fmt.Println(v2, v1, v34)
        }
        for fi39 := 0; fi39 < 6; fi39++ {
            switch "\n\t" {
            case "\n\t":
                v40 := GenS4{F0: 255, F1: 162, F2: 2}
                _ = v40
            case "a":
                v34 = v34
            case "foo bar":
                v4 += (1.5 / float32(1.5))
            default:
                fmt.Println(v2)
            }
            v1 = 42
            v34 = (-(v34))
            _ = fi39
        }
        _ = fi36
    }
    v41 := GenS3{F0: 42, F1: ("a" + ("" + "\n\t"))}
    v32 = append(v32, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v41
}
