package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int16; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 int16 }

func main() {
    var v1 int16 = 1000
    var v2 uint32 = 65535
    var v3 int8 = 0
    var v4 bool = ((true && (!true)) && (4294967295 <= 654))
    switch 127 {
    case 0:
        if v4 {
            fmt.Println(v1)
            var v5 uint = 0
            switch -165 {
            case 9223372036854775807:
                var v6 uint8 = 127
                _ = v6
            }
            _ = v5
        } else {
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := GenS1{F0: -948}
                _ = fi7
                _ = v8
            }
            fmt.Println(v3, v1)
            var v9 uint16 = 65535
            _ = v9
        }
    case 127:
        v10 := GenS1{F0: -1}
        _ = v10
    }
    v11 := GenS1{F0: -1}
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := []int{fi12, 540, 834, (fi12 % ((fi12) | 1))}
        for fi14 := 0; fi14 < 11; fi14++ {
            if false {
                switch "foo bar" {
                case "hello":
                    var v15 int64 = 559
                    _ = v15
                case "\n\t":
                    var v16 uint64 = 2
                    _ = v16
                case "foo bar":
                    v2 = v2
                }
                var v17 int16 = (v1 & int16(v3))
                var v18 float64 = 0.0
                _ = v17
                _ = v18
            } else {
                v4 = v4
                var v19 uint = 255
                var v20 int32 = -98
                var v21 uint64 = 300
                _ = v19
                _ = v20
                _ = v21
            }
            _ = fi14
        }
        if v4 {
            switch fi12 {
            case 42:
                if len(v13) > 0 {
                    v13[0] = fi12
                }
            }
            var v22 int64 = 42
            v2 = v2
            _ = v22
        } else {
            v23 := GenS1{F0: -328}
            _ = v23
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v2, v3)
    if false {
        for fi24 := 0; fi24 < 7; fi24++ {
            for fi25 := 0; fi25 < 17; fi25++ {
                if false {
                    if true {
                        var v26 float32 = 3.14159
                        switch fi24 {
                        case 128:
                            switch (("hello" + "foo bar") + "a") {
                            case "foo bar":
                                var v27 int16 = v1
                                _ = v27
                            case "a":
                                v4 = (-2.25 != (1e10 - ((1e10 + (0.0 / 3.14159)) + (float64((0.0 - (1.5 / v26))) + 1.5))))
                            default:
                                v28 := []int{912, -130}
                                _ = v28
                            }
                        case -433:
                            v2 = (v2 << ((18446744073709551615) & 63))
                        default:
                            if v4 {
                                fmt.Println(fi24, v26, v1)
                                v3 = 1
                            }
                        }
                        fmt.Println(v4, v1)
                        _ = v26
                    }
                    var v29 uint32 = v2
                    v30 := GenS1{F0: -2147483648}
                    _ = v29
                    _ = v30
                } else {
                    if ("hello" <= "\n\t") {
                        var v31 bool = (true || true)
                        var v32 int8 = (-(v3))
                        v33 := []int{(fi24 | fi24)}
                        var v34 int32 = 861
                        _ = v31
                        _ = v32
                        _ = v33
                        _ = v34
                    } else {
                        v2 -= (^(v2))
                        switch fi25 {
                        case 128:
                            var v35 int = -47
                            _ = v35
                        case 802:
                            var v36 int16 = 0
                            _ = v36
                        case 42:
                            v37 := []int{(fi25 | fi24), 9223372036854775807}
                            _ = v37
                        default:
                            v38 := []int{(fi24 & fi25), fi24}
                            _ = v38
                        }
                        v1 = 1
                        v39 := GenS1{F0: 357}
                        var v40 uint16 = 953
                        _ = v39
                        _ = v40
                    }
                    v3 += (v3 - v3)
                    var v41 uint8 = 1
                    _ = v41
                }
                if v4 {
                    var v42 int16 = (v1 / ((-32768) | 1))
                    _ = v42
                } else {
                    v43 := GenS1{F0: -2147483648}
                    if (((float32(3.14159) + 1e10) < float32(0.1)) || (v1 > 1000)) {
                        v44 := []int{}
                        v2 = (v2 ^ 642)
                        _ = v44
                    }
                    fmt.Println(v2, v2)
                    _ = v43
                }
                v45 := GenS2{F0: 12, F1: v1, F2: "foo bar"}
                var v46 uint8 = 127
                _ = fi25
                _ = v45
                _ = v46
            }
            _ = fi24
        }
    } else {
        v2 = v2
        v1 -= v1
    }
    if v4 {
        v47 := GenS1{F0: 807}
        _ = v47
    }
    for fi48 := 0; fi48 < 18; fi48++ {
        var v49 string = "a"
        var v50 uint16 = 0
        if false {
            for fi51 := 0; fi51 < 2; fi51++ {
                v52 := []int{fi51, (fi48 - (fi51 | (fi51 + (fi48 >> ((18446744073709551615) & 63)))))}
                var v53 float32 = float32((1.5 + ((0.0 - 0.1) / 1.0)))
                _ = fi51
                _ = v52
                _ = v53
            }
        } else {
            v3 = (v3 >> ((4294967295) & 63))
            v4 = ((!v4) || ((42 <= 2) && v4))
        }
        _ = fi48
        _ = v49
        _ = v50
    }
    v54 := GenS3{F0: (!v4)}
    v55 := GenS3{F0: (v4 || v4)}
    for fi56 := 0; fi56 < 10; fi56++ {
        for fi57 := 0; fi57 < 11; fi57++ {
            v58 := []int{fi57, (fi57 * fi57), int(v2), 4294967295}
            fmt.Println(fi56)
            _ = fi57
            _ = v58
        }
        _ = fi56
    }
    for fi59 := 0; fi59 < 11; fi59++ {
        for fi60 := 0; fi60 < 13; fi60++ {
            v61 := GenS4{F0: 42, F1: 147, F2: (v3 + v3)}
            _ = fi60
            _ = v61
        }
        _ = fi59
    }
    fmt.Println(v4)
    for fi62 := 0; fi62 < 10; fi62++ {
        for fi63 := 0; fi63 < 17; fi63++ {
            fmt.Println(fi62, fi62, v2)
            var v64 int32 = 0
            if (!v4) {
                var v65 uint8 = 0
                v66 := []int{}
                _ = v65
                _ = v66
            } else {
                v67 := []int{}
                var v68 string = ("a" + (("a" + "a") + "hello"))
                fmt.Println(v3, v3)
                _ = v67
                _ = v68
            }
            _ = fi63
            _ = v64
        }
        _ = fi62
    }
    v3 += (v3 | v3)
    if v4 {
        var v69 uint64 = 42
        var v70 int16 = v1
        v71 := GenS5{F0: uint8(v69)}
        _ = v69
        _ = v70
        _ = v71
    } else {
        if (1 <= 18446744073709551615) {
            switch "" {
            case "foo bar":
                for fi72 := 0; fi72 < 6; fi72++ {
                    var v73 string = "foo bar"
                    for fi74 := 0; fi74 < 18; fi74++ {
                        v73 = ("foo bar" + v73)
                        _ = fi74
                    }
                    _ = fi72
                    _ = v73
                }
            case "a":
                v75 := []int{32768, 128, 254, 32767}
                _ = v75
            default:
                v76 := []int{}
                _ = v76
            }
        }
        if (((-656 != (v1 << ((372) & 63))) || v4) || v4) {
            if v4 {
                switch 128 {
                case 0:
                    v77 := []int{426, 825}
                    _ = v77
                case -76:
                    var v78 uint16 = 255
                    _ = v78
                case 42:
                    v4 = (v4 || (!v4))
                }
                v79 := GenS6{F0: v4, F1: ("foo bar" + "\n\t")}
                v80 := GenS7{F0: v1}
                _ = v79
                _ = v80
            }
            var v81 float32 = (float32(-2.25) / 1.0)
            _ = v81
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
}
