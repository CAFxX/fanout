package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint8; F1 string; F2 bool }

func main() {
    v1 := []uint32{316, 2, 2}
    var v2 int = 4294967295
    v3 := make(map[uint32]string)
    var v4 int = v2
    var v5 bool = (float64(1e10) != 0.1)
    var v6 uint = 4294967295
    if v5 {
        switch ("foo bar" + "\n\t") {
        case "":
            switch 128 {
            case -922:
                v7 := GenS1{F0: 0, F1: ""}
                _ = v7
            default:
                var v8 int32 = -670
                _ = v8
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v9 float32 = float32(0.1)
        v2 = -1
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: "a"}
    var v11 uint32 = 891
    v4 -= v2
    v5 = v5
    v6 = (v6 % ((278) | 1))
    if ("foo bar" != (("\n\t" + ("hello" + "hello")) + ("foo bar" + "\n\t"))) {
        v12 := GenS1{F0: -9223372036854775808, F1: "a"}
        _ = len(v3)
        if v5 {
            var v13 float64 = float64(3.14159)
            v14 := GenS1{F0: -795, F1: "foo bar"}
            v6 = v6
            _ = v13
            _ = v14
        }
        _ = v12
    } else {
        var v15 uint = v6
        _ = v15
    }
    if len(v1) > 0 {
        v1[0] = v11
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        if v5 {
            switch "a" {
            case "\n\t":
                var v17 uint32 = 13
                _ = v17
            case "foo bar":
                var v18 int64 = 9223372036854775807
                _ = v18
            case "a":
                for fi19 := 0; fi19 < 17; fi19++ {
                    switch v4 {
                    case 4294967295:
                        fmt.Println(v2, v5)
                    case -470:
                        switch "\n\t" {
                        case "":
                            fmt.Println(v5)
                        case "foo bar":
                            fmt.Println(v4, v11)
                        default:
                            fmt.Println(v6, v6)
                        }
                    case -9223372036854775808:
                        var v20 int64 = 0
                        _ = v20
                    }
                    fmt.Println(v2, v4, v2)
                    _ = fi19
                }
            default:
                v11 = v11
            }
        } else {
            fmt.Println(v5)
            switch "foo bar" {
            case "\n\t":
                v11 += v11
            case "a":
                var v21 bool = ((v11 << ((65535) & 63)) >= 543)
                _ = v21
            case "foo bar":
                v6 -= (^(v6))
            default:
                if (42 > 864) {
                    v2 = 32768
                } else {
                    switch "" {
                    case "\n\t":
                        v4 = (v2 + v2)
                    case "hello":
                        if ((128 < 255) || false) {
                            v22 := GenS2{F0: "", F1: (v4 & v2)}
                            v23 := GenS3{F0: uint8(v11), F1: "\n\t", F2: ((-(fi16)) < v2)}
                            v24 := v3[uint32(v6)]
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            _ = len(v3)
                        }
                    }
                }
            }
        }
        _ = fi16
    }
    switch "" {
    case "":
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 int16 = 273
            _ = fi25
            _ = v26
        }
    }
    if true {
        switch v2 {
        case 32767:
            fmt.Println(v11)
        default:
            for fi27 := 0; fi27 < 12; fi27++ {
                v5 = ((!(v5 && false)) || (v5 || ((float32(0.1) + (1e10 + 1.5)) == 0.1)))
                for fi28 := 0; fi28 < 7; fi28++ {
                    v6 = (v6 | (v6 << ((v6) & 63)))
                    _ = fi28
                }
                _ = fi27
            }
        }
    } else {
        switch v4 {
        case 127:
            fmt.Println(v2, v6, v2)
        case 32767:
            v29 := GenS1{F0: -405, F1: ""}
            _ = v29
        default:
            v2 = -891
        }
        v3[v11] = "foo bar"
        for fi30 := 0; fi30 < 18; fi30++ {
            var v31 int8 = 127
            if v5 {
                v31 = v31
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                if v5 {
                    v5 = (v6 >= v6)
                    var v32 int = (^(v2))
                    var v33 uint32 = (v11 ^ (^(v11)))
                    _ = v32
                    _ = v33
                } else {
                    var v34 bool = (!false)
                    _ = v34
                }
            }
            _ = fi30
            _ = v31
        }
    }
    if (!(v11 != (^(v11)))) {
        var v35 int8 = -1
        if len(v1) > 0 {
            v1[0] = (^(v11))
        }
        if v5 {
            if v5 {
                var v36 uint32 = (v11 << ((18446744073709551615) & 63))
                if false {
                    v6 = v6
                    var v37 float32 = 0.0
                    v38 := GenS1{F0: 0, F1: ("\n\t" + "a")}
                    _ = v37
                    _ = v38
                }
                _ = v36
            }
        } else {
            v3[v11] = ""
        }
        _ = v35
    } else {
        var v39 int16 = -534
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
