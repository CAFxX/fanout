package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 uint; F2 int64 }
type GenS5 struct { F0 uint8; F1 int32; F2 uint32 }
func genFn0() uint32 {
    var v7 int = 42
    v7 = -1
    var v8 uint16 = 2
    if (!((1 == 317) && false)) {
        var v9 int16 = 42
        if (853 != 42) {
            var v10 uint = 2
            for fi11 := 0; fi11 < 16; fi11++ {
                fmt.Println(v10, v9, fi11)
                _ = fi11
            }
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS1{F0: 689}
                _ = fi12
                _ = v13
            }
            _ = v10
        } else {
            var v14 int8 = -1
            var v15 uint16 = (v8 / (((v8 * (^(v8)))) | 1))
            _ = v14
            _ = v15
        }
        _ = v9
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := GenS1{F0: v8}
        _ = fi16
        _ = v17
    }
    _ = v7
    _ = v8
    return 2147483648
}

func main() {
    var v1 int8 = -128
    var v2 uint = 4294967295
    var v3 uint32 = 0
    var v4 uint = v2
    v5 := []string{("a" + "foo bar"), "a"}
    v6 := []uint32{}
    v4 -= (v2 ^ uint(v3))
    var v18 int32 = 42
    fmt.Println(v3, v18, v1)
    var v19 uint32 = (v3 << ((v2) & 63))
    var v20 uint32 = v3
    var v21 int = -9223372036854775808
    var v22 uint16 = 0
    var v23 int32 = int32(v19)
    fmt.Println(v18)
    fmt.Println(v21, v22, v21)
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS1{F0: v22}
        for fi26 := 0; fi26 < 19; fi26++ {
            if (v3 > v20) {
                v19 = 2147483647
                var v27 uint8 = 127
                _ = v27
            }
            v28 := GenS2{F0: 2, F1: v2}
            var v29 uint64 = 18446744073709551615
            _ = fi26
            _ = v28
            _ = v29
        }
        for fi30 := 0; fi30 < 4; fi30++ {
            switch "a" {
            case "":
                if true {
                    var v31 uint8 = 127
                    switch "foo bar" {
                    case "hello":
                        v32 := GenS2{F0: v31, F1: uint(v22)}
                        _ = v32
                    case "":
                        switch "" {
                        case "":
                            var v33 uint32 = 137
                            _ = v33
                        }
                        if true {
                            if ((-(fi30)) <= 4294967295) {
                                v18 = v18
                                if (int64(v19) > 45) {
                                    var v34 int64 = 674
                                    var v35 uint8 = (v31 << ((v4) & 63))
                                    switch v21 {
                                    case -1:
                                        switch v21 {
                                        case -715:
                                            var v36 bool = true
                                            _ = v36
                                        default:
                                            v1 -= (v1 * (v1 ^ (^(v1))))
                                        }
                                    default:
                                        v5 = append(v5, "a")
                                    }
                                    _ = v34
                                    _ = v35
                                }
                                var v37 uint32 = (v20 / ((genFn0()) | 1))
                                v20 = v3
                                v31 -= v31
                                _ = v37
                            }
                            if (v20 < 2147483647) {
                                v38 := GenS1{F0: 598}
                                _ = v38
                            } else {
                                v18 = (v18 & v23)
                                if (0.0 != 1.5) {
                                    if len(v5) > 0 {
                                        v5[0] = (("hello" + "foo bar") + ("foo bar" + "\n\t"))
                                    }
                                    var v39 float32 = 0.1
                                    v6 = append(v6, (v3 << (((v4 - 2)) & 63)))
                                    v39 = ((0.1 * v39) - 0.0)
                                    switch v21 {
                                    case 867:
                                        v21 -= (fi24 >> ((18446744073709551615) & 63))
                                    case 128:
                                        var v40 string = ""
                                        var v41 uint32 = v19
                                        if len(v6) > 0 {
                                            v6[0] = v20
                                        }
                                        _ = v40
                                        _ = v41
                                    }
                                    _ = v39
                                }
                                v1 -= v1
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        } else {
                            if (!true) {
                                switch 1 {
                                case 127:
                                    var v42 uint8 = 0
                                    if true {
                                        v19 += (v19 ^ (v3 % ((v3) | 1)))
                                    }
                                    v6 = append(v6, 2)
                                    _ = v42
                                case -2147483648:
                                    v6 = append(v6, v3)
                                }
                            } else {
                                var v43 float64 = -2.25
                                _ = v43
                            }
                        }
                        var v44 string = "\n\t"
                        v45 := GenS1{F0: v22}
                        v18 = (v23 % ((-456) | 1))
                        _ = v44
                        _ = v45
                    case "a":
                        var v46 string = ""
                        _ = v46
                    default:
                        if false {
                            switch ("hello" + "a") {
                            case "foo bar":
                                var v47 uint32 = (^(v19))
                                _ = v47
                            case "hello":
                                v48 := GenS3{F0: (v31 ^ (^(v31)))}
                                _ = v48
                            case "":
                                fmt.Println(v18)
                            }
                            _ = len(v5)
                        } else {
                            if (1.5 != 1.5) {
                                if false {
                                    v4 -= uint(v22)
                                } else {
                                    var v49 int64 = 0
                                    _ = v49
                                }
                                if ((v22 << ((v4) & 63)) > 42) {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                } else {
                                    fmt.Println(fi30, v20, v4)
                                }
                            } else {
                                switch -721 {
                                case 0:
                                    var v50 uint32 = (^(v20))
                                    var v51 uint8 = (v31 << (((^(v4))) & 63))
                                    v52 := GenS4{F0: ("hello" + "a"), F1: uint(v22), F2: -9223372036854775808}
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                case -421:
                                    v53 := GenS3{F0: 142}
                                    _ = v53
                                default:
                                    fmt.Println(v3)
                                }
                                v54 := GenS5{F0: (v31 % ((128) | 1)), F1: 2147483647, F2: v20}
                                _ = v54
                            }
                            v4 = (v4 * (v4 - v2))
                            v23 += (v23 << ((384) & 63))
                        }
                    }
                    _ = v31
                } else {
                    fmt.Println(v22)
                    if (!true) {
                        if (true || true) {
                            if false {
                                v2 = (v4 << ((v2) & 63))
                            } else {
                                if len(v5) > 0 {
                                    v5[0] = "hello"
                                }
                            }
                            v55 := GenS1{F0: (v22 % ((815) | 1))}
                            _ = v55
                        } else {
                            v20 += v20
                            var v56 float64 = 1e10
                            fmt.Println(v3)
                            _ = v56
                        }
                        v19 = v3
                    }
                    var v57 int = -1
                    _ = v57
                }
            case "a":
                var v58 uint32 = v19
                _ = v58
            case "\n\t":
                fmt.Println(v1, v18, v1)
            default:
                if (!((!true) && (float32((0.0 * (float64(1e10) * float64((0.0 * 0.1))))) < 1.5))) {
                    var v59 uint32 = v3
                    _ = v59
                } else {
                    var v60 bool = (((1000 >= -207) && ((v23 + int32(fi24)) != v18)) || false)
                    if (v18 != (v18 - v23)) {
                        v61 := GenS5{F0: 1, F1: v23, F2: v3}
                        var v62 uint = (v2 >> ((683) & 63))
                        _ = v61
                        _ = v62
                    }
                    var v63 uint8 = 42
                    _ = v60
                    _ = v63
                }
            }
            _ = fi30
        }
        _ = fi24
        _ = v25
    }
    var v64 float32 = (float32((3.14159 - (1.5 + ((-2.25 / 0.1) + 0.0)))) / 3.14159)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v65 := GenS4{F0: "\n\t", F1: v2, F2: 1}
    switch "" {
    case "\n\t":
        var v66 float32 = 3.14159
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
}
