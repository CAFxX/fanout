package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 int; F2 string }
func genFn0() uint8 {
    var v7 uint64 = 106
    var v8 uint = 219
    if false {
        fmt.Println(v8, v8)
        v7 = 0
    }
    var v9 uint64 = (v7 * v7)
    v9 -= 42
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 1
    var v3 uint16 = 1
    var v4 int64 = -1
    var v5 string = "a"
    var v6 uint64 = 0
    v6 = v6
    v3 = (v3 + v3)
    if (!false) {
        if (true || false) {
            switch "\n\t" {
            case "a":
                v1[42] = (2147483647 > 42)
            case "foo bar":
                if false {
                    var v10 uint64 = uint64(v2)
                    for fi11 := 0; fi11 < 13; fi11++ {
                        delete(v1, 65535)
                        if true {
                            v12 := v1[720]
                            switch fi11 {
                            case -1:
                                var v13 int8 = 0
                                _ = v13
                            default:
                                v2 += v2
                            }
                            v6 += v10
                            _ = v12
                        } else {
                            switch (v5 + (v5 + v5)) {
                            case "hello":
                                var v14 uint16 = (v3 & (v3 | (v3 / ((v3) | 1))))
                                _ = v14
                            case "":
                                var v15 uint = 0
                                _ = v15
                            case "foo bar":
                                _ = len(v1)
                            }
                            v4 += 1
                        }
                        _ = fi11
                    }
                    _ = v10
                } else {
                    v6 -= uint64(v3)
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v4 = (v4 + v4)
                        if true {
                            var v17 uint16 = (v3 - (v3 << ((229) & 63)))
                            if (((v4 & 9223372036854775807) < (v4 >> ((826) & 63))) && true) {
                                var v18 uint32 = 2
                                var v19 float64 = ((((-2.25 / 1.0) + (0.0 + ((1.5 / ((1.5 * -2.25) - float64(3.14159))) - 1e10))) + (1e10 + 0.1)) / 0.1)
                                v20 := GenS1{F0: 2147483648}
                                _ = v18
                                _ = v19
                                _ = v20
                            } else {
                                var v21 bool = false
                                _ = v21
                            }
                            for fi22 := 0; fi22 < 13; fi22++ {
                                var v23 float32 = -2.25
                                v2 = genFn0()
                                _ = fi22
                                _ = v23
                            }
                            _ = v17
                        }
                        v24 := []int{fi16, 32767}
                        _ = fi16
                        _ = v24
                    }
                }
            case "\n\t":
                switch v5 {
                case "":
                    if (2 == 42) {
                        var v25 float32 = 0.1
                        _ = v25
                    }
                case "foo bar":
                    v26 := []int{128, 42, 4294967295}
                    _ = v26
                default:
                    if false {
                        var v27 float64 = 3.14159
                        _ = v27
                    } else {
                        v28 := GenS2{F0: uint16(v2), F1: 286}
                        v29 := []int{}
                        if true {
                            var v30 uint = 1
                            switch 1 {
                            case -2147483648:
                                var v31 float64 = ((-2.25 - 3.14159) * (0.0 - 0.0))
                                _ = v31
                            case 2147483647:
                                if (!(false || true)) {
                                    v30 -= (^(v30))
                                    v32 := v1[635]
                                    if true {
                                        v5 = v5
                                    } else {
                                        v32 = v32
                                    }
                                    _ = v32
                                } else {
                                    var v33 int = 4294967295
                                    if len(v29) >= 2 {
                                        v29 = v29[1:len(v29)]
                                    }
                                    _ = v33
                                }
                            case 88:
                                if (false || true) {
                                    switch 127 {
                                    case 32767:
                                        v4 += v4
                                    default:
                                        v3 -= v3
                                    }
                                    for fi34 := 0; fi34 < 3; fi34++ {
                                        for fi35 := 0; fi35 < 4; fi35++ {
                                            v6 = v6
                                            fmt.Println(v6, fi34, v30)
                                            v3 += v3
                                            _ = fi35
                                        }
                                        v36 := GenS1{F0: 4294967295}
                                        _ = fi34
                                        _ = v36
                                    }
                                }
                            }
                            var v37 uint32 = 4294967295
                            _ = v30
                            _ = v37
                        }
                        _ = v28
                        _ = v29
                    }
                }
            default:
                for fi38 := 0; fi38 < 10; fi38++ {
                    switch (fi38 << ((65535) & 63)) {
                    case 526:
                        v6 = v6
                    }
                    var v39 float32 = -2.25
                    _ = fi38
                    _ = v39
                }
            }
        }
        var v40 int32 = -2147483648
        _ = v40
    }
    fmt.Println(v6, v6)
    v4 += (-(v4))
    if (2147483647 < -2147483648) {
        fmt.Println(v3, v5, v6)
        delete(v1, 2147483648)
    } else {
        for fi41 := 0; fi41 < 11; fi41++ {
            v4 -= 0
            if true {
                v42 := GenS2{F0: (v3 | (v3 ^ (^(v3)))), F1: (-(fi41))}
                _ = v42
            } else {
                for fi43 := 0; fi43 < 6; fi43++ {
                    v44 := GenS1{F0: 2147483648}
                    _ = fi43
                    _ = v44
                }
            }
            if true {
                for fi45 := 0; fi45 < 10; fi45++ {
                    fmt.Println(fi45)
                    if true {
                        var v46 uint64 = 1
                        _ = v46
                    }
                    _ = fi45
                }
                if true {
                    if false {
                        switch (fi41 | (fi41 * (fi41 / ((fi41) | 1)))) {
                        case -1:
                            _ = len(v1)
                        case -569:
                            v6 = (^(v6))
                        case 2147483647:
                            v47 := []int{fi41, (-(fi41))}
                            _ = v47
                        }
                        v48 := v1[0]
                        _ = v48
                    } else {
                        if true {
                            delete(v1, 2147483648)
                        }
                        v49 := GenS1{F0: 2}
                        var v50 string = v5
                        _ = v49
                        _ = v50
                    }
                } else {
                    var v51 uint8 = genFn0()
                    _ = v51
                }
            }
            _ = fi41
        }
        if ((v3 / (((^(v3))) | 1)) <= v3) {
            switch ("foo bar" + v5) {
            case "hello":
                if (true && (true || (!false))) {
                    v52 := GenS2{F0: uint16(v2), F1: 127}
                    delete(v1, uint32(v6))
                    _ = v52
                }
            case "a":
                var v53 uint = 317
                _ = v53
            default:
                v54 := GenS3{F0: 0}
                _ = v54
            }
        }
        var v55 int8 = -1
        _ = v55
    }
    if (false || ((^(v2)) != genFn0())) {
        v56 := []int{935}
        v4 -= v4
        fmt.Println(v4, v5)
        _ = v56
    }
    for fi57 := 0; fi57 < 11; fi57++ {
        for fi58 := 0; fi58 < 11; fi58++ {
            v2 += v2
            _ = fi58
        }
        if (2147483647 < 42) {
            for fi59 := 0; fi59 < 11; fi59++ {
                v60 := GenS4{F0: v5, F1: 0, F2: "foo bar"}
                v61 := []int{(fi57 / (((fi59 | -2147483648)) | 1)), fi57}
                _ = fi59
                _ = v60
                _ = v61
            }
            var v62 int64 = -472
            _ = len(v1)
            _ = v62
        } else {
            for fi63 := 0; fi63 < 4; fi63++ {
                v64 := GenS2{F0: v3, F1: int(v3)}
                _ = fi63
                _ = v64
            }
        }
        for fi65 := 0; fi65 < 14; fi65++ {
            v2 = (v2 - uint8(v3))
            _ = fi65
        }
        _ = fi57
    }
    for fi66 := 0; fi66 < 19; fi66++ {
        var v67 int16 = 1000
        var v68 int64 = int64(v3)
        _ = fi66
        _ = v67
        _ = v68
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
