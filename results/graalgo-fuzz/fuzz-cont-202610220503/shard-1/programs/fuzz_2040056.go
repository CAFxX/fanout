package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int32 }
type GenS2 struct { F0 int8; F1 string; F2 int64 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint32; F1 uint; F2 uint }
type GenS5 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint8, p1 int16, p2 uint) string {
    v4 := []int{0}
    switch (((("foo bar" + (("foo bar" + "hello") + ("foo bar" + "\n\t"))) + "foo bar") + ("a" + "foo bar")) + ((("a" + "hello") + ("a" + "\n\t")) + "a")) {
    case "":
        var v5 uint16 = 65535
        _ = v5
    case "a":
        v6 := GenS1{F0: 4294967295, F1: false, F2: 42}
        _ = v6
    default:
        switch 57 {
        case 361:
            p2 += (^(p2))
        default:
            switch "a" {
            case "foo bar":
                fmt.Println(p2, p2, p1)
            case "hello":
                for fi7 := 0; fi7 < 9; fi7++ {
                    switch "foo bar" {
                    case "a":
                        p1 += 0
                    case "hello":
                        fmt.Println(p2)
                    case "\n\t":
                        var v8 uint64 = 0
                        _ = v8
                    default:
                        fmt.Println(fi7, p1, p1)
                    }
                    var v9 float32 = 1.5
                    for fi10 := 0; fi10 < 10; fi10++ {
                        var v11 int16 = 1000
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v12 int16 = 1
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = fi7
                    _ = v9
                }
            default:
                if len(v4) > 0 {
                    v4[0] = 42
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "hello"
}
func genFn1(p0 uint, p1 bool) uint32 {
    var v13 string = "foo bar"
    if (65535 >= 2) {
        p1 = ((1e10 != 1e10) || true)
        switch 32768 {
        case -9223372036854775808:
            var v14 float64 = 0.0
            _ = v14
        case -2147483648:
            if p1 {
                switch "hello" {
                case "":
                    var v15 int64 = -673
                    _ = v15
                case "a":
                    var v16 uint16 = 255
                    _ = v16
                }
                v17 := GenS2{F0: -128, F1: "a", F2: 42}
                p0 -= p0
                _ = v17
            } else {
                p1 = true
            }
        case 0:
            if (-2.25 <= -2.25) {
                p0 += (^(p0))
            }
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            v13 = v13
            _ = fi18
        }
    } else {
        v19 := GenS3{F0: 32768, F1: 2}
        fmt.Println(p0, v13, p1)
        _ = v19
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        fmt.Println(p0)
        switch (-(fi20)) {
        case 0:
            var v21 uint16 = 186
            _ = v21
        case 32768:
            v22 := GenS4{F0: 4294967295, F1: p0, F2: (p0 * (p0 + p0))}
            _ = v22
        case -509:
            v13 = ("a" + ((("a" + v13) + (v13 + (("\n\t" + v13) + (v13 + v13)))) + ((v13 + "") + "hello")))
        }
        _ = fi20
    }
    p0 += (^(p0))
    if false {
        var v23 uint16 = uint16(p0)
        p1 = ((-2.25 - 1e10) <= 1.5)
        _ = v23
    } else {
        var v24 float64 = (float64(3.14159) / 1.5)
        p1 = ((p0 < 18446744073709551615) && false)
        for fi25 := 0; fi25 < 3; fi25++ {
            if p1 {
                v26 := GenS3{F0: fi25, F1: 1}
                v27 := GenS5{F0: 1, F1: 1}
                _ = v26
                _ = v27
            }
            fmt.Println(v13, v13, v24)
            _ = fi25
        }
        _ = v24
    }
    _ = p0
    _ = p1
    _ = v13
    return 233
}

func main() {
    var v1 bool = ((((1e10 + 1.5) + 0.1) == (-2.25 + 0.0)) && (("a" + ("hello" + "\n\t")) != ("\n\t" + ("foo bar" + "foo bar"))))
    var v2 int = -5
    v3 := []uint32{6}
    var v28 string = (((("foo bar" + "foo bar") + "hello") + (("hello" + "") + (((("" + "\n\t") + "foo bar") + ("\n\t" + ((("foo bar" + "") + "\n\t") + "a"))) + (("hello" + "foo bar") + ("\n\t" + ""))))) + "hello")
    var v29 int16 = -32768
    _ = len(v3)
    var v30 int = (v2 & v2)
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 uint = 1
        _ = fi31
        _ = v32
    }
    fmt.Println(v28, v1)
    fmt.Println(v2, v30)
    var v33 bool = (!((false && false) || (1 == 0)))
    fmt.Println(v28, v1, v28)
    v30 = (-(v2))
    v34 := GenS2{F0: 0, F1: v28, F2: 42}
    v2 = v30
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
}
