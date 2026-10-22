package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }
type GenS2 struct { F0 string; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int64; F2 int64 }
type GenS4 struct { F0 uint32; F1 int; F2 int64 }

func main() {
    v1 := []string{("\n\t" + "\n\t"), "hello", ("" + "")}
    var v2 uint32 = 1
    v3 := make(map[uint32]string)
    var v4 string = "a"
    v5 := []int{9223372036854775807, 0}
    fmt.Println(v2, v2)
    var v6 uint32 = (^(v2))
    if true {
        var v7 uint16 = 0
        _ = v7
    } else {
        var v8 uint8 = 42
        switch (v4 + "foo bar") {
        case "foo bar":
            var v9 bool = true
            _ = v9
        case "\n\t":
            for fi10 := 0; fi10 < 15; fi10++ {
                if (0 == 128) {
                    if (!true) {
                        v6 += (v6 * v2)
                        v2 -= (v2 << ((2) & 63))
                        v8 -= 0
                    } else {
                        fmt.Println(fi10, v4)
                        if ((false && (!false)) || (v2 < v6)) {
                            switch (v4 + "\n\t") {
                            case "hello":
                                var v11 string = v4
                                _ = v11
                            default:
                                v12 := GenS1{F0: 32767, F1: 2147483647, F2: 65535}
                                _ = v12
                            }
                            _ = len(v3)
                        }
                        var v13 uint8 = v8
                        _ = v13
                    }
                    v2 = v2
                    var v14 int16 = 32767
                    _ = v14
                }
                v1 = append(v1, v4)
                _ = fi10
            }
        case "hello":
            v4 = v4
        }
        _ = v8
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        v6 = v6
        if false {
            var v16 bool = false
            _ = v16
        } else {
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 uint = 2
                _ = fi17
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 2; fi19++ {
            v2 -= (v6 / ((v6) | 1))
            _ = fi19
        }
        _ = fi15
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        fmt.Println(v6)
        _ = fi20
    }
    v6 = (^(v6))
    v21 := GenS2{F0: v4, F1: true, F2: 590}
    v22 := v3[1]
    v3[148] = ("hello" + ((((v22 + v22) + v4) + "foo bar") + v22))
    for fi23 := 0; fi23 < 16; fi23++ {
        switch "hello" {
        case "\n\t":
            if false {
                for fi24 := 0; fi24 < 4; fi24++ {
                    fmt.Println(v2, fi23)
                    _ = fi24
                }
            } else {
                var v25 string = v4
                v26 := GenS1{F0: 42, F1: 2147483647, F2: 0}
                _ = v25
                _ = v26
            }
        case "a":
            var v27 int16 = 0
            _ = v27
        default:
            if (!false) {
                if ((fi23 << ((492) & 63)) > fi23) {
                    var v28 float64 = 0.1
                    _ = v28
                }
            } else {
                if false {
                    v2 -= v6
                    delete(v3, v2)
                }
            }
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi23
    }
    v6 += 65535
    v22 = v4
    for fi29 := 0; fi29 < 14; fi29++ {
        if (!false) {
            var v30 uint8 = 128
            _ = len(v1)
            _ = v30
        } else {
            for fi31 := 0; fi31 < 10; fi31++ {
                switch fi29 {
                case -512:
                    var v32 int16 = 599
                    _ = v32
                case 1:
                    fmt.Println(v2, v4)
                case -1:
                    v33 := GenS2{F0: "a", F1: true, F2: v6}
                    var v34 int64 = -9223372036854775808
                    if (32 >= -54) {
                        var v35 uint = 150
                        _ = v35
                    } else {
                        if true {
                            v36 := GenS3{F0: 46, F1: (-(v34)), F2: v34}
                            v22 = v4
                            _ = v36
                        } else {
                            v22 = "a"
                            v22 = ((v4 + "") + ("\n\t" + v22))
                        }
                    }
                    _ = v33
                    _ = v34
                default:
                    var v37 uint16 = 697
                    _ = v37
                }
                fmt.Println(v22)
                var v38 float32 = 3.14159
                _ = fi31
                _ = v38
            }
            var v39 int8 = -116
            _ = v39
        }
        var v40 int64 = -884
        var v41 float64 = 1e10
        _ = fi29
        _ = v40
        _ = v41
    }
    var v42 uint8 = 201
    v43 := GenS3{F0: (^(v42)), F1: -9223372036854775808, F2: -842}
    for fi44 := 0; fi44 < 18; fi44++ {
        for fi45 := 0; fi45 < 18; fi45++ {
            var v46 int8 = 0
            if true {
                switch (((v4 + v22) + "") + v4) {
                case "\n\t":
                    fmt.Println(v6)
                    var v47 uint8 = uint8(v2)
                    _ = v47
                case "hello":
                    fmt.Println(v46, v4, v46)
                default:
                    v48 := GenS3{F0: 112, F1: -590, F2: 1}
                    _ = v48
                }
                fmt.Println(fi45, v22)
                v49 := v3[(v2 & 2)]
                if (fi45 >= fi45) {
                    fmt.Println(v4, fi45)
                    switch fi45 {
                    case -9223372036854775808:
                        v4 = (v22 + "")
                    case 70:
                        v4 = v22
                    }
                    v50 := GenS2{F0: v4, F1: true, F2: (v6 << ((453) & 63))}
                    _ = v50
                }
                v51 := GenS4{F0: v2, F1: fi45, F2: 789}
                var v52 bool = true
                switch 4294967295 {
                case 4294967295:
                    var v53 int64 = -1
                    _ = v53
                case 42:
                    v54 := GenS3{F0: v42, F1: -9223372036854775808, F2: 1}
                    _ = v54
                }
                v6 = (^(v6))
                var v55 uint64 = 511
                var v56 int = -1
                _ = v49
                _ = v51
                _ = v52
                _ = v55
                _ = v56
            }
            _ = fi45
            _ = v46
        }
        _ = fi44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
