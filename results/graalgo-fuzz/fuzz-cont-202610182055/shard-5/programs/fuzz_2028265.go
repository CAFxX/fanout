package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint8; F2 uint }
type GenS3 struct { F0 bool; F1 bool; F2 uint64 }
type GenS4 struct { F0 int; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 string; F2 uint16 }
type GenS7 struct { F0 int32; F1 string }
func genFn0() uint8 {
    if ((3.14159 - 0.0) > 1e10) {
        var v4 float32 = 1.5
        if (-2147483648 >= -2147483648) {
            for fi5 := 0; fi5 < 10; fi5++ {
                v4 += (1.5 - (((v4 * 3.14159) * (v4 - 0.1)) / -2.25))
                _ = fi5
            }
            var v6 uint = 411
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 int64 = 1
                _ = fi7
                _ = v8
            }
            _ = v6
        }
        var v9 uint8 = 4
        _ = v4
        _ = v9
    }
    v10 := GenS1{F0: ("hello" + ("foo bar" + "\n\t"))}
    v11 := GenS2{F0: 9223372036854775807, F1: 47, F2: 43}
    if false {
        var v12 bool = (127 <= 76)
        switch -820 {
        case -2147483648:
            var v13 int = 1
            _ = v13
        }
        _ = v12
    } else {
        for fi14 := 0; fi14 < 5; fi14++ {
            for fi15 := 0; fi15 < 13; fi15++ {
                fmt.Println(fi14)
                fmt.Println(fi14, fi14)
                _ = fi15
            }
            var v16 float32 = (3.14159 * 1.5)
            _ = fi14
            _ = v16
        }
        var v17 int = 893
        v17 = (v17 & v17)
        _ = v17
    }
    _ = v10
    _ = v11
    return 255
}
func genFn1() uint16 {
    var v18 float64 = 1.5
    for fi19 := 0; fi19 < 7; fi19++ {
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 uint = 2
            var v22 uint64 = 0
            v23 := GenS2{F0: 9223372036854775807, F1: 2, F2: 42}
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        v18 = v18
        _ = fi19
    }
    if (false || false) {
        fmt.Println(v18)
    } else {
        switch "\n\t" {
        case "":
            fmt.Println(v18)
        case "a":
            for fi24 := 0; fi24 < 3; fi24++ {
                for fi25 := 0; fi25 < 2; fi25++ {
                    v26 := GenS1{F0: "\n\t"}
                    _ = fi25
                    _ = v26
                }
                v27 := GenS3{F0: true, F1: (!(!(!true))), F2: 0}
                _ = fi24
                _ = v27
            }
        default:
            var v28 bool = true
            _ = v28
        }
    }
    switch 1 {
    case 4294967295:
        var v29 float64 = 1e10
        _ = v29
    case 32768:
        v18 = float64(1e10)
    }
    v18 = v18
    _ = v18
    return 255
}

