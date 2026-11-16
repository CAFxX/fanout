package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 bool; F2 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 uint16; F2 int16 }
type GenS8 struct { F0 int32; F1 uint; F2 int16 }
type GenS9 struct { F0 uint; F1 string; F2 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 42
    var v3 int32 = -2147483648
    var v4 int64 = -9223372036854775808
    var v5 uint8 = 96
    v6 := GenS1{F0: 151, F1: v4}
    var v7 uint64 = 2
    v5 -= (v5 ^ uint8(v7))
    if false {
        switch "" {
        case "":
            _ = len(v1)
        case "foo bar":
            v8 := []int{1, 128, int(v5)}
            _ = v8
        case "a":
            switch "foo bar" {
            case "a":
                v9 := GenS1{F0: v5, F1: v4}
                _ = v9
            case "":
                v5 = 255
            }
        }
        for fi10 := 0; fi10 < 6; fi10++ {
            switch -698 {
            case 32768:
                switch -501 {
                case -1:
                    var v11 bool = true
                    _ = v11
                case -576:
                    v1[((("foo bar" + "hello") + "hello") + ("" + ("foo bar" + "hello")))] = 748
                case 128:
                    for fi12 := 0; fi12 < 15; fi12++ {
                        if false {
                            v13 := GenS1{F0: (v5 % ((v5) | 1)), F1: int64(fi12)}
                            var v14 int64 = (v4 - (v4 + -1))
                            var v15 int = fi10
                            switch fi10 {
                            case -9223372036854775808:
                                fmt.Println(v7)
                            case 521:
                                v16 := GenS1{F0: (v5 % (((^(v5))) | 1)), F1: (-(v14))}
                                _ = v16
                            case -1:
                                switch "foo bar" {
                                case "\n\t":
                                    if (v7 >= 672) {
                                        switch (fi10 * (v15 / (((fi12 | v15)) | 1))) {
                                        case 128:
                                            v17 := GenS1{F0: v5, F1: 9223372036854775807}
                                            _ = v17
                                        case 42:
                                            switch -1 {
                                            case 365:
                                                v14 = int64(fi12)
                                            default:
                                                switch -407 {
                                                case 0:
                                                    v14 += v4
                                                case 936:
                                                    switch (fi12 % (((-(fi12))) | 1)) {
                                                    case -1:
                                                        var v18 float64 = ((1.5 / (1e10 - ((float64(3.14159) - float64(0.1)) / (1e10 + float64(((1e10 / 0.1) - (1.5 / -2.25))))))) - (-2.25 - 0.1))
                                                        _ = v18
                                                    default:
                                                        if false {
                                                            switch 2147483647 {
                                                            case 1:
                                                                v19 := GenS2{F0: 42, F1: 0}
                                                                _ = v19
                                                            default:
                                                                if (fi10 != 127) {
                                                                    var v20 uint = 65535
                                                                    if false {
                                                                        var v21 uint32 = 531
                                                                        fmt.Println(v15, v14)
                                                                        _ = v21
                                                                    } else {
                                                                        if false {
                                                                            v22 := GenS1{F0: 127, F1: v14}
                                                                            _ = v22
                                                                        } else {
                                                                            v15 = fi12
                                                                        }
                                                                        var v23 float32 = 3.14159
                                                                        _ = v23
                                                                    }
                                                                    _ = v20
                                                                }
                                                            }
                                                            v24 := GenS3{F0: 139}
                                                            v15 = fi10
                                                            v3 = (v3 * v3)
                                                            _ = v24
                                                        }
                                                    }
                                                case -502:
                                                    v14 = 42
                                                default:
                                                    v25 := GenS1{F0: (v5 | v5), F1: (v14 >> ((42) & 63))}
                                                    _ = v25
                                                }
                                            }
                                        default:
                                            var v26 int = fi12
                                            _ = v26
                                        }
                                        v4 += (v4 << ((1) & 63))
                                    } else {
                                        if (42 < -128) {
                                            v1["a"] = 0
                                            v27 := []int{(fi10 >> ((42) & 63)), (v15 ^ 127), fi10, fi10}
                                            v28 := GenS1{F0: (v5 ^ (v5 & 1)), F1: (v4 * v14)}
                                            _ = v27
                                            _ = v28
                                        } else {
                                            var v29 int32 = v3
                                            v15 = 9223372036854775807
                                            _ = v29
                                        }
                                        v4 = (v4 / (((v14 * (v14 ^ 42))) | 1))
                                        v30 := []int{v15, (fi10 * fi10), fi10}
                                        v31 := GenS4{F0: v4}
                                        switch (v15 + (fi10 / ((v15) | 1))) {
                                        case 2147483647:
                                            v32 := GenS4{F0: v14}
                                            _ = v32
                                        case 191:
                                            var v33 int = fi10
                                            _ = v33
                                        case -199:
                                            v3 += v3
                                        default:
                                            _ = len(v1)
                                        }
                                        _ = v30
                                        _ = v31
                                    }
                                case "":
                                    v34 := GenS5{F0: true, F1: false, F2: v3}
                                    _ = v34
                                case "a":
                                    v35 := GenS6{F0: (-(v14))}
                                    _ = v35
                                default:
                                    var v36 uint16 = 65535
                                    var v37 int32 = (v3 + (^(v3)))
                                    v38 := GenS2{F0: (v2 & v2), F1: 5}
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                            }
                            switch v15 {
                            case 541:
                                v39 := []int{fi12, (v15 >> ((2) & 63))}
                                _ = v39
                            default:
                                if true {
                                    switch "hello" {
                                    case "a":
                                        var v40 float64 = 1.5
                                        _ = v40
                                    default:
                                        v41 := GenS7{F0: (^(v14)), F1: 255, F2: (-(v2))}
                                        _ = v41
                                    }
                                    var v42 uint32 = 0
                                    _ = v42
                                }
                            }
                            v43 := v1["\n\t"]
                            _ = v13
                            _ = v14
                            _ = v15
                            _ = v43
                        }
                        _ = fi12
                    }
                default:
                    var v44 int8 = -128
                    _ = v44
                }
            default:
                for fi45 := 0; fi45 < 7; fi45++ {
                    var v46 string = ("\n\t" + "")
                    if true {
                        var v47 int8 = -1
                        _ = v47
                    } else {
                        var v48 uint16 = 826
                        v49 := []int{-378}
                        _ = v48
                        _ = v49
                    }
                    var v50 int8 = 27
                    v51 := GenS8{F0: -2147483648, F1: 82, F2: v2}
                    v3 = (v3 << ((2) & 63))
                    var v52 int32 = (v3 << ((752) & 63))
                    switch "\n\t" {
                    case "\n\t":
                        var v53 string = ("a" + v46)
                        _ = v53
                    case "hello":
                        var v54 int64 = (-(v4))
                        var v55 int16 = (v2 | 0)
                        v5 += v5
                        v3 = v52
                        _ = v54
                        _ = v55
                    default:
                        var v56 float32 = 1.5
                        _ = v56
                    }
                    _ = fi45
                    _ = v46
                    _ = v50
                    _ = v51
                    _ = v52
                }
            }
            _ = fi10
        }
        var v57 float64 = float64(1.5)
        _ = v57
    }
    v58 := GenS9{F0: 42, F1: "hello", F2: 1}
    v5 = v5
    v59 := GenS7{F0: v4, F1: 1, F2: v2}
    var v60 int8 = 59
    if true {
        v61 := v1[("" + ("a" + "hello"))]
        switch int(v60) {
        case 330:
            var v62 string = ""
            _ = v62
        case 42:
            for fi63 := 0; fi63 < 9; fi63++ {
                if (true && (true || (((1e10 * 1e10) <= -2.25) || true))) {
                    _ = len(v1)
                    v2 = (v2 | (v2 << ((922) & 63)))
                }
                _ = fi63
            }
        case -2147483648:
            v64 := []int{int(v2), 9223372036854775807}
            _ = v64
        default:
            for fi65 := 0; fi65 < 16; fi65++ {
                var v66 uint = 42
                v1["foo bar"] = v61
                _ = fi65
                _ = v66
            }
        }
        _ = v61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v7
}
