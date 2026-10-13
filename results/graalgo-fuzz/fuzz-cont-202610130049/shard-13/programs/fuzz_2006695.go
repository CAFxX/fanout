package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int8; F1 int8; F2 string }
type GenS4 struct { F0 uint8; F1 uint64; F2 int }
type GenS5 struct { F0 uint32; F1 uint; F2 int8 }
type GenS6 struct { F0 int8; F1 string }
func genFn0(p0 uint32) uint16 {
    p0 += 318
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 42, F1: 42}
        v9 := []int{fi7, fi7, -2147483648, (fi7 ^ fi7)}
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 = p0
    switch "foo bar" {
    case "":
        v10 := GenS1{F0: 4294967295, F1: 0}
        _ = v10
    case "\n\t":
        v11 := []int{1}
        _ = v11
    default:
        switch 9223372036854775807 {
        case 42:
            v12 := GenS2{F0: 423, F1: 32767}
            _ = v12
        case 1:
            var v13 int16 = 1000
            _ = v13
        case 9223372036854775807:
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := []int{(fi14 & fi14), (^(fi14))}
                var v16 float64 = -2.25
                switch fi14 {
                case 42:
                    for fi17 := 0; fi17 < 15; fi17++ {
                        p0 += p0
                        _ = fi17
                    }
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    }
    _ = p0
    return 1
}
func genFn1(p0 int64, p1 int, p2 int16) bool {
    v18 := []int{p1, (^(p1)), p1, (p1 * p1)}
    var v19 uint8 = 1
    if (false && false) {
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    } else {
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        for fi20 := 0; fi20 < 7; fi20++ {
            for fi21 := 0; fi21 < 15; fi21++ {
                fmt.Println(p0)
                v18 = append(v18, (-(fi21)))
                if (!(!true)) {
                    p1 = fi21
                    switch 0 {
                    case -627:
                        if len(v18) > 0 {
                            v18[0] = p1
                        }
                    case 128:
                        v19 = (^(v19))
                    default:
                        fmt.Println(p1, v19)
                    }
                    v22 := GenS1{F0: 255, F1: 785}
                    if (!true) {
                        if (p2 < p2) {
                            fmt.Println(fi20, fi20, v19)
                            var v23 int8 = 24
                            _ = v23
                        }
                    }
                    v19 = (v19 * (v19 & (v19 % ((127) | 1))))
                    _ = v22
                }
                var v24 uint8 = (v19 + (^(v19)))
                _ = fi21
                _ = v24
            }
            switch 2147483647 {
            case 9223372036854775807:
                if false {
                    switch ("" + ("" + "foo bar")) {
                    case "foo bar":
                        var v25 float32 = 1.5
                        _ = v25
                    case "":
                        switch fi20 {
                        case 735:
                            fmt.Println(p0)
                        default:
                            p0 = p0
                        }
                    case "a":
                        p2 = (p2 | (^(p2)))
                    default:
                        if true {
                            if ((4294967295 > 65535) || true) {
                                for fi26 := 0; fi26 < 9; fi26++ {
                                    v27 := GenS2{F0: 1, F1: (-(p2))}
                                    var v28 string = (("\n\t" + ("\n\t" + "a")) + "a")
                                    v29 := GenS3{F0: -22, F1: 100, F2: "\n\t"}
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                                for fi30 := 0; fi30 < 8; fi30++ {
                                    if len(v18) > 0 {
                                        v18[0] = p1
                                    }
                                    _ = fi30
                                }
                            }
                            p1 = -1
                            var v31 bool = true
                            _ = v31
                        }
                    }
                    p0 = p0
                }
            case -711:
                p2 += p2
            case 289:
                if (!false) {
                    v19 = (^(v19))
                } else {
                    var v32 int32 = 1
                    _ = v32
                }
            default:
                var v33 float32 = (0.1 / 1.5)
                _ = v33
            }
            _ = fi20
        }
        var v34 bool = false
        _ = v34
    }
    var v35 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v35
    return (!true)
}

func main() {
    var v1 float32 = (-2.25 * float32(((1e10 + 1.5) * (((1.5 / float64(3.14159)) / -2.25) / 0.1))))
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    var v4 int8 = -128
    v5 := []string{"foo bar", "foo bar", ("\n\t" + (("\n\t" + "\n\t") + ("hello" + (("hello" + "\n\t") + "\n\t")))), ""}
    var v6 uint32 = 2
    switch "\n\t" {
    case "":
        var v36 uint64 = 18446744073709551615
        _ = v36
    case "hello":
        v6 -= 54
    case "\n\t":
        if true {
            v37 := GenS3{F0: v4, F1: 0, F2: (("\n\t" + "foo bar") + "hello")}
            switch "a" {
            case "a":
                var v38 bool = false
                _ = v38
            case "\n\t":
                for fi39 := 0; fi39 < 9; fi39++ {
                    fmt.Println(v6)
                    var v40 int8 = (v4 & int8(v6))
                    _ = fi39
                    _ = v40
                }
            default:
                for fi41 := 0; fi41 < 14; fi41++ {
                    fmt.Println(fi41)
                    v42 := v2[(fi41 & fi41)]
                    for fi43 := 0; fi43 < 19; fi43++ {
                        fmt.Println(v1, v6, fi43)
                        v1 = ((v1 * 3.14159) + (v1 / v1))
                        _ = fi43
                    }
                    _ = fi41
                    _ = v42
                }
            }
            var v44 int = 44
            _ = v37
            _ = v44
        } else {
            if (42 >= 42) {
                fmt.Println(v4)
                switch 636 {
                case 32767:
                    v6 += v6
                case 91:
                    var v45 int = 32767
                    _ = v45
                }
            } else {
                fmt.Println(v6, v6, v4)
                for fi46 := 0; fi46 < 8; fi46++ {
                    var v47 int8 = (v4 ^ 42)
                    v48 := GenS4{F0: 255, F1: 2, F2: 2147483647}
                    fmt.Println(fi46, fi46, v47)
                    _ = fi46
                    _ = v47
                    _ = v48
                }
            }
        }
    }
    var v49 int16 = 32767
    var v50 uint8 = 2
    v51 := GenS2{F0: 0, F1: v49}
    var v52 bool = false
    switch 9223372036854775807 {
    case 32767:
        switch 42 {
        case 32767:
            var v53 int16 = (v49 << ((18446744073709551615) & 63))
            _ = v53
        case 0:
            v6 = v6
        case 2147483647:
            v54 := GenS5{F0: (v6 + (v6 | (v6 - (^(v6))))), F1: 1, F2: (v4 + 69)}
            _ = v54
        }
    }
    fmt.Println(v50, v4, v49)
    delete(v2, 1)
    if v52 {
        v55 := GenS6{F0: 127, F1: ""}
        switch 128 {
        case 0:
            v52 = (32767 < 65535)
        case 345:
            var v56 int64 = 0
            _ = v56
        default:
            v49 = v49
        }
        _ = v55
    }
    switch "a" {
    case "foo bar":
        for fi57 := 0; fi57 < 15; fi57++ {
            var v58 int8 = (v4 - 100)
            _ = len(v5)
            var v59 string = ("" + "")
            _ = fi57
            _ = v58
            _ = v59
        }
    }
    if (v52 && v52) {
        v60 := GenS6{F0: (^(v4)), F1: (("" + ("a" + ("\n\t" + ("" + "hello")))) + "")}
        _ = v60
    }
    for fi61 := 0; fi61 < 14; fi61++ {
        var v62 int = 936
        var v63 string = ""
        _ = fi61
        _ = v62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
