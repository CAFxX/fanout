package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS4 struct { F0 int32; F1 int32; F2 int }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint32; F1 uint16 }
type GenS10 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (0.1 - 1.5)
    v3 := []uint32{2147483648, 2, 1, 4294967295}
    v4 := make(map[uint32]uint32)
    var v5 int8 = 14
    var v6 uint16 = 42
    switch 128 {
    case 9223372036854775807:
        if (50 == 910) {
            var v7 int16 = 42
            v5 = v5
            v8 := GenS1{F0: 32767, F1: (v7 >> ((585) & 63))}
            _ = v7
            _ = v8
        }
    default:
        v9 := v4[612]
        _ = v9
    }
    v10 := GenS2{F0: ("" + (("hello" + "") + "a")), F1: (("hello" + "foo bar") + "a")}
    v11 := GenS3{F0: 84, F1: (0 >= 255), F2: 42}
    var v12 int32 = 678
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            v12 = 145
            fmt.Println(v12, fi14)
            var v15 uint8 = 70
            var v16 int16 = 32767
            v6 += 0
            _ = fi14
            _ = v15
            _ = v16
        }
        v12 += (v12 - (v12 / (((v12 / ((v12) | 1))) | 1)))
        _ = fi13
    }
    switch "\n\t" {
    case "a":
        for fi17 := 0; fi17 < 14; fi17++ {
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 uint64 = 576
                var v20 float64 = 1.5
                _ = fi18
                _ = v19
                _ = v20
            }
            switch (("a" + ((((("a" + "\n\t") + "\n\t") + "\n\t") + "foo bar") + "\n\t")) + ("foo bar" + "foo bar")) {
            case "hello":
                v6 = 42
            }
            _ = fi17
        }
    }
    if false {
        v2 += float32(0.1)
        for fi21 := 0; fi21 < 16; fi21++ {
            v22 := GenS1{F0: -32768, F1: 32767}
            var v23 float32 = v2
            fmt.Println(v5)
            _ = fi21
            _ = v22
            _ = v23
        }
    } else {
        for fi24 := 0; fi24 < 20; fi24++ {
            v12 = (v12 % ((-1) | 1))
            fmt.Println(v2)
            switch ("" + "foo bar") {
            case "":
                v2 -= v2
            default:
                if (!false) {
                    var v25 uint8 = 255
                    if true {
                        switch fi24 {
                        case 646:
                            var v26 int8 = v5
                            _ = v26
                        case 0:
                            for fi27 := 0; fi27 < 11; fi27++ {
                                fmt.Println(v25, v6, fi27)
                                _ = fi27
                            }
                        }
                    } else {
                        _ = len(v1)
                        for fi28 := 0; fi28 < 18; fi28++ {
                            v29 := GenS4{F0: 870, F1: 42, F2: fi24}
                            switch "foo bar" {
                            case "\n\t":
                                switch fi28 {
                                case 560:
                                    v30 := GenS1{F0: 232, F1: 0}
                                    _ = v30
                                case 2147483647:
                                    v6 += v6
                                }
                            case "a":
                                v5 = v5
                            default:
                                if len(v3) > 0 {
                                    v3[0] = 763
                                }
                            }
                            var v31 int16 = 937
                            _ = fi28
                            _ = v29
                            _ = v31
                        }
                        var v32 int32 = v12
                        _ = v32
                    }
                    delete(v1, fi24)
                    _ = v25
                } else {
                    switch fi24 {
                    case 0:
                        var v33 float32 = 0.0
                        _ = v33
                    case 196:
                        if true {
                            for fi34 := 0; fi34 < 4; fi34++ {
                                v35 := GenS5{F0: 2, F1: int8(v6)}
                                v3 = append(v3, 178)
                                _ = fi34
                                _ = v35
                            }
                        } else {
                            var v36 uint64 = 42
                            v36 += (v36 >> ((2) & 63))
                            var v37 int64 = 1
                            _ = v36
                            _ = v37
                        }
                    case 128:
                        fmt.Println(v12, v6)
                    }
                    var v38 uint32 = 284
                    _ = v38
                }
            }
            _ = fi24
        }
        v12 = v12
        var v39 uint64 = 269
        _ = v39
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        if (0.1 < ((v2 / (((1.5 * (1.5 * 1e10)) - 0.0) - -2.25)) + ((v2 / v2) * v2))) {
            if (v12 != v12) {
                for fi41 := 0; fi41 < 3; fi41++ {
                    switch "hello" {
                    case "":
                        if true {
                            v42 := GenS6{F0: (v6 - v6)}
                            _ = v42
                        }
                    }
                    if (200 != 42) {
                        switch (fi41 | -979) {
                        case 4294967295:
                            if false {
                                var v43 int64 = 0
                                var v44 uint16 = (v6 ^ 679)
                                v43 = v43
                                _ = v43
                                _ = v44
                            } else {
                                var v45 uint32 = 929
                                _ = v45
                            }
                        case 1:
                            v6 = (^(v6))
                        }
                        v46 := GenS7{F0: -9223372036854775808, F1: 2}
                        v2 = v2
                        _ = v46
                    }
                    var v47 int8 = v5
                    _ = fi41
                    _ = v47
                }
            }
            var v48 float64 = (float64(float32(float64(v2))) - ((((-2.25 - 0.0) - (3.14159 + 0.0)) * (0.0 - -2.25)) + (((3.14159 + (3.14159 / 0.1)) - (-2.25 * 0.1)) + float64(0.0))))
            var v49 uint8 = 2
            _ = v48
            _ = v49
        } else {
            v50 := GenS2{F0: "foo bar", F1: "\n\t"}
            fmt.Println(fi40, v2, v2)
            _ = v50
        }
        _ = fi40
    }
    if false {
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        delete(v4, 42)
    }
    var v51 int16 = -32768
    if ((v12 ^ (-(v12))) > v12) {
        if false {
            v51 = v51
            v52 := GenS8{F0: "a"}
            _ = v52
        } else {
            v51 = v51
            v51 -= v51
            v53 := GenS9{F0: 65535, F1: (v6 - v6)}
            _ = v53
        }
    }
    if (false && false) {
        v5 -= (v5 & v5)
    }
    fmt.Println(v51)
    v54 := GenS10{F0: (v6 & v6), F1: 42}
    v55 := GenS2{F0: "", F1: "hello"}
    fmt.Println(v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v54
    _ = v55
    _ = v6
}
