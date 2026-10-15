package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint; F2 uint64 }
type GenS3 struct { F0 int32; F1 int16; F2 int64 }
type GenS4 struct { F0 string; F1 int; F2 string }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 uint16; F2 uint32 }
func genFn0(p0 int16, p1 int32) uint16 {
    fmt.Println(p1, p1)
    fmt.Println(p1)
    p0 += (-(p0))
    v6 := GenS1{F0: -416, F1: 1, F2: 1}
    p1 = (p1 >> ((180) & 63))
    _ = p0
    _ = p1
    _ = v6
    return 846
}
func genFn1(p0 uint64) uint64 {
    var v7 uint32 = 2147483647
    v7 = v7
    v8 := GenS1{F0: -1, F1: -1, F2: p0}
    _ = p0
    _ = v7
    _ = v8
    return (p0 + p0)
}

func main() {
    var v1 int = -2147483648
    var v2 int = v1
    var v3 int32 = 676
    var v4 uint = 814
    var v5 int = (v2 << ((v4) & 63))
    var v9 int = 128
    var v10 uint32 = 465
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            fmt.Println(fi11, v1)
            v13 := []int{fi12}
            fmt.Println(v4, fi12, v2)
            _ = fi12
            _ = v13
        }
        fmt.Println(v3, v9)
        _ = fi11
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        if true {
            v15 := GenS1{F0: -9223372036854775808, F1: v9, F2: 138}
            _ = v15
        } else {
            if ((1e10 > 0.0) || false) {
                for fi16 := 0; fi16 < 12; fi16++ {
                    if (!false) {
                        var v17 int32 = v3
                        v18 := []int{v5}
                        var v19 uint64 = 514
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = fi16
                }
                for fi20 := 0; fi20 < 1; fi20++ {
                    var v21 uint8 = 138
                    v21 += 138
                    if false {
                        var v22 uint64 = 820
                        var v23 uint = uint(v22)
                        _ = v22
                        _ = v23
                    } else {
                        var v24 float64 = (float64(3.14159) / 1.0)
                        var v25 uint32 = v10
                        _ = v24
                        _ = v25
                    }
                    _ = fi20
                    _ = v21
                }
                v26 := GenS1{F0: 9223372036854775807, F1: 42, F2: 1}
                _ = v26
            } else {
                v27 := GenS2{F0: -95, F1: (^(v4)), F2: 96}
                switch ("foo bar" + ("foo bar" + "hello")) {
                case "":
                    var v28 uint16 = 32767
                    _ = v28
                case "\n\t":
                    var v29 uint64 = uint64(v10)
                    _ = v29
                case "a":
                    for fi30 := 0; fi30 < 5; fi30++ {
                        v2 += (fi14 * fi14)
                        v31 := []int{}
                        v2 -= (fi14 ^ v5)
                        _ = fi30
                        _ = v31
                    }
                default:
                    v32 := []int{}
                    _ = v32
                }
                _ = v27
            }
            var v33 int16 = 42
            _ = v33
        }
        _ = fi14
    }
    v34 := []int{(v2 - int(v10)), int(v4), 128, (v1 << ((uint(v10)) & 63))}
    var v35 uint8 = 71
    for fi36 := 0; fi36 < 16; fi36++ {
        var v37 uint = uint(v35)
        fmt.Println(v5, v2)
        v9 = (v9 * v1)
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 int32 = 2147483647
        var v40 uint16 = 32767
        var v41 uint16 = genFn0(0, v39)
        _ = fi38
        _ = v39
        _ = v40
        _ = v41
    }
    v42 := GenS3{F0: v3, F1: 626, F2: -9223372036854775808}
    v2 = int(v4)
    v43 := GenS4{F0: "", F1: v2, F2: "\n\t"}
    if (-1 != 59) {
        fmt.Println(v35, v2)
    } else {
        for fi44 := 0; fi44 < 16; fi44++ {
            v9 -= -9223372036854775808
            _ = fi44
        }
        var v45 float64 = float64(0.0)
        _ = v45
    }
    if true {
        switch v2 {
        case 42:
            for fi46 := 0; fi46 < 15; fi46++ {
                if ((1.5 < float32(((0.0 - 0.0) / 3.14159))) && (true && false)) {
                    v47 := GenS5{F0: 32767, F1: v4}
                    fmt.Println(v3)
                    if false {
                        v5 += -881
                        if (false && true) {
                            switch (v5 >> ((1) & 63)) {
                            case 4294967295:
                                if (!false) {
                                    v48 := GenS3{F0: (v3 & 831), F1: -404, F2: 0}
                                    fmt.Println(v1)
                                    _ = v48
                                } else {
                                    if false {
                                        fmt.Println(v5)
                                        _ = len(v34)
                                        v35 += v35
                                    } else {
                                        var v49 uint64 = genFn1(4294967295)
                                        v10 += (v10 << ((255) & 63))
                                        v1 -= int(v35)
                                        _ = v49
                                    }
                                    if (!false) {
                                        if (genFn0(313, -2147483648) <= 65535) {
                                            _ = len(v34)
                                            fmt.Println(v5)
                                        } else {
                                            fmt.Println(v9, fi46)
                                        }
                                    } else {
                                        for fi50 := 0; fi50 < 20; fi50++ {
                                            v10 = 2147483648
                                            if len(v34) > 0 {
                                                v34[0] = (v2 << ((255) & 63))
                                            }
                                            var v51 uint16 = 109
                                            _ = fi50
                                            _ = v51
                                        }
                                    }
                                }
                            case -862:
                                v52 := GenS5{F0: 42, F1: (v4 >> ((v4) & 63))}
                                _ = v52
                            }
                        }
                        var v53 float64 = float64((0.1 * 0.1))
                        _ = v53
                    } else {
                        var v54 string = "foo bar"
                        switch ((("" + v54) + v54) + (v54 + (v54 + v54))) {
                        case "\n\t":
                            v55 := GenS5{F0: 617, F1: v4}
                            _ = v55
                        default:
                            for fi56 := 0; fi56 < 1; fi56++ {
                                v35 = v35
                                var v57 uint64 = 42
                                _ = fi56
                                _ = v57
                            }
                        }
                        _ = v54
                    }
                    _ = v47
                } else {
                    v58 := GenS2{F0: -78, F1: 42, F2: 4294967295}
                    var v59 uint8 = v35
                    _ = v58
                    _ = v59
                }
                var v60 int = (^(v2))
                _ = fi46
                _ = v60
            }
        case 0:
            var v61 int64 = 1
            _ = v61
        }
        v62 := GenS6{F0: -856}
        _ = v62
    }
    if ("foo bar" <= (("a" + "hello") + ("" + ""))) {
        var v63 string = (("a" + ("foo bar" + "\n\t")) + "\n\t")
        for fi64 := 0; fi64 < 10; fi64++ {
            var v65 int16 = -1
            switch fi64 {
            case -875:
                var v66 string = v63
                _ = v66
            }
            v67 := GenS7{F0: "\n\t", F1: 32767, F2: 42}
            _ = fi64
            _ = v65
            _ = v67
        }
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v9
}
