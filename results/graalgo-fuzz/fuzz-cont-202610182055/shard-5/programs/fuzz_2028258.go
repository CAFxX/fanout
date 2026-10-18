package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int8; F2 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string }
func genFn0() int8 {
    var v6 uint8 = 76
    for fi7 := 0; fi7 < 18; fi7++ {
        v6 -= 22
        _ = fi7
    }
    v6 = (v6 & v6)
    v8 := GenS1{F0: (false || ((-128 <= 100) || true))}
    _ = v6
    _ = v8
    return 1
}
func genFn1(p0 bool, p1 string, p2 uint16) bool {
    var v9 uint32 = 2147483647
    var v10 float64 = ((3.14159 + 1e10) * 0.0)
    p1 = (p1 + p1)
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint16 = p2
        v13 := GenS1{F0: p0}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return (!(1.5 < 0.0))
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[string]bool)
    var v3 int32 = 42
    v4 := make(map[uint32]uint32)
    var v5 int = 79
    v14 := v4[844]
    var v15 bool = false
    v1 = v1
    v5 -= -135
    switch -332 {
    case -2147483648:
        var v16 float32 = -2.25
        _ = v16
    default:
        switch "foo bar" {
        case "a":
            v3 = v3
        case "":
            v17 := GenS2{F0: 0, F1: v14, F2: 127}
            _ = v17
        default:
            v18 := GenS1{F0: ((((1 < 1) || (1 < 128)) && true) || false)}
            _ = v18
        }
    }
    var v19 int32 = (v3 + v3)
    v20 := GenS1{F0: v15}
    v21 := GenS3{F0: -1, F1: ((-(v5)) != v5)}
    v15 = false
    v22 := v4[65535]
    switch "hello" {
    case "":
        v19 = -906
    case "hello":
        v14 += (v14 & v22)
    }
    switch v5 {
    case -560:
        if (277 >= 346) {
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 uint = 65535
                switch ("\n\t" + "hello") {
                case "foo bar":
                    var v25 int16 = 32767
                    _ = v25
                case "\n\t":
                    switch ("hello" + "") {
                    case "hello":
                        var v26 uint = (^(v24))
                        _ = v26
                    }
                case "":
                    v14 -= 4294967295
                }
                _ = fi23
                _ = v24
            }
            v27 := []int{}
            _ = v27
        } else {
            v28 := GenS1{F0: v15}
            var v29 uint16 = 0
            _ = v28
            _ = v29
        }
    case -1:
        if true {
            v30 := GenS2{F0: 38, F1: 2147483648, F2: 128}
            _ = v30
        } else {
            for fi31 := 0; fi31 < 15; fi31++ {
                if ("a" <= "a") {
                    var v32 int8 = -1
                    _ = v32
                }
                if v15 {
                    if v15 {
                        if v15 {
                            v33 := []int{v5, v5, fi31, (fi31 / ((v5) | 1))}
                            _ = v33
                        }
                        var v34 float32 = (3.14159 - (0.0 - v1))
                        v35 := GenS4{F0: 224, F1: 222}
                        _ = v34
                        _ = v35
                    }
                    var v36 int32 = (v3 << ((0) & 63))
                    for fi37 := 0; fi37 < 19; fi37++ {
                        var v38 int32 = 1
                        switch -1 {
                        case 460:
                            if (v14 < 2147483648) {
                                v39 := []int{0, fi31, 80}
                                v36 = (v36 + v38)
                                _ = v39
                            } else {
                                switch (v5 - (fi37 % (((fi37 * -2147483648)) | 1))) {
                                case 42:
                                    fmt.Println(v15, v19, fi37)
                                case 1:
                                    v19 += (v38 << ((694) & 63))
                                }
                                var v40 uint32 = (^(v14))
                                fmt.Println(v5, v3, v5)
                                _ = v40
                            }
                        case 128:
                            var v41 uint64 = 42
                            _ = v41
                        default:
                            var v42 uint32 = (v22 | v22)
                            _ = v42
                        }
                        v43 := GenS4{F0: 42, F1: 32767}
                        _ = fi37
                        _ = v38
                        _ = v43
                    }
                    _ = v36
                } else {
                    v44 := []int{(-(v5)), fi31, -145}
                    _ = v44
                }
                _ = fi31
            }
            v15 = v15
            if v15 {
                v5 += 9223372036854775807
            }
        }
    case -535:
        v45 := GenS2{F0: 32767, F1: (v22 | v14), F2: 246}
        _ = v45
    }
    var v46 int8 = genFn0()
    v3 -= 683
    switch ((("a" + "hello") + "foo bar") + "hello") {
    case "foo bar":
        fmt.Println(v19, v15)
    case "hello":
        v47 := GenS5{F0: v14, F1: v46, F2: (v3 ^ (v3 << ((42) & 63)))}
        _ = v47
    case "":
        if ((v46 == v46) && v15) {
            if false {
                var v48 int = v5
                if v15 {
                    v49 := GenS4{F0: 255, F1: 32767}
                    _ = v49
                }
                _ = v48
            }
        }
    default:
        v46 += v46
    }
    if (0 > 18446744073709551615) {
        if v15 {
            if (("" > "foo bar") && v15) {
                v50 := GenS6{F0: v46}
                var v51 int32 = 42
                _ = v50
                _ = v51
            } else {
                v52 := GenS2{F0: 881, F1: v14, F2: 0}
                if v15 {
                    v53 := GenS2{F0: 65535, F1: v14, F2: 128}
                    v54 := v2[("foo bar" + "hello")]
                    var v55 float64 = 1e10
                    _ = v53
                    _ = v54
                    _ = v55
                } else {
                    for fi56 := 0; fi56 < 6; fi56++ {
                        fmt.Println(v14, v22)
                        var v57 string = "\n\t"
                        _ = fi56
                        _ = v57
                    }
                    for fi58 := 0; fi58 < 15; fi58++ {
                        if (-248 == 0) {
                            if (3.14159 <= float64(((0.0 - 1e10) + v1))) {
                                var v59 uint64 = 1
                                var v60 int64 = -49
                                v14 = v14
                                _ = v59
                                _ = v60
                            } else {
                                v22 -= (v14 / ((v22) | 1))
                                var v61 float32 = v1
                                _ = v61
                            }
                            v46 = v46
                        } else {
                            fmt.Println(v46, v14)
                            v62 := GenS7{F0: v15}
                            v3 += (v3 ^ v3)
                            _ = v62
                        }
                        for fi63 := 0; fi63 < 12; fi63++ {
                            switch "foo bar" {
                            case "foo bar":
                                var v64 string = "hello"
                                _ = v64
                            }
                            switch (-(fi63)) {
                            case 32767:
                                var v65 float32 = v1
                                _ = v65
                            }
                            v66 := GenS8{F0: 18446744073709551615}
                            _ = fi63
                            _ = v66
                        }
                        var v67 int16 = 42
                        _ = fi58
                        _ = v67
                    }
                    if v15 {
                        v5 += (v5 & v5)
                        if (!(!true)) {
                            v22 += v14
                        }
                        var v68 int8 = (v46 - (v46 ^ v46))
                        _ = v68
                    } else {
                        v46 -= (v46 / (((v46 / ((int8(v14)) | 1))) | 1))
                    }
                }
                _ = v52
            }
            fmt.Println(v1, v14)
            for fi69 := 0; fi69 < 17; fi69++ {
                var v70 int64 = 0
                v3 = 2147483647
                _ = fi69
                _ = v70
            }
        } else {
            v15 = v15
        }
    } else {
        v15 = false
        v71 := GenS3{F0: (v5 / ((v5) | 1)), F1: v15}
        var v72 int32 = 137
        _ = v71
        _ = v72
    }
    switch "hello" {
    case "\n\t":
        v73 := GenS9{F0: ("foo bar" + "foo bar")}
        _ = v73
    default:
        var v74 uint32 = (v22 >> ((253) & 63))
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
