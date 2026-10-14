package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int }
type GenS2 struct { F0 uint; F1 int8; F2 string }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint64; F1 uint16; F2 int64 }
func genFn0() uint {
    var v7 int32 = -2147483648
    var v8 uint32 = 4294967295
    for fi9 := 0; fi9 < 19; fi9++ {
        v8 -= 1
        v7 += -419
        _ = fi9
    }
    fmt.Println(v7, v8)
    _ = v7
    _ = v8
    return 255
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -39
    v3 := make(map[string]int)
    v4 := make(map[string]int)
    var v5 uint32 = 2
    var v6 uint = 255
    v5 = 2147483647
    v10 := GenS1{F0: (v6 + v6), F1: -208, F2: 1}
    var v11 uint64 = 57
    var v12 int64 = -607
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            if (42 > 1) {
                var v15 int32 = 2147483647
                switch ("\n\t" + "a") {
                case "foo bar":
                    v15 -= (v15 >> ((777) & 63))
                case "hello":
                    var v16 uint8 = 2
                    _ = v16
                case "":
                    switch "hello" {
                    case "\n\t":
                        v2 += 42
                    case "hello":
                        var v17 uint32 = (v5 + uint32(v6))
                        _ = v17
                    case "a":
                        if (v5 < v5) {
                            v2 = v2
                            var v18 uint8 = 0
                            switch "foo bar" {
                            case "":
                                _ = len(v1)
                            default:
                                v12 += (v12 / ((65) | 1))
                            }
                            _ = v18
                        }
                    default:
                        v19 := GenS2{F0: v6, F1: 0, F2: "hello"}
                        _ = v19
                    }
                }
                switch "" {
                case "":
                    if (!(42 != (v12 ^ v12))) {
                        switch (fi13 % (((fi14 & fi13)) | 1)) {
                        case -1:
                            if ((!(0.1 != 3.14159)) || (true || (0 < (v12 % (((^(v12))) | 1))))) {
                                v20 := GenS1{F0: v6, F1: 42, F2: (-(fi13))}
                                _ = v20
                            } else {
                                var v21 int16 = 42
                                switch (fi13 + fi14) {
                                case -379:
                                    var v22 int16 = (v2 << (((v6 | (^(v6)))) & 63))
                                    _ = v22
                                case 9223372036854775807:
                                    var v23 float32 = 3.14159
                                    _ = v23
                                }
                                var v24 uint8 = 255
                                _ = v21
                                _ = v24
                            }
                        case -2147483648:
                            if false {
                                switch ("hello" + ("a" + "hello")) {
                                case "hello":
                                    if ((-2.25 < 0.0) || (!true)) {
                                        var v25 string = "\n\t"
                                        v26 := []int{127, int(v6)}
                                        v2 = v2
                                        var v27 string = v25
                                        var v28 string = "a"
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    } else {
                                        var v29 float32 = 1.5
                                        var v30 uint16 = 1
                                        fmt.Println(v12, v11)
                                        _ = v29
                                        _ = v30
                                    }
                                    var v31 bool = (100 >= 100)
                                    _ = v31
                                case "":
                                    var v32 uint64 = (v11 % (((^(v11))) | 1))
                                    _ = v32
                                }
                                if true {
                                    v5 = v5
                                } else {
                                    var v33 uint16 = 1
                                    v5 = v5
                                    _ = v33
                                }
                                v34 := GenS1{F0: 18446744073709551615, F1: (-(v12)), F2: (fi14 % ((fi14) | 1))}
                                _ = v34
                            }
                        case 127:
                            var v35 int16 = (v2 | (v2 + -571))
                            var v36 uint32 = v5
                            v11 = v11
                            _ = v35
                            _ = v36
                        default:
                            v15 = int32(v6)
                        }
                        if false {
                            switch (fi14 / ((-2147483648) | 1)) {
                            case 370:
                                v2 = 0
                                if true {
                                    if false {
                                        var v37 uint8 = 42
                                        _ = v37
                                    }
                                    v38 := GenS1{F0: (v6 >> (((v6 % ((genFn0()) | 1))) & 63)), F1: 9223372036854775807, F2: int(v11)}
                                    v39 := GenS1{F0: 465, F1: (v12 + v12), F2: (-(fi14))}
                                    _ = v38
                                    _ = v39
                                } else {
                                    v4["a"] = fi13
                                    if true {
                                        v6 -= v6
                                        v40 := GenS3{F0: uint64(v6), F1: (v6 * v6)}
                                        _ = v40
                                    }
                                }
                                v5 = (v5 >> ((43) & 63))
                            }
                            v41 := GenS4{F0: genFn0()}
                            _ = v41
                        }
                    }
                case "\n\t":
                    v42 := GenS2{F0: (v6 | (v6 / (((v6 / ((genFn0()) | 1))) | 1))), F1: -1, F2: ("\n\t" + "hello")}
                    _ = v42
                default:
                    var v43 uint64 = v11
                    _ = v43
                }
                _ = v15
            }
            v11 -= v11
            v12 -= (-(v12))
            _ = fi14
        }
        var v44 bool = (v6 < v6)
        _ = fi13
        _ = v44
    }
    for fi45 := 0; fi45 < 19; fi45++ {
        for fi46 := 0; fi46 < 12; fi46++ {
            var v47 int32 = -460
            _ = fi46
            _ = v47
        }
        for fi48 := 0; fi48 < 14; fi48++ {
            v2 = (v2 - 1000)
            v49 := []int{}
            _ = fi48
            _ = v49
        }
        for fi50 := 0; fi50 < 6; fi50++ {
            var v51 int64 = (-(v12))
            switch (("a" + "a") + "\n\t") {
            case "\n\t":
                v11 = 18446744073709551615
            }
            _ = fi50
            _ = v51
        }
        _ = fi45
    }
    for fi52 := 0; fi52 < 20; fi52++ {
        v53 := v3["hello"]
        if (v53 >= v53) {
            if true {
                v54 := GenS5{F0: 127, F1: 42}
                switch "" {
                case "":
                    if true {
                        switch v53 {
                        case 941:
                            var v55 uint32 = (v5 % ((4294967295) | 1))
                            _ = v55
                        case 9223372036854775807:
                            for fi56 := 0; fi56 < 13; fi56++ {
                                v5 += v5
                                v53 -= v53
                                v53 -= v53
                                _ = fi56
                            }
                        default:
                            _ = len(v1)
                        }
                        fmt.Println(v53)
                        for fi57 := 0; fi57 < 18; fi57++ {
                            v11 += 0
                            _ = fi57
                        }
                    }
                case "hello":
                    for fi58 := 0; fi58 < 16; fi58++ {
                        v59 := GenS1{F0: 255, F1: (v12 >> ((v6) & 63)), F2: 42}
                        v5 -= v5
                        v60 := GenS4{F0: 249}
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                default:
                    var v61 int32 = 279
                    _ = v61
                }
                _ = v54
            }
            v12 += (^(v12))
        } else {
            var v62 int16 = 42
            v62 = v62
            var v63 uint32 = v5
            _ = v62
            _ = v63
        }
        var v64 float64 = float64((0.1 / 1.0))
        _ = fi52
        _ = v53
        _ = v64
    }
    v2 -= 1000
    v11 = (v11 ^ 2)
    fmt.Println(v6, v6)
    fmt.Println(v11, v6)
    v65 := GenS6{F0: (v11 << ((genFn0()) & 63)), F1: 592, F2: v12}
    fmt.Println(v6, v6)
    v66 := []int{int(v11), -2147483648}
    fmt.Println(v5, v6)
    v6 = (^(v6))
    var v67 int16 = (v2 | (v2 >> ((v6) & 63)))
    v11 -= 233
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v66:", len(v66), v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v67
}
