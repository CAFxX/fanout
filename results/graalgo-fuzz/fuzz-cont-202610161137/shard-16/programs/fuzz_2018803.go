package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int32; F1 int8; F2 int64 }
type GenS3 struct { F0 bool; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 bool; F1 int; F2 uint64 }
type GenS6 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS7 struct { F0 int16; F1 int64; F2 int8 }
type GenS8 struct { F0 uint8; F1 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32 }
func genFn0(p0 int16) int {
    var v6 int64 = -1
    switch 2147483647 {
    case 0:
        for fi7 := 0; fi7 < 18; fi7++ {
            p0 = p0
            v8 := GenS1{F0: 165, F1: v6}
            _ = fi7
            _ = v8
        }
    case 616:
        var v9 uint32 = 195
        _ = v9
    }
    _ = p0
    _ = v6
    return 4294967295
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 1e10
    v3 := []string{"foo bar"}
    var v4 int32 = 276
    var v5 uint32 = 239
    v10 := GenS1{F0: 255, F1: 9223372036854775807}
    v11 := GenS2{F0: -2147483648, F1: 1, F2: -996}
    switch 1 {
    case 9223372036854775807:
        switch genFn0(42) {
        case 2147483647:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            v2 -= 0.1
        }
    case 42:
        v5 = 2147483647
    case -952:
        switch "foo bar" {
        case "\n\t":
            var v12 uint64 = 42
            _ = v12
        case "a":
            var v13 int16 = 1000
            _ = v13
        case "hello":
            var v14 uint = 42
            _ = v14
        default:
            var v15 uint = 42
            _ = v15
        }
    }
    var v16 uint16 = 696
    var v17 string = "a"
    switch -2147483648 {
    case -1:
        if ((v16 >> ((813) & 63)) != v16) {
            var v18 int8 = 66
            var v19 int64 = 1
            fmt.Println(v4)
            _ = v18
            _ = v19
        } else {
            v20 := GenS1{F0: (v16 / ((v16) | 1)), F1: -9223372036854775808}
            if (true && (v4 < v4)) {
                for fi21 := 0; fi21 < 7; fi21++ {
                    v22 := GenS3{F0: (!(32767 < v16)), F1: -2147483648, F2: v16}
                    var v23 uint32 = uint32(v16)
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                var v24 int32 = v4
                var v25 uint32 = v5
                _ = v24
                _ = v25
            } else {
                v16 = 644
                var v26 string = (v17 + v17)
                var v27 int16 = -1
                _ = v26
                _ = v27
            }
            if false {
                for fi28 := 0; fi28 < 8; fi28++ {
                    for fi29 := 0; fi29 < 3; fi29++ {
                        v30 := GenS1{F0: v16, F1: -1}
                        fmt.Println(v17, v2, fi28)
                        if (float64(1e10) <= v2) {
                            v31 := GenS2{F0: v4, F1: 42, F2: -1}
                            v32 := GenS2{F0: v4, F1: 0, F2: -1}
                            v17 = (v17 + "\n\t")
                            _ = len(v3)
                            _ = v31
                            _ = v32
                        } else {
                            v1[(v5 % ((v5) | 1))] = (v17 + (((v17 + v17) + ((("" + (((v17 + v17) + "\n\t") + v17)) + v17) + ((v17 + "hello") + v17))) + (v17 + (v17 + (v17 + v17)))))
                            switch (fi29 >> ((189) & 63)) {
                            case 2147483647:
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v4)
                                case "\n\t":
                                    v16 -= v16
                                default:
                                    if false {
                                        v2 = 0.1
                                    }
                                }
                            }
                            var v33 int = (fi28 << ((2) & 63))
                            _ = v33
                        }
                        _ = fi29
                        _ = v30
                    }
                    v34 := GenS1{F0: 42, F1: 55}
                    _ = fi28
                    _ = v34
                }
                switch v17 {
                case "hello":
                    switch 366 {
                    case 32767:
                        if (-345 > 42) {
                            switch genFn0(1000) {
                            case -911:
                                for fi35 := 0; fi35 < 12; fi35++ {
                                    switch v17 {
                                    case "a":
                                        var v36 uint16 = 412
                                        _ = v36
                                    default:
                                        v4 += (v4 ^ v4)
                                    }
                                    _ = fi35
                                }
                            case 4294967295:
                                fmt.Println(v5)
                            case 42:
                                v3 = append(v3, v17)
                            }
                            _ = len(v1)
                        }
                    case 128:
                        _ = len(v3)
                    }
                case "foo bar":
                    v5 -= v5
                }
                v37 := GenS3{F0: false, F1: v4, F2: 255}
                _ = v37
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v16, v16, v4)
            }
            _ = v20
        }
    }
    var v38 string = "hello"
    v17 = v17
    switch 32768 {
    case 42:
        var v39 uint16 = v16
        _ = v39
    case -265:
        switch v17 {
        case "foo bar":
            v38 = v17
        case "\n\t":
            var v40 int16 = 0
            _ = v40
        default:
            switch 2147483647 {
            case 2147483647:
                v5 += v5
            case 128:
                v41 := GenS2{F0: (v4 | (-(v4))), F1: 127, F2: 515}
                _ = v41
            case -9223372036854775808:
                fmt.Println(v16, v4, v4)
            default:
                var v42 float32 = float32(float64(float32(v2)))
                _ = v42
            }
        }
    case 32768:
        _ = len(v3)
    default:
        switch -1 {
        case -921:
            for fi43 := 0; fi43 < 18; fi43++ {
                v16 -= (v16 << ((42) & 63))
                switch "" {
                case "hello":
                    v44 := GenS4{F0: 42, F1: v4}
                    _ = v44
                }
                _ = fi43
            }
        default:
            v45 := GenS3{F0: (v5 != (v5 / ((2147483648) | 1))), F1: v4, F2: v16}
            _ = v45
        }
    }
    if (4294967295 >= uint(v16)) {
        delete(v1, (^(v5)))
    }
    if (!false) {
        for fi46 := 0; fi46 < 19; fi46++ {
            for fi47 := 0; fi47 < 17; fi47++ {
                v3 = append(v3, ("a" + v17))
                _ = fi47
            }
            switch v38 {
            case "":
                var v48 uint32 = (v5 << ((312) & 63))
                _ = v48
            case "hello":
                if true {
                    if true {
                        if false {
                            var v49 int16 = -32768
                            var v50 string = v38
                            switch int(v16) {
                            case 1:
                                v49 = (v49 + v49)
                            case 128:
                                v51 := GenS2{F0: (v4 + -2147483648), F1: 1, F2: 42}
                                _ = v51
                            case -2147483648:
                                for fi52 := 0; fi52 < 2; fi52++ {
                                    var v53 uint64 = 1
                                    _ = fi52
                                    _ = v53
                                }
                            default:
                                if ((32767 != (fi46 >> ((18446744073709551615) & 63))) && false) {
                                    fmt.Println(v38, v49, v5)
                                    var v54 int8 = -1
                                    _ = len(v1)
                                    _ = v54
                                }
                            }
                            _ = v49
                            _ = v50
                        }
                        v55 := GenS5{F0: (42 <= 680), F1: int(v4), F2: 4294967295}
                        fmt.Println(v4, v4)
                        _ = v55
                    } else {
                        var v56 int = genFn0(0)
                        v57 := GenS6{F0: -128, F1: 2, F2: 752}
                        _ = v56
                        _ = v57
                    }
                    for fi58 := 0; fi58 < 3; fi58++ {
                        if ((3.14159 / 0.1) != ((float32(((((((1.5 - (-2.25 - v2)) - float64(1.5)) / (v2 / v2)) * v2) * float64(0.0)) - v2)) * float32((0.0 + float64(((1.5 - 3.14159) * -2.25))))) / 1e10)) {
                            v2 -= v2
                        } else {
                            v4 = v4
                            v59 := GenS1{F0: (v16 >> ((255) & 63)), F1: 0}
                            var v60 string = v17
                            _ = v59
                            _ = v60
                        }
                        if (true || true) {
                            if ((v2 + v2) <= v2) {
                                v1[(v5 >> ((1) & 63))] = v17
                                fmt.Println(fi46)
                            } else {
                                v61 := GenS7{F0: 42, F1: 0, F2: -1}
                                if (!(42 < -1)) {
                                    v62 := GenS5{F0: false, F1: (^(fi46)), F2: 4294967295}
                                    _ = v62
                                }
                                if true {
                                    var v63 int8 = int8(v16)
                                    _ = v63
                                } else {
                                    v4 += v4
                                    v5 -= (v5 << ((390) & 63))
                                    v64 := GenS8{F0: 255, F1: v16}
                                    _ = v64
                                }
                                _ = v61
                            }
                        } else {
                            v5 = 1
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v65 uint16 = (^(v16))
                            _ = v65
                        }
                        if false {
                            v5 -= (v5 ^ (v5 | v5))
                            var v66 bool = true
                            if false {
                                _ = len(v3)
                                if ((42 <= 42) && ((^(v5)) >= 106)) {
                                    var v67 int = genFn0(-452)
                                    switch (v67 >> ((42) & 63)) {
                                    case 9223372036854775807:
                                        fmt.Println(v38)
                                    case -1:
                                        fmt.Println(v67)
                                    }
                                    v68 := GenS9{F0: 32767}
                                    if true {
                                        if (127 > uint8(v5)) {
                                            v67 += 0
                                            switch (^(fi58)) {
                                            case 0:
                                                var v69 uint16 = (^(v16))
                                                _ = v69
                                            case 4294967295:
                                                var v70 string = (v17 + ((v17 + (v17 + (v17 + v17))) + "a"))
                                                _ = v70
                                            default:
                                                var v71 uint8 = 255
                                                v38 = v38
                                                var v72 uint = 2
                                                v5 = 2
                                                _ = v71
                                                _ = v72
                                            }
                                        }
                                        if v66 {
                                            var v73 int16 = 1
                                            v4 = v4
                                            _ = v73
                                        }
                                    }
                                    _ = v67
                                    _ = v68
                                } else {
                                    v66 = true
                                }
                                switch ("\n\t" + v17) {
                                case "":
                                    fmt.Println(fi46, v16, v17)
                                case "a":
                                    v2 -= float64(3.14159)
                                }
                            } else {
                                v5 = (v5 ^ (v5 / ((65535) | 1)))
                                var v74 string = "a"
                                var v75 int64 = 0
                                _ = v74
                                _ = v75
                            }
                            _ = v66
                        } else {
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                            v38 = ("" + "hello")
                            v76 := GenS10{F0: v5}
                            _ = v76
                        }
                        if len(v3) > 0 {
                            v3[0] = v17
                        }
                        _ = fi58
                    }
                }
            }
            var v77 uint16 = v16
            _ = fi46
            _ = v77
        }
        var v78 int = 128
        var v79 float32 = float32((0.1 - (1e10 + 1e10)))
        _ = v78
        _ = v79
    } else {
        _ = len(v1)
        var v80 uint32 = v5
        fmt.Println(v38)
        _ = v80
    }
    for fi81 := 0; fi81 < 16; fi81++ {
        fmt.Println(v16, v38, v16)
        _ = fi81
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
