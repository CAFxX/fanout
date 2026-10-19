package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint16; F1 bool; F2 string }

func main() {
    var v1 uint64 = 1
    var v2 uint32 = 780
    var v3 int16 = -32768
    var v4 string = ("hello" + ("" + "\n\t"))
    v1 = 0
    if false {
        fmt.Println(v1)
    } else {
        switch -9223372036854775808 {
        case 127:
            v5 := []int{}
            _ = v5
        default:
            switch int(v2) {
            case -2147483648:
                if ((((!(-950 > -2147483648)) || true) && (!false)) && (v3 == 42)) {
                    if false {
                        for fi6 := 0; fi6 < 2; fi6++ {
                            for fi7 := 0; fi7 < 9; fi7++ {
                                v4 = (v4 + v4)
                                v2 -= v2
                                _ = fi7
                            }
                            v8 := GenS1{F0: fi6}
                            _ = fi6
                            _ = v8
                        }
                        for fi9 := 0; fi9 < 9; fi9++ {
                            switch -9223372036854775808 {
                            case 32767:
                                fmt.Println(v4, fi9)
                            case -659:
                                for fi10 := 0; fi10 < 8; fi10++ {
                                    if (((!(1e10 == 0.0)) || true) || true) {
                                        if (!false) {
                                            v11 := GenS2{F0: (v4 + "a"), F1: (fi9 & (fi10 & fi10)), F2: 644}
                                            v2 = 587
                                            v12 := GenS3{F0: 0, F1: 268}
                                            v13 := GenS1{F0: 0}
                                            _ = v11
                                            _ = v12
                                            _ = v13
                                        }
                                        var v14 uint32 = 245
                                        if false {
                                            var v15 uint32 = (v2 * (v2 % ((v14) | 1)))
                                            _ = v15
                                        }
                                        _ = v14
                                    }
                                    v3 -= 87
                                    _ = fi10
                                }
                            }
                            _ = fi9
                        }
                        if ((!false) || false) {
                            var v16 int = 127
                            var v17 uint32 = v2
                            _ = v16
                            _ = v17
                        } else {
                            if false {
                                var v18 int = 2147483647
                                var v19 bool = (-414 <= 2147483647)
                                v20 := []int{v18, (v18 % ((2147483647) | 1)), (^(v18)), 127}
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                            v2 = v2
                        }
                    }
                    if (((1.5 * 3.14159) - 1e10) != 0.1) {
                        var v21 uint = 811
                        v22 := GenS2{F0: ("hello" + "\n\t"), F1: -464, F2: 42}
                        _ = v21
                        _ = v22
                    }
                } else {
                    var v23 float32 = 0.0
                    _ = v23
                }
            case -520:
                v24 := []int{4294967295, 32767, 271}
                _ = v24
            }
        }
        var v25 uint32 = (^(v2))
        v26 := GenS4{F0: 65535, F1: false, F2: v4}
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        v28 := []int{32767, 127, (^(fi27)), (fi27 + fi27)}
        _ = fi27
        _ = v28
    }
    var v29 uint16 = 65535
    var v30 string = "a"
    var v31 int8 = 1
    var v32 int32 = 42
    var v33 int = int(v32)
    if (!(v29 <= (^(v29)))) {
        var v34 uint64 = v1
        if true {
            var v35 float32 = float32(0.0)
            for fi36 := 0; fi36 < 6; fi36++ {
                var v37 int16 = v3
                var v38 int32 = v32
                var v39 int32 = 0
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
            }
            _ = v35
        }
        var v40 bool = false
        _ = v34
        _ = v40
    }
    v41 := []int{128}
    var v42 uint = 2
    var v43 int32 = 42
    switch 128 {
    case 9223372036854775807:
        for fi44 := 0; fi44 < 19; fi44++ {
            switch (((v30 + ("\n\t" + "hello")) + "hello") + v4) {
            case "hello":
                var v45 string = (("" + "") + "")
                _ = v45
            case "\n\t":
                v31 -= v31
            case "foo bar":
                for fi46 := 0; fi46 < 18; fi46++ {
                    v33 = fi46
                    _ = fi46
                }
            default:
                var v47 int16 = (v3 / (((v3 << ((v42) & 63))) | 1))
                _ = v47
            }
            for fi48 := 0; fi48 < 20; fi48++ {
                v30 = "hello"
                v30 = (v4 + "")
                if len(v41) > 0 {
                    v41[0] = 0
                }
                _ = fi48
            }
            _ = fi44
        }
    case -2147483648:
        v33 = v33
    }
    v49 := GenS2{F0: "hello", F1: v33, F2: uint16(v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v49
}
