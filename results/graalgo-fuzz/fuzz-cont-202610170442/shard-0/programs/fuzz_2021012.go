package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 int64; F2 int32 }
type GenS6 struct { F0 uint64; F1 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = float32((3.14159 * 1.5))
    v3 := make(map[string]int)
    var v4 uint32 = 42
    var v5 int32 = 42
    var v6 bool = false
    v6 = (true && v6)
    for fi7 := 0; fi7 < 18; fi7++ {
        v5 = 738
        var v8 string = ((("" + "") + ("hello" + "\n\t")) + ("hello" + "a"))
        for fi9 := 0; fi9 < 18; fi9++ {
            v2 = v2
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    switch "\n\t" {
    case "\n\t":
        var v10 int64 = 221
        _ = v10
    default:
        if (v6 || v6) {
            v4 = 0
            var v11 uint = 65535
            for fi12 := 0; fi12 < 18; fi12++ {
                for fi13 := 0; fi13 < 3; fi13++ {
                    switch fi13 {
                    case 731:
                        fmt.Println(fi13)
                    case 2147483647:
                        v14 := GenS1{F0: 0, F1: 0}
                        _ = v14
                    default:
                        var v15 int = (^(fi13))
                        _ = v15
                    }
                    if v6 {
                        v6 = false
                        if (-846 > 0) {
                            switch fi12 {
                            case 128:
                                switch (fi12 >> (((^(v11))) & 63)) {
                                case 9223372036854775807:
                                    switch "hello" {
                                    case "\n\t":
                                        v11 -= (v11 | (^(v11)))
                                    }
                                case 42:
                                    var v16 int16 = 32767
                                    v5 = v5
                                    v17 := GenS1{F0: 0, F1: 0}
                                    v1[("foo bar" + ("hello" + "a"))] = (42 <= (v11 % (((v11 % ((v11) | 1))) | 1)))
                                    var v18 int64 = 0
                                    if ((v6 && true) || v6) {
                                        v18 += v18
                                        v19 := GenS1{F0: 9223372036854775807, F1: -128}
                                        _ = v19
                                    } else {
                                        if (!false) {
                                            v6 = ((-128 == 0) && (!(((v2 * float32(0.1)) - v2) >= 1.5)))
                                            delete(v3, ("hello" + ""))
                                        } else {
                                            v20 := GenS1{F0: (-(v18)), F1: -55}
                                            v6 = false
                                            v18 = v18
                                            _ = v20
                                        }
                                        v16 += (-(v16))
                                    }
                                    if true {
                                        if (v2 != ((v2 - 0.0) - 0.0)) {
                                            var v21 int8 = int8(v18)
                                            var v22 uint = v11
                                            _ = v21
                                            _ = v22
                                        } else {
                                            switch (fi13 & (fi13 * -1)) {
                                            case 32768:
                                                var v23 int = fi13
                                                _ = v23
                                            case 128:
                                                if (true || true) {
                                                    v5 -= int32(v18)
                                                    fmt.Println(v6, v2)
                                                    if (0 > 255) {
                                                        v24 := []int{(fi13 | (fi12 - 127)), (-(fi13))}
                                                        v16 = 42
                                                        var v25 int64 = (v18 ^ -9223372036854775808)
                                                        _ = v24
                                                        _ = v25
                                                    } else {
                                                        v18 += v18
                                                        fmt.Println(v16, v4)
                                                        v26 := GenS2{F0: v11, F1: 3}
                                                        _ = v26
                                                    }
                                                    var v27 uint64 = 129
                                                    v2 = -2.25
                                                    _ = v27
                                                } else {
                                                    v16 = 174
                                                    var v28 uint16 = 1
                                                    fmt.Println(v28)
                                                    v29 := GenS1{F0: (-(v18)), F1: -9}
                                                    _ = v28
                                                    _ = v29
                                                }
                                            }
                                        }
                                        if (true && false) {
                                            var v30 int64 = int64(v5)
                                            v30 = (v30 | int64(v11))
                                            v5 = (v5 % (((v5 | 0)) | 1))
                                            if (v4 > v4) {
                                                if v6 {
                                                    var v31 uint32 = v4
                                                    v32 := GenS3{F0: -128, F1: (v18 + v18)}
                                                    var v33 float64 = 1.5
                                                    v11 += (v11 / ((v11) | 1))
                                                    _ = v31
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    v34 := []int{(fi12 | -1), (fi13 + (fi12 * (fi12 - (fi12 % ((fi12) | 1))))), fi12}
                                                    fmt.Println(v6, v6, v2)
                                                    _ = v34
                                                }
                                            }
                                            _ = v30
                                        } else {
                                            var v35 int16 = (v16 % ((v16) | 1))
                                            _ = v35
                                        }
                                        v16 = (v16 | -204)
                                        v4 = v4
                                        v4 = (^(v4))
                                        var v36 int16 = v16
                                        _ = v36
                                    } else {
                                        if false {
                                            var v37 float32 = v2
                                            v38 := []int{(^(fi12)), int(v16), (fi13 / ((876) | 1))}
                                            var v39 uint16 = 65535
                                            fmt.Println(fi12, v5, v37)
                                            var v40 uint8 = 84
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        } else {
                                            var v41 uint64 = 2
                                            if (!v6) {
                                                var v42 string = "hello"
                                                _ = v42
                                            } else {
                                                var v43 int8 = int8(fi13)
                                                var v44 uint16 = 448
                                                _ = v43
                                                _ = v44
                                            }
                                            v45 := GenS1{F0: v18, F1: -1}
                                            _ = v41
                                            _ = v45
                                        }
                                        switch "hello" {
                                        case "":
                                            v46 := []int{}
                                            _ = v46
                                        case "\n\t":
                                            v6 = v6
                                        case "foo bar":
                                            fmt.Println(v16, v18, fi13)
                                        }
                                    }
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                            default:
                                v47 := []int{-9223372036854775808, fi13, (fi13 + (fi13 << ((v11) & 63))), 0}
                                _ = v47
                            }
                            var v48 int64 = 1
                            _ = v48
                        } else {
                            switch "foo bar" {
                            case "\n\t":
                                v49 := GenS2{F0: v11, F1: 2}
                                _ = v49
                            }
                        }
                        v11 += v11
                        switch ("foo bar" + (("foo bar" + "hello") + (("hello" + "hello") + ("hello" + "hello")))) {
                        case "foo bar":
                            fmt.Println(v6, v5)
                        case "hello":
                            v11 += v11
                        }
                        switch "\n\t" {
                        case "hello":
                            var v50 uint16 = 65535
                            v51 := GenS4{F0: 898}
                            _ = v50
                            _ = v51
                        case "":
                            fmt.Println(v4, v2)
                        }
                    } else {
                        var v52 int8 = 127
                        v52 = (v52 % (((v52 / (((v52 / (((^(v52))) | 1))) | 1))) | 1))
                        _ = v52
                    }
                    v5 = (v5 >> (((v11 ^ v11)) & 63))
                    _ = fi13
                }
                v6 = v6
                for fi53 := 0; fi53 < 7; fi53++ {
                    var v54 uint = v11
                    _ = fi53
                    _ = v54
                }
                _ = fi12
            }
            _ = v11
        }
    }
    switch ("hello" + ("\n\t" + ("foo bar" + ""))) {
    case "foo bar":
        fmt.Println(v5)
    case "":
        var v55 float64 = -2.25
        _ = v55
    }
    for fi56 := 0; fi56 < 14; fi56++ {
        v5 -= (v5 % ((v5) | 1))
        _ = fi56
    }
    var v57 int32 = v5
    fmt.Println(v57)
    for fi58 := 0; fi58 < 13; fi58++ {
        v5 += -372
        switch (-(fi58)) {
        case -2147483648:
            fmt.Println(fi58, fi58)
        default:
            switch (fi58 % ((fi58) | 1)) {
            case 127:
                var v59 int8 = 75
                _ = v59
            case 0:
                v60 := GenS5{F0: (1 >= 1), F1: -9223372036854775808, F2: v57}
                _ = v60
            }
        }
        _ = fi58
    }
    if false {
        v57 = v57
        v4 = v4
        var v61 uint = 255
        _ = v61
    } else {
        var v62 uint64 = 107
        if v6 {
            v63 := GenS2{F0: 245, F1: uint32(v62)}
            _ = v63
        }
        if v6 {
            v57 = v57
        }
        _ = v62
    }
    if (9223372036854775807 == 42) {
        fmt.Println(v6, v57)
    }
    v5 = (v57 & -2147483648)
    fmt.Println(v57, v2)
    switch 4294967295 {
    case 42:
        fmt.Println(v6, v5)
    }
    if v6 {
        v64 := []int{127}
        for fi65 := 0; fi65 < 10; fi65++ {
            v66 := GenS6{F0: 18446744073709551615, F1: (0.0 > ((-2.25 + 0.1) / 1.5))}
            _ = fi65
            _ = v66
        }
        _ = v64
    }
    var v67 bool = v6
    var v68 uint32 = (v4 ^ v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v67
    _ = v68
}
