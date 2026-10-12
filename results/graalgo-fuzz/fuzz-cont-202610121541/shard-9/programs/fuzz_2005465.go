package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 string; F2 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 int }
func genFn0(p0 string) uint32 {
    var v6 bool = (!(987 > 190))
    fmt.Println(v6)
    if (v6 || false) {
        v7 := GenS1{F0: p0}
        var v8 uint8 = 1
        _ = v7
        _ = v8
    } else {
        var v9 float32 = float32(0.0)
        _ = v9
    }
    v6 = ((42 > 952) && v6)
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 string, p1 uint, p2 uint64) uint8 {
    if (3.14159 <= ((-2.25 - float64(1e10)) - 0.1)) {
        p0 = p0
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint8 = 255
            _ = fi10
            _ = v11
        }
    } else {
        fmt.Println(p0, p2, p1)
        var v12 string = ("\n\t" + ("\n\t" + (p0 + ("a" + "\n\t"))))
        _ = v12
    }
    p0 = p0
    if true {
        switch 127 {
        case 32767:
            p2 = p2
        case 700:
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := []int{}
                switch p0 {
                case "hello":
                    var v15 uint8 = 1
                    _ = v15
                case "\n\t":
                    for fi16 := 0; fi16 < 10; fi16++ {
                        if (((("\n\t" + "a") + "a") + (p0 + (((((((p0 + p0) + p0) + "foo bar") + (p0 + p0)) + p0) + (("" + (p0 + p0)) + "a")) + ((("\n\t" + p0) + p0) + ((p0 + "hello") + ((p0 + p0) + (p0 + (p0 + (p0 + ""))))))))) != (p0 + "\n\t")) {
                            var v17 float32 = 1.5
                            _ = v17
                        } else {
                            if (-117 <= -21) {
                                p1 = 42
                                v18 := GenS1{F0: ""}
                                if (!(p1 == (^(p1)))) {
                                    p0 = (((p0 + "\n\t") + "foo bar") + ((("\n\t" + p0) + p0) + p0))
                                }
                                _ = v18
                            } else {
                                var v19 float64 = 3.14159
                                v20 := GenS2{F0: 2}
                                _ = v19
                                _ = v20
                            }
                            v21 := GenS1{F0: (("\n\t" + "a") + p0)}
                            p2 = p2
                            _ = v21
                        }
                        p2 = 799
                        v14 = append(v14, (fi13 >> (((p1 + (p1 / ((p1) | 1)))) & 63)))
                        _ = fi16
                    }
                }
                _ = fi13
                _ = v14
            }
        case 9223372036854775807:
            if true {
                p0 = p0
            } else {
                p1 -= (p1 ^ p1)
                p2 = 42
            }
        }
    }
    if (54 >= 614) {
        v22 := GenS3{F0: 2}
        _ = v22
    } else {
        v23 := []int{-1, 1, 0}
        v24 := GenS1{F0: (p0 + p0)}
        for fi25 := 0; fi25 < 20; fi25++ {
            for fi26 := 0; fi26 < 1; fi26++ {
                switch p0 {
                case "\n\t":
                    v27 := GenS4{F0: 9223372036854775807, F1: ("hello" + (((p0 + (p0 + ((((("" + p0) + "foo bar") + p0) + "\n\t") + "hello"))) + "foo bar") + p0)), F2: 1}
                    _ = v27
                default:
                    v28 := GenS5{F0: 0}
                    _ = v28
                }
                _ = fi26
            }
            _ = fi25
        }
        _ = v23
        _ = v24
    }
    switch ("" + "a") {
    case "\n\t":
        var v29 int8 = 1
        _ = v29
    case "a":
        var v30 bool = true
        _ = v30
    case "foo bar":
        var v31 string = p0
        _ = v31
    }
    _ = p0
    _ = p1
    _ = p2
    return 40
}

func main() {
    var v1 int = -2147483648
    var v2 int64 = -1
    v3 := make(map[uint32]bool)
    v4 := []string{"a", ("foo bar" + "hello"), ("hello" + ("foo bar" + "\n\t")), "a"}
    v5 := []uint32{65535, 2147483648, 537}
    switch v1 {
    case 42:
        var v32 float64 = ((0.0 + ((((1.5 + ((float64(1.5) - -2.25) - 1e10)) + (1.5 * 3.14159)) + ((-2.25 + (3.14159 * ((3.14159 * 1.5) / 1e10))) / (((-2.25 * 1.5) + 1.5) * 0.1))) / 1e10)) + 0.0)
        _ = v32
    case -9223372036854775808:
        if true {
            v1 -= v1
            var v33 uint8 = 1
            v34 := GenS4{F0: v2, F1: ("foo bar" + "foo bar"), F2: 1}
            _ = v33
            _ = v34
        } else {
            var v35 string = "hello"
            v36 := GenS5{F0: 2}
            var v37 float64 = 1.5
            _ = v35
            _ = v36
            _ = v37
        }
    }
    fmt.Println(v1, v1)
    v2 = v2
    var v38 bool = true
    var v39 int = 32767
    switch (v39 << ((18446744073709551615) & 63)) {
    case -793:
        var v40 float64 = -2.25
        _ = v40
    default:
        v3[585] = (genFn0("foo bar") != 2)
    }
    v41 := GenS6{F0: 2, F1: -1}
    v2 += (v2 >> ((143) & 63))
    for fi42 := 0; fi42 < 3; fi42++ {
        var v43 int64 = v2
        var v44 uint64 = 700
        _ = fi42
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v5
}
