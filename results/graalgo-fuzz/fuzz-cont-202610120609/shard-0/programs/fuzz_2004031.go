package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 string; F2 uint16 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint16, p1 int64, p2 int8) string {
    v4 := GenS1{F0: 2147483647, F1: -1, F2: (p1 & 845)}
    fmt.Println(p0, p1, p1)
    v5 := []int{2147483647, int(p1), 128}
    v5 = append(v5, 128)
    var v6 int64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return "hello"
}
func genFn1(p0 uint16) uint16 {
    switch "hello" {
    case "\n\t":
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 14; fi8++ {
                switch ((((((("foo bar" + "") + "a") + ("" + "a")) + "a") + "a") + "hello") + ("hello" + ("hello" + (("foo bar" + "") + "")))) {
                case "hello":
                    var v9 int16 = 1000
                    _ = v9
                case "\n\t":
                    var v10 string = "foo bar"
                    _ = v10
                default:
                    switch "" {
                    case "foo bar":
                        fmt.Println(fi8, fi8)
                    case "\n\t":
                        switch ("foo bar" + "") {
                        case "foo bar":
                            var v11 uint = 1
                            p0 = (^(p0))
                            _ = v11
                        case "":
                            var v12 uint8 = 127
                            _ = v12
                        case "a":
                            v13 := GenS2{F0: 0}
                            _ = v13
                        default:
                            fmt.Println(fi8)
                        }
                    }
                }
                switch ("" + "hello") {
                case "hello":
                    switch "\n\t" {
                    case "":
                        v14 := GenS2{F0: 623}
                        _ = v14
                    case "\n\t":
                        var v15 string = "foo bar"
                        _ = v15
                    default:
                        var v16 string = "a"
                        _ = v16
                    }
                default:
                    v17 := []int{(fi8 & fi7), fi7}
                    _ = v17
                }
                _ = fi8
            }
            var v18 string = (("\n\t" + (("a" + "hello") + ("\n\t" + "hello"))) + "\n\t")
            _ = fi7
            _ = v18
        }
    case "foo bar":
        var v19 int32 = 2147483647
        _ = v19
    case "":
        v20 := GenS3{F0: false}
        _ = v20
    }
    v21 := GenS4{F0: 1, F1: "hello", F2: 716}
    var v22 string = ("\n\t" + ("hello" + (("hello" + "a") + ("\n\t" + ""))))
    fmt.Println(p0, p0)
    _ = p0
    _ = v21
    _ = v22
    return (p0 ^ (p0 + p0))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 1.5
    v3 := make(map[string]bool)
    switch "a" {
    case "":
        if ((0 != -961) || false) {
            for fi23 := 0; fi23 < 1; fi23++ {
                v2 -= (v2 * v2)
                fmt.Println(fi23)
                _ = fi23
            }
        } else {
            v2 = (0.1 - v2)
            v24 := v3[(("hello" + "\n\t") + "")]
            if v24 {
                v2 = 1.5
                v24 = (v24 || (65535 <= 751))
            }
            _ = v24
        }
    case "a":
        v2 += -2.25
    }
    v25 := GenS4{F0: 42, F1: "foo bar", F2: genFn1(1)}
    var v26 uint32 = 65535
    v27 := GenS2{F0: 2}
    if false {
        if true {
            v2 -= ((v2 * -2.25) * 1e10)
            v28 := []int{-176, 2147483647}
            _ = v28
        }
        for fi29 := 0; fi29 < 13; fi29++ {
            if true {
                for fi30 := 0; fi30 < 5; fi30++ {
                    var v31 int = fi29
                    var v32 uint64 = 18446744073709551615
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            } else {
                var v33 uint64 = 18446744073709551615
                _ = v33
            }
            for fi34 := 0; fi34 < 16; fi34++ {
                fmt.Println(v26)
                fmt.Println(fi29)
                _ = fi34
            }
            var v35 bool = (true && (372 < 57))
            _ = fi29
            _ = v35
        }
        var v36 float64 = v2
        _ = v36
    } else {
        switch "a" {
        case "a":
            var v37 float64 = (v2 + v2)
            _ = v37
        case "":
            fmt.Println(v2, v2)
        case "foo bar":
            switch "foo bar" {
            case "foo bar":
                for fi38 := 0; fi38 < 3; fi38++ {
                    v2 = v2
                    fmt.Println(fi38, fi38)
                    _ = fi38
                }
            }
        default:
            v39 := GenS2{F0: 0}
            _ = v39
        }
        switch -171 {
        case -9223372036854775808:
            if ((v2 * v2) == v2) {
                v26 = 2147483647
            }
        }
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        v26 = 131
        _ = fi40
    }
    v2 = (float64(-2.25) + -2.25)
    for fi41 := 0; fi41 < 11; fi41++ {
        v42 := GenS5{F0: 1}
        _ = fi41
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
