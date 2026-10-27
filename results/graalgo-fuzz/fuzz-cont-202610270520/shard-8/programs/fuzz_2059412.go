package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int8) int32 {
    if true {
        var v7 uint16 = p0
        _ = v7
    } else {
        var v8 uint = 18446744073709551615
        if true {
            v9 := GenS1{F0: -2147483648, F1: true, F2: 2}
            _ = v9
        } else {
            v10 := GenS1{F0: 2147483647, F1: false, F2: 4294967295}
            _ = v10
        }
        if false {
            v8 = (v8 + v8)
            var v11 uint64 = 1
            _ = v11
        }
        _ = v8
    }
    p1 = 0
    _ = p0
    _ = p1
    return 2147483647
}
func genFn1(p0 bool) uint8 {
    p0 = false
    p0 = p0
    _ = p0
    return 0
}

func main() {
    var v1 int64 = 42
    v2 := []int{32767, -2147483648, 103, 32767}
    var v3 float32 = (1.5 + ((1e10 + 0.0) / (1e10 + float32(float64(3.14159)))))
    v4 := []string{"foo bar"}
    var v5 float32 = (v3 / v3)
    var v6 uint16 = 255
    switch (("a" + ("a" + "a")) + "foo bar") {
    case "hello":
        if false {
            fmt.Println(v3)
            var v12 float64 = (((((((3.14159 - 1.5) - 1e10) - 0.0) * (float64((1e10 / v5)) + -2.25)) / 1.0) / 1.5) * ((float64((((3.14159 + v5) * v3) - float32((1e10 - ((1e10 / 1.0) + 0.0))))) - (1e10 / float64(float32(3.14159)))) / 1.5))
            var v13 int = 404
            _ = v12
            _ = v13
        }
    case "\n\t":
        switch 127 {
        case 32768:
            if false {
                var v14 bool = true
                v4 = append(v4, "a")
                _ = v14
            }
        }
    default:
        switch (("\n\t" + "hello") + "\n\t") {
        case "a":
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 int64 = (v1 * 0)
                _ = fi15
                _ = v16
            }
        case "foo bar":
            for fi17 := 0; fi17 < 13; fi17++ {
                switch "a" {
                case "":
                    if true {
                        v1 = v1
                    } else {
                        v5 = 3.14159
                    }
                case "a":
                    if true {
                        v3 = v5
                        if true {
                            v18 := GenS2{F0: 2}
                            _ = v18
                        } else {
                            var v19 int8 = 100
                            _ = v19
                        }
                    }
                case "hello":
                    switch "\n\t" {
                    case "hello":
                        v6 = (v6 - (v6 | 0))
                    case "a":
                        v6 = v6
                    case "\n\t":
                        fmt.Println(v3)
                    }
                default:
                    if ((fi17 - fi17) >= int(v6)) {
                        _ = len(v4)
                        var v20 float64 = 3.14159
                        v1 = (v1 >> ((393) & 63))
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 18; fi21++ {
                            if true {
                                var v22 int16 = 32767
                                _ = v22
                            } else {
                                if true {
                                    var v23 string = ("" + ("hello" + ("" + "hello")))
                                    _ = v23
                                }
                                v1 = v1
                            }
                            switch fi21 {
                            case 42:
                                if false {
                                    var v24 string = "foo bar"
                                    var v25 uint16 = v6
                                    var v26 int8 = -128
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                            case 128:
                                v3 -= -2.25
                            }
                            _ = fi21
                        }
                        if (!((0.0 - (3.14159 + 1.5)) >= (float64(((v5 * ((v5 / v5) - 0.0)) + ((v3 * (((-2.25 * v5) - 0.1) / ((-2.25 / v5) - v3))) - v3))) + (((3.14159 + 1.5) + (0.1 / ((0.0 + (float64(((1e10 / (3.14159 / -2.25)) - 1.5)) / 1e10)) + 1.5))) * 0.0)))) {
                            var v27 bool = ((v6 * (v6 % (((v6 >> ((0) & 63))) | 1))) <= 32767)
                            _ = v27
                        }
                        switch (fi17 & (fi17 + fi17)) {
                        case 365:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case -476:
                            _ = len(v2)
                        }
                    }
                }
                _ = fi17
            }
        case "hello":
            v3 += (0.0 - (1e10 / float32(float64(v5))))
        }
    }
    v3 = float32(3.14159)
    switch 306 {
    case 4294967295:
        _ = len(v2)
    default:
        v28 := GenS3{F0: (v6 / (((v6 / (((v6 % ((v6) | 1))) | 1))) | 1)), F1: v6, F2: 0}
        _ = v28
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        if (((42 < 0) && (false || ((!(65535 <= 4294967295)) && ((!(!(100 != -126))) || true)))) && (-402 <= 1)) {
            var v30 string = "foo bar"
            var v31 int = (^(fi29))
            _ = v30
            _ = v31
        }
        if (799 == 168) {
            v1 += 1
        }
        if false {
            var v32 string = ("a" + "foo bar")
            if false {
                v33 := GenS4{F0: 245}
                _ = v33
            } else {
                v4 = append(v4, (v32 + (((("foo bar" + v32) + ((v32 + "foo bar") + (v32 + v32))) + v32) + v32)))
            }
            _ = v32
        } else {
            _ = len(v2)
            _ = len(v2)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi29
    }
    var v34 uint64 = 70
    v1 = v1
    v2 = append(v2, -270)
    switch "hello" {
    case "":
        v35 := GenS1{F0: 42, F1: ((true && true) || true), F2: (v34 >> ((458) & 63))}
        _ = v35
    case "a":
        if false {
            switch "foo bar" {
            case "foo bar":
                var v36 float64 = 1e10
                _ = v36
            case "":
                var v37 string = ("" + "")
                _ = v37
            case "\n\t":
                v6 += v6
            default:
                v1 = 1
            }
            if len(v2) > 0 {
                v2[0] = -243
            }
            v38 := GenS5{F0: 2}
            _ = v38
        }
    }
    fmt.Println(v3, v1)
    var v39 int16 = -781
    fmt.Println(v34, v34)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
