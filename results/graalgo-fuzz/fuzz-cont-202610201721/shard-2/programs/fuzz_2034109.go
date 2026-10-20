package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 uint16; F1 int8; F2 int }
func genFn0() int32 {
    v6 := []int{2147483647, 625, 32768}
    switch "foo bar" {
    case "a":
        v6 = append(v6, 1)
    default:
        v6 = append(v6, 9223372036854775807)
    }
    var v7 uint16 = 788
    _ = v6
    _ = v7
    return -807
}

func main() {
    var v1 int64 = -28
    v2 := []uint32{}
    var v3 uint32 = 2147483648
    v4 := make(map[string]uint32)
    var v5 int32 = -33
    fmt.Println(v1)
    v5 = (v5 / (((v5 % (((v5 - -2147483648)) | 1))) | 1))
    v5 += v5
    if ((-1 == 878) && (32767 != -1)) {
        v8 := v4["foo bar"]
        v9 := GenS1{F0: 2147483647, F1: v1, F2: 221}
        _ = v8
        _ = v9
    } else {
        var v10 int8 = 1
        var v11 uint8 = 174
        var v12 int8 = 34
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 9223372036854775807, F1: (v1 + v1), F2: 1}
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := GenS1{F0: (fi14 - fi14), F1: -9223372036854775808, F2: 2}
        switch "foo bar" {
        case "hello":
            fmt.Println(fi14)
        case "\n\t":
            switch fi14 {
            case 128:
                if false {
                    for fi16 := 0; fi16 < 19; fi16++ {
                        v3 = v3
                        if false {
                            switch fi14 {
                            case 369:
                                var v17 bool = (!((false || (true || (1 == 100))) && false))
                                _ = v17
                            case 127:
                                var v18 string = ("a" + "foo bar")
                                _ = v18
                            default:
                                v19 := GenS1{F0: (fi16 | (fi16 ^ fi16)), F1: v1, F2: 0}
                                _ = v19
                            }
                            v20 := GenS1{F0: fi16, F1: int64(v5), F2: 2}
                            _ = v20
                        }
                        _ = fi16
                    }
                }
            case 960:
                var v21 bool = (v3 >= 2)
                _ = v21
            case 4294967295:
                v3 = 0
            default:
                switch "a" {
                case "":
                    v3 += (^(v3))
                case "a":
                    var v22 bool = true
                    _ = v22
                default:
                    v5 = v5
                }
            }
        }
        v23 := GenS1{F0: fi14, F1: 42, F2: 449}
        _ = fi14
        _ = v15
        _ = v23
    }
    v24 := GenS2{F0: genFn0(), F1: v1}
    fmt.Println(v5, v1, v5)
    _ = len(v4)
    v25 := GenS3{F0: v5}
    if ((270 <= 32767) || false) {
        var v26 string = ""
        v3 = v3
        _ = v26
    } else {
        var v27 bool = true
        for fi28 := 0; fi28 < 18; fi28++ {
            for fi29 := 0; fi29 < 1; fi29++ {
                v1 -= 0
                if v27 {
                    v30 := GenS4{F0: (-(fi29))}
                    switch fi29 {
                    case -1:
                        var v31 string = ""
                        _ = v31
                    case -9223372036854775808:
                        var v32 float32 = ((0.0 - float32(1.5)) - (1.5 / 1e10))
                        _ = v32
                    }
                    v5 -= 1
                    var v33 uint8 = 42
                    v27 = true
                    _ = v30
                    _ = v33
                }
                if len(v2) > 0 {
                    v2[0] = 0
                }
                _ = fi29
            }
            v2 = append(v2, (^(v3)))
            _ = fi28
        }
        _ = v27
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        for fi35 := 0; fi35 < 11; fi35++ {
            v36 := GenS5{F0: 819, F1: 42}
            _ = fi35
            _ = v36
        }
        if (32767 != 255) {
            if true {
                for fi37 := 0; fi37 < 20; fi37++ {
                    var v38 uint16 = 65535
                    v1 = 431
                    v38 -= (^(v38))
                    _ = fi37
                    _ = v38
                }
                for fi39 := 0; fi39 < 20; fi39++ {
                    v5 = genFn0()
                    _ = fi39
                }
            } else {
                var v40 bool = true
                fmt.Println(v5, v40, v40)
                if (((("" + "hello") + "a") + (("hello" + ("hello" + "\n\t")) + ("\n\t" + ("hello" + "\n\t")))) < "hello") {
                    if false {
                        var v41 uint8 = 207
                        _ = v41
                    }
                    var v42 uint64 = 42
                    _ = v42
                }
                _ = v40
            }
            var v43 int64 = -1
            for fi44 := 0; fi44 < 11; fi44++ {
                v5 = (v5 - 0)
                _ = fi44
            }
            _ = v43
        } else {
            v45 := GenS6{F0: 0, F1: 0, F2: 127}
            _ = v45
        }
        for fi46 := 0; fi46 < 18; fi46++ {
            fmt.Println(v5, v1, v5)
            var v47 uint8 = 0
            switch (fi46 / ((int(v3)) | 1)) {
            case 32768:
                switch (^(fi34)) {
                case 128:
                    var v48 float64 = float64(0.1)
                    _ = v48
                }
                v49 := GenS6{F0: 736, F1: 42, F2: fi46}
                fmt.Println(v1)
                _ = v49
            }
            var v50 int = 1
            switch (fi34 >> ((2) & 63)) {
            case 2147483647:
                if (true || (!false)) {
                    switch "" {
                    case "hello":
                        switch 2147483647 {
                        case -827:
                            var v51 int16 = -32768
                            _ = v51
                        case 32768:
                            v52 := v4[""]
                            _ = v52
                        default:
                            v3 = (v3 | 2)
                        }
                    }
                    fmt.Println(v3, fi46)
                }
            case 42:
                v2 = append(v2, 4294967295)
            case -1:
                fmt.Println(fi34)
            default:
                v1 -= v1
            }
            _ = fi46
            _ = v47
            _ = v50
        }
        _ = fi34
    }
    if (false && true) {
        fmt.Println(v3, v1, v5)
        _ = len(v2)
        v53 := GenS6{F0: 65535, F1: -1, F2: -643}
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
