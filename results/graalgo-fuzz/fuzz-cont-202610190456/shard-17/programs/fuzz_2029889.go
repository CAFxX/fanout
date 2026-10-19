package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0() uint16 {
    var v7 uint = 42
    var v8 uint8 = 255
    _ = v7
    _ = v8
    return 2
}
func genFn1() uint32 {
    for fi9 := 0; fi9 < 16; fi9++ {
        if false {
            fmt.Println(fi9)
            switch -1 {
            case -2147483648:
                var v10 uint8 = 2
                _ = v10
            case -745:
                var v11 string = ""
                _ = v11
            case 1:
                if false {
                    var v12 int16 = 32767
                    var v13 float32 = float32(float64(1.5))
                    _ = v12
                    _ = v13
                }
            default:
                var v14 int64 = 0
                _ = v14
            }
        } else {
            var v15 bool = ((42 != 0) || false)
            _ = v15
        }
        _ = fi9
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        fmt.Println(fi16)
        _ = fi16
    }
    if (false && false) {
        v17 := []int{}
        var v18 bool = true
        v17 = append(v17, 42)
        _ = v17
        _ = v18
    } else {
        if false {
            var v19 float32 = 0.0
            _ = v19
        } else {
            v20 := GenS1{F0: 0, F1: 476}
            switch ("a" + (("foo bar" + "hello") + "")) {
            case "\n\t":
                for fi21 := 0; fi21 < 10; fi21++ {
                    var v22 uint16 = 42
                    if (false && (1 < (v22 >> ((255) & 63)))) {
                        for fi23 := 0; fi23 < 18; fi23++ {
                            v24 := []int{(fi23 | 4294967295)}
                            if (false || (false && (-2147483648 >= 1))) {
                                var v25 int = -1
                                fmt.Println(v25)
                                v25 = v25
                                _ = v25
                            }
                            if true {
                                fmt.Println(fi21, v22, v22)
                                switch "\n\t" {
                                case "foo bar":
                                    if true {
                                        if false {
                                            switch (fi21 << ((4294967295) & 63)) {
                                            case 493:
                                                if false {
                                                    v22 = v22
                                                }
                                            default:
                                                v26 := GenS2{F0: 32768, F1: 4294967295}
                                                _ = v26
                                            }
                                            var v27 float32 = float32(((-2.25 - 3.14159) - 1e10))
                                            var v28 uint = 42
                                            _ = v27
                                            _ = v28
                                        } else {
                                            fmt.Println(fi21, v22)
                                        }
                                    } else {
                                        var v29 uint16 = (v22 / (((^(v22))) | 1))
                                        if len(v24) > 0 {
                                            v24[0] = (fi23 + (fi23 % ((4294967295) | 1)))
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            v30 := GenS1{F0: 127, F1: 300}
                                            _ = v30
                                        case "":
                                            v22 = 32767
                                            fmt.Println(fi23, v22)
                                            if len(v24) > 0 {
                                                v24[0] = fi23
                                            }
                                            if len(v24) >= 2 {
                                                v24 = v24[1:len(v24)]
                                            }
                                        case "hello":
                                            if len(v24) >= 2 {
                                                v24 = v24[1:len(v24)]
                                            }
                                        default:
                                            switch 9223372036854775807 {
                                            case 127:
                                                var v31 int32 = 0
                                                _ = v31
                                            default:
                                                switch "" {
                                                case "hello":
                                                    if len(v24) > 0 {
                                                        v24[0] = (fi23 + fi23)
                                                    }
                                                case "\n\t":
                                                    _ = len(v24)
                                                    switch (^(fi23)) {
                                                    case 0:
                                                        v29 = genFn0()
                                                    default:
                                                        switch 664 {
                                                        case -268:
                                                            var v32 uint64 = 893
                                                            _ = v32
                                                        case 89:
                                                            v33 := GenS1{F0: 100, F1: -250}
                                                            _ = v33
                                                        case -581:
                                                            var v34 int = (fi21 % ((32767) | 1))
                                                            _ = v34
                                                        default:
                                                            if false {
                                                                var v35 bool = true
                                                                _ = v35
                                                            }
                                                            fmt.Println(fi21)
                                                        }
                                                    }
                                                    var v36 int16 = 1000
                                                    v22 += (v22 ^ v29)
                                                    _ = v36
                                                case "":
                                                    v29 = (v29 * v29)
                                                }
                                            }
                                        }
                                        _ = v29
                                    }
                                default:
                                    var v37 float64 = 0.1
                                    _ = v37
                                }
                            } else {
                                var v38 float32 = ((0.0 + -2.25) - 3.14159)
                                _ = v38
                            }
                            _ = fi23
                            _ = v24
                        }
                        v22 = 304
                    } else {
                        var v39 string = "foo bar"
                        fmt.Println(fi21)
                        var v40 uint64 = 0
                        _ = v39
                        _ = v40
                    }
                    _ = fi21
                    _ = v22
                }
            case "hello":
                var v41 int16 = -1
                _ = v41
            }
            _ = v20
        }
        v42 := []int{32767, 127, -2147483648, -9223372036854775808}
        var v43 uint = 1
        _ = v42
        _ = v43
    }
    return 65535
}

func main() {
    var v1 int8 = -1
    var v2 int16 = 1000
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    var v5 int8 = (v1 * (v1 / ((v1) | 1)))
    v6 := []string{}
    v5 = -128
    if len(v6) > 0 {
        v6[0] = "foo bar"
    }
    var v44 uint = 42
    v2 = v2
    v44 = (v44 - 255)
    _ = len(v6)
    var v45 int16 = v2
    switch ("hello" + "foo bar") {
    case "hello":
        var v46 uint = v44
        _ = v46
    case "foo bar":
        _ = len(v6)
    }
    v6 = append(v6, "a")
    v45 -= v45
    if (false || true) {
        v47 := GenS1{F0: 42, F1: 2147483647}
        for fi48 := 0; fi48 < 1; fi48++ {
            v1 -= v5
            fmt.Println(v44)
            _ = fi48
        }
        _ = v47
    } else {
        var v49 uint64 = 42
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
