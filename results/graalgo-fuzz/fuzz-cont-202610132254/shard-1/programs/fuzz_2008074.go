package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 string; F2 uint }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 int; F2 uint }
type GenS7 struct { F0 int64; F1 uint32; F2 bool }
type GenS8 struct { F0 uint32; F1 uint }
type GenS9 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 uint16 = 255
    var v4 uint32 = 4294967295
    v5 := []uint32{}
    var v6 int = 359
    if (true && false) {
        for fi7 := 0; fi7 < 16; fi7++ {
            if (42 >= 2) {
                v3 -= uint16(v4)
            } else {
                for fi8 := 0; fi8 < 20; fi8++ {
                    var v9 uint = 4294967295
                    _ = fi8
                    _ = v9
                }
            }
            _ = fi7
        }
        v4 = v4
    }
    var v10 uint16 = 0
    var v11 bool = true
    var v12 float32 = -2.25
    v6 += (v6 - v6)
    if false {
        var v13 float32 = 3.14159
        var v14 int32 = -1
        v6 = v6
        _ = v13
        _ = v14
    } else {
        if true {
            fmt.Println(v4, v4)
            if true {
                v15 := GenS1{F0: 42}
                v16 := GenS2{F0: v4}
                for fi17 := 0; fi17 < 1; fi17++ {
                    switch "\n\t" {
                    case "\n\t":
                        var v18 uint16 = (v3 - v10)
                        _ = v18
                    default:
                        fmt.Println(v4, v6)
                    }
                    if (!v11) {
                        var v19 int16 = 0
                        _ = v19
                    }
                    var v20 uint = 574
                    _ = fi17
                    _ = v20
                }
                _ = v15
                _ = v16
            } else {
                v4 -= v4
                if v11 {
                    if (v11 || v11) {
                        v21 := GenS3{F0: 0, F1: (("hello" + ("" + "a")) + "\n\t"), F2: 539}
                        var v22 int8 = 100
                        var v23 int64 = 806
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        var v24 string = ("\n\t" + "")
                        var v25 uint8 = 255
                        v26 := GenS2{F0: 0}
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    var v27 uint32 = 2147483648
                    var v28 uint32 = 886
                    _ = v27
                    _ = v28
                } else {
                    if false {
                        var v29 int32 = 2147483647
                        for fi30 := 0; fi30 < 13; fi30++ {
                            if true {
                                v12 -= (0.0 - ((v12 / (0.1 / 1e10)) * v12))
                                for fi31 := 0; fi31 < 5; fi31++ {
                                    _ = len(v5)
                                    delete(v2, "foo bar")
                                    _ = fi31
                                }
                                if len(v5) > 0 {
                                    v5[0] = v4
                                }
                            }
                            var v32 int8 = 7
                            _ = fi30
                            _ = v32
                        }
                        v6 -= v6
                        _ = v29
                    }
                    v33 := GenS4{F0: 4294967295, F1: 1}
                    _ = v33
                }
            }
        } else {
            if ((v4 - v4) != 4294967295) {
                if true {
                    for fi34 := 0; fi34 < 8; fi34++ {
                        v35 := GenS1{F0: 127}
                        v36 := GenS5{F0: v3}
                        var v37 bool = false
                        _ = fi34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    v38 := GenS6{F0: 42, F1: -634, F2: 981}
                    _ = v38
                } else {
                    v1[4294967295] = true
                    v12 = v12
                }
                v10 = v3
            }
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                var v39 uint64 = 784
                _ = v39
            default:
                v12 += -2.25
            }
            v10 += v3
        }
    }
    switch (v6 + int(v4)) {
    case -342:
        fmt.Println(v6, v11, v10)
    case -782:
        switch (v6 / (((v6 ^ v6)) | 1)) {
        case 338:
            var v40 int32 = -150
            _ = v40
        case 21:
            if v11 {
                if (0 > 0) {
                    v12 -= (v12 - 1.5)
                    var v41 float32 = 1e10
                    if true {
                        v42 := GenS7{F0: -304, F1: (v4 >> ((861) & 63)), F2: (-164 >= (v6 + v6))}
                        switch (v6 / ((127) | 1)) {
                        case 982:
                            v43 := GenS8{F0: (v4 % ((v4) | 1)), F1: 410}
                            _ = v43
                        default:
                            if (100 < 0) {
                                if (!v11) {
                                    if v11 {
                                        v44 := GenS9{F0: uint64(v3)}
                                        for fi45 := 0; fi45 < 5; fi45++ {
                                            var v46 int = (v6 * 4294967295)
                                            var v47 string = ("" + "\n\t")
                                            _ = fi45
                                            _ = v46
                                            _ = v47
                                        }
                                        _ = v44
                                    } else {
                                        v10 = 255
                                        var v48 float64 = ((float64((v12 / v12)) * (float64(0.0) + (1e10 / 0.1))) / 3.14159)
                                        _ = v48
                                    }
                                } else {
                                    var v49 string = ""
                                    _ = v49
                                }
                                var v50 bool = (!v11)
                                _ = v50
                            }
                        }
                        _ = v42
                    }
                    _ = v41
                }
            }
        }
    case -1:
        var v51 int = (v6 % ((-9223372036854775808) | 1))
        _ = v51
    }
    v4 -= v4
    var v52 float32 = ((v12 - v12) - (v12 - 0.0))
    v11 = (false && v11)
    fmt.Println(v12, v10, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
