package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8; F1 uint64 }
func genFn0(p0 string, p1 uint8) uint8 {
    p0 = "foo bar"
    v5 := GenS1{F0: 42, F1: (p1 * (p1 - 127))}
    var v6 string = p0
    p1 = (p1 << ((2) & 63))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int16 = -32768
    v2 := []uint32{}
    var v3 bool = ((51 == 1) || ((0.0 != ((1e10 + float64(1.5)) + ((1.5 / (1e10 - 0.1)) - 1.5))) && (0.1 > -2.25)))
    v4 := make(map[int]uint32)
    switch 783 {
    case 416:
        var v7 bool = v3
        _ = v7
    default:
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            fmt.Println(v1, v1)
        case "a":
            v8 := GenS2{F0: 9223372036854775807}
            _ = v8
        }
    }
    if v3 {
        var v9 uint = 117
        fmt.Println(v9, v1)
        v2 = append(v2, 2147483647)
        _ = v9
    }
    v10 := v4[42]
    fmt.Println(v10, v3, v1)
    v11 := GenS1{F0: 42, F1: genFn0("hello", genFn0("", 116))}
    v12 := GenS1{F0: 496, F1: 127}
    v13 := GenS1{F0: 919, F1: 179}
    for fi14 := 0; fi14 < 9; fi14++ {
        v3 = v3
        var v15 uint16 = 1
        _ = fi14
        _ = v15
    }
    fmt.Println(v10, v10)
    var v16 int8 = 1
    switch 32768 {
    case 32767:
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 uint16 = 2
            switch "" {
            case "":
                v19 := GenS2{F0: (fi17 >> ((774) & 63))}
                _ = v19
            }
            _ = fi17
            _ = v18
        }
    case 166:
        for fi20 := 0; fi20 < 15; fi20++ {
            if v3 {
                switch "\n\t" {
                case "foo bar":
                    var v21 string = (("" + "hello") + "\n\t")
                    _ = v21
                }
                fmt.Println(fi20)
                if (1 >= 0) {
                    for fi22 := 0; fi22 < 9; fi22++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v2 = append(v2, v10)
                        fmt.Println(v16, v16)
                        _ = fi22
                    }
                    var v23 int8 = (-(v16))
                    _ = v23
                } else {
                    fmt.Println(v10, v1, v10)
                }
            } else {
                for fi24 := 0; fi24 < 10; fi24++ {
                    switch -9223372036854775808 {
                    case -671:
                        var v25 uint64 = 718
                        _ = v25
                    case 1:
                        v26 := GenS3{F0: 728}
                        _ = v26
                    default:
                        var v27 uint16 = 42
                        _ = v27
                    }
                    _ = fi24
                }
                v10 = v10
            }
            for fi28 := 0; fi28 < 12; fi28++ {
                v2 = append(v2, 1)
                _ = fi28
            }
            _ = fi20
        }
    case -2147483648:
        v29 := GenS4{F0: 2, F1: (v10 * 2)}
        _ = v29
    default:
        if (687 >= -888) {
            var v30 int16 = v1
            _ = v30
        }
    }
    if (true || (!v3)) {
        var v31 int8 = v16
        var v32 int16 = 1000
        if v3 {
            v32 += (v32 << ((0) & 63))
        }
        _ = v31
        _ = v32
    } else {
        switch "a" {
        case "hello":
            var v33 bool = false
            _ = v33
        case "foo bar":
            v16 += (v16 & 8)
        case "":
            if (!false) {
                for fi34 := 0; fi34 < 5; fi34++ {
                    switch "hello" {
                    case "\n\t":
                        v1 = v1
                    case "a":
                        v35 := GenS3{F0: 4294967295}
                        _ = v35
                    default:
                        if len(v2) > 0 {
                            v2[0] = v10
                        }
                    }
                    switch 0 {
                    case 42:
                        if (v3 && v3) {
                            var v36 bool = (int64(v10) != -1)
                            var v37 uint32 = v10
                            _ = v36
                            _ = v37
                        }
                    case -933:
                        v38 := GenS5{F0: 2, F1: v10}
                        _ = v38
                    case 4294967295:
                        v10 -= 853
                    default:
                        v2 = append(v2, (v10 * 2147483647))
                    }
                    _ = fi34
                }
                var v39 int16 = v1
                if (0.0 <= (((((1e10 + 1e10) - 0.1) - 1e10) * (-2.25 / 1e10)) * 1.5)) {
                    if (!(v3 || v3)) {
                        var v40 bool = (1 != 717)
                        _ = v40
                    }
                    switch (("" + "\n\t") + "\n\t") {
                    case "\n\t":
                        for fi41 := 0; fi41 < 15; fi41++ {
                            v2 = append(v2, 295)
                            _ = fi41
                        }
                    }
                    if ((((-2.25 * 1.5) - 1.5) == 1.5) && (!(209 > 114))) {
                        v3 = ("" >= "")
                        v39 = v39
                    } else {
                        if v3 {
                            v10 = v10
                            v10 = v10
                        }
                        for fi42 := 0; fi42 < 18; fi42++ {
                            v43 := GenS6{F0: (fi42 ^ (fi42 % (((fi42 / ((fi42) | 1))) | 1)))}
                            if (127 == 135) {
                                v1 = (v39 * (-(v39)))
                                v10 = (v10 & 321)
                                v2 = append(v2, (v10 / ((v10) | 1)))
                            }
                            _ = fi42
                            _ = v43
                        }
                    }
                } else {
                    fmt.Println(v10, v10)
                    v16 -= v16
                }
                _ = v39
            } else {
                var v44 int16 = (v1 >> ((18446744073709551615) & 63))
                fmt.Println(v1)
                v45 := GenS7{F0: (v16 & (v16 | 103)), F1: 42}
                _ = v44
                _ = v45
            }
        }
    }
    var v46 uint64 = 832
    v2 = append(v2, (v10 >> ((1) & 63)))
    for fi47 := 0; fi47 < 3; fi47++ {
        for fi48 := 0; fi48 < 7; fi48++ {
            v2 = append(v2, v10)
            var v49 int16 = (v1 >> ((0) & 63))
            switch (fi48 * fi48) {
            case 9223372036854775807:
                v50 := v4[0]
                _ = v50
            default:
                var v51 uint64 = v46
                _ = v51
            }
            _ = fi48
            _ = v49
        }
        _ = fi47
    }
    var v52 int8 = (v16 >> ((65535) & 63))
    fmt.Println(v3, v52, v52)
    v53 := v4[2147483647]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v52
    _ = v53
}
