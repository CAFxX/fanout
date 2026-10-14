package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int; F1 bool; F2 uint32 }
func genFn0() string {
    var v4 bool = true
    var v5 int = -9223372036854775808
    var v6 int16 = -32768
    if ((^(v5)) < 2147483647) {
        v7 := GenS1{F0: -1}
        _ = v7
    } else {
        switch "hello" {
        case "\n\t":
            v8 := GenS2{F0: -9223372036854775808, F1: v4}
            _ = v8
        default:
            v5 = v5
        }
        if v4 {
            v9 := []int{v5, 32768, -2147483648}
            v5 = (-(v5))
            _ = v9
        }
    }
    _ = v4
    _ = v5
    _ = v6
    return ""
}
func genFn1(p0 int16, p1 uint32) uint {
    switch ("" + (("a" + "\n\t") + "foo bar")) {
    case "foo bar":
        v10 := []int{-1, 127}
        _ = v10
    default:
        fmt.Println(p0, p1)
    }
    v11 := []int{194, 32768, 32768}
    p1 = 893
    _ = p0
    _ = p1
    _ = v11
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 26
    var v3 uint16 = 57
    var v12 uint64 = 18446744073709551615
    var v13 int8 = -1
    v2 += v2
    v3 += (v3 << ((genFn1(-32768, 4294967295)) & 63))
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(v3)
        _ = fi14
    }
    _ = len(v1)
    var v15 uint64 = uint64(v3)
    if true {
        fmt.Println(v12, v12)
        v12 -= (v12 | v12)
        fmt.Println(v3, v12, v15)
    }
    switch (("" + "a") + (("hello" + "hello") + "")) {
    case "\n\t":
        if (((160 == 42) && true) || false) {
            switch "a" {
            case "":
                var v16 uint8 = (v2 / ((v2) | 1))
                _ = v16
            default:
                for fi17 := 0; fi17 < 1; fi17++ {
                    v12 = (v15 - uint64(v3))
                    _ = len(v1)
                    v18 := GenS3{F0: 4294967295, F1: ((false || false) || (("foo bar" + ("foo bar" + "a")) < ("hello" + "\n\t")))}
                    _ = fi17
                    _ = v18
                }
            }
            v2 -= v2
        }
    case "foo bar":
        if false {
            if (!false) {
                if (714 >= 161) {
                    var v19 int64 = 1
                    v19 += v19
                    _ = v19
                } else {
                    v20 := []int{32767, -2147483648}
                    var v21 uint64 = v12
                    _ = v20
                    _ = v21
                }
                if (127 > 4294967295) {
                    switch -9223372036854775808 {
                    case 2147483647:
                        for fi22 := 0; fi22 < 9; fi22++ {
                            var v23 string = "foo bar"
                            _ = fi22
                            _ = v23
                        }
                    case -9223372036854775808:
                        for fi24 := 0; fi24 < 9; fi24++ {
                            var v25 int = (fi24 + (fi24 | 2147483647))
                            for fi26 := 0; fi26 < 13; fi26++ {
                                v27 := GenS1{F0: 9223372036854775807}
                                var v28 float64 = -2.25
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            for fi29 := 0; fi29 < 20; fi29++ {
                                var v30 int = (v25 | v25)
                                v31 := GenS1{F0: -283}
                                v1[""] = 710
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            _ = fi24
                            _ = v25
                        }
                    case 128:
                        var v32 int16 = 1
                        _ = v32
                    default:
                        _ = len(v1)
                    }
                }
                delete(v1, "foo bar")
            }
            var v33 float32 = 0.0
            _ = v33
        }
    default:
        switch 0 {
        case 128:
            var v34 uint8 = (v2 & (v2 | v2))
            _ = v34
        case -681:
            if (v12 <= (v12 ^ (^(v12)))) {
                var v35 uint = genFn1(1000, 502)
                _ = v35
            } else {
                if (42 >= 32767) {
                    if true {
                        v2 -= (v2 ^ v2)
                        if false {
                            v3 = (v3 % (((v3 ^ (v3 & 2))) | 1))
                        }
                        var v36 int8 = 0
                        _ = v36
                    }
                }
                v37 := GenS3{F0: 2, F1: true}
                _ = len(v1)
                _ = v37
            }
        case -693:
            v15 -= v12
        default:
            for fi38 := 0; fi38 < 12; fi38++ {
                if ((3.14159 != 0.1) && (-2.25 != float64(0.0))) {
                    v12 = v15
                }
                v39 := GenS4{F0: genFn1(32767, 2)}
                v40 := GenS5{F0: 2147483648, F1: 211}
                _ = fi38
                _ = v39
                _ = v40
            }
        }
    }
    var v41 string = ((("a" + ("hello" + "foo bar")) + ("foo bar" + "")) + "a")
    if true {
        var v42 uint = 4294967295
        _ = v42
    } else {
        v12 -= v15
    }
    var v43 string = "hello"
    for fi44 := 0; fi44 < 3; fi44++ {
        v41 = (v43 + "")
        _ = fi44
    }
    switch v41 {
    case "\n\t":
        if true {
            var v45 int8 = v13
            v13 = (v13 - v45)
            v46 := GenS6{F0: -554, F1: ((450 >= 2147483647) && (1 < 995)), F2: 2147483647}
            _ = v45
            _ = v46
        }
    case "":
        var v47 int64 = 9223372036854775807
        _ = v47
    }
    fmt.Println(v43, v15, v3)
    v48 := []int{-2147483648, 32768, -9223372036854775808}
    switch 127 {
    case 32767:
        v49 := GenS6{F0: 4294967295, F1: true, F2: 2147483648}
        _ = v49
    case 377:
        switch "a" {
        case "a":
            v50 := GenS4{F0: 2}
            _ = v50
        }
    case 4294967295:
        v12 -= v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v48:", len(v48), v48)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v41
    _ = v43
    _ = v48
}
