package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 int64; F2 int32 }
func genFn0() string {
    var v6 uint64 = 4294967295
    var v7 int = 211
    _ = v6
    _ = v7
    return "foo bar"
}
func genFn1(p0 uint64, p1 string) uint8 {
    var v8 int64 = -1
    var v9 uint = 255
    v8 -= (v8 >> (((v9 / ((v9) | 1))) & 63))
    var v10 string = "foo bar"
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = -67
    var v3 uint = 667
    v4 := []string{"hello", "a", "foo bar", ""}
    var v5 int64 = -714
    v5 += -1
    switch "" {
    case "foo bar":
        v5 += 42
    case "hello":
        if false {
            fmt.Println(v3)
        }
    case "a":
        switch "\n\t" {
        case "a":
            v2 = v2
        case "\n\t":
            v11 := GenS1{F0: true}
            _ = v11
        default:
            v3 = (v3 % ((4294967295) | 1))
        }
    }
    if (true || true) {
        var v12 int8 = (v2 | (v2 + v2))
        _ = v12
    } else {
        v3 += v3
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        if true {
            var v14 uint = v3
            v3 += 18446744073709551615
            v5 = v5
            _ = v14
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            delete(v1, 2147483648)
            var v16 uint32 = 871
            _ = fi15
            _ = v16
        }
        if ((42 != 42) || (494 <= 0)) {
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 uint64 = 42
                _ = fi17
                _ = v18
            }
        }
        _ = fi13
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        v3 -= (^(v3))
        _ = fi19
    }
    v20 := GenS1{F0: false}
    fmt.Println(v3, v2)
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 bool = (!(!(65535 != 0)))
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 int8 = (^(v2))
        _ = len(v4)
        var v25 int8 = v2
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 int32 = 2147483647
    switch -9223372036854775808 {
    case 127:
        v5 += 42
    case 2147483647:
        switch 32768 {
        case 32767:
            v3 += (v3 * (v3 >> ((v3) & 63)))
        case 300:
            for fi27 := 0; fi27 < 9; fi27++ {
                for fi28 := 0; fi28 < 7; fi28++ {
                    v26 += -605
                    if true {
                        fmt.Println(v3, v5, fi27)
                        var v29 uint8 = 252
                        v30 := GenS1{F0: false}
                        v2 -= -1
                        v26 -= v26
                        _ = v29
                        _ = v30
                    }
                    _ = fi28
                }
                var v31 bool = (!true)
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                _ = fi27
                _ = v31
            }
        case -586:
            v32 := GenS2{F0: ("\n\t" + ("\n\t" + (("" + "a") + "a"))), F1: 255, F2: v3}
            _ = v32
        }
    }
    if (0.0 == -2.25) {
        var v33 uint8 = 225
        v34 := GenS1{F0: ((!false) || true)}
        _ = v33
        _ = v34
    } else {
        v35 := GenS3{F0: v2}
        var v36 int16 = -292
        _ = v35
        _ = v36
    }
    if true {
        v5 = (v5 ^ v5)
        switch "hello" {
        case "foo bar":
            v37 := GenS1{F0: (0.0 <= (0.1 * 0.1))}
            _ = v37
        case "hello":
            if false {
                v2 = (-(v2))
                _ = len(v1)
                if (-9223372036854775808 <= 1) {
                    v38 := GenS4{F0: (-(v26)), F1: v5}
                    _ = v38
                } else {
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                    _ = len(v4)
                }
            } else {
                for fi39 := 0; fi39 < 19; fi39++ {
                    switch (fi39 << (((v3 | 987)) & 63)) {
                    case -958:
                        for fi40 := 0; fi40 < 17; fi40++ {
                            v41 := GenS3{F0: -121}
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                            if false {
                                var v42 int = (fi39 / (((fi40 ^ fi40)) | 1))
                                _ = v42
                            }
                            _ = fi40
                            _ = v41
                        }
                    default:
                        fmt.Println(v26)
                    }
                    v2 -= (v2 & (v2 * v2))
                    var v43 int64 = (v5 >> ((v3) & 63))
                    _ = fi39
                    _ = v43
                }
            }
        default:
            for fi44 := 0; fi44 < 10; fi44++ {
                fmt.Println(fi44, v3, v2)
                switch 128 {
                case 244:
                    _ = len(v1)
                case -2147483648:
                    v45 := GenS1{F0: true}
                    _ = v45
                default:
                    _ = len(v4)
                }
                var v46 uint32 = 4294967295
                _ = fi44
                _ = v46
            }
        }
        var v47 int = 963
        _ = v47
    }
    var v48 int8 = 100
    for fi49 := 0; fi49 < 6; fi49++ {
        if true {
            v26 = (v26 << ((65535) & 63))
            for fi50 := 0; fi50 < 18; fi50++ {
                v51 := v1[42]
                _ = fi50
                _ = v51
            }
            for fi52 := 0; fi52 < 8; fi52++ {
                fmt.Println(v2, v26, v2)
                if true {
                    v5 = v5
                }
                v3 = 65535
                _ = fi52
            }
        } else {
            var v53 bool = false
            _ = v53
        }
        if true {
            v54 := v1[42]
            if (!false) {
                var v55 float32 = -2.25
                var v56 float32 = float32((float64(float32((3.14159 * -2.25))) / 1e10))
                _ = v55
                _ = v56
            }
            if false {
                if (!((v26 < 0) || ((("a" + "") + ("\n\t" + ("" + ""))) == ""))) {
                    fmt.Println(v3)
                    v3 = v3
                }
                if (4294967295 != (v54 << ((65535) & 63))) {
                    var v57 int16 = 0
                    v26 += (v26 << (((v3 * uint(v54))) & 63))
                    _ = v57
                } else {
                    for fi58 := 0; fi58 < 2; fi58++ {
                        if ((42 >= 0) || (42 == 1)) {
                            if (!(!(false && true))) {
                                var v59 uint64 = 42
                                var v60 int32 = (v26 / (((v26 - (v26 >> ((v3) & 63)))) | 1))
                                fmt.Println(v60, v3, v5)
                                v61 := GenS5{F0: v59}
                                _ = v59
                                _ = v60
                                _ = v61
                            } else {
                                v5 += v5
                                if true {
                                    switch (fi58 ^ -58) {
                                    case 1:
                                        if (v2 <= v2) {
                                            var v62 int8 = 100
                                            v63 := GenS3{F0: (^(v2))}
                                            var v64 int64 = (v5 | (^(v5)))
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                        } else {
                                            fmt.Println(v2)
                                            v65 := GenS6{F0: v3, F1: 1, F2: v26}
                                            switch (fi58 << (((v3 ^ (v3 % (((^(v3))) | 1)))) & 63)) {
                                            case 282:
                                                if false {
                                                    switch 2147483647 {
                                                    case -2147483648:
                                                        var v66 uint64 = 4294967295
                                                        _ = v66
                                                    case -409:
                                                        v4 = append(v4, "a")
                                                    case -155:
                                                        switch "foo bar" {
                                                        case "a":
                                                            var v67 int64 = (v5 * 996)
                                                            _ = v67
                                                        default:
                                                            var v68 float64 = float64((1e10 - 1.5))
                                                            _ = v68
                                                        }
                                                    }
                                                } else {
                                                    var v69 float32 = (3.14159 / -2.25)
                                                    v5 = v5
                                                    _ = v69
                                                }
                                            case 32767:
                                                var v70 int32 = v26
                                                _ = v70
                                            default:
                                                v71 := GenS5{F0: 0}
                                                _ = v71
                                            }
                                            _ = v65
                                        }
                                    }
                                }
                                var v72 float64 = 0.0
                                _ = v72
                            }
                        }
                        v1[v54] = v54
                        v73 := v1[v54]
                        _ = fi58
                        _ = v73
                    }
                    fmt.Println(v5)
                }
            }
            _ = v54
        }
        _ = fi49
    }
    var v74 int32 = (^(v26))
    var v75 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v74
    _ = v75
}
