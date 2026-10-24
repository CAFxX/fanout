package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 uint32; F2 int }
type GenS5 struct { F0 uint8; F1 uint; F2 int8 }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 string }
func genFn0(p0 int64, p1 int16) string {
    v4 := []int{-1, -1, -109, -624}
    if false {
        switch "hello" {
        case "hello":
            v5 := GenS1{F0: 997, F1: 707}
            _ = v5
        case "":
            fmt.Println(p0, p0)
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        var v6 int = 1
        p0 = p0
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        p1 = (-(p1))
        var v8 int32 = 1
        _ = fi7
        _ = v8
    }
    p1 = (p1 % ((p1) | 1))
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v4
    return (("foo bar" + ("hello" + "")) + "a")
}
func genFn1(p0 int32) uint {
    v9 := []int{347, 769, 9223372036854775807}
    var v10 uint = 1
    var v11 uint32 = 65535
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 158
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 1
    v3 := []int{}
    switch ("" + "\n\t") {
    case "":
        if true {
            switch ("hello" + "\n\t") {
            case "foo bar":
                if false {
                    for fi12 := 0; fi12 < 11; fi12++ {
                        v13 := v1["\n\t"]
                        _ = fi12
                        _ = v13
                    }
                    v14 := GenS1{F0: 946, F1: v2}
                    v15 := GenS2{F0: 0}
                    _ = v14
                    _ = v15
                }
            case "\n\t":
                var v16 int32 = v2
                _ = v16
            default:
                switch 32767 {
                case 42:
                    for fi17 := 0; fi17 < 3; fi17++ {
                        var v18 int8 = -128
                        v18 = v18
                        v18 = (v18 % ((42) | 1))
                        _ = fi17
                        _ = v18
                    }
                case 32767:
                    var v19 uint32 = 1
                    _ = v19
                case 516:
                    v2 = -308
                }
            }
            var v20 int16 = -1
            _ = v20
        }
    case "foo bar":
        for fi21 := 0; fi21 < 9; fi21++ {
            _ = len(v3)
            _ = fi21
        }
    }
    if (70 < 1) {
        for fi22 := 0; fi22 < 19; fi22++ {
            for fi23 := 0; fi23 < 6; fi23++ {
                v2 = (v2 | (^(v2)))
                v2 += 0
                var v24 bool = true
                var v25 int32 = v2
                if (!v24) {
                    v25 += v25
                    v26 := GenS2{F0: genFn1((v25 << ((genFn1(v25)) & 63)))}
                    _ = v26
                } else {
                    fmt.Println(fi22, v25)
                    var v27 uint64 = 778
                    if ((((("hello" + "hello") + "\n\t") + (("\n\t" + "\n\t") + ("\n\t" + "hello"))) < "hello") || (v24 && ((((v24 && (false || v24)) || v24) && (float32(0.0) != -2.25)) || (0 > 415)))) {
                        var v28 int32 = 1
                        _ = v28
                    } else {
                        switch (("" + "\n\t") + ("" + (("a" + "a") + "hello"))) {
                        case "\n\t":
                            switch fi22 {
                            case -342:
                                switch (("" + "") + "foo bar") {
                                case "hello":
                                    v27 = (^(v27))
                                case "foo bar":
                                    v29 := GenS1{F0: 2147483648, F1: v25}
                                    _ = v29
                                case "\n\t":
                                    v30 := GenS3{F0: v25}
                                    _ = v30
                                default:
                                    fmt.Println(v24)
                                }
                            case 127:
                                var v31 int8 = 0
                                _ = v31
                            case -1:
                                v2 -= (v25 / ((v25) | 1))
                            }
                        case "foo bar":
                            switch "" {
                            case "":
                                v2 = -1
                            }
                        }
                        if v24 {
                            if true {
                                switch (fi22 / (((fi22 >> ((2) & 63))) | 1)) {
                                case 937:
                                    v2 = v2
                                case 2147483647:
                                    var v32 int16 = -32768
                                    _ = v32
                                default:
                                    if (0 <= 1) {
                                        var v33 int = (fi22 ^ 32767)
                                        v2 -= v25
                                        _ = v33
                                    } else {
                                        v25 = -1
                                        v24 = (110 <= 42)
                                        v34 := GenS2{F0: genFn1(0)}
                                        _ = v34
                                    }
                                }
                            } else {
                                v24 = v24
                            }
                            var v35 int = fi22
                            _ = v35
                        } else {
                            v36 := GenS4{F0: int16(fi23), F1: 4294967295, F2: fi22}
                            _ = v36
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = v27
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            for fi37 := 0; fi37 < 20; fi37++ {
                v2 = v2
                var v38 uint64 = 4294967295
                v38 = (v38 / ((788) | 1))
                _ = fi37
                _ = v38
            }
            _ = fi22
        }
    } else {
        v2 -= (v2 >> ((2) & 63))
        v2 -= v2
        switch "foo bar" {
        case "":
            v39 := GenS5{F0: 1, F1: genFn1(v2), F2: 42}
            _ = v39
        }
    }
    delete(v1, "hello")
    fmt.Println(v2)
    v2 = (v2 | (v2 & v2))
    v40 := GenS2{F0: 65535}
    if false {
        for fi41 := 0; fi41 < 14; fi41++ {
            v2 += v2
            _ = fi41
        }
        v2 = 42
    }
    v2 -= (-(v2))
    if (2147483648 < 2147483647) {
        var v42 int = 722
        _ = v42
    } else {
        for fi43 := 0; fi43 < 2; fi43++ {
            v2 = 94
            switch (("\n\t" + ("foo bar" + "\n\t")) + "hello") {
            case "foo bar":
                _ = len(v1)
            case "\n\t":
                delete(v1, "\n\t")
            }
            var v44 float64 = (1e10 - (1e10 / (-2.25 - 0.0)))
            _ = fi43
            _ = v44
        }
        _ = len(v1)
        for fi45 := 0; fi45 < 6; fi45++ {
            v46 := GenS5{F0: 42, F1: genFn1(v2), F2: 127}
            v1[""] = "a"
            _ = fi45
            _ = v46
        }
    }
    if (false || (!(1e10 > 0.0))) {
        v47 := GenS1{F0: 285, F1: v2}
        v48 := GenS6{F0: -9223372036854775808, F1: 18446744073709551615}
        _ = v47
        _ = v48
    } else {
        var v49 float64 = float64(float32(0.1))
        if false {
            v50 := GenS7{F0: ("hello" + ("\n\t" + ("\n\t" + "a")))}
            var v51 float32 = 1e10
            _ = v50
            _ = v51
        } else {
            fmt.Println(v2, v2)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v2, v2)
        }
        _ = v49
    }
    v52 := GenS2{F0: genFn1((-(v2)))}
    var v53 uint64 = 295
    var v54 uint8 = 128
    switch "" {
    case "a":
        v1["a"] = ((("\n\t" + (("hello" + "\n\t") + ("\n\t" + "foo bar"))) + "hello") + "")
    case "":
        for fi55 := 0; fi55 < 3; fi55++ {
            v54 -= (v54 / (((v54 >> ((0) & 63))) | 1))
            _ = fi55
        }
    default:
        v2 += v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v52
    _ = v53
    _ = v54
}
