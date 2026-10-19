package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 bool; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint32; F2 int64 }
func genFn0(p0 uint32) uint8 {
    for fi6 := 0; fi6 < 10; fi6++ {
        p0 = (p0 % ((696) | 1))
        switch ((("" + (("" + "\n\t") + "hello")) + ("" + "")) + (("hello" + "a") + "")) {
        case "hello":
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := GenS1{F0: 127, F1: -1}
                _ = fi7
                _ = v8
            }
        }
        v9 := GenS2{F0: false, F1: 299, F2: 128}
        _ = fi6
        _ = v9
    }
    v10 := GenS3{F0: 32767}
    _ = p0
    _ = v10
    return 2
}
func genFn1(p0 int32) string {
    var v11 int = 10
    p0 = (p0 * -885)
    _ = p0
    _ = v11
    return "\n\t"
}

func main() {
    v1 := []uint32{170, 4294967295, 42}
    v2 := make(map[uint32]int)
    v3 := []uint32{2147483647}
    v4 := []string{("\n\t" + ""), "a", "a", "\n\t"}
    var v5 float64 = (-2.25 * float64(1e10))
    switch (("hello" + ("" + ("a" + "foo bar"))) + ("" + "foo bar")) {
    case "\n\t":
        var v12 bool = false
        _ = v12
    default:
        if false {
            v13 := GenS4{F0: "\n\t", F1: 981, F2: 1}
            v5 = (v5 * (((((v5 - v5) - 1.5) / v5) - float64(3.14159)) * (v5 / v5)))
            var v14 int = -1
            _ = v13
            _ = v14
        }
    }
    v5 = (float64(0.1) - 3.14159)
    fmt.Println(v5)
    if (-1 != -32768) {
        var v15 uint64 = 2
        _ = v15
    } else {
        switch "foo bar" {
        case "\n\t":
            var v16 uint8 = 128
            _ = v16
        case "hello":
            for fi17 := 0; fi17 < 4; fi17++ {
                for fi18 := 0; fi18 < 14; fi18++ {
                    v19 := v2[0]
                    _ = fi18
                    _ = v19
                }
                v5 -= ((1e10 / float64(1.5)) / 0.1)
                _ = fi17
            }
        case "":
            v5 += v5
        default:
            var v20 uint64 = 42
            _ = v20
        }
        fmt.Println(v5)
        v5 += v5
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := v2[65535]
        _ = fi21
        _ = v22
    }
    v5 += 0.0
    fmt.Println(v5)
    fmt.Println(v5)
    if true {
        fmt.Println(v5)
        var v23 int32 = 207
        _ = v23
    } else {
        var v24 bool = (genFn0(2147483648) < 42)
        var v25 uint64 = 2
        _ = v24
        _ = v25
    }
    var v26 int8 = 100
    var v27 uint8 = 42
    for fi28 := 0; fi28 < 16; fi28++ {
        if false {
            if (!true) {
                var v29 int = 660
                var v30 uint64 = 18446744073709551615
                _ = v29
                _ = v30
            }
            v27 -= (v27 + 128)
            if true {
                switch fi28 {
                case 128:
                    if (-673 > 42) {
                        if false {
                            switch "foo bar" {
                            case "":
                                v26 -= v26
                            case "a":
                                v26 += -128
                            default:
                                switch ("\n\t" + ("hello" + "a")) {
                                case "hello":
                                    if (true && true) {
                                        fmt.Println(v5, fi28)
                                    }
                                case "\n\t":
                                    v31 := GenS4{F0: ((((("foo bar" + ("foo bar" + "hello")) + "a") + "foo bar") + "hello") + "a"), F1: 72, F2: -1}
                                    _ = v31
                                case "":
                                    fmt.Println(v26)
                                }
                            }
                        } else {
                            fmt.Println(fi28)
                        }
                        v27 += genFn0(2147483647)
                    }
                case -906:
                    v32 := GenS2{F0: ((v27 <= v27) || (1 != 42)), F1: 765, F2: (v27 & (v27 & genFn0(774)))}
                    _ = v32
                }
                switch ("hello" + "") {
                case "foo bar":
                    v33 := GenS4{F0: "\n\t", F1: 352, F2: -9223372036854775808}
                    _ = v33
                case "\n\t":
                    fmt.Println(fi28)
                case "":
                    v34 := GenS3{F0: 631}
                    _ = v34
                default:
                    if true {
                        fmt.Println(fi28, v27, v27)
                    }
                }
                v26 = v26
            } else {
                if false {
                    var v35 bool = ((!false) && true)
                    var v36 uint16 = 756
                    _ = v35
                    _ = v36
                }
                var v37 float64 = v5
                v38 := GenS3{F0: 1}
                _ = v37
                _ = v38
            }
        }
        v26 = (v26 - v26)
        _ = fi28
    }
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
