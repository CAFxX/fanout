package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint32; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8 }

func main() {
    var v1 string = "hello"
    var v2 uint64 = 18446744073709551615
    var v3 int64 = 9223372036854775807
    v4 := make(map[string]string)
    var v5 int32 = -1
    switch v1 {
    case "hello":
        if true {
            var v6 uint32 = 2147483648
            _ = len(v4)
            _ = v6
        }
    case "\n\t":
        v7 := []int{}
        _ = v7
    default:
        v2 = v2
    }
    switch -2147483648 {
    case -352:
        var v8 float64 = float64(1e10)
        _ = v8
    default:
        switch (v1 + ((v1 + (v1 + "a")) + "\n\t")) {
        case "a":
            var v9 float32 = ((3.14159 / 0.1) / 0.1)
            _ = v9
        case "\n\t":
            if true {
                v10 := GenS1{F0: ((-32768 != -1) && false)}
                switch v1 {
                case "hello":
                    v3 -= v3
                case "\n\t":
                    if (128 > 127) {
                        switch v1 {
                        case "a":
                            v4["hello"] = v1
                        case "":
                            v2 = (^(v2))
                        case "foo bar":
                            for fi11 := 0; fi11 < 1; fi11++ {
                                v5 -= (^(v5))
                                _ = fi11
                            }
                        default:
                            fmt.Println(v3)
                        }
                        for fi12 := 0; fi12 < 20; fi12++ {
                            v2 = (v2 * (v2 % ((802) | 1)))
                            v13 := []int{(fi12 | 9223372036854775807)}
                            _ = fi12
                            _ = v13
                        }
                        switch "\n\t" {
                        case "hello":
                            if true {
                                fmt.Println(v2, v5, v5)
                                var v14 float64 = 0.0
                                _ = v14
                            } else {
                                v5 = v5
                                v5 = -2147483648
                            }
                        case "foo bar":
                            v1 = "\n\t"
                        case "\n\t":
                            v15 := GenS1{F0: (2 == 492)}
                            _ = v15
                        default:
                            var v16 int8 = 26
                            _ = v16
                        }
                    } else {
                        var v17 float64 = 3.14159
                        _ = v17
                    }
                default:
                    for fi18 := 0; fi18 < 3; fi18++ {
                        v19 := GenS1{F0: false}
                        _ = fi18
                        _ = v19
                    }
                }
                _ = v10
            }
        default:
            fmt.Println(v2, v2, v1)
        }
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        v4[("\n\t" + v1)] = "\n\t"
        v21 := GenS1{F0: true}
        var v22 bool = true
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint8 = 16
    switch v1 {
    case "":
        v5 = (-(v5))
    case "a":
        v24 := []int{-487}
        _ = v24
    default:
        if (1 != 0) {
            var v25 string = "a"
            if (!true) {
                var v26 uint8 = 195
                _ = v26
            } else {
                v27 := GenS2{F0: v3, F1: 2147483648, F2: "hello"}
                _ = v27
            }
            _ = v25
        } else {
            if ((v23 | (^(v23))) < (v23 * 1)) {
                v3 = v3
            }
            switch (v1 + ("a" + "foo bar")) {
            case "foo bar":
                for fi28 := 0; fi28 < 14; fi28++ {
                    v29 := GenS1{F0: ((v3 & int64(v23)) > v3)}
                    v30 := GenS3{F0: (v3 - 620)}
                    v1 = "a"
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            case "hello":
                var v31 int32 = (v5 + v5)
                _ = v31
            default:
                var v32 bool = (!(v5 > (^(v5))))
                _ = v32
            }
        }
    }
    fmt.Println(v23)
    for fi33 := 0; fi33 < 3; fi33++ {
        v34 := GenS2{F0: (v3 & 1), F1: 2, F2: ("" + v1)}
        var v35 bool = true
        var v36 bool = v35
        _ = fi33
        _ = v34
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        switch v1 {
        case "foo bar":
            for fi38 := 0; fi38 < 10; fi38++ {
                v39 := v4[(v1 + "a")]
                _ = fi38
                _ = v39
            }
        }
        switch (v1 + v1) {
        case "":
            for fi40 := 0; fi40 < 17; fi40++ {
                var v41 uint16 = 42
                var v42 uint = 47
                v43 := GenS4{F0: uint16(v23)}
                _ = fi40
                _ = v41
                _ = v42
                _ = v43
            }
        }
        _ = fi37
    }
    v44 := []int{9223372036854775807, -702, int(v23), 2147483647}
    if (false || (((!(-778 > 0)) || (1e10 >= -2.25)) && false)) {
        var v45 bool = (true && true)
        _ = v45
    }
    var v46 uint = 42
    v44 = append(v44, -120)
    switch v1 {
    case "\n\t":
        v47 := GenS3{F0: (v3 - v3)}
        _ = v47
    default:
        switch 9223372036854775807 {
        case -1:
            if ((v3 / ((int64(v2)) | 1)) >= v3) {
                for fi48 := 0; fi48 < 19; fi48++ {
                    v2 = (v2 % ((0) | 1))
                    if false {
                        fmt.Println(v46)
                    } else {
                        if len(v44) > 0 {
                            v44[0] = (fi48 ^ fi48)
                        }
                        var v49 int64 = v3
                        _ = v49
                    }
                    _ = fi48
                }
                for fi50 := 0; fi50 < 8; fi50++ {
                    fmt.Println(v46, v2, v5)
                    v51 := v4["a"]
                    _ = fi50
                    _ = v51
                }
                if false {
                    v52 := GenS3{F0: (v3 - 42)}
                    switch ("foo bar" + v1) {
                    case "":
                        v5 = -2147483648
                    case "foo bar":
                        switch ("foo bar" + v1) {
                        case "\n\t":
                            v53 := GenS5{F0: -121}
                            _ = v53
                        default:
                            var v54 int64 = v3
                            _ = v54
                        }
                    }
                    var v55 int16 = -1
                    _ = v52
                    _ = v55
                } else {
                    v3 = (^(v3))
                }
            }
        }
    }
    if (!true) {
        v56 := GenS5{F0: int8(v5)}
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v5
}
