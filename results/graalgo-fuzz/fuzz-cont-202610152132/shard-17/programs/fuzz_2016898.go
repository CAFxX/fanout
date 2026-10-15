package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int) int8 {
    v4 := GenS1{F0: 32767, F1: -568, F2: 42}
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            if true {
                p0 = p0
            } else {
                if true {
                    switch "hello" {
                    case "\n\t":
                        if (false || false) {
                            if false {
                                v7 := GenS1{F0: -1, F1: -936, F2: 9223372036854775807}
                                _ = v7
                            } else {
                                switch "\n\t" {
                                case "\n\t":
                                    p0 = p0
                                case "foo bar":
                                    p0 = 4294967295
                                    p0 += 0
                                }
                                p0 = fi5
                            }
                            p0 += 9223372036854775807
                            var v8 bool = true
                            _ = v8
                        } else {
                            p0 -= p0
                            p0 += p0
                            p0 -= 32768
                            p0 -= 0
                            var v9 uint8 = 42
                            switch fi5 {
                            case 127:
                                v10 := GenS1{F0: -1, F1: 830, F2: -152}
                                _ = v10
                            case -308:
                                if ("\n\t" != "foo bar") {
                                    v11 := GenS1{F0: 32767, F1: -1, F2: -1}
                                    switch (fi6 ^ 128) {
                                    case -9223372036854775808:
                                        v12 := []int{}
                                        _ = v12
                                    case 1:
                                        v13 := GenS1{F0: -422, F1: 42, F2: -154}
                                        _ = v13
                                    case 128:
                                        v9 -= 1
                                    default:
                                        var v14 int16 = 926
                                        _ = v14
                                    }
                                    p0 = (p0 - fi5)
                                    _ = v11
                                }
                            }
                            var v15 float64 = ((-2.25 - 0.0) * float64(0.1))
                            _ = v15
                            _ = v9
                        }
                    }
                    fmt.Println(p0, p0)
                }
                fmt.Println(p0, p0)
                p0 = fi5
            }
            v16 := []int{fi6, (p0 - (fi5 >> ((4294967295) & 63)))}
            _ = fi6
            _ = v16
        }
        var v17 int16 = -94
        _ = fi5
        _ = v17
    }
    v18 := GenS1{F0: -1, F1: 0, F2: 1}
    switch ((("\n\t" + "a") + "\n\t") + "\n\t") {
    case "a":
        if ((true || (-532 > 0)) || (236 >= 172)) {
            p0 += 32767
        }
    case "foo bar":
        p0 = (p0 << ((1) & 63))
    case "":
        var v19 int64 = 1
        _ = v19
    default:
        for fi20 := 0; fi20 < 10; fi20++ {
            for fi21 := 0; fi21 < 7; fi21++ {
                v22 := GenS1{F0: -269, F1: 0, F2: -451}
                _ = fi21
                _ = v22
            }
            var v23 uint16 = 65535
            _ = fi20
            _ = v23
        }
    }
    _ = p0
    _ = v18
    _ = v4
    return 127
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    var v3 int64 = 9223372036854775807
    v3 = (v3 >> ((1) & 63))
    v24 := []int{42, 655, 126}
    var v25 int32 = 2147483647
    v3 = v3
    v25 += v25
    fmt.Println(v3)
    v24 = append(v24, 158)
    var v26 uint = 2
    var v27 string = "foo bar"
    v28 := GenS1{F0: 882, F1: (v25 / (((v25 << (((v26 | v26)) & 63))) | 1)), F2: (^(v3))}
    switch 127 {
    case 4294967295:
        switch 128 {
        case 4294967295:
            v29 := GenS1{F0: -32768, F1: v25, F2: (v3 & v3)}
            _ = v29
        case -96:
            var v30 int8 = -1
            _ = v30
        default:
            switch v27 {
            case "foo bar":
                if false {
                    for fi31 := 0; fi31 < 16; fi31++ {
                        v26 -= (v26 / ((4294967295) | 1))
                        _ = fi31
                    }
                } else {
                    for fi32 := 0; fi32 < 19; fi32++ {
                        v27 = "\n\t"
                        _ = fi32
                    }
                    v27 = ""
                }
            case "":
                for fi33 := 0; fi33 < 12; fi33++ {
                    v26 = v26
                    v24 = append(v24, 32767)
                    var v34 uint64 = 4294967295
                    _ = fi33
                    _ = v34
                }
            default:
                fmt.Println(v25, v27, v27)
            }
        }
    default:
        var v35 uint = 201
        _ = v35
    }
    var v36 uint64 = 1
    v37 := GenS1{F0: 306, F1: (v25 - 2147483647), F2: v3}
    for fi38 := 0; fi38 < 18; fi38++ {
        var v39 uint = (v26 & 18446744073709551615)
        v40 := GenS2{F0: (v36 ^ (v36 << (((v26 * 1)) & 63))), F1: 2147483647}
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v37
}
