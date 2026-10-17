package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int }

func main() {
    var v1 string = (("hello" + "") + ("" + "\n\t"))
    var v2 float32 = 0.1
    var v3 float64 = (0.0 - 3.14159)
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    var v6 int32 = 2147483647
    if false {
        var v7 uint64 = 460
        fmt.Println(v1, v7)
        if false {
            var v8 uint8 = 128
            v9 := GenS1{F0: 359}
            _ = v8
            _ = v9
        } else {
            var v10 int32 = v6
            _ = v10
        }
        _ = v7
    }
    if true {
        v2 += 1.5
        var v11 int = -2147483648
        v12 := []int{v11, v11, v11, (-(v11))}
        _ = v11
        _ = v12
    }
    v2 += (v2 - v2)
    if false {
        for fi13 := 0; fi13 < 14; fi13++ {
            switch (fi13 - fi13) {
            case -522:
                _ = len(v4)
            case 127:
                var v14 uint = 246
                _ = v14
            case -256:
                if (("foo bar" != "") && false) {
                    var v15 string = v1
                    for fi16 := 0; fi16 < 9; fi16++ {
                        var v17 uint16 = 398
                        v6 = (v6 % ((v6) | 1))
                        var v18 uint8 = 2
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    _ = v15
                }
            default:
                v6 = (v6 & v6)
            }
            _ = fi13
        }
    }
    switch v1 {
    case "foo bar":
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 uint8 = 63
            _ = fi19
            _ = v20
        }
    case "":
        v3 += v3
    }
    switch 706 {
    case 1:
        fmt.Println(v1, v6, v6)
    case 0:
        var v21 bool = false
        _ = v21
    case 32767:
        for fi22 := 0; fi22 < 11; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 uint32 = 110
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    }
    if (1.5 != (v2 * v2)) {
        var v25 int16 = 42
        _ = v25
    }
    if (255 != 255) {
        var v26 int32 = (v6 % ((v6) | 1))
        _ = v26
    } else {
        for fi27 := 0; fi27 < 19; fi27++ {
            for fi28 := 0; fi28 < 7; fi28++ {
                v29 := GenS1{F0: 907}
                switch (^(fi28)) {
                case -524:
                    if ((!(-2.25 > ((v2 - v2) * 0.0))) && true) {
                        _ = len(v5)
                    } else {
                        delete(v5, (fi28 << ((337) & 63)))
                        v2 -= v2
                        v3 = ((v3 + (v3 - (v3 - 1e10))) - v3)
                    }
                }
                fmt.Println(fi28, fi27)
                _ = fi28
                _ = v29
            }
            v30 := GenS1{F0: 0}
            v31 := GenS2{F0: 1}
            _ = fi27
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 7; fi32++ {
            v6 -= v6
            switch fi32 {
            case 9223372036854775807:
                switch fi32 {
                case 32767:
                    fmt.Println(v1, v1, v2)
                case 1:
                    v3 = v3
                }
            case 127:
                if true {
                    var v33 uint32 = 2147483648
                    _ = v33
                } else {
                    switch fi32 {
                    case 42:
                        fmt.Println(v6, v3)
                    default:
                        var v34 int32 = v6
                        _ = v34
                    }
                }
            case -55:
                v35 := GenS3{F0: 42}
                _ = v35
            }
            _ = fi32
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            fmt.Println(v2)
            var v37 int32 = (v6 - 0)
            switch v1 {
            case "":
                var v38 uint8 = 0
                _ = v38
            case "hello":
                if true {
                    if (!(42 != 9)) {
                        switch (fi36 >> ((427) & 63)) {
                        case 128:
                            var v39 int16 = 1000
                            _ = v39
                        default:
                            switch (fi36 & (fi36 % ((128) | 1))) {
                            case 806:
                                v3 += v3
                            case 0:
                                fmt.Println(v2, v37)
                            case 488:
                                for fi40 := 0; fi40 < 2; fi40++ {
                                    v41 := []int{fi40}
                                    switch (^(fi40)) {
                                    case 42:
                                        fmt.Println(v37, v3, fi40)
                                    case -284:
                                        if (!(0 < 18446744073709551615)) {
                                            var v42 float64 = v3
                                            _ = v42
                                        }
                                    case 4294967295:
                                        if len(v41) > 0 {
                                            v41[0] = (fi36 / (((fi40 % ((2147483647) | 1))) | 1))
                                        }
                                    }
                                    _ = fi40
                                    _ = v41
                                }
                            }
                        }
                        fmt.Println(v3)
                    }
                    fmt.Println(fi36, v37)
                }
            case "a":
                var v43 float64 = float64(float32(v3))
                _ = v43
            default:
                delete(v4, fi36)
            }
            _ = fi36
            _ = v37
        }
    }
    if ((!false) && true) {
        var v44 bool = false
        _ = v44
    } else {
        switch (v1 + v1) {
        case "hello":
            v45 := GenS1{F0: 0}
            _ = v45
        }
        v46 := GenS3{F0: 42}
        v47 := GenS4{F0: -1}
        _ = v46
        _ = v47
    }
    var v48 int8 = 100
    switch v1 {
    case "\n\t":
        if true {
            var v49 uint8 = 255
            var v50 uint64 = 0
            var v51 uint64 = v50
            _ = v49
            _ = v50
            _ = v51
        } else {
            if (!false) {
                v3 = (-2.25 * v3)
                fmt.Println(v1, v48)
                fmt.Println(v3)
            }
            v48 -= (v48 & -128)
            for fi52 := 0; fi52 < 18; fi52++ {
                if false {
                    for fi53 := 0; fi53 < 12; fi53++ {
                        if true {
                            v2 += 3.14159
                            v54 := GenS3{F0: (v6 | (v6 >> ((4294967295) & 63)))}
                            if (100 >= (v48 & v48)) {
                                var v55 int32 = v6
                                switch 100 {
                                case 42:
                                    var v56 uint = 0
                                    if true {
                                        v6 = (v6 & (v6 - int32(fi52)))
                                    }
                                    _ = v56
                                case 9223372036854775807:
                                    var v57 uint = 0
                                    v58 := GenS3{F0: -2147483648}
                                    v1 = v1
                                    v48 = (v48 + v48)
                                    var v59 bool = (1.5 != float64(v2))
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                case 0:
                                    var v60 string = v1
                                    _ = v60
                                }
                                _ = v55
                            }
                            _ = v54
                        } else {
                            var v61 uint8 = 0
                            _ = v61
                        }
                        _ = fi53
                    }
                    v62 := []int{}
                    if (!true) {
                        fmt.Println(v1, v2, v1)
                        fmt.Println(v2, v3)
                        switch -2147483648 {
                        case -825:
                            if (825 > 283) {
                                _ = len(v4)
                                for fi63 := 0; fi63 < 7; fi63++ {
                                    v64 := GenS1{F0: 4294967295}
                                    var v65 uint = 42
                                    v66 := GenS4{F0: -237}
                                    _ = fi63
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                }
                                var v67 int16 = 1
                                _ = v67
                            }
                        case -9223372036854775808:
                            v68 := v4[(fi52 - fi52)]
                            _ = v68
                        case 0:
                            var v69 int16 = -32768
                            _ = v69
                        }
                    }
                    _ = v62
                }
                var v70 uint = 4294967295
                for fi71 := 0; fi71 < 17; fi71++ {
                    v5[479] = true
                    switch (fi71 / ((9223372036854775807) | 1)) {
                    case -31:
                        v2 = 0.1
                    }
                    _ = fi71
                }
                _ = fi52
                _ = v70
            }
        }
    case "foo bar":
        v48 = (v48 % (((v48 << ((785) & 63))) | 1))
    }
    var v72 bool = (-1 >= (v6 / ((2147483647) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v72
}
