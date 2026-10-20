package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint64 }
func genFn0() bool {
    var v4 uint32 = 0
    var v5 bool = true
    v6 := GenS1{F0: 19}
    _ = v4
    _ = v5
    _ = v6
    return false
}
func genFn1() int32 {
    var v7 int8 = 127
    switch "" {
    case "a":
        v8 := GenS2{F0: 42}
        _ = v8
    default:
        if (1 < 0) {
            var v9 uint = 2
            fmt.Println(v7, v7)
            _ = v9
        }
    }
    v7 += -128
    var v10 int16 = -174
    _ = v10
    _ = v7
    return -2147483648
}

func main() {
    v1 := []string{}
    v2 := make(map[string]string)
    var v3 uint = 834
    if false {
        var v11 float64 = 0.0
        v11 -= float64(0.0)
        fmt.Println(v3)
        _ = v11
    } else {
        if (!true) {
            v12 := GenS2{F0: 42}
            _ = v12
        } else {
            if ((true && (!(true && (18446744073709551615 <= 1)))) || false) {
                fmt.Println(v3)
                for fi13 := 0; fi13 < 7; fi13++ {
                    v14 := v2["hello"]
                    _ = fi13
                    _ = v14
                }
            } else {
                if (!true) {
                    var v15 string = "a"
                    switch -240 {
                    case 127:
                        for fi16 := 0; fi16 < 9; fi16++ {
                            v3 = v3
                            var v17 float32 = -2.25
                            _ = fi16
                            _ = v17
                        }
                    default:
                        for fi18 := 0; fi18 < 10; fi18++ {
                            fmt.Println(v15)
                            _ = fi18
                        }
                    }
                    fmt.Println(v3, v3)
                    _ = v15
                }
                var v19 float32 = (1.5 * (float32(1e10) / 3.14159))
                for fi20 := 0; fi20 < 4; fi20++ {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        v19 = v19
                        fmt.Println(fi20, v3)
                        v1 = append(v1, ("hello" + ("a" + "a")))
                        _ = fi21
                    }
                    var v22 int64 = -1
                    var v23 string = ("hello" + "")
                    _ = fi20
                    _ = v22
                    _ = v23
                }
                _ = v19
            }
        }
        v3 -= 255
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 uint32 = 65535
            for fi26 := 0; fi26 < 14; fi26++ {
                fmt.Println(fi24, fi24)
                if ((2 >= 0) || (1 > 302)) {
                    v27 := GenS2{F0: (^(v3))}
                    if false {
                        v3 -= 2
                        var v28 uint8 = 230
                        fmt.Println(v3)
                        delete(v2, (("" + ("foo bar" + "\n\t")) + ("hello" + "")))
                        var v29 uint8 = 163
                        var v30 uint = 622
                        var v31 float32 = -2.25
                        fmt.Println(fi26)
                        _ = v28
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        v3 = v3
                        switch (fi26 << ((0) & 63)) {
                        case 693:
                            if (float64((((3.14159 * (-2.25 - -2.25)) * (1.5 + 1.5)) + 3.14159)) >= (1.5 * 0.1)) {
                                fmt.Println(v3, v3)
                                var v32 uint32 = v25
                                _ = v32
                            } else {
                                if true {
                                    fmt.Println(fi26)
                                    v3 = 4294967295
                                    var v33 float64 = -2.25
                                    switch ("a" + ("" + "hello")) {
                                    case "foo bar":
                                        fmt.Println(fi24, fi24, v25)
                                    default:
                                        v33 -= (3.14159 + float64(1e10))
                                    }
                                    _ = v33
                                } else {
                                    if ((42 == 42) && true) {
                                        v34 := v2["hello"]
                                        _ = v34
                                    }
                                }
                                if true {
                                    if false {
                                        var v35 int64 = 1
                                        var v36 int = fi24
                                        var v37 bool = ((1.5 * -2.25) == 1e10)
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                    v38 := GenS1{F0: 100}
                                    _ = v38
                                }
                                v39 := GenS1{F0: 127}
                                _ = v39
                            }
                        case 128:
                            v1 = append(v1, ("\n\t" + (("hello" + "a") + "hello")))
                        case 536:
                            var v40 uint32 = v25
                            _ = v40
                        default:
                            v25 -= v25
                        }
                    }
                    _ = v27
                }
                _ = fi26
            }
            switch fi24 {
            case 128:
                v25 = v25
            case 211:
                for fi41 := 0; fi41 < 11; fi41++ {
                    fmt.Println(fi41)
                    v42 := GenS3{F0: ("hello" + ("a" + "\n\t")), F1: 319}
                    _ = fi41
                    _ = v42
                }
            case -9223372036854775808:
                switch ("" + "hello") {
                case "hello":
                    v1 = append(v1, "")
                case "":
                    fmt.Println(v25, v3)
                }
            default:
                v43 := GenS1{F0: 127}
                _ = v43
            }
            _ = fi24
            _ = v25
        }
    }
    v2[((("\n\t" + (("\n\t" + ((("foo bar" + "hello") + ("" + "")) + ("\n\t" + "a"))) + "hello")) + ("a" + (("foo bar" + "a") + "foo bar"))) + "\n\t")] = "foo bar"
    v3 = v3
    v44 := GenS1{F0: 0}
    v3 = v3
    v3 = v3
    if (2147483647 >= 4294967295) {
        v1 = append(v1, "hello")
        for fi45 := 0; fi45 < 9; fi45++ {
            v3 = 32
            var v46 int = 127
            _ = fi45
            _ = v46
        }
        var v47 int32 = -1
        _ = v47
    }
    v3 -= v3
    v3 -= v3
    for fi48 := 0; fi48 < 6; fi48++ {
        var v49 uint8 = 1
        v2["foo bar"] = "a"
        fmt.Println(v3, fi48, fi48)
        _ = fi48
        _ = v49
    }
    for fi50 := 0; fi50 < 19; fi50++ {
        var v51 uint16 = 90
        v51 -= v51
        _ = fi50
        _ = v51
    }
    switch "foo bar" {
    case "\n\t":
        v52 := v2[("a" + "")]
        _ = v52
    case "foo bar":
        v3 -= v3
    case "":
        for fi53 := 0; fi53 < 19; fi53++ {
            for fi54 := 0; fi54 < 16; fi54++ {
                _ = len(v1)
                v55 := GenS4{F0: 1}
                _ = fi54
                _ = v55
            }
            if ((v3 - 531) != v3) {
                if true {
                    var v56 int16 = 1000
                    _ = v56
                } else {
                    v3 = v3
                }
                v1 = append(v1, ("" + ""))
            }
            _ = fi53
        }
    }
    v3 = v3
    fmt.Println(v3)
    _ = len(v1)
    v3 -= (v3 & (v3 << (((v3 | (v3 + 42))) & 63)))
    v3 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
}
