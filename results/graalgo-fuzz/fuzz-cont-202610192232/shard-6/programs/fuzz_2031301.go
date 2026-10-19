package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int32; F1 int; F2 bool }
type GenS4 struct { F0 uint32; F1 bool }
func genFn0() uint {
    var v6 int16 = 42
    switch "foo bar" {
    case "a":
        v6 += (v6 | (v6 << ((65535) & 63)))
    }
    switch -2147483648 {
    case -1:
        var v7 uint32 = 539
        _ = v7
    case -9223372036854775808:
        v8 := []int{9223372036854775807}
        _ = v8
    case -2147483648:
        v9 := GenS1{F0: false, F1: 255}
        _ = v9
    default:
        v6 -= 1000
    }
    switch "" {
    case "":
        var v10 int16 = v6
        _ = v10
    }
    _ = v6
    return 1
}

func main() {
    var v1 int64 = -1
    var v2 uint32 = 2147483648
    var v3 float64 = (0.0 * (1e10 + (3.14159 * float64(0.0))))
    var v4 int32 = -2147483648
    var v5 uint64 = 0
    v3 = (-2.25 / v3)
    switch -9223372036854775808 {
    case -639:
        if (!(false && true)) {
            v11 := []int{2147483647, 128, -939}
            fmt.Println(v3, v5, v5)
            _ = v11
        } else {
            switch 9223372036854775807 {
            case -976:
                var v12 int32 = v4
                _ = v12
            case 42:
                v13 := GenS1{F0: true, F1: 2}
                _ = v13
            default:
                var v14 int16 = -184
                _ = v14
            }
            var v15 uint32 = 42
            for fi16 := 0; fi16 < 12; fi16++ {
                switch -27 {
                case -705:
                    for fi17 := 0; fi17 < 20; fi17++ {
                        if false {
                            var v18 float64 = v3
                            var v19 int8 = 100
                            _ = v18
                            _ = v19
                        }
                        _ = fi17
                    }
                default:
                    if false {
                        var v20 uint64 = (^(v5))
                        _ = v20
                    } else {
                        v1 = v1
                        v21 := GenS2{F0: "foo bar", F1: v15}
                        v22 := []int{fi16, int(v2)}
                        _ = v21
                        _ = v22
                    }
                }
                _ = fi16
            }
            _ = v15
        }
    case 1:
        if (56 == -20) {
            for fi23 := 0; fi23 < 3; fi23++ {
                fmt.Println(v1, fi23)
                var v24 int = -58
                _ = fi23
                _ = v24
            }
        } else {
            fmt.Println(v5)
            if (!true) {
                switch 42 {
                case -328:
                    fmt.Println(v2)
                case 32767:
                    var v25 int64 = (v1 / (((^(v1))) | 1))
                    _ = v25
                case 9223372036854775807:
                    if (((^(v5)) >= 0) || false) {
                        if true {
                            v1 -= v1
                            var v26 int64 = (v1 << ((18446744073709551615) & 63))
                            _ = v26
                        }
                    } else {
                        var v27 int = 42
                        v28 := []int{v27, 128, (v27 + 4294967295), -22}
                        var v29 uint32 = (v2 | 221)
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                default:
                    if true {
                        v1 -= -9223372036854775808
                    }
                }
                var v30 uint32 = v2
                v31 := GenS2{F0: ("" + ""), F1: v2}
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 8; fi32++ {
                v4 = (v4 - -373)
                for fi33 := 0; fi33 < 1; fi33++ {
                    var v34 string = "a"
                    v1 -= (v1 % ((v1) | 1))
                    var v35 int8 = 0
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                var v36 int16 = -32768
                _ = fi32
                _ = v36
            }
        }
    default:
        v4 = 236
    }
    switch 2147483647 {
    case -1:
        var v37 uint = 413
        _ = v37
    }
    if (v3 > v3) {
        for fi38 := 0; fi38 < 14; fi38++ {
            switch (fi38 >> ((255) & 63)) {
            case 42:
                for fi39 := 0; fi39 < 20; fi39++ {
                    var v40 float64 = v3
                    if true {
                        v41 := []int{0}
                        fmt.Println(v1, v2)
                        _ = v41
                    }
                    if false {
                        v5 = v5
                        if true {
                            var v42 int64 = (v1 ^ v1)
                            v4 = (v4 << ((genFn0()) & 63))
                            _ = v42
                        } else {
                            var v43 int8 = -1
                            if ((v1 | v1) >= -327) {
                                var v44 uint16 = 2
                                fmt.Println(v2)
                                _ = v44
                            } else {
                                v45 := []int{fi39, fi38, fi38, int(v2)}
                                fmt.Println(v43, v1)
                                if (false || false) {
                                    v2 = (v2 ^ (v2 >> ((18446744073709551615) & 63)))
                                    var v46 int8 = 0
                                    if len(v45) > 0 {
                                        v45[0] = fi39
                                    }
                                    _ = v46
                                } else {
                                    var v47 uint64 = (v5 / ((v5) | 1))
                                    _ = v47
                                }
                                _ = v45
                            }
                            fmt.Println(v40, v2, v4)
                            var v48 int = (fi39 & 4294967295)
                            var v49 float32 = -2.25
                            _ = v43
                            _ = v48
                            _ = v49
                        }
                    }
                    _ = fi39
                    _ = v40
                }
            case 4294967295:
                var v50 int64 = v1
                _ = v50
            default:
                v1 = 42
            }
            fmt.Println(v2, v5, v5)
            var v51 float64 = (0.0 * v3)
            _ = fi38
            _ = v51
        }
        for fi52 := 0; fi52 < 12; fi52++ {
            var v53 int = fi52
            _ = fi52
            _ = v53
        }
        v3 = v3
    } else {
        v54 := GenS2{F0: "", F1: (v2 << ((65535) & 63))}
        v55 := GenS3{F0: (v4 + v4), F1: 579, F2: (!(false || false))}
        v56 := GenS2{F0: "foo bar", F1: v2}
        _ = v54
        _ = v55
        _ = v56
    }
    for fi57 := 0; fi57 < 11; fi57++ {
        v3 -= v3
        _ = fi57
    }
    v58 := GenS4{F0: 0, F1: true}
    var v59 float32 = (0.1 / 1.0)
    v1 -= 664
    var v60 int8 = 1
    var v61 int8 = v60
    v5 += (v5 | v5)
    var v62 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v60
    _ = v61
    _ = v62
}
