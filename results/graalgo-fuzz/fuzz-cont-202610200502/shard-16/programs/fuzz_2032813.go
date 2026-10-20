package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 int16; F1 uint; F2 int16 }
type GenS6 struct { F0 uint8 }
func genFn0() int32 {
    if (68 != -128) {
        var v4 uint32 = 2147483647
        _ = v4
    } else {
        v5 := []int{-868, 1, 42}
        if false {
            var v6 uint32 = 42
            v6 = (^(v6))
            v6 = (v6 * v6)
            _ = v6
        } else {
            var v7 uint = 374
            _ = v7
        }
        var v8 int64 = 9223372036854775807
        _ = v5
        _ = v8
    }
    if ((3.14159 * 1.5) < float64(1e10)) {
        var v9 int32 = 42
        var v10 string = (((((("\n\t" + (("\n\t" + "\n\t") + ("foo bar" + ((("a" + "hello") + "\n\t") + "hello")))) + "foo bar") + "foo bar") + ("foo bar" + "\n\t")) + "a") + "hello")
        _ = v10
        _ = v9
    }
    var v11 float64 = float64(0.1)
    var v12 uint8 = 127
    if (!((0.0 - float32(0.0)) == 1.5)) {
        v12 += v12
        if (!((!true) && false)) {
            v13 := GenS1{F0: 1}
            v12 = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 uint = 0
                _ = fi14
                _ = v15
            }
            var v16 int8 = 0
            _ = v16
        }
        fmt.Println(v12, v12)
    }
    _ = v11
    _ = v12
    return 1
}
func genFn1(p0 bool, p1 uint8) string {
    p0 = true
    if p0 {
        p0 = p0
        if ((1.5 + -2.25) < (1e10 / -2.25)) {
            for fi17 := 0; fi17 < 13; fi17++ {
                v18 := []int{}
                var v19 int64 = 1
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        var v20 uint = 0
        _ = v20
    } else {
        p1 = p1
    }
    _ = p0
    _ = p1
    return "a"
}

func main() {
    v1 := []string{(("a" + ("\n\t" + "hello")) + "a"), ("hello" + "a"), ("a" + ("\n\t" + "foo bar")), "foo bar"}
    var v2 uint64 = 42
    var v3 string = ("hello" + "foo bar")
    v21 := GenS2{F0: 0, F1: 185}
    v3 = ((v3 + (v3 + v3)) + v3)
    fmt.Println(v2)
    for fi22 := 0; fi22 < 6; fi22++ {
        switch fi22 {
        case 127:
            v23 := GenS3{F0: (fi22 + (fi22 >> ((18446744073709551615) & 63))), F1: 2}
            _ = v23
        case -9223372036854775808:
            if (42 < 0) {
                if true {
                    fmt.Println(v2, v3)
                    v2 += (^(v2))
                }
                v3 = v3
                if (165 > 0) {
                    v2 = (v2 - (v2 / ((v2) | 1)))
                } else {
                    switch (v3 + "hello") {
                    case "":
                        for fi24 := 0; fi24 < 3; fi24++ {
                            v25 := GenS4{F0: (v2 / ((315) | 1)), F1: "a"}
                            _ = fi24
                            _ = v25
                        }
                    case "hello":
                        v2 -= v2
                    case "foo bar":
                        for fi26 := 0; fi26 < 4; fi26++ {
                            v3 = "\n\t"
                            v2 += v2
                            switch fi26 {
                            case 192:
                                fmt.Println(v2)
                            case -91:
                                if (1e10 <= ((3.14159 - float64(0.0)) + 1e10)) {
                                    var v27 uint64 = 0
                                    v3 = "hello"
                                    var v28 float32 = 3.14159
                                    _ = v27
                                    _ = v28
                                }
                            }
                            _ = fi26
                        }
                    }
                    var v29 uint32 = 533
                    if (42 != 42) {
                        if len(v1) > 0 {
                            v1[0] = (v3 + "\n\t")
                        }
                        var v30 int32 = -481
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v30
                    }
                    _ = v29
                }
            }
        default:
            switch ((((v3 + (v3 + v3)) + "hello") + v3) + v3) {
            case "\n\t":
                fmt.Println(v2, v2, fi22)
            case "a":
                if (127 > -64) {
                    v2 = 414
                }
            case "hello":
                var v31 int32 = 276
                _ = v31
            }
        }
        _ = fi22
    }
    if (42 > -2147483648) {
        v32 := GenS5{F0: 450, F1: 299, F2: -366}
        v3 = ""
        _ = v32
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v3)
    var v33 int16 = 1000
    if true {
        var v34 uint64 = v2
        v33 = -32768
        switch -246 {
        case -9223372036854775808:
            fmt.Println(v2, v3)
        default:
            switch v3 {
            case "":
                fmt.Println(v3, v34, v3)
            default:
                v35 := GenS6{F0: 2}
                _ = v35
            }
        }
        _ = v34
    }
    fmt.Println(v2, v3)
    var v36 uint8 = 128
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v36
}
