package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 1.5
    v3 := make(map[string]uint32)
    var v4 bool = false
    v5 := make(map[uint32]string)
    v4 = (52 <= 127)
    var v6 string = "hello"
    v7 := []int{}
    v8 := GenS1{F0: 0, F1: v6}
    v9 := GenS1{F0: 18446744073709551615, F1: ""}
    for fi10 := 0; fi10 < 6; fi10++ {
        switch v6 {
        case "hello":
            if v4 {
                v2 = 1e10
                v2 = v2
            } else {
                for fi11 := 0; fi11 < 6; fi11++ {
                    var v12 uint16 = 255
                    if v4 {
                        v13 := v1[fi10]
                        _ = v13
                    }
                    _ = fi11
                    _ = v12
                }
                var v14 int32 = 38
                _ = v14
            }
        }
        if ((-(fi10)) != fi10) {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            switch fi10 {
            case -938:
                switch v6 {
                case "hello":
                    fmt.Println(v2, fi10)
                }
            case -230:
                switch fi10 {
                case 9223372036854775807:
                    v7 = append(v7, fi10)
                case 320:
                    v4 = v4
                default:
                    switch fi10 {
                    case 319:
                        v6 = v6
                    case -9223372036854775808:
                        switch "\n\t" {
                        case "\n\t":
                            var v15 int64 = -1
                            _ = v15
                        case "foo bar":
                            var v16 string = v6
                            _ = v16
                        }
                    case 4294967295:
                        for fi17 := 0; fi17 < 10; fi17++ {
                            v7 = append(v7, (fi17 % ((fi10) | 1)))
                            _ = fi17
                        }
                    }
                }
            case 32768:
                if v4 {
                    switch fi10 {
                    case -390:
                        var v18 uint32 = 42
                        _ = v18
                    default:
                        var v19 int8 = -1
                        _ = v19
                    }
                    if v4 {
                        for fi20 := 0; fi20 < 15; fi20++ {
                            v21 := GenS1{F0: 255, F1: ""}
                            v22 := GenS1{F0: 112, F1: ("a" + v6)}
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 8; fi23++ {
                            var v24 int32 = 575
                            if (v4 && v4) {
                                v25 := GenS1{F0: 2, F1: ""}
                                _ = v25
                            }
                            switch v6 {
                            case "a":
                                var v26 int64 = 9223372036854775807
                                _ = v26
                            case "\n\t":
                                v27 := GenS1{F0: 2, F1: (v6 + ("a" + "a"))}
                                _ = v27
                            case "hello":
                                var v28 bool = v4
                                _ = v28
                            default:
                                var v29 int32 = -470
                                _ = v29
                            }
                            _ = fi23
                            _ = v24
                        }
                        v30 := GenS1{F0: 42, F1: "a"}
                        _ = v30
                    } else {
                        var v31 float64 = v2
                        _ = v31
                    }
                } else {
                    v5[65535] = v6
                    v6 = v6
                    if v4 {
                        v4 = (515 <= 1)
                        v2 = 0.0
                        v32 := GenS1{F0: 42, F1: v6}
                        _ = v32
                    } else {
                        if (-109 < -826) {
                            fmt.Println(fi10, v2, v2)
                            for fi33 := 0; fi33 < 19; fi33++ {
                                v6 = (v6 + "a")
                                var v34 float32 = (float32((v2 * 1.5)) + -2.25)
                                _ = fi33
                                _ = v34
                            }
                        } else {
                            if v4 {
                                v35 := GenS1{F0: 158, F1: v6}
                                _ = v35
                            } else {
                                for fi36 := 0; fi36 < 1; fi36++ {
                                    var v37 int = -9223372036854775808
                                    _ = fi36
                                    _ = v37
                                }
                            }
                            fmt.Println(v4, fi10)
                        }
                        fmt.Println(v2, fi10, v4)
                        for fi38 := 0; fi38 < 17; fi38++ {
                            v2 += 1.5
                            fmt.Println(v6)
                            _ = fi38
                        }
                    }
                }
            default:
                if true {
                    for fi39 := 0; fi39 < 1; fi39++ {
                        if v4 {
                            _ = len(v5)
                            var v40 int = -454
                            _ = v40
                        }
                        _ = fi39
                    }
                    var v41 float32 = float32(v2)
                    _ = v41
                } else {
                    v42 := GenS1{F0: 40, F1: v6}
                    _ = v42
                }
            }
        }
        _ = fi10
    }
    var v43 string = ""
    v43 = (v43 + "a")
    fmt.Println(v2)
    if true {
        var v44 bool = false
        var v45 int16 = -32768
        var v46 float32 = (0.1 * (1.5 * 0.0))
        _ = v44
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 20; fi47++ {
        v4 = (!(-62 < 106))
        _ = fi47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
