package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int8 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint }
type GenS5 struct { F0 uint; F1 int64 }
func genFn0() int16 {
    v5 := []int{}
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(fi6)
        var v7 int = 1
        _ = fi6
        _ = v7
    }
    _ = v5
    return -197
}

func main() {
    var v1 uint8 = 0
    var v2 uint16 = 65535
    var v3 string = "\n\t"
    var v4 int32 = 539
    v8 := GenS1{F0: -334, F1: 0, F2: 100}
    v9 := []int{}
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS1{F0: fi10, F1: 0, F2: 1}
        v4 += v4
        _ = fi10
        _ = v11
    }
    v9 = append(v9, 32767)
    v3 = (v3 + ((v3 + (v3 + (v3 + v3))) + ("" + "\n\t")))
    if false {
        if false {
            var v12 int64 = -1
            v13 := GenS1{F0: 42, F1: 100, F2: -1}
            if len(v9) > 0 {
                v9[0] = 128
            }
            _ = v12
            _ = v13
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    if len(v9) > 0 {
        v9[0] = 1
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 int64 = 297
        _ = fi14
        _ = v15
    }
    if len(v9) > 0 {
        v9[0] = 2147483647
    }
    switch -500 {
    case -9223372036854775808:
        var v16 uint64 = 2
        _ = v16
    }
    v9 = append(v9, -1)
    var v17 float32 = 1.5
    _ = len(v9)
    for fi18 := 0; fi18 < 3; fi18++ {
        if (v3 == v3) {
            if true {
                for fi19 := 0; fi19 < 13; fi19++ {
                    v1 = (v1 - 2)
                    _ = fi19
                }
                var v20 uint32 = 1
                switch fi18 {
                case 128:
                    for fi21 := 0; fi21 < 7; fi21++ {
                        v22 := GenS2{F0: (^(v20)), F1: (true || (false && (false || ((v3 + "") == v3)))), F2: (^(v20))}
                        _ = fi21
                        _ = v22
                    }
                case -375:
                    switch (v3 + v3) {
                    case "\n\t":
                        for fi23 := 0; fi23 < 20; fi23++ {
                            var v24 float32 = (v17 + (-2.25 + v17))
                            _ = fi23
                            _ = v24
                        }
                    case "foo bar":
                        v20 = 65535
                    case "hello":
                        for fi25 := 0; fi25 < 20; fi25++ {
                            v26 := GenS1{F0: (-(fi25)), F1: 127, F2: 0}
                            var v27 int64 = 660
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    default:
                        v2 += (v2 | 2)
                    }
                default:
                    switch ("\n\t" + "hello") {
                    case "foo bar":
                        fmt.Println(fi18, fi18, v20)
                    default:
                        if true {
                            if ((868 <= 1) && (false || (42 >= 834))) {
                                var v28 bool = false
                                _ = v28
                            } else {
                                for fi29 := 0; fi29 < 18; fi29++ {
                                    var v30 uint32 = (v20 | v20)
                                    var v31 uint64 = 42
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                            }
                            if (!(-106 <= 127)) {
                                fmt.Println(v1, v17, v1)
                                for fi32 := 0; fi32 < 9; fi32++ {
                                    if len(v9) > 0 {
                                        v9[0] = (fi18 / ((fi18) | 1))
                                    }
                                    _ = fi32
                                }
                                switch fi18 {
                                case 828:
                                    if false {
                                        var v33 string = (v3 + (v3 + (("\n\t" + "a") + ((("a" + v3) + v3) + ("" + v3)))))
                                        var v34 int = (-(fi18))
                                        switch (v34 & 127) {
                                        case 4294967295:
                                            var v35 string = (v3 + ("\n\t" + v33))
                                            _ = v35
                                        }
                                        _ = v33
                                        _ = v34
                                    }
                                case -774:
                                    var v36 bool = false
                                    _ = v36
                                default:
                                    _ = len(v9)
                                }
                            }
                        }
                    }
                }
                _ = v20
            } else {
                var v37 float32 = v17
                _ = len(v9)
                for fi38 := 0; fi38 < 3; fi38++ {
                    v39 := GenS3{F0: (v1 | v1), F1: 0, F2: 65535}
                    if ((!(32767 != 0)) && true) {
                        switch "\n\t" {
                        case "a":
                            v40 := GenS1{F0: (fi18 << ((18446744073709551615) & 63)), F1: 0, F2: 127}
                            _ = v40
                        default:
                            v3 = v3
                        }
                        v41 := GenS3{F0: (v1 << ((805) & 63)), F1: 516, F2: 2}
                        if (2 < 255) {
                            var v42 int = (-(fi18))
                            _ = v42
                        }
                        _ = v41
                    } else {
                        v43 := GenS2{F0: 65535, F1: false, F2: 453}
                        var v44 int = (fi38 << ((2) & 63))
                        var v45 string = v3
                        v3 = (v3 + v3)
                        v46 := GenS4{F0: v1, F1: (v2 - (^(v2))), F2: 255}
                        v1 += (v1 + v1)
                        var v47 uint = 1
                        _ = v43
                        _ = v44
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    if true {
                        v48 := GenS1{F0: (fi18 / ((fi38) | 1)), F1: 19, F2: -1}
                        if false {
                            if (42 > -9223372036854775808) {
                                v4 = int32(v2)
                                var v49 bool = false
                                _ = v49
                            }
                            fmt.Println(fi18, fi38, v37)
                            switch "a" {
                            case "foo bar":
                                if false {
                                    switch "foo bar" {
                                    case "a":
                                        if ((fi18 * fi18) >= fi18) {
                                            if len(v9) > 0 {
                                                v9[0] = fi38
                                            }
                                        }
                                    case "":
                                        var v50 int8 = -128
                                        _ = v50
                                    case "\n\t":
                                        v1 = v1
                                    }
                                } else {
                                    v17 += (1.5 * (1.5 - v37))
                                }
                            case "hello":
                                v51 := GenS1{F0: fi18, F1: -128, F2: 0}
                                _ = v51
                            case "a":
                                v17 -= (1.5 / (v37 - v37))
                            default:
                                fmt.Println(v3, v2)
                            }
                        } else {
                            switch (fi18 * (fi18 << ((315) & 63))) {
                            case -1:
                                if false {
                                    v4 += 154
                                    switch (fi38 + (fi18 * (fi38 / ((fi38) | 1)))) {
                                    case 721:
                                        var v52 uint8 = (v1 & v1)
                                        _ = v52
                                    case -14:
                                        switch fi38 {
                                        case 2147483647:
                                            var v53 int8 = 0
                                            var v54 uint64 = 42
                                            _ = v53
                                            _ = v54
                                        }
                                    default:
                                        v4 += (v4 & v4)
                                    }
                                    fmt.Println(fi18, v3, fi18)
                                }
                            case -11:
                                v3 = ((((("" + v3) + v3) + v3) + "a") + v3)
                            case 32767:
                                v4 = 1
                            }
                            fmt.Println(v2)
                        }
                        v37 = v17
                        _ = v48
                    }
                    v55 := GenS2{F0: 2147483647, F1: true, F2: 939}
                    fmt.Println(fi38)
                    _ = fi38
                    _ = v39
                    _ = v55
                }
                _ = v37
            }
            fmt.Println(fi18, v2, v3)
        } else {
            v56 := GenS5{F0: 65535, F1: 766}
            v2 = v2
            for fi57 := 0; fi57 < 2; fi57++ {
                v58 := GenS5{F0: 255, F1: -9223372036854775808}
                if (42 >= 42) {
                    var v59 float32 = float32(3.14159)
                    fmt.Println(v2, fi57)
                    var v60 uint = 854
                    var v61 uint32 = 65535
                    var v62 int = (^(fi18))
                    _ = v59
                    _ = v60
                    _ = v61
                    _ = v62
                }
                _ = fi57
                _ = v58
            }
            _ = v56
        }
        if len(v9) > 0 {
            v9[0] = (fi18 - (^(fi18)))
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi18
    }
    if false {
        var v63 int = 0
        _ = v63
    } else {
        v3 = ""
        fmt.Println(v2, v3)
        if true {
            switch ("hello" + v3) {
            case "\n\t":
                for fi64 := 0; fi64 < 10; fi64++ {
                    var v65 int64 = -496
                    v3 = ((("a" + "") + ("a" + v3)) + "")
                    v2 = (v2 << ((42) & 63))
                    _ = fi64
                    _ = v65
                }
            }
            switch 893 {
            case 9223372036854775807:
                var v66 float64 = -2.25
                _ = v66
            }
        }
    }
    _ = len(v9)
    fmt.Println(v17, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
