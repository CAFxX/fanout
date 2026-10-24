package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 uint8; F2 uint8 }

func main() {
    var v1 float64 = 0.1
    var v2 uint64 = 0
    var v3 uint64 = (v2 - 1)
    var v4 string = ("\n\t" + "hello")
    if false {
        v5 := []int{-9223372036854775808, 32768}
        switch -9223372036854775808 {
        case 240:
            var v6 uint8 = 89
            _ = v6
        default:
            for fi7 := 0; fi7 < 20; fi7++ {
                if false {
                    var v8 float64 = 1e10
                    v2 = (v3 | v3)
                    for fi9 := 0; fi9 < 7; fi9++ {
                        v10 := GenS1{F0: 2147483647, F1: 42, F2: 42}
                        if (!(189 >= 2)) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v2 = 796
                            switch ("hello" + v4) {
                            case "\n\t":
                                v11 := GenS2{F0: 1}
                                _ = v11
                            case "hello":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case "foo bar":
                                _ = len(v5)
                            }
                            if true {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v8 = 1.5
                            }
                            v12 := GenS2{F0: -9223372036854775808}
                            _ = v12
                        } else {
                            fmt.Println(v1, v2)
                            var v13 int64 = 42
                            var v14 int8 = 0
                            _ = v13
                            _ = v14
                        }
                        switch v4 {
                        case "a":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if false {
                                var v15 string = ((v4 + ((v4 + v4) + ((v4 + "foo bar") + "foo bar"))) + "a")
                                _ = v15
                            }
                            switch fi9 {
                            case -1:
                                v8 += v1
                                v8 = v8
                                v1 = v1
                            case 32768:
                                v16 := GenS3{F0: 324, F1: 42, F2: 240}
                                _ = v16
                            case -2147483648:
                                var v17 int8 = int8(fi7)
                                _ = v17
                            }
                        case "":
                            v18 := GenS4{F0: 368}
                            _ = v18
                        }
                        var v19 int = (fi9 / ((fi9) | 1))
                        v4 = v4
                        _ = fi9
                        _ = v10
                        _ = v19
                    }
                    _ = v8
                }
                if false {
                    v1 = v1
                }
                _ = fi7
            }
        }
        for fi20 := 0; fi20 < 1; fi20++ {
            v3 -= (v2 % ((v2) | 1))
            _ = fi20
        }
        _ = v5
    } else {
        v4 = (v4 + v4)
        v2 += (v2 * v2)
        v4 = v4
    }
    switch 435 {
    case 4294967295:
        v21 := GenS3{F0: 42, F1: 506, F2: 255}
        _ = v21
    case 32767:
        for fi22 := 0; fi22 < 3; fi22++ {
            for fi23 := 0; fi23 < 8; fi23++ {
                var v24 uint32 = 65535
                if false {
                    v24 = (v24 << ((18446744073709551615) & 63))
                    v3 -= 0
                    if (1e10 != (v1 + float64(-2.25))) {
                        if false {
                            var v25 float64 = 1e10
                            var v26 int16 = 0
                            _ = v25
                            _ = v26
                        } else {
                            v1 = float64(1e10)
                            var v27 string = v4
                            v28 := GenS2{F0: 851}
                            _ = v27
                            _ = v28
                        }
                        if true {
                            switch (^(fi23)) {
                            case 32768:
                                var v29 int32 = 1
                                _ = v29
                            case 768:
                                v30 := GenS3{F0: 481, F1: v24, F2: 112}
                                _ = v30
                            default:
                                switch fi22 {
                                case 96:
                                    v24 -= v24
                                    var v31 int16 = 0
                                    switch fi23 {
                                    case 9223372036854775807:
                                        if (1 < 18446744073709551615) {
                                            var v32 int64 = 42
                                            v31 -= (v31 * -257)
                                            var v33 int32 = -553
                                            _ = v32
                                            _ = v33
                                        } else {
                                            v34 := GenS5{F0: (v4 + v4), F1: -128, F2: 128}
                                            var v35 int = (fi22 >> ((1) & 63))
                                            _ = v34
                                            _ = v35
                                        }
                                    }
                                    v36 := GenS3{F0: 2147483647, F1: v24, F2: 229}
                                    _ = v31
                                    _ = v36
                                case -28:
                                    if true {
                                        if true {
                                            v4 = ((v4 + ("foo bar" + v4)) + v4)
                                        }
                                        if (!false) {
                                            fmt.Println(v24)
                                        }
                                        switch (fi22 + fi23) {
                                        case -9223372036854775808:
                                            switch fi23 {
                                            case 6:
                                                fmt.Println(v2)
                                            default:
                                                if false {
                                                    var v37 uint64 = v2
                                                    var v38 float64 = v1
                                                    _ = v37
                                                    _ = v38
                                                }
                                            }
                                        case 713:
                                            var v39 uint16 = 42
                                            _ = v39
                                        }
                                    } else {
                                        var v40 uint16 = 1
                                        var v41 int8 = -25
                                        var v42 float32 = float32(0.0)
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                case -187:
                                    v43 := GenS6{F0: 32767}
                                    _ = v43
                                }
                            }
                            v44 := []int{(fi22 + fi23), fi22, (fi22 << ((uint(v2)) & 63)), -693}
                            _ = v44
                        }
                    } else {
                        fmt.Println(fi23, fi22)
                    }
                } else {
                    fmt.Println(v4, v2, v24)
                    var v45 int32 = 42
                    var v46 int = 9223372036854775807
                    v47 := []int{fi23}
                    v24 = 42
                    switch 325 {
                    case 127:
                        v48 := GenS3{F0: v45, F1: v24, F2: 1}
                        _ = v48
                    default:
                        fmt.Println(v3)
                    }
                    _ = v45
                    _ = v46
                    _ = v47
                }
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    default:
        v49 := GenS7{F0: -128, F1: 254, F2: 42}
        _ = v49
    }
    var v50 int8 = 35
    v2 = v3
    var v51 int = 1
    if false {
        v4 = v4
        v51 = (v51 * v51)
        var v52 float32 = 0.1
        _ = v52
    }
    var v53 int8 = -1
    switch (v51 - v51) {
    case 42:
        v3 = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v53
}
