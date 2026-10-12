package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 int }
func genFn0(p0 bool, p1 int32, p2 int8) int {
    var v4 bool = p0
    switch 0 {
    case 32768:
        fmt.Println(p1, v4)
    case 2147483647:
        for fi5 := 0; fi5 < 14; fi5++ {
            for fi6 := 0; fi6 < 2; fi6++ {
                v7 := GenS1{F0: 84, F1: ""}
                var v8 float64 = 0.1
                var v9 uint16 = 642
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            var v10 string = "\n\t"
            _ = fi5
            _ = v10
        }
    }
    var v11 float64 = 0.0
    switch "" {
    case "hello":
        var v12 float32 = 0.1
        _ = v12
    case "":
        var v13 bool = (42 < 2)
        _ = v13
    case "a":
        for fi14 := 0; fi14 < 4; fi14++ {
            if p0 {
                switch ("foo bar" + "a") {
                case "foo bar":
                    if ((((("" + "foo bar") + "foo bar") + "foo bar") + "foo bar") != "hello") {
                        for fi15 := 0; fi15 < 15; fi15++ {
                            v16 := []int{}
                            _ = fi15
                            _ = v16
                        }
                        var v17 uint = 170
                        _ = v17
                    }
                case "hello":
                    var v18 string = ("\n\t" + "a")
                    _ = v18
                }
            }
            v19 := GenS2{F0: "a"}
            _ = fi14
            _ = v19
        }
    default:
        v4 = true
    }
    var v20 bool = ((((((0.1 + -2.25) + float32(1e10)) - float32(v11)) + 0.1) / ((-2.25 + 1e10) / 3.14159)) == 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v20
    _ = v4
    return 127
}
func genFn1(p0 uint64, p1 uint8, p2 int64) int16 {
    v21 := GenS1{F0: 18446744073709551615, F1: ("a" + ("\n\t" + ("a" + ("foo bar" + ""))))}
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 uint32 = 1
        p0 = p0
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 float64 = 1.5
            p0 = (^(p0))
            _ = fi24
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    v26 := []int{4294967295, -2147483648}
    var v27 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v26
    _ = v27
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 132
    var v3 float32 = 1e10
    v3 += (1.5 * float32(-2.25))
    for fi28 := 0; fi28 < 2; fi28++ {
        v29 := GenS2{F0: "\n\t"}
        fmt.Println(v3)
        if true {
            v30 := []int{}
            _ = v30
        }
        _ = fi28
        _ = v29
    }
    switch genFn0((0 >= 428), -2147483648, 0) {
    case -9223372036854775808:
        v31 := []int{4294967295}
        _ = v31
    default:
        for fi32 := 0; fi32 < 5; fi32++ {
            if (v3 != v3) {
                v33 := GenS3{F0: v2, F1: (v2 + 1), F2: 32767}
                if false {
                    var v34 uint = 0
                    for fi35 := 0; fi35 < 3; fi35++ {
                        switch fi32 {
                        case 0:
                            v34 += uint(v2)
                        case 4294967295:
                            var v36 uint8 = (v2 - v2)
                            _ = v36
                        case -23:
                            var v37 float64 = 0.0
                            _ = v37
                        }
                        var v38 int64 = -736
                        _ = fi35
                        _ = v38
                    }
                    switch "" {
                    case "a":
                        if true {
                            var v39 bool = true
                            _ = v39
                        } else {
                            v40 := GenS4{F0: fi32}
                            delete(v1, "foo bar")
                            _ = v40
                        }
                    case "":
                        if ((^(v34)) > v34) {
                            v41 := GenS1{F0: v34, F1: "a"}
                            _ = v41
                        }
                    case "\n\t":
                        var v42 uint16 = 617
                        _ = v42
                    }
                    _ = v34
                }
                v43 := GenS1{F0: 0, F1: "hello"}
                _ = v33
                _ = v43
            }
            for fi44 := 0; fi44 < 10; fi44++ {
                v3 = float32((1e10 + (float64((0.0 - (v3 / (0.1 + v3)))) - (-2.25 + 0.0))))
                if false {
                    var v45 string = "a"
                    var v46 int32 = -2147483648
                    v47 := []int{}
                    v48 := GenS4{F0: fi44}
                    _ = v45
                    _ = v46
                    _ = v47
                    _ = v48
                } else {
                    v2 -= (v2 * v2)
                    v3 = v3
                    v49 := GenS5{F0: -1, F1: 42}
                    _ = v49
                }
                fmt.Println(v2, v3, fi32)
                var v50 int = fi32
                v2 = (^(v2))
                _ = fi44
                _ = v50
            }
            _ = fi32
        }
    }
    if (!(4294967295 < 0)) {
        for fi51 := 0; fi51 < 20; fi51++ {
            fmt.Println(v3, fi51, fi51)
            _ = fi51
        }
    } else {
        v52 := GenS1{F0: 18446744073709551615, F1: ""}
        v53 := v1["\n\t"]
        _ = v52
        _ = v53
    }
    var v54 uint = 521
    v3 += -2.25
    fmt.Println(v54)
    v2 += (v2 * v2)
    fmt.Println(v3)
    _ = len(v1)
    var v55 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v55
}
