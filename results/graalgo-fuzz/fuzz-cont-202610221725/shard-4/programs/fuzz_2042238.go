package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint32; F2 uint16 }
func genFn0(p0 uint16) uint32 {
    fmt.Println(p0)
    var v6 int16 = -32768
    var v7 int32 = -2147483648
    _ = p0
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 uint64, p1 string) int8 {
    var v8 int32 = 1
    if (true && (true && (true && (459 != 852)))) {
        if true {
            var v9 float64 = -2.25
            switch (("foo bar" + p1) + p1) {
            case "foo bar":
                v10 := GenS1{F0: 128}
                _ = v10
            case "":
                fmt.Println(p1, v8, p0)
            default:
                if (true && true) {
                    var v11 uint8 = 47
                    for fi12 := 0; fi12 < 18; fi12++ {
                        v11 = v11
                        v13 := []int{fi12, (^(fi12))}
                        switch fi12 {
                        case -9223372036854775808:
                            if len(v13) > 0 {
                                v13[0] = (fi12 | (fi12 ^ fi12))
                            }
                        case -419:
                            fmt.Println(v11, v9)
                        }
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                } else {
                    v14 := []int{2147483647}
                    _ = v14
                }
            }
            _ = v9
        } else {
            v15 := GenS2{F0: true, F1: (-9223372036854775808 != 97)}
            var v16 float64 = 3.14159
            var v17 uint16 = uint16(p0)
            _ = v15
            _ = v16
            _ = v17
        }
        var v18 bool = (true || true)
        v18 = v18
        _ = v18
    } else {
        for fi19 := 0; fi19 < 14; fi19++ {
            p1 = p1
            _ = fi19
        }
        if (false && (1 > genFn0(1))) {
            p0 = (p0 >> ((18446744073709551615) & 63))
            p0 = p0
        }
    }
    switch p1 {
    case "":
        if true {
            var v20 string = p1
            switch "\n\t" {
            case "":
                v21 := GenS2{F0: false, F1: true}
                _ = v21
            case "hello":
                var v22 float32 = (-2.25 - ((1e10 + -2.25) - 1e10))
                _ = v22
            case "foo bar":
                if (42 != 32767) {
                    var v23 float64 = 0.1
                    var v24 bool = true
                    v20 = "a"
                    _ = v23
                    _ = v24
                } else {
                    v25 := []int{-9223372036854775808, 1, 9223372036854775807, -51}
                    p1 = v20
                    _ = v25
                }
            }
            v8 = -3
            _ = v20
        } else {
            v26 := []int{}
            if len(v26) > 0 {
                v26[0] = -1
            }
            fmt.Println(v8, v8)
            _ = v26
        }
    case "foo bar":
        if false {
            v27 := GenS1{F0: 0}
            switch p1 {
            case "foo bar":
                if true {
                    v28 := []int{int(v8)}
                    _ = v28
                }
            case "":
                v29 := GenS1{F0: 32768}
                _ = v29
            case "a":
                fmt.Println(p0, p1, p1)
            default:
                for fi30 := 0; fi30 < 3; fi30++ {
                    switch (-(fi30)) {
                    case 127:
                        var v31 int64 = 993
                        _ = v31
                    case -793:
                        p0 = p0
                    default:
                        var v32 int = -1
                        _ = v32
                    }
                    _ = fi30
                }
            }
            if (1 >= 42) {
                v8 = -1
                var v33 bool = false
                _ = v33
            } else {
                var v34 string = ""
                if (1000 > -749) {
                    v35 := []int{675, -2147483648, 127, 4}
                    _ = v35
                } else {
                    v36 := GenS3{F0: 42, F1: 301}
                    var v37 int = 0
                    if false {
                        fmt.Println(v34)
                        var v38 float32 = ((1.5 / 1e10) - 0.1)
                        fmt.Println(p1)
                        _ = v38
                    } else {
                        var v39 float64 = 0.1
                        _ = v39
                    }
                    _ = v36
                    _ = v37
                }
                for fi40 := 0; fi40 < 16; fi40++ {
                    fmt.Println(fi40)
                    _ = fi40
                }
                _ = v34
            }
            _ = v27
        } else {
            p0 = p0
            var v41 uint8 = 1
            v42 := []int{9223372036854775807, 42, -311, -2147483648}
            _ = v41
            _ = v42
        }
    case "hello":
        fmt.Println(v8)
    default:
        var v43 float64 = (3.14159 / 3.14159)
        _ = v43
    }
    _ = p0
    _ = p1
    _ = v8
    return -36
}

func main() {
    v1 := []int{}
    var v2 uint = 4294967295
    v3 := make(map[uint32]bool)
    var v4 int = -1
    var v5 int16 = 42
    fmt.Println(v4, v5, v2)
    if false {
        if (("" + ("" + ((((((("\n\t" + "hello") + "") + "") + ("\n\t" + "")) + (("foo bar" + "") + "hello")) + "a") + "foo bar"))) <= "a") {
            v2 = (v2 << ((v2) & 63))
            if (false && false) {
                v4 = v4
                v2 -= 2
            } else {
                v5 -= (-(v5))
                v2 += (^(v2))
                var v44 string = "hello"
                _ = v44
            }
        } else {
            v5 += (-(v5))
            v3[2147483647] = true
        }
    } else {
        switch (v4 >> ((967) & 63)) {
        case 106:
            var v45 bool = false
            _ = v45
        }
    }
    fmt.Println(v2, v4)
    for fi46 := 0; fi46 < 2; fi46++ {
        var v47 int16 = v5
        if (100 != 100) {
            v48 := GenS4{F0: 208}
            switch fi46 {
            case 128:
                v2 = (^(v2))
            default:
                for fi49 := 0; fi49 < 15; fi49++ {
                    switch v4 {
                    case 2147483647:
                        v47 = v5
                    case -9223372036854775808:
                        if (false && (0 != -128)) {
                            _ = len(v1)
                            fmt.Println(fi46)
                        } else {
                            fmt.Println(v47, v4, fi49)
                            _ = len(v1)
                            var v50 string = ""
                            fmt.Println(v50, v4)
                            if (v50 == v50) {
                                switch v50 {
                                case "foo bar":
                                    var v51 int8 = 1
                                    _ = v51
                                case "hello":
                                    var v52 uint64 = 4294967295
                                    _ = v52
                                }
                                v47 = (^(v47))
                            }
                            _ = v50
                        }
                    default:
                        if true {
                            _ = len(v1)
                            v2 += v2
                            fmt.Println(v5, fi46, v47)
                            v4 += (fi46 | (fi46 | (fi49 - 42)))
                            v2 -= (v2 | (^(v2)))
                            var v53 uint8 = 127
                            var v54 float64 = (1.5 / 0.1)
                            _ = v53
                            _ = v54
                        } else {
                            if true {
                                if (!(701 < 42)) {
                                    fmt.Println(fi46, fi46)
                                    var v55 uint64 = 18446744073709551615
                                    _ = v55
                                }
                                var v56 int8 = 100
                                _ = v56
                            } else {
                                var v57 float64 = 0.1
                                v58 := GenS1{F0: (fi49 + (fi46 + -542))}
                                var v59 uint32 = 42
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                        }
                    }
                    var v60 uint16 = 0
                    v3[42] = true
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v2 = 65535
                    _ = fi49
                    _ = v60
                }
            }
            v5 -= v5
            _ = v48
        }
        var v61 int32 = 1
        _ = fi46
        _ = v47
        _ = v61
    }
    v5 = 1
    var v62 uint = (v2 << (((^(v2))) & 63))
    switch v4 {
    case 127:
        switch v4 {
        case 2147483647:
            if (127 < (-(v4))) {
                fmt.Println(v62, v5)
                for fi63 := 0; fi63 < 12; fi63++ {
                    for fi64 := 0; fi64 < 2; fi64++ {
                        if len(v1) > 0 {
                            v1[0] = fi63
                        }
                        if false {
                            v65 := GenS4{F0: 118}
                            switch "" {
                            case "foo bar":
                                switch v4 {
                                case -775:
                                    v2 += v2
                                case 128:
                                    v4 += fi64
                                case 42:
                                    v2 += v62
                                }
                            case "a":
                                if true {
                                    fmt.Println(fi64, v62, fi63)
                                    v5 -= (v5 | v5)
                                }
                            case "":
                                v66 := GenS2{F0: (2147483647 <= (^(v4))), F1: false}
                                _ = v66
                            }
                            _ = v65
                        }
                        v67 := GenS5{F0: "hello", F1: 713, F2: 616}
                        delete(v3, 4294967295)
                        _ = fi64
                        _ = v67
                    }
                    for fi68 := 0; fi68 < 5; fi68++ {
                        switch "hello" {
                        case "foo bar":
                            var v69 string = "a"
                            _ = v69
                        case "a":
                            var v70 int32 = 0
                            _ = v70
                        default:
                            var v71 float32 = float32((-2.25 * 1.5))
                            _ = v71
                        }
                        var v72 string = ""
                        _ = fi68
                        _ = v72
                    }
                    if ((^(v62)) != 467) {
                        var v73 uint = 0
                        _ = v73
                    } else {
                        v62 = (^(v2))
                    }
                    _ = fi63
                }
            } else {
                var v74 uint = v2
                _ = v74
            }
        default:
            if true {
                v75 := GenS3{F0: 32767, F1: v5}
                switch v4 {
                case -1:
                    var v76 uint32 = 4294967295
                    _ = v76
                }
                for fi77 := 0; fi77 < 11; fi77++ {
                    var v78 uint = (v62 % (((^(v2))) | 1))
                    var v79 uint32 = 42
                    fmt.Println(v4, fi77, v79)
                    _ = fi77
                    _ = v78
                    _ = v79
                }
                _ = v75
            }
        }
    default:
        switch ("" + ("foo bar" + ("hello" + "\n\t"))) {
        case "\n\t":
            v2 += v62
        case "a":
            fmt.Println(v4, v2, v5)
        default:
            v62 = v62
        }
    }
    v4 = (v4 & v4)
    fmt.Println(v2, v4, v62)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
}
