package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 int32; F1 int16; F2 string }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 uint8; F2 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16 }
func genFn0(p0 int16) uint16 {
    v5 := []int{}
    if (p0 >= 1000) {
        var v6 bool = true
        _ = v6
    } else {
        v7 := GenS1{F0: 0, F1: 32767, F2: -1}
        var v8 uint64 = 611
        _ = v7
        _ = v8
    }
    var v9 uint32 = 42
    for fi10 := 0; fi10 < 15; fi10++ {
        if (-1 >= 127) {
            _ = len(v5)
            v11 := GenS1{F0: 127, F1: 65535, F2: -1}
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 int64 = -9223372036854775808
                if len(v5) > 0 {
                    v5[0] = (fi10 | fi12)
                }
                p0 -= 719
                fmt.Println(v13, fi12, v9)
                _ = fi12
                _ = v13
            }
            _ = v11
        } else {
            var v14 int16 = (p0 * (-(p0)))
            _ = v14
        }
        _ = fi10
    }
    _ = p0
    _ = v5
    _ = v9
    return 368
}

func main() {
    var v1 string = "foo bar"
    var v2 bool = false
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 1
    for fi15 := 0; fi15 < 2; fi15++ {
        v2 = true
        fmt.Println(v2, v2, v1)
        _ = fi15
    }
    switch 9223372036854775807 {
    case 127:
        var v16 uint32 = 0
        _ = v16
    case -628:
        v2 = ((v2 || v2) && (!v2))
    case -2147483648:
        switch 683 {
        case -2147483648:
            fmt.Println(v2, v2, v2)
        case 127:
            if v2 {
                v4 = v4
                v17 := GenS2{F0: 42, F1: 621, F2: "foo bar"}
                var v18 string = v1
                _ = v17
                _ = v18
            } else {
                var v19 bool = v2
                v20 := GenS3{F0: -698, F1: v4}
                var v21 int16 = 42
                _ = v19
                _ = v20
                _ = v21
            }
        case 2147483647:
            v2 = v2
        }
    }
    if (v2 || v2) {
        v22 := GenS4{F0: 1}
        _ = v22
    }
    switch "foo bar" {
    case "a":
        for fi23 := 0; fi23 < 14; fi23++ {
            v24 := GenS3{F0: 9223372036854775807, F1: v4}
            v4 = (^(v4))
            if (-2.25 <= (0.0 - (1e10 * ((0.0 + 1e10) + ((1e10 - -2.25) + 1.5))))) {
                if v2 {
                    switch (fi23 + fi23) {
                    case 1:
                        v25 := []int{fi23, (fi23 - (fi23 | (fi23 - (^(fi23)))))}
                        _ = v25
                    default:
                        switch (v1 + "\n\t") {
                        case "":
                            switch ("a" + v1) {
                            case "hello":
                                var v26 uint32 = 0
                                _ = v26
                            default:
                                v27 := v3[42]
                                _ = v27
                            }
                        default:
                            v28 := []int{}
                            _ = v28
                        }
                    }
                    var v29 uint8 = 127
                    _ = v29
                }
                for fi30 := 0; fi30 < 10; fi30++ {
                    switch -2147483648 {
                    case -631:
                        if (127 >= 100) {
                            fmt.Println(v2, v2)
                            if true {
                                var v31 float32 = 1e10
                                var v32 string = v1
                                _ = v31
                                _ = v32
                            } else {
                                switch 0 {
                                case 245:
                                    v33 := []int{(fi23 + (fi30 - (fi30 - fi30)))}
                                    _ = v33
                                case -442:
                                    v34 := GenS3{F0: 42, F1: v4}
                                    _ = v34
                                case 32768:
                                    v2 = (!v2)
                                default:
                                    fmt.Println(fi30, fi23, v1)
                                }
                                fmt.Println(v2)
                                v35 := []int{fi30}
                                _ = v35
                            }
                            v3[42] = 2147483647
                            v4 += (v4 ^ v4)
                            v4 -= (v4 >> ((0) & 63))
                        }
                    case -2147483648:
                        v36 := []int{}
                        _ = v36
                    default:
                        fmt.Println(fi23)
                    }
                    if v2 {
                        var v37 uint16 = 42
                        v37 = v37
                        v4 = 42
                        _ = v37
                    }
                    var v38 bool = (803 == 0)
                    _ = fi30
                    _ = v38
                }
            }
            _ = fi23
            _ = v24
        }
    case "\n\t":
        fmt.Println(v2)
    case "":
        v39 := GenS1{F0: 100, F1: 0, F2: 0}
        _ = v39
    default:
        if v2 {
            for fi40 := 0; fi40 < 20; fi40++ {
                v3[65535] = 1
                _ = fi40
            }
            switch -2147483648 {
            case 589:
                fmt.Println(v2)
            case -752:
                if (v2 || v2) {
                    if false {
                        for fi41 := 0; fi41 < 9; fi41++ {
                            fmt.Println(v4)
                            _ = fi41
                        }
                    } else {
                        switch -309 {
                        case 1:
                            var v42 float64 = 0.0
                            _ = v42
                        case -545:
                            v43 := GenS3{F0: -1, F1: 42}
                            _ = v43
                        default:
                            switch 32768 {
                            case 127:
                                if v2 {
                                    switch 1 {
                                    case -658:
                                        if v2 {
                                            v44 := GenS5{F0: v1, F1: (^(v4)), F2: -1}
                                            switch -2147483648 {
                                            case 710:
                                                for fi45 := 0; fi45 < 17; fi45++ {
                                                    fmt.Println(v2, v4)
                                                    if ((^(v4)) == v4) {
                                                        v1 = ""
                                                    }
                                                    _ = fi45
                                                }
                                            case 274:
                                                for fi46 := 0; fi46 < 14; fi46++ {
                                                    for fi47 := 0; fi47 < 10; fi47++ {
                                                        v1 = "a"
                                                        var v48 uint8 = v4
                                                        switch v1 {
                                                        case "foo bar":
                                                            var v49 int = (^(fi46))
                                                            _ = v49
                                                        }
                                                        _ = fi47
                                                        _ = v48
                                                    }
                                                    v50 := GenS2{F0: -1, F1: 677, F2: v1}
                                                    _ = fi46
                                                    _ = v50
                                                }
                                            default:
                                                if v2 {
                                                    v2 = ((42 != 8) && (-1 >= 37))
                                                }
                                            }
                                            _ = v44
                                        }
                                    default:
                                        for fi51 := 0; fi51 < 14; fi51++ {
                                            _ = len(v3)
                                            var v52 bool = v2
                                            v4 -= v4
                                            _ = fi51
                                            _ = v52
                                        }
                                    }
                                    if (v2 && (-24 == 100)) {
                                        var v53 float64 = (0.1 + 0.0)
                                        v1 = "\n\t"
                                        _ = v53
                                    } else {
                                        v54 := []int{32768, 456, -769}
                                        if len(v54) >= 2 {
                                            v54 = v54[1:len(v54)]
                                        }
                                        var v55 uint32 = 906
                                        _ = v54
                                        _ = v55
                                    }
                                    for fi56 := 0; fi56 < 9; fi56++ {
                                        if (127 > 1) {
                                            v4 = 255
                                            v4 = (v4 ^ (v4 >> ((255) & 63)))
                                            var v57 int = fi56
                                            _ = v57
                                        } else {
                                            var v58 int8 = -78
                                            _ = v58
                                        }
                                        if (false && (2 == 0)) {
                                            v59 := GenS6{F0: 757}
                                            _ = v59
                                        }
                                        var v60 float32 = -2.25
                                        _ = fi56
                                        _ = v60
                                    }
                                } else {
                                    var v61 float64 = (0.1 * (1.5 / float64(-2.25)))
                                    v4 = v4
                                    _ = v61
                                }
                            }
                        }
                        fmt.Println(v4, v4, v4)
                        v62 := GenS4{F0: 65535}
                        _ = v62
                    }
                    var v63 int32 = -3
                    v64 := GenS7{F0: 42}
                    _ = v63
                    _ = v64
                }
            case -9223372036854775808:
                if (18446744073709551615 != 309) {
                    v65 := []int{-398, 127, 42, -338}
                    _ = v65
                }
            }
        }
    }
    v1 = v1
    switch -1 {
    case 1:
        v66 := []int{}
        _ = v66
    case 42:
        fmt.Println(v2, v2)
    }
    v67 := GenS7{F0: 699}
    v4 += v4
    var v68 bool = false
    for fi69 := 0; fi69 < 8; fi69++ {
        v4 = v4
        v70 := v3[0]
        for fi71 := 0; fi71 < 9; fi71++ {
            var v72 int32 = 42
            fmt.Println(fi69)
            _ = fi71
            _ = v72
        }
        _ = fi69
        _ = v70
    }
    if (-128 > 127) {
        switch 9223372036854775807 {
        case 42:
            var v73 int8 = 109
            _ = v73
        case 9223372036854775807:
            v4 = v4
        default:
            var v74 int32 = 42
            _ = v74
        }
        if (!v2) {
            v75 := GenS4{F0: 1}
            switch (v1 + v1) {
            case "foo bar":
                v76 := []int{935, 1, 4294967295}
                _ = v76
            default:
                fmt.Println(v2, v4, v68)
            }
            v77 := GenS8{F0: 32767}
            _ = v75
            _ = v77
        }
        v2 = v2
    } else {
        fmt.Println(v2, v2)
        v4 += v4
    }
    v2 = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v67
    _ = v68
}
