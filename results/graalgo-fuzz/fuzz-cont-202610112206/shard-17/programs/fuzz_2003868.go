package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint32; F1 int8; F2 int }
func genFn0(p0 uint32) string {
    var v5 uint64 = 42
    for fi6 := 0; fi6 < 16; fi6++ {
        v5 = (^(v5))
        _ = fi6
    }
    _ = p0
    _ = v5
    return (("foo bar" + "\n\t") + ("foo bar" + ("\n\t" + ("a" + ("" + "a")))))
}

func main() {
    var v1 uint64 = 0
    v2 := []string{("hello" + "a"), "a", "foo bar"}
    v3 := make(map[string]int)
    var v4 uint8 = 189
    if (9 <= 42) {
        for fi7 := 0; fi7 < 17; fi7++ {
            delete(v3, "foo bar")
            var v8 int64 = 42
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := v3[("foo bar" + ("hello" + "a"))]
                var v11 uint8 = 255
                var v12 uint64 = (v1 | (v1 / ((v1) | 1)))
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi7
            _ = v8
        }
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            fmt.Println(v4, v1, fi13)
            for fi14 := 0; fi14 < 20; fi14++ {
                v4 -= v4
                v1 = (^(v1))
                fmt.Println(v1, v1)
                _ = fi14
            }
            _ = fi13
        }
        v15 := v3["hello"]
        _ = v15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        v1 = 18446744073709551615
        var v17 float32 = 1e10
        switch (("hello" + "") + "") {
        case "\n\t":
            for fi18 := 0; fi18 < 5; fi18++ {
                v1 = (v1 << ((939) & 63))
                v17 -= v17
                _ = fi18
            }
        case "":
            if (true && true) {
                for fi19 := 0; fi19 < 7; fi19++ {
                    if (-811 == int16(v1)) {
                        if ((v1 ^ (v1 / ((v1) | 1))) >= v1) {
                            v20 := GenS1{F0: -1}
                            _ = v20
                        }
                        var v21 float64 = float64((v17 - v17))
                        v22 := GenS1{F0: -9223372036854775808}
                        _ = v21
                        _ = v22
                    }
                    v1 = v1
                    switch fi16 {
                    case 0:
                        var v23 string = "a"
                        _ = v23
                    case 1:
                        if (-2.25 <= 1.5) {
                            v4 -= (^(v4))
                        } else {
                            v4 -= v4
                        }
                    case 32768:
                        v24 := GenS1{F0: -613}
                        _ = v24
                    }
                    v25 := GenS1{F0: -1}
                    var v26 int16 = 236
                    _ = fi19
                    _ = v25
                    _ = v26
                }
                var v27 int = (fi16 >> ((65535) & 63))
                _ = v27
            }
        case "a":
            var v28 int8 = -128
            _ = v28
        }
        _ = fi16
        _ = v17
    }
    v4 = 92
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 int64 = 612
        v31 := GenS1{F0: v30}
        var v32 int8 = -4
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    v4 += 42
    v33 := GenS1{F0: 42}
    fmt.Println(v1, v4)
    fmt.Println(v4)
    v34 := GenS2{F0: -1, F1: ((float64((((3.14159 + (-2.25 + 1e10)) * 0.1) * ((0.1 + 1.5) / 1e10))) + 1.5) >= 3.14159)}
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 uint64 = v1
        fmt.Println(v4)
        _ = fi35
        _ = v36
    }
    switch 788 {
    case -2147483648:
        v4 += v4
    case 375:
        if ((^(v1)) >= (v1 >> ((255) & 63))) {
            if (2147483647 < -591) {
                var v37 float64 = (0.1 / (-2.25 - 1e10))
                if (!(42 == -423)) {
                    var v38 uint32 = 42
                    _ = v38
                }
                if ((v37 + (float64(float32(v37)) - v37)) > v37) {
                    if (0 >= 0) {
                        var v39 uint32 = 4294967295
                        _ = v39
                    }
                    if (100 > 1) {
                        var v40 uint8 = uint8(v1)
                        v41 := GenS2{F0: 1, F1: (true && false)}
                        switch 2147483647 {
                        case 127:
                            fmt.Println(v40, v37, v37)
                        default:
                            v4 = v4
                        }
                        _ = v40
                        _ = v41
                    }
                } else {
                    for fi42 := 0; fi42 < 20; fi42++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi42
                    }
                }
                _ = v37
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v43 uint32 = 42
                _ = v43
            }
        } else {
            v4 = 253
            if false {
                if true {
                    v1 = 2
                    fmt.Println(v1)
                    if (false || false) {
                        v4 -= v4
                    }
                } else {
                    var v44 bool = (169 != 1)
                    v45 := GenS2{F0: -317, F1: v44}
                    v4 = (^(v4))
                    _ = v44
                    _ = v45
                }
            }
            v46 := GenS3{F0: 2147483648, F1: -128, F2: -10}
            _ = v46
        }
    case -757:
        v1 += (^(v1))
    default:
        delete(v3, ("a" + (((("a" + "a") + "\n\t") + "") + ((((("a" + "\n\t") + "foo bar") + ("hello" + "\n\t")) + "hello") + "\n\t"))))
    }
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
