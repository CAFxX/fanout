package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int8; F1 uint8; F2 bool }
type GenS5 struct { F0 string; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 int; F1 string; F2 uint8 }

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]bool)
    var v3 uint = 42
    var v4 string = "a"
    v2[(("\n\t" + (v4 + (v4 + v4))) + "\n\t")] = true
    if (!(-119 > -1)) {
        v4 = v4
        for fi5 := 0; fi5 < 6; fi5++ {
            if false {
                v3 -= (v3 / (((v3 ^ (v3 | v3))) | 1))
            }
            _ = fi5
        }
        var v6 int = 127
        _ = v6
    } else {
        if true {
            var v7 string = v4
            v8 := []int{}
            _ = v7
            _ = v8
        } else {
            var v9 uint64 = 1
            v10 := GenS1{F0: v3}
            if true {
                v1 += v1
                v11 := GenS2{F0: 0, F1: -138}
                _ = v11
            } else {
                if true {
                    if true {
                        v1 = (v1 | 42)
                        for fi12 := 0; fi12 < 16; fi12++ {
                            v13 := GenS3{F0: 1, F1: 1}
                            _ = fi12
                            _ = v13
                        }
                        if (-119 < -128) {
                            v14 := GenS4{F0: -128, F1: 4, F2: false}
                            _ = v14
                        }
                    } else {
                        v1 = (v1 + 248)
                    }
                    var v15 float64 = (-2.25 / 3.14159)
                    v9 = 2
                    _ = v15
                }
                v16 := []int{}
                switch (v4 + ("" + v4)) {
                case "\n\t":
                    for fi17 := 0; fi17 < 10; fi17++ {
                        if (!false) {
                            fmt.Println(v3, fi17)
                            if (42 > 65535) {
                                switch (fi17 / (((fi17 >> (((v3 - v3)) & 63))) | 1)) {
                                case 514:
                                    fmt.Println(v1, v9)
                                default:
                                    switch fi17 {
                                    case 128:
                                        v18 := v2[v4]
                                        _ = v18
                                    case 42:
                                        v19 := GenS5{F0: v4, F1: (v9 * (v9 / ((uint64(v1)) | 1))), F2: 0}
                                        _ = v19
                                    default:
                                        if false {
                                            var v20 float64 = (1e10 / ((0.1 / (-2.25 / 3.14159)) - ((-2.25 / 3.14159) - 0.1)))
                                            _ = v20
                                        }
                                    }
                                }
                                if (156 == 1) {
                                    switch (-(fi17)) {
                                    case 42:
                                        fmt.Println(fi17, v4, v3)
                                    }
                                } else {
                                    var v21 uint = (v3 + uint(v9))
                                    var v22 int16 = 0
                                    var v23 float64 = 1e10
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                                if true {
                                    var v24 string = "foo bar"
                                    for fi25 := 0; fi25 < 7; fi25++ {
                                        if true {
                                            var v26 int64 = -9223372036854775808
                                            _ = v26
                                        } else {
                                            if (v24 == (v24 + ("" + ((v4 + v4) + v4)))) {
                                                delete(v2, v4)
                                            }
                                            v9 = 18446744073709551615
                                        }
                                        fmt.Println(v4, v3)
                                        v16 = append(v16, fi25)
                                        _ = fi25
                                    }
                                    _ = v24
                                } else {
                                    switch (v4 + v4) {
                                    case "hello":
                                        for fi27 := 0; fi27 < 18; fi27++ {
                                            v28 := GenS6{F0: (v3 + (v3 + v3))}
                                            _ = fi27
                                            _ = v28
                                        }
                                    case "a":
                                        v3 = (v3 | (v3 % ((uint(v9)) | 1)))
                                    case "":
                                        if len(v16) > 0 {
                                            v16[0] = fi17
                                        }
                                    }
                                    for fi29 := 0; fi29 < 13; fi29++ {
                                        _ = len(v16)
                                        switch (v4 + v4) {
                                        case "a":
                                            var v30 bool = false
                                            _ = v30
                                        }
                                        _ = fi29
                                    }
                                }
                            } else {
                                var v31 float32 = 3.14159
                                _ = v31
                            }
                        } else {
                            v32 := GenS1{F0: v3}
                            v33 := GenS4{F0: 1, F1: v1, F2: (v1 != v1)}
                            if false {
                                v34 := v2[v4]
                                delete(v2, v4)
                                var v35 int16 = 42
                                _ = v34
                                _ = v35
                            } else {
                                fmt.Println(v1, v3, v4)
                                var v36 int8 = 1
                                switch v4 {
                                case "foo bar":
                                    if false {
                                        var v37 uint = v3
                                        fmt.Println(v4)
                                        _ = v37
                                    } else {
                                        v38 := GenS7{F0: 249, F1: -2147483648}
                                        fmt.Println(v3, v1)
                                        _ = v38
                                    }
                                case "a":
                                    v9 += v9
                                }
                                _ = v36
                            }
                            _ = v32
                            _ = v33
                        }
                        if false {
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            var v39 uint = v3
                            _ = v39
                        } else {
                            var v40 int32 = 2147483647
                            v4 = "a"
                            _ = v40
                        }
                        _ = fi17
                    }
                case "foo bar":
                    if len(v16) > 0 {
                        v16[0] = 421
                    }
                default:
                    for fi41 := 0; fi41 < 12; fi41++ {
                        if false {
                            switch "hello" {
                            case "foo bar":
                                v3 = (v3 >> ((991) & 63))
                            }
                            v4 = ((v4 + "\n\t") + v4)
                            fmt.Println(v3)
                        }
                        v42 := GenS8{F0: (fi41 >> (((^(v3))) & 63)), F1: "a", F2: v1}
                        _ = fi41
                        _ = v42
                    }
                }
                _ = v16
            }
            _ = v10
            _ = v9
        }
        var v43 int8 = -1
        _ = v43
    }
    v44 := GenS7{F0: 1, F1: 1}
    switch v4 {
    case "":
        for fi45 := 0; fi45 < 6; fi45++ {
            v46 := GenS7{F0: -760, F1: 0}
            var v47 int = (fi45 & 1)
            _ = fi45
            _ = v46
            _ = v47
        }
    case "a":
        fmt.Println(v3, v1)
    }
    for fi48 := 0; fi48 < 2; fi48++ {
        var v49 int32 = 1
        v50 := GenS3{F0: (v1 * v1), F1: 42}
        var v51 int8 = 100
        _ = fi48
        _ = v49
        _ = v50
        _ = v51
    }
    var v52 uint64 = 42
    var v53 string = v4
    _ = len(v2)
    v3 += 722
    switch v53 {
    case "hello":
        v53 = (v53 + (v53 + ((("\n\t" + (((("foo bar" + (((v4 + "hello") + "foo bar") + v4)) + (v53 + ("\n\t" + ("foo bar" + "\n\t")))) + "a") + v4)) + v4) + "")))
    case "\n\t":
        var v54 int64 = -1
        _ = v54
    default:
        v1 -= 146
    }
    delete(v2, "hello")
    v4 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v52
    _ = v53
}
