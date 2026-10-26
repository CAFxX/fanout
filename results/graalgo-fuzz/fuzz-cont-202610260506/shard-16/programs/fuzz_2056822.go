package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint8; F1 string; F2 int64 }
type GenS6 struct { F0 int }
func genFn0(p0 int8, p1 uint32, p2 int64) uint {
    var v4 bool = (false || ("" == ("" + "foo bar")))
    v5 := []int{-1, -2147483648, 1, 510}
    var v6 int64 = p2
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 577
}
func genFn1(p0 string) uint8 {
    for fi7 := 0; fi7 < 12; fi7++ {
        p0 = p0
        v8 := GenS1{F0: 2147483647}
        p0 = ("" + ("foo bar" + p0))
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    switch "hello" {
    case "\n\t":
        fmt.Println(p0)
    }
    var v9 uint8 = 2
    switch "foo bar" {
    case "foo bar":
        for fi10 := 0; fi10 < 8; fi10++ {
            if false {
                v9 += (v9 * 227)
            }
            if true {
                v9 = v9
            }
            _ = fi10
        }
    case "\n\t":
        v9 = v9
    default:
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 bool = (2147483647 <= 885)
            _ = fi11
            _ = v12
        }
    }
    _ = p0
    _ = v9
    return 3
}

func main() {
    var v1 bool = false
    var v2 uint16 = 0
    var v3 uint32 = 0
    switch -27 {
    case -2147483648:
        v3 = v3
    case 128:
        if (-675 <= 496) {
            if v1 {
                for fi13 := 0; fi13 < 11; fi13++ {
                    if true {
                        for fi14 := 0; fi14 < 7; fi14++ {
                            if v1 {
                                var v15 bool = false
                                v16 := GenS1{F0: 0}
                                v17 := []int{(fi14 * fi14), (fi14 - 127)}
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            switch fi13 {
                            case 42:
                                var v18 uint32 = v3
                                _ = v18
                            default:
                                if true {
                                    var v19 uint32 = (^(v3))
                                    _ = v19
                                } else {
                                    v2 = v2
                                    switch "\n\t" {
                                    case "\n\t":
                                        if ((true && v1) && ((v2 % (((v2 + 42)) | 1)) == (^(v2)))) {
                                            var v20 float32 = 1e10
                                            v21 := GenS1{F0: 1}
                                            v22 := []int{}
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        }
                                    case "":
                                        var v23 int = (fi14 + (fi13 * -702))
                                        _ = v23
                                    }
                                    var v24 float64 = 0.0
                                    _ = v24
                                }
                            }
                            fmt.Println(v2)
                            _ = fi14
                        }
                        var v25 int8 = 100
                        var v26 bool = true
                        _ = v25
                        _ = v26
                    } else {
                        switch "hello" {
                        case "a":
                            for fi27 := 0; fi27 < 20; fi27++ {
                                var v28 string = "\n\t"
                                v29 := []int{}
                                switch (v28 + "foo bar") {
                                case "":
                                    v1 = true
                                default:
                                    var v30 bool = (v1 && ((false && (!false)) && (2147483647 > v3)))
                                    _ = v30
                                }
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                        case "foo bar":
                            v31 := GenS2{F0: (v2 ^ v2)}
                            _ = v31
                        }
                        fmt.Println(v1)
                    }
                    if true {
                        if v1 {
                            var v32 string = "\n\t"
                            _ = v32
                        }
                        if (fi13 >= 4294967295) {
                            switch (("foo bar" + ("a" + "a")) + ("foo bar" + ("" + (("a" + "a") + "a")))) {
                            case "":
                                v33 := []int{fi13, 128, fi13, (fi13 << ((genFn0(-1, (v3 + 42), 42)) & 63))}
                                _ = v33
                            default:
                                var v34 uint8 = 85
                                _ = v34
                            }
                            v35 := GenS3{F0: 2, F1: 70}
                            _ = v35
                        } else {
                            v36 := []int{4294967295, fi13}
                            if false {
                                var v37 int = (fi13 - 128)
                                if (((v3 != 65535) && false) || v1) {
                                    var v38 int64 = 288
                                    _ = v38
                                }
                                for fi39 := 0; fi39 < 7; fi39++ {
                                    var v40 uint32 = uint32(v2)
                                    _ = fi39
                                    _ = v40
                                }
                                _ = v37
                            }
                            _ = v36
                        }
                        v41 := []int{}
                        _ = v41
                    } else {
                        v1 = v1
                        v42 := GenS4{F0: genFn1(("\n\t" + "hello")), F1: 0, F2: 0}
                        _ = v42
                    }
                    fmt.Println(v1)
                    _ = fi13
                }
            }
        }
    default:
        if (-1 != 2147483647) {
            fmt.Println(v2)
            var v43 int8 = 1
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        for fi45 := 0; fi45 < 15; fi45++ {
            var v46 bool = (1e10 > 1.5)
            var v47 float32 = float32(-2.25)
            v48 := GenS4{F0: 231, F1: -1, F2: 903}
            _ = fi45
            _ = v46
            _ = v47
            _ = v48
        }
        _ = fi44
    }
    switch "a" {
    case "a":
        v49 := GenS5{F0: 1, F1: ("\n\t" + ("foo bar" + "hello")), F2: 9223372036854775807}
        _ = v49
    case "hello":
        v50 := []int{692, 128}
        _ = v50
    case "foo bar":
        switch 9223372036854775807 {
        case 42:
            var v51 float32 = 1.5
            _ = v51
        case -371:
            if v1 {
                v52 := GenS3{F0: 0, F1: -55}
                v3 = 51
                _ = v52
            } else {
                for fi53 := 0; fi53 < 14; fi53++ {
                    v54 := []int{1, (fi53 * fi53)}
                    _ = fi53
                    _ = v54
                }
            }
        case 128:
            fmt.Println(v1)
        default:
            if (!(!(!v1))) {
                v55 := GenS6{F0: 2147483647}
                fmt.Println(v3, v2)
                fmt.Println(v2, v2, v2)
                _ = v55
            }
        }
    default:
        v2 += (v2 ^ v2)
    }
    var v56 int64 = 9223372036854775807
    v57 := []int{32768, 4294967295, 32767}
    v3 = 65535
    if len(v57) > 0 {
        v57[0] = 0
    }
    if len(v57) > 0 {
        v57[0] = 4294967295
    }
    _ = len(v57)
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v57:", len(v57), v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v56
    _ = v57
}
