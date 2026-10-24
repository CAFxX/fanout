package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int }
type GenS2 struct { F0 int32; F1 uint; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool }
func genFn0(p0 string) uint32 {
    for fi6 := 0; fi6 < 17; fi6++ {
        fmt.Println(p0, fi6)
        p0 = p0
        p0 = (((p0 + "a") + "") + (p0 + ""))
        _ = fi6
    }
    var v7 int = 0
    var v8 int = (^(v7))
    v9 := []int{(^(v7)), 128, (v7 & v7), -823}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 float32 = 0.1
    var v2 int8 = -102
    var v3 uint64 = 941
    var v4 int = -421
    var v5 uint64 = (v3 & (^(v3)))
    fmt.Println(v3, v1, v1)
    v10 := GenS1{F0: 18446744073709551615, F1: 1, F2: v4}
    v11 := GenS1{F0: 424, F1: (-(v2)), F2: (v4 + (-(v4)))}
    v4 += 0
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int16 = -1
        v2 = (v2 | 127)
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 2, F1: (v2 & (v2 / (((v2 / (((v2 / ((-115) | 1))) | 1))) | 1))), F2: v4}
    v4 = (^(v4))
    v2 = (v2 / ((v2) | 1))
    var v15 int8 = (^(v2))
    v16 := GenS2{F0: 2147483647, F1: 0, F2: 146}
    var v17 uint = 1
    var v18 float64 = 1e10
    for fi19 := 0; fi19 < 16; fi19++ {
        if ((-1 >= 0) || ((v17 == (^(v17))) && false)) {
            var v20 int16 = -32768
            if (!(false || (!(v20 != v20)))) {
                v21 := GenS2{F0: 2147483647, F1: 4294967295, F2: -22}
                var v22 uint16 = 2
                _ = v21
                _ = v22
            } else {
                var v23 float32 = (v1 * 1e10)
                for fi24 := 0; fi24 < 8; fi24++ {
                    switch fi24 {
                    case -2147483648:
                        v15 += (v2 | 82)
                    case -186:
                        var v25 string = "hello"
                        _ = v25
                    case 2147483647:
                        if (255 < 617) {
                            if (473 != v5) {
                                if false {
                                    if false {
                                        fmt.Println(v20, v20, v17)
                                        v3 = v5
                                    } else {
                                        switch "hello" {
                                        case "foo bar":
                                            v17 = 65535
                                        case "\n\t":
                                            v18 -= v18
                                        case "a":
                                            switch v4 {
                                            case -2147483648:
                                                switch 127 {
                                                case -2147483648:
                                                    v26 := []int{v4, 42}
                                                    _ = v26
                                                case 128:
                                                    fmt.Println(fi24, v15)
                                                case 1:
                                                    if true {
                                                        var v27 bool = true
                                                        v3 = (^(v3))
                                                        _ = v27
                                                    } else {
                                                        if (true || true) {
                                                            switch (fi19 / ((-113) | 1)) {
                                                            case 32767:
                                                                v28 := []int{(fi24 * (fi24 / (((fi24 / ((fi24) | 1))) | 1)))}
                                                                _ = v28
                                                            default:
                                                                v29 := GenS1{F0: v17, F1: v15, F2: v4}
                                                                _ = v29
                                                            }
                                                            fmt.Println(v3, v23, v23)
                                                            v23 = float32(v18)
                                                            var v30 int = fi24
                                                            _ = v30
                                                        }
                                                    }
                                                }
                                            case 1:
                                                var v31 int8 = (-(v15))
                                                _ = v31
                                            }
                                        }
                                        if (v2 < v2) {
                                            if false {
                                                switch (("foo bar" + "\n\t") + ("foo bar" + "a")) {
                                                case "foo bar":
                                                    var v32 uint64 = (v3 - 2)
                                                    _ = v32
                                                case "a":
                                                    var v33 bool = (((fi19 - 881) < (v4 & -788)) || (!true))
                                                    _ = v33
                                                case "\n\t":
                                                    var v34 float64 = 0.0
                                                    _ = v34
                                                default:
                                                    var v35 uint = (v17 - v17)
                                                    _ = v35
                                                }
                                                if true {
                                                    v36 := GenS3{F0: v20}
                                                    v37 := GenS4{F0: 32767, F1: 42}
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    var v38 int = fi24
                                                    v39 := GenS2{F0: -461, F1: (v17 >> (((v17 & v17)) & 63)), F2: 42}
                                                    _ = v38
                                                    _ = v39
                                                }
                                                var v40 uint = 280
                                                var v41 int64 = -543
                                                v42 := []int{(fi24 & (v4 / (((v4 * (v4 & v4))) | 1)))}
                                                if true {
                                                    v43 := GenS4{F0: 9223372036854775807, F1: 241}
                                                    v3 += v3
                                                    _ = v43
                                                } else {
                                                    var v44 uint16 = 900
                                                    _ = v44
                                                }
                                                var v45 int = fi24
                                                if (1 >= 32767) {
                                                    var v46 uint = v17
                                                    if true {
                                                        v3 = v5
                                                    }
                                                    var v47 int32 = -1
                                                    _ = v46
                                                    _ = v47
                                                }
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                                _ = v45
                                            } else {
                                                v48 := []int{}
                                                v3 += (v5 * v3)
                                                _ = v48
                                            }
                                            v49 := GenS1{F0: v17, F1: (v15 % (((v2 >> ((v17) & 63))) | 1)), F2: 32767}
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v50 uint64 = v5
                                                _ = v50
                                            case "\n\t":
                                                v3 -= 669
                                            case "a":
                                                v51 := []int{}
                                                _ = v51
                                            }
                                            fmt.Println(v15)
                                            _ = v49
                                        }
                                    }
                                    var v52 bool = ((^(v20)) >= v20)
                                    fmt.Println(v2)
                                    _ = v52
                                } else {
                                    switch ("" + "foo bar") {
                                    case "foo bar":
                                        v20 = (v20 & (v20 << ((v17) & 63)))
                                    case "\n\t":
                                        v53 := GenS1{F0: 4294967295, F1: 42, F2: fi19}
                                        _ = v53
                                    case "":
                                        var v54 uint8 = 255
                                        _ = v54
                                    }
                                    fmt.Println(v4, v5, v5)
                                    v2 -= v2
                                }
                            }
                        } else {
                            v55 := []int{fi19, (fi24 >> ((v17) & 63)), 128, int(v17)}
                            var v56 float32 = v23
                            _ = v55
                            _ = v56
                        }
                    default:
                        v23 -= v23
                    }
                    switch int(v17) {
                    case 2147483647:
                        v2 -= v15
                    case 0:
                        v23 -= (v23 * 3.14159)
                    default:
                        var v57 uint64 = v3
                        _ = v57
                    }
                    v58 := []int{}
                    if false {
                        v59 := GenS3{F0: (^(v20))}
                        _ = v59
                    } else {
                        var v60 bool = (v18 > v18)
                        var v61 int8 = 100
                        v62 := GenS5{F0: 572}
                        if ((1 == 2147483647) || v60) {
                            switch "" {
                            case "foo bar":
                                var v63 uint64 = (^(v5))
                                _ = v63
                            case "hello":
                                var v64 int32 = 2147483647
                                _ = v64
                            case "\n\t":
                                var v65 uint32 = genFn0("foo bar")
                                _ = v65
                            }
                            fmt.Println(v3, v61)
                            fmt.Println(v1, v2)
                            v58 = append(v58, (-(fi24)))
                            fmt.Println(v4, v3, v60)
                        } else {
                            v20 = 42
                        }
                        fmt.Println(v18, v15, v4)
                        _ = v60
                        _ = v61
                        _ = v62
                    }
                    _ = fi24
                    _ = v58
                }
                v66 := GenS1{F0: (v17 % (((^(v17))) | 1)), F1: int8(v5), F2: v4}
                _ = v23
                _ = v66
            }
            _ = v20
        } else {
            fmt.Println(fi19)
        }
        fmt.Println(v17)
        _ = fi19
    }
    switch (("foo bar" + "\n\t") + "a") {
    case "\n\t":
        v67 := GenS6{F0: (true && false)}
        _ = v67
    case "hello":
        v5 = v3
    case "":
        var v68 int64 = 9223372036854775807
        _ = v68
    default:
        v69 := []int{(v4 ^ v4), (v4 >> (((v17 << ((4294967295) & 63))) & 63)), 128, (v4 + 32768)}
        _ = v69
    }
    v70 := GenS1{F0: (v17 - (^(v17))), F1: v2, F2: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
}
