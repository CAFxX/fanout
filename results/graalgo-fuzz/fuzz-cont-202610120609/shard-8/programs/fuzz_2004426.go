package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []int{-81}
    var v2 int32 = 42
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    var v5 uint = 1
    if true {
        if ((!true) && (0 <= 1000)) {
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 float32 = -2.25
                _ = fi6
                _ = v7
            }
            _ = len(v1)
        } else {
            v5 = v5
        }
    }
    v8 := GenS1{F0: 42}
    if true {
        v2 -= v2
        var v9 int32 = v2
        _ = v9
    } else {
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
        fmt.Println(v5, v5)
    }
    var v10 bool = (-2.25 != (-2.25 / (-2.25 + (1e10 * (1.5 / 3.14159)))))
    v10 = (((float32(0.1) + 1e10) / 0.1) <= 0.0)
    switch "foo bar" {
    case "hello":
        if (1 > 9223372036854775807) {
            if (v10 || v10) {
                if (false && (4294967295 >= 1)) {
                    for fi11 := 0; fi11 < 6; fi11++ {
                        v1 = append(v1, fi11)
                        _ = fi11
                    }
                    v12 := GenS2{F0: 759}
                    v5 = 0
                    _ = v12
                }
                fmt.Println(v2, v5)
                if v10 {
                    if (true || (v10 || false)) {
                        switch "\n\t" {
                        case "":
                            if (v10 && (42 < 274)) {
                                _ = len(v1)
                                var v13 uint = (^(v5))
                                if len(v1) > 0 {
                                    v1[0] = 42
                                }
                                _ = v13
                            }
                        default:
                            var v14 uint8 = 127
                            _ = v14
                        }
                    }
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v16 := v3[("" + "")]
                        _ = fi15
                        _ = v16
                    }
                } else {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        var v18 string = ("a" + "foo bar")
                        var v19 bool = (!v10)
                        for fi20 := 0; fi20 < 11; fi20++ {
                            switch fi17 {
                            case 376:
                                v19 = v19
                                if false {
                                    fmt.Println(v10, v18)
                                    fmt.Println(v18, v2, v5)
                                }
                            case 150:
                                v18 = v18
                            }
                            var v21 int16 = 635
                            _ = fi20
                            _ = v21
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    v2 = 42
                }
            }
        } else {
            var v22 int64 = -1
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 uint = 42
                _ = fi23
                _ = v24
            }
            v25 := GenS1{F0: 4294967295}
            _ = v22
            _ = v25
        }
    case "\n\t":
        var v26 uint8 = 42
        _ = v26
    case "":
        var v27 float32 = -2.25
        _ = v27
    default:
        switch ("" + ("foo bar" + "\n\t")) {
        case "a":
            var v28 int = 0
            _ = v28
        case "":
            if ((v5 << ((18446744073709551615) & 63)) < (v5 / (((v5 % ((v5) | 1))) | 1))) {
                v2 = (v2 - (v2 / ((v2) | 1)))
            }
        case "foo bar":
            if true {
                v2 -= (v2 / (((^(v2))) | 1))
            } else {
                var v29 uint16 = 42
                _ = v29
            }
        default:
            if false {
                if v10 {
                    var v30 uint32 = 4294967295
                    _ = v30
                }
                var v31 int32 = v2
                _ = v31
            }
        }
    }
    v32 := GenS2{F0: 0}
    v33 := GenS2{F0: -386}
    var v34 float64 = (-2.25 / (3.14159 * -2.25))
    if (true && true) {
        for fi35 := 0; fi35 < 13; fi35++ {
            fmt.Println(v2, v5)
            var v36 uint8 = 128
            _ = len(v3)
            _ = fi35
            _ = v36
        }
        if v10 {
            var v37 uint16 = 32767
            _ = v37
        } else {
            v5 -= v5
            for fi38 := 0; fi38 < 9; fi38++ {
                var v39 int = (fi38 | 9223372036854775807)
                if ((42 < 1) && v10) {
                    v40 := GenS1{F0: fi38}
                    _ = v40
                }
                v5 += v5
                _ = fi38
                _ = v39
            }
            if ((1 >= 293) && true) {
                v10 = v10
                v10 = (!(!(196 > 255)))
            }
        }
        v2 = (v2 / ((v2) | 1))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v8
}
