package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint32 }
func genFn0(p0 int8, p1 int8, p2 int16) int {
    p1 = (-(p1))
    p1 = (^(p0))
    fmt.Println(p0, p1)
    v7 := []int{-9223372036854775808, -2147483648, 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 46
}

func main() {
    var v1 int64 = 42
    v2 := []string{}
    var v3 int64 = -218
    v4 := make(map[string]int)
    var v5 int8 = -1
    v6 := make(map[string]int)
    var v8 int16 = -338
    if false {
        if false {
            var v9 uint32 = 0
            _ = v9
        }
        v3 = (v3 % (((v1 / ((v1) | 1))) | 1))
        v1 = v3
    } else {
        v8 = 1000
        for fi10 := 0; fi10 < 19; fi10++ {
            if (130 == 65535) {
                fmt.Println(v5, fi10, v1)
                switch "hello" {
                case "":
                    for fi11 := 0; fi11 < 13; fi11++ {
                        var v12 float32 = (0.0 - float32(3.14159))
                        switch "" {
                        case "\n\t":
                            fmt.Println(v12)
                        case "foo bar":
                            v13 := GenS1{F0: 1, F1: -2147483648}
                            _ = v13
                        case "hello":
                            fmt.Println(v5)
                        }
                        fmt.Println(v12, fi11)
                        var v14 uint = 4294967295
                        _ = fi11
                        _ = v12
                        _ = v14
                    }
                }
            }
            _ = fi10
        }
    }
    fmt.Println(v1)
    var v15 uint8 = 59
    switch "\n\t" {
    case "\n\t":
        if (!true) {
            var v16 string = ""
            _ = v16
        }
    case "a":
        v17 := GenS1{F0: 255, F1: 2147483647}
        _ = v17
    case "hello":
        v15 -= v15
    }
    var v18 uint8 = (v15 % (((v15 + 127)) | 1))
    if (v5 != (v5 % (((-(v5))) | 1))) {
        if len(v2) > 0 {
            v2[0] = ("hello" + "")
        }
        if false {
            var v19 float64 = 1e10
            v20 := GenS1{F0: 1, F1: 0}
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := GenS1{F0: 32767, F1: 2147483647}
            _ = len(v4)
            _ = fi21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        delete(v6, (("" + ("" + "")) + "\n\t"))
        _ = fi23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v25 := GenS2{F0: (v3 ^ v1)}
        _ = fi24
        _ = v25
    }
    v1 -= (v1 + v1)
    for fi26 := 0; fi26 < 2; fi26++ {
        for fi27 := 0; fi27 < 20; fi27++ {
            _ = len(v4)
            v28 := GenS2{F0: v3}
            _ = fi27
            _ = v28
        }
        v6["foo bar"] = int(v8)
        _ = fi26
    }
    v8 = 94
    var v29 float32 = (3.14159 + (-2.25 - 3.14159))
    switch 499 {
    case 127:
        v30 := GenS2{F0: 9223372036854775807}
        _ = v30
    case 0:
        var v31 int64 = (v3 | -926)
        _ = v31
    case 695:
        _ = len(v4)
    default:
        _ = len(v2)
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        for fi33 := 0; fi33 < 12; fi33++ {
            v34 := GenS1{F0: 255, F1: 42}
            v1 += v3
            if false {
                if (!true) {
                    fmt.Println(v18, v18)
                    v35 := GenS1{F0: 2, F1: 318}
                    switch -761 {
                    case 32768:
                        v36 := GenS3{F0: (v8 ^ 581), F1: (v5 / ((v5) | 1))}
                        _ = v36
                    case 127:
                        var v37 int16 = (v8 & (v8 >> ((0) & 63)))
                        _ = v37
                    default:
                        v38 := GenS2{F0: 222}
                        _ = v38
                    }
                    _ = v35
                } else {
                    var v39 uint16 = 843
                    _ = v39
                }
                v1 = (v1 | v1)
                v2 = append(v2, "a")
            }
            _ = fi33
            _ = v34
        }
        switch (("hello" + "a") + "foo bar") {
        case "hello":
            for fi40 := 0; fi40 < 5; fi40++ {
                if false {
                    var v41 bool = false
                    _ = v41
                }
                _ = fi40
            }
        default:
            switch (fi32 ^ 402) {
            case -164:
                switch fi32 {
                case 2147483647:
                    for fi42 := 0; fi42 < 20; fi42++ {
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            v43 := GenS4{F0: 415, F1: "hello"}
                            fmt.Println(v8)
                            _ = v43
                        }
                        fmt.Println(v18, v8, v5)
                        _ = fi42
                    }
                case 990:
                    if true {
                        v44 := GenS5{F0: 65535, F1: 2, F2: 336}
                        _ = v44
                    }
                default:
                    v1 = v3
                }
            case 751:
                for fi45 := 0; fi45 < 4; fi45++ {
                    if false {
                        v15 += (v15 % ((0) | 1))
                    } else {
                        var v46 float64 = 0.1
                        var v47 string = "foo bar"
                        if (1 < -2147483648) {
                            v2 = append(v2, "foo bar")
                            switch fi32 {
                            case 2147483647:
                                v3 += v3
                            case 127:
                                var v48 uint8 = (v18 * 27)
                                _ = v48
                            case 32767:
                                v49 := GenS5{F0: 4294967295, F1: 0, F2: 2}
                                _ = v49
                            }
                            var v50 int8 = v5
                            _ = v50
                        }
                        _ = v46
                        _ = v47
                    }
                    fmt.Println(v29, v5)
                    switch (fi32 - -9223372036854775808) {
                    case -624:
                        fmt.Println(v3, v29)
                    case -2147483648:
                        v5 += -128
                    case 45:
                        v1 = v1
                    }
                    _ = fi45
                }
            default:
                switch (fi32 << ((618) & 63)) {
                case 32768:
                    fmt.Println(v1)
                case 1:
                    v29 = -2.25
                case -1:
                    v5 = v5
                default:
                    var v51 uint32 = 57
                    _ = v51
                }
            }
        }
        v52 := GenS2{F0: v1}
        _ = fi32
        _ = v52
    }
    var v53 int64 = v3
    fmt.Println(v1, v1, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v8
}
