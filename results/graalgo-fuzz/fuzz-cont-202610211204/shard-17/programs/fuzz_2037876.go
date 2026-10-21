package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 string; F1 int32; F2 uint16 }
type GenS3 struct { F0 string; F1 uint; F2 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 uint16 = 597
    v2 := make(map[uint32]bool)
    v3 := []uint32{2147483647}
    v4 := make(map[string]int)
    var v5 int32 = 67
    var v6 int32 = v5
    fmt.Println(v6, v5)
    switch 1 {
    case 2147483647:
        var v7 uint8 = 1
        _ = v7
    case 712:
        v8 := GenS1{F0: false, F1: false}
        _ = v8
    case -2147483648:
        delete(v4, "hello")
    default:
        if len(v3) > 0 {
            v3[0] = uint32(v1)
        }
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v5 -= (v5 % (((-(v5))) | 1))
        v10 := GenS1{F0: true, F1: ("hello" <= "foo bar")}
        v11 := GenS2{F0: "\n\t", F1: v6, F2: (^(v1))}
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v1, v5)
    case "a":
        _ = len(v3)
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        if false {
            for fi13 := 0; fi13 < 2; fi13++ {
                v2[0] = false
                _ = fi13
            }
        } else {
            if false {
                for fi14 := 0; fi14 < 19; fi14++ {
                    fmt.Println(fi12, v5)
                    var v15 float64 = 1.5
                    if false {
                        v2[65535] = (42 > 0)
                        if (!(false || false)) {
                            if len(v3) > 0 {
                                v3[0] = 65535
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            if false {
                                fmt.Println(fi12, v5)
                            } else {
                                switch -365 {
                                case 42:
                                    switch fi14 {
                                    case 32767:
                                        v16 := v2[2]
                                        _ = v16
                                    case -1:
                                        v17 := GenS3{F0: "\n\t", F1: 65535, F2: 1}
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v18 uint16 = 1
                                            _ = v18
                                        case "a":
                                            if true {
                                                v15 = v15
                                                v3 = append(v3, 525)
                                            }
                                        }
                                        v15 = (((1e10 / 1e10) - (v15 / 1.0)) / v15)
                                        _ = v17
                                    default:
                                        var v19 uint32 = 1
                                        _ = v19
                                    }
                                }
                                var v20 int = fi14
                                v20 -= int(v1)
                                _ = v20
                            }
                            v21 := GenS2{F0: "\n\t", F1: (v5 / (((v6 & v5)) | 1)), F2: (^(v1))}
                            _ = v21
                        }
                    } else {
                        if false {
                            v22 := v2[42]
                            if false {
                                if (((("foo bar" <= "hello") && (2 >= 2147483648)) && v22) && ((42 == 127) || ((-482 <= 1) || v22))) {
                                    if v22 {
                                        v23 := GenS4{F0: 918}
                                        v22 = v22
                                        _ = v23
                                    }
                                }
                                v6 = (v5 - v5)
                            }
                            switch 538 {
                            case 9223372036854775807:
                                var v24 int = fi14
                                _ = v24
                            case 4294967295:
                                var v25 int64 = -1
                                _ = v25
                            case 2147483647:
                                var v26 uint8 = 9
                                _ = v26
                            default:
                                fmt.Println(v22, v6, fi12)
                            }
                            _ = v22
                        } else {
                            if (true && true) {
                                if len(v3) > 0 {
                                    v3[0] = 42
                                }
                                if (32767 <= -576) {
                                    if (true && false) {
                                        var v27 uint32 = 0
                                        var v28 uint32 = v27
                                        v15 -= v15
                                        _ = v27
                                        _ = v28
                                    }
                                    var v29 int64 = -9223372036854775808
                                    _ = v29
                                }
                            } else {
                                fmt.Println(v5)
                                v6 = v6
                                v5 = 0
                            }
                            fmt.Println(fi12)
                        }
                        if true {
                            var v30 int8 = 1
                            fmt.Println(v5, fi12)
                            var v31 float32 = 1.5
                            _ = v30
                            _ = v31
                        } else {
                            var v32 int16 = 238
                            _ = v32
                        }
                        switch (fi14 * fi14) {
                        case 2147483647:
                            fmt.Println(fi14)
                        default:
                            var v33 float64 = v15
                            _ = v33
                        }
                    }
                    _ = fi14
                    _ = v15
                }
                v5 -= 0
                fmt.Println(v5, v6)
            }
        }
        _ = fi12
    }
    v1 -= (^(v1))
    if (v5 < v6) {
        fmt.Println(v1, v5)
        var v34 uint16 = v1
        _ = v34
    } else {
        if len(v3) > 0 {
            v3[0] = 688
        }
    }
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    v35 := GenS2{F0: "foo bar", F1: (v5 ^ -1), F2: v1}
    v36 := GenS5{F0: 682}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v37 uint = 1
    v38 := GenS3{F0: "hello", F1: v37, F2: -32768}
    var v39 uint8 = 127
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    delete(v4, (("" + "\n\t") + (("a" + "") + "\n\t")))
    switch 32768 {
    case -29:
        fmt.Println(v37)
    case 4294967295:
        if false {
            var v40 uint32 = 42
            fmt.Println(v6, v37)
            _ = v40
        }
    case 0:
        v5 = (v6 - v6)
    default:
        fmt.Println(v1, v6, v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
