package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{32768, -759, 32768, 2147483647}
    var v3 uint16 = 123
    v4 := []string{("hello" + "a"), "\n\t", "foo bar", "\n\t"}
    var v5 float64 = 1e10
    v3 -= v3
    v5 = (v5 / float64(-2.25))
    v6 := GenS1{F0: 42, F1: (false || ((!false) && ((!(!true)) && (-189 >= 32767)))), F2: v3}
    v3 = v3
    var v7 bool = ("\n\t" >= ("" + "hello"))
    for fi8 := 0; fi8 < 12; fi8++ {
        v5 = (v5 + 1e10)
        _ = fi8
    }
    delete(v1, "\n\t")
    var v9 int8 = 88
    if v7 {
        _ = len(v2)
        var v10 uint = 65535
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS1{F0: (fi11 + (fi11 + (fi11 ^ -755))), F1: false, F2: (v3 - v3)}
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 int8 = 127
            switch "foo bar" {
            case "a":
                fmt.Println(v7, v3)
            case "\n\t":
                v9 += v14
            case "foo bar":
                v2 = append(v2, (fi13 % (((^(fi13))) | 1)))
            default:
                v5 += (float64(0.0) / v5)
            }
            v7 = v7
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    switch "\n\t" {
    case "\n\t":
        var v15 int32 = 2147483647
        _ = v15
    case "hello":
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 float64 = v5
            var v18 int = fi16
            _ = fi16
            _ = v17
            _ = v18
        }
    case "foo bar":
        if false {
            v3 += (v3 + (v3 << ((480) & 63)))
        } else {
            switch "" {
            case "":
                if len(v4) > 0 {
                    v4[0] = ""
                }
            case "foo bar":
                var v19 int16 = 42
                _ = v19
            }
        }
    }
    delete(v1, "")
    v20 := GenS2{F0: 1000}
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 17; fi22++ {
            var v23 uint64 = 0
            switch (fi21 % ((fi22) | 1)) {
            case -910:
                fmt.Println(fi21, v5, v7)
            default:
                fmt.Println(v7)
            }
            v24 := GenS2{F0: 32767}
            if len(v2) > 0 {
                v2[0] = -9223372036854775808
            }
            if v7 {
                var v25 float32 = 0.0
                _ = v25
            } else {
                fmt.Println(v3)
                v9 = v9
                if v7 {
                    var v26 int16 = -32768
                    v23 -= 2
                    switch fi21 {
                    case 42:
                        fmt.Println(v5, fi22, v5)
                    case 532:
                        var v27 int = (-(fi21))
                        _ = v27
                    default:
                        var v28 string = "\n\t"
                        _ = v28
                    }
                    _ = v26
                } else {
                    if ((166 != -129) || false) {
                        switch fi22 {
                        case 1:
                            fmt.Println(v3, v5)
                            fmt.Println(v9, v3)
                        default:
                            if (!(!(!v7))) {
                                switch (fi22 % (((-(fi21))) | 1)) {
                                case 1:
                                    var v29 int = fi22
                                    _ = v29
                                case 9223372036854775807:
                                    v5 = v5
                                case 128:
                                    v5 -= (3.14159 + ((1e10 / 1.0) - v5))
                                    if false {
                                        v30 := GenS2{F0: 1}
                                        fmt.Println(fi22, v9)
                                        _ = v30
                                    } else {
                                        if v7 {
                                            var v31 uint = 4294967295
                                            _ = v31
                                        } else {
                                            switch "hello" {
                                            case "foo bar":
                                                v32 := GenS1{F0: fi22, F1: true, F2: v3}
                                                _ = v32
                                            case "a":
                                                v3 = 631
                                            default:
                                                v23 -= (v23 * 34)
                                            }
                                            v3 = 971
                                        }
                                    }
                                }
                            } else {
                                v3 = v3
                            }
                            v1[""] = v7
                        }
                        if false {
                            v3 += v3
                            v5 -= v5
                            v33 := GenS1{F0: 127, F1: v7, F2: (v3 << ((4294967295) & 63))}
                            _ = v33
                        } else {
                            var v34 int16 = -32768
                            _ = v34
                        }
                    }
                    switch fi21 {
                    case -829:
                        v35 := GenS1{F0: 1, F1: v7, F2: (v3 / (((v3 << ((65535) & 63))) | 1))}
                        _ = v35
                    case 32767:
                        var v36 int64 = 1
                        _ = v36
                    default:
                        v37 := GenS3{F0: v7, F1: v7}
                        _ = v37
                    }
                    var v38 uint64 = 1
                    v39 := GenS2{F0: -32768}
                    _ = v38
                    _ = v39
                }
            }
            _ = len(v1)
            if (v3 != 101) {
                var v40 float64 = (v5 / float64(0.1))
                if (2 >= 894) {
                    v23 -= (^(v23))
                    var v41 int = (^(fi21))
                    v42 := GenS1{F0: 1, F1: (9223372036854775807 > -1), F2: v3}
                    v43 := GenS1{F0: fi22, F1: v7, F2: v3}
                    _ = v41
                    _ = v42
                    _ = v43
                }
                fmt.Println(v40)
                _ = v40
            }
            fmt.Println(fi22, v3, v23)
            if (!false) {
                v3 -= v3
            }
            _ = fi22
            _ = v23
            _ = v24
        }
        var v44 string = "hello"
        v45 := GenS4{F0: (!(288 >= 690))}
        _ = fi21
        _ = v44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
