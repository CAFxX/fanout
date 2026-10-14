package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int64; F2 uint64 }
func genFn0() uint64 {
    v7 := GenS1{F0: 966, F1: 1}
    v8 := []int{-1}
    var v9 uint32 = 2147483647
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1() uint64 {
    v10 := GenS1{F0: 0, F1: 42}
    switch 42 {
    case 4294967295:
        if false {
            v11 := GenS1{F0: 268, F1: 65535}
            if true {
                var v12 int = 365
                var v13 string = "foo bar"
                _ = v12
                _ = v13
            } else {
                var v14 uint = 42
                v15 := GenS1{F0: 32767, F1: 309}
                _ = v14
                _ = v15
            }
            switch "\n\t" {
            case "a":
                if (!true) {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v17 := GenS2{F0: 2, F1: -296}
                        var v18 int8 = -15
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                } else {
                    var v19 float32 = ((0.1 / 1.0) / float32((1e10 / 3.14159)))
                    _ = v19
                }
            case "\n\t":
                var v20 int32 = 0
                _ = v20
            case "foo bar":
                var v21 uint16 = 1
                _ = v21
            }
            _ = v11
        }
    case -305:
        var v22 int = -256
        _ = v22
    default:
        v23 := GenS3{F0: 821, F1: -76}
        _ = v23
    }
    _ = v10
    return 102
}

func main() {
    v1 := []int{}
    var v2 uint = 375
    var v3 bool = (-32768 < -32768)
    v4 := make(map[uint32]bool)
    var v5 float64 = float64(float32(0.1))
    v6 := []uint32{2147483647}
    switch "a" {
    case "foo bar":
        for fi24 := 0; fi24 < 5; fi24++ {
            v2 -= (^(v2))
            _ = fi24
        }
    case "\n\t":
        if v3 {
            var v25 uint64 = 975
            for fi26 := 0; fi26 < 16; fi26++ {
                for fi27 := 0; fi27 < 13; fi27++ {
                    var v28 uint8 = 243
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
            v25 -= genFn0()
            _ = v25
        }
    case "":
        v29 := GenS3{F0: -1, F1: 100}
        _ = v29
    default:
        v5 += ((v5 / (0.1 + ((v5 - 3.14159) - (0.1 - v5)))) - (float64((0.0 * 3.14159)) / v5))
    }
    v30 := GenS4{F0: 1}
    v31 := GenS5{F0: v2, F1: -833, F2: 0}
    if (v2 == v2) {
        switch ("\n\t" + ((("foo bar" + "\n\t") + "") + "\n\t")) {
        case "\n\t":
            if (v2 != 0) {
                var v32 uint8 = 197
                _ = v32
            } else {
                for fi33 := 0; fi33 < 12; fi33++ {
                    if (97 >= 100) {
                        var v34 uint16 = 32767
                        switch "a" {
                        case "hello":
                            if (!v3) {
                                v35 := GenS4{F0: 130}
                                _ = v35
                            } else {
                                if (!v3) {
                                    var v36 string = ""
                                    v36 = (((v36 + "a") + v36) + v36)
                                    fmt.Println(v34, v34, v3)
                                    _ = v36
                                }
                            }
                        }
                        var v37 uint8 = 17
                        _ = v34
                        _ = v37
                    }
                    if false {
                        v4[60] = true
                    } else {
                        v38 := GenS3{F0: 1, F1: -1}
                        var v39 uint8 = uint8(v2)
                        _ = v38
                        _ = v39
                    }
                    if len(v6) > 0 {
                        v6[0] = uint32(v2)
                    }
                    _ = fi33
                }
                switch ("hello" + ("" + "foo bar")) {
                case "foo bar":
                    if v3 {
                        for fi40 := 0; fi40 < 10; fi40++ {
                            v2 = 65535
                            _ = fi40
                        }
                    } else {
                        var v41 uint8 = 217
                        _ = v41
                    }
                }
                v42 := GenS5{F0: 836, F1: 0, F2: 4294967295}
                _ = v42
            }
        case "":
            var v43 float32 = 3.14159
            _ = v43
        case "hello":
            for fi44 := 0; fi44 < 14; fi44++ {
                v45 := GenS1{F0: 2, F1: 0}
                _ = fi44
                _ = v45
            }
        default:
            v46 := GenS2{F0: v2, F1: 42}
            _ = v46
        }
        var v47 uint16 = 2
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v47
    }
    fmt.Println(v2, v3)
    v2 = (v2 & (v2 ^ (v2 * 42)))
    if v3 {
        if v3 {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            var v48 float32 = 1e10
            _ = v48
        }
    }
    fmt.Println(v5)
    v49 := v4[833]
    if v49 {
        v4[0] = (v5 == v5)
    }
    var v50 int64 = 689
    v5 = 1e10
    v51 := GenS2{F0: 767, F1: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
