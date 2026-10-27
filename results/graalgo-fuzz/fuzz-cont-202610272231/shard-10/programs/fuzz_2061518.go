package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 string; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint64; F1 int8; F2 int32 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 string; F1 uint64; F2 uint32 }
func genFn0() uint64 {
    var v4 uint32 = 42
    fmt.Println(v4)
    _ = v4
    return 18446744073709551615
}
func genFn1() uint16 {
    var v5 float32 = (1e10 - (3.14159 + 1e10))
    if ((("hello" + "") + "\n\t") < "") {
        for fi6 := 0; fi6 < 3; fi6++ {
            v5 = v5
            _ = fi6
        }
    }
    v7 := GenS1{F0: (("a" + ("" + "a")) + ""), F1: 42}
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(v5, fi8)
        v5 -= v5
        _ = fi8
    }
    _ = v5
    _ = v7
    return 429
}

func main() {
    var v1 int = 1
    v2 := make(map[uint32]bool)
    var v3 bool = false
    var v9 uint16 = 65535
    if (2 == 2) {
        v10 := GenS2{F0: "\n\t", F1: 2, F2: -32768}
        fmt.Println(v3, v1, v3)
        _ = v10
    }
    var v11 int32 = 42
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 uint32 = 2147483648
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: ("foo bar" + "\n\t"), F1: 4294967295}
    if v3 {
        v9 = genFn1()
    } else {
        v15 := []int{1}
        if v3 {
            if (!(v3 && (255 != 255))) {
                switch v1 {
                case 1:
                    v16 := v2[2]
                    _ = v16
                case 32768:
                    fmt.Println(v9)
                case 128:
                    v3 = (!(2 == 735))
                default:
                    fmt.Println(v11)
                }
                if (1 < 65535) {
                    var v17 float64 = 1e10
                    _ = v17
                }
                if v3 {
                    var v18 bool = (18446744073709551615 == 1)
                    switch v1 {
                    case 2147483647:
                        fmt.Println(v9, v18, v1)
                    case 4294967295:
                        for fi19 := 0; fi19 < 3; fi19++ {
                            var v20 uint16 = (v9 % ((v9) | 1))
                            _ = fi19
                            _ = v20
                        }
                    case 32768:
                        for fi21 := 0; fi21 < 20; fi21++ {
                            v18 = (uint(v9) < 65535)
                            v22 := GenS3{F0: 2, F1: 54, F2: v11}
                            v3 = (v3 || true)
                            _ = fi21
                            _ = v22
                        }
                    default:
                        for fi23 := 0; fi23 < 18; fi23++ {
                            v9 = v9
                            switch "a" {
                            case "a":
                                fmt.Println(v3)
                            case "hello":
                                for fi24 := 0; fi24 < 1; fi24++ {
                                    v25 := GenS1{F0: "\n\t", F1: 2147483647}
                                    v1 = (fi23 ^ (v1 & (v1 & fi23)))
                                    _ = fi24
                                    _ = v25
                                }
                            case "":
                                for fi26 := 0; fi26 < 15; fi26++ {
                                    switch v1 {
                                    case 128:
                                        var v27 uint64 = 0
                                        _ = v27
                                    }
                                    fmt.Println(fi26)
                                    switch fi26 {
                                    case 32768:
                                        v1 -= 32768
                                    case 127:
                                        v28 := GenS4{F0: 40, F1: ("a" + "hello")}
                                        _ = v28
                                    case 4294967295:
                                        _ = len(v15)
                                    }
                                    _ = fi26
                                }
                            default:
                                fmt.Println(v11, v1, v1)
                            }
                            v11 = (v11 + (v11 * (v11 >> ((42) & 63))))
                            _ = fi23
                        }
                    }
                    _ = v18
                } else {
                    switch (v1 - v1) {
                    case -114:
                        switch ("\n\t" + ("foo bar" + "foo bar")) {
                        case "a":
                            if (!v3) {
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(v11, v1, v1)
                                case "hello":
                                    for fi29 := 0; fi29 < 6; fi29++ {
                                        _ = len(v15)
                                        _ = fi29
                                    }
                                case "a":
                                    delete(v2, 1)
                                default:
                                    v11 = (-(v11))
                                }
                            } else {
                                var v30 int32 = v11
                                var v31 uint8 = 241
                                _ = v30
                                _ = v31
                            }
                        }
                    case 234:
                        _ = len(v15)
                    default:
                        if len(v15) > 0 {
                            v15[0] = v1
                        }
                    }
                    delete(v2, 1)
                    var v32 uint = 0
                    _ = v32
                }
            } else {
                var v33 float64 = 0.0
                var v34 uint16 = v9
                _ = v33
                _ = v34
            }
            fmt.Println(v3, v9, v11)
            var v35 uint8 = 152
            _ = v35
        }
        _ = len(v2)
        _ = v15
    }
    if false {
        v3 = false
        for fi36 := 0; fi36 < 7; fi36++ {
            var v37 uint16 = (v9 >> ((326) & 63))
            if (255 <= 153) {
                switch "\n\t" {
                case "foo bar":
                    v38 := []int{(v1 + (fi36 | fi36)), v1}
                    _ = v38
                }
                switch "\n\t" {
                case "":
                    if false {
                        var v39 uint = 65535
                        if v3 {
                            v1 -= 4294967295
                        } else {
                            var v40 int32 = (v11 / (((v11 | (v11 & v11))) | 1))
                            var v41 uint16 = (v37 / ((v37) | 1))
                            v1 = v1
                            _ = v40
                            _ = v41
                        }
                        var v42 uint = v39
                        _ = v39
                        _ = v42
                    } else {
                        var v43 int32 = v11
                        switch (fi36 - fi36) {
                        case -2147483648:
                            if (!(!((!(115 > 47)) && v3))) {
                                v37 += v37
                                if v3 {
                                    v44 := []int{42}
                                    var v45 int16 = 0
                                    v46 := GenS2{F0: "hello", F1: 42, F2: 32767}
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                } else {
                                    v11 += v11
                                    var v47 uint64 = 42
                                    v48 := GenS5{F0: 42, F1: "\n\t"}
                                    _ = v47
                                    _ = v48
                                }
                                fmt.Println(v9)
                            }
                        case 0:
                            for fi49 := 0; fi49 < 14; fi49++ {
                                if v3 {
                                    var v50 int64 = -9223372036854775808
                                    if v3 {
                                        if v3 {
                                            v3 = (42 != 42)
                                            v9 = (^(v9))
                                            fmt.Println(fi49, v11)
                                            var v51 uint32 = 0
                                            switch (^(fi36)) {
                                            case -158:
                                                if (42 > 0) {
                                                    var v52 float32 = float32(1.5)
                                                    _ = v52
                                                }
                                            case 128:
                                                v37 += (v37 >> ((42) & 63))
                                            case -47:
                                                v50 += -1
                                                v51 = 0
                                                v2[v51] = v3
                                            default:
                                                v53 := []int{32767, 273, (^(fi49))}
                                                _ = v53
                                            }
                                            v54 := GenS4{F0: -1, F1: "hello"}
                                            _ = v51
                                            _ = v54
                                        }
                                        fmt.Println(v37, v3)
                                        v43 = (v11 * v11)
                                        var v55 int16 = 891
                                        _ = v55
                                    }
                                    _ = v50
                                } else {
                                    v56 := GenS3{F0: 2, F1: -86, F2: 42}
                                    var v57 string = "a"
                                    v58 := GenS6{F0: v57, F1: 0, F2: 2147483648}
                                    fmt.Println(v57, v3)
                                    if (1 > 127) {
                                        v3 = (1.5 > 3.14159)
                                    } else {
                                        var v59 bool = (39 < -9223372036854775808)
                                        fmt.Println(v11, fi36)
                                        _ = v59
                                    }
                                    if (668 == 4294967295) {
                                        v11 -= (v43 + v43)
                                        v37 -= v37
                                    }
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                }
                                if v3 {
                                    v37 = genFn1()
                                    switch (("hello" + ("a" + "foo bar")) + ("a" + "")) {
                                    case "a":
                                        var v60 uint8 = 127
                                        v43 += v11
                                        _ = len(v2)
                                        _ = v60
                                    case "":
                                        v61 := []int{}
                                        _ = v61
                                    }
                                    var v62 int = (-(fi36))
                                    v37 -= 65535
                                    v9 = v37
                                    if (672 > 4294967295) {
                                        v62 = 2147483647
                                        if v3 {
                                            if (127 >= 56) {
                                                var v63 float32 = 0.1
                                                _ = v63
                                            }
                                        } else {
                                            if (42 <= 949) {
                                                if v3 {
                                                    v64 := GenS2{F0: ("" + "hello"), F1: 2147483647, F2: 1000}
                                                    var v65 uint64 = 18446744073709551615
                                                    _ = v64
                                                    _ = v65
                                                }
                                            } else {
                                                var v66 int16 = 707
                                                _ = v66
                                            }
                                            var v67 bool = v3
                                            var v68 float32 = 0.0
                                            _ = v67
                                            _ = v68
                                        }
                                        var v69 int16 = 1
                                        _ = v69
                                    }
                                    fmt.Println(v11, fi49)
                                    _ = v62
                                } else {
                                    var v70 uint16 = (v37 + (v9 * 808))
                                    v71 := GenS2{F0: "\n\t", F1: 202, F2: 730}
                                    v72 := []int{fi36, v1}
                                    v72 = append(v72, v1)
                                    _ = v70
                                    _ = v71
                                    _ = v72
                                }
                                _ = fi49
                            }
                        case 164:
                            v37 = v37
                        }
                        var v73 int8 = 127
                        _ = v43
                        _ = v73
                    }
                default:
                    for fi74 := 0; fi74 < 8; fi74++ {
                        var v75 int8 = 0
                        var v76 int32 = -962
                        _ = fi74
                        _ = v75
                        _ = v76
                    }
                }
            }
            var v77 int32 = v11
            _ = fi36
            _ = v37
            _ = v77
        }
        v1 = (v1 + v1)
    }
    var v78 int = 32767
    if ((^(v9)) != (v9 & v9)) {
        fmt.Println(v78)
    }
    var v79 string = (("a" + "hello") + "\n\t")
    var v80 uint8 = 128
    v81 := []int{-332, v1, v78}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("slicelen v81:", len(v81), v81)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v78
    _ = v79
    _ = v80
    _ = v81
    _ = v9
}
