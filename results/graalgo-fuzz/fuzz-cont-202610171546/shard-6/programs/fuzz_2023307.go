package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 string }

func main() {
    var v1 float32 = (1.5 + (1.5 + ((1.5 - (((0.1 / 3.14159) / 1e10) + 0.0)) / 1e10)))
    var v2 string = "foo bar"
    var v3 int8 = 1
    if false {
        v4 := GenS1{F0: 18446744073709551615}
        switch 0 {
        case 670:
            v5 := GenS2{F0: v3}
            _ = v5
        case 0:
            var v6 int32 = 2147483647
            _ = v6
        case 155:
            v3 = v3
        }
        _ = v4
    } else {
        switch (v2 + "") {
        case "a":
            if false {
                if (false && false) {
                    for fi7 := 0; fi7 < 12; fi7++ {
                        var v8 int32 = 0
                        var v9 uint8 = 2
                        v10 := []int{(fi7 | fi7), (^(fi7))}
                        _ = fi7
                        _ = v10
                        _ = v8
                        _ = v9
                    }
                    v1 = v1
                }
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 int16 = 1
                    var v13 int8 = 127
                    if (true || false) {
                        for fi14 := 0; fi14 < 7; fi14++ {
                            var v15 uint = 255
                            v15 += v15
                            var v16 int8 = (v3 ^ 42)
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        var v17 float64 = (float64(0.1) / 0.1)
                        v18 := []int{fi11, fi11, fi11, fi11}
                        _ = v17
                        _ = v18
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                if (!(65535 <= 993)) {
                    v19 := GenS3{F0: 2}
                    v20 := GenS3{F0: 11}
                    _ = v19
                    _ = v20
                }
            } else {
                var v21 uint16 = 571
                _ = v21
            }
        default:
            v22 := GenS1{F0: 255}
            _ = v22
        }
        v2 = "\n\t"
    }
    v23 := []int{-9223372036854775808, 32768, -421}
    for fi24 := 0; fi24 < 18; fi24++ {
        for fi25 := 0; fi25 < 4; fi25++ {
            v1 -= v1
            v2 = v2
            switch (fi25 - (fi24 % (((fi25 - (fi24 * fi24))) | 1))) {
            case 153:
                v1 = v1
            case -307:
                if true {
                    v3 = v3
                }
            case -731:
                if (!(1 <= 2)) {
                    switch (((v2 + v2) + v2) + v2) {
                    case "\n\t":
                        if true {
                            _ = len(v23)
                        }
                    case "foo bar":
                        var v26 uint32 = 2147483647
                        _ = v26
                    case "hello":
                        var v27 int8 = v3
                        _ = v27
                    }
                    switch 2147483647 {
                    case 32767:
                        var v28 string = ((v2 + (("foo bar" + "foo bar") + "\n\t")) + v2)
                        _ = v28
                    default:
                        var v29 bool = false
                        _ = v29
                    }
                    var v30 uint16 = 913
                    _ = v30
                } else {
                    if (65535 < 42) {
                        if true {
                            v31 := GenS4{F0: false, F1: 401}
                            _ = v31
                        }
                        switch v2 {
                        case "foo bar":
                            if len(v23) > 0 {
                                v23[0] = fi24
                            }
                        case "hello":
                            v32 := GenS1{F0: 1}
                            _ = v32
                        case "a":
                            v3 -= 42
                        default:
                            _ = len(v23)
                        }
                        if true {
                            if (v2 <= v2) {
                                v1 += v1
                                var v33 string = v2
                                v34 := GenS4{F0: true, F1: 197}
                                if true {
                                    if true {
                                        var v35 int32 = int32(v3)
                                        var v36 int = 2147483647
                                        if false {
                                            v36 += fi25
                                            v23 = append(v23, (fi25 | fi24))
                                        } else {
                                            if true {
                                                var v37 int32 = (-(v35))
                                                if len(v23) > 0 {
                                                    v23[0] = fi24
                                                }
                                                if true {
                                                    v38 := GenS4{F0: ((true || false) && true), F1: 4294967295}
                                                    v39 := GenS5{F0: 302, F1: 135}
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    _ = len(v23)
                                                    v40 := GenS6{F0: "a"}
                                                    v41 := GenS3{F0: 821}
                                                    _ = v40
                                                    _ = v41
                                                }
                                                var v42 int8 = (^(v3))
                                                _ = v37
                                                _ = v42
                                            } else {
                                                v43 := GenS6{F0: "\n\t"}
                                                _ = v43
                                            }
                                            var v44 float32 = (v1 - (v1 * (0.1 + v1)))
                                            v45 := GenS4{F0: true, F1: 610}
                                            _ = v44
                                            _ = v45
                                        }
                                        _ = v35
                                        _ = v36
                                    }
                                    var v46 bool = ((2147483647 < 65535) || true)
                                    _ = v46
                                }
                                _ = v33
                                _ = v34
                            } else {
                                var v47 bool = (0 < -295)
                                _ = v47
                            }
                        }
                    }
                    if len(v23) > 0 {
                        v23[0] = -373
                    }
                }
            default:
                if len(v23) > 0 {
                    v23[0] = -239
                }
            }
            _ = fi25
        }
        if (1000 >= int16(v3)) {
            _ = len(v23)
            var v48 int16 = 32767
            _ = v48
        }
        _ = fi24
    }
    if true {
        if false {
            var v49 int16 = 128
            var v50 int32 = -1
            if (true && false) {
                for fi51 := 0; fi51 < 1; fi51++ {
                    v3 = int8(v49)
                    _ = fi51
                }
            }
            _ = v49
            _ = v50
        }
        var v52 int16 = -545
        _ = v52
    } else {
        var v53 int8 = (v3 % ((v3) | 1))
        switch 359 {
        case -812:
            var v54 string = v2
            _ = v54
        case -548:
            var v55 int8 = (^(v53))
            _ = v55
        default:
            for fi56 := 0; fi56 < 9; fi56++ {
                var v57 uint = 0
                var v58 bool = false
                _ = fi56
                _ = v57
                _ = v58
            }
        }
        if (!(true && (255 == 2))) {
            fmt.Println(v2)
            for fi59 := 0; fi59 < 3; fi59++ {
                if (42 <= 4294967295) {
                    var v60 int = 0
                    fmt.Println(fi59, v2, v3)
                    _ = v60
                } else {
                    var v61 uint16 = 122
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                    _ = v61
                }
                _ = fi59
            }
            var v62 int16 = 1
            _ = v62
        } else {
            var v63 int16 = -32768
            _ = v63
        }
        _ = v53
    }
    var v64 int = -1
    var v65 string = (v2 + v2)
    var v66 int = v64
    v67 := GenS6{F0: v2}
    for fi68 := 0; fi68 < 19; fi68++ {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        _ = fi68
    }
    var v69 int32 = -2147483648
    v70 := GenS2{F0: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v69
    _ = v70
}
