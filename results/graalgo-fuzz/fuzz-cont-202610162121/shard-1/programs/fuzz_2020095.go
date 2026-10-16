package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint32, p1 bool, p2 int64) string {
    fmt.Println(p1, p0)
    p2 -= -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}

func main() {
    var v1 float32 = ((-2.25 + 0.1) - (((1.5 / (1.5 - -2.25)) * 3.14159) / (0.1 + float32((0.1 - 3.14159)))))
    var v2 int16 = 0
    var v3 float64 = (float64((v1 + (((0.1 + (v1 / 1.5)) - (1.5 + 1e10)) + v1))) / (((3.14159 - 3.14159) + (1.5 - ((float64(-2.25) + (0.0 * 0.0)) - (0.1 * 0.0)))) - -2.25))
    var v4 uint16 = 178
    v5 := []string{}
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint = 2
        switch ("a" + "") {
        case "":
            v8 := GenS1{F0: 214, F1: int64(v7), F2: 4294967295}
            _ = v8
        case "a":
            for fi9 := 0; fi9 < 6; fi9++ {
                fmt.Println(v7, fi6)
                var v10 int16 = v2
                if ((v2 - (v2 + -32768)) >= (v2 << ((v7) & 63))) {
                    var v11 int64 = -9223372036854775808
                    _ = v11
                } else {
                    if false {
                        v12 := GenS2{F0: (v7 * v7), F1: 1}
                        if false {
                            var v13 uint8 = 127
                            if ("" == "foo bar") {
                                if (false && (fi9 <= fi6)) {
                                    v14 := GenS1{F0: v7, F1: 0, F2: 18446744073709551615}
                                    _ = v14
                                } else {
                                    var v15 uint16 = v4
                                    var v16 int16 = 0
                                    v15 = (^(v4))
                                    _ = v15
                                    _ = v16
                                }
                                v1 += v1
                                var v17 float64 = v3
                                if (int16(v4) != v10) {
                                    if ((false || (false || false)) && (false && true)) {
                                        if (false || (true || ((v17 * float64((float32(v3) - float32((v3 - (float64(v1) + (v3 + v17))))))) < v3))) {
                                            v18 := GenS1{F0: (v7 * v7), F1: -1, F2: 106}
                                            fmt.Println(v17)
                                            var v19 float32 = v1
                                            _ = v18
                                            _ = v19
                                        } else {
                                            v20 := GenS2{F0: v7, F1: 829}
                                            var v21 uint8 = v13
                                            fmt.Println(v21, v7)
                                            if false {
                                                if (float32(float64(1e10)) > (v1 + v1)) {
                                                    if len(v5) > 0 {
                                                        v5[0] = ("\n\t" + "a")
                                                    }
                                                }
                                                var v22 float64 = 0.1
                                                switch "hello" {
                                                case "":
                                                    v23 := GenS1{F0: (v7 * (^(v7))), F1: -407, F2: 1}
                                                    _ = v23
                                                }
                                                _ = v22
                                            }
                                            _ = v20
                                            _ = v21
                                        }
                                    } else {
                                        var v24 string = "foo bar"
                                        v25 := GenS1{F0: 687, F1: 1, F2: 1}
                                        v26 := GenS3{F0: v13}
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    }
                                    v5 = append(v5, "foo bar")
                                }
                                var v27 float32 = (v1 + v1)
                                _ = v17
                                _ = v27
                            } else {
                                v5 = append(v5, "foo bar")
                                var v28 int = 42
                                _ = v28
                            }
                            _ = v13
                        }
                        _ = v12
                    } else {
                        var v29 int32 = -1
                        v30 := GenS4{F0: 1, F1: 1}
                        _ = v29
                        _ = v30
                    }
                    switch (fi6 | 128) {
                    case -9223372036854775808:
                        fmt.Println(v10, v10)
                    }
                }
                _ = fi9
                _ = v10
            }
        }
        _ = fi6
        _ = v7
    }
    v2 = (-(v2))
    if (!(v4 <= (v4 + 2))) {
        if false {
            for fi31 := 0; fi31 < 12; fi31++ {
                var v32 int32 = -412
                v2 = (v2 & v2)
                var v33 int = (fi31 - (fi31 | fi31))
                _ = fi31
                _ = v32
                _ = v33
            }
            var v34 uint16 = 1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v34
        }
    } else {
        var v35 float64 = (v3 + v3)
        var v36 int32 = 1
        v3 = float64(3.14159)
        _ = v35
        _ = v36
    }
    if true {
        var v37 int8 = 127
        _ = v37
    } else {
        var v38 bool = (!(false && ((false || ((2 > 0) && ("" == "hello"))) || (587 >= 828))))
        v38 = (v38 && (v38 && v38))
        v2 = v2
        _ = v38
    }
    v39 := GenS3{F0: 157}
    fmt.Println(v1, v1, v3)
    v40 := GenS5{F0: -94}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v41 int32 = 695
    v2 -= (v2 * -32768)
    fmt.Println(v1, v2)
    v4 += v4
    for fi42 := 0; fi42 < 6; fi42++ {
        v4 = 2
        v43 := GenS3{F0: 228}
        if (65535 <= 42) {
            fmt.Println(v1, fi42)
            switch "hello" {
            case "a":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "\n\t":
                v44 := GenS6{F0: 632}
                _ = v44
            }
            var v45 float64 = ((0.1 + 1e10) + -2.25)
            _ = v45
        }
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
