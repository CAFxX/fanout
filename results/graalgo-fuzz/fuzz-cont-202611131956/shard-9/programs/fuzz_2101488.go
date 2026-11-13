package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 uint }
func genFn0() int16 {
    switch ("foo bar" + "hello") {
    case "":
        var v5 string = "\n\t"
        _ = v5
    case "a":
        v6 := GenS1{F0: 42, F1: 2147483647, F2: 127}
        _ = v6
    case "hello":
        var v7 uint16 = 2
        _ = v7
    }
    v8 := []int{}
    switch "hello" {
    case "a":
        switch (("hello" + "hello") + "\n\t") {
        case "foo bar":
            if false {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                var v9 uint32 = 419
                _ = v9
            }
        }
    case "":
        _ = len(v8)
    default:
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    _ = v8
    return 0
}
func genFn1(p0 uint16, p1 uint32) uint32 {
    p0 = p0
    if false {
        switch -2147483648 {
        case 996:
            v10 := GenS1{F0: -1, F1: 1, F2: 0}
            _ = v10
        case -275:
            fmt.Println(p1)
        default:
            v11 := GenS2{F0: 5, F1: "\n\t"}
            _ = v11
        }
    }
    if false {
        var v12 bool = (!true)
        var v13 float32 = (float32(0.0) / 3.14159)
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int64 = 1
            switch (fi14 * (fi14 & 0)) {
            case 32768:
                fmt.Println(v12, v15, fi14)
            case -523:
                var v16 int = (fi14 << ((2) & 63))
                _ = v16
            case -2147483648:
                if true {
                    switch (-(fi14)) {
                    case -2147483648:
                        if (2 >= 42) {
                            if v12 {
                                var v17 int32 = 42
                                _ = v17
                            } else {
                                var v18 float32 = (float32(-2.25) / v13)
                                for fi19 := 0; fi19 < 17; fi19++ {
                                    if v12 {
                                        var v20 bool = false
                                        v21 := []int{fi14, (fi19 | fi19), int(p0), 9223372036854775807}
                                        v15 += v15
                                        _ = v20
                                        _ = v21
                                    } else {
                                        v18 += v18
                                        v22 := GenS2{F0: 255, F1: "hello"}
                                        if (!false) {
                                            fmt.Println(v12, p0, fi19)
                                            v23 := GenS2{F0: uint8(p0), F1: ""}
                                            _ = v23
                                        }
                                        _ = v22
                                    }
                                    v24 := []int{(fi19 + (fi14 / (((fi14 << ((697) & 63))) | 1)))}
                                    var v25 uint = 423
                                    _ = fi19
                                    _ = v24
                                    _ = v25
                                }
                                _ = v18
                            }
                            for fi26 := 0; fi26 < 11; fi26++ {
                                var v27 uint = 4294967295
                                if (v27 <= (v27 | 0)) {
                                    var v28 int = fi14
                                    v29 := GenS3{F0: p0, F1: v12}
                                    _ = v28
                                    _ = v29
                                } else {
                                    v27 -= (^(v27))
                                    var v30 bool = v12
                                    var v31 int = (fi26 << ((v27) & 63))
                                    var v32 int64 = v15
                                    v15 += 9223372036854775807
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                                fmt.Println(v15, p1)
                                _ = fi26
                                _ = v27
                            }
                            switch ("hello" + "a") {
                            case "a":
                                fmt.Println(p0)
                            case "":
                                v13 = v13
                            default:
                                v33 := []int{(-(fi14)), fi14}
                                _ = v33
                            }
                        }
                    default:
                        if true {
                            if v12 {
                                fmt.Println(v12, v13)
                                var v34 string = ("hello" + "")
                                _ = v34
                            }
                        }
                    }
                }
            default:
                if (0 < 69) {
                    var v35 int16 = genFn0()
                    for fi36 := 0; fi36 < 12; fi36++ {
                        var v37 int = (^(fi36))
                        v13 -= 0.0
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                } else {
                    v13 += v13
                }
            }
            _ = fi14
            _ = v15
        }
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -396
    var v3 bool = true
    var v4 int = -2147483648
    var v38 int64 = -1
    var v39 uint32 = 1
    var v40 string = "hello"
    for fi41 := 0; fi41 < 17; fi41++ {
        switch fi41 {
        case 32768:
            switch v40 {
            case "":
                var v42 int64 = v38
                _ = v42
            case "hello":
                for fi43 := 0; fi43 < 9; fi43++ {
                    var v44 uint64 = 2
                    v38 = (v38 ^ (v38 | (^(v38))))
                    v45 := GenS3{F0: 32767, F1: (!(fi43 > (fi41 + (v4 | (v4 + (fi41 - fi41))))))}
                    var v46 int8 = 127
                    _ = fi43
                    _ = v44
                    _ = v45
                    _ = v46
                }
            }
        default:
            if false {
                if (1 > 4294967295) {
                    v47 := GenS3{F0: 255, F1: v3}
                    var v48 float64 = float64(3.14159)
                    for fi49 := 0; fi49 < 4; fi49++ {
                        switch v4 {
                        case 9223372036854775807:
                            var v50 int64 = 0
                            var v51 uint16 = 1
                            _ = v50
                            _ = v51
                        case -1:
                            fmt.Println(v48, v38)
                        case -9223372036854775808:
                            v39 -= (v39 + v39)
                        }
                        v39 = v39
                        _ = fi49
                    }
                    _ = v47
                    _ = v48
                }
            }
        }
        _ = fi41
    }
    fmt.Println(v4, v4)
    for fi52 := 0; fi52 < 7; fi52++ {
        fmt.Println(v38, v40, v40)
        v3 = (true || (!((v3 && (0 > v39)) && (42 != 42))))
        _ = fi52
    }
    if (27 <= 629) {
        v38 += -361
    } else {
        switch v40 {
        case "foo bar":
            var v53 string = "\n\t"
            _ = v53
        case "":
            switch 1 {
            case 127:
                for fi54 := 0; fi54 < 15; fi54++ {
                    var v55 int64 = 1
                    if true {
                        if (32767 < 2) {
                            var v56 int32 = v2
                            _ = v56
                        }
                    }
                    v57 := GenS4{F0: genFn0(), F1: 2}
                    _ = fi54
                    _ = v55
                    _ = v57
                }
            case 32767:
                for fi58 := 0; fi58 < 14; fi58++ {
                    fmt.Println(v3, v3)
                    var v59 int = v4
                    _ = fi58
                    _ = v59
                }
            }
        }
        switch "\n\t" {
        case "hello":
            var v60 int8 = 100
            _ = v60
        case "foo bar":
            var v61 int32 = v2
            _ = v61
        }
    }
    v62 := GenS1{F0: 42, F1: -169, F2: 127}
    fmt.Println(v4, v4)
    switch ((("\n\t" + ("\n\t" + "\n\t")) + "foo bar") + v40) {
    case "a":
        fmt.Println(v4, v40, v39)
    case "":
        for fi63 := 0; fi63 < 16; fi63++ {
            if (1 < 18446744073709551615) {
                if (v3 || (!(0 > -1))) {
                    switch ((v40 + v40) + v40) {
                    case "a":
                        v64 := GenS4{F0: 0, F1: 42}
                        _ = v64
                    case "\n\t":
                        if v3 {
                            delete(v1, v39)
                        }
                    }
                    var v65 uint64 = 657
                    var v66 uint16 = 65535
                    _ = v65
                    _ = v66
                }
                v67 := GenS5{F0: 4294967295}
                _ = v67
            }
            var v68 uint32 = v39
            _ = fi63
            _ = v68
        }
    case "foo bar":
        if v3 {
            v40 = (("hello" + "a") + v40)
            fmt.Println(v39, v2)
        } else {
            v69 := []int{(v4 - (v4 * 32768)), (v4 * 127), 4294967295}
            v70 := GenS6{F0: (v39 ^ v39), F1: 600}
            _ = v69
            _ = v70
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v62
}
