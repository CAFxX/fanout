package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 uint32; F2 string }
func genFn0(p0 uint16) bool {
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            p0 += p0
            var v9 bool = true
            v10 := []int{(fi7 & (fi8 | 9223372036854775807)), fi7, fi8}
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: fi7, F1: 127, F2: "a"}
        _ = fi7
        _ = v11
    }
    switch -2147483648 {
    case 9223372036854775807:
        fmt.Println(p0)
    case 0:
        var v12 uint8 = 247
        _ = v12
    case 4294967295:
        if true {
            v13 := GenS1{F0: 388, F1: -10, F2: "a"}
            switch -247 {
            case 2147483647:
                v14 := GenS2{F0: (!(!true))}
                _ = v14
            case 893:
                switch -698 {
                case -989:
                    p0 = 1
                case 127:
                    for fi15 := 0; fi15 < 3; fi15++ {
                        fmt.Println(fi15, fi15)
                        _ = fi15
                    }
                case -881:
                    if true {
                        for fi16 := 0; fi16 < 20; fi16++ {
                            if (uint8(p0) > 157) {
                                p0 += p0
                                v17 := GenS1{F0: (fi16 >> ((657) & 63)), F1: 109, F2: ("hello" + "foo bar")}
                                var v18 uint8 = 249
                                _ = v17
                                _ = v18
                            } else {
                                v19 := []int{}
                                fmt.Println(fi16, fi16)
                                var v20 float32 = -2.25
                                _ = v19
                                _ = v20
                            }
                            p0 += 0
                            _ = fi16
                        }
                        var v21 uint64 = 2
                        _ = v21
                    }
                default:
                    v22 := GenS3{F0: 32767, F1: 2147483648, F2: (("" + ("hello" + "hello")) + ("a" + ""))}
                    _ = v22
                }
            default:
                v23 := []int{-822, -267, 127, -1}
                _ = v23
            }
            _ = v13
        }
    }
    v24 := GenS1{F0: -1, F1: 1, F2: ("a" + "hello")}
    v25 := []int{}
    var v26 float64 = 3.14159
    _ = p0
    _ = v24
    _ = v25
    _ = v26
    return ((true || (!(true && (-215 != -9223372036854775808)))) && false)
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    var v3 int = 2147483647
    v4 := []uint32{}
    var v5 string = "foo bar"
    v6 := make(map[string]int)
    var v27 uint16 = 32767
    fmt.Println(v3)
    if true {
        v28 := GenS2{F0: ((42 <= 42) || (1 <= 2147483647))}
        if (!(false || true)) {
            v5 = v5
            v5 = v5
        } else {
            var v29 float64 = 3.14159
            for fi30 := 0; fi30 < 8; fi30++ {
                var v31 uint = 175
                _ = fi30
                _ = v31
            }
            var v32 int32 = -1
            _ = v29
            _ = v32
        }
        for fi33 := 0; fi33 < 4; fi33++ {
            v27 -= (^(v27))
            if (((!(0.1 > 1e10)) && (fi33 <= 1)) && (false || (!false))) {
                switch 1 {
                case 2147483647:
                    fmt.Println(v3, fi33, fi33)
                case 127:
                    switch ((((v5 + v5) + v5) + v5) + ("" + v5)) {
                    case "":
                        v34 := GenS3{F0: -266, F1: 0, F2: ((v5 + v5) + ((v5 + "foo bar") + (v5 + v5)))}
                        _ = v34
                    case "foo bar":
                        v35 := GenS3{F0: -32768, F1: uint32(v27), F2: (v5 + v5)}
                        _ = v35
                    default:
                        v3 = (fi33 / (((v3 * (v3 | 128))) | 1))
                    }
                case -325:
                    v36 := GenS3{F0: -226, F1: 463, F2: ((v5 + (v5 + v5)) + "")}
                    _ = v36
                }
                fmt.Println(v27, fi33, fi33)
                _ = len(v6)
            }
            _ = fi33
        }
        _ = v28
    } else {
        v37 := GenS3{F0: 42, F1: 2147483647, F2: (v5 + ((v5 + v5) + ("a" + "foo bar")))}
        v3 = (v3 + (v3 ^ v3))
        if (0 > 9223372036854775807) {
            v5 = v5
        } else {
            var v38 uint8 = 127
            switch v3 {
            case 128:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case -2147483648:
                if true {
                    var v39 int16 = -649
                    for fi40 := 0; fi40 < 13; fi40++ {
                        v3 = (v3 - (fi40 - (fi40 | 1)))
                        switch "hello" {
                        case "foo bar":
                            _ = len(v4)
                        case "\n\t":
                            switch v5 {
                            case "hello":
                                if len(v4) > 0 {
                                    v4[0] = 1
                                }
                            case "":
                                v27 = v27
                            }
                        case "a":
                            v4 = append(v4, 2147483647)
                        }
                        _ = fi40
                    }
                    _ = v39
                }
            }
            fmt.Println(v38)
            _ = v38
        }
        _ = v37
    }
    var v41 uint16 = (^(v27))
    fmt.Println(v5)
    v42 := GenS1{F0: v3, F1: 100, F2: (((v5 + ("" + ("\n\t" + v5))) + v5) + v5)}
    var v43 uint = 4294967295
    var v44 float64 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
