package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32 }
func genFn0() uint8 {
    if ((!true) && (true || (false && (-440 <= 897)))) {
        var v6 string = ("" + "hello")
        var v7 uint8 = 73
        if ((0.0 / 1.0) == 0.0) {
            v8 := []int{800}
            v6 = "foo bar"
            _ = v8
        }
        _ = v6
        _ = v7
    } else {
        var v9 string = "hello"
        _ = v9
    }
    var v10 uint8 = 42
    fmt.Println(v10)
    v10 -= 134
    var v11 int16 = 32767
    _ = v10
    _ = v11
    return (v10 * (v10 * (v10 * v10)))
}
func genFn1(p0 uint16, p1 int64, p2 uint64) int64 {
    var v12 int32 = 0
    var v13 int32 = int32(p2)
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            fmt.Println(v13)
            var v16 int64 = 687
            var v17 uint32 = 2147483647
            var v18 bool = true
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        v19 := []int{fi14, 926}
        _ = fi14
        _ = v19
    }
    v20 := []int{32768, 128, -257, 0}
    var v21 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v20
    _ = v21
    return -171
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 int16 = 1
    var v4 int16 = 0
    var v5 int = 32768
    var v22 int8 = -85
    var v23 uint16 = 2
    for fi24 := 0; fi24 < 11; fi24++ {
        for fi25 := 0; fi25 < 18; fi25++ {
            delete(v2, 712)
            switch v5 {
            case 583:
                fmt.Println(fi25, v23, fi25)
            case 32767:
                v3 = v3
            default:
                switch (fi25 << ((468) & 63)) {
                case -9223372036854775808:
                    switch v5 {
                    case 9223372036854775807:
                        v26 := GenS1{F0: (v23 * v23), F1: -635, F2: 65535}
                        _ = v26
                    case 42:
                        if true {
                            v4 -= (v3 >> ((255) & 63))
                            fmt.Println(fi25, v3, v22)
                            if ((0 < 4294967295) || (!false)) {
                                fmt.Println(fi25, fi24)
                            } else {
                                switch (fi24 - fi25) {
                                case 32768:
                                    var v27 int16 = v4
                                    v28 := []int{fi24, -9223372036854775808, (fi24 & fi25)}
                                    if len(v28) > 0 {
                                        v28[0] = fi25
                                    }
                                    _ = v27
                                    _ = v28
                                case 0:
                                    v3 -= 32767
                                default:
                                    v29 := GenS2{F0: false}
                                    _ = v29
                                }
                            }
                            switch "\n\t" {
                            case "":
                                var v30 int16 = (-(v3))
                                _ = v30
                            case "hello":
                                if false {
                                    fmt.Println(fi24)
                                }
                            default:
                                if (!(255 < 128)) {
                                    v31 := GenS1{F0: (v23 * v23), F1: genFn1((v23 / ((v23) | 1)), genFn1(v23, -1, 1), 18446744073709551615), F2: 0}
                                    var v32 uint64 = 691
                                    _ = v31
                                    _ = v32
                                } else {
                                    v33 := GenS1{F0: v23, F1: -9223372036854775808, F2: 2147483647}
                                    switch -1 {
                                    case -528:
                                        _ = len(v2)
                                    case 4294967295:
                                        switch (("foo bar" + "foo bar") + ((("" + "\n\t") + ((("\n\t" + ("\n\t" + "\n\t")) + "a") + "foo bar")) + "\n\t")) {
                                        case "hello":
                                            if false {
                                                v34 := GenS1{F0: v23, F1: -217, F2: 2147483648}
                                                if true {
                                                    v35 := []int{v5, 32767}
                                                    var v36 uint = 360
                                                    _ = v35
                                                    _ = v36
                                                }
                                                _ = v34
                                            }
                                        case "foo bar":
                                            if false {
                                                v22 += v22
                                            } else {
                                                v37 := GenS3{F0: 2147483647, F1: 65535, F2: 1}
                                                fmt.Println(fi25, v3, fi24)
                                                _ = v37
                                            }
                                        default:
                                            v3 = v3
                                        }
                                    case 127:
                                        v4 += v3
                                    }
                                    v4 -= v4
                                    _ = v33
                                }
                            }
                            fmt.Println(v3, fi24)
                        }
                    default:
                        v38 := []int{(fi24 ^ fi24), fi24, (v5 * 1)}
                        if (2147483647 > 2147483647) {
                            switch 547 {
                            case -34:
                                v4 -= int16(fi25)
                            case 42:
                                v3 = v4
                            case -9223372036854775808:
                                fmt.Println(v23, fi24)
                            }
                            v38 = append(v38, fi25)
                        }
                        _ = v38
                    }
                }
            }
            _ = fi25
        }
        for fi39 := 0; fi39 < 4; fi39++ {
            _ = len(v1)
            switch ("\n\t" + "") {
            case "hello":
                v23 = v23
            }
            _ = fi39
        }
        for fi40 := 0; fi40 < 11; fi40++ {
            fmt.Println(v3)
            _ = fi40
        }
        _ = fi24
    }
    for fi41 := 0; fi41 < 7; fi41++ {
        switch 1 {
        case 1:
            v42 := GenS4{F0: -2147483648}
            _ = v42
        case 862:
            fmt.Println(v5, v5, fi41)
        case 169:
            switch (v5 | fi41) {
            case 196:
                for fi43 := 0; fi43 < 20; fi43++ {
                    switch fi41 {
                    case 2147483647:
                        if (2147483648 <= uint32(v23)) {
                            var v44 uint8 = genFn0()
                            delete(v2, 2)
                            _ = v44
                        }
                    case 906:
                        v45 := GenS4{F0: fi43}
                        _ = v45
                    }
                    v46 := GenS5{F0: 533}
                    switch (^(fi43)) {
                    case 42:
                        v47 := []int{(^(v5)), 0, (fi41 ^ fi43)}
                        _ = v47
                    }
                    v48 := []int{v5, fi41, (fi41 * fi41), (fi43 & v5)}
                    _ = fi43
                    _ = v46
                    _ = v48
                }
            case 2147483647:
                var v49 uint32 = 941
                _ = v49
            case -211:
                var v50 uint32 = 36
                _ = v50
            default:
                v4 = v3
            }
        }
        if (genFn0() >= 255) {
            v3 = 32767
            if true {
                for fi51 := 0; fi51 < 15; fi51++ {
                    fmt.Println(v22)
                    _ = fi51
                }
            }
            var v52 int16 = -956
            _ = v52
        } else {
            v1[103] = 1
            v53 := []int{v5, 9223372036854775807, fi41, fi41}
            _ = v53
        }
        v54 := GenS4{F0: 2147483647}
        _ = fi41
        _ = v54
    }
    v3 = (v3 & -292)
    if ((v23 != (v23 << ((568) & 63))) && (((0.1 / -2.25) * 0.0) == 1.5)) {
        if (0 >= 227) {
            v55 := v1[41]
            v23 -= 0
            for fi56 := 0; fi56 < 10; fi56++ {
                switch "a" {
                case "hello":
                    fmt.Println(v55)
                case "\n\t":
                    switch (fi56 | (fi56 * (^(fi56)))) {
                    case -224:
                        for fi57 := 0; fi57 < 4; fi57++ {
                            v55 += (^(v55))
                            _ = fi57
                        }
                    case 128:
                        var v58 uint8 = genFn0()
                        _ = v58
                    case -1:
                        v5 += v5
                    }
                case "a":
                    var v59 int8 = v22
                    _ = v59
                }
                v60 := []int{v5, (v5 * fi56), v5}
                _ = fi56
                _ = v60
            }
            _ = v55
        } else {
            var v61 int16 = (v3 % ((int16(v23)) | 1))
            v3 = (v3 << ((95) & 63))
            _ = v61
        }
        switch (v5 - (v5 >> ((1) & 63))) {
        case -933:
            _ = len(v2)
        case 42:
            v3 += v4
        }
        v62 := v2[4294967295]
        _ = v62
    }
    switch "" {
    case "\n\t":
        var v63 string = "hello"
        _ = v63
    case "foo bar":
        for fi64 := 0; fi64 < 20; fi64++ {
            var v65 int8 = v22
            v66 := []int{0, v5}
            var v67 bool = true
            _ = fi64
            _ = v65
            _ = v66
            _ = v67
        }
    }
    if false {
        for fi68 := 0; fi68 < 19; fi68++ {
            fmt.Println(v22, v23, v5)
            _ = fi68
        }
        var v69 int8 = v22
        _ = v69
    }
    v3 = 1000
    var v70 float64 = ((((3.14159 + 1.5) - (1e10 - (1e10 / -2.25))) * 1e10) + -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v70
}
