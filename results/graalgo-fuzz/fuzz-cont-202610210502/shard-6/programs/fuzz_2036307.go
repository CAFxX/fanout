package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint16) uint {
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{(-(fi7)), (fi7 + int(p0)), 1}
        if (!(!false)) {
            v8 = append(v8, fi7)
        }
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: false, F2: -1}
    v10 := GenS1{F0: 1, F1: (false && (p0 == p0)), F2: -2147483648}
    fmt.Println(p0)
    if (65535 != 123) {
        v11 := []int{706}
        _ = v11
    }
    _ = p0
    _ = v10
    _ = v9
    return 327
}
func genFn1(p0 int64) bool {
    if true {
        fmt.Println(p0)
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 bool = (!((fi12 * fi12) != fi12))
        v13 = v13
        _ = fi12
        _ = v13
    }
    _ = p0
    return true
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 254
    var v3 int64 = -9223372036854775808
    var v4 bool = (18446744073709551615 > 2)
    var v5 uint8 = v2
    v6 := make(map[string]bool)
    if v4 {
        v4 = false
        var v14 int32 = -783
        _ = v14
    }
    switch "foo bar" {
    case "":
        var v15 string = ("hello" + "foo bar")
        _ = v15
    case "hello":
        var v16 uint64 = 298
        _ = v16
    default:
        v6["hello"] = v4
    }
    if v4 {
        if v4 {
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 bool = (!v4)
                v19 := []int{(fi17 - (fi17 | fi17)), (fi17 ^ -2147483648), 0}
                _ = fi17
                _ = v18
                _ = v19
            }
            for fi20 := 0; fi20 < 19; fi20++ {
                if (!v4) {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        v4 = ((v2 / (((^(v2))) | 1)) > 1)
                        _ = fi21
                    }
                }
                if v4 {
                    var v22 int16 = 42
                    _ = v22
                }
                delete(v6, "")
                _ = fi20
            }
        } else {
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 uint = 0
                fmt.Println(v24, v4, v4)
                _ = fi23
                _ = v24
            }
        }
        v25 := []int{32767}
        _ = v25
    }
    v26 := GenS1{F0: 42, F1: (1 >= 32767), F2: -1}
    var v27 bool = v4
    v28 := []int{-2147483648, -926, 42}
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    switch ("" + "\n\t") {
    case "":
        v5 = v5
    case "a":
        var v29 bool = (-1 >= -1)
        _ = v29
    case "foo bar":
        var v30 uint16 = 32767
        _ = v30
    }
    switch 1 {
    case -356:
        var v31 uint32 = 2
        _ = v31
    case -2147483648:
        var v32 uint16 = 508
        _ = v32
    case 127:
        switch "" {
        case "foo bar":
            if (false && (-9223372036854775808 <= -316)) {
                switch 32767 {
                case 32767:
                    _ = len(v28)
                case -2147483648:
                    fmt.Println(v4)
                default:
                    if len(v28) > 0 {
                        v28[0] = 1
                    }
                }
                switch (((("foo bar" + "foo bar") + "") + "\n\t") + "hello") {
                case "":
                    if (((v4 || (v4 && false)) || v27) && (-142 > -236)) {
                        var v33 float32 = 1.5
                        v33 = v33
                        _ = v33
                    } else {
                        switch 1 {
                        case 42:
                            for fi34 := 0; fi34 < 19; fi34++ {
                                var v35 int8 = 42
                                var v36 int8 = 41
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                        case 2147483647:
                            var v37 uint16 = 2
                            _ = v37
                        default:
                            if (("" >= "\n\t") || (v4 || v27)) {
                                for fi38 := 0; fi38 < 9; fi38++ {
                                    var v39 string = (("foo bar" + (("foo bar" + "a") + "a")) + (("foo bar" + "\n\t") + ("\n\t" + "hello")))
                                    if len(v28) >= 2 {
                                        v28 = v28[1:len(v28)]
                                    }
                                    _ = fi38
                                    _ = v39
                                }
                            } else {
                                if len(v28) > 0 {
                                    v28[0] = 32768
                                }
                                for fi40 := 0; fi40 < 16; fi40++ {
                                    for fi41 := 0; fi41 < 8; fi41++ {
                                        v27 = ("\n\t" >= "\n\t")
                                        v4 = ((v5 << ((2) & 63)) < v5)
                                        v5 += v2
                                        var v42 int16 = 42
                                        _ = fi41
                                        _ = v42
                                    }
                                    for fi43 := 0; fi43 < 11; fi43++ {
                                        fmt.Println(v4, v4, fi40)
                                        if v27 {
                                            v44 := GenS1{F0: -1, F1: v27, F2: -2147483648}
                                            _ = v44
                                        }
                                        var v45 uint32 = 0
                                        _ = fi43
                                        _ = v45
                                    }
                                    _ = fi40
                                }
                            }
                        }
                        for fi46 := 0; fi46 < 3; fi46++ {
                            v5 = v2
                            _ = fi46
                        }
                    }
                case "foo bar":
                    v4 = (0 != 1)
                case "a":
                    v3 += v3
                default:
                    switch 32767 {
                    case 42:
                        var v47 uint32 = 2
                        _ = v47
                    case -358:
                        var v48 uint = 569
                        _ = v48
                    case -9223372036854775808:
                        v49 := GenS1{F0: -34, F1: ((65535 == 65535) && v27), F2: 1}
                        _ = v49
                    }
                }
            }
        case "a":
            var v50 int64 = -9223372036854775808
            _ = v50
        }
    default:
        var v51 uint = 958
        _ = v51
    }
    if (622 > 224) {
        v52 := GenS2{F0: 2}
        _ = v52
    } else {
        v27 = v27
    }
    var v53 int = 0
    switch "a" {
    case "a":
        fmt.Println(v27, v27, v27)
    case "foo bar":
        v5 -= v5
    default:
        fmt.Println(v3, v2, v2)
    }
    switch (("hello" + ("" + "\n\t")) + "") {
    case "a":
        fmt.Println(v27)
    case "":
        var v54 uint8 = v2
        _ = v54
    default:
        v3 = -1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}
