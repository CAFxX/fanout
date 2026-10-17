package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint64 }
func genFn0() int8 {
    var v7 bool = (4294967295 == 756)
    v8 := GenS1{F0: 42, F1: 436}
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    var v3 uint = 658
    v4 := []uint32{42, 0}
    v5 := []uint32{4294967295, 65535, 2147483647, 0}
    v6 := []uint32{1}
    if (-1 <= 89) {
        var v9 int16 = 32767
        v10 := GenS2{F0: 2147483647}
        _ = v10
        _ = v9
    }
    var v11 int32 = -1
    fmt.Println(v3, v11)
    if (-1 < -9223372036854775808) {
        v12 := GenS2{F0: 42}
        if true {
            var v13 int16 = 32767
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v13 = 32767
                fmt.Println(v11, v13)
            } else {
                v13 += v13
            }
            _ = v13
        } else {
            if ("foo bar" == "") {
                for fi14 := 0; fi14 < 6; fi14++ {
                    if (!(float64(1e10) >= 0.1)) {
                        for fi15 := 0; fi15 < 18; fi15++ {
                            v11 = 1
                            v16 := GenS1{F0: 42, F1: 1}
                            v3 = v3
                            _ = fi15
                            _ = v16
                        }
                        if true {
                            fmt.Println(v3)
                            switch fi14 {
                            case 39:
                                var v17 uint32 = 42
                                _ = v17
                            case -497:
                                var v18 int32 = 2147483647
                                _ = v18
                            default:
                                for fi19 := 0; fi19 < 2; fi19++ {
                                    if true {
                                        fmt.Println(v11, fi19, v3)
                                    }
                                    v20 := GenS3{F0: "", F1: 276, F2: 69}
                                    _ = fi19
                                    _ = v20
                                }
                            }
                            fmt.Println(v11, fi14, v3)
                        }
                    } else {
                        v11 = -868
                    }
                    _ = fi14
                }
                v11 = (v11 | 42)
                fmt.Println(v3, v11)
            } else {
                v11 -= -448
                if true {
                    v21 := GenS3{F0: "", F1: 2, F2: genFn0()}
                    if (!false) {
                        var v22 float64 = 1e10
                        var v23 int8 = 100
                        _ = v22
                        _ = v23
                    } else {
                        if false {
                            v24 := GenS4{F0: 18446744073709551615, F1: (v11 / ((v11) | 1)), F2: 32767}
                            _ = v24
                        } else {
                            for fi25 := 0; fi25 < 11; fi25++ {
                                fmt.Println(fi25, fi25)
                                _ = fi25
                            }
                        }
                        switch 32767 {
                        case 345:
                            if true {
                                switch "a" {
                                case "a":
                                    for fi26 := 0; fi26 < 7; fi26++ {
                                        var v27 uint16 = 42
                                        if false {
                                            v28 := GenS5{F0: -924}
                                            switch "" {
                                            case "":
                                                var v29 uint32 = 2
                                                _ = v29
                                            case "\n\t":
                                                fmt.Println(v27)
                                            default:
                                                var v30 uint64 = 0
                                                _ = v30
                                            }
                                            _ = v28
                                        }
                                        switch -459 {
                                        case 32767:
                                            _ = len(v1)
                                        default:
                                            v27 += 32767
                                        }
                                        _ = fi26
                                        _ = v27
                                    }
                                case "foo bar":
                                    var v31 int32 = v11
                                    _ = v31
                                case "":
                                    if false {
                                        v3 += 4294967295
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v3 -= 255
                                    }
                                default:
                                    v32 := v1[42]
                                    _ = v32
                                }
                            } else {
                                v11 -= (v11 % (((v11 / (((-(v11))) | 1))) | 1))
                                v11 -= v11
                                v11 = 42
                            }
                        default:
                            for fi33 := 0; fi33 < 3; fi33++ {
                                var v34 uint16 = 42
                                if (-1 <= 218) {
                                    var v35 uint32 = 895
                                    _ = v35
                                } else {
                                    if (!false) {
                                        if (!true) {
                                            v36 := GenS6{F0: v11, F1: (^(v34)), F2: 42}
                                            _ = v36
                                        }
                                    }
                                    v37 := GenS7{F0: 0}
                                    _ = v37
                                }
                                v2[(("a" + (("\n\t" + (("foo bar" + "\n\t") + "")) + (("" + "foo bar") + ("\n\t" + "")))) + (("a" + "\n\t") + ("a" + "\n\t")))] = 440
                                _ = fi33
                                _ = v34
                            }
                        }
                        for fi38 := 0; fi38 < 13; fi38++ {
                            var v39 float64 = float64(0.1)
                            _ = fi38
                            _ = v39
                        }
                    }
                    _ = v21
                } else {
                    v40 := GenS7{F0: 32}
                    _ = v40
                }
                delete(v2, "")
            }
            var v41 int8 = 1
            v1[0] = false
            _ = v41
        }
        fmt.Println(v3, v3)
        _ = v12
    } else {
        if (!(false && true)) {
            var v42 uint = 65535
            fmt.Println(v42, v42, v42)
            v42 += (^(v42))
            _ = v42
        }
        var v43 int16 = 979
        var v44 int16 = (v43 + -1)
        _ = v43
        _ = v44
    }
    fmt.Println(v3, v3)
    v45 := GenS6{F0: v11, F1: 42, F2: -128}
    switch 42 {
    case 459:
        fmt.Println(v11, v11)
    case -472:
        v3 = 704
    }
    fmt.Println(v3, v11)
    var v46 float64 = 0.1
    v46 = v46
    var v47 uint8 = 108
    v47 += 104
    var v48 int16 = 205
    v1[32768] = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
