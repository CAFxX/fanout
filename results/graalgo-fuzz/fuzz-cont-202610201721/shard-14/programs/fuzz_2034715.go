package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 bool; F2 int8 }
type GenS7 struct { F0 uint16; F1 int64 }
func genFn0(p0 int8) bool {
    if false {
        var v4 int = -443
        v5 := []int{}
        _ = v4
        _ = v5
    } else {
        var v6 bool = (224 < 42)
        _ = v6
    }
    switch "a" {
    case "foo bar":
        var v7 bool = false
        _ = v7
    default:
        p0 -= p0
    }
    _ = p0
    return false
}

func main() {
    v1 := []uint32{2147483648}
    var v2 float64 = 0.0
    var v3 string = "\n\t"
    var v8 float32 = float32(v2)
    var v9 bool = false
    if (v9 || (288 <= -1)) {
        fmt.Println(v2)
        v10 := GenS1{F0: 955}
        if false {
            var v11 float64 = 1.5
            var v12 uint64 = 2
            _ = len(v1)
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 uint = 42
                v15 := GenS2{F0: 1, F1: ((v14 < v14) || (true || (281 >= 2))), F2: v14}
                _ = fi13
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                if ((float32(0.1) <= (v8 + (float32(0.1) + (v8 - (v8 + (0.1 - float32(v2))))))) && false) {
                    switch (v3 + "a") {
                    case "foo bar":
                        v9 = (2 > 1)
                    case "\n\t":
                        switch 9223372036854775807 {
                        case 32768:
                            for fi17 := 0; fi17 < 2; fi17++ {
                                v18 := GenS2{F0: 4294967295, F1: v9, F2: 667}
                                switch (fi16 >> ((2) & 63)) {
                                case 127:
                                    v8 += ((v8 * v8) - (v8 * ((v8 + v8) + v8)))
                                }
                                var v19 uint64 = 18446744073709551615
                                var v20 uint = 18446744073709551615
                                v19 -= 18446744073709551615
                                _ = fi17
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                        case 42:
                            if ((-2.25 + v8) <= (0.0 + 1e10)) {
                                fmt.Println(v3, fi16, fi16)
                                v9 = true
                            }
                        default:
                            if true {
                                if (v9 || false) {
                                    switch "" {
                                    case "\n\t":
                                        v3 = (("hello" + (("\n\t" + ("a" + "\n\t")) + ((((v3 + v3) + v3) + (v3 + "hello")) + v3))) + ("foo bar" + v3))
                                    case "hello":
                                        for fi21 := 0; fi21 < 13; fi21++ {
                                            v22 := GenS2{F0: 2, F1: v9, F2: 514}
                                            v3 = "a"
                                            _ = fi21
                                            _ = v22
                                        }
                                    case "":
                                        var v23 bool = true
                                        _ = v23
                                    }
                                    v2 = (((v2 * v2) + (float64(0.0) - float64(v8))) * (float64(0.0) - float64(v8)))
                                    var v24 float64 = v2
                                    _ = v24
                                }
                                var v25 uint16 = 0
                                _ = v25
                            }
                        }
                    }
                    fmt.Println(fi16, v3, v3)
                }
                _ = len(v1)
                _ = fi16
            }
        }
        _ = v10
    }
    switch 2147483647 {
    case 128:
        fmt.Println(v8)
    default:
        if (0 != -1) {
            for fi26 := 0; fi26 < 8; fi26++ {
                switch (v3 + v3) {
                case "foo bar":
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                case "":
                    if len(v1) > 0 {
                        v1[0] = 2147483648
                    }
                case "hello":
                    fmt.Println(fi26)
                }
                _ = fi26
            }
            if false {
                v27 := GenS3{F0: 18, F1: -9223372036854775808}
                _ = v27
            }
            if (!v9) {
                var v28 uint32 = 0
                _ = v28
            }
        }
    }
    v9 = false
    v8 = float32(v2)
    for fi29 := 0; fi29 < 10; fi29++ {
        _ = len(v1)
        _ = fi29
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 int64 = -761
        switch v3 {
        case "foo bar":
            v9 = v9
        case "\n\t":
            v3 = ((v3 + "") + v3)
        }
        _ = fi30
        _ = v31
    }
    v8 = (0.1 / v8)
    if (false || v9) {
        var v32 int64 = -1
        _ = v32
    } else {
        v33 := GenS4{F0: 639}
        var v34 float32 = v8
        _ = v33
        _ = v34
    }
    if (v8 > v8) {
        if v9 {
            if (272 == 674) {
                var v35 int8 = 1
                _ = v35
            }
        }
    }
    var v36 float32 = (v8 / float32(((1e10 + (v2 + (v2 / (v2 + -2.25)))) / v2)))
    if (!v9) {
        v1 = append(v1, 2)
        v37 := GenS5{F0: 42}
        v8 = v8
        _ = v37
    } else {
        _ = len(v1)
    }
    var v38 string = ((("" + (("\n\t" + (v3 + "a")) + v3)) + "\n\t") + ("" + ("\n\t" + (((v3 + (("a" + v3) + ((v3 + "") + (v3 + v3)))) + v3) + ""))))
    if ((v3 != v3) && v9) {
        fmt.Println(v36, v3)
    } else {
        var v39 float64 = 3.14159
        switch 42 {
        case 0:
            for fi40 := 0; fi40 < 14; fi40++ {
                switch v3 {
                case "a":
                    for fi41 := 0; fi41 < 3; fi41++ {
                        switch "a" {
                        case "":
                            var v42 uint = 0
                            _ = v42
                        case "hello":
                            v43 := GenS6{F0: 32767, F1: true, F2: 0}
                            _ = v43
                        case "foo bar":
                            v44 := GenS2{F0: 1, F1: (305 != 2147483648), F2: 30}
                            _ = v44
                        }
                        v38 = "\n\t"
                        _ = fi41
                    }
                case "":
                    v2 -= float64((v8 + ((v8 - 0.0) + v8)))
                case "\n\t":
                    var v45 uint = 319
                    _ = v45
                }
                _ = fi40
            }
        case 128:
            var v46 float32 = v8
            _ = v46
        case 4294967295:
            v36 -= (((v36 / v8) + v36) - (((0.0 + (v8 - ((v8 * ((-2.25 - 3.14159) + float32(0.1))) - (v36 / float32(1e10))))) + 3.14159) - (float32(v2) * -2.25)))
        }
        if (-2147483648 <= -962) {
            fmt.Println(v2, v8, v38)
            var v47 int = 0
            _ = v47
        }
        _ = v39
    }
    for fi48 := 0; fi48 < 16; fi48++ {
        var v49 int16 = -32768
        v50 := GenS7{F0: 32767, F1: -1}
        switch v3 {
        case "\n\t":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi48
        _ = v49
        _ = v50
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v8
    _ = v9
}
