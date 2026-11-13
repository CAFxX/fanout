package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 int16; F2 int64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 1
    var v3 uint32 = 42
    v4 := make(map[string]string)
    if false {
        v5 := GenS1{F0: 1, F1: 4294967295, F2: uint64(v3)}
        if ((0.0 / 1.5) != -2.25) {
            var v6 uint32 = v3
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 uint16 = 64
                var v9 int16 = -645
                _ = fi7
                _ = v8
                _ = v9
            }
            switch "" {
            case "":
                v3 = (v3 ^ 906)
            }
            _ = v6
        } else {
            switch -2147483648 {
            case -164:
                if false {
                    var v10 uint64 = 0
                    var v11 uint32 = 2147483648
                    var v12 int32 = 385
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    var v13 int = -879
                    _ = v13
                }
            case -9223372036854775808:
                v14 := []int{-480}
                _ = v14
            case 2147483647:
                var v15 uint32 = v3
                _ = v15
            }
            switch -10 {
            case 790:
                for fi16 := 0; fi16 < 14; fi16++ {
                    v17 := []int{(fi16 - fi16)}
                    v2 = (v2 + (v2 ^ v2))
                    if true {
                        var v18 string = ("hello" + "")
                        v2 += (v2 + 83)
                        if (1 == 89) {
                            v18 = "foo bar"
                            v3 -= v3
                            v2 = v2
                        }
                        _ = v18
                    }
                    _ = fi16
                    _ = v17
                }
            default:
                fmt.Println(v3, v2)
            }
        }
        v19 := []int{9223372036854775807, 998, -2147483648}
        _ = v19
        _ = v5
    }
    switch 0 {
    case 32767:
        v2 = (v2 - (-(v2)))
    case 985:
        var v20 uint16 = 65535
        _ = v20
    }
    var v21 uint16 = 607
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 int16 = -871
        v24 := GenS1{F0: 1, F1: 4294967295, F2: 18446744073709551615}
        var v25 int64 = -371
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    if true {
        v3 += v3
        v21 += v21
    }
    var v26 string = "\n\t"
    var v27 int = 0
    if ((false && false) && false) {
        if (0 <= 127) {
            var v28 uint32 = v3
            v29 := []int{v27, -2147483648}
            _ = v28
            _ = v29
        }
        for fi30 := 0; fi30 < 9; fi30++ {
            switch fi30 {
            case 32768:
                v31 := GenS1{F0: 723, F1: 0, F2: 18446744073709551615}
                _ = v31
            default:
                v32 := []int{v27, (^(fi30)), fi30, -843}
                _ = v32
            }
            _ = fi30
        }
        var v33 uint8 = 0
        _ = v33
    }
    var v34 int8 = (v2 * -8)
    switch v26 {
    case "a":
        v35 := GenS1{F0: 4294967295, F1: 4294967295, F2: 42}
        _ = v35
    case "foo bar":
        v1[(v3 << ((2) & 63))] = ("foo bar" + "foo bar")
    case "hello":
        v27 += (v27 & (v27 << ((65535) & 63)))
    }
    switch (^(v27)) {
    case 1:
        var v36 int8 = (v34 ^ v2)
        _ = v36
    case 9223372036854775807:
        if ((^(v21)) > v21) {
            v21 += 609
            fmt.Println(v3)
            v37 := GenS1{F0: 517, F1: 4294967295, F2: uint64(v3)}
            _ = v37
        } else {
            fmt.Println(v2, v27, v34)
        }
    default:
        for fi38 := 0; fi38 < 12; fi38++ {
            if ((!true) || (true || (4294967295 > 1))) {
                v39 := GenS1{F0: 42, F1: 42, F2: 4294967295}
                _ = v39
            } else {
                v21 = (v21 * v21)
                fmt.Println(v21)
                if (false || (((("" + v26) + "a") + v26) > (v26 + (v26 + (v26 + "\n\t"))))) {
                    for fi40 := 0; fi40 < 19; fi40++ {
                        switch (v27 | v27) {
                        case 32767:
                            v26 = ("\n\t" + v26)
                        case 128:
                            v34 = (v34 >> ((1) & 63))
                        }
                        _ = fi40
                    }
                    v41 := []int{}
                    _ = v41
                }
            }
            for fi42 := 0; fi42 < 4; fi42++ {
                switch v26 {
                case "a":
                    v43 := GenS2{F0: v3, F1: 2}
                    _ = v43
                case "\n\t":
                    v44 := GenS3{F0: false, F1: (v21 + v21), F2: 42}
                    _ = v44
                default:
                    v3 += (^(v3))
                }
                if false {
                    var v45 int = int(v34)
                    v46 := []int{fi38, -2147483648, v27, -2147483648}
                    v45 = v45
                    _ = v45
                    _ = v46
                }
                var v47 int = v27
                var v48 uint64 = 2
                _ = fi42
                _ = v47
                _ = v48
            }
            _ = fi38
        }
    }
    if ((799 >= 1) && ((false && (0 > 2)) || ((v3 >> ((2) & 63)) > uint32(v21)))) {
        var v49 uint16 = v21
        var v50 string = "hello"
        _ = v49
        _ = v50
    }
    if true {
        v2 -= v2
        v51 := GenS4{F0: false}
        v27 += v27
        _ = v51
    }
    v52 := v4["a"]
    if (100 < -2) {
        var v53 float64 = ((((0.0 - 1.5) - float64(-2.25)) / 1.5) + -2.25)
        _ = v53
    } else {
        switch (v27 - 211) {
        case 32767:
            v54 := GenS5{F0: 120}
            _ = v54
        case 1:
            var v55 int16 = -87
            _ = v55
        }
        switch "a" {
        case "\n\t":
            v56 := GenS6{F0: 128, F1: 0, F2: 0}
            _ = v56
        default:
            switch int(v3) {
            case 2147483647:
                v3 = v3
            case 9223372036854775807:
                fmt.Println(v21)
            case 32768:
                var v57 int32 = -2147483648
                _ = v57
            }
        }
        v26 = (v26 + (v52 + ""))
    }
    if false {
        fmt.Println(v3, v26, v2)
        for fi58 := 0; fi58 < 12; fi58++ {
            var v59 uint8 = uint8(v3)
            var v60 int16 = 0
            _ = fi58
            _ = v59
            _ = v60
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v52
}
