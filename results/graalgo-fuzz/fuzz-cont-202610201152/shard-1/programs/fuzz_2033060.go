package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int; F1 uint32; F2 uint8 }
func genFn0() uint8 {
    var v6 bool = true
    v6 = v6
    var v7 int64 = 0
    _ = v6
    _ = v7
    return 244
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 0
    v3 := []string{("\n\t" + ("" + "hello")), "foo bar", (("a" + ("foo bar" + (("hello" + "hello") + "foo bar"))) + "")}
    v4 := []uint32{}
    var v5 uint32 = (v2 ^ v2)
    var v8 int16 = 1000
    var v9 float32 = 1.5
    switch 42 {
    case 128:
        switch "foo bar" {
        case "":
            if false {
                v9 += v9
                var v10 uint64 = 42
                _ = v10
            }
        case "a":
            v9 = (v9 / (v9 + (v9 * v9)))
        case "hello":
            if (v8 != (v8 ^ (-(v8)))) {
                v5 -= 4294967295
                var v11 bool = (!(1.5 == (0.0 + 0.1)))
                _ = v11
            }
        default:
            var v12 uint32 = (v2 << ((18446744073709551615) & 63))
            _ = v12
        }
    default:
        var v13 uint8 = 0
        _ = v13
    }
    if (!false) {
        v14 := GenS1{F0: 9223372036854775807, F1: ((v9 * v9) >= 0.1)}
        _ = len(v1)
        _ = v14
    } else {
        v9 += v9
    }
    var v15 int16 = v8
    if (100 != 100) {
        v16 := GenS1{F0: 9223372036854775807, F1: ((!(!true)) && true)}
        if (167 != 65535) {
            var v17 int16 = (v15 | 158)
            var v18 int8 = 117
            _ = v17
            _ = v18
        }
        v9 -= -2.25
        _ = v16
    } else {
        fmt.Println(v9, v9)
        var v19 float32 = ((((((1.5 + (-2.25 / 3.14159)) - v9) / (v9 + v9)) - (1.5 / -2.25)) / -2.25) - float32((-2.25 / ((1.5 * (((((0.1 + (1.5 / 1e10)) * 3.14159) / 0.1) - 0.1) * (1.5 / (float64(v9) / 0.1)))) * (-2.25 - 0.0)))))
        _ = v19
    }
    v1["a"] = "hello"
    if (v8 != -1) {
        if (!(uint16(v2) >= 65535)) {
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 bool = false
                var v22 int8 = 127
                if (!((fi20 & 168) != int(v22))) {
                    v23 := GenS1{F0: -9223372036854775808, F1: false}
                    _ = v23
                } else {
                    var v24 string = (("a" + ("a" + "a")) + "foo bar")
                    _ = v24
                }
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        _ = len(v1)
    }
    _ = len(v1)
    var v25 uint8 = genFn0()
    fmt.Println(v5, v8)
    if (false && (v9 > (v9 - v9))) {
        var v26 int16 = 1
        _ = v26
    } else {
        switch ("\n\t" + "a") {
        case "hello":
            if (true || (true || true)) {
                for fi27 := 0; fi27 < 12; fi27++ {
                    if true {
                        v28 := GenS1{F0: -9223372036854775808, F1: true}
                        for fi29 := 0; fi29 < 11; fi29++ {
                            v9 -= 1e10
                            v8 = (v8 + v8)
                            _ = fi29
                        }
                        _ = v28
                    } else {
                        v25 += v25
                        var v30 int8 = 127
                        switch "" {
                        case "foo bar":
                            v31 := GenS2{F0: (v25 % (((v25 / ((29) | 1))) | 1)), F1: (v5 - v2)}
                            _ = v31
                        case "\n\t":
                            fmt.Println(fi27)
                        case "":
                            var v32 int8 = 42
                            _ = v32
                        }
                        _ = v30
                    }
                    var v33 uint32 = (v2 | 1)
                    _ = fi27
                    _ = v33
                }
                v34 := GenS3{F0: 643, F1: 855, F2: 2}
                _ = v34
            } else {
                var v35 int64 = -962
                v36 := GenS2{F0: v25, F1: 42}
                v37 := GenS4{F0: 128, F1: (v2 << ((854) & 63)), F2: v25}
                _ = v35
                _ = v36
                _ = v37
            }
        case "foo bar":
            var v38 string = ("foo bar" + "")
            _ = v38
        default:
            var v39 float32 = (v9 + 3.14159)
            _ = v39
        }
        v40 := GenS3{F0: 4294967295, F1: 6, F2: 74}
        var v41 uint8 = (v25 * genFn0())
        _ = v40
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
