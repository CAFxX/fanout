package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int8; F1 int8; F2 int16 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int; F1 int32; F2 bool }
type GenS7 struct { F0 bool; F1 int64; F2 uint16 }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 uint16 }
func genFn0() uint32 {
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := []int{}
        var v7 uint64 = 4294967295
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint64 = 18446744073709551615
    if true {
        var v9 int8 = 42
        var v10 int8 = (-(v9))
        if false {
            if (9223372036854775807 != -199) {
                v11 := []int{}
                var v12 float64 = 0.1
                var v13 float64 = v12
                _ = v11
                _ = v12
                _ = v13
            } else {
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 uint64 = (v8 << ((uint(v8)) & 63))
                    _ = fi14
                    _ = v15
                }
                v16 := GenS1{F0: 0}
                if (!((32767 > 42) && false)) {
                    v17 := []int{-920, 32768, -9223372036854775808}
                    v17 = append(v17, -634)
                    _ = v17
                } else {
                    v18 := GenS2{F0: v9, F1: 42}
                    v19 := GenS3{F0: 34, F1: -32768}
                    _ = v18
                    _ = v19
                }
                _ = v16
            }
            var v20 float64 = ((1e10 - (-2.25 - 0.1)) * float64((0.1 * -2.25)))
            v9 += (v9 & v9)
            _ = v20
        } else {
            fmt.Println(v9, v10, v10)
            var v21 uint8 = 72
            _ = v21
        }
        _ = v10
        _ = v9
    } else {
        v8 += (v8 | (v8 + (v8 | v8)))
        v22 := GenS4{F0: -1, F1: -128, F2: -68}
        v8 = 466
        _ = v22
    }
    v23 := []int{}
    if (true && (0.1 < (3.14159 - 0.1))) {
        fmt.Println(v8)
    }
    _ = v23
    _ = v8
    return 863
}
func genFn1(p0 int64, p1 uint32) uint64 {
    v24 := GenS5{F0: -916, F1: -976}
    if false {
        switch ((("foo bar" + "\n\t") + "") + (((((("hello" + "foo bar") + ("\n\t" + "")) + "hello") + "\n\t") + ("\n\t" + "hello")) + "a")) {
        case "a":
            if true {
                p1 += (^(p1))
                p1 = (p1 >> ((42) & 63))
                var v25 int64 = (p0 * p0)
                _ = v25
            } else {
                v26 := GenS5{F0: 501, F1: 1}
                for fi27 := 0; fi27 < 4; fi27++ {
                    v28 := []int{(fi27 / (((-(fi27))) | 1)), fi27}
                    v29 := GenS6{F0: 32768, F1: 1, F2: true}
                    for fi30 := 0; fi30 < 9; fi30++ {
                        v31 := GenS7{F0: false, F1: p0, F2: 211}
                        v32 := GenS4{F0: -96, F1: 1, F2: 777}
                        v33 := GenS4{F0: 127, F1: -1, F2: 985}
                        fmt.Println(p1)
                        switch "" {
                        case "":
                            p1 = p1
                        case "a":
                            var v34 uint32 = 677
                            _ = v34
                        case "foo bar":
                            fmt.Println(p1)
                        }
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                if false {
                    var v35 int = 42
                    for fi36 := 0; fi36 < 8; fi36++ {
                        v37 := GenS8{F0: p1, F1: int64(p1)}
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                } else {
                    var v38 uint32 = (p1 * 5)
                    _ = v38
                }
                _ = v26
            }
        case "hello":
            var v39 uint8 = 88
            _ = v39
        }
        var v40 int = 1
        if true {
            v41 := []int{87}
            fmt.Println(v40, p1, p0)
            _ = v41
        }
        _ = v40
    }
    _ = p0
    _ = p1
    _ = v24
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{-2147483648}
    v3 := []uint32{}
    var v4 float64 = ((1.5 * 0.0) - 1.5)
    v42 := GenS3{F0: -128, F1: 220}
    if true {
        var v43 uint16 = 0
        if true {
            if ((!true) && ((false || true) && true)) {
                var v44 uint = uint(v43)
                fmt.Println(v44)
                _ = v44
            } else {
                for fi45 := 0; fi45 < 18; fi45++ {
                    if false {
                        if true {
                            v46 := GenS9{F0: (v43 ^ v43)}
                            if false {
                                fmt.Println(v4, fi45)
                                v4 = v4
                                switch "foo bar" {
                                case "foo bar":
                                    var v47 bool = false
                                    _ = v47
                                case "hello":
                                    switch -2147483648 {
                                    case 2147483647:
                                        v43 = 42
                                    case 4294967295:
                                        var v48 bool = true
                                        _ = v48
                                    case 128:
                                        for fi49 := 0; fi49 < 13; fi49++ {
                                            var v50 int = (fi45 >> ((18446744073709551615) & 63))
                                            v50 += int(v43)
                                            _ = fi49
                                            _ = v50
                                        }
                                    default:
                                        var v51 string = "a"
                                        _ = v51
                                    }
                                case "":
                                    var v52 int8 = 1
                                    _ = v52
                                }
                            }
                            _ = v46
                        } else {
                            for fi53 := 0; fi53 < 12; fi53++ {
                                var v54 uint = 417
                                _ = fi53
                                _ = v54
                            }
                            switch fi45 {
                            case 127:
                                v55 := GenS6{F0: (fi45 >> ((2) & 63)), F1: 0, F2: (255 <= -1)}
                                _ = v55
                            case -204:
                                var v56 uint32 = 2
                                _ = v56
                            }
                        }
                    }
                    v4 = float64((float32(v4) * 3.14159))
                    switch "\n\t" {
                    case "hello":
                        if len(v2) > 0 {
                            v2[0] = 4294967295
                        }
                    }
                    _ = fi45
                }
                _ = len(v1)
                var v57 string = ""
                _ = v57
            }
            var v58 float64 = 3.14159
            _ = v58
        } else {
            for fi59 := 0; fi59 < 9; fi59++ {
                var v60 uint64 = 4294967295
                _ = fi59
                _ = v60
            }
            var v61 uint16 = v43
            v61 = (v61 >> ((255) & 63))
            _ = v61
        }
        v4 -= v4
        _ = v43
    } else {
        var v62 int = 2147483647
        v4 += -2.25
        fmt.Println(v4)
        _ = v62
    }
    for fi63 := 0; fi63 < 18; fi63++ {
        var v64 uint16 = 65
        _ = fi63
        _ = v64
    }
    delete(v1, ("\n\t" + ""))
    if len(v3) > 0 {
        v3[0] = 65535
    }
    for fi65 := 0; fi65 < 18; fi65++ {
        var v66 bool = true
        _ = fi65
        _ = v66
    }
    fmt.Println(v4)
    var v67 uint64 = 150
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v67
}
