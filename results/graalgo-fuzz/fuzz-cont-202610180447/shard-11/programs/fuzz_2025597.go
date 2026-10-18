package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 string; F2 int }
type GenS4 struct { F0 int16; F1 int16; F2 int8 }
func genFn0(p0 int32, p1 uint16) string {
    v6 := GenS1{F0: ("" + ("" + "")), F1: false}
    if (!(false || true)) {
        v7 := []int{4294967295, -579, 383}
        _ = v7
    } else {
        switch "foo bar" {
        case "a":
            var v8 uint64 = 57
            _ = v8
        default:
            switch 32768 {
            case -9223372036854775808:
                switch (("foo bar" + "a") + "hello") {
                case "":
                    v9 := GenS1{F0: "hello", F1: false}
                    _ = v9
                case "\n\t":
                    v10 := []int{-9223372036854775808, -1, -2147483648}
                    _ = v10
                }
            case 127:
                v11 := []int{9223372036854775807, 103, 523, 2147483647}
                _ = v11
            case -2147483648:
                v12 := []int{0}
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            p1 = p1
            p0 = (-(p0))
            if (false || (p0 == (p0 + (p0 / ((0) | 1))))) {
                for fi14 := 0; fi14 < 1; fi14++ {
                    v15 := []int{fi13, (fi14 ^ fi13)}
                    switch fi13 {
                    case -921:
                        fmt.Println(p1, fi14)
                    case 32767:
                        switch "foo bar" {
                        case "":
                            if (!false) {
                                fmt.Println(p0, fi14)
                                v15 = append(v15, (^(fi14)))
                                fmt.Println(fi13, fi14, p0)
                            }
                        default:
                            v16 := GenS1{F0: "a", F1: true}
                            _ = v16
                        }
                    case 4294967295:
                        p1 -= p1
                        var v17 string = ("hello" + "hello")
                        var v18 int16 = 353
                        fmt.Println(fi14, fi13, fi13)
                        _ = len(v15)
                        if len(v15) > 0 {
                            v15[0] = (fi13 ^ 32768)
                        }
                        if len(v15) > 0 {
                            v15[0] = fi14
                        }
                        _ = v17
                        _ = v18
                    }
                    _ = fi14
                    _ = v15
                }
            } else {
                var v19 float32 = 0.0
                if false {
                    fmt.Println(p0, p0)
                    p1 = p1
                }
                _ = v19
            }
            _ = fi13
        }
    }
    switch "foo bar" {
    case "a":
        var v20 float64 = -2.25
        _ = v20
    case "foo bar":
        p0 = 2147483647
    }
    _ = p0
    _ = p1
    _ = v6
    return (("foo bar" + "foo bar") + "a")
}
func genFn1() uint16 {
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 uint = 2
        var v23 int = 2147483647
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 int = 4294967295
    var v25 int8 = 1
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 string = "a"
        switch (v27 + v27) {
        case "a":
            switch (fi26 >> ((42) & 63)) {
            case 2147483647:
                fmt.Println(fi26, v24, v24)
            case 322:
                v25 += v25
            }
        case "":
            v28 := GenS1{F0: "foo bar", F1: (true && true)}
            _ = v28
        default:
            var v29 uint8 = 1
            _ = v29
        }
        switch v27 {
        case "\n\t":
            if (1 < 2147483648) {
                var v30 string = (("a" + v27) + (v27 + (v27 + "a")))
                for fi31 := 0; fi31 < 10; fi31++ {
                    var v32 uint16 = 2
                    var v33 bool = true
                    v34 := GenS1{F0: v30, F1: (-1 > (v25 - -1))}
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
                _ = v30
            } else {
                v27 = v27
                v35 := []int{(-(fi26)), v24}
                _ = v35
            }
        case "hello":
            v36 := GenS2{F0: 871}
            _ = v36
        }
        _ = fi26
        _ = v27
    }
    v37 := GenS3{F0: -373, F1: (("hello" + (("" + "") + "foo bar")) + ("\n\t" + ("foo bar" + "a"))), F2: (v24 / ((v24) | 1))}
    _ = v24
    _ = v25
    _ = v37
    return 255
}

func main() {
    var v1 string = (("\n\t" + (("hello" + "a") + "a")) + ("" + "hello"))
    v2 := []string{}
    var v3 uint16 = 644
    var v4 uint8 = 127
    var v5 uint32 = 2147483648
    fmt.Println(v5, v5, v1)
    v38 := GenS3{F0: -1, F1: v1, F2: 2147483647}
    var v39 uint = 0
    if len(v2) > 0 {
        v2[0] = ("a" + "")
    }
    var v40 int8 = 1
    for fi41 := 0; fi41 < 16; fi41++ {
        fmt.Println(v40, v5)
        for fi42 := 0; fi42 < 19; fi42++ {
            v43 := GenS1{F0: ((((v1 + v1) + ("" + v1)) + (("a" + v1) + v1)) + ((v1 + "") + "foo bar")), F1: (1 >= genFn1())}
            var v44 float64 = 1.5
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi41
    }
    v3 = (v3 & (v3 % ((v3) | 1)))
    switch (((v1 + "") + v1) + (v1 + v1)) {
    case "\n\t":
        fmt.Println(v40, v1, v39)
    case "a":
        for fi45 := 0; fi45 < 1; fi45++ {
            for fi46 := 0; fi46 < 2; fi46++ {
                var v47 int = (fi46 * (^(fi46)))
                if false {
                    v48 := GenS1{F0: (v1 + "a"), F1: true}
                    v49 := GenS4{F0: 42, F1: -390, F2: 127}
                    _ = v48
                    _ = v49
                } else {
                    var v50 uint16 = (v3 + v3)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v40 -= (v40 + -128)
                    _ = v50
                }
                if (v1 != v1) {
                    fmt.Println(v3, fi46, fi45)
                }
                var v51 float64 = 0.1
                _ = fi46
                _ = v47
                _ = v51
            }
            _ = fi45
        }
    default:
        for fi52 := 0; fi52 < 6; fi52++ {
            if false {
                if true {
                    var v53 uint = 0
                    if (uint64(v4) != 2) {
                        v39 = v53
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        for fi54 := 0; fi54 < 17; fi54++ {
                            v1 = ""
                            _ = fi54
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v53
                }
                for fi55 := 0; fi55 < 13; fi55++ {
                    v5 = v5
                    if true {
                        if ((!false) && false) {
                            v3 += 2
                        } else {
                            v4 -= v4
                            v56 := GenS4{F0: 1000, F1: int16(v3), F2: v40}
                            var v57 int16 = -1
                            _ = v56
                            _ = v57
                        }
                        if (!(("foo bar" + (v1 + v1)) < ("foo bar" + (v1 + v1)))) {
                            if len(v2) > 0 {
                                v2[0] = v1
                            }
                            fmt.Println(fi55)
                        }
                    }
                    _ = fi55
                }
                var v58 int16 = 0
                _ = v58
            }
            if false {
                var v59 uint = (v39 | (v39 + (v39 >> (((v39 + (v39 / ((v39) | 1)))) & 63))))
                for fi60 := 0; fi60 < 14; fi60++ {
                    switch v1 {
                    case "foo bar":
                        v4 = v4
                    case "":
                        var v61 int = fi52
                        _ = v61
                    default:
                        v4 = v4
                        if true {
                            if true {
                                v2 = append(v2, "\n\t")
                            }
                            v62 := GenS2{F0: (v3 << (((^(v39))) & 63))}
                            _ = v62
                        }
                        v4 = uint8(v59)
                    }
                    _ = fi60
                }
                _ = v59
            }
            _ = fi52
        }
    }
    if false {
        v3 -= (^(v3))
        var v63 uint8 = 84
        var v64 uint8 = (v4 - v63)
        _ = v63
        _ = v64
    }
    v39 += 42
    v3 -= genFn1()
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
