package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint32; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int64; F2 int8 }
func genFn0() uint32 {
    switch 42 {
    case 4294967295:
        var v7 uint16 = 124
        _ = v7
    case -2147483648:
        var v8 bool = (-346 == 42)
        _ = v8
    default:
        if true {
            var v9 float32 = 0.0
            _ = v9
        } else {
            switch (("hello" + "hello") + ("" + "foo bar")) {
            case "hello":
                v10 := GenS1{F0: 0, F1: ((!true) && true), F2: 42}
                _ = v10
            default:
                var v11 uint16 = 0
                _ = v11
            }
            if false {
                var v12 int64 = 0
                v12 = v12
                _ = v12
            } else {
                var v13 int8 = -1
                if (0 > 65535) {
                    v14 := GenS1{F0: 264, F1: (0 < 387), F2: 2}
                    var v15 string = "foo bar"
                    _ = v14
                    _ = v15
                } else {
                    var v16 uint64 = 4294967295
                    var v17 uint16 = 88
                    if false {
                        switch ("\n\t" + "") {
                        case "hello":
                            v18 := []int{42}
                            _ = v18
                        case "":
                            v19 := GenS1{F0: 352, F1: true, F2: 95}
                            _ = v19
                        case "\n\t":
                            switch 693 {
                            case 1:
                                var v20 float64 = 0.0
                                _ = v20
                            case 127:
                                var v21 int = 32767
                                _ = v21
                            case -880:
                                v17 += v17
                            }
                        default:
                            for fi22 := 0; fi22 < 18; fi22++ {
                                var v23 uint32 = 65535
                                _ = fi22
                                _ = v23
                            }
                        }
                        v17 += (v17 | (v17 % ((2) | 1)))
                    } else {
                        var v24 int8 = -1
                        fmt.Println(v17)
                        for fi25 := 0; fi25 < 16; fi25++ {
                            v17 = v17
                            switch 32767 {
                            case 9223372036854775807:
                                v17 += 882
                            case -239:
                                if false {
                                    v26 := GenS2{F0: 0, F1: 1, F2: (v16 >> ((2) & 63))}
                                    var v27 bool = (!false)
                                    v28 := GenS3{F0: -982, F1: uint32(v17), F2: (v24 & v13)}
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                } else {
                                    v24 = v13
                                    v24 = (v13 ^ (-(v13)))
                                }
                            case 32767:
                                v29 := GenS3{F0: 42, F1: 65535, F2: v13}
                                _ = v29
                            default:
                                for fi30 := 0; fi30 < 2; fi30++ {
                                    var v31 float32 = 1e10
                                    var v32 uint = 65535
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                            }
                            for fi33 := 0; fi33 < 1; fi33++ {
                                v34 := []int{fi25, fi33, fi33}
                                fmt.Println(fi25)
                                _ = fi33
                                _ = v34
                            }
                            _ = fi25
                        }
                        _ = v24
                    }
                    _ = v16
                    _ = v17
                }
                _ = v13
            }
            var v35 int64 = 9223372036854775807
            _ = v35
        }
    }
    var v36 float64 = 1.5
    fmt.Println(v36)
    _ = v36
    return 42
}

func main() {
    var v1 int = -498
    var v2 uint32 = 2147483647
    v3 := make(map[string]bool)
    var v4 int32 = 1
    v5 := make(map[string]int)
    var v6 int32 = v4
    v4 += (v4 | v4)
    if false {
        for fi37 := 0; fi37 < 18; fi37++ {
            for fi38 := 0; fi38 < 20; fi38++ {
                v6 = (-(v4))
                fmt.Println(fi38, v4)
                switch "a" {
                case "":
                    if (true || (!false)) {
                        var v39 int16 = 441
                        v1 = (fi38 - fi38)
                        v40 := GenS1{F0: -1, F1: (false || (false || (0.0 >= float32((float64(0.1) + -2.25))))), F2: 255}
                        _ = v39
                        _ = v40
                    } else {
                        v41 := GenS2{F0: 2, F1: 2, F2: 2}
                        if ("hello" == "") {
                            v6 = (v4 & (v6 ^ v4))
                        } else {
                            var v42 uint16 = 65535
                            _ = v42
                        }
                        v1 += 315
                        if true {
                            var v43 int32 = (^(v6))
                            switch (v1 >> ((2) & 63)) {
                            case -180:
                                var v44 float32 = ((1e10 * ((0.0 + -2.25) / -2.25)) * 1.5)
                                _ = v44
                            default:
                                switch "" {
                                case "\n\t":
                                    v45 := GenS3{F0: -25, F1: v2, F2: -1}
                                    _ = v45
                                case "foo bar":
                                    v1 = fi37
                                default:
                                    var v46 int8 = 105
                                    _ = v46
                                }
                            }
                            _ = v43
                        }
                        _ = v41
                    }
                }
                _ = fi38
            }
            _ = fi37
        }
        var v47 int32 = v4
        _ = v47
    } else {
        var v48 int16 = 1000
        v49 := GenS4{F0: -121}
        _ = v48
        _ = v49
    }
    var v50 float64 = (float64(1.5) - 1.5)
    v51 := []int{(v1 >> ((255) & 63)), v1, v1, (v1 >> ((65535) & 63))}
    if (((2 >= 0) || (-2.25 != ((0.0 * (1e10 + float64(0.1))) + v50))) || (432 >= 0)) {
        v2 = (v2 & v2)
        v52 := GenS4{F0: 1}
        v53 := v5[""]
        _ = v52
        _ = v53
    } else {
        var v54 bool = false
        var v55 uint16 = 2
        v56 := v3["a"]
        _ = v54
        _ = v55
        _ = v56
    }
    _ = len(v5)
    for fi57 := 0; fi57 < 13; fi57++ {
        switch "foo bar" {
        case "foo bar":
            _ = len(v51)
        default:
            fmt.Println(v1)
        }
        _ = fi57
    }
    v58 := GenS2{F0: 42, F1: 1, F2: 2}
    v6 = -2147483648
    v3["foo bar"] = (42 < 42)
    v59 := GenS5{F0: ("\n\t" + "foo bar"), F1: 412, F2: -27}
    switch (^(v1)) {
    case 32768:
        var v60 uint = 255
        _ = v60
    default:
        var v61 uint = 65535
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v58
    _ = v59
    _ = v6
}
