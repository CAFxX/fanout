package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint; F1 string; F2 string }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int64; F1 int }
func genFn0(p0 int32, p1 uint16) bool {
    for fi7 := 0; fi7 < 9; fi7++ {
        switch ("\n\t" + "a") {
        case "foo bar":
            var v8 int8 = -128
            _ = v8
        default:
            var v9 uint64 = 18446744073709551615
            _ = v9
        }
        v10 := GenS1{F0: int8(fi7), F1: (((("foo bar" + ((("foo bar" + "") + "hello") + ("\n\t" + ("a" + "a")))) + (("hello" + "a") + ("foo bar" + (("a" + ("hello" + "foo bar")) + "")))) + ("" + (("" + "hello") + "hello"))) + "foo bar")}
        _ = fi7
        _ = v10
    }
    v11 := []int{9223372036854775807, 393}
    var v12 int32 = (p0 & p0)
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return (108 > int(p1))
}
func genFn1(p0 int, p1 uint64) string {
    var v13 float64 = 1e10
    var v14 uint = 255
    var v15 uint16 = 42
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return "foo bar"
}

func main() {
    var v1 string = ("hello" + "\n\t")
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v4 int16 = -308
    var v5 bool = (((v1 + ((v1 + v1) + v1)) != "foo bar") || true)
    var v6 float64 = 3.14159
    if v5 {
        v16 := []int{127, -1, -9223372036854775808}
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        v18 := []int{1, (fi17 & fi17), int(v4), fi17}
        var v19 bool = v5
        for fi20 := 0; fi20 < 11; fi20++ {
            v18 = append(v18, 127)
            _ = fi20
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    var v21 uint16 = 2
    switch ("hello" + v1) {
    case "a":
        var v22 uint16 = v21
        _ = v22
    case "":
        v4 = v4
    default:
        if ((v4 % (((v4 >> ((18446744073709551615) & 63))) | 1)) > -780) {
            if (!(v21 < v21)) {
                var v23 uint8 = 29
                _ = v23
            }
            for fi24 := 0; fi24 < 5; fi24++ {
                v25 := GenS2{F0: 1, F1: ((v1 + ((v1 + v1) + "a")) + "foo bar"), F2: "foo bar"}
                _ = fi24
                _ = v25
            }
            for fi26 := 0; fi26 < 18; fi26++ {
                fmt.Println(v5, v4)
                v1 = (("\n\t" + (v1 + v1)) + "")
                _ = fi26
            }
        } else {
            var v27 string = "foo bar"
            for fi28 := 0; fi28 < 8; fi28++ {
                switch ((("a" + v27) + v27) + "foo bar") {
                case "a":
                    if v5 {
                        v6 = ((v6 + float64(((float32(v6) / (1e10 / -2.25)) - -2.25))) - v6)
                        for fi29 := 0; fi29 < 3; fi29++ {
                            v3[(fi28 ^ int(v21))] = ("" + v27)
                            var v30 uint32 = 1
                            var v31 uint16 = (v21 | (v21 / ((v21) | 1)))
                            _ = len(v3)
                            fmt.Println(v4, v6, v27)
                            var v32 int8 = 0
                            fmt.Println(v1)
                            _ = fi29
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    } else {
                        v1 = "a"
                        var v33 string = "\n\t"
                        for fi34 := 0; fi34 < 9; fi34++ {
                            switch fi28 {
                            case 728:
                                delete(v2, 4294967295)
                            case 896:
                                v27 = ((v33 + "a") + ("foo bar" + (v33 + v1)))
                            case 2147483647:
                                v1 = (v33 + (("a" + v33) + (v33 + ((v27 + (v27 + "a")) + v27))))
                            }
                            _ = fi34
                        }
                        _ = v33
                    }
                case "\n\t":
                    for fi35 := 0; fi35 < 11; fi35++ {
                        switch v27 {
                        case "hello":
                            var v36 float64 = 0.0
                            _ = v36
                        }
                        v37 := GenS3{F0: 42, F1: ((v21 - (v21 * v21)) > v21)}
                        if ((0.1 - (-2.25 * 1e10)) != ((-2.25 - float32(v6)) + ((3.14159 - (0.1 + 0.1)) - 0.0))) {
                            v38 := v3[fi28]
                            v5 = v5
                            _ = v38
                        }
                        v39 := GenS2{F0: 4294967295, F1: v1, F2: "a"}
                        _ = fi35
                        _ = v37
                        _ = v39
                    }
                }
                _ = fi28
            }
            _ = v27
        }
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        v1 = v1
        _ = fi40
    }
    v5 = false
    for fi41 := 0; fi41 < 8; fi41++ {
        fmt.Println(v21, fi41, v1)
        _ = fi41
    }
    v42 := GenS4{F0: ((v6 - 3.14159) > v6), F1: 0}
    v43 := v3[127]
    var v44 int = 42
    v44 = (v44 * (v44 + (v44 / ((v44) | 1))))
    if v5 {
        fmt.Println(v1, v44, v6)
    } else {
        if false {
            v45 := GenS5{F0: 914, F1: v44}
            _ = v45
        }
        switch "" {
        case "hello":
            var v46 uint64 = 18446744073709551615
            _ = v46
        case "a":
            if v5 {
                v1 = "a"
                fmt.Println(v6, v21)
            } else {
                var v47 bool = false
                _ = v47
            }
        default:
            v48 := []int{}
            _ = v48
        }
    }
    v43 = (v43 + ("\n\t" + ""))
    for fi49 := 0; fi49 < 1; fi49++ {
        var v50 uint = 0
        _ = fi49
        _ = v50
    }
    v5 = v5
    v51 := GenS4{F0: (65535 != 2147483648), F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v51
    _ = v6
}
