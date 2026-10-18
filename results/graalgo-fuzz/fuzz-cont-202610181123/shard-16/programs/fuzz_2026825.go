package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int64; F1 int; F2 uint32 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 uint32; F1 string }
func genFn0(p0 uint64, p1 int32) int16 {
    v7 := GenS1{F0: 119, F1: 1}
    p0 = 963
    _ = p0
    _ = p1
    _ = v7
    return 995
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]string)
    var v3 bool = (!(!(42 != -1)))
    v4 := make(map[uint32]bool)
    var v5 int = 9223372036854775807
    var v6 int8 = -1
    v8 := v4[10]
    v9 := v2["foo bar"]
    var v10 int32 = -744
    if true {
        if (1 != 2) {
            v11 := GenS2{F0: 1, F1: (v5 << ((v1) & 63)), F2: 0}
            v5 -= v5
            v12 := []int{}
            _ = v11
            _ = v12
        } else {
            fmt.Println(v6, v1, v3)
            for fi13 := 0; fi13 < 6; fi13++ {
                fmt.Println(v3, v10, v8)
                _ = fi13
            }
        }
        _ = len(v4)
        fmt.Println(v1)
    }
    if false {
        v14 := v4[2147483647]
        _ = v14
    } else {
        for fi15 := 0; fi15 < 17; fi15++ {
            switch (v5 | fi15) {
            case 42:
                v16 := GenS3{F0: 83, F1: v3}
                _ = v16
            case 2147483647:
                v1 = (^(v1))
            case 32767:
                v9 = ("a" + v9)
            default:
                v8 = true
            }
            v17 := GenS1{F0: 128, F1: 0}
            _ = fi15
            _ = v17
        }
        if v3 {
            switch 32768 {
            case -2147483648:
                var v18 uint16 = 700
                _ = v18
            case 42:
                for fi19 := 0; fi19 < 2; fi19++ {
                    if false {
                        var v20 uint8 = 127
                        fmt.Println(v5, fi19)
                        if false {
                            for fi21 := 0; fi21 < 15; fi21++ {
                                var v22 int = 42
                                v22 = v5
                                _ = fi21
                                _ = v22
                            }
                            v23 := []int{fi19, v5, (v5 + fi19), v5}
                            _ = v23
                        }
                        _ = v20
                    } else {
                        v24 := GenS4{F0: 2}
                        v3 = v3
                        v8 = v3
                        _ = v24
                    }
                    var v25 float64 = (-2.25 + 0.0)
                    _ = fi19
                    _ = v25
                }
            case 32768:
                var v26 int16 = -37
                _ = v26
            default:
                v6 = v6
            }
            var v27 uint32 = 0
            fmt.Println(v5, v3)
            _ = v27
        } else {
            var v28 int16 = -570
            if v8 {
                for fi29 := 0; fi29 < 6; fi29++ {
                    switch (v9 + v9) {
                    case "":
                        var v30 float64 = 0.1
                        _ = v30
                    case "foo bar":
                        v6 += -63
                    }
                    v31 := GenS2{F0: 0, F1: fi29, F2: 954}
                    _ = fi29
                    _ = v31
                }
                for fi32 := 0; fi32 < 16; fi32++ {
                    _ = len(v2)
                    v33 := []int{-9223372036854775808, (fi32 | 32768), v5}
                    v28 = (v28 % (((-(v28))) | 1))
                    _ = fi32
                    _ = v33
                }
                var v34 int64 = -653
                _ = v34
            }
            delete(v4, 445)
            _ = v28
        }
        switch (v5 / (((v5 + v5)) | 1)) {
        case -698:
            for fi35 := 0; fi35 < 17; fi35++ {
                for fi36 := 0; fi36 < 1; fi36++ {
                    var v37 uint32 = 4294967295
                    switch -205 {
                    case 127:
                        if (411 <= -546) {
                            var v38 int64 = -1
                            fmt.Println(v37, fi35, v37)
                            switch "\n\t" {
                            case "":
                                v39 := GenS2{F0: 789, F1: fi35, F2: v37}
                                _ = v39
                            case "a":
                                v40 := GenS2{F0: v38, F1: fi36, F2: 42}
                                if true {
                                    var v41 uint64 = 2
                                    _ = v41
                                } else {
                                    v42 := GenS4{F0: 937}
                                    _ = v42
                                }
                                _ = v40
                            case "hello":
                                v43 := []int{(fi35 - (-(fi35)))}
                                _ = v43
                            default:
                                v38 += v38
                            }
                            _ = v38
                        }
                    default:
                        var v44 int = fi35
                        _ = v44
                    }
                    var v45 int8 = -128
                    _ = fi36
                    _ = v37
                    _ = v45
                }
                v3 = v3
                _ = fi35
            }
        }
    }
    v46 := GenS3{F0: -532, F1: v8}
    delete(v2, (v9 + ("\n\t" + v9)))
    if (v10 < (v10 | v10)) {
        v5 = (v5 & (v5 * v5))
        v1 = v1
        if v3 {
            var v47 float32 = 3.14159
            switch (((v9 + v9) + v9) + v9) {
            case "foo bar":
                for fi48 := 0; fi48 < 2; fi48++ {
                    var v49 int = v5
                    var v50 uint8 = 0
                    for fi51 := 0; fi51 < 14; fi51++ {
                        v10 = (^(v10))
                        _ = fi51
                    }
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            default:
                fmt.Println(v3, v6, v10)
            }
            fmt.Println(v10, v1)
            _ = v47
        } else {
            var v52 bool = v8
            switch "a" {
            case "":
                v53 := GenS3{F0: 42, F1: v52}
                _ = v53
            case "\n\t":
                v52 = v8
            }
            if v52 {
                switch v5 {
                case 128:
                    for fi54 := 0; fi54 < 10; fi54++ {
                        switch v9 {
                        case "a":
                            v3 = v8
                        case "":
                            v55 := GenS5{F0: v1, F1: 651}
                            _ = v55
                        case "\n\t":
                            v56 := GenS1{F0: 128, F1: 127}
                            _ = v56
                        default:
                            v57 := GenS2{F0: 0, F1: fi54, F2: 2147483647}
                            _ = v57
                        }
                        var v58 uint8 = 114
                        _ = fi54
                        _ = v58
                    }
                default:
                    v59 := []int{v5, 11, (v5 % (((v5 >> ((v1) & 63))) | 1))}
                    _ = v59
                }
                var v60 bool = ((!(((v9 + "hello") == "foo bar") || (!true))) || false)
                _ = v60
            } else {
                fmt.Println(v10, v52)
                switch (v5 + -1) {
                case 1:
                    fmt.Println(v6, v52, v5)
                case 127:
                    if v52 {
                        v61 := GenS1{F0: 55, F1: 128}
                        _ = v61
                    } else {
                        v4[2147483647] = (true || true)
                        var v62 bool = v52
                        _ = v62
                    }
                }
            }
            _ = v52
        }
    } else {
        v63 := []int{v5, (v5 % ((4294967295) | 1)), int(v6), int(v1)}
        if false {
            for fi64 := 0; fi64 < 13; fi64++ {
                v8 = v3
                _ = fi64
            }
        } else {
            if false {
                var v65 int16 = 1
                _ = v65
            }
        }
        v10 = v10
        _ = v63
    }
    var v66 uint32 = 65535
    v67 := GenS6{F0: v66, F1: ""}
    fmt.Println(v66, v8, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v8
    _ = v9
}
