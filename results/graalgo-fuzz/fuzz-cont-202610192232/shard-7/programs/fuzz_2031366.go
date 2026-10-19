package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int8; F1 int; F2 string }
type GenS6 struct { F0 uint32; F1 uint; F2 uint64 }
func genFn0() int8 {
    var v7 uint16 = 255
    switch 580 {
    case 4294967295:
        v7 = (v7 >> ((2) & 63))
    case -1:
        v8 := GenS1{F0: int16(v7), F1: (false && false), F2: 649}
        _ = v8
    default:
        fmt.Println(v7)
    }
    _ = v7
    return 0
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int32 = 1
    var v3 uint32 = 42
    var v4 float64 = ((3.14159 - 0.1) - 0.0)
    var v5 int32 = (v2 / (((v2 / (((^(v2))) | 1))) | 1))
    var v6 int = 32768
    v4 += v4
    switch "" {
    case "hello":
        var v9 uint16 = 319
        _ = v9
    case "foo bar":
        var v10 uint8 = 128
        _ = v10
    case "\n\t":
        v4 = v4
    default:
        fmt.Println(v3)
    }
    if (65535 != 446) {
        fmt.Println(v2, v1, v3)
        if (!true) {
            var v11 int8 = -80
            _ = v11
        }
    }
    if (float32((((1.5 * (1.5 * v4)) - v4) + v4)) <= 0.1) {
        fmt.Println(v6, v1)
    }
    if false {
        var v12 uint = 392
        v4 -= 0.0
        var v13 int32 = 1
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS2{F0: 42}
        fmt.Println(v6)
        var v16 bool = (v4 == -2.25)
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: 127}
    var v18 uint64 = 42
    for fi19 := 0; fi19 < 5; fi19++ {
        if true {
            v4 += (v4 / v4)
            var v20 uint32 = (^(v3))
            _ = v20
        } else {
            for fi21 := 0; fi21 < 6; fi21++ {
                switch "\n\t" {
                case "\n\t":
                    if (!true) {
                        v4 += ((((0.0 * v4) / 1.0) * 0.1) - 1e10)
                        v1 -= v1
                    }
                case "a":
                    v22 := GenS1{F0: -1, F1: true, F2: 1}
                    _ = v22
                case "hello":
                    v4 += 0.1
                }
                var v23 uint32 = 42
                if false {
                    v24 := GenS1{F0: 1000, F1: (!(-1 != 32767)), F2: v1}
                    var v25 int = -145
                    _ = v24
                    _ = v25
                } else {
                    if (!true) {
                        v26 := []int{(fi21 % (((^(fi21))) | 1)), fi21}
                        var v27 int8 = 112
                        _ = v26
                        _ = v27
                    } else {
                        v5 = 0
                        v3 = 678
                        v18 -= 2
                    }
                    switch (fi19 - (fi21 << ((879) & 63))) {
                    case 32768:
                        if ((127 <= 240) || true) {
                            v28 := []int{fi21, fi19, (fi21 & fi19), fi21}
                            switch (v6 / (((fi21 + 32767)) | 1)) {
                            case 0:
                                v5 += v5
                            case 128:
                                if true {
                                    v23 -= v23
                                }
                            }
                            _ = v28
                        }
                    case 99:
                        v29 := []int{}
                        _ = v29
                    }
                    v18 -= v18
                    var v30 int64 = (v1 % ((42) | 1))
                    _ = v30
                }
                _ = fi21
                _ = v23
            }
            var v31 uint64 = v18
            for fi32 := 0; fi32 < 1; fi32++ {
                var v33 int = (fi19 / ((fi19) | 1))
                _ = fi32
                _ = v33
            }
            _ = v31
        }
        switch "\n\t" {
        case "a":
            var v34 int64 = v1
            _ = v34
        case "hello":
            v4 += v4
        }
        for fi35 := 0; fi35 < 13; fi35++ {
            if (!(true && false)) {
                var v36 bool = false
                _ = v36
            }
            _ = fi35
        }
        _ = fi19
    }
    v2 += (v5 << ((18446744073709551615) & 63))
    v37 := GenS1{F0: 0, F1: false, F2: (v1 << ((255) & 63))}
    v38 := []int{(v6 >> ((428) & 63)), v6, (^(v6))}
    if (!false) {
        var v39 int32 = v5
        var v40 uint16 = 931
        _ = v39
        _ = v40
    }
    switch "foo bar" {
    case "a":
        if (0 == 100) {
            if (v6 > v6) {
                if (!true) {
                    v18 = 2
                    v1 = v1
                } else {
                    var v41 float64 = v4
                    _ = v41
                }
            } else {
                fmt.Println(v1, v6)
                v42 := GenS3{F0: (false || false), F1: int8(v3)}
                _ = v42
            }
            for fi43 := 0; fi43 < 8; fi43++ {
                var v44 int64 = (-(v1))
                _ = fi43
                _ = v44
            }
            var v45 uint16 = 47
            _ = v45
        } else {
            v46 := GenS4{F0: (v6 ^ v6), F1: (^(v3))}
            _ = v46
        }
    default:
        v5 = v5
    }
    v6 += int(v18)
    if len(v38) >= 2 {
        v38 = v38[1:len(v38)]
    }
    for fi47 := 0; fi47 < 16; fi47++ {
        for fi48 := 0; fi48 < 3; fi48++ {
            var v49 int32 = v5
            if true {
                var v50 int64 = (v1 / (((^(v1))) | 1))
                if false {
                    if true {
                        var v51 int64 = (-(v50))
                        _ = v51
                    } else {
                        var v52 uint = 18446744073709551615
                        var v53 uint16 = uint16(v3)
                        v54 := GenS3{F0: true, F1: -1}
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                } else {
                    v38 = append(v38, 128)
                }
                if ((v4 - v4) != ((1e10 * ((1.5 / v4) + v4)) - v4)) {
                    if (!false) {
                        v55 := GenS5{F0: 1, F1: int(v49), F2: ((("a" + "\n\t") + "a") + "")}
                        _ = v55
                    } else {
                        var v56 uint32 = (v3 ^ (^(v3)))
                        switch (fi48 | int(v18)) {
                        case 1:
                            var v57 int32 = -1
                            v58 := GenS1{F0: 1000, F1: (true && ((!(!(!false))) && (v49 != -230))), F2: (v1 / (((^(v50))) | 1))}
                            _ = v57
                            _ = v58
                        case 88:
                            var v59 int32 = (v2 >> ((1) & 63))
                            _ = v59
                        default:
                            v60 := GenS6{F0: v56, F1: 1, F2: v18}
                            if len(v38) >= 2 {
                                v38 = v38[1:len(v38)]
                            }
                            _ = v60
                        }
                        _ = v56
                    }
                    var v61 int16 = -1
                    fmt.Println(v3, v4)
                    _ = v61
                }
                _ = v50
            } else {
                if false {
                    var v62 bool = false
                    var v63 int16 = 1
                    v1 = v1
                    _ = v62
                    _ = v63
                } else {
                    var v64 string = "\n\t"
                    var v65 float32 = 3.14159
                    if len(v38) > 0 {
                        v38[0] = -2147483648
                    }
                    _ = len(v38)
                    _ = v64
                    _ = v65
                }
                fmt.Println(v2, v1, v4)
                fmt.Println(v49, v3)
            }
            v3 -= v3
            switch ("hello" + (("a" + "foo bar") + "foo bar")) {
            case "":
                _ = len(v38)
            }
            if false {
                v3 = v3
            } else {
                _ = len(v38)
                fmt.Println(v5, v49)
            }
            _ = fi48
            _ = v49
        }
        _ = fi47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
