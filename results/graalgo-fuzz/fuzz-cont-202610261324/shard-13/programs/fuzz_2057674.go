package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 int; F2 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 int8; F2 uint8 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 uint; F1 string }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 int64 = -839
    v4 := []int{134}
    if (true || false) {
        v3 = v3
    } else {
        var v5 float64 = 1.5
        v5 = (1e10 + (0.0 * float64((1e10 / (1e10 - 1.5)))))
        fmt.Println(v3, v3)
        _ = v5
    }
    switch ("" + ("hello" + ("hello" + ""))) {
    case "foo bar":
        var v6 uint8 = 42
        _ = v6
    case "\n\t":
        v2[("\n\t" + "")] = (("a" + "hello") + "foo bar")
    case "":
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 uint64 = 42
            v8 -= v8
            if true {
                if (!((true || (1 != 0)) || (127 > -128))) {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        var v10 float64 = 3.14159
                        v8 -= v8
                        v11 := GenS1{F0: 1, F1: 4294967295, F2: 2147483648}
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                } else {
                    fmt.Println(fi7, v8, v3)
                    for fi12 := 0; fi12 < 6; fi12++ {
                        if false {
                            v13 := GenS2{F0: 32767, F1: (("a" + "\n\t") + "\n\t")}
                            v14 := GenS3{F0: 1}
                            _ = v13
                            _ = v14
                        } else {
                            if false {
                                v15 := GenS1{F0: 710, F1: 1, F2: 65535}
                                _ = v15
                            } else {
                                _ = len(v1)
                                v3 = v3
                                v16 := GenS3{F0: 2147483647}
                                _ = v16
                            }
                        }
                        _ = fi12
                    }
                }
            }
            _ = fi7
            _ = v8
        }
    }
    v17 := GenS4{F0: 0, F1: -380, F2: 32768}
    v1["\n\t"] = 44
    fmt.Println(v3)
    for fi18 := 0; fi18 < 1; fi18++ {
        fmt.Println(fi18, fi18)
        var v19 int64 = (v3 * (v3 - v3))
        _ = fi18
        _ = v19
    }
    _ = len(v2)
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 string = "a"
        v22 := GenS5{F0: 0}
        if (!false) {
            if false {
                v4 = append(v4, (^(fi20)))
                if true {
                    switch (fi20 * fi20) {
                    case 606:
                        delete(v2, "hello")
                    default:
                        switch ("a" + ((v21 + v21) + "foo bar")) {
                        case "hello":
                            if len(v4) > 0 {
                                v4[0] = (fi20 & fi20)
                            }
                        case "a":
                            for fi23 := 0; fi23 < 14; fi23++ {
                                v3 += (v3 * (v3 / ((0) | 1)))
                                _ = fi23
                            }
                        case "":
                            switch v21 {
                            case "foo bar":
                                _ = len(v4)
                            default:
                                v21 = v21
                            }
                        default:
                            var v24 uint32 = 2
                            _ = v24
                        }
                    }
                }
                fmt.Println(fi20)
            }
        } else {
            v3 = 0
            if false {
                if (true && false) {
                    for fi25 := 0; fi25 < 2; fi25++ {
                        var v26 uint8 = 42
                        if false {
                            if len(v4) > 0 {
                                v4[0] = -2147483648
                            }
                        } else {
                            v4 = append(v4, fi20)
                            var v27 float32 = float32(0.0)
                            if true {
                                var v28 int = 9223372036854775807
                                switch fi25 {
                                case 4294967295:
                                    v29 := GenS6{F0: 42, F1: 34, F2: (v26 + v26)}
                                    _ = v29
                                case 249:
                                    fmt.Println(v21, v28)
                                }
                                _ = v28
                            } else {
                                switch "foo bar" {
                                case "foo bar":
                                    v21 = "foo bar"
                                }
                            }
                            fmt.Println(v21)
                            v3 -= 9223372036854775807
                            _ = v27
                        }
                        _ = fi25
                        _ = v26
                    }
                    var v30 bool = (!(!true))
                    _ = v30
                } else {
                    fmt.Println(v3, fi20, v21)
                    fmt.Println(v3, fi20, v21)
                    switch 1 {
                    case -871:
                        if (true || true) {
                            if len(v4) > 0 {
                                v4[0] = 128
                            }
                            switch (v21 + v21) {
                            case "hello":
                                v31 := GenS4{F0: 42, F1: (-(fi20)), F2: (fi20 - (fi20 << ((264) & 63)))}
                                _ = v31
                            default:
                                switch fi20 {
                                case 151:
                                    _ = len(v2)
                                case 32767:
                                    v32 := GenS7{F0: -2147483648, F1: 511}
                                    _ = v32
                                default:
                                    v33 := GenS8{F0: 18446744073709551615, F1: 42}
                                    _ = v33
                                }
                            }
                        }
                    case 844:
                        if true {
                            v3 += v3
                        } else {
                            fmt.Println(v21)
                            v34 := GenS8{F0: 0, F1: 0}
                            _ = v34
                        }
                    case -799:
                        fmt.Println(v21)
                    default:
                        _ = len(v2)
                    }
                }
            }
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    var v35 uint64 = 2
    v36 := GenS9{F0: 1, F1: "hello"}
    var v37 int = 2147483647
    v4 = append(v4, (v37 - -338))
    v35 = (v35 ^ (v35 - 18446744073709551615))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
