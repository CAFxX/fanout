package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 bool; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int32; F2 string }
func genFn0() string {
    var v4 int32 = -612
    v4 = v4
    _ = v4
    return ("" + "a")
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = -2.25
    v3 := []int{32768, 2147483647, 2147483647}
    var v5 bool = ((496 < 773) || true)
    v6 := GenS1{F0: ("hello" + (v1 + v1)), F1: 42}
    v1 = ((v1 + v1) + v1)
    var v7 string = (("foo bar" + "a") + v1)
    _ = len(v3)
    fmt.Println(v7)
    v1 = v7
    v5 = (65535 > 605)
    if v5 {
        v1 = (v7 + v1)
        if v5 {
            if v5 {
                for fi8 := 0; fi8 < 4; fi8++ {
                    for fi9 := 0; fi9 < 19; fi9++ {
                        switch v7 {
                        case "hello":
                            if v5 {
                                if v5 {
                                    fmt.Println(v1, v7)
                                } else {
                                    fmt.Println(v1, fi9, fi9)
                                    _ = len(v3)
                                }
                                v3 = append(v3, 127)
                                if v5 {
                                    v7 = (v7 + ("foo bar" + (v1 + ((v7 + "a") + v7))))
                                    var v10 uint8 = 0
                                    _ = v10
                                } else {
                                    v1 = (((("\n\t" + "a") + v1) + "\n\t") + "")
                                    v11 := GenS1{F0: v1, F1: 1}
                                    v7 = v1
                                    _ = v11
                                }
                            }
                        case "foo bar":
                            v12 := GenS2{F0: ((fi8 * fi9) != -141), F1: 2, F2: -68}
                            _ = v12
                        default:
                            if v5 {
                                if v5 {
                                    switch fi9 {
                                    case -381:
                                        var v13 uint = 42
                                        _ = v13
                                    default:
                                        v5 = (255 != 32767)
                                    }
                                    if ((v5 && v5) || ((-2.25 + -2.25) > 0.0)) {
                                        v14 := GenS2{F0: v5, F1: 210, F2: 9223372036854775807}
                                        v15 := GenS1{F0: (v1 + ("a" + v1)), F1: 0}
                                        _ = v14
                                        _ = v15
                                    }
                                    if true {
                                        switch 127 {
                                        case 42:
                                            v1 = ("a" + (v1 + v7))
                                        case 1:
                                            switch v1 {
                                            case "\n\t":
                                                v16 := GenS1{F0: v7, F1: 33}
                                                _ = v16
                                            case "foo bar":
                                                if (2 >= 4294967295) {
                                                    var v17 int = fi8
                                                    _ = v17
                                                } else {
                                                    var v18 int = fi9
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v7 = "foo bar"
                                                    _ = v18
                                                }
                                            case "a":
                                                _ = len(v3)
                                            }
                                        }
                                    } else {
                                        fmt.Println(v7, v2, fi8)
                                        v5 = (!v5)
                                        v1 = ""
                                    }
                                } else {
                                    v1 = "\n\t"
                                }
                                v5 = v5
                            }
                        }
                        switch (-(fi8)) {
                        case 2:
                            var v19 uint8 = 128
                            _ = v19
                        case 4294967295:
                            var v20 bool = (fi9 != fi8)
                            _ = v20
                        case -9223372036854775808:
                            v7 = v7
                        default:
                            var v21 int8 = 42
                            _ = v21
                        }
                        if len(v3) > 0 {
                            v3[0] = 730
                        }
                        _ = fi9
                    }
                    _ = fi8
                }
                if (42 < -2147483648) {
                    var v22 uint = 65535
                    switch v1 {
                    case "a":
                        var v23 uint64 = 18446744073709551615
                        _ = v23
                    case "":
                        if len(v3) > 0 {
                            v3[0] = 494
                        }
                    case "\n\t":
                        v24 := GenS3{F0: 2}
                        _ = v24
                    }
                    v7 = v7
                    _ = v22
                }
                switch 42 {
                case 2147483647:
                    switch -9223372036854775808 {
                    case -12:
                        for fi25 := 0; fi25 < 2; fi25++ {
                            v26 := GenS4{F0: 1000, F1: -2147483648, F2: "a"}
                            if (false && v5) {
                                if (0 > 255) {
                                    var v27 int64 = 42
                                    v2 += (v2 - v2)
                                    _ = v27
                                }
                                v2 = v2
                                switch (fi25 - 9223372036854775807) {
                                case -976:
                                    v3 = append(v3, (fi25 & -280))
                                }
                            }
                            _ = fi25
                            _ = v26
                        }
                    case -35:
                        if len(v3) > 0 {
                            v3[0] = 4294967295
                        }
                    }
                case 191:
                    var v28 string = "foo bar"
                    _ = v28
                case -129:
                    v3 = append(v3, -9223372036854775808)
                default:
                    v7 = ""
                }
            }
        } else {
            v29 := GenS4{F0: 32767, F1: 0, F2: "foo bar"}
            var v30 float32 = v2
            _ = v29
            _ = v30
        }
    } else {
        v5 = false
        if len(v3) > 0 {
            v3[0] = -9223372036854775808
        }
    }
    var v31 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v5
    _ = v6
    _ = v7
}
