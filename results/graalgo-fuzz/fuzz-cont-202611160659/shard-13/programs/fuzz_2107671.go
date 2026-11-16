package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 bool; F2 int8 }
type GenS3 struct { F0 int16; F1 int8; F2 uint16 }
type GenS4 struct { F0 int32; F1 string; F2 bool }
type GenS5 struct { F0 int64; F1 int16; F2 int8 }
func genFn0(p0 bool) int32 {
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 bool = false
        _ = fi4
        _ = v5
    }
    if (0 >= 549) {
        switch "hello" {
        case "\n\t":
            p0 = (("foo bar" > "a") && p0)
        case "hello":
            switch "" {
            case "foo bar":
                p0 = (18446744073709551615 <= 0)
            default:
                if p0 {
                    if (p0 || p0) {
                        var v6 int = 1
                        fmt.Println(p0, p0)
                        _ = v6
                    }
                    p0 = false
                }
            }
        default:
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 int = (fi7 & fi7)
                _ = fi7
                _ = v8
            }
        }
    } else {
        if true {
            var v9 uint64 = 42
            _ = v9
        }
        var v10 int = 32767
        _ = v10
    }
    var v11 int16 = 1
    for fi12 := 0; fi12 < 19; fi12++ {
        switch fi12 {
        case -149:
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := []int{}
                if p0 {
                    if p0 {
                        var v15 int32 = 242
                        v16 := GenS1{F0: 65535}
                        if len(v14) > 0 {
                            v14[0] = (fi12 * (fi12 - fi13))
                        }
                        v15 += v15
                        var v17 uint16 = 42
                        _ = v15
                        _ = v16
                        _ = v17
                    } else {
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                    }
                }
                if true {
                    v18 := GenS1{F0: 2147483648}
                    v19 := GenS2{F0: 1, F1: p0, F2: -5}
                    var v20 int64 = 812
                    if p0 {
                        var v21 uint64 = 42
                        v20 -= (v20 * v20)
                        _ = v21
                    } else {
                        p0 = p0
                        v22 := GenS2{F0: 616, F1: p0, F2: -1}
                        v14 = append(v14, fi12)
                        _ = v22
                    }
                    _ = v18
                    _ = v19
                    _ = v20
                }
                fmt.Println(v11, fi12)
                var v23 uint64 = 4294967295
                _ = fi13
                _ = v14
                _ = v23
            }
        default:
            if p0 {
                var v24 int16 = v11
                fmt.Println(fi12, v11, v11)
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 7; fi25++ {
            if (p0 || ((fi12 + (^(fi25))) != (fi12 & fi25))) {
                v26 := GenS1{F0: 2147483647}
                _ = v26
            } else {
                v11 = (v11 << ((4294967295) & 63))
            }
            _ = fi25
        }
        var v27 int32 = 0
        _ = fi12
        _ = v27
    }
    if (2 == 847) {
        var v28 int64 = -141
        v28 -= v28
        for fi29 := 0; fi29 < 3; fi29++ {
            if true {
                v30 := GenS1{F0: 2}
                fmt.Println(v28)
                _ = v30
            } else {
                switch "a" {
                case "":
                    var v31 int8 = 1
                    _ = v31
                default:
                    v32 := []int{(^(fi29)), (fi29 - (fi29 % ((fi29) | 1))), (fi29 ^ -430)}
                    _ = v32
                }
            }
            _ = fi29
        }
        _ = v28
    } else {
        if (true || false) {
            v33 := []int{756}
            v11 += (^(v11))
            _ = v33
        }
        for fi34 := 0; fi34 < 14; fi34++ {
            var v35 int8 = -128
            _ = fi34
            _ = v35
        }
    }
    _ = p0
    _ = v11
    return 42
}

func main() {
    var v1 int8 = 0
    var v2 float64 = (1.5 + -2.25)
    var v3 string = ""
    v36 := []int{4294967295}
    _ = len(v36)
    for fi37 := 0; fi37 < 15; fi37++ {
        v1 = (v1 ^ v1)
        _ = fi37
    }
    v38 := GenS3{F0: -32768, F1: v1, F2: 2}
    switch v3 {
    case "hello":
        for fi39 := 0; fi39 < 3; fi39++ {
            for fi40 := 0; fi40 < 7; fi40++ {
                var v41 int16 = 607
                var v42 bool = true
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = fi39
        }
    }
    var v43 int8 = 100
    for fi44 := 0; fi44 < 20; fi44++ {
        var v45 string = ("hello" + "\n\t")
        switch fi44 {
        case -9223372036854775808:
            fmt.Println(fi44, fi44)
        case 127:
            var v46 int8 = -128
            _ = v46
        case 95:
            switch (fi44 >> ((0) & 63)) {
            case 2147483647:
                v36 = append(v36, -120)
            case 566:
                var v47 int64 = 773
                _ = v47
            default:
                if len(v36) >= 2 {
                    v36 = v36[1:len(v36)]
                }
            }
        }
        _ = fi44
        _ = v45
    }
    v43 = (^(v43))
    for fi48 := 0; fi48 < 9; fi48++ {
        v49 := GenS4{F0: genFn0((1.5 <= 1.5)), F1: "", F2: (4294967295 < 0)}
        _ = fi48
        _ = v49
    }
    if false {
        if true {
            switch (v3 + "") {
            case "a":
                for fi50 := 0; fi50 < 4; fi50++ {
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                    fmt.Println(v1)
                    _ = fi50
                }
            default:
                for fi51 := 0; fi51 < 1; fi51++ {
                    if (true || true) {
                        if (582 != 989) {
                            var v52 uint = 827
                            _ = v52
                        }
                    } else {
                        fmt.Println(v43)
                        var v53 int16 = 581
                        _ = v53
                    }
                    var v54 bool = ((true || true) || (32767 > -26))
                    _ = fi51
                    _ = v54
                }
            }
        } else {
            switch 2147483647 {
            case 4294967295:
                v55 := GenS5{F0: -9223372036854775808, F1: 725, F2: v1}
                _ = v55
            case 9223372036854775807:
                switch "" {
                case "a":
                    v3 = v3
                case "":
                    if ((!true) || false) {
                        v56 := GenS5{F0: 9223372036854775807, F1: 32767, F2: -1}
                        if false {
                            var v57 int = 381
                            _ = v57
                        }
                        _ = v56
                    } else {
                        _ = len(v36)
                        v36 = append(v36, 32767)
                        fmt.Println(v3)
                    }
                case "foo bar":
                    v43 = v43
                }
            case -1:
                for fi58 := 0; fi58 < 9; fi58++ {
                    switch (fi58 - 127) {
                    case -237:
                        _ = len(v36)
                    case -326:
                        fmt.Println(fi58)
                    case -9223372036854775808:
                        fmt.Println(v2)
                    default:
                        fmt.Println(v3)
                    }
                    v2 -= (v2 * 0.1)
                    fmt.Println(v1, v2, v3)
                    _ = fi58
                }
            }
            v1 = v43
            var v59 float64 = 0.1
            _ = v59
        }
        switch ((v3 + v3) + v3) {
        case "a":
            var v60 uint16 = 0
            _ = v60
        case "foo bar":
            var v61 uint32 = 1
            _ = v61
        case "hello":
            v2 += float64(3.14159)
        default:
            var v62 uint32 = 728
            _ = v62
        }
    } else {
        for fi63 := 0; fi63 < 15; fi63++ {
            var v64 int16 = 259
            v2 -= ((3.14159 - -2.25) * (v2 * 0.1))
            fmt.Println(v1)
            _ = fi63
            _ = v64
        }
    }
    v2 += v2
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v43
}
