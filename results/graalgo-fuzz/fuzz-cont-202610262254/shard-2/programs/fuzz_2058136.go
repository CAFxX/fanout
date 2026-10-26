package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int8; F2 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 uint32; F2 bool }

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    var v3 uint16 = 65535
    v4 := make(map[uint32]bool)
    var v5 string = "hello"
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := []int{-2147483648, -497, -704}
        switch v5 {
        case "":
            var v8 int64 = 0
            _ = v8
        case "foo bar":
            v2 = v2
        case "a":
            var v9 float32 = (0.1 + 0.1)
            _ = v9
        }
        fmt.Println(v3)
        _ = fi6
        _ = v7
    }
    v3 += 419
    fmt.Println(v3)
    if v2 {
        var v10 int16 = int16(v3)
        if (v2 && ((!(1 < 2147483648)) || (42 == -609))) {
            switch (v5 + v5) {
            case "hello":
                v10 += (v10 * v10)
            case "foo bar":
                switch ((v5 + "") + v5) {
                case "\n\t":
                    var v11 float32 = 1e10
                    _ = v11
                case "":
                    switch ("" + "a") {
                    case "\n\t":
                        v12 := GenS1{F0: 0, F1: 42, F2: "foo bar"}
                        _ = v12
                    case "hello":
                        v3 = (^(v3))
                    }
                case "foo bar":
                    switch v5 {
                    case "foo bar":
                        switch -1 {
                        case -9223372036854775808:
                            v5 = "\n\t"
                        case 9223372036854775807:
                            v13 := GenS1{F0: 2, F1: -128, F2: v5}
                            _ = v13
                        }
                    case "\n\t":
                        var v14 int = 9223372036854775807
                        _ = v14
                    case "a":
                        var v15 int16 = 1
                        _ = v15
                    default:
                        switch 42 {
                        case 858:
                            for fi16 := 0; fi16 < 6; fi16++ {
                                for fi17 := 0; fi17 < 13; fi17++ {
                                    v3 += v3
                                    _ = fi17
                                }
                                var v18 string = v5
                                _ = fi16
                                _ = v18
                            }
                        case -403:
                            switch "" {
                            case "a":
                                v19 := GenS2{F0: 0}
                                _ = v19
                            default:
                                var v20 bool = (0 > 0)
                                _ = v20
                            }
                        }
                    }
                default:
                    v3 = 255
                }
            }
            fmt.Println(v10, v3, v2)
        }
        _ = v10
    }
    if v2 {
        v21 := GenS1{F0: 652, F1: 1, F2: "foo bar"}
        v2 = v2
        for fi22 := 0; fi22 < 12; fi22++ {
            for fi23 := 0; fi23 < 4; fi23++ {
                v24 := GenS2{F0: 1}
                fmt.Println(v2)
                v3 = (v3 * 65535)
                var v25 bool = (v2 || v2)
                _ = fi23
                _ = v24
                _ = v25
            }
            v2 = ("\n\t" > "")
            _ = fi22
        }
        _ = v21
    } else {
        var v26 float32 = -2.25
        _ = v26
    }
    if v2 {
        var v27 int = 530
        var v28 int32 = 2147483647
        delete(v1, "\n\t")
        _ = v27
        _ = v28
    } else {
        v2 = v2
        if v2 {
            switch v5 {
            case "":
                v3 = (v3 / ((1) | 1))
            case "\n\t":
                fmt.Println(v2, v2)
            case "foo bar":
                for fi29 := 0; fi29 < 19; fi29++ {
                    switch (v5 + v5) {
                    case "a":
                        v4[7] = (0.0 <= 0.0)
                    case "":
                        v30 := GenS1{F0: 78, F1: 1, F2: ((v5 + "a") + (v5 + v5))}
                        _ = v30
                    case "\n\t":
                        v31 := GenS3{F0: "", F1: 127, F2: int64(v3)}
                        _ = v31
                    default:
                        if v2 {
                            if true {
                                var v32 int16 = 803
                                _ = v32
                            } else {
                                v2 = v2
                            }
                            var v33 int32 = 1
                            fmt.Println(v2, v5)
                            _ = v33
                        } else {
                            var v34 string = "hello"
                            _ = v34
                        }
                    }
                    _ = fi29
                }
            }
            v35 := GenS1{F0: 18446744073709551615, F1: 100, F2: "\n\t"}
            v36 := GenS4{F0: 1}
            _ = v35
            _ = v36
        }
        var v37 float32 = float32(0.1)
        _ = v37
    }
    v38 := GenS5{F0: (-1 < -1), F1: 1, F2: true}
    var v39 bool = (116 >= 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
