package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 string; F1 int; F2 uint32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int }
func genFn0(p0 uint16) int32 {
    v6 := []int{210}
    var v7 bool = false
    if len(v6) > 0 {
        v6[0] = 1
    }
    switch "hello" {
    case "":
        var v8 uint = 1
        _ = v8
    case "hello":
        switch 2147483647 {
        case 32767:
            v9 := GenS1{F0: 255, F1: -9223372036854775808}
            _ = v9
        case 1:
            var v10 uint8 = 55
            _ = v10
        }
    case "foo bar":
        if v7 {
            var v11 int8 = -120
            if v7 {
                v12 := GenS2{F0: 42, F1: v11}
                v11 += v11
                switch ("\n\t" + "a") {
                case "\n\t":
                    var v13 int = 9223372036854775807
                    _ = v13
                default:
                    var v14 bool = v7
                    _ = v14
                }
                _ = v12
            } else {
                v6 = append(v6, 42)
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 uint16 = 42
                    _ = fi15
                    _ = v16
                }
            }
            _ = len(v6)
            _ = v11
        }
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS3{F0: 255, F1: 18446744073709551615}
        if false {
            v19 := GenS4{F0: uint64(p0)}
            switch (fi17 << ((653) & 63)) {
            case 4294967295:
                switch (fi17 ^ fi17) {
                case 127:
                    p0 += (p0 % ((p0) | 1))
                default:
                    _ = len(v6)
                }
            case 229:
                if len(v6) > 0 {
                    v6[0] = (fi17 | (fi17 % (((^(fi17))) | 1)))
                }
            }
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = v6
    _ = v7
    return -970
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]bool)
    var v3 int64 = -9223372036854775808
    var v4 uint16 = 2
    var v5 int64 = v3
    v5 = v5
    switch "\n\t" {
    case "foo bar":
        if (891 > 255) {
            var v20 uint16 = 33
            v4 = 529
            for fi21 := 0; fi21 < 8; fi21++ {
                v3 += (v5 >> ((2) & 63))
                if true {
                    fmt.Println(v20, fi21)
                }
                _ = fi21
            }
            _ = v20
        }
    }
    fmt.Println(v5, v4, v4)
    if len(v1) > 0 {
        v1[0] = 494
    }
    var v22 int = -117
    switch v22 {
    case 32767:
        if (((-1 <= -93) && (!false)) || true) {
            if false {
                v23 := GenS4{F0: 4294967295}
                var v24 float64 = 1e10
                var v25 uint8 = uint8(v4)
                _ = v23
                _ = v24
                _ = v25
            } else {
                if true {
                    if (18446744073709551615 >= 370) {
                        if true {
                            v3 = v3
                            if (!true) {
                                v26 := GenS5{F0: 1}
                                _ = v26
                            }
                        }
                        var v27 string = ""
                        if (18446744073709551615 < 2) {
                            if false {
                                var v28 bool = (4294967295 == 65535)
                                _ = v28
                            }
                            var v29 uint16 = v4
                            var v30 float64 = float64(float32(0.0))
                            _ = v29
                            _ = v30
                        } else {
                            if true {
                                switch "foo bar" {
                                case "foo bar":
                                    v31 := GenS6{F0: 255, F1: 417}
                                    _ = v31
                                case "":
                                    if true {
                                        if (!true) {
                                            v27 = (v27 + (("" + ("foo bar" + "")) + "hello"))
                                            fmt.Println(v4)
                                            if len(v1) > 0 {
                                                v1[0] = 0
                                            }
                                        } else {
                                            v32 := GenS2{F0: (v5 / ((v3) | 1)), F1: 0}
                                            fmt.Println(v27, v3)
                                            _ = v32
                                        }
                                    }
                                case "hello":
                                    var v33 uint32 = 2147483647
                                    _ = v33
                                default:
                                    if false {
                                        var v34 int16 = 42
                                        _ = v34
                                    } else {
                                        v35 := GenS7{F0: v4}
                                        v22 -= (v22 >> ((42) & 63))
                                        _ = v35
                                    }
                                }
                                v36 := GenS8{F0: "hello", F1: v22, F2: 561}
                                _ = v36
                            }
                            v5 = 9223372036854775807
                        }
                        _ = v27
                    }
                    for fi37 := 0; fi37 < 15; fi37++ {
                        for fi38 := 0; fi38 < 11; fi38++ {
                            v39 := GenS8{F0: "", F1: (v22 + fi37), F2: 4294967295}
                            _ = fi38
                            _ = v39
                        }
                        _ = fi37
                    }
                    fmt.Println(v5, v4)
                } else {
                    for fi40 := 0; fi40 < 19; fi40++ {
                        v22 = 32767
                        _ = fi40
                    }
                    if true {
                        if len(v1) > 0 {
                            v1[0] = 856
                        }
                        v5 = (^(v3))
                        v41 := GenS9{F0: 4294967295}
                        _ = v41
                    } else {
                        if (float64(1.5) <= 0.1) {
                            v4 = v4
                        }
                        var v42 uint = 679
                        _ = v42
                    }
                    if len(v1) > 0 {
                        v1[0] = 2147483648
                    }
                }
                var v43 int32 = genFn0(v4)
                if len(v1) > 0 {
                    v1[0] = 2
                }
                _ = v43
            }
            for fi44 := 0; fi44 < 6; fi44++ {
                v3 = int64(fi44)
                v45 := GenS10{F0: fi44}
                var v46 int = (fi44 + (fi44 / (((-(fi44))) | 1)))
                _ = fi44
                _ = v45
                _ = v46
            }
            v22 -= v22
        } else {
            var v47 int = v22
            _ = v47
        }
    case 127:
        v48 := GenS7{F0: (^(v4))}
        _ = v48
    default:
        for fi49 := 0; fi49 < 5; fi49++ {
            var v50 int16 = -135
            _ = fi49
            _ = v50
        }
    }
    for fi51 := 0; fi51 < 4; fi51++ {
        v4 -= 255
        _ = fi51
    }
    v4 = (v4 | (v4 - (v4 & v4)))
    var v52 int64 = (v5 & (v3 + v3))
    v4 = v4
    if (1 >= 0) {
        v3 += (v52 + (v3 / (((v52 << ((42) & 63))) | 1)))
        var v53 int8 = 96
        if len(v1) > 0 {
            v1[0] = 42
        }
        _ = v53
    }
    v5 -= 42
    var v54 int32 = 671
    v3 += 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v54
}
