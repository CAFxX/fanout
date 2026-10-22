package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int32; F2 int32 }
type GenS4 struct { F0 string; F1 int; F2 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32; F1 uint64 }
type GenS8 struct { F0 int16; F1 uint; F2 int64 }
type GenS9 struct { F0 int8; F1 bool }
type GenS10 struct { F0 int64; F1 uint; F2 int32 }
type GenS11 struct { F0 int; F1 int8 }
func genFn0() int16 {
    v7 := GenS1{F0: 4294967295, F1: 407, F2: 304}
    var v8 bool = (!false)
    v9 := GenS2{F0: v8}
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int8 = -1
    var v2 string = ""
    var v3 float32 = ((1.5 - 0.1) / ((float32(float64(3.14159)) * 0.1) + 0.1))
    var v4 bool = false
    v5 := make(map[string]uint32)
    var v6 int16 = -45
    v10 := []int{128, int(v6), int(v6), 18}
    var v11 int8 = (-(v1))
    if v4 {
        if v4 {
            v12 := GenS3{F0: 256, F1: 503, F2: 42}
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 uint16 = 2
                switch "" {
                case "\n\t":
                    v5["a"] = 2147483647
                case "":
                    for fi15 := 0; fi15 < 2; fi15++ {
                        switch v2 {
                        case "a":
                            v4 = false
                        default:
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        }
                        v11 -= 0
                        fmt.Println(v3, fi13)
                        v2 = "foo bar"
                        _ = fi15
                    }
                }
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        var v16 string = "a"
        var v17 bool = v4
        _ = v16
        _ = v17
    } else {
        v5[("\n\t" + "a")] = 518
        v18 := GenS4{F0: v2, F1: 128, F2: 4294967295}
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := GenS1{F0: 72, F1: 0, F2: v6}
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    switch 83 {
    case 32767:
        var v21 int = 42
        _ = v21
    default:
        if (0 <= 0) {
            for fi22 := 0; fi22 < 15; fi22++ {
                if v4 {
                    if true {
                        for fi23 := 0; fi23 < 8; fi23++ {
                            fmt.Println(v1, fi22, v4)
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            v6 -= v6
                            _ = fi23
                        }
                    }
                    if false {
                        var v24 int16 = v6
                        v10 = append(v10, (fi22 | int(v24)))
                        _ = v24
                    }
                }
                switch (fi22 >> ((531) & 63)) {
                case 4294967295:
                    v25 := GenS5{F0: (-(v6))}
                    _ = v25
                case 45:
                    switch (fi22 & fi22) {
                    case 32767:
                        switch 42 {
                        case 32768:
                            v6 = genFn0()
                        case 313:
                            var v26 int32 = 1
                            _ = v26
                        case -1:
                            var v27 uint32 = 78
                            _ = v27
                        default:
                            v28 := GenS3{F0: 593, F1: -1, F2: 71}
                            _ = v28
                        }
                    case -9223372036854775808:
                        if false {
                            var v29 int64 = 0
                            _ = v29
                        } else {
                            var v30 uint8 = 192
                            var v31 int32 = 584
                            _ = v30
                            _ = v31
                        }
                    }
                }
                _ = fi22
            }
            if (42 > 2) {
                var v32 int8 = (v1 | (^(v11)))
                _ = v32
            } else {
                for fi33 := 0; fi33 < 17; fi33++ {
                    for fi34 := 0; fi34 < 20; fi34++ {
                        if (65535 > 225) {
                            switch 1 {
                            case 771:
                                switch (fi34 % ((fi34) | 1)) {
                                case 128:
                                    v5[v2] = 952
                                }
                            }
                            _ = len(v5)
                            var v35 uint64 = 42
                            _ = v35
                        }
                        switch fi33 {
                        case -275:
                            if len(v10) > 0 {
                                v10[0] = (^(fi33))
                            }
                        case 127:
                            switch fi34 {
                            case 623:
                                var v36 int16 = (-(v6))
                                _ = v36
                            }
                        default:
                            v1 = v11
                        }
                        if v4 {
                            if v4 {
                                fmt.Println(v1, v2, v2)
                                fmt.Println(fi33, v4)
                                var v37 int = fi33
                                _ = v37
                            }
                            v6 -= genFn0()
                        }
                        _ = fi34
                    }
                    var v38 uint64 = 4294967295
                    if false {
                        for fi39 := 0; fi39 < 11; fi39++ {
                            if v4 {
                                var v40 uint = 791
                                _ = v40
                            } else {
                                var v41 int32 = 2147483647
                                v2 = v2
                                _ = v41
                            }
                            v42 := GenS6{F0: 42}
                            var v43 uint32 = 4294967295
                            v44 := GenS7{F0: 42, F1: (v38 - (v38 & (v38 ^ (v38 | (v38 - v38)))))}
                            v45 := v5["a"]
                            _ = fi39
                            _ = v42
                            _ = v43
                            _ = v44
                            _ = v45
                        }
                        v4 = (!true)
                        fmt.Println(v2)
                    }
                    _ = fi33
                    _ = v38
                }
                for fi46 := 0; fi46 < 3; fi46++ {
                    var v47 int16 = v6
                    switch fi46 {
                    case 809:
                        switch (((v2 + (v2 + v2)) + v2) + "a") {
                        case "\n\t":
                            v6 -= v47
                        }
                    case -2147483648:
                        v6 = v47
                    case -1:
                        if false {
                            _ = len(v10)
                        } else {
                            switch 32768 {
                            case 42:
                                v4 = v4
                            default:
                                var v48 float32 = v3
                                _ = v48
                            }
                            if (!(18446744073709551615 < 461)) {
                                v3 -= float32(3.14159)
                                fmt.Println(v2, v47)
                                if len(v10) > 0 {
                                    v10[0] = (^(fi46))
                                }
                            }
                            v49 := v5["a"]
                            _ = v49
                        }
                    default:
                        v10 = append(v10, (fi46 * fi46))
                    }
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = fi46
                    _ = v47
                }
            }
        } else {
            if v4 {
                fmt.Println(v6, v4, v11)
                var v50 uint64 = 662
                _ = v50
            } else {
                var v51 uint = 0
                fmt.Println(v51, v3)
                for fi52 := 0; fi52 < 12; fi52++ {
                    v53 := GenS8{F0: (^(v6)), F1: (v51 >> ((v51) & 63)), F2: 0}
                    v1 = v1
                    _ = fi52
                    _ = v53
                }
                _ = v51
            }
        }
    }
    var v54 int64 = 601
    var v55 uint16 = 65535
    v56 := GenS5{F0: (v6 - v6)}
    v54 = v54
    for fi57 := 0; fi57 < 18; fi57++ {
        if true {
            v58 := v5[""]
            for fi59 := 0; fi59 < 16; fi59++ {
                var v60 string = (v2 + "a")
                var v61 uint = uint(v55)
                v11 = v11
                _ = fi59
                _ = v60
                _ = v61
            }
            _ = v58
        }
        v10 = append(v10, (-(fi57)))
        v4 = v4
        _ = fi57
    }
    _ = len(v5)
    for fi62 := 0; fi62 < 18; fi62++ {
        var v63 uint64 = 2
        var v64 int16 = v6
        v5[v2] = 42
        _ = fi62
        _ = v63
        _ = v64
    }
    for fi65 := 0; fi65 < 4; fi65++ {
        v66 := GenS9{F0: v1, F1: (float32((0.1 * float64(1e10))) == ((float32(0.0) + (float32(((float64(v3) / 1.0) * 0.1)) + v3)) + -2.25))}
        for fi67 := 0; fi67 < 15; fi67++ {
            v68 := v5["hello"]
            _ = fi67
            _ = v68
        }
        v69 := GenS10{F0: (^(v54)), F1: 454, F2: int32(v6)}
        _ = fi65
        _ = v66
        _ = v69
    }
    if ((42 >= (v6 & genFn0())) || v4) {
        if (819 >= -9) {
            _ = len(v10)
            _ = len(v5)
        }
    } else {
        v70 := GenS11{F0: -1, F1: (v1 << ((0) & 63))}
        v55 -= 42
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v6
}
