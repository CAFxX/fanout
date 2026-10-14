package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int; F2 int32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{-1}
    v3 := []int{}
    var v4 uint = 4294967295
    var v5 uint = (v4 / (((v4 - v4)) | 1))
    if true {
        v6 := GenS1{F0: 32767, F1: 1}
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 bool = ((!(true && true)) && true)
                fmt.Println(fi8, v4)
                _ = fi8
                _ = v9
            }
            switch (-(fi7)) {
            case 545:
                v5 = 255
            case 531:
                fmt.Println(fi7, fi7, v5)
            case 128:
                v4 -= 43
            default:
                fmt.Println(v5)
            }
            for fi10 := 0; fi10 < 18; fi10++ {
                fmt.Println(v5, v4)
                v11 := GenS1{F0: 382, F1: 1}
                _ = fi10
                _ = v11
            }
            _ = fi7
        }
        var v12 float64 = (((0.0 + 1e10) - (-2.25 + (-2.25 * -2.25))) + ((0.1 / 1e10) - ((0.1 + 1e10) + 0.0)))
        _ = v12
        _ = v6
    } else {
        v13 := v1[""]
        var v14 int32 = -1
        _ = v13
        _ = v14
    }
    fmt.Println(v5)
    v15 := GenS1{F0: 32768, F1: -920}
    switch 42 {
    case 9223372036854775807:
        var v16 int64 = 339
        _ = v16
    case 32767:
        if (0 <= 392) {
            v4 -= (v4 & 2)
            if (true && (42 < 48)) {
                v4 = (v4 - 364)
                _ = len(v2)
                var v17 int64 = 1
                _ = v17
            } else {
                switch ("\n\t" + "a") {
                case "hello":
                    if true {
                        var v18 uint64 = 18446744073709551615
                        switch -1 {
                        case 127:
                            if false {
                                v5 += v4
                                fmt.Println(v5)
                            } else {
                                v19 := GenS1{F0: 127, F1: 832}
                                for fi20 := 0; fi20 < 8; fi20++ {
                                    var v21 int8 = -122
                                    v21 = v21
                                    var v22 int32 = 516
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                                _ = v19
                            }
                        case 478:
                            switch ("foo bar" + "hello") {
                            case "":
                                v23 := v1["foo bar"]
                                _ = v23
                            case "foo bar":
                                v24 := GenS1{F0: 32767, F1: int32(v18)}
                                _ = v24
                            }
                        case 32767:
                            for fi25 := 0; fi25 < 10; fi25++ {
                                v5 += (v5 | 964)
                                v18 += 4294967295
                                switch fi25 {
                                case -73:
                                    v26 := GenS2{F0: 577, F1: 215}
                                    _ = v26
                                case 128:
                                    for fi27 := 0; fi27 < 16; fi27++ {
                                        switch (fi27 - (fi25 & (fi25 * fi27))) {
                                        case -1:
                                            if ((fi27 - (fi27 - -604)) != fi27) {
                                                _ = len(v1)
                                                var v28 string = "hello"
                                                v29 := GenS3{F0: 64, F1: 0, F2: 248}
                                                _ = v28
                                                _ = v29
                                            }
                                        case -195:
                                            v1["\n\t"] = true
                                        case 34:
                                            fmt.Println(v5, v18)
                                        default:
                                            var v30 string = (("\n\t" + (((("hello" + ("" + "")) + "") + ("a" + "a")) + ("\n\t" + "a"))) + "foo bar")
                                            _ = v30
                                        }
                                        _ = fi27
                                    }
                                }
                                _ = fi25
                            }
                        }
                        _ = v18
                    } else {
                        fmt.Println(v5, v4)
                        var v31 uint = 1
                        _ = v31
                    }
                case "foo bar":
                    var v32 int = 730
                    _ = v32
                default:
                    for fi33 := 0; fi33 < 17; fi33++ {
                        var v34 uint16 = 2
                        v4 = v4
                        _ = fi33
                        _ = v34
                    }
                }
                switch "" {
                case "\n\t":
                    v4 -= v4
                }
            }
        }
    }
    fmt.Println(v4, v4)
    fmt.Println(v4)
    v35 := GenS3{F0: 39, F1: -1, F2: 6}
    fmt.Println(v5)
    for fi36 := 0; fi36 < 11; fi36++ {
        fmt.Println(v4)
        fmt.Println(v4, v4)
        _ = fi36
    }
    if (!(0 == 915)) {
        v5 += v5
    }
    var v37 uint = (v5 >> ((v4) & 63))
    for fi38 := 0; fi38 < 1; fi38++ {
        v4 = v37
        var v39 int = fi38
        _ = fi38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v5
}
