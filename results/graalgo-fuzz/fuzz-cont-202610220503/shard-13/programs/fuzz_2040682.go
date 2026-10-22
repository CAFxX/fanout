package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint16; F2 string }
type GenS3 struct { F0 int8; F1 int32; F2 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 string; F1 uint16; F2 int16 }
type GenS7 struct { F0 uint8; F1 int8; F2 int32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint }
func genFn0() bool {
    var v7 int32 = 554
    var v8 int32 = v7
    fmt.Println(v8)
    switch 32768 {
    case 2147483647:
        v9 := GenS1{F0: ("a" + ("a" + ("hello" + ""))), F1: 0}
        _ = v9
    default:
        v10 := []int{9223372036854775807}
        _ = v10
    }
    _ = v7
    _ = v8
    return true
}
func genFn1() uint64 {
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 uint64 = 4294967295
        _ = fi11
        _ = v12
    }
    switch "\n\t" {
    case "":
        v13 := []int{-2147483648}
        _ = v13
    default:
        v14 := GenS2{F0: 1, F1: 32767, F2: "hello"}
        _ = v14
    }
    v15 := GenS1{F0: "hello", F1: 28}
    _ = v15
    return 2
}

func main() {
    var v1 bool = false
    var v2 int32 = 1
    v3 := make(map[string]int)
    v4 := []uint32{2, 4294967295, 40, 1}
    var v5 string = "foo bar"
    v6 := []int{1, 32767, -902, -966}
    v16 := GenS3{F0: -128, F1: -130, F2: 1}
    if (!(-1 >= 147)) {
        var v17 string = v5
        v18 := GenS4{F0: genFn1()}
        _ = v17
        _ = v18
    }
    var v19 uint32 = 2
    switch v5 {
    case "\n\t":
        v3[v5] = 633
    case "a":
        if (32768 == 128) {
            if v1 {
                v19 = v19
            }
            _ = len(v3)
        } else {
            if (32767 >= 1) {
                v20 := GenS4{F0: 170}
                v19 = v19
                _ = v20
            } else {
                var v21 uint16 = 1
                _ = v21
            }
        }
    default:
        v22 := GenS3{F0: -128, F1: (^(v2)), F2: 129}
        _ = v22
    }
    var v23 uint8 = 128
    var v24 uint64 = 42
    for fi25 := 0; fi25 < 1; fi25++ {
        v4 = append(v4, (v19 + v19))
        _ = fi25
    }
    v1 = (true || (v2 >= (^(v2))))
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 uint = 65535
        fmt.Println(v1, v2)
        _ = fi26
        _ = v27
    }
    var v28 uint16 = 32767
    for fi29 := 0; fi29 < 19; fi29++ {
        for fi30 := 0; fi30 < 17; fi30++ {
            var v31 int8 = -128
            var v32 int8 = v31
            var v33 uint = 1
            if v1 {
                switch (v5 + v5) {
                case "a":
                    if len(v4) > 0 {
                        v4[0] = (v19 ^ (^(v19)))
                    }
                case "\n\t":
                    if true {
                        switch fi29 {
                        case 9223372036854775807:
                            var v34 bool = false
                            _ = v34
                        case -2147483648:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        default:
                            if true {
                                var v35 int64 = 1
                                switch (fi29 * (fi29 % (((-(fi30))) | 1))) {
                                case 0:
                                    if (v5 != (v5 + (v5 + v5))) {
                                        v24 = v24
                                    } else {
                                        var v36 int64 = -9223372036854775808
                                        v37 := GenS5{F0: -154, F1: v28}
                                        _ = v36
                                        _ = v37
                                    }
                                    v38 := GenS3{F0: (v31 & (v32 / (((v31 % ((100) | 1))) | 1))), F1: int32(v24), F2: 1000}
                                    var v39 bool = (827 > uint64(v23))
                                    if true {
                                        var v40 uint16 = 42
                                        if (((v24 + v24) > uint64(v33)) && (((float32((float64(1.5) - 0.0)) * 0.1) / ((0.1 + -2.25) + float32(-2.25))) != (3.14159 - (0.1 / 1.0)))) {
                                            if (uint32(v24) == (v19 * (^(v19)))) {
                                                var v41 int64 = (v35 ^ v35)
                                                v19 = 893
                                                v42 := GenS6{F0: v5, F1: (v40 + (^(v28))), F2: 0}
                                                _ = v41
                                                _ = v42
                                            } else {
                                                v43 := GenS7{F0: (v23 + 255), F1: (v31 | -128), F2: v2}
                                                _ = v43
                                            }
                                            fmt.Println(v35, v39, v23)
                                        }
                                        if ((int64(v31) < int64(fi30)) || true) {
                                            if v39 {
                                                var v44 uint = 2
                                                _ = v44
                                            }
                                            delete(v3, "")
                                            var v45 int32 = v2
                                            _ = v45
                                        }
                                        _ = v40
                                    } else {
                                        var v46 int32 = v2
                                        if (v39 && true) {
                                            switch (^(fi30)) {
                                            case 787:
                                                var v47 uint64 = v24
                                                _ = v47
                                            case 490:
                                                if (!false) {
                                                    v33 = v33
                                                    var v48 int8 = (v31 & 0)
                                                    _ = v48
                                                }
                                                var v49 string = "a"
                                                _ = v49
                                            }
                                            fmt.Println(v2, v19)
                                            if v39 {
                                                var v50 int32 = v46
                                                _ = v50
                                            }
                                        }
                                        _ = v46
                                    }
                                    _ = v38
                                    _ = v39
                                case 128:
                                    if v1 {
                                        v51 := GenS8{F0: genFn1()}
                                        v23 = (v23 << ((255) & 63))
                                        v52 := GenS7{F0: (v23 * 127), F1: v31, F2: 0}
                                        _ = v51
                                        _ = v52
                                    }
                                case 415:
                                    v53 := GenS9{F0: v1, F1: v23, F2: (v19 % ((v19) | 1))}
                                    _ = v53
                                default:
                                    var v54 int8 = v31
                                    var v55 float32 = 0.1
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = v54
                                    _ = v55
                                }
                                _ = len(v4)
                                _ = v35
                            } else {
                                if (-398 != -510) {
                                    var v56 int8 = -40
                                    _ = v56
                                } else {
                                    var v57 string = ("" + (v5 + (("a" + (v5 + v5)) + v5)))
                                    v24 = (v24 + 1)
                                    _ = v57
                                }
                                switch v5 {
                                case "\n\t":
                                    var v58 int16 = 1
                                    _ = v58
                                case "hello":
                                    var v59 int = fi30
                                    _ = v59
                                }
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                        }
                    }
                }
                if false {
                    fmt.Println(v19)
                } else {
                    v4 = append(v4, (v19 & (v19 << (((^(v33))) & 63))))
                    v32 = v31
                    switch ("hello" + v5) {
                    case "foo bar":
                        v3[v5] = fi30
                    }
                }
                v60 := v3[(("" + "foo bar") + v5)]
                _ = v60
            }
            v61 := GenS10{F0: (v24 / (((v24 % ((v24) | 1))) | 1))}
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
            _ = v61
        }
        if v1 {
            v1 = v1
        } else {
            if v1 {
                var v62 int32 = v2
                var v63 int16 = 0
                _ = v62
                _ = v63
            } else {
                switch int(v24) {
                case 2147483647:
                    switch (fi29 & (^(fi29))) {
                    case 0:
                        v64 := GenS5{F0: v2, F1: v28}
                        _ = v64
                    case 127:
                        if v1 {
                            if len(v4) > 0 {
                                v4[0] = 793
                            }
                            v3["\n\t"] = int(v24)
                            if (!(!(37 < (^(v23))))) {
                                for fi65 := 0; fi65 < 19; fi65++ {
                                    v66 := GenS11{F0: 2}
                                    _ = fi65
                                    _ = v66
                                }
                                v23 = v23
                            }
                        }
                    default:
                        var v67 int32 = (v2 | v2)
                        _ = v67
                    }
                case 32767:
                    v1 = false
                }
                for fi68 := 0; fi68 < 9; fi68++ {
                    fmt.Println(v5)
                    var v69 uint16 = (v28 / (((v28 - (v28 / ((v28) | 1)))) | 1))
                    v70 := GenS8{F0: 517}
                    _ = fi68
                    _ = v69
                    _ = v70
                }
            }
        }
        fmt.Println(v19, v19, fi29)
        _ = fi29
    }
    v19 = 0
    var v71 uint32 = uint32(v28)
    v6 = append(v6, 4294967295)
    fmt.Println(v23, v2, v19)
    for fi72 := 0; fi72 < 7; fi72++ {
        v73 := GenS10{F0: 164}
        v23 = v23
        if len(v4) > 0 {
            v4[0] = 65535
        }
        _ = fi72
        _ = v73
    }
    fmt.Println(v24)
    var v74 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v71:", v71)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v74
}
