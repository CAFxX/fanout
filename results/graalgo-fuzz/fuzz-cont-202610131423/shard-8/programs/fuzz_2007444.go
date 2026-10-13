package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint16; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 uint; F2 string }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int; F1 string }
type GenS11 struct { F0 int }

func main() {
    var v1 uint8 = 128
    v2 := []string{}
    var v3 string = ("hello" + "foo bar")
    v4 := []int{-2147483648, 9223372036854775807}
    var v5 uint32 = 138
    var v6 int64 = -14
    fmt.Println(v1, v6, v3)
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        var v7 int = 127
        _ = len(v2)
        for fi8 := 0; fi8 < 16; fi8++ {
            if (true && (4294967295 < 1)) {
                switch 2147483647 {
                case -286:
                    if ((^(fi8)) == (^(v7))) {
                        if ((false || true) || ((false || false) && false)) {
                            fmt.Println(v7, fi8, v1)
                            _ = len(v4)
                        } else {
                            fmt.Println(v5, v5, v5)
                        }
                        for fi9 := 0; fi9 < 2; fi9++ {
                            var v10 float32 = ((((0.1 * 0.0) * 1.5) - ((-2.25 - 1.5) - float32(3.14159))) / 1.5)
                            var v11 uint = 65535
                            v12 := GenS1{F0: -448, F1: (v6 | v6), F2: 42}
                            _ = fi9
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                    }
                default:
                    v13 := GenS2{F0: v5, F1: 218, F2: v3}
                    _ = v13
                }
            } else {
                v6 = v6
                var v14 bool = false
                _ = v14
            }
            v15 := GenS3{F0: ((v1 != v1) && (((1.5 - (1.5 + (1e10 / 1.0))) * ((0.0 / 1e10) - float64(0.0))) < 0.0))}
            _ = fi8
            _ = v15
        }
        _ = v7
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 uint32 = v5
            switch (-(fi16)) {
            case -2147483648:
                v1 -= 2
            default:
                if true {
                    var v19 int32 = 0
                    v1 += v1
                    _ = v19
                }
            }
            v1 = (v1 + 1)
            var v20 int16 = -32768
            _ = fi17
            _ = v18
            _ = v20
        }
        fmt.Println(v1, v3)
        var v21 int16 = 0
        _ = fi16
        _ = v21
    }
    v5 += v5
    fmt.Println(v1, v3)
    for fi22 := 0; fi22 < 14; fi22++ {
        v3 = ("\n\t" + v3)
        for fi23 := 0; fi23 < 19; fi23++ {
            switch fi22 {
            case -260:
                v24 := GenS4{F0: 42}
                _ = v24
            case 210:
                var v25 float64 = ((0.1 - 1e10) + (-2.25 * (0.0 * 0.1)))
                _ = v25
            default:
                switch (fi22 + (fi23 | fi22)) {
                case 127:
                    if (32 >= -1) {
                        v6 -= v6
                        v1 += v1
                    }
                case -880:
                    if false {
                        v5 = 776
                    } else {
                        v26 := GenS5{F0: ("a" + "\n\t"), F1: 42, F2: v3}
                        v6 = (v6 | (v6 - 1))
                        v3 = v3
                        v6 += (v6 << ((4294967295) & 63))
                        _ = v26
                    }
                case -64:
                    v5 = v5
                }
            }
            if len(v4) > 0 {
                v4[0] = (-(fi22))
            }
            _ = fi23
        }
        _ = fi22
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi27
    }
    var v28 bool = true
    if ((!((-2.25 / 1.0) < 1e10)) && ((float32(-2.25) * 0.0) != float32(0.1))) {
        switch 32767 {
        case -9223372036854775808:
            switch "a" {
            case "foo bar":
                fmt.Println(v5)
            case "hello":
                var v29 uint8 = (v1 + v1)
                _ = v29
            default:
                for fi30 := 0; fi30 < 17; fi30++ {
                    v28 = (177 != 1)
                    v1 = (v1 << ((42) & 63))
                    v31 := GenS4{F0: 1}
                    _ = fi30
                    _ = v31
                }
            }
        case -497:
            for fi32 := 0; fi32 < 10; fi32++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v4 = append(v4, fi32)
                var v33 float32 = 1.5
                _ = fi32
                _ = v33
            }
        }
        if (!(((uint64(v1) >= 668) && (!(-1 < 42))) || (v28 || (true || v28)))) {
            switch 810 {
            case -323:
                for fi34 := 0; fi34 < 2; fi34++ {
                    v35 := GenS1{F0: 0, F1: (-(v6)), F2: 0}
                    v28 = ((false && v28) || (-1 >= -32768))
                    v36 := GenS6{F0: 161, F1: 32767}
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            case -9223372036854775808:
                var v37 int16 = 292
                _ = v37
            case 9223372036854775807:
                v38 := GenS2{F0: v5, F1: 0, F2: v3}
                _ = v38
            default:
                v39 := GenS7{F0: 179, F1: 0, F2: 39}
                _ = v39
            }
            var v40 string = v3
            _ = v40
        } else {
            _ = len(v2)
        }
    }
    v1 = 1
    switch 0 {
    case 0:
        var v41 int64 = (v6 | -1)
        _ = v41
    default:
        v1 = v1
    }
    if (-2.25 > 1.5) {
        if v28 {
            v1 = 221
            fmt.Println(v1, v6)
            if v28 {
                var v42 string = v3
                _ = v42
            }
        } else {
            v43 := GenS2{F0: (v5 * 65535), F1: 65535, F2: ""}
            v44 := GenS8{F0: 516}
            for fi45 := 0; fi45 < 16; fi45++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v46 := GenS3{F0: v28}
                _ = fi45
                _ = v46
            }
            _ = v43
            _ = v44
        }
    }
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v47 uint = 42
    if true {
        v47 = v47
    } else {
        for fi48 := 0; fi48 < 19; fi48++ {
            for fi49 := 0; fi49 < 9; fi49++ {
                if len(v2) > 0 {
                    v2[0] = (v3 + "")
                }
                fmt.Println(v6)
                v28 = true
                _ = fi49
            }
            _ = len(v4)
            _ = fi48
        }
        v2 = append(v2, (v3 + "foo bar"))
        for fi50 := 0; fi50 < 9; fi50++ {
            if (!(9223372036854775807 < (v6 * v6))) {
                var v51 uint16 = 680
                _ = v51
            } else {
                if ((v28 || ((v6 % ((v6) | 1)) < 270)) && v28) {
                    v52 := GenS6{F0: 702, F1: (fi50 - fi50)}
                    _ = len(v4)
                    fmt.Println(v3, v1)
                    _ = v52
                }
                v53 := GenS9{F0: (v5 + (v5 >> ((v47) & 63)))}
                v54 := GenS10{F0: fi50, F1: v3}
                _ = v53
                _ = v54
            }
            v55 := GenS11{F0: (fi50 - fi50)}
            _ = fi50
            _ = v55
        }
    }
    switch (v3 + "\n\t") {
    case "a":
        var v56 int64 = -460
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
