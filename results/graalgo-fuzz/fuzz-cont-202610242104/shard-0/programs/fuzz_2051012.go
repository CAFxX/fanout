package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint; F2 bool }
func genFn0() uint8 {
    switch "" {
    case "hello":
        var v5 uint32 = 746
        _ = v5
    default:
        var v6 uint16 = 1
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int32 = 0
        v8 = v8
        var v9 int8 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{-148, 1}
    v11 := GenS1{F0: true, F1: -2147483648, F2: -989}
    _ = v10
    _ = v11
    return 255
}
func genFn1() string {
    switch 165 {
    case -2147483648:
        var v12 uint8 = 0
        _ = v12
    case -9223372036854775808:
        var v13 float64 = ((0.0 + (1e10 * 1e10)) + 0.1)
        _ = v13
    case -185:
        var v14 float32 = 1.5
        _ = v14
    default:
        var v15 int64 = -606
        _ = v15
    }
    if false {
        var v16 bool = (!false)
        v17 := GenS2{F0: 9223372036854775807}
        var v18 uint = 255
        _ = v16
        _ = v17
        _ = v18
    }
    switch "\n\t" {
    case "a":
        switch 0 {
        case -9223372036854775808:
            v19 := []int{-1, 4294967295, 127}
            _ = v19
        }
    case "\n\t":
        var v20 uint32 = 594
        _ = v20
    }
    return "hello"
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 732
    v3 := []string{"hello", "", "hello"}
    var v4 uint = 255
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 uint = v4
        _ = fi21
        _ = v22
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        v2 += (v2 + (v2 & v2))
        v24 := GenS2{F0: -1}
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        if (false || false) {
            v2 = v2
            v2 -= v2
        } else {
            var v26 int8 = 0
            _ = v26
        }
        _ = fi25
    }
    if false {
        fmt.Println(v2, v4)
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 int16 = -432
            v2 -= v2
            v28 += v28
            _ = fi27
            _ = v28
        }
        v2 -= v2
    } else {
        var v29 int8 = 1
        var v30 uint8 = genFn0()
        var v31 uint32 = 4294967295
        _ = v29
        _ = v30
        _ = v31
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        if false {
            v32 := GenS1{F0: (!(0 <= 777)), F1: 2147483647, F2: 0}
            _ = v32
        } else {
            if (2147483647 >= 42) {
                fmt.Println(v4, v4)
                v4 = (v4 << (((v4 + 18446744073709551615)) & 63))
            } else {
                var v33 int32 = 1
                switch 2147483647 {
                case 1:
                    var v34 uint16 = 32767
                    _ = v34
                case -209:
                    var v35 float32 = ((float32((((1.5 / 0.1) + (1e10 + (((0.0 / 1.0) * 3.14159) / (0.1 / (-2.25 / 1.0))))) / 3.14159)) - 1.5) / (-2.25 / (0.0 - -2.25)))
                    _ = v35
                }
                for fi36 := 0; fi36 < 14; fi36++ {
                    switch (fi36 ^ -662) {
                    case -917:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -2147483648:
                        var v37 uint = (v4 & (v4 >> (((v4 & 18446744073709551615)) & 63)))
                        _ = v37
                    }
                    _ = fi36
                }
                _ = v33
            }
            for fi38 := 0; fi38 < 16; fi38++ {
                v4 = (v4 & 4294967295)
                switch fi38 {
                case -1:
                    _ = len(v3)
                case 4294967295:
                    if (!false) {
                        var v39 int8 = -110
                        for fi40 := 0; fi40 < 14; fi40++ {
                            switch "foo bar" {
                            case "a":
                                var v41 uint8 = genFn0()
                                _ = v41
                            case "":
                                v42 := GenS1{F0: (!true), F1: -2147483648, F2: 1}
                                _ = v42
                            case "\n\t":
                                var v43 int8 = (v39 - -128)
                                _ = v43
                            default:
                                switch "hello" {
                                case "foo bar":
                                    v39 = v39
                                }
                            }
                            _ = fi40
                        }
                        _ = v39
                    } else {
                        var v44 uint64 = 1
                        for fi45 := 0; fi45 < 18; fi45++ {
                            var v46 int64 = -468
                            _ = fi45
                            _ = v46
                        }
                        _ = v44
                    }
                }
                fmt.Println(v2, v4, v4)
                _ = fi38
            }
            var v47 string = (("\n\t" + "foo bar") + "")
            _ = v47
        }
    case 32768:
        v48 := GenS1{F0: true, F1: -2147483648, F2: 32767}
        _ = v48
    default:
        v2 = (^(v2))
    }
    for fi49 := 0; fi49 < 16; fi49++ {
        var v50 string = "\n\t"
        _ = fi49
        _ = v50
    }
    v3 = append(v3, ("a" + (("" + "") + "")))
    var v51 uint8 = 21
    v4 -= (v4 - (v4 % (((v4 << ((v4) & 63))) | 1)))
    v52 := GenS3{F0: 2, F1: v4, F2: true}
    v2 += 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
}
