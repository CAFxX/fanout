package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int64; F1 int64; F2 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 int16 }

func main() {
    var v1 uint8 = 74
    v2 := []string{}
    var v3 int32 = -668
    var v4 int8 = 127
    fmt.Println(v4, v1)
    var v5 uint8 = v1
    v5 += (v5 ^ 42)
    if false {
        v3 -= v3
        for fi6 := 0; fi6 < 11; fi6++ {
            switch "foo bar" {
            case "foo bar":
                v7 := GenS1{F0: 4294967295, F1: "hello", F2: -385}
                _ = v7
            case "\n\t":
                v8 := GenS2{F0: 0}
                _ = v8
            case "hello":
                if false {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        v10 := GenS1{F0: 552, F1: "a", F2: -9223372036854775808}
                        v11 := GenS1{F0: 4294967295, F1: ("\n\t" + ""), F2: -799}
                        if len(v2) > 0 {
                            v2[0] = (("\n\t" + ("\n\t" + "\n\t")) + ("a" + ("hello" + ("\n\t" + "a"))))
                        }
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    if false {
                        v2 = append(v2, "")
                        if (!(false || (9223372036854775807 != -578))) {
                            for fi12 := 0; fi12 < 1; fi12++ {
                                switch "hello" {
                                case "\n\t":
                                    var v13 uint32 = 409
                                    _ = v13
                                case "hello":
                                    fmt.Println(fi6, v1, v4)
                                case "foo bar":
                                    fmt.Println(v3)
                                    var v14 bool = ((^(fi6)) >= 4294967295)
                                    v15 := GenS3{F0: 117, F1: ("a" + ("hello" + "a"))}
                                    if false {
                                        v16 := GenS4{F0: 220, F1: int64(fi12), F2: 13}
                                        switch "hello" {
                                        case "":
                                            v1 -= (v1 / ((v5) | 1))
                                        default:
                                            var v17 uint8 = v1
                                            _ = v17
                                        }
                                        _ = v16
                                    }
                                    if false {
                                        v18 := GenS5{F0: v3}
                                        _ = v18
                                    } else {
                                        v4 = (v4 | v4)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = len(v2)
                                        v3 = 2147483647
                                    }
                                    _ = v14
                                    _ = v15
                                }
                                var v19 int = fi6
                                v20 := GenS2{F0: 1}
                                _ = fi12
                                _ = v19
                                _ = v20
                            }
                            var v21 int8 = v4
                            v21 += (v4 ^ (v4 * v21))
                            _ = v21
                        } else {
                            v5 = v1
                        }
                        var v22 uint64 = 280
                        _ = v22
                    } else {
                        var v23 float64 = 3.14159
                        _ = v23
                    }
                }
            default:
                v24 := GenS6{F0: ("" + "")}
                _ = v24
            }
            v3 += v3
            if false {
                v1 += 248
                fmt.Println(v4)
                for fi25 := 0; fi25 < 7; fi25++ {
                    var v26 float64 = 3.14159
                    _ = fi25
                    _ = v26
                }
            }
            _ = fi6
        }
    }
    v3 = (^(v3))
    switch -2147483648 {
    case -9223372036854775808:
        var v27 string = "hello"
        _ = v27
    case -1:
        fmt.Println(v4, v4)
    case 996:
        var v28 string = "hello"
        _ = v28
    default:
        v3 = (v3 >> ((492) & 63))
    }
    if (18446744073709551615 != 42) {
        var v29 uint16 = 690
        for fi30 := 0; fi30 < 15; fi30++ {
            fmt.Println(v29, v29)
            _ = fi30
        }
        if (2147483647 == v3) {
            if (0 <= 127) {
                v29 = v29
                for fi31 := 0; fi31 < 14; fi31++ {
                    if false {
                        v32 := GenS2{F0: 32767}
                        _ = v32
                    }
                    _ = fi31
                }
            }
        }
        _ = v29
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 = v5
    if (609 > 2) {
        fmt.Println(v5, v4, v1)
        fmt.Println(v3, v1)
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            if (3.14159 >= (1e10 - float64(float32(0.0)))) {
                var v33 float64 = 1e10
                fmt.Println(v33, v33, v5)
                _ = v33
            } else {
                var v34 uint64 = 18446744073709551615
                v35 := GenS5{F0: 0}
                _ = v34
                _ = v35
            }
            fmt.Println(v3, v5)
            v1 = 1
        }
    }
    _ = len(v2)
    var v36 float64 = ((3.14159 * 1e10) - ((0.1 / 1.5) + 0.0))
    for fi37 := 0; fi37 < 11; fi37++ {
        var v38 uint = 2
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        v40 := GenS7{F0: (fi39 ^ 128), F1: -1}
        _ = fi39
        _ = v40
    }
    v41 := GenS3{F0: 668, F1: ("\n\t" + "foo bar")}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v41
    _ = v5
}
