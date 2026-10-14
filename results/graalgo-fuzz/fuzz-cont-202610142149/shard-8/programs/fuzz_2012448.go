package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 uint64 }
func genFn0() bool {
    if false {
        var v6 int = 4294967295
        _ = v6
    } else {
        var v7 int64 = 1
        _ = v7
    }
    switch "a" {
    case "hello":
        var v8 int64 = -9223372036854775808
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 float32 = 1.5
            _ = fi9
            _ = v10
        }
    }
    if true {
        if true {
            var v11 float64 = 1.5
            v11 += v11
            var v12 uint32 = 2147483647
            _ = v11
            _ = v12
        }
    }
    return true
}
func genFn1(p0 uint8) uint16 {
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int64 = 42
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        if false {
            p0 = 42
            fmt.Println(p0, fi15)
            switch "" {
            case "a":
                for fi16 := 0; fi16 < 7; fi16++ {
                    if true {
                        var v17 float64 = 0.1
                        if false {
                            v18 := []int{(^(fi15))}
                            _ = v18
                        } else {
                            var v19 int64 = -1
                            _ = v19
                        }
                        _ = v17
                    }
                    fmt.Println(fi16, fi16, fi16)
                    v20 := GenS1{F0: 2, F1: 394}
                    _ = fi16
                    _ = v20
                }
            case "foo bar":
                switch "\n\t" {
                case "":
                    v21 := GenS2{F0: ("a" + "\n\t"), F1: 127, F2: -128}
                    _ = v21
                case "\n\t":
                    v22 := []int{fi15, (^(fi15)), (fi15 & fi15)}
                    _ = v22
                case "foo bar":
                    var v23 int32 = 0
                    _ = v23
                }
            }
        } else {
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 float32 = 1.5
                _ = fi24
                _ = v25
            }
            p0 += (p0 & (p0 % ((137) | 1)))
        }
        var v26 float32 = 3.14159
        _ = fi15
        _ = v26
    }
    v27 := []int{-1, -293}
    v28 := GenS1{F0: 42, F1: 42}
    var v29 bool = true
    _ = p0
    _ = v27
    _ = v28
    _ = v29
    return 338
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 173
    var v3 bool = (!((^(v2)) >= v2))
    var v4 uint8 = (^(v2))
    v5 := []string{"foo bar", "hello", "\n\t", ("hello" + "\n\t")}
    delete(v1, -169)
    for fi30 := 0; fi30 < 17; fi30++ {
        fmt.Println(fi30, v2)
        v3 = v3
        _ = fi30
    }
    v4 -= (v4 / ((v2) | 1))
    for fi31 := 0; fi31 < 15; fi31++ {
        v2 -= 151
        _ = fi31
    }
    v32 := GenS1{F0: 461, F1: 2}
    for fi33 := 0; fi33 < 10; fi33++ {
        fmt.Println(v2, v2, fi33)
        _ = fi33
    }
    switch ("a" + "hello") {
    case "a":
        if (!(!v3)) {
            for fi34 := 0; fi34 < 5; fi34++ {
                fmt.Println(v3)
                switch -2147483648 {
                case 90:
                    var v35 float64 = ((1.5 - 0.1) + (0.1 * 3.14159))
                    _ = v35
                default:
                    for fi36 := 0; fi36 < 6; fi36++ {
                        v1[fi34] = 1
                        _ = len(v1)
                        if (v3 && ("a" < ("" + "foo bar"))) {
                            var v37 uint64 = 2
                            _ = len(v1)
                            _ = v37
                        }
                        if (!false) {
                            if (((^(v4)) >= v2) || ((v2 | v2) < (v2 - 127))) {
                                v38 := GenS3{F0: 2}
                                fmt.Println(v3)
                                _ = v38
                            }
                            fmt.Println(v3, v3, fi36)
                            switch "\n\t" {
                            case "\n\t":
                                _ = len(v5)
                            default:
                                var v39 bool = true
                                _ = v39
                            }
                        }
                        _ = fi36
                    }
                }
                fmt.Println(v2)
                _ = fi34
            }
        } else {
            if (v3 && (481 > 18446744073709551615)) {
                if (!(18446744073709551615 > 641)) {
                    switch "a" {
                    case "a":
                        if (true && (v3 && (1 > 1))) {
                            v2 = v2
                            for fi40 := 0; fi40 < 12; fi40++ {
                                var v41 uint32 = 516
                                _ = fi40
                                _ = v41
                            }
                            v42 := GenS1{F0: 539, F1: 0}
                            _ = v42
                        } else {
                            fmt.Println(v3, v2, v4)
                        }
                    default:
                        v43 := GenS3{F0: 4294967295}
                        _ = v43
                    }
                }
                var v44 uint32 = 381
                _ = v44
            }
            for fi45 := 0; fi45 < 12; fi45++ {
                if len(v5) > 0 {
                    v5[0] = ("foo bar" + "hello")
                }
                switch (-(fi45)) {
                case 42:
                    for fi46 := 0; fi46 < 11; fi46++ {
                        var v47 int16 = 0
                        _ = fi46
                        _ = v47
                    }
                }
                _ = fi45
            }
        }
    case "hello":
        delete(v1, 42)
    }
    var v48 int64 = -9223372036854775808
    if v3 {
        fmt.Println(v48, v3)
        for fi49 := 0; fi49 < 4; fi49++ {
            for fi50 := 0; fi50 < 2; fi50++ {
                v51 := GenS2{F0: "hello", F1: 1, F2: 127}
                _ = fi50
                _ = v51
            }
            for fi52 := 0; fi52 < 12; fi52++ {
                _ = len(v1)
                _ = fi52
            }
            _ = fi49
        }
    } else {
        var v53 int8 = 127
        fmt.Println(v3, v48, v3)
        var v54 int16 = -931
        _ = v53
        _ = v54
    }
    _ = len(v5)
    fmt.Println(v3)
    for fi55 := 0; fi55 < 16; fi55++ {
        v4 -= 41
        v3 = (712 < 255)
        v5 = append(v5, ("" + ""))
        _ = fi55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v48
    _ = v5
}
