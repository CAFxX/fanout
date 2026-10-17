package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int32; F2 int8 }
type GenS5 struct { F0 int64; F1 string }
func genFn0(p0 int64, p1 uint16, p2 int16) string {
    fmt.Println(p1)
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := []int{fi6, (-(fi6)), 2147483647, 32768}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return ("" + ("hello" + "foo bar"))
}
func genFn1(p0 int64, p1 int64, p2 int32) int64 {
    v8 := []int{128, 42, -2147483648, 42}
    v8 = append(v8, 127)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p0 | 42)
}

func main() {
    v1 := []string{"hello", "a", "a"}
    var v2 string = ("" + "")
    var v3 float64 = 3.14159
    var v4 int32 = 38
    var v5 uint = 259
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: true, F1: (("foo bar" + v2) + (v2 + ((v2 + v2) + v2)))}
        _ = fi9
        _ = v10
    }
    if (!(!((-863 != genFn1(42, 9223372036854775807, 42)) && (false || false)))) {
        v2 = v2
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 bool = (-120 < int64(v5))
            _ = fi11
            _ = v12
        }
    }
    if (!(!true)) {
        if len(v1) > 0 {
            v1[0] = (("" + v2) + v2)
        }
    } else {
        if (255 < 0) {
            for fi13 := 0; fi13 < 1; fi13++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch "a" {
                case "hello":
                    if true {
                        _ = len(v1)
                        for fi14 := 0; fi14 < 9; fi14++ {
                            v4 += v4
                            if len(v1) > 0 {
                                v1[0] = (v2 + v2)
                            }
                            _ = fi14
                        }
                    }
                case "a":
                    v15 := GenS2{F0: 458}
                    _ = v15
                }
                _ = fi13
            }
            var v16 uint8 = 188
            if (true || false) {
                fmt.Println(v4, v2)
                switch 583 {
                case 32768:
                    v3 += (0.1 - 1.5)
                }
            } else {
                v2 = ("a" + ((v2 + (v2 + v2)) + ""))
                v17 := GenS2{F0: -507}
                if ((v2 + "hello") >= (v2 + ((v2 + v2) + (v2 + v2)))) {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        switch (fi18 + fi18) {
                        case 32767:
                            v19 := GenS2{F0: -1}
                            _ = v19
                        case 42:
                            if false {
                                if true {
                                    for fi20 := 0; fi20 < 5; fi20++ {
                                        v5 = v5
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v21 int = -759
                                        _ = fi20
                                        _ = v21
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v16 = v16
                                } else {
                                    v22 := GenS3{F0: 1}
                                    v2 = "foo bar"
                                    v4 = v4
                                    _ = v22
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        case 9223372036854775807:
                            for fi23 := 0; fi23 < 15; fi23++ {
                                if false {
                                    switch v2 {
                                    case "a":
                                        var v24 uint8 = (^(v16))
                                        _ = v24
                                    case "hello":
                                        var v25 float32 = 1e10
                                        _ = v25
                                    case "foo bar":
                                        switch v2 {
                                        case "a":
                                            fmt.Println(fi18, fi18, fi23)
                                        case "foo bar":
                                            v26 := GenS4{F0: (true && (9223372036854775807 <= -9223372036854775808)), F1: v4, F2: -1}
                                            _ = v26
                                        case "\n\t":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        default:
                                            var v27 int64 = -9223372036854775808
                                            _ = v27
                                        }
                                    }
                                    fmt.Println(v4, v3, v5)
                                } else {
                                    var v28 uint64 = 42
                                    switch fi23 {
                                    case 0:
                                        if (!(!(-2147483648 > v4))) {
                                            v16 -= v16
                                        } else {
                                            var v29 float64 = float64(0.1)
                                            var v30 int8 = -128
                                            v16 += v16
                                            _ = v29
                                            _ = v30
                                        }
                                    case -47:
                                        v28 += v28
                                    default:
                                        var v31 float32 = 1e10
                                        _ = v31
                                    }
                                    var v32 float64 = float64((((1e10 - 3.14159) / 1e10) * 0.0))
                                    _ = v28
                                    _ = v32
                                }
                                _ = fi23
                            }
                        }
                        fmt.Println(v3, v5, v16)
                        v33 := GenS4{F0: ((v5 - v5) != (^(v5))), F1: v4, F2: 100}
                        _ = fi18
                        _ = v33
                    }
                    v1 = append(v1, "foo bar")
                    var v34 uint = v5
                    _ = v34
                } else {
                    if true {
                        if true {
                            if (!true) {
                                var v35 float32 = (-2.25 - 1.5)
                                _ = v35
                            }
                            var v36 uint = 42
                            _ = v36
                        }
                        v1 = append(v1, (v2 + "\n\t"))
                    }
                    v4 = (v4 ^ (v4 << (((v5 ^ 42)) & 63)))
                }
                _ = v17
            }
            _ = v16
        }
    }
    fmt.Println(v4, v3)
    for fi37 := 0; fi37 < 3; fi37++ {
        fmt.Println(v4)
        switch ("a" + v2) {
        case "a":
            var v38 int16 = 32767
            _ = v38
        case "foo bar":
            switch "a" {
            case "\n\t":
                if false {
                    switch v2 {
                    case "\n\t":
                        v39 := GenS5{F0: 42, F1: (v2 + v2)}
                        _ = v39
                    default:
                        v40 := GenS1{F0: true, F1: (((v2 + v2) + "hello") + ((("" + ((v2 + v2) + "")) + (v2 + v2)) + ("" + ("\n\t" + (v2 + "hello")))))}
                        _ = v40
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            case "hello":
                fmt.Println(v3, fi37)
            case "":
                switch (-(fi37)) {
                case 74:
                    v5 -= (v5 ^ v5)
                }
            }
        }
        _ = fi37
    }
    v4 = (^(v4))
    fmt.Println(v5)
    var v41 bool = true
    if v41 {
        var v42 float32 = 1e10
        _ = v42
    }
    var v43 float32 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v5
}
