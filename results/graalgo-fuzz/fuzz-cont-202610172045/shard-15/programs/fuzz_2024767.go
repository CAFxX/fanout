package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32 }

func main() {
    var v1 float64 = 1.5
    var v2 uint16 = 32767
    var v3 int32 = 475
    v2 = (v2 ^ (v2 + v2))
    v4 := []int{2147483647, 128}
    if (64 < -44) {
        v5 := GenS1{F0: (!true), F1: 79}
        _ = v5
    } else {
        if ((-2.25 / v1) >= 0.0) {
            fmt.Println(v2)
            _ = len(v4)
            switch "\n\t" {
            case "a":
                v6 := GenS2{F0: 127, F1: ((("a" + ("\n\t" + ("\n\t" + ("hello" + "foo bar")))) + ("foo bar" + "foo bar")) + "foo bar")}
                _ = v6
            case "hello":
                var v7 float32 = (float32(v1) * float32(v1))
                _ = v7
            default:
                v2 = (^(v2))
            }
        } else {
            v1 += (v1 - v1)
        }
    }
    v1 -= 0.1
    v2 = (^(v2))
    v3 = v3
    var v8 uint32 = 42
    switch "\n\t" {
    case "foo bar":
        if (!(((1.5 - (1e10 / 1.0)) * 3.14159) == -2.25)) {
            var v9 int16 = 0
            var v10 uint16 = (^(v2))
            _ = v10
            _ = v9
        } else {
            if len(v4) > 0 {
                v4[0] = 16
            }
            switch "foo bar" {
            case "foo bar":
                switch 128 {
                case 475:
                    var v11 int32 = 2147483647
                    _ = v11
                }
            case "":
                if true {
                    v12 := GenS3{F0: 128, F1: -1}
                    var v13 float64 = v1
                    var v14 uint = 62
                    _ = v12
                    _ = v13
                    _ = v14
                }
            }
        }
    case "a":
        for fi15 := 0; fi15 < 2; fi15++ {
            fmt.Println(v8, fi15, v3)
            if false {
                var v16 uint16 = (v2 ^ (v2 - v2))
                for fi17 := 0; fi17 < 3; fi17++ {
                    if true {
                        var v18 int16 = -1
                        _ = v18
                    } else {
                        var v19 int32 = (v3 ^ v3)
                        if (!false) {
                            switch fi17 {
                            case 1:
                                var v20 float32 = (((1e10 + 1e10) / (0.0 - 1e10)) - (3.14159 / (1.5 + 1.5)))
                                _ = v20
                            case -701:
                                v2 = 1
                            }
                            var v21 uint8 = 0
                            _ = v21
                        }
                        v16 += 2
                        _ = v19
                    }
                    _ = fi17
                }
                v22 := GenS4{F0: 1, F1: 18446744073709551615}
                _ = v16
                _ = v22
            } else {
                if (!(uint64(v2) >= 0)) {
                    if (!false) {
                        var v23 bool = true
                        if len(v4) > 0 {
                            v4[0] = (^(fi15))
                        }
                        for fi24 := 0; fi24 < 4; fi24++ {
                            v2 = (v2 << ((1) & 63))
                            v25 := GenS5{F0: ""}
                            switch fi15 {
                            case 1:
                                v8 = (^(v8))
                            case -2147483648:
                                if true {
                                    fmt.Println(fi15)
                                } else {
                                    v26 := GenS1{F0: v23, F1: 211}
                                    if true {
                                        if false {
                                            fmt.Println(fi24, v23)
                                            var v27 int32 = (v3 / ((int32(v2)) | 1))
                                            v28 := GenS6{F0: (^(v3))}
                                            _ = v27
                                            _ = v28
                                        }
                                        var v29 int64 = -1
                                        if (!true) {
                                            v1 = (v1 * 0.1)
                                        } else {
                                            v30 := GenS5{F0: "foo bar"}
                                            if false {
                                                fmt.Println(v1, v23)
                                                v29 += (-(v29))
                                                fmt.Println(v23)
                                            }
                                            var v31 float64 = (v1 - v1)
                                            _ = v30
                                            _ = v31
                                        }
                                        _ = v29
                                    } else {
                                        var v32 bool = false
                                        switch fi15 {
                                        case 4294967295:
                                            var v33 int8 = 42
                                            _ = v33
                                        case 275:
                                            v23 = v32
                                        case 761:
                                            var v34 float32 = (-2.25 * -2.25)
                                            _ = v34
                                        }
                                        v23 = (int8(v3) <= 42)
                                        v32 = v32
                                        v8 = v8
                                        _ = v32
                                    }
                                    if (v23 || v23) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    _ = v26
                                }
                            case 42:
                                if v23 {
                                    v23 = (0 < v2)
                                }
                            }
                            _ = fi24
                            _ = v25
                        }
                        _ = v23
                    } else {
                        v2 = (v2 << ((4294967295) & 63))
                    }
                    v35 := GenS5{F0: "a"}
                    v36 := GenS5{F0: "hello"}
                    _ = v35
                    _ = v36
                } else {
                    v37 := GenS6{F0: v3}
                    var v38 int32 = int32(v2)
                    _ = v37
                    _ = v38
                }
            }
            _ = fi15
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
