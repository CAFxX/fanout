package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int; F2 int32 }
type GenS4 struct { F0 uint; F1 int64 }
func genFn0(p0 int, p1 string) uint8 {
    var v4 uint32 = 65535
    if (!false) {
        v5 := []int{2147483647}
        switch ("a" + p1) {
        case "foo bar":
            var v6 uint8 = 199
            _ = v6
        default:
            var v7 int16 = -32768
            _ = v7
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint16 = 42
            _ = fi8
            _ = v9
        }
        _ = v5
    } else {
        if (true && false) {
            v10 := []int{p0, (p0 | 42), 2147483647, (p0 / ((654) | 1))}
            p0 = (p0 | (p0 >> ((42) & 63)))
            switch (p0 << ((2) & 63)) {
            case 88:
                var v11 int8 = -73
                _ = v11
            case 9223372036854775807:
                if true {
                    for fi12 := 0; fi12 < 17; fi12++ {
                        p0 += fi12
                        for fi13 := 0; fi13 < 6; fi13++ {
                            v10 = append(v10, -1)
                            _ = fi13
                        }
                        _ = fi12
                    }
                    v4 = v4
                } else {
                    var v14 string = "a"
                    _ = v14
                }
            case 2147483647:
                var v15 bool = false
                _ = v15
            }
            _ = v10
        } else {
            for fi16 := 0; fi16 < 13; fi16++ {
                switch fi16 {
                case 32768:
                    var v17 uint64 = 2
                    _ = v17
                case 0:
                    v18 := []int{fi16}
                    _ = v18
                case -622:
                    if true {
                        p1 = p1
                    }
                default:
                    p0 += (fi16 | 905)
                }
                _ = fi16
            }
        }
        v19 := GenS1{F0: 0}
        _ = v19
    }
    if false {
        var v20 int = (p0 + -519)
        var v21 uint64 = 42
        _ = v20
        _ = v21
    } else {
        p1 = (p1 + (p1 + "a"))
        switch p1 {
        case "\n\t":
            for fi22 := 0; fi22 < 14; fi22++ {
                p0 = fi22
                fmt.Println(p0)
                v23 := GenS2{F0: 65535}
                _ = fi22
                _ = v23
            }
        case "":
            switch (-(p0)) {
            case 898:
                switch ("foo bar" + ("\n\t" + p1)) {
                case "a":
                    if (float32(1e10) == -2.25) {
                        switch p0 {
                        case -2147483648:
                            switch (p0 * p0) {
                            case 32767:
                                var v24 float32 = 0.1
                                _ = v24
                            case -1:
                                var v25 float64 = ((((1.5 * float64(0.1)) - 0.1) - ((float64(-2.25) - ((3.14159 - float64(-2.25)) + (3.14159 - -2.25))) - -2.25)) + 3.14159)
                                _ = v25
                            default:
                                v4 -= v4
                            }
                        case -1:
                            v4 -= v4
                        case 2147483647:
                            for fi26 := 0; fi26 < 3; fi26++ {
                                var v27 uint8 = 116
                                if false {
                                    var v28 uint32 = v4
                                    switch "a" {
                                    case "foo bar":
                                        v29 := []int{(-(fi26))}
                                        _ = v29
                                    case "hello":
                                        v28 = v4
                                    case "a":
                                        v30 := GenS1{F0: 4294967295}
                                        _ = v30
                                    default:
                                        for fi31 := 0; fi31 < 14; fi31++ {
                                            if (!true) {
                                                var v32 float32 = -2.25
                                                _ = v32
                                            }
                                            var v33 int64 = -9223372036854775808
                                            v28 = 0
                                            _ = fi31
                                            _ = v33
                                        }
                                    }
                                    var v34 int8 = -31
                                    _ = v28
                                    _ = v34
                                } else {
                                    fmt.Println(p1)
                                    var v35 uint8 = 2
                                    _ = v35
                                }
                                _ = fi26
                                _ = v27
                            }
                        }
                        v4 = 65535
                        switch p1 {
                        case "":
                            if (-120 >= -1) {
                                var v36 uint16 = 42
                                _ = v36
                            } else {
                                for fi37 := 0; fi37 < 1; fi37++ {
                                    v38 := []int{fi37, 2147483647, fi37, (-(fi37))}
                                    _ = fi37
                                    _ = v38
                                }
                                fmt.Println(p0, p0)
                            }
                        }
                    } else {
                        switch ((p1 + "hello") + p1) {
                        case "\n\t":
                            v39 := GenS2{F0: (v4 * v4)}
                            _ = v39
                        default:
                            var v40 int64 = 717
                            _ = v40
                        }
                        v41 := []int{p0, (p0 + (^(p0))), (p0 - 595)}
                        _ = v41
                    }
                }
            case -338:
                if (true && (127 <= 117)) {
                    var v42 uint8 = 138
                    _ = v42
                }
            }
        default:
            switch 128 {
            case -516:
                for fi43 := 0; fi43 < 13; fi43++ {
                    var v44 float64 = 0.1
                    var v45 uint16 = 0
                    _ = fi43
                    _ = v44
                    _ = v45
                }
            default:
                v46 := []int{-550, -192, -208}
                _ = v46
            }
        }
        if false {
            v47 := []int{}
            _ = v47
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return uint8(v4)
}

func main() {
    var v1 int16 = -32768
    v2 := []uint32{4294967295, 42, 170}
    var v3 string = "a"
    var v48 bool = (1.5 >= float64(-2.25))
    if ((!true) && (!v48)) {
        if true {
            v3 = v3
            var v49 bool = (!(v48 && (true || (2 <= 4294967295))))
            _ = v49
        }
        for fi50 := 0; fi50 < 11; fi50++ {
            fmt.Println(v48, v1, v48)
            var v51 uint8 = 1
            _ = fi50
            _ = v51
        }
    }
    for fi52 := 0; fi52 < 9; fi52++ {
        switch 1 {
        case 32767:
            for fi53 := 0; fi53 < 4; fi53++ {
                v1 -= v1
                v48 = ((v48 && v48) || (-2147483648 > 1))
                if false {
                    if (1 == 4294967295) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    fmt.Println(v1)
                    v1 = v1
                }
                _ = fi53
            }
        }
        for fi54 := 0; fi54 < 2; fi54++ {
            v48 = false
            _ = fi54
        }
        var v55 uint8 = 184
        _ = fi52
        _ = v55
    }
    switch (v3 + (v3 + (("" + ("a" + "")) + v3))) {
    case "hello":
        var v56 bool = v48
        _ = v56
    default:
        v2 = append(v2, 1)
    }
    v3 = v3
    for fi57 := 0; fi57 < 4; fi57++ {
        v48 = v48
        var v58 int = fi57
        _ = fi57
        _ = v58
    }
    v3 = ((("a" + (("\n\t" + "foo bar") + v3)) + "foo bar") + (v3 + v3))
    var v59 int32 = -1
    switch -9223372036854775808 {
    case 885:
        v60 := GenS3{F0: v3, F1: 4294967295, F2: v59}
        _ = v60
    case 2147483647:
        switch 489 {
        case 264:
            v1 = (v1 | (v1 % ((v1) | 1)))
        }
    case 927:
        var v61 int32 = v59
        _ = v61
    }
    v59 = v59
    var v62 bool = (v48 || (71 >= (v1 | (v1 << ((255) & 63)))))
    if v48 {
        if (!v62) {
            var v63 uint = 255
            _ = v63
        }
        fmt.Println(v59)
    } else {
        var v64 float32 = 1.5
        _ = v64
    }
    var v65 int8 = -1
    for fi66 := 0; fi66 < 15; fi66++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi66
    }
    for fi67 := 0; fi67 < 15; fi67++ {
        fmt.Println(v59, v59, v1)
        for fi68 := 0; fi68 < 20; fi68++ {
            _ = len(v2)
            var v69 bool = (!(true || v62))
            v70 := GenS4{F0: 2, F1: 42}
            var v71 int16 = 32767
            fmt.Println(v69, v1, v3)
            _ = fi68
            _ = v69
            _ = v70
            _ = v71
        }
        _ = fi67
    }
    var v72 uint64 = 42
    v59 += 0
    v2 = append(v2, 958)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("v59:", v59)
    fmt.Println("v62:", v62)
    fmt.Println("v65:", v65)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v59
    _ = v62
    _ = v65
    _ = v72
}
