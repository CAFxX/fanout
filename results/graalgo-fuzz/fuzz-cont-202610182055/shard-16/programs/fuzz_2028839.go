package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 uint16; F1 uint; F2 int32 }
type GenS7 struct { F0 int32; F1 int16 }

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    var v4 string = "a"
    var v5 int16 = 1000
    var v6 uint64 = 2
    fmt.Println(v4, v6)
    var v7 string = (v4 + "foo bar")
    var v8 uint64 = v6
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(v4, fi9, v7)
            v1 -= 35
            v4 = "\n\t"
            if false {
                v11 := GenS1{F0: (^(v1))}
                v1 = 42
                _ = v11
            } else {
                v4 = (v4 + "foo bar")
                var v12 uint32 = 529
                _ = v12
            }
            _ = fi10
        }
        _ = fi9
    }
    if true {
        v8 += v6
        v8 += v8
        v8 = v8
    } else {
        if (0.0 >= 1.5) {
            v13 := GenS1{F0: v1}
            var v14 int = -491
            if true {
                var v15 float32 = float32((-2.25 + 1.5))
                _ = v15
            } else {
                if (!(-2147483648 > 0)) {
                    var v16 bool = false
                    _ = v16
                }
                v17 := []int{v14, (^(v14)), v14}
                var v18 int16 = int16(v1)
                _ = v17
                _ = v18
            }
            _ = v13
            _ = v14
        } else {
            var v19 int = -556
            v6 = uint64(v1)
            if (!(v19 == (v19 | (v19 * (^(v19)))))) {
                if false {
                    v20 := GenS1{F0: v1}
                    v21 := []int{v19, -297}
                    _ = v20
                    _ = v21
                } else {
                    if (true && ((65535 == 1) || ((v6 == v8) || (!(false && (128 == 128)))))) {
                        var v22 int8 = 0
                        if false {
                            v23 := []int{}
                            _ = v23
                        }
                        switch v7 {
                        case "\n\t":
                            fmt.Println(v4)
                        case "foo bar":
                            switch v7 {
                            case "":
                                v8 += (v6 - v8)
                            }
                        case "a":
                            v24 := GenS1{F0: v1}
                            _ = v24
                        default:
                            if true {
                                if (536 >= 2) {
                                    if (v5 < 1) {
                                        switch (v19 / ((v19) | 1)) {
                                        case -1:
                                            v25 := GenS2{F0: true}
                                            _ = v25
                                        default:
                                            v26 := GenS3{F0: -126, F1: (v22 | 126)}
                                            _ = v26
                                        }
                                    } else {
                                        v27 := GenS3{F0: -52, F1: (v22 * (v22 - v22))}
                                        var v28 int16 = (^(v5))
                                        for fi29 := 0; fi29 < 8; fi29++ {
                                            if false {
                                                v28 += (v5 - v5)
                                            } else {
                                                if false {
                                                    var v30 int = -2147483648
                                                    _ = v30
                                                }
                                                var v31 int8 = 0
                                                _ = v31
                                            }
                                            v32 := []int{(fi29 % ((301) | 1)), v19}
                                            var v33 bool = true
                                            _ = fi29
                                            _ = v32
                                            _ = v33
                                        }
                                        _ = v27
                                        _ = v28
                                    }
                                    var v34 bool = false
                                    _ = v34
                                }
                                switch v19 {
                                case 128:
                                    var v35 int32 = 1
                                    _ = v35
                                case 4294967295:
                                    if (39 != 128) {
                                        var v36 float64 = float64(-2.25)
                                        var v37 int8 = 45
                                        _ = v36
                                        _ = v37
                                    } else {
                                        switch v7 {
                                        case "foo bar":
                                            for fi38 := 0; fi38 < 14; fi38++ {
                                                var v39 float64 = -2.25
                                                _ = fi38
                                                _ = v39
                                            }
                                        case "hello":
                                            var v40 uint16 = 42
                                            _ = v40
                                        default:
                                            var v41 float64 = 3.14159
                                            _ = v41
                                        }
                                    }
                                case -9223372036854775808:
                                    v7 = v4
                                default:
                                    var v42 float64 = ((1.5 - (0.1 * 0.0)) - (-2.25 + 0.0))
                                    _ = v42
                                }
                            }
                        }
                        _ = v22
                    }
                    for fi43 := 0; fi43 < 18; fi43++ {
                        if true {
                            v44 := GenS4{F0: 42}
                            _ = v44
                        }
                        v8 = v8
                        _ = fi43
                    }
                }
                v19 += v19
            } else {
                var v45 float32 = 0.0
                _ = v45
            }
            _ = v19
        }
        var v46 bool = (-1 == -651)
        switch (v7 + v7) {
        case "\n\t":
            if false {
                var v47 int = -2147483648
                if false {
                    for fi48 := 0; fi48 < 18; fi48++ {
                        v49 := v3[2147483648]
                        fmt.Println(v6, v4)
                        _ = fi48
                        _ = v49
                    }
                }
                if (-613 <= v47) {
                    for fi50 := 0; fi50 < 11; fi50++ {
                        if v46 {
                            var v51 int64 = -1
                            v52 := GenS3{F0: 1, F1: 42}
                            _ = v51
                            _ = v52
                        } else {
                            if ((!v46) || (v46 || v46)) {
                                var v53 uint = 13
                                if v46 {
                                    switch fi50 {
                                    case 1:
                                        var v54 bool = v46
                                        _ = v54
                                    case -9223372036854775808:
                                        if (true && v46) {
                                            var v55 int = v47
                                            var v56 float32 = 0.1
                                            v8 = v8
                                            _ = v55
                                            _ = v56
                                        }
                                    default:
                                        v53 = 18446744073709551615
                                    }
                                    v57 := GenS4{F0: 9223372036854775807}
                                    var v58 uint32 = 2
                                    _ = v57
                                    _ = v58
                                }
                                if (1 < 1) {
                                    fmt.Println(v53, v7, v53)
                                    v59 := GenS1{F0: (v1 / ((v53) | 1))}
                                    var v60 int32 = 678
                                    _ = v59
                                    _ = v60
                                } else {
                                    if ((!v46) || true) {
                                        if (v7 != v4) {
                                            var v61 uint32 = 65535
                                            v62 := GenS5{F0: v53, F1: -128}
                                            _ = v61
                                            _ = v62
                                        }
                                        v63 := []int{(fi50 ^ fi50), int(v1)}
                                        _ = v63
                                    }
                                    if false {
                                        v64 := GenS6{F0: 32767, F1: (v1 * (v53 | (v1 - (v53 ^ (v53 << ((v53) & 63)))))), F2: -2147483648}
                                        _ = v64
                                    }
                                }
                                _ = v53
                            }
                            if v46 {
                                var v65 uint = 4294967295
                                _ = v65
                            } else {
                                for fi66 := 0; fi66 < 13; fi66++ {
                                    v1 -= 186
                                    _ = fi66
                                }
                            }
                            v4 = (("" + ((v4 + v4) + ("foo bar" + "\n\t"))) + "\n\t")
                        }
                        v3[1] = fi50
                        _ = fi50
                    }
                    v7 = v4
                }
                _ = v47
            }
        }
        _ = v46
    }
    fmt.Println(v5)
    if true {
        switch (v7 + (v7 + v7)) {
        case "":
            _ = len(v3)
        case "\n\t":
            var v67 uint16 = 326
            _ = v67
        case "foo bar":
            fmt.Println(v7)
        default:
            if (4294967295 == v8) {
                v68 := GenS7{F0: 1, F1: v5}
                if true {
                    var v69 int16 = v5
                    _ = v69
                }
                _ = v68
            }
        }
    }
    var v70 uint = 65535
    v71 := v3[4294967295]
    var v72 int = (-(v71))
    if false {
        if true {
            var v73 uint8 = 189
            v2[1] = 53
            fmt.Println(v5, v70)
            _ = v73
        }
    } else {
        for fi74 := 0; fi74 < 5; fi74++ {
            if true {
                var v75 int8 = -1
                v76 := GenS6{F0: 0, F1: (^(v1)), F2: -623}
                _ = v75
                _ = v76
            } else {
                fmt.Println(fi74)
            }
            switch v4 {
            case "foo bar":
                if (3.14159 == ((1.5 * (float32(0.0) - -2.25)) + float32(1.5))) {
                    v6 -= v6
                    v77 := GenS4{F0: 9223372036854775807}
                    _ = v77
                } else {
                    v78 := GenS6{F0: 428, F1: (v70 % ((65535) | 1)), F2: 516}
                    v2[714] = 933
                    _ = v78
                }
            case "hello":
                if true {
                    v79 := []int{v72, v71, v72}
                    _ = v79
                }
            default:
                var v80 string = v4
                _ = v80
            }
            _ = fi74
        }
    }
    for fi81 := 0; fi81 < 7; fi81++ {
        var v82 uint = 18446744073709551615
        v72 = -555
        if (!((!(!false)) || (0 != 2))) {
            if (0.0 < -2.25) {
                v4 = v4
                v8 = v6
            }
        }
        _ = fi81
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v70
    _ = v71
    _ = v72
    _ = v8
}
