package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int8; F1 bool; F2 string }
func genFn0(p0 uint16) int8 {
    v7 := GenS1{F0: -1, F1: "a"}
    switch "\n\t" {
    case "":
        if (p0 <= (p0 << ((65535) & 63))) {
            var v8 int64 = -1
            var v9 float64 = (((1.5 - float64(((3.14159 / float32(0.1)) * 0.1))) - -2.25) - 1.5)
            if (0.1 > 3.14159) {
                var v10 int16 = 1000
                _ = v10
            }
            _ = v8
            _ = v9
        } else {
            if (-916 <= 42) {
                switch 42 {
                case 677:
                    if true {
                        fmt.Println(p0)
                        p0 -= p0
                        p0 = 579
                    }
                case -9223372036854775808:
                    p0 = p0
                }
                var v11 float64 = 1e10
                _ = v11
            }
        }
    }
    var v12 int16 = 451
    v13 := []int{4294967295, 127, 42}
    if false {
        p0 += 685
    }
    _ = p0
    _ = v12
    _ = v13
    _ = v7
    return 1
}
func genFn1(p0 int, p1 int32, p2 uint8) uint {
    fmt.Println(p0, p1)
    p1 = -1
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{}
    v3 := []uint32{}
    v4 := []string{("" + ""), (((("hello" + "a") + "") + "a") + "hello"), ("foo bar" + ("hello" + "hello"))}
    v5 := make(map[int]bool)
    var v6 int = 1
    if false {
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS2{F0: (v6 & fi14)}
            switch ("foo bar" + ("a" + "")) {
            case "foo bar":
                v6 += -471
            }
            _ = fi14
            _ = v15
        }
        v6 = 333
    }
    var v16 float32 = (float32(1.5) - (3.14159 * 3.14159))
    v17 := GenS3{F0: false}
    if len(v4) > 0 {
        v4[0] = (("\n\t" + "a") + "foo bar")
    }
    var v18 bool = (((0.1 + (float32((1e10 * 1.5)) * float32(1e10))) / (((float32((0.1 + (-2.25 / 3.14159))) * float32(3.14159)) / (3.14159 + v16)) + 1.5)) < 1.5)
    fmt.Println(v18, v16)
    if true {
        if false {
            switch (("a" + "hello") + "foo bar") {
            case "\n\t":
                var v19 int64 = -178
                _ = v19
            default:
                delete(v5, v6)
            }
            v18 = ((!false) && true)
            fmt.Println(v16)
        } else {
            if (v6 == v6) {
                var v20 int16 = 1
                _ = v20
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                _ = len(v3)
                for fi22 := 0; fi22 < 15; fi22++ {
                    fmt.Println(fi21)
                    switch "hello" {
                    case "a":
                        v6 = fi21
                    case "":
                        switch "foo bar" {
                        case "":
                            v23 := GenS1{F0: 0, F1: ((("a" + "\n\t") + "hello") + "")}
                            _ = v23
                        }
                    case "\n\t":
                        switch "hello" {
                        case "hello":
                            v24 := GenS2{F0: (fi21 >> ((255) & 63))}
                            switch "foo bar" {
                            case "foo bar":
                                switch (-(v6)) {
                                case 128:
                                    fmt.Println(fi21)
                                case 0:
                                    fmt.Println(fi21)
                                default:
                                    v6 -= (-(v6))
                                }
                            case "\n\t":
                                if v18 {
                                    if v18 {
                                        v25 := GenS3{F0: v18}
                                        v26 := GenS1{F0: -2147483648, F1: (("\n\t" + ("foo bar" + "a")) + "\n\t")}
                                        _ = v25
                                        _ = v26
                                    } else {
                                        v18 = (v18 || v18)
                                        var v27 int = 128
                                        v27 = -2147483648
                                        _ = v27
                                    }
                                    v16 += v16
                                } else {
                                    v16 = ((0.0 * v16) - v16)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            default:
                                v6 = (^(fi21))
                            }
                            _ = v24
                        }
                    }
                    switch 42 {
                    case 42:
                        v16 = (float32(-2.25) + (0.0 / v16))
                    default:
                        v28 := GenS3{F0: v18}
                        _ = v28
                    }
                    _ = fi22
                }
                var v29 string = ("foo bar" + "foo bar")
                _ = fi21
                _ = v29
            }
        }
        var v30 string = ("a" + "hello")
        switch (v6 << ((201) & 63)) {
        case 42:
            var v31 float32 = (((v16 * v16) - (0.0 / -2.25)) / v16)
            _ = v31
        default:
            v30 = v30
        }
        _ = v30
    }
    v18 = (1 != -917)
    if (v6 == 1) {
        for fi32 := 0; fi32 < 2; fi32++ {
            var v33 uint32 = 65535
            v34 := GenS3{F0: (909 >= 0)}
            var v35 string = ("a" + "hello")
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
    } else {
        v36 := GenS2{F0: 42}
        fmt.Println(v18)
        var v37 uint = 4294967295
        _ = v36
        _ = v37
    }
    var v38 float64 = -2.25
    if v18 {
        v39 := GenS4{F0: -880, F1: (-(v6))}
        _ = v39
    } else {
        var v40 int32 = 252
        switch (("foo bar" + ((("" + "\n\t") + "") + ("\n\t" + "hello"))) + "\n\t") {
        case "foo bar":
            fmt.Println(v6, v38)
        case "\n\t":
            var v41 int16 = -1
            _ = v41
        default:
            if len(v4) > 0 {
                v4[0] = (("\n\t" + "foo bar") + "foo bar")
            }
        }
        v5[v6] = v18
        _ = v40
    }
    v6 -= v6
    v42 := GenS4{F0: -251, F1: (^(v6))}
    if v18 {
        var v43 uint = 2
        _ = v43
    }
    for fi44 := 0; fi44 < 18; fi44++ {
        fmt.Println(fi44)
        v45 := GenS5{F0: -128, F1: (!v18), F2: "foo bar"}
        _ = fi44
        _ = v45
    }
    v16 -= v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
