package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 uint; F2 int64 }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint32) uint8 {
    v6 := GenS1{F0: -56}
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int8 = 0
            if (!false) {
                var v10 int64 = 9223372036854775807
                switch (fi8 & fi8) {
                case 2147483647:
                    var v11 int32 = 1000
                    _ = v11
                case -9223372036854775808:
                    var v12 string = ("a" + "")
                    _ = v12
                case 9223372036854775807:
                    if (((^(v9)) < (v9 | v9)) && true) {
                        v13 := []int{}
                        if false {
                            fmt.Println(fi8)
                            v9 = v9
                            v13 = append(v13, fi8)
                            var v14 float32 = 0.0
                            _ = v14
                        } else {
                            p0 += p0
                            switch "hello" {
                            case "hello":
                                var v15 uint64 = 0
                                _ = v15
                            case "a":
                                v9 -= (v9 / ((-81) | 1))
                            default:
                                fmt.Println(fi8)
                            }
                        }
                        _ = v13
                    }
                default:
                    if true {
                        var v16 bool = (0 >= int32(fi8))
                        _ = v16
                    } else {
                        var v17 uint8 = 1
                        if ("hello" == "\n\t") {
                            p0 = 65535
                            fmt.Println(fi8, v17, fi7)
                            if (true || (4294967295 <= 104)) {
                                v17 = (v17 / ((v17) | 1))
                                v18 := GenS2{F0: (255 <= (^(v17))), F1: 2}
                                v19 := GenS3{F0: p0, F1: v10, F2: 0}
                                _ = v18
                                _ = v19
                            } else {
                                v20 := []int{}
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                                var v21 float32 = -2.25
                                _ = v20
                                _ = v21
                            }
                        } else {
                            v10 -= v10
                            v17 -= (v17 % ((v17) | 1))
                        }
                        var v22 uint64 = 1
                        _ = v17
                        _ = v22
                    }
                }
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 uint16, p1 int16) int32 {
    var v23 int32 = 2147483647
    var v24 int = -1
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return 2147483647
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 0
    var v3 float32 = 1.5
    var v4 bool = (-882 != 1)
    var v5 int8 = 1
    for fi25 := 0; fi25 < 1; fi25++ {
        fmt.Println(v5)
        _ = fi25
    }
    if (!false) {
        v1[("\n\t" + ("foo bar" + "hello"))] = ((^(v2)) == 35)
    } else {
        var v26 int32 = 487
        var v27 uint64 = 626
        var v28 bool = (!(!v4))
        _ = v26
        _ = v27
        _ = v28
    }
    v29 := []int{-1}
    if v4 {
        var v30 uint8 = 2
        var v31 uint8 = genFn0(42)
        if v4 {
            var v32 int64 = 9223372036854775807
            delete(v1, ("a" + "\n\t"))
            var v33 uint32 = 0
            _ = v32
            _ = v33
        } else {
            var v34 uint = 832
            _ = v34
        }
        _ = v30
        _ = v31
    }
    switch "\n\t" {
    case "hello":
        if v4 {
            switch "foo bar" {
            case "a":
                var v35 int64 = 42
                _ = v35
            case "foo bar":
                var v36 bool = false
                _ = v36
            default:
                for fi37 := 0; fi37 < 2; fi37++ {
                    if v4 {
                        v38 := v1[""]
                        var v39 int16 = -679
                        if v38 {
                            v29 = append(v29, 127)
                        } else {
                            fmt.Println(fi37, v4, v5)
                            var v40 float32 = v3
                            v41 := GenS3{F0: 1, F1: -225, F2: 18446744073709551615}
                            _ = v40
                            _ = v41
                        }
                        _ = v38
                        _ = v39
                    }
                    switch (fi37 + 2147483647) {
                    case 4294967295:
                        if (2 >= 1) {
                            if false {
                                fmt.Println(v3)
                                v4 = v4
                            } else {
                                var v42 uint8 = genFn0(42)
                                v2 = (v42 % ((genFn0(2147483648)) | 1))
                                _ = v42
                            }
                            fmt.Println(v3)
                        } else {
                            v43 := v1["foo bar"]
                            _ = v43
                        }
                    }
                    v1[""] = false
                    _ = fi37
                }
            }
        }
    case "":
        var v44 float64 = float64((0.1 * (v3 * ((float32(1.5) - (0.1 - ((((1.5 / 3.14159) + v3) / ((((v3 / 3.14159) / v3) - ((1e10 - float32(((0.1 + 3.14159) / 1e10))) - v3)) * v3)) / v3))) + float32((((-2.25 * 3.14159) - -2.25) - -2.25))))))
        _ = v44
    default:
        fmt.Println(v3)
    }
    v2 = v2
    if len(v29) > 0 {
        v29[0] = -2147483648
    }
    v4 = ((4294967295 > 959) || true)
    v45 := GenS4{F0: 1}
    var v46 float32 = v3
    if (true && ((false || (!(v4 || ((v4 || (-1 == 0)) || (633 >= 0))))) && v4)) {
        v2 += v2
        v47 := GenS4{F0: 32767}
        var v48 float64 = -2.25
        _ = v47
        _ = v48
    } else {
        v2 += 42
        v49 := GenS5{F0: 18446744073709551615, F1: 42, F2: 42}
        _ = v49
    }
    v3 += 1.5
    var v50 uint64 = 980
    for fi51 := 0; fi51 < 11; fi51++ {
        if true {
            v50 = (v50 + (^(v50)))
        }
        if v4 {
            if ((!v4) || ((!true) || (2 >= 1))) {
                var v52 bool = true
                var v53 int8 = v5
                var v54 uint8 = 1
                _ = v52
                _ = v53
                _ = v54
            } else {
                v55 := GenS4{F0: -32768}
                _ = v55
            }
            switch fi51 {
            case 9223372036854775807:
                switch fi51 {
                case -2147483648:
                    var v56 int32 = 0
                    _ = v56
                default:
                    for fi57 := 0; fi57 < 10; fi57++ {
                        v58 := GenS6{F0: (^(v5)), F1: 42}
                        if len(v29) >= 2 {
                            v29 = v29[1:len(v29)]
                        }
                        _ = fi57
                        _ = v58
                    }
                }
            default:
                if (42 < 0) {
                    if ((42 == 223) && (1e10 >= 3.14159)) {
                        var v59 uint = 255
                        fmt.Println(v2)
                        var v60 string = "a"
                        _ = v59
                        _ = v60
                    }
                } else {
                    var v61 uint64 = v50
                    _ = v61
                }
            }
            v5 += v5
        } else {
            if (151 >= 2) {
                var v62 bool = v4
                if true {
                    v5 -= -117
                } else {
                    v46 += 0.1
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                }
                _ = v62
            }
        }
        _ = fi51
    }
    v63 := GenS7{F0: v50, F1: 0}
    var v64 bool = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v50
    _ = v63
    _ = v64
}