func main() {
    var v1 float32 = 1.5
    var v2 bool = ((0.0 / 1.0) >= (1e10 / float64(float32(1e10))))
    var v3 bool = (-386 != -224)
    v30 := []int{-443, -1, -9223372036854775808}
    for fi31 := 0; fi31 < 18; fi31++ {
        for fi32 := 0; fi32 < 9; fi32++ {
            v2 = ((1e10 + -2.25) >= 1e10)
            var v33 float32 = (float32(1e10) - v1)
            if (v2 || v3) {
                v2 = true
                var v34 uint32 = 65535
                _ = v34
            }
            _ = fi32
            _ = v33
        }
        switch (fi31 >> ((944) & 63)) {
        case -258:
            v35 := GenS4{F0: 0, F1: 2, F2: 4294967295}
            _ = v35
        case 4294967295:
            fmt.Println(v1)
        default:
            v2 = true
        }
        _ = fi31
    }
    var v36 int = -575
    fmt.Println(v36)
    if (!(-1 == -77)) {
        if true {
            v3 = ((535 == 752) || (1.5 > -2.25))
        } else {
            switch (v36 - v36) {
            case 32767:
                for fi37 := 0; fi37 < 13; fi37++ {
                    v38 := GenS5{F0: 42}
                    _ = fi37
                    _ = v38
                }
            case -9223372036854775808:
                v1 -= v1
            case -18:
                if (((42 != 255) || (789 <= 0)) && v3) {
                    var v39 int64 = 1
                    if v2 {
                        v40 := GenS3{F0: v3, F1: v3, F2: 849}
                        if v2 {
                            v41 := GenS5{F0: 18446744073709551615}
                            _ = v41
                        } else {
                            if len(v30) > 0 {
                                v30[0] = v36
                            }
                            v36 += (v36 ^ (v36 * v36))
                        }
                        _ = v40
                    }
                    for fi42 := 0; fi42 < 19; fi42++ {
                        v43 := GenS3{F0: v2, F1: v2, F2: 689}
                        _ = fi42
                        _ = v43
                    }
                    _ = v39
                } else {
                    fmt.Println(v3, v36, v2)
                    var v44 bool = v2
                    _ = v44
                }
            }
        }
        for fi45 := 0; fi45 < 17; fi45++ {
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            v2 = (4294967295 > 0)
            _ = fi45
        }
        v36 += (-(v36))
    }
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    if false {
        v3 = (("a" + "foo bar") < "\n\t")
        for fi46 := 0; fi46 < 3; fi46++ {
            var v47 float64 = 3.14159
            if v3 {
                switch "" {
                case "foo bar":
                    v48 := GenS5{F0: 803}
                    _ = v48
                case "a":
                    v30 = append(v30, (^(v36)))
                case "\n\t":
                    var v49 int32 = -1
                    _ = v49
                }
                if false {
                    v47 = v47
                }
                fmt.Println(v2)
            }
            _ = fi46
            _ = v47
        }
    }
    if false {
        v2 = false
        var v50 float64 = (1.5 * float64(-2.25))
        _ = v50
    }
    if true {
        v30 = append(v30, v36)
        if v2 {
            if len(v30) > 0 {
                v30[0] = v36
            }
            v30 = append(v30, v36)
        } else {
            for fi51 := 0; fi51 < 1; fi51++ {
                var v52 int64 = -1
                _ = fi51
                _ = v52
            }
        }
    } else {
        switch "foo bar" {
        case "hello":
            var v53 uint = 255
            _ = v53
        case "foo bar":
            var v54 int16 = -678
            _ = v54
        default:
            v1 = float32(-2.25)
        }
        if (false || (!v2)) {
            var v55 float64 = -2.25
            _ = v55
        }
        v2 = (!(!(v2 || v3)))
    }
    if (32767 >= -1) {
        for fi56 := 0; fi56 < 7; fi56++ {
            for fi57 := 0; fi57 < 15; fi57++ {
                if v2 {
                    v3 = v2
                    _ = len(v30)
                    var v58 uint8 = 0
                    _ = v58
                } else {
                    if (true && v2) {
                        v30 = append(v30, 9223372036854775807)
                    } else {
                        v3 = v3
                        _ = len(v30)
                    }
                    if len(v30) > 0 {
                        v30[0] = (fi57 << ((674) & 63))
                    }
                }
                v59 := GenS6{F0: (("foo bar" + "") + "\n\t"), F1: "foo bar", F2: 2}
                fmt.Println(fi57, fi57, fi56)
                v30 = append(v30, -2147483648)
                _ = fi57
                _ = v59
            }
            v36 += fi56
            _ = fi56
        }
        v36 = (v36 ^ v36)
        switch "" {
        case "foo bar":
            v1 = (3.14159 / (v1 - v1))
        default:
            v60 := GenS7{F0: 0, F1: "a"}
            _ = v60
        }
    }
    for fi61 := 0; fi61 < 17; fi61++ {
        var v62 int = v36
        v1 -= 1.5
        fmt.Println(fi61, v62, v3)
        _ = fi61
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
}
